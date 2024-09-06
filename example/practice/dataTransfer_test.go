package practice

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/base"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/simplelog/logging"
	"reflect"
	"testing"
)

func init() {
	gdaodemo.RootDir = "../../"
	gdao.Init(gdaodemo.DataSource.Sqlite(), gdao.SQLITE)
}

// 保存表数据到指定数据库
func saveByClass[T any](dbhandle base.DBhandle, list []*T) (err error) {
	t := new(T)
	if table, ok := any(t).(gdao.GStruct[*T, T]); ok {
		table.ToGdao()
		table.UseDBHandle(dbhandle)
		for _, v := range list {
			table.Copy(v)
			table.AddBatch()
		}
		_, err = table.ExecBatch()
	} else {
		return fmt.Errorf("error: %v is not a gdao GStruct type", reflect.TypeOf(t).Elem())
	}
	return err
}

// 迁移源数据库表到指定数据库
func transferByClass[T any](dbhandle base.DBhandle) error {
	step := int64(10)
	startindex := int64(0)
	mysqlDBHandle := gdao.NewDBHandle(gdaodemo.DataSource.Mysql(), gdao.MYSQL)
	t := new(T)
	if _, ok := any(t).(gdao.GStruct[*T, T]); !ok {
		return fmt.Errorf("error: %v is not a gdao GStruct type", reflect.TypeOf(t).Elem())
	}
	for {
		t := new(T)
		table, _ := any(t).(gdao.GStruct[*T, T])
		table.ToGdao()
		table.UseDBHandle(mysqlDBHandle)
		table.Limit2(startindex, step)
		var list []*T
		if list, _ = table.Selects(); len(list) > 0 {
			if err := saveByClass[T](dbhandle, list); err != nil {
				logging.Error(err)
			}
		}
		if len(list) < int(step) {
			break
		}
		startindex = startindex + step
	}
	return nil
}

func TestTransferHstest1(t *testing.T) {
	postgreDBHandle := gdao.NewDBHandle(gdaodemo.DataSource.PostgrepSql(), gdao.POSTGRESQL)
	deleteTableData[dao.Hstest1](postgreDBHandle)
	transferByClass[dao.Hstest1](postgreDBHandle)
	fmt.Println("--------------Complete Table Hstest1 data move from mysql to postgreSql---------------")
	fmt.Println("Query hstest1 data records in the postgreSql database")
	showTableData[dao.Hstest1](postgreDBHandle)
}

// delete from table
func deleteTableData[T any](dbhandle base.DBhandle) error {
	t := new(T)
	if table, ok := any(t).(gdao.GStruct[*T, T]); ok {
		table.ToGdao()
		table.UseDBHandle(dbhandle)
		table.Delete()
		fmt.Println("delete all data from table:", table.TableName())
	} else {
		return fmt.Errorf("error: %v is not a gdao GStruct type", reflect.TypeOf(t).Elem())
	}
	return nil
}

// show all data from table
func showTableData[T any](dbhandle base.DBhandle) error {
	t := new(T)
	if table, ok := any(t).(gdao.GStruct[*T, T]); ok {
		table.ToGdao()
		table.UseDBHandle(dbhandle)
		if list, err := table.Selects(); err == nil {
			for _, h := range list {
				fmt.Println(h)
			}
		}
	} else {
		return fmt.Errorf("error: %v is not a gdao GStruct type", reflect.TypeOf(t).Elem())
	}
	return nil
}
