package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntDivideElectron struct {
	lock *sync.Mutex
}

// pairFracQuotientNoNumSeps - Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2' is considered
// the divisor.
//
// 'BigIntPair.maxPrecision' is used to control the maximum precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				maxPrecision			uint			 // Controls Precision
//	}
//
// This method performs a division operation on BigIntNum parameters 'dividend'
// (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
//
//	Dividend (BigIntPair.Big1) divided Divisor (BigIntPair.Big2) = quotient
//
// The resulting quotient is returned as a BigIntNum type representing the result
// of the division operation expressed as integer and fractional digits. The
// maximum number of fractional digits output to the result is controlled by
// BigIntPair.maxPrecision. Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the right of the
// decimal place.
//
// Examples:
// =========
//
// Note: For all examples BigIntPair.maxPrecision is specified as '15'.
// ----------------------------------------------------------------------------
//
//																					   Quotient
//	 Dividend		divided by	Divisor		=		BigIntNum Integer 	Precision	 Result
//	 -------- 	  ----------	--------				-----------------	  ---------	 ------
//		 10.5  				/ 				2 				= 			525  							  2  			 5.25
//		 10    				/ 				2 				= 			5	  							  0  			 5
//	  11.5  				/         2.5				=  			46								  1				 4.6
//	   2.5					/				 12.555			=				199123855037834	   15				 0.199123855037834
//		-12.555 			/ 				2.5 			= 		 -5022							  3				-5.022
//	 -12.555				/    			2  			  = 		 -62775							  4				-6.2775
//	 - 2.5					/ 			 12.555		  = 		 -199123855037834	   15				-0.199123855037834
//		 12.555				/ 			- 2.5			  =			 -5022								3				-5.022
//	  12.555				/ 			- 2 				= 		 -62775								4				-6.2775
//	   2.5					/ 			-12.555		  = 		 -199123855037834	   15				-0.199123855037834
//		-12.555 			/ 			- 2.5 			= 			5022								3				 5.022
//	 -12.555				/    		- 2 				= 		  62775								4				 6.2775
//	 - 2.5	 				/ 			-12.555		  = 		  199123855037834	   15				 0.199123855037834
//	 -10						/				- 2					=				5														 5
//
// The returned BigIntNum division result 'fracQuotient' will contain default
// numeric separators (decimal separator, thousands separator and currency
// symbol).
func (bIDivideElec *bigIntDivideElectron) pairFracQuotientNoNumSeps(
	bPair BigIntPair) (fracQuotient BigIntNum, err error) {

	if bIDivideElec.lock == nil {
		bIDivideElec.lock = new(sync.Mutex)
	}

	bIDivideElec.lock.Lock()

	defer bIDivideElec.lock.Unlock()

	ePrefix := "bigIntDivideElectron.pairFracQuotientNoNumSeps() "

	fracQuotient = BigIntNum{}.New()

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {

		err = fmt.Errorf("%v\n"+
			"Attempted Divide by ZERO!\n"+
			"'bPair.Big2' has a ZERO value.\n", ePrefix)

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

		err = fmt.Errorf("%v\n"+
			"Error returned by BigIntNum{}.NewNumStr(numStr).\n"+
			"numStr='%v'\nmaxPrecision='%v'\nError= %v\n",
			ePrefix,
			numStr,
			bPair.MaxPrecision,
			errx.Error())

		return fracQuotient, err
	}

	fracQuotient.TrimTrailingFracZeros()

	fracQuotient.SetNumericSeparatorsToDefaultIfEmpty()

	err = nil

	return fracQuotient, err
}

// pairIntQuotientNoNumSeps - Performs integer division on two BigIntNum types passed as
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
//
//																					Return Value
//	 Divisor	divided by	Dividend		=		Integer Quotient
//			5 				/ 				 2 				= 				 2
//	    5.25			/ 				 2  			= 				 2
//	    2 				/ 				 4				= 				 0
//		 -5					/ 				 2 				= 				-2
//	   -5.25			/    			 2  			= 				-2
//	   -2 				/ 				 4				= 				 0
//			5 				/ 				-2 				=					-2
//	    5.25			/ 				-2 				= 				-2
//	    2 				/ 				-4				= 				 0
//		 -5					/ 				-2 				= 				 2
//	   -5.25			/    			-2 				= 				 2
//	   -2 				/ 				-4				= 				 0
//	    12.555		/ 			  -2.5			=			    -5
//	   -12.555		/ 			  -2.5			=			     5
//	    12.555		/ 			  -2				=			    -6
//
// The returned BigIntNum division result 'intQuotient' will contain default
// numeric separators (decimal separator, thousands separator and currency
// symbol).
func (bIDivideElec *bigIntDivideElectron) pairIntQuotientNoNumSeps(
	bPair BigIntPair) (intQuotient BigIntNum, err error) {

	if bIDivideElec.lock == nil {
		bIDivideElec.lock = new(sync.Mutex)
	}

	bIDivideElec.lock.Lock()

	defer bIDivideElec.lock.Unlock()

	intQuotient = BigIntNum{}.New()

	ePrefix := "bigIntDivideElectron.pairIntQuotientNoNumSeps()"

	if bPair.Big2.IsZero() {

		err = fmt.Errorf("%v\n"+
			"Error: Attempted to divide by zero!\n"+
			"'bPair.Big2' has a zero value.\n", ePrefix)

		return intQuotient, err
	}

	bPair.MakePrecisionsEqual()

	bigIQuotient := big.NewInt(0).Quo(bPair.Big1.bigInt, bPair.Big2.bigInt)

	intQuotient = BigIntNum{}.NewBigInt(bigIQuotient, 0)

	intQuotient.SetNumericSeparatorsToDefaultIfEmpty()

	err = nil

	return intQuotient, err
}

