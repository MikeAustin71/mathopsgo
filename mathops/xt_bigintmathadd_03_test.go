package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathAdd_AddNumStr_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStr_01"
	n1Str := "123456.789"
	n2Str := "987.123456"
	// Result = 	124443.912456
	expectedFinalResult := "124443.912456"
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	dto := &NumericSeparatorDto{}
	expectedNumSeps := dto.New()

	result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)\n"+
			"n1Str= '%v'\n"+
			"n2Str= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			n1Str,
			n2Str,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStr_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStr_02"
	n1Str := "123456.789"
	n2Str := "-987.123456"

	expectedResultStr := "122469665544"
	expectedFinalResult := "122469.665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	dto := &NumericSeparatorDto{}
	expectedNumSeps := dto.New()

	result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)\n"+
			"n1Str= '%v'\n"+
			"n2Str= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			n1Str,
			n2Str,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStr_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStr_03"
	n1Str := "-123456.789"
	n2Str := "987.123456"

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedFinalResult := "-122469.665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	dto := &NumericSeparatorDto{}
	expectedNumSeps := dto.New()

	result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)\n"+
			"n1Str= '%v'\n"+
			"n2Str= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			n1Str,
			n2Str,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStr_04(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStr_04"
	n1Str := "-123456.789"
	n2Str := "-987.123456"

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedFinalResult := "-124443.912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.SetDefaultsIfEmpty()

	result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)\n"+
			"n1Str= '%v'\n"+
			"n2Str= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			n1Str,
			n2Str,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStr_05(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStr_05"
	n1Str := "123456.789"
	n2Str := "987.123456"
	// Result = 	124443.912456
	expectedFinalResult := "124443.912456"
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)\n"+
			"n1Str= '%v'\n"+
			"n2Str= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			n1Str,
			n2Str,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStr_06(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStr_06"
	n1Str := "123456,789"
	n2Str := "987,123456"
	// Result = 	124443,912456
	expectedFinalResult := "124443,912456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStr(n1Str, n2Str, expectedNumSeps)\n"+
			"n1Str= '%v'\n"+
			"n2Str= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			n1Str,
			n2Str,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStrOutputToArray_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrOutputToArray_01"
	var err error

	// addendStr = 5
	addendStr := "5"

	// numStrsArray
	numStrsArray := []string{
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

	lenArray := len(numStrsArray)

	dto := &NumericSeparatorDto{}
	expectedNumSeps := dto.New()

	resultArray, err := new(BigIntMathAdd).AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathAdd).AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)\n"+
			"addendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			addendStr,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	var actualNumStr string

	for j := 0; j < lenArray; j++ {

		actualNumStr = resultArray[j]

		if expectedNumStrs[j] != actualNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected actualNumStr[%d] = '%v'\n"+
				"Instead, actualNumStr[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				actualNumStr)

			return
		}
	}

	return
}

func TestBigIntMathAdd_AddNumStrOutputToArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrOutputToArray_02"
	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// numStrsArray
	numStrsArray := []string{
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

	lenArray := len(numStrsArray)

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.SetDefaultsIfEmpty()

	resultArray, err :=
		new(BigIntMathAdd).AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathAdd).AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)\n"+
			"addendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			addendStr,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	var actualNumStr string

	for j := 0; j < lenArray; j++ {

		actualNumStr = resultArray[j]

		if expectedNumStrs[j] != actualNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected actualNumStr[%d] = '%v'\n"+
				"Instead, actualNumStr[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				actualNumStr)

			return
		}
	}

	return
}

func TestBigIntMathAdd_AddNumStrOutputToArray_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrOutputToArray_03"
	var err error

	// addendStr = 3.1
	addendStr := "3,1"

	// numStrsArray
	numStrsArray := []string{
		"5",
		"10,123",
		"0",
		"253,692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"8,1",
		"13,223",
		"3,1",
		"256,792",
		"38,1",
		"58,1",
	}

	lenArray := len(numStrsArray)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	resultArray, err :=
		new(BigIntMathAdd).AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathAdd).AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)\n"+
			"addendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			addendStr,
			expectedNumSeps.String(),
			ePrefix,
			err.Error())
		return
	}

	var actualNumStr string

	for j := 0; j < lenArray; j++ {

		actualNumStr = resultArray[j]

		if expectedNumStrs[j] != actualNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected actualNumStr[%d] = '%v'\n"+
				"Instead, actualNumStr[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				actualNumStr)

			return
		}
	}
}

func TestBigIntMathAdd_AddNumStrArray_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrArray_01"
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
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	dto := &NumericSeparatorDto{}
	expectedNumSeps := dto.New()

	total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

func TestBigIntMathAdd_AddNumStrArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrArray_02"
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedNumSeps := new(NumericSeparatorDto).New()

	expectedNumSeps.SetDefaultsIfEmpty()

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

