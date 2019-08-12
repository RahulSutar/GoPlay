//Main file to execute sorting programs
//Author 	: Rahul Sutar

package main

import (
	"fmt"
	"GoPlay/sorting"
)

func main() {

	intArray := []int{12, 4, 5, 25, 3, 17, 65, 31}

	fmt.Println("Before sorting : ", intArray)

	//Selection Sort
	// resultArray := sorting.SelectionSort(intArray)

	//Bubble Sort
	resultArray := sorting.BubbleSort(intArray)

	fmt.Println("After sorting : ", resultArray)

}
