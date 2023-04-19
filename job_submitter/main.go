package main

import (
    "job_submitter/api"
)

func main() {
    router := api.SetupRouter()
    router.Run(":8080")
}