package main

import (
	"database/sql"
	"fmt"
	goutil "github.com/donnie4w/gofer/util"
	"github.com/donnie4w/simplelog/logging"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

var logger = logging.NewLogger().SetFormat(logging.FORMAT_LEVELFLAG | logging.FORMAT_TIME).SetFormatter("{time}{level}{message}\n")

type ConfBean struct {
	DbType string `json:"dbtype"`
	DbHost string `json:"dbhost"`
	DbPort int    `json:"dbport"`
	DbName string `json:"dbname"`
	DbUser string `json:"dbuser"`
	DbPwd  string `json:"dbpwd"`
}

func getDataSource(path string) (db *sql.DB, err error) {
	var bs []byte
	if bs, err = goutil.ReadFile(path); err == nil {
		var config *ConfBean
		if config, err = goutil.JsonDecode[*ConfBean](bs); err == nil {
			return openDB(config.DbType, config)
		}
	}
	return
}

func openDB(driver string, config *ConfBean) (DB *sql.DB, err error) {
	dataSourceName := ""
	switch strings.ToLower(driver) {
	case "mysql", "mariadb":
		dataSourceName = config.DbUser + ":" + config.DbPwd + "@tcp(" + config.DbHost + ":" + fmt.Sprint(config.DbPort) + ")/" + config.DbName
		DB, err = sql.Open("mysql", dataSourceName)
	case "postgresql":
		dataSourceName = "host=" + config.DbHost + " port=" + fmt.Sprint(config.DbPort) + " user=" + config.DbUser + " password=" + config.DbPwd + " dbname=" + config.DbName + " sslmode=disable"
		DB, err = sql.Open("postgres", dataSourceName)
	case "sqlite":
		dataSourceName = config.DbName
		DB, err = sql.Open("sqlite3", dataSourceName)
	default:
		err = fmt.Errorf("Unsupported driver: %s", driver)
	}
	logger.Info("connect to ", driver, "[", dataSourceName, "]")
	if err == nil {
		if err = DB.Ping(); err != nil {
			logger.Error("Ping DB failed:", err)
		}
	}
	return
}
