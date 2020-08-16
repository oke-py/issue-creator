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
Update %s to follow v1.18 of the original (English) text.

**Why is this needed**
content/ja/%s is outdated.

**Comments**
/language ja
/good-first-issue

File to update
https://github.com/kubernetes/website/blob/dev-1.18-ja.1/content/ja/%s

Original
https://github.com/kubernetes/website/blob/fb6364d/content/en/%s

diff between translated and v1.18
%s

**Note**
Currently, we use **dev-1.18-ja.1** branch. Please open a PR targeting the branch.

ref
https://kubernetes.io/ja/docs/contribute/localization/

If you have a question, feel free to ask us at slack.k8s.io #kubernetes-docs-ja channel.`,
		path, path, path, path, gistURL,
	)
}

// Create creates a new issue and return the URL
func Create(path string, gistURL string) string {
	return fmt.Sprintf("https://github.com/kubernetes/website/issues/new?title=%s&body=%s",
		url.QueryEscape(fmt.Sprintf("ja: Make %s follow v1.18 of the original text", path)),
		url.QueryEscape(CreateBody(path, gistURL)),
	)
}
