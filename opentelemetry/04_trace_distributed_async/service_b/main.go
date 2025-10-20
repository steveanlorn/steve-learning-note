package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	otelSDK "github.com/steveanlorn/steve-learning-note/opentelemetry/04_trace_distributed_async/sdk/otel"
)

const (
	serviceName         = "service-b"
	serviceVersion      = "1.2.0"
	instrumentationName = "github.com/steve-learning-note/opentelemetry/04_trace_distributed_async/service_b/main"
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

	log.Println("run the subscriber")

	_, err = nc.Subscribe("demo", func(msg *nats.Msg) {
		ctx := context.Background()

		log.Printf("Received trace context in headers: %+v", msg.Header)

		link, err := createSpanLinkFromHeaders(msg.Header)
		if err != nil {
			log.Printf("failed to create span link from headers: %v", err)
		}

		_, span := tracer.Start(ctx, "handle-nats-message", trace.WithLinks(link), trace.WithSpanKind(trace.SpanKindConsumer))
		defer span.End()

		span.SetAttributes(
			semconv.MessagingSystemKey.String("nats"),
			semconv.MessagingDestinationNameKey.String(msg.Subject),
			semconv.MessagingOperationTypeReceive,
		)

		time.Sleep(100 * time.Millisecond)

		msg.Ack()
		span.AddEvent("ack", trace.WithAttributes(
			semconv.MessagingOperationNameKey.String("ack"),
		))
	})
	if err != nil {
		log.Fatalf("subscribe error: %v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("shutting down subscriber...")
}

func createSpanLinkFromHeaders(headers nats.Header) (trace.Link, error) {
	httpHeader := http.Header(headers)
	carrier := propagation.HeaderCarrier(httpHeader)

	log.Println(httpHeader)
	log.Println(carrier.Keys())
	log.Println(httpHeader.Get("traceparent"))
	log.Println(headers.Get("traceparent"))

	traceparent := headers.Get("traceparent")

	var link trace.Link
	if traceparent != "" {
		parts := strings.Split(traceparent, "-")
		if len(parts) >= 4 {
			traceIDStr := parts[1]
			spanIDStr := parts[2]
			flagStr := parts[3]

			traceID, err := trace.TraceIDFromHex(traceIDStr)
			if err != nil {
				return link, err
			}

			spanID, err := trace.SpanIDFromHex(spanIDStr)
			if err != nil {
				return link, err
			}

			var traceFlags trace.TraceFlags
			if flags, err := strconv.ParseUint(flagStr, 16, 8); err != nil {
				return link, err
			} else {
				traceFlags = trace.TraceFlags(flags)
			}

			spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
				TraceID:    traceID,
				SpanID:     spanID,
				TraceFlags: traceFlags,
				Remote:     true,
			})

			link = trace.Link{
				SpanContext: spanCtx,
				Attributes: []attribute.KeyValue{
					semconv.MessagingSystemKey.String("nats"),
				},
			}

			log.Println("link is generated")
		} else {
			log.Println("less than 4")
		}
	} else {
		log.Println("empty traceparent")
	}

	return link, nil
}
