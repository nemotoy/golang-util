package slice_test

import (
	"fmt"

	"github.com/nemotoy/golang-util/slice"
)

func ExampleUInt() {
	var in = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice.UInt(in))
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExampleUIntASC() {
	var in = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(slice.UIntASC(in))
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}
