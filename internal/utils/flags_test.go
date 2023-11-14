package utils

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestFlagsWithEnvs(t *testing.T) {
	input := []cli.Flag{
		&cli.BoolFlag{Name: "test_bool"},
		&cli.StringFlag{Name: "test_string"},
		&cli.IntFlag{Name: "test_int"},
		&cli.Int64Flag{Name: "test_int64"},
		&cli.Uint64Flag{Name: "test_uint64"},
		&cli.DurationFlag{Name: "test_duration"},
		&cli.StringSliceFlag{Name: "test_string_slice"},
	}
	excepted := []cli.Flag{
		&cli.BoolFlag{Name: "test_bool", EnvVars: []string{"TEST_BOOL"}},
		&cli.StringFlag{Name: "test_string", EnvVars: []string{"TEST_STRING"}},
		&cli.IntFlag{Name: "test_int", EnvVars: []string{"TEST_INT"}},
		&cli.Int64Flag{Name: "test_int64", EnvVars: []string{"TEST_INT64"}},
		&cli.Uint64Flag{Name: "test_uint64", EnvVars: []string{"TEST_UINT64"}},
		&cli.DurationFlag{Name: "test_duration", EnvVars: []string{"TEST_DURATION"}},
		&cli.StringSliceFlag{Name: "test_string_slice", EnvVars: []string{"TEST_STRING_SLICE"}},
	}
	actual := FlagsWithEnvs(input)

	assert.Equal(t, excepted, actual, "must be equal flag with env")
}

type unknownFlag struct{}

func (u unknownFlag) String() string                { return "" }
func (u unknownFlag) Apply(set *flag.FlagSet) error { return nil }
func (u unknownFlag) Names() []string               { return []string{} }
func (u unknownFlag) IsSet() bool                   { return true }

func TestFlagsWithEnvsThrowPanic(t *testing.T) {
	assert.Panics(t, func() {
		FlagsWithEnvs([]cli.Flag{&unknownFlag{}})
	}, "must be throw panic")
}
