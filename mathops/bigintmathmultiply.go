package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
)

// BigIntMathMultiply
//
// The methods associated with this structute are used to perform
// multiplication operations using the *big.Int Type.
//
// If you are unfamiliar with the *big.Int type, reference:
//
//	https://golang.org/pkg/math/big/
type BigIntMathMultiply struct {
	Input  BigIntPair
	Result BigIntNum
}

// BigIntMultiply
//
//	This method receives two *big.Int numbers and their associated
//	precision specifications. This method then proceeds to perform
//	a multiplication operation by multiplying the 'multiplier' times
//	the 'multiplicand' to generate the 'product'.
//
//	'multiplier', 'multiplicand' and 'product' are configured as
//	pairs of *big.Int integer numbers and precision specifications.
//
//	Taken together, an integer number and precision specification
//	are used to define a fixed length floating point number.
//
//	Multiplication Operation
//	========================
//
//	In the multiplication operation, the number to be multiplied is
//	called the "multiplicand", while the number of times the
//	multiplicand is to be multiplied comes from the "multiplier".
//	Usually, the multiplier is placed first and the multiplicand is
//	placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	or result.
//
//	  multiplier x multiplicand = product or result
//
//	This method performs the multiplication operation described
//	above and afterward returns the result or 'product' as a
//	BigIntNum type.
//
//	Examples
//	========
//
//	Consider the following multiplication example.
//
//	  752.314 x 21.67894 = product
//
//	'multiplier' and 'multiplicand' would be configured as follows:
//
//	    multiplier            = 752314
//	    multiplierPrecision   = 3
//	    multiplicand          = 2167894
//	    multiplicandPrecision = 5
//
//	The 'product' value (16309.37006716) of 'multiplier' and
//	'multiplicand' would be calculated and configured as
//	follows:
//
//	    product           = 1630937006716
//	    productPrecision  = 8
//	    product value     = 16309.37006716
//
//	Input Parameters
//	================
//
//	multiplier               *big.Int
//	  The number to be multiplied by 'multiplicand'
//
//	multiplierPrecision      *big.Int
//	  The 'multiplier' precision or numeric digits to
//	  the right of the decimal point.
//
//	multiplicand             *big.Int
//	  The number to be multiplied by the 'multiplier'.
//
//	multiplicandPrecision    *big.Int
//	  The 'multiplicand' precision or numeric digits to
//	  the right of the decimal point.
//
//	Return Values
//	=============
//
//	product                  *big.Int
//	  The product of the multiplier multiplied by the multiplicand.
//
//
//	productPrecision         *big.Int
//	  The precision specification for the returned product. Here,
//	  the term precision is defined as the number of fractional
//	  digits to the right of the decimal place in the returned
//	  'product'. 'productPrecision' is always equal to or greater
//	  than zero.
//
//
//	Note
//	====
//
//	This method removes trailing fractional zeros from the result.
//
//	  Example: 3.1200 is returned as 3.12
func (bMultiply *BigIntMathMultiply) BigIntMultiply(
	multiplier *big.Int,
	multiplierPrecision *big.Int,
	multiplicand *big.Int,
	multiplicandPrecision *big.Int) (product *big.Int, productPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.BigIntMultiply",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return new(bigIntMathMultiplyElectron).multiplyBigInt(
		multiplier, multiplierPrecision, multiplicand, multiplicandPrecision, ePrefix)
}

// BigIntMultiplyByTenToPower
//
//		 Overview
//		 ========
//
//		 This method Multiplies a *big.Int number by ten to the power
//		 of 'exponent'. 'exponent' is an input parameter of type
//		 *big.Int.
//
//		   product = multiplier x 10^exponent
//
//		 The result is returned as a *big.Int type ('product') with a
//		 precision specification ('productPrecision').
//
//		 Example
//		 =======
//
//		 In this example, input parameter 'multiplier' is set to
//		 a value of '53285' with a 'multiplierPrecision' of '2' .
//
//		 Input parameter exponent is set to a value of '4'.
//
//		 The calculation result or product is:
//
//		     product:            5328500
//		     product precision:        0
//		     numeric value:    5,328,500
//	      calculation:  532.85 x 10^4 = 5,328,500
//
//			Input Parameters
//			================
//
//			multiplier            *big.Int
//
//			'multiplier' will be multiplied by 10 to power of 'exponent'
//			to generate the result or 'product'.
//
//			multiplierPrecision   *big.Int
//
//			The precision specification for 'multiplier'. Precision
//			specifies the number of digits to the right of the decimal
//			place in 'multiplier'. This value must be greater than or
//			equal to zero.
//
//			exponent              *big.Int
//
//			Ten will be raised to the power of exponent and multiplied by
//			'multiplier' to generate the result or product. 'exponent'
//			can be a negative value.
//
//			Return Values
//			=============
//
//			product               *big.Int
//
//			The result generated by multiplying 'multiplier' by ten to the
//			power of 'exponent'.
//
//
//			productPrecision      *big.Int
//
//			The precision specification for 'product'. Precision specifies
//			the number of digits to the right of the decimal place in
//			'product'. This value will always be greater than or equal to
//			zero.
//
//
//			err                   error
//
//			If 'multiplierPrecision' or 'exponent' are less than zero, an
//			error will be returned. If no processing errors are encountered
//			during execution, value of 'nil' will be returned for 'err'.
//
//
//			Note
//			====
//
//			This method will remove trailing fractional zeros from the final
//			result (product).
func (bMultiply *BigIntMathMultiply) BigIntMultiplyByTenToPower(
	multiplier,
	multiplierPrecision,
	exponent *big.Int) (product *big.Int, productPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.BigIntMultiplyByTenToPower",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return new(bigIntMathMultiplyMechanics).multiplyByTenToPowerBigInt(
		multiplier, multiplierPrecision, exponent, ePrefix)
}

// BigIntMultiplyByTwoToPower
//
//	This method multiplies a *big.Int number by powers of two and
//	returns the result (product) as a *big.Int value with an
//	associated precision specification.
//
//		   product = multiplier X 2^exponent
//
//
//	Three input parameters are required: the multiplier, multiplier
//	precision and the exponent value.
//
//	Examples
//	========
//
//	------  multiplier ---------  --------- product ------------
//	integer  precision  exponent  integer  precision   value
//
//	12345        5         15     40452096      4      4045.2096
//
//	            (0.12345 x 2^15 = 4045.2096)
//
//	571          1          8     146176        1      14617.6
//
//	              (57.1 x 2^8 = 14617.6)
//
//
//	Input Parameters
//	================
//
//	multiplier               *big.Int
//	  'multiplier' will be multiplied by two to power of 'exponent'
//	  to generate the result or 'product'.
//
//
//	multiplierPrecision	     *big.Int
//	  The precision specification for 'multiplier'. Precision
//	  specifies the number of digits to the right of the decimal
//	  place in 'multiplier'. This value must be greater than or
//	  equal to zero.
//
//	exponent                 uint
//	  Two will be raised to the power of exponent and multiplied
//	  by 'multiplier' to generate the result or product. Since
//	  'exponent' is a type uint (unsigned integer), the exponent
//	  must be a positive number. Negative exponents are not
//	  supported in this calculation.
//
//
//	Return Values
//	=============
//
//	product                  *big.Int
//	  The multiplication result. product will be set equal to
//	  multiplier times two to the power of exponent.
//
//
//	productPrecision         *big.Int
//	  The precision specification for 'product'. Precision
//	  specifies the number of digits to the right of the
//	  decimal place in 'product'. This value will always be
//	  greater than or equal to zero.
//
//	err                      error
//	  If 'multiplierPrecision' is less than zero, an error will
//	  be returned. If no processing errors are encountered, this
//	  error return value will be set to 'nil'.
//
//
//	Note
//	====
//
//	This method will delete trailing fractional zeros from the
//	returned result (product).
func (bMultiply *BigIntMathMultiply) BigIntMultiplyByTwoToPower(
	multiplier,
	multiplierPrecision *big.Int,
	exponent uint) (product *big.Int, productPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.BigIntMultiplyByTwoToPower",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	product = big.NewInt(0)
	productPrecision = big.NewInt(0)
	err = nil

	product,
		productPrecision,
		err = new(bigIntMathMultiplyNanobot).
		multiplyByTwoToPowerBigInt(
			multiplier,
			multiplierPrecision,
			exponent,
			ePrefix)

	return product, productPrecision, err
}

// FixedDecimalMultiply
//
//		This method receives two BigIntFixedDecimal types and then
//		proceeds to perform a multiplication operation by multiplying
//		the 'multiplier' by the 'multiplicand' to generate the
//		'product'.
//
//		Multiplication Operation
//		========================
//
//		In the multiplication operation, the number to be multiplied is
//		called the "multiplicand", while the number of times the
//		multiplicand is to be multiplied comes from the "multiplier".
//		Usually the multiplier is placed first and the multiplicand is
//		placed second.
//
//		For example, in the problem 5 x 3 equals 15, the 5 is the
//		'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//		or result.
//
//		  multiplier x multiplicand = product or result
//
//		'multiplier', 'multiplicand' and 'product' are
//		BigIntFixedDecimal types which may be used to defined fixed
//		length floating point numbers.
//
//		BigIntFixedDecimal
//		==================
//
//		The BigIntFixedDecimal structure is defined as
//
//		type BigIntFixedDecimal struct {
//		  integerNum *big.Int  -  All the numeric digits, both integer and fractional,
//		                          necessary to define a fixed length floating point number.
//		                          The number of digits to the right of the decimal place
//		                          is specified by the data field,
//		                          BigIntFixedDecimal.precision.
//
//		  precision  uint      -  Specifies the number of digits to the right of the decimal
//		                          place in the series of numeric digits represented by data
//		                          field BigIntFixedDecimal.integerNum.
//
//		}
//
//		To represent the floating point number 52.459, a
//		BigIntFixedDecimal Structure would be configured as follows:
//
//		   BigIntFixedDecimal.integerNum = 52459
//		   BigIntFixedDecimal.precision  = 3
//
//		Consider the following multiplication example:
//		   product = 752.314 x 21.67894 = 16309.37006716
//
//		'multiplier' and 'multiplicand' would be configured as follows:
//
//		   multiplier.integerNum   = 752314
//		   multiplier.precision    = 3
//		   multiplicand.integerNum = 2167894
//		   multiplicand.precision  = 5
//
//		The 'product' would be calculated as follows:
//
//		   product.integerNum  = 1630937006716
//		   product.precision   = 8
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//		them into numeric values.
//
//		Numeric Separator characters are typically encapsulated in a
//		NumericSeparatorDto type.
//
//		The BigIntFixedDecimal multiplication 'Product' returned by
//		this method will contain numeric separators (Decimal Separator,
//		Thousands Separator and Currency Symbol) derived from one of
//		two possible sources.
//
//		Users have the option to supply an input parameter,
//		'outputNumSeps' of type NumericSeparatorDto. If this optional
//		parameter is provided, it will be used to configure the
//		BigIntFixedDecimal multiplication 'product' returned from this
//		method. Note that the first valid NumericSeparatorDto in the
//		'outputNumSeps' series will be selected and used. There is no
//		need to provide more than one valid NumericSeparatorDto object
//		for input parameter 'outputNumSeps'.
//
//		If the optional input parameter 'outputNumSeps' is NOT
//		provided, the returned BigIntFixedDecimal instance will be
//		configured using numeric separators copied from input
//		parameter, 'multiplier'.
//
//		Input Parameters
//		================
//
//		multiplier               BigIntFixedDecimal
//		  The number to be multiplied by 'multiplicand'
//
//		multiplicand             BigIntFixedDecimal
//		  The number to be multiplied by the 'multiplier'.
//
//		outputNumSeps            ... NumericSeparatorDto
//		  This method is defined as a variadic function in that
//		  'outputNumSeps' is configured as an optional input parameter
//		  meaning that it is NOT required. The user can choose to
//		  provide a value for 'outputNumSeps', or not.
//
//		  If the user chooses to provide a valid 'NumericSeparatorDto'
//		  object for this parameter, it will be used to configure the
//		  'BigIntNum' product value returned by this method.
//
//		  Note that only the first valid NumericSeparatorDto in the
//		  'outputNumSeps' series will be selected and used. There is no
//		  need to provide more than one valid NumericSeparatorDto object
//		  for parameter 'outputNumSeps'.
//
//		  Be advised that if the user chooses NOT to provide this
//		  optional parameter, the 'BigIntNum' value returned by this
//		  method will be configued using the 'NumericSeparatorDto'
//		  copied from the 'multiplier' input parameter.
//
//		Return Values
//		=============
//
//		product                  BigIntFixedDecimal
//	   The product of the 'multiplier' multiplied by	the
//	   'multiplicand'.
//
//		err                      error
//	   If no errors are encountered during execution, this returned
//	   error value will be set to 'nil'.
func (bMultiply *BigIntMathMultiply) FixedDecimalMultiply(
	multiplier BigIntFixedDecimal,
	multiplicand BigIntFixedDecimal,
	outputNumSeps ...NumericSeparatorDto) (product BigIntFixedDecimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.FixedDecimalMultiply",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntFixedDecimal{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntFixedDecimal{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return new(bigIntMathMultiplyMechanics).multiplyBigIntFixedDecimals(
		multiplier, multiplicand, finalOutputNumSeps, ePrefix)
}

