package models

// 基础DB模型

type Action string

const (
	UpdateAction Action = "update"
	InsertAction Action = "insert"
	DeleteAction Action = "delete"
)

// event
type MQEvent struct {
	Table  *Table          `json:"table"`
	Action Action          `json:"action"`
	Rows   [][]interface{} `json:"rows"`
	Header *EventHeader    `json:"header"`
}

// Table
type Table struct {
	DBName    string `json:"db_name"`
	TableName string `json:"table_name"`

	Columns   []TableColumn `json:"columns"`
	Indexes   []*Index      `json:"indexes"`
	PKColumns []int         `json:"pk_columns"`

	UnsignedColumns []int `json:"unsigned_columns"`
}

type TableColumn struct {
	Name       string   `json:"name"`
	Type       int      `json:"type"`
	Collation  string   `json:"collation"`
	RawType    string   `json:"raw_type"`
	IsAuto     bool     `json:"is_auto"`
	IsUnsigned bool     `json:"is_unsigned"`
	IsVirtual  bool     `json:"is_virtual"`
	EnumValues []string `json:"enum_values"`
	SetValues  []string `json:"set_values"`
	FixedSize  uint     `json:"fixed_size"`
	MaxSize    uint     `json:"max_size"`
}

type Index struct {
	Name        string   `json:"name"`
	Columns     []string `json:"columns"`
	Cardinality []uint64 `json:"cardinality"`
}

type EventHeader struct {
	Timestamp uint32    `json:"timestamp"`
	EventType EventType `json:"event_type"`
	ServerID  uint32    `json:"server_id"`
	EventSize uint32    `json:"event_size"`
	LogPos    uint32    `json:"log_pos"`
	Flags     uint16    `json:"flags"`
}

type EventType byte

const (
	UNKNOWN_EVENT EventType = iota
	START_EVENT_V3
	QUERY_EVENT
	STOP_EVENT
	ROTATE_EVENT
	INTVAR_EVENT
	LOAD_EVENT
	SLAVE_EVENT
	CREATE_FILE_EVENT
	APPEND_BLOCK_EVENT
	EXEC_LOAD_EVENT
	DELETE_FILE_EVENT
	NEW_LOAD_EVENT
	RAND_EVENT
	USER_VAR_EVENT
	FORMAT_DESCRIPTION_EVENT
	XID_EVENT
	BEGIN_LOAD_QUERY_EVENT
	EXECUTE_LOAD_QUERY_EVENT
	TABLE_MAP_EVENT
	WRITE_ROWS_EVENTv0
	UPDATE_ROWS_EVENTv0
	DELETE_ROWS_EVENTv0
	WRITE_ROWS_EVENTv1
	UPDATE_ROWS_EVENTv1
	DELETE_ROWS_EVENTv1
	INCIDENT_EVENT
	HEARTBEAT_EVENT
	IGNORABLE_EVENT
	ROWS_QUERY_EVENT
	WRITE_ROWS_EVENTv2
	UPDATE_ROWS_EVENTv2
	DELETE_ROWS_EVENTv2
	GTID_EVENT
	ANONYMOUS_GTID_EVENT
	PREVIOUS_GTIDS_EVENT
	TRANSACTION_CONTEXT_EVENT
	VIEW_CHANGE_EVENT
	XA_PREPARE_LOG_EVENT
)
