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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_03(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//    2.5 					/ 				 	12.555		= 	   0							 2.5
	dividendStr 			:= "2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "2.5"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_06(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  - 2.5 					/ 				 	12.555		= 		 0							-2.5
	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "-2.5"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_09(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//    2.5 				  / 				 -12.555		= 		 0							 2.5

	dividendStr 			:= "2.5"
	divisorStr  			:= "-12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "2.5"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_12(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  - 2.5	 					/ 				 -12.555		= 		 0							-2.5

	dividendStr 			:= "-2.5"
	divisorStr  			:= "-12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "-2.5"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_13(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  0	 					   / 				  -12.555			= 		 0							0

	dividendStr 			:= "0"
	divisorStr  			:= "-12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "0"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_14(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  0	 					   / 				   12.555			= 		 0							0

	dividendStr 			:= "0"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "0"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumQuotientMod_15(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  4	 					   / 				    2					= 		 2							0

	dividendStr 			:= "4"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "2"
	expectedModuloStr	:= "0"
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

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumQuotientMod" +
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

func TestBigIntMathDivide_BigIntNumIntQuotient_01(t *testing.T) {
	// Dividend	 divided by		Divisor		=		   Quotient
	// 		 5 				/ 				 2 				= 				 2

	dividendStr 			:= "5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_02(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//     5.25 		/ 				 2  			= 				 2

	dividendStr 			:= "5.25"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_03(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//     2 				/ 				 4				= 				 0

	dividendStr 			:= "2"
	divisorStr  			:= "4"
	expectedQuoStr 		:= "0"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_04(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	// 		-5 				/ 				 2 				= 				-2

	dividendStr 			:= "-5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_05(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//    -5.25     /    			 2  			= 				-2

	dividendStr 			:= "-5.25"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_06(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//    -2 				/ 				 4				= 				 0

	dividendStr 			:= "-2"
	divisorStr  			:= "4"
	expectedQuoStr 		:= "0"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_07(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	// 		 5 				/ 				-2 				=					-2

	dividendStr 			:= "5"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_08(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//     5.25 		/ 				-2 				= 				-2

	dividendStr 			:= "5.25"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_09(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//     2 				/ 				-4				= 				 0

	dividendStr 			:= "2"
	divisorStr  			:= "-4"
	expectedQuoStr 		:= "0"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_10(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	// 		-5 				/ 				-2 				= 				 2

	dividendStr 			:= "-5"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_11(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//    -5.25     /    			-2 				= 				 2

	dividendStr 			:= "-5.25"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "2"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_12(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//    -2 				/ 				-4				= 				 0

	dividendStr 			:= "-2"
	divisorStr  			:= "-4"
	expectedQuoStr 		:= "0"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_13(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//  12.555			/ 			  -2.5			=			    -5

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "-5"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_14(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	// -12.555			/ 			  -2.5			=			     5
	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumIntQuotient_15(t *testing.T) {
	//																					 Integer
	// Dividend	 divided by		Divisor		=		     Quotient
	//  12.555			/ 			  -2				=			    -6
	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-6"


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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumIntQuotient(dividend, divisor)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumIntQuotient" +
			"(dividend, divisor).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}


}

func TestBigIntMathDivide_BigIntNumModulo_01(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555					%						 2.5			=			 0.055
	dividendStr 			:= "12.555"
	divisorStr  			:= "2.5"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_02(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555  	 			% 				 	 2  			= 		 0.555

	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_03(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    2.5 					% 				 	12.555		= 	   2.500

	dividendStr 			:= "2.5"
	divisorStr  			:= "12.555"
	expectedModuloStr	:= "2.5"

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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_04(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//	-12.555 				% 				   2.5 			= 		-0.055

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_05(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  -12.555     		%    			 	 2  			= 		-0.555

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_06(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  - 2.5 					% 				 	12.555		= 		-2.500

	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedModuloStr	:= "-2.5"

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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_07(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	// 	 12.555					% 				 - 2.5			=			 0.055

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_08(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555 				% 				 - 2 				= 		 0.555

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_09(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr 			:= "2.5"
	divisorStr  			:= "-12.555"
	expectedModuloStr	:= "2.5"

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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_10(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	// 	-12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_11(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  -12.555     		%    			 - 2 				= 		-0.555

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2"
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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_12(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  - 2.5	 					% 				 -12.555		= 		-2.5

	dividendStr 			:= "-2.5"
	divisorStr  			:= "-12.555"
	expectedModuloStr	:= "-2.5"

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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_13(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    0	 					% 				  -12.555		  = 		 0

	dividendStr 			:= "0"
	divisorStr  			:= "-12.555"
	expectedModuloStr	:= "0"

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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumModulo_14(t *testing.T) {
	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    0	 					% 				   12.555		  = 		0

	dividendStr 			:= "0"
	divisorStr  			:= "12.555"
	expectedModuloStr	:= "0"

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


	expectedModulo, err := BigIntNum{}.NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedModuloStr). " +
			"expectedModuloStr='%v' Error='%v' ",
			expectedModuloStr, err.Error())
	}

	modulo, err :=
		BigIntMathDivide{}.BigIntNumModulo(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumModulo" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedModulo.Equal(modulo) {
		t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
			expectedModulo.GetNumStr(), modulo.GetNumStr())
	}

}

func TestBigIntMathDivide_BigIntNumFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10    				/ 					2 				= 			5

	dividendStr 			:= "10"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//   11.5  				/         	2.5				=  		4.6

	dividendStr 			:= "11.5"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "4.6"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    2.5					/				 	12.555			=		0.199123855037834

	dividendStr 			:= "2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0.199123855037834"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_05(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5.022"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_06(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  -12.555     	/    				2  			  = 	 -6.2775

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6.2775"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_07(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "-0.199123855037834"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_08(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 12.555				/ 				- 2.5			  =		 -5.022

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "-5.022"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_09(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//   12.555 			/ 				 -2 				=    -6.2775

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-6.2775"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_10(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//    2.5 				/ 				-12.555		  = 	-0.199123855037834

	dividendStr 			:= "2.5"
	divisorStr  			:= "-12.555"
	expectedQuoStr 		:= "-0.199123855037834"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_11(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5.022"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_12(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  -12.555     	/    			 -2 				= 		6.2775

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "6.2775"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_13(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  - 2.5	 				/ 				-12.555		  = 		0.199123855037834

	dividendStr 			:= "-2.5"
	divisorStr  			:= "-12.555"
	expectedQuoStr 		:= "0.199123855037834"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_14(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  -10						/					- 2					=						5

	dividendStr 			:= "-10"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "5"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_15(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  0							/					- 2					=						0

	dividendStr 			:= "0"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "0"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotient_16(t *testing.T) {
	// Dividend		divided by		Divisor			=					Quotient
	//  0							/					  2					=						0

	dividendStr 			:= "0"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "0"
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

	quotient, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_BigIntNumFracQuotientArray_01(t *testing.T) {

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

	lenDividends := len(dividendArrayStr)

	divisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
			"divisor='%v' Error='%v' ",
				divisorStr, err.Error())
	}

	dividends := make([]BigIntNum, lenDividends)
	expectedResults := make([]BigIntNum, lenDividends)

	for i:=0; i < lenDividends; i++ {

		dividends[i], err = BigIntNum{}.NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendArrayStr[i]). " +
				"dividendArrayStr[%v]='%v' Error='%v' ",
					i, dividendArrayStr[i], err.Error())
		}

		expectedResults[i], err = BigIntNum{}.NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedArrayStr[i]). " +
				"expectedArrayStr[%v]='%v' Error='%v' ",
				i, expectedArrayStr[i], err.Error())
		}

	}

	resultArray, err := BigIntMathDivide{}.BigIntNumFracQuotientArray(dividends, divisor, maxPrecision )

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotientArray"+
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

	}
}

func TestBigIntMathDivide_DecimalQuotientMod_01(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555					/						 2.5			=			 5							 0.055
	dividendStr 			:= "12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalQuotientMod" +
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

func TestBigIntMathDivide_DecimalQuotientMod_02(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalQuotientMod" +
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

func TestBigIntMathDivide_DecimalQuotientMod_03(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "-0.055"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalQuotientMod" +
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

func TestBigIntMathDivide_DecimalQuotientMod_04(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalQuotientMod" +
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

func TestBigIntMathDivide_DecimalQuotientMod_05(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	// 	 12.555					/ 				 - 2.5			=			-5							 0.055

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalQuotientMod" +
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

func TestBigIntMathDivide_DecimalQuotientMod_06(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555 				/ 				 - 2 				= 		-6							 0.555

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalQuotientMod" +
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

func TestBigIntMathDivide_DecimalFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_DecimalFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5.022"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_DecimalFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "-0.199123855037834"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_DecimalFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5.022"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.DecimalFracQuotient(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.DecimalFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_DecimalFracQuotientArray_01(t *testing.T) {

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

	lenDividends := len(dividendArrayStr)

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
			"divisor='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	dividends := make([]Decimal, lenDividends)
	expectedResults := make([]Decimal, lenDividends)

	for i:=0; i < lenDividends; i++ {

		dividends[i], err = Decimal{}.NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(dividendArrayStr[i]). " +
				"dividendArrayStr[%v]='%v' Error='%v' ",
				i, dividendArrayStr[i], err.Error())
		}

		expectedResults[i], err = Decimal{}.NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedArrayStr[i]). " +
				"expectedArrayStr[%v]='%v' Error='%v' ",
				i, expectedArrayStr[i], err.Error())
		}

	}

	resultArray, err := BigIntMathDivide{}.DecimalFracQuotientArray(dividends, divisor, maxPrecision )

	if err != nil {
		t.Errorf("Error returned by DecimalFracQuotientArray{}.DecimalFracQuotientArray"+
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

	}
}

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

}

func TestBigIntMathDivide_ModuloToIntAry_01(t *testing.T) {
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

	iaModulo, err := BigIntMathDivide{}.ModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToIntAry_02(t *testing.T) {
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

	iaModulo, err := BigIntMathDivide{}.ModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToIntAry_03(t *testing.T) {
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

	iaModulo, err := BigIntMathDivide{}.ModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToIntAry_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
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

	iaModulo, err := BigIntMathDivide{}.ModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToIntAry(iaDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := iaModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected iaModulo='%v'. Instead iaModulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_INumMgrQuotientMod_01(t *testing.T) {
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

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.INumMgrQuotientMod(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrQuotientMod" +
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

func TestBigIntMathDivide_INumMgrQuotientMod_02(t *testing.T) {
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

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.INumMgrQuotientMod(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrQuotientMod" +
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

func TestBigIntMathDivide_INumMgrQuotientMod_03(t *testing.T) {
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

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.INumMgrQuotientMod(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrQuotientMod" +
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

func TestBigIntMathDivide_INumMgrQuotientMod_04(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
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
		BigIntMathDivide{}.INumMgrQuotientMod(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrQuotientMod" +
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

func TestBigIntMathDivide_INumMgrFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
	maxPrecision			:= uint(15)

	dividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.INumMgrFracQuotient(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_INumMgrFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5.022"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
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
		BigIntMathDivide{}.INumMgrFracQuotient(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_INumMgrFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "-0.199123855037834"
	maxPrecision			:= uint(15)

	dividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.INumMgrFracQuotient(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_INumMgrFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5.022"
	maxPrecision			:= uint(15)

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
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

	quotient, err :=
		BigIntMathDivide{}.INumMgrFracQuotient(&dividend, &divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrFracQuotient" +
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}


func TestBigIntMathDivide_INumMgrFracQuotientArray_01(t *testing.T) {

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

	lenDividends := len(dividendArrayStr)

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
			"divisor='%v' Error='%v' ",
			divisorStr, err.Error())
	}

	dividends := make([]INumMgr, lenDividends)
	expectedResults := make([]BigIntNum, lenDividends)

	for i:=0; i < lenDividends; i++ {

		dec, err := Decimal{}.NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(dividendArrayStr[i]). " +
				"dividendArrayStr[%v]='%v' Error='%v' ",
				i, dividendArrayStr[i], err.Error())
		}

		dividends[i] = &dec

		expectedResults[i], err = BigIntNum{}.NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedArrayStr[i]). " +
				"expectedArrayStr[%v]='%v' Error='%v' ",
				i, expectedArrayStr[i], err.Error())
		}

	}

	resultArray, err := BigIntMathDivide{}.INumMgrFracQuotientArray(dividends, &divisor, maxPrecision )

	if err != nil {
		t.Errorf("Error returned by INumMgrFracQuotientArray{}.INumMgrFracQuotientArray"+
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

		expectedNumStr := expectedResults[k].GetNumStr()
		actualNumStr := resultArray[k].GetNumStr()

		if expectedNumStr != actualNumStr {
			t.Errorf("Error: Expected Value='%v'. Actual Value='%v' k='%v'",
				expectedNumStr, actualNumStr, k)
		}

	}
}

func TestBigIntMathDivide_INumMgrModulo_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decimalDividend, err := Decimal{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.INumMgrModulo(&decimalDividend, &nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrModulo(&decimalDividend, " +
			"&nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_INumMgrModulo_02(t *testing.T) {
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

	bINumDivisor, err := BigIntNum{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.INumMgrModulo(&iaDividend, &bINumDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrModulo(&iaDividend, " +
			"&bINumDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_INumMgrModulo_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	bINumDividend, err := BigIntNum{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decimalDivisor, err := Decimal{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err :=
		BigIntMathDivide{}.INumMgrModulo(&bINumDividend, &decimalDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrModulo(&bINumDividend, " +
			"&decimalDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_INumMgrModulo_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	iaDivisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err :=
		BigIntMathDivide{}.INumMgrModulo(&nDtoDividend, &iaDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.INumMgrModulo(nDtoDividend, " +
			"iaDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
			expectedModuloStr, actualModuloStr)
	}

}


func TestBigIntMathDivide_NumStrQuotientMod_01(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555					/						 2.5			=			 5							 0.055
	dividendStr 			:= "12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)


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
		BigIntMathDivide{}.NumStrQuotientMod(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrQuotientMod" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' Error='%v' ",
			dividendStr, divisorStr, err.Error())
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

func TestBigIntMathDivide_NumStrQuotientMod_02(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

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
		BigIntMathDivide{}.NumStrQuotientMod(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrQuotientMod" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' Error='%v' ",
			dividendStr, divisorStr, err.Error())
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

func TestBigIntMathDivide_NumStrQuotientMod_03(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "-0.055"
	maxPrecision			:= uint(15)

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
		BigIntMathDivide{}.NumStrQuotientMod(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrQuotientMod" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' Error='%v' ",
			dividendStr, divisorStr, err.Error())
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

func TestBigIntMathDivide_NumStrQuotientMod_04(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

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
		BigIntMathDivide{}.NumStrQuotientMod(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrQuotientMod" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' Error='%v' ",
			dividendStr, divisorStr, err.Error())
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

func TestBigIntMathDivide_NumStrQuotientMod_05(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	// 	 12.555					/ 				 - 2.5			=			-5							 0.055

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

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
		BigIntMathDivide{}.NumStrQuotientMod(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrQuotientMod" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' Error='%v' ",
			dividendStr, divisorStr, err.Error())
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

func TestBigIntMathDivide_NumStrQuotientMod_06(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555 				/ 				 - 2 				= 		-6							 0.555

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

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
		BigIntMathDivide{}.NumStrQuotientMod(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrQuotientMod" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' Error='%v' ",
			dividendStr, divisorStr, err.Error())
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

func TestBigIntMathDivide_NumStrFracQuotient_01(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

	dividendStr 			:= "10.5"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "5.25"
	maxPrecision			:= uint(15)

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.NumStrFracQuotient(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrFracQuotient" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
			dividendStr, divisorStr, maxPrecision,err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_NumStrFracQuotient_02(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5.022"
	maxPrecision			:= uint(15)

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.NumStrFracQuotient(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrFracQuotient" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
			dividendStr, divisorStr, maxPrecision,err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_NumStrFracQuotient_03(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "-0.199123855037834"
	maxPrecision			:= uint(15)

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.NumStrFracQuotient(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrFracQuotient" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
			dividendStr, divisorStr, maxPrecision,err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}

func TestBigIntMathDivide_NumStrFracQuotient_04(t *testing.T) {
	// Dividend		divided by		Divisor			=		Quotient
	// 	-12.555 			/ 				 -2.5 			= 	 5.022

	dividendStr 			:= "-12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "5.022"
	maxPrecision			:= uint(15)

	expectedQuo, err := BigIntNum{}.NewNumStr(expectedQuoStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedQuoStr). " +
			"expectedQuoStr='%v' Error='%v' ",
			expectedQuoStr, err.Error())
	}

	quotient, err :=
		BigIntMathDivide{}.NumStrFracQuotient(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrFracQuotient" +
			"(dividendStr, divisorStr, maxPrecision).  " +
			"dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
			dividendStr, divisorStr, maxPrecision,err.Error())
	}

	if !expectedQuo.Equal(quotient) {
		t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
			expectedQuo.GetNumStr(), quotient.GetNumStr())
	}
}


func TestBigIntMathDivide_NumStrFracQuotientArray_01(t *testing.T) {

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

	lenDividends := len(dividendArrayStr)

	resultArray, err :=
		BigIntMathDivide{}.NumStrFracQuotientArray(dividendArrayStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrFracQuotientArray"+
			"(dividendArrayStr, divisorStr, maxPrecision). " +
			"divisor='%v' maxPrecision='%v' Error='%v' ",
			divisorStr, maxPrecision, err.Error())
	}

	lenResultArray := len(resultArray)

	if lenDividends != lenResultArray {
		t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
			lenDividends, lenResultArray)
	}

	for k:=0; k < lenDividends; k++ {

		if expectedArrayStr[k] != resultArray[k].GetNumStr() {
			t.Errorf(	"Expected Value='%v'. Actual Value='%v' k='%v'",
				expectedArrayStr[k], resultArray[k].GetNumStr(), k)
		}

	}
}

func TestBigIntMathDivide_NumStrModulo_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	moduloBINum, err := BigIntMathDivide{}.NumStrModulo(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(dividendStr, " +
			"divisorStr, maxPrecision). " +
			"dividendStr='%v' divisorStr='%v' Error='%v'",
			dividendStr, divisorStr, err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrModulo_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	moduloBINum, err := BigIntMathDivide{}.NumStrModulo(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(dividendStr, " +
			"divisorStr, maxPrecision). " +
			"dividendStr='%v' divisorStr='%v' Error='%v'",
			dividendStr, divisorStr, err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrModulo_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	moduloBINum, err := BigIntMathDivide{}.NumStrModulo(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(dividendStr, " +
			"divisorStr, maxPrecision). " +
			"dividendStr='%v' divisorStr='%v' Error='%v'",
			dividendStr, divisorStr, err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrModulo_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
	maxPrecision := uint(15)

	moduloBINum, err := BigIntMathDivide{}.NumStrModulo(dividendStr, divisorStr, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(dividendStr, " +
			"divisorStr, maxPrecision). " +
			"dividendStr='%v' divisorStr='%v' Error='%v'",
			dividendStr, divisorStr, err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToNumStrDto_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)


	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloNDto, err := BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloNDto.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToNumStrDto_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)


	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloNDto, err := BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloNDto.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToNumStrDto_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)


	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloNDto, err := BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloNDto.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_ModuloToNumStrDto_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloNDto, err := BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.ModuloToNumStrDto(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloNDto.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected moduloNDto='%v'. Instead moduloNDto='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrDtoModulo_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrDtoModulo_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)


	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrDtoModulo_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)


	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrDtoModulo_04(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
	expectedModuloStr := "2.5"
	maxPrecision := uint(15)

	nDtoDividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	nDtoDivisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, nDtoDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoModulo(nDtoDividend, " +
			"nDtoDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_NumStrDtoQuotientMod_01(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  - 2.5 					/ 				 	12.555		= 		 0							-2.5
	dividendStr 			:= "-2.5"
	divisorStr  			:= "12.555"
	expectedQuoStr 		:= "0"
	expectedModuloStr	:= "-2.5"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.NumStrDtoQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoQuotientMod" +
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_02(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr 			:= "12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.NumStrDtoQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoQuotientMod" +
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_03(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "-0.055"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.NumStrDtoQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoQuotientMod" +
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_04(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr 			:= "-12.555"
	divisorStr  			:= "2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "-0.555"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.NumStrDtoQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoQuotientMod" +
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_05(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	// 	 12.555					/ 				 - 2.5			=			-5							 0.055

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2.5"
	expectedQuoStr 		:= "-5"
	expectedModuloStr	:= "0.055"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.NumStrDtoQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoQuotientMod" +
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_06(t *testing.T) {
	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555 				/ 				 - 2 				= 		-6							 0.555

	dividendStr 			:= "12.555"
	divisorStr  			:= "-2"
	expectedQuoStr 		:= "-6"
	expectedModuloStr	:= "0.555"
	maxPrecision			:= uint(15)

	dividend, err := NumStrDto{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	divisor, err := NumStrDto{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(divisorStr). " +
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
		BigIntMathDivide{}.NumStrDtoQuotientMod(dividend, divisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathDivide{}.NumStrDtoQuotientMod" +
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
