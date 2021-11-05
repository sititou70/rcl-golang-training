// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "true")
		} else {
			fmt.Fprint(buf, "false")
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
		fmt.Fprintf(buf, `"#C(%g %g)"`, real(c), imag(c))

	case reflect.String:
		fmt.Fprintf(buf, `%q`, v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteString(",")
			}

			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteString(",")
			}

			buf.WriteString(`"` + v.Type().Field(i).Name + `":`)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i != 0 {
				buf.WriteString(",")
			}

			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteString(`:`)
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Chan:
		fmt.Fprintf(buf, `"#Chan(%s)"`, getTypeIdentifire(v.Type()))

	case reflect.Func:
		fmt.Fprintf(buf, `"#Func(%s)"`, getTypeIdentifire(v.Type()))

	case reflect.Interface:
		i := v.Interface()
		fmt.Fprintf(buf, `"#Iface((%T) `, i)
		if err := encode(buf, reflect.ValueOf(i)); err != nil {
			return err
		}
		buf.WriteByte('"')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func getTypeIdentifire(t reflect.Type) string {
	return t.PkgPath() + ":" + t.String()
}
