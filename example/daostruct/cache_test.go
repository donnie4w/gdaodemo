package gdaodemo

import (
	"github.com/donnie4w/gdao/gdaoCache"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/go-logger/logger"
	"testing"
)

func TestCacheTablename(t *testing.T) {
	//设置表hstest使用缓存
	gdaoCache.BindTableNames("hstest")
	hs := dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(1)
	if hs, err := hs.Select(); err == nil { //第一次查询，缓冲池没有数据，则结果集放入缓冲池
		logger.Debug(hs)
	}
	logger.Debug("----------------------Set Cache----------------------")
	println()
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(1)
	if hs, err := hs.Select(); err == nil { //第二次查询，缓冲池有数据，则返回缓存数据
		logger.Debug(hs)
	}
	logger.Debug("----------------------Get Cache----------------------")
	println()
	gdaoCache.UnbindTableNames("hstest") // 解绑缓存设置
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(1)
	if hs, err := hs.Select(); err == nil { //第三次查询，已解绑，不再设置缓存数据
		logger.Debug(hs)
	}
	logger.Debug("----------------------No Use Cache----------------------")
}

func TestCacheClass(t *testing.T) {
	gdaoCache.BindClassWithCacheHandle[*dao.Hstest](gdaoCache.NewCacheHandle().SetExpire(100).SetStoreMode(gdaoCache.STRONG)) //set cache for Hstest
	hs := dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil { //第一次查询，缓冲池没有数据，则结果集放入缓冲池
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Set Cache----------------------")
	println()
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil { //第二次查询，缓冲池有数据，则返回缓存数据
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Get Cache----------------------")
	println()
	gdaoCache.ClearClass[dao.Hstest]() //清理缓存数据
	hs = dao.NewHstest()
	hs.Where((hs.ID.Between(0, 2)).Or(hs.ID.Between(10, 15)))
	hs.Limit(3)
	if hslist, err := hs.Selects(); err == nil { //第三次查询，缓存数据已经清理，则重新将结果放入缓冲池
		for _, hs := range hslist {
			logger.Debug(hs)
		}
	}
	logger.Debug("----------------------Set Cache Again----------------------")
}
