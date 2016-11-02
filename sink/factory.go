package sink

import (
	"dd-server/types"
	"fmt"
)

type SinkDriver interface {
	Write(*types.MetricPayload) error
}

var sinkDriver SinkDriver

func InitSinkDriver(opts map[string]string) error {
	driver, ok := opts["sink-driver"]

	if !ok {
		return fmt.Errorf("sink driver not provided")
	}

	var err error

	if driver == "opentsdb" {
		sinkDriver, err = NewOpentsdbSink(opts)
	} else if driver == "kafka" {
		sinkDriver, err = NewKafkaSink(opts)
	} else {
		err = fmt.Errorf("driver [%s] not support!", driver)
	}

	return err
}

func GetSinkDriver() SinkDriver {
	return sinkDriver
}
