package mathops

import (
  "fmt"
  "math/big"
  "sync"

  ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathNthRootMechanics struct {
  lock *sync.Mutex
}

// getNthRoot
//
//	Calculates the Nth Root of a real number ('radicand') passed to
//	the method as Type BigIntNum.  The calling function must supply
//	input parameters for 'radicand', 'nthRoot' and 'maxPrecision'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First, they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'radicand' to the returned instance of
//	BigIntNum.
//
//	Input Parameters
//	================
//
//	radicand                 BigIntNum
//	  The radicand value from which the nth Root will be taken.
//	            nthRootResult^nthRoot = radicand
//
//	nthRoot                  BigIntNum
//	  Specifies the root which will be calculated for parameter,
//	  'radicand'. Examples: square root, cube root, 4th root, 9th
//	  root, etc.
//
//	  'nthRoot' is a BigIntNum Type, which may be a positive or
//	  negative number. In addition, the nthRoot may be either an
//	  integer number or a fractional number.
//
//	  The nthRoot must be a numeric value greater than one ('1') or
//	  less than minus one (-1). nthRoots with a value of zero will
//	  always return an nthRoot result of zero. Nth Root values of +1
//	  or -1 will generate an error.
//
//	  If the radicand is negative and the nthRoot value is an even
//	  number (evenly divisible by 2 with no remainder), an error will
//	  be returned since the result of such a calculation is an
//	  imaginary number.
//
//	maxPrecision             uint
//	  Specifies the maximum number of decimals to the right of the
//	  decimal point to which the Nth root result will be calculated.
//
//	errPrefDto               *ePref.ErrPrefixDto
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
//		BigIntNum
//	  If the calculation is successful, the nth root result will be
//	  returned as a BigIntNum type. This returned BigIntNum nth root
//	  will contain numeric separators (decimal separator, thousands
//	  separator and currency symbol) copied from input parameter,
//	  'radicand'.
//
//	error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (bIMathNthrtMech *bigIntMathNthRootMechanics) getNthRoot(
  nthrt *BigIntMathNthRoot,
  radicand *BigIntNum,
  validateRadicand bool,
  nthRoot *BigIntNum,
  validateNthRoot bool,
  maxPrecision uint,
  errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

  if bIMathNthrtMech.lock == nil {
    bIMathNthrtMech.lock = new(sync.Mutex)
  }

  bIMathNthrtMech.lock.Lock()

  defer bIMathNthrtMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathNthRootMechanics.getNthRoot",
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

  if validateRadicand {

    err = radicand.IsValid(ePrefix.XCpy("Validating 'radicand'").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = radicand.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'radicand'\").String()",
          ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
            "'radicand' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNthRoot {

    err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

    if err != nil {

      return BigIntNum{},
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

  radicandSignValue, err := radicand.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandSignValue, err := radicand.GetSign()",
        ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  if radicandSignValue == -1 {
    // 'radicand' is a negative number

    isEvenNum, err := nthRoot.IsEvenNumber()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "isEvenNum, err := nthRoot.IsEvenNumber()",
          ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
          ErrMessage: err.Error(),
        }
    }

    if isEvenNum {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: fmt.Sprintf("radicand= '%v'\n"+
            "nthRoot= '%v'", radicandNumStr, nthRootNumStr),
          ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
            "Cannot calculate nthRoot of a negative radicand when nthRoot is even.",
        }
    }
  }

  // If the radicand is zero, the result will always be zero

  radicandIsZero, err := radicand.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandIsZero, err := radicand.IsZero()",
        ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  if radicandIsZero {

    radicandCopy, err := radicand.CopyOut()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "radicandCopy, err := radicand.CopyOut()",
          ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
          ErrMessage: err.Error(),
        }
    }

    return radicandCopy, nil
  }

  numSeps, err := radicand.GetNumericSeparatorsDto()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := radicand.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !validateRadicand {

    err = numSeps.IsValid(ePrefix.XCpy("Validating radicand numSeps").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating radicand numSeps\").String())",
          ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
            "'radicand' numeric separators FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  bigINumOne, err := new(BigIntNum).NewBigIntNumSeps(big.NewInt(1), 0, numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "bigINumOne, err := new(BigIntNum).\n" +
          "  NewBigIntNumSeps(big.NewInt(1),  0, numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nthRootIsZero, err := nthRoot.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootIsZero, err := nthRoot.IsZero()",
        ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
        ErrMessage: err.Error(),
      }
  }

  // If nthRoot is zero, the result will always be '1'
  if nthRootIsZero {

    return bigINumOne, nil
  }

  nthRootCmpBigINumOne, err := nthRoot.Cmp(bigINumOne)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootCmpBigINumOne, err := nthRoot.Cmp(bigINumOne)",
        ErrContext: fmt.Sprintf("nthRoot= '%v'\n"+
          "bigINumOne= '1'", nthRootNumStr),
        ErrMessage: err.Error(),
      }
  }

  // Error if nthRoot == 1
  if nthRootCmpBigINumOne == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
        ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
          "'nthRoot' cannot be equal to '1'",
      }
  }

  var nthRootResult BigIntNum

  nthRootSignValue, err := nthRoot.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootSignValue, err := nthRoot.GetSign()",
        ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
        ErrMessage: err.Error(),
      }
  }

  if nthRootSignValue == -1 {

    nthRootResult, err = new(bigIntMathNthRootMacrobot).
      calcNegativeNthRoot(nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "nthRootResult, err = new(bigIntMathNthRootMacrobot).\n" +
            " calcNegativeNthRoot( nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)",
          ErrContext: fmt.Sprintf("radicand= '%v'\n"+
            "nthRoot= '%v'\n"+
            "maxPrecision= '%v'",
            radicandNumStr, nthRootNumStr, maxPrecision),
          ErrMessage: err.Error(),
        }
    }

  } else {

    nthRootResult, err = new(bigIntMathNthRootMacrobot).calcPositiveNthRoot(nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "nthRootResult, err = new(bigIntMathNthRootMacrobot).\n" +
            " calcPositiveNthRoot( nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)",
          ErrContext: fmt.Sprintf("radicand= '%v'\n"+
            "nthRoot= '%v'\n"+
            "maxPrecision= '%v'",
            radicandNumStr, nthRootNumStr, maxPrecision),
          ErrMessage: err.Error(),
        }
    }
  }

  err = nthRootResult.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthRootResult.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "Error setting Numeric Separators for final calculated result",
        ErrMessage: err.Error(),
      }
  }

  err = nthRootResult.IsValid(ePrefix.XCpy("Validating final result 'nthRootResult'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthRootResult.IsValid(ePrefix)",
        ErrContext: "Error: Final calculated result 'nthRootResult' is INVALID!\n" +
          "'nthRootResult' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return nthRootResult, nil
}
