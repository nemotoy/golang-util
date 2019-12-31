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

/*
	$ go test -count 3 -benchmem -bench .
	goos: darwin
	goarch: amd64
	pkg: github.com/nemotoy/golang-util/slice
	BenchmarkUInt-4            22629             53242 ns/op           32921 B/op          7 allocs/op
	BenchmarkUInt-4            22537             53996 ns/op           32921 B/op          7 allocs/op
	BenchmarkUInt-4            21912             59558 ns/op           32921 B/op          7 allocs/op
	PASS
	ok    github.com/nemotoy/golang-util/slice      5.373s
*/
func BenchmarkUInt(b *testing.B) {
	var ii []int
	for i := 0; i < 1000; i++ {
		ii = append(ii, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ii = UInt(ii)
	}
}

func TestUIntASC(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "descending order slice of int",
			in:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "duplicated slice of int",
			in:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UIntASC(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UIntASC() = %v, want %v", got, tt.want)
			}
		})
	}
}
