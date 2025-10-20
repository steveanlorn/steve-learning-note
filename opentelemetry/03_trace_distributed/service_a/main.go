package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	otelSDK "github.com/steveanlorn/steve-learning-note/opentelemetry/03_trace_distributed/sdk/otel"
	"github.com/steveanlorn/steve-learning-note/opentelemetry/03_trace_distributed/sdk/postgres"
)

const (
	serviceName         = "service-a"
	serviceVersion      = "1.2.0"
	instrumentationName = "github.com/steve-learning-note/opentelemetry/03_trace_distributed/service_a/main"
)

var tracer = otel.GetTracerProvider().Tracer(
	instrumentationName,
	trace.WithInstrumentationVersion(serviceVersion),
	trace.WithSchemaURL(semconv.SchemaURL),
)

func main() {
	ctx := context.Background()

	shutdown, err := otelSDK.InitTracer(ctx, serviceName, serviceVersion)
	if err != nil {
		log.Fatalf("could not install tracer pipeline")
	}
	defer shutdown(context.WithoutCancel(ctx))

	db, err := postgres.Open(ctx, postgres.Config{
		User:     "demo",
		Password: "demo",
		Host:     "localhost:5432",
		DBName:   "demo",
	})
	if err != nil {
		log.Fatal("could not open postgres connection")
	}
	defer db.Close()

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	itemDetailHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemName, err := getItemFromDB(r.Context(), db, r.PathValue("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Call user service
		userName, err := getUserName(r.Context(), client)
		if err != nil {
			http.Error(w, "failed to call user service: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"item": "%s", "user": %s}`, itemName, userName)
	})

	var mux http.ServeMux
	mux.Handle("/item/{name}", itemDetailHandler)

	if err := http.ListenAndServe(":6666",
		otelhttp.NewHandler(&mux, "server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		),
	); err != nil {
		log.Fatal(err)
	}
}

func getItemFromDB(ctx context.Context, db *pgxpool.Pool, name string) (string, error) {
	ctx, span := tracer.Start(ctx, "GetItemFromDB")
	defer span.End()

	var itemName string
	if err := db.QueryRow(ctx, "SELECT @name", pgx.NamedArgs{"name": name}).Scan(&itemName); err != nil {
		return "", err
	}

	return itemName, nil
}

func getUserName(ctx context.Context, httpClient http.Client) (string, error) {
	ctx, span := tracer.Start(ctx, "getUserName")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:7777/user/steve", nil)
	if err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
