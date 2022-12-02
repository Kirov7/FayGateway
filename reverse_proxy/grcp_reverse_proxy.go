package reverse_proxy

import (
	"context"
	"github.com/Kirov7/FayGateway/grpc_proxy"
	"github.com/Kirov7/FayGateway/reverse_proxy/load_balance"
	"google.golang.org/grpc"
	"log"
)

func NewGrpcLoadBalanceHandler(lb load_balance.LoadBalance) grpc.StreamHandler {
	return func() grpc.StreamHandler {
		nextAddr, err := lb.Get("")
		if err != nil {
			log.Fatal("get next addr fail")
		}
		director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
			c, err := grpc.DialContext(ctx, nextAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
			return ctx, c, err
		}
		return proxy.TransparentHandler(director)
	}()
}
