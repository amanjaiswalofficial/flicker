package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/health", GetHandler)
    r.POST("/submit", SubmitJob)

    return r
}