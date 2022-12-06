package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/climan"
	"moul.io/motd"
	"moul.io/srand"
	"moul.io/u"
	"moul.io/zapconfig"

	"github.com/adapterkit/adapterkit/pkg/generate"
	"github.com/peterbourgon/ff/v3"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintf(os.Stderr, "error: %v+\n", err)
		}
		os.Exit(1)
	}
}

var opts struct {
	Debug      bool
	rootLogger *zap.Logger
}

func run(args []string) error {
	// parse CLI
	root := &climan.Command{
		Name:           "adapterkit",
		ShortUsage:     "adapterkit [global flags] <subcommand> [flags] [args]",
		ShortHelp:      "More info on https://moul.io/adapterkit.",
		FlagSetBuilder: func(fs *flag.FlagSet) { fs.BoolVar(&opts.Debug, "debug", opts.Debug, "debug mode") },
		Exec:           doRoot,
		FFOptions:      []ff.Option{ff.WithEnvVarPrefix("adapterkit")},
		Subcommands: []*climan.Command{
			// subcommands
			generate.Cmd(),
		},
		// LongHelp: "",
	}
	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	// init runtime
	{
		// prng
		rand.Seed(srand.Fast())

		// concurrency
		// runtime.GOMAXPROCS(1)

		// logger
		config := zapconfig.New().SetPreset("light-console")
		if opts.Debug {
			config = config.SetLevel(zapcore.DebugLevel)
		} else {
			config = config.SetLevel(zapcore.InfoLevel)
		}
		var err error
		opts.rootLogger, err = config.Build()
		if err != nil {
			return fmt.Errorf("logger init: %w", err)
		}
	}

	// run
	if err := root.Run(context.Background()); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func doRoot(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	opts.rootLogger.Debug("init", zap.Strings("args", args), zap.Any("opts", opts))
	fmt.Print(motd.Default())
	fmt.Println(u.PrettyJSON(args))
	return nil
}
