# Plumber MySQL OR MariaDB CDC 

### TO => Kafka or NSQ or RabbitMQ or MongoDB

### 开发进度
- [x] 基础设施
- [x] Kafka
- [x] NSQ
- [x] RabbitMQ
- [ ] MongoDB

### 部署
- `docker run --name plumber -p 8089:8089 dollarkiller/plumber:latest`
- docker-compose
- k8s

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

### CDC MESSAGE:
``` 
{
    "table":{
        "db_name":"cpx",
        "table_name":"kvs",
        "columns":[
            {
                "name":"key",
                "type":5,
                "collation":"utf8mb4_0900_ai_ci",
                "raw_type":"varchar(255)",
                "is_auto":false,
                "is_unsigned":false,
                "is_virtual":false,
                "enum_values":null,
                "set_values":null,
                "fixed_size":0,
                "max_size":255
            },
            {
                "name":"value",
                "type":5,
                "collation":"utf8mb4_0900_ai_ci",
                "raw_type":"varchar(255)",
                "is_auto":false,
                "is_unsigned":false,
                "is_virtual":false,
                "enum_values":null,
                "set_values":null,
                "fixed_size":0,
                "max_size":255
            }
        ],
        "indexes":[
            {
                "name":"PRIMARY",
                "columns":[
                    "key"
                ],
                "cardinality":[
                    1
                ]
            },
            {
                "name":"idx_kvs_value",
                "columns":[
                    "value"
                ],
                "cardinality":[
                    1
                ]
            }
        ],
        "pk_columns":[
            0
        ],
        "unsigned_columns":null
    },
    "action":"update",   #  update, insert, delete
    "rows":"[{"key":"sd","value":"sd"},{"key":"sd","value":"000"}]",  # 如果是insert , rows 数量为1， rows[0]为当前插入的数据.
                                                                        如果是update，rows数量为2，rows[0] 为旧数据 rows[1]为新数据
                                                                        如果是delete, rows 数量为1， rows[0]为删除的数据
    "header":{
        "timestamp":1614762803,
        "event_type":31,
        "server_id":1,
        "event_size":55,
        "log_pos":4939,
        "flags":0
    }
}
```

### 使用Kafka作为MQ
POST BODY:
```json
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
```

### 使用NSQ作为MQ
POST BODY:
```json
{
    "engine":  "MySQL",
    "mq_engine": "NSQ",
    "cdc_start_timestamp": 0,
    "db_config": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "root"
    },
    "nsq_config": {
        "addr": ["127.0.0.1:4150"],
        "topic": "test1"
    }
}
```

### 使用RabbitMQ作为MQ
POST BODY:
```json
{
    "engine":  "MySQL",
    "mq_engine": "RabbitMQ",
    "cdc_start_timestamp": 0,
    "db_config": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "root"
    },
    "rabbit_mq_config": {
        "uri": "amqp://admin:admin@127.0.0.1:5672/",
        "queue": "test1"
    }
}
```

### 开发依赖
- mysql-client  `sudo apt-get install mysql-client`

#### seo
- MySQL cdc Kafka
- MySQL cdc NSQ
- MySQL cdc RabbitMQ
- MySQL cdc MongoDB
- MariaDB cdc Kafka
- MariaDB cdc NSQ
- MariaDB cdc RabbitMQ
- MariaDB cdc MongoDB
