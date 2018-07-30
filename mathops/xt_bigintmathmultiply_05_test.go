package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := IntAry{}.NewNumStrDto(nDtoArray[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrDto(nDtoArray[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ",
				i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoArray(" +
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoArray(" +
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoArray(" +
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoArray(" +
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_05(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664,26376575446779648
	expectedNumStr := "11995826664,26376575446779648"

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoArray(" +
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v' " +
			"Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}


func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"2",
		"4",
		"6",
		"8",
		"10",
		"12",
	}


	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray" +
			"(multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_02(t *testing.T) {

	var err error

	// multiplier = 8
	multiplierStr := "8"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"800.8",
		"-208",
		"31.392",
		"64",
		"42376.984",
		"-39.168",
	}


	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray" +
			"(multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_03(t *testing.T) {

	var err error

	// multiplier = -31.2
	multiplierStr := "-31.2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"-3123.12",
		"811.2",
		"-122.4288",
		"-249.6",
		"-165270.2376",
		"152.7552",
	}


	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray" +
			"(multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_04(t *testing.T) {

	var err error

	// multiplier = 283
	multiplierStr := "283"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"0",
		"-26",
		"0",
		"8",
		"5297.123",
		"0",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"0",
		"-7358",
		"0",
		"2264",
		"1499085.809",
		"0",
	}


	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray" +
			"(multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_05(t *testing.T) {

	var err error

	// multiplier = 0
	multiplierStr := "0"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"5",
		"-26",
		"9",
		"8",
		"5297.123",
		"37",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}


	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray" +
			"(multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoOutputToArray_06(t *testing.T) {

	var err error

	// multiplier = 8
	multiplierStr := "8"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"800,8",
		"-208",
		"31,392",
		"64",
		"42376,984",
		"-39,168",
	}

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoOutputToArray" +
			"(multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'. Index='%v'",
				expectedNumSeps.String(), actualNumSeps.String(), j)
		}

	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5],)


	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoSeries(multiplierNumStrDto," +
			" ...) multiplierNumStrDto='%v'  Error='%v'. ", multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoSeries(multiplierNumStrDto," +
			" ...) multiplierNumStrDto='%v'  Error='%v'. ", multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoSeries(multiplierNumStrDto," +
			" ...) multiplierNumStrDto='%v'  Error='%v'. ", multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoSeries_05(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664,26376575446779648
	expectedNumStr := "11995826664,26376575446779648"

	multiplierNumStrDto, err := NumStrDto{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		nDtoArray[i], err = 	NumStrDto{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrDtoSeries(
		multiplierNumStrDto,
		nDtoArray[0],
		nDtoArray[1],
		nDtoArray[2],
		nDtoArray[3],
		nDtoArray[4],
		nDtoArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrDtoSeries(multiplierNumStrDto," +
			" ...) multiplierNumStrDto='%v'  Error='%v'. ", multiplierNumStrDto.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v' " +
			"Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyPair_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	result := BigIntMathMultiply{}.MultiplyPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyPair_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	result := BigIntMathMultiply{}.MultiplyPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyPair_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	result := BigIntMathMultiply{}.MultiplyPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyPair_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	result := BigIntMathMultiply{}.MultiplyPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyPair_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedNumStr := "0"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	result := BigIntMathMultiply{}.MultiplyPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyPair_06(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875,94572"

	multiplierBiNum, err := BigIntNum{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	multiplicandBiNum, err := BigIntNum{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	bPair := BigIntPair{}.NewBigIntNum(multiplierBiNum, multiplicandBiNum)

	result := BigIntMathMultiply{}.MultiplyPair(bPair)

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}
