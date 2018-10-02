package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathSubtract_SubtractPair_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"
	minuendPrecision := uint(2)

	// subtrahend = 23.321
	subtrahendStr := "23.321"
	subtrahendPrecision := uint(3)

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"
	minuendPrecision := uint(4)

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"
	subtrahendPrecision := uint(2)

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"
	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"
	minuendPrecision := uint(8)

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"
	subtrahendPrecision := uint(3)

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"
	expectedBigINumSign := -1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"
	minuendPrecision := uint(6)

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"
	subtrahendPrecision := uint(3)

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"
	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_05(t *testing.T) {
	// minuend = 0
	minuendStr := "0"
	minuendPrecision := uint(0)

	// subtrahend = 0
	subtrahendStr := "0"
	subtrahendPrecision := uint(0)

	// result = 0
	expectedBigINumStr := "0"
	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_06(t *testing.T) {
	// minuend = 270.1
	minuendStr := "270.1"
	minuendPrecision := uint(1)

	// subtrahend = 0
	subtrahendStr := "0"
	subtrahendPrecision := uint(0)

	// result = 270.1
	expectedBigINumStr := "270.1"
	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_07(t *testing.T) {
	// minuend = 0
	minuendStr := "0"
	minuendPrecision := uint(0)

	// subtrahend = 270.1
	subtrahendStr := "270.1"
	subtrahendPrecision := uint(1)

	// result = -270.1
	expectedBigINumStr := "-270.1"
	expectedBigINumSign := -1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_08(t *testing.T) {
	// minuend = 0
	minuendStr := "0"
	minuendPrecision := uint(0)

	// subtrahend = -270.1
	subtrahendStr := "-270.1"
	subtrahendPrecision := uint(1)

	// result = 270.1
	expectedBigINumStr := "270.1"
	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_09(t *testing.T) {
	// minuend = 2.5
	minuendStr := "2.5"
	minuendPrecision := uint(0)

	// subtrahend = 2.5
	subtrahendStr := "2.5"
	subtrahendPrecision := uint(1)

	// result = 0.0
	expectedBigINumStr := "0.0"
	expectedBigINumSign := 1

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("Error: Expected Big1.Precision='%v'. Instead, Big1.Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("Error: Expected Big2.Precision='%v'. Instead, Big2.Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathSubtract{}.SubtractPair(bPair)

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

func TestBigIntMathSubtract_SubtractPair_10(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"
	minuendPrecision := uint(2)

	// subtrahend = 23.321
	subtrahendStr := "23.321"
	subtrahendPrecision := uint(3)

	// result = 99.999
	expectedNumStr := "99,999"

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(minuendBiNum, subtrahendBiNum)

	result := BigIntMathSubtract{}.SubtractPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr= '%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathSubtract_BigIntSubtract_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}

}

func TestBigIntMathSubtract_BigIntSubtract_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_05(t *testing.T) {

	minuendStr := "5"


	subtrahendStr := "5"


	expectedBigINumStr := "0"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_06(t *testing.T) {

	minuendStr := "-5"
	subtrahendStr := "5"
	expectedBigINumStr := "-10"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_07(t *testing.T) {

	minuendStr := "0"
	subtrahendStr := "0"
	expectedBigINumStr := "0"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_08(t *testing.T) {

	minuendStr := "50.0"
	subtrahendStr := "2.60134"
	expectedBigINumStr := "47.39866"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_09(t *testing.T) {

	minuendStr := "50.0"
	subtrahendStr := "50.0000"
	expectedBigINumStr := "0"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_BigIntSubtract_10(t *testing.T) {

	minuendStr := "0.0"
	subtrahendStr := "0.0000"
	expectedBigINumStr := "0"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result, resultPrecision := BigIntMathSubtract{}.BigIntSubtract(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != resultPrecision {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), resultPrecision)
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_05(t *testing.T) {

	minuendStr := "5"

	subtrahendStr := "5"

	expectedBigINumStr := "0"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_06(t *testing.T) {

	minuendStr := "-5"
	subtrahendStr := "5"
	expectedBigINumStr := "-10"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_07(t *testing.T) {

	minuendStr := "0"
	subtrahendStr := "0"
	expectedBigINumStr := "0"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}

func TestBigIntMathSubtract_FixedDecimalSubtract_08(t *testing.T) {

	minuendStr := "50.0"
	subtrahendStr := "2.60134"
	expectedBigINumStr := "47.39866"

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}


	bIMinuend :=
		BigIntFixedDecimal{}.New(
			minuendBiNum.GetIntegerValue(),
			minuendBiNum.GetPrecisionUint())

	biSubtrahend :=
		BigIntFixedDecimal{}.New(
			subtrahendBiNum.GetIntegerValue(),
			subtrahendBiNum.GetPrecisionUint())

	result := BigIntMathSubtract{}.FixedDecimalSubtract(
		bIMinuend,
		biSubtrahend)

	expectedBI := expectedBigINum.GetIntegerValue()

	if expectedBI.Cmp(result.GetInteger()) != 0 {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBI.Text(10), result.GetInteger().Text(10))
	}

	if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
			expectedBigINum.GetPrecisionUint(), result.GetPrecision())
	}
}
