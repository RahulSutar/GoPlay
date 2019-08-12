//Support function required for general purpose
//Author	: Rahul Sutar

package helpers

//SwapInArrayByIndex - Function to swap two values in an array by element index no
func SwapInArrayByIndex(intArray []int, sourceIndex, destinationIndex int) {
	intArray[sourceIndex], intArray[destinationIndex] = intArray[destinationIndex], intArray[sourceIndex]
}

//SwapTwoMemoryLocations - Swap values of two memmory locations
func SwapTwoMemoryLocations(source, destination *int) {
	*source, *destination = *destination, *source
}
