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
// The returned BigIntNum division 'result' (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'
//
func (bIDivide BigIntMathDivide) BigIntNumQuotientMod(
								dividend,
									divisor BigIntNum,
										maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.BigIntNumQuotientMod() "


	if divisor.IsZero() {
		quotient = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		modulo = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		err = fmt.Errorf(ePrefix + "Error: Attempted to divide by zero!")
		return quotient, modulo, err
	}

	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)

	bPair.MaxPrecision = maxPrecision

	var err2 error

	quotient, modulo, err2 = BigIntMathDivide{}.PairQuotientMod(bPair)

	if err2 != nil {
		quotient = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		modulo = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}.pairQuotientMod(bPair) " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
					bPair.MaxPrecision, err2.Error())

		return quotient, modulo, err
	}

	err = nil

	return quotient, modulo, err
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
// The returned BigIntNum division 'result' (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) BigIntNumIntQuotient(
				dividend,
					divisor BigIntNum) (intQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.BigIntNumIntQuotient() "
	intQuotient = BigIntNum{}.NewBigInt(big.NewInt(0), 0)

  if divisor.IsZero() {
		err = fmt.Errorf(ePrefix + "Error: Attempted to divide by zero!")
		return intQuotient, err
	}

	numSeps := dividend.GetNumericSeparatorsDto()

	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)

	var errx error

	intQuotient, errx = BigIntMathDivide{}.pairIntQuotient(bPair)

	if errx != nil {
		intQuotient = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}.pairIntQuotient(bPair). " +
			"dividend='%v' divisor='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), errx.Error())

		return intQuotient, err
	}

	errx = intQuotient.SetNumericSeparatorsDto(numSeps)

	if errx != nil {
		intQuotient = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by intQuotient.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v'", errx.Error())

		return intQuotient, err
	}

	err = nil

	return intQuotient, err
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
// The returned BigIntNum division 'result' (modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) BigIntNumModulo(
						dividend,
							divisor BigIntNum,
								maxPrecision uint) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.BigIntNumQuotientMod() "


	if divisor.IsZero() {
		modulo = BigIntNum{}.NewBigInt(big.NewInt(0), 0)
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair := BigIntPair{}.NewBigIntNum(dividend, divisor)

	bPair.MaxPrecision = maxPrecision

	var errx error

	modulo, errx = BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
					bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
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
// The returned BigIntNum division result ('fracQuotient') will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
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

	bPair.MaxPrecision = maxPrecision

  var errx error

	fracQuotient, errx =
				BigIntMathDivide{}.PairFracQuotient(bPair)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair). " +
				"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

		return fracQuotient, err
	}


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
// Each element of the returned BigIntNum array resulting from this division operation
// will contain contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from the first element of the input parameter 'dividends'
// array.
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

	numSeps := dividends[0].GetNumericSeparatorsDto()

	fracQuoArray = make([]BigIntNum, lenAry, lenAry+20)
	var errx error

	for i:=0; i < lenAry; i++ {

		if i > 0 {
			errx = dividends[i].SetNumericSeparatorsDto(numSeps)

			if errx != nil {
				fracQuoArray = []BigIntNum{}

				err =	fmt.Errorf(ePrefix +
					"Error returned by dividends[i].SetNumericSeparatorsDto(numSeps). " +
					"Index='%v' Error='%v'", i, errx.Error())

				return fracQuoArray, err
			}
		}

		bPair := BigIntPair{}.NewBigIntNum(dividends[i], divisor)

		bPair.MaxPrecision = maxPrecision

		fracQuoArray[i], errx =
					BigIntMathDivide{}.PairFracQuotient(bPair)

		if errx != nil {
			fracQuoArray = []BigIntNum{}
			err =	fmt.Errorf(ePrefix +
					"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair). " +
					"dividend='%v' divisor='%v' maxPrecision='%v' Index='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr() , bPair.MaxPrecision, i, errx.Error())

			return fracQuoArray, err
		}

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
// The returned BigIntNum division results ('quotient' and 'modulo') will
// contain numeric separators (decimal separator, thousands separator and
// currency symbol) copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) DecimalQuotientMod(
								dividend,
									divisor Decimal,
										maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.DecimalQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	bPair, errx := BigIntPair{}.NewDecimal(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewDecimal(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
			dividend.GetNumStr() ,divisor.GetNumStr(), errx.Error())

		return quotient, modulo, err
	}

	if bPair.Big2.IsZero() {
		err = errors.New(ePrefix + "Error: Attempted divide by ZERO!")
		return quotient, modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, errx =
			BigIntMathDivide{}.PairQuotientMod(bPair)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"pairQuotientMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

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
// The returned BigIntNum division result ('fracQuotient') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input parameter,
// 'dividend'.
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
				dividend.GetNumStr(), divisor.GetNumStr(), errx.Error())

		return fracQuotient, err
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, errx =
		BigIntMathDivide{}.PairFracQuotient(bPair)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

		return fracQuotient, err
	}

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
// Each element in the returned division results array ('fracQuoArray') will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// copied from the first element of the input parameter 'dividends' array.
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

	numSeps := dividends[0].GetNumericSeparatorsDto()

	fracQuoArray = make([]Decimal, lenAry, lenAry+20)

	var errx error
	var bPair BigIntPair

	for i:=0; i < lenAry; i++ {

		if i > 0 {

			errx = dividends[i].SetNumericSeparatorsDto(numSeps)

			if errx != nil {
				fracQuoArray = []Decimal{}
				err =	fmt.Errorf(ePrefix +
					"Error returned by dividends[%v].SetNumericSeparatorsDto(numSeps). " +
					"Error='%v'", i, errx.Error())

				return fracQuoArray, err
			}
		}

		bPair, errx = BigIntPair{}.NewDecimal(dividends[i], divisor)

		if errx != nil {
			fracQuoArray = []Decimal{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewDecimal(dividends[i], divisor). " +
				"dividends[%v]='%v' divisor='%v' Error='%v'",
				i, dividends[i].GetNumStr(), divisor.GetNumStr(), errx.Error())

			return fracQuoArray, err
		}

		bPair.MaxPrecision = maxPrecision

		bINum, errx := BigIntMathDivide{}.PairFracQuotient(bPair)

		if errx != nil {
			fracQuoArray = []Decimal{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair). " +
				"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())
			return fracQuoArray, err
		}

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

// DecimalModulo - Performs a modulo operation on Decimal type
// input parameters, 'dividend' and 'divisor'. The result of this division
// operation is returned as a type BigIntNum.
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one Decimal value: 'modulo'.
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
// The returned BigIntNum division result ('modulo') will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) DecimalModulo(
									dividend,
										divisor Decimal,
											maxPrecision uint) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.DecimalModulo() "

	modulo = BigIntNum{}.New()

	bPair, errx := BigIntPair{}.NewDecimal(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewDecimal(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(),
			maxPrecision, errx.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	modulo, errx = BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// DecimalModuloToDecimal - Performs a modulo operation on Decimal input
// parameters 'dividend' and 'divisor'.
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one Decimal value: 'modulo'.
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
// The returned BigIntNum division result ('modulo') will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
//
// The difference between this method and BigIntMathDivide.DecimalModulo()
// above, is that this method returns the resulting modulo value as a type
// 'Decimal'.
//
func (bIDivide BigIntMathDivide) DecimalModuloToDecimal(
										dividend,
											divisor Decimal,
												maxPrecision uint) (modulo Decimal, err error) {

	ePrefix := "BigIntMathDivide.DecimalModuloToDecimal() "

	bPair, errx := BigIntPair{}.NewDecimal(dividend, divisor)

	if errx != nil {
		modulo = Decimal{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewDecimal(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(),
			maxPrecision, errx.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		modulo = Decimal{}.New()
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	bINumModulo, errx := BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = Decimal{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	modulo, errx = bINumModulo.GetDecimal()

	if errx != nil {
		modulo = Decimal{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by bINumModulo.GetDecimal(). " +
			"Error='%v'", errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
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
// The returned BigIntNum division results ('quotient' and 'modulo') will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) IntAryQuotientMod(
														dividend,
															divisor IntAry,
																maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.IntAryQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	bPair, errx := BigIntPair{}.NewIntAry(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewIntAry(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), errx.Error())

		return quotient, modulo, err
	}

	if bPair.Big2.IsZero() {
		err = errors.New(ePrefix + "Error: Attempted divide by ZERO!")
		return quotient, modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, errx =
		BigIntMathDivide{}.PairQuotientMod(bPair)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"pairQuotientMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v'  Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

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
// The returned BigIntNum division result ('fracQuotient') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input parameter,
// 'dividend'.
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
			dividend.GetNumStr(), divisor.GetNumStr(), errx.Error())

		return fracQuotient, err
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, errx =
		BigIntMathDivide{}.PairFracQuotient(bPair)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

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
// The returned []IntAry division result ('fracQuoArray') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from the first
// element of the input parameter 'dividends' array.
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

	numSeps := dividends[0].GetNumericSeparatorsDto()

	fracQuoArray = make([]IntAry, lenAry, lenAry+20)

	var bPair BigIntPair
	var errx error

	for i:=0; i < lenAry; i++ {

		if i > 0 {

			errx = dividends[i].SetNumericSeparatorsDto(numSeps)

			if errx != nil {

				fracQuoArray = []IntAry{}

				err =	fmt.Errorf(ePrefix +
					"Error returned by dividends[%v].SetNumericSeparatorsDto(numSeps). " +
					"Error='%v'", i, errx.Error())

				return fracQuoArray, err
			}

		}

		bPair, errx = BigIntPair{}.NewIntAry(dividends[i], divisor)

		if errx != nil {

			fracQuoArray = []IntAry{}

			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewIntAry(dividends[i], divisor). " +
				"dividends[%v]='%v' divisor='%v' Error='%v'",
				i, dividends[i].GetNumStr(), divisor.GetNumStr(), errx.Error())

			return fracQuoArray, err
		}

		bPair.MaxPrecision = maxPrecision

		bINum, errx := BigIntMathDivide{}.pairFracQuotient(bPair)

		if errx != nil {
			fracQuoArray = []IntAry{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair). " +
				"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())
			return fracQuoArray, err
		}

		fracQuoArray[i], errx = bINum.GetIntAry()

		if errx != nil {
			fracQuoArray = []IntAry{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by bINum.GetIntAryElements(). Error='%v'",
				errx.Error())
			return fracQuoArray, err
		}

	}

	return fracQuoArray, nil
}

// IntAryModulo - Performs a modulo operation on IntAry input
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
// The returned BigIntNum division result ('modulo') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input
// parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) IntAryModulo(
						dividend,
							divisor IntAry,
								maxPrecision uint) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.IntAryModulo() "

	bPair, errx := BigIntPair{}.NewIntAry(dividend, divisor)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewIntAry(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(),
			maxPrecision, errx.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	modulo, errx = BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// IntAryModuloToIntAry - Performs a modulo operation on IntAry input
// parameters 'dividend' and 'divisor'.
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one IntAry value: 'modulo'.
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
// The returned IntAry division result ('modulo') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input
// parameter, 'dividend'.
//
// The difference between this method and BigIntMathDivide.IntAryModulo() above,
// is that the division result, 'modulo', is returned as Type IntAry.
//
func (bIDivide BigIntMathDivide) IntAryModuloToIntAry(
						dividend,
							divisor IntAry,
								maxPrecision uint) (modulo IntAry, err error) {

	ePrefix := "BigIntMathDivide.IntAryModuloToIntAry() "

	bPair, errx := BigIntPair{}.NewIntAry(dividend, divisor)

	if errx != nil {
		modulo = IntAry{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewIntAry(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(),
			maxPrecision, errx.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		modulo = IntAry{}.New()
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	bINumModulo, errx := BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = IntAry{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.pairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	modulo, errx = bINumModulo.GetIntAry()

	if errx != nil {
		modulo = IntAry{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by bINumModulo.GetIntAryElements(). " +
			"Error='%v'", errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
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
// The returned BigIntNum division results ('quotient' and 'modulo') will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied from
// input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) INumMgrQuotientMod(
											dividend,
												divisor INumMgr,
													maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.INumMgrQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	bPair, errx := BigIntPair{}.NewINumMgr(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewINumMgr(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(), maxPrecision, errx.Error())

		return quotient, modulo, err
	}

	if bPair.Big2.IsZero() {
		err = errors.New(ePrefix + "Error: Attempted divide by ZERO!")
		return quotient, modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, errx =
		BigIntMathDivide{}.PairQuotientMod(bPair)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.PairQuotientMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

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
// The returned BigIntNum division result ('fracQuotient') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input parameter,
// 'dividend'.
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

	// Validity tests are performed on 'dividend' and 'divisor'
	bPair, errx := BigIntPair{}.NewINumMgr(dividend, divisor)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewINumMgr(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(), errx.Error())

		return fracQuotient, err
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, errx =
		BigIntMathDivide{}.PairFracQuotient(bPair)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.PairFracQuotient(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

		return fracQuotient, err
	}

	return fracQuotient, nil
}

// INumMgrFracQuotientArray - Performs a division operation on input parameters
// 'dividends' and 'divisor' which implement the INumMgr interface. 'dividends'
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
// The returned INumMgr array division result ('fracQuoArray') will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied
// from the first element of input parameter, 'dividends' array.
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

	numSeps := dividends[0].GetNumericSeparatorsDto()

	fracQuoArray = make([]INumMgr, lenAry, lenAry+20)

	var errx error
	var bPair BigIntPair
	var bINum BigIntNum

	for i:=0; i < lenAry; i++ {

		if i > 0 {
			errx = dividends[i].SetNumericSeparatorsDto(numSeps)

			if errx != nil {
				fracQuoArray = []INumMgr{}
				err = fmt.Errorf(ePrefix +
					"Error returned by dividends[%v].SetNumericSeparatorsDto(numSeps). " +
					"Error='%v'", i, errx.Error())

				return fracQuoArray, err
			}

		}

		bPair, errx = BigIntPair{}.NewINumMgr(dividends[i], divisor)

		if errx != nil {
			fracQuoArray = []INumMgr{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewINumMgr(dividends[i], divisor). " +
				"dividends[%v]='%v' divisor='%v' Error='%v'",
				i, dividends[i].GetNumStr(), divisor.GetNumStr(), errx.Error())

			return fracQuoArray, err
		}

		bPair.MaxPrecision = maxPrecision

		bINum, errx = BigIntMathDivide{}.PairFracQuotient(bPair)

		if errx != nil {
			fracQuoArray = []INumMgr{}
			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.PairFracQuotient(bPair). " +
				"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
				bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

			return fracQuoArray, err
		}

		fracQuoArray[i] = &bINum

	}

	err = nil

	return fracQuoArray, err
}

// INumMgrModulo - Performs a modulo operation on input parameters
// 'dividend' and 'divisor'. Both parameters must implement the INumMgr
// interface.
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
// The returned BigIntNum division result ('modulo') will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) INumMgrModulo(
											dividend,
												divisor INumMgr,
													maxPrecision uint) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.INumMgrModulo() "

	bPair, errx := BigIntPair{}.NewINumMgr(dividend, divisor)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewINumMgr(" +
			"dividend, divisor) dividend='%v' divisor='%v' Error='%v'",
				dividend.GetNumStr(), divisor.GetNumStr(), err.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	modulo, errx = BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.PairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// NumStNumber strings are strings of numeric digits which may,
//// or may not, include a decimal separator ('.') used to separate integer and fractional digits..rQuotientMod - Performs a division operation on input parameters 'dividend' and 'divisor'
// which are comprised as number strings. Number strings are strings of numeric digits which may,
// or may not, include a decimal separator ('.') used to separate integer and fractional digits..
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
// The returned BigIntNum division results ('quotient' and 'modulo') will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bIDivide BigIntMathDivide) NumStrQuotientMod(
					dividend,
						divisor string,
							maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.NumStrQuotientMod() "

	quotient = BigIntNum{}.NewBigInt(big.NewInt(0), 0)

	modulo = BigIntNum{}.NewBigInt(big.NewInt(0), 0)

	bPair, errx := BigIntPair{}.NewNumStr(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(dividend, divisor) " +
			"dividend='%v' divisor='%v' Error='%v' ",
				dividend ,divisor, errx.Error())
		return quotient, modulo, err
	}

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {
		err = errors.New(ePrefix + "Error: Attempted Divide By ZERO!")
		return quotient, modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, errx =
		BigIntMathDivide{}.PairQuotientMod(bPair)

	if errx != nil {
		quotient = BigIntNum{}.New()
		modulo = BigIntNum{}.New()

		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"pairQuotientMod(bPair). " +
		  " dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

		return quotient, modulo, err
	}

	return quotient, modulo, nil
}

// NumStrFracQuotient - Performs a division operation on NumStr Type input
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
// The returned BigIntNum division result ('fracQuotient') will contain default numeric
// separators (decimal separator, thousands separator and currency symbol).
//
func (bIDivide BigIntMathDivide) NumStrFracQuotient(
															dividend,
																divisor string,
																	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.NumStrFracQuotient() "

	bPair, errx := BigIntPair{}.NewNumStr(dividend, divisor)

	if errx != nil {
		fracQuotient = BigIntNum{}
		err = fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(dividend, divisor) " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v' ",
			dividend, divisor, maxPrecision, errx.Error())

		return fracQuotient, err
	}

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {
		fracQuotient = BigIntNum{}.New()
		err = errors.New(ePrefix + "Error: Attempted divide by zero!")
		return fracQuotient, err
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, errx =
		BigIntMathDivide{}.pairFracQuotient(bPair)

	if errx != nil {
		fracQuotient = BigIntNum{}
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}.pairFracQuotient(" +
			"bPair)" +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v' ",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

		return fracQuotient, err
	}

	err = nil

	return fracQuotient, err
}

// NumStrFracQuotientArray - Performs a division operation on input parameters
// 'dividends' and 'divisor' which are formatted as number strings.
//
// 	Number strings are strings of numeric digits which may, or may not, include
//  a decimal separator ('.') to separate integer and fractional digits.
//
// 'dividends' is an array of number strings. The division operation is performed
// on each element of the 'dividends' array using a single 'divisor'.
//
// The resulting quotients are returned as an array of BigIntNum types. The values
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
// Each element of the returned BigIntNum Array ('fracQuoArray') will contain default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (bIDivide BigIntMathDivide) NumStrFracQuotientArray(
					dividends []string,
						divisor string,
							maxPrecision uint) (fracQuoArray [] BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.NumStrFracQuotientArray() "

	lenAry := len(dividends)

	if lenAry == 0 {
		fracQuoArray = []BigIntNum{}
		err = errors.New(ePrefix + "Error: Input Parameter 'dividends' is an EMPTY Array!")
		return fracQuoArray, err
	}

	fracQuoArray = make([]BigIntNum, lenAry, lenAry+20)
	var errx error
	var bPair BigIntPair

	for i:=0; i < lenAry; i++ {

		bPair, errx = BigIntPair{}.NewNumStr(dividends[i], divisor)

		if errx != nil {
			fracQuoArray = []BigIntNum{}
			err = fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewNumStr(dividends[i], divisor) " +
				"dividends[%v]='%v' divisor='%v' Error='%v' ",
				i, dividends[i], divisor, errx.Error())

			return fracQuoArray, err
		}

		bPair.MaxPrecision = maxPrecision

		fracQuoArray[i], errx =
			BigIntMathDivide{}.pairFracQuotient(bPair)

		if errx != nil {
			fracQuoArray = []BigIntNum{}
			err = fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.pairFracQuotient(bPair) " +
				"dividend['%v']='%v' divisor='%v' maxPrecision='%v' Error='%v' ",
				i, bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), maxPrecision, errx.Error())

			return fracQuoArray, err
		}

	}

	err = nil

	return fracQuoArray, err
}


// NumStrModulo - Performs a modulo operation on input parameters 'dividend'
// and 'divisor'. Both input parameters are formatted a number strings. Number
// strings are strings of numeric digits which may,/ or may not, include a decimal
// separator ('.') used to separate integer and fractional digits.
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
// The returned BigIntNum division result ('modulo') will contain default
// numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bIDivide BigIntMathDivide) NumStrModulo(
											dividend,
												divisor string,
													maxPrecision uint) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.NumStrModulo() "

	bPair, errx := BigIntPair{}.NewNumStr(dividend, divisor)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewNumStr(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend, divisor, maxPrecision, errx.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	modulo, errx = BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.PairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// NumStrDtoModulo - Performs a modulo operation on NumStrDto type
// input parameters, 'dividend' and 'divisor'. The result of this division
// operation is returned as a type BigIntNum.
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one NumStrDto value: 'modulo'.
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
// The returned BigIntNum division result ('modulo') will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input
// parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) NumStrDtoModulo(
										dividend,
											divisor NumStrDto,
												maxPrecision uint) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.NumStrDtoModulo() "

	modulo = BigIntNum{}.New()

	// This method will test the validity of dividend and divisor.
	bPair, errx := BigIntPair{}.NewNumStrDto(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewNumStrDto(dividend, divisor). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			dividend.GetNumStr(), divisor.GetNumStr(),
			maxPrecision, errx.Error())

		return modulo, err
	}

	if bPair.Big2.IsZero() {
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	modulo, errx = BigIntMathDivide{}.PairMod(bPair)

	if errx != nil {
		modulo = BigIntNum{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.PairMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(),
			bPair.MaxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// NumStrDtoModuloToNumStrDto - Performs a modulo operation on NumStrDto type
// input parameters, 'dividend' and 'divisor'. The result of this division
// operation is returned as a type NumStrDto.
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one NumStrDto value: 'modulo'.
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
// The returned NumStrDto division result ('modulo') will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter,'dividend'.
//
// This method is different from method BigIntMathDivide.NumStrDtoModulo()
// above, in that this method returns 'modulo' as Type NumStrDto.
//
func (bIDivide BigIntMathDivide) NumStrDtoModuloToNumStrDto(
													dividend,
														divisor NumStrDto,
															maxPrecision uint) (modulo NumStrDto, err error) {

	ePrefix := "BigIntMathDivide.NumStrDtoModuloToNumStrDto() "

	modulo = NumStrDto{}.New()

	if divisor.IsZero() {
		err = fmt.Errorf(ePrefix + "Error: Attempted to mod by zero!")
		return modulo, err
	}

	bINumModulo, errx :=
		BigIntMathDivide{}.NumStrDtoModulo(dividend, divisor, maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.NumStrDtoModulo(dividend, divisor," +
			" maxPrecision). Error='%v'", errx.Error())

		return modulo, err
	}

	modulo, errx = bINumModulo.GetNumStrDto()

	if errx != nil {
		modulo = NumStrDto{}.New()
		err = fmt.Errorf(ePrefix +
			"Error returned by bINumModulo.GetNumStrDto(). Error='%v'",
			errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// NumStrDtoQuotientMod - Performs a division operation on NumStrDto type input
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
// The returned BigIntNum division results ('quotient' and 'modulo') will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
//
func (bIDivide BigIntMathDivide) NumStrDtoQuotientMod(
											dividend,
												divisor NumStrDto,
													maxPrecision uint) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.NumStrDtoQuotientMod() "

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	err = nil

	// This method will test the validity of dividend and divisor
	bPair, errx := BigIntPair{}.NewNumStrDto(dividend, divisor)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntPair{}.NewNumStrDto(dividend, divisor). " +
			"dividend='%v' divisor='%v' Error='%v'",
			dividend.GetNumStr() ,divisor.GetNumStr(), errx.Error())

		return quotient, modulo, err
	}

	if bPair.Big2.IsZero() {
		err = errors.New(ePrefix + "Error: Attempted divide by ZERO!")
		return quotient, modulo, err
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, errx =
		BigIntMathDivide{}.PairQuotientMod(bPair)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}."+
			"PairQuotientMod(bPair). " +
			"dividend='%v' divisor='%v' maxPrecision='%v' Error='%v'",
			bPair.Big1.GetNumStr(), bPair.Big2.GetNumStr(), bPair.MaxPrecision, errx.Error())

		quotient = BigIntNum{}.New()
		modulo = BigIntNum{}.New()

		return quotient, modulo, err
	}

	err = nil

	return quotient, modulo, err
}

// PairQuotientMod -  Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor.
// 'BigIntPair.MaxPrecision' is used to control the precision of the resulting
// fractional quotient. Be advised that this method is capable of calculating
// quotients with very long strings of fractional digits. Therefore, the user
// is advised to set a relevant 'BigIntPair.MaxPrecision' value.
//
// 	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				MaxPrecision			uint			 // Controls Precision
//	}
//
// The method performs a division operation on BigIntNum input
// parameters 'dividend' (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
// The result is a quotient and modulo (remainder). These values are returned as
// two BigIntNum types, 'quotient' and 'modulo'.
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
// The returned BigIntNum division results (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from the the input parameter Dividend,
// 'BigIntPair.Big1'.
//
func (bIDivide BigIntMathDivide) PairQuotientMod(
	bPair BigIntPair) (quotient, modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.PairQuotientMod() "

	numSeps := bPair.Big1.GetNumericSeparatorsDto()

	var err2 error

	quotient, modulo, err2 = bIDivide.pairQuotientMod(bPair)

	if err2 != nil {
		quotient = BigIntNum{}.NewZero(0)
		modulo = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix + "Error returned by bIDivide.pairQuotientMod(bPair). " +
			"Error='%v'", err2.Error())

		return quotient, modulo, err
	}

	err2 = quotient.SetNumericSeparatorsDto(numSeps)

	if err2 != nil {
		quotient = BigIntNum{}.NewZero(0)
		modulo = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix + "Error returned by quotient.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v'", err2.Error())

		return quotient, modulo, err
	}

	err2 = modulo.SetNumericSeparatorsDto(numSeps)

	if err2 != nil {
		quotient = BigIntNum{}.NewZero(0)
		modulo = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix + "Error returned by modulo.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v'", err2.Error())

		return quotient, modulo, err
	}

	err = nil

	return quotient, modulo, err
}


// PairMod - Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor.
// The method proceeds to performs a modulo operation on input parameters 'dividend'
// (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2). The modulo result is returned
// as a type BigIntNum.
//
// 'BigIntPair.MaxPrecision' is used to control the precision of the resulting
// fractional modulo returned by this method. Be advised that this method is capable
// of calculating modulo values with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant 'BigIntPair.MaxPrecision' value.
//
// 	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				MaxPrecision			uint			 // Controls Precision
//	}
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
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
// The returned BigIntNum division result 'modulo' will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from the the input parameter Dividend, 'BigIntPair.Big1'.
//
func (bIDivide BigIntMathDivide) PairMod(
																	bPair BigIntPair) (modulo BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.PairMod() "

	numSeps := bPair.Big1.GetNumericSeparatorsDto()

	var err2 error

	modulo, err2 = bIDivide.pairMod(bPair)

	if err2 != nil {
		modulo = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by bIDivide.pairMod(bPair). " +
			"Error='%v' ", err2.Error())

		return modulo, err
	}

	err2 = modulo.SetNumericSeparatorsDto(numSeps)

	if err2 != nil {
		modulo = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by modulo.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' ", err2.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// PairIntQuotient - Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor. This method
// performs integer division on input parameters, 'Dividend' (BigIntPair.Big1) and
// 'Divisor' (BigIntPair.Big2).
//
// The result of this division operation returns an integer quotient of Dividend
// (BigIntPair.Big1) divided by Divisor (BigIntPair.Big2).
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
//  Dividend divided by	Divisor			=		Integer Quotient
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
// The returned BigIntNum division result 'intQuotient' will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from the input parameter Dividend, 'BigIntPair.Big1'.
//
func (bIDivide BigIntMathDivide) PairIntQuotient(bPair BigIntPair) (intQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.PairIntQuotient() "

	numSeps := bPair.Big1.GetNumericSeparatorsDto()

	var err2 error

	intQuotient, err2 = bIDivide.pairIntQuotient(bPair)

	if err2 != nil {
		intQuotient = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by bIDivide.pairIntQuotient(bPair). " +
			"Error='%v' ", err2.Error())

		return intQuotient, err
	}

	err2 = intQuotient.SetNumericSeparatorsDto(numSeps)

	if err2 != nil {
		intQuotient = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by intQuotient.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' ", err2.Error())

		return intQuotient, err
	}

	err = nil

	return intQuotient, err
}

// PairFracQuotient - Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2' is considered
// the divisor.
//
// 'BigIntPair.MaxPrecision' is used to control the maximum precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant 'BigIntPair.MaxPrecision' value.
//
// 	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				MaxPrecision			uint			 // Controls Precision
//	}
//
// This method performs a division operation on BigIntNum parameters 'dividend'
// (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
//
// 		Dividend (BigIntPair.Big1) divided Divisor (BigIntPair.Big2) = quotient
//
// The resulting quotient is returned as a BigIntNum type representing the result
// of the division operation expressed as integer and fractional digits. The
// maximum number of fractional digits output to the result is controlled by
// BigIntPair.MaxPrecision. Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the right of the
// decimal place.
//
// Examples:
// =========
//
// Note: For all examples BigIntPair.MaxPrecision is specified as '15'.
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
// The returned BigIntNum division result 'fracQuotient' will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied
// from the input parameter Dividend, 'BigIntPair.Big1'.
//
func (bIDivide BigIntMathDivide) PairFracQuotient(
									bPair BigIntPair) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.pairFracQuotient() "

	numSeps := bPair.Big1.GetNumericSeparatorsDto()

	var err2 error

	fracQuotient, err2 = bIDivide.pairFracQuotient(bPair)

	if err2 != nil {
		fracQuotient = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by bIDivide.pairFracQuotient(bPair). " +
			"Error='%v' ", err2.Error())

		return fracQuotient, err
	}

	err2 = fracQuotient.SetNumericSeparatorsDto(numSeps)

	if err2 != nil {
		fracQuotient = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned by fracQuotient.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' ", err2.Error())

		return fracQuotient, err
	}

	err = nil

	return fracQuotient, err
}


// pairFracQuotient - Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2' is considered
// the divisor.
//
// 'BigIntPair.MaxPrecision' is used to control the maximum precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant 'BigIntPair.MaxPrecision' value.
//
// 	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				MaxPrecision			uint			 // Controls Precision
//	}
//
// This method performs a division operation on BigIntNum parameters 'dividend'
// (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
//
// 		Dividend (BigIntPair.Big1) divided Divisor (BigIntPair.Big2) = quotient
//
// The resulting quotient is returned as a BigIntNum type representing the result
// of the division operation expressed as integer and fractional digits. The
// maximum number of fractional digits output to the result is controlled by
// BigIntPair.MaxPrecision. Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the right of the
// decimal place.
//
// Examples:
// =========
//
// Note: For all examples BigIntPair.MaxPrecision is specified as '15'.
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
// The returned BigIntNum division result 'fracQuotient' will contain default
// numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bIDivide BigIntMathDivide) pairFracQuotient(
										bPair BigIntPair) (fracQuotient BigIntNum, err error) {

	ePrefix := "BigIntMathDivide.pairFracQuotient() "
	fracQuotient = BigIntNum{}.New()

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {
		err = errors.New(ePrefix + "Attempted Divide by ZERO!")
		return fracQuotient, err
	}

	bPair.MakePrecisionsEqual()

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)
	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(int(bPair.MaxPrecision))

	fracQuotient, errx :=
		BigIntNum{}.NewNumStr(numStr)

	if errx != nil {
		fracQuotient = BigIntNum{}.New()
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(numStr). " +
			"numStr='%v' maxPrecision='%v' Error='%v'",
			numStr, bPair.MaxPrecision, errx.Error())

		return fracQuotient, err
	}

	fracQuotient.TrimTrailingFracZeros()

	fracQuotient.SetNumericSeparatorsToDefaultIfEmpty()

  err = nil

	return fracQuotient, err
}

