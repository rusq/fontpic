package charset

import "unicode/utf8"

type Charset string

func (c Charset) Translate(s string) []byte {
	b := make([]byte, utf8.RuneCountInString(s))
	for i, r := range []rune(s) {
		b[i] = c.TranslateRune(r)
	}
	return b
}

func (c Charset) TranslateRune(r rune) byte {
	if r < 0x80 {
		return byte(r)
	}
	for i, rr := range []rune(c) {
		if rr == r {
			return byte(i) + 0x80
		}
	}
	return c[r]
}
