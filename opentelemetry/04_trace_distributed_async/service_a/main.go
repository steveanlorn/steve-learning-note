package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	otelSDK "github.com/steveanlorn/steve-learning-note/opentelemetry/04_trace_distributed_async/sdk/otel"
)

const (
	serviceName         = "service-a"
	serviceVersion      = "1.1.0"
	instrumentationName = "github.com/steve-learning-note/opentelemetry/04_trace_distributed_async/service_a/main"
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

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("nats connect: %v", err)
	}
	defer nc.Drain()

	sendHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := sendEvent(ctx, nc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message sent successfully to service B"))
	})

	var mux http.ServeMux
	mux.Handle("/send", sendHandler)

	if err := http.ListenAndServe(":7777",
		otelhttp.NewHandler(&mux, "server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		),
	); err != nil {
		log.Fatal(err)
	}
}

func sendEvent(ctx context.Context, nc *nats.Conn) error {
	ctx, span := tracer.Start(ctx, "sendEvent", trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	const natsSubject string = "demo"

	span.SetAttributes(
		semconv.MessagingSystemKey.String("nats"),
		semconv.MessagingDestinationNameKey.String(natsSubject),
		semconv.MessagingOperationNameKey.String("send"),
	)

	message := []byte("Hello ffrom service A!")

	headers := nats.Header{}
	httpHeader := http.Header(headers)
	carrier := propagation.HeaderCarrier(httpHeader)
	otel.GetTextMapPropagator().Inject(ctx, carrier)

	log.Printf("Sending trace context in headers:%+v", headers)

	msg := &nats.Msg{
		Subject: natsSubject,
		Data:    message,
		Header:  headers,
	}

	if err := nc.PublishMsg(msg); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "could not send event")
		return err
	}

	span.AddEvent("Message published to NATS", trace.WithAttributes(
		semconv.MessagingDestinationNameKey.String(natsSubject),
	))

	return nil
}
