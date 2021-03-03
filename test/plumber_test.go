package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dollarkillerx/plumber/internal/utils"
	"github.com/siddontang/go-mysql/canal"
)

func TestBase(t *testing.T) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3305"
	cfg.User = "root"
	cfg.Password = "root"
	// We only care table canal_test in test db
	//cfg.Dump.TableDB = "test"
	//cfg.Dump.Tables = []string{"canal_test"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	if e == nil {
		return nil
	}
	if e.Header == nil {
		return nil
	}

	if int64(e.Header.Timestamp) < time.Now().Unix() {
		return nil
	}
	//log.Printf("%s %v\n", e.Action, e.Rows)
	//marshal, err := json.Marshal(e)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//
	//fmt.Println(string(marshal))

	//c := make([]string, 0)
	//for _, v := range e.Table.Columns {
	//	c = append(c, v.Name)
	//}
	//r := make([]map[string]interface{}, 0)
	//for _, v := range e.Rows {
	//	vc := map[string]interface{}{}
	//	for k, vv := range v {
	//		vc[c[k]] = vv
	//	}
	//
	//	r = append(r, vc)
	//}
	//
	//marshal, err := json.Marshal(r)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//fmt.Println(string(marshal))

	event := utils.PkgMQEvent(e)
	marshal, err := json.Marshal(event)
	if err == nil {
		fmt.Println(string(marshal))
	}
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}
