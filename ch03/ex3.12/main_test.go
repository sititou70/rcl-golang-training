// page 83
package main

import (
	"testing"
)

func TestCheckAnagram(t *testing.T) {
	if checkAnagram("a gentleman", "elegant man") != true {
		t.Fail()
	}
	if checkAnagram("a gentleman!", "elegant man") != false {
		t.Fail()
	}
	if checkAnagram("eleven plus two", "twelve plus one") != true {
		t.Fail()
	}
	if checkAnagram("いろはにほへとちりぬるをわかよたれそつねならむうゐのおくやまけふこえてあさきゆめみしゑひもせす", "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわゐゑを") != true {
		t.Fail()
	}
	if checkAnagram("いろはにほへとちりぬるをわかよたれそつねならむうゐのおくやまけふこえてあさきゆめみしゑひもせす", "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわゐゑをん") != false {
		t.Fail()
	}
}
