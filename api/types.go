package api

type Metric struct {
	Name      string            `json:"name"`
	Timestamp int64             `json:"timestamp"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type RequestBody struct {
	ApiKey           string        `json:"apiKey"`
	InternalHostname string        `json:"internalHostname"`
	CpuIdle          float64       `json:"cpuIdle"`
	SystemLoad1      float64       `json:"system.load.1"`
	SystemLoad5      float64       `json:"system.load.5"`
	SystemLoad15     float64       `json:"system.load.15"`
	Metrics          []interface{} `json:"metrics"`
	Gohai            string        `json:"gohai"`
}
