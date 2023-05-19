package session

import (
    "fmt"
    "github.com/go-gota/gota/dataframe"
    "job_submitter/job"
)

// flickerdf represents a dataframe in the Flicker framework.
// It contains the actual dataframe, source information, a task set,
// and a flickerdfWriter instance for writing operations.
type flickerdf struct {
    df           dataframe.DataFrame
    sourceInfo   map[string]string
    taskSet      []map[string]interface{}
    Write        flickerdfWriter
}

// Take prints the first row of the dataframe.
func (fdf flickerdf) Take() {
    fmt.Println(fdf.df.Subset([]int{0}))
}

// Filter adds a filter task to the task set.
// It takes a condition string and returns the modified flickerdf.
func (fdf flickerdf) Filter(condition string) flickerdf {
    existingLen := len(fdf.taskSet)
    newItem := map[string]interface{}{
        "task_count": existingLen + 1,
        "task_name":  "Filter",
        "task_args": map[string]interface{}{
            "condition": condition,
        },
    }
    fdf.taskSet = append(fdf.taskSet, newItem)
    
    return fdf
}

// GroupBy adds a group by task to the task set.
// It takes column names and an aggregate map as arguments.
// It returns the modified flickerdf.
func (fdf flickerdf) GroupBy(columnNames []string, aggregate map[string]interface{}) flickerdf {
    existingLen := len(fdf.taskSet)
    newItem := map[string]interface{}{
        "task_count": existingLen + 1,
        "task_name":  "GroupBy",
        "task_args": map[string]interface{}{
            "group_by_columns": columnNames,
            "aggregate":        aggregate,
        },
    }
    fdf.taskSet = append(fdf.taskSet, newItem)
    
    return fdf
}

// WithColumnRenamed adds a column rename task to the task set.
// It takes the old and new column names as arguments.
// It returns the modified flickerdf.
func (fdf flickerdf) WithColumnRenamed(oldColumnName string, newColumnName string) flickerdf {
    existingLen := len(fdf.taskSet)
    newItem := map[string]interface{}{
        "task_count": existingLen + 1,
        "task_name":  "GroupBy",
        "task_args": map[string]interface{}{
            oldColumnName: oldColumnName,
            newColumnName: newColumnName,
        },
    }
    fdf.taskSet = append(fdf.taskSet, newItem)
    
    return fdf
}

// PrintTaskSet prints the task set of the flickerdf.
func (fdf flickerdf) PrintTaskSet() {
    for _, value := range fdf.taskSet {
        for key, val := range value {
            fmt.Println(key, val)
        }
    }
}

// Execute converts the flickerdf to a job JSON file, initializes a job runner,
// and executes the job. It returns a boolean value indicating success and an error, if any.
func (fdf flickerdf) Execute() (bool, error) {
    jobJSONPath := ConvertJobToJSON(fdf)
    jobRunner := job.Initialize(jobJSONPath)
    jobRunner.Execute()
    return true, nil
}
