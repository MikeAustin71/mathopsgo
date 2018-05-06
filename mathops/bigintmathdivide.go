package mathops

import (
	"math/big"
	"fmt"
	"errors"
)

// BigIntMathDivide - This type is comprised of methods used to perform the
// division operation using *big.Int numeric types.
//
// Reference the 'big' math package: https://golang.org/pkg/math/big/
//
type BigIntMathDivide struct {
	Input  BigIntPair
									// BigIntPair.Big1 = Dividend
									// BigIntPair.Big2 = Divisor

	Result BigIntNum // BigIntPair.Big1 = Quotient
									 // BigIntPair.Big2 = Modulo

	ResultFracQuo	BigIntNum		// Quotient expressed with fractional digits
														// to the right of the decimal place.
}

// BigIntNumQuotientMod - Performs a division operation on BigIntNum input
// parameters 'dividend' and 'divisor'.
//
// There are two BigIntNum return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division (Truncate Division).
// See "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of Utrecht
// Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//   						q = D div d = f(D/d) r = D mod d = D − d ·q
//
// 'quotient' is the integer result of dividing the 'dividend' by the 'divisor'
//
// 'modulo' - The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values.
//
// Examples:
// =========
//
// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
//   12.555					/						 2.5			=			 5							 0.055
//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
//    2.5 					/ 				 	12.555		= 	   0							 2.500
//	-12.555 				/ 				   2.5 			= 		-5							-0.055
//  -12.555     		/    			 	 2  			= 		-6							-0.555
//  - 2.5 					/ 				 	12.555		= 		 0							-2.500
// 	 12.555					/ 				 - 2.5			=			-5							 0.055
//   12.555 				/ 				 - 2 				= 		-6							 0.555
//    2.5 				  / 				 -12.555		= 		 0							 2.500
// 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055
//  -12.555     		/    			 - 2 				= 		 6							-0.555
//  - 2.5	 					/ 				 -12.555		= 		 0							-2.500
//
func (bIDivide BigIntMathDivide) BigIntNumQuotientMod(
								dividend,
									divisor BigIntNum,
										maxPrecision uint) (quotient, modulo BigIntNum, err error) {


	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)
	bPair.MakePrecisionsEqual()

	if divisor.IsZero() {
		quotient = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		modulo = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		ePrefix := "BigIntMathDivide.BigIntNumQuotientMod() "
		err = fmt.Errorf(ePrefix + "Error: Attempted to divide by zero!")
		return quotient, modulo, err
	}

	quotientBigI, moduloBigI := big.NewInt(0).QuoRem(
					bPair.Big1.bigInt, bPair.Big2.bigInt, big.NewInt(0))

	quotient = BigIntNum{}.NewBigInt(quotientBigI, 0)

	modulo = BigIntNum{}.NewBigInt(moduloBigI, bPair.Big2.GetPrecisionUint())

	if modulo.precision > maxPrecision {
		modulo.RoundToDecPlace(maxPrecision)
	}

	return quotient, modulo, nil
}

// BigIntNumIntQuotient - Performs integer division on two BigIntNum types passed as
// input parameters, 'dividend' and 'divisor'.
//
// The division operation performed by this method is T-Division or truncated division.
// See "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of Utrecht
// Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
//
// After completing the division operation, an integer quotient of type BigIntNum
// is returned.
//
// Examples:
// =========
//																				Return Value
//  Divisor	divided by	Dividend		=		Integer Quotient
// 		 5 				/ 				 2 				= 				 2
//     5.25 		/ 				 2  			= 				 2
//     2 				/ 				 4				= 				 0
// 		-5 				/ 				 2 				= 				-2
//    -5.25     /    			 2  			= 				-2
//    -2 				/ 				 4				= 				 0
// 		 5 				/ 				-2 				=					-2
//     5.25 		/ 				-2 				= 				-2
//     2 				/ 				-4				= 				 0
// 		-5 				/ 				-2 				= 				 2
//    -5.25     /    			-2 				= 				 2
//    -2 				/ 				-4				= 				 0
//     12.555		/ 			  -2.5			=			    -5
//    -12.555		/ 			  -2.5			=			     5
//     12.555		/ 			  -2				=			    -6
//
func (bIDivide BigIntMathDivide) BigIntNumIntQuotient(
				dividend,
					divisor BigIntNum) (intQuotient BigIntNum, err error) {


  if divisor.IsZero() {
		intQuotient = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		ePrefix := "BigIntMathDivide.BigIntNumIntQuotient() "
		err = fmt.Errorf(ePrefix + "Error: Attempted to divide by zero!")
		return intQuotient, err
	}

	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)
	bPair.MakePrecisionsEqual()

	quotient := big.NewInt(0).Quo(bPair.Big1.bigInt, bPair.Big2.bigInt)

	intQuotient = BigIntNum{}.NewBigInt(quotient, 0)

	return intQuotient, nil
}


