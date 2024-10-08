package gdaodemo

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoCache"
	"github.com/donnie4w/gdaodemo/dao"
	"testing"
	"time"
)

func init() {
	gdao.Init(DataSource.Sqlite(), gdao.SQLITE)
	gdao.SetLogger(true)
	logger.Info("datasource init")
}

func TestSelect(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.EQ(3))
	h, err := hs.Select(hs.ID, hs.VALUE, hs.ROWNAME)
	logger.Debug(h, err)
	bean := gdao.ExecuteQueryBean("select id,value,rowname from hstest where id=?", 11)
	logger.Debug(bean)
}

func TestSelect2(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ROWNAME.RLIKE(1)).GroupBy(hs.ID).Having(hs.ID.Count().LT(2)).Limit(2)
	hslist, _ := hs.Selects()
	for _, h := range hslist {
		logger.Debug(h)
	}
	bean := gdao.ExecuteQueryBean("select id,value,rowname from hstest where id=?", 11)
	logger.Debug(bean)
}

func TestUpdate(t *testing.T) {
	hs := dao.NewHstest()
	hs.SetRowname("hello10")
	hs.Where(hs.ID.EQ(10))
	hs.Update()
	bean := gdao.ExecuteQueryBean("select id,value,rowname from hstest where id=?", 13)
	logger.Debug(bean)
}

func TestDelete(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.EQ(50))
	hs.Update()
}

func TestInsert(t *testing.T) {
	hs := dao.NewHstest()
	hs.SetValue("hello123")
	hs.SetLevel(12345)
	hs.SetBody([]byte("hello"))
	hs.SetRowname("hello1234")
	hs.SetUpdatetime(time.Now())
	hs.SetFloa(123456)
	hs.SetAge(123)
	hs.Insert()
	bean := gdao.ExecuteQueryBean("select * from hstest order by id desc limit ?", 1)
	logger.Debug(bean)
}

func TestSelects(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where((hs.ID.Between(0, 5)).Or(hs.ID.Between(10, 60)))
	hs.GroupBy(hs.ID)
	hs.Having(hs.ID.Count().LE(2))
	hs.OrderBy(hs.ID.Asc())
	hs.Limit2(1, 10)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	} else {
		logger.Error("err>>>>", err)
	}
}

func TestSelects2(t *testing.T) {
	hs := dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(5)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	} else {
		logger.Debug("err>>>>", err)
	}
	logger.Debug("---------------------Gdao Hstest1---------------------")
	if beans := gdao.ExecuteQueryBeans("select id,rowname,value,goto from hstest1 where id between ? and ? order by id desc  LIMIT ? OFFSET ?", 0, 5, 10, 1); beans.GetError() == nil {
		for _, bean := range beans.Beans {
			logger.Debug(bean)
		}
	}
	logger.Debug("---------------------DataBean---------------------")
}

func TestSelectScan(t *testing.T) {
	var hslist []*dao.Hs1
	if beans := gdao.ExecuteQueryBeans("select id,rowname,value,goto from hstest1 where id between ? and ? order by id desc  LIMIT ? OFFSET ?", 0, 5, 10, 1); beans.GetError() == nil {
		fmt.Println(beans.Scan(&hslist))
	}
	logger.Debug("---------------------DataBean---------------------")

	for _, h1 := range hslist {
		logger.Debug(h1)
	}
	logger.Debug("---------------------scan []*dao.Hs1---------------------")
}

func TestSelectScan2(t *testing.T) {
	hslist := make([]*dao.Hs1, 0)
	if beans := gdao.ExecuteQueryBeans("select * from hstest where id between ? and ? order by id desc  LIMIT ? OFFSET ?", 0, 5, 10, 1); beans.GetError() == nil {
		fmt.Println(beans.Scan(&hslist))
	}
	logger.Debug("---------------------DataBean---------------------")

	for _, h1 := range hslist {
		logger.Debug(h1)
		logger.Debug(h1.GetUpdatetime().Format("2006-01-02 15:04:05"))
	}
	logger.Debug("---------------------scan []*dao.Hs1---------------------")

}

