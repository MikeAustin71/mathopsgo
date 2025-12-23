package mathops

import (
	"testing"
)

func TestPrecisionScaleMode_String_01(t *testing.T) {

	ePrefix := "TestPrecisionScaleMode_String_01"

	actualConstantParam := SCALEPRECISIONRIGHT

	expectedStr := "ScalePrecisionRight"

	resultConstantParamStr := actualConstantParam.String()

	if expectedStr != resultConstantParamStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Parameter Strings DON'T MATCH!\n"+
			"Because expectedStr != resultConstantParamStr\n"+
			"Expected resultConstantParamStr = '%v'\n"+
			"  Actual resultConstantParamStr = '%v'\n\n",
			ePrefix, expectedStr, resultConstantParamStr)

		return
	}

	return
}

func TestPrecisionScaleMode_String_02(t *testing.T) {

	ePrefix := "TestPrecisionScaleMode_String_02"

	actualConstantParam := SCALEPRECISIONLEFT

	expectedStr := "ScalePrecisionLeft"

	resultConstantParamStr := actualConstantParam.String()

	if expectedStr != resultConstantParamStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Parameter Strings DON'T MATCH!\n"+
			"Because expectedStr != resultConstantParamStr\n"+
			"Expected resultConstantParamStr = '%v'\n"+
			"  Actual resultConstantParamStr = '%v'\n\n",
			ePrefix, expectedStr, resultConstantParamStr)

		return
	}

	return
}

func TestPrecisionScaleMode_Value_01(t *testing.T) {

	ePrefix := "TestPrecisionScaleMode_Value_01"

	expectedConstantInt := 0

	var precisionScaleModeFromConstant PrecisionScaleMode

	var resultConstantInt int

	precisionScaleModeFromConstant = SCALEPRECISIONRIGHT

	resultConstantInt = int(precisionScaleModeFromConstant)

	if expectedConstantInt != resultConstantInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Constant Integer Values DON'T MATCH!\n"+
			"Because expectedConstantInt != resultConstantInt\n"+
			"SCALEPRECISIONRIGHT Value is Wrong!\n"+
			"Expected resultConstantInt = '%v'\n"+
			"  Actual resultConstantInt = '%v'\n\n",
			ePrefix, expectedConstantInt, resultConstantInt)

		return
	}

	return
}

func TestPrecisionScaleMode_Value_02(t *testing.T) {

	ePrefix := "TestPrecisionScaleMode_Value_02"

	expectedConstantInt := 1

	var precisionScaleModeFromConstant PrecisionScaleMode

	var resultConstantInt int

	precisionScaleModeFromConstant = SCALEPRECISIONLEFT

	resultConstantInt = int(precisionScaleModeFromConstant)

	if expectedConstantInt != resultConstantInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Constant Integer Values DON'T MATCH!\n"+
			"Because expectedConstantInt != resultConstantInt\n"+
			"SCALEPRECISIONLEFT Value is Wrong!\n"+
			"Expected resultConstantInt = '%v'\n"+
			"  Actual resultConstantInt = '%v'\n\n",
			ePrefix, expectedConstantInt, resultConstantInt)

		return
	}

	return
}
