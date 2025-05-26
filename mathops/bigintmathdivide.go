package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
)

// BigIntMathDivide - This type contains methods used to perform the
// division operation using *big.Int numeric types.
//
// Reference the 'big' math package: https://golang.org/pkg/math/big/
type BigIntMathDivide struct {
	Input BigIntPair
	// BigIntPair.Big1 = Dividend
	// BigIntPair.Big2 = Divisor

	Result BigIntNum // BigIntPair.Big1 = Quotient
	// BigIntPair.Big2 = Modulo

	ResultFracQuo BigIntNum // Quotient expressed with fractional digits
	// to the right of the decimal place.
}

// BigIntDividedByTwoToPower
//
// Performs integer division by two using a 'right-shift' technique.
// Remainders from this division operation are discarded, only the
// integer quotient is returned.
//
//	Example
//	=======
//
//	  quotient =  dividend / 2^(exponent)
//
//	In the example of 33,333 / 2^8:
//
//	(1) The fractional quotient of 33,333/256 (or 2^8) is 130.
//
//	(2) This method will use a right shift technique 33,333 / 2^(8) to generate
//	    a quotient of 130.
func (bIDivide *BigIntMathDivide) BigIntDividedByTwoToPower(
	dividend *big.Int,
	exponent uint) (integerQuotient *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntDividedByTwoToPower",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if dividend == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'dividend'",
		}

		return big.NewInt(0), err
	}

	integerQuotient = big.NewInt(0)

	integerQuotient.Rsh(dividend, exponent)

	return integerQuotient, err
}

// BigIntFracQuotient
//
// Performs a division on integers of type *big.Int. The result is
// returned as a type *big.Int with an accompanying precision
// specification. Taken together, the returned *big.Int quotient
// and precision specification describe a floating point numeric
// value with a fixed number of digits after the decimal place.
// The input parameter, 'maxPrecision' is used to configure the
// maximum number of fractional digits to the right of the decimal
// place in the returned quotient.
//
//	 Example
//	 =======
//
//	 This division operation will produce a quotient which may include
//	 a fixed number of fractional digits to the right of the decimal
//	 place:
//
//	   quotient = dividend / divisor
//
//	 For the example
//
//	   quotient =	752.314 / 21.67894
//
//	 'dividend' and 'divisor' would be configured as follows:
//
//	   dividend           =  752314
//	   dividendPrecision  =       3
//	   divisor            =  2167894
//	   divisorPrecision   =       5
//
//	 Assuming a 'maxPrecision' value of '30', the quotient would be
//	 calculated as follows:
//
//	   quotient           = 34702526968569496479071393712054
//	   quotientPrecision  = 30
//
//	 Input Parameters
//	 ================
//
//	 dividend            *big.Int
//
//	 The 'dividend' value will be divided by the 'divisor'
//		to produce a 'quotient'.
//
//
//	 dividendPrecision   uint
//
//	 This unsigned integer value is a precision specification
//	 associated with input parameter 'dividend'. The precision
//	 specifies the number if digits to right of the decimal
//	 place in the series of integer digits contained in
//	 'dividend'.
//
//
//	 divisor             *big.Int
//
//	 The 'dividend' value will be divided by the 'divisor'
//	 to produce a 'quotient'.
//
//
//	 divisorPrecision    uint
//
//	 This unsigned integer value is a precision specification
//	 associated with input parameter 'divisor'. The precision
//	 specifies the number if digits to right of the decimal
//	 place in the series of integer digits contained in
//	 'divisor'.
//
//
//	 'maxPrecision'      uint
//
//	 Maximum precision specifies the maximum number of
//	 decimal digits to which the result or 'quotient'
//	 is calculated and returned to the caller. The
//	 quotient may consist of actual fractional digits
//	 which number less than 'maxPrecision'. However, if
//	 the number of digits to the right of the decimal
//	 place exceeds 'maxPrecision', the returned quotient
//	 will be rounded to 'maxPrecision' fractional digits
//	 to the right of the decimal place.
//
//	 Return Values
//	 =============
//
//	 quotient            *big.Int
//
//	 The result of the division operation expressed
//	 as an integer.
//
//
//	 quotientPrecision   uint
//
//	 An unsigned integer which specifies the number of
//	 fractional digits to the right of the decimal place
//	 in the series of integer digits defined by 'quotient'
//
//
//	 err                 error
//
//	 If an error is encountered, this function will
//	 return a quotient set equal to zero and an error
//	 object will be returned containing an appropriate
//	 error message. If the function completes the division
//	 operation successfully, the returned 'quotient' will
//	 be populated with the correct result and 'err' will
//	 be set equal to 'nil'.
func (bIDivide *BigIntMathDivide) BigIntFracQuotient(
	dividend *big.Int,
	dividendPrecision *big.Int,
	divisor *big.Int,
	divisorPrecision *big.Int,
	maxPrecision *big.Int) (quotient *big.Int, quotientPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	quotient = big.NewInt(0)

	quotientPrecision = big.NewInt(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntFracQuotient",
		"")

	if err != nil {
		return quotient, quotientPrecision, err
	}

	quotient = big.NewInt(0)
	quotientPrecision = big.NewInt(0)
	err = nil

	if dividend == nil {

		return quotient, quotientPrecision, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'dividend'",
		}
	}

	if divisor == nil {

		return quotient, quotientPrecision, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'divisor'",
		}
	}

	if dividendPrecision == nil {

		return quotient, quotientPrecision, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'dividendPrecision'",
		}
	}

	if divisorPrecision == nil {

		return quotient, quotientPrecision, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'divisorPrecision'",
		}
	}

	bigZero := big.NewInt(0)

	if divisor.Cmp(bigZero) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error - Divide by ZERO! Input parameter 'divisor' is ZERO!\n",
			ePrefix.String())

		return quotient, quotientPrecision, err
	}

	if dividend.Cmp(bigZero) == 0 {
		err = nil
		return quotient, quotientPrecision, err
	}

	// Prepare divisor
	// Setting to absolute value of divisor
	denomnatrDivsr := big.NewInt(0).Set(divisor)
	denomnatrDivsrSign := int64(1)

	if divisor.Cmp(bigZero) == -1 {
		denomnatrDivsrSign = -1
		denomnatrDivsr.Mul(denomnatrDivsr, big.NewInt(denomnatrDivsrSign))
	}

	// Prepare dividend
	// Setting to absolute value of dividend
	bigTen := big.NewInt(10)

	numratrDivdnd := big.NewInt(0).Set(dividend)
	numratrDivdndSign := int64(1)

	if dividend.Cmp(bigZero) == -1 {
		numratrDivdndSign = -1
		numratrDivdnd.Mul(numratrDivdnd, big.NewInt(numratrDivdndSign))
	}

	denomnatrDivsrShift := big.NewInt(0).Set(divisorPrecision)
	numratrDivdndShift := big.NewInt(0).Set(dividendPrecision)
	scale := big.NewInt(0)

	if numratrDivdndShift.Cmp(bigZero) == 1 {
		denomnatrDivsrShift.Sub(denomnatrDivsrShift, numratrDivdndShift)
		if denomnatrDivsrShift.Cmp(bigZero) == -1 {
			scale = big.NewInt(0).Exp(bigTen, big.NewInt(0).Mul(denomnatrDivsrShift, big.NewInt(-1)), nil)
			denomnatrDivsr.Mul(denomnatrDivsr, scale)
		}

		numratrDivdndShift = big.NewInt(0)
	}

	if denomnatrDivsrShift.Cmp(bigZero) == 1 {
		numratrDivdndShift.Sub(numratrDivdndShift, denomnatrDivsrShift)

		if numratrDivdndShift.Cmp(bigZero) == -1 {
			scale = big.NewInt(0).Exp(bigTen, big.NewInt(0).Mul(numratrDivdndShift, big.NewInt(-1)), nil)
			numratrDivdnd.Mul(numratrDivdnd, scale)
		}
	}

	// Do integer division
	scratch := big.NewInt(0)
	intQuotient, intRemndr := big.NewInt(0).QuoRem(numratrDivdnd, denomnatrDivsr, scratch)

	quotient = big.NewInt(0).Set(intQuotient)

	// Calculate fractional digits out to maxPrecision + 1
	bigOne := big.NewInt(1)
	iMaxPrecision := big.NewInt(0).Set(maxPrecision)
	iMaxPrecision.Add(iMaxPrecision, bigOne)
	iCnt := big.NewInt(0)
	lastNonZeroDigitIdx := big.NewInt(-1)

	for iCnt.Cmp(iMaxPrecision) == -1 {

		numratrDivdnd = big.NewInt(0).Mul(intRemndr, bigTen)
		intQuotient, intRemndr = big.NewInt(0).QuoRem(numratrDivdnd, denomnatrDivsr, scratch)

		if intQuotient.Cmp(bigZero) == 1 {
			lastNonZeroDigitIdx.Set(iCnt)
		}

		quotient.Mul(quotient, bigTen)
		quotient.Add(quotient, intQuotient)

		iCnt.Add(iCnt, bigOne)
	}

	if lastNonZeroDigitIdx.Cmp(big.NewInt(-1)) == 0 {

		scale = big.NewInt(0).Exp(bigTen, iMaxPrecision, nil)
		quotient.Quo(quotient, scale)
		quotientPrecision = big.NewInt(0)

	} else if lastNonZeroDigitIdx.Cmp(big.NewInt(0).Sub(iMaxPrecision, bigOne)) == -1 {
		// else if lastNonZeroDigitIdx < (i64MaxPrecision - 1)
		factor := big.NewInt(0).Sub(iMaxPrecision, lastNonZeroDigitIdx)
		factor.Sub(factor, bigOne)
		scale =
			big.NewInt(0).Exp(bigTen, factor, nil)

		quotient.Quo(quotient, scale)

		// uint(lastNonZeroDigitIdx + 1)
		quotientPrecision = big.NewInt(0).Add(lastNonZeroDigitIdx, bigOne)

	} else {

		//fmt.Println("before quotient: %v", quotient.Text(10))
		quotient.Add(quotient, big.NewInt(5))

		quotient.Quo(quotient, bigTen)
		//fmt.Println("after quotient: %v", quotient.Text(10))
		iMaxPrecision.Sub(iMaxPrecision, bigOne)

		// = uint(i64MaxPrecision)
		quotientPrecision.Set(iMaxPrecision)

	}

	if numratrDivdndSign != denomnatrDivsrSign {
		quotient.Mul(quotient, big.NewInt(-1))
	}

	return quotient, quotientPrecision, err
}

// BigIntNumQuotientMod
//
// Performs a division operation on BigIntNum input parameters
// 'dividend' and 'divisor'.
//
// There are two BigIntNum return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht/ Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at ../notes/divmodnote-letter.pdf.
//
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient' is the integer result of dividing the 'dividend' by the 'divisor'
//
//	'modulo' - The modulo operation finds the remainder after division of one
//	number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the
// resulting 'modulo'. Precision is defined as the number of fractional digits
// to the right of the decimal point. Be advised that these calculations can support
// very large precision values.
//
//	Examples
//	=========
//
//	Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//	 12.555        /         2.5   =      5          0.055
//	 12.555        /         2     =      6          0.555
//	 2.5           /        12.555 =      0          2.5
//	-12.555        /         2.5   =     -5         -0.055
//	-12.555        /         2     =     -6         -0.555
//	- 2.5          /        12.555 =      0         -2.5
//	 12.555        /       - 2.5   =     -5          0.055
//	 12.555        /       - 2     =     -6          0.555
//	 2.5           /       -12.555 =      0          2.5
//	-12.555        /       - 2.5   =      5         -0.055
//	-12.555        /       - 2     =      6         -0.555
//	-2.5           /       -12.555 =      0         -2.5
//
// The returned BigIntNum division 'result' (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'
//
//		Numeric Separators
//		==================
//
//	  Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	  instance. A NumericSeparatorDto contains symbols or characters
//	  for the decimal separator, thousands separator and currency
//	  symbol. These separators are used when presenting numeric
//	  values in number strings.
//
//		If any of the 'numSeps' Numeric Separator Components are set
//		to zero, those components will be automatically reset to USA
//		default values.
//
//		The returned values ('quotient' and 'modulo') will be
//		configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumQuotientMod(
	dividend BigIntNum,
	divisor BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	quotient = new(BigIntNum).New()
	modulo = new(BigIntNum).New()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return quotient, modulo, err
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIsZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		quotient, err = new(BigIntNum).NewBigInt(big.NewInt(0), 0)

		if err != nil {

			return quotient, modulo,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "quotient, err = new(BigIntNum).NewBigInt(big.NewInt(0), 0)",
					ErrMessage: err.Error(),
				}
		}

		modulo, err = new(BigIntNum).NewBigInt(big.NewInt(0), 0)

		if err != nil {

			return quotient, modulo,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "modulo, err = new(BigIntNum).NewBigInt(big.NewInt(0), 0)",
					ErrMessage: err.Error(),
				}
		}

		return quotient, modulo, err
	}

	bPair, err := new(BigIntPair).NewBigIntNum(dividend, divisor)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(dividend, divisor)",
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, err = new(bigIntMathDivideNanobot).
		pairQuotientMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"pairQuotientMod(&bPair, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// BigIntNumIntQuotient
