package mathops

import (
	"fmt"
	"testing"
)

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_01"

	var err error

	// multiplier = 2
	multiplierStr := "2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"2",
		"4",
		"6",
		"8",
		"10",
		"12",
	}

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

	} // End of Loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because  expectedNumStrs[%d] != resultNumStr\n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, j, expectedNumStrs[j], resultNumStr)

			return
		}

	} // End of Loop

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_02"

	var err error

	// multiplier = 8
	multiplierStr := "8"

	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"800.8",
		"-208",
		"31.392",
		"64",
		"42376.984",
		"-39.168",
	}

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

	} // End of Loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because  expectedNumStrs[%d] != resultNumStr\n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, j, expectedNumStrs[j], resultNumStr)

			return
		}

	} // End of Loop

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_03"

	var err error

	// multiplier = -31.2
	multiplierStr := "-31.2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"-3123.12",
		"811.2",
		"-122.4288",
		"-249.6",
		"-165270.2376",
		"152.7552",
	}

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

	} // End of Loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because  expectedNumStrs[%d] != resultNumStr\n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, j, expectedNumStrs[j], resultNumStr)

			return
		}

	} // End of Loop

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_04"

	var err error

	// multiplier = 283
	multiplierStr := "283"

	// multiplicandStrs
	multiplicandStrs := []string{
		"0",
		"-26",
		"0",
		"8",
		"5297.123",
		"0",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"0",
		"-7358",
		"0",
		"2264",
		"1499085.809",
		"0",
	}

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

	} // End of Loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because  expectedNumStrs[%d] != resultNumStr\n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, j, expectedNumStrs[j], resultNumStr)

			return
		}

	} // End of Loop

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_05"

	var err error

	// multiplier = 0
	multiplierStr := "0"

	// multiplicandStrs
	multiplicandStrs := []string{
		"5",
		"-26",
		"9",
		"8",
		"5297.123",
		"37",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}

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

	} // End of Loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray[...])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			multiplierNumStrDtoNumStr,
			err.Error())

		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because  expectedNumStrs[%d] != resultNumStr\n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, j, expectedNumStrs[j], resultNumStr)

			return
		}

	} // End of Loop

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_06"

	var err error

	// multiplier = 8
	multiplierStr := "8"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"800,8",
		"-208",
		"31,392",
		"64",
		"42376,984",
		"-39,168",
	}

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

	} // End of Loop

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray[...], expectedNumSeps)\n"+
			"multiplierNumStrDto= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierNumStrDtoNumStr, expectedNumSeps.String(), err.Error())

		return
	}

	var resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for j := 0; j < lenArray; j++ {

		resultNumSeps, err = result[j].GetNumericSeparatorsDto()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
				"Error= '%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if !expectedNumSeps.Equal(resultNumSeps) {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedNumSeps NOT EQUAL to resultNumSeps\n"+
				"Expected resultNumSeps[%d] = '%v'\n"+
				"  Actual resultNumSeps[%d] = '%v'\n\n",
				ePrefix, j, expectedNumSeps.String(), j, resultNumSeps.String())

			return
		}

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because  expectedNumStrs[%d] != resultNumStr[%d]\n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, j, j, expectedNumStrs[j], resultNumStr)

			return
		}

	} // End of Loop

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoSeries_01"

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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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

	err = iaResult.IsValid(ePrefix + "Validating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	var nDtoArrayNumStr, iaNumStr string

	var ia IntAry

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

		nDtoArrayNumStr, err = nDtoArray[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArrayNumStr, err =  nDtoArray[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"nDtoArray[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, nDtoArrayNumStr, err.Error())
			return
		}

		err = ia.IsValid(ePrefix + fmt.Sprintf(" Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia.IsValid(ePrefix)\n"+
				"Cycle No= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				i,
				iaNumStr,
				iaResultNumStr,
				err.Error())

			return
		}

		err = iaResult.IsValid(ePrefix + "Validating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult.GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of loop

	iaResultNumSeps, err := iaResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumNumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(\n"+
			"  multiplierNumStrDto, nDtoArray[0-5])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStrDtoNumStr, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Numbers Not Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values Not Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because iaResultNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(iaResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != iaResultNumSeps \n"+
			"Expected iaResultNumSeps = '%v'\n"+
			"  Actual iaResultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), iaResultNumSeps.String())

		return
	}

	if !usaNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoSeries_02"

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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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

	err = iaResult.IsValid(ePrefix + "Validating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	var nDtoArrayNumStr, iaNumStr string

	var ia IntAry

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

		nDtoArrayNumStr, err = nDtoArray[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArrayNumStr, err =  nDtoArray[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"nDtoArray[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, nDtoArrayNumStr, err.Error())
			return
		}

		err = ia.IsValid(ePrefix + fmt.Sprintf(" Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia.IsValid(ePrefix)\n"+
				"Cycle No= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				i,
				iaNumStr,
				iaResultNumStr,
				err.Error())

			return
		}

		err = iaResult.IsValid(ePrefix + "Validating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult.GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of loop

	iaResultNumSeps, err := iaResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumNumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(\n"+
			"  multiplierNumStrDto, nDtoArray[0-5])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStrDtoNumStr, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Numbers Not Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values Not Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because iaResultNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(iaResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != iaResultNumSeps \n"+
			"Expected iaResultNumSeps = '%v'\n"+
			"  Actual iaResultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), iaResultNumSeps.String())

		return
	}

	if !usaNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoSeries_03"

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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	err = iaResult.IsValid(ePrefix + "Validating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	var nDtoArrayNumStr, iaNumStr string

	var ia IntAry

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

		nDtoArrayNumStr, err = nDtoArray[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArrayNumStr, err =  nDtoArray[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"nDtoArray[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, nDtoArrayNumStr, err.Error())
			return
		}

		err = ia.IsValid(ePrefix + fmt.Sprintf(" Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia.IsValid(ePrefix)\n"+
				"Cycle No= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				i,
				iaNumStr,
				iaResultNumStr,
				err.Error())

			return
		}

		err = iaResult.IsValid(ePrefix + "Validating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult.GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of loop

	iaResultNumSeps, err := iaResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumNumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(\n"+
			"  multiplierNumStrDto, nDtoArray[0-5])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStrDtoNumStr, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Numbers Not Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values Not Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because iaResultNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(iaResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != iaResultNumSeps \n"+
			"Expected iaResultNumSeps = '%v'\n"+
			"  Actual iaResultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), iaResultNumSeps.String())

		return
	}

	if !usaNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoSeries_04"

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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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

	err = iaResult.IsValid(ePrefix + "Validating iaResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	var nDtoArrayNumStr, iaNumStr string

	var ia IntAry

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

		nDtoArrayNumStr, err = nDtoArray[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArrayNumStr, err =  nDtoArray[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		ia, err = nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = nDtoArray[%d].GetIntAry()\n"+
				"nDtoArray[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, nDtoArrayNumStr, err.Error())
			return
		}

		err = ia.IsValid(ePrefix + fmt.Sprintf(" Validating ia[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia.IsValid(ePrefix)\n"+
				"Cycle No= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult[%d].MultiplyThisBy(&ia, -1, -1)\n"+
				"ia= '%v'\n"+
				"iaResult= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				i,
				iaNumStr,
				iaResultNumStr,
				err.Error())

			return
		}

		err = iaResult.IsValid(ePrefix + "Validating iaResult")

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

		iaResultNumStr, err = iaResult.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaResultNumStr, err := iaResult.GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

	} // End of loop

	iaResultNumSeps, err := iaResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumNumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(\n"+
			"  multiplierNumStrDto, nDtoArray[0-5])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStrDtoNumStr, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Numbers Not Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values Not Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if expectedBigINumNumberStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedBigINumNumberStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because iaResultNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(iaResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != iaResultNumSeps \n"+
			"Expected iaResultNumSeps = '%v'\n"+
			"  Actual iaResultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), iaResultNumSeps.String())

		return
	}

	if !usaNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyNumStrDtoSeries_05"

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
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStrWithNumSeps(multiplicandStrs[i], &usaNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).\n"+
				"  NewNumStrWithNumSeps(multiplicandStrs[%d], &usaNumSeps)\n"+
				"multiplicandStrs[%d]= '%v'\n"+
				"usaNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())

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

	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrDtoSeries(\n"+
			"  multiplierNumStrDto, nDtoArray[0-5])\n"+
			"multiplierNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStrDtoNumStr, err.Error())
		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
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
			"Error: Numeric Separator Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyPair_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyPair_01"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplierBiNum.IsValid(ePrefix + "Validating multiplierBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplicandStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplicandBiNum.IsValid(ePrefix + "Validating multiplicandBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, usaNumSeps.String(), err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumNumberStr)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).\n"+
			"  NewBigIntNum(multiplierBiNum, multiplicandBiNum)\n"+
			"multiplierBiNum= '%v'\n"+
			"multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyPair(bPair)\n"+
			"bPair.multiplierBiNum= '%v'\n"+
			"bPair.multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyPair_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyPair_02"

	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplierBiNum.IsValid(ePrefix + "Validating multiplierBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplicandStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplicandBiNum.IsValid(ePrefix + "Validating multiplicandBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, usaNumSeps.String(), err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumNumberStr)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).\n"+
			"  NewBigIntNum(multiplierBiNum, multiplicandBiNum)\n"+
			"multiplierBiNum= '%v'\n"+
			"multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyPair(bPair)\n"+
			"bPair.multiplierBiNum= '%v'\n"+
			"bPair.multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyPair_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyPair_03"

	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplierBiNum.IsValid(ePrefix + "Validating multiplierBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplicandStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplicandBiNum.IsValid(ePrefix + "Validating multiplicandBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, usaNumSeps.String(), err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumNumberStr)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).\n"+
			"  NewBigIntNum(multiplierBiNum, multiplicandBiNum)\n"+
			"multiplierBiNum= '%v'\n"+
			"multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyPair(bPair)\n"+
			"bPair.multiplierBiNum= '%v'\n"+
			"bPair.multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyPair_04(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyPair_04"

	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplierBiNum.IsValid(ePrefix + "Validating multiplierBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplicandStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplicandBiNum.IsValid(ePrefix + "Validating multiplicandBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, usaNumSeps.String(), err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumNumberStr)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).\n"+
			"  NewBigIntNum(multiplierBiNum, multiplicandBiNum)\n"+
			"multiplierBiNum= '%v'\n"+
			"multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyPair(bPair)\n"+
			"bPair.multiplierBiNum= '%v'\n"+
			"bPair.multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyPair_05(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyPair_05"

	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedNumStr := "0"

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	multiplierBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplierBiNum.IsValid(ePrefix + "Validating multiplierBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplicandStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplicandBiNum.IsValid(ePrefix + "Validating multiplicandBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &usaNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, usaNumSeps.String(), err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumNumberStr)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).\n"+
			"  NewBigIntNum(multiplierBiNum, multiplicandBiNum)\n"+
			"multiplierBiNum= '%v'\n"+
			"multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyPair(bPair)\n"+
			"bPair.multiplierBiNum= '%v'\n"+
			"bPair.multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyPair_06(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyPair_06"

	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
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

	multiplierBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplierStr, &usaNumSeps)\n"+
			"multiplierStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplierBiNum.IsValid(ePrefix + "Validating multiplierBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplierBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(multiplicandStr, &usaNumSeps)\n"+
			"multiplicandStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplicandStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = multiplicandBiNum.IsValid(ePrefix + "Validating multiplicandBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = multiplicandBiNum.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
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

	expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != expectedBigINumNumberStr\n"+
			"Expected expectedBigINumNumberStr = '%v'\n"+
			"  Actual expectedBigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumNumberStr)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).\n"+
			"  NewBigIntNum(multiplierBiNum, multiplicandBiNum)\n"+
			"multiplierBiNum= '%v'\n"+
			"multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyPair(bPair)\n"+
			"bPair.multiplierBiNum= '%v'\n"+
			"bPair.multiplicandBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			multiplierBiNumStr,
			multiplicandBiNumStr,
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid(ePrefix + "Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultNumStr, err.Error())
		return
	}

	expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualsResult, err :=\n"+
			"  expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'"+
			"result= '%v'"+
			"Error='%v'\n\n", ePrefix, expectedBigINumNumberStr,
			resultNumStr, err.Error())
		return
	}

	if !expectedBigINumEqualsResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumEqualsResult = 'false'\n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumNumberStr, resultNumStr)

		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_01"

	baseStr := "35"

	tenExponentStr := "3"

	expectedStr := "35,000"

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseBINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(baseStr, &usaNumSeps)\n"+
			"baseStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = baseBINum.IsValid(ePrefix + "Validating baseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = baseBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	baseBINumNumberStr, err := baseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'
	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = alternateNumSeps.IsValid("Validating alternateNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating alternateNumSeps')\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, alternateNumSeps.String(), err.Error())
		return
	}

	tenExpBINum, err := new(BigIntNum).NewNumStrWithNumSeps(tenExponentStr, &alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"tenExpBINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(tenExponentStr, &alternateNumSeps)\n"+
			"tenExponentStr= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			tenExponentStr,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = tenExpBINum.IsValid(ePrefix + "Validating tenExpBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = tenExpBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	tenExpBINumNumberStr, err := tenExpBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"tenExpBINumNumberStr, err := tenExpBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	// Numseps taken from baseBINum
	result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, 100)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, 100)\n"+
			"baseBINum= '%v'\n"+
			"tenExpBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBINumNumberStr,
			tenExpBINumNumberStr,
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	if !usaNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because usaNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, usaNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_02"

	baseStr := "35,9657"

	tenExponentStr := "3"

	expectedStr := "35965,7"

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
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseBINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(baseStr, &expectedNumSeps)\n"+
			"baseStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = baseBINum.IsValid(ePrefix + "Validating baseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = baseBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	baseBINumNumberStr, err := baseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	tenExpBINum, err := new(BigIntNum).NewNumStrWithNumSeps(tenExponentStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"tenExpBINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(tenExponentStr, &usaNumSeps)\n"+
			"tenExponentStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			tenExponentStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = tenExpBINum.IsValid(ePrefix + "Validating tenExpBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = tenExpBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	tenExpBINumNumberStr, err := tenExpBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"tenExpBINumNumberStr, err := tenExpBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, 100)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, 100)\n"+
			"baseBINum= '%v'\n"+
			"tenExpBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBINumNumberStr,
			tenExpBINumNumberStr,
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_02"

	baseStr := "35"

	tenExponentStr := "3.9"
	// 10^3.9 = 7943.2823472428150206591828283639

	expectedStr := "278014,88215349852572307139899274"

	// 278014.8821534985257230713989927365
	maxPrecision := uint(26)

	expectedNumSeps := new(NumericSeparatorDto).New()
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
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseBINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(baseStr, &expectedNumSeps)\n"+
			"baseStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = baseBINum.IsValid(ePrefix + "Validating baseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = baseBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	baseBINumNumberStr, err := baseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	tenExpBINum, err := new(BigIntNum).NewNumStrWithNumSeps(tenExponentStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"tenExpBINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(tenExponentStr, &usaNumSeps)\n"+
			"tenExponentStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			tenExponentStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = tenExpBINum.IsValid(ePrefix + "Validating tenExpBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = tenExpBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	tenExpBINumNumberStr, err := tenExpBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"tenExpBINumNumberStr, err := tenExpBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).\n"+
			"  MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, maxPrecision)\n"+
			"baseBINum= '%v'\n"+
			"tenExpBINum= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBINumNumberStr,
			tenExpBINumNumberStr,
			maxPrecision,
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

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}
