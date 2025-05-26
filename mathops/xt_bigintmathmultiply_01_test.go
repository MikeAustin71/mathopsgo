package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathMultiply_BigIntMultiply_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedBigINumStr := "2875.94572"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'\n Error='%v'\n\n", multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'\nError='%v'\n\n", expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathMultiply).\n"+
			"BigIntMultiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(result) != 0 {
		t.Errorf("Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n\n",
			expectedBigINumBigInt.Text(10), result.Text(10))
		return
	}

	expectedBigINumPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	uintResultPrecision := uint(resultPrecision.Uint64())

	if expectedBigINumPrecision != uintResultPrecision {
		t.Errorf("Error: Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			expectedBigINumPrecision, uintResultPrecision)
		return
	}
}

func TestBigIntMathMultiply_BigIntMultiply_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedBigINumStr := "30987680500513189125.14259702468435"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by:"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by:"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.\n"+
			"  GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.\n"+
			"  GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumInt.Cmp(result) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
			expectedBigINumInt.Text(10), result.Text(10))
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumPrecisionBigInt, err :=\n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecision) != 0 {
		t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
			expectedBigINumPrecisionBigInt.Text(10), resultPrecision.Text(10))
		return
	}
}

func TestBigIntMathMultiply_BigIntMultiply_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedBigINumStr := "-30987680500513189125.14259702468435"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n",
			multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStr(multiplicandStr)\n\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by:"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathMultiply).\n"+
			"BigIntMultiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumInt.Cmp(result) != 0 {
		t.Errorf("Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n\n",
			expectedBigINumInt.Text(10), result.Text(10))
		return
	}

	expectedBigINumPrecisionInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumPrecisionInt, err := expectedBigINum.\n"+
			"  GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumPrecisionInt.Cmp(resultPrecision) != 0 {
		t.Errorf("Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			expectedBigINumPrecisionInt.Text(10), resultPrecision.Text(10))

	}

	return
}

