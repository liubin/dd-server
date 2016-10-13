package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"net/http"
)

func IntakeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := DecodeRequestBody(r)

	if err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}
	decoder := json.NewDecoder(bytes.NewReader(body))

	var req RequestBody
	if err := decoder.Decode(&req); err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}

	metrics := make([]*Metric, 0)
	for _, metric := range req.Metrics {
		if m, e := decodeMetric(metric); e != nil {
		} else {
			//b, _ := json.Marshal(m)
			//fmt.Println(string(b))
			metrics = append(metrics, m)
		}
	}

	//processes := ParseProcesses(&req.Processes)
	//fmt.Println(processes)

	//events := ParseEvents(&req)
	//fmt.Println(events)

	//agentCheck := ParseAgentChecks(&req)
	//fmt.Println(agentCheck)

	serviceCheck, _ := ParseServiceChecks(&req)
	fmt.Println(serviceCheck)
}

func decodeMetric(v interface{}) (*Metric, error) {
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

	return &Metric{
		Name:      name,
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
