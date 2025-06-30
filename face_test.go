package fontpic

import (
	"reflect"
	"testing"
)

func Test_bits2byte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name       string
		args       args
		wantStride [8]byte
	}{
		{"all 1s", args{0xff}, [8]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{"fst 1 ", args{0x80}, [8]byte{0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"lst 1 ", args{0x01}, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff}},
		{"intrlv", args{0x55}, [8]byte{0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff}},
		{"all 1s", args{0xaa}, [8]byte{0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStride := bits2bytes(tt.args.b); !reflect.DeepEqual(gotStride, tt.wantStride) {
				t.Errorf("bits2byte() = %v, want %v", gotStride, tt.wantStride)
			}
		})
	}
}

func Test_uint16ToUint8(t *testing.T) {
	type args struct {
		data []uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "converts 16-bit slice to 8-bit big-endian representation",
			args: args{
				[]uint16{0x0123, 0x4567, 0x89AB, 0xCDEF},
			},
			want: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uint16ToUint8(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uint16ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint16ToUint8Rev(t *testing.T) {
	type args struct {
		data []uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "reverses bits in each word",
			args: args{
				[]uint16{0b0000_0000_0000_0001, 0b0011_0111_1101_0011},
			},
			want: []byte{0b1000_0000, 0b0000_0000, 0b1100_1011, 0b1110_1100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uint16ToUint8Rev(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uint16ToUint8Rev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_everySecond(t *testing.T) {
	type args struct {
		data []uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "inserts zero every second element",
			args: args{
				data: []uint8{1, 2, 3, 4},
			},
			want: []uint8{1, 0, 2, 0, 3, 0, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := everySecond(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("everySecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