// pairIntQuotient - Performs integer division on two BigIntNum types passed as
// BigIntPair input parameters. 'dividend' is BigIntPair.Big1 and 'divisor' is
// BigIntPair.Big2.
//
// The result of this division operation returns an integer quotient of Dividend
// (BigIntPair.Big1) divided by Divisor (BigIntPair.Big2).
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
// The returned BigIntNum division result 'intQuotient' will contain default
// numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bIDivide BigIntMathDivide) pairIntQuotient(
												bPair BigIntPair) (intQuotient BigIntNum, err error) {

	intQuotient = BigIntNum{}.New()

	if bPair.Big2.IsZero() {
		ePrefix := "BigIntMathDivide.pairIntQuotient() "
		err = fmt.Errorf(ePrefix + "Error: Attempted to divide by zero!")
		return intQuotient, err
	}

	bPair.MakePrecisionsEqual()

	bigIQuotient := big.NewInt(0).Quo(bPair.Big1.bigInt, bPair.Big2.bigInt)

	intQuotient = BigIntNum{}.NewBigInt(bigIQuotient, 0)

	intQuotient.SetNumericSeparatorsToDefaultIfEmpty()

	err = nil

	return intQuotient, err
}

// pairMod - Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor.
// The method proceeds to performs a modulo operation on input parameters 'dividend'
// (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2). The modulo result is returned
// as a type BigIntNum.
//
// 'BigIntPair.MaxPrecision' is used to control the precision of the resulting
// fractional modulo returned by this method. Be advised that this method is capable
// of calculating modulo values with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant 'BigIntPair.MaxPrecision' value.
//
// 	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				MaxPrecision			uint			 // Controls Precision
//	}
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
// 				Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
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
// The returned BigIntNum division results (quotient and modulo) will
// contain default numeric separators (decimal separator, thousands
// separator and currency symbol).
//
func (bIDivide BigIntMathDivide) pairMod(bPair BigIntPair) (modulo BigIntNum, err error) {

	modulo = BigIntNum{}.New()

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {
		ePrefix := "BigIntMathDivide.pairMod() "
		err = errors.New(ePrefix + "Error: Attempted divide by ZERO!")
		return modulo, err
	}

	bPair.MakePrecisionsEqual()

	moduloBigI := big.NewInt(0).Rem(bPair.Big1.bigInt, bPair.Big2.bigInt)

	modulo = BigIntNum{}.NewBigInt(moduloBigI, bPair.Big2.GetPrecisionUint())

	if modulo.precision > bPair.MaxPrecision {
		modulo.RoundToDecPlace(bPair.MaxPrecision)
	}

	modulo.TrimTrailingFracZeros()

	modulo.SetNumericSeparatorsToDefaultIfEmpty()

	err = nil

	return modulo, err
}

