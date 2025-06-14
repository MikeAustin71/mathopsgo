package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathDivide_BigIntDividedByTwoToPower_01(t *testing.T) {
	ePrefix := "TestBigIntMathDivide_BigIntDividedByTwoToPower_01"
	num := int64(33333)
	exponent := uint(8)
	expectedNum := int64(130)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient, err := new(BigIntMathDivide).BigIntDividedByTwoToPower(dividend, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intQuotient, err := new(BigIntMathDivide).\n"+
			"  BigIntDividedByTwoToPower(dividend, exponent)\n"+
			"dividend= '%v'\n"+
			"exponent= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividend.Text(10), exponent, err.Error())
		return
	}

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedValue.Cmp(intQuotient) != 0\n"+
			"Expected intQuotient = '%v'\n"+
			"Instead, intQuotient = '%v'\n\n",
			ePrefix, expectedValue.Text(10), intQuotient.Text(10))
	}

	return
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_02(t *testing.T) {
	ePrefix := "TestBigIntMathDivide_BigIntDividedByTwoToPower_02"
	num := int64(4)
	exponent := uint(9)
	expectedValue := big.NewInt(int64(0))

	dividend := big.NewInt(num)

	intQuotient, err := new(BigIntMathDivide).BigIntDividedByTwoToPower(dividend, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intQuotient, err := new(BigIntMathDivide).\n"+
			"  BigIntDividedByTwoToPower(dividend, exponent)\n"+
			"dividend= '%v'\n"+
			"exponent= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividend.Text(10), exponent, err.Error())
		return
	}

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedValue.Cmp(intQuotient) != 0\n"+
			"Expected intQuotient = '%v'\n"+
			"Instead, intQuotient = '%v'\n\n",
			ePrefix, expectedValue.Text(10), intQuotient.Text(10))
	}

	return
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntDividedByTwoToPower_03"
	num := int64(-8123456789012345)
	exponent := uint(12)
	expectedNum := int64(-1983265817630)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient, err := new(BigIntMathDivide).BigIntDividedByTwoToPower(dividend, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intQuotient, err := new(BigIntMathDivide).\n"+
			"  BigIntDividedByTwoToPower(dividend, exponent)\n"+
			"dividend= '%v'\n"+
			"exponent= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividend.Text(10), exponent, err.Error())
		return
	}

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedValue.Cmp(intQuotient) != 0\n"+
			"Expected intQuotient = '%v'\n"+
			"Instead, intQuotient = '%v'\n\n",
			ePrefix, expectedValue.Text(10), intQuotient.Text(10))
	}

	return
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_04(t *testing.T) {
	ePrefix := "TestBigIntMathDivide_BigIntDividedByTwoToPower_04"
	num := int64(8123456789012345)
	exponent := uint(12)
	expectedNum := int64(1983265817629)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient, err := new(BigIntMathDivide).BigIntDividedByTwoToPower(dividend, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intQuotient, err := new(BigIntMathDivide).\n"+
			"  BigIntDividedByTwoToPower(dividend, exponent)\n"+
			"dividend= '%v'\n"+
			"exponent= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividend.Text(10), exponent, err.Error())
		return
	}

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedValue.Cmp(intQuotient) != 0\n"+
			"Expected intQuotient = '%v'\n"+
			"Instead, intQuotient = '%v'\n\n",
			ePrefix, expectedValue.Text(10), intQuotient.Text(10))
	}

	return
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_05(t *testing.T) {
	ePrefix := "TestBigIntMathDivide_BigIntDividedByTwoToPower_05"
	num := int64(4)
	exponent := uint(1)
	expectedNum := int64(2)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient, err := new(BigIntMathDivide).BigIntDividedByTwoToPower(dividend, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intQuotient, err := new(BigIntMathDivide).\n"+
			"  BigIntDividedByTwoToPower(dividend, exponent)\n"+
			"dividend= '%v'\n"+
			"exponent= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividend.Text(10), exponent, err.Error())
		return
	}

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedValue.Cmp(intQuotient) != 0\n"+
			"Expected intQuotient = '%v'\n"+
			"Instead, intQuotient = '%v'\n\n",
			ePrefix, expectedValue.Text(10), intQuotient.Text(10))
	}

	return
}

