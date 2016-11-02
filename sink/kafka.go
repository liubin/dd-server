package sink

import (
	"dd-server/types"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"strings"
	"time"
)

type KafkaSink struct {
	Brokers  string
	Topic    string
	Producer sarama.AsyncProducer
}

func NewKafkaSink(opts map[string]string) (SinkDriver, error) {

	brokers, ok1 := opts["brokers"]
	topic, ok2 := opts["topic"]

	var acks sarama.RequiredAcks
	var compression sarama.CompressionCodec
	var flushFrequency int64

	if s, ok := opts["acks"]; ok {
		if s == "local" {
			acks = sarama.WaitForLocal
		} else if s == "all" {
			acks = sarama.WaitForAll
		}
	} else {
		acks = sarama.NoResponse
	}

	if s, ok := opts["compression"]; ok {
		if s == "snappy" {
			compression = sarama.CompressionSnappy
		} else if s == "gzip" {
			compression = sarama.CompressionGZIP
		}
	} else {
		compression = sarama.CompressionNone
	}

	if s, ok := opts["flush-frequency"]; ok {
		if flushFrequency, _ = strconv.ParseInt(s, 10, 64); flushFrequency < 1 {
			flushFrequency = 500
		}
	} else {
		flushFrequency = 500
	}

	if ok1 && ok2 && brokers != "" && topic != "" {
		brokerList := strings.Split(brokers, ",")

		config := sarama.NewConfig()

		// Only wait for the leader to ack
		config.Producer.RequiredAcks = acks
		// Compress messages
		config.Producer.Compression = compression

		// Flush batches every 500ms
		config.Producer.Flush.Frequency = time.Duration(flushFrequency) * time.Millisecond

		producer, err := sarama.NewAsyncProducer(brokerList, config)

		if err != nil {
			log.Fatalln("Failed to start Sarama producer:", err)
			return nil, err
		}
		return &KafkaSink{
			Brokers:  brokers,
			Topic:    topic,
			Producer: producer,
		}, nil
	} else {
		return nil, fmt.Errorf("brokers or topic not provided!")
	}
}

func (kafka *KafkaSink) Write(metrics *types.MetricPayload) error {
	kafka.Producer.Input() <- &sarama.ProducerMessage{
		Topic: kafka.Topic,
		Key:   sarama.StringEncoder(metrics.Metrics[0].Metric), // TODO set to what?
		Value: metrics,
	}

	return nil // TODO how to get error?
}
