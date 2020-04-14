package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type metadataReaderWriter struct {
	metadata.MD
}

// OpenTracingClientInterceptor 实现GRPC client一元拦截器
func OpenTracingClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var parentCtx opentracing.SpanContext
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}

		clientSpan := tracer.StartSpan(method, opentracing.ChildOf(parentCtx), opentracing.Tags{}, ext.SpanKindRPCClient)
		defer clientSpan.Finish()

		// 将之前放入context中的metadata数据出去 如果没有则新建一个
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		if err := tracer.Inject(clientSpan.Context(), opentracing.TextMap, metadataReaderWriter{md}); err != nil {
			grpclog.Errorf("inject to metadata err: %v", err)
		}

		// md -> ctx
		ctx = metadata.NewOutgoingContext(ctx, md)
		if err := invoker(ctx, method, req, reply, cc, opts...); err != nil {
			clientSpan.LogFields(log.String("err", err.Error()))
			return err
		}

		return nil
	}
}

func OpentracingServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		spanContext, err := tracer.Extract(opentracing.TextMap, metadataReaderWriter{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		}

		serverSpan := tracer.StartSpan(info.FullMethod, ext.RPCServerOption(spanContext), opentracing.Tags{}, ext.SpanKindRPCServer) // tag
		defer serverSpan.Finish()

		ctx = opentracing.ContextWithSpan(ctx, serverSpan)

		return handler(ctx, req)
	}
}
