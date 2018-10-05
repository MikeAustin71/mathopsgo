package mathops

import "testing"

func TestBigIntFixedDecimal_RoundToDecPlace_01(t *testing.T) {

	expectedNumStr := "-123.57"
	num := -123567
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_RoundToDecPlace_02(t *testing.T) {

	expectedNumStr := "123.57"
	num := 123567
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_03(t *testing.T) {

	expectedNumStr := "123.567"
	num := 123567
	precision := uint(3)
	roundToDec := uint(3)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_04(t *testing.T) {

	expectedNumStr := "123.5670"
	num := 123567
	precision := uint(3)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_05(t *testing.T) {

	expectedNumStr := "-123.5670"
	num := -123567
	precision := uint(3)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_06(t *testing.T) {

	expectedNumStr := "0.00"
	num := 0
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_07(t *testing.T) {

	expectedNumStr := "654.123"
	num := 654123456
	precision := uint(6)
	roundToDec := uint(3)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_08(t *testing.T) {

	expectedNumStr := "654.1235"
	num := 654123456
	precision := uint(6)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_09(t *testing.T) {

	num := -654123456
	expectedNumStr := "-654.123"
	precision := uint(6)
	roundToDec := uint(3)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_10(t *testing.T) {

	num := -654123456
	expectedNumStr := "-654.1235"
	precision := uint(6)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_11(t *testing.T) {

	num := 654
	expectedNumStr := "654.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_12(t *testing.T) {

	num := 654123
	expectedNumStr := "654.123000000"
	precision := uint(3)
	roundToDec := uint(9)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_13(t *testing.T) {

	num := -654123
	expectedNumStr := "-654.123000000"
	precision := uint(3)
	roundToDec := uint(9)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_14(t *testing.T) {

	num := -654
	expectedNumStr := "-654.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_15(t *testing.T) {

	num := 0
	expectedNumStr := "0.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := BigIntNum{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

