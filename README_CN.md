# Plumber MySQL OR MariaDB CDC 

### TO => Kafka or NSQ or RabbitMQ or MongoDB

### 开发进度
- [x] 基础设施
- [x] Kafka
- [ ] NSQ
- [ ] RabbitMQ
- [ ] MongoDB


### 开始CDC监听
POST: `http://127.0.0.1:8089/new_monitor`
Body:
```json 

# 监听当前数据库所有DB和Table
{
    "engine":  "MySQL",
    "mq_engine": "Kafka",
    "cdc_start_timestamp": 0,
    "db_config": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "root"
    },
    "kafka_config": {
        "enable_sasl": false,
        "brokers": ["127.0.0.1:9082"],
        "topic": "test1"
    }
}

# 监听当前数据库特定DB和Table
{
    "engine":  "MySQL",
    "mq_engine": "Kafka",
    "cdc_start_timestamp": 0,
    "db_config": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "root",
        "db_name": "xxx",    # 可选, 监听指定DB
        "table_name": "xxx"  # 可选, 监听指定Table
    },
    "kafka_config": {
        "enable_sasl": false,
        "brokers": ["127.0.0.1:9082"],
        "topic": "test1"
    }
}
```

### 获取进行中的CDC
GET: `http://127.0.0.1:8089/all_monitor`

### 结束某个CDC的监听
POST: `http://127.0.0.1:8089/stop_monitor/:task_id`

### 开发依赖
- mysql-client  `sudo apt-get install mysql-client`