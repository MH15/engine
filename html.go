package main

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// Parse the html tree
func Parse(r io.Reader) Element {

	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	node := MakeTree(doc)

	return node

}

// MakeTree creates an Element Tree
func MakeTree(n *html.Node) Element {
	// if n.Type == html.ElementNode && n.Data == "a" {
	// 	for _, a := range n.Attr {
	// 		if a.Key == "href" {
	// 			fmt.Println(a.Val)
	// 			break
	// 		}
	// 	}
	// }

	root := MakeElementNode("a")
	root.SetData(n.Data)

	if n.Type == html.ElementNode {
		root.nodeType = ElementNode
		// fmt.Println("element node: '" + n.Data + "' " + n.DataAtom.String())
		for _, attribute := range n.Attr {
			root.attributes[attribute.Key] = attribute.Val
		}
	} else {
		if len(SpaceFieldsJoin(n.Data)) > 0 {
			root.nodeType = TextNode
		}
		// fmt.Println("text node: '" + n.Data + "' " + n.DataAtom.String())
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if len(SpaceFieldsJoin(c.Data)) > 0 {
			root.Add(MakeTree(c))
		}
	}

	return root
}

// SpaceFieldsJoin quickly removes all whitespace from a string
// @author: https://stackoverflow.com/a/32081891
func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}
