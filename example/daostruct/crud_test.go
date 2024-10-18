package gdaodemo

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/base"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/go-logger/logger"
	"testing"
	"time"
)

func init() {
	logger.SetFormatter("{time}>>> {message}\n")
	gdaodemo.RootDir = "../../"
	gdao.Init(gdaodemo.DataSource.PostgrepSql(), gdao.POSTGRESQL)
	gdao.SetLogger(true)
}

func TestSelect(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.IN(5, 4, 3), hs.ID.LT(10))
	h, err := hs.Select()
	logger.Debug(h, err)
}

func TestSelect2(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ROWNAME.RLIKE(1)).GroupBy(hs.ID).Having(hs.ID.Count().LT(2)).Limit(2)
	hslist, _ := hs.Selects()
	for _, h := range hslist {
		logger.Debug(h)
	}
}

func TestSelect3(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.IN(5, 4, 3))
	h, err := hs.Select(base.Col("1"))
	logger.Debug(h, err)
}

func TestUpdate(t *testing.T) {
	hs := dao.NewHstest()
	hs.SetRowname("hello10")
	hs.Where(hs.ID.EQ(10))
	hs.Update()
}

func TestDelete(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.EQ(50))
	hs.Delete()
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
	}
}

func Test_select_setdb(t *testing.T) {
	hs := dao.NewHstest()
	hs.Limit(5)
	list, _ := hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
	gdao.BindDataSource(gdaodemo.DataSource.Mysql(), gdao.MYSQL, "hstest")
	fmt.Println("---------------修改表hstest的数据源:mysql----------------")
	hs = dao.NewHstest()
	hs.Limit(5)
	list, _ = hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
	gdao.BindDataSourceWithClass[dao.Hstest](gdaodemo.DataSource.PostgrepSql(), gdao.POSTGRESQL)
	fmt.Println("---------------修改表hstest的数据源:postgresql----------------")
	hs = dao.NewHstest()
	hs.Limit(5)
	list, _ = hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
	fmt.Println("---------------修改表hstest的数据源:sqlite----------------")
	hs = dao.NewHstest()
	hs.UseDBHandle(gdao.NewDBHandle(gdaodemo.DataSource.Sqlite(), gdao.SQLITE))
	hs.Limit(5)
	list, _ = hs.Selects()
	for _, h := range list {
		fmt.Println(h)
	}
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

	tx.Rollback()
}

// 序列化
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

// 类数据复制
func Test_copy(t *testing.T) {
	hs := dao.NewHstest()
	hs.Where(hs.ID.EQ(1))
	hs1, _ := hs.Select() //任意查询一条数据，作为f复杂的数据准备
	logger.Debug(hs1)
	hs2 := dao.NewHstest()
	hs2.Copy(hs1)
	logger.Debug(hs2)
}

// 结构体绑定数据源
func TestDatasource(t *testing.T) {
	gdao.BindDataSourceWithClass[dao.Hstest1](gdaodemo.DataSource.Mysql(), gdao.MYSQL)
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