//
// Performs integer division on two BigIntNum types passed as
// input parameters, 'dividend' and 'divisor'.
//
// The division operation performed by this method is T-Division
// or truncated division. See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// After completing the division operation, an integer quotient of
// type BigIntNum is returned.
//
//	Examples
//	=========
//	                                    Return Value
//	Divisor  divided by  Dividend  =  Integer Quotient
//
//	   5          /         2      =          2
//	   5.25       /         2      =          2
//	   2          /         4      =          0
//	 - 5          /         2      =         -2
//	 - 5.25       /         2      =         -2
//	 - 2          /         4      =          0
//	   5          /         2      =         -2
//	   5.25       /         2      =         -2
//	   2          /         4      =          0
//	 - 5          /         2      =          2
//	 - 5.25       /        -2      =          2
//	 - 2          /         4      =          0
//	  12.555      /        -2.5    =         -5
//	 -12.555      /        -2.5    =          5
//	  12.555      /        -2      =         -6
//
// The returned BigIntNum division 'result' (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'.
//
//		Numeric Separators
//		==================
//
//	  Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	  instance. A NumericSeparatorDto contains symbols or characters
//	  for the decimal separator, thousands separator and currency
//	  symbol. These separators are used when presenting numeric
//	  values in number strings.
//
//		If any of the 'numSeps' Numeric Separator Components are set
//		to zero, those components will be automatically reset to USA
//		default values.
//
//		The returned value ('intQuotient') will be configured with
//	 	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumIntQuotient(
	dividend BigIntNum,
	divisor BigIntNum,
	numSeps NumericSeparatorDto) (intQuotient BigIntNum, err error) {

	intQuotient = new(BigIntNum).New()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumIntQuotient",
		"")

	if err != nil {
		return intQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing input parameter 'dividend'").String())

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = dividend.IsValid(ePrefix + \"Testing input parameter 'dividend'\")",
				ErrContext: "Input parameter 'dividend' is invalid.",
				ErrMessage: err.Error(),
			}
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing input parameter 'divisor'").String())

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = divisor.IsValid(ePrefix + \"Testing input parameter 'divisor'\")",
				ErrContext: "Input parameter 'divisor' is invalid.",
				ErrMessage: err.Error(),
			}
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "intQuotient, err = new(BigIntNum).NewBigInt(big.NewInt(0), 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if divisorIsZero {",
				ErrMessage: "Input parameter 'divisor' has a ZERO value.",
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(dividend, divisor)

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(dividend, divisor)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	intQuotient, err = new(bigIntMathDivideNanobot).
		pairIntQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "intQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"    pairIntQuotient(&bPair, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return intQuotient, nil
}

// BigIntNumModulo
//
// Performs a modulo operation on BigIntNum input parameters
// 'dividend' and 'divisor'.
//
// The modulo operation finds the remainder after division of
// one number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one BigIntNum value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large
// precision values.
//
//	Examples:
//	=========
//
//	Dividend  mod by  Divisor  =  Modulo/Remainder
//	--------  ------  -------     ----------------
//
//	  12.555     %       2.5   =       0.055
//	  12.555     %       2     =       0.555
//	  2.5        %      12.555 =       2.5
//	 -12.555     %       2.5   =      -0.055
//	 -12.555     %       2     =      -0.555
//	 - 2.5       %      12.555 =      -2.5
//	  12.555     %     - 2.5   =       0.055
//	  12.555     %     - 2     =       0.555
//	   2.5       %     -12.555 =       2.5
//	 -12.555     %     - 2.5   =      -0.055
//	 -12.555     %     - 2     =      -0.555
//	 - 2.5       %     -12.555 =      -2.5
//
// The returned BigIntNum division 'result' (modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'.
//
//			Numeric Separators
//			==================
//
//	    Input parameter, 'numSeps' consits of a NumericSeparatorDto
//			instance. A NumericSeparatorDto contains symbols or characters
//			for the decimal separator, thousands separator and currency
//			symbol. These separators are used when presenting numeric
//			values in number strings.
//
//			If any of the 'numSeps' Numeric Separator Components are set
//			to zero, those components will be automatically reset to USA
//			default values.
//
//			The returned value 'modulo' will be configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumModulo(
	dividend BigIntNum,
	divisor BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	modulo = new(BigIntNum).New()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntDividedByTwoToPower",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {
		return modulo, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {
		return modulo, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	isZero, err := divisor.IsZero()

	if err != nil {
		return modulo, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "isZero, err := divisor.IsZero()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if isZero {

		return modulo, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "divisor==0",
			ErrMessage: "Error: Attempted to mod by zero!\n" +
				"Input parameter 'divisor' has a ZERO value.",
		}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(dividend, divisor)

	if err != nil {
		return modulo, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(dividend, divisor)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bPair.MaxPrecision = maxPrecision

	modulo, err = new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {
		return modulo, &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
				"pairMod(&bPair, numSeps, ePrefix)",
			ErrContext: fmt.Sprintf("dividend='%v'\ndivisor='%v'\nmaxPrecision='%v'",
				dividendNumStr, divisorNumStr, maxPrecision),
			ErrMessage: err.Error(),
		}
	}

	return modulo, nil
}

// BigIntNumFracQuotient
//
// Performs a division operation on BigIntNum input parameters
// 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type representing
// the result of the division operation expressed as integer and
// fractional digits as appropriate. Remember that the BigIntNum type
// specifies 'precision'. Precision is defined as the number of
// fractional digits to the right of the decimal place.
//
//		Examples
//		========
//
//		Note: For all examples maximum precision is specified as '15'.
//
//		                                        Quotient
//		Dividend  divided by  Divisor  =  BigIntNum Integer    Precision  Result
//		--------  ----------  -------     -----------------    ---------  ------
//
//		  10.5        /         2     =                 525        2      5.25
//		  10          /         2     =                   5        0      5
//		  11.5        /        2.5    =                  46        1      4.6
//		   2.5        /        12.555 =     199123855037834       15      0.199123855037834
//		 -12.555      /         2.5   =    -           5022        3     -5.022
//		 -12.555      /         2     =    -          62775        4     -6.2775
//		   2.5        /        12.555 =    -199123855037834       15     -0.199123855037834
//		  12.555      /       - 2.5   =    -           5022        3     -5.022
//		  12.555      /       - 2     =    -          62775        4     -6.2775
//		   2.5        /        12.555 =    -199123855037834       15     -0.199123855037834
//		 -12.55       /       - 2.5   =                5022        3      5.022
//		 -12.555      /       - 2     =               62775        4      6.2775
//	   - 2.5        /       -12.555 =     199123855037834       15      0.199123855037834
//		 -10          /       - 2     =                   5        5      5.00000
//
// The input parameter 'maxPrecision' is used to control the precision of the
// resulting fractional quotient. Be advised that this method is capable of
// calculating quotients with very long strings of fractional digits.
//
// The returned BigIntNum division result ('fracQuotient') will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter, 'dividend'.
//
//			 Numeric Separators
//			 ==================
//
//	    Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	    instance. A NumericSeparatorDto contains symbols or characters
//	    for the decimal separator, thousands separator and currency
//	    symbol. These separators are used when presenting numeric
//	    values in number strings.
//
//	    If any of the 'numSeps' Numeric Separator Components are set
//	    to zero, those components will be automatically reset to USA
//	    default values.
//
//	    The returned value ('fracQuotient') will be configured with
//	    'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumFracQuotient(
	dividend BigIntNum,
	divisor BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntDividedByTwoToPower",
		"")

	if err != nil {
		return fracQuotient, err
	}

	fracQuotient, err = new(BigIntNum).NewZero(0)

	if err != nil {
		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	binDividend, err := dividend.GetBigInt()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "binDividend, err := dividend.GetBigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	binDividendPrecision, err := dividend.GetPrecisionBigInt()

	if err != nil {
		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "binDividendPrecision, err := dividend.GetPrecisionBigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	binDivisor, err := divisor.GetBigInt()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "binDivisor, err := divisor.GetBigInt()",
				ErrMessage: err.Error(),
			}
	}

	binDivisorPrecision, err := divisor.GetPrecisionBigInt()

	if err != nil {
		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "binDivisorPrecision, err := divisor.GetPrecisionBigInt()",
				ErrMessage: err.Error(),
			}

	}

	biMaxPrecision := big.NewInt(0).SetUint64(uint64(maxPrecision))

	fracQuo, fracQuoPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(
		binDividend,
		binDividendPrecision,
		binDivisor,
		binDivisorPrecision,
		biMaxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuo, fracQuoPrecision, err := BigIntMathDivide{}.BigIntFracQuotient(\n" +
					"    binDividend,binDividendPrecision,binDivisor,binDivisorPrecision,biMaxPrecision)",
				ErrMessage: err.Error(),
			}
	}

	err = fracQuotient.SetBigInt(fracQuo, uint(fracQuoPrecision.Uint64()))

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "  err = fracQuotient.SetBigInt(fracQuo, uint(fracQuoPrecision.Uint64()))",
				ErrMessage: err.Error(),
			}
	}

	err = fracQuotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fracQuotient.SetNumericSeparatorsDto(numSepsDto)",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, err
}

// BigIntNumFracQuotientArray
//
// Performs a division operation on BigIntNum input parameters
// 'dividends' and 'divisor'. 'dividends' is an array of BigIntNum
// types and the division is operation is performed on each element
// of the array using a single 'divisor'.
//
// The resulting quotients are returned as an array BigIntNum types
// representing the result of each division operation expressed as
// integer and fractional digits. Remember that the BigIntNum type
// specifies 'precision'. Precision is defined as the number of
// fractional digits to the right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotients. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//			Examples
//			=========
//
//		 Note: For all examples maximum precision is specified as '15'.
//
//		                                              Returned
//		   Dividend    divided by  Divisor     =       Array          =    Result
//		   --------    ----------  --------       -----------------       ---------
//
//	      10.5          /         2.5       =   fracQuoArray[0]    =       4.2
//	      10            /         2.5       =   fracQuoArray[1]    =       4
//	      11.5          /         2.5       =   fracQuoArray[2]    =       4.6
//	      2.5           /         2.5       =   fracQuoArray[3]    =       1
//	    -12.555         /         2.5       =   fracQuoArray[4]    =   -   5.022
//	   -  2.5           /         2.5       =   fracQuoArray[5]    =   -   1
//	     12.555         /         2.5       =   fracQuoArray[6]    =       5.022
//	  - 122.783         /         2.5       =   fracQuoArray[7]    =   -  49.1132
//	  -6847.231         /         2.5       =   fracQuoArray[8]    =   -2738.8924
//	  -   2.5           /         2.5       =   fracQuoArray[9]    =   -   1
//	  -  10             /         2.5       =   fracQuoArray[10]   =   -   4
//	  -  10.5           /         2.5       =   fracQuoArray[11]   =   -   4.2
//
// Each element of the returned BigIntNum array resulting from this division operation
// will contain contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from the first element of the input parameter 'dividends'
// array.
//
//			 Numeric Separators
//			 ==================
//
//	    Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	    instance. A NumericSeparatorDto contains symbols or characters
//	    for the decimal separator, thousands separator and currency
//	    symbol. These separators are used when presenting numeric
//	    values in number strings.
//
//	    If any of the 'numSeps' Numeric Separator Components are set
//	    to zero, those components will be automatically reset to USA
//	    default values.
//
//	    The returned value ('fracQuoArray') will be configured with
//	    'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumFracQuotientArray(
	dividends []BigIntNum,
	divisor BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuoArray []BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumFracQuotientArray",
		"")

	if err != nil {
		return fracQuoArray, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuoArray, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	isZero, err := divisor.IsZero()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "isZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if isZero {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor==0",
				ErrMessage: "Error: Attempted divide by zero!\n" +
					"Input parameter 'divisor' has a ZERO value.",
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenAry := len(dividends)

	if lenAry == 0 {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividends)==0",
				ErrMessage: "Error: Input Parameter 'dividends' is an EMPTY Array!",
			}
	}

	fracQuoArray = make([]BigIntNum, lenAry, lenAry+20)

	var dividendsNumStr string

	for i := 0; i < lenAry; i++ {

		err = dividends[i].IsValid(ePrefix.XCpy(
			fmt.Sprintf("Testing dividends[%d]", i)).String())

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = dividends[%d].IsValid(ePrefix)", i),
					ErrContext: fmt.Sprintf("dividends[%d] is INVALID!", i),
					ErrMessage: err.Error(),
				}
		}

		dividendsNumStr, err = dividends[i].GetNumStr()

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("dividendsNumStr, err = dividends[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(dividends[i], divisor)

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewBigIntNum(dividends[%d], divisor)", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair.MaxPrecision = maxPrecision

		fracQuoArray[i], err =
			new(bigIntMathDivideNanobot).
				pairFracQuotient(&bPair, numSeps, ePrefix)

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("fracQuoArray[%d], err = new(bigIntMathDivideNanobot).\n"+
						"pairFracQuotient(&bPair, numSeps, ePrefix)", i),
					ErrContext: fmt.Sprintf("dividend='%v'\ndivisor='%v'\nmaxPrecision='%v'\nIndex='%v'",
						dividendsNumStr, divisorNumStr, bPair.MaxPrecision, i),
					ErrMessage: err.Error(),
				}
		}
	}

	return fracQuoArray, nil
}

