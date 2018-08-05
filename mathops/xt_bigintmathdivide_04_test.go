package mathops

import "testing"


func TestBigIntMathDivide_IntAryQuotientMod_01(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555					/						 2.5			=			 5							 0.055
	dividendStr 			:= "12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	quotient, modulo, err :=
		BigIntMathDivide{}.IntAryQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryQuotientMod" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_IntAryQuotientMod_02(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	quotient, modulo, err :=
		BigIntMathDivide{}.IntAryQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryQuotientMod" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_IntAryQuotientMod_03(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "-0.055"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	quotient, modulo, err :=
		BigIntMathDivide{}.IntAryQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryQuotientMod" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_IntAryQuotientMod_04(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	quotient, modulo, err :=
		BigIntMathDivide{}.IntAryQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryQuotientMod" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected quotient NumSeps='%v'. Instead, quotient NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	actualNumSeps = modulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected modulo NumSeps='%v'. Instead, modulo NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryQuotientMod_05(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12,555  	 			/ 				 	 2  			= 		 6							 0,555
	dividendStr 			:= "12,555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0,555"
	maxPrecision			:= uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	dividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr,expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr,expectedNumSeps). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStrWithNumSeps(expectedQuoStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(" +
			"expectedQuoStr, expectedNumSeps). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedModulo, err := BigIntNum{}.NewNumStrWithNumSeps(expectedModuloStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(" +
			"expectedModuloStr, expectedNumSeps). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	quotient, modulo, err :=
		BigIntMathDivide{}.IntAryQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryQuotientMod" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected quotient NumSeps='%v'. Instead, quotient NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	actualNumSeps = modulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected modulo NumSeps='%v'. Instead, modulo NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryQuotientMod_06(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr,expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr,expectedNumSeps). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStrWithNumSeps(expectedQuoStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(" +
			"expectedQuoStr, expectedNumSeps). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	quotient, modulo, err :=
		BigIntMathDivide{}.IntAryQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryQuotientMod" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected quotient NumSeps='%v'. Instead, quotient NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	actualNumSeps = modulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected modulo NumSeps='%v'. Instead, modulo NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_IntAryFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5.022"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_IntAryFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "-0.199123855037834"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_IntAryFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5.022"
	maxPrecision			:= uint(15)


	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryFracQuotient_05(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
	maxPrecision			:= uint(15)


	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryFracQuotient_06(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10,5  				/ 					2 				= 	 5,25

	dividendStr 			:= "10,5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5,25"
	maxPrecision			:= uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	dividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr,expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr,expectedNumSeps). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStrWithNumSeps(expectedQuoStr,expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedQuoStr,expectedNumSeps). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryFracQuotient_07(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
	maxPrecision			:= uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr,expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr,expectedNumSeps). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	expectedQuo, err := BigIntNum{}.NewNumStrWithNumSeps(expectedQuoStr,expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedQuoStr,expectedNumSeps). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.IntAryFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	actualNumSeps := quotient.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_IntAryFracQuotientArray_01(t *testing.T) {

	divisorStr := "2.5"
	maxPrecision := uint(15)

	dividendArrayStr := [] string {
		"10.5",
		"10",
		"11.5",
		"2.5",
		"-12.555",
		"-2.5",
		"12.555",
		"-122.783",
		"-6847.231",
		"-2.5",
		"-10",
		"-10.5",
	}

	expectedArrayStr := [] string {
		"4.2",
		"4",
		"4.6",
		"1",
		"-5.022",
		"-1",
		"5.022",
		"-49.1132",
		"-2738.8924",
		"-1",
		"-4",
		"-4.2",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	lenDividends := len(dividendArrayStr)

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisor='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	dividends := make([]IntAry, lenDividends)
	expectedResults := make([]IntAry, lenDividends)

	for i:=0; i < lenDividends; i++ {

		dividends[i], err = IntAry{}.NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(dividendArrayStr[i]). " +
				"dividendArrayStr[%v]='%v' Error='%v' ",
				i, dividendArrayStr[i], err.Error())
		}

		expectedResults[i], err = IntAry{}.NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedArrayStr[i]). " +
				"expectedArrayStr[%v]='%v' Error='%v' ",
				i, expectedArrayStr[i], err.Error())
		}

	}

	resultArray, err := BigIntMathDivide{}.IntAryFracQuotientArray(dividends, divisor, maxPrecision )

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotientArray"+
			"(dividends, divisor, maxPrecision ). " +
			"divisor='%v' maxPrecision='%v' Error='%v' ",
			divisor.GetNumStr(), maxPrecision, err.Error())
	}

	lenResultArray := len(resultArray)

	if lenDividends != lenResultArray {
		t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
			lenDividends, lenResultArray)
	}

	for k:=0; k < lenDividends; k++ {

		if !resultArray[k].Equal(expectedResults[k]) {
			t.Errorf("Error: Expected Result NOT Equal to Actual Result! " +
				"Expected Value='%v'. Actual Value='%v' k='%v'",
				expectedResults[k].GetNumStr(), resultArray[k].GetNumStr(), k)
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
				expectedNumSeps.String(), actualNumSeps.String())
		}

	}
}

