package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryNanobot struct {
  lock sync.Mutex
}

// getNumericSeparatorsDto
//
//	 Receives a pointer to IntAry and extracts the Numeric
//	 Separators. These separators are consolidated and returned as
//	 a NumericSeparatorDto structure containing the character or
//	 'rune' values for decimal point separator, thousands
//	 separator and currency symbol.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//		them into numeric values.
func (iaNanobot *intAryNanobot) getNumericSeparatorsDto(
  intAry *IntAry,
  errPrefDto *ePref.ErrPrefixDto) (NumericSeparatorDto, error) {

  iaNanobot.lock.Lock()

  defer iaNanobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNanobot.getNumericSeparatorsDto()",
    "")

  if err != nil {
    return NumericSeparatorDto{}, err
  }

  if intAry == nil {

    return NumericSeparatorDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  numSeps := NumericSeparatorDto{}
  numSeps.DecimalSeparator = intAry.GetDecimalSeparator()
  numSeps.ThousandsSeparator = intAry.GetThousandsSeparator()
  numSeps.CurrencySymbol = intAry.GetCurrencySymbol()

  return numSeps, nil
}

// hasFractionalDigits
//
//	This method examines the current intAry object to determine if
//	there are non-zero digits to the right of the decimal place. If
//	all digits to the right of the decimal place are zero, this
//	method returns 'false'
//
//	If non-zero digits are present to the right of the decimal
//	place, the method returns 'true'.
func (iaNanobot *intAryNanobot) hasFractionalDigits(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (bool, error) {

  iaNanobot.lock.Lock()

  defer iaNanobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNanobot.hasFractionalDigits",
    "")

  if err != nil {
    return false, err
  }

  if validateIntAry {
    err = new(intAryElectron).isValidIntAry(
      intAry,
      ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return false,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "  intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  if intAry.precision == 0 {
    return false, nil
  }

  err = new(intAryElectron).setIntAryLength(
    intAry, ePrefix.XCpy("Setting 'intAry' IntAry Length"))

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).setIntAryLength(\n" +
          "intAry, ePrefix.XCpy(Setting 'intAry' IntAry Length))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  intLen := intAry.intAryLen - intAry.precision

  if intLen < 1 {
    return false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "intLen < 1",
        ErrMessage: fmt.Sprintf("Error - Int Array integer length is less than 1.\n"+
          "intLen= '%v'", intLen),
      }
  }

  for i := intLen; i < intAry.intAryLen; i++ {

    if intAry.intAry[i] > 0 {

      return true, nil
    }
  }

  return false, nil
}

// setIntAryUint64Exponent
//
//	Returns a new IntAry instance based on input parameters,
//	'uint64Num', 'signValue' and 'exponent'.
//
//	The returned IntAry numeric value is set using an uint64 value
//	multiplied by 10 raised to the power of the 'exponent' parameter.
//
//		    Result Numeric Value = uint64 X 10^exponent
//
//	Usage
//	=====
//
//	This method is may be used with the 'new' keyword syntax.
//
//	  iAry := new(IntAry).NewUint64Exponent(123456, -3)
//	  -- iAry is now equal to "123.456", precision = 3
//
//	  iAry := new(IntAry).NewUint64Exponent(123456, 3)
//	  -- iAry is now equal to "123456.000", precision = 3
//
//	Examples
//	========
//
//	uint64Num      exponent      IntAry Result
//
//	  123456          -3              123.456
//	  123456           3           123456.000
//	  123456           0           123456
//
//	Input Parameters
//	================
//
//	uint64Num               uint64
//	  The uint64 holds the numeric digits which will make up the
//	  returned IntAry numeric value.
//
//	signValue                int
//	  This parameter must be set to one of two possible values:
//	  +1 or -1.
//
//	  Final numeric values less than zero must be tagged with
//	  signValue= -1.
//
//	  Final numeric values greater than or equal to zero must be
//	  tagged with signValue= +1.
//
//	exponent                 int
//	  This value will be used to determine the numeric digits in
//	  'uint64Num' which will be assigned to the right of the
//	  decimal point in the final calculation result returned as
//	  IntAry instance.
//
//	Return Values
//	=============
//
//	IntAry
//	  This returned IntAry object will be configured with the
//	  numeric value computed from input parameters 'uint64Num',
//	  signValue and 'exponent'.
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (iaNanobot *intAryNanobot) setIntAryUint64Exponent(
  ia *IntAry,
  numSepsSrcIntAry *IntAry,
  nsProfile NumSepsProfileSelection,
  uint64Num uint64,
  signValue int,
  exponent int,
  validateResult bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaNanobot.lock.Lock()

  defer iaNanobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNanobot.setInt64Exponent()",
    "")

  if err != nil {
    return err
  }

  if ia == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'ia'",
    }
  }

  uint64Ten := uint64(10)

  if exponent > 0 {
    for i := 0; i < exponent; i++ {

      uint64Num *= uint64Ten

    }
  }

  if exponent < 0 {

    exponent = exponent * -1

  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryGluon).setIntAryWithUint64(
    &iAry,
    numSepsSrcIntAry,
    nsProfile,
    uint64Num,
    signValue,
    uint(exponent),
    validateResult,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryGluon).setIntAryWithUint64(\n" +
        "  &iAry, ia, nsProfile, uint64Num, signValue, uint(exponent),\n" +
        "  validateResult, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// setInternalFlags
//
//	 Sets Array Lengths and test for zero values
//
//		IMPORTANT
//		=========
//
//		The calling function is responsible for verifying the validity
//		of 'ia', the IntAry object.
func (iaNanobot *intAryNanobot) setInternalFlags(
  ia *IntAry,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaNanobot.lock.Lock()

  defer iaNanobot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNanobot.setInternalFlags()",
    "")

  if err != nil {
    return err
  }

  if ia == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'ia'",
    }
  }

  err = new(intAryElectron).setSignificantDigitIdxs(
    ia,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryElectron).\n" +
        "  setSignificantDigitIdxs( ia, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// setInternalFlagsNoErrors
//
//	This method differs from 'intAryNanobot.setInternalFlags' in
//	that this method does NOT return an error.
func (iaNanobot *intAryNanobot) setInternalFlagsNoErrors(
  ia *IntAry) {

  iaNanobot.lock.Lock()

  defer iaNanobot.lock.Unlock()

  new(intAryElectron).setSignificantDigitIdxsNoErrors(ia)

  return
}

// setIsZeroValue
//
//	Analyzes the value of the intAry and sets a flag if the value
//	of intAry evaluates to zero.
func (iaNanobot *intAryNanobot) setIsZeroValue(
  ia *IntAry) {

  iaNanobot.lock.Lock()

  defer iaNanobot.lock.Unlock()

  if ia == nil {
    return
  }

  ia.intAryLen = len(ia.intAry)

  ia.isZeroValue = true

  ia.isIntegerZeroValue = true

  intLen := ia.intAryLen - ia.precision

  for i := 0; i < ia.intAryLen; i++ {

    if i < intLen && ia.intAry[i] > 0 {

      ia.isIntegerZeroValue = false
    }

    if ia.intAry[i] > 0 {

      ia.isZeroValue = false

      return
    }
  }

  // ia.isZeroValue == true
  // signVal must be 1
  ia.signVal = 1

  return
}
