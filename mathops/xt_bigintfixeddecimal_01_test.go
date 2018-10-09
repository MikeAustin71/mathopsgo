package mathops

import (
	"math/big"
	"testing"
)

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

func TestBigIntFixedDecimal_FormatNumStr_01(t *testing.T) {

	expectedNumStr := "-123.45"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_02(t *testing.T) {

	expectedNumStr := "123.45"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_03(t *testing.T) {

	expectedNumStr := "(123.45)"
	mode := PARENTHESESNEGVALFMTMODE
	originalNum := -12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_04(t *testing.T) {

	expectedNumStr := "-1234.56"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := -123456
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_05(t *testing.T) {

	expectedNumStr := "1234.56"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := 123456
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_06(t *testing.T) {

	expectedNumStr := "(1234.56)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -123456
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_07(t *testing.T) {

	expectedNumStr := "0"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 0
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_08(t *testing.T) {

	expectedNumStr := "0.000"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 0000
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_09(t *testing.T) {

	expectedNumStr := "0.000"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 0000
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_12(t *testing.T) {

	expectedNumStr := "12345"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_13(t *testing.T) {

	expectedNumStr := "(12345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_14(t *testing.T) {

	expectedNumStr := "-12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_15(t *testing.T) {

	expectedNumStr := "0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_16(t *testing.T) {

	expectedNumStr := "0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_17(t *testing.T) {

	expectedNumStr := "-0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_18(t *testing.T) {

	expectedNumStr := "-0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_19(t *testing.T) {

	expectedNumStr := "(0.00012345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_20(t *testing.T) {

	expectedNumStr := "(0.12345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_21(t *testing.T) {

	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_22(t *testing.T) {

	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_23(t *testing.T) {

	expectedNumStr := "0012345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(7)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_24(t *testing.T) {

	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_GetNumStr_01(t *testing.T) {

	expectedNumStr := "123.45"

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.GetNumStr()

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_GetNumStr_02(t *testing.T) {

	expectedNumStr := "-123.45"

	originalNum := -12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.GetNumStr()

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_GetNumStr_03(t *testing.T) {

	expectedNumStr := "0.000"

	originalNum := 0
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.GetNumStr()

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}