// MultiplyBigInt2ToPowerBigIntNum
//
//	Multiplies a *big.Int number by powers of two and returns
//	the result as a BigIntNum type.
//
//	This method returns a BigIntNum type for the result.
//
//	Examples
//	========
//
//		product = multiplier X 2^exponent
//
//	multiplier   multiplierPrecision    exponent       product
//
//	12345                 5                15          4045.2096
//	           (0.12345 x 2^15 = 4045.2096)
//
//
//	571                   1                 8          14617.6
//	             (57.1 x 2^8 = 14617.6)
//
//	Trailing Fractional Zeros
//	=========================
//
//	This method will delete trailing fractional zeros from the
//	returned product.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	Numeric Separator characters are typically encapsulated in a
//	NumericSeparatorDto type.
//
//	For this method, users may provide an optional input parameter
//	labeled 'outputNumSeps' of type NumericSeparatorDto. If this
//	optional input parameter is provided, it will be used to format
//	the BigIntNum object returned by this method. Note that the
//	first valid NumericSeparatorDto in the 'outputNumSeps' series
//	will be selected and used. There is no need to provide more
//	than one valid NumericSeparatorDto object for parameter
//	'outputNumSeps'.
//
//	If optional input parameter 'outputNumSeps' is NOT provided,
//	the returned instance of BigIntNum will be configured with
//	default USA Numeric Separators.
//
//	Input Parameters
//	================
//
//	multiplier               *big.Int
//		The product value computed by this method will be calculated
//	  by multiplying 'multiplier' times 2 to power or 'exponent'.
//
//	      product = multiplier X 2^exponent
//
//	multiplierPrecision      uint
//	  Precision is defined as the number of fractional numeric
//	  digits to the right of the decimal place.
//	  'multiplierPrecision' therefore defines the number of digits
//	  in the 'multiplier' value which are fractional digits to the
//	  right of the decimal place. See the examples above.
//
//	exponent                 uint
//	  This value defines the exponent to used in multiplication
//	  operation:
//
//	      product = multiplier X 2^exponent
//
//	outputNumSeps            ... NumericSeparatorDto
//
//	  This method is defined as a variadic function in that
//	  'outputNumSeps' is configured as an optional input parameter
//	  meaning that it is NOT required. The user can choose to
//	  provide a value for 'outputNumSeps', or not.
//
//	  If the user chooses to provide a valid 'NumericSeparatorDto'
//	  object for this parameter, it will be used to configure the
//	  'BigIntNum' product value returned by this method.
//
//	  Note that the only first valid NumericSeparatorDto in the
//	  'outputNumSeps' series will be selected and used. There is no
//	  need to provide more than one valid NumericSeparatorDto object
//	  for parameter 'outputNumSeps'.
//
//	  Be advised that if the user chooses NOT to provide this
//	  optional parameter, the 'BigIntNum' value returned by this
//	  method will be configued using default USA Numeric Separators.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The product of the multiplication operation described above is
//	  returned as BigIntNum type.
//
//	error
//	  If no errors are encountered during method execution, this
//	  return parameter is set to 'nil'.
//
//	Be Advised
//	==========
//
//	This method differs BigIntMathMultiply.BigIntMultiplyByTwoToPower.
//	BigIntMathMultiply.BigIntMultiplyByTwoToPower returns *big.Int
//	types.
func (bMultiply *BigIntMathMultiply) MultiplyBigInt2ToPowerBigIntNum(
	multiplier *big.Int,
	multiplierPrecision uint,
	exponent uint,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {
	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigInt2ToPowerBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	var finalOutputNumSeps NumericSeparatorDto

	usaNumSeps := new(NumericSeparatorDto)

	usaNumSeps.SetUSADefaults()

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"output",
			"usa",
			usaNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"numStr\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of bPairBig1NumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			usaNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'usaNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &numStrNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'numStrNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return new(bigIntMathMultiplyMechanics).multiplyByTwoToPowerBigIntNum(
		multiplier, multiplierPrecision, exponent, finalOutputNumSeps, ePrefix)
}

// New
//
// Creates a BigIntMathMultiply instance with data variables
// initialized to zero.
func (bMultiply *BigIntMathMultiply) New() (BigIntMathMultiply, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.New",
		"")

	if err != nil {
		return BigIntMathMultiply{}, err
	}

	b2Math, err := new(bigIntMathMultiplyElectron).
		newBigIntMathMultiplyZero(ePrefix)

	return b2Math, err
}

// NewBigIntPairResult
//
// Creates a new BigIntMathMultiply based on input parameter type,
// 'BigIntPair'.
func (bMultiply *BigIntMathMultiply) NewBigIntPairResult(
	bPair BigIntPair) (BigIntMathMultiply, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.NewBigIntPairResult",
		"")

	if err != nil {
		return BigIntMathMultiply{}, err
	}

	b2Math := BigIntMathMultiply{}

	bPair2, err := bPair.CopyOut()

	if err != nil {

		return BigIntMathMultiply{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair2, err := bPair.CopyOut()",
				ErrMessage: err.Error(),
			}
	}

	b2Math.Input = bPair2

	return b2Math, nil
}

// MultiplyBigIntsBigIntNum
//
// Receives two *big.Int numbers and their associated precision
// specifications. This method then proceeds to perform a
// multiplication operation by multiplying the 'multiplier' times
// the 'multiplicand'.
//
// In the multiplication operation, the number to be multiplied
// is called the "multiplicand", while the number of times the
// multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is
// placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the
// 'multiplier' and 3 is the 'multiplicand', The result of this
// calculation, '15', is designated as the 'product'.
//
//					multiplier x multiplicand = product or result
//
//	 Input Parameters
//	 ================
//
//	 multiplier               *big.Int
//	   The number to be multiplied by 'multiplicand'
//
//	 multiplierPrecision      uint
//	   The 'multiplier' precision or numeric digits after the
//	   decimal point in 'multiplier'.
//
//	 multiplicand             *big.Int
//	   The number to be multiplied by the 'multiplier'.
//
//	 multiplicandPrecision    uint
//	   The 'multiplicand' precision or numeric digits after the
//	   decimal point.
//
//	 Return Values
//	 =============
//
//	 BigIntNum
//	   This method performs the multiplication operation and
//	   returns the result or 'product' as a BigIntNum type.
//
//	 Numeric Separators
//	 ==================
//
//	 Numeric sepatators include the Decimal separator, Thousands
//	 separator and Currency symbol characters. Numeric separators
//	 are used to parse number strings and display numeric values
//	 formatted as number strings.
//
//	 This method will configure the returned BigIntFixedDecimal
//	 'product' with Numeric Separators current configured in the
//	 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntsBigIntNum(
	multiplier *big.Int,
	multiplierPrecision uint,
	multiplicand *big.Int,
	multiplicandPrecision uint,
	numSepsDto NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigIntsBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntMathMultiplyMechanics).multiplyBigIntsBigIntNum(
		multiplier,
		multiplierPrecision,
		multiplicand,
		multiplicandPrecision,
		numSepsDto,
		ePrefix)
}

