package mathops

import (
	"testing"
	"math/big"
)

func TestBigIntMathAdd_AddNumStr_01(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "987.123456"
	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
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

}

func TestBigIntMathAdd_AddNumStr_02(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "-987.123456"

	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	result,err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
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

}

func TestBigIntMathAdd_AddNumStr_03(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "987.123456"

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
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

}

func TestBigIntMathAdd_AddNumStr_04(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "-987.123456"

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
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

}

func TestBigIntMathAdd_AddNumStr_05(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "987.123456"
	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	result, err := BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
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

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

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
	numStrsArray :=  [] string {
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}

	lenArray := len(numStrsArray)

	resultArray, err := BigIntMathAdd{}.AddNumStrOutputToArray(addendStr, numStrsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrOutputToArray(" +
			"addendStr, numStrsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != resultArray[j]  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, resultArray[j])
		}
	}
}

func TestBigIntMathAdd_AddNumStrOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// numStrsArray
	numStrsArray :=  [] string {
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	lenArray := len(numStrsArray)

	resultArray, err := BigIntMathAdd{}.AddNumStrOutputToArray(addendStr, numStrsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrOutputToArray(" +
			"addendStr, numStrsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != resultArray[j]  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, resultArray[j])
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
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

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
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

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrArray(numStrAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrArray(numStrAry). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}


	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

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

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
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

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
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

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	total, err := BigIntMathAdd{}.AddNumStrSeries(
		numStrAry[0],
		numStrAry[1],
		numStrAry[2],
		numStrAry[3],
		numStrAry[4],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}


	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

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

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
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

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
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

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
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

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
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

	expectedResultStr := 	"124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	n1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
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
		t.Errorf("Error returned by n1Dto.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddNumStrDto(n1Dto, n2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStr(AddNumStrDto(n1Dto, n2Dto)). " +
			"n1Dto.GetNumStr()='%v' n2Dto.GetNumStr()='%v' Error='%v' ",
			n1Dto.GetNumStr(), n2Dto.GetNumStr(), err.Error())
	}

	actualResultStr:= result.GetNumStr()

	if expectedResultStr!=actualResultStr {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			expectedResultStr,actualResultStr)
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). " +
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). " +
			"Length numStrDtoAry='%v' Error='%v' ", len(numStrDtoAry), err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
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

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
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
		t.Errorf("Error returned by numStrDtoAry[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	total, err := BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoArray(numStrDtoAry). " +
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

func TestBigIntMathAdd_AddNumStrDtoOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// numStrs
	numStrs :=  [] string {
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}


	nDtoAddend, err := NumStrDto{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(numStrs)
	nDtoArray := make([]NumStrDto, lenArray)


	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrs[i]) " +
				"i='%v'  numStrs[i]='%v'  Error='%v'. ", i, numStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoOutputToArray(" +
			"nDtoAddend, nDtoArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddNumStrDtoOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// numStrs
	numStrs :=  [] string {
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	nDtoAddend, err := NumStrDto{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(numStrs)
	decsArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		decsArray[i], err =	NumStrDto{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrs[i]) " +
				"i='%v'  numStrs[i]='%v'  Error='%v'. ", i, numStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddNumStrDtoOutputToArray(nDtoAddend, decsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoOutputToArray(" +
			"nDtoAddend, decsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddNumStrDtoOutputToArray_03(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// numStrs
	numStrs :=  [] string {
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"10",
		"15,123",
		"20",
		"258,692",
		"40",
		"60",
	}

	nDtoAddend, err := NumStrDto{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
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
		t.Errorf("Error returned by nDtoAddend.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(numStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrs[i]) " +
				"i='%v'  numStrs[i]='%v'  Error='%v'. ", i, numStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddNumStrDtoOutputToArray(nDtoAddend, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoOutputToArray(" +
			"nDtoAddend, nDtoArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error Index[%v]: Expected numSeps='%v'. Instead, numSeps='%v'",
				j, expectedNumSeps.String(), actualNumSeps.String())
		}
	}
}

func TestBigIntMathAdd_AddNumStrDtoSeries_01(t *testing.T) {
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4],	)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrDtoSeries_02(t *testing.T) {

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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	lenNStrAry := len(numStrAry)

	numStrDtoAry := make([]NumStrDto, lenNStrAry)

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
				"i='%v' numStrAry[i]='%v' Error='%v'",
				i, numStrAry[i], err.Error())
		}

	}

	total, err := BigIntMathAdd{}.AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4],	)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoSeries(...). " +
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddNumStrDtoSeries_03(t *testing.T) {
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

	for i:=0; i < lenNStrAry; i++ {

		numStrDtoAry[i], err = NumStrDto{}.NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]). " +
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
			"Error='%v'", err.Error())
	}

	total, err := BigIntMathAdd{}.AddNumStrDtoSeries(
		numStrDtoAry[0],
		numStrDtoAry[1],
		numStrDtoAry[2],
		numStrDtoAry[3],
		numStrDtoAry[4],	)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddNumStrDtoSeries(...). " +
			"Error='%v' ", err.Error())
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddPair_01(t *testing.T) {
	// n1Str := 123456.789
	b1Str := "123456.789"
	b1Precision := uint(3)

	// n2Str := 987.123456
	b2Str := "987.123456"
	b2Precision := uint(6)

	expectedResultStr := "124443.912456"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_02(t *testing.T) {
	// n1Str := 123456.789
	b1Str := "123456.789"
	b1Precision := uint(3)

	// n2Str := -987.123456
	b2Str := "-987.123456"
	b2Precision := uint(6)

	expectedResultStr := "122469.665544"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_03(t *testing.T) {
	// n1Str := -123456.789
	b1Str := "-123456.789"
	b1Precision := uint(3)

	// n2Str := 987.123456
	b2Str := "987.123456"
	b2Precision := uint(6)

	expectedResultStr := "-122469.665544"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_04(t *testing.T) {
	// n1Str := -123456.789
	b1Str := "-123456.789"
	b1Precision := uint(3)

	// n2Str := -987.123456
	b2Str := "-987.123456"
	b2Precision := uint(6)

	expectedResultStr := "-124443.912456"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_05(t *testing.T) {
	// n1Str := 0
	b1Str := "0"
	b1Precision := uint(0)

	// n2Str := -987.123456
	b2Str := "-987.123456"
	b2Precision := uint(6)

	expectedResultStr := "-987.123456"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_06(t *testing.T) {
	// n1Str := -123456.789
	b1Str := "-123456.789"
	b1Precision := uint(3)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "-123456.789"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_07(t *testing.T) {
	// n1Str := 7
	b1Str := "7"
	b1Precision := uint(0)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "7"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_08(t *testing.T) {
	// n1Str := 0
	b1Str := "0"
	b1Precision := uint(0)

	// n2Str := 7
	b2Str := "7"
	b2Precision := uint(0)

	expectedResultStr := "7"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_09(t *testing.T) {
	// n1Str := 1.7
	b1Str := "1.7"
	b1Precision := uint(1)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "1.7"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_10(t *testing.T) {
	// n1Str := 0
	b1Str := "0"
	b1Precision := uint(0)

	// n2Str := 0
	b2Str := "0"
	b2Precision := uint(0)

	expectedResultStr := "0"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}

	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedResultStr)

	if err != nil {
		t.Errorf("Error returned by expectedBigINum, err := " +
			"BigIntNum{}.NewNumStr(expectedResultStr) " +
			"expectedResultStr='%v' Error='%v",
			expectedResultStr, err.Error())
	}

	maxPrecision := b1Precision

	if b2Precision > b1Precision {
		maxPrecision = b2Precision
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	bPair.MakePrecisionsEqual()

	if maxPrecision != bPair.Big1.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big1 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big1.GetPrecisionUint())
	}

	if maxPrecision != bPair.Big2.GetPrecisionUint() {
		t.Errorf("BigIntPair Error: Expected bPair.Big2 Precision= '%v'. " +
			"Instead, Precision='%v'.",
			maxPrecision, bPair.Big2.GetPrecisionUint())
	}

	result := BigIntMathAdd{}.AddPair(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}

func TestBigIntMathAdd_AddPair_11(t *testing.T) {
	// n1Str := 123456.789
	b1Str := "123456.789"
	b1Precision := uint(3)

	// n2Str := 987.123456
	b2Str := "987.123456"
	b2Precision := uint(6)

	expectedResultStr := "124443,912456"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by b1Num, err := " +
			"BigIntNum{}.NewNumStr(b1Str) b1Str='%v' Error='%v",
			b1Str, err.Error())
	}


	if b1Precision != b1Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b1Precision='%v'. "+
			"Instead, b1Precision='%v.",
			b1Precision, b1Num.GetPrecisionUint())
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
		t.Errorf("Error returned by b1Num.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v'", err.Error())
	}

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	if err != nil {
		t.Errorf("Error returned by b2Num, err := " +
			"BigIntNum{}.NewNumStr(b2Str) b2Str='%v' Error='%v",
			b2Str, err.Error())
	}

	if b2Precision != b2Num.GetPrecisionUint() {
		t.Errorf("Setup Error: Expected b2Precision='%v'. "+
			"Instead, b2Precision='%v.",
			b2Precision, b2Num.GetPrecisionUint())
	}

	bPair := BigIntPair{}.NewBigIntNum(b1Num, b2Num)

	result := BigIntMathAdd{}.AddPair(bPair)

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedResultStr, actualResultStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