func TestBigIntMathAdd_AddNumStrArray_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrArray_03"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedNumSeps := new(NumericSeparatorDto).New()
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

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
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddNumStrArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStrArray_04"
	numStrAry := []string{
		"45,8",
		"1,45962",
		"58,71",
		"-37,62174",
		"89,8",
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	expectedTotalStr := "158,14788"

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrArray(numStrAry, expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

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
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddNumStrSeries_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrSeries_01"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedNumSeps := new(NumericSeparatorDto).New()

	expectedNumSeps.SetDefaultsIfEmpty()

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

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
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			expectedNumSeps.String(),
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
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

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
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddNumStrSeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStrSeries_02"
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

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
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			expectedNumSeps.String(),
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
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

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
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddNumStrSeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStrSeries_03"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

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
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			expectedNumSeps.String(),
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
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

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
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddNumStrSeries_04(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrSeries_04"
	numStrAry := []string{
		"45,8",
		"1,45962",
		"58,71",
		"-37,62174",
		"89,8",
	}

	expectedTotalStr := "158,14788"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedTotalStr, &expectedNumSeps)\n"+
			"expectedTotalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			expectedNumSeps.String(),
			err.Error())
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
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	total, err := new(BigIntMathAdd).AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddNumStrSeries(...). "+
			"Error='%v' ", err.Error())
	}

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
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			numStrAry[0],
			numStrAry[1],
			numStrAry[2],
			numStrAry[3],
			numStrAry[4],
			expectedNumSeps.String(),
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
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

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
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddNumStrDto_01(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStrDto_01"
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedFinalResult := "124443.912456"
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	n1Dto, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	n1DtoNumStr, err := n1Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1DtoNumStr, err := n1Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	n2Dto, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	n2DtoNumStr, err := n2Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2DtoNumStr, err := n2Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)\n"+
			"n1Dto= '%v'\n"+
			"n2Dto= '%v'\n"+
			"Error='%v'\n\n",
			n1DtoNumStr,
			n2DtoNumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddNumStrDto_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDto_02"
	n1Str := "123456.789"
	n2Str := "-987.123456"
	expectedFinalResult := "122469.665544"
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	n1DtoNumStr, err := n1Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1DtoNumStr, err := n1Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	n2Dto, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	n2DtoNumStr, err := n2Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2DtoNumStr, err := n2Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)\n"+
			"n1Dto= '%v'\n"+
			"n2Dto= '%v'\n"+
			"Error='%v'\n\n",
			n1DtoNumStr,
			n2DtoNumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddNumStrDto_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDto_03"
	n1Str := "-123456.789"
	n2Str := "987.123456"
	// Result := -122469.665544
	expectedFinalResult := "-122469.665544"
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	n1Dto, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	n1DtoNumStr, err := n1Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1DtoNumStr, err := n1Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	n2Dto, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	n2DtoNumStr, err := n2Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2DtoNumStr, err := n2Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)\n"+
			"n1Dto= '%v'\n"+
			"n2Dto= '%v'\n"+
			"Error='%v'\n\n",
			n1DtoNumStr,
			n2DtoNumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddNumStrDto_04(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_04"
	n1Str := "-123456.789"
	n2Str := "-987.123456"
	// Result := -124443.912456
	expectedFinalResult := "-124443.912456"
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	n1Dto, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	n1DtoNumStr, err := n1Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1DtoNumStr, err := n1Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	n2Dto, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	n2DtoNumStr, err := n2Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2DtoNumStr, err := n2Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)\n"+
			"n1Dto= '%v'\n"+
			"n2Dto= '%v'\n"+
			"Error='%v'\n\n",
			n1DtoNumStr,
			n2DtoNumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddNumStrDto_05(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStrDto_05"
	n1Str := "123456.789"
	n2Str := "987.123456"

	expectedFinalResult := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	n1Dto, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	n1DtoNumStr, err := n1Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n1DtoNumStr, err := n1Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = n1Dto.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = n1Dto.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	n2Dto, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	n2DtoNumStr, err := n2Dto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"n2DtoNumStr, err := n2Dto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddNumStrDto(n1Dto, n2Dto)\n"+
			"n1Dto= '%v'\n"+
			"n2Dto= '%v'\n"+
			"Error='%v'\n\n",
			n1DtoNumStr,
			n2DtoNumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

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

func TestBigIntMathAdd_AddNumStrDtoArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddNumStrDtoArray_01"
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

	total, err := new(BigIntMathAdd).AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrDtoArray(numStrDtoAry)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathAdd_AddNumStrDtoArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoArray_02"
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
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			err.Error())
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

	total, err := new(BigIntMathAdd).AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrDtoArray(numStrDtoAry)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathAdd_AddNumStrDtoArray_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddNumStrDtoArray_03"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	var err error

	expectedTotalStr := "158,14788"

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

	total, err := new(BigIntMathAdd).AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddNumStrDtoArray(numStrDtoAry)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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

}
