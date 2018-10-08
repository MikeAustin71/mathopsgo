package mathops

import "testing"

func TestBigIntNum_TrimTrailingFracZeros_01(t *testing.T) {

	nStr := "-123.000"
	expectedNumStr := "-123"

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TrimTrailingFracZeros()

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntNum_TrimTrailingFracZeros_02(t *testing.T) {

	nStr := "123.000"
	expectedNumStr := "123"

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TrimTrailingFracZeros()

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntNum_TrimTrailingFracZeros_03(t *testing.T) {

	nStr := "123.0090"
	expectedNumStr := "123.009"

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TrimTrailingFracZeros()

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntNum_TrimTrailingFracZeros_04(t *testing.T) {

	nStr := "-123.0090"
	expectedNumStr := "-123.009"

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TrimTrailingFracZeros()

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntNum_TrimTrailingFracZeros_05(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0"

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TrimTrailingFracZeros()

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntNum_TruncToDecPlace_01(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.56"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_02(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.56"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_03(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.567"
	truncToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_04(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.5670"
	truncToPlace := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToPlace)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_05(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.5670"
	truncToPlace := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToPlace)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_06(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0.00"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_07(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.123"
	truncToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_08(t *testing.T) {

	nStr := "654.123456789"
	expectedNumStr := "654.1234"
	truncToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_09(t *testing.T) {

	nStr := "654.123456789"
	expectedNumStr := "654"
	truncToDec := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_10(t *testing.T) {

	nStr := "654"
	expectedNumStr := "654.00000"
	truncToDec := uint(5)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_11(t *testing.T) {

	nStr := "654.123"
	expectedNumStr := "654.123000000"
	truncToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_12(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0.000000"
	truncToDec := uint(6)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_13(t *testing.T) {

	nStr := "0.000000"
	expectedNumStr := "0"
	truncToDec := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_TruncToDecPlace_14(t *testing.T) {

	nStr := "654.123456789015"
	expectedNumStr := "654.12345678901"
	truncToDec := uint(11)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

