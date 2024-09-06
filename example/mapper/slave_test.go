package mapper

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdao/gdaoSlave"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"testing"
)

func init() {
	gdao.SetLogger(true)
	gdaodemo.RootDir = "../../"
}

// Read-write separation example 读写分离示例
func Test_mapper_slave(t *testing.T) {
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	}
	println("----------------------use master---------------------------")
	println()
	gdaoSlave.BindMapper("user", gdaodemo.DataSource.Mysql(), gdao.MYSQL)
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	}
	println("----------------------use slave---------------------------")
	println()
	gdaoSlave.UnbindMapperId("user", "selectHstest1")
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	}
	println("----------------------unbind slave---------------------------")
	println()
	gdaoSlave.BindMapperId("user", "selectHstest1", gdaodemo.DataSource.Mysql(), gdao.MYSQL)
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	}
	println("----------------------use slave---------------------------")
	println()
	gdaoSlave.UnbindMapper("user")
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	}
	println("----------------------unbind slave---------------------------")
	println()
}
