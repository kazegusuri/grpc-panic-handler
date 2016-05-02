package panichandler

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ grpc.UnaryServerInterceptor = UnaryPanicHandler
var _ grpc.StreamServerInterceptor = StreamPanicHandler

func toPanicError(r interface{}) error {
	return grpc.Errorf(codes.Internal, "panic: %v", r)
}

func UnaryPanicHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer handleCrash(func(r interface{}) {
		err = toPanicError(r)
	})

	return handler(ctx, req)
}

func StreamPanicHandler(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer handleCrash(func(r interface{}) {
		err = toPanicError(r)
	})

	return handler(srv, stream)
}

var additionalHandlers []func(interface{})

func InstallPanicHandler(handler func(interface{})) {
	additionalHandlers = append(additionalHandlers, handler)
}

func handleCrash(handler func(interface{})) {
	if r := recover(); r != nil {
		handler(r)

		if additionalHandlers != nil {
			for _, fn := range additionalHandlers {
				fn(r)
			}
		}
	}
}
