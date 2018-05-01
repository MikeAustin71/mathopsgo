package mathops

import (
	"math/big"
	"fmt"
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

	Result BigIntNum
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
//    2.5 					/ 				 	12.555		= 	   0							 2.5
//	-12.555 				/ 				   2.5 			= 		-5							-0.055
//  -12.555     		/    			 	 2  			= 		-6							-0.555
//  - 2.5 					/ 				 	12.555		= 		 0							-2.5
// 	 12.555					/ 				 - 2.5			=			-5							 0.055
//   12.555 				/ 				 - 2 				= 		-6							 0.555
//    2.5 				  / 				 -12.555		= 		 0							 2.5
// 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055
//  -12.555     		/    			 - 2 				= 		 6							-0.555
//  - 2.5	 					/ 				 -12.555		= 		 0							-2.5
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
//    2.5 					% 				 	12.555		= 	   2.5
//	-12.555 				% 				   2.5 			= 		-0.055
//  -12.555     		%    			 	 2  			= 		-0.555
//  - 2.5 					% 				 	12.555		= 		-2.5
// 	 12.555					% 				 - 2.5			=			 0.055
//   12.555 				% 				 - 2 				= 		 0.555
//    2.5 				  % 				 -12.555		= 		 2.5
// 	-12.555 				% 				 - 2.5 			= 		-0.055
//  -12.555     		%    			 - 2 				= 		-0.555
//  - 2.5	 					% 				 -12.555		= 		-2.5
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
//  Divisor		divided by	Dividend		=		BigIntNum Integer 	Precision	 Result
//  ------- 	----------	--------				-----------------	  ---------	 ------
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
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error: Attempted divide by zero!")
	}

	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)
	bPair.MakePrecisionsEqual()

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(int(maxPrecision))


	fracQuotient, err =
				BigIntNum{}.NewNumStrMaxPrecision(numStr, maxPrecision)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrMaxPrecision(numStr, maxPrecision). " +
				"numStr='%v' maxPrecision='%v' Error='%v'",
				numStr, maxPrecision, err.Error())
	}

	return fracQuotient, nil
}