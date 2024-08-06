package main

import (
	"context"
	"github.com/donnie4w/gdaodemo/dao"
)

// orm 持久层 抽象层 示例

// TableModel 接口定义了数据访问的基本操作
// 持久层 的抽象层示例
type TableModel[T any] interface {
	FindByCriteria(ctx context.Context, t *T) ([]*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	FindById(ctx context.Context, id any) (*T, error)
	Update(ctx context.Context, t *T) (int64, error)
	Delete(ctx context.Context, t *T) (int64, error)
	Save(ctx context.Context, t *T) (int64, error)
}

type HstestService struct {
}

func (h *HstestService) FindByCriteria(ctx context.Context, hstest *dao.Hstest) ([]*dao.Hstest, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return hstest.Selects()
}

func (h *HstestService) FindAll(ctx context.Context) ([]*dao.Hstest, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	hs := dao.NewHstest()
	return hs.Selects()
}

func (h *HstestService) FindById(ctx context.Context, id any) (*dao.Hstest, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	hs := dao.NewHstest()
	hs.Where(hs.Id.EQ(id))
	return hs.Select()
}

func (h *HstestService) Update(ctx context.Context, t *dao.Hstest) (int64, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return t.Update()
}

func (HstestService) Delete(ctx context.Context, t *dao.Hstest) (int64, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return t.Delete()
}

func (HstestService) Save(ctx context.Context, t *dao.Hstest) (int64, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return t.Insert()
}

func GetHstestService() TableModel[dao.Hstest] {
	return new(HstestService)
}
