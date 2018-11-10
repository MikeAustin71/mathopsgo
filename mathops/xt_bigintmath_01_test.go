package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMath_ArithmeticGeometricMean_01(t *testing.T) {
	aNum := big.NewInt(24)
	aNumPrecision := big.NewInt(0)
	gNum := big.NewInt(6)
	gNumPrecision := big.NewInt(0)
	maxInternalPrecision := big.NewInt(150)
	targetPrecision := big.NewInt(46)
	expectedValue := "13.4581714817256154207668131569743992430538388544"

	agMean, agMeanPrecision, gValue, gValuePrecision, _, err :=
		BigIntMath{}.ArithmeticGeometricMean(
			aNum,
			aNumPrecision,
			gNum,
			gNumPrecision,
			maxInternalPrecision,
			targetPrecision)


	if err!=nil {
		t.Errorf("Error: %v", err.Error())
		return
	}

	binAGMean, err :=BigIntNum{}.NewBigIntPrecision(agMean, agMeanPrecision)

	if err!=nil {
		t.Errorf("%v", err.Error())
		return
	}

	binGValue, err := BigIntNum{}.NewBigIntPrecision(gValue, gValuePrecision)

	if err!=nil {
		t.Errorf( "%v", err.Error())
		return
	}

	if expectedValue != binAGMean.GetNumStr() {
		t.Errorf("Error: Expected agMean='%v'.  Instead, agMean='%v'. ",
			expectedValue, binAGMean.GetNumStr())
	}

	if expectedValue != binGValue.GetNumStr() {
		t.Errorf("Error: Expected gMean='%v'.  Instead, gMean='%v'. ",
			expectedValue, binGValue.GetNumStr())
	}

}