// MultiplyBigIntNums - Receives two BigIntNum types as input parameters and then
// performs a 'multiplication' operation on these BigIntNum's.
//
// In the multiplication operation, the number to be multiplied is called the
// "multiplicand", while the number of times the multiplicand is to be multiplied
// comes from the "multiplier". Usually the multiplier is placed first and the
// multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//	multiplier x multiplicand = product or result
//
// This method performs the multiplication operation and afterward returns the
// result or 'product' as a BigIntNum type.
//
// The returned BigIntNum multiplication 'result' will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from
// input parameter, 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNums(
	multiplier BigIntNum,
	multiplicand BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathMultiply.MultiplyBigIntNums()"

	bPair, err := new(BigIntPair).NewBigIntNum(multiplier, multiplicand)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(multiplier, multiplicand)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	finalResult, err := bMultiply.MultiplyPair(bPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "finalResult, err := bMultiply.MultiplyPair(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// MultiplyBigIntNumArray
//
// Receives one BigIntNum which is classified as the 'multiplier'.
// The second input parameter is an array of BigIntNum Types
// labeled, 'multiplicands'. The first element of the
// 'multiplicands' array is multiplied by the 'multiplier' to
// produce a 'product'. That 'product' replaces the 'multiplier'
// and is multiplied by the next element in the multiplicands
// array. This process is continued through the last element in
// the array and the combined, final 'product' is returned as a
// Type 'BigIntNum'.
//
// In the multiplication operation, the number to be multiplied
// is called the "multiplicand", while the number of times the
// multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand
// is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the
// 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
// or result.
//
//	multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described
// above and afterward returns the result or 'product' as a
// BigIntNum type.
//
// The returned BigIntNum multiplication 'result' will contain
// numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter,
// 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumArray(
	multiplier BigIntNum,
	multiplicands []BigIntNum) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigIntNumArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var finalResult BigIntNum

	finalResult, err = multiplier.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err := multiplier.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' array is empty!",
			}
	}

	var numSeps NumericSeparatorDto

	numSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var bPair BigIntPair

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.XCpy(
			fmt.Sprintf("Validating 'multiplicands[%d]'", i)).String())

		if err != nil {
			return BigIntNum{}, err
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicands[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("bPair, err = new(BigIntPair).NewBigIntNum(\n"+
						"finalResult, multiplicands[%d])", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, numSeps,
				ePrefix.XCpy("bPair & numSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"    multiplyPairNoNumSeps(bPair, numSeps,\n" +
						"    ePrefix.XCpy(\"bPair & numSeps\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return finalResult, err
}

// MultiplyBigIntNumOutputToArray - Receives one input parameter of Type BigIntNum which
// is classified as the 'multiplier'. The second input parameter is an array of BigIntNum
// Types labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
//
//	Multiplicands												Output
//
// Multiplier				    	Array														Array
//
//	3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//	3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//	3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//	3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//	3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//	3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
// This method performs the multiplication operation described above and afterward returns the
// result or 'product' in an Array of 'BigIntNums' ([] BigIntNums).
//
// Each element of the returned BigIntNum array 'result' will contain
// numeric separators (decimal separator, thousands separator and
// currency symbol) copied from input parameter, 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumOutputToArray(
	multiplier BigIntNum,
	multiplicands []BigIntNum) ([]BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigIntNumOutputToArray",
		"")

	if err != nil {
		return []BigIntNum{}, err
	}

	var bINumInterimResult BigIntNum

	bINumInterimResult, err = multiplier.CopyOut()

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return []BigIntNum{}, nil
	}

	var numSeps NumericSeparatorDto

	numSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return []BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultArray := make([]BigIntNum, lenMultiplicands)

	var bPair BigIntPair

	for i := 0; i < lenMultiplicands; i++ {

		bPair, err = new(BigIntPair).NewBigIntNum(bINumInterimResult, multiplicands[i])

		if err != nil {

			return []BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("bPair, err = new(BigIntPair).NewBigIntNum(\n"+
						"bINumInterimResult, multiplicands[%d]))", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		resultArray[i], err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, numSeps,
				ePrefix.XCpy("bPair & numSeps"))

		if err != nil {

			return []BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("resultArray[%d], err = new(bigIntMathMultiplyNanobot).\n"+
						"    multiplyPairWithNumSeps(bPair, numSeps,\n"+
						"    ePrefix.XCpy(\"bPair & numSeps\"))", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return resultArray, err
}

// MultiplyBigIntNumSeries - Receives one input parameter of Type BigIntNum which is classified
//
//	as the 'multiplier'. The second input parameter is a series of BigIntNum Types labeled,
//
// 'multiplicands'. The first element of the 'multiplicands' series is multiplied by the 'multiplier'
// to produce a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next
// element in the multiplicand series. This process is continued through the last element in the
// series. Afterward, the combined final 'product' is returned as a Type 'BigIntNum'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//	multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterward returns the
// result or 'product' as a BigIntNum type.
//
// The returned BigIntNum multiplication 'result' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input parameter,
// 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumSeries(
	multiplier BigIntNum,
	multiplicands ...BigIntNum) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.BigIntMultiply",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	var finalResult BigIntNum

	finalResult, err = multiplier.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err = multiplier.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' array is empty!",
			}
	}

	var numSeps NumericSeparatorDto

	numSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var bPair BigIntPair

	for idx, multiplicand := range multiplicands {

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicand)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"    finalResult, multiplicand)",
					ErrContext: fmt.Sprintf("multiplicand idx= %d", idx),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, numSeps,
				ePrefix.XCpy("bPair & numSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, numSeps,\n" +
						"ePrefix.XCpy(\"bPair & numSeps\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return finalResult, err
}

// MultiplyBigIntNumByTwo - Receives a BigIntNum input parameter 'base' and then
// proceeds to multiply this value times two (2).
//
//	product = base X 2
//
// The product of this multiplication operation is returned as a BigIntNum.
// This returned BigIntNum 'product' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input
// parameter,'base'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByTwo(base BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathMultiply.MultiplyBigIntNumByTwo()"

	result := big.NewInt(0).Lsh(base.bigInt, 1)

	basePrecision, err := base.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "basePrecision, err := base.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bINumResult, err := new(BigIntNum).NewBigInt(result, basePrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bINumResult, err := new(BigIntNum).NewBigInt(result, basePrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	baseNumSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "baseNumSeps, err := base.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = bINumResult.SetNumericSeparatorsDto(baseNumSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = bINumResult.SetNumericSeparatorsDto(baseNumSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bINumResult, nil
}

// MultiplyBigIntNumByTwoToPower - Receives a BigIntNum input parameter 'base' and then
// proceeds to multiply this value times two to the power of 'exponent'.
//
//	product = base X 2^exponent
//
// The product of this multiplication operation is returned as a BigIntNum.
// This returned BigIntNum 'product' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input
// parameter,'base'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByTwoToPower(
	base BigIntNum, exponent uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigIntNumByTwoToPower",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	result := big.NewInt(0).Lsh(base.bigInt, exponent)

	var uintPrecision uint

	uintPrecision, err = base.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "uintPrecision, err = base.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bINumResult, err := new(BigIntNum).NewBigInt(result, uintPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINumResult, err := new(BigIntNum).\n" +
					"NewBigInt(result, uintPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	baseNumSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseNumSeps, err := base.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = bINumResult.SetNumericSeparatorsDto(baseNumSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bINumResult.SetNumericSeparatorsDto(baseNumSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bINumResult, nil
}

// MultiplyBigIntNumByThree - Receives a BigIntNum input parameter 'base' and then
// proceeds to multiply this value times three (3).
//
//	product = base X 3
//
// The product of this multiplication operation is returned as a BigIntNum.
// This returned BigIntNum 'product' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input
// parameter,'base'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByThree(bigIntNum1 BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathMultiply.MultiplyBigIntNumByThree()"

	bigINumThree, err := new(BigIntNum).NewThree(0)

	if err != nil {

		return bigIntNum1,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bigINumThree, err :=\tnew(BigIntNum).NewThree(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(bigIntNum1, bigINumThree)

	if err != nil {

		return bigIntNum1,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(base, bigINumThree)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bMultiply2 := new(BigIntMathMultiply)

	result, err := bMultiply2.MultiplyPair(bPair)

	if err != nil {

		return bigIntNum1,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "result, err := bMultiply2.MultiplyPair(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}

// MultiplyBigIntNumByFive - Receives a BigIntNum input parameter 'base' and then
// proceeds to multiply this value times five (5).
//
//	product = base X 5
//
// The product of this multiplication operation is returned as a BigIntNum.
// This returned BigIntNum 'product' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input
// parameter,'base'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByFive(base BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathMultiply.MultiplyBigIntNumByFive()"

	newFiveBigInt, err := new(BigIntNum).NewFive(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "newFiveBigInt, err := new(BigIntNum).NewFive(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(base, newFiveBigInt)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(base, newFiveBigInt)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	mulResult, err := bMultiply.MultiplyPair(bPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "mulResult, err := bMultiply.MultiplyPair(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	return mulResult, nil
}

// MultiplyBigIntNumByTen - Receives a BigIntNum input parameter 'base' and then
// proceeds to multiply this value times ten (10).
//
//	product = base X 10
//
// The product of this multiplication operation is returned as a BigIntNum.
// This returned BigIntNum 'product' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input
// parameter,'base'.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByTen(base BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathMultiply.MultiplyBigIntNumByTen()"

	tenBigInt, err := new(BigIntNum).NewTen(0)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "tenBigInt, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(base, tenBigInt)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(\n" +
					"    base, tenBigInt)",
				ErrMessage: err.Error(),
			}
	}

	mulResult, err := bMultiply.MultiplyPair(bPair)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "mulResult, err := bMultiply.MultiplyPair(bPair)",
				ErrMessage: err.Error(),
			}
	}

	return mulResult, nil
}

// MultiplyBigIntNumByTenToPower - Receives two BigIntNum input parameters, 'base'
// and 'tenExponent'. This method then proceeds to multiply 'base' time 10 to the
// exponent, 'tenExponent'.
//
//	result = base X 10^tenExponent
//
// The exponent can be a negative value and/or a fractional value.
//
// Input parameter 'maxPrecision' is used to control the maximum precision for the
// result returned by this method. Precision is defined as the number of fractional
// digits to the right of the decimal place. Maximum precision therefore controls
// the maximum number of decimal digits to the right of the decimal place. If the
// returned result from this operation contains a number of fractional digits which
// is greater than 'maxPrecision' the result will be rounded to 'maxPrecision' decimal
// places. Be advised that these calculations can support very large precision values.
//
// Return Value
// ============
// The return value is of type BigIntNum and represents the result of the
// multiplication operation described above. This returned BigIntNum multiplication
// 'result' will contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter,'base'.
//
// If the precision of the return value precision exceeds input parameter 'maxPrecision',
// the return value will be rounded to 'maxPrecision' decimal places.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByTenToPower(
	base, tenExponent BigIntNum,
	maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigIntNumByTenToPower",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bigINumTen, err := new(BigIntNum).NewTen(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigINumTen, err := new(BigIntNum).NewTen(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	scale, err := new(BigIntMathPower).BigIntNumPwr(bigINumTen, tenExponent, maxPrecision+20)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "scale, err := new(BigIntMathPower).BigIntNumPwr(\n" +
					"    bigINumTen, tenExponent, maxPrecision+20)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// result numSeps are copied from 'base'
	result, err := bMultiply.MultiplyBigIntNums(base, scale)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "result, err := bMultiply.\n" +
					"    MultiplyBigIntNums(base, scale)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if result.precision > maxPrecision {

		err = result.RoundToDecPlace(maxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.RoundToDecPlace(maxPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return result, nil
}

// MultiplyBigIntNumByTenToIntPower - Receives a BigIntNum input parameter, 'base'
// and an uint64 input parameter, 'tenExponent'. This method then proceeds to multiply
// 'base' time 10 to the exponent, 'tenExponent'.
//
//	result = base X 10^tenExponent
//
// Input parameter 'maxPrecision' is used to control the maximum precision for the
// result returned by this method. Precision is defined as the number of fractional
// digits to the right of the decimal place. Maximum precision therefore controls
// the maximum number of decimal digits to the right of the decimal place. If the
// returned result from this operation contains a number of fractional digits which
// is greater than 'maxPrecision' the result will be rounded to 'maxPrecision' decimal
// places. Be advised that these calculations can support very large precision values.
//
// Return Value
// ============
// The return value is of type BigIntNum and represents the result of the
// multiplication operation described above. This returned BigIntNum multiplication
// 'result' will contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter,'base'.
//
// If the precision of the return value precision exceeds input parameter 'maxPrecision',
// the return value will be rounded to 'maxPrecision' decimal places.
func (bMultiply *BigIntMathMultiply) MultiplyBigIntNumByTenToIntPower(
	base BigIntNum,
	tenExponent uint64,
	maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyBigIntNumByTenToIntPower",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	scale :=
		big.NewInt(0).Exp(big.NewInt(10), big.NewInt(0).SetUint64(tenExponent), nil)

	bINumScale, err := new(BigIntNum).NewBigInt(scale, 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumScale, err := new(BigIntNum).NewBigInt(scale, 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// result numSeps are copied from 'base'
	result, err := bMultiply.MultiplyBigIntNums(base, bINumScale)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "result, err := bMultiply.MultiplyBigIntNums(base, bINumScale)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if result.precision > maxPrecision {

		err = result.RoundToDecPlace(maxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.RoundToDecPlace(maxPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return result, nil
}

// MultiplyDecimal - Receives two Decimal instances and multiplies their
// numeric values. The result or 'product' is returned as a 'BigIntNum'
// type.
//
// In the multiplication operation, the number to be multiplied is called the
// "multiplicand", while the number of times the multiplicand is to be multiplied
// comes from the "multiplier". Usually the multiplier is placed first and the
// multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//	multiplier x multiplicand = product or result
//
// This method performs the multiplication operation and afterward returns the
// result or 'product' as a BigIntNum type.
//
// The returned BigIntNum multiplication 'result' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input parameter,
// 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyDecimal(
	multiplier,
	multiplicand Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathMultiply.MultiplyDecimal() "

	// This method tests the validity of multiplier and multiplicand.
	bPair, err := new(BigIntPair).NewDecimal(multiplier, multiplicand)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix+
				"Error returned by new(BigIntPair).NewDecimal(multiplier, multiplicand). "+
				"Error='%v' ", err.Error())
	}

	finalResult, err := bMultiply.MultiplyPair(bPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "finalResult, err := bMultiply.MultiplyPair(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// MultiplyDecimalArray - Receives one Decimal which is classified as the 'multiplier'.
// The second input parameter is an array of Decimal Types labeled, 'multiplicands'. The
// first element of the 'multiplicands' array is multiplied by the 'multiplier' to produce
// a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next element
// in the multiplicands array. This process is continued through the last element in the array
// when the combined, final 'product' is returned as a Type 'BigIntNum'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//	multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterward returns the
// result or 'product' as a BigIntNum type.
//
// The returned BigIntNum multiplication 'result' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input parameter,
// 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyDecimalArray(
	multiplier Decimal,
	multiplicands []Decimal) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyDecimalArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(ePrefix.\n" +
					"XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Error: Input parameter 'multiplier' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' array is empty!",
			}
	}

	var finalResult BigIntNum

	// This method tests the validity of 'multiplier'.
	finalResult, err = new(BigIntNum).NewDecimal(multiplier)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err = new(BigIntNum).NewDecimal(multiplier)",
				ErrContext: fmt.Sprintf("multiplier='%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var numSeps NumericSeparatorDto

	numSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: fmt.Sprintf("multiplier='%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var multiplicandBINum BigIntNum
	var multiplicandBINumStr string
	var bPair BigIntPair

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.
			XCpy(fmt.Sprintf("Validating multiplicands[%d]", i)).String())

		if err != nil {
			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = multiplicands[i].IsValid(ePrefix.\n" +
						"  XCpy(fmt.Sprintf(\"Validating multiplicands[%d]\", i)).String())",
					ErrContext: fmt.Sprintf("Error: multiplicands[%d] is INVALID!", i),
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: fmt.Sprintf("Initial Result: finalResult = multiplier='%v'", multiplierNumStr),
					ErrMessage: err.Error(),
				}
		}

		// This method tests the validity of multiplicands[i]
		multiplicandBINum, err = new(BigIntNum).NewDecimal(multiplicands[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("multiplicandBINum.NewDecimal(multiplicands[%d])",
						i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplicandBINumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandBINumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicands index=[%d])", i),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult= '%v'; multiplicandBINum= '%v'; multiplicand index= '%d'",
						multiplicandBINumStr, multiplicandBINumStr, i),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, numSeps,
				ePrefix.XCpy("bPair & numSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, numSeps,\n" +
						"ePrefix.XCpy(\"bPair & numSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v'; bPair.Big2= '%v'\n"+
						"multiplicand index= '%d'",
						multiplierNumStr, multiplicandBINumStr, i),
					ErrMessage: err.Error(),
				}
		}
	}

	return finalResult, err
}

// MultiplyDecimalOutputToArray - Receives one input parameter of Type Decimal which
// is classified as the 'multiplier'. The second input parameter is an array of Decimal
// Types labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
//
//	Multiplicands												Output
//
// Multiplier				    	Array														Array
//
//	3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//	3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//	3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//	3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//	3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//	3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
// This method performs the multiplication operation described above and afterward returns the
// result or 'product' in an Array of 'Decimals' ([] Decimals).
//
// The returned Decimal Array ([]Decimal) will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from input parameter, 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyDecimalOutputToArray(
	multiplier Decimal,
	multiplicands []Decimal) ([]Decimal, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyDecimalOutputToArray",
		"")

	if err != nil {
		return []Decimal{}, err
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return []Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "lenMultiplicands == 0",
				ErrMessage: "Error: multiplicands array is Empty!",
			}
	}

	var multiplicandNumStr, multiplierBiNumStr string

	// This method tests the validity of multiplier
	multiplierBINum, err := multiplier.GetBigIntNum()

	if err != nil {
		return []Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierBINum, err := multiplier.GetBigIntNum()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplierBiNumStr, err = multiplierBINum.GetNumStr()

	if err != nil {

		return []Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierBiNumStr, err = multiplierBINum.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultArray := make([]Decimal, lenMultiplicands)

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating multiplicands[%d]", i)).String())

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = multiplicands[%d].IsValid(ePrefix.XCpy(fmt.Sprintf(\"Validating multiplicands[%d]\", i)).String())", i, i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicands[i].GetNumStr()

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("multiplicandNumStr, err = multiplicands[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// This method tests the validity of multiplicands[i]
		multiplicandBINum, err := new(BigIntNum).NewDecimal(multiplicands[i])

		if err != nil {
			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicands[%v]='%v'",
						i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(\n" +
						"    multiplierBINum, multiplicandBINum)",
					ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicands[%v]='%v'",
						multiplierBiNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		result, err := bMultiply.MultiplyPair(bPair)

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "result, err := bMultiply.MultiplyPair(bPair)",
					ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicands[%v]='%v'",
						multiplierBiNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		resultArray[i], err = result.GetDecimal()

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("resultArray[%d], err = result.GetDecimal()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return resultArray, nil
}

// MultiplyDecimalSeries
//
//	Overview
//	========
//
//	This method receives one input parameter of Type Decimal which
//	is classified as the 'multiplier'. The second input parameter
//	is a comma-delimited series of Decimal Types labeled,
//	'multiplicands'. The first element of the 'multiplicands'
//	series is multiplied by the 'multiplier' to produce a
//	'product'. That 'product' replaces the 'multiplier' and is
//	multiplied by the next element in the 'multiplicands' series.
//	This process is repeated through the last element in the
//	series. Afterward, the combined final 'product' is returned as
//	a Type 'BigIntNum'.
//
//	Multiplication Operation
//	========================
//
//	In the multiplication operation, the number to be multiplied is
//	called the "multiplicand", while the number of times the
//	multiplicand is to be multiplied comes from the "multiplier".
//	Usually, the multiplier is placed first and the multiplicand is
//	placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	or result.
//
//	    multiplier x multiplicand = product or result
//
//	This method performs the multiplication operation described above
//	and afterward returns the result or 'product' as a BigIntNum type.
//
// The returned BigIntNum multiplication 'result' will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from input parameter,
// 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyDecimalSeries(
	multiplier Decimal,
	multiplicands ...Decimal) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyDecimalSeries",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: " err = multiplier.IsValid(ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	if len(multiplicands) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: Input parameter 'multiplicands' array is empty!",
			}
	}

	var numSeps NumericSeparatorDto

	numSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// This method will test the validity of 'multiplier'
	finalResult, err := new(BigIntNum).NewDecimal(multiplier)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err := new(BigIntNum).NewDecimal(multiplier)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var multiplicandBINum BigIntNum
	var multiplicandNumStr, multiplierNumStr string
	var bPair BigIntPair

	for idx, multiplicand := range multiplicands {

		err = multiplicand.IsValid(ePrefix.XCpy(
			fmt.Sprintf("Validating 'multiplicand' idx=%d", idx)).String())

		if err != nil {
			return BigIntNum{}, err
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicand index= '%v'", idx),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicand.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicand.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicand index= '%v'", idx),
					ErrMessage: err.Error(),
				}
		}

		// This method will test the validity of 'multiplicand'
		multiplicandBINum, err = new(BigIntNum).NewDecimal(multiplicand)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err = new(BigIntNum).NewDecimal(multiplicand)",
					ErrContext: fmt.Sprintf("multiplicand='%v'; multiplicand index= '%d'",
						multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: " multiplicandNumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicand range index=[%d])", idx),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult= '%v'; multiplicandBINum= '%v'; Index= '%d'",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, numSeps,
				ePrefix.XCpy("bPair & numSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, numSeps,\n" +
						"ePrefix.XCpy(\"bPair & numSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v'; bPair.Big2= '%v'; Index= '%v'",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}
	}

	return finalResult, err
}

// MultiplyIntAry
//
//	 Overview
//	 ========
//
//	 This method receives two IntAry instances and multiplies their
//	 numeric values. The result or 'product' is returned as a
//	 'BigIntNum' type.
//
//	   multiplier x multiplicand = product or result
//
//	 Be Advised
//	 ==========
//
//	 IntAry's can accommodate very, very large numbers.
//
//	 Multiplication Operation
//	 ========================
//
//	 In the multiplication operation, the number to be multiplied is
//	 called the "multiplicand", while the number of times the
//	 multiplicand is to be multiplied comes from the "multiplier".
//	 Usually, the multiplier is placed first and the multiplicand is
//	 placed second.
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the
//	 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	 or result.
//
//	   multiplier x multiplicand = product or result
//
//	 This method performs the multiplication operation described
//	 above and afterward returns the result or 'product' as a
//	 BigIntNum type.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 The BigIntNum multiplication 'product' returned by this method
//	 will contain numeric separators (Decimal Separator, Thousands
//	 Separator and Currency Symbol) derived from one of two possible
//	 sources.
//
//		Users have the option to supply an input parameter,
//		'outputNumSeps' of type NumericSeparatorDto. If this optional
//	 parameter is provided, it will be used to configure the
//	 returned BigIntNum multiplication 'product' from this method.
//	 Note that the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for input parameter 'outputNumSeps'.
//
//		If the optional input parameter 'outputNumSeps' is NOT
//		provided, the returned BigIntNum instance will be
//		configured using numeric separators copied from input
//		parameter, 'multiplier'.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               IntAry
//
//		This IntAry instance serves as the initial 'multiplier' in the
//	 multiplication operation defined above.
//
//	 multiplicand             IntAry
//
//	 This instace of IntAry serves as the 'multiplicand' for the
//	 multiplication operation defined above.
//
//	 outputNumSeps            ... NumericSeparatorDto
//
//	 This method is defined as a variadic function in that
//	 'outputNumSeps' is configured as an optional input parameter
//	 meaning that it is NOT required. The user can choose to
//	 provide a value for 'outputNumSeps', or not.
//
//	 If the user chooses to provide a valid 'NumericSeparatorDto'
//	 object for this parameter, it will be used to configure the
//	 'BigIntNum' product value returned by this method with the
//	 Numeric Separators copied from 'outputNumSeps'.
//
//	 Note that only the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//	 If the user chooses NOT to provide this optional parameter, the
//	 'BigIntNum' value returned by this  method will be configued
//	 using the 'NumericSeparatorDto' copied from the 'multiplier'
//	 input parameter.
//
//	 Return Values
//	 =============
//
//	 BigIntNum
//
//	 The product of the multiplication operation described above is
//	 returned as BigIntNum type.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyIntAry(
	multiplier IntAry,
	multiplicand IntAry,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyIntAry",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(ePrefix.\n" +
					"XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Error: Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	err = multiplicand.IsValid(ePrefix.XCpy(
		"Validating 'multiplicand'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplicand.IsValid(ePrefix.\n" +
					"XCpy(\"Validating 'multiplicand'\").String())",
				ErrContext: "Error: Input parameter 'multiplicand' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr, multiplicandNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplicandNumStr, err = multiplicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplicandNumStr, err = multiplicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	// This method will test the validity of 'multiplier' and
	// 'multiplicand'.
	bPair, err := new(BigIntPair).NewIntAry(multiplier, multiplicand)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).\n" +
					"  NewIntAry(multiplier, multiplicand)",
				ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicand= '%v'",
					multiplierNumStr, multiplicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	var finalResult BigIntNum

	finalResult, err = new(bigIntMathMultiplyNanobot).
		multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
			ePrefix.XCpy("bPair & finalOutputNumSeps"))

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
					"  multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n" +
					"  ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
				ErrContext: fmt.Sprintf("bPair.Big1= '%v'; bPair.Big2= '%v'",
					multiplierNumStr, multiplicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// MultiplyIntAryArray
//
//	This method receives one IntAry instance which is classified as
//	the 'multiplier'. The second input parameter is an array of
//	IntAry Types labeled, 'multiplicands'. The first element of the
//	'multiplicands' array is multiplied by the 'multiplier' to
//	produce a 'product'. That 'product' then replaces the
//	'multiplier' and is multiplied by the next element in the
//	'multiplicands' array. This process is continued through the
//	last element in the array when the combined, final 'product' is
//	returned as a Type 'BigIntNum'.
//
//	  Example:
//	    multiplier = 3
//	    multiplicands = [3]IntAry{2,3,4}
//
//	    (1) 3 x 2 = 6
//	    (2) 6 x 3 = 18
//	    (3) 18 x 4 = 72
//	    The returned product is 72
//
//	Multiplication Operation
//	========================
//
//	In the multiplication operation, the number to be multiplied is
//	called the "multiplicand", while the number of times the
//	multiplicand is to be multiplied comes from the "multiplier".
//	Usually, the multiplier is placed first and the multiplicand is
//	placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	or result.
//
//	  multiplier x multiplicand = product or result
//
//	This method performs the multiplication operation described
//	above and afterward returns the result or 'product' as a
//	BigIntNum type.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	Numeric Separator characters are typically encapsulated in a
//	NumericSeparatorDto type.
//
//	The BigIntNum multiplication 'Product' returned by this method
//	will contain numeric separators (Decimal Separator, Thousands
//	Separator and Currency Symbol) derived from one of two possible
//	sources.
//
//	Users have the option to supply an input parameter,
//	'outputNumSeps' of type NumericSeparatorDto. If this optional
//	parameter is provided, it will be used to configure the
//	BigIntNum multiplication 'product' returned from this method.
//	Note that the first valid NumericSeparatorDto in the
//	'outputNumSeps' series will be selected and used. There is no
//	need to provide more than one valid NumericSeparatorDto object
//	for input parameter 'outputNumSeps'.
//
//	If the optional input parameter 'outputNumSeps' is NOT
//	provided, the returned BigIntNum instance will be
//	configured using numeric separators copied from input
//	parameter, 'multiplier'.
//
//	Input Parameters
//	================
//
//	multiplier               IntAry
//	  This IntAry instance serves as the initial 'multiplier' in
//	  the multiplication operation defined above.
//
//	multiplicands           []IntAry
//	  This array of IntAry objects serves as 'multiplicands' for
//	  the multiplication operation defined above.
//
//	outputNumSeps           ... NumericSeparatorDto
//	  This method is defined as a variadic function in that
//	  'outputNumSeps' is configured as an optional input parameter
//	  meaning that it is NOT required. The user can choose to
//	  provide a value for 'outputNumSeps', or not.
//
//	  If the user chooses to provide a valid 'NumericSeparatorDto'
//	  object for this parameter, it will be used to configure the
//	  'BigIntNum' product value returned by this method.
//
//	  Note that only the first valid NumericSeparatorDto in the
//	  'outputNumSeps' series will be selected and used. There is no
//	  need to provide more than one valid NumericSeparatorDto
//	  object for parameter 'outputNumSeps'.
//
//	  Be advised that if the user chooses NOT to provide this
//	  optional parameter, the 'BigIntNum' value returned by this
//	  method will be configued using the 'NumericSeparatorDto'
//	  copied from the 'multiplier' input parameter.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The product of the multiplication operation described above
//	  is returned as BigIntNum type.
//
//	error
//	  If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyIntAryArray(
	multiplier IntAry,
	multiplicands []IntAry,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyIntAryArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' ([]IntAry) array is empty!",
			}
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// This method will test the validity of 'multiplier'
	var finalResult BigIntNum

	finalResult, err = new(BigIntNum).NewIntAry(multiplier)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "finalResult, err := new(BigIntNum).\n" +
					"    NewIntAry(multiplier)",
				ErrContext: fmt.Sprintf("Multiplier = finalResult = '%s'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	var multiplicandBINum BigIntNum
	var multiplicandNumStr string
	var bPair BigIntPair

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.
			XCpy(fmt.Sprintf("Validating multiplicands[%d]", i)).String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = multiplicands[%d].IsValid(ePrefix.\n", i) +
						"XCpy(\"Validating 'multiplicands[i]'\").String())",
					ErrContext: fmt.Sprintf("'multiplicands[%d]' is invalid!", i),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicands[i].GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = multiplicands[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// This method will test the validity of multiplicands[i]
		multiplicandBINum, err = new(BigIntNum).NewIntAry(multiplicands[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err = new(BigIntNum).\n" +
						fmt.Sprintf("    NewIntAry(multiplicands[%d])", i),
					ErrContext: fmt.Sprintf("Multiplier '%s'; multiplicands[%d]= '%v' ",
						multiplierNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicands index = '%d'", i),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult '%s'; multiplicands[%d]= '%v' ",
						multiplierNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n " +
						"ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\nmultiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, i),
					ErrMessage: err.Error(),
				}
		}

	}

	return finalResult, err
}

// MultiplyIntAryOutputToArray
//
//	 Overview
//	 ========
//
//	 This method receives one input parameter of Type IntAry which
//	 is classified as the 'multiplier'. The second input parameter
//	 is an array of IntAry Types labeled, 'multiplicands'.
//
//	 Each element of the 'multiplicands' array is multiplied by the
//	 'multiplier'. The result or 'product' is then stored in a
//	 results array which is returned to the calling function.
//
//	 Example:
//
//	                Multiplicands                 Output
//	 Multiplier         Array                      Array
//
//	     3       x   multiplicands[0] = 2   =   outputarray[0] =  6
//	     3       x   multiplicands[1] = 3   =   outputarray[1] =  9
//	     3       x   multiplicands[2] = 4   =   outputarray[2] = 12
//	     3       x   multiplicands[3] = 5   =   outputarray[3] = 15
//	     3       x   multiplicands[4] = 6   =   outputarray[4] = 18
//	     3       x   multiplicands[5] = 7   =   outputarray[5] = 21
//
//
//	 Multiplication Operation
//	 ========================
//
//	 In the multiplication operation, the number to be multiplied is
//	 called the "multiplicand", while the number of times the
//	 multiplicand is to be multiplied comes from the "multiplier".
//	 Usually, the multiplier is placed first and the multiplicand is
//	 placed second.
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the
//	 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	 or result.
//
//	   multiplier x multiplicand = product or result
//
//	 This method performs the multiplication operation described
//	 above and afterward returns the result or 'product' as a
//	 BigIntNum type.
//
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 Each element in the array of IntAry objects returned by this
//	 method will contain Numeric Separators (Decimal Separator,
//	 Thousands Separator and Currency Symbol) copid input parameter
//	 'multiplier'.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               IntAry
//
//		This IntAry instance serves as the initial 'multiplier' in the
//	 multiplication operation defined above.
//
//	 multiplicands           []IntAry
//
//	 This array of IntAry objects serves as 'multiplicands' for the
//	 multiplication operation defined above.
//
//	 Return Values
//	 =============
//
//	 []IntAry
//
//	 An array of IntAry instances containing the product valus from
//	 the multiplication operation described above. Each element in
//	 this returned array will contain Numeric Seprators copied from
//	 input parameter 'multiplier'.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyIntAryOutputToArray(
	multiplier IntAry,
	multiplicands []IntAry,
	outputNumSeps ...NumericSeparatorDto) ([]IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyIntAryOutputToArray",
		"")

	if err != nil {
		return []IntAry{}, err
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' ([]IntAry) array is empty!",
			}
	}

	// This method will test the validity of 'multiplier'
	multiplierBINum, err := multiplier.GetBigIntNum()

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: " multiplierBINum, err := multiplier.GetBigIntNum()",
				ErrContext: fmt.Sprintf("multiplier= '%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	resultArray := make([]IntAry, lenMultiplicands)

	var multiplicandBINum, finalResult BigIntNum

	var multiplicandsNumStr string

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating multiplicands[%d]", i)).String())

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = multiplicands[%d].IsValid(\n"+
						"ePrefix.XCpy(\"Validating multiplicands[i]\").String())", i),
					ErrContext: fmt.Sprintf("Error: multiplicands[%d] is Invalid!", i),
					ErrMessage: err.Error(),
				}
		}

		multiplicandsNumStr, err = multiplicands[i].GetNumStr()

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("multiplicandsNumStr, err = multiplicands[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// This method will test the validity of multiplicands[i]
		multiplicandBINum, err = new(BigIntNum).NewIntAry(multiplicands[i])

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("multiplicandBINum, err = new(BigIntNum).\n"+
						"  NewIntAry(multiplicands[%d])", i),
					ErrContext: fmt.Sprintf("multiplicands[%d]= '%v'",
						i, multiplicandsNumStr),
					ErrMessage: err.Error(),
				}
		}

		multiplicandsNumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandsNumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)",
					ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicand= '%v'",
						multiplierNumStr, multiplicandsNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"  multiplyPairWithNumSeps( bPair, finalOutputNumSeps,\n" +
						"  ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v'; bPari.Big2= '%v'",
						multiplierNumStr, multiplicandsNumStr),
					ErrMessage: err.Error(),
				}
		}

		resultArray[i], err = finalResult.GetIntAry()

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "resultArray[i], err = finalResult.GetIntAry()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return resultArray, err
}

