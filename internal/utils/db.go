package utils

import (
	"database/sql"
	"database/sql/driver"

	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/schema"
	"github.com/urfave/cli/v2"
)

func NewBunIDB(bunDB *bun.DB) bun.IDB {
	return bunDB
}

func NewBunDB(sqlDB *sql.DB, schemaDialect schema.Dialect) *bun.DB {
	return bun.NewDB(sqlDB, schemaDialect)
}

func NewSqlDB(sqlDriverConnector driver.Connector) *sql.DB {
	return sql.OpenDB(sqlDriverConnector)
}

func NewSqlDriverConnector(c *cli.Context) driver.Connector {
	return pgdriver.NewConnector(pgdriver.WithDSN(c.String("dbDsn")))
}

func NewSchemaDialect() schema.Dialect {
	return pgdialect.New()
}
