// Copyright (c) 2024, donnie <donnie4w@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// github.com/donnie4w/gdao
//
// datetime :2024-08-03 11:09:32
// gdao version 1.1.0
// dbtype:mysql ,database:hstest ,tablename:hstest

package dao

import (
	"fmt"
	"github.com/donnie4w/gdao/base"
	"github.com/donnie4w/gdao"
	"time"
)
type hstest_Id[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest_Id[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Id[T]) Value() any {
	return t.fieldValue
}

type hstest_Age[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest_Age[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Age[T]) Value() any {
	return t.fieldValue
}

type hstest_Rowname[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *string
}

func (t *hstest_Rowname[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Rowname[T]) Value() any {
	return t.fieldValue
}

type hstest_Value[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *string
}

func (t *hstest_Value[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Value[T]) Value() any {
	return t.fieldValue
}

type hstest_Updatetime[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue time.Time
}

func (t *hstest_Updatetime[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Updatetime[T]) Value() any {
	return t.fieldValue
}

type hstest_Body[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue []byte
}

func (t *hstest_Body[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Body[T]) Value() any {
	return t.fieldValue
}

type hstest_Floa[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *float64
}

func (t *hstest_Floa[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Floa[T]) Value() any {
	return t.fieldValue
}

type hstest_Level[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest_Level[T]) Name() string {
	return t.fieldName
}

func (t *hstest_Level[T]) Value() any {
	return t.fieldValue
}

type Hstest struct {
	gdao.Table[Hstest]

	Id		*hstest_Id[Hstest]
	Age		*hstest_Age[Hstest]
	Rowname		*hstest_Rowname[Hstest]
	Value		*hstest_Value[Hstest]
	Updatetime		*hstest_Updatetime[Hstest]
	Body		*hstest_Body[Hstest]
	Floa		*hstest_Floa[Hstest]
	Level		*hstest_Level[Hstest]
}

func (u *Hstest) GetId() (_r int64){
	if u.Id.fieldValue != nil {
		_r = *u.Id.fieldValue
	}
	return
}

func (u *Hstest) SetId(arg int64) *Hstest{
	u.Put0(u.Id.fieldName, arg)
	u.Id.fieldValue = &arg
	return u
}

func (u *Hstest) GetAge() (_r int64){
	if u.Age.fieldValue != nil {
		_r = *u.Age.fieldValue
	}
	return
}

func (u *Hstest) SetAge(arg int64) *Hstest{
	u.Put0(u.Age.fieldName, arg)
	u.Age.fieldValue = &arg
	return u
}

func (u *Hstest) GetRowname() (_r string){
	if u.Rowname.fieldValue != nil {
		_r = *u.Rowname.fieldValue
	}
	return
}

func (u *Hstest) SetRowname(arg string) *Hstest{
	u.Put0(u.Rowname.fieldName, arg)
	u.Rowname.fieldValue = &arg
	return u
}

func (u *Hstest) GetValue() (_r string){
	if u.Value.fieldValue != nil {
		_r = *u.Value.fieldValue
	}
	return
}

func (u *Hstest) SetValue(arg string) *Hstest{
	u.Put0(u.Value.fieldName, arg)
	u.Value.fieldValue = &arg
	return u
}

func (u *Hstest) GetUpdatetime() (_r time.Time){
	_r = u.Updatetime.fieldValue
	return
}

func (u *Hstest) SetUpdatetime(arg time.Time) *Hstest{
	u.Put0(u.Updatetime.fieldName, arg)
	u.Updatetime.fieldValue = arg
	return u
}

func (u *Hstest) GetBody() (_r []byte){
	_r = u.Body.fieldValue
	return
}

func (u *Hstest) SetBody(arg []byte) *Hstest{
	u.Put0(u.Body.fieldName, arg)
	u.Body.fieldValue = arg
	return u
}

func (u *Hstest) GetFloa() (_r float64){
	if u.Floa.fieldValue != nil {
		_r = *u.Floa.fieldValue
	}
	return
}

func (u *Hstest) SetFloa(arg float64) *Hstest{
	u.Put0(u.Floa.fieldName, arg)
	u.Floa.fieldValue = &arg
	return u
}

func (u *Hstest) GetLevel() (_r int64){
	if u.Level.fieldValue != nil {
		_r = *u.Level.fieldValue
	}
	return
}

func (u *Hstest) SetLevel(arg int64) *Hstest{
	u.Put0(u.Level.fieldName, arg)
	u.Level.fieldValue = &arg
	return u
}


func (u *Hstest) Scan(fieldname string, value any) {
	switch fieldname {
	case "id":
		u.SetId(base.AsInt64(value))
	case "age":
		u.SetAge(base.AsInt64(value))
	case "rowname":
		u.SetRowname(base.AsString(value))
	case "value":
		u.SetValue(base.AsString(value))
	case "updatetime":
		if t, err := base.AsTime(value); err == nil {
			u.SetUpdatetime(t)
		}
	case "body":
		u.SetBody(base.AsBytes(value))
	case "floa":
		u.SetFloa(base.AsFloat64(value))
	case "level":
		u.SetLevel(base.AsInt64(value))
	}
}

func (t *Hstest) ToGdao() {
	_t := NewHstest()
	*t = *_t
}

func (t *Hstest) Copy(h *Hstest) *Hstest{
	t.SetId(h.GetId())
	t.SetAge(h.GetAge())
	t.SetRowname(h.GetRowname())
	t.SetValue(h.GetValue())
	t.SetUpdatetime(h.GetUpdatetime())
	t.SetBody(h.GetBody())
	t.SetFloa(h.GetFloa())
	t.SetLevel(h.GetLevel())
	return t
}

func (t *Hstest) String() string {
	return fmt.Sprint("Id:",t.GetId(), ",","Age:",t.GetAge(), ",","Rowname:",t.GetRowname(), ",","Value:",t.GetValue(), ",","Updatetime:",t.GetUpdatetime(), ",","Body:",t.GetBody(), ",","Floa:",t.GetFloa(), ",","Level:",t.GetLevel())
}

func NewHstest(tablename ...string) (_r *Hstest) {

	id := &hstest_Id[Hstest]{fieldName: "id"}
	id.Field.FieldName = "id"

	age := &hstest_Age[Hstest]{fieldName: "age"}
	age.Field.FieldName = "age"

	rowname := &hstest_Rowname[Hstest]{fieldName: "rowname"}
	rowname.Field.FieldName = "rowname"

	value := &hstest_Value[Hstest]{fieldName: "value"}
	value.Field.FieldName = "value"

	updatetime := &hstest_Updatetime[Hstest]{fieldName: "updatetime"}
	updatetime.Field.FieldName = "updatetime"

	body := &hstest_Body[Hstest]{fieldName: "body"}
	body.Field.FieldName = "body"

	floa := &hstest_Floa[Hstest]{fieldName: "floa"}
	floa.Field.FieldName = "floa"

	level := &hstest_Level[Hstest]{fieldName: "level"}
	level.Field.FieldName = "level"

	_r = &Hstest{Id:id,Age:age,Rowname:rowname,Value:value,Updatetime:updatetime,Body:body,Floa:floa,Level:level}
	s := "hstest"
	if len(tablename) > 0 && tablename[0] != "" {
		s = tablename[0]
	}
	_r.Init(s, []base.Column[Hstest]{id,age,rowname,value,updatetime,body,floa,level})
	return
}

func (t *Hstest) Encode() ([]byte, error) {
	m := make(map[string]any, 0)
	m["id"] = t.GetId()
	m["age"] = t.GetAge()
	m["rowname"] = t.GetRowname()
	m["value"] = t.GetValue()
	m["updatetime"] = t.GetUpdatetime()
	m["body"] = t.GetBody()
	m["floa"] = t.GetFloa()
	m["level"] = t.GetLevel()
	return t.Table.Encode(m)
}

func (t *Hstest) Decode(bs []byte) (err error) {
	var m map[string]any
	if m, err = t.Table.Decode(bs); err == nil {
		if !t.IsInit() {
			t.ToGdao()
		}
		for name, bean := range m {
			t.Scan(name, bean)
		}
	}
	return
}

