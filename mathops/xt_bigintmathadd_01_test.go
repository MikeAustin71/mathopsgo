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
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
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
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
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
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
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
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
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
