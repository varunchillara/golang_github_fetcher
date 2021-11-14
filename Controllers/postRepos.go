package Controllers

import (
	"encoding/json"
	"golang_github_fetcher/Helpers"

	"github.com/gin-gonic/gin"
)

type Body struct {
	User string `json:user`
}

func PostRepos(c *gin.Context) {
	body := Body{}
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(500, err)
	}
	json.Unmarshal(jsonData, &body)
	repos, err := Helpers.GithubAPI(body.User)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, repos)
}
