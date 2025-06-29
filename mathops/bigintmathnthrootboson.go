package mathops

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathNthRootBoson struct {
	lock sync.Mutex
}

// Experimental 2
func (bIMathNthrtBoson *bigIntMathNthRootBoson) setupBundles(
	radicand *BigIntNum,
	nthRoot *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (setupRadicand BigIntNum,
	intBundleRadicand BigIntNum,
	fracBundleRadicand BigIntNum,
	precisionAdjustment *big.Int,
	err error) {

	bIMathNthrtBoson.lock.Lock()

	defer bIMathNthrtBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootBoson.setupBundles",
		"")

	if err != nil {
		return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0), err
	}

	intBundleRadicand = new(BigIntNum).NewZero(0)

	fracBundleRadicand = new(BigIntNum).NewZero(0)

	precisionAdjustment = big.NewInt(0)

	var errx error

	modX := big.NewInt(0)
	bigTen := big.NewInt(10)
	scaleValue := big.NewInt(0)

	setupRadicand =
		new(BigIntNum).NewBigInt(
			radicand.GetAbsoluteBigIntValue(),
			radicand.GetPrecisionUint())

	setupRadicand.TrimTrailingFracZeros()
	setupRadicand.SetExpectedToActualNumberOfDigits()

	if setupRadicand.GetPrecisionUint() == 0 {
		intBundleRadicand = setupRadicand.CopyOut()

		/*
		   fmt.Println("setupRadicand: ", setupRadicand.GetNumStr())
		   fmt.Println("intBundleRadicand:", intBundleRadicand.GetNumStr())
		   fmt.Println("fracBundleRadicand:", fracBundleRadicand.GetNumStr())
		*/

		return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	setupRadicandTotalDigits, _, errx := setupRadicand.GetActualNumberOfDigits()

	if errx != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by setupRadicand.GetActualNumberOfDigits(). "+
			"Error='%v' ", errx.Error())

		return new(BigIntNum).NewZero(0), intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	// Precision must be greater than zero
	radicandPrecision := setupRadicand.GetPrecisionBigInt()

	setupRadicandIntegerDigits :=
		big.NewInt(0).Sub(setupRadicandTotalDigits, radicandPrecision)

	expectedFractionalDigits := big.NewInt(0).Set(radicandPrecision)

	scaleValue = big.NewInt(0).Exp(bigTen, radicandPrecision, nil)

	intBundleRadicandBigInt, fracBundleRadicandBigInt :=
		big.NewInt(0).QuoRem(setupRadicand.GetAbsoluteBigIntValue(), scaleValue, modX)

	intBundleRadicand = new(BigIntNum).NewBigInt(intBundleRadicandBigInt, 0)

	intBundleRadicand.SetExpectedNumberOfDigits(setupRadicandIntegerDigits)

	mod := big.NewInt(0).Rem(radicandPrecision, nthRoot.GetAbsoluteBigIntValue())

	if mod.Cmp(big.NewInt(0)) == 1 {
		delta := big.NewInt(0).Sub(nthRoot.GetAbsoluteBigIntValue(), mod)
		scaleVal := big.NewInt(0).Exp(big.NewInt(10), delta, nil)
		fracBundleRadicandBigInt = big.NewInt(0).Mul(fracBundleRadicandBigInt, scaleVal)
		expectedFractionalDigits = big.NewInt(0).Add(expectedFractionalDigits, delta)
	}

	fracBundleRadicand = new(BigIntNum).NewBigInt(fracBundleRadicandBigInt, 0)

	if intBundleRadicand.IsZero() &&
		fracBundleRadicand.IsZero() {
		err = errors.New(ePrefix + "Error: Both intBundleRadicand and fracBundleRadicand are ZERO!")
		return new(BigIntNum).NewZero(0), intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	fracBundleRadicand.SetExpectedNumberOfDigits(expectedFractionalDigits)

	err = nil

	return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
}

