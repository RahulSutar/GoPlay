//Program to perform selection sort
//Author 	: Rahul Sutar

package sorting

import "GoPlay/helpers"

//Algorithm Reference - https://www.youtube.com/watch?v=EdUWyka7kpI

//SelectionSort - Sorting algorithm of selection sort
func SelectionSort(intArr []int) []int {

	for i := 0; i < len(intArr)-1; i++ {

		minAtIndex := i

		for j := (i + 1); j < len(intArr); j++ {

			if intArr[j] < intArr[minAtIndex] {
				minAtIndex = j
			}
		}
		helpers.SwapTwoMemoryLocations(&intArr[i], &intArr[minAtIndex])
		// SwapInArrayByIndex(intArr, i, minAtIndex)

	}

	return intArr
}