func TestBigIntMath_GetMagnitude_01(t *testing.T) {

	target := big.NewInt(98327123)
	expectedMagnitude := big.NewInt(7)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_02(t *testing.T) {

	target := big.NewInt(0)
	expectedMagnitude := big.NewInt(0)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_03(t *testing.T) {

	target := big.NewInt(82)
	expectedMagnitude := big.NewInt(1)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_04(t *testing.T) {

	target := big.NewInt(-5)
	expectedMagnitude := big.NewInt(0)
	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}
}

func TestBigIntMath_GetMagnitude_05(t *testing.T) {

	target := big.NewInt(2)
	expectedMagnitude := big.NewInt(0)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_06(t *testing.T) {

	target := big.NewInt(85652647928)
	expectedMagnitude := big.NewInt(10)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_07(t *testing.T) {


	numStr := "8565264792812345678901234567890"
	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())
	expectedMagnitude := big.NewInt(30)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_08(t *testing.T) {

	numStr := "-8565264792812345678901234567890"
	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())
	expectedMagnitude := big.NewInt(30)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_09(t *testing.T) {

	numStr := "131072"
	expectedMagnitude := big.NewInt(5)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_10(t *testing.T) {

	numStr := "131071"
	expectedMagnitude := big.NewInt(5)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_11(t *testing.T) {

	numStr := "999999999999"
	expectedMagnitude := big.NewInt(11)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_12(t *testing.T) {

	numStr := "1999999999999"
	expectedMagnitude := big.NewInt(12)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_13(t *testing.T) {

	numStr := "9999999999999"
	expectedMagnitude := big.NewInt(12)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_14(t *testing.T) {

	numStr := "7"
	expectedMagnitude := big.NewInt(0)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_15(t *testing.T) {

	numStr := "8"
	expectedMagnitude := big.NewInt(0)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_16(t *testing.T) {

	numStr := "9"
	expectedMagnitude := big.NewInt(0)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_17(t *testing.T) {

	numStr := "10"
	expectedMagnitude := big.NewInt(1)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_BigIntPrecisionCmp_01(t *testing.T) {
	num1 := big.NewInt(5)
	num1Precision := big.NewInt(0)
	num2 := big.NewInt(2)
	num2Precision := big.NewInt(0)
	expectedResult := 1

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntMath_BigIntPrecisionCmp_02(t *testing.T) {
	num1 := big.NewInt(52)
	num1Precision := big.NewInt(1)
	num2 := big.NewInt(5105)
	num2Precision := big.NewInt(3)
	expectedResult := 1

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntMath_BigIntPrecisionCmp_03(t *testing.T) {
	num1 := big.NewInt(-52)
	num1Precision := big.NewInt(1)
	num2 := big.NewInt(5105)
	num2Precision := big.NewInt(3)
	expectedResult := -1

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntMath_BigIntPrecisionCmp_04(t *testing.T) {
	num1 := big.NewInt(52346789)
	num1Precision := big.NewInt(1)
	num2 := big.NewInt(8234001)
	num2Precision := big.NewInt(3)
	expectedResult := 1

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntMath_BigIntPrecisionCmp_05(t *testing.T) {
	num1 := big.NewInt(52346789)
	num1Precision := big.NewInt(1)
	num2 := big.NewInt(52346789)
	num2Precision := big.NewInt(1)
	expectedResult := 0

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntMath_BigIntPrecisionCmp_06(t *testing.T) {
	num1 := big.NewInt(-51)
	num1Precision := big.NewInt(1)
	num2 := big.NewInt(-52)
	num2Precision := big.NewInt(1)
	expectedResult := 1

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntMath_BigIntPrecisionCmp_07(t *testing.T) {
	num1 := big.NewInt(-51)
	num1Precision := big.NewInt(1)
	num2 := big.NewInt(-51)
	num2Precision := big.NewInt(1)
	expectedResult := 0

	cmpResult :=
		BigIntMath{}.BigIntPrecisionCmp(
			num1,
			num1Precision,
			num2,
			num2Precision)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}


func TestBigIntMath_RoundToMaxPrecision_01(t *testing.T) {
	num1 := big.NewInt(762939453125)
	num1Precision := big.NewInt(17)
	maxPrecision := big.NewInt(16)
	expectedResult := "0.0000076293945313"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_RoundToMaxPrecision_02(t *testing.T) {
	binNum1, err := BigIntNum{}.NewNumStr("276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(29)
	expectedResult := "0.00276213586400995126719079829"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_RoundToMaxPrecision_03(t *testing.T) {

	binNum1, err := BigIntNum{}.NewNumStr("276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(30)
	expectedResult := "0.002762135864009951267190798289"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_RoundToMaxPrecision_04(t *testing.T) {

	binNum1, err := BigIntNum{}.NewNumStr("-276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "-0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(30)
	expectedResult := "-0.002762135864009951267190798289"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_RoundToMaxPrecision_05(t *testing.T) {
	binNum1, err := BigIntNum{}.NewNumStr("-276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "-0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(29)
	expectedResult := "-0.00276213586400995126719079829"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_RoundToMaxPrecision_06(t *testing.T) {

	num1 := big.NewInt(1230000)

	num1Precision := big.NewInt(5)
	// "12.30000" -> "
	maxPrecision := big.NewInt(3)
	expectedResult := "12.3"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, true)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_RoundToMaxPrecision_07(t *testing.T) {

	num1 := big.NewInt(-1230000)

	num1Precision := big.NewInt(5)
	// "-12.30000" -> "
	maxPrecision := big.NewInt(3)
	expectedResult := "-12.3"

	result, resultPrecision, err :=
		BigIntMath{}.RoundToMaxPrecision(num1, num1Precision, maxPrecision, true)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}
}


func TestBigIntMath_TruncateToMaxPrecision_01(t *testing.T) {
	num1 := big.NewInt(762939453125)
	num1Precision := big.NewInt(17)
	maxPrecision := big.NewInt(16)
	expectedResult := "0.0000076293945312"

	result, resultPrecision, err :=
		BigIntMath{}.TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from Truncate to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_TruncateToMaxPrecision_02(t *testing.T) {
	binNum1, err := BigIntNum{}.NewNumStr("276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(29)
	expectedResult := "0.00276213586400995126719079828"

	result, resultPrecision, err :=
		BigIntMath{}.TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from Truncate to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_TruncateToMaxPrecision_03(t *testing.T) {

	binNum1, err := BigIntNum{}.NewNumStr("276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(30)
	expectedResult := "0.002762135864009951267190798289"

	result, resultPrecision, err :=
		BigIntMath{}.TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from Truncate to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_TruncateToMaxPrecision_04(t *testing.T) {

	binNum1, err := BigIntNum{}.NewNumStr("-276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "-0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(30)
	expectedResult := "-0.002762135864009951267190798289"

	result, resultPrecision, err :=
		BigIntMath{}.TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}

func TestBigIntMath_TruncateToMaxPrecision_05(t *testing.T) {
	binNum1, err := BigIntNum{}.NewNumStr("-276213586400995126719079828947")

	if err != nil {
		t.Errorf("Error returned from binNum1: %v", err.Error())
	}

	num1 := binNum1.GetIntegerValue()

	num1Precision := big.NewInt(32)
	// "-0.00276213586400995126719079828947"
	maxPrecision := big.NewInt(29)
	expectedResult := "-0.00276213586400995126719079828"

	result, resultPrecision, err :=
		BigIntMath{}.TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewBigIntPrecision: %v",
			err.Error())
	}

	if expectedResult != binResult.GetNumStr() {
		t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
			expectedResult, binResult.GetNumStr())
	}

}