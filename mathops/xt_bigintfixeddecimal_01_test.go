package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntFixedDecimal_Ceiling_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_01"

	numStr := "5.95"

	expectedNumStr := "6"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("Error returned by new(BigIntFixedDecimal)."+
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())

		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_02"
	numStr := "5.05"
	expectedNumStr := "6"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(numStr, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_03"
	numStr := "5"
	expectedNumStr := "5"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			" NewNumStr(numStr, '.')\n"+
			"  numStr'='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_04(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_04"
	numStr := "-5.05"
	expectedNumStr := "-5"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err :=.NewNumStr("+
			"  numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_05(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_05"
	numStr := "2.4"
	expectedNumStr := "3"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_06(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_06"

	numStr := "2.9"
	expectedNumStr := "3"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_07(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_07"

	numStr := "-2.7"
	expectedNumStr := "-2"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_08(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_08"
	numStr := "-2"
	expectedNumStr := "-2"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_09(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_09"

	numStr := "0"
	expectedNumStr := "0"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ia.GetNumStr()== '%v'\n"+
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_10(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_10"
	numStr := "0.00000"
	expectedNumStr := "0"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(numStr, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_11(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_11"
	numStr := "159876231.9999999999"
	expectedNumStr := "159876232"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
		return
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_12(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_12"
	numStr := "-159876231.9999999999"
	expectedNumStr := "-159876231"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(numStr, '.')\n"+
			" numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_13(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_13"

	numStr := "159876231.0000000000000001"

	expectedNumStr := "159876232"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_14(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_14"

	numStr := "-159876231.0000000000000001"

	expectedNumStr := "-159876231"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_15(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_15"

	numStr := "-0.0000000000000001"

	expectedNumStr := "0"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Ceiling_16(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_16"
	numStr := "0.0000000000000001"
	expectedNumStr := "1"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
		return
	}
}

func TestBigIntFixedDecimal_Ceiling_17(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Ceiling_17"
	numStr := "-0.0000000000000001"
	expectedNumStr := "0"

	fixDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	ceiling, err := fixDec.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceiling, err := fixDec.Ceiling()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ceilingNumStr, err := ceiling.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ceilingNumStr, err := ceiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != ceilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected ceiling.GetNumStr()== '%v'\n"+
			"Instead ceiling.GetNumStr() == '%v'\n\n",
			ePrefix, expectedNumStr, ceilingNumStr)
	}

	return
}

func TestBigIntFixedDecimal_ChangeSign_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_ChangeSign_01"

	numStr := "859"
	expectedStr := "-859"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.ChangeSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDec.GetNumStr()== '%v'\n"+
			"Instead fixedDec.GetNumStr() == '%v'\n\n",
			ePrefix, expectedStr, fixedDecNumStr)
	}

	return
}

func TestBigIntFixedDecimal_ChangeSign_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_ChangeSign_02"

	numStr := "-859"
	expectedStr := "859"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.ChangeSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDec.GetNumStr()== '%v'\n"+
			"Instead fixedDec.GetNumStr() == '%v'\n\n",
			ePrefix, expectedStr, fixedDecNumStr)
	}

	return
}

func TestBigIntFixedDecimal_ChangeSign_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_ChangeSign_03"
	numStr := "859.123456"
	expectedStr := "-859.123456"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.ChangeSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDec.GetNumStr()== '%v'\n"+
			"Instead fixedDec.GetNumStr() == '%v'\n\n",
			ePrefix, expectedStr, fixedDecNumStr)
	}

	return
}

func TestBigIntFixedDecimal_ChangeSign_04(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_ChangeSign_04"

	numStr := "-859.123456"
	expectedStr := "859.123456"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.ChangeSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDec.GetNumStr()== '%v'\n"+
			"Instead fixedDec.GetNumStr() == '%v'\n\n",
			ePrefix, expectedStr, fixedDecNumStr)
	}

	return
}

