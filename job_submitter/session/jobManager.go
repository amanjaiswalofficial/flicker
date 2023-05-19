package session

import (
	"os"
	"encoding/json"
)

// createJSONObj creates a JSON object by combining the source information and taskSet
// from the given flickerdf.
func createJSONObj(fdf flickerdf) map[string]interface{} {
	// Create an empty JSON object
	jsonObj := make(map[string]interface{})
	// Add the source information and taskSet to the JSON object
	jsonObj["sourceInfo"] = fdf.sourceInfo
	jsonObj["taskSet"] = fdf.taskSet

	return jsonObj
}

// ConvertJobToJSON converts the given flickerdf to JSON and saves it to a file.
// It returns the path of the saved JSON file.
func ConvertJobToJSON(fdf flickerdf) (jobDataPath string) {
	// Create a JSON object from the flickerdf
	jsonObj := createJSONObj(fdf)
	// Define the path for the JSON file
	// filename := "..\\archive\\job.json"
	filename := "path_to_write_job_json"
	// Marshal the JSON object to JSON data
	if jsonData, err := json.Marshal(jsonObj); err != nil {
		panic(err)
	} else {
		// Create the JSON file
		file, err := os.Create(filename)
		if err != nil {
			print(err)
			return ""
		}
		defer file.Close()

		// Create a JSON encoder and set the indentation
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		jsonDataInString := json.RawMessage(jsonData)

		// Encode the JSON data to the file
		if err := encoder.Encode(jsonDataInString); err != nil {
			print(err)
			return ""
		}
	}

	return filename
}
