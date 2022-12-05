package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"test/pkg/tictac_gen"

	tictac "github.com/Doozers/adapterkit-module-tictac"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func initSvc() tictac.TictacSvcServer {
	return tictac.New()
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
			callback := func(res *tictac.CountdownRes, err error) error {
				// you can modify this callback function
				if err != nil {
					return err
				}
				fmt.Println(res)
				return nil
			}
			err := tictac_gen.SvcCountdown(count, msg, initSvc(), callback)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
