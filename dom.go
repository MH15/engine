// By Matt Hall 2018

// This code represents a DOM tree
package main

/**
 * Node
 */

// Node type
type Node interface {
	SetData(string)
	ToString() string
	NumberOfChildren() int
	Value(string)
	Add(Node)
	Child(int) Node
}

// NodeType const enum is better than just using strings
type NodeType uint8

const (
	// TextNode is for text
	TextNode NodeType = iota
	// ElementNode is for elements
	ElementNode
)

// Element composition upon Text
type Element struct {
	value      string
	nodeType   NodeType
	children   []Element
	attributes map[string]string
}

// SetData of the TextNode
func (node *Element) SetData(value string) {
	node.value = value
}

// Child returns the child at index i
func (node *Element) Child(i int) Element {
	return node.children[i]
}

// Add a node(s) to the tree
func (node *Element) Add(child Element) {
	node.children = append(node.children, child)
}

// NumberOfChildren returns the length of array children in struct ElementNode
func (node *Element) NumberOfChildren() int {
	if node.nodeType == TextNode {
		return 0
	}
	return len(node.children)
}

// ToString that the node is an ElementNode
func (node *Element) ToString(indentation ...int) string {
	var pretty string

	if len(indentation) == 0 {
		indentation = []int{0}
	}

	indent := 0

	if node.NumberOfChildren() > 0 {
		for i := 0; i < node.NumberOfChildren(); i++ {
			indent = indentation[0]
			child := node.Child(i)

			if child.nodeType == ElementNode {
				pretty += spaces(indent)
				pretty += "<" + child.value
				pretty += attributes(child.attributes)
				// pretty = pretty + node.name
				if child.NumberOfChildren() > 0 {
					indent++
					if child.Child(0).nodeType == ElementNode {
						pretty += ">\n"
						pretty += spaces(indent - 2)
						pretty += child.ToString(indent)
						pretty += spaces(indent - 2)
					} else {
						pretty += ">"
						pretty += child.ToString(indent)
					}
					pretty += "</" + child.value + ">" + "\n"
				} else {
					pretty += "/>" + "\n"
				}
			} else {
				pretty += child.value
			}

		}
	}

	return pretty
}
func spaces(step int) string {
	indent := ""
	for i := 0; i < step; i++ {
		indent += "   "
	}
	return indent
}
func attributes(m map[string]string) string {
	attr := ""
	for key, value := range m {
		attr += " " + key + "=\"" + value + "\""
	}
	return attr
}

// MakeTextNode returns a new Text Node
func MakeTextNode(value string) Element {
	return Element{value: value, nodeType: TextNode}
}

// MakeElementNode returns a new Element Node
func MakeElementNode(name string) Element {
	return Element{value: name, nodeType: ElementNode, attributes: make(map[string]string)}
}
