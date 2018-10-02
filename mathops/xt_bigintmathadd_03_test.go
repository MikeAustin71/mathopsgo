package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathAdd_AddNumStr_01(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "987.123456"
	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedNumStr := "124443.912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	expectedNumSeps := NumericSeparatorDto{}.New()

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Result sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Result NumSeps='%v'. Instead, Result NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStr_02(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "-987.123456"

	expectedResultStr := "122469665544"
	expectedNumStr := "122469.665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	expectedNumSeps := NumericSeparatorDto{}.New()

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Result NumSeps='%v'. Instead, Result NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStr_03(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "987.123456"

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedNumStr := "-122469.665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	expectedNumSeps := NumericSeparatorDto{}.New()

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Result NumSeps='%v'. Instead, Result NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStr_04(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "-987.123456"

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedNumStr := "-124443.912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	expectedNumSeps := NumericSeparatorDto{}

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
			"Error='%v' ", err.Error())
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Result NumSeps='%v'. Instead, Result NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStr_05(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "987.123456"
	// Result = 	124443.912456
	expectedNumStr := "124443.912456"
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'.  Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathAdd_AddNumStr_06(t *testing.T) {

	n1Str := "123456,789"
	n2Str := "987,123456"
	// Result = 	124443,912456
	expectedNumStr := "124443,912456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
			"Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'.  Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathAdd_AddNumStrOutputToArray_01(t *testing.T) {

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

	expectedNumSeps := NumericSeparatorDto{}.New()

	resultArray, err := BigIntMathAdd{}.AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrOutputToArray("+
			"addendStr, numStrsArray, expectedNumSeps) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		actualNumStr := resultArray[j]

		if expectedNumStrs[j] != actualNumStr {
			t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' j='%v'",
				expectedNumStrs[j], actualNumStr, j)
		}

	}
}

func TestBigIntMathAdd_AddNumStrOutputToArray_02(t *testing.T) {

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

	resultArray, err :=
		BigIntMathAdd{}.AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrOutputToArray("+
			"addendStr, numStrsArray, expectedNumSeps) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	for j := 0; j < lenArray; j++ {

		actualNumStr := resultArray[j]

		if expectedNumStrs[j] != actualNumStr {
			t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' j='%v'",
				expectedNumStrs[j], actualNumStr, j)
		}

	}
}

func TestBigIntMathAdd_AddNumStrOutputToArray_03(t *testing.T) {

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
		BigIntMathAdd{}.AddNumStrOutputToArray(addendStr, numStrsArray, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrOutputToArray("+
			"addendStr, numStrsArray, expectedNumSeps) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		actualNumStr := resultArray[j]

		if expectedNumStrs[j] != actualNumStr {
			t.Errorf("Error: Expected Result NumStr='%v'. Instead, Result NumStr='%v' j='%v'",
				expectedNumStrs[j], actualNumStr, j)
		}

	}
}

func TestBigIntMathAdd_AddNumStrArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	expectedNumSeps := NumericSeparatorDto{}.New()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry, expNumSeps). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedNumSeps := NumericSeparatorDto{}

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry,expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrArray_03(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedTotalStr, expectedNumSeps). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'.  Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStrArray_04(t *testing.T) {
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

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedTotalStr, expectedNumSeps). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'.  Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStrSeries_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedNumSeps := NumericSeparatorDto{}

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedTotalStr, expectedNumSeps). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())
	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). "+
			"Error='%v' ", err.Error())
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStrSeries_02(t *testing.T) {

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

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedTotalStr, expectedNumSeps). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStrSeries_03(t *testing.T) {
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

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedTotalStr, expectedNumSeps). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'.  Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddNumStrSeries_04(t *testing.T) {
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

	expectedBNum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedTotalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedTotalStr, expectedNumSeps). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		expectedNumSeps,
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4])

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'.  Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathAdd_AddNumStrDto_01(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). "+
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddNumStrDto_02(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "-987.123456"
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). "+
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}
}

func TestBigIntMathAdd_AddNumStrDto_03(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "987.123456"
	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). "+
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}
}

func TestBigIntMathAdd_AddNumStrDto_04(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "-987.123456"
	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). "+
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}
}

func TestBigIntMathAdd_AddNumStrDto_05(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "987.123456"

	expectedResultStr := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
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
		t.Errorf("Error returned by n1Dto.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). "+
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			expectedResultStr, actualResultStr)
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathAdd_AddNumStrDtoArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i := 0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). "+
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). "+
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrDtoArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i := 0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). "+
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). "+
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrDtoArray_03(t *testing.T) {
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

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). "+
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
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
		t.Errorf("Error returned by numStrDtoAry[0].SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). "+
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
