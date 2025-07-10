package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
)

// BigIntMathSubtract - Contains methods used to perform subtraction
// operations on *big.Int numeric types.
//
//	minuend − subtrahend = difference
type BigIntMathSubtract struct {
	Input  BigIntPair // BigIntPair.Big1 = minuend  BigIntPair.Big2 = subtrahend
	Result BigIntNum  // The result of the subtraction otherwise known as the 'difference'
}

// BigIntSubtract
//
//	Performs the subtraction operation on two numeric values. The
//	'minuend' is the number from which the 'subtrahend' is subtracted
//	in order to generate a result or difference between the two
//	numbers.
//
//	In the subtraction operation:
//
//	  'minuend' - 'subtrahend' = difference or result
//
//	This method provides for the subtraction of fixed length floating
//	point values by means of integer and precision specification
//	pairs.
//
//	As an example, consider the following subtraction operation:
//
//	  752.314 - 21.67894 = 730.63506 = difference
//
//	In this case the 'minuend', 'subtrahend' and 'difference' would be
//	configured as follows:
//
//	    minuend             = 752314
//	    minuendPrecision    = 3
//	    subtrahend          = 2167894
//	    subtrahendPrecision = 5
//
//	    difference          = 73063506
//	    differencePrecision = 5
//
//	In this way, the method uses integer, precision pairs to define
//	fixed length floating point numbers.
//
//	Note
//	====
//
//	This function will delete all trailing fractional zeros from
//	the result or difference.
//
//	Input Parameters
//	================
//
//	minuend                  *big.Int
//	  The number from which the subtrahend will be subtracted.
//
//	minuendPrecision         *big.Int
//	  The 'minuend' precision or numeric digits after the decimal
//	  point. 'minuendPrecision' must be greater than or equal to zero.
//
//	subtrahend               *big.Int
//	  The number to be subtracted from the 'minuend'.
//
//	subtrahendPrecision      *big.Int
//	  The 'subtrahend' precision or numeric digits after the decimal
//	  point. 'subtrahendPrecision' must be greater than or equal to
//	  zero.
//
//	Return Values
//	=============
//
//	difference               *big.Int
//	  The difference or result of the subtraction operation returned
//	  as a *big.Int type.
//
//	differencePrecision      *big.Int
//	  The precision specification for the returned subtraction
//	  'result'.
//
//	  Precision specifies the number of fractional digits to the right
//	  of the decimal place. 'differencePrecision' will always be
//	  greater than or equal to zero.
//
//	  Taken together, 'difference' and 'differencePrecision' can
//	  define a fixed length floating point number.
//
//	err                      error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (bSubtract *BigIntMathSubtract) BigIntSubtract(
	minuend *big.Int,
	minuendPrecision *big.Int,
	subtrahend *big.Int,
	subtrahendPrecision *big.Int) (difference *big.Int, differencePrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.BigIntSubtract",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return new(bigIntMathSubtractNanobot).bigIntSubtract(
		minuend, minuendPrecision, subtrahend, subtrahendPrecision, ePrefix)
}

// FixedDecimalSubtract
//
//	Performs the subtraction operation on two BigIntFixedDecimal
//	types. The subtraction result is also returned as a
//	BigIntFixedDecimal type.
//
//	Examples
//	=========
//
//	In the subtraction operation:
//
//	'minuend' - 'subtrahend' = 'difference' or result
//
//	752.314   -   21.67894   = 730.63506 = difference
//
//	For this method 'minuend', 'subtrahend' and 'difference' are
//	configured as BigIntFixedDecimal types.
//
//	The BigIntFixedDecimal is used to defined fixed length floating
//	point numbers and is defined as follows:
//
//	type BigIntFixedDecimal struct {
//
//	  integerNum  *big.Int
//	    All numeric digits, both integer and fractional, necessary to
//	    define a fixed length floating point number.
//
//	    The number of digits to the right of the decimal place is
//	    specified by the data field, 'BigIntFixedDecimal.precision'.
//
//	  precision   uint
//	    Specifies the number of digits to the right of the decimal
//	    place in the series of numeric digits represented by data
//	    field BigIntFixedDecimal.integerNum.
//
//	}
//
//	To represent the floating point number 52.459	a BigIntDecimal
//	Structure  would be configured as follows:
//
//	    BigIntFixedDecimal.integerNum = 52459
//	    BigIntFixedDecimal.precision  = 3
//
//	As an example consider the following subtraction operation:
//
//	    752.314 - 21.67894 = 730.63506 = difference
//
//	In this case the 'minuend', 'subtrahend' and 'difference' consist
//	of BigIntFixedDecimal types configured as follows:
//
//	           minuend.integerNum     = 752314
//	           minuend.precision      = 3
//	           subtrahend.integerNum  = 2167894
//	           subtrahend.precision   = 5
//
//	           difference.integerNum  = 73063506
//	           difference.precision   = 5
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' to the returned instance of
//	'difference' (type BigIntFixedDecimal).
//
//	Input Parameters
//	================
//
//	minuend                  BigIntFixedDecimal
//	  The number from which the subtrahend will be subtracted.
//
//	subtrahend               BigIntFixedDecimal
//	  The number to be subtracted from the 'minuend'.
//
//	Return Values
//	=============
//
//	difference               BigIntFixedDecimal
//	  The difference or result of the subtraction operation returned
//	  as a BigIntFixedDecimal type.
//	        'minuend' - 'subtrahend' = 'difference'
//
//	err                      error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (bSubtract *BigIntMathSubtract) FixedDecimalSubtract(
	minuend BigIntFixedDecimal,
	subtrahend BigIntFixedDecimal) (difference BigIntFixedDecimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.FixedDecimalSubtract",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
					"'minuend' contains invalid Numeric Separators.",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).fixedDecimalSubtract(
		numSeps, false, minuend, true, subtrahend, true, ePrefix)
}

