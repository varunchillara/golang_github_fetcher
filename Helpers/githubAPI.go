package Helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Users []struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:owner`
	HtmlUrl  string `json:"html_url"`
	Watchers int    `json:"watchers"`
}

func GithubAPI(user string) Users {
	resp, err := http.Get("https://api.github.com/users/" + user + "/repos")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	users := Users{}
	json.Unmarshal(body, &users)
	return users
}
