package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Repretsts anchor tag in html documents
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	document, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(document, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