// BigIntNumModulo - Performs a modulo operation on BigIntNum input
// parameters 'dividend' and 'divisor'.
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one BigIntNum value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate Division). See
// "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of
// Utrecht Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//   						q = D div d = f(D/d)
// 							r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values.
//
// Examples:
// =========
//
// Dividend			  mod by			Divisor			=			Modulo/Remainder
// --------				------			-------						----------------
//   12.555					%						 2.5			=			 0.055
//   12.555  	 			% 				 	 2  			= 		 0.555
//    2.5 					% 				 	12.555		= 	   2.500
//	-12.555 				% 				   2.5 			= 		-0.055
//  -12.555     		%    			 	 2  			= 		-0.555
//  - 2.5 					% 				 	12.555		= 		-2.500
// 	 12.555					% 				 - 2.5			=			 0.055
//   12.555 				% 				 - 2 				= 		 0.555
//    2.5 				  % 				 -12.555		= 		 2.500
// 	-12.555 				% 				 - 2.5 			= 		-0.055
//  -12.555     		%    			 - 2 				= 		-0.555
//  - 2.5	 					% 				 -12.555		= 		-2.500
//
func (bIDivide BigIntMathDivide) BigIntNumModulo(
						dividend,
							divisor BigIntNum,
								maxPrecision uint) (modulo BigIntNum, err error) {


	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)
	bPair.MakePrecisionsEqual()

	if divisor.IsZero() {
		modulo = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		ePrefix := "BigIntMathDivide.BigIntNumQuotientMod() "
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	moduloBigI := big.NewInt(0).Rem(bPair.Big1.bigInt, bPair.Big2.bigInt)

	modulo = BigIntNum{}.NewBigInt(moduloBigI, bPair.Big2.GetPrecisionUint())

	if modulo.precision > maxPrecision {
		modulo.RoundToDecPlace(maxPrecision)
	}

	return modulo, nil
}

// BigIntNumFracQuotient - Performs a division operation on BigIntNum input
// parameters 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type representing the
// result of the division operation expressed as integer and fractional digits
// as appropriate. Remember that the BigIntNum type specifies 'precision'. Precision
// is defined as the number of fractional digits to the right of the decimal place.
//
// Examples:
// =========
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//																				   Quotient
//  Dividend		divided by	Divisor		=		BigIntNum Integer 	Precision	 Result
//  -------- 	  ----------	--------				-----------------	  ---------	 ------
// 	 10.5  				/ 				2 				= 			525  							  2  			 5.25
// 	 10    				/ 				2 				= 			5	  							  0  			 5
//   11.5  				/         2.5				=  			46								  1				 4.6
//    2.5					/				 12.555			=				199123855037834	   15				 0.199123855037834
//	-12.555 			/ 				2.5 			= 		 -5022							  3				-5.022
//  -12.555     	/    			2  			  = 		 -62775							  4				-6.2775
//  - 2.5 				/ 			 12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	 12.555				/ 			- 2.5			  =			 -5022								3				-5.022
//   12.555 			/ 			- 2 				= 		 -62775								4				-6.2775
//    2.5 				/ 			-12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	-12.555 			/ 			- 2.5 			= 			5022								3				 5.022
//  -12.555     	/    		- 2 				= 		  62775								4				 6.2775
//  - 2.5	 				/ 			-12.555		  = 		  199123855037834	   15				 0.199123855037834
//  -10						/				- 2					=				5														 5
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits.
//
func (bIDivide BigIntMathDivide) BigIntNumFracQuotient(
										dividend,
											divisor BigIntNum,
												maxPrecision uint) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.BigIntNumFracQuotient() "

	if divisor.IsZero() {
		fracQuotient = BigIntNum{}.New()
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuotient, err
	}

	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)
	bPair.MakePrecisionsEqual()

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(int(maxPrecision))


	fracQuotient, errx :=
				BigIntNum{}.NewNumStr(numStr)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(numStr, maxPrecision). " +
				"numStr='%v' maxPrecision='%v' Error='%v'",
				numStr, maxPrecision, err.Error())

		return fracQuotient, err
	}

	fracQuotient.TrimTrailingFracZeros()

	return fracQuotient, nil
}

