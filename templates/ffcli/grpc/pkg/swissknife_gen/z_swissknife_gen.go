package swissknife_gen

import (
    "context"
    "log"

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

    _ "moul.io/adapterkit/pkg/lib"

    swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)

func SvcConvHexa(input string) (*swissknife.ConvHexaRes, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
		log.Fatal(err)
	}
    defer conn.Close()

    client := swissknife.NewSwissknifeSvcClient(conn)
    req := &swissknife.ConvHexaReq{Input: input}
    res, err := client.ConvHexa(ctx, req)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func SvcConvBase64(input string) (*swissknife.ConvBase64Res, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
		log.Fatal(err)
	}
    defer conn.Close()

    client := swissknife.NewSwissknifeSvcClient(conn)
    req := &swissknife.ConvBase64Req{Input: input}
    res, err := client.ConvBase64(ctx, req)
    if err != nil {
        return nil, err
    }

    return res, nil
}
