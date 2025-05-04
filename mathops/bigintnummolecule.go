package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntNumMolecule struct {
	lock *sync.Mutex
}

// getActualNumberOfDigits - Returns the number of numeric digits
// in the absolute value of this BigIntNum instance. In addition,
// a boolean value is returned indicating whether the absolute value
// is zero.
//
// Examples
// ========
//
//	      123.45														5
//	1,234,567                               7
//
// -1,234,567                               7
//
//					 0															1
//	         0.00                           1
//	       012.34                           4
//	         0.1234													4
//	       - 0.1234													4
//	         0.123400												4
//	         0.0123400											4
//	 1,234,567.800												  8
//	         5                              1
func (bIntMolecule *bigIntNumMolecule) getActualNumberOfDigits(
	bNum *BigIntNum,
	callingMethodChain string) (
	numberOfDigits *big.Int, isZeroValue bool, err error) {

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	ePrefix := "Active Method: bigIntNumMolecule.getActualNumberOfDigits()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	numberOfDigits = big.NewInt(0)

	isZeroValue = false

	err = nil

	if bNum == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)

		return numberOfDigits, isZeroValue, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {

		return numberOfDigits, isZeroValue, err

	}

	numOfDigits, errx := BigIntMath{}.GetMagnitude(bNum.absBigInt)

	if errx != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			"numOfDigits, errx := BigIntMath{}.GetMagnitude(bNum.absBigInt)\n"+
			"bNum.absBigInt='%v' Error='%v' ",
			ePrefix,
			bNum.absBigInt.Text(10),
			errx.Error())

		return numberOfDigits, isZeroValue, err
	}

	numberOfDigits = big.NewInt(0).Add(numOfDigits, big.NewInt(1))

	if bNum.absBigInt.Cmp(big.NewInt(0)) == 0 {

		isZeroValue = true

	}

	return numberOfDigits, isZeroValue, err
}

// isBIntNumZero - Returns a boolean signaling whether a
// BigIntNum value is zero.
//
// NOTE:
// This method will first test the passed instance of BigIntNum
// to determine if that instance is valid, or not.
func (bIntMolecule *bigIntNumMolecule) isBIntNumZero(
	bNum *BigIntNum,
	callingMethodChain string) (bool, error) {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	ePrefix := "Active Method: bigIntNumMolecule.newOne()"

	var err error

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return false,
			fmt.Errorf("%v\n"+
				"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
				ePrefix)
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		"BigIntNum.IsZero() Vallidity Test on 'bNum'")

	if err != nil {
		return false, err
	}

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}

	return false, nil
}

// newOne - Returns a BigIntNum Type with a value equal to '1' (one).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								1
//			1								1.0
//			2								1.00
//			3								1.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bIntMolecule *bigIntNumMolecule) newOne(
	precision uint,
	callingMethodChain string) (BigIntNum, error) {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	ePrefix := "Active Method: bigIntNumMolecule.newOne()"

	var err error

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	b, err := new(bigIntNumMechanics).newZero(
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	bNumNanobot := new(bigIntNumNanobot)

	if precision == 0 {

		err = bNumNanobot.setBigInt(
			&b,
			big.NewInt(1),
			0,
			ePrefix)

		if err != nil {

			return BigIntNum{}, err

		}

		return b, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	newVal := big.NewInt(0).Mul(big.NewInt(1), scaleVal)

	err = bNumNanobot.setBigInt(
		&b,
		newVal,
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return b, nil
}

// setBigIntExponent - Sets the numeric value using an integer
// multiplied by 10 raised to the power of the 'exponent'
// parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'bigI' is of type *big.Int.
//
// Input parameter 'exponent' is of type int.
//
// If exponent is less than +1, precision is set equal to exponent and
// bigI is unchanged. Example:
//
//	   bigI				exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//
// If exponent is greater than 0, bigI is multiplied by 10 raised to the
// power of exponent and precision is set equal to exponent.
//
//	   bigI				exponent			BigIntNum Result
//		 123456		 		   3							123456.000
func (bIntMolecule *bigIntNumMolecule) setBigIntExponent(
	bNum *BigIntNum,
	bigI *big.Int,
	exponent int,
	callingMethodChain string) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	ePrefix := "bigIntNumMolecule.setBigIntExponent()"

	var err error

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if bigI == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigI' is a nil pointer!\n",
			ePrefix)

	}

	if exponent < 1 {

		precision := uint(exponent * -1)

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			bigI,
			precision,
			ePrefix)

		return err
	}

	// exponent must be greater than zero.
	// scale left exponent places and set precision to zero

	big10 := big.NewInt(10)
	scale := big.NewInt(int64(exponent))
	scaleValue := big.NewInt(0).Exp(big10, scale, nil)
	newBigI := big.NewInt(0).Mul(bigI, scaleValue)

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		newBigI,
		uint(exponent),
		ePrefix)

	return err
}