// calcBundleLength - Calculates the final bundle length. These are bundles
// of integers that are packaged for submission to the Nth Root calculation
// routine.
//
// The final bundle length is equal to the actual number of bundles associated
// with the radicand plus bundles added on for additional precision in the
// final nthRoot result.
func (bIMathNthrtBoson *bigIntMathNthRootBoson) calcBundleLength(
	radicand *BigIntNum,
	nthRoot *BigIntNum,
	bundleAddonPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (totalBundleLength *big.Int,
	fracBundleLength *big.Int, err error) {

	bIMathNthrtBoson.lock.Lock()

	defer bIMathNthrtBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootBoson.calcBundleLength",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	totalBundleLength = big.NewInt(0)
	fracBundleLength = big.NewInt(0)

	err = nil
	var errx error

	modX := big.NewInt(0)
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	numOfDigits := big.NewInt(1)
	intBundleLength := big.NewInt(0)
	intBundleLengthMod := big.NewInt(0)
	bigIntNthRoot := nthRoot.GetAbsoluteBigIntValue()

	scaleValue := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(radicand.GetPrecisionUint())),
		nil)

	radicandBigInt := radicand.GetAbsoluteBigIntValue()

	// break radicand into integer digits
	intRadicand :=
		big.NewInt(0).Quo(
			radicandBigInt,
			scaleValue)

	if intRadicand.Cmp(bigZero) == 1 {

		numOfDigits, errx = BigIntMath{}.GetMagnitude(intRadicand)

		if errx != nil {
			err = fmt.Errorf(ePrefix+
				"Error returned by BigIntMath{}.GetMagnitudeDigits(intRadicand). "+
				"intRadicand='%v' Error='%v'",
				intRadicand.Text(10), errx.Error())

			return totalBundleLength, fracBundleLength, err
		}

		// Convert integer radicand to number of digits
		numOfDigits = big.NewInt(0).Add(numOfDigits, bigOne)

		intBundleLength, intBundleLengthMod = big.NewInt(0).QuoRem(
			numOfDigits,
			bigIntNthRoot,
			modX)

		// For the integer bundle calculation add one to bundle length if
		// there are any remaining digits.
		// Example  5 / 3 = Quotient of 1 Mod of 2. Add 1 to bundle size
		if intBundleLengthMod.Cmp(bigZero) == 1 {
			intBundleLength = big.NewInt(0).Add(intBundleLength, bigOne)
		}

	}

	radicandPrecision := big.NewInt(int64(radicand.GetPrecisionUint()))

	if radicandPrecision.Cmp(bigZero) == 1 {
		fracBundleLength = big.NewInt(0).Quo(radicandPrecision, bigIntNthRoot)
	}

	/*
		fmt.Println("        intBundleLength: ", intBundleLength.Text(10))
		fmt.Println("       fracBundleLength: ", fracBundleLength.Text(10))
		fmt.Println( "  bundleAddonPrecision: ", bundleAddonPrecision.Text(10))
	*/

	totalBundleLength = big.NewInt(0).Add(intBundleLength, fracBundleLength)

	totalBundleLength = big.NewInt(0).Add(totalBundleLength, bundleAddonPrecision)

	err = nil

	return totalBundleLength, fracBundleLength, err
}

// findNextRoot - Called by doRootExtraction() to find the next
// root value.
func (bIMathNthrtBoson *bigIntMathNthRootBoson) findNextRoot(
	nthrt *BigIntMathNthRoot,
	errPrefDto *ePref.ErrPrefixDto) error {

	bIMathNthrtBoson.lock.Lock()

	defer bIMathNthrtBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootBoson.findNextRoot",
		"")

	if err != nil {
		return err
	}

	var bundle *big.Int

	bundle, nthrt.IntBundleRadicand, nthrt.FracBundleRadicand, err =
		nthrt.getNextBundleBigIntValue(nthrt.IntBundleRadicand,
			nthrt.FracBundleRadicand,
			nthrt.NthRoot)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by nthrt.getNextBundleBigIntValue(nthrt.IntBundleRadicand). "+
			"Error='%v' ", err.Error())
	}

	// alpha = next n-digits of radicand
	nthrt.Alpha = big.NewInt(0).Set(bundle)

	nthrt.RPrime = big.NewInt(-1)

	itatr := big.NewInt(9)
	term_1a := big.NewInt(0)
	term_1b := big.NewInt(0)

	term_2a1 := big.NewInt(0)
	term_2a2 := big.NewInt(0)
	term_2a := big.NewInt(0)

	term_2b := big.NewInt(0)
	term_2b1 := big.NewInt(0)
	term_2b2 := big.NewInt(0)

	term_1a = big.NewInt(0).Mul(nthrt.Big10ToNthPower, nthrt.R)
	term_1b = big.NewInt(0).Set(nthrt.Alpha)
	nthrt.Minuend = big.NewInt(0).Add(term_1a, term_1b)

	for itatr.Cmp(nthrt.BigZero) > -1 &&
		nthrt.RPrime.Cmp(nthrt.BigZero) == -1 {

		nthrt.Beta = big.NewInt(0).Set(itatr)
		nthrt.YPrime = big.NewInt(0).Mul(nthrt.Y, nthrt.Big10)
		nthrt.YPrime = big.NewInt(0).Add(nthrt.YPrime, nthrt.Beta)

		term_2a1 = big.NewInt(0).Mul(nthrt.BaseNum, nthrt.Y)
		term_2a2 = big.NewInt(0).Add(term_2a1, nthrt.Beta)

		term_2a = big.NewInt(0).Exp(term_2a2,
			big.NewInt(0).Set(nthrt.NthRoot.GetAbsoluteBigIntValue()),
			nil)

		term_2b1 = big.NewInt(0).Set(nthrt.Big10ToNthPower)

		term_2b2 = big.NewInt(0).Exp(nthrt.Y,
			big.NewInt(0).Set(nthrt.NthRoot.GetAbsoluteBigIntValue()),
			nil)

		term_2b = big.NewInt(0).Mul(term_2b1, term_2b2)

		nthrt.Subtrahend = big.NewInt(0).Sub(term_2a, term_2b)

		nthrt.RPrime = big.NewInt(0).Sub(nthrt.Minuend, nthrt.Subtrahend)

		itatr = big.NewInt(0).Sub(itatr, nthrt.BigOne)
	}

	nthrt.R = big.NewInt(0).Set(nthrt.RPrime)
	nthrt.Y = big.NewInt(0).Set(nthrt.YPrime)
	nthrt.ResultBInt = big.NewInt(0).Mul(nthrt.ResultBInt, nthrt.Big10)
	nthrt.ResultBInt = big.NewInt(0).Add(nthrt.ResultBInt, nthrt.Beta)

	return nil
}
