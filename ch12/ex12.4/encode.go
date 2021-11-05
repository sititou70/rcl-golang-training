// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Pretty(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), ""); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value, indent string) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "t")
		} else {
			fmt.Fprint(buf, "nil")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		f := v.Float()
		fmt.Fprintf(buf, "%g", f)

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%g %g)", real(c), imag(c))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		indent += " "
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteString(indent)
			}

			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}

			if i != v.Len()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		indent += " "
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteString(indent)
			}

			buf.WriteString("(" + v.Type().Field(i).Name + " ")
			if err := encode(buf, v.Field(i), indent+strings.Repeat(" ", len(v.Type().Field(i).Name)+2)); err != nil {
				return err
			}

			buf.WriteByte(')')
			if i != v.NumField()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		indent += " "
		for i, key := range v.MapKeys() {
			if i != 0 {
				buf.WriteString(indent)
			}

			buf.WriteByte('(')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}

			buf.WriteByte(')')
			if i != len(v.MapKeys())-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')

	case reflect.Chan:
		fmt.Fprintf(buf, `#Chan("%s")`, getTypeIdentifire(v.Type()))

	case reflect.Func:
		fmt.Fprintf(buf, `#Func("%s")`, getTypeIdentifire(v.Type()))

	case reflect.Interface:
		i := v.Interface()
		fmt.Fprintf(buf, `#Iface("%T" `, i)
		if err := encode(buf, reflect.ValueOf(i), indent); err != nil {
			return err
		}
		buf.WriteByte(')')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func getTypeIdentifire(t reflect.Type) string {
	return t.PkgPath() + ":" + t.String()
}
