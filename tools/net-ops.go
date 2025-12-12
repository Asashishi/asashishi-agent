package tools

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var SkipTags map[string]bool = map[string]bool{
	"script": true, "style": true, "noscript": true,
	"head": true, "meta": true, "links": true,
	"iframe": true, "object": true, "embed": true, "applet": true, "img": true, "image": true,
	"svg": true, "canvas": true, "ins": true, "video": true, "audio": true,
	"form": true, "input": true, "button": true,
	"nav": true, "footer": true, "aside": true, "header": true,
}

func WebContentSearch(url string) string {
	var (
		err         error
		body        []byte
		bodyStr     string
		bodyRoot    *html.Node
		resp        *http.Response
		textBuilder strings.Builder
		gText       func(node *html.Node)
	)
	if resp, err = http.Get(url); err != nil {
		return ""
	}
	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return ""
	}
	bodyStr = string(body)
	if bodyRoot, err = html.Parse(strings.NewReader(bodyStr)); err != nil {
		return ""
	}
	gText = func(node *html.Node) {
		var (
			text  string
			child *html.Node
		)
		switch node.Type {
		case html.ElementNode:
			if SkipTags[node.Data] {
				return
			}
		case html.TextNode:
			text = strings.TrimSpace(node.Data)
			if len(text) > 0 {
				textBuilder.WriteString(text + "\n")
			}
		}
		for child = node.FirstChild; child != nil; child = child.NextSibling {
			gText(child)
		}
	}
	gText(bodyRoot)
	return textBuilder.String()
}
