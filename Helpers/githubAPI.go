package Helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Repos []struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:owner`
	HtmlUrl  string `json:"html_url"`
	Watchers int    `json:"watchers"`
}

func GithubAPI(user string) (Repos, error) {
	resp, err := http.Get("https://api.github.com/users/" + user + "/repos")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	repos := Repos{}
	json.Unmarshal(body, &repos)
	return repos, nil
}