// pairModNoNumSeps - Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor.
// The method proceeds to performs a modulo operation on input parameters 'dividend'
// (BigIntPair.Big1) and 'divisor' (BigIntPair.Big2). The modulo result is returned
// as a type BigIntNum.
//
// 'BigIntPair.maxPrecision' is used to control the precision of the resulting
// fractional modulo returned by this method. Be advised that this method is capable
// of calculating modulo values with very long strings of fractional digits. Therefore,
// the user is advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				maxPrecision			uint			 // Controls Precision
//	}
//
// The modulo operation finds the remainder after division of one number
// by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// The calculation of 'modulo' is based on T-Division (Truncate Division). See
// "Division and Modulus for Computer Scientists", DAAN LEIJEN, University of
// Utrecht Dept. of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	  						q = D div d = f(D/d)
//								r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Examples:
// =========
//
// Dividend			  mod by			Divisor			=			Modulo/Remainder
// --------				------			-------						----------------
//
//	  12.555				%						 2.5			=			 0.055
//	  12.555				% 				 	 2  			= 		 0.555
//	   2.5 					% 				 	12.555		= 	   2.5
//		-12.555				% 				   2.5 			= 		-0.055
//	 -12.555     		%    			 	 2  			= 		-0.555
//	 - 2.5 					% 				 	12.555		= 		-2.5
//		 12.555				% 				 - 2.5			=			 0.055
//	  12.555 				% 				 - 2 				= 		 0.555
//	   2.5 				  % 				 -12.555		= 		 2.5
//		-12.555				% 				 - 2.5 			= 		-0.055
//	 -12.555     		%    			 - 2 				= 		-0.555
//	 - 2.5					% 				 -12.555		= 		-2.5
//
// The returned BigIntNum division results (quotient and modulo) will
// contain default numeric separators (decimal separator, thousands
// separator and currency symbol).
func (bIDivideElec *bigIntDivideElectron) pairModNoNumSeps(
	bPair BigIntPair) (modulo BigIntNum, err error) {

	if bIDivideElec.lock == nil {
		bIDivideElec.lock = new(sync.Mutex)
	}

	bIDivideElec.lock.Lock()

	defer bIDivideElec.lock.Unlock()

	modulo = BigIntNum{}.New()

	err = nil

	ePrefix := "bigIntDivideElectron.pairModNoNumSeps()"

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Attempted divide by ZERO!\n"+
			"'bPair.Big2.bigInt' has a zero value.\n", ePrefix)

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

	return modulo, err
}

// pairQuotientModNoNumSeps -  Receives a BigIntPair type as an input parameter. 'BigIntPair.Big1'
// is treated as the Dividend. 'BigIntPair.Big2' is considered the Divisor.
// 'BigIntPair.maxPrecision' is used to control the precision of the resulting
// fractional quotient. Be advised that this method is capable of calculating
// quotients with very long strings of fractional digits. Therefore, the user
// is advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//				Big1							BigIntNum  // The Dividend
//				Big2							BigIntNum	 // The Divisor
//				maxPrecision			uint			 // Controls Precision
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
//
//	q = D div d = f(D/d) r = D mod d = D − d ·q
//
// 'quotient' is the integer result of dividing the 'dividend' by the 'divisor'
//
// 'modulo' - The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values.
//
// Examples:
// =========
//
// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
//
//	  12.555				/						 2.5			=			 5							 0.055
//	  12.555 	 			/ 				 	 2  			= 		 6							 0.555
//	   2.5 					/ 				 	12.555		= 	   0							 2.5
//	 -12.555 				/ 				   2.5 			= 		-5							-0.055
//	 -12.555     		/    			 	 2  			= 		-6							-0.555
//	 - 2.5 					/ 				 	12.555		= 		 0							-2.5
//		 12.555				/ 				 - 2.5			=			-5							 0.055
//	  12.555 				/ 				 - 2 				= 		-6							 0.555
//	   2.5 				  / 				 -12.555		= 		 0							 2.5
//		-12.555				/ 				 - 2.5 			= 		 5							-0.055
//	 -12.555     		/    			 - 2 				= 		 6							-0.555
//	 - 2.5	 				/ 				 -12.555		= 		 0							-2.5
//
// The returned BigIntNum division results (quotient and modulo) will
// contain default numeric separators (decimal separator, thousands
// separator and currency symbol).
func (bIDivideElec *bigIntDivideElectron) pairQuotientModNoNumSeps(
	bPair BigIntPair) (quotient, modulo BigIntNum, err error) {

	if bIDivideElec.lock == nil {
		bIDivideElec.lock = new(sync.Mutex)
	}

	bIDivideElec.lock.Lock()

	defer bIDivideElec.lock.Unlock()

	quotient = BigIntNum{}.New()
	modulo = BigIntNum{}.New()
	ePrefix := "bigIntDivideElectron.pairQuotientModNoNumSeps()"

	if bPair.Big2.bigInt.Cmp(big.NewInt(0)) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Attempted Divide By ZERO!\n"+
			"'bPair.Big2.bigInt' has a zero value.\n", ePrefix)

		return quotient, modulo, err
	}

	bPair.MakePrecisionsEqual()

	scratch := big.NewInt(0)

	quotientBigI, moduloBigI := big.NewInt(0).QuoRem(
		bPair.Big1.bigInt,
		bPair.Big2.bigInt,
		scratch)

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
