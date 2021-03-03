# Plumber MySQL OR MariaDB CDC 

### TO => Kafka or NSQ or RabbitMQ or MongoDB

### Project Progress
- [x] Infrastructure
- [x] Kafka
- [x] NSQ
- [x] RabbitMQ
- [ ] MongoDB

### Deploy
- `docker run --name plumber -p 8089:8089 dollarkiller/plumber:latest`
- docker-compose
- k8s

### Establishing CDC Listening
POST: `http://127.0.0.1:8089/new_monitor`
Body:
```json

# Listening to all DBs and tables
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

# Listening to the current database specifying DB and Table
{
    "engine":  "MySQL",
    "mq_engine": "Kafka",
    "cdc_start_timestamp": 0,
    "db_config": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "root",
        "db_name": "xxx",    # Optionally listens to a specific database
        "table_name": "xxx"  # Optionally, listen to the specified table of the specified database
    },
    "kafka_config": {
        "enable_sasl": false,
        "brokers": ["127.0.0.1:9082"],
        "topic": "test1"
    }
}
```

### Get the CDC in progress
GET: `http://127.0.0.1:8089/all_monitor`

### Ending the listening of a CDC
POST: `http://127.0.0.1:8089/stop_monitor/:task_id`

### CDC MESSAGE:
```json
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
    "rows":"[{"key":"sd","value":"sd"},{"key":"sd","value":"000"}]",  # If it is insert , the number of rows is 1 and rows[0] is the current inserted data.
                                                                        If it is update, the number of rows is 2, rows[0] is the old data rows[1] is the new data
                                                                        If it is delete, the number of rows is 1 and rows[0] is the deleted data.
    "original_row": [["sd", "sd"], ["sd", "000"]], # original row
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

### Use Kafka
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

### Use NSQ
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

### Use RabbitMQ
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

### Dev Rely
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