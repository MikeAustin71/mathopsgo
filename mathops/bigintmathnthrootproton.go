package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathNthRootProton struct {
	lock sync.Mutex
}

// initializeBigIntMathNthRoot
//
//	Initializes the data fields of the BigIntMathNthRoot structure and
//	validates the original number passed to the Nth Root Calculation.
//	This method assumes that radicand and nthRoot have already been
//	validated.
func (bIMathNthrtProton *bigIntMathNthRootProton) initializeBigIntMathNthRoot(
	nthrt *BigIntMathNthRoot,
	radicand *BigIntNum,
	nthRoot *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (setupRadicand BigIntNum, err error) {

	bIMathNthrtProton.lock.Lock()

	defer bIMathNthrtProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootProton.initializeBigIntMathNthRoot",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if nthrt == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthrt'",
			}
	}

	if radicand == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'radicand'",
			}
	}

	if nthRoot == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthRoot'",
			}
	}

	radicandNumStr, err := radicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "radicandNumStr, err := radicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthRootNumStr, err := nthRoot.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	precisionAdjustment := big.NewInt(0)

	nthrt.OriginalRadicand, err = radicand.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthrt.OriginalRadicand, err = radicand.CopyOut()",
				ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	nthrt.NthRoot, err = nthRoot.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthrt.NthRoot, err = nthRoot.CopyOut()",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	nthrt.RequestedPrecision = maxPrecision

	nthrt.SetupRadicand, nthrt.IntBundleRadicand, nthrt.FracBundleRadicand, precisionAdjustment, err =
		new(bigIntMathNthRootBoson).setupBundles(&nthrt.OriginalRadicand,
			&nthrt.NthRoot, ePrefix)

	if err != nil {
		setupRadicand = new(BigIntNum).NewZero(0)
		err = fmt.Errorf(ePrefix+
			"Error returned from nthrt.setupBundles(OriginalRadicand, NthRootInt). "+
			"OriginalRadicand='%v' NthRootInt='%v' Error= %v",
			nthrt.OriginalRadicand.GetNumStr(), nthrt.NthRoot.GetNumStr(), errx.Error())
		return setupRadicand, err
	}

	nthrt.BundleAddOnPrecision =
		big.NewInt(int64(maxPrecision + uint(2)))

	nthrt.FracBundleLength = big.NewInt(0)
	nthrt.TotalBundleLength = big.NewInt(0)

	nthrt.TotalBundleLength, nthrt.FracBundleLength, err = new(bigIntMathNthRootBoson).
		calcBundleLength(
			&nthrt.SetupRadicand,
			&nthrt.NthRoot,
			nthrt.BundleAddOnPrecision,
			ePrefix)

	if err != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned from nthrt.calcBundleLength(SetupRadicand, NthRootInt, AddOnPrecision). "+
			"SetupRadicand='%v' NthRootInt='%v' AddOnPrecision='%v' Error= %v",
			nthrt.SetupRadicand.GetNumStr(), nthrt.NthRoot.GetNumStr(),
			nthrt.BundleAddOnPrecision.Text(10), errx.Error())
		setupRadicand = new(BigIntNum).NewZero(0)
		return setupRadicand, err
	}

	radicandPrecision := big.NewInt(int64(nthrt.SetupRadicand.GetPrecisionUint()))

	nthrt.ActualResultPrecision, errx =
		nthrt.calcPrecision(radicandPrecision,
			nthrt.FracBundleLength,
			precisionAdjustment,
			nthrt.BundleAddOnPrecision,
			nthrt.NthRoot.GetAbsoluteBigIntValue())

	if errx != nil {
		err = fmt.Errorf(ePrefix+"Error returned from nthrt.calcPrecision(). Error= %v", errx)
		setupRadicand = new(BigIntNum).NewZero(0)
		return setupRadicand, err
	}

	// Set constants for calculations
	nthrt.FracPrecision = big.NewInt(0)
	nthrt.BigOne = big.NewInt(1)
	nthrt.Big10 = big.NewInt(10)
	nthrt.Big10ToNthPower = big.NewInt(0).Exp(
		nthrt.Big10,
		nthrt.NthRoot.GetAbsoluteBigIntValue(),
		nil)
	nthrt.BigZero = big.NewInt(0)
	nthrt.ResultBInt = big.NewInt(0)
	nthrt.ResultBINum = new(BigIntNum).NewZero(0)
	nthrt.Y = big.NewInt(0)
	nthrt.YPrime = big.NewInt(0)
	nthrt.R = big.NewInt(0)
	nthrt.RPrime = big.NewInt(0)
	nthrt.BaseNum = big.NewInt(10)
	nthrt.Alpha = big.NewInt(0)
	nthrt.Beta = big.NewInt(0)

	return nthrt.SetupRadicand, nil
}

// doRootExtraction - Manages the root extraction process
func (bIMathNthrtProton *bigIntMathNthRootProton) doRootExtraction(
	nthrt *BigIntMathNthRoot,
	errPrefDto *ePref.ErrPrefixDto) error {

	bIMathNthrtProton.lock.Lock()

	defer bIMathNthrtProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootProton.doRootExtraction",
		"")

	if err != nil {
		return err
	}

	nthrt.Y = big.NewInt(0)
	nthrt.Minuend = big.NewInt(0)
	nthrt.Subtrahend = big.NewInt(0)
	nthrt.R = big.NewInt(0)
	bigOne := big.NewInt(1)

	nthrtBoson := new(bigIntMathNthRootBoson)

	for i := big.NewInt(0); i.Cmp(nthrt.TotalBundleLength) == -1; i = big.NewInt(0).Add(i, bigOne) {

		err := nthrtBoson.findNextRoot(nthrt, ePrefix)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned by nthrt.findNextRoot(i). "+
				"Error='%v' ", err.Error())
		}

	}

	if nthrt.OriginalRadicand.GetSign() < 1 {
		nthrt.ResultBInt = big.NewInt(0).Neg(nthrt.ResultBInt)
	}

	actualPrecision := uint(nthrt.ActualResultPrecision.Int64())

	nthrt.ResultBINum = new(BigIntNum).NewBigInt(nthrt.ResultBInt, actualPrecision)

	nthrt.ResultBINum.RoundToDecPlace(nthrt.RequestedPrecision)

	return nil
}