// BigIntNumFracQuotientArray - Performs a division operation on BigIntNum input
// parameters 'dividends' and 'divisor'. 'dividends' is an array of BigIntNum types
// and the division is operation is performed on each element of the array using a
// single 'divisor'.
//
// The resulting quotients are returned as an array BigIntNum types representing the
// result of each division operation expressed as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'. Precision is defined
// as the number of fractional digits to the right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotients. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//    	                                       Returned
//  	Dividend		divided by	Divisor		=	       Array						=  	 Result
//  	-------- 	  ----------	--------				-----------------	   		---------
//	 	 10.5  				 / 				2.5 			= 		 fracQuoArray[0]		=   	 4.2
//	 	 10    				 / 				2.5 			= 		 fracQuoArray[1]		=	 	   4
//	   11.5  				 /        2.5				=  		 fracQuoArray[2]		=		   4.6
//	    2.5					 /				2.5			  =			 fracQuoArray[3]    =      1
//		-12.555 			 / 				2.5 			= 		 fracQuoArray[4]    =  -   5.022
//	  - 2.5 				 / 			  2.5		    = 		 fracQuoArray[5]    =  -   1
//	   12.555 			 / 			  2.5 			= 		 fracQuoArray[6]    =      5.022
//	 -122.783 			 / 			  2.5 			= 		 fracQuoArray[7]    =  -  49.1132
//	-6847.231   	   /    	  2.5 			= 		 fracQuoArray[8]    =  -2738.8924
//	  - 2.5	 				 / 			  2.5		    = 		 fracQuoArray[9]    =  -   1
//	  -10						 /			  2.5				=			 fracQuoArray[10]   =  -   4
//	  -10.5					 /			  2.5				=			 fracQuoArray[11]   =  -   4.2
//
func (bIDivide BigIntMathDivide) BigIntNumFracQuotientArray(
										dividends []BigIntNum,
											divisor BigIntNum,
												maxPrecision uint) (fracQuoArray [] BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.BigIntNumFracQuotientArray() "

	if divisor.IsZero() {
		fracQuoArray = []BigIntNum{}
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuoArray, err
	}

	lenAry := len(dividends)

	if lenAry == 0 {
		fracQuoArray = []BigIntNum{}
		err = errors.New(ePrefix + "Error: Input Parameter 'dividends' is an EMPTY Array!")
		return fracQuoArray, err
	}

	fracQuoArray = make([]BigIntNum, lenAry, lenAry+20)
	var errx error

	for i:=0; i < lenAry; i++ {

		bPair := BigIntPair{}.NewBigIntNum(dividends[i], divisor)
		bPair.MakePrecisionsEqual()

		rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
		rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
		rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
		numStr := rQuotient.FloatString(int(maxPrecision))

		fracQuoArray[i], errx =
					BigIntNum{}.NewNumStr(numStr)

		if errx != nil {
			fracQuoArray = []BigIntNum{}
			err =	fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(numStr, maxPrecision). " +
					"numStr='%v' maxPrecision='%v' Error='%v'",
					numStr, maxPrecision, errx.Error())
			return fracQuoArray, err
		}

		fracQuoArray[i].TrimTrailingFracZeros()
	}

	return fracQuoArray, nil
}

// DecimalQuotientMod - Performs a division operation on Decimal type input
// parameters, 'dividend' and 'divisor'.
//
// There are two BigIntNum Type return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division (Truncate Division).
// See "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of Utrecht
// Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//   						q = D div d = f(D/d) r = D mod d = D − d ·q
//
// 'quotient' is the integer result of dividing the 'dividend' by the 'divisor'
//
// 'modulo' - The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values. Therefore, the user is advised to set a relevant value
// for 'maxPrecision'.
//
// Examples:
// =========
//
// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
//   12.555					/						 2.5			=			 5							 0.055
//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
//    2.5 					/ 				 	12.555		= 	   0							 2.500
//	-12.555 				/ 				   2.5 			= 		-5							-0.055
//  -12.555     		/    			 	 2  			= 		-6							-0.555
//  - 2.5 					/ 				 	12.555		= 		 0							-2.500
// 	 12.555					/ 				 - 2.5			=			-5							 0.055
//   12.555 				/ 				 - 2 				= 		-6							 0.555
//    2.5 				  / 				 -12.555		= 		 0							 2.500
// 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055
//  -12.555     		/    			 - 2 				= 		 6							-0.555
//  - 2.5	 					/ 				 -12.555		= 		 0							-2.500
//
func (bIDivide BigIntMathDivide) DecimalQuotientMod(
								dividend,
									divisor Decimal,
										maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.DecimalQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	bINumDividend, errx := dividend.GetBigIntNum()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by dividend.GetBigIntNum(). " +
			"dividend='%v' Error='%v'",
			dividend.GetNumStr(), errx.Error())
		return quotient, modulo, err
	}


	bINumDivisor, errx := divisor.GetBigIntNum()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by divisor.GetBigIntNum(). " +
			"divisor='%v' Error='%v'",
			divisor.GetNumStr(), errx.Error())

		return quotient, modulo, err
	}


	quotient, modulo, errx =
			BigIntMathDivide{}.BigIntNumQuotientMod(bINumDividend, bINumDivisor, maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"BigIntNumQuotientMod(bINumDividend, bINumDivisor, maxPrecision). " +
			"Error='%v'",	errx.Error())

		quotient = BigIntNum{}.New()
		modulo = BigIntNum{}.New()

		return quotient, modulo, err
	}

	err = nil

	return quotient, modulo, err
}

