package mathops

import "testing"

func TestIntAryMathAdd_Add_01(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "12.851"
	expectedStr := "58.613"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}
}

func TestIntAryMathAdd_Add_02(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "-12.851"
	expectedStr := "32.911"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}
}

func TestIntAryMathAdd_Add_03(t *testing.T) {

	nStr1 := "-45.762"
	nStr2 := "-12.851"
	expectedStr := "-58.613"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}
}

func TestIntAryMathAdd_Add_04(t *testing.T) {

	nStr1 := "0"
	nStr2 := "12.851"
	expectedStr := "12.851"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}
}

func TestIntAryMathAdd_Add_05(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "0"
	expectedStr := "45.762"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}
}

func TestIntAryMathAdd_Add_06(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "12,851"
	expectedStr := "58.613"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, expectedNumSeps). "+
			"nStr1='%v' expectedNumSeps='%v' Error='%v' ",
			nStr1, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}

	actualNumSeps := ia3.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestIntAryMathAdd_Add_07(t *testing.T) {

	nStr1 := "45,762"
	nStr2 := "12.851"
	expectedStr := "58,613"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, expectedNumSeps). "+
			"nStr1='%v' expectedNumSeps='%v' Error='%v' ",
			nStr1, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr())
	}

	actualNumSeps := ia3.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestIntAryMathAdd_RunTotal_01(t *testing.T) {
	total := IntAry{}.NewZero(0)

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). "+
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	ia4, err := IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). "+
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

}

func TestIntAryMathAdd_RunTotal_02(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}.New()
	totalStr := "0"
	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	nStr1 := "5,1"
	nStr2 := "21,452"
	nStr3 := "8"
	nStr4 := "-6,7"
	expectedStr := "27.852"

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAryMathAdd_RunTotal_03(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"
	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"

	expectedStr := "27,852"

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAryMathAdd_AddManyArray_01(t *testing.T) {

	total := IntAry{}.NewZero(0)

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

	iaArray := make([]IntAry, 4)

	var err error

	iaArray[0], err = IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	iaArray[1], err = IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	iaArray[2], err = IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). "+
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	iaArray[3], err = IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). "+
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.AddManyArray(&total, iaArray)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

}

func TestIntAryMathAdd_AddManyArray_02(t *testing.T) {

	var err error

	expectedNumSeps := NumericSeparatorDto{}.New()

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5,1"
	nStr2 := "21,452"
	nStr3 := "8"
	nStr4 := "-6,7"
	expectedStr := "27.852"

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	iaArray := make([]IntAry, 4)

	iaArray[0], err = IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	iaArray[1], err = IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	iaArray[2], err = IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	iaArray[3], err = IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddManyArray(&total, iaArray)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestIntAryMathAdd_AddManyArray_03(t *testing.T) {

	var err error

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27,852"

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	iaArray := make([]IntAry, 4)

	iaArray[0], err = IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	iaArray[1], err = IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	iaArray[2], err = IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	iaArray[3], err = IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddManyArray(&total, iaArray)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestIntAryMathAdd_AddMany_01(t *testing.T) {
	total := IntAry{}.NewZero(0)

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). "+
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	ia4, err := IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). "+
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.AddMany(&total, &ia1, &ia2, &ia3, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}
}

func TestIntAryMathAdd_AddMany_02(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5,1"
	nStr2 := "21,452"
	nStr3 := "8"
	nStr4 := "-6,7"

	expectedStr := "27.852"

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddMany(&total, &ia1, &ia2, &ia3, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAryMathAdd_AddMany_03(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"

	expectedStr := "27,852"

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddMany(&total, &ia1, &ia2, &ia3, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
