package mathops

import (
	"testing"
	"math/big"
)

func TestBigIntMathAdd_AddINumMgr_01(t *testing.T) {
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1


	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10) ")
	}

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&ia1, &dec2). " +
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

func TestBigIntMathAdd_AddINumMgr_02(t *testing.T) {

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

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&ia1, &dec2). " +
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

func TestBigIntMathAdd_AddINumMgr_03(t *testing.T) {

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

	nDto1, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(
		nDto1.GetThisPointer(),
		dec2.GetThisPointer())

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(nDto1.GetThisPointer(), " +
			"dec2.GetThisPointer()). Error='%v' ",
			err.Error())
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

func TestBigIntMathAdd_AddINumMgr_04(t *testing.T) {

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

	nDto1, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&nDto1, &ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&nDto1, &ia2). " +
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


func TestBigIntMathAdd_AddINumMgr_05(t *testing.T) {
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}


	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = ia1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by ia1.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgr(&ia1, &dec2). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'. ",
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
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddINumMgrArray_01(t *testing.T) {

	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	inumMgrAry := make([]INumMgr, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		if i < 2 {
			dec, err := Decimal{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &dec
		} else if i > 1 && i < 4 {
			nDto, err := NumStrDto{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &nDto

		} else {
			// i must be 4
			ia, err := IntAry{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &ia
		}

	}

	total, err := BigIntMathAdd{}.AddINumMgrArray(inumMgrAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrArray(inumMgrAry). " +
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

func TestBigIntMathAdd_AddINumMgrArray_02(t *testing.T) {

	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	inumMgrAry := make([]INumMgr, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		if i < 2 {
			dec, err := Decimal{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &dec
		} else if i > 1 && i < 4 {
			nDto, err := NumStrDto{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &nDto

		} else {
			// i must be 4
			ia, err := IntAry{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &ia
		}

	}

	total, err := BigIntMathAdd{}.AddINumMgrArray(inumMgrAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrArray(inumMgrAry). " +
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


func TestBigIntMathAdd_AddINumMgrArray_03(t *testing.T) {

	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158,14788"


	inumMgrAry := make([]INumMgr, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		if i < 2 {
			dec, err := Decimal{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &dec
		} else if i > 1 && i < 4 {
			nDto, err := NumStrDto{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &nDto

		} else {
			// i must be 4
			ia, err := IntAry{}.NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

			inumMgrAry[i] = &ia
		}

	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := inumMgrAry[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by inumMgrAry[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	total, err := BigIntMathAdd{}.AddINumMgrArray(inumMgrAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrArray(inumMgrAry). " +
			"Error='%v' ", err.Error())
	}

	actualTotalStr := total.GetNumStr()

	if expectedTotalStr != actualTotalStr {
		t.Errorf("Error: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedTotalStr, actualTotalStr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddINumMgrOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// numMgrStrs
	numMgrStrs :=  [] string {
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


	iNumMgrAddend, err := Decimal{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(numMgrStrs)
	iNumMgrArray := make([]INumMgr, lenArray)


	for i:=0; i < lenArray; i++ {

		nDto, err := 	NumStrDto{}.NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numMgrStrs[i]) " +
				"i='%v'  numMgrStrs[i]='%v'  Error='%v'. ", i, numMgrStrs[i], err.Error())
		}

		iNumMgrArray[i] = &nDto
	}

	result, err := BigIntMathAdd{}.AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrOutputToArray(" +
			"&iNumMgrAddend, iNumMgrArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// numMgrStrs
	numMgrStrs :=  [] string {
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

	iNumMgrAddend, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(numMgrStrs)
	iNumMgrsArray := make([]INumMgr, lenArray)

	for i:=0; i < lenArray; i++ {

		ia, err :=	IntAry{}.NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("Error returned by INumMgr{}.NewNumStr(numMgrStrs[i]) " +
				"i='%v'  numMgrStrs[i]='%v'  Error='%v'. ", i, numMgrStrs[i], err.Error())
		}

		iNumMgrsArray[i] = &ia

	}

	result, err := BigIntMathAdd{}.AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrOutputToArray(" +
			"iNumMgrAddend, iNumMgrsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_03(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// numMgrStrs
	numMgrStrs :=  [] string {
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

	iNumMgrAddend, err := Decimal{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iNumMgrAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iNumMgrAddend.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(numMgrStrs)
	iNumMgrArray := make([]INumMgr, lenArray)

	for i:=0; i < lenArray; i++ {

		nDto, err := 	NumStrDto{}.NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(numMgrStrs[i]) " +
				"i='%v'  numMgrStrs[i]='%v'  Error='%v'. ", i, numMgrStrs[i], err.Error())
		}

		iNumMgrArray[i] = &nDto
	}

	result, err := BigIntMathAdd{}.AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrOutputToArray(" +
			"&iNumMgrAddend, iNumMgrArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v' index='%v' ",
				expectedNumSeps.String(), actualNumSeps.String(), j)
		}
	}
}

func TestBigIntMathAdd_AddINumMgrSeries_01(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	nDto2, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by  NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	bigINum4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	dec5, err := Decimal{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrSeries(&dec1, " +
			"&nDto2, &ia3, &bigINum4, &dec5). Error='%v' ", err.Error())
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

func TestBigIntMathAdd_AddINumMgrSeries_02(t *testing.T) {
	n1Str := "-978425.648941"
	n2Str := "33.12"
	n3Str := "-804.1"
	n4Str := "32567"
	n5Str := "-41.859"
	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	bINum5, err := BigIntNum{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddINumMgrSeries(&dec1, &ia2, &nDto3, &dec4, &bINum5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalSeries(&dec1, " +
			"&ia2, &nDto3, &dec4, &bINum5). Error='%v' ", err.Error())
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

func TestBigIntMathAdd_AddINumMgrSeries_03(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158,14788"


	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = dec1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by dec1.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	nDto2, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by  NumStrDto{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	bigINum4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	dec5, err := Decimal{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddINumMgrSeries(&dec1, " +
			"&nDto2, &ia3, &bigINum4, &dec5). Error='%v' ", err.Error())
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

func TestBigIntMathAdd_AddIntAry_01(t *testing.T) {
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

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
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

func TestBigIntMathAdd_AddIntAry_02(t *testing.T) {

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

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
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

func TestBigIntMathAdd_AddIntAry_03(t *testing.T) {

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

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
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

func TestBigIntMathAdd_AddIntAry_04(t *testing.T) {

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

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
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


func TestBigIntMathAdd_AddIntAry_05(t *testing.T) {
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}


	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = ia1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by ia1.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v'", err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAry(ia1, ia2). " +
			"Error='%v' ", err.Error())
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
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddIntAryArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryArray(iaArray). " +
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

func TestBigIntMathAdd_AddIntAryArray_02(t *testing.T) {
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryArray(iaArray). " +
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

func TestBigIntMathAdd_AddIntAryArray_03(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158,14788"

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaArray[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v'", err.Error())
	}

	total, err := BigIntMathAdd{}.AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryArray(iaArray). " +
			"Error='%v' ", err.Error())
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddIntAryOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// iaNumStrs
	iaNumStrs :=  [] string {
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


	iaAddend, err := IntAry{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)


	for i:=0; i < lenArray; i++ {

		iaArray[i], err = 	IntAry{}.NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(iaNumStrs[i]) " +
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryOutputToArray(" +
			"iaAddend, iaArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddIntAryOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// iaNumStrs
	iaNumStrs :=  [] string {
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

	iaAddend, err := IntAry{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(iaNumStrs)
	decsArray := make([]IntAry, lenArray)

	for i:=0; i < lenArray; i++ {

		decsArray[i], err =	IntAry{}.NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(iaNumStrs[i]) " +
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddIntAryOutputToArray(iaAddend, decsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryOutputToArray(" +
			"iaAddend, decsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddIntAryOutputToArray_03(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// iaNumStrs
	iaNumStrs :=  [] string {
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


	iaAddend, err := IntAry{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}


	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iaAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaAddend.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)

	for i:=0; i < lenArray; i++ {

		iaArray[i], err = 	IntAry{}.NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(iaNumStrs[i]) " +
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntAryOutputToArray(" +
			"iaAddend, iaArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
				expectedNumSeps.String(), actualNumSeps.String())
		}
	}
}

func TestBigIntMathAdd_AddIntArySeries_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntArySeries(...). " +
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

func TestBigIntMathAdd_AddIntArySeries_02(t *testing.T) {
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := BigIntNum{}.NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedTotalStr). " +
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := BigIntMathAdd{}.AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntArySeries(...). " +
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

func TestBigIntMathAdd_AddIntArySeries_03(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158,14788"

	iaArray := make([]IntAry, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		ia, err := IntAry{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by IntAry{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia
	}


	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaArray[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	total, err := BigIntMathAdd{}.AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddIntArySeries(...). " +
			"Error='%v' ", err.Error())
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

