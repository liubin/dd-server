package sink

import (
	"bytes"
	"dd-server/types"
	"dd-server/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ElasticSearchSink struct {
	Endpoint  string
	IndexName string
	TypeName  string
	Path      string
}

func NewElasticSearchSink(opts map[string]string) (*ElasticSearchSink, error) {

	var ok bool
	var endpoint, indexName, typeName string
	endpoint, ok = opts["endpoint"]

	if !ok || endpoint == "" {
		return nil, fmt.Errorf("endpoint not provided!")
	}

	if !ok || indexName == "" {
		indexName = "ddserver"
	}

	if !ok || typeName == "" {
		typeName = "events"
	}
	path := fmt.Sprintf("/%s/%s/_bulk", indexName, typeName)
	return &ElasticSearchSink{
		Endpoint:  endpoint,
		IndexName: indexName,
		TypeName:  typeName,
		Path:      path,
	}, nil

}

func (es *ElasticSearchSink) WriteEvents(events []*types.Event) error {

	str := `{"index":{}}`
	var buffer bytes.Buffer
	for _, event := range events {
		eventStr, err := json.Marshal(event)
		if err == nil {
			buffer.WriteString(strings.Join([]string{str, string(eventStr)}, "\n"))
			buffer.WriteString("\n")
		}
	}
	resp, code, err := utils.POST(es.Endpoint, es.Path, buffer.String(), nil)
	if code == http.StatusOK && err == nil {
		return nil
	}

	log.Printf("ElasticSearchSink.WriteEvents response code %d", code)
	log.Printf("ElasticSearchSink.WriteEvents response error %s", err)
	log.Printf("ElasticSearchSink.WriteEvents response body %s", resp)
	return err
}
