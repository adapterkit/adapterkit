package lib

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func Dialer[T interface{}](t T, reg func(grpc.ServiceRegistrar, T)) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()

	reg(server, t)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
