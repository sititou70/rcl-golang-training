// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	type fieldEntry struct {
		value  reflect.Value
		regexp regexp.Regexp
	}
	fields := make(map[string]fieldEntry)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		regStr := tag.Get("regexp")
		var reg *regexp.Regexp
		if regStr != "" {
			var err error
			reg, err = regexp.Compile(tag.Get("regexp"))
			if err != nil {
				return fmt.Errorf("Unpack: regexp compile error")
			}
		}
		fields[name] = fieldEntry{v.Field(i), *reg}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		field := fields[name]
		if !field.value.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if !field.regexp.Match([]byte(value)) {
				return fmt.Errorf("Unpack: %s did not match regexp rule of %s", value, name)
			}
			if field.value.Kind() == reflect.Slice {
				elem := reflect.New(field.value.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				field.value.Set(reflect.Append(field.value, elem))
			} else {
				if err := populate(field.value, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate

type Data struct {
	Mail   string `http:"m" regexp:"^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\\.)+[a-zA-Z]{2,}$"`
	Credit string `http:"c" regexp:"^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"`
	Phone  string `http:"p" regexp:"^(\\+\\d{1,2}\\s)?\\(?\\d{3}\\)?[\\s.-]\\d{3}[\\s.-]\\d{4}$"`
}

func TestCorrectRequest(t *testing.T) {
	data := Data{
		Mail:   "test@example.com",
		Credit: "378282246310005",
		Phone:  "123-456-7890",
	}
	u := &url.URL{}
	u.RawQuery = fmt.Sprintf("m=%s&c=%s&p=%s", data.Mail, data.Credit, data.Phone)
	r := &http.Request{URL: u}

	var parsed Data
	err := Unpack(r, &parsed)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(parsed, data) {
		t.Fatalf("Unpack() = %v, want %v", parsed, data)
	}
}

func TestWrongRequest1(t *testing.T) {
	data := Data{
		Mail:   "testexample.com",
		Credit: "378282246310005",
		Phone:  "123-456-7890",
	}
	u := &url.URL{}
	u.RawQuery = fmt.Sprintf("m=%s&c=%s&p=%s", data.Mail, data.Credit, data.Phone)
	r := &http.Request{URL: u}

	var parsed Data
	err := Unpack(r, &parsed)

	if err == nil {
		t.Fatal("wrong request did not cause error!")
	}
}
func TestWrongRequest2(t *testing.T) {
	data := Data{
		Mail:   "test@example.com",
		Credit: "37882246310005",
		Phone:  "123-456-7890",
	}
	u := &url.URL{}
	u.RawQuery = fmt.Sprintf("m=%s&c=%s&p=%s", data.Mail, data.Credit, data.Phone)
	r := &http.Request{URL: u}

	var parsed Data
	err := Unpack(r, &parsed)

	if err == nil {
		t.Fatal("wrong request did not cause error!")
	}
}
func TestWrongRequest3(t *testing.T) {
	data := Data{
		Mail:   "test@example.com",
		Credit: "378282246310005",
		Phone:  "123-4567890",
	}
	u := &url.URL{}
	u.RawQuery = fmt.Sprintf("m=%s&c=%s&p=%s", data.Mail, data.Credit, data.Phone)
	r := &http.Request{URL: u}

	var parsed Data
	err := Unpack(r, &parsed)

	if err == nil {
		t.Fatal("wrong request did not cause error!")
	}
}
