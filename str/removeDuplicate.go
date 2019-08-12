//Program to remove duplicate characters from given string
//Author 	: Rahul Sutar

package str

import (
	"fmt"
	"strings"
)

//RemoveDuplicateCharsFromString - Remove duplicate characters from given string
func RemoveDuplicateCharsFromString() {

	stringData := ""

	fmt.Print("Enter a string : ")
	fmt.Scan(&stringData)

	if strings.TrimSpace(stringData) == "" {
		fmt.Println("Enter a valid string")
		return
	}

	var unduplicatedStr = ""

	//If case is to be ignored
	stringData = strings.ToLower(stringData)
	for i := 0; i < len(stringData); i++ {
		if !findIfExist(unduplicatedStr, stringData[i]) {
			unduplicatedStr += string(stringData[i])
		}
	}

	fmt.Println("Unupliated string : ", unduplicatedStr)
}

func findIfExist(intputStr string, charToFind byte) bool {
	for i := 0; i < len(intputStr); i++ {
		if intputStr[i] == charToFind {
			return true
		}
	}
	return false
}

// func findIfExist(intputStr []rune, charToFind rune) bool {
// 	for i := 0; i < len(intputStr); i++ {
// 		if intputStr[i] == charToFind {
// 			return true
// 		}
// 	}
// 	return false
// }
