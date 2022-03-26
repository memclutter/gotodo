package hooks

import (
	"database/sql"
	"github.com/memclutter/gotodo/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/urfave/cli/v2"
)

func Before(c *cli.Context) error {
	models.DB = bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.String("dsnDb")))), pgdialect.New())
	return nil
}
