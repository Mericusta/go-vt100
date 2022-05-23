package shape

import (
	"reflect"
	"testing"

	"github.com/Mericusta/go-vt100/core"
)

func TestNewLine(t *testing.T) {
	type args struct {
		p Point
		l uint
		d core.Direction
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		// TODO: Add test cases.
		{
			"horizontal line, length is 5, point rune is '❤'",
			args{
				p: NewPoint('❤'),
				l: 5,
				d: core.Horizontal,
			},
			NewLine(NewPoint('❤'), 5, core.Horizontal),
		},
		{
			"vertical line, length is 5, point rune is '❤'",
			args{
				p: NewPoint('❤'),
				l: 5,
				d: core.Vertical,
			},
			NewLine(NewPoint('❤'), 5, core.Vertical),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLine(tt.args.p, tt.args.l, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_Width(t *testing.T) {
	tests := []struct {
		name string
		l    Line
		want uint
	}{
		// TODO: Add test cases.
		{
			"horizontal line, length is 5, point rune is '❤'",
			NewLine(NewPoint('❤'), 5, core.Horizontal),
			10,
		},
		{
			"vertical line, length is 5, point rune is '❤'",
			NewLine(NewPoint('❤'), 5, core.Vertical),
			1,
		},
		{
			"horizontal line, length is 5, point rune is '*'",
			NewLine(NewPoint('*'), 5, core.Horizontal),
			5,
		},
		{
			"vertical line, length is 5, point rune is '*'",
			NewLine(NewPoint('*'), 5, core.Vertical),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Width(); got != tt.want {
				t.Errorf("Line.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_Height(t *testing.T) {
	tests := []struct {
		name string
		l    Line
		want uint
	}{
		// TODO: Add test cases.
		{
			"horizontal line, length is 5, point rune is '❤'",
			NewLine(NewPoint('❤'), 5, core.Horizontal),
			1,
		},
		{
			"vertical line, length is 5, point rune is '❤'",
			NewLine(NewPoint('❤'), 5, core.Vertical),
			5,
		},
		{
			"horizontal line, length is 5, point rune is '*'",
			NewLine(NewPoint('*'), 5, core.Horizontal),
			1,
		},
		{
			"vertical line, length is 5, point rune is '*'",
			NewLine(NewPoint('*'), 5, core.Vertical),
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Height(); got != tt.want {
				t.Errorf("Line.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
