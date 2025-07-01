package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type bigIntMathSubtractMacrobot struct {
	lock sync.Mutex
}

// fixedDecimalSubtract
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
//	for input parameter 'numSeps' to the returned instance of
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
//	errPrefDto					*ePref.ErrPrefixDto
//
//	  This object encapsulates an error prefix string
//	  which is included in all returned error
//	  messages. Usually, it contains the name of the
//	  calling method or methods listed as a function
//	  chain.
//
//	  If no error prefix information is needed, set
//	  this parameter to 'nil'.
//
//	  Type ErrPrefixDto is included in the 'errpref'
//	  software package:
//	    "github.com/MikeAustin71/errpref".
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
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) fixedDecimalSubtract(
	numSeps NumericSeparatorDto,
	validateNumSeps bool,
	minuend BigIntFixedDecimal,
	validateMinuend bool,
	subtrahend BigIntFixedDecimal,
	validateSubtrahend bool,
	errPrefDto *ePref.ErrPrefixDto) (difference BigIntFixedDecimal, err error) {

	bIMathSubMacrobot.lock.Lock()

	defer bIMathSubMacrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathSubtractMacrobot.fixedDecimalSubtract",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	difference = new(BigIntFixedDecimal).NewZero(0)

	if validateMinuend {

		err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

		if err != nil {

			return difference,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
						"  \"Validating 'minuend'\").String())",
					ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
						"'minuend' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateSubtrahend {

		err = subtrahend.IsValid(ePrefix.XCpy("Validating 'subtrahend'").String())

		if err != nil {

			return difference,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = subtrahend.IsValid(ePrefix.XCpy(\n" +
						"  \"Validating 'subtrahend'\").String())",
					ErrContext: "Error: Input parameter 'subtrahend' is invalid.\n" +
						"'subtrahend' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	minuendBigInt, err := minuend.GetIntegerValue()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "minuendBigInt, err := minuend.GetIntegerValue()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	minuendPrecisionBigInt, err := minuend.GetPrecisionBigInt()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "minuendPrecisionBigInt, err := minuend.GetPrecisionBigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	subtrahendBigInt, err := subtrahend.GetIntegerValue()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "subtrahendBigInt, err := subtrahend.GetIntegerValue()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	subtrahendPrecisionBigInt, err := subtrahend.GetPrecisionBigInt()

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "subtrahendPrecisionBigInt, err := subtrahend.GetPrecisionBigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if validateNumSeps {

		err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

		if err != nil {

			return difference,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
						"  \"Validating 'numSeps'\").String())",
					ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
						"'numSeps' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	// The error is ignored because minuend precision and
	// subtrahend precision will never be less than zero
	difBigInt, difBigIntPrecision, err :=
		new(bigIntMathSubtractNanobot).bigIntSubtract(
			minuendBigInt,
			minuendPrecisionBigInt,
			subtrahendBigInt,
			subtrahendPrecisionBigInt,
			ePrefix)

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "difBigInt, difBigIntPrecision, err := \n" +
					"  new(bigIntMathSubtractNanobot).bigIntSubtract(\n" +
					"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt,\n" +
					"  subtrahendPrecisionBigInt, ePrefix)",
				ErrContext: fmt.Sprintf("minuendBigInt= '%v\n"+
					"minuendPrecisionBigInt= '%v\n"+
					"subtrahendBigInt= '%v\n"+
					"subtrahendPrecisionBigInt= '%v\n",
					minuendBigInt.Text(10), minuendPrecisionBigInt.Text(10),
					subtrahendBigInt.Text(10),
					subtrahendPrecisionBigInt.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	err = difference.SetNumericValue(difBigInt, uint(difBigIntPrecision.Uint64()))

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = difference.SetNumericValue(\n" +
					"  difBigInt, uint(difBigIntPrecision.Uint64()))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = difference.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = difference.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = difference.IsValid(ePrefix.XCpy("Validating final result 'difference'").String())

	if err != nil {

		return difference,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = difference.IsValid(ePrefix)",
				ErrContext: "Error: Final calculated result 'difference' is invalid.\n" +
					"'difference' FAILED final validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return difference, nil
}
