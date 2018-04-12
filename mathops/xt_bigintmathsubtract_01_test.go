package mathops

import (
	"testing"
	"math/big"
)

func TestBigIntMathSubtract_SubtractBigIntNums_01(t *testing.T) {

	minuendStr := "12332"
	minuendPrecision := uint(2)

	subtrahendStr := "23321"
	subtrahendPrecision := uint(3)

	expectedBigINumStr := "99999"
	expectedBigINumPrecision := uint(3)
	expectedBigINumSign := 1

	bMinuend, oK := big.NewInt(0).SetString(minuendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(minuendStr, 10) " +
			"minuendStr='%v' ", minuendStr)
	}

	bSubtrahend, oK := big.NewInt(0).SetString(subtrahendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(subtrahendStr, 10) " +
			"subtrahendStr='%v' ", subtrahendStr)
	}

	expectedBigI, oK := big.NewInt(0).SetString(expectedBigINumStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(expectedBigINumStr, 10) " +
			"expectedBigINumStr='%v' ", expectedBigINumStr)
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	subtrahendBiNum := BigIntNum{}.NewBigInt(bSubtrahend, subtrahendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

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