// DecimalFracQuotient - Performs a division operation on Decimal Type input
// parameters 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type representing the
// result of the division operation expressed as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'. Precision is defined
// as the number of fractional digits to the right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
//
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//																				     Quotient
//  Dividend		divided by	Divisor		=		  BigIntNum Integer 	Precision	 Result
//  -------- 	  ----------	--------				-----------------	  ---------	 ------
// 	 10.5  				/ 				2 				= 			525  							  2  			 5.25
// 	 10    				/ 				2 				= 			5	  							  0  			 5
//   11.5  				/         2.5				=  			46								  1				 4.6
//    2.5					/				 12.555			=				199123855037834	   15				 0.199123855037834
//	-12.555 			/ 				2.5 			= 		 -5022							  3				-5.022
//  -12.555     	/    			2  			  = 		 -62775							  4				-6.2775
//  - 2.5 				/ 			 12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	 12.555				/ 			- 2.5			  =			 -5022								3				-5.022
//   12.555 			/ 			- 2 				= 		 -62775								4				-6.2775
//    2.5 				/ 			-12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	-12.555 			/ 			- 2.5 			= 			5022								3				 5.022
//  -12.555     	/    		- 2 				= 		  62775								4				 6.2775
//  - 2.5	 				/ 			-12.555		  = 		  199123855037834	   15				 0.199123855037834
//  -10						/				- 2					=				5														 5
//
func (bIDivide BigIntMathDivide) DecimalFracQuotient(
															dividend,
																divisor Decimal,
																	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.DecimalFracQuotient() "

	if divisor.IsZero() {
		fracQuotient = BigIntNum{}.New()
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuotient, err
	}

	bPair, errx := BigIntPair{}.NewDecimal(dividend, divisor)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewDecimal(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
				dividend.GetNumStr(), divisor.GetNumStr(), err.Error())

		return fracQuotient, err
	}

	bPair.MakePrecisionsEqual()

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(int(maxPrecision))

	fracQuotient, errx =
		BigIntNum{}.NewNumStr(numStr)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(numStr, maxPrecision). " +
			"numStr='%v' maxPrecision='%v' Error='%v'",
			numStr, maxPrecision, err.Error())

		return fracQuotient, err
	}

	fracQuotient.TrimTrailingFracZeros()

	return fracQuotient, nil
}

// DecimalFracQuotientArray - Performs a division operation on Decimal input
// parameters 'dividends' and 'divisor'. 'dividends' is an array of Decimal
// Types. The division operation is performed on each element of the 'dividends'
// array using a single 'divisor'.
//
// The resulting quotients are returned as an array of Decimal Types. The values
// represent result of each division operation expressed as integer and fractional
// digits.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Precision is defined as the number of numeric
// digits to the right of the decimal point. Be advised that this method is
// capable of calculating quotients with very long strings of fractional digits.
// Therefore the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//    	                                       Returned
//  	Dividend		divided by	Divisor		=	       Array						=  	 Result
//  	-------- 	  ----------	--------				-----------------	   		---------
//	 	 10.5  				 / 				2.5 			= 		 fracQuoArray[0]		=   	 4.2
//	 	 10    				 / 				2.5 			= 		 fracQuoArray[1]		=	 	   4
//	   11.5  				 /        2.5				=  		 fracQuoArray[2]		=		   4.6
//	    2.5					 /				2.5			  =			 fracQuoArray[3]    =      1
//		-12.555 			 / 				2.5 			= 		 fracQuoArray[4]    =  -   5.022
//	  - 2.5 				 / 			  2.5		    = 		 fracQuoArray[5]    =  -   1
//	   12.555 			 / 			  2.5 			= 		 fracQuoArray[6]    =      5.022
//	 -122.783 			 / 			  2.5 			= 		 fracQuoArray[7]    =  -  49.1132
//	-6847.231   	   /    	  2.5 			= 		 fracQuoArray[8]    =  -2738.8924
//	  - 2.5	 				 / 			  2.5		    = 		 fracQuoArray[9]    =  -   1
//	  -10						 /			  2.5				=			 fracQuoArray[10]   =  -   4
//	  -10.5					 /			  2.5				=			 fracQuoArray[11]   =  -   4.2
//
func (bIDivide BigIntMathDivide) DecimalFracQuotientArray(
											dividends []Decimal,
												divisor Decimal,
													maxPrecision uint) (fracQuoArray [] Decimal, err error) {

	ePrefix := "BigIntMathDivide.DecimalFracQuotientArray() "

	if divisor.IsZero() {
		fracQuoArray = []Decimal{}
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuoArray, err
	}

	lenAry := len(dividends)

	if lenAry == 0 {
		fracQuoArray = []Decimal{}
		err = errors.New(ePrefix + "Error: Input Parameter 'dividends' is an EMPTY Array!")
		return fracQuoArray, err
	}

	fracQuoArray = make([]Decimal, lenAry, lenAry+20)

	for i:=0; i < lenAry; i++ {

		bPair, errx := BigIntPair{}.NewDecimal(dividends[i], divisor)

		if errx != nil {
			fracQuoArray = []Decimal{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewDecimal(dividends[i], divisor). " +
				"dividends[%v]='%v' divisor='%v' Error='%v'",
				i, dividends[i].GetNumStr(), divisor.GetNumStr(), errx.Error())
			return fracQuoArray, err
		}

		bPair.MakePrecisionsEqual()

		rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
		rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
		rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
		numStr := rQuotient.FloatString(int(maxPrecision))

		bINum, errx := BigIntNum{}.NewNumStr(numStr)

		if errx != nil {
			fracQuoArray = []Decimal{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(numStr). " +
				"numStr='%v' maxPrecision='%v' Error='%v'",
				numStr, maxPrecision, errx.Error())
			return fracQuoArray, err
		}

		bINum.TrimTrailingFracZeros()

		fracQuoArray[i], errx = bINum.GetDecimal()

		if errx != nil {
			fracQuoArray = []Decimal{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by bINum.GetDecimal(). Error='%v'",
				errx.Error())
			return fracQuoArray, err
		}

	}

	return fracQuoArray, nil
}

