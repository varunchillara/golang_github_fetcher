package main

import (
	"golang_github_fetcher/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/repos", Controllers.GetRepos)
	r.POST("/repos", Controllers.PostRepos)

	r.Run(":3000")
}
