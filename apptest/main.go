package main

import (
	"../mathops"
	"fmt"
)

func main() {

	nStr := "123,456"

	expectedNumSeps := mathops.NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	n1 := mathops.NumStrDto{}

	n1.SetNumericSeparatorsDto(expectedNumSeps)

	n2, err := n1.ParseNumStr(nStr)

	if err != nil {
		fmt.Println("Error returned by n1.ParseNumStr(nStr). Error=", err.Error())
		return
	}

	fmt.Println(" n2 Result= ", n2.GetNumStr())
	fmt.Println("Frac Runes= ", string(n2.GetAbsFracRunes()))
	fmt.Println(" Int Runes= ", string(n2.GetAbsIntRunes()))
}



