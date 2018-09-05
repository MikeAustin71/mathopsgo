package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_01(t *testing.T) {

	baseStr := "35"

	tenExponentStr := "3"

	expectedStr := "35000"

	expectedNumSeps := NumericSeparatorDto{}.New()

	baseBINum, err := BigIntNum{}.NewNumStrWithNumSeps(baseStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(baseStr, expectedNumSeps). "+
			"baseStr='%v', expectedNumSeps='%v' Error='%v'",
			baseStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'
	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	tenExpBINum, err := BigIntNum{}.NewNumStrWithNumSeps(tenExponentStr, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(tenExponentStr, alternateNumSeps). "+
			"tenExponentStr='%v', alternateNumSeps='%v' Error='%v'",
			tenExponentStr, alternateNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, 100)

	actualNumStr := result.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_02(t *testing.T) {

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

	baseBINum, err := BigIntNum{}.NewNumStrWithNumSeps(baseStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(baseStr, expectedNumSeps). "+
			"baseStr='%v', expectedNumSeps='%v' Error='%v'",
			baseStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	tenExpBINum, err := BigIntNum{}.NewNumStrWithNumSeps(tenExponentStr, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(tenExponentStr, alternateNumSeps). "+
			"tenExponentStr='%v', alternateNumSeps='%v' Error='%v'",
			tenExponentStr, alternateNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, 100)

	actualNumStr := result.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTenToPower_03(t *testing.T) {

	baseStr := "35"

	tenExponentStr := "3.9"
	// 10^3.9 = 7943.2823472428150206591828283639

	expectedStr := "278014,88215349852572307139899274"
	// 278014.8821534985257230713989927365
	maxPrecision := uint(26)

	expectedNumSeps := NumericSeparatorDto{}.New()
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'
	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	baseBINum, err := BigIntNum{}.NewNumStrWithNumSeps(baseStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(baseStr, expectedNumSeps). "+
			"baseStr='%v', expectedNumSeps='%v' Error='%v'",
			baseStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	tenExpBINum, err := BigIntNum{}.NewNumStrWithNumSeps(tenExponentStr, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(tenExponentStr, alternateNumSeps). "+
			"tenExponentStr='%v', alternateNumSeps='%v' Error='%v'",
			tenExponentStr, alternateNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyBigIntNumByTenToPower(baseBINum, tenExpBINum, maxPrecision)

	actualNumStr := result.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}
