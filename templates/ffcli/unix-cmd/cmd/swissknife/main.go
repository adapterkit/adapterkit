package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "os"

    "test/pkg/swissknife_gen"

    "github.com/peterbourgon/ff/v3"
    "github.com/peterbourgon/ff/v3/ffcli"
    swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)

func initSvc() swissknife.SwissknifeSvcServer {
	return swissknife.New()
}

func main() {
    if err := swissknifeRun(os.Args[1:]); err != nil {
        log.Fatal(err)
    }
    return
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
            result, err := swissknife_gen.SvcConvHexa(input, initSvc())
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
            result, err := swissknife_gen.SvcConvBase64(input, initSvc())
            if err != nil {
                return err
            }

            fmt.Println(result)
            return nil
        },
    }
}

