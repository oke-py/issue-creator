package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GistResponse is a part of JSON response body
type GistResponse struct {
	URL string `json:"html_url"`
}

// GistRequestBody is request body to create a new gist
type GistRequestBody struct {
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Files       struct {
		Diff struct {
			Content string `json:"content"`
		} `json:"en.diff"`
	} `json:"files"`
}

func main() {
	url := "https://api.github.com/gists"

	client := &http.Client{}
	token := os.Getenv("ISSUE_CREATOR_TOKEN")

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	gist := &GistRequestBody{
		Description: "[beta] git diff upstream/release-1.17 upstream/release-1.18 <file>",
		Public:      true,
		Files: struct {
			Diff struct {
				Content string "json:\"content\""
			} "json:\"en.diff\""
		}{
			Diff: struct {
				Content string "json:\"content\""
			}{
				Content: string(stdin),
			},
		},
	}

	j, _ := json.Marshal(gist)
	b := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", url, b)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	req.Header.Add("Content-Type", "application/vnd.github.v3.text+json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var r GistResponse
	if err := json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("URL: %s\n", r.URL)
}