// SubtractBigInts
//
//	Performs the subtraction operation and returns the 'difference' as
//	a type BigIntNum.
//
//	In the subtraction operation:
//
//	  b1 - b2 = difference or result
//	  'minuend' - 'subtrahend' = difference or result
//	  b1 = 'minuend'
//	  b2 = 'subtrahend'
//
//	Input Parameters
//	================
//
//	minuend                  *big.Int
//	  The number from which the subtrahend will be subtracted.
//
//	minuendPrecision         uint
//	  The 'minuend' precision or numeric digits after the decimal
//	  point.
//
//	subtrahend               *big.Int
//	  The number to be subtracted from the 'minuend'.
//
//	subtrahendPrecision      uint
//	  The 'subtrahend' precision or numeric digits after the decimal
//	  point.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  After the subtraction operation, the 'difference' or 'result' is
//	  returned as a Type BigIntNum.
//
//	  The returned BigIntNum 'result' will contain USA default numeric
//	  separators (decimal separator, thousands separator and currency
//	  symbol).
//
//	err                      error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (bSubtract *BigIntMathSubtract) SubtractBigInts(
	minuend *big.Int,
	minuendPrecision uint,
	subtrahend *big.Int,
	subtrahendPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractBigInts",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps := new(NumericSeparatorDto).NewUSADefaults()

	return new(bigIntMathSubtractMacrobot).subtractBigInts(
		numSeps, minuend, minuendPrecision, subtrahend, subtrahendPrecision, ePrefix)
}

// SubtractBigIntNums
//
//	Receives two 'BigIntNum' instances and proceeds to subtract
//	'subtrahend' from 'minuend' returning 'difference' as a type
//	BigIntNum.
//
//	Subtraction Operation
//	=====================
//
//	In the subtraction operation:
//
//	      'minuend' - 'subtrahend' = difference or result
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' to the returned instance of
//	'difference' (type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractBigIntNums(
	minuend BigIntNum, subtrahend BigIntNum) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractBigIntNums",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		&minuend,
		ePrefix.XCpy(" Validating input parameter 'minuend'"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	difference, err = new(bigIntMathSubtractMacrobot).subtractBigIntNums(numSeps, true, &minuend, true, &subtrahend, true, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "difference, err = new(bigIntMathSubtractMacrobot).subtractBigIntNums(\n" +
					"  numSeps, true, &minuend, true, &subtrahend, true,  ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return difference, nil
}

// SubtractBigIntNumArray
//
//	Receives one BigIntNum which is classified as the 'minuend'.
//
//	The second input parameter is an array of BigIntNum Types labeled,
//	'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend'. The summary result of this subtraction operation is
//	then returned as a BigIntNum type.
//
//	In the subtraction operation:
//
//	  b1 = 'minuend'
//	  b2 = 'subtrahend'
//	  b1 - b2 = difference or result
//	  'minuend' - 'subtrahend' = difference or result
//
//	In this method, the 'subtrahend' is an array of BigIntNum Types.
//
//	After the subtraction operation, the 'difference' or 'result' is
//	returned as a Type BigIntNum.
//
//	                  subtrahends
//	minuend              Array                   difference
//
//	  50      -     subtrahends[0] = 2
//	  48      -     subtrahends[1] = 3
//	  45      -     subtrahends[2] = 4
//	  41      -     subtrahends[3] = 5
//	  36      -     subtrahends[4] = 6
//	  30      -     subtrahends[5] = 9
//
//	                      Final Returned 'difference' = 21
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' to the returned	instance of
//	'difference' (type BigIntFixedDecimal).
func (bSubtract *BigIntMathSubtract) SubtractBigIntNumArray(
	minuend BigIntNum,
	subtrahends []BigIntNum) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractBigIntNumArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractBigIntNumArray(
		numSeps, true, minuend, true, subtrahends, true, ePrefix)
}