// IntAryQuotientMod - Performs a division operation on IntAry type input
// parameters, 'dividend' and 'divisor'.
//
// There are two BigIntNum Type return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division (Truncate Division).
// See "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of Utrecht
// Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//   						q = D div d = f(D/d) r = D mod d = D − d ·q
//
// 'quotient' is the integer result of dividing the 'dividend' by the 'divisor'
//
// 'modulo' - The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values.
//
// Examples:
// =========
//
// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
//   12.555					/						 2.5			=			 5							 0.055
//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
//    2.5 					/ 				 	12.555		= 	   0							 2.500
//	-12.555 				/ 				   2.5 			= 		-5							-0.055
//  -12.555     		/    			 	 2  			= 		-6							-0.555
//  - 2.5 					/ 				 	12.555		= 		 0							-2.500
// 	 12.555					/ 				 - 2.5			=			-5							 0.055
//   12.555 				/ 				 - 2 				= 		-6							 0.555
//    2.5 				  / 				 -12.555		= 		 0							 2.500
// 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055
//  -12.555     		/    			 - 2 				= 		 6							-0.555
//  - 2.5	 					/ 				 -12.555		= 		 0							-2.500
//
func (bIDivide BigIntMathDivide) IntAryQuotientMod(
																dividend,
																	divisor IntAry,
																		maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.IntAryQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	bINumDividend, errx := dividend.GetBigIntNum()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by dividend.GetBigIntNum(). " +
			"dividend='%v' Error='%v'",
			dividend.GetNumStr(), errx.Error())
		return quotient, modulo, err
	}


	bINumDivisor, errx := divisor.GetBigIntNum()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by divisor.GetBigIntNum(). " +
			"divisor='%v' Error='%v'",
			divisor.GetNumStr(), errx.Error())

		return quotient, modulo, err
	}

	quotient, modulo, errx =
		BigIntMathDivide{}.BigIntNumQuotientMod(bINumDividend, bINumDivisor, maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"BigIntNumQuotientMod(bINumDividend, bINumDivisor, maxPrecision). " +
			"Error='%v'",	errx.Error())

		quotient = BigIntNum{}.New()
		modulo = BigIntNum{}.New()

		return quotient, modulo, err
	}

	err = nil

	return quotient, modulo, err
}