// MultiplyIntArySeries
//
//	 Overview
//	 ========
//
//	 This method receives one input parameter of Type IntAry which
//	 is classified as the 'multiplier'. The second input parameter
//	 is a comma-delimited series of IntAry Types passed through the
//	 variadic function parameter labeled, 'multiplicands'. The first
//	 element of the 'multiplicands' series is multiplied by the
//	 'multiplier' to produce a 'product'. That 'product' then
//	 replaces the 'multiplier' and is multiplied by the next element
//	 in the 'multiplicands' series. This process is continued through
//	 the last element in the series. Afterward, the combined final
//	 'product' is returned as a Type 'BigIntNum'.
//
//	   Example:
//	     multiplier = 3
//	     multiplicands = 2,3,4
//
//	     (1) 3 x 2 = 6
//	     (2) 6 x 3 = 18
//	     (3) 18 x 4 = 72
//	      The returned product is 72
//
//	 Multiplication Operation
//	 ========================
//
//	 In the multiplication operation, the number to be multiplied is
//	 called the "multiplicand", while the number of times the
//	 multiplicand is to be multiplied comes from the "multiplier".
//	 Usually, the multiplier is placed first and the multiplicand is
//	 placed second.
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the
//	 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	 or result.
//
//	     multiplier x multiplicand = product or result
//
//	 This method performs the multiplication operation described
//	 above and afterward returns the result or 'product' as a
//	 BigIntNum type.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 The BigIntNum multiplication 'Product' returned by this method
//	 will contain numeric separators (Decimal Separator, Thousands
//	 Separator and Currency Symbol) copied from input parameter,
//	 'multiplier'.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               IntAry
//
//	 This IntAry numeric value serves as the 'multiplier' in the
//	 multiplication operation described above.
//
//	 multiplicands            ...IntAry
//
//	 A series of IntAry objects will serve as the multiplicands in
//	 the multiplication operation described above.
//
//	 This method is defined as a variadic function meaning that this
//	 'multiplicands' may consist of any number of IntAry objects.
//
//	 Be advised that if this parameter is empty and no IntAry objects
//	 are provided, an error will be returned.
//
//	 Return Values
//	 =============
//
//	 BigIntNum
//
//	 The product of the multiplication operation described above is
//	 returned as BigIntNum type.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyIntArySeries(
	multiplier IntAry,
	multiplicands ...IntAry) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyIntArySeries",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier' (IntAry)").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier' (IntAry)\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var finalResult BigIntNum
	// This method will test the validity of 'multiplier'.
	finalResult, err = new(BigIntNum).NewIntAry(multiplier)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err = new(BigIntNum).NewIntAry(multiplier)",
				ErrContext: fmt.Sprintf("multiplier='%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps NumericSeparatorDto

	finalOutputNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalOutputNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(multiplicands) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' is empty!",
			}
	}

	var multiplicandNumStr string
	var bPair BigIntPair

	for idx, multiplicand := range multiplicands {

		err = multiplicand.IsValid(ePrefix.XCpy(
			fmt.Sprintf("Validating 'multiplicands' Index= %d", idx)).String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = multiplicand.IsValid(ePrefix.XCpy(Validating 'multiplicands').String())",
					ErrContext: fmt.Sprintf("multiplicands Index= '%d' is Invalid!", idx),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicand.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicand.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicand index= '%d'", idx),
					ErrMessage: err.Error(),
				}
		}

		// This method will test the validity of 'multiplicand'
		multiplicandBINum, err := new(BigIntNum).NewIntAry(multiplicand)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err := new(BigIntNum).NewIntAry(multiplicand)",
					ErrContext: fmt.Sprintf("multiplicand='%v'", multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"    finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult= '%v'; multiplicand= '%v'\nmultiplicand index= '%d'",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, finalOutputNumSeps, ePrefix)",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\nmultiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

	}

	return finalResult, err
}

