// Package job provides utility functions for working with JSON files.

package job

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSONFile reads a JSON file and returns its content as a map[string]interface{} or a specified return type.
// If a returnType argument is provided, the JSON data is converted to the specified type.
// The function returns the JSON data and any error that occurred during the process.
func ReadJSONFile(filePath string, returnType ...interface{}) (interface{}, error) {
	// Read the JSON file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a map[string]interface{}
	var jsonData map[string]interface{}
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		return nil, err
	}

	// Check if a specific return type is requested
	if len(returnType) > 0 {
		// Return the JSON data converted to the specified return type
		return convertToType(jsonData, returnType[0]), nil
	}

	// Return the JSON data as is
	return jsonData, nil
}

// convertToType performs type conversion of the JSON data based on the requested return type.
func convertToType(data interface{}, returnType interface{}) interface{} {
	// Perform type conversion based on the requested return type
	switch returnType.(type) {
	case map[string]interface{}:
		// Return as map[string]interface{}
		return data
	case []interface{}:
		// Return as []interface{}
		return []interface{}{data}
	default:
		// Return as is (unknown type)
		return data
	}
}
