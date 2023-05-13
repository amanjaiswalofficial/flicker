package session

import (
    "fmt"
    "github.com/go-gota/gota/dataframe"
)

type flickerdf struct {
    df dataframe.DataFrame
    taskSet []map[string]interface{}
    Write flickerdfWriter
}

func (fdf flickerdf) Take() {
    fmt.Println(fdf.df.Subset([]int{0}))
}

func (fdf flickerdf) Filter(condition string) flickerdf {
    existingLen := len(fdf.taskSet)
    newItem := map[string]interface{}{
        "task_count": existingLen + 1,
        "task_name": "Filter",
        "task_args": map[string]interface{}{
            "condition": condition,
        },
    }
    fdf.taskSet = append(fdf.taskSet, newItem)
    
    return fdf
}

func (fdf flickerdf) GroupBy(columnNames []string, aggregate map[string]interface{}) flickerdf {
    existingLen := len(fdf.taskSet)
    newItem := map[string]interface{}{
        "task_count": existingLen + 1,
        "task_name": "GroupBy",
        "task_args": map[string]interface{}{
            "group_by_columns": columnNames,
            "aggregate": aggregate,
        },
    }
    fdf.taskSet = append(fdf.taskSet, newItem)
    
    return fdf
}

func (fdf flickerdf) WithColumnRenamed(oldColumnName string, newColumnName string) flickerdf {
    existingLen := len(fdf.taskSet)
    newItem := map[string]interface{}{
        "task_count": existingLen + 1,
        "task_name": "GroupBy",
        "task_args": map[string]interface{}{
            oldColumnName: oldColumnName,
            newColumnName: newColumnName,
        },
    }
    fdf.taskSet = append(fdf.taskSet, newItem)
    
    return fdf
}

func (fdf flickerdf) Execute() (bool, error){
    /*
    Do the whole job execution
    */
    fmt.Println("Submitting job for execution")
    fmt.Println("Job finished post execution")
    return true, nil
}

func (fdf flickerdf) PrintTaskSet() {
    for _, value := range fdf.taskSet {
        for key, val := range value {
            fmt.Println(key, val)
        }
    }
}