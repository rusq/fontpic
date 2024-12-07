package charset

import (
	"reflect"
	"testing"
)

func TestCharset_TranslateRune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		c    Charset
		args args
		want byte
	}{
		{
			name: "0x00",
			c:    CP866,
			args: args{r: 0x00},
			want: 0x00,
		},
		{
			name: "Cyrillic A",
			c:    CP866,
			args: args{r: 'А'},
			want: 0x80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.TranslateRune(tt.args.r); got != tt.want {
				t.Errorf("Charset.TranslateRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharset_Translate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		c    Charset
		args args
		want []byte
	}{
		{
			name: "АБВГД",
			c:    CP866,
			args: args{s: "АБВГД"},
			want: []byte{0x80, 0x81, 0x82, 0x83, 0x84},
		},
		{
			name: "Привет",
			c:    CP866,
			args: args{s: "Привет из 1989"},
			want: []byte{0x8F, 0xe0, 0xA8, 0xA2, 0xA5, 0xE2, 0x20, 0xA8, 0xA7, 0x20, 0x31, 0x39, 0x38, 0x39},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Translate(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Charset.Translate() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
