package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Running...")
	var div = MakeElementNode("div")

	var h1 = MakeElementNode("h1")
	h1.Add(MakeTextNode("sample"))
	// h1.attributes["style"] = "font-weight: 400;"

	var h2 = MakeElementNode("h2")
	h2.Add(MakeTextNode("etc"))
	h2.attributes["class"] = "subheader"

	div.Add(h1)
	div.Add(h2)

	var d = MakeElementNode("div")
	d.Add(h1)
	d.Add(h2)

	div.Add(d)
	// fmt.Println(div)
	// fmt.Println(div.ToString())

	f, err := os.Open("data/test.html")
	// // defer f.Close()
	check(err)

	// _, err = f.Seek(0, 0)
	// check(err)

	r := io.Reader(f)
	node := Parse(r)
	fmt.Println("Pretty Print:")
	fmt.Println(node.ToString())
	fmt.Println("Regular Print:")
	fmt.Println(node)

}
