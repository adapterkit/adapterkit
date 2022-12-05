package tictac_gen //nolint

import (
	"context"
	"io"
    "log"

	tictac "github.com/Doozers/adapterkit-module-tictac"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SvcCountdown(count int64, msg string) (chan *tictac.CountdownRes, error) {
  ctx := context.Background()

  conn, err := grpc.DialContext(ctx, "127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatal(err)
  }
  client := tictac.NewTictacSvcClient(conn)
  chann := make (chan *tictac.CountdownRes, 1000)

  req := &tictac.CountdownReq{Count: count, Msg: msg}
  c, err := client.Countdown(ctx, req)
  if err != nil {
    return nil, err
  }

  go func() {
    for {
  	  res, err := c.Recv()
  	  if err == io.EOF {
  		  conn.Close()
  		  close(chann)
  		  break
  	  }
  	  if err != nil {
  		  return
  	  }
  	  chann <- res
    }
  }()

  return chann, nil
}
