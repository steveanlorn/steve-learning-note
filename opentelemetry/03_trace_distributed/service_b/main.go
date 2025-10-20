package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	otelSDK "github.com/steveanlorn/steve-learning-note/opentelemetry/03_trace_distributed/sdk/otel"
	"github.com/steveanlorn/steve-learning-note/opentelemetry/03_trace_distributed/sdk/postgres"
)

const (
	serviceName         = "user-service"
	serviceVersion      = "1.1.0"
	instrumentationName = "github.com/steve-learning-note/opentelemetry/03_trace_distributed/user_service/main"
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

	userHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "GetUserFromDB", trace.WithSpanKind(trace.SpanKindInternal))
		defer span.End()

		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
			}
		}

		var name string
		if err := db.QueryRow(ctx, "SELECT @name", pgx.NamedArgs{"name": r.PathValue("name")}).Scan(&name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]any{"user": name})
	})

	var mux http.ServeMux
	mux.Handle("/user/{name}", userHandler)

	if err := http.ListenAndServe(":7777",
		otelhttp.NewHandler(&mux, "server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		),
	); err != nil {
		log.Fatal(err)
	}
}
