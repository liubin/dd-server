package api

import (
	"dd-server/types"
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
func ParseServiceChecks(req *types.RequestBody) ([]types.ServiceCheckOutput, error) {

	var inputs []types.ServiceCheckInput
	if err := mapstructure.Decode(req.ServiceChecks, &inputs); err != nil {
		return nil, err
	}

	outputs := make([]types.ServiceCheckOutput, 0)

	for _, input := range inputs {
		output := types.ServiceCheckOutput{
			ServiceCheckBasic: input.ServiceCheckBasic,
			Tags:              ParseStringTag(input.Tags)}
		outputs = append(outputs, output)
	}
	return outputs, nil
}
