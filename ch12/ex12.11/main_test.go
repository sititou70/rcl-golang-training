package main

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func Pack(data interface{}) string {
	v := reflect.ValueOf(data)
	q := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Array, reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				q.Add(v.Type().Field(i).Tag.Get("http"), fmt.Sprint(f.Index(j)))
			}
		default:
			q.Add(v.Type().Field(i).Tag.Get("http"), fmt.Sprint(f))
		}
	}

	return q.Encode()
}

func Test(t *testing.T) {
	type Data struct {
		Label     []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}

	url := Pack(Data{
		Label:     []string{"label1", "label2", "label3"},
		MaxResult: 123,
		Exact:     true,
	})

	want := "l=label1&l=label2&l=label3&max=123&x=true"
	if url != want {
		t.Fatalf("url = %v, want %v", url, want)
	}
}
