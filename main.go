package main

import (
	"fmt"
	"os"

	"github.com/oke-py/issue-creator/gist"
	"github.com/oke-py/issue-creator/issue"
)

func main() {
	gistURL := gist.Create()

	issueBody := issue.CreateBody(os.Getenv("IC_FILEPATH"), gistURL)
	fmt.Printf("%s\n", issueBody)
}
