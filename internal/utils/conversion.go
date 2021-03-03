package utils

import (
	"log"

	"encoding/json"
	"github.com/dollarkillerx/plumber/pkg/models"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
)

func PkgMQEvent(event *canal.RowsEvent) *models.MQEvent {
	if event == nil {
		return nil
	}

	return &models.MQEvent{
		Table:  PkgTable(event.Table),
		Action: models.Action(event.Action),
		Rows:   PkgRows(event),
		Header: PkgEventHeader(event.Header),
	}
}

func PkgRows(event *canal.RowsEvent) (resp string) {
	resp = ""

	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	if event.Table == nil || len(event.Rows) == 0 {
		return resp
	}

	c := make([]string, 0)
	for _, v := range event.Table.Columns {
		c = append(c, v.Name)
	}

	r := make([]map[string]interface{}, 0)
	for _, v := range event.Rows {
		vc := map[string]interface{}{}
		for k, vv := range v {
			vc[c[k]] = vv
		}

		r = append(r, vc)
	}

	marshal, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return resp
	}

	return string(marshal)
}

func PkgEventHeader(header *replication.EventHeader) *models.EventHeader {
	if header == nil {
		return nil
	}

	return &models.EventHeader{
		Timestamp: header.Timestamp,
		EventType: models.EventType(header.EventType),
		ServerID:  header.ServerID,
		EventSize: header.EventSize,
		LogPos:    header.LogPos,
		Flags:     header.Flags,
	}
}

func PkgIndex(event *schema.Index) *models.Index {
	if event == nil {
		return nil
	}

	return &models.Index{
		Name:        event.Name,
		Columns:     event.Columns,
		Cardinality: event.Cardinality,
	}
}

func PkgTableColumn(column schema.TableColumn) models.TableColumn {
	return models.TableColumn{
		Name:       column.Name,
		Type:       column.Type,
		Collation:  column.Collation,
		RawType:    column.RawType,
		IsAuto:     column.IsAuto,
		IsUnsigned: column.IsUnsigned,
		IsVirtual:  column.IsVirtual,
		EnumValues: column.EnumValues,
		SetValues:  column.SetValues,
		FixedSize:  column.FixedSize,
		MaxSize:    column.MaxSize,
	}
}

func PkgTable(event *schema.Table) *models.Table {
	if event == nil {
		return nil
	}

	columns := make([]models.TableColumn, 0)
	for _, v := range event.Columns {
		columns = append(columns, PkgTableColumn(v))
	}

	indexes := make([]*models.Index, 0)
	for _, v := range event.Indexes {
		indexes = append(indexes, PkgIndex(v))
	}

	return &models.Table{
		DBName:          event.Schema,
		TableName:       event.Name,
		Columns:         columns,
		Indexes:         indexes,
		PKColumns:       event.PKColumns,
		UnsignedColumns: event.UnsignedColumns,
	}
}

//func PkgRows(event *canal.RowsEvent) (resp []map[string]interface{}) {
//	resp = make([]map[string]interface{}, 0)
//
//	defer func() {
//		if err := recover(); err != nil {
//			return
//		}
//	}()
//
//	if event.Table == nil || len(event.Rows) == 0 {
//		return resp
//	}
//
//	c := make([]string, 0)
//	for _, v := range event.Table.Columns {
//		c = append(c, v.Name)
//	}
//
//	for _, v := range event.Rows {
//		vc := map[string]interface{}{}
//		for k, vv := range v {
//			vc[c[k]] = vv
//		}
//
//		resp = append(resp, vc)
//	}
//
//	return resp
//}
