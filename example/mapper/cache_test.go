package mapper

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoCache"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"testing"
)

func init() {
	gdao.SetLogger(true)
	gdaodemo.RootDir = "../../"
}

func Test_mapper_selectcache(t *testing.T) {
	gdaoCache.BindMapperId("user", "selectHstest1")
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}

	fmt.Println("-----------------------Set Cache-------------------------")
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}

	fmt.Println("-----------------------Get Cache-------------------------")
	gdaoCache.UnbindMapperId("user", "selectHstest1")
	if hs1, err := gdaoMapper.Select[dao.Hstest1]("user.selectHstest1", 1); err == nil {
		fmt.Println(hs1)
	} else {
		fmt.Println(err)
	}
	fmt.Println("-----------------------No Use Cache-------------------------")
}

func Test_mapper_selectcachehandle(t *testing.T) {
	gdaoCache.BindMapperWithCacheHandle("user", gdaoCache.NewCacheHandle().SetExpire(100))
	mapper := gdaoMapper.NewInstance()
	if db := mapper.SelectBean("user.selectHstest1", 1); db.GetError() == nil {
		fmt.Println(db)
	} else {
		fmt.Println(db.GetError())
	}
	fmt.Println("-----------------------Set Cache-------------------------")
	if db := mapper.SelectBean("user.selectHstest1", 1); db.GetError() == nil {
		fmt.Println(db)
	} else {
		fmt.Println(db.GetError())
	}
	fmt.Println("-----------------------Get Cache-------------------------")
	//gdaoCache.UnbindMapper("user")
	gdaoCache.UnbindMapperId("user", "selectHstest1")
	if db := mapper.SelectBean("user.selectHstest1", 1); db.GetError() == nil {
		fmt.Println(db)
	} else {
		fmt.Println(db.GetError())
	}
	fmt.Println("-----------------------No Use Cache-------------------------")
}

func Test_mapper_SelectBeansCache(t *testing.T) {
	gdaoCache.BindMapperId("user", "selectHstest1")
	mapper := gdaoMapper.NewInstance()
	if dbs := mapper.SelectBeans("user.selectHstest1", 3); dbs.GetError() == nil {
		for _, db := range dbs.Beans {
			fmt.Println(db)
		}
	} else {
		fmt.Println(dbs.GetError())
	}
	fmt.Println("-----------------------Set Cache-------------------------")
	if dbs := mapper.SelectBeans("user.selectHstest1", 3); dbs.GetError() == nil {
		for _, db := range dbs.Beans {
			fmt.Println(db)
		}
	} else {
		fmt.Println(dbs.GetError())
	}
	fmt.Println("-----------------------Get Cache-------------------------")
	gdaoCache.UnbindMapperId("user", "selectHstest1")
	if dbs := mapper.SelectBeans("user.selectHstest1", 3); dbs.GetError() == nil {
		for _, db := range dbs.Beans {
			fmt.Println(db)
		}
	} else {
		fmt.Println(dbs.GetError())
	}
	fmt.Println("-----------------------No Use Cache-------------------------")
}
