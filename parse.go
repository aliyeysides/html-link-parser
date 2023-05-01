package link

import (
	"io"
	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
  Href string
  Text string
}

// Parse will take in an HTML document and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
  doc, err := html.Parse(r)
  if err != nil {
    return nil, err
  }
  dfs(doc)
  return nil, nil
}

func dfs(n *html.Node) {
  if n.Type == html.ElementNode && n.Data == "a" {
    println(n.Data)
  }
  for c := n.FirstChild; c != nil; c = c.NextSibling {
    dfs(c)
  }
}
