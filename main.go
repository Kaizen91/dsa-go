package main

import (
	"github.com/Kaizen91/DSA-Go/Search"
	"fmt"
)

func main() {
	res := search.BinarySearch(1, []int{1,2,3,4,5})
	fmt.Printf("%v\n", res)
}
