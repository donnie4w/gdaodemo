package gdaodemo

import (
	"database/sql"
	"path/filepath"
)

type dataSource struct {
}

var DataSource = &dataSource{}

func (d *dataSource) Mysql() *sql.DB {
	cfg := "mysql.json"
	if RootDir != "" {
		cfg = filepath.Join(RootDir, cfg)
	}
	if db, err := getDataSource(cfg); err == nil {
		return db
	} else {
		panic(err)
	}
}

func (d *dataSource) Sqlite() *sql.DB {
	cfg := "sqlite.json"
	if RootDir != "" {
		cfg = filepath.Join(RootDir, cfg)
	}
	if db, err := getDataSource(cfg); err == nil {
		return db
	} else {
		panic(err)
	}
}

func (d *dataSource) PostgrepSql() *sql.DB {

	cfg := "postgres.json"
	if RootDir != "" {
		cfg = filepath.Join(RootDir, cfg)
	}
	if db, err := getDataSource(cfg); err == nil {
		return db
	} else {
		panic(err)
	}
}
