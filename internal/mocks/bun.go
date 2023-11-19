package mocks

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type BunDB struct {
	mock.Mock
}

func (m *BunDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) Dialect() schema.Dialect {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewValues(model interface{}) *bun.ValuesQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewSelect() *bun.SelectQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewInsert() *bun.InsertQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewUpdate() *bun.UpdateQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewDelete() *bun.DeleteQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewRaw(query string, args ...interface{}) *bun.RawQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewCreateTable() *bun.CreateTableQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewDropTable() *bun.DropTableQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewCreateIndex() *bun.CreateIndexQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewDropIndex() *bun.DropIndexQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewTruncateTable() *bun.TruncateTableQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewAddColumn() *bun.AddColumnQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) NewDropColumn() *bun.DropColumnQuery {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (bun.Tx, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BunDB) RunInTx(ctx context.Context, opts *sql.TxOptions, f func(ctx context.Context, tx bun.Tx) error) error {
	//TODO implement me
	panic("implement me")
}