// BigIntNumDivideByTwoQuoMod
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value two ('2'). There are two BigIntNum return values:
// 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at mathoopsgo/notes/divmodnote-letter.pdf.
//
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//		q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	 'quotient'  -  The integer result of dividing the 'dividend' by two
//
//	 'modulo'    -  The modulo operation finds the remainder after division
//	                of one number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum precision
// of the resulting 'modulo'. Precision is defined as the number of
// fractional digits to the right of the decimal point. Be advised that
// these calculations can support very large precision values.
//
//	Examples
//	=========
//
//	Dividend    divided by    Divisor    =    Quotient    Modulo/Remainder
//
//	    4            /           2       =        2             0
//	    5            /           2       =        2             1
//	    2.5          /           2       =        1             0.5
//	  -12.555        /           2       =       -6            -0.555
//	    0            /           2       =        0             0
//	  -19            /           2       =       -9            -1
//
// The returned BigIntNum division 'result' (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'
//
//			 Numeric Separators
//			 ==================
//
//	    Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	    instance. A NumericSeparatorDto contains symbols or characters
//	    for the decimal separator, thousands separator and currency
//	    symbol. These separators are used when presenting numeric
//	    values in number strings.
//
//	    If any of the 'numSeps' Numeric Separator Components are set
//	    to zero, those components will be automatically reset to USA
//	    default values.
//
//	    The returned values ('quotient' and 'modulo') will be
//	    configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTwoQuoMod(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTwoQuoMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumTwo, err := new(BigIntNum).NewTwo(0)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumTwo, err := new(BigIntNum).NewTwo(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	quotient, modulo, err = new(BigIntMathDivide).
		BigIntNumQuotientMod(dividend, bINumTwo, numSeps, maxPrecision)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(BigIntMathDivide).BigIntNumQuotientMod(\n" +
					"    dividend, bINumTwo, numSeps, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// BigIntNumDivideByThreeQuoMod
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value three ('3'). There are two BigIntNum return values:
// 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
// Also available at mathopsgo/notes/divmodnote-letter.pdf.
//
//	So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo':
//
//	q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend'
//	               by three.
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as
// the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
//	Examples
//	=========
//
//	Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//	   4         /           3     =      1            1
//	   5         /           3     =      1            2
//	   8         /           3     =      2            2
//	  12         /           3     =      4            0
//
// The returned BigIntNum division 'result' (quotient and modulo) will
// contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter, 'dividend'
//
//			 Numeric Separators
//			 ==================
//
//	    Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	    instance. A NumericSeparatorDto contains symbols or characters
//	    for the decimal separator, thousands separator and currency
//	    symbol. These separators are used when presenting numeric
//	    values in number strings.
//
//	    If any of the 'numSeps' Numeric Separator Components are set
//	    to zero, those components will be automatically reset to USA
//	    default values.
//
//	    The returned values ('quotient' and 'modulo') will be
//	    configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByThreeQuoMod(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByThreeQuoMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumThree, err := new(BigIntNum).NewThree(0)

	if err != nil {
		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumThree, err := new(BigIntNum).NewThree(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	quotient, modulo, err = new(BigIntMathDivide).
		BigIntNumQuotientMod(dividend, bINumThree, numSeps, maxPrecision)

	if err != nil {
		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(BigIntMathDivide).\n" +
					"    BigIntNumQuotientMod(dividend, bINumThree, numSeps, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// BigIntNumDivideByFiveQuoMod
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value five ('5'). There are two BigIntNum return
// values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on
// T-Division (Truncate Division). See "Division and Modulus for
// Computer Scientists", DAAN LEIJEN, University of Utrecht Dept.
// of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// Also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// So for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	  q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend'
//	               by five.
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo' value. Precision is defined
// as the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
//	Examples
//	=========
//
//	Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//	   6         /          5      =      1           1
//	  12         /          5      =      2           2
//	  16         /          5      =      3           1
//
//
//			 Numeric Separators
//			 ==================
//
//	    Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	    instance. A NumericSeparatorDto contains symbols or characters
//	    for the decimal separator, thousands separator and currency
//	    symbol. These separators are used when presenting numeric
//	    values in number strings.
//
//	    If any of the 'numSeps' Numeric Separator Components are set
//	    to zero, those components will be automatically reset to USA
//	    default values.
//
//	    The returned values ('quotient' and 'modulo') will be
//	    configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByFiveQuoMod(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByFiveQuoMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumFive, err := new(BigIntNum).NewFive(0)

	if err != nil {
		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumFive, err := new(BigIntNum).NewFive(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	quotient, modulo, err = bIDivide.BigIntNumQuotientMod(
		dividend, bINumFive, numSeps, maxPrecision)

	if err != nil {
		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = bIDivide.BigIntNumQuotientMod(\n" +
					"    dividend, bINumFive, numSeps, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// BigIntNumDivideByTenQuoMod
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value 10 ('10'). There are two BigIntNum return values:
// 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on
// T-Division (Truncate Division). See "Division and Modulus for
// Computer Scientists", DAAN LEIJEN, University of Utrecht Dept.
// of Computer Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// Also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	  q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend' by 10
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo' value. Precision is defined
// as the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
//		Examples
//		=========
//
//		Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//		   16         /         10     =      1           6
//		   32         /         10     =      3           2
//		   96         /         10     =      9           6
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned values ('quotient' and 'modulo') will be
//	 configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTenQuoMod(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTenQuoMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumTen, err := new(BigIntNum).NewTen(0)

	if err != nil {
		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumTen, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	quotient, modulo, err = bIDivide.
		BigIntNumQuotientMod(dividend, bINumTen, numSeps, maxPrecision)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = bIDivide.BigIntNumQuotientMod(\n" +
					"    dividend, bINumTen, numSeps, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// BigIntNumDivideByTenToPowerQuoMod
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the numeric
// value of ten (10) to the power of input parameter 'exponent'.
// There are two BigIntNum return values: 'quotient' and 'modulo'.
//
//	quotient, modulo = dividend / (10^exponent)
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	  q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend'
//	               by 10.
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo' value. Precision is defined
// as the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
// The returned BigIntNum division 'result' (quotient and modulo) will
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTenToPowerQuoMod(
	dividend BigIntNum,
	exponent BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTenToPowerQuoMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy("Testing dividend").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = dividend.IsValid(ePrefix.XCpy(\"Testing dividend\").String())",
				ErrContext: "Input parameter 'dividend' (BigIntNum) is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	err = exponent.IsValid(ePrefix.XCpy("Testing exponent").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = exponent.IsValid(ePrefix.XCpy(\"Testing exponent\").String())",
				ErrContext: "Input parameter 'exponent' (BigIntNum) is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bigITen, err := new(BigIntNum).NewTen(0)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigITen, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newMaxPrecision := maxPrecision + 10

	scaleValue, err :=
		new(BigIntMathPower).Pwr(bigITen, exponent, newMaxPrecision)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "scaleValue, err := new(BigIntMathPower).Pwr(\n" +
					"    bigITen, exponent, newMaxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	quotient, modulo, err = new(BigIntMathDivide).
		BigIntNumQuotientMod(dividend, scaleValue, numSeps, maxPrecision)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(BigIntMathDivide).\n" +
					"    BigIntNumQuotientMod(dividend, scaleValue, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// BigIntNumDivideByTwoFracQuo
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value two ('2').
//
// The result of this division operation is returned as a BigIntNum
// type representing quotient as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'. Precision
// is defined as the number of fractional digits to the right of
// the decimal place.
//
//	Examples
//	=========
//
//	Dividend  divided by  Divisor  =  Precision  Result
//
//	  10.5        /          2     =      2        5.25
//	  10          /          2     =      0        5
//	 -12.555      /          2     =      4       -6.2775
//
// The input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits.
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTwoFracQuo(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTwoFracQuo",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumTwo, err := new(BigIntNum).NewTwo(0)

	if err != nil {
		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumTwo, err := new(BigIntNum).NewTwo(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fracQuotient, err = new(BigIntMathDivide).
		BigIntNumFracQuotient(dividend, bINumTwo, numSeps, maxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumFracQuotient(dividend, bINumTwo, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// BigIntNumDivideByThreeFracQuo
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value three ('3').
//
// The result of this division operation is returned as a
// BigIntNum type representing quotient as integer and fractional
// digits. Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
//	Examples
//	========
//
//	For this example, assume that 'maxPrecision' = 15
//
//	Dividend  divided by  Divisor  =  Precision  Result
//
//	  9.5         /          3     =      2      3.166666666666667
//	 10           /          3     =      0      3.333333333333333
//	 12           /          3     =      0      4
//	-12           /          3     =      0     -4
//
// The input parameter 'maxPrecision' is used to control the
// maximum precision of the resulting fractional quotient. Be
// advised that this method is capable of calculating quotients
// with very long strings of fractional digits.
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByThreeFracQuo(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByThreeFracQuo",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumThree, err := new(BigIntNum).NewThree(0)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumThree, err := new(BigIntNum).NewThree(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fracQuotient, err = new(BigIntMathDivide).
		BigIntNumFracQuotient(dividend, bINumThree, numSeps, maxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumFracQuotient(dividend, bINumThree, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// BigIntNumDivideByFiveFracQuo
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value five ('5').
//
// The result of this division operation is returned as a BigIntNum
// type representing quotient as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'. Precision
// is defined as the number of fractional digits to the right of
// the decimal place.
//
//	Examples
//	========
//
//	Dividend  divided by  Divisor  =  Precision  Result
//
//	 16.2          /        5      =      2        3.24
//	 10            /        5      =      0        2
//	 12            /        5      =      1        2.4
//	-12            /        5      =      1       -2.4
//
// The input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits.
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByFiveFracQuo(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByFiveFracQuo",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumFive, err := new(BigIntNum).NewFive(0)

	if err != nil {
		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumFive, err := new(BigIntNum).NewFive(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fracQuotient, err = new(BigIntMathDivide).
		BigIntNumFracQuotient(dividend, bINumFive, numSeps, maxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).(\n" +
					"    BigIntNumFracQuotient(dividend, bINumFive, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// BigIntNumDivideByTenFracQuo
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the integer
// numeric value ten ('10').
//
// The result of this division operation is returned as a BigIntNum
// type representing quotient as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
//	Examples
//	=========
//
//	For this example, assume that 'maxPrecision' = 15
//
//	Dividend  divided by  Divisor  =  Precision  Result
//
//	  16.2         /         10    =       2      1.62
//	  10           /         10    =       0      1
//	  12           /         10    =       1      1.2
//	 -12           /         10    =       1     -1.2
//
// The input parameter 'maxPrecision' is used to control the
// maximum precision of the resulting fractional quotient. Be
// advised that this method is capable of calculating quotients
// with very long strings of fractional digits.
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTenFracQuo(
	dividend BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTenFracQuo",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINumTen, err := new(BigIntNum).NewTen(0)

	if err != nil {
		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumTen, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fracQuotient, err = new(BigIntMathDivide).
		BigIntNumFracQuotient(dividend, bINumTen, numSeps, maxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).(\n" +
					"    BigIntNumFracQuotient(dividend, bINumTen, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// BigIntNumDivideByTenToPowerFracQuo
//
// Performs a division operation on BigIntNum input parameter,
// 'dividend'. This method will divide 'dividend' by the numeric
// value ten ('10') to the power of input parameter 'exponent'.
//
//	fractional quotient = dividend / (10^exponent)
//
// The result of this division operation is returned as a BigIntNum
// type representing quotient as integer and fractional digits.
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits.
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTenToPowerFracQuo(
	dividend BigIntNum,
	exponent BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTenToPowerFracQuo",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	err = exponent.IsValid(ePrefix.XCpy(
		"Testing exponent").String())

	if err != nil {
		return fracQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	fracQuotient, err = new(BigIntNum).NewZero(0)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bIntNum2, err := new(BigIntNum).NewTen(0)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bIntNum2, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newMaxPrecision := maxPrecision + 10

	scaleValue, err :=
		new(BigIntMathPower).Pwr(bIntNum2, exponent, newMaxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "scaleValue, err := new(BigIntMathPower).Pwr(bIntNum2, exponent, newMaxPrecision)",
				ErrContext: fmt.Sprintf("newMaxPrecision= '%v'", newMaxPrecision),
				ErrMessage: err.Error(),
			}
	}

	fracQuotient, err = new(BigIntMathDivide).
		BigIntNumFracQuotient(dividend, scaleValue, numSeps, maxPrecision)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).(\n" +
					"    BigIntNumFracQuotient(dividend, scaleValue, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// BigIntNumDivideByTenToPowerIntQuo
//
// Performs a division operation on BigIntNum input parameter
// 'dividend'. This method will divide 'dividend' by the numeric
// value ten ('10') to the power of input parameter 'exponent'.
//
//	integer quotient = dividend / (10^exponent)
//
// The result of this division operation is returned as a BigIntNum
// type representing 'quotient' as an integer value.
//
// The input parameter 'maxPrecision' is used to control the maximum
// precision of the internal calculations for those cases with
// 'exponent' is a fractional or floating point number. Be advised
// that this method is capable of calculating quotients with very
// long strings of fractional digits.
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('intQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTenToPowerIntQuo(
	dividend BigIntNum,
	exponent BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (intQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTenToPowerIntQuo",
		"")

	if err != nil {
		return intQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return intQuotient, err
	}

	err = exponent.IsValid(ePrefix.XCpy(
		"Testing exponent").String())

	if err != nil {
		return intQuotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINum10, err := new(BigIntNum).NewTen(0)

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINum10, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newScalePrecision := maxPrecision + 10

	scaleValue, err := new(BigIntMathPower).
		Pwr(bINum10, exponent, newScalePrecision)

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "scaleValue, err := new(BigIntMathPower).\n" +
					"    Pwr(bINum10, exponent, newScalePrecision)",
				ErrContext: fmt.Sprintf("exponent= '%v' newScalePrecision= '%v'",
					exponent, newScalePrecision),
				ErrMessage: err.Error(),
			}
	}

	intQuotient, err = new(BigIntMathDivide).
		BigIntNumIntQuotient(dividend, scaleValue, numSeps)

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "intQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumIntQuotient(dividend, scaleValue)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return intQuotient, nil
}

// BigIntNumDivideByTenToPowerMod
//
// Performs a modulo operation on BigIntNum input parameters
// 'dividend' and ten (10) to the power of exponent.
//
//	modulo = dividend / (10^exponent)
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one BigIntNum value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo' value. Precision is defined
// as the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
//	Examples
//	=========
//
//	  Dividend  mod by    Power    Divisor  =  Modulo/Remainder
//
//	  1200.555     %        2        100    =         0.555
//	 10235.555     %        3       1000    =       235.555
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('modulo') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) BigIntNumDivideByTenToPowerMod(
	dividend BigIntNum,
	exponent BigIntNum,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.BigIntNumDivideByTenToPowerMod",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = exponent.IsValid(ePrefix.XCpy(
		"Testing exponent").String())

	if err != nil {
		return modulo, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bINum10, err := new(BigIntNum).NewTen(0)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINum10, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newScalePrecision := maxPrecision + 10

	scaleValue, err := new(BigIntMathPower).
		Pwr(bINum10, exponent, newScalePrecision)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "scaleValue, err := new(BigIntMathPower).\n" +
					"    Pwr(bINum10, exponent, newScalePrecision)",
				ErrContext: fmt.Sprintf("exponent= '%v' newScalePrecision= '%v'",
					exponent, newScalePrecision),
				ErrMessage: err.Error(),
			}
	}

	modulo, err = new(BigIntMathDivide).
		BigIntNumModulo(dividend, scaleValue, numSeps, maxPrecision)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(BigIntMathDivide).\n" +
					"    BigIntNumModulo(dividend, scaleValue, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// DecimalQuotientMod
//
// Performs a division operation on Decimal type input parameters,
// 'dividend' and 'divisor'.
//
// There are two BigIntNum Type return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//		q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	 'quotient'  -  The integer result of dividing the 'dividend' by
//	                the 'divisor'.
//
//	 'modulo'    -  The modulo operation finds the remainder after
//	                division of one number by another.
//	                (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point. Be
// advised that these calculations can support very large precision
// values. Therefore, the user is advised to set a relevant value
// for 'maxPrecision'.
//
//		 Examples
//		 =========
//
//		 Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//		  12.555        /         2.5   =       5          0.055
//		  12.555        /         2     =       6          0.555
//		  2.5           /        12.555 =       0          2.5
//		 -12.555        /         2.5   =      -5         -0.055
//		 -12.555        /         2     =      -6         -0.555
//		 - 2.5          /        12.555 =       0         -2.5
//		  12.555        /       - 2.5   =      -5          0.055
//		  12.555        /       - 2     =      -6          0.555
//		   2.5          /       -12.555 =       0          2.5
//		 -12.555        /       - 2.5   =       5         -0.055
//		 -12.555        /       - 2     =       6         -0.555
//		 - 2.5          /       -12.555 =       0         -2.5
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned values ('quotient' and 'modulo') will be
//	 configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) DecimalQuotientMod(
	dividend Decimal,
	divisor Decimal,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.DecimalQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return quotient, modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, err =
		new(bigIntMathDivideNanobot).
			pairQuotientMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairQuotientMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision='%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// DecimalFracQuotient
//
// Performs a division operation on Decimal Type input parameters
// 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type
// representing the result of the division operation expressed as
// integer and fractional digits.
//
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//	Examples
//	=========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	                                       Quotient
//	Dividend  divided by  Divisor  =  BigIntNum Integer  Precision  Result
//
//	  10.5        /         2      =   525                  2        5.25
//	  10          /         2      =     5                  0        5
//	  11.5        /         2.5    =    46                  1        4.6
//	  2.5         /        12.555  =    199123855037834    15        0.199123855037834
//	-12.555       /         2.5    =   -5022                3       -5.022
//	-12.555       /         2      =   -62775               4       -6.2775
//	- 2.5         /        12.555  =   -199123855037834    15       -0.199123855037834
//	 12.555       /       - 2.5    =   -5022                3       -5.022
//	 12.555       /       - 2      =   -62775               4       -6.2775
//	  2.5         /       -12.555  =   -199123855037834    15       -0.199123855037834
//	-12.555       /       - 2.5    =    5022                3        5.022
//	-12.555       /       - 2      =    62775               4        6.2775
//	- 2.5         /       -12.555  =    199123855037834    15        0.199123855037834
//	-10           /       - 2      =    5                   5        5.00000
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) DecimalFracQuotient(
	dividend Decimal,
	divisor Decimal,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.DecimalFracQuotient",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuotient, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIsZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a zero value.",
			}
	}

	bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, err = new(bigIntMathDivideNanobot).
		pairFracQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"pairFracQuotient(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// DecimalFracQuotientArray
//
// Performs a division operation on Decimal input parameters
// 'dividends' and 'divisor'. 'dividends' is an array of Decimal
// Types. The division operation is performed on each element of
// the 'dividends' array using a single 'divisor'.
//
// The resulting quotients are returned as an array of Decimal
// Types. The values represent result of each division operation
// expressed as integer and fractional digits.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Precision is
// defined as the number of numeric digits to the right of the
// decimal point. Be advised that this method is capable of
// calculating quotients with very long strings of fractional
// digits. Therefore, the user is advised to set a relevant value
// for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	                                                       Returned
//	Dividend  divided by  Divisor  =   Array             =   Result
//
//	   10.5        /       2.5    =   fracQuoArray[0]   =       4.2
//	   10          /       2.5    =   fracQuoArray[1]   =       4
//	   11.5        /       2.5    =   fracQuoArray[2]   =       4.6
//	    2.5        /       2.5    =   fracQuoArray[3]   =       1
//	  -12.555      /       2.5    =   fracQuoArray[4]   =      -5.022
//	   -2.5        /       2.5    =   fracQuoArray[5]   =      -1
//	   12.555      /       2.5    =   fracQuoArray[6]   =       5.022
//	 -122.783      /       2.5    =   fracQuoArray[7]   =     -49.1132
//	-6847.231      /       2.5    =   fracQuoArray[8]   =   -2738.8924
//	   -2.5        /       2.5    =   fracQuoArray[9]   =      -1
//	  -10          /       2.5    =   fracQuoArray[10]  =      -4
//	  -10.5        /       2.5    =   fracQuoArray[11]  =      -4.2
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuoArray') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) DecimalFracQuotientArray(
	dividends []Decimal,
	divisor Decimal,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuoArray []Decimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.DecimalFracQuotientArray",
		"")

	if err != nil {
		return fracQuoArray, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuoArray, err
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIsZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a zero value.",
			}
	}

	lenAry := len(dividends)

	if lenAry == 0 {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividends) == 0",
				ErrMessage: "Error: Input Parameter 'dividends' is an EMPTY Array!",
			}
	}

	fracQuoArray = make([]Decimal, lenAry, lenAry+20)

	var dividendsNumStr string

	for i := 0; i < lenAry; i++ {

		dividendsNumStr, err = dividends[i].GetNumStr()

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("dividendsNumStr, err = dividends[%v].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewDecimal(dividends[i], divisor)

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewDecimal(dividends[%d], divisor)", i),
					ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'",
						dividendsNumStr, dividendsNumStr),
					ErrMessage: err.Error(),
				}
		}

		bPair.MaxPrecision = maxPrecision

		bINum, err := new(bigIntMathDivideNanobot).
			pairFracQuotient(&bPair, numSeps, ePrefix)

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bINum, err := new(bigIntMathDivideNanobot).\n" +
						"pairFracQuotient(&bPair, numSeps, ePrefix)",
					ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision='%v'",
						dividendsNumStr, dividendsNumStr, maxPrecision),
					ErrMessage: err.Error(),
				}
		}

		fracQuoArray[i], err = bINum.GetDecimal()

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("fracQuoArray[%d], err = bINum.GetDecimal()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return fracQuoArray, err
}

// DecimalModulo
//
// Performs a modulo operation on Decimal type input parameters,
// 'dividend' and 'divisor'. The result of this division operation
// is returned as a type BigIntNum.
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one Decimal value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo' value. Precision is defined
// as the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
//		Examples
//		=========
//
//	 Dividend  mod by  Divisor    =  Modulo/Remainder
//
//	  12.555     %       2.5      =      0.055
//	  12.555     %       2        =      0.555
//	   2.5       %      12.555    =      2.5
//	             %                =
//	 -12.555     %       2.5      =     -0.055
//	 -12.555     %       2        =     -0.555
//	  -2.5       %      12.555    =     -2.5
//	             %                =
//	  12.555     %      -2.5      =      0.055
//	  12.555     %      -2        =      0.555
//	   2.5       %     -12.555    =      2.5
//	             %                =
//	 -12.555     %      -2.5      =     -0.055
//	 -12.555     %      -2        =     -0.555
//	  -2.5       %     -12.555    =     -2.5
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned value ('modulo') will be configured with
//	 'numSeps'.
func (bIDivide *BigIntMathDivide) DecimalModulo(
	dividend Decimal,
	divisor Decimal,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.DecimalModulo",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted to mod by zero!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	modulo, err = new(bigIntMathDivideNanobot).pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
					"pairMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision='%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// DecimalModuloToDecimal
//
// Performs a modulo operation on Decimal input parameters
// 'dividend' and 'divisor'.
//
// The modulo operation finds the remainder after division of
// one number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one Decimal value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large
// precision values.
//
// The difference between this method and BigIntMathDivide.DecimalModulo()
// above, is that this method returns the resulting modulo value as
// a type 'Decimal'.
//
//		 Examples
//		 ========
//
//		 Dividend  mod by  Divisor    =    Modulo/Remainder
//
//	   12.555      %       2.5     =        0.055
//	   12.555      %       2       =        0.555
//	    2.5        %      12.555   =        2.5
//	  -12.555      %       2.5     =       -0.055
//	  -12.555      %       2       =       -0.555
//	   -2.5        %      12.555   =       -2.5
//	   12.555      %      -2.5     =        0.055
//	   12.555      %      -2       =        0.555
//	    2.5        %     -12.555   =        2.5
//	  -12.555      %      -2.5     =       -0.055
//	  -12.555      %      -2       =       -0.555
//	   -2.5        %     -12.555   =       -2.5
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned value ('modulo') will be configured with
//	 'numSeps'.
func (bIDivide *BigIntMathDivide) DecimalModuloToDecimal(
	dividend Decimal,
	divisor Decimal,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo Decimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.DecimalModuloToDecimal",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewDecimal(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted to mod by zero!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	bINumModulo, err := new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINumModulo, err := new(bigIntMathDivideNanobot).\n" +
					"pairMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision='%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	modulo, err = bINumModulo.GetDecimal()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "modulo, err = bINumModulo.GetDecimal()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// FixedDecimalFracQuotient
//
// Performs a division operation on objects of type
// BigIntFixedDecimal. The result is also returned a type of
// BigIntFixedDecimal.
//
//				Examples
//				========
//
//				This division operation will produce a quotient which may
//				include a fixed length floating point number:
//
//				  quotient = dividend / divisor
//
//				The BigIntFixedDecimal structure is defined as
//
//				type BigIntFixedDecimal struct {
//
//				  integerNum *big.Int  -  All the numeric digits, both integer and fractional,
//				                          necessary to define a fixed length floating point number.
//				                          The number of digits to the right of the decimal place
//				                          is specified by the data field,
//				                          BigIntFixedDecimal.precision.
//
//				  precision  uint      -  Specifies the number of digits to the right of the decimal
//				                          place in the series of numeric digits represented by data
//				                          field BigIntFixedDecimal.integerNum.
//
//				}
//
//				To represent the floating point number 52.459
//				a BigIntDecimal Structure would be configured as follows:
//
//				      BigIntFixedDecimal.integerNum	= 52459
//
//				      BigIntFixedDecimal.precision	= 3
//
//				Consider the following division example:
//
//				      quotient =	752.314 / 21.67894
//
//			 'dividend' and 'divisor' would be configured as follows:
//
//			       dividend.integerNum = 752314
//			       dividend.precision  = 3
//			       divisor.integerNum  = 2167894
//			       divisor.precision   = 5
//
//			 Assuming a 'maxPrecision' value of '30', the quotient would
//			 be calculated as follows:
//
//	          Decimal Digit Count-> XX123456789012345678921234567893
//			       quotient.integerNum = 34702526968569496479071393712054
//			       quotient.precision  = 30
//
//		 Input Parameters
//		 ================
//
//		 dividend        BigIntFixedDecimal
//
//		 The 'dividend' value will be divided by the 'divisor' to
//		 produce a 'quotient'.
//
//
//		 divisor         BigIntFixedDecimal
//
//		 The 'dividend' value will be divided by the 'divisor' to
//		 produce a 'quotient'.
//
//
//		 numSeps         NumericSeparatorDto
//
//		 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//		 instance. A NumericSeparatorDto structure contains symbols
//		 or characters (runes) for the decimal separator, thousands
//		 separator and currency symbol. These separators are used
//		 when presenting numeric values in number strings.
//
//		 If any of the 'numSeps' Numeric Separator Components are set
//		 to zero, those components will be automatically reset to USA
//		 default values.
//
//		 The returned values ('quotient' and 'modulo') will be
//		 configured with 'numSeps'.
//
//
//		 'maxPrecision'  uint
//
//		 Maximum precision will determine the maximum number of decimal
//		 digits to which the result or 'quotient' Will be calculated
//		 and returned to the caller. The quotient may consist of actual
//		 fractional digits which number less than 'maxPrecision'.
//		 However, if the number of digits to the right of the decimal
//		 place exceeds 'maxPrecision', the returned quotient will be
//		 rounded to 'maxPrecision' fractional digits to the right of
//		 the decimal place.
//
//		 Return Values
//		 =============
//
//		 quotient        BigIntFixedDecimal
//
//		 The result of the division operation expressed as a type
//		 BigIntFixedDecimal.
//
//
//		 err             error
//
//		 If an error is encountered, this function will return a
//		 quotient set equal to zero and an error object will be
//		 returned containing an appropriate error message. If the
//		 function completes the division operation successfully,
//		 the returned 'quotient' will be populated with the
//		 correct result, and 'err' will be set equal to 'nil'.
func (bIDivide *BigIntMathDivide) FixedDecimalFracQuotient(
	dividend BigIntFixedDecimal,
	divisor BigIntFixedDecimal,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntFixedDecimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.FixedDecimalFracQuotient",
		"")

	if err != nil {
		return quotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return quotient, err
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	dividendInt, err := dividend.GetIntegerValue()

	if err != nil {

		return quotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendInt, err := dividend.GetIntegerValue()",
				ErrContext: "dividend is type BigIntFixedDecimal",
				ErrMessage: err.Error(),
			}
	}

	dividendPrecision := dividend.GetPrecisionBigInt()

	divisorInt, err := divisor.GetIntegerValue()

	if err != nil {

		return quotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorInt, err := divisor.GetIntegerValue()",
				ErrContext: "divisor is type BigIntFixedDecimal",
				ErrMessage: err.Error(),
			}
	}

	divisorPrecision := divisor.GetPrecisionBigInt()

	bigIntMaxPrecision := big.NewInt(0).SetUint64(uint64(maxPrecision))

	result, resultPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			dividendInt,
			dividendPrecision,
			divisorInt,
			divisorPrecision,
			bigIntMaxPrecision)

	if err != nil {

		return quotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "result, resultPrecision, err := new(BigIntMathDivide).BigIntFracQuotient(\n" +
					"    dividendInt, dividendPrecision, divisorInt, divisorPrecision,bigIntMaxPrecision)",
				ErrContext: fmt.Sprintf("dividendInt= '%v'; dividendPrecision= '%v';\n"+
					"divisorInt= '%v'\ndivisorPrecision= '%v'; bigIntMaxPrecision= '%v'",
					dividendInt.Text(10), dividendPrecision.Text(10),
					divisorInt.Text(10), divisorPrecision.Text(10), bigIntMaxPrecision.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	err = quotient.SetNumericValue(result, uint(resultPrecision.Uint64()))

	if err != nil {

		return quotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = quotient.SetNumericValue(\n" +
					"result, uint(resultPrecision.Uint64())",
				ErrContext: fmt.Sprintf("result= '%v'; resultPrecision= '%v'",
					result.Text(10), resultPrecision.Uint64()),
				ErrMessage: err.Error(),
			}
	}

	err = quotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return quotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = quotient.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, nil
}

// IntAryQuotientMod
//
// Performs a division operation on IntAry type input parameters,
// 'dividend' and 'divisor'.
//
// There are two BigIntNum Type return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or
// 'modulo':
//
//		q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	 'quotient' - The integer result of dividing the 'dividend' by
//	              the 'divisor'
//
//	 'modulo'   - The modulo operation finds the remainder after
//	              division of one number by another.
//	              (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo' value. Precision is defined
// as the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very large
// precision values.
//
//		Examples
//		========
//
//		Dividend  divided by  Divisor   =  Quotient  Modulo/Remainder
//
//		 12.555       /          2.5    =     5           0.055
//		 12.555       /          2      =     6           0.555
//		  2.5         /         12.555  =     0           2.5
//		-12.555       /          2.5    =    -5          -0.055
//		-12.555       /          2      =    -6          -0.555
//		 -2.5         /         12.555  =     0          -2.5
//		 12.555       /         -2.5    =    -5           0.055
//		 12.555       /         -2      =    -6           0.555
//		  2.5         /        -12.555  =     0           2.5
//		-12.555       /         -2.5    =     5          -0.055
//		-12.555       /         -2      =     6          -0.555
//		 -2.5         /        -12.555  =     0          -2.5
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned values ('quotient' and 'modulo') will be
//	 configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) IntAryQuotientMod(
	dividend IntAry,
	divisor IntAry,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.IntAryQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return quotient, modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, err =
		new(bigIntMathDivideNanobot).
			pairQuotientMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairQuotientMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// IntAryFracQuotient
//
// Performs a division operation on IntAry Type input parameters
// 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type
// representing the result of the division operation expressed as
// integer and fractional digits.
//
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	                                     Quotient
//	Dividend  divided by  Divisor   =  BigIntNum Integer  Precision Result
//
//	 10.5         /         2       =    525                   2     5.25
//	 10           /         2       =    5                     0     5
//	 11.5         /         2.5     =    46                    1     4.6
//	  2.5         /        12.555   =    199123855037834      15     0.199123855037834
//	 12.555       /         2.5     =   -5022                  3    -5.022
//	 12.555       /         2       =   -62775                 4    -6.2775
//	 -2.5         /        12.555   =   -199123855037834      15    -0.199123855037834
//	 12.555       /        -2.5     =   -5022                  3    -5.022
//	 12.555       /        -2       =   -62775                 4    -6.2775
//	  2.5         /       -12.555   =   -199123855037834      15    -0.199123855037834
//	 12.555       /        -2.5     =    5022                  3     5.022
//	 12.555       /        -2       =    62775                 4     6.2775
//	 -2.5         /       -12.555   =    199123855037834       15    0.199123855037834
//	 10           /        -2       =    5                      5    5.00000
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) IntAryFracQuotient(
	dividend IntAry,
	divisor IntAry,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.IntAryFracQuotient",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuotient, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	divisorIszero, err := divisor.IsZero()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIszero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIszero {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a ZERO value.",
			}
	}

	bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend='%v'; divisor='%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, err = new(bigIntMathDivideNanobot).
		pairFracQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"    new(bigIntMathDivideNanobot).",
				ErrContext: fmt.Sprintf("dividend='%v'; divisor='%v'; maxPrecision='%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	err = fracQuotient.TrimTrailingFracZeros()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fracQuotient.TrimTrailingFracZeros()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// IntAryFracQuotientArray
//
// Performs a division operation on IntAry input parameters
// 'dividends' and 'divisor'. 'dividends' is an array of IntAry
// Types. The division operation is performed on each element
// of the 'dividends' array using a single 'divisor'.
//
// The resulting quotients are returned as an array of IntAry
// Types. The values represent result of each division operation
// expressed as integer and fractional digits.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Precision is
// defined as the number of numeric digits to the right of the
// decimal point. Be advised that this method is capable of
// calculating quotients with very long strings of fractional
// digits. Therefore, the user is advised to set a relevant
// value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	           divided                 Returned
//	Dividend     by      Divisor   =     Array           =   Result
//
//	   10.5        /        2.5     =   fracQuoArray[0]  =      4.2
//	   10          /        2.5     =   fracQuoArray[1]  =      4
//	   11.5        /        2.5     =   fracQuoArray[2]  =      4.6
//	    2.5        /        2.5     =   fracQuoArray[3]  =      1
//	  -12.555      /        2.5     =   fracQuoArray[4]  =      5.022
//	   -2.5        /        2.5     =   fracQuoArray[5]  =      1
//	   12.555      /        2.5     =   fracQuoArray[6]  =      5.022
//	 -122.783      /        2.5     =   fracQuoArray[7]  =    -49.1132
//	-6847.231      /        2.5     =   fracQuoArray[8]  =  -2738.8924
//	   -2.5        /        2.5     =   fracQuoArray[9]  =     -1
//	  -10          /        2.5     =   fracQuoArray[10] =     -4
//	  -10.5        /        2.5     =   fracQuoArray[11] =     -4.2
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuoArray') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) IntAryFracQuotientArray(
	dividends []IntAry,
	divisor IntAry,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuoArray []IntAry, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.IntAryFracQuotientArray",
		"")

	if err != nil {
		return fracQuoArray, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuoArray, err
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorIszero, err := divisor.IsZero()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIszero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIszero {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a ZERO value.",
			}
	}

	lenAry := len(dividends)

	if lenAry == 0 {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividends) == 0",
				ErrMessage: "Error: Input Parameter 'dividends' is an EMPTY Array!",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	var dividendsNumStr string

	fracQuoArray = make([]IntAry, lenAry, lenAry+20)

	for i := 0; i < lenAry; i++ {

		dividendsNumStr, err = dividends[i].GetNumStr()

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("dividendsNumStr, err = dividends[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewIntAry(dividends[i], divisor)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewIntAry(dividends[%d], divisor)", i),
					ErrContext: fmt.Sprintf("dividends[%d]= '%v'; divisor= '%v'",
						i, dividendsNumStr, divisorNumStr),
					ErrMessage: err.Error(),
				}
		}

		bPair.MaxPrecision = maxPrecision

		bINum, err := new(bigIntMathDivideNanobot).
			pairFracQuotient(&bPair, numSeps, ePrefix)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bINum, err := new(bigIntMathDivideNanobot).\n" +
						"pairFracQuotient(&bPair, numSeps, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		fracQuoArray[i], err = bINum.GetIntAry()

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("fracQuoArray[%d], err = bINum.GetIntAry()",
						i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return fracQuoArray, err
}

// IntAryModulo
//
// Performs a modulo operation on IntAry input parameters 'dividend'
// and 'divisor'.
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one BigIntNum value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as
// the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very
// large precision values. Therefore, the user is advised to set a relevant
// value for 'maxPrecision'.
//
//	Examples
//	========
//
//	 Dividend  mod by  Divisor  =  Modulo/Remainder
//
//	  12.555     %      2.5     =       0.055
//	  12.555     %      2       =       0.555
//	   2.5       %     12.555   =       2.5
//	 -12.555     %      2.5     =      -0.055
//	 -12.555     %      2       =      -0.555
//	  -2.5       %     12.555   =      -2.5
//	  12.555     %     -2.5     =       0.055
//	  12.555     %      2       =       0.555
//	   2.5       %    -12.555   =       2.5
//	 -12.555     %      2.5     =      -0.055
//	 -12.555     %     -2       =      -0.555
//	  -2.5       %    -12.555   =      -2.5
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned value ('modulo') will be configured with
//	 'numSeps' Numeric Separators.
func (bIDivide *BigIntMathDivide) IntAryModulo(
	dividend IntAry,
	divisor IntAry,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.IntAryModulo",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	modulo, err = new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// IntAryModuloToIntAry
//
// Performs a modulo operation on IntAry input parameters
// 'dividend' and 'divisor'.
//
// The modulo operation finds the remainder after division of
// one number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one IntAry value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This informatino is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point. Be
// advised that these calculations can support very large precision
// values.
//
// The difference between this method and BigIntMathDivide.IntAryModulo()
// above, is that the division result, 'modulo', is returned as Type
// IntAry.
//
//	Examples
//	========
//
//	Dividend    mod by    Divisor    =    Modulo/Remainder
//
//	 12.555       %         2.5      =         0.055
//	 12.555       %         2        =         0.555
//	  2.5         %        12.555    =         2.5
//	-12.555       %         2.5      =        -0.055
//	-12.555       %         2        =        -0.555
//	 -2.5         %        12.555    =        -2.5
//	 12.555       %        -2.5      =         0.055
//	 12.555       %        -2        =         0.555
//	  2.5         %       -12.555    =         2.5
//	-12.555       %        -2.5      =        -0.055
//	-12.555       %        -2        =        -0.555
//	 -2.5         %       -12.555    =        -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('modulo') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) IntAryModuloToIntAry(
	dividend IntAry,
	divisor IntAry,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo IntAry, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.IntAryModuloToIntAry",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	bINumModulo, err := new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINumModulo, err := new(bigIntMathDivideNanobot).\n" +
					"    pairMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	modulo, err = bINumModulo.GetIntAry()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "modulo, err = bINumModulo.GetIntAry()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// INumMgrQuotientMod
//
// Performs a division operation on types implementing the INumMgr
// interface. Input parameters, 'dividend' and 'divisor' must
// therefore implement the INumMgr interface.
//
// There are two BigIntNum Type return values: 'quotient' and
// 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	  q = D div d = f(D/d)
//
//	  r = D mod d = D − d ·q
//
//	'quotient' -  The integer result of dividing the 'dividend' by
//	              the 'divisor'.
//
//	'modulo'   -  The modulo operation finds the remainder after
//	              division of one number by another.
//	              (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large
// precision values. Therefore, the user is advised to set a
// relevant value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Dividend  divided by  Divisor    =  Quotient  Modulo/Remainder
//
//	 12.555        /         2.5     =     5           0.055
//	 12.555        /         2       =     6           0.555
//	  2.5          /        12.555   =     0           2.5
//	-12.555        /         2.5     =    -5          -0.055
//	-12.555        /         2       =    -6          -0.555
//	 -2.5          /        12.555   =     0          -2.5
//	 12.555        /        -2.5     =    -5           0.055
//	 12.555        /        -2       =    -6           0.555
//	  2.5          /       -12.555   =     0           2.5
//	-12.555        /        -2.5     =     5          -0.055
//	-12.555        /        -2       =     6          -0.555
//	 -2.5          /       -12.555   =     0          -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) INumMgrQuotientMod(
	dividend INumMgr,
	divisor INumMgr,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.INumMgrQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return quotient, modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return quotient, modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewINumMgr(dividend, divisor)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter bPair.Big2 has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, err = new(bigIntMathDivideNanobot).
		pairQuotientMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairQuotientMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// INumMgrFracQuotient
//
// Performs a division operation on input parameters 'dividend'
// and 'divisor' which implement the INumMgr interface.
//
// The resulting quotient is returned as a BigIntNum type
// representing the result of the division operation expressed
// as integer and fractional digits.
//
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to
// the right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Be advised
// that this method is capable of calculating quotients with
// very long strings of fractional digits. Therefore, the user
// is advised to set a relevant value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	          divided                 Quotient
//	Dividend    by    Divisor  =  BigIntNum Integer  Precision  Result
//
//	  10.5      /       2      =    525                 2        5.25
//	  10        /       2      =    5                   0        5
//	  11.5      /       2.5    =    46                  1        4.6
//	   2.5      /      12.555  =    199123855037834    15        0.199123855037834
//	 -12.555    /       2.5    =   -5022                3       -5.022
//	 -12.555    /       2      =   -62775               4       -6.2775
//	  -2.5      /      12.555  =   -199123855037834    15       -0.199123855037834
//	  12.555    /      -2.5    =   -5022                3       -5.022
//	  12.555    /      -2      =   -62775               4       -6.2775
//	   2.5      /     -12.555  =   -199123855037834    15       -0.199123855037834
//	 -12.555    /      -2.5    =    5022                3        5.022
//	 -12.555    /      -2      =    62775               4        6.2775
//	  -2.5      /     -12.555  =    199123855037834    15        0.199123855037834
//	 -10        /      -2      =    5                   5        5.00000
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) INumMgrFracQuotient(
	dividend INumMgr,
	divisor INumMgr,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.INumMgrFracQuotient",
		"")

	if err != nil {
		return fracQuotient, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return fracQuotient, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuotient, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorIszero, err := divisor.IsZero()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIszero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIszero {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter divisor has a ZERO value.",
			}
	}

	// Validity tests are performed on 'dividend' and 'divisor'
	bPair, err := new(BigIntPair).NewINumMgr(dividend, divisor)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, err = new(bigIntMathDivideNanobot).
		pairFracQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"    pairFracQuotient(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// INumMgrFracQuotientArray
//
// Performs a division operation on input parameters 'dividends'
// and 'divisor' which implement the INumMgr interface.
//
// 'dividends' is an array of types implementing the INumMgr
// interface. The division operation is performed on each element
// of the 'dividends' array using a single 'divisor'.
//
// The resulting quotients are returned as an array of types
// implementing the INumMgr Interface. However, be aware that
// the single underlying type is BigIntNum. The returned values
// represent the results of each division operation expressed
// as integer and fractional digits.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Precision is
// defined as the number of numeric digits to the right of the
// decimal point. Be advised that this method is capable of
// calculating quotients with very long strings of fractional
// digits. Therefore, the user is advised to set a relevant
// value for 'maxPrecision'.
//
//	Examples
//	=========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	            divided              Returned
//	Dividend      by     Divisor  =    Array          =    Result
//
//	   10.5       /        2.5    =  fracQuoArray[0]  =       4.2
//	   10         /        2.5    =  fracQuoArray[1]  =       4
//	   11.5       /        2.5    =  fracQuoArray[2]  =       4.6
//	    2.5       /        2.5    =  fracQuoArray[3]  =       1
//	  -12.555     /        2.5    =  fracQuoArray[4]  =      -5.022
//	   -2.5       /        2.5    =  fracQuoArray[5]  =      -1
//	   12.555     /        2.5    =  fracQuoArray[6]  =       5.022
//	 -122.783     /        2.5    =  fracQuoArray[7]  =     -49.1132
//	-6847.231     /        2.5    =  fracQuoArray[8]  =   -2738.8924
//	   -2.5       /        2.5    =  fracQuoArray[9]  =      -1
//	  -10         /        2.5    =  fracQuoArray[10] =      -4
//	  -10.5       /        2.5    =  fracQuoArray[11] =      -4.2
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuoArray') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) INumMgrFracQuotientArray(
	dividends []INumMgr,
	divisor INumMgr,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuoArray []INumMgr, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.INumMgrFracQuotientArray",
		"")

	if err != nil {
		return fracQuoArray, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuoArray, err
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorIszero, err := divisor.IsZero()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIszero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIszero {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a ZERO value.",
			}
	}

	lenAry := len(dividends)

	if lenAry == 0 {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividends) == 0",
				ErrMessage: "Error: Input Parameter 'dividends' is an EMPTY Array!",
			}
	}

	fracQuoArray = make([]INumMgr, lenAry, lenAry+20)

	var dividendsNumStr string

	for i := 0; i < lenAry; i++ {

		dividendsNumStr, err = dividends[i].GetNumStr()

		if err != nil {

			return []INumMgr{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("dividendsNumStr, err = dividends[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewINumMgr(dividends[i], divisor)

		if err != nil {

			return []INumMgr{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewINumMgr(dividends[%d], divisor)", i),
					ErrContext: fmt.Sprintf("dividends[%d]= '%v'; divisor= '%v';",
						i, dividendsNumStr, divisorNumStr),
					ErrMessage: err.Error(),
				}
		}

		bPair.MaxPrecision = maxPrecision

		bINum, err := new(bigIntMathDivideNanobot).
			pairFracQuotient(&bPair, numSeps, ePrefix)

		if err != nil {

			return []INumMgr{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bINum, err := new(bigIntMathDivideNanobot).\n" +
						"pairFracQuotient(&bPair, numSeps, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		fracQuoArray[i] = &bINum
	}

	return fracQuoArray, err
}

// INumMgrModulo
//
// Performs a modulo operation on input parameters 'dividend' and
// 'divisor'. Both parameters must implement the INumMgr interface.
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one BigIntNum value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This informatin is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo':
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as
// the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very
// large precision values. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Dividend    mod by    Divisor    =    Modulo/Remainder
//
//
//	 12.555       /         2.5      =        0.055
//	 12.555       /         2        =        0.555
//	  2.5         /        12.555    =        2.5
//	-12.555       /         2.5      =       -0.055
//	-12.555       /         2        =       -0.555
//	 -2.5         /        12.555    =       -2.5
//	 12.555       /        -2.5      =        0.055
//	 12.555       /        -2        =        0.555
//	  2.5         /       -12.555    =        2.5
//	-12.555       /        -2.5      =       -0.055
//	-12.555       /        -2        =       -0.555
//	 -2.5         /       -12.555    =       -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('modulo') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) INumMgrModulo(
	dividend INumMgr,
	divisor INumMgr,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.INumMgrModulo",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewINumMgr(dividend, divisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bigpairBig2Iszero, err := bPair.Big2.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigpairBig2Iszero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigpairBig2Iszero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	modulo, err = new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// NumStrQuotientMod
//
// Performs a division operation on input parameters 'dividend'
// and 'divisor' which are comprised as number strings. Number
// strings are strings of numeric digits representing a numeric
// value. Number strings may include a leading minus sign (-), or
// surrounding prentheses '()', indicating a negative numeric
// value. Number strings may also include a decimal separator used
// to separate integer and fractional digits. The decimal
// separator character is specified by input parameter,
// 'numStrNumSepsDto'.
//
// There are two BigIntNum return values: 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on
// T-Division (Truncate Division). See "Division and Modulus for
// Computer Scientists", DAAN LEIJEN, University of Utrecht Dept.
// of Computer Science, PO.Box 80.089, 3508 TB Utrecht The
// Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This informatino is also available at:
//
//	mathops/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//		q = D div d = f(D/d)
//
//		r = D mod d = D − d ·q
//
//	 'quotient'  -  The integer result of dividing the 'dividend' by
//	                the 'divisor'.
//
//	 'modulo'    -  The modulo operation finds the remainder after
//	                division of one number by another.
//	                (r = D mod d = D − d ·q)
//
// Input parameter 'numStrNumSepsDto' is a type NumericSeparatorDto
// and is used to parse the dividend and divisor number strings.
// 'numStrNumSepsDto' specifies the applicable decimal separator,
// thousands separator and currency symbol.
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large precision
// values.  Therefore, the user is advised to set a relevant value
// for 'maxPrecision'.
//
//	Examples
//	=========
//
//	          divided                              Modulo
//	Dividend    by      Divisor   =   Quotient    Remainder
//
//	 12.555      /        2.5     =       5         0.055
//	 12.555      /        2       =       6         0.555
//	  2.5        /       12.555   =       0         2.5
//	-12.555      /        2.5     =      -5        -0.055
//	-12.555      /        2       =      -6        -0.555
//	 -2.5        /       12.555   =       0        -2.5
//	 12.555      /       -2.5     =      -5         0.055
//	 12.555      /       -2       =      -6         0.555
//	  2.5        /      -12.555   =       0         2.5
//	-12.555      /       -2.5     =       5        -0.055
//	-12.555      /       -2       =       6        -0.555
//	 -2.5        /      -12.555   =       0        -2.5
//
//	Numeric Separators
//	==================
//
//	This method recieves two input parameters of type
//	NumericSeparatorDto: 'numStrNumSepsDto' and 'outputNumSepsDto'.
//	A NumericSeparatorDto is a structre containing the numeric
//	separator symbols for Thousands separator, Decimal separator
//	and Currency Symbol. These NumericSeparatorDto components
//	are used to parse number strings and display numeric values
//	formatted as number strings.
//
//	Input parameter 'numStrNumSepsDto' contains the numeric
//	separators used to parse the number strings. Currently,
//	only the Decimal separtor is used to distinguish integer
//	and fractional parts of a number string.
//
//	Input parameter 'outputNumSepsDto' will be used to configure
//	the 'quotient' and 'modulo' BigIntNum values returned by this
//	method.
//
//	If any of the 'numStrNumSepsDto' or 'outputNumSepsDto' Numeric
//	Separator characters are found to be invalid, they will be
//	automatically reset to USA default values.
func (bIDivide *BigIntMathDivide) NumStrQuotientMod(
	dividend string,
	divisor string,
	numStrNumSepsDto NumericSeparatorDto,
	outputNumSepsDto NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	if len(dividend) == 0 {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividend) == 0",
				ErrMessage: "Error: Input parameter 'dividend' is an empty string!",
			}
	}

	if len(divisor) == 0 {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(divisor) == 0",
				ErrMessage: "Error: Input parameter 'divisor' is an empty string!",
			}
	}

	bigIDividend, err := new(BigIntNum).NewNumStrWithNumSeps(
		dividend, numStrNumSepsDto, outputNumSepsDto)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigIDividend, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
					"dividend, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("dividend= '%v'", dividend),
				ErrMessage: err.Error(),
			}
	}

	bigIDivisor, err := new(BigIntNum).NewNumStrWithNumSeps(
		divisor, numStrNumSepsDto, outputNumSepsDto)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigIDividend, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
					"divisor, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("divisor= '%v'", divisor),
				ErrMessage: err.Error(),
			}
	}

	bigIDivisorIszero, err := bigIDivisor.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigIDivisorIszero, err := bigIDivisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bigIDivisorIszero {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bigIDivisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bigIDivisor' has a ZERO value.",
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(bigIDividend, bigIDivisor)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(bigIDividend, bigIDivisor)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, err = new(bigIntMathDivideNanobot).
		pairQuotientMod(&bPair, outputNumSepsDto, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairQuotientMod(&bPair, outputNumSepsDto, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividend, divisor, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// NumStrFracQuotient
//
// Performs a division operation on input parameters 'dividend'
// and 'divisor' and returns the 'quotient' of this division
// operation.
//
// The input parameters, 'dividend' and 'divisor', are number
// strings consisting of strings of numeric digits representing
// a specific numeric value. Number strings may include a leading
// minus sign (-), or surronding parentheses '()', indicating a
// negative numeric value. In addition, number strings may
// include decimal separators used to separate integer and
// fractional digits. The decimal separator character
// is specified by the input parameter, 'numStrNumSepsDto'.
//
// The resulting quotient is returned as a BigIntNum type
// representing the result of the division operation expressed as
// integer and fractional digits.
//
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	          divided                 Quotient
//	Dividend    by    Divisor  =  BigIntNum Integer  Precision  Result
//
//	 10.5        /       2          525                  2       5.25
//	 10          /       2          5                    0       5
//	 11.5        /       2.5        46                   1       4.6
//	  2.5        /      12.555      199123855037834     15       0.199123855037834
//	-12.555      /       2.5       -5022                 3      -5.022
//	-12.555      /       2         -62775                4      -6.2775
//	 -2.5        /      12.555     -199123855037834     15      -0.199123855037834
//	 12.555      /      -2.5       -5022                 3      -5.022
//	 12.555      /      -2         -62775                4      -6.2775
//	  2.5        /     -12.555     -199123855037834     15      -0.199123855037834
//	-12.555      /      -2.5        5022                 3       5.022
//	-12.555      /      -2          62775                4       6.2775
//	 -2.5        /     -12.555      199123855037834     15       0.199123855037834
//	-10          /      -2          5                    0       5
//
//	Numeric Separators
//	==================
//
//	This method recieves two input parameters of type
//	NumericSeparatorDto, 'numStrNumSepsDto' and 'outputNumSepsDto'.
//	A NumericSeparatorDto is a structre containing the numeric
//	separator symbols for Thousands separator, Decimal separator
//	and Currency Symbol. These NumericSeparatorDto components
//	are used to parse number strings and display numeric values
//	formatted as number strings.
//
//	Input parameter 'numStrNumSepsDto' contains the numeric
//	separators used to parse the number strings. Currently,
//	only the Decimal separtor is used to distinguish integer
//	and fractional parts of a number string.
//
//	Input parameter 'outputNumSepsDto' will be used to configure
//	the 'fracQuotient' BigIntNum value returned by this	method.
//
//	If any of the 'numStrNumSepsDto' or 'outputNumSepsDto' Numeric
//	Separator characters are found to be invalid, they will be
//	automatically reset to USA default values.
func (bIDivide *BigIntMathDivide) NumStrFracQuotient(
	dividend string,
	divisor string,
	numStrNumSepsDto NumericSeparatorDto,
	outputNumSepsDto NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrFracQuotient",
		"")

	if err != nil {
		return fracQuotient, err
	}

	if len(divisor) == 0 {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(divisor) == 0",
				ErrMessage: "Error: Input Parameter 'divisor' is an EMPTY string!",
			}
	}

	if len(dividend) == 0 {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividend) == 0",
				ErrMessage: "Error: Input Parameter 'dividend' is an EMPTY string!",
			}
	}

	numStrNumSepsDto.SetDefaultsIfEmpty()

	outputNumSepsDto.SetDefaultsIfEmpty()

	bigIDividend, err := new(BigIntNum).
		NewNumStrWithNumSeps(
			dividend,
			numStrNumSepsDto,
			outputNumSepsDto)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigIDividend, err := new(BigIntNum).\n" +
					"    NewNumStrWithNumSeps(dividend, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("dividend= '%v'", dividend),
				ErrMessage: err.Error(),
			}
	}

	bigIDivisor, err := new(BigIntNum).
		NewNumStrWithNumSeps(divisor, numStrNumSepsDto, outputNumSepsDto)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigIDivisor, err := new(BigIntNum).\n" +
					"    NewNumStrWithNumSeps((divisor, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("divisor= '%v'", divisor),
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(bigIDividend, bigIDivisor)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair)." +
					"    NewBigIntNum(bigIDividend, bigIDivisor)",
				ErrContext: fmt.Sprintf("bigIDividend= '%v' bigIDivisor= '%v'",
					dividend, divisor),
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, err = new(bigIntMathDivideNanobot).
		pairFracQuotient(&bPair, outputNumSepsDto, ePrefix)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"    pairFracQuotient(&bPair, outputNumSepsDto, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividend, divisor, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, err
}

// NumStrFracQuotientArray
//
// Performs a division operation on input parameters 'dividends'
// and 'divisor' which are formatted as number strings.
//
// Number strings are strings of numeric digits representing a
// specific numeric value. Number strings may include a leading
// minus sign (-), or surrounding parentheses '()', indicating a
// negative numeric value. In addition, number strings may include
// a decimal separator used to separate integer and fractional
// digits. The decimal separator character is specified by the
// input parameter, 'numStrNumSepsDto'.
//
// 'dividends' is an array of number strings. The division
// operation is performed on each element of the 'dividends' array
// using a single 'divisor'.
//
// The resulting quotients are returned as an array of BigIntNum
// types. The values represent result of each division operation
// expressed as integer and fractional digits.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Precision is
// defined as the number of numeric digits to the right of the
// decimal point. Be advised that this method is capable of
// calculating quotients with very long strings of fractional
// digits. Therefore, the user is advised to set a relevant value
// for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	         divided                Returned
//	Dividend   by    Divisor  =      Array        =   Result
//
//	   10.5    /      2.5     =  fracQuoArray[0]  =      4.2
//	   10      /      2.5     =  fracQuoArray[1]  =      4
//	   11.5    /      2.5     =  fracQuoArray[2]  =      4.6
//	    2.5    /      2.5     =  fracQuoArray[3]  =      1
//	  -12.555  /      2.5     =  fracQuoArray[4]  =     -5.022
//	   -2.5    /      2.5     =  fracQuoArray[5]  =     -1
//	   12.555  /      2.5     =  fracQuoArray[6]  =      5.022
//	 -122.783  /      2.5     =  fracQuoArray[7]  =    -49.1132
//	-6847.231  /      2.5     =  fracQuoArray[8]  =  -2738.8924
//	   -2.5    /      2.5     =  fracQuoArray[9]  =     -1
//	   -10     /      2.5     =  fracQuoArray[10] =     -4
//	   -10.5   /      2.5     =  fracQuoArray[11] =     -4.2
//
//	Numeric Separators
//	==================
//
//	This method recieves two input parameters of type
//	NumericSeparatorDto, 'numStrNumSepsDto' and 'outputNumSepsDto'.
//	A NumericSeparatorDto is a structre containing the numeric
//	separator symbols for Thousands separator, Decimal separator
//	and Currency Symbol. These NumericSeparatorDto components
//	are used to parse number strings and display numeric values
//	formatted as number strings.
//
//	Input parameter 'numStrNumSepsDto' contains the numeric
//	separators used to parse the number strings. Currently,
//	only the Decimal separtor is used to distinguish integer
//	and fractional parts of a number string.
//
//	Input parameter 'outputNumSepsDto' will be used to configure
//	the 'fracQuoArray' BigIntNum value returned by this method.
//
//	If any of the 'numStrNumSepsDto' or 'outputNumSepsDto' Numeric
//	Separator characters are found to be invalid, they will be
//	automatically reset to USA default values.
func (bIDivide *BigIntMathDivide) NumStrFracQuotientArray(
	dividends []string,
	divisor string,
	numStrNumSepsDto NumericSeparatorDto,
	outputNumSepsDto NumericSeparatorDto,
	maxPrecision uint) (fracQuoArray []BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrFracQuotientArray",
		"")

	if err != nil {
		return fracQuoArray, err
	}

	numStrNumSepsDto.SetDefaultsIfEmpty()

	outputNumSepsDto.SetDefaultsIfEmpty()

	lenAry := len(dividends)

	if lenAry == 0 {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividends) == 0",
				ErrMessage: "Error: Input Parameter 'dividends' is an EMPTY Array!",
			}
	}

	if len(divisor) == 0 {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(divisor) == 0",
				ErrMessage: "Error: Input Parameter 'divisor' is an EMPTY string!",
			}
	}

	fracQuoArray = make([]BigIntNum, lenAry, lenAry+20)

	bigINumDivisor, err := new(BigIntNum).
		NewNumStrWithNumSeps(divisor, numStrNumSepsDto, outputNumSepsDto)

	if err != nil {

		return []BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigINumDivisor, err := new(BigIntNum)." +
					"    NewNumStrWithNumSeps(divisor, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("divisor= '%v'", divisor),
				ErrMessage: err.Error(),
			}
	}

	for i := 0; i < lenAry; i++ {

		if len(dividends[i]) == 0 {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: fmt.Sprintf("len(dividends[%d]) == 0", i),
					ErrMessage: fmt.Sprintf("Error: Array element 'dividends[%d]' is an EMPTY string!", i),
				}
		}

		bigINumDividend, err := new(BigIntNum).
			NewNumStrWithNumSeps(dividends[i], numStrNumSepsDto, outputNumSepsDto)

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bigINumDividend, err := new(BigIntNum).\n" +
						"    NewNumStrWithNumSeps(dividends[i], numStrNumSepsDto, outputNumSepsDto)",
					ErrContext: fmt.Sprintf("dividends[%d]= '%v'",
						i, dividends[i]),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(bigINumDividend, bigINumDivisor)

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(\n" +
						"    bigINumDividend, bigINumDivisor)",
					ErrContext: fmt.Sprintf("dividends[%d]= '%v'; divisor= '%v'",
						i, dividends[i], divisor),
					ErrMessage: err.Error(),
				}
		}

		bPair.MaxPrecision = maxPrecision

		fracQuoArray[i], err = new(bigIntMathDivideNanobot).
			pairFracQuotient(&bPair, outputNumSepsDto, ePrefix)

		if err != nil {

			return fracQuoArray,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("fracQuoArray[%d], err =\n"+
						"   new(bigIntMathDivideNanobot).pairFracQuotient(&bPair, outputNumSepsDto, ePrefix)", i),
					ErrContext: fmt.Sprintf("dividends[%d]= '%v'; divisor= '%v'; maxPrecision= '%v'",
						i, dividends[i], divisor, maxPrecision),
					ErrMessage: err.Error(),
				}
		}
	}

	return fracQuoArray, err
}