// SubtractBigIntNumOutputToArray
//
//	The first input parameter to this method is a BigIntNum Type
//	labeled, 'minuend'.  The second input parameter is an array of
//	BigIntNum types labeled 'subtrahends'.
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend' value with the 'dfference', or result, output to another
//	'results' array of BigIntNum types which is then returned to the
//	calling function.
//
//	Example
//	=======
//
//	              subtrahends                 Output
//	Minuend          Array                     Array
//
//	  10      -  subtrahends[0] = 2    =   outputarray[0] = 8
//	  10      -  subtrahends[1] = 3    =   outputarray[1] = 7
//	  10      -  subtrahends[2] = 4    =   outputarray[2] = 6
//	  10      -  subtrahends[3] = 5    =   outputarray[3] = 5
//	  10      -  subtrahends[4] = 6    =   outputarray[4] = 4
//	  10      -  subtrahends[5] = 9    =   outputarray[5] = 1
//
//	Each of the BigIntNum instances included in the array of BigIntNum
//	subtraction results returned by this method, will contain numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) copied from input parameter 'minuend'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to each member of the returned BigIntNum
//	array.
func (bSubtract *BigIntMathSubtract) SubtractBigIntNumOutputToArray(
	minuend BigIntNum,
	subtrahends []BigIntNum) ([]BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractBigIntNumOutputToArray",
		"")

	if err != nil {
		return []BigIntNum{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return []BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractBigIntNumOutputToArray(
		numSeps, true, minuend, true, subtrahends, true, ePrefix)
}

// SubtractBigIntNumSeries
//
//	Receives one BigIntNum which is classified as the 'minuend'.
//
//	The second  parameter is a series of BigIntNum instances which
//	resents the 'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each member of the 'subtrahends' series is subtracted from the
//	'minuend' to produce a cumulative result.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = cumulative result or 'difference'
//
//	In this method, the 'subtrahend' is a series of BigIntNum objects.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	               subtrahends
//	minuend          series               difference
//
//	  50       -        2
//	  48       -        3
//	  45       -        4
//	  41       -        5
//	  36       -        6
//	  30       -        9
//
//	              Final Returned 'difference' = 21
//
//	Variadic Function
//	=================
//
//	In this method, the 'subtrahends' parameter is defined as a series
//	of Type BigIntNum objects.
//
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide one value, multiple values,
//	or zero values. Note, that if the user chooses NOT to submit at
//	least one valid BigIntNum object for the 'subtrahends' parameter,
//	an error will be returned.
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
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractBigIntNumSeries(
	minuend BigIntNum,
	subtrahends ...BigIntNum) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractBigIntNumSeries",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractBigIntNumSeries(
		numSeps, false, minuend, true, true, ePrefix, subtrahends...)
}

// SubtractDecimals
//
//	Performs the subtraction operation on two Decimal Types.
//
//	Subtraction Operation
//	=====================
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//	    decMinuend     = 'minuend'
//	    decSubtrahend  = 'subtrahend'
//	    decMinuend  -  decSubtrahend  =  difference or result
//
//	After the subtraction operation, the 'difference' or 'result' is
//	returned as a Type BigIntNum.
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
//	This method will copy the Numeric Separators configured for input
//	parameter 'decMinuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractDecimals(
	decMinuend Decimal,
	decSubtrahend Decimal) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractDecimals",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps, err := decMinuend.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := decMinuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractDecimals(
		numSeps, true, &decMinuend, true,
		&decSubtrahend, true, ePrefix)
}

// SubtractDecimalArray
//
//	Receives one Decimal parameter which is classified as the
//	'minuend'. The second input parameter is an array of Decimal Types
//	labeled, 'subtrahends'.
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend'. The summary result of this subtraction operation is
//	then returned as a BigIntNum type.
//
//	Subtraction Operation
//	=====================
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahends' = difference or result
//	    d1 = 'minuend'
//	    d2 = 'subtrahend'
//	    d1 - d2 = difference or result
//
//	In this method, the 'subtrahends' is an array of Decimal Types.
//	Each element in the 'subtrahends' array is subtracted from the
//	value of 'minuend'. After the subtraction operation, the
//	cumulative 'difference' or 'result' is returned as a Type
//	BigIntNum.
//
//	                  subtrahends
//	minuend              Array                   difference
//
//	  50      -     subtrahends[0] = 2
//	  48      -     subtrahends[1] = 3
//	  45      -     subtrahends[2] = 4
//	  41      -     subtrahends[3] = 5
//	  36      -     subtrahends[4] = 6
//	  30      -     subtrahends[5] = 9
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
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractDecimalArray(
	minuend Decimal,
	subtrahends []Decimal) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractDecimalArray",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractDecimalArray(
		numSeps, true, minuend, true, subtrahends, true, ePrefix)
}

// SubtractDecimalOutputToArray
//
//	The first input parameter to this method is a Decimal Type
//	labeled, 'minuend'.  The second input parameter is an array of
//	Decimal types labeled 'subtrahends'.
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend' value with the 'dfference', or result, output to another
//	'results' array of Decimal types which is then returned to the
//	calling function.
//
//	Example
//	=======
//
//	                     subtrahends                 Output
//	  minuend              Array                      Array
//
//	    10      -     subtrahends[0] = 2    =    outputarray[0] =  8
//	    10      -     subtrahends[1] = 3    =    outputarray[1] =  7
//	    10      -     subtrahends[2] = 4    =    outputarray[2] =  6
//	    10      -     subtrahends[3] = 5    =    outputarray[3] =  5
//	    10      -     subtrahends[4] = 6    =    outputarray[4] =  4
//	    10      -     subtrahends[5] = 9    =    outputarray[5] =  1
//
//	Numeric Separators
//	==================
//
//	Each array element of the []Decimal 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'result' array ([]Decimal) returned by this subtraction
//	operation will contain array elements with numeric separators
//	(decimal separator, thousands separator and currency symbol) which
//	have been copied from input parameter 'minuend'.
func (bSubtract *BigIntMathSubtract) SubtractDecimalOutputToArray(
	minuend Decimal,
	subtrahends []Decimal) (results []Decimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractDecimalOutputToArray",
		"")

	if err != nil {
		return results, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return results,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractDecimalOutputToArray(
		numSeps, true, minuend, true, subtrahends,
		true, ePrefix)
}

