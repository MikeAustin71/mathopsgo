package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type intAryNeutron struct {
  lock *sync.Mutex
}

// ceiling
//
//	Returns an IntAry which constitutes the mathematical ceiling of
//	the current IntAry.
//
//	Examples
//	========
//
//	      Initial      Ceiling
//	       Value        Value
//	      -------      -------
//	        5.95          6
//	        5.05          6
//	        5             5
//	       -5.05         -5
//	        2.4           3
//	        2.9           3
//	       -2.7          -2
//	       -2            -2
func (iaNeutron *intAryNeutron) ceiling(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getBigInt()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if intAry == nil {

    return IntAry{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry,
      ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "Input parameter 'intAry' is INVALID!\n" +
            "'intAry' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  iAry2 := new(intAryElectron).newIntAry()

  intLen := intAry.intAryLen - intAry.precision

  intIdx := intLen - 1

  hasFracDigits, err := new(intAryNanobot).hasFractionalDigits(
    intAry, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "hasFracDigits, err := new(intAryNanobot).\n" +
          "  hasFractionalDigits(ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !hasFracDigits {

    // iAry2, err = ia.CopyOut()
    err = new(intAryProton).copy(&iAry2, intAry, false, true, ePrefix)

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = new(intAryProton).copy(&iAry2, intAry, false, true, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    return iAry2, nil
  }

  if intAry.signVal < 0 {

    t := make([]uint8, intAry.intAryLen)

    for i := 0; i < intLen; i++ {
      t[i] = intAry.intAry[i]
    }

    iAry2.intAry = t[0:]
    iAry2.intAryLen = intAry.intAryLen
    iAry2.precision = intAry.precision
    iAry2.signVal = intAry.signVal
    return iAry2, nil
  }

  t := make([]uint8, intAry.intAryLen+1)

  n1 := 0
  n2 := 0
  carry := 0
  adjFac := 1 * intAry.signVal
  for i := intIdx; i >= 0; i-- {

    n1 = int(intAry.intAry[i])

    if i == intIdx {

      if n1+adjFac < 0 {

        n2 = 10 + n1 + adjFac
        carry = -1

      } else if n1+adjFac > 9 {

        n2 = n1 + adjFac - 10
        carry = 1

      } else {

        n2 = n1 + adjFac
        carry = 0

      }

    } else {

      if n1+carry < 0 {

        n2 = 10 + n1
        carry = -1

      } else if n1+carry > 9 {

        n2 = n1 - 10
        carry = 1

      } else {

        n2 = n1 + carry
        carry = 0
      }
    }

    t[i+1] = uint8(n2)

  }

  if carry != 0 {

    t[0] = uint8(carry)
    iAry2.intAry = t[0 : intAry.intAryLen+1]

  } else {

    iAry2.intAry = t[1 : intAry.intAryLen+1]
  }

  iAry2.intAryLen = len(iAry2.intAry)

  iAry2.precision = intAry.precision

  iAry2.signVal = intAry.signVal

  return iAry2, nil
}

// changeSign
//
//	Changes the sign of the IntAry instance passed as input
//	parameter 'intAry'.
//
//	If the 'intAry' numeric value is positive (+), this method will
//	change the sign value to negative (-).
//
//	Conversely, if the 'intAry' is a negative (-)	numeric value,
//	this method will change the sign to	positive (+).
func (iaNeutron *intAryNeutron) changeSign(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.changeSign()",
    "")

  if err != nil {
    return err
  }

  if intAry == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intAry'",
    }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry,
      ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "intAry, ePrefix.XCpy(Validating 'intAry').String())",
        ErrContext: "Input parameter 'intAry' is INVALID!\n" +
          "'intAry' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
    }
  }

  if intAry.isZeroValue {

    intAry.signVal = 1

    return nil
  }

  if intAry.signVal < 1 {

    intAry.signVal = 1

  } else {

    intAry.signVal = -1

  }

  return nil
}

// divideByTwo
//
//	Divides the numeric value of the input parameter 'intAry' by 2.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) divideByTwo(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.divideByTwo()",
    "")

  if err != nil {
    return err
  }

  if intAry == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intAry'",
    }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "   intAry, ePrefix.XCpy(Validating 'intAry').String())",
        ErrContext: "Input parameter 'intAry' is INVALID!\n" +
          "'intAry' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
    }
  }

  err = new(IntAryMathDivide).DivideByTwo(intAry)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathDivide).DivideByTwo(ia)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// divideByInt64
