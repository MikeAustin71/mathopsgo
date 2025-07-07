package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type bigIntMathSubtractMacrobot struct {
  lock sync.Mutex
}

// fixedDecimalSubtract
//
//	Performs the subtraction operation on two BigIntFixedDecimal
//	types. The subtraction result is also returned as a
//	BigIntFixedDecimal type.
//
//	Examples
//	=========
//
//	In the subtraction operation:
//
//	'minuend' - 'subtrahend' = 'difference' or result
//
//	752.314   -   21.67894   = 730.63506 = difference
//
//	For this method 'minuend', 'subtrahend' and 'difference' are
//	configured as BigIntFixedDecimal types.
//
//	The BigIntFixedDecimal is used to defined fixed length floating
//	point numbers and is defined as follows:
//
//	type BigIntFixedDecimal struct {
//
//	  integerNum  *big.Int
//	    All numeric digits, both integer and fractional, necessary to
//	    define a fixed length floating point number.
//
//	    The number of digits to the right of the decimal place is
//	    specified by the data field, 'BigIntFixedDecimal.precision'.
//
//	  precision   uint
//	    Specifies the number of digits to the right of the decimal
//	    place in the series of numeric digits represented by data
//	    field BigIntFixedDecimal.integerNum.
//
//	}
//
//	To represent the floating point number 52.459	a BigIntDecimal
//	Structure  would be configured as follows:
//
//	    BigIntFixedDecimal.integerNum = 52459
//	    BigIntFixedDecimal.precision  = 3
//
//	As an example consider the following subtraction operation:
//
//	    752.314 - 21.67894 = 730.63506 = difference
//
//	In this case the 'minuend', 'subtrahend' and 'difference' consist
//	of BigIntFixedDecimal types configured as follows:
//
//	           minuend.integerNum     = 752314
//	           minuend.precision      = 3
//	           subtrahend.integerNum  = 2167894
//	           subtrahend.precision   = 5
//
//	           difference.integerNum  = 73063506
//	           difference.precision   = 5
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
//	This method will copy the Numeric Separators configured
//	for input parameter 'numSeps' to the returned instance of
//	'difference' (type BigIntFixedDecimal).
//
//	Input Parameters
//	================
//
//	minuend                  BigIntFixedDecimal
//	  The number from which the subtrahend will be subtracted.
//
//	subtrahend               BigIntFixedDecimal
//	  The number to be subtracted from the 'minuend'.
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
//	difference               BigIntFixedDecimal
//	  The difference or result of the subtraction operation returned
//	  as a BigIntFixedDecimal type.
//	        'minuend' - 'subtrahend' = 'difference'
//
//	err                      error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) fixedDecimalSubtract(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend BigIntFixedDecimal,
  validateMinuend bool,
  subtrahend BigIntFixedDecimal,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntFixedDecimal, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.fixedDecimalSubtract",
    "")

  if err != nil {
    return BigIntFixedDecimal{}, err
  }

  difference = new(BigIntFixedDecimal).NewZero(0)

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateSubtrahend {

    err = subtrahend.IsValid(ePrefix.XCpy("Validating 'subtrahend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = subtrahend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'subtrahend'\").String())",
          ErrContext: "Error: Input parameter 'subtrahend' is invalid.\n" +
            "'subtrahend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendBigInt, err := minuend.GetIntegerValue()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendBigInt, err := minuend.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  minuendPrecisionBigInt, err := minuend.GetPrecisionBigInt()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendPrecisionBigInt, err := minuend.GetPrecisionBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  subtrahendBigInt, err := subtrahend.GetIntegerValue()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "subtrahendBigInt, err := subtrahend.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  subtrahendPrecisionBigInt, err := subtrahend.GetPrecisionBigInt()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "subtrahendPrecisionBigInt, err := subtrahend.GetPrecisionBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  // The error is ignored because minuend precision and
  // subtrahend precision will never be less than zero
  difBigInt, difBigIntPrecision, err :=
    new(bigIntMathSubtractNanobot).bigIntSubtract(
      minuendBigInt,
      minuendPrecisionBigInt,
      subtrahendBigInt,
      subtrahendPrecisionBigInt,
      ePrefix)

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "difBigInt, difBigIntPrecision, err := \n" +
          "  new(bigIntMathSubtractNanobot).bigIntSubtract(\n" +
          "  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt,\n" +
          "  subtrahendPrecisionBigInt, ePrefix)",
        ErrContext: fmt.Sprintf("minuendBigInt= '%v\n"+
          "minuendPrecisionBigInt= '%v\n"+
          "subtrahendBigInt= '%v\n"+
          "subtrahendPrecisionBigInt= '%v\n",
          minuendBigInt.Text(10), minuendPrecisionBigInt.Text(10),
          subtrahendBigInt.Text(10),
          subtrahendPrecisionBigInt.Text(10)),
        ErrMessage: err.Error(),
      }
  }

  err = difference.SetNumericValue(difBigInt, uint(difBigIntPrecision.Uint64()))

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = difference.SetNumericValue(\n" +
          "  difBigInt, uint(difBigIntPrecision.Uint64()))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = difference.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = difference.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = difference.IsValid(ePrefix.XCpy("Validating final result 'difference'").String())

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = difference.IsValid(ePrefix)",
        ErrContext: "Error: Final calculated result 'difference' is invalid.\n" +
          "'difference' FAILED final validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return difference, nil
}

// subtractBigInts
//
//	Performs the subtraction operation and returns the 'difference' as
//	a type BigIntNum.
//
//	In the subtraction operation:
//
//	  b1 - b2 = difference or result
//	  'minuend' - 'subtrahend' = difference or result
//	  b1 = 'minuend'
//	  b2 = 'subtrahend'
//
//	Input Parameters
//	================
//
//	minuend                  *big.Int
//	  The number from which the subtrahend will be subtracted.
//
//	minuendPrecision         uint
//	  The 'minuend' precision or numeric digits after the decimal
//	  point.
//
//	subtrahend               *big.Int
//	  The number to be subtracted from the 'minuend'.
//
//	subtrahendPrecision      uint
//	  The 'subtrahend' precision or numeric digits after the decimal
//	  point.
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
//	BigIntNum
//	  After the subtraction operation, the 'difference' or 'result' is
//	  returned as a Type BigIntNum.
//
//	  The returned BigIntNum 'result' will contain USA default numeric
//	  separators (decimal separator, thousands separator and currency
//	  symbol).
//
//	err                      error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractBigInts(
  numSeps NumericSeparatorDto,
  minuend *big.Int,
  minuendPrecision uint,
  subtrahend *big.Int,
  subtrahendPrecision uint,
  errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractBigInts",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if minuend == nil {

    return BigIntNum{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'minuend'",
      }
  }

  if subtrahend == nil {

    return BigIntNum{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'subtrahend'",
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
          "\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
          "'numSeps' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  // Returned error is ignored because minuend and subtrahend precisions
  // will never be less than zero.
  result, resultPrecision, err :=
    new(bigIntMathSubtractNanobot).bigIntSubtract(
      minuend,
      big.NewInt(0).SetUint64(uint64(minuendPrecision)),
      subtrahend,
      big.NewInt(0).SetUint64(uint64(subtrahendPrecision)),
      ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "result, resultPrecision, err :=\n" +
          "  new(bigIntMathSubtractNanobot).bigIntSubtract(\n" +
          "minuend, big.NewInt(0).SetUint64(uint64(minPrecision)),\n" +
          "subtrahend, big.NewInt(0).SetUint64(uint64(subPrecision)), ePrefix)",
        ErrContext: fmt.Sprintf("minuend= '%v'\n"+
          "minuendPrecision= '%v'\n"+
          "subtrahend= '%v'\n"+
          "subtrahendPrecision= '%v'",
          minuend.Text(10), minuendPrecision,
          subtrahend.Text(10), subtrahendPrecision),
        ErrMessage: err.Error(),
      }
  }

  biNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "biNum, err := new(BigIntNum).\n" +
          "  NewBigIntBigPrecision(result, resultPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = biNum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = biNum.SetNumericSeparatorsDto(numSeps)",
        ErrContext: fmt.Sprintf("numSeps= '%v'",
          numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  err = biNum.IsValid(ePrefix.XCpy("Validating final result 'biNum'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = biNum.IsValid(ePrefix)",
        ErrContext: "Error: Final caclulated result 'biNum' is invalid.\n" +
          "'biNum' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return biNum, nil
}

// subtractBigIntNumArray
//
//	Receives one BigIntNum which is classified as the 'minuend'.
//
//	The second input parameter is an array of BigIntNum Types labeled,
//	'subtrahends'.
//
//	The array of 'subtrahends' is subtracted from the 'minuend'.
//
//	In the subtraction operation:
//
//	  b1 = 'minuend'
//	  b2 = 'subtrahend'
//	  b1 - b2 = difference or result
//	  'minuend' - 'subtrahend' = difference or result
//
//	In this method, the 'subtrahend' is an array of BigIntNum Types.
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
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' will be copied to the returned
//	instance of 'difference' (type BigIntFixedDecimal).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractBigIntNumArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend BigIntNum,
  validateMinuend bool,
  subtrahends []BigIntNum,
  validateSubtrahends bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractBigIntNumArray",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  finalResult, err := minuend.CopyOut()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "finalResult , err := minuend.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    err = finalResult.SetNumericSeparatorsDto(numSeps)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    return finalResult, nil
  }

  var bPair BigIntPair

  var finalResultNumStr, subtrahendNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahends {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating subtrahends[%d]", i)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: fmt.Sprintf("err = subtrahends[%d].IsValid(ePrefix)", i),
            ErrContext: fmt.Sprintf("Error: subtrahends[%d] is invalid.\n"+
              "subtrahends[%d] FAILED validation tests.", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    finalResultNumStr, err = finalResult.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "finalResultNumStr, err = finalResult.GetNumStr()",
          ErrContext: fmt.Sprintf("Error: finalResult.GetNumStr() failed. Cicle # %d", i),
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = subtrahends[i].GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("subtrahendNumStr, err = subtrahends[%d].GetNumStr()", i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(finalResult, subtrahends[i])

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewBigIntNum(finalResult, subtrahends[%d])", i),
          ErrContext: fmt.Sprintf("finalResult= '%v'\n"+
            "subtrahends[%d]= '%v'",
            finalResultNumStr, i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    finalResult, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "finalResult, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(\n" +
            "  numSeps, false, &bPair, true, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  difference, err = finalResult.CopyOut()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = finalResult.CopyOut()",
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

// subtractBigIntNumOutputToArray
//
//	The first input parameter to this method is a BigIntNum Type
//	labeled, 'minuend'.  The second input parameter is an array of
//	BigIntNum types labeled 'subtrahends'.
//
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend' value with the 'dfference', or result, output to another
//	'results' array of BigIntNum types which is then returned to the
//	calling function.
//
//	Example
//	=======
//
//	              subtrahends                 Output
//	Minuend          Array                     Array
//
//	  10      -  subtrahends[0] = 2    =   outputarray[0] = 8
//	  10      -  subtrahends[1] = 3    =   outputarray[1] = 7
//	  10      -  subtrahends[2] = 4    =   outputarray[2] = 6
//	  10      -  subtrahends[3] = 5    =   outputarray[3] = 5
//	  10      -  subtrahends[4] = 6    =   outputarray[4] = 4
//	  10      -  subtrahends[5] = 9    =   outputarray[5] = 1
//
//	Each of the BigIntNum instances included in the array of BigIntNum
//	subtraction results returned by this method, will contain numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) copied from input parameter 'minuend'.
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractBigIntNumOutputToArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend BigIntNum,
  validateMinuend bool,
  subtrahends []BigIntNum,
  validateSubtrahends bool,
  errPrefDto *ePref.ErrPrefixDto) ([]BigIntNum, error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractBigIntNumOutputToArray",
    "")

  if err != nil {
    return []BigIntNum{}, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return []BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return []BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return []BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is invalid.\n" +
          "'subtrahends' is a zero length array.",
      }
  }

  resultsArray := make([]BigIntNum, lenSubtrahends)

  var bPair BigIntPair

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return []BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  var subtrahendNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahends {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating subtrahends[%d]", i)).String())

      if err != nil {

        return []BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: fmt.Sprintf("err = subtrahends[%d].IsValid(ePrefix)", i),
            ErrContext: fmt.Sprintf("Error: subtrahends[%d] is invalid.\n"+
              "subtrahends[%d] FAILED validation tests.", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = subtrahends[i].GetNumStr()

    if err != nil {

      return []BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("subtrahendNumStr, err = subtrahends[%d].GetNumStr()", i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(minuend, subtrahends[i])

    if err != nil {

      return []BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bPair, err = new(BigIntPair).NewBigIntNum(minuend, subtrahends[%d])", i),
          ErrContext: fmt.Sprintf("minuend= '%v'\n"+
            "subtrahends[%d]= '%v'",
            minuendNumStr, i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    resultsArray[i], err = new(bigIntMathSubtractNanobot).subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return []BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("resultsArray[%d], err = new(bigIntMathSubtractNanobot).\n"+
            "subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)", i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return resultsArray, nil
}

// subtractBigIntNums
//
//	 Receives two 'BigIntNum' instances and proceeds to subtract
//	 'subtrahend' from 'minuend' returning 'difference' as a type
//	 BigIntNum.
//
//	       'minuend' - 'subtrahend' = difference or result
//
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and
//		convert them into numeric values.
//
//		This method will copy the Numeric Separators configured
//		for input parameter 'numSeps' to the returned instance of
//		'difference' (type BigIntFixedDecimal).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractBigIntNums(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend *BigIntNum,
  validateMinuend bool,
  subtrahend *BigIntNum,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractBigIntNums",
    "")

  if err != nil {
    return difference, err
  }

  if minuend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'minuend'",
      }
  }

  if subtrahend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'minuend'",
      }
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateSubtrahend {

    err = subtrahend.IsValid(ePrefix.XCpy("Validating 'subtrahend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = subtrahend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'subtrahend'\").String())",
          ErrContext: "Error: Input parameter 'subtrahend' is invalid.\n" +
            "'subtrahend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  subtrahendNumStr, err := subtrahend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "subtrahendNumStr, err := subtrahend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bPair, err := new(BigIntPair).NewBigIntNum(*minuend, *subtrahend)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(*minuend, *subtrahend)",
        ErrContext: fmt.Sprintf("minuend= '%v'\n"+
          "subtrahend= '%v'\n",
          minuendNumStr, subtrahendNumStr),
        ErrMessage: err.Error(),
      }
  }

  difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
          "  subtractBigIntPair(numSeps, true, &bPair, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return difference, nil
}

// subtractBigIntNumSeries
//
//	Receives one BigIntNum which is classified as the 'minuend'. The
//	second input parameter, 'subtrahends' is a series of Type
//	BigIntNum's.
//
//	The 'subtrahends' series is subtracted from the 'minuend'.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//	    b1 = 'minuend'
//		  b2 = 'subtrahend'
//	    b1 - b2 = difference or result
//
//	In this method, the 'subtrahend' is a series of BigIntNum Types.
//	This method is defined as a variadic function in that 'subtrahend'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide a value for 'subtrahend',
//	or not.
//
//	If no value is provided for 'subtrahend', an error is returned.
//
//	Therefore, the user MUST provide one or more 'subtrahend' values,
//	separated by commas.
//
//	After subtracting all 'subtrahend' values from 'minuend', the
//	resulting cumulative 'difference' value is returned as a Type
//	BigIntNum.
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
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' to the returned instance of
//	'difference' (type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractBigIntNumSeries(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend BigIntNum,
  validateMinuend bool,
  validateSubtrahends bool,
  errPrefDto *ePref.ErrPrefixDto,
  subtrahends ...BigIntNum) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractBigIntNumSeries",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  beginningMinuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "beginningMinuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.CopyOut()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  foundSubtrahends := false

  var bPair BigIntPair
  var subtrahendNumStr string

  for idx, subtrahend := range subtrahends {

    foundSubtrahends = true

    if validateSubtrahends {

      err = subtrahend.IsValid(ePrefix.XCpy(fmt.Sprintf("Validating subtrahend[%d]", idx)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: fmt.Sprintf("err = subtrahend[%d].IsValid(ePrefix)", idx),
            ErrContext: fmt.Sprintf("Error: subtrahend[%d] is invalid.\n"+
              "subtrahend[%d] FAILED validation tests.", idx, idx),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = subtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "difference, err = minuend.CopyOut()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(minuend, subtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bPair, err = new(BigIntPair).NewBigIntNum(minuend, subtrahend[%d])", idx),
          ErrContext: fmt.Sprintf("Beginning 'minuend'= '%v'\n"+
            "subtrahend[%d]= '%v'",
            beginningMinuendNumStr, idx, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(\n" +
            "  numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("Beginning 'minuend'= '%v'\n"+
            "subtrahend[%d]= '%v'",
            beginningMinuendNumStr, idx, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }
  }

  if !foundSubtrahends {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.CopyOut()",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'subtrahends' is invalid.\n" +
          "'subtrahends' is empty and conains zero subtrahend values.",
      }

  }

  return difference, nil
}

// subtractDecimals
//
//	Performs the subtraction operation on two Decimal Types.
//
//	    decMinuend - decSubtrahend = difference
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//	    decimal1  = 'minuend'
//	    decimal2  = 'subtrahend'
//	    decimal1  -  decimal2  =  difference or result
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
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured for input
//	parameter 'decMinuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractDecimals(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  decMinuend *Decimal,
  validateMinuend bool,
  decSubtrahend *Decimal,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractDecimals",
    "")

  if err != nil {
    return difference, err
  }

  if decMinuend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'decMinuend'",
      }
  }

  if decSubtrahend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'decSubtrahend'",
      }
  }

  if validateMinuend {

    err = decMinuend.IsValid(ePrefix.XCpy("Validating 'decMinuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = decMinuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'decMinuend'\").String())",
          ErrContext: "Error: Input parameter 'decMinuend' is invalid.\n" +
            "'decMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateSubtrahend {

    err = decSubtrahend.IsValid(ePrefix.XCpy("Validating 'decSubtrahend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = decSubtrahend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'decSubtrahend'\").String())",
          ErrContext: "Error: Input parameter 'decSubtrahend' is invalid.\n" +
            "'decSubtrahend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decMinuendNumStr, err := decMinuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decSubtrahendNumStr, err := decSubtrahend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // This method will test the validity of decMinuend and decSubtrahend.

  bPair, err := new(BigIntPair).NewDecimal(*decMinuend, *decSubtrahend)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewDecimal(*decMinuend, *decSubtrahend)",
        ErrContext: fmt.Sprintf("decMinuend= '%v'\n"+
          "decSubtrahend= '%v'",
          decMinuendNumStr, decSubtrahendNumStr),
        ErrMessage: err.Error(),
      }
  }

  difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
    numSeps, false, &bPair, true, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(\n" +
          "  numSeps, false, &bPair, true, ePrefix)",
        ErrContext: fmt.Sprintf("decMinuend= '%v'\n"+
          "decSubtrahend= '%v'\n"+
          "numSeps= '%v'",
          decMinuendNumStr, decSubtrahendNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  return difference, nil
}

// subtractDecimalArray
//
//	Receives one Decimal parameter which is classified as the
//	'minuend'. The second input parameter is an array of Decimal Types
//	labeled, 'subtrahends'.
//
//	The array of 'subtrahends' is subtracted from the 'minuend'.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahends' = difference or result
//	    d1 = 'minuend'
//	    d2 = 'subtrahend'
//	    d1 - d2 = difference or result
//
//	In this method, the 'subtrahends' is an array of Decimal Types.
//	Each element in the 'subtrahends' array is subtracted from the
//	value of 'minuend'. After the subtraction operation, the
//	cumulative 'difference' or 'result' is returned as a Type
//	BigIntNum.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractDecimalArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend Decimal,
  validateMinuend bool,
  subtrahends []Decimal,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractDecimalArray",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'decMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.GetBigIntNum()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  differenceNumStr, err := difference.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "differenceNumStr, err := difference.GetNumStr()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array with zero elements.",
      }
  }

  var bigINumSubtrahend BigIntNum
  var subtrahendNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahend {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahends[%d]'", i)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahends[%d]' is Invalid!\n"+
              "'subtrahends[%d]' Failed Validation Tests", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: fmt.Sprintf("bigINumSubtrahend[%d]", i),
          ErrMessage: err.Error(),
        }
    }

    bigINumSubtrahend, err = subtrahends[i].GetBigIntNum()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err =\n"+
            "subtrahends[%d].GetBigIntNum()", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'",
            i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err := difference.GetNumStr()",
          ErrContext: fmt.Sprintf("Cycle #= '%v'", i),
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bigINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(difference, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "subtrahend[%d]= '%v'",
            differenceNumStr, i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No.= %v",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}

// subtractDecimalOutputToArray
//
//	The first input parameter to this method is a Decimal Type
//	labeled, 'minuend'.  The second input parameter is an array of
//	Decimal types labeled 'subtrahends'. The 'minuend' is subtracted
//	from each element of the 'subtrahends' array with the result
//	output to a second array of Decimal types labeled, 'results',
//	which is then returned to the calling function.
//
//	Example
//	=======
//
//	                     subtrahends                 Output
//	  minuend              Array                      Array
//
//	    10      -     subtrahends[0] = 2    =    outputarray[0] =  8
//	    10      -     subtrahends[1] = 3    =    outputarray[1] =  7
//	    10      -     subtrahends[2] = 4    =    outputarray[2] =  6
//	    10      -     subtrahends[3] = 5    =    outputarray[3] =  5
//	    10      -     subtrahends[4] = 6    =    outputarray[4] =  4
//	    10      -     subtrahends[5] = 9    =    outputarray[5] =  1
//
//	Numeric Separators
//	==================
//
//	The array element of the []Decimal 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'result' array ([]Decimal) returned by this subtraction
//	operation will contain array elements with numeric separators
//	(decimal separator, thousands separator and currency symbol) which
//	have been copied from input parameter 'minuend'.
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractDecimalOutputToArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend Decimal,
  validateMinuend bool,
  subtrahends []Decimal,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (results []Decimal, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  results = []Decimal{}

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractDecimalOutputToArray",
    "")

  if err != nil {
    return results, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bINumMinuend, err := minuend.GetBigIntNum()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  minuendNumStr, err = bINumMinuend.GetNumStr()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := bINumMinuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array with zero elements.",
      }
  }

  resultsArray := make([]Decimal, lenSubtrahends)

  var bigINumSubtrahend, result BigIntNum
  var subtrahendNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahend {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahends[%d]'", i)).String())

      if err != nil {

        return results,
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahends[%d]' is Invalid!\n"+
              "'subtrahends[%d]' Failed Validation Tests", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = subtrahends[i].GetNumStr()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("subtrahendNumStr, err = subtrahends[%d].GetNumStr()", i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bigINumSubtrahend, err = subtrahends[i].GetBigIntNum()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err =\n"+
            "subtrahends[%d].GetBigIntNum()", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'",
            i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: fmt.Sprintf("bigINumSubtrahend[%d]", i),
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(bINumMinuend, bigINumSubtrahend)

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(bINumMinuend, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("bINumMinuend= '%v'\n"+
            "bigINumSubtrahend= '%v'\n"+
            "Cycle No. = '%v'",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }

    result, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No. = '%v'",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }

    resultsArray[i], err = result.GetDecimal()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("resultsArray[%d], err = result.GetDecimal()",
            i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return resultsArray, nil
}

// subtractDecimalSeries
//
//	Receives one Decimal Type which is classified as the 'minuend'.
//	The second input parameter, 'subtrahends' is a series of Type Decimal.
//
//	The 'subtrahends' series is subtracted from the 'minuend'.
//
//	In the subtraction operation:
//
//	  'minuend' - 'subtrahend' = difference or result
//	  b1 = 'minuend'
//	  b2 = 'subtrahends'
//	  b1 - b2 = difference or result
//
//	In this method, the 'subtrahend' is a series of Decimal Types.
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide a value for 'subtrahend',
//	or not.
//
//	After subtracting all 'subtrahend' values from 'minuend', the
//	resulting cumulative 'difference' value is returned as a Type
//	BigIntNum.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' to the returned instance of
//	'difference' (type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractDecimalSeries(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend Decimal,
  validateMinuend bool,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto,
  subtrahends ...Decimal) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = BigIntNum{}

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractDecimalSeries",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'decMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.GetBigIntNum()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  minuendNumStr, err = difference.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err = difference.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if len(subtrahends) == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is empty and contains zero elements.",
      }
  }

  var bINumSubtrahend BigIntNum
  var bPair BigIntPair
  var subtrahendNumStr, differenceNumStr string

  for idx, subtrahend := range subtrahends {

    if validateSubtrahend {

      err = subtrahend.IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahend[%d]'", idx)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahend[%d]' is Invalid!\n"+
              "'subtrahend[%d]' Failed Validation Tests", idx, idx),
            ErrMessage: err.Error(),
          }
      }
    }

    bINumSubtrahend, err = subtrahend.GetBigIntNum()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bINumSubtrahend, err =\n"+
            "subtrahend[%d].GetBigIntNum()", idx),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = subtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = subtrahend.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err = difference.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(difference, bINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "bINumSubtrahend= '%v'", differenceNumStr, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "  bPair2= '%v'", minuendNumStr, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}

// subtractIntAry
//
//	Performs the subtraction operation on two instances of type
//	IntAry.
//
//	    iaMinuend - iaSubtrahend = difference
//
//	In the subtraction operation:
//
//	'minuend' - 'subtrahend' = difference or result
//	b1 = 'minuend'
//	b2 = 'subtrahend'
//	b1 - b2 = difference or result
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
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured	for
//	input parameter 'iaMinuend' to the returned instance of
//	'difference' (type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractIntAry(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  iaMinuend *IntAry,
  validateMinuend bool,
  iaSubtrahend *IntAry,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractIntAry",
    "")

  if err != nil {
    return difference, err
  }

  if iaMinuend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'iaMinuend'",
      }
  }

  if iaSubtrahend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'iaSubtrahend'",
      }
  }

  if validateMinuend {

    err = iaMinuend.IsValid(ePrefix.XCpy("Validating 'iaMinuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = iaMinuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'iaMinuend'\").String())",
          ErrContext: "Error: Input parameter 'iaMinuend' is invalid.\n" +
            "'iaMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateSubtrahend {

    err = iaSubtrahend.IsValid(ePrefix.XCpy("Validating 'iaSubtrahend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = iaSubtrahend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'iaSubtrahend'\").String())",
          ErrContext: "Error: Input parameter 'iaSubtrahend' is invalid.\n" +
            "'iaSubtrahend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaMinuendNumStr, err := iaMinuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // Method NewIntAry will test validity of iaMinuend and iaSubtrahend
  bPair, err := new(BigIntPair).NewIntAry(*iaMinuend, *iaSubtrahend)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(*iaMinuend, *iaSubtrahend)",
        ErrContext: fmt.Sprintf("iaMinuend= '%v'\n"+
          "iaSubtrahend= '%v'", iaMinuendNumStr, iaSubtrahendNumStr),
        ErrMessage: err.Error(),
      }
  }

  difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
    numSeps, false, &bPair, true, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(*iaMinuend, *iaSubtrahend)",
        ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
          "bPair2= '%v'\n"+
          "numSeps= '%v'", iaMinuendNumStr, iaSubtrahendNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  return difference, nil
}

// subtractIntAryArray
//
//	Receives one IntAry parameter which is classified as the
//	'minuend'. The second input parameter is an array of IntAry Types
//	labeled, 'subtrahends'.
//
//	The array of 'subtrahends' is subtracted from the 'minuend'.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//	    b1 = 'minuend'
//	    b2 = 'subtrahend'
//	    b1 - b2 = difference or result
//
//	In this method, the 'subtrahends' is an array of Decimal Types.
//	Each element in the 'subtrahends' array is subtracted from the
//	value of 'minuend'. After the subtraction operation, the
//	cumulative 'difference' or 'result' is returned as a Type
//	BigIntNum.
//
//	After the subtraction operation, the cumulative 'difference' or
//	'result' is returned as a Type BigIntNum.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractIntAryArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend IntAry,
  validateMinuend bool,
  subtrahends []IntAry,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractIntAryArray",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'decMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.GetBigIntNum()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  differenceNumStr, err := difference.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "differenceNumStr, err := difference.GetNumStr()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array with zero elements.",
      }
  }

  var bigINumSubtrahend BigIntNum
  var subtrahendNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahend {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahends[%d]'", i)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahends[%d]' is Invalid!\n"+
              "'subtrahends[%d]' Failed Validation Tests", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: fmt.Sprintf("bigINumSubtrahend[%d]", i),
          ErrMessage: err.Error(),
        }
    }

    bigINumSubtrahend, err = new(BigIntNum).NewIntAry(subtrahends[i])

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err =\n"+
            "new(BigIntNum).NewIntAry(subtrahends[%d])", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'",
            i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err := difference.GetNumStr()",
          ErrContext: fmt.Sprintf("Cycle #= '%v'", i),
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bigINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(difference, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "subtrahend[%d]= '%v'",
            differenceNumStr, i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No.= %v",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}

// subtractIntAryOutputToArray
//
//	The first input parameter to this method is a IntAry Type labeled,
//	'minuend'.  The second input parameter is an array of IntAry types
//	labeled 'subtrahends'.
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend' value with the 'dfference', or result, output to another
//	'results' array of IntAry types which is then returned to the
//	calling function.
//
//	Example
//	=======
//
//	                subtrahends                   Output
//	minuend            Array                       Array
//
//	  10      -    subtrahends[0] = 2     =     outputarray[0] =  8
//	  10      -    subtrahends[1] = 3     =     outputarray[1] =  7
//	  10      -    subtrahends[2] = 4     =     outputarray[2] =  6
//	  10      -    subtrahends[3] = 5     =     outputarray[3] =  5
//	  10      -    subtrahends[4] = 6     =     outputarray[4] =  4
//	  10      -    subtrahends[5] = 9     =     outputarray[5] =  1
//
//	Numeric Separators
//	==================
//
//	Each array element of the []IntAry 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'result' array ([]IntAry) returned by this subtraction
//	operation will contain array elements with numeric separators
//	(decimal separator, thousands separator and currency symbol) which
//	have been copied from input parameter 'minuend'.
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractIntAryOutputToArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend IntAry,
  validateMinuend bool,
  subtrahends []IntAry,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (results []IntAry, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  results = []IntAry{}

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractIntAryOutputToArray",
    "")

  if err != nil {
    return results, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bINumMinuend, err := minuend.GetBigIntNum()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  minuendNumStr, err = bINumMinuend.GetNumStr()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := bINumMinuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array with zero elements.",
      }
  }

  resultsArray := make([]IntAry, lenSubtrahends)

  var bigINumSubtrahend, result BigIntNum
  var subtrahendNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahend {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahends[%d]'", i)).String())

      if err != nil {

        return results,
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahends[%d]' is Invalid!\n"+
              "'subtrahends[%d]' Failed Validation Tests", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = subtrahends[i].GetNumStr()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("subtrahendNumStr, err = subtrahends[%d].GetNumStr()", i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    // Method NewIntAry will test validity of subtrahends[i]
    bigINumSubtrahend, err = subtrahends[i].GetBigIntNum()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err =\n"+
            "subtrahends[%d].GetBigIntNum()", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'",
            i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: fmt.Sprintf("bigINumSubtrahend[%d]", i),
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(bINumMinuend, bigINumSubtrahend)

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(bINumMinuend, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("bINumMinuend= '%v'\n"+
            "bigINumSubtrahend= '%v'\n"+
            "Cycle No. = '%v'",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }

    result, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No. = '%v'",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }

    resultsArray[i], err = result.GetIntAry()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("resultsArray[%d], err = result.GetIntAry()",
            i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return resultsArray, nil
}

// subtractIntArySeries
//
//	Receives one IntAry Type which is classified as the 'minuend'. The
//	second input parameter, 'subtrahends' is a series of Type IntAry .
//
//	The 'subtrahends' series is subtracted from the 'minuend' and the
//	net result is returned in parameter 'difference' (type BigIntNum).
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//	    b1 = 'minuend'
//	    b2 = 'subtrahend'
//	    b1 - b2 = difference or result
//
//	In this method, the 'subtrahend' is a series of IntAry Types.
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide a value for 'subtrahend',
//	or not.
//
//	After subtracting all 'subtrahend' values from 'minuend', the
//	resulting cumulative 'difference' value is returned as a Type
//	BigIntNum.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'minuend' to the returned instance of
//	'difference' (type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractIntArySeries(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend IntAry,
  validateMinuend bool,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto,
  subtrahends ...IntAry) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = BigIntNum{}

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractIntArySeries",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'decMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.GetBigIntNum()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  minuendNumStr, err = difference.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err = difference.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if len(subtrahends) == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is empty and contains zero elements.",
      }
  }

  var bINumSubtrahend BigIntNum
  var bPair BigIntPair
  var subtrahendNumStr, differenceNumStr string

  for idx, subtrahend := range subtrahends {

    if validateSubtrahend {

      err = subtrahend.IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahend[%d]'", idx)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahend[%d]' is Invalid!\n"+
              "'subtrahend[%d]' Failed Validation Tests", idx, idx),
            ErrMessage: err.Error(),
          }
      }
    }

    bINumSubtrahend, err = subtrahend.GetBigIntNum()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bINumSubtrahend, err =\n"+
            "subtrahend[%d].GetBigIntNum()", idx),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = subtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = subtrahend.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err = difference.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(difference, bINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "bINumSubtrahend= '%v'", differenceNumStr, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "  bPair2= '%v'", minuendNumStr, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}