func Test_select_setdb(t *testing.T) {
	hs := dao.NewHstest()
	hs.Limit(5)
	list, _ := hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
	gdao.BindDataSource(DataSource.Mysql(), gdao.MYSQL, "hstest")
	fmt.Println("---------------修改表hstest的数据源:mysql----------------")
	hs = dao.NewHstest()
	hs.Limit(5)
	list, _ = hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
	gdao.BindDataSourceWithClass[dao.Hstest](DataSource.PostgrepSql(), gdao.POSTGRESQL)
	fmt.Println("---------------修改表hstest的数据源:postgresql----------------")
	hs = dao.NewHstest()
	hs.Limit(5)
	list, _ = hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
	fmt.Println("---------------修改表hstest的数据源:sqlite----------------")
	hs = dao.NewHstest()
	hs.UseDBHandle(gdao.NewDBHandle(DataSource.Sqlite(), gdao.SQLITE))
	hs.Limit(5)
	list, _ = hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
}

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
	dbhandle := gdao.NewDBHandle(DataSource.Mysql(), gdao.MYSQL)
	dbs := dbhandle.ExecuteQueryBeans("call proc_hs(?)", 3)
	fmt.Println(dbs.GetError())
	for _, db := range dbs.Beans {
		fmt.Println(db)
	}
}

func TestCall2(t *testing.T) {
	dbhandle := gdao.NewDBHandle(DataSource.Mysql(), gdao.MYSQL)
	var Outparam string
	var Inoutparam = 3
	i, err := dbhandle.ExecuteUpdate("call proc_test(?,?,?)", 1, &Outparam, &Inoutparam)
	fmt.Println(i, err)
}

func TestBatch(t *testing.T) {
	hs := dao.NewHstest2()
	hs.SetAge(100)
	hs.SetName("www")
	hs.SetCreatetime(time.Now())
	hs.SetFloa(1.1)
	hs.AddBatch()

	hs.SetAge(1000)
	hs.SetName("wwww")
	hs.SetCreatetime(time.Now())
	hs.SetFloa(1.11)
	hs.AddBatch()
	hs.ExecBatch()

	hs2 := dao.NewHstest2()
	if hslist, err := hs2.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
}

func TestBatch2(t *testing.T) {
	gdao.ExecuteBatch("insert into hstest1 (Rowname,Value) values (?,?)", [][]any{{"row1", "hello1"}, {"row2", "hello2"}})
	fmt.Println(gdao.ExecuteQueryBeans("select * from hstest1 order by id desc limit 2"))
}

func Test_transaction(t *testing.T) {
	//获取事务对象
	tx, _ := gdao.NewTransaction()
	hs := dao.NewHstest2()
	//传入事务对象tx
	hs.UseTransaction(tx)
	hs.SetAge(100)
	hs.SetName("www123")
	hs.Where(hs.ID.EQ(1))
	hs.Update()

	hs2 := dao.NewHstest2()
	//传入事务对象tx
	hs2.UseTransaction(tx)
	hs2.SetAge(101)
	hs2.SetName("www234")
	hs2.Where(hs.ID.EQ(2))
	hs2.Update()

	//事务对象可以直接调用CRUD函数
	tx.ExecuteUpdate("update hstest set age=? where id=?", 100, 1)

	tx.Rollback()
	logger.Debug("-------------Transaction  Rollback-----------------")
	//检查是否回滚成功
	fmt.Println(gdao.ExecuteQueryBean("select id,age,name from hstest2 where id=?", 1))
	fmt.Println(gdao.ExecuteQueryBean("select id,age,name from hstest2 where id=?", 2))
	fmt.Println(gdao.ExecuteQueryBean("select id,age from hstest where id=?", 2))
}

func TestGdao(t *testing.T) {
	if datas := gdao.ExecuteQueryBeans("select * from hstest where id  < ?", 5); datas.GetError() == nil {
		for _, data := range datas.Beans {
			logger.Debug(data)
		}
	} else {
		logger.Debug("err>>>", datas.GetError())
	}
}

