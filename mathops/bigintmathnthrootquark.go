package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathNthRootQuark struct {
	lock sync.Mutex
}

// getNextBundleBigIntValue - Calculate and return the next bundle of
// numeric digits for nth Root extraction calculations.
func (bIMathNthrtQuark *bigIntMathNthRootQuark) getNextBundleBigIntValue(
	nthrt *BigIntMathNthRoot,
	intBundleRadicand *BigIntNum,
	validateIintBundleRadicand bool,
	fracBundleRadicand *BigIntNum,
	validateFracBundleRadicand bool,
	nthRoot *BigIntNum,
	validateNthRoot bool,
	errPrefDto *ePref.ErrPrefixDto) (
	nextBundleValue *big.Int,
	newIntBundleRadicand BigIntNum,
	newFracBundleRadicand BigIntNum,
	err error) {

	bIMathNthrtQuark.lock.Lock()

	defer bIMathNthrtQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootQuark.getNextBundleBigIntValue",
		"")

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{}, err
	}

	if nthrt == nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthrt'",
			}
	}

	if intBundleRadicand == nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'intBundleRadicand'",
			}
	}

	if fracBundleRadicand == nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'fracBundleRadicand'",
			}
	}

	if nthRoot == nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthRoot'",
			}
	}

	if validateIintBundleRadicand {

		err = intBundleRadicand.IsValid(ePrefix.XCpy("Validating 'intBundleRadicand'").String())

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = intBundleRadicand.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'intBundleRadicand'\").String()",
					ErrContext: "Error: Input parameter 'intBundleRadicand' is INVALID!\n" +
						"'intBundleRadicand' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateFracBundleRadicand {

		err = fracBundleRadicand.IsValid(ePrefix.XCpy("Validating 'fracBundleRadicand'").String())

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = fracBundleRadicand.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'fracBundleRadicand'\").String()",
					ErrContext: "Error: Input parameter 'fracBundleRadicand' is INVALID!\n" +
						"'fracBundleRadicand' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateNthRoot {

		err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = nthRoot.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'nthRoot'\").String()",
					ErrContext: "Error: Input parameter 'nthRoot' is INVALID!\n" +
						"'nthRoot' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	intBundleRadicandNumStr, err := intBundleRadicand.GetNumStr()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "intBundleRadicandNumStr, err := intBundleRadicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fracBundleRadicandNumStr, err := fracBundleRadicand.GetNumStr()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracBundleRadicandNumStr, err := fracBundleRadicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthRootNumStr, err := nthRoot.GetNumStr()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nextBundleValue = big.NewInt(0)

	newIntBundleRadicand = new(BigIntNum).New()

	err = newIntBundleRadicand.SetExpectedToActualNumberOfDigits()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = newIntBundleRadicand.SetExpectedToActualNumberOfDigits()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newFracBundleRadicand = new(BigIntNum).New()

	err = newFracBundleRadicand.SetExpectedToActualNumberOfDigits()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = newFracBundleRadicand.SetExpectedToActualNumberOfDigits()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	modX := big.NewInt(0)

	var magnitude *big.Int

	var numberOfDigits *big.Int

	var tempRadicand *big.Int

	var digitsQuotient *big.Int

	var exponent *big.Int

	var divisor *big.Int

	intExpectedNumOfDigits, err := intBundleRadicand.GetExpectedNumberOfDigits()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "intExpectedNumOfDigits, err := \n" +
					"  intBundleRadicand.GetExpectedNumberOfDigits()",
				ErrContext: fmt.Sprintf("intBundleRadicand= '%v'", intBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	intActualNumOfDigits, _, err := intBundleRadicand.GetActualNumberOfDigits()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "intActualNumOfDigits, _, err := intBundleRadicand.GetActualNumberOfDigits()",
				ErrContext: fmt.Sprintf("intBundleRadicand= '%v'", intBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	intExpectedActualDelta := big.NewInt(0).Sub(intExpectedNumOfDigits, intActualNumOfDigits)

	nthRootAbsoluteBigInt, err := nthRoot.GetAbsoluteBigIntValue()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthRootAbsoluteBigInt, err := \n" +
					"  nthRoot.GetAbsoluteBigIntValue()",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	intModExpectedNthRoot := big.NewInt(0).Rem(intExpectedNumOfDigits, nthRootAbsoluteBigInt)

	intBundleRadicandIsZeroValue, err := intBundleRadicand.IsZero()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "intBundleRadicandIsZeroValue, err := intBundleRadicand.IsZero()",
				ErrContext: fmt.Sprintf("intBundleRadicand= '%v'", intBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	intBundleRadicandAbsoluteBigInt, err := intBundleRadicand.GetAbsoluteBigIntValue()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " intBundleRadicandAbsoluteBigInt, err :=\n" +
					"  intBundleRadicand.GetAbsoluteBigIntValue()",
				ErrContext: fmt.Sprintf("intBundleRadicand= '%v'", intBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	nthrtNthRootAbsoluteBigInt, err := nthrt.NthRoot.GetAbsoluteBigIntValue()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthrtNthRootAbsoluteBigInt, err := \n" +
					"  nthrt.NthRoot.GetAbsoluteBigIntValue()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigZero := big.NewInt(0)

	fracBundleRadicandIsZeroValue, err := fracBundleRadicand.IsZero()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " fracBundleRadicandIsZeroValue, err := \n" +
					"  fracBundleRadicand.IsZero()",
				ErrContext: fmt.Sprintf("fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	fracBundleRadicandAbsoluteBigInt, err := fracBundleRadicand.GetAbsoluteBigIntValue()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracBundleRadicandAbsoluteBigInt, err :=\n" +
					"  fracBundleRadicand.GetAbsoluteBigIntValue()",
				ErrContext: fmt.Sprintf("fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	fracExpectedNumOfDigits, err := fracBundleRadicand.GetExpectedNumberOfDigits()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracExpectedNumOfDigits, err := \n" +
					"  fracBundleRadicand.GetExpectedNumberOfDigits()",
				ErrContext: fmt.Sprintf("fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	fracActualNumOfDigits, _, err := fracBundleRadicand.GetActualNumberOfDigits()

	if err != nil {

		return big.NewInt(0), BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracActualNumOfDigits, _, err := \n" +
					"  fracBundleRadicand.GetActualNumberOfDigits()",
				ErrContext: fmt.Sprintf("fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	fracExpectedActualDelta := big.NewInt(0).Sub(fracExpectedNumOfDigits, fracActualNumOfDigits)

	if (intExpectedActualDelta.Cmp(bigZero) > 0 &&
		intExpectedActualDelta.Cmp(nthRootAbsoluteBigInt) >= 0) ||
		(intExpectedActualDelta.Cmp(bigZero) > 0 &&
			intBundleRadicandIsZeroValue) {

		// fmt.Println("Calc nthRoot IntegerDeltaDigits >=0")

		nextBundleValue = big.NewInt(0)

		newIntBundleRadicand, err = intBundleRadicand.CopyOut()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "newIntBundleRadicand, err = \n" +
						"  intBundleRadicand.CopyOut()",
					ErrContext: fmt.Sprintf("Calc nthRoot IntegerDeltaDigits >=0\n"+
						"intBundleRadicand= '%v'", intBundleRadicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		newIntBundleRadicandNumStr, err := newIntBundleRadicand.GetNumStr()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "newIntBundleRadicandNumStr, err := \n" +
						"  newIntBundleRadicand.GetNumStr()",
					ErrContext: "Calc nthRoot IntegerDeltaDigits >=0",
					ErrMessage: err.Error(),
				}
		}

		newFracBundleRadicand, err = fracBundleRadicand.CopyOut()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: " newFracBundleRadicand, err = \n" +
						"  fracBundleRadicand.CopyOut()",
					ErrContext: fmt.Sprintf("Calc nthRoot IntegerDeltaDigits >=0\n"+
						"fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		intExpectedNumOfDigits = big.NewInt(0).Sub(intExpectedNumOfDigits, nthRootAbsoluteBigInt)

		if intExpectedNumOfDigits.Cmp(intActualNumOfDigits) == -1 {

			intExpectedNumOfDigits = big.NewInt(0).Set(intActualNumOfDigits)

		}

		err = newIntBundleRadicand.SetExpectedNumberOfDigits(intExpectedNumOfDigits)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = newIntBundleRadicand.\n" +
						"  SetExpectedNumberOfDigits(intExpectedNumOfDigits)",
					ErrContext: fmt.Sprintf("Calc nthRoot IntegerDeltaDigits >=0\n"+
						"newIntBundleRadicand= '%v'", newIntBundleRadicandNumStr),
					ErrMessage: err.Error(),
				}
		}

	} else if !intBundleRadicandIsZeroValue {

		// fmt.Println("Calc inBundleRadicand is NOT zero")

		magnitude, err = new(BigIntMath).GetMagnitude(intBundleRadicandAbsoluteBigInt)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "magnitude, err = new(BigIntMath).\n" +
						" GetMagnitude(intBundleRadicandAbsoluteBigInt)",
					ErrContext: fmt.Sprintf("Calc inBundleRadicand is NOT zero\n"+
						"intBundleRadicand= '%v'\n"+
						"intBundleRadicandAbsoluteBigInt= '%v'",
						intBundleRadicandNumStr, intBundleRadicandAbsoluteBigInt.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		intActualNumOfDigits = big.NewInt(0).Add(magnitude, big.NewInt(1))

		digitsQuotient = big.NewInt(0).Quo(magnitude, nthrtNthRootAbsoluteBigInt)

		exponent = big.NewInt(0).Mul(digitsQuotient, nthrtNthRootAbsoluteBigInt)

		divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

		nextBundleValue, tempRadicand =
			big.NewInt(0).QuoRem(intBundleRadicandAbsoluteBigInt, divisor, modX)

		numberOfDigits, err = new(BigIntMath).GetMagnitude(nextBundleValue)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "numberOfDigits, err = new(BigIntMath).GetMagnitude(nextBundleValue)",
					ErrContext: fmt.Sprintf("Calc inBundleRadicand is NOT zero\n"+
						"nextBundleValue = '%v'", nextBundleValue.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		numberOfDigits = big.NewInt(0).Add(numberOfDigits, big.NewInt(1))

		if intModExpectedNthRoot.Cmp(bigZero) > 0 {

			intExpectedNumOfDigits = big.NewInt(0).Sub(intExpectedNumOfDigits, numberOfDigits)

		} else {

			intExpectedNumOfDigits = big.NewInt(0).Sub(intExpectedNumOfDigits, nthrtNthRootAbsoluteBigInt)

		}

		newIntBundleRadicand, err = new(BigIntNum).NewBigInt(tempRadicand, 0)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "newIntBundleRadicand, err = \n" +
						"  new(BigIntNum).NewBigInt(tempRadicand, 0)",
					ErrContext: fmt.Sprintf("Calc inBundleRadicand is NOT zero\n"+
						"tempRadicand = '%v'", tempRadicand.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		err = newIntBundleRadicand.SetExpectedNumberOfDigits(intExpectedNumOfDigits)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = newIntBundleRadicand.\n" +
						"  SetExpectedNumberOfDigits(intExpectedNumOfDigits)",
					ErrContext: fmt.Sprintf("Calc inBundleRadicand is NOT zero\n"+
						"intExpectedNumOfDigits = '%v'", intExpectedNumOfDigits.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		newFracBundleRadicand, err = fracBundleRadicand.CopyOut()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "newFracBundleRadicand, err = fracBundleRadicand.CopyOut()",
					ErrContext: fmt.Sprintf("Calc inBundleRadicand is NOT zero\n"+
						"fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
					ErrMessage: err.Error(),
				}
		}

	} else if (fracExpectedActualDelta.Cmp(bigZero) > 0 &&
		fracExpectedActualDelta.Cmp(nthRootAbsoluteBigInt) >= 0) ||
		(fracExpectedActualDelta.Cmp(bigZero) > 0 && fracBundleRadicandIsZeroValue) {

		// fmt.Println("Calc nthRoot FracDeltaDigits >= 0")

		nextBundleValue = big.NewInt(0)

		newIntBundleRadicand, err = intBundleRadicand.CopyOut()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "newIntBundleRadicand, err = intBundleRadicand.CopyOut()",
					ErrContext: fmt.Sprintf("Calc nthRoot FracDeltaDigits >= 0\n"+
						"intBundleRadicand= '%v'", intBundleRadicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		newFracBundleRadicand, err = fracBundleRadicand.CopyOut()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "newFracBundleRadicand, err = fracBundleRadicand.CopyOut()",
					ErrContext: fmt.Sprintf("Calc nthRoot FracDeltaDigits >= 0\n"+
						"fracBundleRadicand= '%v'", fracBundleRadicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		fracExpectedNumOfDigits = big.NewInt(0).Sub(fracExpectedNumOfDigits, nthRootAbsoluteBigInt)

		if fracExpectedNumOfDigits.Cmp(fracActualNumOfDigits) == -1 {

			fracExpectedNumOfDigits = big.NewInt(0).Set(fracActualNumOfDigits)

		}

		err = newFracBundleRadicand.SetExpectedNumberOfDigits(fracExpectedNumOfDigits)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = newFracBundleRadicand.\n" +
						"  SetExpectedNumberOfDigits(fracExpectedNumOfDigits)",
					ErrContext: fmt.Sprintf("Calc nthRoot FracDeltaDigits >= 0\n"+
						"fracExpectedNumOfDigits = '%v'", fracExpectedNumOfDigits.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		nthrt.FracPrecision = big.NewInt(0).Add(nthrt.FracPrecision, big.NewInt(1))
		err = nil

	} else if !fracBundleRadicandIsZeroValue {

		// fmt.Println("Calc fracBundleRadicand Is NOT zero")

		magnitude, err = new(BigIntMath).GetMagnitude(fracBundleRadicandAbsoluteBigInt)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "magnitude, err = new(BigIntMath).GetMagnitude(fracBundleRadicandAbsoluteBigInt)",
					ErrContext: fmt.Sprintf("Calc fracBundleRadicand Is NOT zero\n"+
						"fracBundleRadicandAbsoluteBigInt= '%v'",
						fracBundleRadicandAbsoluteBigInt.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		//digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
		digitsQuotient = big.NewInt(0).Quo(magnitude, nthrtNthRootAbsoluteBigInt)

		exponent = big.NewInt(0).Mul(digitsQuotient, nthrtNthRootAbsoluteBigInt)

		divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

		nextBundleValue, tempRadicand =
			big.NewInt(0).QuoRem(fracBundleRadicandAbsoluteBigInt, divisor, modX)

		newIntBundleRadicand, err = new(BigIntNum).NewZero(0)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "Calc fracBundleRadicand Is NOT zero\n" +
						"newIntBundleRadicand, err = new(BigIntNum).NewZero(0)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = newIntBundleRadicand.SetExpectedToActualNumberOfDigits()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = newIntBundleRadicand.SetExpectedToActualNumberOfDigits()",
					ErrContext: "Calc fracBundleRadicand Is NOT zero\n",
					ErrMessage: err.Error(),
				}
		}

		fracExpectedNumOfDigits = big.NewInt(0).Sub(fracExpectedNumOfDigits, nthrtNthRootAbsoluteBigInt)

		newFracBundleRadicand, err = new(BigIntNum).NewBigInt(tempRadicand, 0)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "newFracBundleRadicand, err = new(BigIntNum).NewBigInt(tempRadicand, 0)",
					ErrContext: fmt.Sprintf("Calc fracBundleRadicand Is NOT zero\n"+
						"tempRadicand= '%v'", tempRadicand.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		err = newFracBundleRadicand.SetExpectedNumberOfDigits(fracExpectedNumOfDigits)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = newFracBundleRadicand.SetExpectedNumberOfDigits(fracExpectedNumOfDigits)",
					ErrContext: fmt.Sprintf("Calc fracBundleRadicand Is NOT zero\n"+
						"fracExpectedNumOfDigits='%v'",
						fracExpectedNumOfDigits.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		nthrt.FracPrecision = big.NewInt(0).Add(nthrt.FracPrecision, big.NewInt(1))

	} else {

		// fmt.Println("Calc else zero")

		nextBundleValue = big.NewInt(0)

		newIntBundleRadicand, err = new(BigIntNum).NewZero(0)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "newIntBundleRadicand, err = new(BigIntNum).NewZero(0)",
					ErrContext: "Calc else zero",
					ErrMessage: err.Error(),
				}
		}

		err = newIntBundleRadicand.SetExpectedToActualNumberOfDigits()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = newIntBundleRadicand.\n" +
						"  SetExpectedToActualNumberOfDigits()",
					ErrContext: "Calc else zero",
					ErrMessage: err.Error(),
				}
		}

		newFracBundleRadicand, err = new(BigIntNum).NewZero(0)

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "newFracBundleRadicand, err = new(BigIntNum).NewZero(0)",
					ErrContext: "Calc else zero",
					ErrMessage: err.Error(),
				}
		}

		err = newFracBundleRadicand.SetExpectedToActualNumberOfDigits()

		if err != nil {

			return big.NewInt(0), BigIntNum{}, BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = newFracBundleRadicand.SetExpectedToActualNumberOfDigits()",
					ErrContext: "Calc else zero",
					ErrMessage: err.Error(),
				}
		}

		nthrt.FracPrecision = big.NewInt(0).Add(nthrt.FracPrecision, big.NewInt(1))
	}

	return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, nil
}