func TestBigIntMathDivide_IntAryFracQuotientArray_02(t *testing.T) {

	divisorStr := "2.5"
	maxPrecision := uint(15)

	dividendArrayStr := [] string {
		"10.5",
		"10",
		"11.5",
		"2.5",
		"-12.555",
		"-2.5",
		"12.555",
		"-122.783",
		"-6847.231",
		"-2.5",
		"-10",
		"-10.5",
	}

	expectedArrayStr := [] string {
		"4,2",
		"4",
		"4,6",
		"1",
		"-5,022",
		"-1",
		"5,022",
		"-49,1132",
		"-2738,8924",
		"-1",
		"-4",
		"-4,2",
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	lenDividends := len(dividendArrayStr)

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisor='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	dividends := make([]IntAry, lenDividends)
	expectedResults := make([]IntAry, lenDividends)

	for i:=0; i < lenDividends; i++ {

		dividends[i], err = IntAry{}.NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(dividendArrayStr[i]). " +
				"dividendArrayStr[%v]='%v' Error='%v' ",
				i, dividendArrayStr[i], err.Error())
		}

		expectedResults[i], err = IntAry{}.NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps" +
				"(expectedArrayStr[i], expectedNumSeps). " +
				"expectedArrayStr[%v]='%v' Error='%v' ",
				i, expectedArrayStr[i], err.Error())
		}

	}

	err = dividends[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by dividends[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	resultArray, err := BigIntMathDivide{}.IntAryFracQuotientArray(dividends, divisor, maxPrecision )

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotientArray"+
			"(dividends, divisor, maxPrecision ). " +
			"divisor='%v' maxPrecision='%v' Error='%v' ",
			divisor.GetNumStr(), maxPrecision, err.Error())
	}

	lenResultArray := len(resultArray)

	if lenDividends != lenResultArray {
		t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
			lenDividends, lenResultArray)
	}

	for k:=0; k < lenDividends; k++ {

		if !resultArray[k].Equal(expectedResults[k]) {
			t.Errorf("Error: Expected Result NOT Equal to Actual Result! " +
				"Expected Value='%v'. Actual Value='%v' k='%v'",
				expectedResults[k].GetNumStr(), resultArray[k].GetNumStr(), k)
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
				expectedNumSeps.String(), actualNumSeps.String())
		}

	}
}