// SubtractDecimalSeries
//
//	Receives one Decimal Type which is classified as the 'minuend'.
//	The second input parameter, 'subtrahends' is a series of Type Decimal.
//
//	Subtraction Operation
//	=====================
//
//	Each member of the 'subtrahends' series is subtracted from the
//	'minuend' to produce a cumulative result.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = cumulative result or 'difference'
//
//	In this method, the 'subtrahend' is a series of Decimal objects.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	               subtrahends
//	minuend          series               difference
//
//	  50       -        2
//	  48       -        3
//	  45       -        4
//	  41       -        5
//	  36       -        6
//	  30       -        9
//
//	              Final Returned 'difference' = 21
//
//	Variadic Function
//	=================
//
//	In this method, the 'subtrahends' parameter is defined as a series
//	of Type Decimal objects.
//
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide one value, multiple values,
//	or zero values. Note, that if the user chooses NOT to submit at
//	least one valid Decimal object for the 'subtrahends' parameter, an
//	error will be returned.
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
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractDecimalSeries(
	minuend Decimal,
	subtrahends ...Decimal) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractDecimalSeries",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractDecimalSeries(
		numSeps, true, minuend, true, true, ePrefix, subtrahends...)
}

// SubtractIntAry
//
//	Performs the subtraction operation on two instances of type
//	IntAry.
//
//	Subtraction Operation
//	=====================
//
//	In the subtraction operation:
//
//	'minuend' - 'subtrahend' = difference or result
//	iaMinuend = 'minuend'
//	iaSubtrahend = 'subtrahend'
//	iaMinuend - iaSubtrahend = difference or result
//
//	After the subtraction operation, the 'difference' or 'result' is
//	returned as a Type BigIntNum.
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
//	This method will copy the Numeric Separators configured	for
//	input parameter 'iaMinuend' to the returned instance of
//	'difference' (type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractIntAry(
	iaMinuend IntAry,
	iaSubtrahend IntAry) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractIntAry",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := iaMinuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := iaMinuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractIntAry(
		numSeps, true, &iaMinuend, true,
		&iaSubtrahend, true, ePrefix)
}

// SubtractIntAryArray
//
//	Receives one IntAry parameter which is classified as the
//	'minuend'. The second input parameter is an array of IntAry Types
//	labeled, 'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend'. The summary result of this subtraction operation is
//	then returned as a BigIntNum type.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = cumulative difference or result
//	    b1 = 'minuend'
//	    b2 = 'subtrahend'
//	    b1 - b2 = difference or result
//
//	In this method, the 'subtrahends' is an array of Decimal Types.
//	Each element in the 'subtrahends' array is subtracted from the
//	value of 'minuend'. After the subtraction operation, the
//	cumulative 'difference' or 'result' is returned as a Type
//	BigIntNum.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	                  subtrahends
//	minuend              Array                   difference
//
//	  50      -     subtrahends[0] = 2
//	  48      -     subtrahends[1] = 3
//	  45      -     subtrahends[2] = 4
//	  41      -     subtrahends[3] = 5
//	  36      -     subtrahends[4] = 6
//	  30      -     subtrahends[5] = 9
//
//	                      Final Returned 'difference' = 21
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
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractIntAryArray(
	minuend IntAry,
	subtrahends []IntAry) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractIntAryArray",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractIntAryArray(
		numSeps, true, minuend, true,
		subtrahends, true, ePrefix)
}

// SubtractIntAryOutputToArray
//
//	This method receives two input parameters. The first input
//	parameter is an IntAry Type labeled, 'minuend'.  The second input
//	parameter is an array of IntAry types labeled 'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend' value with the 'dfference', or result, output to another
//	'results' array of IntAry types which is then returned to the
//	calling function.
//
//	                subtrahends                   Output
//	minuend            Array                       Array
//
//	  10      -    subtrahends[0] = 2     =     outputarray[0] =  8
//	  10      -    subtrahends[1] = 3     =     outputarray[1] =  7
//	  10      -    subtrahends[2] = 4     =     outputarray[2] =  6
//	  10      -    subtrahends[3] = 5     =     outputarray[3] =  5
//	  10      -    subtrahends[4] = 6     =     outputarray[4] =  4
//	  10      -    subtrahends[5] = 9     =     outputarray[5] =  1
//
//	Numeric Separators
//	==================
//
//	Each array element of the []IntAry 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'result' array ([]IntAry) returned by this subtraction
//	operation will contain array elements with numeric separators
//	(decimal separator, thousands separator and currency symbol) which
//	have been copied from input parameter 'minuend'.
func (bSubtract *BigIntMathSubtract) SubtractIntAryOutputToArray(
	minuend IntAry,
	subtrahends []IntAry) ([]IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractIntAryOutputToArray",
		"")

	if err != nil {
		return []IntAry{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractIntAryOutputToArray(
		numSeps, true, minuend, true,
		subtrahends, true, ePrefix)
}

// SubtractIntArySeries
//
//	This method receives two input parameters. The first input
//	parameter is an IntAry Type labeled, 'minuend'.  The second
//	input parameter is a series of IntAry types labeled 'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each member of the 'subtrahends' series is subtracted from the
//	'minuend' to produce a cumulative result.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = cumulative result or 'difference'
//
//	In this method, the 'subtrahend' is a series of IntAry objects.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	               subtrahends
//	minuend          series               difference
//
//	  50       -        2
//	  48       -        3
//	  45       -        4
//	  41       -        5
//	  36       -        6
//	  30       -        9
//
//	              Final Returned 'difference' = 21
//
//	Variadic Function
//	=================
//
//	In this method, the 'subtrahends' parameter is as a series of
//	IntAry objects.
//
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide one value, multiple values,
//	or zero values. Note, that if the user chooses NOT to submit at
//	least one valid IntAry object for the 'subtrahends' parameter, an
//	error will be returned.
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
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractIntArySeries(
	minuend IntAry,
	subtrahends ...IntAry) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractIntArySeries",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractIntArySeries(
		numSeps, true, minuend, true, true, ePrefix, subtrahends...)
}