// subtractINumMgr
//
//	Receives two objects which implement the INumMgr Interface and
//	subtracts their numeric values.
//
//	The 'subtrahend' numeric value is subtracted from the 'minuend'
//	numeric value.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = difference or result
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry.
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
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of	'difference'
//	(type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractINumMgr(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend INumMgr,
  validateMinuend bool,
  subtrahend INumMgr,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractINumMgr",
    "")

  if err != nil {
    return difference, err
  }

  if minuend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'minuend'",
      }
  }

  if subtrahend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'subtrahend'",
      }
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateSubtrahend {

    err = subtrahend.IsValid(ePrefix.XCpy("Validating 'subtrahend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = subtrahend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'subtrahend'\").String())",
          ErrContext: "Error: Input parameter 'subtrahend' is invalid.\n" +
            "'subtrahend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  subtrahendNumStr, err := subtrahend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "subtrahendNumStr, err := subtrahend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bPair, err := new(BigIntPair).NewINumMgr(minuend, subtrahend)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(minuend, subtrahend)",
        ErrContext: fmt.Sprintf("minuend= '%v'\n"+
          "subtrahend= '%v'",
          minuendNumStr, subtrahendNumStr),
        ErrMessage: err.Error(),
      }
  }

  difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
    numSeps, false, &bPair, true, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(\n" +
          "  numSeps, false, &bPair, true, ePrefix)",
        ErrContext: fmt.Sprintf("minuend= '%v'\n"+
          "subtrahend= '%v'\n"+
          "numSeps= '%v'",
          minuendNumStr, subtrahendNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  return difference, nil
}

// subtractINumMgrArray
//
//	Receives two input parameters. The first parameter is an INumMgr
//	instance which is classified as the 'minuend'. The second
//	parameter is an array of INumMgr instances which resents the
//	'subtrahends'.
//
//	A subtraction operation is performed on the 'minuend' and the
//	'subtrahends' array. The numeric value of the 'subtrahends' is
//	subtracted from the 'minuend'.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahends' = cumulative difference or result
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry. This allows the user to mix different types
//	in a single array and add their numeric values.
//
//	In this method, the 'subtrahends' is an array of INumMgr Types.
//	Each element in the 'subtrahends' array is subtracted from the
//	value of 'minuend'. After the subtraction operation, the
//	cumulative 'difference' or 'result' is returned as a Type
//	BigIntNum.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractINumMgrArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend INumMgr,
  validateMinuend bool,
  subtrahends []INumMgr,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractINumMgrArray",
    "")

  if err != nil {
    return difference, err
  }

  if minuend == nil {

    return difference,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'minuend'",
      }
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.GetBigIntNum()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  differenceNumStr, err := difference.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "differenceNumStr, err := difference.GetNumStr()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array with zero elements.",
      }
  }

  var bigINumSubtrahend BigIntNum
  var subtrahendNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahend {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahends[%d]'", i)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahends[%d]' is Invalid!\n"+
              "'subtrahends[%d]' Failed Validation Tests", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: fmt.Sprintf("bigINumSubtrahend[%d]", i),
          ErrMessage: err.Error(),
        }
    }

    bigINumSubtrahend, err = new(BigIntNum).NewINumMgr(subtrahends[i])

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err =\n"+
            "new(BigIntNum).NewINumMgr(subtrahends[%d])", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'",
            i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err := difference.GetNumStr()",
          ErrContext: fmt.Sprintf("Cycle #= '%v'", i),
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bigINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(difference, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "subtrahend[%d]= '%v'",
            differenceNumStr, i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No.= %v",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}

// subtractINumMgrOutputToArray
//
//	The first input parameter to this method is an object which
//	implements the INumMgr interface ('minuend').  The second input
//	parameter is an array of INumMgr interface types labeled
//	'subtrahends'. The 'minuend' is subtracted from each element of
//	the 'subtrahends' array with the result output to another
//	'results' array of INumMgr interface types which is then returned
//	to the calling function.
//
//	Example
//	=======
//
//	                subtrahends                 Output
//	Minuend            Array                     Array
//
//	  10     -    subtrahends[0] = 2    =   outputarray[0] =  8
//	  10     -    subtrahends[1] = 3    =   outputarray[1] =  7
//	  10     -    subtrahends[2] = 4    =   outputarray[2] =  6
//	  10     -    subtrahends[3] = 5    =   outputarray[3] =  5
//	  10     -    subtrahends[4] = 6    =   outputarray[4] =  4
//	  10     -    subtrahends[5] = 9    =   outputarray[5] =  1
//
//	Note: The underlying type for the returned results array is
//	'BigIntNum', a type which implements the INumMgr Interface.
//
//	Numeric Separators
//	==================
//
//	Each array element of the []INumMgr 'result' returned by this
//	subtraction operation will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	input parameter 'minuend'.
//
//	The 'result' array ([]INumMgr) returned by this subtraction
//	operation will contain array elements with numeric separators
//	(decimal separator, thousands separator and currency symbol) which
//	have been copied from input parameter 'minuend'.
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractINumMgrOutputToArray(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend INumMgr,
  validateMinuend bool,
  subtrahends []INumMgr,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto) (results []INumMgr, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  results = []INumMgr{}

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractINumMgrOutputToArray",
    "")

  if err != nil {
    return results, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bINumMinuend, err := minuend.GetBigIntNum()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  minuendNumStr, err = bINumMinuend.GetNumStr()

  if err != nil {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := bINumMinuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return results,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array with zero elements.",
      }
  }

  resultsArray := make([]INumMgr, lenSubtrahends)

  var bigINumSubtrahend, result BigIntNum
  var subtrahendNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    if validateSubtrahend {

      err = subtrahends[i].IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahends[%d]'", i)).String())

      if err != nil {

        return results,
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahends[%d]' is Invalid!\n"+
              "'subtrahends[%d]' Failed Validation Tests", i, i),
            ErrMessage: err.Error(),
          }
      }
    }

    subtrahendNumStr, err = subtrahends[i].GetNumStr()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: fmt.Sprintf("subtrahendNumStr, err = subtrahends[%d].GetNumStr()", i),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    // Method NewIntAry will test validity of subtrahends[i]
    bigINumSubtrahend, err = subtrahends[i].GetBigIntNum()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err =\n"+
            "subtrahends[%d].GetBigIntNum()", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'",
            i, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: fmt.Sprintf("bigINumSubtrahend[%d]", i),
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(bINumMinuend, bigINumSubtrahend)

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(bINumMinuend, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("bINumMinuend= '%v'\n"+
            "bigINumSubtrahend= '%v'\n"+
            "Cycle No. = '%v'",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }

    result, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return results,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No. = '%v'",
            minuendNumStr, subtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }

    resultsArray[i] = &result
  }

  return resultsArray, nil
}

