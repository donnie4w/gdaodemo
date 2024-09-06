package nativesql

import (
	"fmt"
	. "github.com/donnie4w/gdao/sqlBuilder"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/go-logger/logger"
	"testing"
)

// 动态SQL 示例
// dynamic sql examples

func Test_sqlBuilder_if(t *testing.T) {
	context := map[string]any{"id": 12}

	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest where").
		AppendIf("id>0", context, "id=?", context["id"])

	bean := builder.SelectOne()
	logger.Debug(bean)

}

func Test_sqlBuilder_2(t *testing.T) {
	hs := dao.Hs1{Id: 12}
	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest where").
		AppendIf("id>0", hs, "id=?", hs.GetId())

	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_sqlBuilder_trim(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      15,
	}

	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest").
		AppendTrim("WHERE ", "", "AND", "", func(trimBuilder *SqlBuilder) {
			trimBuilder.AppendIf("rowname != nil", context, "AND rowname = ?", context["rowname"]).
				AppendIf("id > 18", context, "AND id = ?", context["id"])
		}).
		Append("ORDER BY id ASC")
	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_sqlBuilder_choose(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      15,
	}
	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest").
		AppendChoose(context, func(chooseBuilder *ChooseBuilder) {
			chooseBuilder.When("rowname != nil", "AND rowname = ?", context["rowname"]).
				When("id > 0", "AND id = ?", context["id"]).
				Otherwise("id=?", 1)
		}).
		Append("ORDER BY id ASC")
	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_sqlBuilder_foreach(t *testing.T) {
	context := []int{11, 12, 13, 14, 15}
	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest").
		Append("where id in").
		AppendForeach("", context, "id", ",", "(", ")", func(foreach *ForeachBuilder) {
			foreach.Body("#{id}")
		}).
		Append("ORDER BY id ASC")
	fmt.Println(builder.GetSql())
	fmt.Println(builder.GetParameters())
	beans := builder.SelectList()
	for _, bean := range beans.Beans {
		logger.Debug(bean)
	}
}

func Test_sqlBuilder_foreach2(t *testing.T) {
	context := []int{11, 12, 13, 14, 15}
	m := map[string]any{"ids": context}
	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest").
		Append("where id in").
		AppendForeach("ids", m, "id", ",", "(", ")", func(foreach *ForeachBuilder) {
			foreach.Body("#{id}")
		}).
		Append("ORDER BY id ASC")
	beans := builder.SelectList()
	for _, bean := range beans.Beans {
		logger.Debug(bean)
	}
}

func Test_AppendChoose(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      13,
	}
	builder := &SqlBuilder{}
	builder.Append("SELECT * FROM hstest where 1=1").
		AppendChoose(context, func(chooseBuilder *ChooseBuilder) {
			chooseBuilder.When("id > 0", "AND id = ?", context["id"])
			chooseBuilder.When("rowname != nil", "AND rowname = ?", context["rowname"])
			chooseBuilder.Otherwise("AND id=1")
		}).
		Append("limit 1")

	bean := builder.SelectOne()
	logger.Debug(bean)
}

func Test_AppendSet(t *testing.T) {
	context := map[string]any{
		"rowname": "hello",
		"id":      15,
	}
	builder := &SqlBuilder{}
	builder.Append("UPDATE hstest").
		AppendSet(func(setBuilder *SqlBuilder) {
			setBuilder.AppendIf("rowname != nil", context, "rowname = ?,", context["rowname"]).
				AppendIf("id >0 ", context, "id = ?,", context["id"])
		}).
		Append("WHERE id = ?", context["id"])
	builder.Exec()
}
