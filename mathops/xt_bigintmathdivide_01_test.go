package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathDivide_BigIntDividedByTwoToPower_01(t *testing.T) {

	num := int64(33333)
	exponent := uint(8)
	expectedNum := int64(130)
  expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient := BigIntMathDivide{}.BigIntDividedByTwoToPower(dividend, exponent)

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.Text(10), intQuotient.Text(10))
	}
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_02(t *testing.T) {

	num := int64(4)
	exponent := uint(9)
  expectedValue := big.NewInt(int64(0))

	dividend := big.NewInt(num)

	intQuotient := BigIntMathDivide{}.BigIntDividedByTwoToPower(dividend, exponent)

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.Text(10), intQuotient.Text(10))
	}

}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_03(t *testing.T) {

	num := int64(-8123456789012345)
	exponent := uint(12)
	expectedNum := int64(-1983265817630)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient := BigIntMathDivide{}.BigIntDividedByTwoToPower(dividend, exponent)

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.Text(10), intQuotient.Text(10))
	}
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_04(t *testing.T) {

	num := int64(8123456789012345)
	exponent := uint(12)
	expectedNum := int64(1983265817629)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient := BigIntMathDivide{}.BigIntDividedByTwoToPower(dividend, exponent)

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.Text(10), intQuotient.Text(10))
	}
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_05(t *testing.T) {

	num := int64(4)
	exponent := uint(1)
	expectedNum := int64(2)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient := BigIntMathDivide{}.BigIntDividedByTwoToPower(dividend, exponent)

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.Text(10), intQuotient.Text(10))
	}
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_06(t *testing.T) {

	num := int64(-4)
	exponent := uint(1)
	expectedNum := int64(-2)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient := BigIntMathDivide{}.BigIntDividedByTwoToPower(dividend, exponent)

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.Text(10), intQuotient.Text(10))
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr := "10.5"
	divisorStr := "2"
	expectedQuoStr := "5.25"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
									BigIntMathDivide{}.BigIntFracQuotient(
										dividend.GetIntegerValue(),
										dividend.GetPrecisionBigInt(),
										divisor.GetIntegerValue(),
										divisor.GetPrecisionBigInt(),
										maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10    				/ 					2 				= 			5

	dividendStr := "10"
	divisorStr := "2"
	expectedQuoStr := "5"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}


	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//   11.5  				/         	2.5				=  		4.6

	dividendStr := "11.5"
	divisorStr := "2.5"
	expectedQuoStr := "4.6"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}


	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}


	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    2.5					/				 	12.555			=		0.199123855037834

	dividendStr := "2.5"
	divisorStr := "12.555"
	expectedQuoStr := "0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_05(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr := "-12.555"
	divisorStr := "2.5"
	expectedQuoStr := "-5.022"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_06(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  -12.555     	/    				2  			  = 	 -6.2775

	dividendStr := "-12.555"
	divisorStr := "2"
	expectedQuoStr := "-6.2775"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_07(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr := "-2.5"
	divisorStr := "12.555"
	expectedQuoStr := "-0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_08(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 12.555				/ 				- 2.5			  =		 -5.022

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedQuoStr := "-5.022"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_09(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//   12.555 			/ 				 -2 				=    -6.2775

	dividendStr := "12.555"
	divisorStr := "-2"
	expectedQuoStr := "-6.2775"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_10(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    2.5 				/ 				-12.555		  = 	-0.199123855037834

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedQuoStr := "-0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_11(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedQuoStr := "5.022"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_12(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  -12.555     	/    			 -2 				= 		6.2775

	dividendStr := "-12.555"
	divisorStr := "-2"
	expectedQuoStr := "6.2775"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_13(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  - 2.5	 				/ 				-12.555		  = 		0.199123855037834

	dividendStr := "-2.5"
	divisorStr := "-12.555"
	expectedQuoStr := "0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_14(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  -10						/					- 2					=						5

	dividendStr := "-10"
	divisorStr := "-2"
	expectedQuoStr := "5"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_15(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  0							/					- 2					=						0

	dividendStr := "0"
	divisorStr := "-2"
	expectedQuoStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_16(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  0							/					  2					=						0

	dividendStr := "0"
	divisorStr := "2"
	expectedQuoStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_17(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  11.5				  /				  2.5					=						4.6

	dividendStr := "11.5"
	divisorStr := "2.5"
	expectedQuoStr := "4.6"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntFracQuotient_18(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  11.5				  /				  2.5					=						4.6

	dividendStr := "11.5"
	divisorStr := "2.5"
	expectedQuoStr := "4.6"
	maxPrecision := big.NewInt(15)


	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntFracQuotient_19(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    0 				 / 	  			 12.555		  = 	0

	dividendStr := "0"
	divisorStr := "12.555"
	expectedQuoStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). "+
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(...) " +
			"Error='%v' ", err.Error())
	}

	actualQuo, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedQuo.GetNumStr() != actualQuo.GetNumStr() {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), actualQuo.GetNumStr())
	}
}


func TestBigIntMathDivide_BigIntFracQuotient_20(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    15.8 			 / 	  			 0		  		= 	ERROR

	dividendStr := "15.8"
	divisorStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	_, _, err =
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err == nil {
		t.Error("Error - Expected Divide By zero Error. NO ERROR Returned! ")
	}

}

func TestBigIntMathDivide_BigIntFracQuotient_21(t *testing.T) {
	// Dividend		 divided by		Divisor							=		Quotient
	// 0.000009218 		 /        35829.8234	     		= 2.572717118108932683156903307539e-10

	dividendStr := "0.000009218"
	divisorStr := "35829.8234"
	expectedResult, err :=
		BigIntNum{}.NewNumStr("0.0000000002572717118108932683156903307539")

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	maxPrecision := expectedResult.GetPrecisionBigInt()

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(). "+
			"Error='%v' ", err.Error())
	}

	actualResult, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult.GetNumStr() != actualResult.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedResult.GetNumStr(), actualResult.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntFracQuotient_22(t *testing.T) {
	// Dividend		 divided by		Divisor							=		Quotient
	// 35829.8234 	 /        	 0.000009218     		= 3886941136.9060533738338034280755

	dividendStr := "35829.8234"
	divisorStr := "0.000009218"
	expectedResultStr := "3886941136.9060533738338034280755"
	expectedResult, err :=
		BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedResultStr). "+
			"expectedResultStr='%v' Error='%v' ",
			expectedResultStr, err.Error())
	}

	maxPrecision := expectedResult.GetPrecisionBigInt()

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		BigIntMathDivide{}.BigIntFracQuotient(
			dividend.GetIntegerValue(),
			dividend.GetPrecisionBigInt(),
			divisor.GetIntegerValue(),
			divisor.GetPrecisionBigInt(),
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntFracQuotient(). "+
			"Error='%v' ", err.Error())
	}

	actualResult, err := BigIntNum{}.NewBigIntPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult.GetNumStr() != actualResult.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedResult.GetNumStr(), actualResult.GetNumStr())
	}
}