//
//		Divide the current value of the IntAry parameter 'intAry' by an
//		int64 'divisor'	parameter passed to the method. The result or
//	 'quotient' produced by this divsion operation is stored in the
//	 original 'intAry' object.
//
//		If the quotient has a number of decimal places to the right of
//		the decimal point which is greater than 'maxPrecision', the
//		result is rounded to 'maxPrecision' decimal places.
//
//		If 'maxPrecision' is set equal to -1, 'maxPrecision' is
//		automatically set to 4,096.
//
//		If 'maxPrecision' is less than -1, an error will be returned.
func (iaNeutron *intAryNeutron) divideByInt64(
  intAry *IntAry,
  validateIntAry bool,
  divisor int64,
  maxPrecision int,
  errPrefDto *ePref.ErrPrefixDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.divideByInt64()",
    "")

  if err != nil {
    return err
  }

  if intAry == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intAry'",
    }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry,
      ePrefix.XCpy("Validating intAry").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "  intAry, ePrefix.XCpy(Validating 'intAry').String())",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  } else {

    err = new(intAryNanobot).setInternalFlags(
      intAry, ePrefix.XCpy("Setting 'intAry' Flags"))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
          "  ia, ePrefix.XCpy(Setting 'intAry' Flags))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

  }

  intAryNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "intAryNumStr, err := new(intAryAtom).getRawNumStr(\n" +
        "  getRawNumStr(intAry, false, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathDivide).DivideByInt64(intAry, divisor, maxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathDivide).DivideByInt64(\n" +
        "intAry, divisor, maxPrecision)",
      ErrContext: fmt.Sprintf("intAry = '%v'\ndivisor = '%v'\nmaxPrecision= '%v'",
        intAryNumStr, divisor, maxPrecision),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// divideByTenToPower
//
//	Divide the numerical value of an IntAry instance by
//	10 raised to the power of the input parameter, 'exponent'.
//
//	             'ia'
//	    ia =  -------------
//	         (10^exponent)
//
//	The result or quotient is stored in the IntAry instance passed
//  as input parameter 'intAry'.
func (iaNeutron *intAryNeutron) divideByTenToPower(
  intAry *IntAry,
  validateIntAry bool,
  exponent uint,
  errPrefDto *ePref.ErrPrefixDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.divideByTenToPower()",
    "")

  if err != nil {
    return err
  }

  if intAry == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intAry'",
    }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry,
      ePrefix.XCpy("Validating intAry").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "  intAry, ePrefix.XCpy(Validating 'intAry').String())",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  } else {

    err = new(intAryNanobot).setInternalFlags(
      intAry, ePrefix.XCpy("Setting 'intAry' Flags"))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
          "  ia, ePrefix.XCpy(Setting 'intAry' Flags))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  intAryNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "intAryNumStr, err := new(intAryAtom).getRawNumStr(\n" +
        "  getRawNumStr(intAry, false, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathDivide).DivideByTenToPower(intAry, exponent)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathDivide).DivideByTenToPower(intAry, exponent)",
      ErrContext: fmt.Sprintf("intAry = '%v'\nexponent= '%v'",
        intAryNumStr, exponent),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// getBigInt
//
//	Returns the current value of this intAry object expressed
//	as a signed integer number of type *big.Int.
func (iaNeutron *intAryNeutron) getBigInt(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

  if iaNeutron.lock == nil {
    iaNeutron.lock = new(sync.Mutex)
  }

  iaNeutron.lock.Lock()

  defer iaNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getBigInt()",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  if intAry == nil {

    return big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return big.NewInt(0),
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "   intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "Input parameter 'intAry' is INVALID!\n" +
            "'intAry' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  lenIntAry := len(intAry.intAry)

  if lenIntAry != intAry.intAryLen {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: " if lenIntAry != ia.intAryLen",
        ErrMessage: "Error: The actual length of 'intAry' does not match ia.intAryLen.\n" +
          "This instance of 'ia' is INVALID!",
      }

  }

  if lenIntAry == 0 {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: " if lenIntAry != ia.intAryLen",
        ErrMessage: "Error: The actual length of 'intAry' is ZERO.\n" +
          "This instance of 'ia' is INVALID!",
      }
  }

  result := big.NewInt(0).SetInt64(0)

  big10 := big.NewInt(0).SetInt64(10)

  for i := 0; i < intAry.intAryLen; i++ {
    result = big.NewInt(0).Mul(result, big10)
    result = big.NewInt(0).Add(result, big.NewInt(0).SetInt64(int64(intAry.intAry[i])))

  }

  if intAry.signVal == -1 {

    result = big.NewInt(0).Neg(result)
  }

  return result, nil
}