// NumStrModulo
//
// Performs a modulo operation on input parameters 'dividend' and
// 'divisor'. Both input parameters are formatted as number
// strings. Number strings are strings of numeric digits which
// represent a specific numeric value.
//
// Number strings may include a leading minus sign, or surronding
// parentheses '()', indicating a negative numeric value.
//
// In addition, number strings may include decimal separators used
// to separate integer and fractional digits. The decimal separator
// character is specified by the input parameter, 'numStrNumSepsDto'.
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one BigIntNum value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as
// the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very
// large precision values. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//	Examples
//	=========
//
//	           mod
//	Dividend   by     Divisor  =  Modulo/Remainder
//
//	 12.555    %        2.5    =       0.055
//	 12.555    %        2      =       0.555
//	  2.5      %       12.555  =       2.5
//	-12.555    %        2.5    =      -0.055
//	-12.555    %        2      =      -0.555
//	 -2.5      %       12.555  =      -2.5
//	 12.555    %       -2.5    =       0.055
//	 12.555    %       -2      =       0.555
//	  2.5      %      -12.555  =       2.5
//	-12.555    %       -2.5    =      -0.055
//	-12.555    %       -2      =      -0.555
//	 -2.5      %      -12.555  =      -2.5
//
//	Numeric Separators
//	==================
//
//	This method recieves two input parameters of type
//	NumericSeparatorDto, 'numStrNumSepsDto' and 'outputNumSepsDto'.
//	A NumericSeparatorDto is a structre containing the numeric
//	separator symbols for Thousands separator, Decimal separator
//	and Currency Symbol. These NumericSeparatorDto components
//	are used to parse number strings and display numeric values
//	formatted as number strings.
//
//	Input parameter 'numStrNumSepsDto' contains the numeric
//	separators used to parse the number strings. Currently,
//	only the Decimal separtor is used to distinguish integer
//	and fractional parts of a number string.
//
//	Input parameter 'outputNumSepsDto' will be used to configure
//	the 'modulo' BigIntNum value returned by this	method.
//
//	If any of the 'numStrNumSepsDto' or 'outputNumSepsDto' Numeric
//	Separator characters are found to be invalid, they will be
//	automatically reset to USA default values.
func (bIDivide *BigIntMathDivide) NumStrModulo(
	dividend string,
	divisor string,
	numStrNumSepsDto NumericSeparatorDto,
	outputNumSepsDto NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrModulo",
		"")

	if err != nil {
		return modulo, err
	}

	if len(divisor) == 0 {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(divisor) == 0",
				ErrMessage: "Error: Input Parameter 'divisor' is an EMPTY string!",
			}
	}

	if len(dividend) == 0 {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(dividend) == 0",
				ErrMessage: "Error: Input Parameter 'dividend' is an EMPTY string!",
			}
	}

	numStrNumSepsDto.SetDefaultsIfEmpty()

	outputNumSepsDto.SetDefaultsIfEmpty()

	bigIDividend, err := new(BigIntNum).
		NewNumStrWithNumSeps(dividend, numStrNumSepsDto, outputNumSepsDto)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigIDividend, err := new(BigIntNum).\n" +
					"    NewNumStrWithNumSeps(dividend, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("dividend= '%v'", dividend),
				ErrMessage: err.Error(),
			}
	}

	bigIDivisor, err := new(BigIntNum).
		NewNumStrWithNumSeps(divisor, numStrNumSepsDto, outputNumSepsDto)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigIDivisor, err := new(BigIntNum).\n" +
					"    NewNumStrWithNumSeps(divisor, numStrNumSepsDto, outputNumSepsDto)",
				ErrContext: fmt.Sprintf("divisor= '%v'", divisor),
				ErrMessage: err.Error(),
			}
	}

	bigIDivisorIsZero, err := bigIDivisor.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigIDivisorIsZero, err = bigIDivisor.IsZero()",
				ErrContext: fmt.Sprintf("divisor= '%v'", divisor),
				ErrMessage: err.Error(),
			}
	}

	if bigIDivisorIsZero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bigIDivisor == 0",
				ErrMessage: "Error: Attempted Divide By ZERO!\n" +
					"'bidIDivisor' has a ZERO value.",
			}
	}

	bPair, err := new(BigIntPair).
		NewBigIntNum(bigIDividend, bigIDivisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair)." +
					"    NewBigIntNum(bigIDividend, bigIDivisor)",
				ErrContext: fmt.Sprintf("dividend= '%v' divisor= '%v'",
					dividend, divisor),
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	modulo, err = new(bigIntMathDivideNanobot).
		pairMod(&bPair, outputNumSepsDto, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividend, divisor, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// NumStrDtoQuotientMod
//
// Performs a division operation on NumStrDto type input
// parameters, 'dividend' and 'divisor'.
//
// There are two BigIntNum Type return values: 'quotient' and
// 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on
// T-Division (Truncate Division). See "Division and Modulus for
// Computer Scientists", DAAN LEIJEN, University of Utrecht Dept.
// of Computer Science, PO.Box 80.089, 3508 TB Utrecht The
// Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo':
//
//	  q = D div d = f(D/d)
//
//	  r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend' by
//	               the 'divisor'.
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large
// precision values. Therefore, the user is advised to set a
// relevant value for 'maxPrecision'.
//
//	Examples
//	=========
//
//	Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//	 12.555        /         2.5   =      5        0.055
//	 12.555        /         2     =      6        0.555
//	  2.5          /        12.555 =      0        2.5
//	-12.555        /         2.5   =     -5       -0.055
//	-12.555        /         2     =     -6       -0.555
//	 -2.5          /        12.555 =      0       -2.5
//	 12.555        /       - 2.5   =     -5        0.055
//	 12.555        /       - 2     =     -6        0.555
//	  2.5          /       -12.555 =      0        2.5
//	-12.555        /       - 2.5   =      5       -0.055
//	-12.555        /       - 2     =      6       -0.555
//	 -2.5          /       -12.555 =      0       -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) NumStrDtoQuotientMod(
	dividend NumStrDto,
	divisor NumStrDto,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrDtoQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy("Testing dividend").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = dividend.IsValid(ePrefix.XCpy(\"Testing dividend\").String())",
				ErrContext: "Input parameter 'dividend' (NumStrDto) is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	err = divisor.IsValid(ePrefix.XCpy("Testing divisor").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = divisor.IsValid(ePrefix.XCpy(\"Testing divisor\").String())",
				ErrContext: "Input parameter 'divisor' (NumStrDto) is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// This method will test the validity of dividend and divisor
	bPair, err := new(BigIntPair).NewNumStrDto(dividend, divisor)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair)." +
					"    NewNumStrDto(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend= '%v' divisor= '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bPairBig2IsZero, err := bPair.Big2.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2IsZero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bPairBig2IsZero {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2 == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	quotient, modulo, err = new(bigIntMathDivideNanobot).
		pairQuotientMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairQuotientMod(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor= '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, nil
}

// NumStrDtoFracQuotient
//
// Performs a division operation on NumStrDto Type input parameters
// 'dividend' and 'divisor'.
//
// The resulting quotient is returned as a BigIntNum type
// representing the result of the division operation expressed as
// integer and fractional digits.
//
// Remember that the BigIntNum type specifies 'precision'.
// Precision is defined as the number of fractional digits to the
// right of the decimal place.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//
//	            divided                   Quotient
//	Dividend       by      Divisor  =  BigIntNum Integer    Precision  Result
//
//	 10.5         /          2      =   525                      2      5.25
//	 10           /          2      =   5                        0      5
//	 11.5         /          2.5    =  	46                       1      4.6
//	  2.5         /         12.555  =   199123855037834         15      0.199123855037834
//	-12.555       /          2.5    =  -5022                     3      5.022
//	-12.555       /          2      =  -62775                    4      6.2775
//	- 2.5         /         12.555  =  -199123855037834         15      0.199123855037834
//	 12.555       /        - 2.5    =  -5022                     3      5.022
//	 12.555       /        - 2      =  -62775                    4      6.2775
//	  2.5         /        -12.555  =  -199123855037834         15      0.199123855037834
//	-12.555       /        - 2.5    =   5022                     3      5.022
//	-12.555       /        - 2      =   62775                    4      6.2775
//	- 2.5         /        -12.555  =   199123855037834         15      0.199123855037834
//	-10           /        - 2      =   5                        5      5.00000
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuotient') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) NumStrDtoFracQuotient(
	dividend NumStrDto,
	divisor NumStrDto,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrFracQuotientArray",
		"")

	if err != nil {
		return fracQuotient, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIsZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a zero value.",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	bPair, err := new(BigIntPair).NewNumStrDto(dividend, divisor)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).NewNumStrDto(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'",
					dividendNumStr, divisorNumStr),
				ErrMessage: err.Error(),
			}
	}

	bPair.MaxPrecision = maxPrecision

	fracQuotient, err = new(bigIntMathDivideNanobot).
		pairFracQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"pairFracQuotient(&bPair, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = fracQuotient.TrimTrailingFracZeros()

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fracQuotient.TrimTrailingFracZeros()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// NumStrDtoFracQuotientArray
//
// Performs a division operation on NumStrDto input parameters
// 'dividends' and 'divisor'. 'dividends' is an array of NumStrDto
// Types. The division operation is performed on each element of
// the 'dividends' array using a single 'divisor'.
//
// The resulting quotients are returned as an array of NumStrDto
// Types. These values represent the result of each division
// operation expressed as integer and fractional digits.
//
// The input parameter 'maxPrecision' is used to control the
// precision of the resulting fractional quotient. Precision is
// defined as the number of numeric digits to the right of the
// decimal point. Be advised that this method is capable of
// calculating quotients with very long strings of fractional
// digits. Therefore, the user is advised to set a relevant
// value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Note: For all examples maximum precision is specified as '15'.
//
//	             divided                      Returned
//	Dividend       by      Divisor    =        Array            =     Result
//
//	    10.5       /         2.5      =      fracQuoArray[0]    =       4.2
//	    10         /         2.5      =      fracQuoArray[1]    =       4
//	    11.5       /         2.5      =      fracQuoArray[2]    =       4.6
//	    2.5        /         2.5      =      fracQuoArray[3]    =       1
//	  -12.555      /         2.5      =      fracQuoArray[4]    =      -5.022
//	   -2.5        /         2.5      =      fracQuoArray[5]    =      -1
//	   12.555      /         2.5      =      fracQuoArray[6]    =       5.022
//	 -122.783      /         2.5      =      fracQuoArray[7]    =     -49.1132
//	-6847.231      /         2.5      =      fracQuoArray[8]    =   -2738.8924
//	   -2.5        /         2.5      =      fracQuoArray[9]    =      -1
//	  -10          /         2.5      =      fracQuoArray[10]   =      -4
//	  -10.5        /         2.5      =      fracQuoArray[11]   =      -4.2
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('fracQuoArray') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) NumStrDtoFracQuotientArray(
	dividends []NumStrDto,
	divisor NumStrDto,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (fracQuoArray []NumStrDto, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrDtoFracQuotientArray",
		"")

	if err != nil {
		return fracQuoArray, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return fracQuoArray, err
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIsZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a zero value.",
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return fracQuoArray,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenAry := len(dividends)

	if lenAry == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'dividends' is an EMPTY Array!\n",
			ePrefix)

		return []NumStrDto{}, err
	}

	numSeps.SetDefaultsIfEmpty()

	var dividendNumStr string

	fracQuoArray = make([]NumStrDto, lenAry, lenAry+20)

	for i := 0; i < lenAry; i++ {

		err = dividends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Testing dividends[%v]", i)).String())

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = dividends[i].IsValid(ePrefix.XCpy(\n" +
						"    fmt.Sprintf(\"Testing dividends[%v]\", i)).String())",
					ErrContext: fmt.Sprintf("dividends[%v] is INVALID", i),
					ErrMessage: err.Error(),
				}
		}

		dividendNumStr, err = dividends[i].GetNumStr()

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, errx := new(BigIntPair).NewNumStrDto(dividends[i], divisor)

		if errx != nil {

			fracQuoArray = []NumStrDto{}

			err = fmt.Errorf("%v\n"+
				"Error returned by new(BigIntPair).NewNumStrDto(dividends[i], divisor).\n"+
				"dividends[%v]='%v'\ndivisor='%v'\nError= %v\n",
				ePrefix,
				i,
				dividendNumStr,
				divisorNumStr,
				errx.Error())

			return fracQuoArray, err
		}

		bPair.MaxPrecision = maxPrecision

		bINum, err := new(bigIntMathDivideNanobot).
			pairFracQuotient(&bPair, numSeps, ePrefix)

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bINum, err := new(bigIntMathDivideNanobot).\n" +
						"pairFracQuotient(&bPair, numSeps, ePrefix)",
					ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision= '%v'",
						dividendNumStr, divisorNumStr, maxPrecision),
					ErrMessage: err.Error(),
				}
		}

		fracQuoArray[i], err = bINum.GetNumStrDto()

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("fracQuoArray[%d], err = bINum.GetNumStrDto()\n", i) +
						"    pairFracQuotient(&bPair, numSeps, ePrefix)",
					ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision= '%v'",
						dividendNumStr, divisorNumStr, maxPrecision),
					ErrMessage: err.Error(),
				}
		}
	}

	return fracQuoArray, nil
}

