package main

import (
	"../mathops"
	"fmt"
)

func main() {

	nInt := 15
	rInt := -2
	numOfItems := mathops.NumStrDto{}.NewInt(nInt, 0)
	numOfItemsPicked := mathops.NumStrDto{}.NewInt(rInt, 0)
	allowRepetitions := true


	result, err := mathops.Probability{}.PermutationsNumStrDto(
							numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		fmt.Printf("Error: error returned from Probability{}.PermutationsNumStrDto(n, r) "+
			"n='%v' r='%v' Error='%v'", nInt, rInt, err.Error())
		return
	}

	actualResultStr := result.GetNumStr()

	fmt.Println("       Num Of Items: ", numOfItems)
	fmt.Println("Num Of Items Picked: ", numOfItemsPicked)
	fmt.Println("  Allow Repetitions: ", allowRepetitions)
	fmt.Println(" Permutation Result: ", actualResultStr)

}
