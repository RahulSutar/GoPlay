//Program to perform selection sort
//Author 	: Rahul Sutar

package sorting

import "GoPlay/helpers"

//Algorithm Reference - https://www.youtube.com/watch?v=NiyEqLZmngY

//BubbleSort - Sorting algorithm of selection sort
func BubbleSort(intArr []int) []int {

	arrLen := len(intArr)
	swapped := false
	sorted := false

	for !sorted {

		swapped = false

		for i := 0; i < arrLen-1; i++ {

			if intArr[i] > intArr[i+1] {
				helpers.SwapTwoMemoryLocations(&intArr[i], &intArr[i+1])
				swapped = true
			}

		}

		if !swapped {
			sorted = true
		}

		arrLen = arrLen - 1
	}

	return intArr

}
