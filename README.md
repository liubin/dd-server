# dd-server

Receive metrics from dd-agent ( datadog agent ), and save other store ( like Opentsdb ).

# How to run

## Opentsdb

```
./dd-server server --port 8800 --sink-driver opentsdb  \
    --sink-opts endpoint=http://192.168.33.10:4242 \
    --license-validator http://localhost:8081/api/v1/users
```

## Kafka

```
./dd-server server --port 8800 \
    --sink-driver kafka -sink-opts brokers=192.168.33.10:9092 \
    --sink-opts topic=metrics \
    --sink-opts acks=local \
    --sink-opts compression=none \
    --sink-opts flush-frequency=500 \
    --license-validator http://localhost:8081/api/v1/users
```

options:

- `acks`: default none
- `compression`: default none
- `flush-frequency`: default 500ms

## Save events to elasticsearch

### Create index and type


```
curl -XPUT 'localhost:9200/ddserver?pretty' -d @hack/events_mapping.json
```

### Run with elasticsearch

```
./dd-server server --port 8800 --sink-driver opentsdb  \
    --sink-opts endpoint=http://192.168.33.10:4242 \
    --license-validator http://localhost:8081/api/v1/users \
    --elasticsearch endpoint=http://192.168.33.10:9200 \
    [--es-index=ddserver --es-type=events]
```

Items:
* `--elasticsearch endpoint`: HTTP API endpoint.
* `--elasticsearch index`: index name, default `ddserver`
* `--elasticsearch events`: type name, defalut `events`


## Implement your license validator.

License validator is used to do token authentication, and use the result as tags.

The validator should:

- Receive a POST request with the body:

```
{"license": "license-code-here"}
```

- And send a response with(a type of map with the key and value all are string):

```
{
  "user": "12345",
  "any-tag": "tag-value"
}
```

The result from license validator will be added to metric's tags.

