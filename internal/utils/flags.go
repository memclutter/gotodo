package utils

import (
	"log"
	"strings"

	"github.com/urfave/cli/v2"
)

func FlagsWithEnvs(flags []cli.Flag) []cli.Flag {
	for _, flag := range flags {
		switch typedFlag := flag.(type) {
		case *cli.BoolFlag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		case *cli.StringFlag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		case *cli.IntFlag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		case *cli.Int64Flag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		case *cli.Uint64Flag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		case *cli.DurationFlag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		case *cli.StringSliceFlag:
			typedFlag.EnvVars = []string{strings.ToUpper(typedFlag.Name)}
		default:
			log.Panicf("unknown flag type %T", typedFlag)
		}
	}
	return flags
}
