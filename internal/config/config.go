package config

var CONF = &baseConfig{}

type baseConfig struct {
	Engine   Engine
	MQEngine MQEngine
	DBConfig DBConfig

	KafkaConfig    *KafkaConfig
	NSQConfig      *NSQConfig
	RabbitMQConfig *RabbitMQConfig
}

type Engine string

var (
	MySQL   Engine = "MySQL"
	MariaDB Engine = "MariaDB"
)

type MQEngine string

var (
	Kafka    MQEngine = "Kafka"
	NSQ      MQEngine = "NSQ"
	RabbitMQ Engine   = "RabbitMQ"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type KafkaConfig struct {
	Brokers    []string
	User       string
	Password   string
	Topic      string
	EnableSASL bool
}

type NSQConfig struct {
}

type RabbitMQConfig struct {
}