func TestBigIntFixedDecimal_ChangeSign_05(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_ChangeSign_05"

	numStr := "0"
	expectedStr := "0"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.ChangeSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDec.GetNumStr()== '%v'\n"+
			"Instead fixedDec.GetNumStr() == '%v'\n\n",
			ePrefix, expectedStr, fixedDecNumStr)
	}

	return
}

func TestBigIntFixedDecimal_ChangeSign_06(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_ChangeSign_06"

	numStr := "0.000"
	expectedStr := "0.000"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.ChangeSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDec.GetNumStr()== '%v'\n"+
			"Instead fixedDec.GetNumStr() == '%v'\n\n",
			ePrefix, expectedStr, fixedDecNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_01"

	num1Str := "5"
	num2Str := "2"
	expectedResult := 1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\nError='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual BigIntFixedDecimal fd1NumStr1 = %v\n"+
			"Actual BigIntFixedDecimal fd2NumStr1 = %v\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_02"
	num1Str := "5.2"
	num2Str := "5.1"
	expectedResult := 1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\nError='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\nError='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual BigIntFixedDecimal fd1NumStr1 = %v\n"+
			"Actual BigIntFixedDecimal fd2NumStr1 = %v\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_03"
	num1Str := "5.2"
	num2Str := "5.2"
	expectedResult := 0

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\nError='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\nError='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual BigIntFixedDecimal fd1NumStr1 = %v\n"+
			"Actual BigIntFixedDecimal fd2NumStr1 = %v\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_04(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_04"

	num1Str := "837123.4"
	num2Str := "837123.5"
	expectedResult := -1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual BigIntFixedDecimal fd1NumStr1 = %v\n"+
			"Actual BigIntFixedDecimal fd2NumStr1 = %v\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_05(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_05"

	num1Str := "0"
	num2Str := "0.1"
	expectedResult := -1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result == '%v'\n"+
			"Instead  Compare Result == '%v'\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual BigIntFixedDecimal fd1NumStr1 = %v\n"+
			"Actual BigIntFixedDecimal fd2NumStr1 = %v\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_06(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_06"
	num1Str := "35.123456"
	num2Str := "40.5"
	expectedResult := -1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual BigIntFixedDecimal fd1NumStr == '%v'\n"+
			"Actual BigIntFixedDecimal fd2NumStr == '%v'\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_07(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_07"
	num1Str := "35.123456"
	num2Str := "2.5"
	expectedResult := 1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual fd1NumStr == '%v'\n"+
			"Actual fd2NumStr == '%v'\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_08(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_08"
	num1Str := "35.123456"
	num2Str := "2.123456789012345"
	expectedResult := 1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual fd1NumStr == '%v'\n"+
			"Actual fd2NumStr == '%v'\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

}

func TestBigIntFixedDecimal_Cmp_09(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_09"
	num1Str := "-35.123456"
	num2Str := "2.123456789012345"
	expectedResult := -1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str,'.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual fd1NumStr == '%v'\n"+
			"Actual fd2NumStr == '%v'\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_10(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_10"
	num1Str := "-35.123456"
	num2Str := "-35.123455"
	expectedResult := -1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"cmpResult, err := fd1.Cmp(fd2)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual fd1NumStr == '%v'\n"+
			"Actual fd2NumStr == '%v'\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_Cmp_11(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Cmp_11"
	num1Str := "-35.123455"
	num2Str := "-35.123456"
	expectedResult := 1

	fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1, err := new(BigIntFixedDecimal).NewNumStr(num1Str, '.')\n"+
			"num1Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num1Str, err.Error())
		return
	}

	fd1NumStr, err := fd1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd1NumStr, err := fd1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2, err := new(BigIntFixedDecimal).NewNumStr(num2Str, '.')\n"+
			"num2Str='%v'\n"+
			"Error='%v'\n\n", ePrefix, num2Str, err.Error())
		return
	}

	fd2NumStr, err := fd2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fd2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	cmpResult, err := fd1.Cmp(fd2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != cmpResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Original numStr1 = %v\n"+
			"Original numStr2 = %v\n"+
			"Actual fd1NumStr == '%v'\n"+
			"Actual fd2NumStr == '%v'\n\n",
			ePrefix, expectedResult, cmpResult, num1Str, num2Str, fd1NumStr, fd2NumStr)
	}

}

