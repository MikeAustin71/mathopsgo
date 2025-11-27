package mathops

import (
	"testing"
)

func TestIntAry_SetIntAryWithInt32_01(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt32_01"

	originalNumInt32 := int32(123456789)

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "123456.789"

	expectedPrecisionUint := uint(3)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt32(originalNumInt32, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt32(\n"+
			"  originalNumInt32, expectedPrecisionUint)\n"+
			"originalNumInt32= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			expectedPrecisionUint,
			err.Error())

		return
	}
	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt32_02(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt32_02"

	originalNumInt32 := int32(-123456789)

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "-12345.6789"

	expectedPrecisionUint := uint(4)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt32(originalNumInt32, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt32(\n"+
			"  originalNumInt32, expectedPrecisionUint)\n"+
			"originalNumInt32= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			expectedPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt32_03(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt32_03"

	originalNumInt32 := int32(0)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0000"

	expectedPrecisionUint := uint(4)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt32(originalNumInt32, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt32(\n"+
			"  originalNumInt32, expectedPrecisionUint)\n"+
			"originalNumInt32= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			expectedPrecisionUint,
			err.Error())

		return
	}
	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt32_04(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt32_04"

	originalNumInt32 := int32(32)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0032"

	expectedPrecisionUint := uint(4)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt32(originalNumInt32, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt32(\n"+
			"  originalNumInt32, expectedPrecisionUint)\n"+
			"originalNumInt32= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			expectedPrecisionUint,
			err.Error())

		return
	}
	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt32_05(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt32_05"

	originalNumInt32 := int32(-32)

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr := "-32"

	expectedPrecisionUint := uint(0)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt32(originalNumInt32, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt32(\n"+
			"  originalNumInt32, expectedPrecisionUint)\n"+
			"originalNumInt32= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			expectedPrecisionUint,
			err.Error())

		return
	}
	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt32_06(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt32_06"

	originalNumInt32 := int32(32)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.32"

	expectedPrecisionUint := uint(2)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt32(originalNumInt32, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt32(\n"+
			"  originalNumInt32, expectedPrecisionUint)\n"+
			"originalNumInt32= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			expectedPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint64_01(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint64_01"

	originalNumUint64 := uint64(123456789)

	originalPrecisionUint := uint(3)

	originalSignValue := 1

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "123456.789"

	expectedPrecisionUint := uint(3)

	expectedPrecisionInt := 3

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint64(originalNumUint64, originalSignValue, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint64(\n"+
			"  originalNumUint64, originalSignValue, originalPrecisionUint)n"+
			"originalNumUint64= '%v'\n"+
			"originalSignValue= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint64,
			originalSignValue,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint64_02(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint64_02"

	originalNumUint64 := uint64(123456789)

	originalPrecisionIntValue := 5

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr := "1234.56789"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint64(originalNumUint64, originalSignValue, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint64(\n"+
			"  originalNumUint64, originalSignValue, originalPrecisionUint)n"+
			"originalNumUint64= '%v'\n"+
			"originalSignValue= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint64,
			originalSignValue,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint64_03(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint64_03"

	originalNumUint64 := uint64(0)

	originalPrecisionIntValue := 4

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0000"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint64(originalNumUint64, originalSignValue, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint64(\n"+
			"  originalNumUint64, originalSignValue, originalPrecisionUint)n"+
			"originalNumUint64= '%v'\n"+
			"originalSignValue= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint64,
			originalSignValue,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint64_04(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint64_04"

	originalNumUint64 := uint64(32)

	originalPrecisionIntValue := 4

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0032"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint64(originalNumUint64, originalSignValue, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint64(\n"+
			"  originalNumUint64, originalSignValue, originalPrecisionUint)n"+
			"originalNumUint64= '%v'\n"+
			"originalSignValue= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint64,
			originalSignValue,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint64_05(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint64_05"

	originalNumUint64 := uint64(32)

	originalPrecisionIntValue := 0

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "32"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint64(originalNumUint64, originalSignValue, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint64(\n"+
			"  originalNumUint64, originalSignValue, originalPrecisionUint)n"+
			"originalNumUint64= '%v'\n"+
			"originalSignValue= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint64,
			originalSignValue,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint64_06(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint64_06"

	originalNumUint64 := uint64(32)

	originalPrecisionIntValue := 2

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.32"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint64(originalNumUint64, originalSignValue, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint64(\n"+
			"  originalNumUint64, originalSignValue, originalPrecisionUint)n"+
			"originalNumUint64= '%v'\n"+
			"originalSignValue= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint64,
			originalSignValue,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt64_01(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt64_01"

	originalNumInt64 := int64(123456789)

	originalPrecisionIntValue := 3

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "123456.789"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)\n"+
			"originalNumInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt64_02(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt64_02"

	originalNumInt64 := int64(-123456789)

	originalPrecisionIntValue := 4

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := -1

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "-12345.6789"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)\n"+
			"originalNumInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt64_03(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt64_03"

	originalNumInt64 := int64(0)

	originalPrecisionIntValue := 4

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                         1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0000"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)\n"+
			"originalNumInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt64_04(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt64_04"

	originalNumInt64 := int64(32)

	originalPrecisionIntValue := 4

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0032"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)\n"+
			"originalNumInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt64_05(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt64_05"

	originalNumInt64 := int64(-32)

	originalPrecisionIntValue := 0

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := -1

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr := "-32"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)\n"+
			"originalNumInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithInt64_06(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithInt64_05"

	originalNumInt64 := int64(32)

	originalPrecisionIntValue := 2

	originalPrecisionUint := uint(originalPrecisionIntValue)

	originalSignValue := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.32"

	expectedPrecisionUint := uint(originalPrecisionIntValue)

	expectedPrecisionInt := originalPrecisionIntValue

	expectedSignValue := originalSignValue

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithInt64(originalNumInt64, originalPrecisionUint)\n"+
			"originalNumInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithIntAry_01(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithIntAry_01"

	intAryDigits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	intAryDigitsLen := len(intAryDigits)

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "123456.789"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithIntAry(intAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithIntAry(\n"+
			"  intAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"intAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if intAryDigitsLen != iaLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAryDigitsLen != iaLen\n"+
			"Expected iaLen = '%v'\n"+
			"  Actual iaLen = '%v'\n\n",
			ePrefix, intAryDigitsLen, iaLen)

		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if intAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, intAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = uint8(intAryDigits[i])
		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithIntAry_02(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithIntAry_02"

	intAryDigits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	intAryDigitsLen := len(intAryDigits)

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "-12345.6789"

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithIntAry(intAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithIntAry(\n"+
			"  intAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"intAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if intAryDigitsLen != iaLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAryDigitsLen != iaLen\n"+
			"Expected iaLen = '%v'\n"+
			"  Actual iaLen = '%v'\n\n",
			ePrefix, intAryDigitsLen, iaLen)

		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if intAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, intAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = uint8(intAryDigits[i])
		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithIntAry_03(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithIntAry_03"

	originalIntAryDigits := []int{3, 2}

	expectedIntAryDigits := []int{0, 0, 0, 3, 2}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0032"

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithIntAry(originalIntAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithIntAry(\n"+
			"  originalIntAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"originalIntAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalIntAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedIntAryDigitsLen != iaLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != iaLen\n"+
			"Expected iaLen = '%v'\n"+
			"  Actual iaLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, iaLen)

		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = uint8(expectedIntAryDigits[i])
		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithIntAry_04(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithIntAry_04"

	originalIntAryDigits := []int{3, 2}

	expectedIntAryDigits := []int{0, 3, 2}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.32"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithIntAry(originalIntAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithIntAry(\n"+
			"  originalIntAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"originalIntAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalIntAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedIntAryDigitsLen != iaLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != iaLen\n"+
			"Expected iaLen = '%v'\n"+
			"  Actual iaLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, iaLen)

		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, actualIntAryLen)

		return
	}

	if expectedIntAryDigitsLen != iaLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != iaLen\n"+
			"Expected iaLen = '%v'\n"+
			"  Actual iaLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, iaLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = uint8(expectedIntAryDigits[i])
		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithIntAryObj_01(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithIntAryObj_01"

	originalNumberStr1 := "0000589432.607528000"

	originalNumberStr2 := "-028.3700"

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry1.SetIntAryWithNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr1 != intAry1NumberStr \n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAryFinal := new(IntAry).New()

	err = intAryFinal.SetIntAryWithNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.SetIntAryWithNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = intAryFinal.IsValid("Validating initial intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating initial intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"intAryFinal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr2 != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAryFinalNumberStr)

		return
	}

	intAry1PrecisionInt := intAry1.GetPrecision()

	intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1PrecisionUint, err :=\n"+
			"  intAry1.GetPrecisionUint()\n"+
			"intAry1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry1SignValue, err := intAry1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1SignValue, err := intAry1.GetSign()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
			"intAry1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAryFinal.SetIntAryWithIntAryObj(&intAry1, false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.SetIntAryWithIntAryObj(\n"+
			"  &intAry1, false)\n"+
			"intAry1= '%v'\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAryFinalNumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating final intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating final intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err = intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err = intAryFinal.GetNumStr()\n"+
			"intAryFinal set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if !intAry1.Equals(&intAryFinal) {
		t.Errorf("%v\n"+
			"Error: intAry1 is NOT EQUAL to intAryFinal!!!\n"+
			"Because!!!!\n"+
			"Expected intAryFinal = '%v'\n"+
			"  Actual intAryFinal = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAryFinalNumberStr)

		return
	}

	if intAry1NumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because intAry1NumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAryFinalNumberStr)

		return
	}

	if intAry1PrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: intAry1 & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because intAry1PrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, intAry1PrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if intAry1PrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because intAry1PrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, intAry1PrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if intAry1SignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: intAry1 & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because intAry1SignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, intAry1SignValue, intAryFinalSignValue)

		return
	}

	if !intAry1NumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because intAry1NumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, intAry1NumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithIntAryObj_02(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithIntAryObj_02"

	originalNumberStr1 := "0000589432.607528000"

	originalNumberStr2 := "-028.3700"

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry1.SetIntAryWithNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr1 != intAry1NumberStr \n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAryFinal := new(IntAry).New()

	err = intAryFinal.SetIntAryWithNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.SetIntAryWithNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = intAryFinal.IsValid("Validating initial intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating initial intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"intAryFinal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr2 != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAryFinalNumberStr)

		return
	}

	intAry1PrecisionInt := intAry1.GetPrecision()

	intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1PrecisionUint, err :=\n"+
			"  intAry1.GetPrecisionUint()\n"+
			"intAry1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry1SignValue, err := intAry1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1SignValue, err := intAry1.GetSign()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
			"intAry1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAryFinal.SetIntAryWithIntAryObj(&intAry1, false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.SetIntAryWithIntAryObj(\n"+
			"  &intAry1, false)\n"+
			"intAry1= '%v'\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAryFinalNumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating final intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating final intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err = intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err = intAryFinal.GetNumStr()\n"+
			"intAryFinal set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if !intAry1.Equals(&intAryFinal) {
		t.Errorf("%v\n"+
			"Error: intAry1 is NOT EQUAL to intAryFinal!!!\n"+
			"Because!!!!\n"+
			"Expected intAryFinal = '%v'\n"+
			"  Actual intAryFinal = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAryFinalNumberStr)

		return
	}

	if intAry1NumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because intAry1NumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAryFinalNumberStr)

		return
	}

	if intAry1PrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: intAry1 & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because intAry1PrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, intAry1PrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if intAry1PrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because intAry1PrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, intAry1PrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if intAry1SignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: intAry1 & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because intAry1SignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, intAry1SignValue, intAryFinalSignValue)

		return
	}

	if !intAry1NumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because intAry1NumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, intAry1NumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetIntAryWithUint8Ary_01(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint8Ary_01"

	originalIntAryDigits := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expectedIntAryDigits := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "123456.789"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint8Ary(originalIntAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint8Ary(\n"+
			"  originalIntAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"originalIntAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalIntAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = expectedIntAryDigits[i]
		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithUint8Ary_02(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint8Ary_02"

	originalIntAryDigits := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expectedIntAryDigits := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	expectedNumberStr := "-12345.6789"

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint8Ary(originalIntAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint8Ary(\n"+
			"  originalIntAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"originalIntAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalIntAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = expectedIntAryDigits[i]

		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithUint8Ary_03(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint8Ary_03"

	originalIntAryDigits := []uint8{3, 2}

	expectedIntAryDigits := []uint8{0, 0, 0, 3, 2}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0032"

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint8Ary(originalIntAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint8Ary(\n"+
			"  originalIntAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"originalIntAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalIntAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = expectedIntAryDigits[i]

		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetIntAryWithUint8Ary_04(t *testing.T) {

	ePrefix := "TestIntAry_SetIntAryWithUint8Ary_04"

	originalIntAryDigits := []uint8{3, 2}

	expectedIntAryDigits := []uint8{0, 3, 2}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.32"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithUint8Ary(originalIntAryDigits, expectedPrecisionUint, expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithUint8Ary(\n"+
			"  originalIntAryDigits, expectedPrecisionUint, expectedSignValue)\n"+
			"originalIntAryDigits= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalIntAryDigits,
			expectedPrecisionUint,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	iaStats := intAry.GetIntAryStats()

	intAryElements, iaLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	actualIntAryLen := iaStats.IntAryLen

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, expectedIntAryDigitsLen, actualIntAryLen)

		return
	}

	var expectedIntAryDigit, actualIntAryDigit uint8

	for i := 0; i < iaLen; i++ {

		expectedIntAryDigit = expectedIntAryDigits[i]

		actualIntAryDigit = intAryElements[i]

		if expectedIntAryDigit != actualIntAryDigit {
			t.Errorf("%v\n"+
				"Error: Expected and Actual IntAry Digits DO NOT MATCH!\n"+
				"Because expectedIntAryDigit != actualIntAryDigit\n"+
				"Expected actualIntAryDigit = '%v'\n"+
				"  Actual actualIntAryDigit = '%v'\n"+
				"Cycle Number i= '%v' \n\n",
				ePrefix, expectedIntAryDigit, actualIntAryDigit, i)

			return
		}

	}

	return
}

func TestIntAry_SetSign_01(t *testing.T) {

	ePrefix := "TestIntAry_SetSign_01"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	originalNumberStr := "123.456"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr := "-123.456"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetSign(expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetSign(expectedSignValue)\n"+
			"intAry= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetSign_02(t *testing.T) {

	ePrefix := "TestIntAry_SetSign_02"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	originalNumberStr := "123.456"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr := "123.456"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetSign(expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetSign(expectedSignValue)\n"+
			"intAry= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetSign_03(t *testing.T) {

	ePrefix := "TestIntAry_SetSign_03"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	originalNumberStr := "-123.456"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr := "123.456"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetSign(expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetSign(expectedSignValue)\n"+
			"intAry= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetSign_04(t *testing.T) {

	ePrefix := "TestIntAry_SetSign_04"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	originalNumberStr := "-123.456"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr := "-123.456"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetSign(expectedSignValue)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetSign(expectedSignValue)\n"+
			"intAry= '%v'\n"+
			"expectedSignValue= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedSignValue,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetSign_05(t *testing.T) {

	ePrefix := "TestIntAry_SetSign_05"

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	originalNumberStr := "0.0"

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "0.0"

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	attemptedSignVal := -1

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	// This should fail to reset the sign value
	// because the underlying numeric value is zero.
	// No error should be produced.
	err = intAry.SetSign(attemptedSignVal)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetSign(attemptedSignVal)\n"+
			"intAry= '%v'\n"+
			"attemptedSignVal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			attemptedSignVal,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetWithNumStr_01(t *testing.T) {

	ePrefix := "TestIntAry_SetWithNumStr_01"

	originalNumberStr := "123.456"

	expectedIntAryDigits := []uint8{1, 2, 3, 4, 5, 6}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	expectedNumberStr := originalNumberStr

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	//
	//err = intAry.SetIntAryLength()
	//
	// if err != nil {
	//   t.Errorf("%v\n" +
	//     "Error returned by:\n"+
	//     "intAry.SetIntAryLength()\n"+
	//     "intAry= '%v'\n" +
	//     "Error= '%v'\n\n",
	//     ePrefix, originalNumberStr, err.Error())
	//   return
	// }

	err = intAry.IsValid("Validating intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	actualIntAryLen := intAry.GetIntAryLength()

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: IntAry Digit Lengths ARE NOT EQUAL!`\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, actualIntAryLen, actualIntAryLen)

		return
	}

	var actualElementUint8, expectedElementUint8 uint8

	for i := 0; i < expectedIntAryDigitsLen; i++ {

		actualElementUint8, err = intAry.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualElementUint8, err = intAry.GetIntAryElement(i)\n"+
				"intAry= '%v'\n"+
				"Cycle Number i = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		expectedElementUint8 = expectedIntAryDigits[i]

		if expectedElementUint8 != actualElementUint8 {
			t.Errorf("%v\n"+
				"Error: IntAry Element Values DO NOT MATCH!\n"+
				"Because expectedElementUint8 != actualElementUint8\n"+
				"Expected actualElementUint8 = '%v'\n"+
				"  Actual actualElementUint8 = '%v'\n\n",
				ePrefix, expectedElementUint8, actualElementUint8)

			return
		}

	}

	return
}

func TestIntAry_SetWithNumStr_02(t *testing.T) {

	ePrefix := "TestIntAry_SetWithNumStr_02"

	originalNumberStr := "-12345.9"

	expectedIntAryDigits := []uint8{1, 2, 3, 4, 5, 9}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	expectedNumberStr := originalNumberStr

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.SetIntAryLength()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry.SetIntAryLength()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	actualIntAryLen := intAry.GetIntAryLength()

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: IntAry Digit Lengths ARE NOT EQUAL!`\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, actualIntAryLen, actualIntAryLen)

		return
	}

	var actualElementUint8, expectedElementUint8 uint8

	for i := 0; i < expectedIntAryDigitsLen; i++ {

		actualElementUint8, err = intAry.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualElementUint8, err = intAry.GetIntAryElement(i)\n"+
				"intAry= '%v'\n"+
				"Cycle Number i = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		expectedElementUint8 = expectedIntAryDigits[i]

		if expectedElementUint8 != actualElementUint8 {
			t.Errorf("%v\n"+
				"Error: IntAry Element Values DO NOT MATCH!\n"+
				"Because expectedElementUint8 != actualElementUint8\n"+
				"Expected actualElementUint8 = '%v'\n"+
				"  Actual actualElementUint8 = '%v'\n\n",
				ePrefix, expectedElementUint8, actualElementUint8)

			return
		}

	}

	return
}

func TestIntAry_SetWithNumStr_03(t *testing.T) {

	ePrefix := "TestIntAry_SetWithNumStr_03"

	originalNumberStr := "-123  45.9"

	expectedIntAryDigits := []uint8{1, 2, 3, 4, 5, 9}

	expectedIntAryDigitsLen := len(expectedIntAryDigits)

	expectedNumberStr := "-12345.9"

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.SetIntAryLength()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry.SetIntAryLength()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	actualIntAryLen := intAry.GetIntAryLength()

	if expectedIntAryDigitsLen != actualIntAryLen {
		t.Errorf("%v\n"+
			"Error: IntAry Digit Lengths ARE NOT EQUAL!`\n"+
			"Because expectedIntAryDigitsLen != actualIntAryLen\n"+
			"Expected actualIntAryLen = '%v'\n"+
			"  Actual actualIntAryLen = '%v'\n\n",
			ePrefix, actualIntAryLen, actualIntAryLen)

		return
	}

	var actualElementUint8, expectedElementUint8 uint8

	for i := 0; i < expectedIntAryDigitsLen; i++ {

		actualElementUint8, err = intAry.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualElementUint8, err = intAry.GetIntAryElement(i)\n"+
				"intAry= '%v'\n"+
				"Cycle Number i = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		expectedElementUint8 = expectedIntAryDigits[i]

		if expectedElementUint8 != actualElementUint8 {
			t.Errorf("%v\n"+
				"Error: IntAry Element Values DO NOT MATCH!\n"+
				"Because expectedElementUint8 != actualElementUint8\n"+
				"Expected actualElementUint8 = '%v'\n"+
				"  Actual actualElementUint8 = '%v'\n\n",
				ePrefix, expectedElementUint8, actualElementUint8)

			return
		}

	}

	return
}

func TestIntAry_SetPrecision_01(t *testing.T) {

	ePrefix := "TestIntAry_SetPrecision_01"

	originalNumberStr := "99.995"

	expectedNumberStr := "99.99"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var roundResult bool

	roundResult = false

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetPrecision(expectedPrecisionInt, roundResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetPrecision(expectedPrecisionInt, roundResult)\n"+
			"intAry= '%v'\n"+
			"expectedPrecisionInt= '%v'\n"+
			"roundResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedPrecisionInt,
			roundResult,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetPrecision_02(t *testing.T) {

	ePrefix := "TestIntAry_SetPrecision_02"

	originalNumberStr := "100.00"

	expectedNumberStr := originalNumberStr

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var roundResult bool

	roundResult = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetPrecision(expectedPrecisionInt, roundResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetPrecision(expectedPrecisionInt, roundResult)\n"+
			"intAry= '%v'\n"+
			"expectedPrecisionInt= '%v'\n"+
			"roundResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedPrecisionInt,
			roundResult,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetPrecision_03(t *testing.T) {

	ePrefix := "TestIntAry_SetPrecision_03"

	originalNumberStr := "-0"

	expectedNumberStr := "0.00"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var roundResult bool

	roundResult = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetPrecision(expectedPrecisionInt, roundResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetPrecision(expectedPrecisionInt, roundResult)\n"+
			"intAry= '%v'\n"+
			"expectedPrecisionInt= '%v'\n"+
			"roundResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedPrecisionInt,
			roundResult,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetPrecision_04(t *testing.T) {

	ePrefix := "TestIntAry_SetPrecision_04"

	originalNumberStr := "-999.995"

	expectedNumberStr := "-1000.00"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var roundResult bool

	roundResult = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetPrecision(expectedPrecisionInt, roundResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetPrecision(expectedPrecisionInt, roundResult)\n"+
			"intAry= '%v'\n"+
			"expectedPrecisionInt= '%v'\n"+
			"roundResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedPrecisionInt,
			roundResult,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetPrecision_05(t *testing.T) {

	ePrefix := "TestIntAry_SetPrecision_05"

	originalNumberStr := "-999.995"

	expectedNumberStr := "-999.995"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var roundResult bool

	roundResult = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetPrecision(expectedPrecisionInt, roundResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetPrecision(expectedPrecisionInt, roundResult)\n"+
			"intAry= '%v'\n"+
			"expectedPrecisionInt= '%v'\n"+
			"roundResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedPrecisionInt,
			roundResult,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_SetPrecision_06(t *testing.T) {

	ePrefix := "TestIntAry_SetPrecision_05"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	originalNumberStr := "-999.995"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr := "-999.995000"

	expectedPrecisionInt := 6

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var roundResult bool

	roundResult = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.SetPrecision(expectedPrecisionInt, roundResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetPrecision(expectedPrecisionInt, roundResult)\n"+
			"intAry= '%v'\n"+
			"expectedPrecisionInt= '%v'\n"+
			"roundResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedPrecisionInt,
			roundResult,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to final value 'after rounding'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionInt := intAry.GetPrecision()

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionInt != intAryPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryPrecisionInt\n"+
			"Expected intAryPrecisionInt = '%v'\n"+
			"  Actual intAryPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}