// subtractINumMgrSeries
//
//	Receives two input parameters. The first parameter is an INumMgr
//	instance which is classified as the 'minuend'.
//
//	The second  parameter is a series of INumMgr instances which
//	resents the 'subtrahends'.
//
//	The 'subtrahends' series is subtracted from the 'minuend' and the
//	net result is returned in parameter 'difference' (type BigIntNum).
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahend' = net difference or result
//	    b1 = 'minuend'
//	    b2 = 'subtrahend'
//	    b1 - b2 = net difference or result
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry.
//
//	In this method, the 'subtrahends' is a series of INumMgr Types.
//	This method is defined as a variadic function in that 'subtrahends'
//	is configured as an optional input parameter meaning that it is NOT
//	required. The user can choose to provide a value for 'subtrahend',
//	or not. Note that if the user chooses NOT to submit any valid
//	INumMgr objects for the 'subtrahends' parameter, an error will be
//	returned.
//
//	After subtracting all 'subtrahend' values from 'minuend', the
//	resulting net 'difference' value is returned as a Type BigIntNum.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators configured	for input
//	parameter 'minuend' to the returned instance of 'difference'
//	(type BigIntNum).
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractINumMgrSeries(
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  minuend INumMgr,
  validateMinuend bool,
  validateSubtrahend bool,
  errPrefDto *ePref.ErrPrefixDto,
  subtrahends ...INumMgr) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = BigIntNum{}

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractIntArySeries",
    "")

  if err != nil {
    return difference, err
  }

  if validateMinuend {

    err = minuend.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = minuend.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'decMinuend' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  minuendNumStr, err := minuend.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err := minuend.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  difference, err = minuend.GetBigIntNum()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "difference, err = minuend.GetBigIntNum()",
        ErrContext: fmt.Sprintf("minuend= '%v'", minuendNumStr),
        ErrMessage: err.Error(),
      }
  }

  minuendNumStr, err = difference.GetNumStr()

  if err != nil {

    return difference,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minuendNumStr, err = difference.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if len(subtrahends) == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is empty and contains zero elements.",
      }
  }

  var bINumSubtrahend BigIntNum
  var bPair BigIntPair
  var subtrahendNumStr, differenceNumStr string

  for idx, subtrahend := range subtrahends {

    if validateSubtrahend {

      err = subtrahend.IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahend[%d]'", idx)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = subtrahends[i].IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: 'subtrahend[%d]' is Invalid!\n"+
              "'subtrahend[%d]' Failed Validation Tests", idx, idx),
            ErrMessage: err.Error(),
          }
      }
    }

    bINumSubtrahend, err = subtrahend.GetBigIntNum()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bINumSubtrahend, err =\n"+
            "subtrahend[%d].GetBigIntNum()", idx),
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    subtrahendNumStr, err = subtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "subtrahendNumStr, err = subtrahend.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err = difference.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "  NewBigIntNum(difference, bINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "bINumSubtrahend= '%v'", differenceNumStr, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "  bPair2= '%v'", minuendNumStr, subtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}

