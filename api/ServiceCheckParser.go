package api

import (
	"github.com/mitchellh/mapstructure"
)

// Service checks
//"service_checks": [
//{
//"status": 3,
//"tags": null,
//"timestamp": 1476195870.704261,
//"check": "ntp.in_sync",
//"host_name": "bogon",
//"message": null,
//"id": 1
//},
func ParseServiceChecks(req *RequestBody) ([]ServiceCheckOutput, error) {

	var inputs []ServiceCheckInput
	if err := mapstructure.Decode(req.ServiceChecks, &inputs); err != nil {
		return nil, err
	}

	outputs := make([]ServiceCheckOutput, 0)

	for _, input := range inputs {
		output := ServiceCheckOutput{
			ServiceCheckBasic: input.ServiceCheckBasic,
			Tags:              ParseStringTag(input.Tags)}
		outputs = append(outputs, output)
	}
	return outputs, nil
}
