package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	xmlselect(os.Stdin, os.Args[1:], func(s []xml.StartElement, tok xml.CharData) {
		stack := []string{}
		for _, item := range s {
			stack = append(stack, item.Name.Local)
		}
		fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
	})
}

func xmlselect(r io.Reader, q []string, onMatch func(stack []xml.StartElement, tok xml.CharData)) {
	dec := xml.NewDecoder(r)

	// parse query
	queries := []xml.StartElement{}
	for _, query := range q {
		queries = append(queries, parseSelectQuery(query))
	}

	// select
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, queries) {
				onMatch(stack, tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, q []xml.StartElement) bool {
	for len(q) <= len(x) {
		if len(q) == 0 {
			return true
		}
		if matchElement(x[0], q[0]) {
			q = q[1:]
		}
		x = x[1:]
	}
	return false
}

func matchElement(e, q xml.StartElement) bool {
	if e.Name.Local != q.Name.Local {
		return false
	}

OUTER:
	for _, qAttr := range q.Attr {
		for _, eAttr := range e.Attr {
			if qAttr.Name.Local == eAttr.Name.Local && qAttr.Value == eAttr.Value {
				continue OUTER
			}
		}

		return false
	}

	return true
}

func parseSelectQuery(s string) xml.StartElement {
	r := regexp.MustCompile("^[^\\[]+")
	name := r.FindString(s)

	attr := []xml.Attr{}
	r = regexp.MustCompile("\\[.+?\\]")
	attrQuery := r.FindAllString(s, -1)
	for _, query := range attrQuery {
		r2 := regexp.MustCompile("^\\[(.+)='(.+)'\\]$")
		m := r2.FindAllStringSubmatch(query, -1)
		if len(m[0]) <= 1 {
			continue
		}

		attr = append(attr, xml.Attr{Name: xml.Name{Local: m[0][1]}, Value: m[0][2]})
	}

	return xml.StartElement{Name: xml.Name{Local: name}, Attr: attr}
}
