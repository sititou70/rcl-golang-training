// page 104
// unicodeおよびUTF-8の理解のため，できるだけフルスクラッチで解いてみます
package main

import (
	"errors"
	"testing"
	"unicode"
)

func reverseUTF8(s []byte) error {
	var i int
	for i < len(s) {
		bytes, err := getUTF8CharSize(s, i)
		if err != nil {
			return err
		}

		reverseBytes(s[i : i+bytes])
		i += bytes
	}
	reverseBytes(s)

	return nil
}
func getUTF8CharSize(s []byte, i int) (int, error) {
	_, bytes, err := decodeUTF8Rune(s, i)
	if err != nil {
		return -1, err
	}

	additionalBytes := 0
	i += bytes
	for i < len(s) {
		r, bytes, err := decodeUTF8Rune(s, i)
		if err != nil {
			return -1, err
		}

		if isMark(r) {
			additionalBytes += bytes
			i += bytes
		} else {
			break
		}
	}

	return bytes + additionalBytes, nil
}
func isMark(r rune) bool {
	rune32 := uint32(r)
	if rune32 <= uint32(unicode.Mark.R16[len(unicode.Mark.R16)-1].Hi) {
		rune16 := uint16(r)
		for _, range16 := range unicode.Mark.R16 {
			if range16.Lo <= rune16 && rune16 <= range16.Hi && (range16.Stride == 1 || (rune16-range16.Lo)%range16.Stride == 0) {
				return true
			}
		}
	} else {
		for _, range32 := range unicode.Mark.R32 {
			if range32.Lo <= rune32 && rune32 <= range32.Hi && (range32.Stride == 1 || (rune32-range32.Lo)%range32.Stride == 0) {
				return true
			}
		}
	}

	return false
}
func decodeUTF8Rune(s []byte, i int) (rune, int, error) {
	switch {
	case s[i]&0x80 == 0x00:
		return rune(s[i]), 1, nil
	case s[i]&0xE0 == 0xC0:
		return rune(s[i]&0x1F)<<6 | rune(s[i+1]&0x3F), 2, nil
	case s[i]&0xF0 == 0xE0:
		return rune(s[i]&0x0F)<<12 | rune(s[i+1]&0x3F)<<6 | rune(s[i+2]&0x3F), 3, nil
	case s[i]&0xF8 == 0xF0:
		return rune(s[i]&0x07)<<18 | rune(s[i+1]&0x3F)<<12 | rune(s[i+2]&0x3F)<<6 | rune(s[i+3]&0x3F), 4, nil
	}

	return 0, -1, errors.New("invalid UTF-8 encoding!")
}
func reverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestReverseUTF8(t *testing.T) {
	var cases = []struct {
		orig   string
		expect string
	}{
		{"12345", "54321"},
		{"a©あ😄", "😄あ©a"},
		// Combining character check: स् (U+0938 + U+094D), see also: https://unicode-table.com/jp/blocks/devanagari/
		{"स्", "स्"},
		// Variant form check: 禰󠄀 (U+79B0 + U+E0100), see also: https://747.github.io/vsselector/
		{"竈門禰󠄀豆子", "子豆禰󠄀門竈"},
		{"Hello你好안녕하세요สวัสดีค่ะनमस्कारជម្រាបសួរ។नमस्तेவணக்கம்ನಮಸ್ಕಾರგამარჯობათPäivääЗдравствуйтеጤናይስጥልኝ!Gâitsēs", "sēstiâG!ኝልጥስይናጤетйувтсвардЗääviäPთაბოჯრამაგರಕಾಸ್ಮನம்கக்ணவतेस्मन។រសួបរាម្ជरकास्मनะค่ดีสวัส요세하녕안好你olleH"},
		{"しんぶんし", "しんぶんし"},
		{"新聞紙", "紙聞新"},
		{"なかきよのとおのねふりのみなめさめなみのりふねのおとのよきかな", "なかきよのとおのねふりのみなめさめなみのりふねのおとのよきかな"},
		// from: yonashiro-san
		//
		//
		{"Z͑ͫ̓ͪ̂ͫ̽͏̴̙̤̞͉͚̯̞̠͍A̴̵̜̰͔ͫ͗͢L̠ͨͧͩ͘G̴̻͈͍͔̹̑͗̎̅͛́Ǫ̵̹̻̝̳͂̌̌͘!͖̬̰̙̗̿̋ͥͥ̂ͣ̐́́͜͞", "!͖̬̰̙̗̿̋ͥͥ̂ͣ̐́́͜͞Ǫ̵̹̻̝̳͂̌̌͘G̴̻͈͍͔̹̑͗̎̅͛́L̠ͨͧͩ͘A̴̵̜̰͔ͫ͗͢Z͑ͫ̓ͪ̂ͫ̽͏̴̙̤̞͉͚̯̞̠͍"},
	}

	for _, c := range cases {
		origBytes := []byte(c.orig)
		err := reverseUTF8(origBytes)
		if err != nil {
			t.Errorf("--------------------\n")
			t.Errorf("%s\n", err.Error())
			t.Errorf("expected: %v\n", c.expect)
		}

		reversed := string(origBytes)
		if reversed != c.expect {
			t.Errorf("--------------------\n")
			t.Errorf("expected: %v, but actual: %v\n", c.expect, reversed)
			t.Errorf("expect: % x\n", c.expect)
			t.Errorf("actual: % x\n", reversed)
		}
	}
}
