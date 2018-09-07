package main

import (
	"../mathops"
	"fmt"
	"math/big"
)

func main() {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "1820"

	result, err := mathops.Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		fmt.Printf("Error returned by Probability{}.CombinationsNoRepsBigInt(" +
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
		return
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		fmt.Printf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
		return
	}

	fmt.Println("       Num Of Items: ", numOfItems)
	fmt.Println("Num Of Items Chosen: ", numOfItemsChosen)
	fmt.Println(" CombinationsNoReps: ", actualResultStr)
	fmt.Println("    Expected Result: ", expectedResultStr)
}