func TestBigIntMathMultiply_BigIntMultiply_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedBigINumStr := "22197234145.3632"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'\nError='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'\n"+
			"\nError='%v'\n\n", multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'\n"+
			"\nError='%v'\n\n", multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n", multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplier.Multiply(...)"+
			"\nError='%v'\n\n", err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\n"+
			"\nError='%v'\n\n", expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.\n"+
			"  GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumInt, err := expectedBigINum.\n"+
			"  GetIntegerValue()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumInt.Cmp(result) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
			expectedBigINumInt.Text(10), result.Text(10))
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumPrecisionBigInt, err := expectedBigINum.\n"+
			"  GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecision) != 0 {
		t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
			expectedBigINumPrecisionBigInt.Text(10), resultPrecision.Text(10))
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiply_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"\nError='%v'\n\n", multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"\nError='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"\nError='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"\terr = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return

	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\n"+
			"Error='%v'\n\n", expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathMultiply).\n"+
			"BigIntMultiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumInt.Cmp(result) != 0 {
		t.Errorf("Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n",
			expectedBigINumInt.Text(10), result.Text(10))
		return
	}

	expectedBigINumBigIPrecision, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumBigIPrecision, err :=\n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumBigIPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
			expectedBigINumBigIPrecision, resultPrecision.Text(10))
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiply_06(t *testing.T) {

	multiplierStr := "0.00000"

	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n", multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\n"+
			"Error='%v'\n\n", expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"multiplicandPrecision := multiplicandBiNum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiply(...)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(result) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
			expectedBigINumBigInt.Text(10), result.Text(10))
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecision) != 0 {
		t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
			expectedBigINumPrecisionBigInt, resultPrecision.Text(10))
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_01"

	expectedNumStr := "105.6752"
	num := 1056752
	precision := uint(4)
	exponent := big.NewInt(0)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).\n"+
			"  NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_02"

	expectedNumStr := "-105.6752"
	num := -1056752
	precision := uint(4)
	exponent := big.NewInt(0)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n+"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).\n"+
			"  NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'.\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_03"

	expectedNumStr := "1056.752"
	num := 1056752
	precision := uint(4)
	exponent := big.NewInt(1)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err :=new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).\n"+
			"  NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_04"

	expectedNumStr := "-1056.752"
	num := -1056752
	precision := uint(4)
	exponent := big.NewInt(1)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).\n"+
			"  NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"  actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_05"

	expectedNumStr := "10567.52"
	num := 1056752
	precision := uint(4)
	exponent := big.NewInt(2)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_06"

	expectedNumStr := "-10567.52"
	num := -1056752
	precision := uint(4)
	exponent := big.NewInt(2)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt, \n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).\n"+
			"  NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_07(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_07"

	expectedNumStr := "10567520000"
	num := 1056752
	precision := uint(4)
	exponent := big.NewInt(8)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}
	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_08(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_08"

	expectedNumStr := "0"
	num := 0
	precision := uint(4)
	exponent := big.NewInt(5)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}
	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_09(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_09"

	expectedNumStr := "10.56752"
	num := 1056752
	precision := uint(0)
	exponent := big.NewInt(-5)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_10(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_10"

	expectedNumStr := "0.01056752"
	num := 1056752
	precision := uint(3)
	exponent := big.NewInt(-5)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'\nInstead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_11(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyByTenToPwr_11"

	expectedNumStr := "-0.01056752"
	num := -1056752
	precision := uint(3)
	exponent := big.NewInt(-5)

	fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

	fixDecBigInt, err := fixDec.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecBigInt, err := fixDec.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDecPrecisionBigInt, err := fixDec.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
			fixDecBigInt, fixDecPrecisionBigInt, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTenToPower(fixDecBigInt,\n"+
			"    fixDecPrecisionBigInt, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'\n\n",
			ePrefix, expectedNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_01"

	// multiplicand = 23.321
	multiplicandBInt := big.NewInt(23321)
	multiplicandPrecision := big.NewInt(3)
	exponent := uint(5)
	expectedResult := "746.272"

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			multiplicandBInt, multiplicandPrecision, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTwoToPower(multiplicandBInt,\n"+
			"    multiplicandPrecision, exponent)n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result='%v'.\n\n",
			ePrefix, expectedResult, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_02"

	// multiplicand = 8
	multiplicandBInt := big.NewInt(8)
	multiplicandPrecision := big.NewInt(0)
	exponent := uint(10)
	expectedResult := "8192"

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			multiplicandBInt, multiplicandPrecision, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTwoToPower(multiplicandBInt,\n"+
			"    multiplicandPrecision, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err = result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err = result.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result='%v'.\n\n",
			ePrefix, expectedResult, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_03"

	// multiplicand = 9.871234
	multiplicandBInt := big.NewInt(9871234)
	multiplicandPrecision := big.NewInt(6)
	exponent := uint(1)
	expectedResult := "19.742468"

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			multiplicandBInt, multiplicandPrecision, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTwoToPower(multiplicandBInt,\n"+
			"    multiplicandPrecision, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result='%v'.\n\n",
			ePrefix, expectedResult, resultNumStr)

	}

	return
}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_04"

	// multiplicand = -9.871234
	multiplicandBInt := big.NewInt(-9871234)
	multiplicandPrecision := big.NewInt(6)
	exponent := uint(3)
	expectedResult := "-78.969872"

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			multiplicandBInt, multiplicandPrecision, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTwoToPower(\n"+
			"    tmultiplicandBInt, multiplicandPrecision, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result='%v'.\n\n",
			ePrefix, expectedResult, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_05"

	// multiplicand = 8
	multiplicandBInt := big.NewInt(8)
	multiplicandPrecision := big.NewInt(0)
	exponent := uint(0)
	expectedResult := "8"

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			multiplicandBInt, multiplicandPrecision, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := new(BigIntMathMultiply).\n"+
			"  BigIntMultiplyByTwoToPower(\n"+
			"    multiplicandBInt, multiplicandPrecision, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
			"    product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result='%v'.\n\n",
			ePrefix, expectedResult, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_06"

	// (0.12345 x 2^15 = 4045.2096)
	multiplicandBInt := big.NewInt(12345)
	multiplicandPrecision := big.NewInt(5)
	exponent := uint(15)
	expectedResult := "4045.2096"

	product, productPrecision, err :=
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			multiplicandBInt, multiplicandPrecision, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"product, productPrecision, err := \n"+
			"  new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(\n"+
			"    multiplicandBInt, multiplicandPrecision, exponent)\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result='%v'.\n\n",
			ePrefix, expectedResult, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_01"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedBigINumStr := "2875.94572"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBINumBigInt, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBINumBigInt, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBINumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBINumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBINumBigInt, err := multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBINumBigInt, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBINumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBINumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBINumBigInt,
			multiplierBINumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := := new(BigIntFixedDecimal).New(\n"+
			"  multiplierBINumBigInt, multiplierBINumUintPrecision)\n"+
			"multiplierBINumBigInt= '%v'\n"+
			"multiplierBINumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBINumBigInt.Text(10),
			multiplierBINumUintPrecision,
			err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBINumBigInt,
			multiplicandBINumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).New(\n"+
			"  multiplicandBINumBigInt, multiplicandBINumUintPrecision)\n"+
			"multiplicandBiNumBigIntNum= '%v'\n"+
			"multiplicandBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBINumBigInt.Text(10),
			multiplicandBINumUintPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  FixedDecimalMultiply(\n"+
			"    multiplier,multiplicand)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigIntNumBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigIntNumBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigIntNumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigIntNumBigInt.Text(10))
		return
	}

	expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumUintPrecision, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumUintPrecision, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumUintPrecision != resultBigINumUintPrecision {
		t.Errorf("%v\n"+
			"Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			ePrefix, expectedBigINumUintPrecision, resultBigINumUintPrecision)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_02"

	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedBigINumStr := "30987680500513189125.14259702468435"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBiNumBigIntNum, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumBigIntNum, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBiNumBigIntNum,
			multiplierBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := new(BigIntFixedDecimal).New(\n"+
			"  multiplierBiNumBigIntNum, multiplierBiNumUintPrecision)\n"+
			"multiplierBiNumBigIntNum= '%v'\n"+
			"multiplierBiNumUintPrecision= '%v'\nError='%v'\n\n",
			ePrefix,
			multiplierBiNumBigIntNum.Text(10),
			multiplierBiNumUintPrecision,
			err.Error())
		return
	}

	multiplicandBiNumBigIntNum, err :=
		multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" multiplicandBiNumBigIntNum, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumUintPrecision, err :=
		multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBiNumBigIntNum,
			multiplicandBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).New(\n"+
			"  multiplicandBiNumBigIntNum, multiplicandBiNumUintPrecision)\n"+
			"multiplicandBiNumBigIntNum= '%v'\n"+
			"multiplicandBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBiNumBigIntNum.Text(10),
			multiplicandBiNumUintPrecision,
			err.Error())

		return
	}

	multiplierBINumStr, err := multiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBINumStr, err := multiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBINumStr, err := multiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBINumStr, err := multiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).FixedDecimalMultiply(\n"+
			"    multiplier,multiplicand)\n"+
			"multiplier= '%v'\n"+
			"multiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBINumStr,
			multiplicandBINumStr,
			err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBitINumBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBitINumBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBitINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result= '%s'. ",
			ePrefix,
			expectedBigINumBigInt.Text(10),
			resultBitINumBigInt.Text(10))
		return
	}

	expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultUintPrecision, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultUintPrecision, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumUintPrecision != resultUintPrecision {
		t.Errorf("%v\n"+
			"Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			ePrefix, expectedBigINumUintPrecision, resultUintPrecision)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_03"

	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedBigINumStr := "-30987680500513189125.14259702468435"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBiNumBigInt,
			multiplierBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplierBiNumBigInt, multiplierBiNumUintPrecision)\n"+
			"multiplierBiNumBigInt= '%v'\n"+
			"multiplierBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBiNumBigInt.Text(10),
			multiplierBiNumUintPrecision,
			err.Error())

		return
	}

	multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumUintPrecision, err := \n"+
			"  multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBiNumBigInt,
			multiplicandBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplicandBiNumBigInt, multiplicandBiNumUintPrecision)\n"+
			"multiplicandBiNumBigInt= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBiNumBigInt.Text(10),
			multiplicandBiNumUintPrecision,
			err.Error())

		return
	}

	multiplierNumStr, err := multiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStr, err := multiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStr, err := multiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStr, err := multiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  FixedDecimalMultiply(multiplier, multiplicand)\n"+
			"multiplier= '%v'\n"+
			"multiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStr,
			multiplicandNumStr,
			err.Error())

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultUintPrecision, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultUintPrecision, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumUintPrecision != resultUintPrecision {
		t.Errorf("%v\n"+
			"Expected result precision= '%v'.\n"+
			"Instead, result precision= '%v'.\n\n",
			ePrefix, expectedBigINumUintPrecision, resultUintPrecision)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_04"

	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedBigINumStr := "22197234145.3632"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBiNumBigInt,
			multiplierBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplierBiNumBigInt, multiplierBiNumUintPrecision)\n"+
			"multiplierBiNumBigInt= '%v'\n"+
			"multiplierBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBiNumBigInt.Text(10),
			multiplierBiNumUintPrecision,
			err.Error())

		return
	}

	multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBiNumBigInt,
			multiplicandBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplicandBiNumBigInt, multiplicandBiNumUintPrecision)\n"+
			"multiplicandBiNumBigInt= '%v'\n"+
			"multiplicandBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBiNumBigInt.Text(10),
			multiplicandBiNumUintPrecision,
			err.Error())

		return
	}

	multiplierNumStr, err := multiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStr, err := multiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStr, err := multiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStr, err := multiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"    FixedDecimalMultiply(multiplier, multiplicand)\n"+
			"multiplier= '%v'\n"+
			"multiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStr,
			multiplicandNumStr,
			err.Error())

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\nInstead, result= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultUintPrecision, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultUintPrecision, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumUintPrecision != resultUintPrecision {
		t.Errorf("%v\n"+
			"Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			ePrefix, expectedBigINumUintPrecision, resultUintPrecision)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_05"

	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBiNumBigInt,
			multiplierBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplierBiNumBigInt, multiplierBiNumUintPrecision)\n"+
			"multiplierBiNumBigInt= '%v'\n"+
			"multiplierBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBiNumBigInt.Text(10),
			multiplierBiNumUintPrecision,
			err.Error())

		return
	}

	multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBiNumBigInt,
			multiplicandBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplicandBiNumBigInt, multiplicandBiNumUintPrecision)\n"+
			"multiplicandBiNumBigInt= '%v'\n"+
			"multiplicandBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBiNumBigInt.Text(10),
			multiplicandBiNumUintPrecision,
			err.Error())

		return
	}

	multiplierNumStr, err := multiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStr, err := multiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStr, err := multiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStr, err := multiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"    FixedDecimalMultiply(multiplier, multiplicand)\n"+
			"multiplier= '%v'\n"+
			"multiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStr,
			multiplicandNumStr,
			err.Error())

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultUintPrecision, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultUintPrecision, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumUintPrecision != resultUintPrecision {
		t.Errorf("%v\n"+
			"Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			ePrefix, expectedBigINumUintPrecision, resultUintPrecision)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_06"

	multiplierStr := "0.00000"

	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBiNumBigInt,
			multiplierBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplierBiNumBigInt, multiplierBiNumUintPrecision)\n"+
			"multiplierBiNumBigInt= '%v'\n"+
			"multiplierBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBiNumBigInt.Text(10),
			multiplierBiNumUintPrecision,
			err.Error())

		return
	}

	multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBiNumBigInt,
			multiplicandBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplicandBiNumBigInt, multiplicandBiNumUintPrecision)\n"+
			"multiplicandBiNumBigInt= '%v'\n"+
			"multiplicandBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBiNumBigInt.Text(10),
			multiplicandBiNumUintPrecision,
			err.Error())

		return
	}

	multiplierNumStr, err := multiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStr, err := multiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStr, err := multiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStr, err := multiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"    FixedDecimalMultiply(multiplier, multiplicand)\n"+
			"multiplier= '%v'\n"+
			"multiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStr,
			multiplicandNumStr,
			err.Error())

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10),
			resultBigInt.Text(10))
	}

	expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumUintPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultUintPrecision, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultUintPrecision, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumUintPrecision != resultUintPrecision {
		t.Errorf("%v\n"+
			"Expected result precision='%v'.\n"+
			"Instead, result precision='%v'\n\n",
			ePrefix, expectedBigINumUintPrecision, resultUintPrecision)
	}

	return
}

