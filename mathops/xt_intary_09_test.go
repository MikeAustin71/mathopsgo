package mathops

import (
	"math/big"
	"testing"
)

func TestIntAry_NewNumStrWithNumSeps2_01(t *testing.T) {

	ePrefix := "TestIntAry_NewNumStrWithNumSeps2_01"

	// -------------------------------------
	//  input Values Setup

	inputNumStr := "401.2345"

	inputIntAryNumStr := inputNumStr

	inputIntAryBigIntNumStr := "4012345"

	inputIntAryBigInt, isOk := big.NewInt(0).SetString(inputIntAryBigIntNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inputIntAryBigInt, isOk :=\n"+
			"  big.NewInt(0).SetString(inputIntAryBigIntNumStr, 10)\n"+
			"expectedBigIntStr= '%v'\n"+
			"base= '10'\n"+
			"Error= \n\n",
			ePrefix,
			inputIntAryBigIntNumStr)

		return
	}

	inputArrayUint8 := []uint8{4, 0, 1, 2, 3, 4, 5}

	inputArrayUint8Len := len(inputArrayUint8)

	inputArrayMagnitudeInt := 2 // 10^2 == 100

	inputArrayPrecisionInt := 4

	inputArrayPrecisionUint := uint(inputArrayPrecisionInt)

	inputArraySignValue := 1

	inputArrayNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	// -------------------------------
	// Expected Values Setup

	expectedNumStr := "401.2345"

	expectedBigIntStr := "4012345"

	expectedBigInt, isOk := big.NewInt(0).SetString(expectedBigIntStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, isOk := big.NewInt(0).SetString(expectedBigIntStr, 10)\n"+
			"expectedBigIntStr= '%v'\n"+
			"base= '10'\n"+
			"Error= \n\n",
			ePrefix,
			expectedBigIntStr)

		return
	}

	expectedArrayUint8 := []uint8{4, 0, 1, 2, 3, 4, 5}

	expectedArrayUint8Len := len(expectedArrayUint8)

	expectedMagnitudeInt := 2 // 10^2 == 100

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedIntAry.IsValid("Validating expectedIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedIntAry.IsValid(\"Validating expectedIntAry\")\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedIntAryNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedIntAryNumStr\n"+
			"Expected expectedIntAryNumStr = '%v'\n"+
			"  Actual expectedIntAryNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedIntAryNumStr)

		return
	}

	expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryBigInt, err :=\n"+
			"  expectedIntAry.GetBigInt()\n"+
			"expectedIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedBigInt.Cmp(expectedIntAryBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected IntAry Big Int is INVALID!\n"+
			"Because expectedBigInt.Cmp(expectedIntAryBigInt) != 0\n"+
			"Expected expectedIntAryBigInt = '%v'\n"+
			"  Actual expectedIntAryBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), expectedIntAryBigInt.Text(10))

		return
	}

	expectedIntAryMagnitude, err := expectedIntAry.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryMagnitude, err :=\n"+
			"  expectedIntAry.GetMagnitude()\n"+
			"expectedIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedMagnitudeInt != expectedIntAryMagnitude {
		t.Errorf("%v\n"+
			"Error: expectedIntAry Magnitude is INVALID!\n"+
			"Because expectedMagnitudeInt != expectedIntAryMagnitude\n"+
			"Expected expectedIntAryMagnitude = '%v'\n"+
			"  Actual expectedIntAryMagnitude = '%v'\n\n",
			ePrefix, expectedMagnitudeInt, expectedIntAryMagnitude)

		return
	}

	expectedIntAryPrecisionInt := expectedIntAry.GetPrecision()

	if expectedPrecisionInt != expectedIntAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedPrecisionInt != expectedIntAryPrecisionInt\n"+
			"Expected expectedIntAryPrecisionInt = '%v'\n"+
			"  Actual expectedIntAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, expectedIntAryPrecisionInt)

		return
	}

	expectedIntAryPrecisionUint, err := expectedIntAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryPrecisionUint, err :=\n"+
			"  expectedIntAry.GetPrecisionUint()\n"+
			"expectedIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedIntAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Int Ary Precision Uint Invalid!\n"+
			"Because expectedPrecisionUint != expectedIntAryPrecisionUint\n"+
			"Expected expectedIntAryPrecisionUint = '%v'\n"+
			"  Actual expectedIntAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedIntAryPrecisionUint)

		return
	}

	expectedIntArySignValue, err := expectedIntAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntArySignValue, err := expectedIntAry.GetSign()\n"+
			"expectedIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedSignValue != expectedIntArySignValue {
		t.Errorf("%v\n"+
			"Error: expectedIntArySignValue is INVALID!\n"+
			"Because expectedSignValue != expectedIntArySignValue\n"+
			"Expected expectedIntArySignValue = '%v'\n"+
			"  Actual expectedIntArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, expectedIntArySignValue)

		return
	}

	expectedIntAryNumSeps, err := expectedIntAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryNumSeps, err :=\n"+
			"  expectedIntAry.GetNumericSeparatorsDto()\n"+
			"expectedIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedIntAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != expectedIntAryNumSeps\n"+
			"Expected expectedIntAryNumSeps = '%v'\n"+
			"  Actual expectedIntAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedIntAryNumSeps.String())

		return
	}

	expectedIntAryUint8Elements, expectedIntAryUint8Len, err :=
		expectedIntAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedIntAryUint8Elements, expectedIntAryUint8Len, err :=\n"+
			"  expectedIntAry.GetIntAryElements()\n"+
			"expectedIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedIntAryNumStr, err.Error())
		return
	}

	if expectedArrayUint8Len != expectedIntAryUint8Len {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
			"Because expectedArrayUint8Len != expectedIntAryUint8Len\n"+
			"Expected expectedIntAryUint8Len = '%v'\n"+
			"  Actual expectedIntAryUint8Len = '%v'\n\n",
			ePrefix, expectedArrayUint8Len, expectedIntAryUint8Len)

		return
	}

	for i := 0; i < expectedIntAryUint8Len; i++ {

		if expectedArrayUint8[i] != expectedIntAryUint8Elements[i] {
			t.Errorf("%v\n"+
				"Error: Expected and Actual Array Elements Don't Match!\n"+
				"Because expectedArrayUint8[%v] != expectedIntAryUint8Elements[%v]\n"+
				"Expected expectedIntAryUint8Elements[%v] = '%v'\n"+
				"  Actual expectedIntAryUint8Elements[%v] = '%v'\n"+
				"Expected Number String = '%v'\n"+
				"ExpectedIntAry Number String   = '%v'\n",
				ePrefix, i, i, i, expectedArrayUint8[i], i, expectedIntAryUint8Elements[i],
				expectedNumStr, expectedIntAryNumStr)

			return
		}

	}

	actualIntAry, err := new(IntAry).NewNumStr(inputNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAry, err := new(IntAry).NewNumStr(inputNumStr)\n"+
			"inputNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr, err.Error())
		return
	}

	err = actualIntAry.IsValid("Validating actualIntAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualIntAry.IsValid(\"Validating actualIntAry\")\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	actualIntAryNumStr, err := actualIntAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAryNumStr, err := actualIntAry.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputIntAryNumStr != actualIntAryNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because inputIntAryNumStr != actualIntAryNumStr\n"+
			"Expected actualIntAryNumStr = '%v'\n"+
			"  Actual actualIntAryNumStr = '%v'\n\n",
			ePrefix, inputIntAryNumStr, actualIntAryNumStr)

		return
	}

	actualIntAryBigInt, err := actualIntAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAryBigInt, err :=\n"+
			"  actualIntAry.GetBigInt()\n"+
			"actualIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualIntAryNumStr, err.Error())
		return
	}

	if inputIntAryBigInt.Cmp(actualIntAryBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Actual IntAry Big Int is INVALID!\n"+
			"Because inputIntAryBigInt.Cmp(actualIntAryBigInt) != 0\n"+
			"Expected actualIntAryBigInt = '%v'\n"+
			"  Actual actualIntAryBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), actualIntAryBigInt.Text(10))

		return
	}

	actualIntAryMagnitude, err := actualIntAry.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAryMagnitude, err :=\n"+
			"  actualIntAry.GetMagnitude()\n"+
			"actualIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualIntAryNumStr, err.Error())
		return
	}

	if inputArrayMagnitudeInt != actualIntAryMagnitude {
		t.Errorf("%v\n"+
			"Error: actualIntAry Magnitude is INVALID!\n"+
			"Because inputArrayMagnitudeInt != actualIntAryMagnitude\n"+
			"Expected actualIntAryMagnitude = '%v'\n"+
			"  Actual actualIntAryMagnitude = '%v'\n\n",
			ePrefix, inputArrayMagnitudeInt, actualIntAryMagnitude)

		return
	}

	actualIntAryPrecisionInt := actualIntAry.GetPrecision()

	if inputArrayPrecisionInt != actualIntAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because inputArrayPrecisionInt != actualIntAryPrecisionInt\n"+
			"Expected actualIntAryPrecisionInt = '%v'\n"+
			"  Actual actualIntAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, actualIntAryPrecisionInt)

		return
	}

	actualIntAryPrecisionUint, err := actualIntAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAryPrecisionUint, err :=\n"+
			"  actualIntAry.GetPrecisionUint()\n"+
			"actualIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualIntAryNumStr, err.Error())
		return
	}

	if inputArrayPrecisionUint != actualIntAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Actual Int Ary Precision Uint Invalid!\n"+
			"Because inputArrayPrecisionUint != actualIntAryPrecisionUint\n"+
			"Expected actualIntAryPrecisionUint = '%v'\n"+
			"  Actual actualIntAryPrecisionUint = '%v'\n\n",
			ePrefix, inputArrayPrecisionUint, actualIntAryPrecisionUint)

		return
	}

	actualIntArySignValue, err := actualIntAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntArySignValue, err := actualIntAry.GetSign()\n"+
			"actualIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualIntAryNumStr, err.Error())
		return
	}

	if inputArraySignValue != actualIntArySignValue {
		t.Errorf("%v\n"+
			"Error: actualIntArySignValue is INVALID!\n"+
			"Because inputArraySignValue != actualIntArySignValue\n"+
			"Expected actualIntArySignValue = '%v'\n"+
			"  Actual actualIntArySignValue = '%v'\n\n",
			ePrefix, inputArraySignValue, actualIntArySignValue)

		return
	}

	actualIntAryNumSeps, err := actualIntAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAryNumSeps, err :=\n"+
			"  actualIntAry.GetNumericSeparatorsDto()\n"+
			"actualIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualIntAryNumStr, err.Error())
		return
	}

	if !inputArrayNumSeps.Equal(actualIntAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because inputArrayNumSeps != actualIntAryNumSeps\n"+
			"Expected actualIntAryNumSeps = '%v'\n"+
			"  Actual actualIntAryNumSeps = '%v'\n\n",
			ePrefix, inputArrayNumSeps.String(), actualIntAryNumSeps.String())

		return
	}

	actualIntAryUint8Elements, actualIntAryUint8Len, err :=
		actualIntAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIntAryUint8Elements, actualIntAryUint8Len, err :=\n"+
			"  actualIntAry.GetIntAryElements()\n"+
			"actualIntAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualIntAryNumStr, err.Error())
		return
	}

	if inputArrayUint8Len != actualIntAryUint8Len {
		t.Errorf("%v\n"+
			"Error: Input and Actual Array Lengths ARE NOT EQUAL!\n"+
			"Because inputArrayUint8Len != actualIntAryUint8Len\n"+
			"Expected actualIntAryUint8Len = '%v'\n"+
			"  Actual actualIntAryUint8Len = '%v'\n\n",
			ePrefix, inputArrayUint8Len, actualIntAryUint8Len)

		return
	}

	for i := 0; i < actualIntAryUint8Len; i++ {

		if inputArrayUint8[i] != actualIntAryUint8Elements[i] {
			t.Errorf("%v\n"+
				"Error: Input and Actual Array Elements Don't Match!\n"+
				"Because inputArrayUint8[%v] != actualIntAryUint8Elements[%v]\n"+
				"Expected actualIntAryUint8Elements[%v] = '%v'\n"+
				"  Actual actualIntAryUint8Elements[%v] = '%v'\n"+
				"Expected Number String = '%v'\n"+
				"ExpectedIntAry Number String  = '%v'\n",
				ePrefix, i, i, i, expectedArrayUint8[i], i, actualIntAryUint8Elements[i],
				expectedNumStr, actualIntAryNumStr)

			return
		}

	}

	return
}
