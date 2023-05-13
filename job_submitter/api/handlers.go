package api

import "github.com/gin-gonic/gin"
import "net/http"

func GetHandler(c *gin.Context) {
    c.JSON(200, gin.H{"health": "is ok"})
}

func SubmitJob(c *gin.Context) {
    var job Job
    if err := c.ShouldBindJSON(&job); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": job})
}