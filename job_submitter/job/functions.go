package job

import (
	"os"
	"fmt"
	"encoding/json"	
	"io/ioutil"
	"github.com/go-gota/gota/dataframe"
)

type jobController struct {
	jobJSONPath string
	SourceInfo map[string]string
    TaskSet []map[string]interface{}
	DestInfo map[string]string
	executorMetadata map[string]string
}

func getRowsPerExecutor(row_count int, num_executor int) []int {
	quotient := row_count / num_executor
	remainder := row_count % num_executor

	values := make([]int, num_executor)

	for i := 0; i < num_executor-1; i++ {
		values[i] = quotient
	}
	values[num_executor-1] = quotient + remainder

	return values
}

func (jctrlr jobController) Execute() {
	file, err := ioutil.ReadFile(jctrlr.jobJSONPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &jctrlr)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}
	
	f, err := os.Open(jctrlr.SourceInfo["path"])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	totalRecords := len(df.Records())
	availableNodes := 3 // TODO: get this from some configuration
	rowsPerExecutor := getRowsPerExecutor(totalRecords, availableNodes)
	fmt.Println(rowsPerExecutor)
}

func Initialize(jobJSONPath string) (jobController) {
	fmt.Println("Reading job from", jobJSONPath)
	jctrlr := jobController{}
	jctrlr.jobJSONPath = jobJSONPath
	return jctrlr
}

