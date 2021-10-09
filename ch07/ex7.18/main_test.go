package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func parseXml(r io.Reader) *Element {
	dec := xml.NewDecoder(r)

	root := Element{
		Type:     xml.Name{Local: "root"},
		Attr:     []xml.Attr{},
		Children: []Node{},
	}
	stack := []*Element{&root}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "parseXml: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			latest := stack[len(stack)-1]
			newElement := Element{
				Type:     tok.Name,
				Attr:     tok.Attr,
				Children: []Node{},
			}
			latest.Children = append(latest.Children, &newElement)
			stack = append(stack, &newElement) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			latest := stack[len(stack)-1]
			latest.Children = append(latest.Children, CharData(tok))
		}
	}

	return stack[0]
}

func TestParseXml(t *testing.T) {
	f, err := os.Open("./assets/sample.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	e := parseXml(f)
	if e.Children[1].(*Element).Children[3].(*Element).Children[1].(*Element).Children[1].(*Element).Children[2].(*Element).Children[3].(*Element).Children[0].(CharData) != "W3C" {
		t.Fail()
	}
}
