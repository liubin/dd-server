package types

import (
	"encoding/json"
)

// Transferred between sinks.
type MetricPayload struct {
	Metrics []*Metric `json:"metrics"`

	encoded []byte
	err     error
}

type Metric struct {
	Metric    string            `json:"metric"`
	Timestamp int64             `json:"timestamp"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type Process struct {
	Process string `json:"process"`
	User    string `json:"user"`
}

type ProcessStruct struct {
	ApiKey    string        `json:"apiKey"`
	Host      string        `json:"host"`
	Processes []interface{} `json:"processes"`
}

type Event struct {
	Name  string                 `json:"name"`
	Props map[string]interface{} `json:"props"`
}

type ServiceCheckBasic struct {
	Check     string  `json:"check"`
	Timestamp float64 `json:"timestamp"`
	Status    int     `json:"status"`
	Id        int     `json:"id"`
	Message   string  `json:"message"`
}

type ServiceCheckInput struct {
	ServiceCheckBasic `mapstructure:",squash"`
	Tags              []string `json:"tags"`
}

type ServiceCheckOutput struct {
	ServiceCheckBasic `mapstructure:",squash"`
	Tags              map[string]string `json:"tags"`
}

type RequestBody struct {
	ApiKey              string                   `json:"apiKey"`
	InternalHostname    string                   `json:"internalHostname"`
	CpuIdle             float64                  `json:"cpuIdle"`
	SystemLoad1         float64                  `json:"system.load.1"`
	SystemLoad5         float64                  `json:"system.load.5"`
	SystemLoad15        float64                  `json:"system.load.15"`
	CollectionTimestamp float64                  `json:"collection_timestamp"`
	Metrics             []interface{}            `json:"metrics"`
	Gohai               string                   `json:"gohai"`
	Processes           ProcessStruct            `json:"processes"`
	Events              map[string][]interface{} `json:"events"`
	AgentChecks         []interface{}            `json:"agent_checks"`
	ServiceChecks       []interface{}            `json:"service_checks"`
}

func (ml *MetricPayload) ensureEncoded() {
	if ml.encoded == nil && ml.err == nil {
		ml.encoded, ml.err = json.Marshal(ml)
	}
}

func (ml *MetricPayload) Length() int {
	ml.ensureEncoded()
	return len(ml.encoded)
}

func (ml *MetricPayload) Encode() ([]byte, error) {
	ml.ensureEncoded()
	return ml.encoded, ml.err
}
