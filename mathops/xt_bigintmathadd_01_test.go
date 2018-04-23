package mathops

import (
	"testing"
	"math/big"
	"errors"
	"fmt"
)

func TestBigIntMathAdd_AddBigInts_01(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := 124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigInts_02(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := 122469.665544
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigInts_03(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigInts_04(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigIntNums_01(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := 124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigIntNums_02(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := 122469.665544
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigIntNums_03(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigIntNums_04(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddBigIntNumArray_01(t *testing.T) {

	numStrs := []string{"45.8",
											"1.45962",
											"58.71",
											"-37.62174",
											"89.8",
										}

	expectedTotalStr := "158.14788"
	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	const lenBigNums = 5

	bNums := make([]BigIntNum, lenBigNums)

	for i:=0; i < lenBigNums; i++ {

		bNums[i], err = BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
				"i='%v' numStrs[i]='%v' Error='%v'",
					i, numStrs[i], err.Error())
		}

	}

  results := BigIntMathAdd{}.AddBigIntNumArray(bNums)

  actualResultNumStr, err := results.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by results.Result.GetNumStrErr(). Error='%v'",
			err.Error())
	}


	if expectedResultNumStr != actualResultNumStr {
  	t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
			expectedResultNumStr, actualResultNumStr)
	}

	if !expectedBNum.Equal(results.Result) {
		t.Errorf("Error: Expected BigIntNum to equal actual BigIntNum. It Did NOT! " +
			"BigIntTotal='%s'. Instead, BigIntTotal='%s'. ",
			expectedBNum.BigInt.Text(10), results.Result.BigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumArray_02(t *testing.T) {

	numStrs := []string{"-978425.648941",
											"33.12",
											"-804.1",
											"32567",
											"-41.859",
										}

	expectedTotalStr := "-946671.487941"
	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	const lenBigNums = 5

	bNums := make([]BigIntNum, lenBigNums)

	for i:=0; i < lenBigNums; i++ {

		bNums[i], err = BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
				"i='%v' numStrs[i]='%v' Error='%v'",
					i, numStrs[i], err.Error())
		}

	}

  results := BigIntMathAdd{}.AddBigIntNumArray(bNums)

  actualResultNumStr, err := results.Result.GetNumStrErr()

  if err != nil {
  	t.Errorf("Error returned by results.Result.GetNumStrErr(). " +
  		"Error='%v' ", err.Error())
	}


  if expectedResultNumStr != actualResultNumStr {
  	t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
			expectedResultNumStr, actualResultNumStr)
	}

	if !expectedBNum.Equal(results.Result) {
		t.Errorf("Error: Expected BigIntNum != actual BigIntNum. " +
			"BigIntTotal='%v'. Instead, BigIntTotal='%v'. ",
			expectedBNum.BigInt.Text(10), results.Result.BigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_01(t *testing.T) {

	var err error

	// multiplier = 0
	addendStr := "5"
	// bNumStrs
	bNumStrs :=  [] string {
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}


	addendBiNum, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(bNumStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(bNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(bNumStrs[i]) " +
				"i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
		}

	}

	result := BigIntMathAdd{}.AddBigIntNumOutputToArray(addendBiNum, bINumArray)

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_02(t *testing.T) {

	var err error

	// multiplier = 0
	addendStr := "3.1"
	// bNumStrs
	bNumStrs :=  [] string {
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	addendBiNum, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(bNumStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(bNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(bNumStrs[i]) " +
				"i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
		}

	}

	result := BigIntMathAdd{}.AddBigIntNumOutputToArray(addendBiNum, bINumArray)

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathAdd_AddBigIntNumSeries_01(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	b1, err := BigIntNum{}.NewNumStr(n1Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'.", n1Str, err.Error())
	}
	
	b2, err := BigIntNum{}.NewNumStr(n2Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'.", n2Str, err.Error())
	}
	
	
	b3, err := BigIntNum{}.NewNumStr(n3Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'.", n3Str, err.Error())
	}
	
	b4, err := BigIntNum{}.NewNumStr(n4Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'.", n4Str, err.Error())
	}
	
	b5, err := BigIntNum{}.NewNumStr(n5Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'.", n5Str, err.Error())
	}

	total := BigIntMathAdd{}.AddBigIntNumSeries(b1, b2, b3, b4, b5)

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedResultNumStr, total.Result.BigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumSeries_02(t *testing.T) {
	n1Str := "-978425.648941"
	n2Str := "33.12"
	n3Str := "-804.1"
	n4Str := "32567"
	n5Str := "-41.859"
	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	b1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'.", n1Str, err.Error())
	}

	b2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'.", n2Str, err.Error())
	}


	b3, err := BigIntNum{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'.", n3Str, err.Error())
	}

	b4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'.", n4Str, err.Error())
	}

	b5, err := BigIntNum{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'.", n5Str, err.Error())
	}

	total := BigIntMathAdd{}.AddBigIntNumSeries(b1, b2, b3, b4, b5)

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedResultNumStr, total.Result.BigInt.Text(10))
	}

}


func TestBigIntMathAdd_AddDecimal_01(t *testing.T) {
	n1Str := "123456.789"

	n2Str := "987.123456"

	expectedNumStr 		:= "124443.912456"
	expectedPrecision := uint(6)
	expectedSign := 1


	dec1, err := Decimal{}.NewNumStr(n1Str)
	
	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)
	
	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}


	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStr(). " +
			"Error='%v' ", err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'.  Instead, Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Result Sign='%v'. Instead, Sign='%v'. ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddDecimal_02(t *testing.T) {
	n1Str := "123456.789"

	n2Str := "-987.123456"

	expectedNumStr 		:= "122469.665544"
	expectedPrecision := uint(6)
	expectedSign := 1


	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}


	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStr(). " +
			"Error='%v' ", err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'.  Instead, Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Result Sign='%v'. Instead, Sign='%v'. ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddDecimal_03(t *testing.T) {
	n1Str := "-123456.789"

	n2Str := "987.123456"

	expectedNumStr 		:= "-122469.665544"
	expectedPrecision := uint(6)
	expectedSign := -1


	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}


	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStr(). " +
			"Error='%v' ", err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'.  Instead, Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Result Sign='%v'. Instead, Sign='%v'. ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddDecimalArray_01(t *testing.T) {

	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	decAry := make([]Decimal, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		dec, err := Decimal{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				fmt.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		decAry[i] = dec

	}

	total, err := BigIntMathAdd{}.AddDecimalArray(decAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddDecimalArray(decAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddDecimalArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	decAry := make([]Decimal, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		dec, err := Decimal{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				fmt.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		decAry[i] = dec

	}

	total, err := BigIntMathAdd{}.AddDecimalArray(decAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddDecimalArray(decAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddDecimalOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// decNumStrs
	decNumStrs :=  [] string {
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}


	decAddend, err := Decimal{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(decNumStrs)
	decArray := make([]Decimal, lenArray)


	for i:=0; i < lenArray; i++ {

		decArray[i], err = 	Decimal{}.NewNumStr(decNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) " +
				"i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddDecimalOutputToArray(decAddend, decArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalOutputToArray(" +
			"decAddend, decArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddDecimalOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// decNumStrs
	decNumStrs :=  [] string {
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	decAddend, err := Decimal{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(decNumStrs)
	decsArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decsArray[i], err =	Decimal{}.NewNumStr(decNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) " +
				"i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddDecimalOutputToArray(decAddend, decsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalOutputToArray(" +
			"decAddend, decsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddDecimalSeries_01(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	dec1, err := Decimal{}.NewNumStr(n1Str)
	
	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " + 
			"n1Str='%v' Error='%v'. ", 
				n1Str, err.Error())
	}
	
	dec2, err := Decimal{}.NewNumStr(n2Str)
	
	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). " + 
			"n2Str='%v' Error='%v'. ", 
				n2Str, err.Error())
	}
	
	dec3, err := Decimal{}.NewNumStr(n3Str)
	
	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). " + 
			"n3Str='%v' Error='%v'. ", 
				n3Str, err.Error())
	}
	
	dec4, err := Decimal{}.NewNumStr(n4Str)
	
	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). " + 
			"n4Str='%v' Error='%v'. ", 
				n4Str, err.Error())
	}
	
	dec5, err := Decimal{}.NewNumStr(n5Str)
	
	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " + 
			"n5Str='%v' Error='%v'. ", 
				n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddDecimalSeries(dec1, " +
			"dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddDecimalSeries_02(t *testing.T) {
	n1Str := "-978425.648941"
	n2Str := "33.12"
	n3Str := "-804.1"
	n4Str := "32567"
	n5Str := "-41.859"
	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ",
				n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
				n2Str, err.Error())
	}

	dec3, err := Decimal{}.NewNumStr(n3Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
				n3Str, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(n4Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
				n4Str, err.Error())
	}

	dec5, err := Decimal{}.NewNumStr(n5Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
				n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddDecimalSeries(dec1, " +
			"dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStrErr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddINumMgr_01(t *testing.T) {
	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&ia1, &dec2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddINumMgr_02(t *testing.T) {

	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&ia1, &dec2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}


	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}


	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddINumMgr_03(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	nDto1, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(
										nDto1.GetThisPointer(),
										dec2.GetThisPointer())

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(nDto1.GetThisPointer(), " +
			"dec2.GetThisPointer()). Error='%v' ",
			err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}
}

func TestBigIntMathAdd_AddINumMgr_04(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	nDto1, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&nDto1, &ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&nDto1, &ia2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if b2Precision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			b2Precision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddINumMgrArray_01(t *testing.T) {

	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	inumMgrAry := make([]INumMgr, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		if i < 2 {
			dec, err := Decimal{}.NewNumStr(numStrAry[i])

			if err != nil {
					fmt.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
						"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &dec
		} else if i > 1 && i < 4 {
			nDto, err := NumStrDto{}.NewNumStr(numStrAry[i])

			if err != nil {
				fmt.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &nDto

		} else {
			// i must be 4
			ia, err := IntAry{}.NewNumStr(numStrAry[i])

			if err != nil {
				fmt.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &ia
		}

	}

	total, err := BigIntMathAdd{}.AddINumMgrArray(inumMgrAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrArray(inumMgrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddINumMgrArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	inumMgrAry := make([]INumMgr, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		if i < 2 {
			dec, err := Decimal{}.NewNumStr(numStrAry[i])

			if err != nil {
				fmt.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &dec
		} else if i > 1 && i < 4 {
			nDto, err := NumStrDto{}.NewNumStr(numStrAry[i])

			if err != nil {
				fmt.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &nDto

		} else {
			// i must be 4
			ia, err := IntAry{}.NewNumStr(numStrAry[i])

			if err != nil {
				fmt.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &ia
		}

	}

	total, err := BigIntMathAdd{}.AddINumMgrArray(inumMgrAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrArray(inumMgrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddINumMgrSeries_01(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	nDto2, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		fmt.Errorf("Error returned by  NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(n3Str)

	if err != nil {
		fmt.Errorf("Error returned by IntAry{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	bigINum4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		fmt.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	dec5, err := Decimal{}.NewNumStr(n5Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrSeries(&dec1, " +
			"&nDto2, &ia3, &bigINum4, &dec5). Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}
}

func TestBigIntMathAdd_AddINumMgrSeries_02(t *testing.T) {
	n1Str := "-978425.648941"
	n2Str := "33.12"
	n3Str := "-804.1"
	n4Str := "32567"
	n5Str := "-41.859"
	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		fmt.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(n3Str)

	if err != nil {
		fmt.Errorf("Error returned by NumStrDto{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(n4Str)

	if err != nil {
		fmt.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	bINum5, err := BigIntNum{}.NewNumStr(n5Str)

	if err != nil {
		fmt.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddINumMgrSeries(&dec1, &ia2, &nDto3, &dec4, &bINum5)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddDecimalSeries(&dec1, " +
			"&ia2, &nDto3, &dec4, &bINum5). Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStrErr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}
}

func TestBigIntMathAdd_AddIntAry_01(t *testing.T) {
	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddIntAry_02(t *testing.T) {

	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}


	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}


	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}


}

func TestBigIntMathAdd_AddIntAry_03(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}
}

func TestBigIntMathAdd_AddIntAry_04(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if b2Precision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			b2Precision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddIntArySeries_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				fmt.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntArySeries(
									iaArray[0],
									iaArray[1],
									iaArray[2],
									iaArray[3],
									iaArray[4],)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddIntArySeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntArySeries_02(t *testing.T) {
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				fmt.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntArySeries(
									iaArray[0],
									iaArray[1],
									iaArray[2],
									iaArray[3],
									iaArray[4],)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddIntArySeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntAryArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				fmt.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntAryArray(iaArray)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddIntAryArray(iaArray). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntAryArray_02(t *testing.T) {
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				fmt.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntAryArray(iaArray)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddIntAryArray(iaArray). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}


	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntAryOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// iaNumStrs
	iaNumStrs :=  [] string {
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}


	iaAddend, err := IntAry{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)


	for i:=0; i < lenArray; i++ {

		iaArray[i], err = 	IntAry{}.NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(iaNumStrs[i]) " +
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryOutputToArray(" +
			"iaAddend, iaArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddIntAryOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// iaNumStrs
	iaNumStrs :=  [] string {
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	iaAddend, err := IntAry{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(iaNumStrs)
	decsArray := make([]IntAry, lenArray)

	for i:=0; i < lenArray; i++ {

		decsArray[i], err =	IntAry{}.NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(iaNumStrs[i]) " +
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddIntAryOutputToArray(iaAddend, decsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryOutputToArray(" +
			"iaAddend, decsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddNumStrSeries_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}


	total, err := BigIntMathAdd{}.AddNumStrSeries(
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4],)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrSeries_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4],)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStr_01(t *testing.T) {

	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddNumStr_02(t *testing.T) {

	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result,err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}


	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}


	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddNumStr_03(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}


	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddNumStr_04(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}


	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}


	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if b2Precision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			b2Precision, result.Input.Big2.Precision)
	}


	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddNumStrDto_01(t *testing.T) {

	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}

}

func TestBigIntMathAdd_AddNumStrDto_02(t *testing.T) {

	n1Str := "123456.789"
	b1ReconciledStr := "123456789000"
	b1Precision := uint(3)

	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1ReconciledStr, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}
}

func TestBigIntMathAdd_AddNumStrDto_03(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "987.123456"
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			maxPrecision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}
}

func TestBigIntMathAdd_AddNumStrDto_04(t *testing.T) {

	n1Str := "-123456.789"
	b1ReconciledStr := "-123456789000"
	b1Precision := uint(3)
	b1ReconciledBig, oK := big.NewInt(0).SetString(b1ReconciledStr, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	n2Str := "-987.123456"
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		errors.New("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if b1ReconciledBig.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1ReconciledBig.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if maxPrecision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			maxPrecision, result.Input.Big1.Precision)
	}

	if b2Big.Cmp(result.Input.Big2.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig2='%v'. Instead, InputBig2='%v'. " +
			b2Big.Text(10), result.Input.Big2.BigInt.Text(10))
	}

	if b2Precision != result.Input.Big2.Precision {
		t.Errorf("Error: Expected InputBig2Precision='%v'.  Instead, " +
			"InputBig2Precision='%v' ",
			b2Precision, result.Input.Big2.Precision)
	}

	if biExpectedResult.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedPrecision != result.Result.Precision {
		t.Errorf("Error: Expected Result Precision='%v'. Instead, Result Precision='%v'. ",
			expectedPrecision, result.Result.Precision)
	}

	if expectedSign != result.Result.Sign {
		t.Errorf("Error: Expected Recult Sign='%v'. Instead, Result Sign='%v' ",
			expectedSign, result.Result.Sign)
	}
}

func TestBigIntMathAdd_AddNumStrDtoArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). " +
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrDtoArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}
	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). " +
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}


func TestBigIntMathAdd_AddNumStrDtoSeries_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoSeries(
					numStrDtoAry[0],
					numStrDtoAry[1],
					numStrDtoAry[2],
					numStrDtoAry[3],
					numStrDtoAry[4],	)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrDtoSeries_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr, err := expectedBNum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by expectedBNum.GetNumStr(). Error='%v'",
			err.Error())
	}
	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4],	)

	if err != nil {
		fmt.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total.Result) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.BigInt.Text(10), total.Result.BigInt.Text(10))
	}

	actualTotalNumstr, err := total.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by total.Result.GetNumStr() " +
			"Error='%v' ", err.Error())
	}

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

