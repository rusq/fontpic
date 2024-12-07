package fontpic

import (
	_ "embed"
	"testing"
)

func Test_charStride(t *testing.T) {
	type args struct {
		width int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"8", args{8}, 1},
		{"16", args{16}, 2},
		{"32", args{32}, 4},
		{"1", args{1}, 1},
		{"0", args{0}, 1},
		{"15", args{15}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charStride(tt.args.width); got != tt.want {
				t.Errorf("charStride() = %v, want %v", got, tt.want)
			}
		})
	}
}
