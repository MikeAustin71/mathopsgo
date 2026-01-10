package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathAddNanobot struct {
	lock *sync.Mutex
}

// addPairNoNumSeps
//
//	Receives a BigIntPair and proceeds to add the values
//	b1.BigIntNum to b2.BigIntNum.
//
//	The result of this addition is returned as a type BigIntNum.
//
//	The BigIntNum 'result' returned by this addition operation
//	will contain default numeric separators (decimal separator,
//	thousands separator and currency symbol).
func (bIMathAddNano *bigIntMathAddNanobot) addPairNoNumSeps(
	bPair BigIntPair,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIMathAddNano.lock == nil {
		bIMathAddNano.lock = new(sync.Mutex)
	}

	bIMathAddNano.lock.Lock()

	defer bIMathAddNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathAddNanobot.addPairNoNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = bPair.IsValid("Validating 'bPair'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.IsValid(\"Validating 'bPair'\")",
				ErrContext: "Input parameter 'bPair' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err := bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigI1, err := bPair.GetBig1BigInt()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigI1, err := bPair.GetBig1BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigI2, err := bPair.GetBig2BigInt()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigI2, err := bPair.GetBig2BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b3 := big.NewInt(0).Add(bigI1, bigI2)

	big2Precision, err := bPair.Big2.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "big2Precision, err := bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bResult, err := new(BigIntNum).NewBigInt(b3, big2Precision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bResult, err := new(BigIntNum).NewBigInt(b3, big2Precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = bResult.IsValid("Validating 'bResult'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bResult.IsValid(\"Validating 'bResult'\")",
				ErrContext: "Final result 'bResult' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	return bResult, nil
}

// bigIntAdd
//
//	Adds two fixed length floating point numbers and generates the
//	sum or total resulting from that addition. The two numbers added
//	together are configured as pairs of *big.Int integer numbers and
//	precision specifications. Each integer precision pair is used to
//	define a fixed length floating point number.
//
//	Examples:
//	=========
//
//	In the addition operation:
//
//	    b1 + b2 = total or sum
//
//	This method provides for the addition of fixed length floating
//	point values by means of integer and precision specification
//	pairs.
//
//	As an example, consider the following addition operation
//
//	    752.314 + 21.67894 = 773.99294 = total
//	     b1     +   b2     =  total
//
//	In this case 'b1', 'b2' and 'total' would be configured as
//	integer precision pairs:
//
//	       b1             = 752314
//	       b1Precision    = 3
//	       b2             = 2167894
//	       b2Precision    = 5
//
//	       total          = 77399294
//	       totalPrecision = 5
//
//	In this way, the method uses integer, precision pairs to define
//	fixed length floating point numbers.
//
//	Input Parameters
//	================
//
//	b1              *big.Int
//
//	  The first number which will be added to 'b2' to generate
//	  a total.
//
//	b1Precision     *big.Int
//
//	  Specifies the precision for input parameter 'b1'.
//	  Precision defines the number of fractional digits after
//	  the decimal place. 'b1Precision' must be equal to or
//	  greater than zero.
//
//	b2              *big.Int
//
//	  The second number which is added to 'b1' in order to
//	  generate a total.
//
//	b2Precision     *big.Int
//
//	  The 'b2' precision or the number of fractional digits
//	  after the decimal place. 'b2Precision' must be equal to
//	  or greater than zero.
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
//	total           *big.Int
//	  The sum or total of 'b1' and 'b2' input values.
//
//	totalPrecision  *big.Int
//	  The 'total' precision or the number of fractional
//	  digits after the decimal place.
//
//	err             error
//	  If input parameters 'b1Precision' or 'b2Precision'
//	  are less than zero, an error will be returned.
//
//
//	Taken together, 'total' and 'totalPrecision' can define a fixed
//	length floating point number.
func (bIMathAddNano *bigIntMathAddNanobot) bigIntAdd(
	b1 *big.Int,
	b1Precision *big.Int,
	b2 *big.Int,
	b2Precision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (
	total *big.Int, totalPrecision *big.Int, err error) {

	if bIMathAddNano.lock == nil {
		bIMathAddNano.lock = new(sync.Mutex)
	}

	bIMathAddNano.lock.Lock()

	defer bIMathAddNano.lock.Unlock()

	total = big.NewInt(0)

	totalPrecision = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathAddNanobot.bigIntAdd",
		"")

	if err != nil {
		return total, totalPrecision, err
	}

	if b1 == nil {

		return total, totalPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'b1'",
			}
	}

	if b1Precision == nil {

		return total, totalPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'b1'",
			}
	}

	if b2 == nil {

		return total, totalPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'b2'",
			}
	}

	if b2Precision == nil {

		return total, totalPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'b2Precision'",
			}
	}

	bigZero := big.NewInt(0)

	if b1Precision.Cmp(bigZero) == -1 {

		return total, totalPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("b1Precision='%v'", b1Precision.Text(10)),
				ErrMessage: "Error: Input parameter 'b1Precision' is LESS THAN ZERO!",
			}
	}

	if b2Precision.Cmp(bigZero) == -1 {

		return total, totalPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("b2Precision='%v'", b2Precision.Text(10)),
				ErrMessage: "Error: Input parameter 'b2Precision' is LESS THAN ZERO!",
			}
	}

	if b1.Cmp(bigZero) == 0 &&
		b2.Cmp(bigZero) == 0 {

		total = big.NewInt(0)

		totalPrecision = big.NewInt(0)

		return total, totalPrecision, nil
	}

	bigTen := big.NewInt(10)

	delta := big.NewInt(0)

	scale := big.NewInt(0)

	if b1Precision.Cmp(b2Precision) == 0 {

		total = big.NewInt(0).Add(b1, b2)

		totalPrecision = big.NewInt(0).Set(b1Precision)

	} else if b1Precision.Cmp(b2Precision) == 1 {
		// b1Precision > b2Precision
		delta = big.NewInt(0).Sub(b1Precision, b2Precision)

		scale = big.NewInt(0).Exp(bigTen, delta, nil)

		b2ToScale := big.NewInt(0).Mul(b2, scale)

		total = big.NewInt(0).Add(b1, b2ToScale)

		totalPrecision = big.NewInt(0).Set(b1Precision)

	} else {
		// b2Precision must be GREATER than b1Precision
		delta = big.NewInt(0).Sub(b2Precision, b1Precision)

		scale = big.NewInt(0).Exp(bigTen, delta, nil)

		b1ToScale := big.NewInt(0).Mul(b1, scale)

		total = big.NewInt(0).Add(b1ToScale, b2)

		totalPrecision = big.NewInt(0).Set(b2Precision)

	}

	if total.Cmp(bigZero) == 0 {

		totalPrecision = big.NewInt(0)
	}

	// Delete trailing fractional zeros
	if totalPrecision.Cmp(bigZero) == 1 {
		//totalPrecision > 0

		scrap := big.NewInt(0)

		biBase10 := big.NewInt(10)

		biBaseZero := big.NewInt(0)

		newTotal, mod10 := big.NewInt(0).QuoRem(total, biBase10, scrap)

		bigOne := big.NewInt(1)

		for mod10.Cmp(biBaseZero) == 0 && totalPrecision.Cmp(bigZero) == 1 {

			total.Set(newTotal)

			totalPrecision.Sub(totalPrecision, bigOne)

			newTotal, mod10 = big.NewInt(0).QuoRem(total, biBase10, scrap)
		}
	}

	return total, totalPrecision, nil
}
