package swissknife_gen

import (
    "context"
    "log"

    "google.golang.org/grpc"
    "moul.io/adapterkit/pkg/lib"

    swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)


func SvcConvHexa(input string, svc swissknife.SwissknifeSvcServer) (string, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithInsecure(),
 		grpc.WithContextDialer(lib.Dialer(svc, swissknife.RegisterSwissknifeSvcServer)))
    if err != nil {
		log.Fatal(err)
	}
    defer conn.Close()
    
    client := swissknife.NewSwissknifeSvcClient(conn)
    req := &swissknife.ConvHexaReq{Input: input}
    res, err := client.ConvHexa(ctx, req)
    return res.Output, err
}


func SvcConvBase64(input string, svc swissknife.SwissknifeSvcServer) (string, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithInsecure(),
 		grpc.WithContextDialer(lib.Dialer(svc, swissknife.RegisterSwissknifeSvcServer)))
    if err != nil {
		log.Fatal(err)
	}
    defer conn.Close()
    
    client := swissknife.NewSwissknifeSvcClient(conn)
    req := &swissknife.ConvBase64Req{Input: input}
    res, err := client.ConvBase64(ctx, req)
    return res.Output, err
}


