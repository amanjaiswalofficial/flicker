package api

import "github.com/gin-gonic/gin"

func GetHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Hello, World!"})
}