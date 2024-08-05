package main

import (
	"fmt"
	"github.com/donnie4w/gdao/gdaoBuilder"
	"testing"
)

// gdaoBuidler test
func Test_build(t *testing.T) {
	db := DataSource.Sqlite()
	// tablename:hstest1
	// dbtype : mysql
	// dbname : hsetst
	// package : dao
	if err := gdaoBuilder.Build("hstest1", "mysql", "hstest", "dao", db); err != nil {
		fmt.Println(err)
	}
}

// gdaoBuidler test
func Test_buildDir(t *testing.T) {
	db := DataSource.Sqlite()
	// dir: Path for storing the generated file.
	// tablename: hstest1
	// dbtype : mysql
	// dbname : hsetst
	// package : dao
	if err := gdaoBuilder.BuildDir("ddd", "hstest1", "mysql", "hstest", "dao", db); err != nil {
		fmt.Println(err)
	}
}

// gdaoBuidler test
func Test_buildDirWithTables(t *testing.T) {
	db := DataSource.Sqlite()
	tables := []string{"hstest", "hstest1", "hstest2", "hstest3"}

	for _, table := range tables {
		if err := gdaoBuilder.BuildDir("ddd", table, "mysql", "hstest", "dao", db); err != nil {
			fmt.Println(err)
		}
	}

}
