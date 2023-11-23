package database

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSONFromFile reads JSON data from a file and returns a slice of maps.
func ReadJSONFromFile(filename string)  ([]map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
