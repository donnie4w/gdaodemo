package main

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoCache"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdao/gdaoSlave"
	"github.com/donnie4w/gdaodemo/dao"
	"testing"
)

func init() {
	gdaoMapper.Builder("mappers.xml")
}

func Test_mapper_database(t *testing.T) {
	database, err := gdaoMapper.SelectBean("user.selectHstest1", 1)
	fmt.Println(database, err)

	databases, err := gdaoMapper.SelectsBean("user.selectHstest1", 3)
	for _, database := range databases {
		fmt.Println(database)
	}
}

func Test_mapper_select(t *testing.T) {
	hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1)
	fmt.Println(hs1, err)

	id, err := gdaoMapper.Select[int]("user.selectHstest1", 1)
	fmt.Println(*id, err)
}

func Test_mapper_select2(t *testing.T) {
	hss, err := gdaoMapper.SelectsAny[dao.Hstest]("picture.selectHstest", []int{10, 20})
	fmt.Println(err)
	for _, hs := range hss {
		fmt.Println(hs)
	}

	hslist, _ := gdaoMapper.Selects[dao.Hstest]("picture.selectHstestLimit", 5)
	for _, hs := range hslist {
		fmt.Println(hs)
	}
}

func Test_mapper_selectMap(t *testing.T) {
	m := make(map[string]interface{})
	m["id"] = 5
	m["age"] = 25
	mapper := gdaoMapper.NewInstance()
	if db, err := mapper.SelectAny("user.selectHstestById", m); err == nil {
		fmt.Println(db)
	} else {
		fmt.Println(err)
	}
}

func Test_mapper_selectcache(t *testing.T) {
	gdaoCache.BindMapperId("user", "selectHstest1")
	if hs1, err := gdaoMapper.SelectAny[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}

	fmt.Println("-----------------------Set Cache-------------------------")
	if hs1, err := gdaoMapper.SelectAny[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}

	fmt.Println("-----------------------Get Cache-------------------------")
	gdaoCache.UnbindMapperId("user", "selectHstest1")
	if hs1, err := gdaoMapper.SelectAny[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------No Use Cache-------------------------")
}

func Test_mapper_selectcachehandle(t *testing.T) {
	gdaoCache.BindMapperWithCacheHandle("user", gdaoCache.NewCacheHandle().SetExpire(100))
	mapper := gdaoMapper.NewInstance()
	if db, err := mapper.SelectAny("user.selectHstest1", 1); err == nil {
		fmt.Println(db)
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------Set Cache-------------------------")
	if db, err := mapper.SelectAny("user.selectHstest1", 1); err == nil {
		fmt.Println(db)
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------Get Cache-------------------------")
	//gdaoCache.UnbindMapper("user")
	gdaoCache.UnbindMapperId("user", "selectHstest1")
	if db, err := mapper.SelectAny("user.selectHstest1", 1); err == nil {
		fmt.Println(db)
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------No Use Cache-------------------------")
}

func Test_mapper_SelectsAny(t *testing.T) {
	mapper := gdaoMapper.NewInstance()
	if dbs, err := mapper.SelectsAny("user.selectHstest1", 3); err == nil {
		for _, db := range dbs {
			fmt.Println(db)
		}
	} else {
		fmt.Println(err)
	}
}

func Test_mapper_SelectsAnyCache(t *testing.T) {
	gdaoCache.BindMapperId("user", "selectHstest1")
	mapper := gdaoMapper.NewInstance()
	if dbs, err := mapper.SelectsAny("user.selectHstest1", 3); err == nil {
		for _, db := range dbs {
			fmt.Println(db)
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------Set Cache-------------------------")
	if dbs, err := mapper.SelectsAny("user.selectHstest1", 3); err == nil {
		for _, db := range dbs {
			fmt.Println(db)
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------Get Cache-------------------------")
	gdaoCache.UnbindMapperId("user", "selectHstest1")
	if dbs, err := mapper.SelectsAny("user.selectHstest1", 3); err == nil {
		for _, db := range dbs {
			fmt.Println(db)
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------No Use Cache-------------------------")
}

func Test_mapper_selectMapper(t *testing.T) {
	if id, err := gdaoMapper.Select[int]("user.selectHstest1", 1); err == nil {
		fmt.Println(*id)
	} else {
		fmt.Println(err)
	}

	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}
}

func Test_mapper_insert(t *testing.T) {
	hstest1 := dao.NewHstest1()
	hstest1.SetRowname("rowname>>>12345678")
	hstest1.SetValue([]byte("hello gdao"))
	hstest1.SetGoto([]byte("12345"))
	if _, err := gdaoMapper.InsertAny("user.insertHstest1", hstest1); err != nil {
		fmt.Println(err)
	}
	println("--------------------------InsertAny-----------------------------")
	println()
	hstest2 := dao.NewHstest1()
	hstest2.OrderBy(hstest2.Id.Desc()).Limit(1)
	hstest2, _ = hstest2.Select()
	fmt.Println(hstest2) //查询最后插入的数据

	hstest2.SetRowname("rowname>>>123456789")
	if _, err := gdaoMapper.UpdateAny("user.updateHstest1", hstest2); err != nil {
		fmt.Println(err)
	}
	fmt.Println("--------------------------UpdateAny-----------------------------")
	hstest2 = dao.NewHstest1()
	hstest2.OrderBy(hstest2.Id.Desc()).Limit(1)
	hstest2, _ = hstest2.Select()
	fmt.Println(hstest2) //查询最后更新的数据
}

func Test_mapper_slave(t *testing.T) {
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	}
	println("----------------------use master---------------------------")
	println()
	gdaoSlave.BindMapper("user", DataSource.Mysql(), gdao.MYSQL)
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
	gdaoSlave.BindMapperId("user", "selectHstest1", DataSource.Mysql(), gdao.MYSQL)
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

func Test_mapper_Scan(t *testing.T) {
	mapper := gdaoMapper.NewInstance()
	if databean, err := mapper.SelectAny("user.selectHstest1", 1); err == nil {
		fmt.Println("databean>>>", databean)
	}

	i, _ := gdaoMapper.SelectAnyWithGdaoMapper[int](mapper, "user.selectHstest1", 1)
	fmt.Println("i>>>", *i)

	s, _ := gdaoMapper.SelectWithGdaoMapper[string](mapper, "user.selectHstest1", 1)
	fmt.Println("s>>>", *s)

	hs1, _ := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1)
	fmt.Println("hs1>>>", hs1)
}

func Test_mapper_datasource(t *testing.T) {
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}
	println("----------------------sqlite datasource---------------------------")
	println()
	gdao.BindMapperDataSource("user", DataSource.Mysql(), gdao.MYSQL)
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}
	println("----------------------mysql datasource---------------------------")
}
