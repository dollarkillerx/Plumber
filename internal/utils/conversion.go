package utils

import (
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
		Rows:   event.Rows,
		Header: PkgEventHeader(event.Header),
	}
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