func TestBigIntMathMultiply_FixedDecimalMultiply_07(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_FixedDecimalMultiply_07"

	// multiplier = 0.12345
	multiplierStr := "0.12345"

	// multiplicand = 32768
	multiplicandStr := "32768"

	// product = 4045.2096
	expectedBigINumStr := "4045.2096"

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumBigInt, err := multiplierBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumUintPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplier, err :=
		new(BigIntFixedDecimal).New(
			multiplierBiNumBigInt,
			multiplierBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplier, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplierBiNumBigInt, multiplierBiNumUintPrecision)\n"+
			"multiplierBiNumBigInt= '%v'\n"+
			"multiplierBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierBiNumBigInt.Text(10),
			multiplierBiNumUintPrecision,
			err.Error())

		return
	}

	multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumBigInt, err := multiplicandBiNum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumUintPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicand, err :=
		new(BigIntFixedDecimal).New(
			multiplicandBiNumBigInt,
			multiplicandBiNumUintPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicand, err := new(BigIntFixedDecimal).\n"+
			"  New(multiplicandBiNumBigInt, multiplicandBiNumUintPrecision)\n"+
			"multiplicandBiNumBigInt= '%v'\n"+
			"multiplicandBiNumUintPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplicandBiNumBigInt.Text(10),
			multiplicandBiNumUintPrecision,
			err.Error())

		return
	}

	multiplierNumStr, err := multiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStr, err := multiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStr, err := multiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStr, err := multiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).FixedDecimalMultiply(
		multiplier,
		multiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"    FixedDecimalMultiply(multiplier, multiplicand)\n"+
			"multiplier= '%v'\n"+
			"multiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStr,
			multiplicandNumStr,
			err.Error())

		return
	}

	expectedBigINumNumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected result='%v'.\n"+
			"Instead, result= '%v'.\n\n",
			ePrefix, expectedBigINumNumStr, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigInts_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigInts_01"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedBigINumStr := "2875.94572"

	expectedBigINumSign := 1

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision,
		numSepsDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Comparison Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected number sign='%v'.\n"+
			"Instead, number sign='%v'\n\n",
			ePrefix, expectedBigINumSign, result.sign)
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigInts_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigInts_02"

	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedBigINumStr := "30987680500513189125.14259702468435"

	expectedBigINumSign := 1

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision,
		numSepsDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Comparison Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected number sign='%v'.\n"+
			"Instead, number sign='%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected actualNumStr='%v'\n"+
			"Instead, actualNumStr='%v'\n\n",
			ePrefix, iaResultNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigInts_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigInts_03"

	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedBigINumStr := "-30987680500513189125.14259702468435"

	expectedBigINumSign := -1

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	numSepsDto := NumericSeparatorDto{}
	numSepsDto.SetDefaultsIfEmpty()

	result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision,
		numSepsDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Comparison Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected number sign='%v'.\n"+
			"Instead, number sign='%v'\n\n",
			ePrefix,
			expectedBigINumSign, result.sign)
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected actualNumStr='%v'\n"+
			"Instead, actualNumStr='%v'\n\n",
			ePrefix, iaResultNumStr, actualNumStr)

	}

	return
}

func TestBigIntMathMultiply_MultiplyBigInts_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigInts_04"

	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedBigINumStr := "22197234145.3632"

	expectedBigINumSign := 1

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision,
		numSepsDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Comparison Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected number sign='%v'.\n"+
			"Instead, number sign='%v'\n\n",
			ePrefix,
			expectedBigINumSign, resultSignValue)
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResultNumStr, actualNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigInts_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigInts_05"

	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\nError='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\nError='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr='%v'\nError='%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

	multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierPrecision, err := multiplierBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

	multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandPrecision, err := multiplicandBiNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
		bIMultiplier,
		multiplierPrecision,
		biMultiplicand,
		multiplicandPrecision,
		numSepsDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(...)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Comparison Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected number sign='%v'.\n"+
			"Instead, number sign='%v'\n\n",
			ePrefix,
			expectedBigINumSign, result.sign)
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != actualNumStr {
		t.Errorf("%v\n"+
			"Error: Expected actualNumStr='%v'\n"+
			"Instead, actualNumStr='%v'\n\n",
			ePrefix, iaResultNumStr, actualNumStr)
	}

	return
}
