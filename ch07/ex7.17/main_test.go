package main

import (
	"encoding/xml"
	"os"
	"reflect"
	"testing"
)

func TestParseSelectQuery(t *testing.T) {
	e := parseSelectQuery("name[attr1='val1'][attr2='val2']")
	if e.Name.Local != "name" {
		t.Fail()
	}
	if len(e.Attr) != 2 {
		t.Fail()
	}
	if e.Attr[0].Name.Local != "attr1" || e.Attr[0].Value != "val1" {
		t.Fail()
	}
	if e.Attr[1].Name.Local != "attr2" || e.Attr[1].Value != "val2" {
		t.Fail()
	}
}

func TestXmlselect1(t *testing.T) {
	f, err := os.Open("./assets/sample.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := []xml.CharData{}
	xmlselect(f, []string{"a[href='/']"}, func(stack []xml.StartElement, tok xml.CharData) {
		data = append(data, tok)
	})

	if len(data) != 2 {
		t.Fail()
	}
	if reflect.DeepEqual(data[0], []byte("W3C")) {
		t.Fail()
	}
	if reflect.DeepEqual(data[0], []byte("HOME")) {
		t.Fail()
	}
}

func TestXmlselect2(t *testing.T) {
	f, err := os.Open("./assets/sample.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := []xml.CharData{}
	xmlselect(f, []string{"div", "div", "h2"}, func(stack []xml.StartElement, tok xml.CharData) {
		data = append(data, tok)
	})

	if len(data) != 2 {
		t.Fail()
	}
	if reflect.DeepEqual(data[0], []byte("Navigation")) {
		t.Fail()
	}
	if reflect.DeepEqual(data[0], []byte("Navigation")) {
		t.Fail()
	}
}

func TestXmlselect3(t *testing.T) {
	f, err := os.Open("./assets/sample.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := []xml.CharData{}
	xmlselect(f, []string{"div[id='w3c_sidenavigation'][class='w3c_leftCol']", "a[href='/Help/Account/']"}, func(stack []xml.StartElement, tok xml.CharData) {
		data = append(data, tok)
	})

	if len(data) != 1 {
		t.Fail()
	}
	if reflect.DeepEqual(data[0], []byte("W3C User Account Management")) {
		t.Fail()
	}
}
