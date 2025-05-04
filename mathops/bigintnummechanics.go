package mathops

import (
	"fmt"
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
func (bIngNumMech *bigIntNumMechanics) new() BigIntNum {

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
func (bIngNumMech *bigIntNumMechanics) newBigInt(
	bigI *big.Int,
	precision uint,
	callingMethodChain string) (BigIntNum, error) {

	if bIngNumMech.lock == nil {
		bIngNumMech.lock = new(sync.Mutex)
	}

	bIngNumMech.lock.Lock()

	defer bIngNumMech.lock.Unlock()

	ePrefix := "bigIntNumMechanics.newBigInt()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bigI == nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'bigI' is a nil pointer!\n",
				ePrefix)

	}

	b := new(BigIntNum)

	new(bigIntNumElectron).empty(b)

	err := new(bigIntNumNanobot).setBigInt(
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
func (bIngNumMech *bigIntNumMechanics) newInt64Exponent(
	int64Num int64, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewInt64()"

	bigI := big.NewInt(int64Num)

	b := BigIntNum{}

	new(bigIntNumElectron).empty(&b)

	err := new(bigIntNumNanobot).setBigInt(
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
				ePrefix,
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
func (bIngNumMech *bigIntNumMechanics) newZero(
	precision uint,
	callingMethodChain string) (BigIntNum, error) {

	if bIngNumMech.lock == nil {
		bIngNumMech.lock = new(sync.Mutex)
	}

	bIngNumMech.lock.Lock()

	defer bIngNumMech.lock.Unlock()

	ePrefix := "bigIntNumMechanics.newZero()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	b := BigIntNum{}

	new(bigIntNumElectron).empty(&b)

	err := new(bigIntNumNanobot).setBigInt(
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
func (bIngNumMech *bigIntNumMechanics) newBigIntNum() BigIntNum {

	b := new(BigIntNum)

	new(bigIntNumElectron).empty(b)

	return *b
}
