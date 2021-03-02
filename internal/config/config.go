package config

import "github.com/spf13/viper"

var CONF = &BaseConfig{}

type BaseConfig struct {
	Engine   Engine
	MQEngine MQEngine
	DBConfig DBConfig

	CDCStartTimestamp int64 // is 0 then real time data

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
	RabbitMQ MQEngine = "RabbitMQ"
)

func (m MQEngine) String() string {
	return string(m)
}

type DBConfig struct {
	Host     string
	Port     int64
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

// InitConfiguration ...
func InitConfiguration(configName string, configPaths []string, config interface{}) error {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AutomaticEnv()
	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	err := vp.Unmarshal(config)
	if err != nil {
		return err
	}

	return nil
}
