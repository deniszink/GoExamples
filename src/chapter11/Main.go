package main

import (
	"fmt"
	alias "chapter11/math"
	"math"
)

func main() {
	xs := []int{1, 2, 3, 4}
	avg := alias.Average(xs); //we can add alias to package import if we have duplicate names,
	// as example "math" from go library and own math
	fmt.Println(math.Min(23,65))
	fmt.Print(avg)
}
