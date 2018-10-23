package mathops

import (
	"math/big"
	"testing"
)

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_01(t *testing.T) {
	radicand := big.NewInt(842567)
	radicandPrecision := big.NewInt(3)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "9.4449895520307989885751143526805"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_02(t *testing.T) {
	radicand := big.NewInt(-842567)
	radicandPrecision := big.NewInt(3)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "-9.4449895520307989885751143526805"

	fdNr :=  FixedDecimalNthRoot{}.New()

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}


func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_03(t *testing.T) {
	radicand := big.NewInt(357)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "7.0939709447507098368744825374903"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}


func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_04(t *testing.T) {
	radicand := big.NewInt(-357)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "-7.0939709447507098368744825374903"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_05(t *testing.T) {
	radicand := big.NewInt(16)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(2)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "4"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_06(t *testing.T) {
	radicand := big.NewInt(531441)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(12)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "3"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_07(t *testing.T) {
	radicand := big.NewInt(78123456)
	radicandPrecision := big.NewInt(6)
	nthRoot := big.NewInt(4)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "2.9730030983116531418548508649862"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_08(t *testing.T) {
	radicand := big.NewInt(-78123456)
	radicandPrecision := big.NewInt(6)
	nthRoot := big.NewInt(4)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)


	fdNr :=  FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error("Expected an error return from even nthRoot and negative radicand. NO ERROR RETURNED! ")
	}

}


func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_09(t *testing.T) {
	radicand := big.NewInt(-78123456)
	radicandPrecision := big.NewInt(6)
	nthRoot := big.NewInt(5)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(30)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "-2.390871799107522442017212407745"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_10(t *testing.T) {
	radicand := big.NewInt(1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(5)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(30)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "1"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_11(t *testing.T) {
	radicand := big.NewInt(1)
	radicandPrecision := big.NewInt(5)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.02154434690031883721759293566519"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_12(t *testing.T) {
	radicand := big.NewInt(-1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "-1"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_13(t *testing.T) {
	radicand := big.NewInt(-1)
	radicandPrecision := big.NewInt(5)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "-0.02154434690031883721759293566519"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_14(t *testing.T) {
	radicand := big.NewInt(0)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(3)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

	//                             1         2         3
	//                    1234567890123456789012345678901234567
	expectedResult := "0"

	fdNr :=  FixedDecimalNthRoot{}

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
	}

	resultBiNum, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := resultBiNum.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_15(t *testing.T) {
	radicand := big.NewInt(37)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(0)
	nthRootPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

	fdNr :=  FixedDecimalNthRoot{}

	_, _, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err == nil {
		t.Error ("Expected error return from nthRoot==0. NO ERROR RETURNED! ")
	}

}
