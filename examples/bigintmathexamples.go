package examples

import (
	"../mathops"
	"fmt"
	"math/big"
)


func ExampleBigIntAdd_01() {
	//n1Str := "0.000123"

	//n2Str := ".001"

	b1 := big.NewInt(123)
	b2 := big.NewInt(1)

	expectedResultStr:= "1123"
	expectedPrecision := uint(6)

	result := mathops.BigIntMathAdd{}.AddBigInts(b1, 6, b2, 3 )


	ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

}

func ExampleBigIntAddNumStr_01() {
	n1Str := "0.000123"


	n2Str := "1"


	// Result = 	1.000123
	expectedResultStr := "1000123"
	expectedPrecision := uint(6)


	result, err := mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
		return
	}

	ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

	fmt.Println("Successful Completion")
}

func ExampleBigIntAddNumStr_02() {
	n1Str := "0.000123"


	n2Str := ".001"


	// Result = 	"0.001123"
	expectedResultStr := "001123"
	expectedPrecision := uint(6)


	result, err := mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
		return
	}

	ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

}

func ExamplePrintBasicMathResult(expectedResultStr string,
																	expectedPrecision uint,
																		result mathops.BigIntBasicMathResult) {

  getNumStr, _ := result.Result.GetNumStr()

	fmt.Println("            expected result: ", expectedResultStr)
	fmt.Println("              result.BigInt: ", result.Result.BigInt.Text(10))
	fmt.Println("         result.GetNumStr(): ", getNumStr)
	fmt.Println("")
	fmt.Println("         Expected Precision: ", expectedPrecision)
	fmt.Println("           result.Precision: ", result.Result.Precision)
	fmt.Println("")
	fmt.Println("          result.Input.Big1: ", result.Input.Big1.BigInt.Text(10))
	fmt.Println("result.Input.Big1.Precision: ", result.Input.Big1.Precision)
	fmt.Println("")
	fmt.Println("          result.Input.Big2: ", result.Input.Big2.BigInt.Text(10))
	fmt.Println("result.Input.Big2.Precision: ", result.Input.Big2.Precision)


}
