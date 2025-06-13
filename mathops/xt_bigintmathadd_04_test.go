package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathAdd_AddNumStrDtoOutputToArray_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoOutputToArray_01"
	var err error

	// addendStr = 5
	addendStr := "5"

	// numStrs
	numStrs := []string{
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}

	nDtoAddend, err := new(NumStrDto).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoAddend, err := new(NumStrDto).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n",
			addendStr,
			ePrefix,
			err.Error())
		return
	}

	nDtoAddendNumStr, err := nDtoAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoAddendNumStr, err := nDtoAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if addendStr != nDtoAddendNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected nDtoAddendNumStr = '%v'\n"+
			"Instead, nDtoAddendNumStr = '%v'\n\n",
			ePrefix, addendStr, nDtoAddendNumStr)

		return
	}

	lenArray := len(numStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).NewNumStr(numStrs[%d])\n"+
				"numStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, numStrs[i], err.Error())
			return
		}
	}

	result, err := new(BigIntMathAdd).AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)\n"+
			"nDtoAddend= '%v'\n"+
			"Error='%v'\n\n",
			nDtoAddendNumStr,
			ePrefix,
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
				"Expected actualNumStr[%d] = '%v'\n"+
				"Instead, actualNumStr[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)
			return
		}
	}

	return
}

func TestBigIntMathAdd_AddNumStrDtoOutputToArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoOutputToArray_02"
	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// numStrs
	numStrs := []string{
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	nDtoAddend, err := new(NumStrDto).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoAddend, err := new(NumStrDto).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n",
			addendStr,
			ePrefix,
			err.Error())
		return
	}

	nDtoAddendNumStr, err := nDtoAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoAddendNumStr, err := nDtoAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if addendStr != nDtoAddendNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected nDtoAddendNumStr = '%v'\n"+
			"Instead, nDtoAddendNumStr = '%v'\n\n",
			ePrefix, addendStr, nDtoAddendNumStr)

		return
	}

	lenArray := len(numStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).NewNumStr(numStrs[%d])\n"+
				"numStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, numStrs[i], err.Error())
			return
		}

	}

	result, err := new(BigIntMathAdd).AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)\n"+
			"nDtoAddend= '%v'\n"+
			"Error='%v'\n\n",
			nDtoAddendNumStr,
			ePrefix,
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
				"Expected actualNumStr[%d] = '%v'\n"+
				"Instead, actualNumStr[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)
			return
		}
	}

	return
}

func TestBigIntMathAdd_AddNumStrDtoOutputToArray_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoOutputToArray_03"
	var err error

	// addendStr = 5
	addendStr := "5"

	// numStrs
	numStrs := []string{
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"10",
		"15,123",
		"20",
		"258,692",
		"40",
		"60",
	}

	nDtoAddend, err := new(NumStrDto).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoAddend, err := new(NumStrDto).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n",
			addendStr,
			ePrefix,
			err.Error())
		return
	}

	nDtoAddendNumStr, err := nDtoAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoAddendNumStr, err := nDtoAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if addendStr != nDtoAddendNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected nDtoAddendNumStr = '%v'\n"+
			"Instead, nDtoAddendNumStr = '%v'\n\n",
			ePrefix, addendStr, nDtoAddendNumStr)

		return
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = nDtoAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoAddend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenArray := len(numStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDtoArray[%d], err = new(NumStrDto).NewNumStr(numStrs[%d])\n"+
				"numStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, numStrs[i], err.Error())
			return
		}
	}

	result, err := new(BigIntMathAdd).AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)\n"+
			"nDtoAddend= '%v'\n"+
			"Error='%v'\n\n",
			nDtoAddendNumStr,
			ePrefix,
			err.Error())
		return
	}

	var resultNumStr string
	var actualNumSeps NumericSeparatorDto

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
				"Expected actualNumStr[%d] = '%v'\n"+
				"Instead, actualNumStr[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)
			return
		}

		actualNumSeps, err = result[j].GetNumericSeparatorsDto()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected total NumSeps = '%v'\n"+
				"Instead, total NumSeps = '%v'\n"+
				"Cycle Index= '%v'\n\n",
				ePrefix,
				expectedNumSeps.String(),
				actualNumSeps.String(),
				j)
			return
		}
	}
	return
}

