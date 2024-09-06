package gdaodemo

import (
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoSlave"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/go-logger/logger"
	"testing"
)

func TestSlave(t *testing.T) {
	gdaoSlave.BindClass[dao.Hstest1](gdaodemo.DataSource.Mysql(), gdao.MYSQL) //绑定Hstest1的读操作数据源
	hs := dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Bind the hstest table to the slave data source mysql----------------------")
	println()

	gdaoSlave.UnbindClass[dao.Hstest1]() //解除绑定Hstest1的读操作数据源

	hs = dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------gdaoslave unbinds the hstest table----------------------")
	println()
}

func TestSlaveTable(t *testing.T) {
	gdaoSlave.BindTable(gdaodemo.DataSource.Mysql(), gdao.MYSQL, "hstest1", "hstest2") //绑定Hstest1，hstest2 的读操作数据源
	hs := dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Bind the hstest1, hstest2 table to the slave data source mysql----------------------")
	println()
	gdaoSlave.UnbindTable("hstest1", "hstest2") //解除绑定Hstest1,Hstest2 的读操作数据源
	hs = dao.NewHstest1()
	hs.Where(hs.ID.Between(0, 5))
	hs.OrderBy(hs.ID.Desc())
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil {
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------gdaoslave unbinds the hstest1 table----------------------")
}