func TestBigIntFixedDecimal_CmpZero_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CmpZero_01"
	numStr := "123.45"
	expectedCmpResult := 1

	fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fdNumStr, err := fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdNumStr, err := fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualCmp, err := fd.CmpZero()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCmp, err := fd.CmpZero()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Expected fdNumStr = '%v'\n"+
			"Actual fdNumStr = '%v'\n\n",
			ePrefix, expectedCmpResult, actualCmp, numStr, fdNumStr)
	}

	return
}

func TestBigIntFixedDecimal_CmpZero_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CmpZero_02"
	numStr := "-123.45"
	expectedCmpResult := -1

	fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fdNumStr, err := fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdNumStr, err := fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualCmp, err := fd.CmpZero()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCmp, err := fd.CmpZero()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Expected fdNumStr = '%v'\n"+
			"Actual fdNumStr = '%v'\n\n",
			ePrefix, expectedCmpResult, actualCmp, numStr, fdNumStr)
	}

	return
}

func TestBigIntFixedDecimal_CmpZero_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CmpZero_03"
	numStr := "0"
	expectedCmpResult := 0

	fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fdNumStr, err := fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdNumStr, err := fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualCmp, err := fd.CmpZero()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCmp, err := fd.CmpZero()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Expected fdNumStr = '%v'\n"+
			"Actual fdNumStr = '%v'\n\n",
			ePrefix, expectedCmpResult, actualCmp, numStr, fdNumStr)
	}

	return
}

func TestBigIntFixedDecimal_CmpZero_04(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CmpZero_04"
	numStr := "0.00"
	expectedCmpResult := 0

	fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fdNumStr, err := fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdNumStr, err := fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualCmp, err := fd.CmpZero()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCmp, err := fd.CmpZero()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Expected fdNumStr = '%v'\n"+
			"Actual fdNumStr = '%v'\n\n",
			ePrefix, expectedCmpResult, actualCmp, numStr, fdNumStr)
	}

	return
}

func TestBigIntFixedDecimal_CmpZero_05(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CmpZero_05"
	numStr := "8"
	expectedCmpResult := 1

	fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fdNumStr, err := fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdNumStr, err := fd.GetNumStr()n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualCmp, err := fd.CmpZero()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCmp, err := fd.CmpZero()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Expected fdNumStr = '%v'\n"+
			"Actual fdNumStr = '%v'\n\n",
			ePrefix, expectedCmpResult, actualCmp, numStr, fdNumStr)
	}

	return
}

