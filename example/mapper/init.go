package mapper

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdaodemo"
	"path/filepath"
)

func init() {
	gdao.SetLogger(true)
	gdaodemo.RootDir = "../../"
	gdao.Init(gdaodemo.DataSource.Sqlite(), gdao.SQLITE)
	e1 := gdaoMapper.Builder(filepath.Join(gdaodemo.RootDir, "mapper.xml"))
	e2 := gdaoMapper.Builder(filepath.Join(gdaodemo.RootDir, "mapper2.xml"))
	e3 := gdaoMapper.Builder(filepath.Join(gdaodemo.RootDir, "dynamic.xml"))
	if e1 != nil || e2 != nil || e3 != nil {
		panic(fmt.Sprint(e1, e2, e3))
	}
}
