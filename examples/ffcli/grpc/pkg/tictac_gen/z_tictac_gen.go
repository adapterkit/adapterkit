package tictac_gen //nolint

import (
	"context"
  "io"
  

	tictac "github.com/Doozers/adapterkit-module-tictac"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SvcCountdown(count int64, msg string ,callback func (*tictac.CountdownRes, error) error) error {
  ctx := context.Background()

  conn, err := grpc.DialContext(ctx, "127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
