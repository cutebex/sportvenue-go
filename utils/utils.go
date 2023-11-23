package utils

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

// ReadJSONFromFile reads JSON data from a file and returns a slice of maps.
func ReadJSONFromFile(filename string) ([]map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read data from %s: %v", filename, err)
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON data: %v", err)
	}

	return result, nil
}