func TestBigIntMathAdd_AddNumStrDtoSeries_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoSeries_01"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
			"expectedTotalStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedTotalStr, err.Error())
		return
	}

	expectedResultNumStr, err := expectedBNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedTotalStr != expectedResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bNum expectedResultNumStr = '%v'\n"+
			"Instead, bNum expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i := 0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = new(NumStrDto).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"numStrDtoAry[%d], err = new(NumStrDto).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, numStrAry[i], err.Error())
			return
		}
	}

	total, err := new(BigIntMathAdd).AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrSeries("+
			"  numStrAry[0], numStrAry[1], numStrAry[2],\n"+
			"  numStrAry[3], numStrAry[4])\n"+
			"numStrAry[0]= '%v'\n"+
			"numStrAry[1]= '%v'\n"+
			"numStrAry[2]= '%v'\n"+
			"numStrAry[3]= '%v'\n"+
			"numStrAry[4]= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			err.Error())
		return
	}

	totalNumStr, err := total.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalNumStr, err := total.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBNumEqualsTotal, err := expectedBNum.Equal(total)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNumEqualsTotal, err := expectedBNum.Equal(total)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBNumEqualsTotal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix, expectedResultNumStr, totalNumStr)
	}

	return
}

func TestBigIntMathAdd_AddNumStrDtoSeries_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoSeries_02"
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
			"expectedTotalStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedTotalStr, err.Error())
		return
	}

	expectedResultNumStr, err := expectedBNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedTotalStr != expectedResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bNum expectedResultNumStr = '%v'\n"+
			"Instead, bNum expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i := 0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = new(NumStrDto).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"numStrDtoAry[%d], err = new(NumStrDto).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, numStrAry[i], err.Error())
			return
		}

	}

	total, err := new(BigIntMathAdd).AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrSeries("+
			"  numStrAry[0], numStrAry[1], numStrAry[2],\n"+
			"  numStrAry[3], numStrAry[4])\n"+
			"numStrAry[0]= '%v'\n"+
			"numStrAry[1]= '%v'\n"+
			"numStrAry[2]= '%v'\n"+
			"numStrAry[3]= '%v'\n"+
			"numStrAry[4]= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			err.Error())
		return
	}

	totalNumStr, err := total.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalNumStr, err := total.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBNumEqualsTotal, err := expectedBNum.Equal(total)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNumEqualsTotal, err := expectedBNum.Equal(total)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBNumEqualsTotal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix, expectedResultNumStr, totalNumStr)
	}

	return
}

