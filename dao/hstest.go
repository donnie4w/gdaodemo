// Copyright (c) 2024, donnie <donnie4w@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// github.com/donnie4w/gdao
//
// datetime :2024-08-16 13:52:41
// gdao version 1.1.0
// dbtype:sqlite ,database:hstest.db ,tablename:hstest

package dao

import (
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/base"
	"time"
)

type Hstest struct {
	gdao.Table[Hstest]

	ID          *base.Field[Hstest]
	AGE         *base.Field[Hstest]
	ROWNAME     *base.Field[Hstest]
	VALUE       *base.Field[Hstest]
	UPDATETIME  *base.Field[Hstest]
	BODY        *base.Field[Hstest]
	FLOA        *base.Field[Hstest]
	LEVEL       *base.Field[Hstest]
	_ID         *int64
	_AGE        *int64
	_ROWNAME    *string
	_VALUE      *string
	_UPDATETIME time.Time
	_BODY       []byte
	_FLOA       *float64
	_LEVEL      *int64
}

var _Hstest_ID = &base.Field[Hstest]{"id"}
var _Hstest_AGE = &base.Field[Hstest]{"age"}
var _Hstest_ROWNAME = &base.Field[Hstest]{"rowname"}
var _Hstest_VALUE = &base.Field[Hstest]{"value"}
var _Hstest_UPDATETIME = &base.Field[Hstest]{"updatetime"}
var _Hstest_BODY = &base.Field[Hstest]{"body"}
var _Hstest_FLOA = &base.Field[Hstest]{"floa"}
var _Hstest_LEVEL = &base.Field[Hstest]{"level"}

func (u *Hstest) GetId() (_r int64) {
	if u._ID != nil {
		_r = *u._ID
	}
	return
}

func (u *Hstest) SetId(arg int64) *Hstest {
	u.Put0(u.ID.FieldName, arg)
	u._ID = &arg
	return u
}

func (u *Hstest) GetAge() (_r int64) {
	if u._AGE != nil {
		_r = *u._AGE
	}
	return
}

func (u *Hstest) SetAge(arg int64) *Hstest {
	u.Put0(u.AGE.FieldName, arg)
	u._AGE = &arg
	return u
}

func (u *Hstest) GetRowname() (_r string) {
	if u._ROWNAME != nil {
		_r = *u._ROWNAME
	}
	return
}

func (u *Hstest) SetRowname(arg string) *Hstest {
	u.Put0(u.ROWNAME.FieldName, arg)
	u._ROWNAME = &arg
	return u
}

func (u *Hstest) GetValue() (_r string) {
	if u._VALUE != nil {
		_r = *u._VALUE
	}
	return
}

func (u *Hstest) SetValue(arg string) *Hstest {
	u.Put0(u.VALUE.FieldName, arg)
	u._VALUE = &arg
	return u
}

func (u *Hstest) GetUpdatetime() (_r time.Time) {
	_r = u._UPDATETIME
	return
}

func (u *Hstest) SetUpdatetime(arg time.Time) *Hstest {
	u.Put0(u.UPDATETIME.FieldName, arg)
	u._UPDATETIME = arg
	return u
}

func (u *Hstest) GetBody() (_r []byte) {
	_r = u._BODY
	return
}

func (u *Hstest) SetBody(arg []byte) *Hstest {
	u.Put0(u.BODY.FieldName, arg)
	u._BODY = arg
	return u
}

func (u *Hstest) GetFloa() (_r float64) {
	if u._FLOA != nil {
		_r = *u._FLOA
	}
	return
}

func (u *Hstest) SetFloa(arg float64) *Hstest {
	u.Put0(u.FLOA.FieldName, arg)
	u._FLOA = &arg
	return u
}

func (u *Hstest) GetLevel() (_r int64) {
	if u._LEVEL != nil {
		_r = *u._LEVEL
	}
	return
}

func (u *Hstest) SetLevel(arg int64) *Hstest {
	u.Put0(u.LEVEL.FieldName, arg)
	u._LEVEL = &arg
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
	t.init("hstest")
}

func (t *Hstest) Copy(h *Hstest) *Hstest {
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
	return fmt.Sprint("Id:", t.GetId(), ",", "Age:", t.GetAge(), ",", "Rowname:", t.GetRowname(), ",", "Value:", t.GetValue(), ",", "Updatetime:", t.GetUpdatetime(), ",", "Body:", t.GetBody(), ",", "Floa:", t.GetFloa(), ",", "Level:", t.GetLevel())
}

func (t *Hstest) init(tablename string) {
	t.ID = _Hstest_ID
	t.AGE = _Hstest_AGE
	t.ROWNAME = _Hstest_ROWNAME
	t.VALUE = _Hstest_VALUE
	t.UPDATETIME = _Hstest_UPDATETIME
	t.BODY = _Hstest_BODY
	t.FLOA = _Hstest_FLOA
	t.LEVEL = _Hstest_LEVEL
	t.Init(tablename, []base.Column[Hstest]{t.ID, t.AGE, t.ROWNAME, t.VALUE, t.UPDATETIME, t.BODY, t.FLOA, t.LEVEL})
}

func NewHstest(tablename ...string) (_r *Hstest) {
	_r = &Hstest{}
	s := "hstest"
	if len(tablename) > 0 && tablename[0] != "" {
		s = tablename[0]
	}
	_r.init(s)
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
