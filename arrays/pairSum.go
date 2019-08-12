//Program to find pairs from give array whose sum is equal to user input
//Author 	: Rahul Sutar

package arrays

import "fmt"

//PairSum - Find pair sum
func PairSum() {

	numbers := []int{2, 3, 4, 7, 5, 1}
	requiredPairTotal := 0

	fmt.Println("Array is : ", numbers)

	fmt.Print("Input pair total : ")
	fmt.Scan(&requiredPairTotal)

	//If duplicate pairs are to be considered
	// for i := 0; i < len(numbers); i++ {
	// 	for j := 0; j < len(numbers); j++ {
	// 		if numbers[i]+numbers[j] == requiredPairTotal {
	// 			fmt.Println(numbers[i], " and ", numbers[j], " add to given total")
	// 		}
	// 	}
	// }

	//If duplicate pairs are not to be considered
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == requiredPairTotal {
				fmt.Println(numbers[i], " and ", numbers[j], " add to given total")
			}
		}
	}

}