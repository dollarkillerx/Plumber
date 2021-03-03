# Plumber MySQL OR MariaDB CDC 

### TO => Kafka or NSQ or RabbitMQ or MongoDB

### Project Progress
- [x] Infrastructure
- [x] Kafka
- [ ] NSQ
- [ ] RabbitMQ
- [ ] MongoDB

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
        "brokers": ["127.0.0.1:9082"]
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
        "brokers": ["127.0.0.1:9082"]
    }
}
```

### Get the CDC in progress
GET: `http://127.0.0.1:8089/all_monitor`

### Ending the listening of a CDC
POST: `http://127.0.0.1:8089/stop_monitor/:task_id`

### Dev Rely
- mysql-client  `sudo apt-get install mysql-client`