func TestBigIntMathDivide_BigIntDividedByTwoToPower_06(t *testing.T) {
	ePrefix := "TestBigIntMathDivide_BigIntDividedByTwoToPower_06"
	num := int64(-4)
	exponent := uint(1)
	expectedNum := int64(-2)
	expectedValue := big.NewInt(expectedNum)

	dividend := big.NewInt(num)

	intQuotient, err := new(BigIntMathDivide).BigIntDividedByTwoToPower(dividend, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intQuotient, err := new(BigIntMathDivide).\n"+
			"  BigIntDividedByTwoToPower(dividend, exponent)\n"+
			"dividend= '%v'\n"+
			"exponent= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividend.Text(10), exponent, err.Error())
		return
	}

	if expectedValue.Cmp(intQuotient) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedValue.Cmp(intQuotient) != 0\n"+
			"Expected intQuotient = '%v'\n"+
			"Instead, intQuotient = '%v'\n\n",
			ePrefix, expectedValue.Text(10), intQuotient.Text(10))
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_01"
	dividendStr := "10.5"
	divisorStr := "2"
	expectedQuoStr := "5.25"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10    				/ 					2 				= 			5

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_02"
	dividendStr := "10"
	divisorStr := "2"
	expectedQuoStr := "5"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//   11.5  				/         	2.5				=  		4.6

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_03"
	dividendStr := "11.5"
	divisorStr := "2.5"
	expectedQuoStr := "4.6"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    2.5					/				 	12.555			=		0.199123855037834

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_04"
	dividendStr := "2.5"
	divisorStr := "12.555"
	expectedQuoStr := "0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_05(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_05"
	dividendStr := "-12.555"
	divisorStr := "2.5"
	expectedQuoStr := "-5.022"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_06(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  -12.555     	/    				2  			  = 	 -6.2775

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_06"
	dividendStr := "-12.555"
	divisorStr := "2"
	expectedQuoStr := "-6.2775"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_07(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_07"
	dividendStr := "-2.5"
	divisorStr := "12.555"
	expectedQuoStr := "-0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_08(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 12.555				/ 				- 2.5			  =		 -5.022

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_08"
	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedQuoStr := "-5.022"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_09(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//   12.555 			/ 				 -2 				=    -6.2775

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_09"
	dividendStr := "12.555"
	divisorStr := "-2"
	expectedQuoStr := "-6.2775"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_10(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    2.5 				/ 				-12.555		  = 	-0.199123855037834

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_10"
	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedQuoStr := "-0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_11(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_11"

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedQuoStr := "5.022"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_12(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  -12.555     	/    			 -2 				= 		6.2775

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_12"
	dividendStr := "-12.555"
	divisorStr := "-2"
	expectedQuoStr := "6.2775"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_13(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  - 2.5	 				/ 				-12.555		  = 		0.199123855037834
	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_13"
	dividendStr := "-2.5"
	divisorStr := "-12.555"
	expectedQuoStr := "0.199123855037834"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_14(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  -10						/					- 2					=						5

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_14"
	dividendStr := "-10"
	divisorStr := "-2"
	expectedQuoStr := "5"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_15(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  0							/					- 2					=						0

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_15"
	dividendStr := "0"
	divisorStr := "-2"
	expectedQuoStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_16(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  0							/					  2					=						0

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_16"
	dividendStr := "0"
	divisorStr := "2"
	expectedQuoStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)

		return
	}

	actualNumSeps, err := actualQuo.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err = actualQuo.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedNumSeps != actualNumSeps\n"+
			"Expected actualNumSeps = '%v'\n"+
			"Instead, actualNumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_17(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  11.5				  /				  2.5					=						4.6

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_17"
	dividendStr := "11.5"
	divisorStr := "2.5"
	expectedQuoStr := "4.6"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)

		return
	}

	actualNumSeps, err := actualQuo.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err = actualQuo.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedNumSeps != actualNumSeps\n"+
			"Expected actualNumSeps = '%v'\n"+
			"Instead, actualNumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_18(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  11.5				  /				  2.5					=						4.6

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_18"
	dividendStr := "11.5"
	divisorStr := "2.5"
	expectedQuoStr := "4.6"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_19(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    0 				 / 	  			 12.555		  = 	0

	dividendStr := "0"
	divisorStr := "12.555"
	expectedQuoStr := "0"
	maxPrecision := big.NewInt(15)
	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_19"

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
			"expectedQuoStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedQuoStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_20(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    15.8 			 / 	  			 0		  		= 	ERROR

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_20"
	dividendStr := "15.8"
	divisorStr := "0"
	maxPrecision := big.NewInt(15)

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	_, _, err =
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected Divide By zero Error.\n"+
			"However,  NO ERROR Was Returned!\n\n",
			ePrefix)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_21(t *testing.T) {
	// Dividend		 divided by		Divisor				=		Quotient
	// 0.000009218 		 /        35829.8234		= 2.572717118108932683156903307539e-10

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_21"
	dividendStr := "0.000009218"
	divisorStr := "35829.8234"

	expectedQuo, err :=
		new(BigIntNum).NewNumStr("0.0000000002572717118108932683156903307539")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(\n"+
			"  \"0.0000000002572717118108932683156903307539\")\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	maxPrecision, err := expectedQuo.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"maxPrecision, err := expectedQuo.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}

func TestBigIntMathDivide_BigIntFracQuotient_22(t *testing.T) {
	// Dividend		 divided by		Divisor							=		Quotient
	// 35829.8234 	 /        	 0.000009218     		= 3886941136.9060533738338034280755

	ePrefix := "TestBigIntMathDivide_BigIntFracQuotient_22"
	dividendStr := "35829.8234"
	divisorStr := "0.000009218"
	expectedResultStr := "3886941136.9060533738338034280755"

	expectedQuo, err :=
		new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	maxPrecision, err := expectedQuo.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"maxPrecision, err := expectedQuo.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividend, err := new(BigIntNum).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	dividenBigInt, err := dividend.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividenBigInt, err := dividend.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendPrecisionBigInt, err := dividend.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendPrecisionBigInt, err = dividend.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorBigInt, err := divisor.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorBigInt, err := divisor.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorPrecisionBigInt, err := divisor.GetPrecisionBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bIntQuotient, bIntQuotientPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividenBigInt,
			dividendPrecisionBigInt,
			divisorBigInt,
			divisorPrecisionBigInt,
			maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIntQuotient, bIntQuotientPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n"+
			"  dividenBigInt, dividendPrecisionBigInt, divisorBigInt,\n"+
			"  divisorPrecisionBigInt, maxPrecision)\n"+
			"dividenBigInt= '%v'\n"+
			"dividendPrecisionBigInt= '%v'\n"+
			"divisorBigInt= '%v'\n"+
			"divisorPrecisionBigInt= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividenBigInt.Text(10),
			dividendPrecisionBigInt.Text(10),
			divisorBigInt.Text(10),
			divisorPrecisionBigInt.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuo, err := new(BigIntNum).NewBigIntBigPrecision(bIntQuotient, bIntQuotientPrecision)\n"+
			"bIntQuotient= '%v'\n"+
			"bIntQuotientPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bIntQuotient.Text(10),
			bIntQuotientPrecision.Text(10),
			err.Error())
		return
	}

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualQuoNumStr, err := actualQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedQuoNumStr != actualQuoNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedQuoNumStr != actualQuoNumStr\n"+
			"Expected actualQuoNumStr = '%v'\n"+
			"Instead, actualQuoNumStr = '%v'\n\n",
			ePrefix, expectedQuoNumStr, actualQuoNumStr)
	}

	return
}