// pairQuotientMod -  Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor.
// 'BigIntPair.MaxPrecision' is used to control the precision of the resulting
// fractional quotient. Be advised that this method is capable of calculating
// quotients with very long strings of fractional digits. Therefore, the user
// is advised to set a relevant 'BigIntPair.MaxPrecision' value.
//
// 	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				MaxPrecision			uint			 // Controls Precision
//	}
//
// The method performs a division operation on BigIntNum input
// parameters 'dividend' (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
// The result is a quotient and modulo (remainder). These values are returned as
// two BigIntNum types, 'quotient' and 'modulo'.
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
// The returned BigIntNum division results (quotient and modulo) will
// contain default numeric separators (decimal separator, thousands
// separator and currency symbol).
//
func (bIDivide BigIntMathDivide) pairQuotientMod(
								bPair BigIntPair) (quotient, modulo BigIntNum, err error) {

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {

		ePrefix := "BigIntMathDivide.pairQuotientMod() "

		err = errors.New(ePrefix +
			"Error: Attempted Divide By ZERO!")

		return quotient, modulo, err
	}

	bPair.MakePrecisionsEqual()

	quotientBigI, moduloBigI := big.NewInt(0).QuoRem(
			bPair.Big1.bigInt,
				bPair.Big2.bigInt,
					big.NewInt(0))

	quotient = BigIntNum{}.NewBigInt(quotientBigI, 0)

	modulo = BigIntNum{}.NewBigInt(moduloBigI, bPair.Big2.GetPrecisionUint())

	if modulo.precision > bPair.MaxPrecision {
		modulo.RoundToDecPlace(bPair.MaxPrecision)
	}

	modulo.TrimTrailingFracZeros()

	quotient.SetNumericSeparatorsToDefaultIfEmpty()

	modulo.SetNumericSeparatorsToDefaultIfEmpty()

	err = nil

	return quotient, modulo, err
}
