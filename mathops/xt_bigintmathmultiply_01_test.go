package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathMultiply_MultiplyBigInts_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedBigINumStr := "2875.94572"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathMultiply_MultiplyBigInts_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedBigINumStr := "30987680500513189125.14259702468435"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigInts_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedBigINumStr := "-30987680500513189125.14259702468435"

	expectedBigINumSign := -1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigInts_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedBigINumStr := "22197234145.3632"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigInts_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNums_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedBigINumStr := "2875.94572"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNums_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedBigINumStr := "30987680500513189125.14259702468435"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}
}

func TestBigIntMathMultiply_MultiplyBigIntNums_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedBigINumStr := "-30987680500513189125.14259702468435"

	expectedBigINumSign := -1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNums_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedBigINumStr := "22197234145.3632"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNums_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0
	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	iaResult.OptimizeIntArrayLen(true)

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNums_06(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875,94572"

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_01(t *testing.T) {
	num1Int := 242
	precision := uint(2)
	expectedNum1 := "2.42"
	expectedResult:= "4.84"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwo(num1)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_02(t *testing.T) {
	num1Int := 921123
	precision := uint(3)
	expectedNum1 := "921.123"
	expectedResult:= "1842.246"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwo(num1)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_03(t *testing.T) {
	num1Int := -4751
	precision := uint(1)
	expectedNum1 := "-475.1"
	expectedResult:= "-950.2"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwo(num1)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_04(t *testing.T) {
	num1Int := 475
	precision := uint(0)
	expectedNum1 := "475"
	expectedResult:= "950"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwo(num1)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_05(t *testing.T) {
	num1Int := 0
	precision := uint(0)
	expectedNum1 := "0"
	expectedResult:= "0"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwo(num1)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_01(t *testing.T) {
	num1Int := 242
	precision := uint(2)
	exponent := uint(1)
	expectedNum1 := "2.42"
	expectedResult:= "4.84"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwoToPower(num1, exponent)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_02(t *testing.T) {
	num1Int := 242
	precision := uint(2)
	exponent := uint(2)
	expectedNum1 := "2.42"
	expectedResult:= "9.68"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwoToPower(num1, exponent)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_03(t *testing.T) {
	num1Int := 97123
	precision := uint(3)
	exponent := uint(4)
	expectedNum1 := "97.123"
	expectedResult:= "1553.968"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwoToPower(num1, exponent)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_04(t *testing.T) {
	num1Int := -97123
	precision := uint(3)
	exponent := uint(4)
	expectedNum1 := "-97.123"
	expectedResult:= "-1553.968"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwoToPower(num1, exponent)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_05(t *testing.T) {
	num1Int := 7
	precision := uint(0)
	exponent := uint(10)
	expectedNum1 := "7"
	expectedResult:= "7168"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwoToPower(num1, exponent)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_06(t *testing.T) {
	num1Int := 0
	precision := uint(0)
	exponent := uint(10)
	expectedNum1 := "0"
	expectedResult:= "0"

	num1 := BigIntNum{}.NewInt(num1Int, precision)

	result := BigIntMathMultiply{}.MultiplyBigIntNumByTwoToPower(num1, exponent)

	if expectedResult != result.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResult, result.GetNumStr())
	}

	if expectedNum1 != num1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNum1, num1.GetNumStr())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	expectedBigINumStr := "128"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedBigINumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedBigINumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	iaResult.OptimizeIntArrayLen(true)

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs := []string{
		"1.879",
		"3.824",
		"21.756",
		"2.1234567",
		"6",
		"2",
	}

	// product = -20408.5138429311978576052224
	expectedBigINumStr := "-20408.5138429311978576052224"

	expectedBigINumSign := -1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_05(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664,26376575446779648
	expectedNumStr := "11995826664,26376575446779648"

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"2",
		"4",
		"6",
		"8",
		"10",
		"12",
	}

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_02(t *testing.T) {

	var err error

	// multiplier = 8
	multiplierStr := "8"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"800.8",
		"-208",
		"31.392",
		"64",
		"42376.984",
		"-39.168",
	}

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_03(t *testing.T) {

	var err error

	// multiplier = -31.2
	multiplierStr := "-31.2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"-3123.12",
		"811.2",
		"-122.4288",
		"-249.6",
		"-165270.2376",
		"152.7552",
	}

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_04(t *testing.T) {

	var err error

	// multiplier = 283
	multiplierStr := "283"
	// multiplicandStrs
	multiplicandStrs := []string{
		"0",
		"-26",
		"0",
		"8",
		"5297.123",
		"0",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"0",
		"-7358",
		"0",
		"2264",
		"1499085.809",
		"0",
	}

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_05(t *testing.T) {

	var err error

	// multiplier = 0
	multiplierStr := "0"
	// multiplicandStrs
	multiplicandStrs := []string{
		"5",
		"-26",
		"9",
		"8",
		"5297.123",
		"37",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_06(t *testing.T) {

	var err error

	// multiplier = -31.2
	multiplierStr := "-31.2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"-3123,12",
		"811,2",
		"-122,4288",
		"-249,6",
		"-165270,2376",
		"152,7552",
	}

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. Index='%v'",
				expectedNumSeps.String(), actualNumSeps.String(), j)
		}

	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	expectedBigINumStr := "128"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumSeries(
		multiplierBiNum,
		bINumArray[0],
		bINumArray[1],
		bINumArray[2],
		bINumArray[3],
		bINumArray[4],
		bINumArray[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedBigINumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumSeries(
		multiplierBiNum,
		bINumArray[0],
		bINumArray[1],
		bINumArray[2],
		bINumArray[3],
		bINumArray[4],
		bINumArray[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedBigINumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumSeries(
		multiplierBiNum,
		bINumArray[0],
		bINumArray[1],
		bINumArray[2],
		bINumArray[3],
		bINumArray[4],
		bINumArray[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	iaResult.OptimizeIntArrayLen(true)

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs := []string{
		"1.879",
		"3.824",
		"21.756",
		"2.1234567",
		"6",
		"2",
	}

	// product = -20408.5138429311978576052224
	expectedBigINumStr := "-20408.5138429311978576052224"

	expectedBigINumSign := -1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumSeries(
		multiplierBiNum,
		bINumArray[0],
		bINumArray[1],
		bINumArray[2],
		bINumArray[3],
		bINumArray[4],
		bINumArray[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_05(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedNumStr := "2212352,1767579232"

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		bINumArray[i], err = BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumSeries(
		multiplierBiNum,
		bINumArray[0],
		bINumArray[1],
		bINumArray[2],
		bINumArray[3],
		bINumArray[4],
		bINumArray[5])

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
