package swissknife_gen //nolint

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"moul.io/adapterkit/pkg/lib"

	swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)

func SvcConvHexa(input string, svc swissknife.SwissknifeSvcServer) (*swissknife.ConvHexaRes, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(lib.Dialer(svc, swissknife.RegisterSwissknifeSvcServer)))
	if err != nil {
		return nil, err
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

func SvcConvBase64(input string, svc swissknife.SwissknifeSvcServer) (*swissknife.ConvBase64Res, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(lib.Dialer(svc, swissknife.RegisterSwissknifeSvcServer)))
	if err != nil {
		return nil, err
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
