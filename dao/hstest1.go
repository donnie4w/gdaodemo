// Copyright (c) 2024, donnie <donnie4w@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// github.com/donnie4w/gdao
//
// datetime :2024-07-26 16:30:05
// gdao version 1.1.0
// dbtype:sqlite ,database:hstest.db ,tablename:hstest1

package dao

import (
	"fmt"
	"github.com/donnie4w/gdao/base"
	"github.com/donnie4w/gdao"
	
)
type hstest1_Id[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *int64
}

func (t *hstest1_Id[T]) Name() string {
	return t.fieldName
}

func (t *hstest1_Id[T]) Value() any {
	return t.fieldValue
}

type hstest1_Rowname[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue *string
}

func (t *hstest1_Rowname[T]) Name() string {
	return t.fieldName
}

func (t *hstest1_Rowname[T]) Value() any {
	return t.fieldValue
}

type hstest1_Value[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue []byte
}

func (t *hstest1_Value[T]) Name() string {
	return t.fieldName
}

func (t *hstest1_Value[T]) Value() any {
	return t.fieldValue
}

type hstest1_Goto[T any] struct {
	base.Field[T]
	fieldName  string
	fieldValue []byte
}

func (t *hstest1_Goto[T]) Name() string {
	return t.fieldName
}

func (t *hstest1_Goto[T]) Value() any {
	return t.fieldValue
}

type Hstest1 struct {
	gdao.Table[Hstest1]

	Id		*hstest1_Id[Hstest1]
	Rowname		*hstest1_Rowname[Hstest1]
	Value		*hstest1_Value[Hstest1]
	Goto		*hstest1_Goto[Hstest1]
}

func (u *Hstest1) GetId() (_r int64){
	if u.Id.fieldValue != nil {
		_r = *u.Id.fieldValue
	}
	return
}

func (u *Hstest1) SetId(arg int64) *Hstest1{
	u.Put0(u.Id.fieldName, arg)
	u.Id.fieldValue = &arg
	return u
}

func (u *Hstest1) GetRowname() (_r string){
	if u.Rowname.fieldValue != nil {
		_r = *u.Rowname.fieldValue
	}
	return
}

func (u *Hstest1) SetRowname(arg string) *Hstest1{
	u.Put0(u.Rowname.fieldName, arg)
	u.Rowname.fieldValue = &arg
	return u
}

func (u *Hstest1) GetValue() (_r []byte){
	_r = u.Value.fieldValue
	return
}

func (u *Hstest1) SetValue(arg []byte) *Hstest1{
	u.Put0(u.Value.fieldName, arg)
	u.Value.fieldValue = arg
	return u
}

func (u *Hstest1) GetGoto() (_r []byte){
	_r = u.Goto.fieldValue
	return
}

func (u *Hstest1) SetGoto(arg []byte) *Hstest1{
	u.Put0(u.Goto.fieldName, arg)
	u.Goto.fieldValue = arg
	return u
}


func (u *Hstest1) Scan(fieldname string, value any) {
	switch fieldname {
	case "id":
		u.SetId(base.AsInt64(value))
	case "rowname":
		u.SetRowname(base.AsString(value))
	case "value":
		u.SetValue(base.AsBytes(value))
	case "goto":
		u.SetGoto(base.AsBytes(value))
	}
}

func (t *Hstest1) ToGdao() {
	_t := NewHstest1()
	*t = *_t
}

func (t *Hstest1) Copy(h *Hstest1) *Hstest1{
	t.SetId(h.GetId())
	t.SetRowname(h.GetRowname())
	t.SetValue(h.GetValue())
	t.SetGoto(h.GetGoto())
	return t
}

func (t *Hstest1) String() string {
	return fmt.Sprint("Id:",t.GetId(), ",","Rowname:",t.GetRowname(), ",","Value:",t.GetValue(), ",","Goto:",t.GetGoto())
}

func NewHstest1(tablename ...string) (_r *Hstest1) {

	id := &hstest1_Id[Hstest1]{fieldName: "id"}
	id.Field.FieldName = "id"

	rowname := &hstest1_Rowname[Hstest1]{fieldName: "rowname"}
	rowname.Field.FieldName = "rowname"

	value := &hstest1_Value[Hstest1]{fieldName: "value"}
	value.Field.FieldName = "value"

	goto_ := &hstest1_Goto[Hstest1]{fieldName: "goto"}
	goto_.Field.FieldName = "goto"

	_r = &Hstest1{Id:id,Rowname:rowname,Value:value,Goto:goto_}
	s := "hstest1"
	if len(tablename) > 0 && tablename[0] != "" {
		s = tablename[0]
	}
	_r.Init(s, []base.Column[Hstest1]{id,rowname,value,goto_})
	return
}

func (t *Hstest1) Encode() ([]byte, error) {
	m := make(map[string]any, 0)
	m["id"] = t.GetId()
	m["rowname"] = t.GetRowname()
	m["value"] = t.GetValue()
	m["goto"] = t.GetGoto()
	return t.Table.Encode(m)
}

func (t *Hstest1) Decode(bs []byte) (err error) {
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