// SubtractINumMgr
//
//	Receives two objects which implement the INumMgr Interface and
//	subtracts their numeric values.
//
//	Subtraction Operation
//	=====================
//
//	The 'subtrahend' numeric value is subtracted from the 'minuend'
//	numeric value.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry.
//
//	After the subtraction operation, the 'difference' or 'result' is
//	returned as a Type BigIntNum.
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
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractINumMgr(
	minuend,
	subtrahend INumMgr) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractINumMgr",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractINumMgr(
		numSeps, true, minuend, true,
		subtrahend, true, ePrefix)
}

// SubtractINumMgrArray
//
//	Receives two input parameters. The first parameter is an INumMgr
//	instance which is classified as the 'minuend'. The second
//	parameter is an array of INumMgr instances which resents the
//	'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	A subtraction operation is performed on the 'minuend' and the
//	'subtrahends' array. Each numeric value of the 'subtrahends' array
//	is subtracted from the 'minuend'. The summary result of this
//	subractioin operation is returned as a BigIntNum type.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahends' = cumulative difference or result
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry. This allows the user to mix different types
//	in a single array and add their numeric values.
//
//	In this method, the 'subtrahends' is an array of INumMgr Types.
//	Each element in the 'subtrahends' array is subtracted from the
//	value of 'minuend'. After the subtraction operation, the
//	cumulative 'difference' or 'result' is returned as a Type
//	BigIntNum.
//
//	                  subtrahends
//	minuend              Array                   difference
//
//	  50      -     subtrahends[0] = 2
//	  48      -     subtrahends[1] = 3
//	  45      -     subtrahends[2] = 4
//	  41      -     subtrahends[3] = 5
//	  36      -     subtrahends[4] = 6
//	  30      -     subtrahends[5] = 9
//
//	                      Final Returned 'difference' = 21
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
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractINumMgrArray(
	minuend INumMgr,
	subtrahends []INumMgr) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractINumMgrArray",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractINumMgrArray(
		numSeps, true, minuend, true,
		subtrahends, true, ePrefix)
}

// SubtractINumMgrOutputToArray
//
//	The first input parameter to this method is an object which
//	implements the INumMgr interface ('minuend').  The second input
//	parameter is an array of INumMgr interface types labeled
//	'subtrahends'. The 'minuend' is subtracted from each element of
//	the 'subtrahends' array with the result output to another
//	'results' array of INumMgr interface types which is then returned
//	to the calling function.
//
//	Example
//	=======
//
//	                subtrahends                 Output
//	Minuend            Array                     Array
//
//	  10     -    subtrahends[0] = 2    =   outputarray[0] =  8
//	  10     -    subtrahends[1] = 3    =   outputarray[1] =  7
//	  10     -    subtrahends[2] = 4    =   outputarray[2] =  6
//	  10     -    subtrahends[3] = 5    =   outputarray[3] =  5
//	  10     -    subtrahends[4] = 6    =   outputarray[4] =  4
//	  10     -    subtrahends[5] = 9    =   outputarray[5] =  1
//
//	Note: The underlying type for the returned results array is
//	'BigIntNum', a type which implements the INumMgr Interface.
//
//	Numeric Separators
//	==================
//
//	Each array element of the []INumMgr 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'result' array ([]INumMgr) returned by this subtraction
//	operation will contain array elements with numeric separators
//	(decimal separator, thousands separator and currency symbol) which
//	have been copied from input parameter 'minuend'.
func (bSubtract *BigIntMathSubtract) SubtractINumMgrOutputToArray(
	minuend INumMgr,
	subtrahends []INumMgr) (result []INumMgr, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractINumMgrOutputToArray",
		"")

	if err != nil {
		return []INumMgr{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return []INumMgr{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractINumMgrOutputToArray(
		numSeps, true, minuend, true,
		subtrahends, true, ePrefix)
}

// SubtractINumMgrSeries
//
//	Receives two input parameters. The first parameter is an INumMgr
//	instance which is classified as the 'minuend'.
//
//	The second  parameter is a series of INumMgr instances which
//	resents the 'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	The 'subtrahends' series is subtracted from the 'minuend' and the
//	net result is returned in parameter 'difference' (type BigIntNum).
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = net difference or result
//	    b1 = 'minuend'
//	    b2 = 'subtrahend'
//	    b1 - b2 = net difference or result
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry.
//
//	In this method, the 'subtrahends' is a series of INumMgr Types.
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide a value for 'subtrahend',
//	or not. Note that if the user chooses NOT to submit any valid
//	INumMgr objects for the 'subtrahends' parameter, an error will be
//	returned.
//
//	After subtracting all 'subtrahend' values from 'minuend', the
//	resulting net 'difference' value is returned as a Type BigIntNum.
//
//	               subtrahends
//	minuend          series               difference
//
//	  50       -        2
//	  48       -        3
//	  45       -        4
//	  41       -        5
//	  36       -        6
//	  30       -        9
//
//	              Final Returned 'difference' = 21
//
//	Variadic Function
//	=================
//
//	In this method, the 'subtrahends' parameter is defined as a series
//	of Type INumMgr objects.
//
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide one value, multiple values,
//	or zero values. Note, that if the user chooses NOT to submit at
//	least one valid INumMgr object for the 'subtrahends' parameter, an
//	error will be returned.
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
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractINumMgrSeries(
	minuend INumMgr,
	subtrahends ...INumMgr) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractINumMgrSeries",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractINumMgrSeries(
		numSeps, true, minuend, true,
		true, ePrefix, subtrahends...)
}

