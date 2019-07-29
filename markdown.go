package main

import (
	"github.com/shurcooL/github_flavored_markdown"
)

func MarkDown(input []byte) []byte {
	html := github_flavored_markdown.Markdown(input)

	return html
}
