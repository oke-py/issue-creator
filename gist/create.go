package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ResponseBody is a part of JSON response body
type ResponseBody struct {
	URL string `json:"html_url"`
}

// RequestBody is request body to create a new gist
type RequestBody struct {
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Files       struct {
		Diff struct {
			Content string `json:"content"`
		} `json:"en.diff"`
	} `json:"files"`
}

// CreateRequestBody returns RequestBody for POST /gist
func CreateRequestBody(before string, after string, content string) RequestBody {
	gist := &RequestBody{
		Description: fmt.Sprintf("git diff %s %s <file>", before, after),
		Public:      true,
		Files: struct {
			Diff struct {
				Content string "json:\"content\""
			} "json:\"en.diff\""
		}{
			Diff: struct {
				Content string "json:\"content\""
			}{
				Content: content,
			},
		},
	}
	return *gist
}

// Create creates a new gist and return the URL
func Create() string {
	url := "https://api.github.com/gists"

	client := &http.Client{}
	token := os.Getenv("ISSUE_CREATOR_TOKEN")

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	gist := CreateRequestBody("upstream/release-1.15", "upstream/release-1.17", string(stdin))

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

	var r ResponseBody
	if err := json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	return r.URL
}
