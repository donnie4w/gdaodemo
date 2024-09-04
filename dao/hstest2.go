// Copyright (c) 2024, donnie <donnie4w@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// github.com/donnie4w/gdao
//
// datetime :2024-08-16 13:52:41
// gdao version 1.1.0
// dbtype:sqlite ,database:hstest.db ,tablename:hstest2

package dao

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/base"
	"time"
)

type Hstest2 struct {
	gdao.Table[Hstest2]

	ID      *base.Field[Hstest2]
	NAME      *base.Field[Hstest2]
	AGE      *base.Field[Hstest2]
	CREATETIME      *base.Field[Hstest2]
	MONEY      *base.Field[Hstest2]
	BYTES      *base.Field[Hstest2]
	FLOA      *base.Field[Hstest2]
	LEVEL      *base.Field[Hstest2]
	TYPE      *base.Field[Hstest2]
	FLOG      *base.Field[Hstest2]
	_ID      *int64
	_NAME      *string
	_AGE      *int64
	_CREATETIME      time.Time
	_MONEY      *float64
	_BYTES      []byte
	_FLOA      *float64
	_LEVEL      *int64
	_TYPE      *float64
	_FLOG      *float64
}

var _Hstest2_ID = &base.Field[Hstest2]{"id"}
var _Hstest2_NAME = &base.Field[Hstest2]{"name"}
var _Hstest2_AGE = &base.Field[Hstest2]{"age"}
var _Hstest2_CREATETIME = &base.Field[Hstest2]{"createtime"}
var _Hstest2_MONEY = &base.Field[Hstest2]{"money"}
var _Hstest2_BYTES = &base.Field[Hstest2]{"bytes"}
var _Hstest2_FLOA = &base.Field[Hstest2]{"floa"}
var _Hstest2_LEVEL = &base.Field[Hstest2]{"level"}
var _Hstest2_TYPE = &base.Field[Hstest2]{"type"}
var _Hstest2_FLOG = &base.Field[Hstest2]{"flog"}

func (u *Hstest2) GetId() (_r int64){
	if u._ID != nil {
		_r = *u._ID
	}
	return
}

func (u *Hstest2) SetId(arg int64) *Hstest2{
	u.Put0(u.ID.FieldName, arg)
	u._ID = &arg
	return u
}

func (u *Hstest2) GetName() (_r string){
	if u._NAME != nil {
		_r = *u._NAME
	}
	return
}

func (u *Hstest2) SetName(arg string) *Hstest2{
	u.Put0(u.NAME.FieldName, arg)
	u._NAME = &arg
	return u
}

func (u *Hstest2) GetAge() (_r int64){
	if u._AGE != nil {
		_r = *u._AGE
	}
	return
}

func (u *Hstest2) SetAge(arg int64) *Hstest2{
	u.Put0(u.AGE.FieldName, arg)
	u._AGE = &arg
	return u
}

func (u *Hstest2) GetCreatetime() (_r time.Time){
	_r = u._CREATETIME
	return
}

func (u *Hstest2) SetCreatetime(arg time.Time) *Hstest2{
	u.Put0(u.CREATETIME.FieldName, arg)
	u._CREATETIME = arg
	return u
}

func (u *Hstest2) GetMoney() (_r float64){
	if u._MONEY != nil {
		_r = *u._MONEY
	}
	return
}

func (u *Hstest2) SetMoney(arg float64) *Hstest2{
	u.Put0(u.MONEY.FieldName, arg)
	u._MONEY = &arg
	return u
}

func (u *Hstest2) GetBytes() (_r []byte){
	_r = u._BYTES
	return
}

func (u *Hstest2) SetBytes(arg []byte) *Hstest2{
	u.Put0(u.BYTES.FieldName, arg)
	u._BYTES = arg
	return u
}

func (u *Hstest2) GetFloa() (_r float64){
	if u._FLOA != nil {
		_r = *u._FLOA
	}
	return
}

func (u *Hstest2) SetFloa(arg float64) *Hstest2{
	u.Put0(u.FLOA.FieldName, arg)
	u._FLOA = &arg
	return u
}

func (u *Hstest2) GetLevel() (_r int64){
	if u._LEVEL != nil {
		_r = *u._LEVEL
	}
	return
}

func (u *Hstest2) SetLevel(arg int64) *Hstest2{
	u.Put0(u.LEVEL.FieldName, arg)
	u._LEVEL = &arg
	return u
}

func (u *Hstest2) GetType() (_r float64){
	if u._TYPE != nil {
		_r = *u._TYPE
	}
	return
}

func (u *Hstest2) SetType(arg float64) *Hstest2{
	u.Put0(u.TYPE.FieldName, arg)
	u._TYPE = &arg
	return u
}

func (u *Hstest2) GetFlog() (_r float64){
	if u._FLOG != nil {
		_r = *u._FLOG
	}
	return
}

func (u *Hstest2) SetFlog(arg float64) *Hstest2{
	u.Put0(u.FLOG.FieldName, arg)
	u._FLOG = &arg
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
	t.init("hstest2")
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

func (t *Hstest2)init(tablename string) {
	t.ID = _Hstest2_ID
	t.NAME = _Hstest2_NAME
	t.AGE = _Hstest2_AGE
	t.CREATETIME = _Hstest2_CREATETIME
	t.MONEY = _Hstest2_MONEY
	t.BYTES = _Hstest2_BYTES
	t.FLOA = _Hstest2_FLOA
	t.LEVEL = _Hstest2_LEVEL
	t.TYPE = _Hstest2_TYPE
	t.FLOG = _Hstest2_FLOG
	t.Init(tablename, []base.Column[Hstest2]{t.ID,t.NAME,t.AGE,t.CREATETIME,t.MONEY,t.BYTES,t.FLOA,t.LEVEL,t.TYPE,t.FLOG})
}

func NewHstest2(tablename ...string) (_r *Hstest2) {
	_r = &Hstest2{}
	s := "hstest2"
	if len(tablename) > 0 && tablename[0] != "" {
		s = tablename[0]
	}
	_r.init(s)
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

