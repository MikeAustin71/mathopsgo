package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

// BigIntMathSubtract
type bigIntMathSubtractNanobot struct {
	lock sync.Mutex
}

// bigIntSubtract
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
func (bIMathSubNanobot *bigIntMathSubtractNanobot) bigIntSubtract(
	minuend *big.Int,
	minuendPrecision *big.Int,
	subtrahend *big.Int,
	subtrahendPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (
	difference *big.Int, differencePrecision *big.Int, err error) {

	bIMathSubNanobot.lock.Lock()

	defer bIMathSubNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathSubtractNanobot.bigIntSubtract",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	if minuend == nil {

		return big.NewInt(0), big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'minuend'",
			}
	}

	if minuendPrecision == nil {

		return big.NewInt(0), big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'minuendPrecision'",
			}
	}

	if subtrahend == nil {

		return big.NewInt(0), big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'subtrahend'",
			}
	}

	if subtrahendPrecision == nil {

		return big.NewInt(0), big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'subPrecision'",
			}
	}

	difference = big.NewInt(0)

	differencePrecision = big.NewInt(0)

	bigZero := big.NewInt(0)

	if minuendPrecision.Cmp(bigZero) == -1 {

		return difference, differencePrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("minuendPrecision= '%v'", minuendPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'minuendPrecision' is INVALID!\n" +
					"'minuendPrecision' has a value less than zero!",
			}
	}

	if subtrahendPrecision.Cmp(bigZero) == -1 {

		return difference, differencePrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("subtrahendPrecision= '%v'", subtrahendPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'subtrahendPrecision' is INVALID!\n" +
					"'subtrahendPrecision' has a value less than zero!",
			}
	}

	base10 := big.NewInt(10)

	deltaPrecision := big.NewInt(0)

	scale := big.NewInt(0)

	if minuendPrecision.Cmp(subtrahendPrecision) == 0 {
		// Precisions are equal.

		difference = big.NewInt(0).Sub(minuend, subtrahend)

		differencePrecision = big.NewInt(0).Set(minuendPrecision)

	} else if minuendPrecision.Cmp(subtrahendPrecision) == 1 {
		//  minPrecision > subPrecision

		deltaPrecision = big.NewInt(0).Sub(minuendPrecision, subtrahendPrecision)

		scale = big.NewInt(0).Exp(base10, deltaPrecision, nil)

		newSubInt := big.NewInt(0).Mul(subtrahend, scale)

		difference = big.NewInt(0).Sub(minuend, newSubInt)

		differencePrecision.Set(minuendPrecision)

	} else {
		// subPrecision must be GREATER THAN minPrecision

		deltaPrecision = big.NewInt(0).Sub(subtrahendPrecision, minuendPrecision)

		scale = big.NewInt(0).Exp(base10, deltaPrecision, nil)

		newMinuendInt := big.NewInt(0).Mul(minuend, scale)

		difference = big.NewInt(0).Sub(newMinuendInt, subtrahend)

		differencePrecision.Set(subtrahendPrecision)
	}

	if difference.Cmp(bigZero) == 0 {

		differencePrecision = big.NewInt(0)
	}

	// Delete trailing fractional zeros
	if differencePrecision.Cmp(bigZero) == 1 {
		// differencePrecision > 0

		scrap := big.NewInt(0)

		biBase10 := big.NewInt(10)

		biBaseZero := big.NewInt(0)

		newDifference, mod10 := big.NewInt(0).QuoRem(difference, biBase10, scrap)

		bigOne := big.NewInt(1)

		for mod10.Cmp(biBaseZero) == 0 && differencePrecision.Cmp(bigZero) == 1 {

			difference.Set(newDifference)

			differencePrecision.Sub(differencePrecision, bigOne)

			newDifference, mod10 = big.NewInt(0).QuoRem(difference, biBase10, scrap)
		}
	}

	return difference, differencePrecision, nil
}

// subtractPairNoNumSeps
//
//	Performs the subtraction operation. This method receives a type
//	'BigIntPair' and proceeds to subtract bPair.Big1 from bPair.Big2.
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
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy Numeric Separators to the returned instance
//	of 'difference' (type BigIntNum).
func (bIMathSubNanobot *bigIntMathSubtractNanobot) subtractBigIntPair(
	numSeps NumericSeparatorDto,
	validateNumSeps bool,
	bPair *BigIntPair,
	validateBigIntPair bool,
	errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

	bIMathSubNanobot.lock.Lock()

	defer bIMathSubNanobot.lock.Unlock()

	difference = new(BigIntNum).New()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathSubtractNanobot.subtractBigIntPair",
		"")

	if err != nil {
		return difference, err
	}

	if bPair == nil {

		return difference,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bPair'",
			}
	}

	if validateBigIntPair {

		err = bPair.IsValid(ePrefix.XCpy("Validating 'bPair'").String())

		if err != nil {

			return difference,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: " err = bPair.IsValid(ePrefix.XCpy(\n" +
						"\"Validating 'bPair'\").String())",
					ErrContext: "Error: Input parameter 'bPair' is INVALID!\n" +
						"'bPair' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateNumSeps {

		err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = numSeps.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'numSeps'\").String())",
					ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
						"'numSeps' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPairBig1BigInt, err := bPair.GetBig1BigInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig1BigInt, err := bPair.GetBig1BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPairBig2BigInt, err := bPair.GetBig2BigInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2BigInt, err := bPair.GetBig2BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b3Difference := big.NewInt(0).Sub(bPairBig1BigInt, bPairBig2BigInt)

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	difference, err = new(BigIntNum).NewBigInt(b3Difference, bPairBig2PrecisionUint)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "difference, err = new(BigIntNum).\n" +
					"  NewBigInt(b3Difference, bPairBig2PrecisionUint)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = difference.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = difference.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = difference.IsValid(ePrefix.XCpy("Validating final result 'difference'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = difference.IsValid(ePrefix.XCpy(\n" +
					"  \"Validating final result 'difference'\").String())",
				ErrContext: "Error: Final calculated result 'difference' is INVALID!\n" +
					"'difference' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return difference, nil
}
