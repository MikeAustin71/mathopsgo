package main

import (
	"../examples"
	"../mathops"
)

func main() {


	intStr := "123"
	fracStr := "456"
	expectedResult := "123.456"
	signVal := 1
	numSeps := mathops.NumericSeparatorDto{}
	numSeps.DecimalSeparator = '.'
	numSeps.ThousandsSeparator = ','
	numSeps.CurrencySymbol = '$'

	examples.ExampleBINumSetFromIntFracStrings(intStr, fracStr, expectedResult, signVal, numSeps)

}