func Test_Scan(t *testing.T) {
	var hss []*dao.Hstest1
	if err := gdao.ExecuteQueryBeans("select * from hstest1 limit ?", 5).Scan(&hss); err == nil {
		for _, hs := range hss {
			fmt.Println(hs)
		}
	}
}

func Test_Scan2(t *testing.T) {
	type Hs struct {
		Id         int64
		Name       string
		Age        int64
		Createtime time.Time
		Money      float64
		Bytes      []byte
		Floa       float64
		Level      int64
		Type       float64
		Flog       float64
	}
	hslist, err := gdao.ExecuteQueryList[Hs]("select * from hstest2 limit ?,?", 10, 5)
	if err != nil {
		logger.Debug(err.Error())
	}
	for _, hs := range hslist {
		logger.Debug(hs)
	}
}

func Test_Scan3(t *testing.T) {
	type Hs struct {
		Id         int64
		Name       string
		Age        int64
		Createtime time.Time
		Money      float64
		Bytes      []byte
		Floa       float64
		Level      int64
		Type       float64
	}
	if gbs := gdao.ExecuteQueryBeans("select * from hstest2 limit ?,?", 10, 5); gbs.GetError() == nil {
		for _, dataBean := range gbs.Beans {
			var hs dao.Hs1
			dataBean.Scan(&hs)
			logger.Debug(&hs)
		}
	}
}

func Test_serialize(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.EQ(1))
	hs1, _ := hs.Select() //任意查询一条数据，作为序列化的数据准备
	bs, _ := hs1.Encode() //调用Encode 实现对象序列化
	logger.Debug("encode len(bs):", len(bs))
	logger.Debug(hs1)
	logger.Debug("----------Encode-----------")
	hs2 := dao.NewHstest()
	hs2.Decode(bs) //调用Decode 实现数据反序列化，并赋值给hs2
	logger.Debug(hs2)
	logger.Debug("----------Decode-----------")
}

func Test_copy(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.EQ(1))
	hs1, _ := hs.Select() //任意查询一条数据，作为f复杂的数据准备
	logger.Debug(hs1)
	hs2 := dao.NewHstest()
	hs2.Copy(hs1)
	logger.Debug(hs2)
}

func TestCacheTablename(t *testing.T) {
	gdaoCache.BindTableNames("hstest") //set cache for Hstest
	hs := dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(1)
	if hs, err := hs.Select(); err == nil {
		logger.Debug(hs)
	}
	logger.Debug("----------------------Set Cache----------------------")
	println()
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(1)
	if hs, err := hs.Select(); err == nil {
		logger.Debug(hs)
	}
	logger.Debug("----------------------Get Cache----------------------")
	println()
	gdaoCache.UnbindTableNames("hstest")
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(1)
	if hs, err := hs.Select(); err == nil {
		logger.Debug(hs)
	}
	logger.Debug("----------------------No Use Cache----------------------")
}

func TestCacheClass(t *testing.T) {
	gdaoCache.BindClassWithCacheHandle[*dao.Hstest](gdaoCache.NewCacheHandle().SetExpire(100).SetStoreMode(gdaoCache.STRONG)) //set cache for Hstest
	hs := dao.NewHstest()

	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Set Cache----------------------")
	println()
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Get Cache----------------------")
	println()
	//gdaoCache.UnbindClass[dao.Hstest]()
	gdaoCache.ClearClass[dao.Hstest]()
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------No Use Cache----------------------")
}

func TestDatasource(t *testing.T) {
	gdao.BindDataSourceWithClass[dao.Hstest1](DataSource.Mysql(), gdao.MYSQL)
	hs := dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	//unbind  data source
	gdao.UnbindDataSourceWithClass[dao.Hstest1]()
	hs = dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
}

func TestDatasource2(t *testing.T) {
	gdao.BindDataSource(DataSource.Mysql(), gdao.MYSQL, "hstest1")
	hs := dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	//unbind  data source
	gdao.UnbindDataSource("hstest1")
	hs = dao.NewHstest1()
	hs.UseDBHandle(gdao.NewDBHandle(DataSource.Mysql(), gdao.MYSQL))
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
}
