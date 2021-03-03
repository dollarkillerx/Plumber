package newsletter

type Engine string

var (
	MySQL   Engine = "MySQL"
	MariaDB Engine = "MariaDB"
)

type MQEngine string

var (
	Kafka    MQEngine = "Kafka"
	NSQ      MQEngine = "NSQ"
	RabbitMQ MQEngine = "RabbitMQ"
)

func (m MQEngine) String() string {
	return string(m)
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`

	DBName    *string `json:"db_name"`    // 可选
	TableName *string `json:"table_name"` // 可选
}

type KafkaConfig struct {
	EnableSASL bool     `json:"enable_sasl"`
	Brokers    []string `json:"brokers"`
	User       string   `json:"user"`
	Password   string   `json:"password"`
	Topic      string   `json:"topic"`
}

type NSQConfig struct {
	Addr  []string `json:"addr"`
	Topic string   `json:"topic"`
}

type RabbitMQConfig struct {
	Uri   string `json:"uri"` // "amqp://guest:guest@127.0.0.1:5672/"
	Queue string `json:"queue"`
}

type TaskConfig struct {
	Engine   Engine   `json:"engine"`
	MQEngine MQEngine `json:"mq_engine"`
	DBConfig DBConfig `json:"db_config"`

	CDCStartTimestamp int64 `json:"cdc_start_timestamp"` // is 0 then real time data

	KafkaConfig    *KafkaConfig    `json:"kafka_config"`
	NSQConfig      *NSQConfig      `json:"nsq_config"`
	RabbitMQConfig *RabbitMQConfig `json:"rabbit_mq_config"`
}

type TaskResponse struct {
	TaskID  string `json:"task_id"`
	Success bool   `json:"success"`
}
