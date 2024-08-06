package main

import (
	"context"
	"errors"
	"github.com/donnie4w/gdaodemo/dao"
	"log"
	"testing"
	"time"
)

// Hstest 持久层的使用测试示例
func TestBusinessWithHstest(t *testing.T) {
	var m = GetHstestService()
	hs, err := m.FindById(context.TODO(), "1")
	logger.Debug(hs, err)

	h := dao.NewHstest()
	h.Where(h.Id.Between(1, 3), h.Id.LT(100))
	h.OrderBy(h.Id.Desc())
	hss, _ := m.FindByCriteria(context.TODO(), h)
	for _, hs := range hss {
		logger.Debug(hs)
	}
}

func TestBusinessHstestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	var model = GetHstestService()

	h := dao.NewHstest()
	h.Where(h.Id.Between(1, 10), h.Id.LT(100)).OrderBy(h.Id.Desc()).GroupBy(h.Id)
	if hss, err := model.FindByCriteria(ctx, h); err == nil {
		for _, hs := range hss {
			logger.Debug(hs)
		}
	} else {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Operation timed out")
		} else if errors.Is(err, context.Canceled) {
			log.Println("Operation was canceled")
		} else {
			log.Printf("Failed to create record: %v", err)
		}
	}
}
