package sexpr

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title    string            `sexpr:"title"`
		Subtitle string            `sexpr:"subtitle"`
		Year     int               `sexpr:"year"`
		Actor    map[string]string `sexpr:"actor"`
		Oscars   []string          `sexpr:"oscars"`
		Sequel   *string           `sexpr:"sequel"`
	}
	input := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s", data)

	var actual Movie
	err = Unmarshal(data, &actual, TypeMap{})
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(input, actual) {
		t.Fatal("not equal")
	}
}
