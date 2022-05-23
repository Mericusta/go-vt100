package shape

import (
	"reflect"
	"testing"

	"github.com/Mericusta/go-vt100/core"
)

func TestNewRectangle(t *testing.T) {
	type args struct {
		l Line
		c uint
	}
	tests := []struct {
		name string
		args args
		want Rectangle
	}{
		// TODO: Add test cases.
		{
			"5x5 rectangle by using horizontal line, point rune is '❤'",
			args{
				l: NewLine(NewPoint('❤'), 5, core.Horizontal),
				c: 5,
			},
			NewRectangle(NewLine(NewPoint('❤'), 5, core.Horizontal), 5),
		},
		{
			"5x5 rectangle by using vertical line, point rune is '❤'",
			args{
				l: NewLine(NewPoint('❤'), 5, core.Vertical),
				c: 5,
			},
			NewRectangle(NewLine(NewPoint('❤'), 5, core.Vertical), 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRectangle(tt.args.l, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Width(t *testing.T) {
	tests := []struct {
		name string
		l    Rectangle
		want uint
	}{
		// TODO: Add test cases.
		{
			"5x5 rectangle by using horizontal line, point rune is '❤'",
			NewRectangle(NewLine(NewPoint('❤'), 5, core.Horizontal), 5),
			10,
		},
		{
			"5x5 rectangle by using vertical line, point rune is '❤'",
			NewRectangle(NewLine(NewPoint('❤'), 5, core.Vertical), 5),
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Width(); got != tt.want {
				t.Errorf("Rectangle.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Height(t *testing.T) {
	tests := []struct {
		name string
		l    Rectangle
		want uint
	}{
		// TODO: Add test cases.
		{
			"5x5 rectangle by using horizontal line, point rune is '❤'",
			NewRectangle(NewLine(NewPoint('❤'), 5, core.Horizontal), 5),
			5,
		},
		{
			"5x5 rectangle by using vertical line, point rune is '❤'",
			NewRectangle(NewLine(NewPoint('❤'), 5, core.Vertical), 5),
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Height(); got != tt.want {
				t.Errorf("Rectangle.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