// IntAryFracQuotient - Performs a division operation on IntAry Type input
// parameters 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type representing the
// result of the division operation expressed as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'. Precision is defined
// as the number of fractional digits to the right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
//
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//																				      Quotient
//  Dividend		divided by	Divisor		=		  BigIntNum Integer 	Precision	 Result
//  -------- 	  ----------	--------				-----------------	  ---------	 ------
// 	 10.5  				/ 				2 				= 			525  							  2  			 5.25
// 	 10    				/ 				2 				= 			5	  							  0  			 5
//   11.5  				/         2.5				=  			46								  1				 4.6
//    2.5					/				 12.555			=				199123855037834	   15				 0.199123855037834
//	-12.555 			/ 				2.5 			= 		 -5022							  3				-5.022
//  -12.555     	/    			2  			  = 		 -62775							  4				-6.2775
//  - 2.5 				/ 			 12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	 12.555				/ 			- 2.5			  =			 -5022								3				-5.022
//   12.555 			/ 			- 2 				= 		 -62775								4				-6.2775
//    2.5 				/ 			-12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	-12.555 			/ 			- 2.5 			= 			5022								3				 5.022
//  -12.555     	/    		- 2 				= 		  62775								4				 6.2775
//  - 2.5	 				/ 			-12.555		  = 		  199123855037834	   15				 0.199123855037834
//  -10						/				- 2					=				5														 5
//
func (bIDivide BigIntMathDivide) IntAryFracQuotient(
														dividend,
															divisor IntAry,
																maxPrecision uint) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.IntAryFracQuotient() "

	if divisor.IsZero() {
		fracQuotient = BigIntNum{}.New()
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuotient, err
	}

	bPair, errx := BigIntPair{}.NewIntAry(dividend, divisor)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewIntAry(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(), err.Error())

		return fracQuotient, err
	}

	bPair.MakePrecisionsEqual()

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(int(maxPrecision))

	fracQuotient, errx =
		BigIntNum{}.NewNumStr(numStr)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(numStr, maxPrecision). " +
			"numStr='%v' maxPrecision='%v' Error='%v'",
			numStr, maxPrecision, err.Error())

		return fracQuotient, err
	}

	fracQuotient.TrimTrailingFracZeros()

	return fracQuotient, nil
}

// IntAryFracQuotientArray - Performs a division operation on IntAry input
// parameters 'dividends' and 'divisor'. 'dividends' is an array of IntAry
// Types. The division operation is performed on each element of the 'dividends'
// array using a single 'divisor'.
//
// The resulting quotients are returned as an array of IntAry Types. The values
// represent result of each division operation expressed as integer and fractional
// digits.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Precision is defined as the number of numeric
// digits to the right of the decimal point. Be advised that this method is
// capable of calculating quotients with very long strings of fractional digits.
// Therefore the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//    	                                       Returned
//  	Dividend		divided by	Divisor		=	       Array						=  	 Result
//  	-------- 	  ----------	--------				-----------------	   		---------
//	 	 10.5  				 / 				2.5 			= 		 fracQuoArray[0]		=   	 4.2
//	 	 10    				 / 				2.5 			= 		 fracQuoArray[1]		=	 	   4
//	   11.5  				 /        2.5				=  		 fracQuoArray[2]		=		   4.6
//	    2.5					 /				2.5			  =			 fracQuoArray[3]    =      1
//		-12.555 			 / 				2.5 			= 		 fracQuoArray[4]    =  -   5.022
//	  - 2.5 				 / 			  2.5		    = 		 fracQuoArray[5]    =  -   1
//	   12.555 			 / 			  2.5 			= 		 fracQuoArray[6]    =      5.022
//	 -122.783 			 / 			  2.5 			= 		 fracQuoArray[7]    =  -  49.1132
//	-6847.231   	   /    	  2.5 			= 		 fracQuoArray[8]    =  -2738.8924
//	  - 2.5	 				 / 			  2.5		    = 		 fracQuoArray[9]    =  -   1
//	  -10						 /			  2.5				=			 fracQuoArray[10]   =  -   4
//	  -10.5					 /			  2.5				=			 fracQuoArray[11]   =  -   4.2
//
func (bIDivide BigIntMathDivide) IntAryFracQuotientArray(
	dividends []IntAry,
	divisor IntAry,
	maxPrecision uint) (fracQuoArray [] IntAry, err error) {

	ePrefix := "BigIntMathDivide.IntAryFracQuotientArray() "

	if divisor.IsZero() {
		fracQuoArray = []IntAry{}
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuoArray, err
	}

	lenAry := len(dividends)

	if lenAry == 0 {
		fracQuoArray = []IntAry{}
		err = errors.New(ePrefix + "Error: Input Parameter 'dividends' is an EMPTY Array!")
		return fracQuoArray, err
	}

	fracQuoArray = make([]IntAry, lenAry, lenAry+20)

	for i:=0; i < lenAry; i++ {

		bPair, errx := BigIntPair{}.NewIntAry(dividends[i], divisor)

		if errx != nil {
			fracQuoArray = []IntAry{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewIntAry(dividends[i], divisor). " +
				"dividends[%v]='%v' divisor='%v' Error='%v'",
				i, dividends[i].GetNumStr(), divisor.GetNumStr(), errx.Error())
			return fracQuoArray, err
		}

		bPair.MakePrecisionsEqual()

		rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
		rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
		rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
		numStr := rQuotient.FloatString(int(maxPrecision))

		bINum, errx := BigIntNum{}.NewNumStr(numStr)

		if errx != nil {
			fracQuoArray = []IntAry{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(numStr). " +
				"numStr='%v' maxPrecision='%v' Error='%v'",
				numStr, maxPrecision, errx.Error())
			return fracQuoArray, err
		}

		bINum.TrimTrailingFracZeros()

		fracQuoArray[i], errx = bINum.GetIntAry()

		if errx != nil {
			fracQuoArray = []IntAry{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by bINum.GetIntAry(). Error='%v'",
				errx.Error())
			return fracQuoArray, err
		}

	}

	return fracQuoArray, nil
}

