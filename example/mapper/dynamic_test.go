package mapper

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"path/filepath"
	"testing"
)

func init() {
	gdaodemo.RootDir = "../../"
	gdaoMapper.Builder(filepath.Join(gdaodemo.RootDir, "dynamic.xml"))
	gdao.Init(gdaodemo.DataSource.Sqlite(), gdao.SQLITE)
}

func Test_demo1(t *testing.T) {
	hs := dao.NewHstest()
	hs.SetId(12)
	database := gdaoMapper.SelectBean("dynamic.demo1", hs)
	fmt.Println(database)
}

func Test_demo2(t *testing.T) {
	hs := dao.NewHstest()
	hs.SetId(12)
	database := gdaoMapper.SelectBean("dynamic.demo2", hs)
	fmt.Println(database)
}

func Test_demo3(t *testing.T) {
	as := []int{11, 12, 13, 14, 15}
	database := gdaoMapper.SelectBean("dynamic.demo3", as)
	fmt.Println(database)
}

func Test_demo4_1(t *testing.T) {
	hs := dao.NewHstest()
	hs.SetId(12)
	hs2 := dao.NewHstest()
	hs2.SetId(13)
	hs3 := dao.NewHstest()
	hs3.SetId(14)
	nodes := []*dao.Hstest{hs, hs2, hs3}
	m := map[string][]*dao.Hstest{"hstest": nodes}
	database := gdaoMapper.SelectBean("dynamic.demo4", m)
	fmt.Println(database)
}

func Test_demo4_2(t *testing.T) {
	hs := &dao.Hs1{Id: 12}
	hs1 := &dao.Hs1{Id: 13}
	hs2 := &dao.Hs1{Id: 14}
	nodes := []*dao.Hs1{hs, hs1, hs2}
	m := map[string][]*dao.Hs1{"hstest": nodes}
	database := gdaoMapper.SelectBean("dynamic.demo4", m)
	fmt.Println(database)
}

func Test_demo5_1(t *testing.T) {
	database := gdaoMapper.SelectBean("dynamic.demo5", []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	fmt.Println(database)
}

func Test_demo5_2(t *testing.T) {
	database := gdaoMapper.SelectBean("dynamic.demo5", 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(database)
}

func Test_demo6(t *testing.T) {
	m := map[string]any{}
	m["rowname"] = "hello"
	m["id"] = 11
	database := gdaoMapper.SelectBean("dynamic.demo6", m)
	fmt.Println(database)
}