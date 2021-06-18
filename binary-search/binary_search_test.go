package binary_search_test

import (
	"testing"

	binary_search "github.com/iAziz786/algosandds/binary-search"
)

func TestSearch(t *testing.T) {
	type args struct {
		array  []interface{}
		target interface{}
		fn     binary_search.Compare
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with odd items",
			args: args{
				array:  []interface{}{3, 33, 333, 3333, 33333, 333333, 3333333},
				target: 3333,
				fn: func(a, b interface{}) int {
					x := a.(int)
					y := b.(int)
					return x - y
				},
			},
			want: 3,
		},
		{
			name: "with even items",
			args: args{
				array:  []interface{}{3, 33, 333, 3333, 33333, 333333},
				target: 33333,
				fn: func(a, b interface{}) int {
					x := a.(int)
					y := b.(int)
					return x - y
				},
			},
			want: 4,
		},
		{
			name: "with item not present",
			args: args{
				array:  []interface{}{3, 33, 333, 3333, 33333, 333333},
				target: 666,
				fn: func(a, b interface{}) int {
					x := a.(int)
					y := b.(int)
					return x - y
				},
			},
			want: -1,
		},
		{
			name: "with empty item",
			args: args{
				array:  []interface{}{},
				target: 666,
				fn: func(a, b interface{}) int {
					x := a.(int)
					y := b.(int)
					return x - y
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary_search.Search(tt.args.array, tt.args.target, tt.args.fn); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
