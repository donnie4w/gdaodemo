// Copyright (c) 2024, donnie <donnie4w@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// github.com/donnie4w/gdao
//
// datetime :2024-07-26 16:30:05
// gdao version 1.1.0
// dbtype:sqlite ,database:hstest.db ,tablename:hstest2

package dao

import (
	"fmt"
	"github.com/donnie4w/gdao/base"
	"github.com/donnie4w/gdao"
	"time"
)
type hstest2_Id[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest2_Id[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Id[T]) Value() any {
	return t.fieldValue
}

type hstest2_Name[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *string
}

func (t *hstest2_Name[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Name[T]) Value() any {
	return t.fieldValue
}

type hstest2_Age[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest2_Age[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Age[T]) Value() any {
	return t.fieldValue
}

type hstest2_Createtime[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue time.Time
}

func (t *hstest2_Createtime[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Createtime[T]) Value() any {
	return t.fieldValue
}

type hstest2_Money[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *float64
}

func (t *hstest2_Money[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Money[T]) Value() any {
	return t.fieldValue
}

type hstest2_Bytes[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue []byte
}

func (t *hstest2_Bytes[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Bytes[T]) Value() any {
	return t.fieldValue
}

type hstest2_Floa[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *float64
}

func (t *hstest2_Floa[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Floa[T]) Value() any {
	return t.fieldValue
}

type hstest2_Level[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest2_Level[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Level[T]) Value() any {
	return t.fieldValue
}

type hstest2_Type[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *float64
}

func (t *hstest2_Type[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Type[T]) Value() any {
	return t.fieldValue
}

type hstest2_Flog[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *float64
}

func (t *hstest2_Flog[T]) Name() string {
	return t.fieldName
}

func (t *hstest2_Flog[T]) Value() any {
	return t.fieldValue
}

type Hstest2 struct {
	gdao.Table[Hstest2]

	Id		*hstest2_Id[Hstest2]
	Name		*hstest2_Name[Hstest2]
	Age		*hstest2_Age[Hstest2]
	Createtime		*hstest2_Createtime[Hstest2]
	Money		*hstest2_Money[Hstest2]
	Bytes		*hstest2_Bytes[Hstest2]
	Floa		*hstest2_Floa[Hstest2]
	Level		*hstest2_Level[Hstest2]
	Type		*hstest2_Type[Hstest2]
	Flog		*hstest2_Flog[Hstest2]
}

func (u *Hstest2) GetId() (_r int64){
	if u.Id.fieldValue != nil {
		_r = *u.Id.fieldValue
	}
	return
}

func (u *Hstest2) SetId(arg int64) *Hstest2{
	u.Put0(u.Id.fieldName, arg)
	u.Id.fieldValue = &arg
	return u
}

func (u *Hstest2) GetName() (_r string){
	if u.Name.fieldValue != nil {
		_r = *u.Name.fieldValue
	}
	return
}

func (u *Hstest2) SetName(arg string) *Hstest2{
	u.Put0(u.Name.fieldName, arg)
	u.Name.fieldValue = &arg
	return u
}

func (u *Hstest2) GetAge() (_r int64){
	if u.Age.fieldValue != nil {
		_r = *u.Age.fieldValue
	}
	return
}

func (u *Hstest2) SetAge(arg int64) *Hstest2{
	u.Put0(u.Age.fieldName, arg)
	u.Age.fieldValue = &arg
	return u
}

func (u *Hstest2) GetCreatetime() (_r time.Time){
	_r = u.Createtime.fieldValue
	return
}

func (u *Hstest2) SetCreatetime(arg time.Time) *Hstest2{
	u.Put0(u.Createtime.fieldName, arg)
	u.Createtime.fieldValue = arg
	return u
}

func (u *Hstest2) GetMoney() (_r float64){
	if u.Money.fieldValue != nil {
		_r = *u.Money.fieldValue
	}
	return
}

func (u *Hstest2) SetMoney(arg float64) *Hstest2{
	u.Put0(u.Money.fieldName, arg)
	u.Money.fieldValue = &arg
	return u
}

func (u *Hstest2) GetBytes() (_r []byte){
	_r = u.Bytes.fieldValue
	return
}

func (u *Hstest2) SetBytes(arg []byte) *Hstest2{
	u.Put0(u.Bytes.fieldName, arg)
	u.Bytes.fieldValue = arg
	return u
}

func (u *Hstest2) GetFloa() (_r float64){
	if u.Floa.fieldValue != nil {
		_r = *u.Floa.fieldValue
	}
	return
}

func (u *Hstest2) SetFloa(arg float64) *Hstest2{
	u.Put0(u.Floa.fieldName, arg)
	u.Floa.fieldValue = &arg
	return u
}

func (u *Hstest2) GetLevel() (_r int64){
	if u.Level.fieldValue != nil {
		_r = *u.Level.fieldValue
	}
	return
}

func (u *Hstest2) SetLevel(arg int64) *Hstest2{
	u.Put0(u.Level.fieldName, arg)
	u.Level.fieldValue = &arg
	return u
}

func (u *Hstest2) GetType() (_r float64){
	if u.Type.fieldValue != nil {
		_r = *u.Type.fieldValue
	}
	return
}

func (u *Hstest2) SetType(arg float64) *Hstest2{
	u.Put0(u.Type.fieldName, arg)
	u.Type.fieldValue = &arg
	return u
}

func (u *Hstest2) GetFlog() (_r float64){
	if u.Flog.fieldValue != nil {
		_r = *u.Flog.fieldValue
	}
	return
}

func (u *Hstest2) SetFlog(arg float64) *Hstest2{
	u.Put0(u.Flog.fieldName, arg)
	u.Flog.fieldValue = &arg
	return u
}


func (u *Hstest2) Scan(fieldname string, value any) {
	switch fieldname {
	case "id":
		u.SetId(base.AsInt64(value))
	case "name":
		u.SetName(base.AsString(value))
	case "age":
		u.SetAge(base.AsInt64(value))
	case "createtime":
		if t, err := base.AsTime(value); err == nil {
			u.SetCreatetime(t)
		}
	case "money":
		u.SetMoney(base.AsFloat64(value))
	case "bytes":
		u.SetBytes(base.AsBytes(value))
	case "floa":
		u.SetFloa(base.AsFloat64(value))
	case "level":
		u.SetLevel(base.AsInt64(value))
	case "type":
		u.SetType(base.AsFloat64(value))
	case "flog":
		u.SetFlog(base.AsFloat64(value))
	}
}

func (t *Hstest2) ToGdao() {
	_t := NewHstest2()
	*t = *_t
}

func (t *Hstest2) Copy(h *Hstest2) *Hstest2{
	t.SetId(h.GetId())
	t.SetName(h.GetName())
	t.SetAge(h.GetAge())
	t.SetCreatetime(h.GetCreatetime())
	t.SetMoney(h.GetMoney())
	t.SetBytes(h.GetBytes())
	t.SetFloa(h.GetFloa())
	t.SetLevel(h.GetLevel())
	t.SetType(h.GetType())
	t.SetFlog(h.GetFlog())
	return t
}

func (t *Hstest2) String() string {
	return fmt.Sprint("Id:",t.GetId(), ",","Name:",t.GetName(), ",","Age:",t.GetAge(), ",","Createtime:",t.GetCreatetime(), ",","Money:",t.GetMoney(), ",","Bytes:",t.GetBytes(), ",","Floa:",t.GetFloa(), ",","Level:",t.GetLevel(), ",","Type:",t.GetType(), ",","Flog:",t.GetFlog())
}

func NewHstest2(tablename ...string) (_r *Hstest2) {

	id := &hstest2_Id[Hstest2]{fieldName: "id"}
	id.Field.FieldName = "id"

	name := &hstest2_Name[Hstest2]{fieldName: "name"}
	name.Field.FieldName = "name"

	age := &hstest2_Age[Hstest2]{fieldName: "age"}
	age.Field.FieldName = "age"

	createtime := &hstest2_Createtime[Hstest2]{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"

	money := &hstest2_Money[Hstest2]{fieldName: "money"}
	money.Field.FieldName = "money"

	bytes := &hstest2_Bytes[Hstest2]{fieldName: "bytes"}
	bytes.Field.FieldName = "bytes"

	floa := &hstest2_Floa[Hstest2]{fieldName: "floa"}
	floa.Field.FieldName = "floa"

	level := &hstest2_Level[Hstest2]{fieldName: "level"}
	level.Field.FieldName = "level"

	type_ := &hstest2_Type[Hstest2]{fieldName: "type"}
	type_.Field.FieldName = "type"

	flog := &hstest2_Flog[Hstest2]{fieldName: "flog"}
	flog.Field.FieldName = "flog"

	_r = &Hstest2{Id:id,Name:name,Age:age,Createtime:createtime,Money:money,Bytes:bytes,Floa:floa,Level:level,Type:type_,Flog:flog}
	s := "hstest2"
	if len(tablename) > 0 && tablename[0] != "" {
		s = tablename[0]
	}
	_r.Init(s, []base.Column[Hstest2]{id,name,age,createtime,money,bytes,floa,level,type_,flog})
	return
}

func (t *Hstest2) Encode() ([]byte, error) {
	m := make(map[string]any, 0)
	m["id"] = t.GetId()
	m["name"] = t.GetName()
	m["age"] = t.GetAge()
	m["createtime"] = t.GetCreatetime()
	m["money"] = t.GetMoney()
	m["bytes"] = t.GetBytes()
	m["floa"] = t.GetFloa()
	m["level"] = t.GetLevel()
	m["type"] = t.GetType()
	m["flog"] = t.GetFlog()
	return t.Table.Encode(m)
}

func (t *Hstest2) Decode(bs []byte) (err error) {
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

