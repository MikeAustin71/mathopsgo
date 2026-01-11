package mathops

import (
  "fmt"
  "sync"

  ePref "github.com/MikeAustin71/errpref"
)

type intAryMathAddMechanics struct {
  lock sync.Mutex
}

// RunTotal
//
//	Adds to IntAry input parameters and returns the results in the
//	first parameter, 'ia1'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIa1' is set to true, this
//	method will subject 'ia1' to validation tests.
//
//	Likewise, if input parameter 'validateIa2' is set to true, this
//	method will subject 'ia2' to validation tests.
//
//	Numeric Separators
//	==================
//
//	The returned addition 'result' in 'ia1' will contain numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) copied from the original 'ia' structure. In other
//	words, the 'ia1' numeric separators will remain unchanged.
func (iaAddMech *intAryMathAddMechanics) runTotal(
  ia1 *IntAry,
  validateIa1 bool,
  ia2 *IntAry,
  validateIa2 bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaAddMech.lock.Lock()

  defer iaAddMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMathAddMechanics.runTotal()",
    "")

  if err != nil {
    return err
  }

  if ia1 == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'ia1'",
    }
  }

  if ia2 == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'ia2'",
    }
  }

  if validateIa1 {

    err = ia1.IsValid(ePrefix.XCpy("Validating 'ia1'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ia1.IsValid(ePrefix.XCpy(Validating 'ia1').String())",
        ErrContext: "Input parameter 'ia1' is INVALID!\n" +
          "'ia1' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
    }

  } else {

    err = ia1.SetInternalFlags()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ia1.SetInternalFlags()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  if validateIa2 {

    err = ia2.IsValid(ePrefix.XCpy("Validating 'ia2'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ia2.IsValid(ePrefix.XCpy(Validating 'ia2').String())",
        ErrContext: "Input parameter 'ia2' is INVALID!\n" +
          "'ia2' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
    }

  } else {

    err = ia2.SetInternalFlags()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ia2.SetInternalFlags()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  err = new(intAryAtom).setEqualArrayLengths(
    ia1, false, ia2, false, false, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = ia1.SetEqualArrayLengths(ia2)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if len(ia1.intAry) != len(ia2.intAry) {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "ia1 & ia2 intAry lengths ARE NOT EQUAL!\n" +
        "After running new(intAryAtom).setEqualArrayLengths()\n" +
        "Array lengths are not equal." +
        fmt.Sprintf("ia1 len = %d\nia2 len = %d\n"+
          "",
          len(ia1.intAry), len(ia2.intAry)),
    }
  }

  if ia2.isZeroValue {
    return nil
  }

  ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "ia1NumSeps, err := ia.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  compare, err := ia1.CompareAbsoluteValues(ia2)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "compare, err := ia.CompareAbsoluteValues(ia2)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  newSignVal := ia1.signVal
  doAdd := true
  isZeroResult := false
  doReverseNums := false

  if compare == 1 {
    // compare == + 1
    // Absolute Value: N1 > N2

    if ia1.signVal == 1 && ia2.signVal == 1 {
      doAdd = true
      newSignVal = 1
    } else if ia1.signVal == -1 && ia2.signVal == 1 {
      doAdd = false
      newSignVal = -1
    } else if ia1.signVal == -1 && ia2.signVal == -1 {
      doAdd = true
      newSignVal = -1
    } else {
      // Must Be ia.signVal == 1 && ia2.signVal == -1
      doAdd = false
      newSignVal = 1
    }

  } else if compare == -1 {
    // Absolute Values: N2 > N1
    if ia1.signVal == 1 && ia2.signVal == 1 {
      doAdd = true
      newSignVal = 1
    } else if ia1.signVal == -1 && ia2.signVal == 1 {
      doAdd = false
      doReverseNums = true
      newSignVal = 1
    } else if ia1.signVal == -1 && ia2.signVal == -1 {
      doAdd = true
      newSignVal = -1
    } else {
      // Must Be ia.signVal == 1 && ia2.signVal == -1
      doAdd = false
      doReverseNums = true
      newSignVal = -1
    }

  } else {
    // Must be compare == 0
    // Absolute Values: N1==N2
    if ia1.signVal == 1 && ia2.signVal == 1 {
      doAdd = true
      newSignVal = 1
    } else if ia1.signVal == -1 && ia2.signVal == 1 {
      doAdd = false
      newSignVal = 1
      isZeroResult = true
    } else if ia1.signVal == -1 && ia2.signVal == -1 {
      doAdd = true
      newSignVal = -1
    } else {
      // Must Be ia.signVal == 1 && ia2.signVal == -1
      doAdd = false
      newSignVal = 1
      isZeroResult = true
    }

  }

  err = new(intAryMathAddUtility).addToSubtract(
    ia1, false, ia2, false, newSignVal, doAdd, isZeroResult, doReverseNums, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryMathAddUtility).addToSubtract(\n" +
        "ia, ia2, newSignVal, doAdd, isZeroResult, doReverseNums, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = ia2.SetNumericSeparatorsDto(ia1NumSeps)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = ia.SetNumericSeparatorsDto(numSeps)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