func TestBigIntMathAdd_AddNumStrDtoSeries_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoSeries_03"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158,14788"

	var err error

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i := 0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = new(NumStrDto).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"numStrDtoAry[%d], err = new(NumStrDto).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, numStrAry[i], err.Error())
			return
		}

	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = numStrDtoAry[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoAry[0].SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrSeries("+
			"  numStrAry[0], numStrAry[1], numStrAry[2],\n"+
			"  numStrAry[3], numStrAry[4])\n"+
			"numStrAry[0]= '%v'\n"+
			"numStrAry[1]= '%v'\n"+
			"numStrAry[2]= '%v'\n"+
			"numStrAry[3]= '%v'\n"+
			"numStrAry[4]= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			err.Error())
		return
	}

	totalNumStr, err := total.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalNumStr, err := total.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedTotalStr != totalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix, expectedTotalStr, totalNumStr)
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualNumSeps = '%v'\n"+
			"Instead, actualNumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddPair_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_01"
	// n1Str := 123456.789
	b1Str := "123456.789"
	b1Precision := uint(3)

	// n2Str := 987.123456
	b2Str := "987.123456"
	b2Precision := uint(6)

	expectedResultStr := "124443.912456"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_02"
	// n1Str := 123456.789
	b1Str := "123456.789"
	b1Precision := uint(3)

	// n2Str := -987.123456
	b2Str := "-987.123456"
	b2Precision := uint(6)

	expectedResultStr := "122469.665544"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_03"
	// n1Str := -123456.789
	b1Str := "-123456.789"
	b1Precision := uint(3)

	// n2Str := 987.123456
	b2Str := "987.123456"
	b2Precision := uint(6)

	expectedResultStr := "-122469.665544"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_04(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_04"
	// n1Str := -123456.789
	b1Str := "-123456.789"
	b1Precision := uint(3)

	// n2Str := -987.123456
	b2Str := "-987.123456"
	b2Precision := uint(6)

	expectedResultStr := "-124443.912456"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_05(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_05"
	// n1Str := 0
	b1Str := "0"
	b1Precision := uint(0)

	// n2Str := -987.123456
	b2Str := "-987.123456"
	b2Precision := uint(6)

	expectedResultStr := "-987.123456"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_06(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_06"
	// n1Str := -123456.789
	b1Str := "-123456.789"
	b1Precision := uint(3)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "-123456.789"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_07(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_07"
	// n1Str := 7
	b1Str := "7"
	b1Precision := uint(0)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "7"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_08(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_08"
	// n1Str := 0
	b1Str := "0"
	b1Precision := uint(0)

	// n2Str := 7
	b2Str := "7"
	b2Precision := uint(0)

	expectedResultStr := "7"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_09(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_09"
	// n1Str := 1.7
	b1Str := "1.7"
	b1Precision := uint(1)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "1.7"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_10(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_10"
	// n1Str := 0
	b1Str := "0"
	b1Precision := uint(0)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "0"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedResultStr)\n"+
			"expectedResultStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedResultStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != expectedBigINumStr\n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"Instead, expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedResultStr, expectedBigINumStr)

		return
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig1PrecisionUint, err :=\tbPair.Big1.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"Instead, bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bPairBig2PrecisionUint, err :=\tbPair.Big2.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"Instead, bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumIsEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumIsEqualResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedBigINum != result\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)
	}

	return
}

func TestBigIntMathAdd_AddPair_11(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddPair_11"
	// n1Str := 123456.789
	b1Str := "123456.789"
	b1Precision := uint(3)

	// n2Str := 987.123456
	b2Str := "987.123456"
	b2Precision := uint(6)

	expectedResultStr := "124443,912456"

	b1Num, err := new(BigIntNum).NewNumStr(b1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
			"b1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Str, err.Error())
		return
	}

	b1NumStr, err := b1Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumStr, err := b1Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Str != b1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumStr = '%v'\n"+
			"Instead, b1NumStr = '%v'\n\n",
			ePrefix, b1Str, b1NumStr)

		return
	}

	b1NumPrecisionUint, err := b1Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1NumPrecisionUint, err := b1Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b1Precision != b1NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b1NumPrecisionUint = '%v'\n"+
			"Instead, b1NumPrecisionUint = '%v'\n\n",
			ePrefix, b1Precision, b1NumPrecisionUint)

		return
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = b1Num.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = b1Num.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	b2Num, err := new(BigIntNum).NewNumStr(b2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
			"b2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Str, err.Error())
		return
	}

	b2NumStr, err := b2Num.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumStr, err := b2Num.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Str != b2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumStr = '%v'\n"+
			"Instead, b2NumStr = '%v'\n\n",
			ePrefix, b2Str, b2NumStr)

		return
	}

	b2NumPrecisionUint, err := b2Num.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2NumPrecisionUint, err := b2Num.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if b2Precision != b2NumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected b2NumPrecisionUint = '%v'\n"+
			"Instead, b2NumPrecisionUint = '%v'\n\n",
			ePrefix, b2Precision, b2NumPrecisionUint)

		return
	}

	bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(b1Num, b2Num)\n"+
			"b1Num= '%v'\n"+
			"b2Num= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1NumStr, b2NumStr, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
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

	if expectedResultStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedResultStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"Instead, resultNumStr = '%v'\n\n",
			ePrefix, expectedResultStr, resultNumStr)

		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err := total.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualNumSeps = '%v'\n"+
			"Instead, actualNumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_BigIntAdd_01"
	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := big.NewInt(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := big.NewInt(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := 124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := big.NewInt(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_BigIntAdd_02"
	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := big.NewInt(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := big.NewInt(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := 122469.665544
	expectedResultStr := "122469665544"
	expectedPrecision := big.NewInt(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_03(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_BigIntAdd_03"
	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := big.NewInt(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := big.NewInt(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := big.NewInt(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_04(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_BigIntAdd_04"
	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := big.NewInt(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := big.NewInt(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := big.NewInt(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_05(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_BigIntAdd_05"
	b1Str := "0000"
	b1Precision := big.NewInt(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	b2Str := "0000000"
	b2Precision := big.NewInt(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	expectedResultStr := "0"
	expectedPrecision := big.NewInt(0)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_06(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_BigIntAdd_06"
	b1Str := "51"
	b1Precision := big.NewInt(1)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	b2Str := "-51"
	b2Precision := big.NewInt(1)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	expectedResultStr := "0"
	expectedPrecision := big.NewInt(0)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_BigIntAdd_07(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_BigIntAdd_07"
	b1Str := "51"
	b1Precision := big.NewInt(0)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	b2Str := "100"
	b2Precision := big.NewInt(2)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	expectedResultStr := "52"
	expectedPrecision := big.NewInt(0)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(b1Big, b1Precision, b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathAdd).BigIntAdd(\n"+
			"b1Big, b1Precision, b2Big, b2Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			b1Big.Text(10),
			b1Precision,
			b2Big.Text(10),
			b2Precision,
			err.Error())

		return
	}

	if biExpectedResult.Cmp(result) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(result) != 0\n"+
			"Expected result = '%v'\n"+
			"Instead, result = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), result.Text(10))

		return
	}

	if expectedPrecision.Cmp(resultPrecision) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision.Cmp(resultPrecision) != 0\n"+
			"Expected resultPrecision = '%v'\n"+
			"Instead, resultPrecision = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecision)
	}

	return
}

func TestBigIntMathAdd_FixedDecimalAdd_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_FixedDecimalAdd_01"
	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := 124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Big.Text(10), b1Precision, err.Error())
		return
	}

	b1FdNumStr, err := b1Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1FdNumStr, err := b1Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Big.Text(10), b2Precision, err.Error())
		return
	}

	b2FdNumStr, err := b2Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2FdNumStr, err := b2Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)\n"+
			"b1Fd= '%v'\n"+
			"b2Fd= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1FdNumStr, b2FdNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if biExpectedResult.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(resultBigInt) != 0"+
			"Expected resultBigInt = '%v'\n"+
			"Instead, resultBigInt = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), resultBigInt.Text(10))

		return
	}

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision != resultPrecisionUint\n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"Instead, resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)
	}

	return
}

func TestBigIntMathAdd_FixedDecimalAdd_02(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_FixedDecimalAdd_02"
	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := 122469.665544
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Big.Text(10), b1Precision, err.Error())
		return
	}

	b1FdNumStr, err := b1Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1FdNumStr, err := b1Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Big.Text(10), b2Precision, err.Error())
		return
	}

	b2FdNumStr, err := b2Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2FdNumStr, err := b2Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)\n"+
			"b1Fd= '%v'\n"+
			"b2Fd= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1FdNumStr, b2FdNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if biExpectedResult.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(resultBigInt) != 0"+
			"Expected resultBigInt = '%v'\n"+
			"Instead, resultBigInt = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), resultBigInt.Text(10))

		return
	}

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision != resultPrecisionUint\n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"Instead, resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)
	}

	return
}

