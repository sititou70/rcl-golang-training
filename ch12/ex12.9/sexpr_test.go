// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	data := `((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964) (Actor (("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden") ("Maj. T.J. \"King\" Kong" "Slim Pickens"))) (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)")))`
	d := NewDecoder(bytes.NewReader([]byte(data)))
	out := strings.Builder{}
	depth := 0
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Token() failed: %v", err)
		}

		getIndent := func(depth int) string { return strings.Repeat("  ", depth) }
		switch tok.(type) {
		case StartList:
			out.WriteString(getIndent(depth) + "(\n")
			depth++
		case EndList:
			depth--
			out.WriteString(getIndent(depth) + ")\n")
		case String:
			out.WriteString(fmt.Sprintf("%s%q\n", getIndent(depth), tok.(String).Value))
		case Int:
			out.WriteString(fmt.Sprintf("%s%d\n", getIndent(depth), tok.(Int).Value))
		case Symbol:
			out.WriteString(fmt.Sprintf("%s%s\n", getIndent(depth), tok.(Symbol).Name))
		}
	}

	want, _ := os.ReadFile("want.txt")
	if out.String() != string(want) {
		t.Fatal()
	}

	println(out.String())
}
