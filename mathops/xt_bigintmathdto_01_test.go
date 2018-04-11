package mathops

import (
	"testing"
	"math/big"
	"errors"
)

func TestBigIntMathDto_AddBigInts_01(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathDto_AddBigInts_02(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathDto_AddBigInts_03(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathDto_AddBigInts_04(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathDto_AddBigIntNums_01(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathDto_AddBigIntNums_02(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathDto_AddBigIntNums_03(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathDto_AddBigIntNums_04(t *testing.T) {

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

	result := BigIntMathDto{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathDto_AddNumStr_01(t *testing.T) {

	n1Str := "123456.789"
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

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

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result, err := BigIntMathDto{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDto{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1Big.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1Big.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if b1Precision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			b1Precision, result.Input.Big1.Precision)
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

func TestBigIntMathDto_AddNumStr_02(t *testing.T) {

	n1Str := "123456.789"
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

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

	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result,err := BigIntMathDto{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDto{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1Big.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1Big.Text(10), result.Input.Big1.BigInt.Text(10))
	}


	if b1Precision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			b1Precision, result.Input.Big1.Precision)
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

func TestBigIntMathDto_AddNumStr_03(t *testing.T) {

	n1Str := "-123456.789"
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

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
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result, err := BigIntMathDto{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDto{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1Big.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1Big.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if b1Precision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			b1Precision, result.Input.Big1.Precision)
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

func TestBigIntMathDto_AddNumStr_04(t *testing.T) {

	n1Str := "-123456.789"
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

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
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result, err := BigIntMathDto{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDto{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
	}

	if b1Big.Cmp(result.Input.Big1.BigInt) != 0 {
		t.Errorf("Error: Expected InputBig1='%v'. Instead, InputBig1='%v'. " +
			b1Big.Text(10), result.Input.Big1.BigInt.Text(10))
	}

	if b1Precision != result.Input.Big1.Precision {
		t.Errorf("Error: Expected InputBig1Precision='%v'.  Instead, " +
			"InputBig1Precision='%v' ",
			b1Precision, result.Input.Big1.Precision)
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
