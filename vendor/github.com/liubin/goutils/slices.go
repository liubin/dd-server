package goutils

import (
	"fmt"
	"strings"
)

// Convert a slice like {"a=b", "c=d", } to map {"a"=>"b", "c"=>"d"}
func SliceToMap(slices []string) (map[string]string, error) {
	result := map[string]string{}
	for _, v := range slices {
		pair := strings.Split(v, "=")
		if len(pair) != 2 {
			return nil, fmt.Errorf("Format error: %s", v)
		}
		result[pair[0]] = pair[1]
	}
	return result, nil
}
