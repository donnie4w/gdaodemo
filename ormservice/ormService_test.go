package ormservice

import (
	"context"
	"errors"
	"github.com/donnie4w/gdaodemo"
	"github.com/donnie4w/gdaodemo/dao"
	"log"
	"testing"
	"time"
)

// Hstest 持久层的使用测试示例
func TestBusinessWithHstest(t *testing.T) {
	var m = GetHstestService()
	hs, err := m.FindById(context.TODO(), "1")
	main.logger.Debug(hs, err)

	h := dao.NewHstest()
	h.Where(h.ID.Between(1, 3), h.ID.LT(100))
	h.OrderBy(h.ID.Desc())
	hss, _ := m.FindByCriteria(context.TODO(), h)
	for _, hs := range hss {
		main.logger.Debug(hs)
	}
}

func TestBusinessHstestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	var model = GetHstestService()

	h := dao.NewHstest()
	h.Where(h.ID.Between(1, 10), h.ID.LT(100)).OrderBy(h.ID.Desc()).GroupBy(h.ID)
	if hss, err := model.FindByCriteria(ctx, h); err == nil {
		for _, hs := range hss {
			main.logger.Debug(hs)
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
