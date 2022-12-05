package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"test/pkg/tictac_gen"

	"google.golang.org/grpc"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	tictac "github.com/Doozers/adapterkit-module-tictac"
)

func initSvc() tictac.TictacSvcServer {
	return tictac.New() // you should need to modify it depending on your service
}

func main() {
	if err := tictacRun(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func tictacRun(args []string) error {
	rootFlagSet := flag.NewFlagSet("tictac", flag.ExitOnError)

	root := &ffcli.Command{
		FlagSet:    rootFlagSet,
		ShortUsage: "tictac [flags] <command> [args...]",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Subcommands: []*ffcli.Command{
			countdown(),
			server(),
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), args)
}

func countdown() *ffcli.Command {
	var count int64
    var msg string
	countdownFlagSet := flag.NewFlagSet("countdown", flag.ExitOnError)
	countdownFlagSet.Int64Var(&count, "Count", 0, "")
    countdownFlagSet.StringVar(&msg, "Msg", "", "")
	return &ffcli.Command{
		Name:       "countdown",
		ShortUsage: "tictac countdown ${input}",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		FlagSet:    countdownFlagSet,
		Exec: func(_ context.Context, _ []string) error {
        result, err := tictac_gen.SvcCountdown(count, msg)
        if err != nil {
          return err
			  }

        for {
        	res, ok := <-result
        	if !ok {
        		break
        	}

        	fmt.Println(res)
        }

        return nil
      },
    }
  }

func server() *ffcli.Command {
	return &ffcli.Command{
		Name:       "server",
		ShortUsage: "tictac start",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			lis, err := net.Listen("tcp", "127.0.0.1:9314")
			if err != nil {
				return err
			}
			grpcServer := grpc.NewServer()

			tictac.RegisterTictacSvcServer(grpcServer, initSvc())

			fmt.Println("starting demo-mod_gen server on port 9314:")
			return grpcServer.Serve(lis)
		},
	}
}
