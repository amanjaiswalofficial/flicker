package job

import (
	"os"
	"fmt"
	"encoding/json"	
	"io/ioutil"
	"github.com/go-gota/gota/dataframe"
)

// jobController represents the controller for executing a job.
type jobController struct {
	jobJSONPath      string              // Path to the job JSON file
	SourceInfo       map[string]string   // Information about the data source
	TaskSet          []map[string]interface{} // List of tasks to be executed
	DestInfo         map[string]string   // Information about the destination
	executorMetadata map[string]string   // Metadata about the executor
}

// getRowsPerExecutor calculates the number of rows per executor based on the total row count and the number of executors.
// It returns a slice containing the number of rows per executor.
func getRowsPerExecutor(rowCount int, numExecutors int) []int {
	quotient := rowCount / numExecutors
	remainder := rowCount % numExecutors

	values := make([]int, numExecutors)

	for i := 0; i < numExecutors-1; i++ {
		values[i] = quotient
	}
	values[numExecutors-1] = quotient + remainder

	return values
}

// Execute executes the job by reading the job details from the JSON file and performing the necessary operations.
func (jctrlr jobController) Execute() {
	// Read the job JSON file
	file, err := ioutil.ReadFile(jctrlr.jobJSONPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into the job controller
	err = json.Unmarshal(file, &jctrlr)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}
	
	// Open the data source file
	f, err := os.Open(jctrlr.SourceInfo["path"])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the data into a dataframe
	df := dataframe.ReadCSV(f)
	totalRecords := len(df.Records())
	
	executors := getExecutorConfiguration()
	availableNodes := len(executors.Config)
	rowsPerExecutor := getRowsPerExecutor(totalRecords, availableNodes)
	executorsRef := &executors
	executorsRef.attachSourceInformation(jctrlr.SourceInfo["path"], rowsPerExecutor)
	/*
		Remaining code coming soon....
		3. Keep waiting for incoming message from the broker, sent by executor node
		4. Perform the required message forwarding from one executor to other executors as required
		5. Until receiving a message of job completion, continue doing steps 3-4
		6. Once receiving a message that states completion of the job, stop listening
		7. Construct the output and return it back to the user
	*/
}

// Initialize initializes the job controller by setting the job JSON path.
// It returns an instance of the job controller.
func Initialize(jobJSONPath string) jobController {
	fmt.Println("Reading job from", jobJSONPath)
	jctrlr := jobController{}
	jctrlr.jobJSONPath = jobJSONPath
	return jctrlr
}
