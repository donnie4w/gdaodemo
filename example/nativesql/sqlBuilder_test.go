package nativesql

import (
	. "github.com/donnie4w/gdao/sqlBuilder"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/go-logger/logger"
	"testing"
)

func Test_Append(t *testing.T) {
	bean := NewSqlBuilder().Append("SELECT * FROM hstest where").
		Append("id=?", 11).
		Append("and rowname<>?", "hello").SelectOne()
	logger.Debug(bean)
}

func Test_AppendScan(t *testing.T) {
	var hs dao.Hstest
	NewSqlBuilder().Append("SELECT * FROM hstest where").
		Append("id=?", 11).
		Append("and rowname<>?", "hello").SelectOne().Scan(&hs)
	logger.Debug(&hs)
}

// 动态SQL 示例
// dynamic sql examples

func Test_AppendIf(t *testing.T) {
	context := map[string]any{"id": 12}
	builder := NewSqlBuilder()
	builder.Append("SELECT * FROM hstest where").
		AppendIf("id>0", context, "id=?", context["id"])

	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_AppendIf2(t *testing.T) {
	hs := dao.Hs1{Id: 12}
	builder := NewSqlBuilder()
	builder.Append("SELECT * FROM hstest where").
		AppendIf("id>0", hs, "id=?", hs.GetId())

	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_AppendTrim(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      15,
	}

	builder := NewSqlBuilder()
	builder.Append("SELECT * FROM hstest").
		AppendTrim("WHERE ", "", "AND", "", func(trimBuilder SqlBuilder) {
			trimBuilder.AppendIf("rowname != nil", context, "AND rowname = ?", context["rowname"]).
				AppendIf("id > 18", context, "AND id = ?", context["id"])
		}).
		Append("ORDER BY id ASC")
	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_AppendChoose(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      13,
	}
	builder := NewSqlBuilder()
	builder.Append("SELECT * FROM hstest where 1=1").
		AppendChoose(context, func(chooseBuilder ChooseBuilder) {
			chooseBuilder.When("id > 0", "AND id = ?", context["id"])
			chooseBuilder.When("rowname != nil", "AND rowname = ?", context["rowname"])
			chooseBuilder.Otherwise("AND id=1")
		}).
		Append("limit 1")

	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_AppendForeach(t *testing.T) {
	context := []int{11, 12, 13, 14, 15}
	builder := NewSqlBuilder()
	builder.Append("SELECT * FROM hstest").
		Append("where id in").
		AppendForeach("", context, "id", ",", "(", ")", func(foreach ForeachBuilder) {
			foreach.Body("#{id}")
		}).
		Append("ORDER BY id ASC")
	beans := builder.SelectList()
	for _, bean := range beans.Beans {
		logger.Debug(bean)
	}
}

func Test_AppendForeach2(t *testing.T) {
	context := []int{11, 12, 13, 14, 15}
	m := map[string]any{"ids": context}
	builder := NewSqlBuilder()
	builder.Append("SELECT * FROM hstest").
		Append("where id in").
		AppendForeach("ids", m, "id", ",", "(", ")", func(foreach ForeachBuilder) {
			foreach.Body("#{id}")
		}).
		Append("ORDER BY id ASC")
	beans := builder.SelectList()
	for _, bean := range beans.Beans {
		logger.Debug(bean)
	}
}

func Test_AppendSet(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      15,
	}
	builder := NewSqlBuilder()
	builder.Append("UPDATE hstest").
		AppendSet(func(setBuilder SqlBuilder) {
			setBuilder.AppendIf("rowname != nil", context, "rowname = ?,", context["rowname"]).
				AppendIf("id >0 ", context, "id = ?,", context["id"])
		}).
		Append("WHERE id = ?", context["id"])
	i, err := builder.Exec()
	logger.Debug(i, err)
}
