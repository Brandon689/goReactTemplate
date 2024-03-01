package db

import (
	"context"
	"database/sql"
	"github.com/Brandon689/goReactTemplate/types"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Database struct {
	db *bun.DB
}

func NewDatabase() (*Database, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:mydatabase.db?cache=shared")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return &Database{db: db}, nil
}

func (d *Database) Migrate(ctx context.Context) error {
	_, err := d.db.NewCreateTable().Model(&types.User{}).IfNotExists().Exec(ctx)
	return err
}

func CreateSchema(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*types.User)(nil),
		(*types.Profile)(nil),
	}
	for _, model := range models {
		if _, err := db.NewCreateTable().Model(model).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) InsertUsers(ctx context.Context, users []*types.User) error {
	if _, err := d.db.NewInsert().Model(&users).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetUserByName(ctx context.Context, profiles []*types.Profile) error {
	if _, err := d.db.NewInsert().Model(&profiles).Exec(ctx); err != nil {
		return err
	}
	return nil
}
