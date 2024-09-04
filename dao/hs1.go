package dao

import (
	"time"
)

type Hs1 struct {
	Id         int64
	Rowname    string
	Value      []byte
	Goto       []byte
	Updatetime time.Time
}

func (this *Hs1) GetId() int64 {
	return this.Id
}
func (this *Hs1) GetRowname() string {
	return this.Rowname
}
func (this *Hs1) GetValue() []byte {
	return this.Value
}
func (this *Hs1) GetGoto() []byte {
	return this.Goto
}
func (this *Hs1) SetValue(v []byte) {
	this.Value = v
}
func (this *Hs1) SetGoto(v []byte) {
	this.Goto = v
}
func (this *Hs1) SetId(id int64) {
	this.Id = id
}
func (this *Hs1) SetRowname(rowname string) {
	this.Rowname = rowname
}
func (this *Hs1) SetUpdatetime(t time.Time) {
	this.Updatetime = t
}
func (this *Hs1) GetUpdatetime() time.Time {
	return this.Updatetime
}
