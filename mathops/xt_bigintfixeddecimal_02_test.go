package mathops

import "testing"

func TestBigIntFixedDecimal_MultiplyByTenToPwr_01(t *testing.T) {

	expectedNumStr := "105.6752"
	num := 1056752
	precision := uint(4)
	exponent := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_02(t *testing.T) {

	expectedNumStr := "-105.6752"
	num := -1056752
	precision := uint(4)
	exponent := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_03(t *testing.T) {

	expectedNumStr := "1056.752"
	num := 1056752
	precision := uint(4)
	exponent := uint(1)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_04(t *testing.T) {

	expectedNumStr := "-1056.752"
	num := -1056752
	precision := uint(4)
	exponent := uint(1)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_05(t *testing.T) {

	expectedNumStr := "10567.52"
	num := 1056752
	precision := uint(4)
	exponent := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_06(t *testing.T) {

	expectedNumStr := "-10567.52"
	num := -1056752
	precision := uint(4)
	exponent := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_07(t *testing.T) {

	expectedNumStr := "10567520000"
	num := 1056752
	precision := uint(4)
	exponent := uint(8)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_08(t *testing.T) {

	expectedNumStr := "0"
	num := 0
	precision := uint(4)
	exponent := uint(5)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_RoundToDecPlace_01(t *testing.T) {

	expectedNumStr := "-123.57"
	num := -123567
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

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

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_01(t *testing.T) {

	num := 456123000
	precision := uint(6)
	expectedNumStr := "456.123"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_02(t *testing.T) {

	num := -456123000
	precision := uint(6)
	expectedNumStr := "-456.123"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_03(t *testing.T) {

	num := 0
	precision := uint(3)
	expectedNumStr := "0"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_04(t *testing.T) {

	num := 70
	precision := uint(1)
	expectedNumStr := "7"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_05(t *testing.T) {

	num := uint64(7000000000000000000)
	precision := uint(17)
	expectedNumStr := "70"

	fixedDec := BigIntFixedDecimal{}.NewUInt64(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TruncToDecPlace_01(t *testing.T) {

	num := -123567
	precision := uint(3)
	expectedNumStr := "-123.56"
	truncToDec := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_02(t *testing.T) {
	num := 123567
	precision := uint(3)
	expectedNumStr := "123.56"
	truncToDec := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_03(t *testing.T) {

	num := 123567
	precision := uint(3)
	expectedNumStr := "123.567"
	truncToDec := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_04(t *testing.T) {

	num := 123567
	precision := uint(3)
	expectedNumStr := "123.5670"
	truncToDec := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_05(t *testing.T) {

	num := -123567
	precision := uint(3)
	expectedNumStr := "-123.5670"
	truncToDec := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_06(t *testing.T) {
	num := 0
	precision := uint(3)
	expectedNumStr := "0.00"
	truncToDec := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_07(t *testing.T) {

	num := 654123456
	precision := uint(6)
	expectedNumStr := "654.123"
	truncToDec := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_08(t *testing.T) {

	num := 654123456789
	precision := uint(9)
	expectedNumStr := "654.1234"
	truncToDec := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_09(t *testing.T) {

	num := 654123456789
	precision := uint(9)
	expectedNumStr := "654"
	truncToDec := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_10(t *testing.T) {

	num := 654
	precision := uint(0)
	expectedNumStr := "654.00000"
	truncToDec := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_11(t *testing.T) {

	num := 654123
	precision := uint(3)
	expectedNumStr := "654.123000000"
	truncToDec := uint(9)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_12(t *testing.T) {

	num := 0
	precision := uint(0)
	expectedNumStr := "0.000000"
	truncToDec := uint(6)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_13(t *testing.T) {

	num := 0
	precision := uint(6)
	expectedNumStr := "0"
	truncToDec := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_14(t *testing.T) {
	num := 654123456789015
	precision := uint(12)
	expectedNumStr := "654.12345678901"
	truncToDec := uint(11)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}
