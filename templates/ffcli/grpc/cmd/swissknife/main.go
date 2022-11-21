package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"test/pkg/swissknife_gen"

	"google.golang.org/grpc"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)

func initSvc() swissknife.SwissknifeSvcServer {
	return swissknife.New() // you should need to modify it depending on your service
}

func main() {
	if err := swissknifeRun(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func swissknifeRun(args []string) error {
	rootFlagSet := flag.NewFlagSet("swissknife", flag.ExitOnError)

	root := &ffcli.Command{
		FlagSet:    rootFlagSet,
		ShortUsage: "swissknife [flags] <command> [args...]",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Subcommands: []*ffcli.Command{
			convHexa(),
			convBase64(),
			server(),
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), args)
}

func convHexa() *ffcli.Command {
	var input string
	convHexaFlagSet := flag.NewFlagSet("convHexa", flag.ExitOnError)
	convHexaFlagSet.StringVar(&input, "Input", "", "")
	return &ffcli.Command{
		Name:       "convHexa",
		ShortUsage: "swissknife convHexa ${input}",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		FlagSet:    convHexaFlagSet,
		Exec: func(_ context.Context, _ []string) error {
			result, err := swissknife_gen.SvcConvHexa(input)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}
}

func convBase64() *ffcli.Command {
	var input string
	convBase64FlagSet := flag.NewFlagSet("convBase64", flag.ExitOnError)
	convBase64FlagSet.StringVar(&input, "Input", "", "")
	return &ffcli.Command{
		Name:       "convBase64",
		ShortUsage: "swissknife convBase64 ${input}",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		FlagSet:    convBase64FlagSet,
		Exec: func(_ context.Context, _ []string) error {
			result, err := swissknife_gen.SvcConvBase64(input)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}
}

func server() *ffcli.Command {
	return &ffcli.Command{
		Name:       "server",
		ShortUsage: "swissknife start",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			lis, err := net.Listen("tcp", "127.0.0.1:9314")
			if err != nil {
				return err
			}
			grpcServer := grpc.NewServer()

			swissknife.RegisterSwissknifeSvcServer(grpcServer, initSvc())

			fmt.Println("starting demo-mod_gen server on port 9314:")
			return grpcServer.Serve(lis)
		},
	}
}
