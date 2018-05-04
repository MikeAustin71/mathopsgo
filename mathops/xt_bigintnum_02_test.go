package mathops

import "testing"

func TestBigIntNum_RoundToDecPlace_01(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.57"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_02(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.57"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_03(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.567"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_04(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.5670"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_05(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.5670"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_06(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0.00"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_07(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.123"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_08(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.1235"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_09(t *testing.T) {

	nStr := "-654.123456"
	expectedNumStr := "-654.123"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_10(t *testing.T) {

	nStr := "-654.123456"
	expectedNumStr := "-654.1235"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_11(t *testing.T) {

	nStr := "654"
	expectedNumStr := "654.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_12(t *testing.T) {

	nStr := "654.123"
	expectedNumStr := "654.123000000"
	roundToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_13(t *testing.T) {

	nStr := "-654.123"
	expectedNumStr := "-654.123000000"
	roundToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_14(t *testing.T) {

	nStr := "-654"
	expectedNumStr := "-654.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}


func TestBigIntNum_RoundToDecPlace_15(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_SetPrecision_01(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654.123"
	newPrecision := uint(3)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_02(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654.1235"
	newPrecision := uint(4)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_03(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654"
	newPrecision := uint(0)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_04(t *testing.T) {

	num1Str := "-654.123456"
	expectedNumStr := "-654.123"
	newPrecision := uint(3)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_05(t *testing.T) {

	num1Str := "-654.123456"
	expectedNumStr := "-654.1235"
	newPrecision := uint(4)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_06(t *testing.T) {

	num1Str := "654"
	expectedNumStr := "654.000"
	newPrecision := uint(3)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_07(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654.123456000"
	newPrecision := uint(9)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_08(t *testing.T) {

	num1Str := "-654"
	expectedNumStr := "-654.000000000"
	newPrecision := uint(9)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_09(t *testing.T) {

	num1Str := "-654.123456"
	expectedNumStr := "-654.123456000"
	newPrecision := uint(9)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_10(t *testing.T) {

	num1Str := "0"
	expectedNumStr := "0.0000"
	newPrecision := uint(4)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_11(t *testing.T) {

	num1Str := "0.0000"
	expectedNumStr := "0"
	newPrecision := uint(0)

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum1='%v'" +
			"precision='%v' sign='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr(), bNum1.GetPrecisionUint(), bNum1.GetSign())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_TruncToDecPlace_01(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.56"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_02(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.56"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}


func TestBigIntNum_TruncToDecPlace_03(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.567"
	truncToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_04(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.5670"
	truncToPlace := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToPlace)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_05(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.5670"
	truncToPlace := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToPlace)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_06(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0.00"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_07(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.123"
	truncToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_08(t *testing.T) {

	nStr := "654.123456789"
	expectedNumStr := "654.1234"
	truncToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_09(t *testing.T) {

	nStr := "654.123456789"
	expectedNumStr := "654"
	truncToDec := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_10(t *testing.T) {

	nStr := "654"
	expectedNumStr := "654.00000"
	truncToDec := uint(5)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_11(t *testing.T) {

	nStr := "654.123"
	expectedNumStr := "654.123000000"
	truncToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_12(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0.000000"
	truncToDec := uint(6)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_13(t *testing.T) {

	nStr := "0.000000"
	expectedNumStr := "0"
	truncToDec := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_14(t *testing.T) {

	nStr := "654.123456789015"
	expectedNumStr := "654.12345678901"
	truncToDec := uint(11)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

