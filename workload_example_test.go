package testutil

import (
	"fmt"
	"strings"
)

func ExampleWorkload() {

	datasetSize := 10
	times := 50

	got := Workload("zipf", datasetSize, times)

	arr := make([]int, datasetSize)
	for _, idx := range got {
		arr[idx]++
	}

	fmt.Println("Accessed:")
	for _, v := range arr {
		fmt.Println("|" + strings.Repeat("*", v))
	}

	// Output:
	//
	// Accessed:
	// |****
	// |
	// |
	// |*
	// |****************
	// |****
	// |***
	// |***********
	// |***
	// |********

}
