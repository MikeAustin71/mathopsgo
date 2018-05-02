package mathops

import "testing"

func TestBigIntMathDivide_BigIntNumQuotientMod_01(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555					/						 2.5			=			 5							 0.055
	dividendStr 			:= "12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}		
				
	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}		
				
}

func TestBigIntMathDivide_BigIntNumQuotientMod_02(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_03(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//    2.5 					/ 				 	12.555		= 	   0							 2.500
	dividendStr 			:= "2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "2.500"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_04(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "-0.055"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_05(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_06(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  - 2.5 					/ 				 	12.555		= 		 0							-2.500
	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "-2.500"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_07(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	// 	 12.555					/ 				 - 2.5			=			-5							 0.055

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_08(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555 				/ 				 - 2 				= 		-6							 0.555

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_09(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//    2.5 				  / 				 -12.555		= 		 0							 2.500

	dividendStr 			:= "2.5"
	divisorStr  			:= "-12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "2.500"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_10(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	// 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5"
	expectedModuloStr	:= "-0.055"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumQuotientMod_11(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 - 2 				= 		 6							-0.555

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
				dividendStr, err.Error())
	}

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
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
			BigIntMathDivide{}.BigIntNumQuotientMod(dividend, divisor, maxPrecision)

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}
