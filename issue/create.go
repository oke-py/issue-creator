package issue

import (
	"fmt"
	"net/url"
)

// CreateBody returns RequestBody for POST /repos/kubernetes/website/issues
func CreateBody(path string, gistURL string) string {
	return fmt.Sprintf(`**This is a Feature Request**

<!-- Please only use this template for submitting feature/enhancement requests -->
<!-- See https://kubernetes.io/docs/contribute/start/ for guidance on writing an actionable issue description. -->

**What would you like to be added**
Update %s to follow v1.17 of the original (English) text.

**Why is this needed**
content/ja/%s is outdated.

**Comments**
/language ja
/good-first-issue

File to update
https://github.com/kubernetes/website/blob/dev-1.17-ja.2/content/ja/%s

Original
https://github.com/kubernetes/website/blob/release-1.17/content/en/%s

diff between translated and v1.17
%s

**Note**
Currently, we use dev-1.17-ja.2 branch. Please open a PR targeting the branch.

ref
[How To Contribute](https://kubernetes-docs-ja.kibe.la/shared/entries/c5878aa5-ad1f-4f29-a5bb-25853cbc14ec)
[翻訳スタイルガイド](https://kubernetes-docs-ja.kibe.la/shared/entries/5efe4fa7-d2a1-4a1d-8bc3-ce7ccdc064a6)

If you have a question, feel free to ask us at slack.k8s.io #kubernetes-docs-ja channel.`,
		path, path, path, path, gistURL,
	)
}

// Create creates a new issue and return the URL
func Create(path string, gistURL string) string {
	return fmt.Sprintf("https://github.com/kubernetes/website/issues/new?title=%s&body=%s",
		url.QueryEscape(fmt.Sprintf("ja: Make %s follow v1.17 of the original text", path)),
		url.QueryEscape(CreateBody(path, gistURL)),
	)
}
