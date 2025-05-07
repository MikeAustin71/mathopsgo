package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumMechanics struct {
	lock *sync.Mutex
}

// new - returns a new BigIntNum instance initialized to zero.
//
// The BigIntNum instance returned by this method will contain USA
// default numeric separators (decimal separator, thousands separator
// and currency symbol).
func (bIntNumMech *bigIntNumMechanics) new() BigIntNum {

	b := new(BigIntNum)

	new(bigIntNumElectron).empty(b)

	return *b
}

// newBigInt - Creates a new BigIntNum instance using a *big.Int type and its
// associated precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
// # Input Parameters
//
// bigI *big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without decimal
//	digits.
//
// precision int
//
//	This unsigned integer (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places.
//
//	Example:
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bIntNumMech *bigIntNumMechanics) newBigInt(
	bigI *big.Int,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumMech.lock == nil {
		bIntNumMech.lock = new(sync.Mutex)
	}

	bIntNumMech.lock.Lock()

	defer bIntNumMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMechanics.newBigInt()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bigI == nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'bigI' is a nil pointer!\n",
				ePrefix.String())

	}

	b := new(BigIntNum)

	new(bigIntNumElectron).empty(b)

	err = new(bigIntNumNanobot).setBigInt(
		b,
		bigI,
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	return *b, nil
}

// NewInt64Exponent -This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewInt64Exponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewInt64Exponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 int64Num		 exponent			  BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (bIntNumMech *bigIntNumMechanics) newInt64Exponent(
	int64Num int64,
	exponent int,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumMech.lock == nil {
		bIntNumMech.lock = new(sync.Mutex)
	}

	bIntNumMech.lock.Lock()

	defer bIntNumMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMechanics.newInt64Exponent()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bigI := big.NewInt(int64Num)

	b := BigIntNum{}

	new(bigIntNumElectron).empty(&b)

	err = new(bigIntNumNanobot).setBigInt(
		&b,
		big.NewInt(0),
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	//err = b.SetBigIntExponent(bigI, exponent)

	err = new(bigIntNumMolecule).
		setBigIntExponent(&b, bigI, exponent, ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err = new(bigIntNumMolecule).\n"+
				"   setBigIntExponent(&b, bigI, exponent, ePrefix)\n"+
				"Error= %v\n",
				ePrefix.String(),
				err.Error())
	}

	return b, nil
}

// newZero - Returns a BigIntNum instance with a value equal to zero.
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '0', set 'precision' equal
// to zero (0).
//
// 'precision'
//
//	  value 					Result
//			0								0
//			2								0.00
//			3								0.000
func (bIntNumMech *bigIntNumMechanics) newZero(
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumMech.lock == nil {
		bIntNumMech.lock = new(sync.Mutex)
	}

	bIntNumMech.lock.Lock()

	defer bIntNumMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMechanics.newZero()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	b := BigIntNum{}

	new(bigIntNumElectron).empty(&b)

	err = new(bigIntNumNanobot).setBigInt(
		&b,
		big.NewInt(0),
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	return b, nil
}

// newBigIntNum - returns a new BigIntNum instance initialized to zero.
//
// The BigIntNum instance returned by this method will contain USA
// default numeric separators (decimal separator, thousands separator
// and currency symbol).
func (bIntNumMech *bigIntNumMechanics) newBigIntNum() BigIntNum {

	b := new(BigIntNum)

	new(bigIntNumElectron).empty(b)

	return *b
}
