package main

import (
	"fmt"
	"job_submitter/session"
)

func main(){
	flicker := session.Flicker()
	// sourceFilePath := "..\\archive\\data.csv"
	sourceFilePath := "path_to_amazon_reviews_dataset.csv"
	fdf := flicker.Read("local").Format("csv").Option("path", sourceFilePath).Load()
	fdf = fdf.Filter("fdf.star_rating >= 4")
	aggCondition := map[string]interface{}{"star_rating": "avg", "helpful_votes": "sum"}
	groupByColumn := []string{"product_title"}
	fdf = fdf.GroupBy(groupByColumn, aggCondition)
	fdf = fdf.WithColumnRenamed("AVG(star_rating)", "avg_star_rating")
	fdf = fdf.WithColumnRenamed("SUM(helpful_votes)", "total_helpful_votes")
	jobStatus, _ := fdf.Execute()
	fmt.Println("Job successful:", jobStatus)
	fdf.Write.Mode("overwrite").Csv("output.csv")
	flicker.Stop()
}