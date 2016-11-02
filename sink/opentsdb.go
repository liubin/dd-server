package sink

import (
	"dd-server/types"
	"dd-server/utils"
	"fmt"
	"log"
	"net/http"
)

type OpentsdbSink struct {
	Endpoint string
}

func NewOpentsdbSink(opts map[string]string) (SinkDriver, error) {

	endpoint, ok := opts["endpoint"]

	if ok && endpoint != "" {
		return &OpentsdbSink{
			Endpoint: endpoint,
		}, nil
	} else {
		return nil, fmt.Errorf("endpoint not provided!")
	}
}

// Write points to opentsdb.
// http://opentsdb.net/docs/build/html/api_http/put.html
// Data points:
// [
//  {
//    "metric": "sys.cpu.nice",
//    "timestamp": 1346846400,
//    "value": 18,
//    "tags": {
//      "host": "web01",
//      "dc": "lga"
//    }
//  },
//  {
//    "metric": "sys.cpu.nice",
//    "timestamp": 1346846400,
//    "value": 9,
//    "tags": {
//      "host": "web02",
//      "dc": "lga"
//    }
//  }
// ]
func (opentsdb *OpentsdbSink) Write(metrics *types.MetricPayload) error {
	resp, code, err := utils.POST(opentsdb.Endpoint, "/api/put", metrics.Metrics, nil)
	if code == http.StatusNoContent {
		return nil
	}

	// Error response(?summary or ?details):
	//{
	//	"errors": [
	//	{
	//	"datapoint": {
	//	"metric": "sys.cpu.nice",
	//	"timestamp": 1365465600,
	//	"value": "NaN",
	//	"tags": {
	//	"host": "web01"
	//	}
	//	},
	//	"error": "Unable to parse value to a number"
	//	}
	//	],
	//	"failed": 1,
	//	"success": 0
	//}
	log.Printf("OpentsdbSink.Write response code %d", code)
	log.Printf("OpentsdbSink.Write response error %s", err)
	log.Printf("OpentsdbSink.Write response body %s", resp)
	return err
}
