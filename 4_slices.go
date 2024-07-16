package main

import (
	"fmt"
	"slices"
)

func main() {

	//
	//func main() {
	//Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, "nil:", s == nil, "len:", len(s))

	s = make([]string, 5)
	fmt.Println("empty:", s, "nil:", s == nil, "len:", len(s))

	//We can set and get just like with arrays.

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"

	fmt.Println("set:", s)
	fmt.Println("get:", s[1])

	//In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "f")
	s = append(s, "g", "h")
	fmt.Println("append:", s)

	//Slices can also be copied. Here we create an empty slice c of the same length as s and copy into c from s.
	copied := make([]string, len(s))
	copy(copied, s)
	fmt.Println("copy:", copied)

	//Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4]
	slice := s[2:5]
	fmt.Println("slice:", slice)

	//This slices up to (but excluding) s[5]
	slice = s[:5]
	fmt.Println("slice:", slice)

	//And this slices up from (and including) s[2].
	slice = s[2:]
	fmt.Println("slice:", slice)

	//We can declare and initialize a variable for slice in a single line as well.
	sliceInitialized := []string{"a", "a", "b", "c"}
	fmt.Println("initialized:", sliceInitialized)

	sliceInitialized2 := []string{"a", "a", "c", "b", "a"}

	//The slices package contains a number of useful utility functions for slices.
	if slices.Equal(sliceInitialized, sliceInitialized2) {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
	}

	if slices.Compare(sliceInitialized, sliceInitialized2) == 0 {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
	}

	sliceCopy := slices.Clone(sliceInitialized)
	sliceCopy = append(sliceCopy[:1], sliceCopy[2:]...)
	fmt.Println("element removed from copy:", sliceCopy)
	fmt.Println("original slice:", sliceInitialized)
	sliceForCompact := []string{"a", "b", "c"}

	fmt.Println("Original compacted: ", sliceForCompact)
	slices.Compact(sliceForCompact)
	fmt.Println("Slice compacted: ", sliceForCompact)

	slices.Sort(sliceInitialized2)

	fmt.Println("Sorted: ", sliceInitialized2)

	sliceConcat := slices.Concat(sliceInitialized, sliceInitialized2)

	fmt.Println("Concat: ", sliceConcat, len(sliceConcat), "=", len(sliceInitialized), "+", len(sliceInitialized2))

	sliceConcat = slices.Compact(sliceConcat)
	fmt.Println("Slice concatenated compacted:", sliceConcat, len(sliceConcat))

}