// NumStrDtoModulo
//
// Performs a modulo operation on NumStrDto type input parameters,
// 'dividend' and 'divisor'. The result of this division operation
// is returned as a type BigIntNum. The returned instance of
// BigIntNum ('modulo') will be configured with the numeric
// seprators passed through input parameter, 'numSeps'
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one NumStrDto value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo':
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as
// the number of fractional digits to the right of the decimal
// point. Be advised that these calculations can support very
// large precision values.
//
//	Examples
//	========
//
//	Dividend    mod by    Divisor  =  Modulo/Remainder
//
//	 12.555       %         2.5    =      0.055
//	 12.555       %         2      =      0.555
//	  2.5         %        12.555  =      2.5
//	-12.555       %         2.5    =     -0.055
//	-12.555       %         2      =     -0.555
//	 -2.5         %        12.555  =     -2.5
//	 12.555       %        -2.5    =      0.055
//	 12.555       %        -2      =      0.555
//	  2.5         %       -12.555  =      2.5
//	-12.555       %        -2.5    =     -0.055
//	-12.555       %        -2      =     -0.555
//	 -2.5         %       -12.555  =     -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('modulo') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) NumStrDtoModulo(
	dividend NumStrDto,
	divisor NumStrDto,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrDtoModulo",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// This method will test the validity of dividend and divisor.
	bPair, err := new(BigIntPair).NewNumStrDto(dividend, divisor)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).\n" +
					"    NewNumStrDto(dividend, divisor)",
				ErrContext: fmt.Sprintf("dividend='%v'\ndivisor='%v'\nmaxPrecision='%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	bPairBig2IsZero, err := bPair.Big2.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2IsZero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bPairBig2IsZero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "bPair.Big2.bigInt == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a zero value.",
			}
	}

	bPair.MaxPrecision = maxPrecision

	// This will call SetDefaultsIfEmpty
	// on 'numSeps'
	modulo, err = new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
					"pairFracQuotient(&bPair, numSeps, ePrefix)",
				ErrContext: fmt.Sprintf("dividend = '%v'; divisor = '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// NumStrDtoModuloToNumStrDto
//
// Performs a modulo operation on NumStrDto type input parameters,
// 'dividend' and 'divisor'. The result of this division operation
// is returned as a type NumStrDto ('modulo'). This returned value
// will be configured with the numeric separators passed through
// input parameter, 'numSeps'.
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// This method returns one NumStrDto value: 'modulo'.
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo':
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large
// precision values.  Therefore, the user is advised to set a relevant
// value for 'maxPrecision'.
//
//	Examples
//	========
//
//	Dividend     mod by    Divisor    =    Modulo/Remainder
//
//	 12.555         %          2.5    =        0.055
//	 12.555         %          2      =        0.555
//	  2.5           %         12.555  =        2.5
//	-12.555         %          2.5    =       -0.055
//	-12.555         %          2      =       -0.555
//	 -2.5           %         12.555  =       -2.5
//	 12.555         %         -2.5    =        0.055
//	 12.555         %         -2      =        0.555
//	  2.5           %        -12.555  =        2.5
//	-12.555         %         -2.5    =       -0.055
//	-12.555         %         -2      =       -0.555
//	 -2.5           %        -12.555  =       -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value ('modulo') will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) NumStrDtoModuloToNumStrDto(
	dividend NumStrDto,
	divisor NumStrDto,
	numSeps NumericSeparatorDto,
	maxPrecision uint) (modulo NumStrDto, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.NumStrDtoModuloToNumStrDto",
		"")

	if err != nil {
		return modulo, err
	}

	err = dividend.IsValid(ePrefix.XCpy(
		"Testing dividend").String())

	if err != nil {
		return modulo, err
	}

	err = divisor.IsValid(ePrefix.XCpy(
		"Testing divisor").String())

	if err != nil {
		return modulo, err
	}

	dividendNumStr, err := dividend.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := dividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := divisor.GetNumStr()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := divisor.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorIsZero, err := divisor.IsZero()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorIsZero, err := divisor.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisorIsZero {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor == 0",
				ErrMessage: "Error: Attempted divide by ZERO!\n" +
					"Input parameter 'divisor' has a zero value.",
			}
	}

	bINumModulo, err := new(BigIntMathDivide).
		NumStrDtoModulo(dividend, divisor, numSeps, maxPrecision)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINumModulo, err :=\tnew(BigIntMathDivide).\n" +
					"NumStrDtoModulo(dividend, divisor, numSeps, maxPrecision)",
				ErrContext: fmt.Sprintf("dividend= '%v'; divisor = '%v'; maxPrecision= '%v'",
					dividendNumStr, divisorNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	modulo, err = bINumModulo.GetNumStrDto()

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "modulo, err = bINumModulo.GetNumStrDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// PairFracQuotient
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the divisor.
//
// 'BigIntPair.maxPrecision' is used to control the maximum
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//	  Big1          BigIntNum  // The Dividend
//	  Big2          BigIntNum  // The Divisor
//	  maxPrecision  uint       // Controls Precision
//	}
//
// This method performs a division operation on BigIntNum parameters
// 'dividend'(BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
//
//	Dividend (BigIntPair.Big1) divided Divisor (BigIntPair.Big2) = quotient
//
// The resulting quotient is returned as a BigIntNum type
// representing the result 0f the division operation expressed as
// integer and fractional digits. The maximum number of fractional
// digits output to the result is controlled by BigIntPair.maxPrecision.
// Remember that the BigIntNum type specifies 'precision'. Precision
// is defined as the number of fractional digits to the right of the
// decimal place.
//
//		Examples
//		=========
//
//		Note: For all examples BigIntPair.maxPrecision is specified as '15'.
//
//		                                        Quotient
//		Dividend  divided by  Divisor   =   BigIntNum Integer    Precision    Result
//
//		  10.5         /         2      =                 525         2      5.25
//		  10           /         2      =                   5         0      5
//		  11.5         /         2.5    =                  46         1      4.6
//		  2.5          /        12.555  =     199123855037834        15      0.199123855037834
//		-12.555        /         2.5    =               -5022         3     -5.022
//		-12.555        /         2      =              -62775         4     -6.2775
//		 -2.5          /        12.555  =    -199123855037834        15     -0.199123855037834
//		 12.555        /       - 2.5    =               -5022         3     -5.022
//		 12.555        /       - 2      =              -62775         4     -6.2775
//		  2.5          /       -12.555  =    -199123855037834        15     -0.199123855037834
//		-12.555        /       - 2.5    =                5022         3      5.022
//		-12.555        /       - 2      =               62775         4      6.2775
//		 -2.5          /       -12.555  =     199123855037834        15      0.199123855037834
//		-10            /       - 2      =                   5         5      5.00000
//
//		Numeric Separators
//		==================
//
//		Input parameter, 'numSeps' consits of a NumericSeparatorDto
//		instance. A NumericSeparatorDto contains symbols or characters
//		for the decimal separator, thousands separator and currency
//		symbol. These separators are used when presenting numeric
//		values in number strings.
//
//		If any of the 'numSeps' Numeric Separator Components are set
//		to zero, those components will be automatically reset to USA
//		default values.
//
//		The returned value ('fracQuotient') will be configured with
//	 'numSeps'.
func (bIDivide *BigIntMathDivide) PairFracQuotient(
	bPair BigIntPair,
	numSeps NumericSeparatorDto) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.PairFracQuotient",
		"")

	if err != nil {
		return fracQuotient, err
	}

	fracQuotient, err = new(bigIntMathDivideNanobot).
		pairFracQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"pairFracQuotientNoNumSeps(bPair, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// PairIntQuotient
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the Divisor. This method performs integer division
// on input parameters, 'Dividend' (BigIntPair.Big1) and 'Divisor'
// (BigIntPair.Big2).
//
// The result of this division operation returns an integer quotient
// of Dividend (BigIntPair.Big1) divided by Divisor
// (BigIntPair.Big2).
//
// The division operation performed by this method is T-Division or
// truncated division. See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathops/notes/divmodnote-letter.pdf.
//
// After completing the division operation, an integer quotient of
// type BigIntNum is returned.
//
//	Examples
//	=========
//
//	                                    Return Value
//	Dividend  divided by  Divisor  =  Integer Quotient
//
//	    5          /          2    =           2
//	    5.25       /          2    =           2
//	    2          /          4    =           0
//	  - 5          /          2    =          -2
//	  - 5.25       /          2    =          -2
//	  - 2          /          4    =           0
//	    5          /         -2    =          -2
//	    5.25       /         -2    =          -2
//	    2          /         -4    =           0
//	  - 5          /         -2    =           2
//	  - 5.25       /         -2    =           2
//	  - 2          /         -4    =           0
//	   12.555      /         -2.5  =          -5
//	  -12.555      /         -2.5  =           5
//	   12.555      /         -2    =          -6
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value 'intQuotient' will be configured with
//	'numSeps'.
func (bIDivide *BigIntMathDivide) PairIntQuotient(
	bPair BigIntPair,
	numSeps NumericSeparatorDto) (intQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.PairIntQuotient",
		"")

	if err != nil {
		return intQuotient, err
	}

	intQuotient, err = new(bigIntMathDivideNanobot).
		pairIntQuotient(&bPair, numSeps, ePrefix)

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "intQuotient, err = new(bigIntMathDivideNanobot).\n" +
					"pairIntQuotientNoNumSeps(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = intQuotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = intQuotient.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return intQuotient, nil
}

