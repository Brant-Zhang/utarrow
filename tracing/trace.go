package tracing

import (
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/stats"
)

var Gtracer *zipkin.Tracer

func NewZipkin(url string, servicename string, local string) *zipkin.Tracer {
	reporter := reco.NewReporter(url)
	ep, err := zipkin.NewEndpoint(servicename, local)
	if err != nil {
		return nil
	}
	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(ep))
	if err != nil {
		return nil
	}
	Gtracer = tracer
	return tracer
}

func NewTraceServer(tags map[string]string) *grpc.Server {
	h := zipkinrpc.NewServerHandler(Gtracer, zipkinrpc.ServerTags(tags))
	op := grpc.StatsHandler(h)
	s := grpc.NewServer(op)
	return s
}

func NewClientHandler() grpc.DialOption {
	return grpc.WithStatsHandler(zipkinrpc.NewClientHandler(Gtracer))
}
