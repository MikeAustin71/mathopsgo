package mathops

import (
	"testing"
)

func TestBigIntMathMultiply_MultiplyIntAry_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_01"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	expectedSignValue := 1

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierIntAry.IsValid("Validating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
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
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier='%v'\n"+
			"iaMultiplicand='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, multiplicandStr, err.Error())
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

	expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
			"multiplierIntAry, multiplicandIntAry)\n"+
			"multiplierIntAry= '%v'\n"+
			"multiplicandIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierIntAryNumStr,
			multiplicandIntAryNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
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

	if expectedIntAryNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedIntAryNumStr, resultNumStr)

		return
	}

	expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
			"expectedIntAry='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
			"Expected resultBigIntNum = '%v'\n"+
			"  Actual resultBigIntNum = '%v'\n\n",
			ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedSignValue != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != result.sign\n"+
			"Expected result.sign = '%v'\n"+
			"  Actual result.sign = '%v'\n\n",
			ePrefix, expectedSignValue, result.sign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAry_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_02"

	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	expectedSignValue := 1

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierIntAry.IsValid("Validating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
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
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier='%v'\n"+
			"iaMultiplicand='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, multiplicandStr, err.Error())
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

	expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
			"multiplierIntAry, multiplicandIntAry)\n"+
			"multiplierIntAry= '%v'\n"+
			"multiplicandIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierIntAryNumStr,
			multiplicandIntAryNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
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

	if expectedIntAryNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedIntAryNumStr, resultNumStr)

		return
	}

	expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
			"expectedIntAry='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
			"Expected resultBigIntNum = '%v'\n"+
			"  Actual resultBigIntNum = '%v'\n\n",
			ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedSignValue != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != result.sign\n"+
			"Expected result.sign = '%v'\n"+
			"  Actual result.sign = '%v'\n\n",
			ePrefix, expectedSignValue, result.sign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAry_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_03"

	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	expectedSignValue := -1

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierIntAry.IsValid("Validating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
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
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier='%v'\n"+
			"iaMultiplicand='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, multiplicandStr, err.Error())
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

	expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
			"multiplierIntAry, multiplicandIntAry)\n"+
			"multiplierIntAry= '%v'\n"+
			"multiplicandIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierIntAryNumStr,
			multiplicandIntAryNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
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

	if expectedIntAryNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedIntAryNumStr, resultNumStr)

		return
	}

	expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
			"expectedIntAry='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
			"Expected resultBigIntNum = '%v'\n"+
			"  Actual resultBigIntNum = '%v'\n\n",
			ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedSignValue != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != result.sign\n"+
			"Expected result.sign = '%v'\n"+
			"  Actual result.sign = '%v'\n\n",
			ePrefix, expectedSignValue, result.sign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAry_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_04"

	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	expectedSignValue := 1

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierIntAry.IsValid("Validating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
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
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier='%v'\n"+
			"iaMultiplicand='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, multiplicandStr, err.Error())
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

	expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
			"multiplierIntAry, multiplicandIntAry)\n"+
			"multiplierIntAry= '%v'\n"+
			"multiplicandIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierIntAryNumStr,
			multiplicandIntAryNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
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

	if expectedIntAryNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedIntAryNumStr, resultNumStr)

		return
	}

	expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
			"expectedIntAry='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
			"Expected resultBigIntNum = '%v'\n"+
			"  Actual resultBigIntNum = '%v'\n\n",
			ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedSignValue != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != result.sign\n"+
			"Expected result.sign = '%v'\n"+
			"  Actual result.sign = '%v'\n\n",
			ePrefix, expectedSignValue, result.sign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	iaResultActualBigInt, err := iaResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultActualBigInt, err := iaResult.GetBigInt()\n"+
			"iaResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, iaResultNumStr, err.Error())
		return
	}

	if resultActualBigInt.Cmp(iaResultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because resultActualBigInt.Cmp(iaResultActualBigInt) != 0\n"+
			"Expected iaResultActualBigInt = '%v'\n"+
			"  Actual iaResultActualBigInt = '%v'\n\n",
			ePrefix, resultActualBigInt.Text(10), iaResultActualBigInt.Text(10))

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAry_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_05"

	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0
	expectedNumStr := "0"

	expectedSignValue := 1

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierIntAry.IsValid("Validating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAry, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
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
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier='%v'\n"+
			"iaMultiplicand='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, multiplicandStr, err.Error())
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

	expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
			"multiplierIntAry, multiplicandIntAry)\n"+
			"multiplierIntAry= '%v'\n"+
			"multiplicandIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierIntAryNumStr,
			multiplicandIntAryNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
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

	if expectedIntAryNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedIntAryNumStr, resultNumStr)

		return
	}

	expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
			"expectedIntAry='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
			"Expected resultBigIntNum = '%v'\n"+
			"  Actual resultBigIntNum = '%v'\n\n",
			ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedSignValue != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != result.sign\n"+
			"Expected result.sign = '%v'\n"+
			"  Actual result.sign = '%v'\n\n",
			ePrefix, expectedSignValue, result.sign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	iaResultActualBigInt, err := iaResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultActualBigInt, err := iaResult.GetBigInt()\n"+
			"iaResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, iaResultNumStr, err.Error())
		return
	}

	if resultActualBigInt.Cmp(iaResultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because resultActualBigInt.Cmp(iaResultActualBigInt) != 0\n"+
			"Expected iaResultActualBigInt = '%v'\n"+
			"  Actual iaResultActualBigInt = '%v'\n\n",
			ePrefix, resultActualBigInt.Text(10), iaResultActualBigInt.Text(10))

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAry_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_06"

	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145,3632"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierIntAry, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplierIntAry.IsValid("Validating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAry, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAry, err := new(IntAry).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).\n"+
			"  NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).\n"+
			"  NewNumStr(multiplicandStr)\n"+
			"multiplicandStr='%v'\n"+
			"Error='%v'\n\n",
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
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand\n"+
			"iaMultiplier= '%v'\n"+
			"iaMultiplicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, multiplicandStr, err.Error())
		return
	}

	err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

	expectedIntAry, err := new(IntAry).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
			" multiplierIntAry, multiplicandIntAry, expectedNumSeps)\n"+
			"multiplierIntAry= '%v'\n"+
			"multiplicandIntAry= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, multiplierIntAryNumStr,
			multiplicandIntAryNumStr, err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
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

	if expectedIntAryNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedIntAryNumStr, resultNumStr)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because !expectedNumSeps.Equal(actualNumSeps)\n"+
			"Expected actualNumSeps = '%v'\n"+
			"  Actual actualNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), actualNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyIntAryArray_01"

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierIntAry.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	multiplicandIntArray := make([]IntAry, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	var ia IntAry

	for i := 0; i < lenArray; i++ {

		multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"multiplicandIntArray[%d], err = new(IntAry).\n"+
				"  NewNumStr(multiplicandStrs[%d])\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], err.Error())
			return
		}

		ia, err = multiplicandIntArray[i].CopyOut()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = multiplicandIntArray[%d].CopyOut()\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
			return
		}

	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)\n"+
			"multiplierIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierIntAryNumStr,
			err.Error())

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

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigIntNumberStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
			"Expected resultActualBigInt = '%v'\n"+
			"  Actual resultActualBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

		return
	}

	resultSign, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSign, err := result.GetSign()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSign\n"+
			"Expected resultSign = '%v'\n"+
			"  Actual resultSign = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_02(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia := decimalArray[i]

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(IntAry).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryArray("+
			"multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected IntAry='%s'. Instead, IntAry= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected IntAry='%s'. Instead, IntAry= '%s'. ",
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

func TestBigIntMathMultiply_MultiplyIntAryArray_03(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia := decimalArray[i]

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(IntAry).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryArray("+
			"multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected IntAry='%s'. Instead, IntAry= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected IntAry='%s'. Instead, IntAry= '%s'. ",
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

func TestBigIntMathMultiply_MultiplyIntAryArray_04(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia := decimalArray[i]

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(IntAry).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryArray("+
			"multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected IntAry='%s'. Instead, IntAry= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected IntAry='%s'. Instead, IntAry= '%s'. ",
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

func TestBigIntMathMultiply_MultiplyIntAryArray_05(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierIntAry.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia := decimalArray[i]

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryArray("+
			"multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

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

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_01(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryOutputToArray"+
			"(multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_02(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryOutputToArray"+
			"(multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_03(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryOutputToArray"+
			"(multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_04(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryOutputToArray"+
			"(multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_05(t *testing.T) {

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

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryOutputToArray"+
			"(multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_06(t *testing.T) {

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
		"800,8",
		"-208",
		"31,392",
		"64",
		"42376,984",
		"-39,168",
	}

	multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierIntAry.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyIntAryOutputToArray"+
			"(multiplierIntAry, decimalArray) multiplierIntAry='%v'  Error='%v'. ",
			multiplierIntAry.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v' Index='%v'",
				expectedNumSeps.String(), actualNumSeps.String(), j)
		}
	}
}
