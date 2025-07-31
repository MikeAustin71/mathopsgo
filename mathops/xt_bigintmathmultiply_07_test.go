package mathops

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBigIntMathMultiply_MultiplyNumStrSeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrSeries_01"

	var err error

	// multiplier = 2
	multiplierStr := "2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	expectedBigINumStr := "128"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var iaMultiplicand IntAry

	var iaMultiplicandNumStr string

	var expectedBigINum, resultBigINum BigIntNum

	for i := 0; i < lenArray; i++ {

		iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
				"multiplicandStrs[%d], expectedNumSeps)\n"+
				"multiplicandStrs[%v]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaMultiplicand.IsValid(ePrefix + "Validating iaMultiplicand")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaMultiplicand.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)\n"+
				"iaResult= '%v'\n"+
				"iaMultiplicand= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResult.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err = iaResult.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

	} // end of For loop

	expectedBigINum, err = new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:"+
			"expectedBigINum, err = new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
			"    expectedNumSeps, expectedNumSeps, multiplierStr, multiplicandStrs[0-5])\n"+
			"expectedNumSeps= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), multiplierStr, err.Error())
		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var isEqualToResult bool

	isEqualToResult, err = expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEqualToResult, err = expectedBigINum.Equal(resultBigINum)\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if !isEqualToResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isEqualToResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	var resultBigInt *big.Int

	resultBigInt, err = resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err = resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, err = expectedBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	var resultSignValue int

	resultSignValue, err = resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultSignValue, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSignValue\n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrSeries_02"

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedBigINumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var iaMultiplicand IntAry

	var iaMultiplicandNumStr string

	var expectedBigINum, resultBigINum BigIntNum

	for i := 0; i < lenArray; i++ {

		iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
				"multiplicandStrs[%d], expectedNumSeps)\n"+
				"multiplicandStrs[%v]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaMultiplicand.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)\n"+
				"iaResult= '%v'\n"+
				"iaMultiplicand= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResult.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err = iaResult.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err = new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:"+
			"expectedBigINum, err = new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigIntNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigIntNumStr\n"+
			"Expected expectedBigIntNumStr = '%v'\n"+
			"  Actual expectedBigIntNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigIntNumStr)

		return
	}

	resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
			"    expectedNumSeps, expectedNumSeps, multiplierStr, multiplicandStrs[0-5])\n"+
			"expectedNumSeps= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), multiplierStr, err.Error())
		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var isEqualToResult bool

	isEqualToResult, err = expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEqualToResult, err = expectedBigINum.Equal(resultBigINum)\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if !isEqualToResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isEqualToResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	var resultBigInt *big.Int

	resultBigInt, err = resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err = resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, err = expectedBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	var resultSignValue int

	resultSignValue, err = resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultSignValue, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSignValue\n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrSeries_03"

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"

	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedBigINumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var iaMultiplicand IntAry

	var iaMultiplicandNumStr string

	var expectedBigINum, resultBigINum BigIntNum

	for i := 0; i < lenArray; i++ {

		iaMultiplicand, err = new(IntAry).
			NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
				"multiplicandStrs[%d], expectedNumSeps)\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaMultiplicand.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)\n"+
				"iaResult= '%v'\n"+
				"iaMultiplicand= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResult.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err = iaResult.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err = new(BigIntNum).
		NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:"+
			"expectedBigINum, err = new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigIntNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigIntNumStr\n"+
			"Expected expectedBigIntNumStr = '%v'\n"+
			"  Actual expectedBigIntNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigIntNumStr)

		return
	}

	resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
			"    expectedNumSeps, expectedNumSeps, multiplierStr, multiplicandStrs[0-5])\n"+
			"expectedNumSeps= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), multiplierStr, err.Error())
		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var isEqualToResult bool

	isEqualToResult, err = expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEqualToResult, err = expectedBigINum.Equal(resultBigINum)\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if !isEqualToResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isEqualToResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	var resultBigInt *big.Int

	resultBigInt, err = resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err = resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, err = expectedBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	var resultSignValue int

	resultSignValue, err = resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultSignValue, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSignValue\n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrSeries_04"

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"

	// multiplicandStrs
	multiplicandStrs := []string{
		"1.879",
		"3.824",
		"21.756",
		"2.1234567",
		"6",
		"2",
	}

	// product = -20408.5138429311978576052224
	expectedBigINumStr := "-20408.5138429311978576052224"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).
		NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = iaResult.IsValid(ePrefix + "Validating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var iaMultiplicand IntAry

	var iaMultiplicandNumStr string

	var expectedBigINum, resultBigINum BigIntNum

	for i := 0; i < lenArray; i++ {

		iaMultiplicand, err = new(IntAry).
			NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
				"multiplicandStrs[%d], expectedNumSeps)\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaMultiplicand.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)\n"+
				"iaResult= '%v'\n"+
				"iaMultiplicand= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResult.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err = iaResult.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err = new(BigIntNum).
		NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:"+
			"expectedBigINum, err = new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigIntNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigIntNumStr\n"+
			"Expected expectedBigIntNumStr = '%v'\n"+
			"  Actual expectedBigIntNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigIntNumStr)

		return
	}

	resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
			"    expectedNumSeps, expectedNumSeps, multiplierStr, multiplicandStrs[0-5])\n"+
			"expectedNumSeps= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), multiplierStr, err.Error())
		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var isEqualToResult bool

	isEqualToResult, err = expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEqualToResult, err = expectedBigINum.Equal(resultBigINum)\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if !isEqualToResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isEqualToResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	var resultBigInt *big.Int

	resultBigInt, err = resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err = resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, err = expectedBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	var resultSignValue int

	resultSignValue, err = resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultSignValue, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSignValue\n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrSeries_05"

	var err error

	// multiplier = -5,123456
	multiplierStr := "-5,123456"

	// multiplicandStrs
	multiplicandStrs := []string{
		"1,879",
		"3,824",
		"21,756",
		"2,1234567",
		"6",
		"2",
	}

	// product = -20408,5138429311978576052224
	expectedBigINumStr := "-20408,5138429311978576052224"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'
	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var iaMultiplicand IntAry

	var iaMultiplicandNumStr string

	var expectedBigINum BigIntNum

	for i := 0; i < lenArray; i++ {

		iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
				"multiplicandStrs[%d], expectedNumSeps)\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaMultiplicand.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)\n"+
				"iaResult= '%v'\n"+
				"iaMultiplicand= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResult.IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err = iaResult.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err = new(BigIntNum).NewNumStrWithNumSeps(
		expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:"+
			"expectedBigINum, err = new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigIntNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigIntNumStr\n"+
			"Expected expectedBigIntNumStr = '%v'\n"+
			"  Actual expectedBigIntNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigIntNumStr)

		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err = new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
			"    expectedNumSeps, expectedNumSeps, multiplierStr, multiplicandStrs[0-5])\n"+
			"expectedNumSeps= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), multiplierStr, err.Error())
		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var isEqualToResult bool

	isEqualToResult, err = expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEqualToResult, err = expectedBigINum.Equal(resultBigINum)\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if !isEqualToResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isEqualToResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	var resultBigInt *big.Int

	resultBigInt, err = resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err = resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, err = expectedBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	var resultSignValue int

	resultSignValue, err = resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultSignValue, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSignValue\n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_01"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedBigINumStr := "2875.94572"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = iaMultiplier.IsValid(ePrefix + "Validating iaMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = iaMultiplicand.IsValid(ePrefix + "Validating iaMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplicand.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iaMultiplierNumStr,
			iaMultiplicandNumStr,
			err.Error())

		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid(ePrefix + "Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	resultBigInt, err := resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedNumStrDtoBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultBigINumSignVal, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultBigINumSignVal, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedSignValue != resultBigINumSignVal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != resultBigINumSignVal\n"+
			"Expected resultBigINumSignVal = '%v'\n"+
			"  Actual resultBigINumSignVal = '%v'\n\n",
			ePrefix, expectedSignValue, resultBigINumSignVal)

		return
	}

	if iaResultNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultBigINumStr)

		return
	}

	if expectedNumStrDtoNumStr != iaResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoNumStr != iaResultNumStr\n"+
			"Expected iaResultNumStr = '%v'\n"+
			"  Actual iaResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, iaResultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_02"

	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = iaMultiplier.IsValid(ePrefix + "Validating iaMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = iaMultiplicand.IsValid(ePrefix + "Validating iaMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplicand.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iaMultiplierNumStr,
			iaMultiplicandNumStr,
			err.Error())

		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid(ePrefix + "Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultBigINumStr)

		return
	}

	resultBigInt, err := resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedNumStrDtoBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultBigINumSignVal, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultBigINumSignVal, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedSignValue != resultBigINumSignVal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != resultBigINumSignVal\n"+
			"Expected resultBigINumSignVal = '%v'\n"+
			"  Actual resultBigINumSignVal = '%v'\n\n",
			ePrefix, expectedSignValue, resultBigINumSignVal)

		return
	}

	if iaResultNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultBigINumStr)

		return
	}

	if expectedNumStrDtoNumStr != iaResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoNumStr != iaResultNumStr\n"+
			"Expected iaResultNumStr = '%v'\n"+
			"  Actual iaResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, iaResultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_03"

	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	expectedSignValue := -1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = iaMultiplier.IsValid(ePrefix + "Validating iaMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = iaMultiplicand.IsValid(ePrefix + "Validating iaMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplicand.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iaMultiplierNumStr,
			iaMultiplicandNumStr,
			err.Error())

		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid(ePrefix + "Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultBigINumStr)

		return
	}

	resultBigInt, err := resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedNumStrDtoBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultBigINumSignVal, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultBigINumSignVal, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedSignValue != resultBigINumSignVal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != resultBigINumSignVal\n"+
			"Expected resultBigINumSignVal = '%v'\n"+
			"  Actual resultBigINumSignVal = '%v'\n\n",
			ePrefix, expectedSignValue, resultBigINumSignVal)

		return
	}

	if iaResultNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultBigINumStr)

		return
	}

	if expectedNumStrDtoNumStr != iaResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoNumStr != iaResultNumStr\n"+
			"Expected iaResultNumStr = '%v'\n"+
			"  Actual iaResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, iaResultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_04"

	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = iaMultiplier.IsValid(ePrefix + "Validating iaMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = iaMultiplicand.IsValid(ePrefix + "Validating iaMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplicand.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iaMultiplierNumStr,
			iaMultiplicandNumStr,
			err.Error())

		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid(ePrefix + "Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultBigINumStr)

		return
	}

	resultBigInt, err := resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	if expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedNumStrDtoBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultBigINumSignVal, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultBigINumSignVal, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedSignValue != resultBigINumSignVal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != resultBigINumSignVal\n"+
			"Expected resultBigINumSignVal = '%v'\n"+
			"  Actual resultBigINumSignVal = '%v'\n\n",
			ePrefix, expectedSignValue, resultBigINumSignVal)

		return
	}

	if iaResultNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultBigINumStr)

		return
	}

	if expectedNumStrDtoNumStr != iaResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoNumStr != iaResultNumStr\n"+
			"Expected iaResultNumStr = '%v'\n"+
			"  Actual iaResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, iaResultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_05"

	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0
	expectedNumStr := "0"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = iaMultiplier.IsValid(ePrefix + "Validating iaMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, err.Error())
		return
	}

	err = iaMultiplicand.IsValid(ePrefix + "Validating iaMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplicand.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"  iaMultiplier.Multiply()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMultiplier.Multiply(\n"+
			"  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
			"iaMultiplier= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iaMultiplierNumStr,
			iaMultiplicandNumStr,
			err.Error())

		return
	}

	err = expectedNumStrDto.IsValid(ePrefix + "Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultBigINumStr)

		return
	}

	resultBigInt, err := resultBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := resultBigINum.GetBigInt()\n"+
			"Error='%v'", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()\n"+
			"expectedNumStrDto= '%v'\n"+
			"Error='%v'", ePrefix, expectedNumStrDtoNumStr, err.Error())
		return
	}

	if expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedNumStrDtoBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	iaResultBigInt, err := iaResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultBigInt, err := iaResult.GetBigInt()\n"+
			"iaResult= '%v'\n"+
			"Error='%v'", ePrefix, iaResultNumStr, err.Error())
		return
	}

	if expectedNumStrDtoBigInt.Cmp(iaResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStrDtoBigInt.Cmp(iaResultBigInt) != 0\n"+
			"Expected iaResultBigInt = '%v'\n"+
			"  Actual iaResultBigInt = '%v'\n\n",
			ePrefix, expectedNumStrDtoBigInt.Text(10), iaResultBigInt.Text(10))

		return
	}

	resultBigINumSignVal, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"resultBigINumSignVal, err = reresultBigINumsult.GetSign()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedSignValue != resultBigINumSignVal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedSignValue != resultBigINumSignVal\n"+
			"Expected resultBigINumSignVal = '%v'\n"+
			"  Actual resultBigINumSignVal = '%v'\n\n",
			ePrefix, expectedSignValue, resultBigINumSignVal)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_06"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875,94572
	expectedNumStr := "2875,94572"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto, expectedNumSeps)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps NOT Equal to resultBigINumSeps\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_07(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDto_06"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875,94572
	expectedNumStr := "2875,94572"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplicandNumStrDto.IsValid(ePrefix + "Validating multiplicandNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandNumStrDtoNumStr, err := multiplicandNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto, expectedNumSeps)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"multiplicandNumStrDto= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			multiplicandNumStrDtoNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != resultBigINumStr\n"+
			"Expected resultBigINumStr = '%v'\n"+
			"  Actual resultBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultBigINumStr)

		return
	}

	resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps NOT Equal to resultBigINumSeps\n"+
			"Expected resultBigINumSeps = '%v'\n"+
			"  Actual resultBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoArray_01"

	var err error

	// multiplier = 2
	multiplierStr := "2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	expectedBigINumStr := "128"

	expectedBigINumSign := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var ia IntAry

	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(multiplicandStrs[%d]))\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], err.Error())

			return
		}

		err = nDtoArray[i].IsValid(ePrefix + fmt.Sprintf("Validating multiplierIntAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				fmt.Sprintf("err = nDtoArray[%d].IsValid(ePrefix)\n", i)+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		ia, err = new(IntAry).NewNumStrDto(nDtoArray[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStrDto(nDtoArray[%d])\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = ia.IsValid(fmt.Sprintf("Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Cycle Number= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaNumStr, i, iaResultNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + fmt.Sprintf(" Validating iaResult[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

	} // End of Loop

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)\n"+
			"multiplierIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigIntNumberStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
			"Expected resultActualBigInt = '%v'\n"+
			"  Actual resultActualBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

		return
	}

	resultSign, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSign, err := result.GetSign()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSign\n"+
			"Expected resultSign = '%v'\n"+
			"  Actual resultSign = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoArray_02"

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedBigINumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var ia IntAry

	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(multiplicandStrs[%d]))\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], err.Error())

			return
		}

		err = nDtoArray[i].IsValid(ePrefix + fmt.Sprintf("Validating multiplierIntAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				fmt.Sprintf("err = nDtoArray[%d].IsValid(ePrefix)\n", i)+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = ia.IsValid(fmt.Sprintf("Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Cycle Number= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaNumStr, i, iaResultNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + fmt.Sprintf(" Validating iaResult[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

	} // End of Loop

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigIntNumberStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
			"Expected resultActualBigInt = '%v'\n"+
			"  Actual resultActualBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

		return
	}

	resultSign, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSign, err := result.GetSign()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSign\n"+
			"Expected resultSign = '%v'\n"+
			"  Actual resultSign = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoArray_03"

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"

	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedBigINumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var ia IntAry

	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(multiplicandStrs[%d]))\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], err.Error())

			return
		}

		err = nDtoArray[i].IsValid(ePrefix + fmt.Sprintf("Validating multiplierIntAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				fmt.Sprintf("err = nDtoArray[%d].IsValid(ePrefix)\n", i)+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = ia.IsValid(fmt.Sprintf("Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Cycle Number= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaNumStr, i, iaResultNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + fmt.Sprintf(" Validating iaResult[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierIntAry= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigIntNumberStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
			"Expected resultActualBigInt = '%v'\n"+
			"  Actual resultActualBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

		return
	}

	resultSign, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSign, err := result.GetSign()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSign\n"+
			"Expected resultSign = '%v'\n"+
			"  Actual resultSign = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoArray_04"

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"

	// multiplicandStrs
	multiplicandStrs := []string{
		"1.879",
		"3.824",
		"21.756",
		"2.1234567",
		"6",
		"2",
	}

	// product = -20408.5138429311978576052224
	expectedBigINumStr := "-20408.5138429311978576052224"

	expectedBigINumSign := -1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
			"multiplierStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	var ia IntAry

	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(multiplicandStrs[%d])\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], err.Error())

			return
		}

		err = nDtoArray[i].IsValid(ePrefix + fmt.Sprintf("Validating multiplierIntAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				fmt.Sprintf("err = nDtoArray[%d].IsValid(ePrefix)\n", i)+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = ia.IsValid(fmt.Sprintf("Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Cycle Number= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, iaNumStr, i, iaResultNumStr, err.Error())
			return
		}

		err = iaResult.IsValid(ePrefix + fmt.Sprintf(" Validating iaResult[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].IsValid(ePrefix)\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid(ePrefix + "Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigIntNumberStr, resultNumStr)

		return
	}

	resultActualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultActualBigInt, err := result.GetBigInt()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
			"Expected resultActualBigInt = '%v'\n"+
			"  Actual resultActualBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

		return
	}

	resultSign, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSign, err := result.GetSign()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != resultSign\n"+
			"Expected resultSign = '%v'\n"+
			"  Actual resultSign = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSign)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because iaResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoArray_05"

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664,26376575446779648
	expectedNumStr := "11995826664,26376575446779648"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierNumStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDto, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
		return
	}

	err = multiplierNumStrDto.IsValid(ePrefix + "Validating multiplierNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierNumStrDto.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierNumStrDtoNumStr, err := multiplierNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(multiplicandStrs)

	nDtoArray := make([]NumStrDto, lenArray)

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStrWithNumSeps(multiplicandStrs[i], &usaNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(multiplicandStrs[%d], &usaNumSeps)\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"usaNumSeps= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())

			return
		}

	} // End of for loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray[...], expectedNumSeps)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr NOT EQUAL to resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps NOT EQUAL to resultNumSeps\n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}