func TestBigIntMathDivide_IntAryFracQuotientArray_03(t *testing.T) {

	divisorStr := "2.5"
	maxPrecision := uint(15)

	dividendArrayStr := [] string {
		"10.5",
		"10",
		"11.5",
		"2.5",
		"-12.555",
		"-2.5",
		"12.555",
		"-122.783",
		"-6847.231",
		"-2.5",
		"-10",
		"-10.5",
	}

	expectedArrayStr := [] string {
		"4.2",
		"4",
		"4.6",
		"1",
		"-5.022",
		"-1",
		"5.022",
		"-49.1132",
		"-2738.8924",
		"-1",
		"-4",
		"-4.2",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	lenDividends := len(dividendArrayStr)

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisor='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	dividends := make([]IntAry, lenDividends)
	expectedResults := make([]IntAry, lenDividends)

	for i:=0; i < lenDividends; i++ {

		dividends[i], err = IntAry{}.NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(dividendArrayStr[i]). " +
				"dividendArrayStr[%v]='%v' Error='%v' ",
				i, dividendArrayStr[i], err.Error())
		}

		expectedResults[i], err = IntAry{}.NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps" +
				"(expectedArrayStr[i], expectedNumSeps). " +
				"expectedArrayStr[%v]='%v' Error='%v' ",
				i, expectedArrayStr[i], err.Error())
		}

	}

	err = dividends[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by dividends[0].SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	resultArray, err := BigIntMathDivide{}.IntAryFracQuotientArray(dividends, divisor, maxPrecision )

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryFracQuotientArray"+
			"(dividends, divisor, maxPrecision ). " +
			"divisor='%v' maxPrecision='%v' Error='%v' ",
			divisor.GetNumStr(), maxPrecision, err.Error())
	}

	lenResultArray := len(resultArray)

	if lenDividends != lenResultArray {
		t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
			lenDividends, lenResultArray)
	}

	for k:=0; k < lenDividends; k++ {

		if !resultArray[k].Equal(expectedResults[k]) {
			t.Errorf("Error: Expected Result NOT Equal to Actual Result! " +
				"Expected Value='%v'. Actual Value='%v' k='%v'",
				expectedResults[k].GetNumStr(), resultArray[k].GetNumStr(), k)
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
				expectedNumSeps.String(), actualNumSeps.String())
		}

	}
}


func TestBigIntMathDivide_IntAryModulo_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.IntAryModulo(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModulo(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_IntAryModulo_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.IntAryModulo(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModulo(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_IntAryModulo_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.IntAryModulo(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModulo(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_IntAryModulo_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
	maxPrecision := uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.IntAryModulo(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModulo(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathDivide_IntAryModulo_05(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12,555					%						 2.5			=			 0,055

	dividendStr := "12,555"
	divisorStr := "2.5"
	expectedModuloStr := "0,055"
	maxPrecision := uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	iaDividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr, expectedNumSeps). " +
			"dividendStr='%v' Error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.IntAryModulo(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModulo(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathDivide_IntAryModulo_06(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	iaDividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr, expectedNumSeps). " +
			"dividendStr='%v' Error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.IntAryModulo(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModulo(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathDivide_IntAryModuloToIntAry_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	iaModulo, err := BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_IntAryModuloToIntAry_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	iaModulo, err := BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_IntAryModuloToIntAry_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	iaModulo, err := BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_IntAryModuloToIntAry_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
	maxPrecision := uint(15)

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	iaDividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	iaModulo, err := BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected iaModulo='%v'. Instead iaModulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := iaModulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathDivide_IntAryModuloToIntAry_05(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12,555					%						 2.5			=			 0,055

	dividendStr := "12,555"
	divisorStr := "2.5"
	expectedModuloStr := "0,055"
	maxPrecision := uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	iaDividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr, expectedNumSeps). " +
			"dividendStr='%v' Error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	iaModulo, err := BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := iaModulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathDivide_IntAryModuloToIntAry_06(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	iaDividend, err := IntAry{}.NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(" +
			"dividendStr, expectedNumSeps). " +
			"dividendStr='%v' Error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	iaModulo, err := BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.IntAryModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := iaModulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

