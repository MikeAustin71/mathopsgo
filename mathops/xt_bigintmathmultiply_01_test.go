package mathops

import (
	"testing"
	"math/big"
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.BigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.BigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
			bIMultiplier,
				multiplierPrecision,
					biMultiplicand,
						multiplicandPrecision)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

  err =	iaMultiplier.Multiply(
  										&iaMultiplier,
  											&iaMultiplicand,
  												&iaResult,
  												-1,
  													-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.BigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.BigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
			bIMultiplier,
				multiplierPrecision,
					biMultiplicand,
						multiplicandPrecision)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

  err =	iaMultiplier.Multiply(
  										&iaMultiplier,
  											&iaMultiplicand,
  												&iaResult,
  												-1,
  													-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.BigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.BigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
			bIMultiplier,
				multiplierPrecision,
					biMultiplicand,
						multiplicandPrecision)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.BigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.BigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
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
	expectedBigINumStr := "0.000000"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		6,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.BigInt)
	multiplierPrecision := multiplierBiNum.GetPrecisionUint()
	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.BigInt)
	multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

	result := BigIntMathMultiply{}.MultiplyBigInts(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

  err =	iaMultiplier.Multiply(
  										&iaMultiplier,
  											&iaMultiplicand,
  												&iaResult,
  												-1,
  													-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

  err =	iaMultiplier.Multiply(
  										&iaMultiplier,
  											&iaMultiplicand,
  												&iaResult,
  												-1,
  													-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
						&iaMultiplier,
							&iaMultiplicand,
								&iaResult,
									-1,
										-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNums_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0.000000"

	expectedBigINumSign := 1

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		6,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	bINumArray := make([]BigIntNum, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := bINumArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by bINumArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathMultiply{}.MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

	actualNumStr, err := result.Result.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by result.Result.GetNumStrErr() " +
			"Error='%v'. ", err.Error())
	}

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