// subtractNumStr
//
//		Receives two number strings and proceeds to subtract subtrahend
//		'n2' from minuend 'n1'.
//
//		Number Strings
//		==============
//
//		Number strings are strings of numeric digits. The digit characters
//		which comprise a number string are formatted to facilitate
//		conversion to a corresponding numeric value.
//
//		The number string parameters ('n1' and 'n2') passed to this method
//		must consist of a string of numeric digits representing a numeric
//		value. A leading minus sign (-), or surrounding parentheses '()',
//		may be included in this number string to indicate a negative
//		numeric value.
//
//		The number string of numeric digits may also include a delimiting
//		decimal separator to identify fractional digits to the right of
//		the decimal separator. This method uses the default USA Decimal
//		Separator ('.') to parse 'numStr' and identify any existing
//		fractional digits.
//
//		Input parameter 'numSeps' is a type NumericSeparatorDto and is
//		used to parse the number strings 'n1' and 'n2' to extract their
//		numeric values. 'numSeps' encapsulates the applicable decimal
//		separator, thousands separator and currency symbol.
//
//		'numSeps' is also used in configuring the BigIntNum return value
//		for this subtraction operation. This method will copy the Numeric
//		Separators contained in 'numSeps' to the returned instance of
//		'difference' (type BigIntNum).
//
//		Subtraction Operation
//	 =====================
//
//		    'minuend' - 'subtrahend' = difference or result
//		    n1 - n2 = difference or result
//		    n1 = 'minuend'
//		    n2 = 'subtrahend'
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractNumStr(
  n1 string,
  validateN1 bool,
  n2 string,
  validateN2 bool,
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractNumStr",
    "")

  if err != nil {
    return difference, err
  }

  if len(n1) == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'n1' is Invalid!\n" +
          "'n1' is an empty string",
      }

  }

  if len(n2) == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'n2' is Invalid!\n" +
          "'n2' is an empty string.",
      }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  n1BigIntNum, err := new(BigIntNum).NewNumStrWithNumSeps(n1, &numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n1BigIntNum, err := new(BigIntNum).NewNumStrWithNumSeps(n1, &numSeps)",
        ErrContext: fmt.Sprintf("Error converting input parameter 'n1' to BigIntNum!\n"+
          "n1= '%v'", n1),
        ErrMessage: err.Error(),
      }
  }

  if validateN1 {

    err = n1BigIntNum.IsValid(ePrefix.XCpy("Validating 'n1'").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = n1BigIntNum.IsValid(ePrefix.XCpy(\"Validating 'n1'\").String())",
          ErrContext: "Error: Input parameter 'n1' is invalid.\n" +
            "'n1' BigIntNum FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  n1BigIntNumStr, err := n1BigIntNum.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n1BigIntNumStr, err := n1BigIntNum.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  n2BigIntNum, err := new(BigIntNum).NewNumStrWithNumSeps(n2, &numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n2BigIntNum, err := new(BigIntNum).NewNumStrWithNumSeps(n2, &numSeps)",
        ErrContext: "Error converting input parameter 'n2' to BigIntNum!",
        ErrMessage: err.Error(),
      }
  }

  if validateN2 {

    err = n2BigIntNum.IsValid(ePrefix.XCpy("Validating 'n2'").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = n2BigIntNum.IsValid(ePrefix.XCpy(\"Validating 'n2'\").String())",
          ErrContext: "Error: Input parameter 'n2' is invalid.\n" +
            "'n2' BigIntNum FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  n2BigIntNumStr, err := n2BigIntNum.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n2BigIntNumStr, err := n2BigIntNum.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bPair, err := new(BigIntPair).NewBigIntNum(n1BigIntNum, n2BigIntNum)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(n1BigIntNum, n2BigIntNum)",
        ErrContext: fmt.Sprintf("n1BigIntNum= '%v'\n"+
          "n2BigIntNum= '%v'", n1BigIntNumStr, n2BigIntNumStr),
        ErrMessage: err.Error(),
      }
  }

  difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
    numSeps, false, &bPair, true, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewIntAry(*iaMinuend, *iaSubtrahend)",
        ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
          "bPair2= '%v'\n"+
          "numSeps= '%v'", n1BigIntNumStr, n2BigIntNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  return difference, nil
}

// subtractNumStrArray
//
//	Receives one number string input parameter which is classified as
//	the 'minuend'. The second input parameter is an array of number
//	strings labeled, 'subtrahends'.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. The digit characters
//	which comprise a number string are formatted to facilitate
//	conversion to a corresponding numeric value.
//
//	The number string parameters ('minuend' and 'subtrahends') passed
//	to this method must consist of stringd of numeric digits
//	representing a numeric value. A leading minus sign (-), or
//	surrounding parentheses '()', may be included in this number
//	string to indicate a negative numeric value.
//
//	The number string of numeric digits may also include a delimiting
//	decimal separator to identify fractional digits to the right of
//	the decimal separator. This method uses the Decimal Separator
//	extracted from input parameter 'numSeps' to parse 'numStr' and
//	identify any existing fractional digits.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	This method will copy the Numeric Separators provided by input
//	parameter 'numSeps'. The Numeric Separators configured for the
//	returned instance of 'difference' (type BigIntNum) will also be
//	copied from input parameter 'numSeps'.
//
//	Subtraction Operation
//	=====================
//
//	Each element in the 'subtrahends' array is subtracted from the
//	'minuend'. The summary result of this subtraction operation is
//	then returned as a BigIntNum type.
//
//	In the subtraction operation:
//
//	    'minuend' - 'subtrahends' = cumulative difference or result
//
//	In this method, 'subtrahends' is an array of number strings
//	(Type string).
//
//	After all subtraction operations have been completed, the cumulative
//	'difference' or 'result' is returned as a Type BigIntNum.
//
//	                  subtrahends
//	minuend              Array                   difference
//
//	  50      -     subtrahends[0] = 2
//	  48      -     subtrahends[1] = 3
//	  45      -     subtrahends[2] = 4
//	  41      -     subtrahends[3] = 5
//	  36      -     subtrahends[4] = 6
//	  30      -     subtrahends[5] = 9
//
//	                      Final Returned 'difference' = 21
func (bIMathSubMacrobot *bigIntMathSubtractMacrobot) subtractNumStrArray(
  minuend string,
  validateMinuend bool,
  subtrahends []string,
  validateSubtrahends bool,
  numSeps NumericSeparatorDto,
  validateNumSeps bool,
  errPrefDto *ePref.ErrPrefixDto) (difference BigIntNum, err error) {

  bIMathSubMacrobot.lock.Lock()

  defer bIMathSubMacrobot.lock.Unlock()

  difference = new(BigIntNum).New()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathSubtractMacrobot.subtractNumStrArray",
    "")

  if err != nil {
    return difference, err
  }

  if len(minuend) == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'minuend' is Invalid!\n" +
          "'minuend' is an empty string",
      }

  }

  lenSubtrahends := len(subtrahends)

  if lenSubtrahends == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(subtrahends) == 0",
        ErrMessage: "Error: Input parameter 'subtrahends' is Invalid!\n" +
          "'subtrahends' is an empty array.",
      }
  }

  if validateNumSeps {

    err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

    if err != nil {

      return difference,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
            "  \"Validating 'numSeps'\").String())",
          ErrContext: "Error: Input parameter 'numSeps' is invalid.\n" +
            "'numSeps' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  difference, err = new(BigIntNum).NewNumStrWithNumSeps(minuend, &numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "difference, err := new(BigIntNum).\n" +
          "  NewNumStrWithNumSeps(minuend, &numSeps)",
        ErrContext: fmt.Sprintf("Error converting input parameter 'minuend' to BigIntNum!\n"+
          "minuend= '%v'", minuend),
        ErrMessage: err.Error(),
      }
  }

  if validateMinuend {

    err = difference.IsValid(ePrefix.XCpy("Validating 'minuend'").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = differenceBigIntNum.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'minuend'\").String())",
          ErrContext: "Error: Input parameter 'minuend' is invalid.\n" +
            "'minuend' BigIntNum (differenceBigIntNum) FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  var bigINumSubtrahend BigIntNum
  var bigInumSubtrahendNumStr, differenceNumStr string
  var bPair BigIntPair

  for i := 0; i < lenSubtrahends; i++ {

    bigINumSubtrahend, err = new(BigIntNum).NewNumStrWithNumSeps(subtrahends[i], &numSeps)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: fmt.Sprintf("bigINumSubtrahend, err := new(BigIntNum).\n"+
            "  NewNumStrWithNumSeps(subtrahends[%d], &numSeps)", i),
          ErrContext: fmt.Sprintf("subtrahends[%d]= '%v'\n"+
            "numSeps= '%v'", i, subtrahends[i], numSeps.String()),
          ErrMessage: err.Error(),
        }
    }

    if validateSubtrahends {

      err = bigINumSubtrahend.IsValid(ePrefix.XCpy(fmt.Sprintf("Validating 'subtrahend[%d]'", i)).String())

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = bigINumSubtrahend.IsValid(ePrefix)",
            ErrContext: fmt.Sprintf("Error: subtrahend[%d] Failed Validation Tests.\n"+
              "subtrahend[%d]= '%v'\n"+
              "numSpes= '%v'", i, i, subtrahends[i], numSeps.String()),
            ErrMessage: err.Error(),
          }
      }
    }

    bigInumSubtrahendNumStr, err = bigINumSubtrahend.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "bigInumSubtrahendNumStr, err = bigINumSubtrahend.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    differenceNumStr, err = difference.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "differenceNumStr, err = difference.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    bPair, err = new(BigIntPair).NewBigIntNum(difference, bigINumSubtrahend)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "bPair, err = new(BigIntPair).\n" +
            "NewBigIntNum(difference, bigINumSubtrahend)",
          ErrContext: fmt.Sprintf("difference= '%v'\n"+
            "bigINumSubtrahend= '%v'",
            differenceNumStr, bigInumSubtrahendNumStr),
          ErrMessage: err.Error(),
        }
    }

    difference, err = new(bigIntMathSubtractNanobot).subtractBigIntPair(
      numSeps, false, &bPair, true, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "difference, err = new(bigIntMathSubtractNanobot).\n" +
            "  subtractBigIntPair(numSeps, false, &bPair, true, ePrefix)",
          ErrContext: fmt.Sprintf("bPair1= '%v'\n"+
            "bPair2= '%v'\n"+
            "Cycle No.= %v",
            differenceNumStr, bigInumSubtrahendNumStr, i),
          ErrMessage: err.Error(),
        }
    }
  }

  return difference, nil
}
