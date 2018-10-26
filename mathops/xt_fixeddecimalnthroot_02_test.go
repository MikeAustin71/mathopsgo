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
		t.Error("Error: Expected error return from zero nthroot. Instead, NO ERROR RETURNED! ")
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
		t.Error("Error: Expected error return from zero nthroot. Instead, NO ERROR RETURNED! " )
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
		t.Error("Error: Expected error return from negative nthroot. Instead, NO ERROR RETURNED! ")
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
		t.Error("Error: Expected error return from integer nthroot. Instead, NO ERROR RETURNED! " )
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_16(t *testing.T) {
	radicand := big.NewInt(15)
	radicandPrecision := big.NewInt(-1)
	nthRoot := big.NewInt(12)
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
		t.Error("Error: Expected error return from negative radicand precision. Instead, NO ERROR RETURNED! " )
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_17(t *testing.T) {
	radicand := big.NewInt(15)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(12)
	nthRootPrecision := big.NewInt(-2)
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
		t.Error("Error: Expected error return from negative nthRoot precision. Instead, NO ERROR RETURNED! " )
	}

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_01(t *testing.T) {

	radicand := big.NewInt(38432)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-372)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.05855387604340335560518923124485"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_02(t *testing.T) {

	radicand := big.NewInt(-38432)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-372)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "-0.05855387604340335560518923124485"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_03(t *testing.T) {

	radicand := big.NewInt(945384)
	radicandPrecision := big.NewInt(5)
	nthRoot := big.NewInt(-5723)
	nthRootPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.6753494112877003591875450895632"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_04(t *testing.T) {

	radicand := big.NewInt(-945384)
	radicandPrecision := big.NewInt(5)
	nthRoot := big.NewInt(-5723)
	nthRootPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.6753494112877003591875450895632"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_05(t *testing.T) {

	radicand := big.NewInt(64)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-22)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.15101118055055591416115527630953"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_06(t *testing.T) {

	radicand := big.NewInt(-6475)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-382)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.10052943765859122324489541725654"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_07(t *testing.T) {

	radicand := big.NewInt(-5291)
	radicandPrecision := big.NewInt(2)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.36052149780869144809980868847077"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_08(t *testing.T) {

	radicand := big.NewInt(5291)
	radicandPrecision := big.NewInt(2)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.36052149780869144809980868847077"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_09(t *testing.T) {

	radicand := big.NewInt(-5291)
	radicandPrecision := big.NewInt(2)
	nthRoot := big.NewInt(-2)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

}


func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_10(t *testing.T) {

	radicand := big.NewInt(0)
	radicandPrecision := big.NewInt(3)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_11(t *testing.T) {

	radicand := big.NewInt(1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_12(t *testing.T) {

	radicand := big.NewInt(-1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1"

	fdNr := FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) " +
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

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_13(t *testing.T) {

	radicand := big.NewInt(58971)
	radicandPrecision := big.NewInt(-1)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)


	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return due to negative radicand precision. NO ERROR RETURNED!")
	}

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_14(t *testing.T) {

	radicand := big.NewInt(58971)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(-389)
	nthRootPrecision := big.NewInt(-2)
	maxPrecision := big.NewInt(32)


	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return due to negative nthRoot precision. NO ERROR RETURNED!")
	}

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_15(t *testing.T) {

	radicand := big.NewInt(58971)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(0)
	nthRootPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(32)


	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return due to zero nthRoot. NO ERROR RETURNED!")
	}

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_16(t *testing.T) {

	radicand := big.NewInt(58971)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(52)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(32)


	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return due to positive nthRoot. NO ERROR RETURNED!")
	}

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_17(t *testing.T) {

	radicand := big.NewInt(58971)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(52)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)


	fdNr := FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return due to integer nthRoot. NO ERROR RETURNED!")
	}

}
