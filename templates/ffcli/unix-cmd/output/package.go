package demomod_gen

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"moul.io/adapterkit/pkg/lib"

	demomod "moul.io/adapterkit-module-demo"
)

func SvcSum(a int64, b int64) (int64, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithInsecure(),
		grpc.WithContextDialer(lib.Dialer(demomod.DemomodSvcServer(&demomod.Service{}), demomod.RegisterDemomodSvcServer)))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demomod.NewDemomodSvcClient(conn)
	req := &demomod.SumRequest{A: a, B: b}
	res, err := client.Sum(ctx, req)
	return res.C, err
}

func SvcSayHello() (string, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithInsecure(),
		grpc.WithContextDialer(lib.Dialer(demomod.DemomodSvcServer(&demomod.Service{}), demomod.RegisterDemomodSvcServer)))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demomod.NewDemomodSvcClient(conn)
	req := &demomod.Empty{}
	res, err := client.SayHello(ctx, req)
	return res.Msg, err
}

func SvcEchoStream(text string) (string, error) {
	panic("not implemented")
}
