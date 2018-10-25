package mathops

import (
	"math/big"
	"testing"
)

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_01(t *testing.T) {
	radicand := big.NewInt(4567)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(666)
	nthRootPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(26)

	//                                 1         2         3
	//                        1234567890123456789012345678901234567
	expectedResult := "312565.80127327376619628998765762"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_02(t *testing.T) {
	radicand := big.NewInt(45123)
	radicandPrecision := big.NewInt(3)
	nthRoot := big.NewInt(25)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "4.5894346095427589142057717684748"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_03(t *testing.T) {
	radicand := big.NewInt(-45123)
	radicandPrecision := big.NewInt(3)
	nthRoot := big.NewInt(25)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "4.5894346095427589142057717684748"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_04(t *testing.T) {
	radicand := big.NewInt(659)
	radicandPrecision := big.NewInt(1)
	nthRoot := big.NewInt(251)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1.1815865924660052902023955150749"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_05(t *testing.T) {
	radicand := big.NewInt(679)
	radicandPrecision := big.NewInt(1)
	nthRoot := big.NewInt(353)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "3.3032639583686665394923710881736"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_06(t *testing.T) {
	radicand := big.NewInt(-679)
	radicandPrecision := big.NewInt(1)
	nthRoot := big.NewInt(353)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "3.3032639583686665394923710881736"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_07(t *testing.T) {
	radicand := big.NewInt(9645321)
	radicandPrecision := big.NewInt(1)
	nthRoot := big.NewInt(912)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(30)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1.163101209960992659162495181682"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_08(t *testing.T) {
	radicand := big.NewInt(-9645321)
	radicandPrecision := big.NewInt(1)
	nthRoot := big.NewInt(9124)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "-1.1630241704961631471493809157491"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_09(t *testing.T) {
	radicand := big.NewInt(0)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(52)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "0"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_10(t *testing.T) {
	radicand := big.NewInt(1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(52)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "1"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_11(t *testing.T) {
	radicand := big.NewInt(-1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(52)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "1"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != resultBiNum.GetNumStr() {
		t.Errorf("Error: expected result='%v'. Instead, result='%v'",
			expectedResult, resultBiNum.GetNumStr())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_12(t *testing.T) {
	radicand := big.NewInt(15)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(0)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected erro return from zero nthroot. Instead, NO ERROR RETURNED! " +
			"Error='%v' ", err.Error())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_13(t *testing.T) {
	radicand := big.NewInt(15)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(0)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected erro return from zero nthroot. Instead, NO ERROR RETURNED! " +
			"Error='%v' ", err.Error())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_14(t *testing.T) {
	radicand := big.NewInt(15)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-67)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected erro return from negative nthroot. Instead, NO ERROR RETURNED! " +
			"Error='%v' ", err.Error())
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_15(t *testing.T) {
	radicand := big.NewInt(15)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(12)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected erro return from integer nthroot. Instead, NO ERROR RETURNED! " +
			"Error='%v' ", err.Error())
	}

}
