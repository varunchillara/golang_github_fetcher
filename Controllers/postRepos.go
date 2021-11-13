package Controllers

import (
	"github.com/gin-gonic/gin"
)

func PostRepos(c *gin.Context) {
	c.String(200, "inside post repos controller")
}