// PairMod
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the Divisor. The method proceeds to perform a
// modulo operation on input parameters 'dividend' (BigIntPair.Big1)
// and 'divisor' (BigIntPair.Big2). The modulo result is returned
// as a type BigIntNum.
//
// 'BigIntPair.maxPrecision' is used to control the precision of
// the resulting fractional modulo returned by this method. Be
// advised that this method is capable of calculating modulo values
// with very long strings of fractional digits. Therefore, the user
// is advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//	  Big1              BigIntNum  // The Dividend
//	  Big2              BigIntNum  // The Divisor
//	  maxPrecision      uint       // Controls Precision
//	}
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
//	Examples
//	========
//
//	Dividend    mod by    Divisor    =  Modulo/Remainder
//
//	  12.555      %           2.5    =        0.055
//	  12.555      %           2      =        0.555
//	   2.5        %          12.555  =        2.5
//	 -12.555      %           2.5    =       -0.055
//	 -12.555      %           2      =       -0.555
//	 - 2.5        %          12.555  =       -2.5
//	  12.555      %         - 2.5    =        0.055
//	  12.555      %         - 2      =        0.555
//	  2.5         %         -12.555  =        2.5
//	 -12.555      %         - 2.5    =       -0.055
//	 -12.555      %         - 2      =       -0.555
//	 - 2.5        %         -12.555  =       -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) PairMod(
	bPair BigIntPair,
	numSeps NumericSeparatorDto) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.PairMod",
		"")

	if err != nil {
		return modulo, err
	}

	modulo, err = new(bigIntMathDivideNanobot).
		pairMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairMod(bPair, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, nil
}

