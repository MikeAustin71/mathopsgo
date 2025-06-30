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

	nthrtOriginalRadicandNumStr, err := nthrt.OriginalRadicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrtOriginalRadicandNumStr, err :=\n" +
					"nthrt.OriginalRadicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthrtNthRootNumStr, err := nthrt.NthRoot.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrtNthRootNumStr, err :=\n" +
					"nthrt.NthRoot.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthrt.SetupRadicand,
		nthrt.IntBundleRadicand,
		nthrt.FracBundleRadicand,
		precisionAdjustment,
		err =
		new(bigIntMathNthRootBoson).setupBundles(
			&nthrt.OriginalRadicand,
			true,
			&nthrt.NthRoot,
			true,
			ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrt.SetupRadicand, nthrt.IntBundleRadicand,\n" +
					"nthrt.FracBundleRadicand, precisionAdjustment, err =\n" +
					"  new(bigIntMathNthRootBoson).setupBundles(\n" +
					"  &nthrt.OriginalRadicand, true, &nthrt.NthRoot, true\n" +
					"  ePrefix)",
				ErrContext: fmt.Sprintf("nthrt.OriginalRadicand= '%v'\n"+
					"nthrt.NthRoot= '%v'", nthrtOriginalRadicandNumStr, nthrtNthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	nthrt.BundleAddOnPrecision =
		big.NewInt(int64(maxPrecision + uint(2)))

	nthrt.FracBundleLength = big.NewInt(0)

	nthrt.TotalBundleLength = big.NewInt(0)

	nthrtSetupRadicandNumStr, err := nthrt.SetupRadicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrtSetupRadicandNumStr, err :=\n" +
					"  nthrt.SetupRadicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthrt.TotalBundleLength, nthrt.FracBundleLength, err = new(bigIntMathNthRootBoson).
		calcBundleLength(
			&nthrt.SetupRadicand,
			true,
			&nthrt.NthRoot,
			false,
			nthrt.BundleAddOnPrecision,
			ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrt.TotalBundleLength, nthrt.FracBundleLength, err = \n" +
					"  new(bigIntMathNthRootBoson).calcBundleLength(\n" +
					"  &nthrt.SetupRadicand, true, &nthrt.NthRoot, false\n" +
					"  nthrt.BundleAddOnPrecision, ePrefix)",
				ErrContext: fmt.Sprintf("SetupRadicand= '%v'\nNthRootInt= '%v'\nAddOnPrecision='%v'",
					nthrtSetupRadicandNumStr, nthrtNthRootNumStr, nthrt.BundleAddOnPrecision),
				ErrMessage: err.Error(),
			}
	}

	nthrtSetupRadicandPrecisionUint, err := nthrt.SetupRadicand.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrtSetupRadicandPrecisionUint, err := \n" +
					"  nthrt.SetupRadicand.GetPrecisionUint()",
				ErrContext: fmt.Sprintf("nthrt.SetupRadicand= '%v'", nthrtSetupRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	radicandPrecision := big.NewInt(int64(nthrtSetupRadicandPrecisionUint))

	nthrtNthRootAbsoluteBigInt, err := nthrt.NthRoot.GetAbsoluteBigIntValue()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrtNthRootAbsoluteBigInt, err :=\n" +
					"  nthrt.NthRoot.GetAbsoluteBigIntValue()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthrt.ActualResultPrecision, err =
		new(bigIntMathNthRootQuark).calcPrecision(
			radicandPrecision,
			nthrt.FracBundleLength,
			precisionAdjustment,
			nthrt.BundleAddOnPrecision,
			nthrtNthRootAbsoluteBigInt,
			ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrt.ActualResultPrecision, err =\n" +
					" new(bigIntMathNthRootQuark).calcPrecision(\n" +
					" radicandPrecision, nthrt.FracBundleLength, precisionAdjustment,\n" +
					" nthrt.BundleAddOnPrecision, nthrtNthRootAbsoluteBigInt, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// Set constants for calculations
	nthrt.FracPrecision = big.NewInt(0)

	nthrt.BigOne = big.NewInt(1)

	nthrt.Big10 = big.NewInt(10)

	nthrt.Big10ToNthPower = big.NewInt(0).Exp(
		nthrt.Big10,
		nthrtNthRootAbsoluteBigInt,
		nil)

	nthrt.BigZero = big.NewInt(0)

	nthrt.ResultBInt = big.NewInt(0)

	nthrt.ResultBINum, err = new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthrt.ResultBINum, err = new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

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

		err = nthrtBoson.findNextRoot(nthrt, ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(bigIntMathNthRootBoson).findNextRoot(nthrt, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	nthrtOriginalRadicandSignValue, err := nthrt.OriginalRadicand.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "nthrtOriginalRadicandSignValue, err := nthrt.OriginalRadicand.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if nthrtOriginalRadicandSignValue < 1 {
		nthrt.ResultBInt = big.NewInt(0).Neg(nthrt.ResultBInt)
	}

	actualPrecision := uint(nthrt.ActualResultPrecision.Int64())

	nthrt.ResultBINum, err = new(BigIntNum).NewBigInt(nthrt.ResultBInt, actualPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "nthrt.ResultBINum, err = new(BigIntNum).\n" +
				"  NewBigInt(nthrt.ResultBInt, actualPrecision)",
			ErrContext: fmt.Sprintf("nthrt.ResultBInt= '%v'\n"+
				"actualPrecision= '%v'",
				nthrt.ResultBInt.Text(10), actualPrecision),
			ErrMessage: err.Error(),
		}
	}

	nthrtResultBINumStr, err := nthrt.ResultBINum.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "nthrtResultBINumStr, err := nthrt.ResultBINum.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = nthrt.ResultBINum.RoundToDecPlace(nthrt.RequestedPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = nthrt.ResultBINum.RoundToDecPlace(nthrt.RequestedPrecision)",
			ErrContext: fmt.Sprintf("nthrt.RequestedPrecision= '%v'\n"+
				"nthrt.ResultBINum= '%v'",
				nthrt.RequestedPrecision, nthrtResultBINumStr),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// empty
//
//	Resets all the internal member variables to their initial or zero
//	values for the current instance of BigIntMathNthRoot.
func (bIMathNthrtProton *bigIntMathNthRootProton) empty(
	nthrt *BigIntMathNthRoot) {

	bIMathNthrtProton.lock.Lock()

	defer bIMathNthrtProton.lock.Unlock()

	if nthrt == nil {
		return
	}

	nthrt.NthRoot = new(BigIntNum).New()

	nthrt.OriginalRadicand = new(BigIntNum).New()

	nthrt.IntBundleRadicand = new(BigIntNum).New()

	nthrt.FracBundleRadicand = new(BigIntNum).New()

	nthrt.BundleAddOnPrecision = big.NewInt(0)

	nthrt.FracBundleLength = big.NewInt(0)

	nthrt.TotalBundleLength = big.NewInt(0)

	nthrt.ResultBInt = big.NewInt(0)

	nthrt.ResultBINum = new(BigIntNum).New()

	nthrt.RequestedPrecision = 0

	nthrt.BigOne = big.NewInt(0)

	nthrt.Big10 = big.NewInt(0)

	nthrt.Big10ToNthPower = big.NewInt(0)

	nthrt.BigZero = big.NewInt(0)

	nthrt.YPrime = big.NewInt(0)

	nthrt.Minuend = big.NewInt(0)

	nthrt.Subtrahend = big.NewInt(0)

	nthrt.R = big.NewInt(0)

	nthrt.RPrime = big.NewInt(0)

	nthrt.BaseNum = big.NewInt(0)

	nthrt.Alpha = big.NewInt(0)

	nthrt.Beta = big.NewInt(0)
}
