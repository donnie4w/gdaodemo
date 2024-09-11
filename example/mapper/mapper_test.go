package mapper

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"testing"
)

func Test_selectHstest1_1(t *testing.T) {
	hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1)
	fmt.Println(hs1, err)
}

func Test_selectHstest1_2(t *testing.T) {
	id, err := gdaoMapper.Select[int]("user.selectHstest1", 1)
	fmt.Println(*id, err)
}

func Test_selectHstest1_3(t *testing.T) {
	mapper := gdaoMapper.NewInstance()
	if dbs := mapper.SelectBeans("user.selectHstest1", 3); dbs.GetError() == nil {
		for _, db := range dbs.Beans {
			fmt.Println(db)
		}
	} else {
		fmt.Println(dbs.GetError())
	}
}

func Test_selectHstest1_4(t *testing.T) {
	mapper := gdaoMapper.NewInstance()
	var hs []*dao.Hstest1
	err := mapper.SelectBeans("user.selectHstest1", 3).Scan(&hs)
	fmt.Println(err)
	for _, h := range hs {
		fmt.Println(h)
	}
}

func Test_selectHstest(t *testing.T) {
	hss, err := gdaoMapper.Selects[dao.Hstest]("picture.selectHstest", []int{10, 20})
	fmt.Println(err)
	for _, hs := range hss {
		fmt.Println(hs)
	}
}

func Test_selectHstestLimit(t *testing.T) {
	hslist, _ := gdaoMapper.Selects[dao.Hstest]("picture.selectHstestLimit", 5)
	for _, hs := range hslist {
		fmt.Println(hs)
	}
}

func Test_selectHstestById(t *testing.T) {
	m := make(map[string]interface{})
	m["id"] = 5
	m["age"] = 25
	mapper := gdaoMapper.NewInstance()
	if db := mapper.SelectBean("user.selectHstestById", m); db.GetError() == nil {
		fmt.Println(db)
	} else {
		fmt.Println(db.GetError())
	}
}

func Test_insertHstest1(t *testing.T) {
	hstest1 := dao.NewHstest1()
	hstest1.SetRowname("rowname>>>12345678")
	hstest1.SetValue([]byte("hello gdao"))
	hstest1.SetGoto([]byte("12345"))
	if _, err := gdaoMapper.Insert("user.insertHstest1", hstest1); err != nil {
		fmt.Println(err)
	}
}

func Test_updateHstest1(t *testing.T) {
	hstest2 := dao.NewHstest1()
	hstest2.OrderBy(hstest2.ID.Desc()).Limit(1)
	hstest2, _ = hstest2.Select()
	fmt.Println(hstest2) //查询最后插入的数据

	hstest2.SetRowname("rowname>>>123456789")
	if _, err := gdaoMapper.Update("user.updateHstest1", hstest2); err != nil {
		fmt.Println(err)
	}
}

func Test_mapper_Scan(t *testing.T) {
	mapper := gdaoMapper.NewInstance()
	if databean := mapper.SelectBean("user.selectHstest1", 1); databean.GetError() == nil {
		fmt.Println("databean>>>", databean)
	}

	i := gdaoMapper.SelectBean("user.selectHstest1", 1).ToInt64()
	fmt.Println("i>>>", i)

	i2, _ := gdaoMapper.Select[int]("user.selectHstest1", 1)
	fmt.Println("i2>>>", *i2)

	s := gdaoMapper.SelectBean("user.selectHstest1", 1).ToString()
	fmt.Println("s>>>", s)

	var hs dao.Hstest1
	err := gdaoMapper.SelectBean("user.selectHstest1", 1).Scan(&hs)
	fmt.Println("hs>>>", &hs, err)

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
	gdao.BindMapperDataSource("user", gdaodemo.DataSource.Mysql(), gdao.MYSQL)
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}
	println("----------------------mysql datasource---------------------------")
}
