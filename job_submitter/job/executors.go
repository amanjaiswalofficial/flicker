// Package job provides functionality related to job configuration and execution.

package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Executor represents an executor configuration.
type Executor struct {
	ID         string `json:"id"`
	IPAddress  string `json:"ipAddress"`
	Hostname   string `json:"hostname"`
	Status     string `json:"status"`
	RowCount   int    // Number of rows
	SourcePath string // Source file path
}

// Executors represents a collection of executor configurations.
type Executors struct {
	Config []Executor `json:"config"`
}

// getExecutorConfiguration reads the executor configuration from a JSON file.
// It returns the Executors object containing the configuration.
func getExecutorConfiguration() (exc Executors) {
	// Read the JSON file
	fileContent, err := ioutil.ReadFile("..\\archive\\config.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Unmarshal the JSON data into a Config struct
	var executorsConfig Executors
	err = json.Unmarshal(fileContent, &executorsConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	return executorsConfig
}

// updateRowCount updates the RowCount field of the Executor.
func (exc *Executor) updateRowCount(rowCount int) {
	exc.RowCount = rowCount
}

// updateSourcePath updates the SourcePath field of the Executor.
func (exc *Executor) updateSourcePath(sourcePath string) {
	exc.SourcePath = sourcePath
}

// attachSourceInformation attaches source information to each executor in the Executors object.
// It updates the RowCount and SourcePath fields of each Executor.
func (exc *Executors) attachSourceInformation(sourcePath string, rowsToRead []int) {
	for i := range exc.Config {
		exc.Config[i].updateRowCount(rowsToRead[i])
		exc.Config[i].updateSourcePath(sourcePath)
	}
}

// printExecutorConfiguration prints the executor configuration to the console.
func (exc Executors) printExecutorConfiguration() {
	// Access and use the data from the Config struct
	for _, executor := range exc.Config {
		fmt.Printf("ID: %s, IPAddress: %s, Hostname: %s, Status: %s Row Count: %d Source Path: %s\n",
			executor.ID, executor.IPAddress, executor.Hostname, executor.Status, executor.RowCount, executor.SourcePath)
	}
}
