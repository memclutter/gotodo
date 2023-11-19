package utils

import (
	"crypto/tls"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/urfave/cli/v2"
)

func TestNewSqlDriverConnector(t *testing.T) {
	app := cli.NewApp()
	app.Flags = append(app.Flags, &cli.StringFlag{Name: "dbDsn"})
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.String("dbDsn", "postgres://user:secret@localhost:5432/db?sslmode=disable", "Test usage")
	c := cli.NewContext(app, flagSet, nil)
	sqlDriverConnector := NewSqlDriverConnector(c)
	pgdriverConnector := sqlDriverConnector.(*pgdriver.Connector)
	require.Equal(t, "localhost:5432", pgdriverConnector.Config().Addr, "Must be correct address")
	require.Equal(t, "db", pgdriverConnector.Config().Database, "Must be correct database name")
	require.Equal(t, "user", pgdriverConnector.Config().User, "Must be correct user")
	require.Equal(t, "secret", pgdriverConnector.Config().Password, "Must be correct password")
	require.Equal(t, (*tls.Config)(nil), pgdriverConnector.Config().TLSConfig, "Must be correct ssl config")

}
