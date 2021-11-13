package Controllers

import (
	"github.com/gin-gonic/gin"
)

func GetRepos(c *gin.Context) {
	c.String(200, "inside get repos controller")
}