func TestBigIntMathAdd_FixedDecimalAdd_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_FixedDecimalAdd_03"
	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Big.Text(10), b1Precision, err.Error())
		return
	}

	b1FdNumStr, err := b1Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1FdNumStr, err := b1Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Big.Text(10), b2Precision, err.Error())
		return
	}

	b2FdNumStr, err := b2Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2FdNumStr, err := b2Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)\n"+
			"b1Fd= '%v'\n"+
			"b2Fd= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1FdNumStr, b2FdNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if biExpectedResult.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(resultBigInt) != 0"+
			"Expected resultBigInt = '%v'\n"+
			"Instead, resultBigInt = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), resultBigInt.Text(10))

		return
	}

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision != resultPrecisionUint\n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"Instead, resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)
	}

	return
}

func TestBigIntMathAdd_FixedDecimalAdd_04(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_FixedDecimalAdd_04"
	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)

	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
			"b1Str='%v'\n\n", ePrefix, b1Str)
		return
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)

	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
			"b2Str='%v'\n\n", ePrefix, b2Str)
		return
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1Fd, err := new(BigIntFixedDecimal).New(b1Big, b1Precision)\n"+
			"b1Big= '%v'\n"+
			"b1Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1Big.Text(10), b1Precision, err.Error())
		return
	}

	b1FdNumStr, err := b1Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b1FdNumStr, err := b1Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2Fd, err := new(BigIntFixedDecimal).New(b2Big, b2Precision)\n"+
			"b2Big= '%v'\n"+
			"b2Precision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b2Big.Text(10), b2Precision, err.Error())
		return
	}

	b2FdNumStr, err := b2Fd.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"b2FdNumStr, err := b2Fd.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).FixedDecimalAdd(b1Fd, b2Fd)\n"+
			"b1Fd= '%v'\n"+
			"b2Fd= '%v'\n"+
			"Error='%v'\n\n", ePrefix, b1FdNumStr, b2FdNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetIntegerValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetIntegerValue()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if biExpectedResult.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"biExpectedResult.Cmp(resultBigInt) != 0"+
			"Expected resultBigInt = '%v'\n"+
			"Instead, resultBigInt = '%v'\n\n",
			ePrefix, biExpectedResult.Text(10), resultBigInt.Text(10))

		return
	}

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"expectedPrecision != resultPrecisionUint\n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"Instead, resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)
	}

	return
}
