package example

import (
	"context"
	"github.com/joshcarp/grpcexample/proto/examplepb"
)

type FooServer struct {
	examplepb.UnimplementedFooAPIServer
}

func (f FooServer) Hello(ctx context.Context, example *examplepb.ExampleRequest) (*examplepb.ExampleResponse, error) {
	return &examplepb.ExampleResponse{
		Message: "Hello " + example.Message,
	}, nil
}

type BarServer struct{
	examplepb.UnimplementedBarAPIServer
}

func (f BarServer) ListBars(ctx context.Context, example *examplepb.BarRequest) (*examplepb.BarResponse, error) {
	return &examplepb.BarResponse{
		Message: "Barserver " + example.Message,
	}, nil
}
