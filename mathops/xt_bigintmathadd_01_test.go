package mathops

import (
	"testing"
	"math/big"
		)

func TestBigIntMathAdd_AddBigInts_01(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := 124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathAdd_AddBigInts_02(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
	t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := 122469.665544
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathAdd_AddBigInts_03(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathAdd_AddBigInts_04(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

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

func TestBigIntMathAdd_AddBigIntNums_01(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := 124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathAdd_AddBigIntNums_02(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := 122469.665544
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathAdd_AddBigIntNums_03(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := 987.123456
	b2Str := "987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathAdd_AddBigIntNums_04(t *testing.T) {

	// n1Str := -123456.789
	b1Str := "-123456789"
	b1Precision := uint(3)
	b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
	}

	b1Num := BigIntNum{}.NewBigInt(b1Big, b1Precision)

	// n2Str := -987.123456
	b2Str := "-987123456"
	b2Precision := uint(6)
	b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

	if !oK {
		t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
	}

	b2Num := BigIntNum{}.NewBigInt(b2Big, b2Precision)

	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1
	biExpectedResult, oK :=  big.NewInt(0).SetString(expectedResultStr, 10)

	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

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

func TestBigIntMathAdd_AddBigIntNums_05(t *testing.T) {

	// n1Str := 123456.789
	b1Str := "123456.789"

	b1Num, err := BigIntNum{}.NewNumStr(b1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(b1Str). " +
			"b1Str='%v' Error='%v'", b1Str, err.Error())
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

	// n2Str := 987.123456
	b2Str := "987.123456"

	b2Num, err := BigIntNum{}.NewNumStr(b2Str)

	// Result := 124443.912456
	expectedResultStr := "124443,912456"


	result := BigIntMathAdd{}.AddBigIntNums(b1Num, b2Num)

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}


func TestBigIntMathAdd_AddBigIntNumArray_01(t *testing.T) {

	numStrs := []string{"45.8",
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

	const lenBigNums = 5

	bNums := make([]BigIntNum, lenBigNums)

	for i:=0; i < lenBigNums; i++ {

		bNums[i], err = BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
				"i='%v' numStrs[i]='%v' Error='%v'",
					i, numStrs[i], err.Error())
		}

	}

  results := BigIntMathAdd{}.AddBigIntNumArray(bNums)

  actualResultNumStr := results.GetNumStr()


	if expectedResultNumStr != actualResultNumStr {
  	t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
			expectedResultNumStr, actualResultNumStr)
	}

	if !expectedBNum.Equal(results) {
		t.Errorf("Error: Expected BigIntNum to equal actual BigIntNum. It Did NOT! " +
			"BigIntTotal='%s'. Instead, BigIntTotal='%s'. ",
			expectedBNum.bigInt.Text(10), results.bigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumArray_02(t *testing.T) {

	numStrs := []string{"-978425.648941",
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

	const lenBigNums = 5

	bNums := make([]BigIntNum, lenBigNums)

	for i:=0; i < lenBigNums; i++ {

		bNums[i], err = BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
				"i='%v' numStrs[i]='%v' Error='%v'",
					i, numStrs[i], err.Error())
		}

	}

  results := BigIntMathAdd{}.AddBigIntNumArray(bNums)

  actualResultNumStr := results.GetNumStr()

  if expectedResultNumStr != actualResultNumStr {
  	t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
			expectedResultNumStr, actualResultNumStr)
	}

	if !expectedBNum.Equal(results) {
		t.Errorf("Error: Expected BigIntNum != actual BigIntNum. " +
			"BigIntTotal='%v'. Instead, BigIntTotal='%v'. ",
			expectedBNum.bigInt.Text(10), results.bigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumArray_03(t *testing.T) {

	numStrs := []string{"-978425.648941",
											"33.12",
											"-804.1",
											"32567",
											"-41.859",
										}

	expectedTotalStr := "-946671,487941"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	const lenBigNums = 5
	var err error
	bNums := make([]BigIntNum, lenBigNums)

	for i:=0; i < lenBigNums; i++ {

		bNums[i], err = BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
				"i='%v' numStrs[i]='%v' Error='%v'",
					i, numStrs[i], err.Error())
		}

	}

	err = bNums[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by bNums[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v'", err.Error())
	}

  results := BigIntMathAdd{}.AddBigIntNumArray(bNums)

  actualResultNumStr := results.GetNumStr()

  if expectedTotalStr != actualResultNumStr {
  	t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
			expectedTotalStr, actualResultNumStr)
	}

  actualNumSeps := results.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
  	t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_01(t *testing.T) {

	var err error

	// multiplier = 0
	addendStr := "5"
	// bNumStrs
	bNumStrs :=  [] string {
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


	addendBiNum, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(bNumStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(bNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(bNumStrs[i]) " +
				"i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
		}

	}

	result := BigIntMathAdd{}.AddBigIntNumOutputToArray(addendBiNum, bINumArray)

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_02(t *testing.T) {

	var err error

	// multiplier = 0
	addendStr := "3.1"
	// bNumStrs
	bNumStrs :=  [] string {
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

	addendBiNum, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(bNumStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(bNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(bNumStrs[i]) " +
				"i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
		}

	}

	result := BigIntMathAdd{}.AddBigIntNumOutputToArray(addendBiNum, bINumArray)

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_03(t *testing.T) {

	var err error

	// multiplier = 0
	addendStr := "5"
	// bNumStrs
	bNumStrs :=  [] string {
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


	addendBiNum, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = addendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by addendBiNum." +
			"SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenArray := len(bNumStrs)
	bINumArray := make([]BigIntNum, lenArray)

	for i:=0; i < lenArray; i++ {

		bINumArray[i], err = 	BigIntNum{}.NewNumStr(bNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(bNumStrs[i]) " +
				"i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
		}

	}

	result := BigIntMathAdd{}.AddBigIntNumOutputToArray(addendBiNum, bINumArray)

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v' Instead, numSeps='%v'",
				expectedNumSeps.String(), actualNumSeps.String())
		}


	}

}


func TestBigIntMathAdd_AddBigIntNumSeries_01(t *testing.T) {
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

	b1, err := BigIntNum{}.NewNumStr(n1Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'.", n1Str, err.Error())
	}
	
	b2, err := BigIntNum{}.NewNumStr(n2Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'.", n2Str, err.Error())
	}
	
	
	b3, err := BigIntNum{}.NewNumStr(n3Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'.", n3Str, err.Error())
	}
	
	b4, err := BigIntNum{}.NewNumStr(n4Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'.", n4Str, err.Error())
	}
	
	b5, err := BigIntNum{}.NewNumStr(n5Str)
	
	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'.", n5Str, err.Error())
	}

	total := BigIntMathAdd{}.AddBigIntNumSeries(b1, b2, b3, b4, b5)

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedResultNumStr, total.bigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumSeries_02(t *testing.T) {
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

	b1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'.", n1Str, err.Error())
	}

	b2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'.", n2Str, err.Error())
	}


	b3, err := BigIntNum{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'.", n3Str, err.Error())
	}

	b4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'.", n4Str, err.Error())
	}

	b5, err := BigIntNum{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'.", n5Str, err.Error())
	}

	total := BigIntMathAdd{}.AddBigIntNumSeries(b1, b2, b3, b4, b5)

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedResultNumStr, total.bigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumSeries_03(t *testing.T) {
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

	b1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'.", n1Str, err.Error())
	}

	b2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'.", n2Str, err.Error())
	}


	b3, err := BigIntNum{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'.", n3Str, err.Error())
	}

	b4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'.", n4Str, err.Error())
	}

	b5, err := BigIntNum{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'.", n5Str, err.Error())
	}

	total := BigIntMathAdd{}.AddBigIntNumSeries(b1, b2, b3, b4, b5)

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedResultNumStr, total.bigInt.Text(10))
	}

}

func TestBigIntMathAdd_AddBigIntNumSeries_04(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedResultNumStr := "158,14788"


	b1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v'.", n1Str, err.Error())
	}


	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = b1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by b1.SetNumericSeparatorsDto(expectedNumSeps) " +
			"Error='%v' ", err.Error())
	}

	b2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'.", n2Str, err.Error())
	}


	b3, err := BigIntNum{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'.", n3Str, err.Error())
	}

	b4, err := BigIntNum{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'.", n4Str, err.Error())
	}

	b5, err := BigIntNum{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'.", n5Str, err.Error())
	}

	total := BigIntMathAdd{}.AddBigIntNumSeries(b1, b2, b3, b4, b5)

	actualNumStr := total.GetNumStr()

	if expectedResultNumStr != actualNumStr {
		t.Errorf("Error: Expected total='%v'. Instead, " +
			"total='%v'. ",
			expectedResultNumStr, actualNumStr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddDecimal_01(t *testing.T) {
	n1Str := "123456.789"

	n2Str := "987.123456"

	expectedNumStr 		:= "124443.912456"
	expectedPrecision := uint(6)
	expectedSign := 1


	dec1, err := Decimal{}.NewNumStr(n1Str)
	
	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)
	
	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}


	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddDecimal_02(t *testing.T) {
	n1Str := "123456.789"

	n2Str := "-987.123456"

	expectedNumStr 		:= "122469.665544"
	expectedPrecision := uint(6)
	expectedSign := 1


	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}


	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddDecimal_03(t *testing.T) {
	n1Str := "-123456.789"

	n2Str := "987.123456"

	expectedNumStr 		:= "-122469.665544"
	expectedPrecision := uint(6)
	expectedSign := -1


	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}


	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddDecimal_04(t *testing.T) {
	n1Str := "123456.789"

	n2Str := "987.123456"

	expectedNumStr 		:= "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1


	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"n1Str='%v' Error='%v' ", n1Str, err.Error())
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
		t.Errorf("Error returned by dec1.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v'", err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v' ", n2Str, err.Error())
	}

	result, err := BigIntMathAdd{}.AddDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimal(dec1, dec2). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
			expectedSign, result.sign)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(),actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddDecimalArray_01(t *testing.T) {

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

	decAry := make([]Decimal, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		dec, err := Decimal{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		decAry[i] = dec

	}

	total, err := BigIntMathAdd{}.AddDecimalArray(decAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalArray(decAry). " +
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

func TestBigIntMathAdd_AddDecimalArray_02(t *testing.T) {

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

	decAry := make([]Decimal, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		dec, err := Decimal{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		decAry[i] = dec

	}

	total, err := BigIntMathAdd{}.AddDecimalArray(decAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalArray(decAry). " +
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


func TestBigIntMathAdd_AddDecimalArray_03(t *testing.T) {

	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry:= len(numStrAry)

	expectedTotalStr := "158,14788"



	decAry := make([]Decimal, lenStrAry)

	for i:=0; i < lenStrAry; i++ {

		dec, err := Decimal{}.NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) " +
					"i='%v' numStrAry[i]='%v' Error='%v' ",i, numStrAry[i], err.Error())
			}

		}

		decAry[i] = dec

	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := decAry[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by decAry[0].SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	total, err := BigIntMathAdd{}.AddDecimalArray(decAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalArray(decAry). " +
			"Error='%v' ", err.Error())
	}

	actualTotalNumStr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumStr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumStr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddDecimalOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// decNumStrs
	decNumStrs :=  [] string {
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


	decAddend, err := Decimal{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(decNumStrs)
	decArray := make([]Decimal, lenArray)


	for i:=0; i < lenArray; i++ {

		decArray[i], err = 	Decimal{}.NewNumStr(decNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) " +
				"i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddDecimalOutputToArray(decAddend, decArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalOutputToArray(" +
			"decAddend, decArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddDecimalOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// decNumStrs
	decNumStrs :=  [] string {
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

	decAddend, err := Decimal{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) " +
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(decNumStrs)
	decsArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decsArray[i], err =	Decimal{}.NewNumStr(decNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) " +
				"i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
		}

	}

	result, err := BigIntMathAdd{}.AddDecimalOutputToArray(decAddend, decsArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalOutputToArray(" +
			"decAddend, decsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddDecimalSeries_01(t *testing.T) {
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
	
	dec2, err := Decimal{}.NewNumStr(n2Str)
	
	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ", 
				n2Str, err.Error())
	}
	
	dec3, err := Decimal{}.NewNumStr(n3Str)
	
	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ", 
				n3Str, err.Error())
	}
	
	dec4, err := Decimal{}.NewNumStr(n4Str)
	
	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ", 
				n4Str, err.Error())
	}
	
	dec5, err := Decimal{}.NewNumStr(n5Str)
	
	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ", 
				n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalSeries(dec1, " +
			"dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
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

func TestBigIntMathAdd_AddDecimalSeries_02(t *testing.T) {
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

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). " +
			"n2Str='%v' Error='%v'. ",
				n2Str, err.Error())
	}

	dec3, err := Decimal{}.NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). " +
			"n3Str='%v' Error='%v'. ",
				n3Str, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). " +
			"n4Str='%v' Error='%v'. ",
				n4Str, err.Error())
	}

	dec5, err := Decimal{}.NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). " +
			"n5Str='%v' Error='%v'. ",
				n5Str, err.Error())
	}

	total, err := BigIntMathAdd{}.AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathAdd{}.AddDecimalSeries(dec1, " +
			"dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

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

	result := BigIntMathAdd{}.addPairNoNumSeps(bPair)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Addition Result='%v'. " +
			"Instead, Result='%v'",
			expectedBigINum.GetNumStr(), result.GetNumStr())
	}

}