// MultiplyNumStr
//
//	This method receives two number strings and multiplies their
//	numeric values to produce a 'product' which is returned as
//	type BigIntNum.
//
//	    multiplier x multiplicand = product
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. These digits
//	must be formatted in a way that facilitates conversion to a
//	corresponding numeric value. This method receives two number
//	string input parameters labeled, 'multiplier' and
//	'multiplicand' both of which are elements in the
//	multiplication operation performed by this method.
//
//	Number String Negative Values
//	=============================
//
//	The 'multiplier' and 'multiplicand' string parameters passed
//	to this method are number strings which consist of strings of
//	numeric digits representing a numeric value. A leading minus
//	sign (-), or surrounding parentheses '()', may be included in
//	these number strings to indicate a negative numeric value.
//
//	Fractional Digits in Number Strings
//	===================================
//
//	The 'multiplier' and 'multiplicand' strings of numeric
//	digits may also include a delimiting decimal separator to
//	identify fractional digits to the right of the decimal
//	separator. In the USA, the default decimal separator is the
//	period character ('.').
//
//	These number strings are parsed based on the Numeric
//	Separators, including the Decimal Separator character, which
//	are specified by input parameter 'numStrNumSeps'.
//
//	Multiplication Operation
//	========================
//
//	In the multiplication operation, the number to be multiplied
//	is called the "multiplicand", while the number of times the
//	multiplicand is to be multiplied comes from the "multiplier".
//	Usually. the multiplier is placed first and the multiplicand
//	is placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	or result.
//
//	    multiplier x multiplicand = product or result
//
//	This method performs the multiplication operation described
//	above and afterward returns the cumulative result or 'product'
//	as a BigIntNum type.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	Numeric Separator characters are typically encapsulated in a
//	NumericSeparatorDto type.
//
//	Input parameter 'inputNumSeps', of type NumericSeparatorDto,
//	contains numeric separators specifying the Decimal Separator,
//	Thousands Separator, and Currency Symbol. These numeric
//	separators are required to parse the 'multiplier' and
//	'multiplicand' number strings. This parsing operation is in
//	turn used to convert those strings into numeric values for
//	calculation purposes.
//
//	In addition, users may provide another optional input
//	parameter labeled 'outputNumSeps' of type NumericSeparatorDto.
//	If this optional parameter is provided, it will be used to
//	format the BigIntNum object returned by this method. Note that
//	the first valid NumericSeparatorDto in the 'outputNumSeps'
//	series will be selected and used. There is no need to provide
//	more than one valid NumericSeparatorDto object for parameter
//	'outputNumSeps'.
//
//	If optional input parameter 'outputNumSeps' is NOT provided,
//	or if the provided 'outputNumSeps' is invalid, the returned
//	instance of BigIntNum will be configured with Numeric
//	Separators supplied by mandatory input parameter,
//	'inputNumSeps'.
//
//	Input Parameters
//	================
//
//	multiplier               string
//		This number string contains the numeric value which serves as
//		the initial multiplier in the multiplication operation defined
//		above.
//
//	multiplicand             string
//	  This string contains the multiplicand numeric value used in
//	  the multiplication operation defined above.
//
//	inputNumSeps           NumericSeparatorDto
//	  This instance of NumericSeparatorDto contains the Decimal
//	  Separator character, Thousands Separator character and
//	  Currency Symbol character used to parse the number strings
//	  passed through input paramters 'multiplier' and
//	  'multiplicand'.
//
//	outputNumSeps            ... NumericSeparatorDto
//
//	  This method is defined as a variadic function in that
//	  'outputNumSeps' is configured as an optional input parameter
//	  meaning that it is NOT required. The user can choose to
//	  provide a value for 'outputNumSeps', or not.
//
//	  If the user chooses to provide a valid 'NumericSeparatorDto'
//	  object for this parameter, it will be used to configure the
//	  'BigIntNum' product value returned by this method.
//
//	  Note that the only first valid NumericSeparatorDto in the
//	  'outputNumSeps' series will be selected and used. There is no
//	  need to provide more than one valid NumericSeparatorDto object
//	  for parameter 'outputNumSeps'.
//
//	  Be advised that if the user chooses NOT to provide this
//	  optional parameter, the 'BigIntNum' value returned by this
//	  method will be configued using the 'NumericSeparatorDto'
//	  copied from 'inputNumSeps'
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The product of the multiplication operation described above is
//	  returned as BigIntNum type.
//
//	error
//	  If no errors are encountered during method execution, this
//	  return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStr(
	multiplier string,
	multiplicand string,
	inputNumSeps NumericSeparatorDto,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if len(multiplier) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "'multiplier' is an empty string.\n" +
					"    len(multiplier) == 0",
				ErrMessage: "Error: Input parameter 'multiplier' is INVALID!",
			}

	}

	if len(multiplicand) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "'multiplicand' is an empty string.\n" +
					"    len(multiplicand) == 0",
				ErrMessage: "Error: Input parameter 'multiplicand' is INVALID!",
			}

	}

	err = inputNumSeps.IsValid(ePrefix.XCpy("Vaidating 'inputNumSeps'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = inputNumSeps.IsValid(ePrefix.XCpy(\"Vaidating 'inputNumSeps'\").String())",
				ErrContext: "Error: Input Parameter 'inputNumSeps' is Invalid!",
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps NumericSeparatorDto

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"output",
			"input",
			&inputNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"numStr\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of bPairBig1NumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&inputNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'inputNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &numStrNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'numStrNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	numSepsPair := NumericSeparatorPairDto{}

	err = numSepsPair.InputSeparators.CopyIn(&inputNumSeps, true)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.InputSeparators.CopyIn(&inputNumSeps, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = numSepsPair.OutputSeparators.CopyIn(&finalOutputNumSeps, true)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.OutputSeparators.CopyIn(&finalOutputNumSeps, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var multiplierBINum, multiplicandBINum, finalResult BigIntNum

	multiplierBINum, err = new(BigIntNum).NewNumStrWithNumSeps(multiplier, &numSepsPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "multiplierBINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
					"    multiplier, numStrNumSeps, finalOutputNumSeps)",
				ErrContext: fmt.Sprintf("multiplier = '%v'\nnumStrNumSeps= '%v'\nfinalOutputNumSeps=  '%v'",
					multiplier, inputNumSeps.String(), finalOutputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	multiplicandBINum, err = new(BigIntNum).NewNumStrWithNumSeps(multiplicand, &numSepsPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "multiplicandBINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
					"    multiplicand, numStrNumSeps, finalOutputNumSeps)",
				ErrContext: fmt.Sprintf("multiplicand = '%v'\nnumStrNumSeps= '%v'\nfinalOutputNumSeps= '%v'",
					multiplicand, inputNumSeps.String(), finalOutputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	bPair, err := new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)",
				ErrContext: fmt.Sprintf("multiplier= '%s'; multiplicand= '%s' ",
					multiplier, multiplicand),
				ErrMessage: err.Error(),
			}
	}

	finalResult, err = new(bigIntMathMultiplyNanobot).
		multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
			ePrefix.XCpy("bPair & finalOutputNumSeps"))

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
					"  multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n" +
					"  ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
				ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'",
					multiplier, multiplicand),
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// MultiplyNumStrArray
//
//	 Overview
//	 ========
//
//	 This method receives one number string which is classified as
//	 the 'multiplier'. The second input parameter is an array of
//	 number strings labeled, 'multiplicands'. The first element of
//	 the 'multiplicands' array is multiplied by the 'multiplier' to
//	 produce a 'product'. That 'product' then replaces the
//	 'multiplier' and is multiplied by the next element in the
//	 'multiplicands' array. This process is repeated through the
//	 last element in the array when the combined, final 'product'
//	 is returned as a Type 'BigIntNum'.
//
//	   Example:
//	     multiplier = 3
//	     multiplicands = [3]string{2,3,4}
//
//	     (1) 3 x 2 = 6
//	     (2) 6 x 3 = 18
//	     (3) 18 x 4 = 72
//	      The returned product is 72
//
//	 Number String Negative Values
//	 =============================
//
//	 The 'multiplier' string and 'multiplicands' string array
//	 parameters passed to this method are number strings which
//	 consist of a string of numeric digits representing a numeric
//	 value. A leading minus sign (-), or surrounding parentheses
//	 '()', may be included in these number strings to indicate a
//	 negative numeric value.
//
//	 Fractional Digits in Number Strings
//	 ===================================
//
//	 'multiplier' string and 'multiplicands' strings of numeric
//	 digits may also include a delimiting decimal separator to
//	 identify fractional digits to the right of the decimal
//	 separator. These number strings are parsed based on the Numeric
//	 Separators, including the decimal separator character, which
//	 are specified by input parameter 'numStrNumSeps'.
//
//	 Multiplication Operation
//	 ========================
//
//	 In the multiplication operation, the number to be multiplied
//	 is called the "multiplicand", while the number of times the
//	 multiplicand is to be multiplied comes from the "multiplier".
//	 Usually. the multiplier is placed first and the multiplicand
//	 is placed second.
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the
//	 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	 or result.
//
//	     multiplier x multiplicand = product or result
//
//	 This method performs the multiplication operation described
//	 above and afterward returns the cumulative result or 'product'
//	 as a BigIntNum type.
//
//	 Numeric Separators
//	 ==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 Input parameter 'numStrNumSeps', of type NumericSeparatorDto,
//	 contains numeric separators specifying the Decimal Separator,
//	 Thousands Separator, and Currency Symbol. These numeric
//	 separators are required to parse the 'multiplier' and
//	 'multiplicands' number strings. This parsing operation is in
//	 turn used to convert those strings into numeric values for
//	 calculation purposes.
//
//	 In addition, users may provide another optional input
//	 parameter labeled 'outputNumSeps' of type NumericSeparatorDto.
//	 If this optional input parameter is provided, it will be used
//	 to format the BigIntNum object returned by this method. Note
//	 that only the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//	 If optional input parameter 'outputNumSeps' is NOT provided,
//	 or if the provided 'outputNumSeps' is invalid, the returned
//	 instance of BigIntNum will be configured with Numeric
//	 Separators supplied by mandatory input parameter,
//	 'numStrNumSeps'.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               string
//
//		This number string contains the numeric value which serves as
//		the initial multiplier in the multiplication operation defined
//		above.
//
//	 multiplicands            []string
//
//	 This array of strings contains the multiplicand numeric values
//	 for the multiplication operation defined above.
//
//	 numStrNumSeps            NumericSeparatorDto
//
//	 This instance of NumericSeparatorDto contains the Decimal
//	 Separator character, Thousands Separator character and the
//	 Currency Symbol character used to parse the number strings
//	 passed through input paramters 'multiplier' and
//	 'multiplicands'.
//
//	 outputNumSeps            ... NumericSeparatorDto
//
//	 This method is defined as a variadic function in that
//	 'outputNumSeps' is configured as an optional input parameter
//	 meaning that it is NOT required. The user can choose to
//	 provide a value for 'outputNumSeps', or not.
//
//	 If the user chooses to provide a valid 'NumericSeparatorDto'
//	 object for this parameter, it will be used to configure the
//	 'BigIntNum' product value returned by this method.
//
//	 Note that the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//	 Be advised that if the user chooses NOT to provide this
//	 optional parameter, the 'BigIntNum' value returned by this
//	 method will be configued using the 'NumericSeparatorDto'
//	 copied from 'numStrNumSeps'
//
//	 Return Values
//	 =============
//
//	 BigIntNum
//
//	 The product of the multiplication operation described above is
//	 returned as BigIntNum type.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStrArray(
	multiplier string,
	multiplicands []string,
	numStrNumSeps NumericSeparatorDto,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if len(multiplier) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "'multiplier' is an empty string.",
				ErrMessage: "Error: Input parameter 'multiplier' is INVALID!",
			}

	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' array is empty!",
			}
	}

	var finalOutputNumSeps NumericSeparatorDto

	finalOutputNumSeps.SetDefaultsIfEmpty()

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"numStr",
			&numStrNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"numStr\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}

	} else {

		err = finalOutputNumSeps.CopyIn(&numStrNumSeps, false)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = finalOutputNumSeps.CopyIn(&numStrNumSeps, true)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	numSepsPair := NumericSeparatorPairDto{}

	err = numSepsPair.InputSeparators.CopyIn(&numStrNumSeps, true)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.InputSeparators.CopyIn(&numStrNumSepsDto, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = numSepsPair.OutputSeparators.CopyIn(&finalOutputNumSeps, true)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.OutputSeparators.CopyIn(&numStrNumSepsDto, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	finalResult, err := new(BigIntNum).NewNumStrWithNumSeps(multiplier, &numSepsPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "finalResult, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
					"    multiplier, numStrNumSeps, finalOutputNumSeps)",
				ErrContext: fmt.Sprintf("multiplier = '%v'\nnumStrNumSeps= '%v'\nfinalOutputNumSeps=  '%v'",
					multiplier, numStrNumSeps.String(), finalOutputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	var multiplicandBINum BigIntNum
	var multiplierNumStr, multiplicandNumStr string
	var bPair BigIntPair

	for i := 0; i < lenMultiplicands; i++ {

		if len(multiplicands[i]) == 0 {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: fmt.Sprintf("len(multiplicands[%d]) == 0", i),
					ErrMessage: fmt.Sprintf("Error: 'multiplicands[%d]' is an empty string!", i),
				}
		}

		multiplicandBINum, err = new(BigIntNum).
			NewNumStrWithNumSeps(multiplicands[i], &numSepsPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("multiplicandBINum, err = new(BigIntNum).NewNumStrWithNumSeps(\n"+
						"    multiplicands[%d], numStrNumSeps, finalOutputNumSeps)", i),
					ErrContext: fmt.Sprintf("multiplicands[%v]= '%v'\nnumStrNumSeps= '%v'\nfinalOutputNumSeps= '%v'",
						i, multiplicands[i], numStrNumSeps.String(), finalOutputNumSeps.String()),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult '%s'; multiplicands[%d]= '%v' ",
						multiplierNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"  multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n" +
						"  ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\n"+
						"multiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, i),
					ErrMessage: err.Error(),
				}
		}
	}

	return finalResult, err
}

// MultiplyNumStrOutputToArray
//
//	 Overview
//	 ========
//
//	 This method receives one input parameter of Type string which
//	 is classified as the 'multiplier'. The second input parameter
//	 is an array of strings labeled, 'multiplicands'.
//
//	 Each element of the 'multiplicands' array is multiplied by the
//	 'multiplier'. The result or 'product' is then stored as a
//	 number string in a results string array which is returned
//	 to the calling function.
//
//	 Example:
//
//	                Multiplicands                 Output
//	 Multiplier         Array                      Array
//
//	     3       x   multiplicands[0] = 2   =   outputarray[0] =  6
//	     3       x   multiplicands[1] = 3   =   outputarray[1] =  9
//	     3       x   multiplicands[2] = 4   =   outputarray[2] = 12
//	     3       x   multiplicands[3] = 5   =   outputarray[3] = 15
//	     3       x   multiplicands[4] = 6   =   outputarray[4] = 18
//	     3       x   multiplicands[5] = 7   =   outputarray[5] = 21
//
//	 Number Strings
//	 ==============
//
//	 Number strings are strings of numeric digits. These digits
//	 must be formatted in way that allows conversion to a
//	 corresponding numeric value. This method receives two
//	 number string input parameters labeled, 'multiplier' and
//	 'multiplicand'
//
//	 Number String Negative Values
//	 =============================
//
//	 The 'multiplier' and 'multiplicand' string parameters passed
//	 to this method are number strings which consist of strings of
//	 numeric digits representing a numeric value. A leading minus
//	 sign (-), or surrounding parentheses '()', may be included in
//	 these number strings to indicate a negative numeric value.
//
//	 Fractional Digits in Number Strings
//	 ===================================
//
//	 The 'multiplier' and 'multiplicand' strings of numeric
//	 digits may also include a delimiting decimal separator to
//	 identify fractional digits to the right of the decimal
//	 separator. In the USA, the default decimal separator is the
//	 period character ('.').
//
//	 These number strings are parsed based on the Numeric
//	 Separators, including the Decimal Separator character, which
//	 are specified by input parameter 'numStrNumSeps'.
//
//	 Multiplication Operation
//	 ========================
//
//	 In the multiplication operation, the number to be multiplied
//	 is called the "multiplicand", while the number of times the
//	 multiplicand is to be multiplied comes from the "multiplier".
//	 Usually. the multiplier is placed first and the multiplicand
//	 is placed second.
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the
//	 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	 or result.
//
//	     multiplier x multiplicand = product or result
//
//	 This method performs the multiplication operation described
//	 above and afterward returns the cumulative result or 'product'
//	 as a BigIntNum type.
//
//	 Numeric Separators
//	 ==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 Input parameter 'numStrNumSeps', of type NumericSeparatorDto,
//	 contains numeric separators specifying the Decimal Separator,
//	 Thousands Separator, and Currency Symbol. These numeric
//	 separators are required to parse the 'multiplier' and
//	 'multiplicand' number strings. This parsing operation is in
//	 turn used to convert those strings into numeric values for
//	 calculation purposes.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               string
//
//		This number string contains the numeric value which serves as
//		the initial multiplier in the multiplication operation defined
//		above.
//
//	 multiplicands            []string
//
//	 This array of strings contains the multiplicand numeric values
//	 for the multiplication operation defined above.
//
//	 numStrNumSeps            NumericSeparatorDto
//
//	 This instance of NumericSeparatorDto contains the Decimal
//	 Separator character, Thousands Separator character and the
//	 Currency Symbol character used to parse the number strings
//	 passed through input paramters 'multiplier' and
//	 'multiplicands'.
//
//	 Return Values
//	 =============
//
//	 []string
//
//	 The products of the multiplication operation described above is
//	 returned as an array of number strings.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStrOutputToArray(
	multiplier string,
	multiplicands []string,
	inputNumSeps NumericSeparatorDto,
	outputNumSeps ...NumericSeparatorDto) ([]string, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrOutputToArray",
		"")

	if err != nil {
		return []string{}, err
	}

	err = inputNumSeps.IsValid(ePrefix.XCpy("Vaidating 'inputNumSeps'").String())

	if err != nil {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = inputNumSeps.IsValid(ePrefix.XCpy(\"Vaidating 'inputNumSeps'\").String())",
				ErrContext: "Error: Input Parameter 'inputNumSeps' is Invalid!",
				ErrMessage: err.Error(),
			}
	}

	if len(multiplier) == 0 {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "'multiplier' is an empty string.",
				ErrMessage: "Error: Input parameter 'multiplier' is INVALID!",
			}

	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "'multiplicands' is an empty array.",
				ErrMessage: "Error: Input parameter 'multiplicands' is INVALID!",
			}
	}

	var finalOutputNumSeps NumericSeparatorDto

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"output",
			"input",
			&inputNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"numStr\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of bPairBig1NumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&inputNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'inputNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &inputNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'inputNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	numSepsPair := NumericSeparatorPairDto{}

	err = numSepsPair.InputSeparators.CopyIn(&inputNumSeps, true)

	if err != nil {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.InputSeparators.CopyIn(&inputNumSeps, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = numSepsPair.OutputSeparators.CopyIn(&finalOutputNumSeps, true)

	if err != nil {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.OutputSeparators.CopyIn(&finalOutputNumSeps, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var multiplierBINum BigIntNum

	multiplierBINum, err = new(BigIntNum).NewNumStrWithNumSeps(multiplier, &numSepsPair)

	if err != nil {

		return []string{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "multiplierBINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
					"    multiplier, numStrNumSeps, numStrNumSeps)",
				ErrContext: fmt.Sprintf("multiplier= '%v'\nnumStrNumSeps= '%v'",
					multiplier, inputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplierBINum.GetNumStr()

	if err != nil {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplierBINum.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultArray := make([]string, lenMultiplicands)
	var multiplicandNumStr string
	var bPair BigIntPair
	var multiplicandBINum, finalResult BigIntNum

	for i := 0; i < lenMultiplicands; i++ {

		if len(multiplicands[i]) == 0 {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: fmt.Sprintf("Error: 'multiplicands[%d]' is an empty string!", i),
				}

		}

		multiplicandBINum, err = new(BigIntNum).
			NewNumStrWithNumSeps(multiplicands[i], &numSepsPair)

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("multiplicandBINum, err := new(BigIntNum).NewNumStrWithNumSeps(multiplicands[%d], numStrNumSeps, numStrNumSeps)", i),
					ErrContext: fmt.Sprintf("multiplicands[%d]= '%v'; numStrNumSeps= '%v'",
						i, multiplicands[i], inputNumSeps.String()),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplierBINum.GetNumStr()

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplierBINum.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)",
					ErrContext: fmt.Sprintf("multiplierBINum= '%v; multiplicandBINum= '%v'",
						multiplierNumStr, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"  multiplyPairWithNumSeps(bPair, numStrNumSeps,\n" +
						"  ePrefix.XCpy(\"bPair & numStrNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v'; bPair.Big2= '%v'\n"+
						" multiplicand index= '%d'",
						multiplierNumStr, multiplicandNumStr, i),
					ErrMessage: err.Error(),
				}
		}

		resultArray[i], err = finalResult.GetNumStr()

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "resultArray[i], err = finalResult.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return resultArray, err
}

// MultiplyNumStrSeries
//
//	This method receives one input parameter of Type string which
//	is classified as the 'multiplier'. The second input parameter
//	is a comma-delimited series of strings labeled,
//	'multiplicands'. The first element of the 'multiplicands'
//	series is multiplied by the 'multiplier' to produce a
//	'product'. That 'product' then replaces the 'multiplier' and is
//	multiplied by the next element in the 'multiplicands' series.
//	This process is repeated through the last element in the
//	'multiplicands' series.
//
//	Upon completion of the multiplication calculation, the
//	combined, final 'product' is returned as a Type 'BigIntNum'.
//
//	  Example:
//	    multiplier = 3
//	    multiplicands = strings "2","3","4"
//
//	    (1) 3 x 2 = 6
//	    (2) 6 x 3 = 18
//	    (3) 18 x 4 = 72
//	     The returned product is 72
//
//	Multiplication Operation
//	========================
//
//	In the multiplication operation, the number to be multiplied is
//	called the 'multiplicand', while the number of times the
//	multiplicand is multiplied comes from the 'multiplier'.
//	Usually, the 'multiplier' is placed first and the
//	'multiplicand' is placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	or result.
//
//			multiplier x multiplicand = product or result
//
//
//	Negative Values in Number Strings
//	=================================
//
//	The 'multiplier' string and 'multiplicands' string series
//	passed to this method are number strings which consist of a
//	string of numeric digits representing a numeric value. A
//	leading minus sign(-), or surrounding parentheses '()', may be
//	included to indicate a negative numeric value.
//
//	Fractional Digits in Number Strings
//	===================================
//
//	'multiplier' string and 'multiplicands' strings of numeric
//	digits may also include a delimiting decimal separator to
//	identify fractional digits to the right of the decimal
//	separator. These number strings are parsed based on the Numeric
//	Separators, including the decimal separator character, which
//	are specified by input parameter 'inputNumSeps'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	Numeric Separator characters are typically encapsulated in a
//	NumericSeparatorDto type.
//
//	Input parameter 'inputNumSeps', of type NumericSeparatorDto,
//	contains numeric separators specifying the Decimal Separator,
//	Thousands Separator, and Currency Symbol. These numeric
//	separators are required to parse the 'multiplier' and
//	'multiplicands' number strings. This parsing operation is in
//	turn used to convert those strings into numeric values for
//	calculation purposes.
//
//	Input parameter 'outputNumSeps' of type NumericSeparatorDto,
//	also contains numeric separators specifying the Decimal
//	Separator, Thousands Separator, and Currency Symbol.
//	The returned BigIntNum multiplication 'product' will contain
//	Numeric Separators specified by input parameter,
//	'outputNumSeps'.
//
//	Input Parameters
//	================
//
//	inputNumSeps            NumericSeparatorDto
//	  This instance of NumericSeparatorDto contains the Decimal
//	  Separator character, Thousands Separator character and the
//	  Currency Symbol character used to parse the number strings
//	  passed through input paramters 'multiplier' and
//	  'multiplicands'.
//
//	outputNumSeps            NumericSeparatorDto
//	  This instance of NumericSeparatorDto contains the Decimal
//	  Separator character, Thousands Separator character and the
//	  Currency Symbol character used to format the BigIntNum product
//	  returned by this method.
//
//	multiplier               string
//	  This number string contains the numeric value which serves as
//	  the initial multiplier in the multiplication operation defined
//	  above.
//
//	multiplicands            ...string
//	  This method is defined as a variadic function in that
//	  'multiplicands' is configured to accept multiple number strings
//	  which will serve as multiplicand values for the multiplication
//	  operation defined above.
//
//	  Although users have the option to supply zero strings for this
//	  parameter, failure to provide at least one valid number string
//	  will trigger an error.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The product of the multiplication operation described above is
//	  returned as BigIntNum type.
//
//	error
//	  If no errors are encountered during method execution, this
//	  return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStrSeries(
	inputNumSeps NumericSeparatorDto,
	outputNumSeps NumericSeparatorDto,
	multiplier string,
	multiplicands ...string) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrSeries",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if len(multiplier) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplier) == 0",
				ErrMessage: "Error: Input parameter 'multiplier' string is empty!",
			}
	}

	if len(multiplicands) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' parameter series is empty!",
			}
	}

	err = inputNumSeps.IsValid(ePrefix.
		XCpy("Validating numStrNumSeps").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = numStrNumSeps.IsValid(ePrefix.\n" +
					"XCpy(\"Validating numStrNumSeps\").String())",
				ErrContext: "Input parameter 'numStrNumSeps' is invalid!\n" +
					fmt.Sprintf("'numStrNumSeps'= '%v'",
						inputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = outputNumSeps.IsValid(ePrefix.
		XCpy("Validating outputNumSeps").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = outputNumSeps.IsValid(ePrefix.\n" +
					"    XCpy(\"Validating outputNumSeps\").String())",
				ErrContext: "Input parameter 'outputNumSeps' is invalid!\n" +
					fmt.Sprintf("'outputNumSeps'= '%v'",
						outputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	numSepsPair := NumericSeparatorPairDto{}

	err = numSepsPair.InputSeparators.CopyIn(&inputNumSeps, true)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.InputSeparators.CopyIn(&inputNumSeps, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = numSepsPair.OutputSeparators.CopyIn(&outputNumSeps, true)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSepsPair.OutputSeparators.CopyIn(&finalOutputNumSeps, true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var finalResult BigIntNum

	finalResult, err = new(BigIntNum).NewNumStrWithNumSeps(multiplier, &numSepsPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "multiplierBigINum, err = new(BigIntNum).\n" +
					"    NewNumStrWithNumSeps(multiplier, inputNumSeps, outputNumSeps)",
				ErrContext: fmt.Sprintf("'multiplier' number string = '%v'\n"+
					"numStrNumSeps= '%v'\noutputNumSeps= '%v'",
					multiplier, inputNumSeps.String(), outputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr, multiplicandNumStr string

	var bPair BigIntPair

	for idx, multiplicand := range multiplicands {

		multiplicandBINum, err := new(BigIntNum).
			NewNumStrWithNumSeps(multiplicand, &numSepsPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err := new(BigIntNum).\n" +
						"    NewNumStrWithNumSeps(multiplicand,\n" +
						"    inputNumSeps, outputNumSeps)",
					ErrContext: fmt.Sprintf("multiplicand= '%v' multiplicand index= '%d'\nnumStrNumSeps= '%v'\noutputNumSeps= '%v'",
						multiplicand, idx, inputNumSeps.String(), outputNumSeps.String()),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicandBINum.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicandBINum.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicand index= '%d'", idx),
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: fmt.Sprintf("multiplicand index= '%d'", idx),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).
			NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"    multiplierBigINum, multiplicandBINum)",
					ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicand= '%v'",
						multiplierNumStr, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, outputNumSeps,
				ePrefix.XCpy("bPair & outputNumSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"  multiplyPairWithNumSeps(bPair, outputNumSeps,\n" +
						"  ePrefix.XCpy(\"bPair & outputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\n"+
						"multiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

	}

	return finalResult, nil
}

// MultiplyNumStrDto
//
//	This method receives two NumStrDto instances and multiplies
//	their numeric values. The result or 'product' is returned as a
//	'BigIntNum' type.
//
//	  multiplier x multiplicand = product or result
//
//	Multiplication Operation
//	========================
//
//	In the multiplication operation, the number to be multiplied
//	is called the "multiplicand", while the number of times the
//	multiplicand is to be multiplied comes from the "multiplier".
//	Usually. the multiplier is placed first and the multiplicand
//	is placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	or result.
//
//	    multiplier x multiplicand = product or result
//
//	This method performs the multiplication operation described
//	above and afterward returns the cumulative result or 'product'
//	as a BigIntNum type.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	Numeric Separator characters are typically encapsulated in a
//	NumericSeparatorDto type.
//
//	The BigIntNum multiplication 'product' returned by this method
//	will contain numeric separators (Decimal Separator, Thousands
//	Separator and Currency Symbol) derived from one of two possible
//	sources.
//
//	Users have the option to supply an input parameter,
//	'outputNumSeps' of type NumericSeparatorDto. If this optional
//	parameter is provided, it will be used to configure the
//	returned BigIntNum multiplication 'product' from this method.
//	Note that the first valid NumericSeparatorDto in the
//	'outputNumSeps' series will be selected and used. There is no
//	need to provide more than one valid NumericSeparatorDto object
//	for input parameter 'outputNumSeps'.
//
//	If the optional input parameter 'outputNumSeps' is NOT
//	provided, the returned BigIntNum instance will be
//	configured using numeric separators copied from input
//	parameter, 'multiplier'.
//
//	Input Parameters
//	================
//
//	multiplier               NumStrDto
//	  This NumStrDto instance serves as the initial 'multiplier' in
//	  the multiplication operation defined above.
//
//	multiplicand             NumStrDto
//	  This instace of NumStrDto serves as the 'multiplicand' for the
//	  multiplication operation defined above.
//
//	outputNumSeps            ... NumericSeparatorDto
//	  This method is defined as a variadic function in that
//	  'outputNumSeps' is configured as an optional input parameter
//	  meaning that it is NOT required. The user can choose to
//	  provide a value for 'outputNumSeps', or not.
//
//	  If the user chooses to provide a valid 'NumericSeparatorDto'
//	  object for this parameter, it will be used to configure the
//	  'BigIntNum' product value returned by this method with the
//	  Numeric Separators copied from 'outputNumSeps'.
//
//	  Note that only the first valid NumericSeparatorDto in the
//	  'outputNumSeps' series will be selected and used. There is no
//	  need to provide more than one valid NumericSeparatorDto object
//	  for parameter 'outputNumSeps'.
//
//	  If the user chooses NOT to provide this optional parameter, the
//	  'BigIntNum' value returned by this  method will be configued
//	  using the 'NumericSeparatorDto' copied from the 'multiplier'
//	  input parameter.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The product of the multiplication operation described above is
//	  returned as BigIntNum type.
//
//	error
//	  If no errors are encountered during method execution, this
//	  return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStrDto(
	multiplier NumStrDto,
	multiplicand NumStrDto,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrDto",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(ePrefix.\n" +
					"XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Error: Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	err = multiplicand.IsValid(ePrefix.XCpy(
		"Validating 'multiplicand'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplicand.IsValid(ePrefix.\n" +
					"XCpy(\"Validating 'multiplicand'\").String())",
				ErrContext: "Error: Input parameter 'multiplicand' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr, multiplicandNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplicandNumStr, err = multiplicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplicandNumStr, err = multiplicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	// This method tests the validity of 'multiplier' and 'multiplicand'
	bPair, err := new(BigIntPair).NewNumStrDto(multiplier, multiplicand)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPair, err := new(BigIntPair).\n" +
					"  NewNumStrDto(multiplier, multiplicand)",
				ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicand= '%v'",
					multiplierNumStr, multiplicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	var finalResult BigIntNum

	finalResult, err = new(bigIntMathMultiplyNanobot).
		multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
			ePrefix.XCpy("bPair & finalOutputNumSeps"))

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
					"  multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n" +
					"  ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
				ErrContext: fmt.Sprintf("bPair.Big1= '%v'; bPair.Big2= '%v'",
					multiplierNumStr, multiplicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// MultiplyNumStrDtoArray
//
//		Overview
//		========
//
//		Receives one NumStrDto instance which is classified as the
//		'multiplier'. The second input parameter is an array of
//		NumStrDto Types labeled, 'multiplicands'. The first element of
//		the 'multiplicands' array is multiplied by the 'multiplier' to
//		produce a 'product'. That 'product' then replaces the
//		'multiplier' and is multiplied by the next element in the
//		'multiplicands' array. This process is repeated through the
//		last element in the array when the combined, final 'product'
//		is returned as a Type 'BigIntNum'.
//
//	   Example:
//	     multiplier = 3
//	     multiplicands = [3]NumStrDto{2,3,4}
//
//	     (1) 3 x 2 = 6
//	     (2) 6 x 3 = 18
//	     (3) 18 x 4 = 72
//	      The returned product is 72
//
//		Multiplication Operation
//		========================
//
//		In the multiplication operation, the number to be multiplied is
//		called the "multiplicand", while the number of times the
//		multiplicand is to be multiplied comes from the "multiplier".
//		Usually, the multiplier is placed first and the multiplicand is
//		placed second.
//
//		For example, in the problem 5 x 3 equals 15, the 5 is the
//		'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//		or result.
//
//		  multiplier x multiplicand = product or result
//
//		This method performs the multiplication operation described
//		above and afterward returns the result or 'product' as a
//		BigIntNum type.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal separator character,
//		Thousands separator character and Currency Symbol characters.
//		These separator characters are used to format and display
//		numeric values as number strings.
//
//		Users have the option to supply an input parameter,
//		'outputNumSeps' of type NumericSeparatorDto, which will be
//		used to configure the returned BigIntNum multiplication
//		'result' from this method. Note that the first valid
//		NumericSeparatorDto in the 'outputNumSeps' series will be
//		selected and used. There is no need to provide more than
//	 one valid NumericSeparatorDto object for parameter
//	 'outputNumSeps'.
//
//		If the optional input parameter 'outputNumSeps' is NOT
//		provided, the returned BigIntNum instance will be
//		configured using numeric separators copied from input
//		parameter, 'multiplier'.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               NumStrDto
//
//		This NumStrDto instance serves as the initial multiplier in the
//	 multiplication operation defined above.
//
//	 multiplicands           []NumStrDto
//
//	 This array of NumStrDto objects serves as multiplicands in the
//	 multiplication operation defined above.
//
//	 outputNumSeps            ... NumericSeparatorDto
//
//	 This method is defined as a variadic function in that
//	 'outputNumSeps' is configured as an optional input parameter
//	 meaning that it is NOT required. The user can choose to
//	 provide a value for 'outputNumSeps', or not.
//
//	 If the user chooses to provide a valid 'NumericSeparatorDto'
//	 object for this parameter, it will be used to configure the
//	 'BigIntNum' product value returned by this method.
//
//	 Note that the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//	 Be advised that if the user chooses NOT to provide this
//	 optional parameter, the 'BigIntNum' value returned by this
//	 method will be configued using the 'NumericSeparatorDto'
//	 copied from 'multiplier'
func (bMultiply *BigIntMathMultiply) MultiplyNumStrDtoArray(
	multiplier NumStrDto,
	multiplicands []NumStrDto,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrDtoArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' array is empty!",
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: "Error: Failed to acquire 'multiplier' numSeps",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	// This method tests the validity of 'multiplier'
	finalResult, err := new(BigIntNum).NewNumStrDto(multiplier)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err := new(BigIntNum).NewNumStrDto(multiplier)",
				ErrContext: fmt.Sprintf("multiplier= '%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var multiplicandBINum BigIntNum
	var multiplicandNumStr string
	var bPair BigIntPair

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.XCpy(
			fmt.Sprintf("Validating multiplicands[%d]", i)).String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = multiplicands[i].IsValid(ePrefix.XCpy(\n" +
						fmt.Sprintf("Validating multiplicands[%d])).String()", i),
					ErrContext: fmt.Sprintf("multiplicands[%d] is INVALID!", i),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicands[i].GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf(
						"multiplicandNumStr, err = multiplicands[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// This method tests the validity of multiplicands[i]
		multiplicandBINum, err = new(BigIntNum).NewNumStrDto(multiplicands[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err = new(BigIntNum).NewNumStrDto(\n" +
						fmt.Sprintf("multiplicands[%d])", i),
					ErrContext: fmt.Sprintf("multiplicands[%d]= '%v'",
						i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"    finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult= '%v'; multiplicands[%d]= '%v'",
						multiplierNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n " +
						"ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\nmultiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, i),
					ErrMessage: err.Error(),
				}
		}

	}

	return finalResult, err
}

// MultiplyNumStrDtoOutputToArray
//
//		Overview
//		========
//
//		This method receives one input parameter of Type NumStrDto
//		which is classified as the 'multiplier'. The second input
//		parameter is an array of NumStrDto Types labeled,
//		'multiplicands'.
//
//		Each element of the 'multiplicands' array is multiplied by the
//		'multiplier'. The result or 'product' is then stored in a
//		results array which is returned to the calling function.
//
//		  Example:
//
//
//		  Multiplicands                                     Output
//		   Multiplier             Array                     Array
//
//		      3       x    multiplicands[0] = 2   =   outputarray[0] =  6
//		      3       x    multiplicands[1] = 3   =   outputarray[1] =  9
//		      3       x    multiplicands[2] = 4   =   outputarray[2] = 12
//		      3       x    multiplicands[3] = 5   =   outputarray[3] = 15
//		      3       x    multiplicands[4] = 6   =   outputarray[4] = 18
//		      3       x    multiplicands[5] = 7   =   outputarray[5] = 21
//
//		Multiplation Operation
//		======================
//
//		In the multiplication operation, the number to be multiplied is
//		called the "multiplicand", while the number of times the
//		multiplicand is to be multiplied comes from the "multiplier".
//		Usually, the multiplier is placed first and the multiplicand is
//		placed second.
//
//		For example, in the problem 5 x 3 equals 15, the 5 is the
//		'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//		or result.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 Each element of the []NumStrDto multiplication 'Product' array
//	 returned by this method will contain numeric separators
//	 (Decimal Separator, Thousands Separator and Currency Symbol)
//	 derived from one of two possible sources.
//
//		Users have the option to supply an input parameter,
//		'outputNumSeps' of type NumericSeparatorDto, which will be
//		used to configure the returned []NumStrDto multiplication
//		'product' array from this method. If the optional input
//		parameter 'outputNumSeps' is provided, it will be used to
//		format each element in the []NumStrDto array returned by
//		this method. Note that the only first valid NumericSeparatorDto
//		in the 'outputNumSeps' series will be selected and used. There
//		is no need to provide more than one valid NumericSeparatorDto
//		object for input parameter 'outputNumSeps'.
//
//		If the optional input parameter 'outputNumSeps' is NOT
//		provided, the returned BigIntNum instance will be
//		configured using numeric separators copied from input
//		parameter, 'multiplier'.
//
//	 Input Parameters
//	 ================
//
//	 multiplier               NumStrDto
//
//		This NumStrDto instance serves as the initial 'multiplier' in
//		the multiplication operation defined above.
//
//	 multiplicands           []NumStrDto
//
//	 This array of NumStrDto objects serves as 'multiplicands' for
//	 the multiplication operation defined above.
//
//	 outputNumSeps            ... NumericSeparatorDto
//
//	 This method is defined as a variadic function in that
//	 'outputNumSeps' is configured as an optional input parameter
//	 meaning that it is NOT required. The user can choose to
//	 provide a value for 'outputNumSeps', or not.
//
//	 If the user chooses to provide a valid 'NumericSeparatorDto'
//	 object for this parameter, it will be used to configure each
//	 element in the '[]NumStrDto' product value returned by this
//	 method.
//
//	 Note that only the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//	 Be advised that if the user chooses NOT to provide this
//	 optional parameter, each element in the '[]NumStrDto' array
//	 returned by this method will be configued using the
//	 'NumericSeparatorDto' copied from the 'multiplier' input
//	 parameter.
//
//	 Return Values
//	 =============
//
//	 []NumStrDto
//
//	 The product of the multiplication operation described above is
//	 returned as an array of type NumStrDto.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStrDtoOutputToArray(
	multiplier NumStrDto,
	multiplicands []NumStrDto,
	outputNumSeps ...NumericSeparatorDto) ([]NumStrDto, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrDtoOutputToArray",
		"")

	if err != nil {
		return []NumStrDto{}, err
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' array is empty!",
			}
	}

	var finalOutputNumSeps, multiplierNumSeps NumericSeparatorDto

	multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err = multiplier.GetNumericSeparatorsDto()",
				ErrContext: fmt.Sprintf("multiplier= '%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"multiplier",
			&multiplierNumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "foundPrimaryNumSep, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of multiplierNumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&multiplierNumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'multiplierNumSeps' -> 'finalOutputNumSeps'"))

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &multiplierNumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'multiplierNumSeps' -> 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	var multiplierBINum, multiplicandBINum BigIntNum

	// This method will test the validity of 'multiplier'
	multiplierBINum, err = new(BigIntNum).NewNumStrDto(multiplier)

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierBINum, err = new(BigIntNum).NewNumStrDto(multiplier)",
				ErrContext: fmt.Sprintf("multiplier= '%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	multiplierNumStr, err = multiplierBINum.GetNumStr()

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplierBINum.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultArray := make([]NumStrDto, lenMultiplicands)

	var multiplicandNumStr string
	var bPair BigIntPair
	var finalResult BigIntNum

	for i := 0; i < lenMultiplicands; i++ {

		err = multiplicands[i].IsValid(ePrefix.XCpy(
			fmt.Sprintf("Validating multiplicands[%d]", i)).String())

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = multiplicands[i].IsValid(ePrefix.XCpy(\n" +
						fmt.Sprintf("Validating multiplicands[%d])).String()", i),
					ErrContext: fmt.Sprintf("multiplicands[%d] is INVALID!", i),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicands[i].GetNumStr()

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf(
						"multiplicandNumStr, err = multiplicands[%d].GetNumStr()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// This method will test the validity of multiplicands[i]
		multiplicandBINum, err = new(BigIntNum).NewNumStrDto(multiplicands[i])

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err = new(BigIntNum).NewNumStrDto(\n" +
						fmt.Sprintf("multiplicands[%d])", i),
					ErrContext: fmt.Sprintf("multiplicands[%d]= '%v'",
						i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(multiplierBINum, multiplicandBINum)

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(\n" +
						"    multiplierBINum, multiplicandBINum)",
					ErrContext: fmt.Sprintf("multiplier= '%v'; multiplicands[%d]= '%v'",
						multiplierNumStr, i, multiplicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n " +
						"ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\nmultiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, i),
					ErrMessage: err.Error(),
				}
		}

		resultArray[i], err = finalResult.GetNumStrDto()

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: fmt.Sprintf("resultArray[%d], err = finalResult.GetNumStrDto()", i),
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return resultArray, err
}

// MultiplyNumStrDtoSeries
//
//	 Overview
//	 ========
//
//	 This method receives one input parameter of Type NumStrDto
//	 which is classified as the 'multiplier'. The second input
//	 parameter is a comma-delimited series of NumStrDto objects
//	 labeled, 'multiplicands'. The first element of the
//	 'multiplicands' series is multiplied by the 'multiplier' to
//	 produce a 'product'. That 'product' then replaces the
//	 'multiplier' and is multiplied by the next element in the
//	 'multiplicands' series. This process is repeated through the
//	 last element in the series. Afterward, the combined final
//	 'product' is returned as a Type 'BigIntNum'.
//
//	   Example:
//	     multiplier = 3
//	     multiplicands = [3]IntAry{2,3,4}
//
//	     (1) 3 x 2 = 6
//	     (2) 6 x 3 = 18
//	     (3) 18 x 4 = 72
//	     The returned product is 72
//
//	 Multiplication Operation
//	 ========================
//
//	 In the multiplication operation, the number to be multiplied
//	 is called the "multiplicand", while the number of times the
//	 multiplicand is to be multiplied comes from the "multiplier".
//	 Usually, the multiplier is placed first and the multiplicand
//	 is placed second.
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
//	 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//	     multiplier x multiplicand = product or result
//
//	 This method performs the multiplication operation described above and afterward returns the
//	 result or 'product' as a BigIntNum type.
//
//	 Numeric Separators
//	 ==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 The BigIntNum multiplication 'product' returned by this method
//	 will contain numeric separators (Decimal Separator, Thousands
//	 Separator and currency symbol) copied from input parameter,
//
// 'multiplier'.
func (bMultiply *BigIntMathMultiply) MultiplyNumStrDtoSeries(
	multiplier NumStrDto,
	multiplicands ...NumStrDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyNumStrDtoSeries",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = multiplier.IsValid(
		ePrefix.XCpy("Validating 'multiplier'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(\n" +
					"ePrefix.XCpy(\"Validating 'multiplier'\").String())",
				ErrContext: "Input parameter 'multiplier' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var multiplierNumStr string

	multiplierNumStr, err = multiplier.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumStr, err = multiplier.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	finalOutputNumSeps, err := multiplier.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalOutputNumSeps, err := multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = finalOutputNumSeps.IsValid(ePrefix.XCpy("Testing 'multiplier' NumSeps").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierNumSeps, err := multiplier.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(multiplicands) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(multiplicands) == 0",
				ErrMessage: "Error: 'multiplicands' variadic func series is empty!",
			}
	}

	var finalResult, multiplicandBINum BigIntNum

	// This method will test the validity of 'multiplier'
	finalResult, err = new(BigIntNum).NewNumStrDto(multiplier)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "finalResult, err := new(BigIntNum).NewNumStrDto(multiplier)",
				ErrContext: fmt.Sprintf("multiplier = '%v'", multiplierNumStr),
				ErrMessage: err.Error(),
			}
	}

	var bPair BigIntPair

	var multiplicandNumStr string

	for idx, multiplicand := range multiplicands {

		err = multiplicand.IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'multiplicand' idx=%d", idx)).String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = multiplicand.IsValid(ePrefix.XCpy(\n" +
						"    fmt.Sprintf(\"Validating 'multiplicand' idx=%d\",idx))\n" +
						"    .String())",
					ErrContext: fmt.Sprintf("'multiplicand' (idx=%d) is invalid!", idx),
					ErrMessage: err.Error(),
				}
		}

		multiplicandNumStr, err = multiplicand.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandNumStr, err = multiplicand.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		multiplierNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplierNumStr, err = finalResult.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// This method will test the validity of multiplicand
		multiplicandBINum, err = new(BigIntNum).NewNumStrDto(multiplicand)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "multiplicandBINum, err = new(BigIntNum).NewNumStrDto(multiplicand)",
					ErrContext: fmt.Sprintf("'multiplicand'= '%v'; 'multiplicand' idx=%d",
						multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

		bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bPair, err = new(BigIntPair).NewBigIntNum(finalResult, multiplicandBINum)",
					ErrContext: fmt.Sprintf("finalResult= '%v'; 'multiplicand'= '%v'; 'multiplicand' idx=%d",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = new(bigIntMathMultiplyNanobot).
			multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
				ePrefix.XCpy("bPair & finalOutputNumSeps"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "finalResult, err = new(bigIntMathMultiplyNanobot).\n" +
						"multiplyPairWithNumSeps(bPair, finalOutputNumSeps,\n " +
						"ePrefix.XCpy(\"bPair & finalOutputNumSeps\"))",
					ErrContext: fmt.Sprintf("bPair.Big1= '%v' bPair.Big2= '%v'\nmultiplicands index= '%d'",
						multiplierNumStr, multiplicandNumStr, idx),
					ErrMessage: err.Error(),
				}
		}
	}

	return finalResult, err
}

// MultiplyPair
//
//	 Multiplication Operation
//	 ========================
//
//	 This method receives a 'BigIntPair' instance labeled 'bPair'
//	 and proceeds to multiply 'bPair.Big1' times 'bPair.Big2'. Both
//	 'bPair.Big1' and 'bPair.Big2' are of type 'BigIntNum'.
//
//	   multiplier x multiplicand = product or result
//	   bPair.Big1 x bPair.Big2   = product or result
//
//	 The result of this multiplication operation (a.k.a. 'product')
//	 is returned as a BigIntNum type.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 Numeric Separator characters are typically encapsulated in a
//	 NumericSeparatorDto type.
//
//	 The returned BigIntNum multiplication 'Product' will contain
//	 numeric separators (Decimal Separator, Thousands Separator and
//	 Currency Symbol) derived from one of two possible sources.
//
//		Users have the option to supply an input parameter,
//		'outputNumSeps' of type NumericSeparatorDto. If this optional
//	 parameter is provided, it will be used to configure the
//	 returned BigIntNum multiplication 'product' from this method.
//	 Note that the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//		If the optional input parameter 'outputNumSeps' is NOT
//		provided, the returned BigIntNum instance will be
//		configured using numeric separators copied from input
//		parameter, 'bPair.Big1'.
//
//	 Input Parameters
//	 ================
//
//	 bPair                    BigIntPair
//
//		This instance of BigIntPair contains two numeric values of type
//	 'BigIntNum'. These two 'BigIntNum' values are referred to as
//	 bPair.Big1 and bPair.Big2. This method is designed to multiply
//	 bPair.Big1 times bPair.Big2 and return the product as a
//	 'BigIntNum' type.
//
//	 outputNumSeps            ... NumericSeparatorDto
//
//	 This method is defined as a variadic function in that
//	 'outputNumSeps' is configured as an optional input parameter
//	 meaning that it is NOT required. The user can choose to
//	 provide a value for 'outputNumSeps', or not.
//
//	 If the user chooses to provide a valid 'NumericSeparatorDto'
//	 object for this parameter, it will be used to configure the
//	 'BigIntNum' product value returned by this method.
//
//	 Note that the first valid NumericSeparatorDto in the
//	 'outputNumSeps' series will be selected and used. There is no
//	 need to provide more than one valid NumericSeparatorDto object
//	 for parameter 'outputNumSeps'.
//
//	 Be advised that if the user chooses NOT to provide this
//	 optional parameter, the 'BigIntNum' value returned by this
//	 method will be configued using the 'NumericSeparatorDto'
//	 copied from 'bPair.Big1'
//
//	 Return Values
//	 =============
//
//	 BigIntNum
//
//	 The product of the multiplication operation described above is
//	 returned as BigIntNum type.
//
//	 error
//
//	 If no errors are encountered during method execution, this
//	 return parameter is set to 'nil'.
func (bMultiply *BigIntMathMultiply) MultiplyPair(
	bPair BigIntPair,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathMultiply.MultiplyPair",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = bPair.IsValid(ePrefix.XCpy("Validating 'bPair'.").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = bPair.IsValid(ePrefix.XCpy(\n" +
					"\"Validating 'bPair'.\").String())",
				ErrContext: "Input parameter 'bPair' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, bPairBig1NumSeps NumericSeparatorDto

	bPairBig1NumSeps, err = bPair.Big1.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPairBig1NumSeps, err := bPair.Big1.\n" +
					"GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"bPairBig1",
			&bPairBig1NumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of bPairBig1NumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&bPairBig1NumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'bPairBig1NumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &bPairBig1NumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'bPairBig1NumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return new(bigIntMathMultiplyNanobot).
		multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
			ePrefix.XCpy("bPair & finalOutputNumSeps"))
}