// SubtractNumStr
//
//	Receives two number strings and proceeds to subtract subtrahend
//	'n2' from minuend 'n1'.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. The digit characters
//	which comprise a number string are formatted to facilitate
//	conversion to a corresponding numeric value.
//
//	The number string parameters ('n1' and 'n2') passed to this method
//	must consist of a string of numeric digits representing a numeric
//	value. A leading minus sign (-), or surrounding parentheses '()',
//	may be included in this number string to indicate a negative
//	numeric value.
//
//	The number string of numeric digits may also include a delimiting
//	decimal separator to identify fractional digits to the right of
//	the decimal separator. This method uses the Decimal Separator
//	extracted from input parameter 'numSeps' to parse 'numStr' and
//	identify any existing fractional digits.
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
//	This method will use the Numeric Separators provided by input
//	parameter 'numSeps' to parse the 'n1' and 'n2' numbers strings.
//	In addtition, the Numeric Separators configured for the returned
//	instance of 'difference' (type BigIntNum) will be copied from
//	input parameter 'numSeps'.
//
//	Subtraction Operation
//	=====================
//
//	In the subtraction operation:
//
//	'minuend' - 'subtrahend' = difference or result
//	n1 = 'minuend'
//	n2 = 'subtrahend'
//	n1 - n2 = difference or result
func (bSubtract *BigIntMathSubtract) SubtractNumStr(
	n1 string,
	n2 string,
	numSeps NumericSeparatorDto) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStr",
		"")

	if err != nil {
		return difference, err
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStr(
		n1, true, n2, true, numSeps, true, ePrefix)
}

// SubtractNumStrArray
//
//	Receives one number string input parameter which is classified as
//	the 'minuend'. The second input parameter is an array of number
//	strings labeled, 'subtrahends'.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. The digit characters
//	which comprise a number string are formatted to facilitate
//	conversion to a corresponding numeric value.
//
//	The number string parameters ('minuend' and 'subtrahends') passed
//	to this method must consist of stringd of numeric digits
//	representing a numeric value. A leading minus sign (-), or
//	surrounding parentheses '()', may be included in this number
//	string to indicate a negative numeric value.
//
//	The number string of numeric digits may also include a delimiting
//	decimal separator to identify fractional digits to the right of
//	the decimal separator. This method uses the Decimal Separator
//	extracted from input parameter 'numSeps' to parse 'numStr' and
//	identify any existing fractional digits.
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
//	This method will use the Numeric Separators provided by input
//	parameter 'numSeps' to parse the input parameter number strings.
//	In addition, the Numeric Separators configured for the returned
//	instance of 'difference' (type BigIntNum) will also becopied from
//	input parameter 'numSeps'.
//
//	Subtraction Operation
//	=====================
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend'. The summary result of this subtraction operation is
//	then returned as a BigIntNum type.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahends' = cumulative difference or result
//
//	In this method, 'subtrahends' is an array of number strings
//	(Type string).
//
//	After all subtraction operations have been completed, the cumulative
//	'difference' or 'result' is returned as a Type BigIntNum.
//
//	                  subtrahends
//	minuend              Array                   difference
//
//	  50      -     subtrahends[0] = 2
//	  48      -     subtrahends[1] = 3
//	  45      -     subtrahends[2] = 4
//	  41      -     subtrahends[3] = 5
//	  36      -     subtrahends[4] = 6
//	  30      -     subtrahends[5] = 9
//
//	                      Final Returned 'difference' = 21
func (bSubtract *BigIntMathSubtract) SubtractNumStrArray(
	minuend string,
	subtrahends []string,
	numSeps NumericSeparatorDto) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrArray",
		"")

	if err != nil {
		return difference, err
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrArray(
		minuend, true, subtrahends, true, numSeps, true, ePrefix)
}

