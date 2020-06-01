package gist

import "testing"

func TestCreateRequestBody(t *testing.T) {
	gist1 := CreateRequestBody("upstream/release-1.17", "upstream/release-1.18", "Hello, World!")

	if gist1.Description != "git diff upstream/release-1.17 upstream/release-1.18 <file>" {
		t.Fatal("failed test")
	}

	gist2 := CreateRequestBody("eb7d44fa6360b8c0c53ebcd5b4eeb2a8fc69c62c", "upstream/release-1.18", "Hello, World!")
	if gist2.Description != "git diff eb7d44fa6360b8c0c53ebcd5b4eeb2a8fc69c62c upstream/release-1.18 <file>" {
		t.Fatal("failed test")
	}
}