// setExpectedNumberOfDigits - Sets the number of expected digits associated with the
// Absolute Value of this 'BigIntNum.absBigInt'. The value is stored in the data
// field, 'BigIntNum.numberOfExpectedDigits'.
//
// Useful in tracking leading zeros.
func (bIntMolecule *bigIntNumMolecule) setExpectedNumberOfDigits(
	bNum *BigIntNum,
	numOfDigits *big.Int,
	callingMethodChain string) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	ePrefix := "Active Method: bigIntNumMolecule.setExpectedNumberOfDigits()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	var err error

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if numOfDigits == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfDigits' is a nil pointer!\n",
			ePrefix)

	}

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			bNum.precision,
			ePrefix)

		if err != nil {
			return err
		}

	}

	bNum.numberOfExpectedDigits = big.NewInt(0).Set(numOfDigits)

	return nil
}

// setNumStr - Initializes the BigIntNum instance
// for the numeric value of the number string input parameter.
// A number string is a string of numeric digits which may
// or may not be prefixed with a minus sign ('-'). The numeric
// string of digits may also contain a decimal separator such
// as a period ('.'). The decimal separator may be set by the
// user. See Method BigIntNum.SetDecimalSeparator(). The decimal
// separator is used to separate integer and fractional numeric
// digits within the number string.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bIntMolecule *bigIntNumMolecule) setNumStr(
	bNum *BigIntNum,
	numStr string,
	callingMethodChain string) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	ePrefix := "Active Method: bigIntNumMolecule.setNumStr()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	var err error

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			0,
			ePrefix)

		if err != nil {
			return err
		}

	}

	if len(numStr) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is an EMPTY string!\n",
			ePrefix)
	}

	baseRunes := []rune(numStr)
	lBaseRunes := len(baseRunes)

	numSeps := NumericSeparatorDto{}
	numSeps.DecimalSeparator = bNum.decimalSeparator
	numSeps.ThousandsSeparator = bNum.thousandsSeparator
	numSeps.CurrencySymbol = bNum.currencySymbol

	newSign := 1

	newPrecision := uint(0)

	newAbsBigInt := big.NewInt(0)

	baseTen := big.NewInt(10)

	isStartNumericDigits := false

	isEndNumericDigits := false

	isFractionalValue := false

	hasMinusSign := false

	hasLeftParen := false

	hasRightParen := false

	numOfNumericDigits := 0

	for i := 0; i < lBaseRunes; i++ {

		if isEndNumericDigits {
			continue
		}

		if baseRunes[i] == ',' && bNum.decimalSeparator != ',' {
			continue
		}

		if baseRunes[i] == '-' {
			hasMinusSign = true
			continue
		}

		if baseRunes[i] == '(' {

			if isStartNumericDigits == false {
				hasLeftParen = true
			}
			continue
		}

		if baseRunes[i] == ')' {

			if isStartNumericDigits == true &&
				hasLeftParen == true {
				hasRightParen = true
				isEndNumericDigits = true
			}

			continue
		}

		if baseRunes[i] == bNum.decimalSeparator {
			isFractionalValue = true
			continue
		}

		if baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			newAbsBigInt = big.NewInt(0).Mul(newAbsBigInt, baseTen)

			newAbsBigInt = big.NewInt(0).Add(newAbsBigInt,
				big.NewInt(int64(baseRunes[i]-48)))

			isStartNumericDigits = true
			numOfNumericDigits++

			if isFractionalValue {
				newPrecision++
			}

		}

	}

	if numOfNumericDigits == 0 {
		return fmt.Errorf("%v\n"+
			"Error: No numeric digits were found in input parameter 'numStr'.\n"+
			"numStr='%v'\n",
			ePrefix,
			numStr)
	}

	if hasMinusSign == true ||
		(hasLeftParen == true && hasRightParen == true) {
		newSign = -1
	}

	bNum.Empty()
	bNum.sign = newSign
	bNum.precision = newPrecision
	bNum.absBigInt = big.NewInt(0).Set(newAbsBigInt)

	if bNum.sign == 1 {
		bNum.bigInt = big.NewInt(0).Set(newAbsBigInt)
	} else {
		bNum.bigInt = big.NewInt(0).Neg(newAbsBigInt)
	}

	bNum.scaleFactor = big.NewInt(0).Exp(baseTen,
		big.NewInt(int64(newPrecision)),
		nil)

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	return nil
}
