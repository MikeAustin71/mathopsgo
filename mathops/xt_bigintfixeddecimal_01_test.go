package mathops

import (
	"math/big"
	"testing"
)


func TestBigIntFixedDecimal_Ceiling_01(t *testing.T) {

	numStr := "5.95"

	expectedNumStr := "6"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_02(t *testing.T) {

	numStr := "5.05"
	expectedNumStr := "6"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_03(t *testing.T) {

	numStr := "5"
	expectedNumStr := "5"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_04(t *testing.T) {

	numStr := "-5.05"
	expectedNumStr := "-5"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_05(t *testing.T) {

	numStr := "2.4"
	expectedNumStr := "3"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}

}

func TestBigIntFixedDecimal_Ceiling_06(t *testing.T) {

	numStr := "2.9"
	expectedNumStr := "3"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_07(t *testing.T) {

	numStr := "-2.7"
	expectedNumStr := "-2"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_08(t *testing.T) {

	numStr := "-2"
	expectedNumStr := "-2"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_09(t *testing.T) {

	numStr := "0"
	expectedNumStr := "0"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_10(t *testing.T) {

	numStr := "0.00000"
	expectedNumStr := "0"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_11(t *testing.T) {

	numStr := "159876231.9999999999"
	expectedNumStr := "159876232"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_12(t *testing.T) {

	numStr := "-159876231.9999999999"
	expectedNumStr := "-159876231"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	ceiling := fixDec.Ceiling()

	if expectedNumStr != ceiling.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, ceiling.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_13(t *testing.T) {

	numStr := "159876231.0000000000000001"

	expectedNumStr := "159876232"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Ceiling()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_14(t *testing.T) {

	numStr := "-159876231.0000000000000001"

	expectedNumStr := "-159876231"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Ceiling()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_15(t *testing.T) {

	numStr := "-0.0000000000000001"

	expectedNumStr := "0"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Ceiling()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_16(t *testing.T) {

	numStr := "0.0000000000000001"
	expectedNumStr := "1"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Ceiling()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Ceiling_17(t *testing.T) {

	numStr := "-0.0000000000000001"
	expectedNumStr := "0"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Ceiling()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}


func TestBigIntFixedDecimal_ChangeSign_01(t *testing.T) {

	numStr := "859"
	expectedStr := "-859"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	fixedDec.ChangeSign()


	if expectedStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_ChangeSign_02(t *testing.T) {

	numStr := "-859"
	expectedStr := "859"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	fixedDec.ChangeSign()


	if expectedStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_ChangeSign_03(t *testing.T) {

	numStr := "859.123456"
	expectedStr := "-859.123456"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	fixedDec.ChangeSign()


	if expectedStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_ChangeSign_04(t *testing.T) {

	numStr := "-859.123456"
	expectedStr := "859.123456"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	fixedDec.ChangeSign()


	if expectedStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_ChangeSign_05(t *testing.T) {

	numStr := "0"
	expectedStr := "0"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	fixedDec.ChangeSign()


	if expectedStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedStr, fixedDec.GetNumStr())
	}

}


func TestBigIntFixedDecimal_ChangeSign_06(t *testing.T) {

	numStr := "0.000"
	expectedStr := "0.000"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	fixedDec.ChangeSign()


	if expectedStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_Cmp_01(t *testing.T) {

	num1Str := "5"
	num2Str := "2"
	expectedResult := 1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_02(t *testing.T) {

	num1Str := "5.2"
	num2Str := "5.1"
	expectedResult := 1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_03(t *testing.T) {

	num1Str := "5.2"
	num2Str := "5.2"
	expectedResult := 0

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_04(t *testing.T) {

	num1Str := "837123.4"
	num2Str := "837123.5"
	expectedResult := -1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_05(t *testing.T) {

	num1Str := "0"
	num2Str := "0.1"
	expectedResult := -1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_06(t *testing.T) {

	num1Str := "35.123456"
	num2Str := "40.5"
	expectedResult := -1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_07(t *testing.T) {

	num1Str := "35.123456"
	num2Str := "2.5"
	expectedResult := 1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_08(t *testing.T) {

	num1Str := "35.123456"
	num2Str := "2.123456789012345"
	expectedResult := 1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_09(t *testing.T) {

	num1Str := "-35.123456"
	num2Str := "2.123456789012345"
	expectedResult := -1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_10(t *testing.T) {

	num1Str := "-35.123456"
	num2Str := "-35.123455"
	expectedResult := -1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_Cmp_11(t *testing.T) {

	num1Str := "-35.123455"
	num2Str := "-35.123456"
	expectedResult := 1

	fd1, err := BigIntFixedDecimal{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num1Str). " +
			"num1Str='%v' Error='%v'", num1Str, err.Error())
	}

	fd2, err := BigIntFixedDecimal{}.NewNumStr(num2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(num2Str). " +
			"num2Str='%v' Error='%v'", num2Str, err.Error())
	}

	cmpResult := fd1.Cmp(fd2)

	if expectedResult != cmpResult {
		t.Errorf("Error: Expected compare result='%v'. " +
			"Instead, compare result='%v' ",
			expectedResult, cmpResult)
	}

}

func TestBigIntFixedDecimal_CmpZero_01(t *testing.T) {
	numStr := "123.45"
	expectedCmpResult := 1

	fd, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returnd by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	actualCmp := fd.CmpZero()

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

}

func TestBigIntFixedDecimal_CmpZero_02(t *testing.T) {
	numStr := "-123.45"
	expectedCmpResult := -1

	fd, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returnd by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	actualCmp := fd.CmpZero()

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

}

func TestBigIntFixedDecimal_CmpZero_03(t *testing.T) {
	numStr := "0"
	expectedCmpResult := 0

	fd, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returnd by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	actualCmp := fd.CmpZero()

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

}

func TestBigIntFixedDecimal_CmpZero_04(t *testing.T) {
	numStr := "0.00"
	expectedCmpResult := 0

	fd, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returnd by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	actualCmp := fd.CmpZero()

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

}

func TestBigIntFixedDecimal_CmpZero_05(t *testing.T) {
	numStr := "8"
	expectedCmpResult := 1

	fd, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returnd by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	actualCmp := fd.CmpZero()

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

}

func TestBigIntFixedDecimal_CmpZero_06(t *testing.T) {
	numStr := "-8"
	expectedCmpResult := -1

	fd, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returnd by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	actualCmp := fd.CmpZero()

	if expectedCmpResult != actualCmp {
		t.Errorf("Error: Expected compare result='%v'. Instead, actual compare result='%v'. ",
			expectedCmpResult, actualCmp)
	}

}

func TestBigIntFixedDecimal_CopyIn_01(t *testing.T) {
	expectedNumStr := "-123.45"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_02(t *testing.T) {
	expectedNumStr := "123.45"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_03(t *testing.T) {
	expectedNumStr := "12345"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_04(t *testing.T) {
	expectedNumStr := "-12345"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_05(t *testing.T) {
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_01(t *testing.T) {
	expectedNumStr := "-123.45"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_02(t *testing.T) {
	expectedNumStr := "123.45"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_03(t *testing.T) {
	expectedNumStr := "12345"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_04(t *testing.T) {
	expectedNumStr := "-12345"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_05(t *testing.T) {
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyOut_01(t *testing.T) {
	expectedNumStr := "894.1234"

	originalNum := big.NewInt(8941234)
	originalNumPrecision := uint(4)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := fixedDec.CopyOut()

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyOut_02(t *testing.T) {
	expectedNumStr := "-894.1234"

	originalNum := big.NewInt(-8941234)
	originalNumPrecision := uint(4)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := fixedDec.CopyOut()

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyOut_03(t *testing.T) {
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := fixedDec.CopyOut()

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_01(t *testing.T) {

	expectedNumStr := "-0.12345"

	originalNum := -12345
	originalNumPrecision := uint(2)
	exponent := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_02(t *testing.T) {

	expectedNumStr := "12.345"

	originalNum := 12345
	originalNumPrecision := uint(2)
	exponent := uint(1)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_03(t *testing.T) {

	expectedNumStr := "-12.345"

	originalNum := -12345
	originalNumPrecision := uint(2)
	exponent := uint(1)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_04(t *testing.T) {

	expectedNumStr := "0.00012345"

	originalNum := 12345
	originalNumPrecision := uint(0)
	exponent := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_05(t *testing.T) {

	expectedNumStr := "-0.00012345"

	originalNum := -12345
	originalNumPrecision := uint(0)
	exponent := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_06(t *testing.T) {

	expectedNumStr := "0.0057"

	originalNum := 57
	originalNumPrecision := uint(0)
	exponent := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_07(t *testing.T) {

	expectedNumStr := "0"

	originalNum := 0
	originalNumPrecision := uint(0)
	exponent := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPower(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_01(t *testing.T) {

	num := int64(33333)
	numPrecision := uint(0)
	exponent := uint(8)
	expectedNum := int64(130)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_02(t *testing.T) {

	num := int64(33123456)
	numPrecision := uint(3)
	exponent := uint(8)
	expectedNum := int64(129388)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_03(t *testing.T) {

	num := int64(4)
	numPrecision := uint(0)
	exponent := uint(9)
	expectedNum := int64(0)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_04(t *testing.T) {

	num := int64(-8123456789012345)
	numPrecision := uint(0)
	exponent := uint(12)
	expectedNum := int64(-1983265817630)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_05(t *testing.T) {

	num := int64(8123456789012345)
	exponent := uint(12)
	expectedNum := int64(1983265817629)
	numPrecision := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_06(t *testing.T) {

	num := int64(4)
	exponent := uint(1)
	expectedNum := int64(2)
	numPrecision := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_07(t *testing.T) {

	num := int64(-4)
	exponent := uint(1)
	expectedNum := int64(-2)
	numPrecision := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_08(t *testing.T) {

	// Fixed Decimal Initial Value = -40579.123456
	num := int64(-40579123456)
	numPrecision := uint(6)
	exponent := uint(3)
	expectedNum := int64(-5072390432)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_09(t *testing.T) {

	// Fixed Decimal Initial Value = 40579.123456
	num := int64(40579123456)
	numPrecision := uint(6)
	exponent := uint(3)
	expectedNum := int64(5072390432)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_BigIntDividedByTwoToPower_10(t *testing.T) {

	// Fixed Decimal Initial Value = 67.1234
	num := int64(671234)
	numPrecision := uint(4)
	exponent := uint(2)
	expectedNum := int64(167808)

	fixDec := BigIntFixedDecimal{}.NewInt64(num, numPrecision)

	expectedValue := BigIntFixedDecimal{}.NewInt64(expectedNum, 0)

	fixDec.DivideByTwoToPower(exponent)

	if expectedValue.GetNumStr() != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'.",
			expectedValue.GetNumStr(), fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_01(t *testing.T) {
	num := 0
	precision := uint(0)

	expectedNumStr := "0"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_02(t *testing.T) {
	num := 4
	precision := uint(0)

	expectedNumStr := "4"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_03(t *testing.T) {
	num := 32
	precision := uint(1)

	expectedNumStr := "3"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}

}

func TestBigIntFixedDecimal_Floor_04(t *testing.T) {
	num := 29
	precision := uint(1)

	expectedNumStr := "2"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_05(t *testing.T) {
	num := -27
	precision := uint(1)

	expectedNumStr := "-3"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_06(t *testing.T) {
	num := -2
	precision := uint(0)

	expectedNumStr := "-2"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_07(t *testing.T) {
	num := 595
	precision := uint(2)

	expectedNumStr := "5"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_08(t *testing.T) {
	num := 505
	precision := uint(2)

	expectedNumStr := "5"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_09(t *testing.T) {
	num := -505
	precision := uint(2)

	expectedNumStr := "-6"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_10(t *testing.T) {
	num := 29
	precision := uint(1)

	expectedNumStr := "2"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_11(t *testing.T) {
	num := -27
	precision := uint(1)

	expectedNumStr := "-3"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_12(t *testing.T) {
	num := 0
	precision := uint(0)

	expectedNumStr := "0"

	fixDec := BigIntFixedDecimal{}.NewInt(num,precision)

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_13(t *testing.T) {

	numStr := "18972.0000000000001"

	expectedNumStr := "18972"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_14(t *testing.T) {

	numStr := "-18972.0000000000001"

	expectedNumStr := "-18973"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_15(t *testing.T) {

	numStr := "0.0000000000001"

	expectedNumStr := "0"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_16(t *testing.T) {

	numStr := "-189765342891.0000000000001"
	expectedNumStr := "-189765342892"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}

func TestBigIntFixedDecimal_Floor_17(t *testing.T) {

	numStr := "189765342891.0000000000001"
	expectedNumStr := "189765342891"

	fixDec, err :=
		BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}." +
			"NewNumStr(numStr). numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	floor := fixDec.Floor()

	if expectedNumStr != floor.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedNumStr, floor.GetNumStr())
	}
}
