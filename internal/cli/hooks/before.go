package hooks

import (
	"database/sql"
	"fmt"
	"github.com/memclutter/gotodo/internal/config"
	"github.com/memclutter/gotodo/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/urfave/cli/v2"
)

func Before(c *cli.Context) error {
	config.Config.Secret = c.String("secret")
	config.Config.Debug = c.Bool("debug")
	if err := config.Config.SetLogLevel(c.String("logLevel")); err != nil {
		return fmt.Errorf("cli hooks before error: %v", err)
	}
	logrus.SetLevel(config.Config.LogLevel)
	models.DB = bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.String("dsnDb")))), pgdialect.New())
	return nil
}