// INumMgrQuotientMod - Performs a division operation on types implementing the INumMgr
// interface. Input parameters, 'dividend' and 'divisor' must therefore implement the
// INumMgr interface.
//
// There are two BigIntNum Type return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division (Truncate Division).
// See "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of Utrecht
// Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//   						q = D div d = f(D/d) r = D mod d = D − d ·q
//
// 'quotient' is the integer result of dividing the 'dividend' by the 'divisor'
//
// 'modulo' - The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values. Therefore, the user is advised to set a relevant value
// for 'maxPrecision'.
//
// Examples:
// =========
//
// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
//   12.555					/						 2.5			=			 5							 0.055
//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
//    2.5 					/ 				 	12.555		= 	   0							 2.500
//	-12.555 				/ 				   2.5 			= 		-5							-0.055
//  -12.555     		/    			 	 2  			= 		-6							-0.555
//  - 2.5 					/ 				 	12.555		= 		 0							-2.500
// 	 12.555					/ 				 - 2.5			=			-5							 0.055
//   12.555 				/ 				 - 2 				= 		-6							 0.555
//    2.5 				  / 				 -12.555		= 		 0							 2.500
// 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055
//  -12.555     		/    			 - 2 				= 		 6							-0.555
//  - 2.5	 					/ 				 -12.555		= 		 0							-2.500
//
func (bIDivide BigIntMathDivide) INumMgrQuotientMod(
											dividend,
												divisor INumMgr,
													maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.INumMgrQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	bIntDividend, errx := dividend.GetBigInt()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by dividend.GetBigInt(). " +
			"dividend='%v' Error='%v'",
			dividend.GetNumStr(), errx.Error())
		return quotient, modulo, err
	}

	bINumDividend := BigIntNum{}.NewBigInt(bIntDividend, dividend.GetPrecisionUint())

	bIntDivisor, errx := divisor.GetBigInt()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by divisor.GetBigInt(). " +
			"divisor='%v' Error='%v'",
			divisor.GetNumStr(), errx.Error())
		return quotient, modulo, err
	}

	bINumDivisor := BigIntNum{}.NewBigInt(bIntDivisor, divisor.GetPrecisionUint())

	quotient, modulo, errx =
		BigIntMathDivide{}.BigIntNumQuotientMod(bINumDividend, bINumDivisor, maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"BigIntNumQuotientMod(bINumDividend, bINumDivisor, maxPrecision). " +
			"Error='%v'",	errx.Error())

		quotient = BigIntNum{}.New()
		modulo = BigIntNum{}.New()

		return quotient, modulo, err
	}

	err = nil

	return quotient, modulo, err
}

// INumMgrFracQuotient - Performs a division operation on input
// parameters 'dividend' and 'divisor' which implement the INumMgr
// interface.
//
// The resulting quotient is returned as a BigIntNum type representing the
// result of the division operation expressed as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'. Precision is defined
// as the number of fractional digits to the right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
//
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//																				     Quotient
//  Dividend		divided by	Divisor		=		  BigIntNum Integer 	Precision	 Result
//  -------- 	  ----------	--------				-----------------	  ---------	 ------
// 	 10.5  				/ 				2 				= 			525  							  2  			 5.25
// 	 10    				/ 				2 				= 			5	  							  0  			 5
//   11.5  				/         2.5				=  			46								  1				 4.6
//    2.5					/				 12.555			=				199123855037834	   15				 0.199123855037834
//	-12.555 			/ 				2.5 			= 		 -5022							  3				-5.022
//  -12.555     	/    			2  			  = 		 -62775							  4				-6.2775
//  - 2.5 				/ 			 12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	 12.555				/ 			- 2.5			  =			 -5022								3				-5.022
//   12.555 			/ 			- 2 				= 		 -62775								4				-6.2775
//    2.5 				/ 			-12.555		  = 		 -199123855037834	   15				-0.199123855037834
// 	-12.555 			/ 			- 2.5 			= 			5022								3				 5.022
//  -12.555     	/    		- 2 				= 		  62775								4				 6.2775
//  - 2.5	 				/ 			-12.555		  = 		  199123855037834	   15				 0.199123855037834
//  -10						/				- 2					=				5														 5
//
func (bIDivide BigIntMathDivide) INumMgrFracQuotient(
														dividend,
															divisor INumMgr,
																maxPrecision uint) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.INumMgrFracQuotient() "

	if divisor.IsZero() {
		fracQuotient = BigIntNum{}.New()
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuotient, err
	}

	bPair, errx := BigIntPair{}.NewINumMgr(dividend, divisor)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewINumMgr(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(), err.Error())

		return fracQuotient, err
	}

	bPair.MakePrecisionsEqual()

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(int(maxPrecision))

	fracQuotient, errx =
		BigIntNum{}.NewNumStr(numStr)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(numStr, maxPrecision). " +
			"numStr='%v' maxPrecision='%v' Error='%v'",
			numStr, maxPrecision, err.Error())

		return fracQuotient, err
	}

	fracQuotient.TrimTrailingFracZeros()

	return fracQuotient, nil
}

