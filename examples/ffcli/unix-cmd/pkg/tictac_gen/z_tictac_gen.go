package tictac_gen //nolint

import (
	"context"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"moul.io/adapterkit/pkg/lib"

	tictac "github.com/Doozers/adapterkit-module-tictac"
)

func SvcCountdown(count int64, msg string, svc tictac.TictacSvcServer, callback func(*tictac.CountdownRes, error) error) error {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(lib.Dialer(svc, tictac.RegisterTictacSvcServer)))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := tictac.NewTictacSvcClient(conn)

	req := &tictac.CountdownReq{Count: count, Msg: msg}
	c, err := client.Countdown(ctx, req)
	if err != nil {
		return err
	}

	for {
		res, err := c.Recv()
		if err == io.EOF {
			break
		}
		err = callback(res, err)
		if err != nil {
			return err
		}
	}

	return nil
}
