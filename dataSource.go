package main

import "database/sql"

type dataSource struct {
}

var DataSource *dataSource = &dataSource{}

func (d *dataSource) Mysql() *sql.DB {
	if db, err := getDataSource("mysql.json"); err == nil {
		return db
	} else {
		panic(err)
	}
}

func (d *dataSource) Sqlite() *sql.DB {
	if db, err := getDataSource("sqlite.json"); err == nil {
		return db
	} else {
		panic(err)
	}
}

func (d *dataSource) PostgrepSql() *sql.DB {
	if db, err := getDataSource("postgre.json"); err == nil {
		return db
	} else {
		panic(err)
	}
}