// SubtractNumStrOutputToArray
//
//	The first input parameter to this method is a number string
//	labeled, 'minuend'.  The second input parameter is an array of
//	number strings labeled 'subtrahends'.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. The digit characters
//	which comprise a number string are formatted to facilitate
//	conversion to a corresponding numeric value.
//
//	The number string parameters ('minuend' and 'subtrahends') passed
//	to this method must consist of stringd of numeric digits
//	representing a numeric value. A leading minus sign (-), or
//	surrounding parentheses '()', may be included in this number
//	string to indicate a negative numeric value.
//
//	The number string of numeric digits may also include a delimiting
//	decimal separator to identify fractional digits to the right of
//	the decimal separator. This method uses the Decimal Separator
//	extracted from input parameter 'numSeps' to parse 'numStr' and
//	identify any existing fractional digits.
//
//	Subtraction Operation
//	=====================
//
//	Each element of the 'subtrahends' array is subtracted from
//	'minuend' with the result of each subtraction output to another
//	'results' array of number strings which is then returned to the
//	calling function.
//
//	                   subtrahends                     Output
//	minuend               Array                         Array
//
//	  10      -      subtrahends[0] = 2      =      outputarray[0] = 8
//	  10      -      subtrahends[1] = 3      =      outputarray[1] = 7
//	  10      -      subtrahends[2] = 4      =      outputarray[2] = 6
//	  10      -      subtrahends[3] = 5      =      outputarray[3] = 5
//	  10      -      subtrahends[4] = 6      =      outputarray[4] = 4
//	  10      -      subtrahends[5] = 9      =      outputarray[5] = 1
func (bSubtract *BigIntMathSubtract) SubtractNumStrOutputToArray(
	minuend string,
	subtrahends []string,
	numSeps NumericSeparatorDto) (results []string, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrOutputToArray",
		"")

	if err != nil {
		return results, err
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrOutputToArray(
		minuend, true, subtrahends, true, numSeps, true, ePrefix)
}

// SubtractNumStrSeries
//
//	Performs a subtraction operation on two number string input
//	parameters. The first parameter is labeled 'minuend'. The second
//	input parameter, 'subtrahends', is a series of number strings.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. The digit characters
//	which comprise a number string are formatted to facilitate
//	conversion to a corresponding numeric value.
//
//	The number string parameters ('minuend' and 'subtrahends') passed
//	to this method must consist of stringd of numeric digits
//	representing a numeric value. A leading minus sign (-), or
//	surrounding parentheses '()', may be included in this number
//	string to indicate a negative numeric value.
//
//	The number string of numeric digits may also include a delimiting
//	decimal separator to identify fractional digits to the right of
//	the decimal separator. This method uses the Decimal Separator
//	extracted from input parameter 'numSeps' to parse 'numStr' and
//	identify any existing fractional digits.
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
//	This method will use the Numeric Separators provided by input
//	parameter 'numSeps' to parse the input parameter number strings.
//	In addition, the Numeric Separators configured for the returned
//	instance of 'difference' (type BigIntNum) will also becopied from
//	input parameter 'numSeps'.
//
//	Subtraction Operation
//	=====================
//
//	Each member of the 'subtrahends' series is subtracted from the
//	'minuend' to produce a cumulative result.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = cumulative result or 'difference'
//
//	In this method, the 'subtrahend' is a series of strings.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	               subtrahends
//	minuend          series               difference
//
//	  50       -        2
//	  48       -        3
//	  45       -        4
//	  41       -        5
//	  36       -        6
//	  30       -        9
//
//	              Final Returned 'difference' = 21
//
//	Variadic Function
//	=================
//
//	In this method, the 'subtrahends' parameter is defined as a series
//	of number strings (Type string).
//
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide one value, multiple values,
//	or zero values. Note, that if the user chooses NOT to submit at
//	least one valid number string for the 'subtrahends' parameter, an
//	error will be returned.
func (bSubtract *BigIntMathSubtract) SubtractNumStrSeries(
	numSeps NumericSeparatorDto,
	minuend string,
	subtrahends ...string) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrSeries",
		"")

	if err != nil {
		return difference, err
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrSeries(
		numSeps, true, minuend, true, true, ePrefix, subtrahends...)
}

// SubtractNumStrDto
//
//	Receives two 'NumStrDto' instances and proceeds to subtract
//	'nDtoSubtrahend' from 'nDtoMinuend'.
//
//	Subtraction Operation
//	=====================
//
//	In the subtraction operation:
//
//	        'minuend' - 'subtrahend' = difference or result
//	                  'nDtoMinuend'  = minuend
//									 'nDtoSubtrahend'  = subtrahend
//	    nDtoMinuend - nDtoSubtrahend = difference or result
//
//	After the subtraction operation, the 'result' or 'difference' is
//	returned as a Type BigIntNum.
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
//	This method will copy the Numeric Separators provided by input
//	parameter 'numSeps'. The Numeric Separators configured for the
//	returned instance of 'difference' (type BigIntNum) will also be
//	copied from input parameter 'numSeps'.
func (bSubtract *BigIntMathSubtract) SubtractNumStrDto(
	nDtoMinuend NumStrDto,
	nDtoSubtrahend NumStrDto) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrDto",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := nDtoMinuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := nDtoMinuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrDto(
		numSeps, true, &nDtoMinuend, true,
		&nDtoSubtrahend, true, ePrefix)
}

