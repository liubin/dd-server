# dd-server

Receive metrics from dd-agent ( datadog agent ), and save other store ( like Opentsdb ).

# How to run

## Opentsdb

```
./dd-server server --port 8800 --sink-driver opentsdb  \
    --sink-opts endpoint=http://192.168.33.10:4242
```

## Kafka

```
./dd-server server --port 8800 \
    --sink-driver kafka -sink-opts brokers=192.168.33.10:9092 \
    --sink-opts topic=metrics \
    -sink-opts acks=local \
    -sink-opts compression=none \
    -sink-opts flush-frequency=500
```

options:

- acks: default none
- compression: default none
- flush-frequency: default 500ms
