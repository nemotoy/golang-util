package slice

import (
	"reflect"
	"testing"
)

func TestUInt(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "hasn't duplicated values",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "has duplicated values",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
