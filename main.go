package main

import (
	"fmt"

	"github.com/oke-py/issue-creator/gist"
)

func main() {
	gistURL := gist.Create()
	fmt.Printf("Created: %s\n", gistURL)
}
