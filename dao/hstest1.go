// Copyright (c) 2024, donnie <donnie4w@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// github.com/donnie4w/gdao
//
// datetime :2024-08-16 13:52:41
// gdao version 1.1.0
// dbtype:sqlite ,database:hstest.db ,tablename:hstest1

package dao

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/base"
	
)

type Hstest1 struct {
	gdao.Table[Hstest1]

	ID      *base.Field[Hstest1]
	ROWNAME      *base.Field[Hstest1]
	VALUE      *base.Field[Hstest1]
	GOTO      *base.Field[Hstest1]
	_ID      *int64
	_ROWNAME      *string
	_VALUE      []byte
	_GOTO      []byte
}

var _Hstest1_ID = &base.Field[Hstest1]{"id"}
var _Hstest1_ROWNAME = &base.Field[Hstest1]{"rowname"}
var _Hstest1_VALUE = &base.Field[Hstest1]{"value"}
var _Hstest1_GOTO = &base.Field[Hstest1]{"goto"}

func (u *Hstest1) GetId() (_r int64){
	if u._ID != nil {
		_r = *u._ID
	}
	return
}

func (u *Hstest1) SetId(arg int64) *Hstest1{
	u.Put0(u.ID.FieldName, arg)
	u._ID = &arg
	return u
}

func (u *Hstest1) GetRowname() (_r string){
	if u._ROWNAME != nil {
		_r = *u._ROWNAME
	}
	return
}

func (u *Hstest1) SetRowname(arg string) *Hstest1{
	u.Put0(u.ROWNAME.FieldName, arg)
	u._ROWNAME = &arg
	return u
}

func (u *Hstest1) GetValue() (_r []byte){
	_r = u._VALUE
	return
}

func (u *Hstest1) SetValue(arg []byte) *Hstest1{
	u.Put0(u.VALUE.FieldName, arg)
	u._VALUE = arg
	return u
}

func (u *Hstest1) GetGoto() (_r []byte){
	_r = u._GOTO
	return
}

func (u *Hstest1) SetGoto(arg []byte) *Hstest1{
	u.Put0(u.GOTO.FieldName, arg)
	u._GOTO = arg
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
	t.init("hstest1")
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

func (t *Hstest1)init(tablename string) {
	t.ID = _Hstest1_ID
	t.ROWNAME = _Hstest1_ROWNAME
	t.VALUE = _Hstest1_VALUE
	t.GOTO = _Hstest1_GOTO
	t.Init(tablename, []base.Column[Hstest1]{t.ID,t.ROWNAME,t.VALUE,t.GOTO})
}

func NewHstest1(tablename ...string) (_r *Hstest1) {
	_r = &Hstest1{}
	s := "hstest1"
	if len(tablename) > 0 && tablename[0] != "" {
		s = tablename[0]
	}
	_r.init(s)
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

