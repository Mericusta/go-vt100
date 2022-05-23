package shape

import (
	"reflect"
	"testing"

	"github.com/Mericusta/go-vt100/core"
)

func TestNewGrid(t *testing.T) {
	type args struct {
		content map[int]map[int][]byte
	}
	tests := []struct {
		name string
		args args
		want *Grid
	}{
		// TODO: Add test cases.
		{
			"1x1 grid with 1 width content",
			args{
				content: map[int]map[int][]byte{
					1: {
						1: []byte("0"),
					},
				},
			},
			&Grid{
				col: 1,
				row: 1,
				content: map[int]map[int][]byte{
					1: {
						1: []byte("0"),
					},
				},
				size: core.Size{
					Width:  3,
					Height: 3,
				},
				maxContentWidth: 1,
			},
		},
		// TODO: Add test cases.
		{
			"3x3 grid with 3 width content",
			args{
				content: map[int]map[int][]byte{
					3: {
						3: []byte(" 0 "),
					},
				},
			},
			&Grid{
				col: 3,
				row: 3,
				content: map[int]map[int][]byte{
					3: {
						3: []byte(" 0 "),
					},
				},
				size: core.Size{
					Width:  13,
					Height: 7,
				},
				maxContentWidth: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrid(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
