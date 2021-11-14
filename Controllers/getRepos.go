package Controllers

import (
	"github.com/gin-gonic/gin"
)

func GetRepos(c *gin.Context) {
	c.JSON(200, "get")
}
