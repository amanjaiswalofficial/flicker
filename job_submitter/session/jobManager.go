package session

import (
	"os"
	"encoding/json"
)

func createJSONObj (fdf flickerdf) (map[string]interface{}) {
	// add both source information and taskSet
	jsonObj := make(map[string]interface{})
	jsonObj["sourceInfo"] = fdf.sourceInfo
	jsonObj["taskSet"] = fdf.taskSet

	return jsonObj
}

func ConvertJobToJSON (fdf flickerdf) (jobDataPath string) {
	
	jsonObj := createJSONObj(fdf)
	filename := "..\\archive\\job.json"
	if jsonData, err := json.Marshal(jsonObj); err != nil {
        panic(err)
    } else {
		file, err := os.Create(filename)
		if err != nil {
			print(err)
			return ""
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		jsonDataInString := json.RawMessage(jsonData)

		if err := encoder.Encode(jsonDataInString); err != nil {
			print(err)
			return ""
		}
    }

	

    return filename
}