// getFractionalDigits
//
//	Examines the current IntAry instanace and returns a new IntAry
//	object consisting of the fractional digits to the right of the
//	decimal point from the original, current IntAry object.
//
//	Note: The sign Value of the returned int Ary is always
//	positive or +1.
//
//	The returned IntAry instance will display fractional digits
//	with a leading integer digit of zero. Example '0.5678'
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getFractionalDigits(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getFractionalDigits()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if intAry == nil {

    return IntAry{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "   intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "Input parameter 'intAry' is INVALID!\n" +
            "'intAry' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, ePrefix.XCpy("Numeric Separators intAry -> numSeps"))

  if err != nil {
    return IntAry{}, err
  }

  iAry2 := new(intAryElectron).newIntAry()

  err = new(intAryQuark).setIntAryToZero(
    &iAry2, 0, numSeps, ePrefix.XCpy("Setting iAry2 to Zero"))

  if err != nil {
    return IntAry{}, err
  }

  if intAry.precision == 0 {
    return iAry2, nil
  }

  fracIdx := intAry.intAryLen - intAry.precision

  iAry2.intAry = make([]uint8, intAry.precision+1)
  idx := 1

  for i := fracIdx; i < intAry.intAryLen; i++ {

    iAry2.intAry[idx] = intAry.intAry[i]

    idx++
  }

  iAry2.precision = intAry.precision

  iAry2.signVal = 1

  err = new(intAryNanobot).setInternalFlags(
    &iAry2, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setInternalFlags(&iAry2, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }
  return iAry2, nil
}

// getIntegerDigits
//
//	Examines the current intAry object and returns a new intAry
//	consisting of only the integer digits to the left of the
//	decimal point in the current intAry object.
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
//	The IntAry object returned by this method will be configured
//	with Numeric Separators copied from the original instance of
//	IntAry, 'intAry'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getIntegerDigits(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getIntegerDigits()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if intAry == nil {

    return IntAry{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "   intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "Input parameter 'intAry' is INVALID!\n" +
            "'intAry' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, ePrefix.XCpy("Numeric Separators intAry -> numSeps"))

  if err != nil {
    return IntAry{}, err
  }

  iAry2 := new(intAryElectron).newIntAry()

  if intAry.isZeroValue {

    err = new(intAryQuark).setIntAryToZero(
      &iAry2, 0, numSeps, ePrefix.XCpy("Set iAry2 to Zero"))

    return iAry2, err
  }

  intLen := intAry.intAryLen - intAry.precision

  iAry2.intAry = make([]uint8, intLen)

  for i := 0; i < intLen; i++ {
    iAry2.intAry[i] = intAry.intAry[i]
  }

  iAry2.signVal = intAry.signVal
  iAry2.precision = 0

  err = new(intAryNanobot).setInternalFlags(
    &iAry2, ePrefix.XCpy("Setting Flags on 'iAry2"))

  if err != nil {
    return IntAry{}, err
  }

  if iAry2.isZeroValue {
    iAry2.signVal = 1
  }

  return iAry2, nil
}

// getNumStrDto
//
//	 Converts the IntAry input parameter ('') to a returned instance
//	 of NumStrDto.
//
//	 The returned NumStrDto will contain numeric separators
//	 (decimal separator, thousands separator and currency symbol)
//	 copied from the current IntAry instance.
//
//		Validation Testing
//		==================
//
//		If input parameter 'validateIntAry' is set to true, this
//		method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getNumStrDto(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  if iaNeutron.lock == nil {
    iaNeutron.lock = new(sync.Mutex)
  }

  iaNeutron.lock.Lock()

  defer iaNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getNumStrDto()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if intAry == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "  intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  numSeps, err := new(intAryPhoton).
    getNumericSeparatorsDto(intAry, ePrefix.XCpy("numSeps<-intAry"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := new(intAryNanobot).getNumericSeparatorsDto(intAry, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // iaNumStr, err := ia.GetNumStr()
  iaNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix.XCpy("iaNumStr<-intAry"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "iaNumStr, err := new(intAryAtom).getRawNumStr(\n" +
          "  intAry, false, ePrefix.XCpy(iaNumStr<-intAry))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nDto, err := new(NumStrDto).NewNumStrWithNumSeps(iaNumStr, numSeps)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nDto, err := new(NumStrDto).NewNumStrWithNumSeps(iaNumStr, numSeps)",
        ErrContext: fmt.Sprintf("iaNumStr= '%v'\nnumSeps= '%v'", iaNumStr, numSeps),
        ErrMessage: err.Error(),
      }
  }

  return nDto, nil
}

// setSign
//
//	Used to change the sign value of the intAry  object passed as
//	input parameter 'intAry'. The new sign value will be set
//	according the input parameter, 'signVal'.
//
//	'signVal' has only two valid values, -1 or +1. If
//	any value other than -1 or +1 is detected, an error
//	will be returned.
func (iaNeutron *intAryNeutron) setSign(
  intAry *IntAry,
  signVal int,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaNeutron.lock == nil {
    iaNeutron.lock = new(sync.Mutex)
  }

  iaNeutron.lock.Lock()

  defer iaNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.equal()",
    "")

  if err != nil {
    return err
  }

  if intAry == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intAry'",
    }
  }

  if signVal != -1 && signVal != 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: fmt.Sprintf("'signVal' == %d", signVal),
      ErrMessage: "Error: Input parameter 'signVal' is INVALID.\n" +
        "'signVal' must be either -1 or +1.",
    }

  }

  err = new(intAryNanobot).setInternalFlags(
    intAry, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(intAry, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if intAry.isZeroValue {
    return nil
  }

  intAry.signVal = signVal

  return nil
}
