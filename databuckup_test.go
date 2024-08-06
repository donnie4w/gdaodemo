package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdaodemo/dao"
	"os"
	"reflect"
	"testing"
)

var filePath = "backupfile.db"
var file *os.File

func init() {
	var err error
	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
}
func AppendToFile(bs []byte) (err error) {
	if len(bs) == 0 {
		return fmt.Errorf("encode result is empty")
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, int32(len(bs)))
	binary.Write(buf, binary.LittleEndian, bs)
	_, err = file.Write(buf.Bytes())
	return
}

func Close() error {
	if file != nil {
		return file.Close()
	}
	return nil
}

func recoverByClass[T any]() (err error) {
	var bs []byte
	if bs, err = os.ReadFile(filePath); err == nil {
		buf := bytes.NewReader(bs)
		for buf.Len() > 0 {
			t := new(T)
			var headlen int32
			binary.Read(buf, binary.LittleEndian, &headlen)
			body := make([]byte, headlen)
			buf.Read(body)
			if table, ok := any(t).(gdao.GStruct[*T, T]); ok {
				if err = table.Decode(body); err == nil {
					fmt.Println(table)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
	return
}

func dataBackUpByClass[T any]() (err error) {
	step := int64(10)
	startindex := int64(0)
	t := new(T)
	if _, ok := any(t).(gdao.GStruct[*T, T]); !ok {
		return fmt.Errorf("error: %v is not a gdao GStruct type", reflect.TypeOf(t).Elem())
	}
	for {
		t := new(T)
		table, _ := any(t).(gdao.GStruct[*T, T])
		table.ToGdao()
		table.Limit2(startindex, step)
		list, _ := table.Selects()
		for _, h := range list {
			if h1, ok := any(h).(gdao.GStruct[*T, T]); ok {
				if bs, _ := h1.Encode(); len(bs) > 0 {
					err = AppendToFile(bs)
				}
			}
		}
		if len(list) < int(step) {
			break
		}
		startindex = startindex + step
	}
	return err
}

func TestHstestBackup(t *testing.T) {
	err := dataBackUpByClass[dao.Hstest]()
	fmt.Println(err)
}

func TestHstestReCover(t *testing.T) {
	err := recoverByClass[dao.Hstest]()
	fmt.Println(err)
}

func TestHstest(t *testing.T) {
	dataBackUpByClass[dao.Hstest]()
	recoverByClass[dao.Hstest]()
}

func m() {
	m := make(map[int]func() error)
	m[1] = recoverByClass[dao.Hstest]
	m[2] = recoverByClass[dao.Hstest1]
}
