package api

import (
	"bytes"
	"dd-server/sink"
	"dd-server/types"
	"encoding/json"
	"fmt"
	"github.com/liubin/goutils"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

func IntakeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := DecodeRequestBody(r)

	if err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}
	decoder := json.NewDecoder(bytes.NewReader(body))

	var req types.RequestBody
	if err := decoder.Decode(&req); err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}

	licenseTags, err := validateLicense(req.ApiKey)
	if err != nil {
		fmt.Printf("licnese validate error %s", err.Error())
		return
	}
	log.Printf("licenseTags : %v", licenseTags)

	metrics := make([]*types.Metric, 0)
	for _, metric := range req.Metrics {
		if m, e := decodeMetric(metric); e != nil {
		} else {
			//b, _ := json.Marshal(m)
			//fmt.Println(string(b))
			m.Tags = goutils.MergeStringMap(m.Tags, licenseTags)
			metrics = append(metrics, m)
		}
	}

	metrics = addOtherMetrics(req, metrics)

	mSink := sink.GetSinkDriver()
	mSink.Write(&types.MetricPayload{Metrics: metrics})

	//processes := ParseProcesses(&req.Processes)
	//fmt.Println(processes)

	//events := ParseEvents(&req)
	//fmt.Println(events)

	//agentCheck := ParseAgentChecks(&req)
	//fmt.Println(agentCheck)

	serviceCheck, _ := ParseServiceChecks(&req)
	fmt.Println(serviceCheck)
}

func addOtherMetrics(req types.RequestBody, metrics []*types.Metric) []*types.Metric {
	tags := map[string]string{"host": req.InternalHostname}

	t := int64(req.CollectionTimestamp)
	m := &types.Metric{
		Metric:    "system.load.1",
		Timestamp: t,
		Value:     req.SystemLoad1,
		Tags:      tags,
	}

	metrics = append(metrics, m)

	m = &types.Metric{
		Metric:    "system.load.5",
		Timestamp: t,
		Value:     req.SystemLoad5,
		Tags:      tags,
	}

	metrics = append(metrics, m)

	m = &types.Metric{
		Metric:    "system.load.15",
		Timestamp: t,
		Value:     req.SystemLoad15,
		Tags:      tags,
	}

	metrics = append(metrics, m)

	return metrics
}

func decodeMetric(v interface{}) (*types.Metric, error) {
	ma, ok := v.([]interface{})

	if !ok {
		return nil, fmt.Errorf("Not a valid dd metric array")
	}
	l := len(ma)
	if l != 4 {
		return nil, fmt.Errorf("Not a valid dd metric array(length is not 4)")
	}

	name := cast.ToString(ma[0])
	ts := cast.ToInt64(ma[1])
	value := cast.ToFloat64(ma[2])

	var tags map[string]string

	if tagsMap, e := ma[3].(map[string]interface{}); e {
		tags = parseMapTag(tagsMap)
	}

	return &types.Metric{
		Metric:    name,
		Timestamp: ts,
		Value:     value,
		Tags:      tags,
	}, nil
}

func parseMapTag(sa map[string]interface{}) map[string]string {
	tags := make(map[string]string, 0)
	for k, v := range sa {
		if k != "tags" {
			sv, _ := v.(string)
			tags[k] = sv
			continue
		}
		va, ok := v.([]string)
		if !ok {
			continue
		}

		m := ParseStringTag(va)
		tags = MergeMap(tags, m)
	}
	return tags
}

func MergeMap(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}

	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}
