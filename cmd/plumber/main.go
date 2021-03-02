package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/siddontang/go-mysql/canal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	t2()

	//cfg := replication.BinlogSyncerConfig{
	//	ServerID: 0606,
	//	Flavor:   "mysql",
	//	Host:     "127.0.0.1",
	//	Port:     3305,
	//	User:     "root",
	//	Password: "root",
	//}
	//
	//syncer := replication.NewBinlogSyncer(cfg)
	//// Start sync with specified binlog file and position
	//streamer, _ := syncer.StartSync(mysql.Position{binlogFile, binlogPos})
	//
	//// or you can start a gtid replication like
	//// streamer, _ := syncer.StartSyncGTID(gtidSet)
	//// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	//// the mariadb GTID set likes this "0-1-100"
	//
	//for {
	//	ev, _ := streamer.GetEvent(context.Background())
	//	// Dump event
	//	ev.Dump(os.Stdout)
	//}
	//
	//// or we can use a timeout context
	//for {
	//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//	ev, err := s.GetEvent(ctx)
	//	cancel()
	//
	//	if err == context.DeadlineExceeded {
	//		// meet timeout
	//		continue
	//	}
	//
	//	ev.Dump(os.Stdout)
	//}
}

func t2() {
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

func (h *MyEventHandler) OnTableChanged(schema, table string) error {
	fmt.Println("TableChange: ", schema, "  ", table)
	return nil
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
	marshal, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(string(marshal))
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

// sudo apt-get install mysql-client  依赖mysql-client