func TestBigIntFixedDecimal_CmpZero_06(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_CmpZero_06"
	numStr := "-8"
	expectedCmpResult := -1

	fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fdNumStr, err := fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdNumStr, err := fd.GetNumStr()n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualCmp, err := fd.CmpZero()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCmp, err := fd.CmpZero()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedCmpResult != actualCmp {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected Compare Result = %v\n"+
			"Actual Compare Result = %v\n"+
			"Expected fdNumStr = '%v'\n"+
			"Actual fdNumStr = '%v'\n\n",
			ePrefix, expectedCmpResult, actualCmp, numStr, fdNumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyIn_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyIn_01"

	expectedNumStr := "-123.45"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyIn(fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyIn(fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyIn_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyIn_02"

	expectedNumStr := "123.45"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(2)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyIn(fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyIn(fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)

		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyIn_03(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_CopyIn_03"
	expectedNumStr := "12345"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(0)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyIn(fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyIn(fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyIn_04(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_CopyIn_04"

	expectedNumStr := "-12345"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(0)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyIn(fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyIn(fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyIn_05(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyIn_05"
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyIn(fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyIn(fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyInPtr_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyIn_01"

	expectedNumStr := "-123.45"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyInPtr(&fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyInPtr(&fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyInPtr_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyIn_02"

	expectedNumStr := "123.45"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(2)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyInPtr(&fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyInPtr(&fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyInPtr_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyIn_03"
	expectedNumStr := "12345"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(0)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyInPtr(&fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyInPtr(&fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyInPtr_04(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_CopyIn_04"
	expectedNumStr := "-12345"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(0)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyInPtr(&fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyInPtr(&fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyInPtr_05(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_CopyIn_05"
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2 := BigIntFixedDecimal{}

	err = fD2.CopyInPtr(&fixedDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fD2.CopyInPtr(&fixedDec)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyOut_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyOut_01"

	expectedNumStr := "894.1234"

	originalNum := big.NewInt(8941234)
	originalNumPrecision := uint(4)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2, err := fixedDec.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fD2, err := fixedDec.CopyOut()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyOut_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyOut_02"

	expectedNumStr := "-894.1234"

	originalNum := big.NewInt(-8941234)
	originalNumPrecision := uint(4)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2, err := fixedDec.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fD2, err := fixedDec.CopyOut()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_CopyOut_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_CopyOut_03"
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fD2, err := fixedDec.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fD2, err := fixedDec.CopyOut()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fd2NumStr, err := fD2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fd2NumStr, err := fD2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead fixedDecNumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr)
		return
	}

	if expectedNumStr != fd2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fd2NumStr = '%v'\n"+
			"Instead fd2NumStr == '%v'\n\n",
			ePrefix, expectedNumStr, fd2NumStr)
	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_01"
	expectedNumStr := "-0.12345"

	originalNum := -12345
	originalNumPrecision := uint(2)
	exponent := uint(3)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_02"
	expectedNumStr := "12.345"

	originalNum := 12345
	originalNumPrecision := uint(2)
	exponent := uint(1)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_03"
	expectedNumStr := "-12.345"

	originalNum := -12345
	originalNumPrecision := uint(2)
	exponent := uint(1)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_04(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_04"
	expectedNumStr := "0.00012345"

	originalNum := 12345
	originalNumPrecision := uint(0)
	exponent := uint(8)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_05(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_05"

	expectedNumStr := "-0.00012345"

	originalNum := -12345
	originalNumPrecision := uint(0)
	exponent := uint(8)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_06(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_06"

	expectedNumStr := "0.0057"

	originalNum := 57
	originalNumPrecision := uint(0)
	exponent := uint(4)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_DivideByTenToPwr_07(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_DivideByTenToPwr_07"
	expectedNumStr := "0"

	originalNum := 0
	originalNumPrecision := uint(0)
	exponent := uint(4)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTenToPower(exponent)\n"+
			"exponent= '%v'\nOriginal fixedDec= '%v'\nError='%v'\n\n",
			ePrefix, exponent, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedNumStr, fixedDecNumStr2)

	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_01(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_01"

	num := int64(33333)
	numPrecision := uint(0)
	exponent := uint(8)
	expectedNum := int64(130)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_02(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_02"
	num := int64(33123456)
	numPrecision := uint(3)
	exponent := uint(8)
	expectedNum := int64(129388)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_03(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_03"
	num := int64(4)
	numPrecision := uint(0)
	exponent := uint(9)
	expectedNum := int64(0)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_04(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_04"
	num := int64(-8123456789012345)
	numPrecision := uint(0)
	exponent := uint(12)
	expectedNum := int64(-1983265817630)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_05(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_05"
	num := int64(8123456789012345)
	exponent := uint(12)
	expectedNum := int64(1983265817629)
	numPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_06(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_06"
	num := int64(4)
	exponent := uint(1)
	expectedNum := int64(2)
	numPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_07(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_07"
	num := int64(-4)
	exponent := uint(1)
	expectedNum := int64(-2)
	numPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_08(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_08"
	// Fixed Decimal Initial Value = -40579.123456
	num := int64(-40579123456)
	numPrecision := uint(6)
	exponent := uint(3)
	expectedNum := int64(-5072390432)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_09(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_09"
	// Fixed Decimal Initial Value = 40579.123456
	num := int64(40579123456)
	numPrecision := uint(6)
	exponent := uint(3)
	expectedNum := int64(5072390432)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_10(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_BigIntDividedByTwoToPower_10"
	// Fixed Decimal Initial Value = 67.1234
	num := int64(671234)
	numPrecision := uint(4)
	exponent := uint(2)
	expectedNum := int64(167808)

	fixedDec := new(BigIntFixedDecimal).NewInt64(num, numPrecision)

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedValue := new(BigIntFixedDecimal).NewInt64(expectedNum, 0)

	expectedValueNumStr, err := expectedValue.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedValueNumStr, err := expectedValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = fixedDec.DivideByTwoToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.DivideByTwoToPower(exponent)\n"+
			"Original expectedValueNumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	fixedDecNumStr2, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValueNumStr != fixedDecNumStr2 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr2 = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedValueNumStr, fixedDecNumStr2)
	}

	return
}

func TestBigIntFixedDecimal_Floor_01(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_01"
	num := 0
	precision := uint(0)

	expectedNumStr := "0"

	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_02(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_02"
	num := 4
	precision := uint(0)

	expectedNumStr := "4"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_03(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_03"
	num := 32
	precision := uint(1)

	expectedNumStr := "3"

	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_04(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_04"
	num := 29
	precision := uint(1)

	expectedNumStr := "2"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_05(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_05"
	num := -27
	precision := uint(1)

	expectedNumStr := "-3"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_06(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_06"
	num := -2
	precision := uint(0)

	expectedNumStr := "-2"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_07(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_07"
	num := 595
	precision := uint(2)

	expectedNumStr := "5"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_08(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_08"
	num := 505
	precision := uint(2)

	expectedNumStr := "5"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_09(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_09"
	num := -505
	precision := uint(2)

	expectedNumStr := "-6"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_10(t *testing.T) {
	ePrefix := "TestBigIntFixedDecimal_Floor_10"
	num := 29
	precision := uint(1)

	expectedNumStr := "2"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_11(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_11"
	num := -27
	precision := uint(1)

	expectedNumStr := "-3"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_12(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_12"
	num := 0
	precision := uint(0)

	expectedNumStr := "0"
	fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_13(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_13"
	numStr := "18972.0000000000001"

	expectedNumStr := "18972"

	fixedDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if numStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead, fixedDecNumStr = '%v'\n\n",
			ePrefix, numStr, fixedDecNumStr)

		return
	}

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_14(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_14"
	numStr := "-18972.0000000000001"

	expectedNumStr := "-18973"

	fixedDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if numStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead, fixedDecNumStr = '%v'\n\n",
			ePrefix, numStr, fixedDecNumStr)

		return
	}

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floor, err := fixedDec.Floor()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_15(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_15"
	numStr := "0.0000000000001"

	expectedNumStr := "0"

	fixedDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if numStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead, fixedDecNumStr = '%v'\n\n",
			ePrefix, numStr, fixedDecNumStr)

		return
	}

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Original Fixed Dec NumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_16(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_16"
	numStr := "-189765342891.0000000000001"
	expectedNumStr := "-189765342892"

	fixedDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if numStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead, fixedDecNumStr = '%v'\n\n",
			ePrefix, numStr, fixedDecNumStr)

		return
	}

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Original Fixed Dec NumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}

func TestBigIntFixedDecimal_Floor_17(t *testing.T) {

	ePrefix := "TestBigIntFixedDecimal_Floor_17"
	numStr := "189765342891.0000000000001"
	expectedNumStr := "189765342891"

	fixedDec, err :=
		new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if numStr != fixedDecNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fixedDecNumStr = '%v'\n"+
			"Instead, fixedDecNumStr = '%v'\n\n",
			ePrefix, numStr, fixedDecNumStr)

		return
	}

	floor, err := fixedDec.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Original Fixed Dec NumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, fixedDecNumStr, err.Error())
		return
	}

	floorNumStr, err := floor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorNumStr, err := floor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != floorNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected floorNumStr = '%v'\n"+
			"Instead, floorNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorNumStr)
	}

	return
}