// SubtractNumStrDtoArray
//
//	 Receives two NumStrDto input parameters. The first is labeled as
//	 the 'minuend'.  The second is an array of NumStrDto Types labeled,
//	 'subtrahends'.
//
//	 Subtraction Operation
//	 =====================
//
//	 In the subtraction operation, the array of 'subtrahends' is
//	 subtracted from the 'minuend'.
//
//	   'minuend' - 'subtrahends' = cumulative difference or result
//
//	 In this method, the 'subtrahends' is an array of NumStrDto Types.
//
//	 After all subtraction operations have been completed, the cumulative
//	 'difference' or 'result' is returned as a Type BigIntNum.
//
//		                  subtrahends
//		minuend              Array                   difference
//
//		  50      -     subtrahends[0] = 2
//		  48      -     subtrahends[1] = 3
//		  45      -     subtrahends[2] = 4
//		  41      -     subtrahends[3] = 5
//		  36      -     subtrahends[4] = 6
//		  30      -     subtrahends[5] = 9
//
//		                      Final Returned 'difference' = 21
//
//	 Numeric Separators
//	 ==================
//
//	 Numeric Separators define the Decimal Separator character,
//	 Thousands Separator character, and Currency Symbol character.
//	 These separator characters serve two purposes. First they are
//	 used to format and display numeric values as number strings.
//	 Second, they are also used to parse number strings and convert
//	 them into numeric values.
//
//	 The Numeric Separators configured for input parameter 'minuend'
//	 will be copied to the returned instance of 'difference'
//	 (type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractNumStrDtoArray(
	minuend NumStrDto,
	subtrahends []NumStrDto) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrDtoArray",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrDtoArray(
		numSeps, true, minuend, true,
		subtrahends, true, ePrefix)
}

// SubtractNumStrDtoOutputToArray
//
//	This method receives two input parameters. The first input
//	parameter is a NumStrDto Type labeled, 'minuend'.  The second
//	input parameter is an array of NumStrDto types labeled 'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend' value with the 'dfference', or result, output to another
//	'results' array of NumStrDto types which is then returned to the
//	calling function.
//
//	                subtrahends                   Results
//	minuend            Array                       Array
//
//	  10      -    subtrahends[0] = 2     =     outputarray[0] =  8
//	  10      -    subtrahends[1] = 3     =     outputarray[1] =  7
//	  10      -    subtrahends[2] = 4     =     outputarray[2] =  6
//	  10      -    subtrahends[3] = 5     =     outputarray[3] =  5
//	  10      -    subtrahends[4] = 6     =     outputarray[4] =  4
//	  10      -    subtrahends[5] = 9     =     outputarray[5] =  1
//
//	Numeric Separators
//	==================
//
//	Each array element of the []IntAry 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'resultsArray' ([]NumStrDto) parameter returned by this
//	subtraction operation will contain array elements with numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) which have been copied from input parameter 'minuend'.
func (bSubtract *BigIntMathSubtract) SubtractNumStrDtoOutputToArray(
	minuend NumStrDto,
	subtrahends []NumStrDto) (resultsArray []NumStrDto, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrDtoOutputToArray",
		"")

	if err != nil {
		return []NumStrDto{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrDtoOutputToArray(
		numSeps, true, minuend, true,
		subtrahends, true, ePrefix)
}

// SubtractNumStrDtoSeries
//
//	This method receives two input parameters. The first input
//	parameter is a NumStrDto Type labeled, 'minuend'.  The second
//	input parameter is a series of NumStrDto types labeled
//	'subtrahends'.
//
//	Subtraction Operation
//	=====================
//
//	Each member of the 'subtrahends' series is subtracted from the
//	'minuend' to produce a cumulative result.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = cumulative result or 'difference'
//
//	In this method, the 'subtrahend' is a series of NumStrDto objects.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	               subtrahends
//	minuend          series               difference
//
//	  50       -        2
//	  48       -        3
//	  45       -        4
//	  41       -        5
//	  36       -        6
//	  30       -        9
//
//	              Final Returned 'difference' = 21
//
//	Variadic Function
//	=================
//
//	In this method, the 'subtrahends' parameter is as a series of
//	NumStrDto Types.
//
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide one value, multiple values,
//	or zero values. Note, that if the user chooses NOT to submit at
//	least one valid NumStrDto object for the 'subtrahends' parameter, an
//	error will be returned.
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
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bSubtract *BigIntMathSubtract) SubtractNumStrDtoSeries(
	minuend NumStrDto,
	subtrahends ...NumStrDto) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractNumStrDtoSeries",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps, err := minuend.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := minuend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractMacrobot).subtractNumStrDtoSeries(
		numSeps, true, minuend, true,
		true, ePrefix, subtrahends...)
}

// SubtractPair
//
//	Performs the subtraction operation. This method receives a type
//	'BigIntPair' and proceeds to subtract bPair.Big2 from bPair.Big1.
//
//	After the subtraction operation, the 'difference' or 'result' is
//	returned as a Type BigIntNum.
//
//	The BigIntNum 'result' returned by this subtraction operation will
//	contain numeric separators (decimal separator, thousands separator
//	and currency symbol) which were copied from input parameter bPair.Big1,
//	the minuend.
func (bSubtract *BigIntMathSubtract) SubtractPair(bPair BigIntPair) (difference BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathSubtract.SubtractPair",
		"")

	if err != nil {
		return difference, err
	}

	numSeps, err := bPair.Big1.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := bPair.Big1.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(bigIntMathSubtractNanobot).subtractBigIntPair(
		numSeps, true, &bPair, true, ePrefix)
}