// INumMgrFracQuotientArray - Performs a division operation on input parameters
// 'dividends' and 'divisor' which implement the INumMgr interfce. 'dividends'
// is an array of types implementing the INumMgr interface. The division operation
// is performed on each element of the 'dividends' array using a single 'divisor'.
//
// The resulting quotients are returned as an array of types implementing the INumMgr
// Interface. The values represent the results of each division operation expressed
// as integer and fractional digits.
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Precision is defined as the number of numeric
// digits to the right of the decimal point. Be advised that this method is
// capable of calculating quotients with very long strings of fractional digits.
// Therefore the user is advised to set a relevant value for 'maxPrecision'.
//
// Examples:
// =========
// Note: For all examples maximum precision is specified as '15'.
// ----------------------------------------------------------------------------
//    	                                       Returned
//  	Dividend		divided by	Divisor		=	       Array						=  	 Result
//  	-------- 	  ----------	--------				-----------------	   		---------
//	 	 10.5  				 / 				2.5 			= 		 fracQuoArray[0]		=   	 4.2
//	 	 10    				 / 				2.5 			= 		 fracQuoArray[1]		=	 	   4
//	   11.5  				 /        2.5				=  		 fracQuoArray[2]		=		   4.6
//	    2.5					 /				2.5			  =			 fracQuoArray[3]    =      1
//		-12.555 			 / 				2.5 			= 		 fracQuoArray[4]    =  -   5.022
//	  - 2.5 				 / 			  2.5		    = 		 fracQuoArray[5]    =  -   1
//	   12.555 			 / 			  2.5 			= 		 fracQuoArray[6]    =      5.022
//	 -122.783 			 / 			  2.5 			= 		 fracQuoArray[7]    =  -  49.1132
//	-6847.231   	   /    	  2.5 			= 		 fracQuoArray[8]    =  -2738.8924
//	  - 2.5	 				 / 			  2.5		    = 		 fracQuoArray[9]    =  -   1
//	  -10						 /			  2.5				=			 fracQuoArray[10]   =  -   4
//	  -10.5					 /			  2.5				=			 fracQuoArray[11]   =  -   4.2
//
func (bIDivide BigIntMathDivide) INumMgrFracQuotientArray(
												dividends []INumMgr,
													divisor INumMgr,
														maxPrecision uint) (fracQuoArray [] INumMgr, err error) {

	ePrefix := "BigIntMathDivide.INumMgrFracQuotientArray() "

	if divisor.IsZero() {
		fracQuoArray = []INumMgr{}
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuoArray, err
	}

	lenAry := len(dividends)

	if lenAry == 0 {
		fracQuoArray = []INumMgr{}
		err = errors.New(ePrefix + "Error: Input Parameter 'dividends' is an EMPTY Array!")
		return fracQuoArray, err
	}

	fracQuoArray = make([]INumMgr, lenAry, lenAry+20)

	for i:=0; i < lenAry; i++ {

		bPair, errx := BigIntPair{}.NewINumMgr(dividends[i], divisor)

		if errx != nil {
			fracQuoArray = []INumMgr{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewINumMgr(dividends[i], divisor). " +
				"dividends[%v]='%v' divisor='%v' Error='%v'",
				i, dividends[i].GetNumStr(), divisor.GetNumStr(), errx.Error())
			return fracQuoArray, err
		}

		bPair.MakePrecisionsEqual()

		rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
		rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
		rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
		numStr := rQuotient.FloatString(int(maxPrecision))

		bINum, errx := BigIntNum{}.NewNumStr(numStr)

		if errx != nil {
			fracQuoArray = []INumMgr{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(numStr). " +
				"numStr='%v' maxPrecision='%v' Error='%v'",
				numStr, maxPrecision, errx.Error())
			return fracQuoArray, err
		}

		bINum.TrimTrailingFracZeros()

		fracQuoArray[i] = &bINum

	}

	return fracQuoArray, nil
}
