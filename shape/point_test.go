package shape

import (
	"reflect"
	"testing"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		// TODO: Add test cases.
		{
			"point rune is '❤'",
			args{
				r: '❤',
			},
			NewPoint('❤'),
		},
		{
			"point rune is '*'",
			args{
				r: '*',
			},
			NewPoint('*'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Width(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		want uint
	}{
		// TODO: Add test cases.
		{
			"point rune is '❤'",
			NewPoint('❤'),
			2,
		},
		{
			"point rune is '*'",
			NewPoint('*'),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Width(); got != tt.want {
				t.Errorf("Point.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Height(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		want uint
	}{
		// TODO: Add test cases.
		{
			"point rune is '❤'",
			NewPoint('❤'),
			1,
		},
		{
			"point rune is '*'",
			NewPoint('*'),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Height(); got != tt.want {
				t.Errorf("Point.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