// PairQuotientMod
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the Divisor. 'BigIntPair.maxPrecision' is used to
// control the precision of the resulting fractional quotient. Be
// advised that this method is capable of calculating quotients with
// very long strings of fractional digits. Therefore, the user is
// advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//	  Big1              BigIntNum  // The Dividend
//	  Big2              BigIntNum  // The Divisor
//	  maxPrecision      uint       // Controls Precision
//	}
//
// The method performs a division operation on BigIntNum input
// parameters 'dividend' (BigIntPair.Big1) and 'divisor'
// (BigIntPair.Big2). The result is a quotient and modulo
// (remainder). These values are returned as two BigIntNum types,
// 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend'
//	               by the 'divisor'.
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large precision
// values.
//
//	Examples
//	=========
//
//	 Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//	  12.555        /        2.5    =      5           0.055
//	  12.555        /        2      =      6           0.555
//	   2.5          /       12.555  =      0           2.5
//	 -12.555        /        2.5    =     -5          -0.055
//	 -12.555        /        2      =     -6          -0.555
//	 - 2.5          /       12.555  =      0          -2.5
//	  12.555        /      - 2.5    =     -5           0.055
//	  12.555        /      - 2      =     -6           0.555
//	   2.5          /      -12.555  =      0           2.5
//	 -12.555        /      - 2.5    =      5          -0.055
//	 -12.555        /      - 2      =      6          -0.555
//	 - 2.5          /      -12.555  =      0          -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIDivide *BigIntMathDivide) PairQuotientMod(
	bPair BigIntPair,
	numSeps NumericSeparatorDto) (quotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathDivide.PairQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	quotient, modulo, err = new(bigIntMathDivideNanobot).
		pairQuotientMod(&bPair, numSeps, ePrefix)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, modulo, err = new(bigIntMathDivideNanobot).\n" +
					"    pairQuotientMod(&bPair, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	return quotient, modulo, nil
}
