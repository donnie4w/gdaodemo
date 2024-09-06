package nativesql

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/go-logger/logger"
	"testing"
)

func init() {
	logger.SetFormatter("{time}>>> {message}\n")
	gdaodemo.RootDir = "../../"
	gdao.Init(gdaodemo.DataSource.Sqlite(), gdao.SQLITE)
	gdao.SetLogger(true)
}

func Test_qurey(t *testing.T) {
	hs, err := gdao.ExecuteQuery[dao.Hstest]("select * from hstest where id=?", 12)
	logger.Debug(err)
	logger.Debug(hs)
}

func Test_qureylist(t *testing.T) {
	hsList, err := gdao.ExecuteQueryList[dao.Hstest]("select * from hstest where id in (?,?,?,?)", 12, 13, 14, 15)
	logger.Debug(err)
	for _, hs := range hsList {
		logger.Debug(hs)
	}
}

func Test_qureybean(t *testing.T) {
	bean := gdao.ExecuteQueryBean("select * from hstest where id=?", 12)
	logger.Debug(bean.GetError())
	logger.Debug(bean)
}

func Test_qureybean2(t *testing.T) {
	var Hs dao.Hstest
	err := gdao.ExecuteQueryBean("select * from hstest where id=?", 12).Scan(&Hs)
	logger.Debug(err, &Hs)
}

func Test_qureybeans(t *testing.T) {
	beans := gdao.ExecuteQueryBeans("select * from hstest where id in (?,?,?,?)", 12, 13, 14, 15)
	for _, bean := range beans.Beans {
		logger.Debug(bean)
	}
}

func Test_qureybeans2(t *testing.T) {
	var Hslist []*dao.Hstest
	err := gdao.ExecuteQueryBeans("select * from hstest where id in (?,?,?,?)", 12, 13, 14, 15).Scan(&Hslist)
	logger.Debug(err, len(Hslist))
	for _, hst := range Hslist {
		logger.Debug(hst)
	}
}

func Test_update(t *testing.T) {
	gdao.ExecuteUpdate("update hstest set id=? where id=?", 12)
}

//存储过程示例
/*
DELIMITER $$
USE `hstest`$$
DROP PROCEDURE IF EXISTS `proc_hs`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_hs`(

	IN in_param INT         -- 输入参数

)
BEGIN
SELECT *  FROM hstest WHERE id < in_param;
END$$
DELIMITER ;
*/
func TestCall(t *testing.T) {
	dbhandle := gdao.NewDBHandle(gdaodemo.DataSource.Mysql(), gdao.MYSQL)
	dbs := dbhandle.ExecuteQueryBeans("call proc_hs(?)", 3)
	fmt.Println(dbs.GetError())
	for _, db := range dbs.Beans {
		fmt.Println(db)
	}
}

// 存储过程示例
func TestCall2(t *testing.T) {
	dbhandle := gdao.NewDBHandle(gdaodemo.DataSource.Mysql(), gdao.MYSQL)
	var Outparam string
	var Inoutparam = 3
	i, err := dbhandle.ExecuteUpdate("call proc_test(?,?,?)", 1, &Outparam, &Inoutparam)
	fmt.Println(i, err)
}

// 批处理示例
func TestBatch(t *testing.T) {
	gdao.ExecuteBatch("insert into hstest1 (Rowname,Value) values (?,?)", [][]any{{"row1", "hello1"}, {"row2", "hello2"}})
	fmt.Println(gdao.ExecuteQueryBeans("select * from hstest1 order by id desc limit 2"))
}

// 事务
func Test_transaction(t *testing.T) {
	//获取事务对象
	tx, _ := gdao.NewTransaction()
	//事务对象可以直接调用CRUD函数
	tx.ExecuteUpdate("update hstest set age=? where id=?", 100, 1)
	tx.ExecuteUpdate("update hstest set age=? where id=?", 100, 2)

	//事务回滚
	tx.Rollback()
}
