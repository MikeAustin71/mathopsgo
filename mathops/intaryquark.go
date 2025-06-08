package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryQuark struct {
  lock *sync.Mutex
}

// compareAbsoluteValues
//
//	Compares the absolute values of two IntAry instances.
//
//	Returns:
//	 0  = Current IntAry value is equal to the passed IntAry value.
//	 1  = Current IntAry value is greater than the passed IntAry value.
//	-1  = Current IntAry value is less than the passed IntAry value.
func (iaQuark *intAryQuark) compareAbsoluteValues(
  intAry1 *IntAry,
  validateIntAry1 bool,
  intAry2 *IntAry,
  validateIntAry2 bool,
  errPrefDto *ePref.ErrPrefixDto) (int, error) {

  if iaQuark.lock == nil {
    iaQuark.lock = new(sync.Mutex)
  }

  iaQuark.lock.Lock()

  defer iaQuark.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryQuark.compareAbsoluteValues()",
    "")

  if err != nil {
    return -1, err
  }

  if intAry1 == nil {

    return -1,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry1'",
      }
  }

  if intAry2 == nil {

    return -1,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry2'",
      }
  }

  iAryElectron := new(intAryElectron)
  iAryNanobot := new(intAryNanobot)

  if validateIntAry1 {

    err = iAryElectron.isValidIntAry(intAry1, ePrefix.XCpy("Validating 'intAry1'").String())

    if err != nil {

      return -1,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = iAryElectron.isValidIntAry(\n" +
            "  intAry1, ePrefix.XCpy(Validating 'intAry1').String())",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  } else {

    err = iAryNanobot.setInternalFlags(
      intAry1, ePrefix.XCpy("Setting 'intAry1' Flags"))

    if err != nil {

      return -1,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = iAryNanobot.setInternalFlags(\n" +
            "  intAry1, ePrefix.XCpy(Setting 'intAry1' Flags))",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  if validateIntAry2 {

    err = iAryElectron.isValidIntAry(intAry2, ePrefix.XCpy("Validating 'intAry2'").String())

    if err != nil {

      return -1,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = iAryElectron.isValidIntAry(\n" +
            "  intAry2, ePrefix.XCpy(Validating 'intAry2').String())",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  } else {

    err = iAryNanobot.setInternalFlags(
      intAry2, ePrefix.XCpy("Setting 'intAry2' Flags"))

    if err != nil {

      return -1,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = iAryNanobot.setInternalFlags(\n" +
            "  intAry2, ePrefix.XCpy(Setting 'intAry2' Flags))",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  err = iAryElectron.setIntAryLength(intAry1, ePrefix.XCpy("Setting 'ia' IntAry Length"))

  if err != nil {

    return -1,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = iAryElectron.setIntAryLength(\n" +
          "ia, ePrefix.XCpy(Setting 'ia' IntAry Length))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = iAryElectron.setIntAryLength(intAry2, ePrefix.XCpy("Setting 'iAry2' IntAry Length"))

  if err != nil {

    return -1,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = iAryElectron.setIntAryLength(\n" +
          "iAry2, ePrefix.XCpy(Setting 'iAry2' IntAry Length))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // Not needed. Validation sets internal flags
  //iAry2.SetIsZeroValue()
  //ia.SetIsZeroValue()

  if intAry1.isZeroValue && intAry2.isZeroValue {
    return 0, nil
  }

  iaIntLen := intAry1.intAryLen - intAry1.precision

  iAry2IntLen := intAry2.intAryLen - intAry2.precision

  // Integer Lengths are Equal
  if iaIntLen == iAry2IntLen {
    for i := 0; i < iaIntLen; i++ {
      if intAry1.intAry[i] > intAry2.intAry[i] {
        return 1, nil
      }

      if intAry2.intAry[i] > intAry1.intAry[i] {
        return -1, nil
      }
    }
  }

  deltaStartIdx := 0

  // ia Integer Length is Greater than IAry2 Integer Length
  if iaIntLen > iAry2IntLen {
    deltaStartIdx = iaIntLen - iAry2IntLen

    for j := 0; j < iaIntLen; j++ {

      if j < deltaStartIdx {

        if intAry1.intAry[j] > 0 {
          return 1, nil
        }

      } else {
        // 'i' must be >= deltaStartIdx

        if intAry1.intAry[j] > intAry2.intAry[j-deltaStartIdx] {
          return 1, nil
        }

        if intAry2.intAry[j-deltaStartIdx] > intAry1.intAry[j] {
          return -1, nil
        }
      }
    }
  }

  // iAry2 Integer Length is Greater Than ia Integer Length
  if iAry2IntLen > iaIntLen {
    deltaStartIdx = iAry2IntLen - iaIntLen

    for k := 0; k < iAry2IntLen; k++ {

      if k < deltaStartIdx {
        if intAry2.intAry[k] > 0 {
          return -1, nil
        }

      } else {
        // 'i' must be >= deltaStartIdx

        if intAry2.intAry[k] > intAry1.intAry[k-deltaStartIdx] {
          return -1, nil
        }

        if intAry1.intAry[k-deltaStartIdx] > intAry2.intAry[k] {
          return 1, nil
        }
      }
    }
  }

  // If precision is zero, the intAry's are equivalent
  if intAry1.precision == 0 && intAry2.precision == 0 {
    return 0, nil
  }

  // Integer Values are Equivalent. Now test
  // digits to the right of the decimal point.

  // Test fractional digits to right of decimal point
  iaFracIdx := iaIntLen
  iAry2FracIdx := iAry2IntLen
  // Test for case of Equal precision
  if intAry1.precision == intAry2.precision {
    for m := 0; m < intAry1.precision; m++ {

      if intAry1.intAry[iaFracIdx] > intAry2.intAry[iAry2FracIdx] {
        return 1, nil
      }

      if intAry2.intAry[iAry2FracIdx] > intAry1.intAry[iaFracIdx] {
        return -1, nil
      }

      iaFracIdx++
      iAry2FracIdx++
    }
  }

  iaFracIdx = iaIntLen
  iAry2FracIdx = iAry2IntLen
  // Test for case where ia precision Greater than iAry2 precision
  if intAry1.precision > intAry2.precision {

    for i := 0; i < intAry1.precision; i++ {

      if i < intAry2.precision {

        if intAry1.intAry[iaFracIdx] > intAry2.intAry[iAry2FracIdx] {
          return 1, nil
        }

        if intAry2.intAry[iAry2FracIdx] > intAry1.intAry[iaFracIdx] {
          return -1, nil
        }

        iaFracIdx++
        iAry2FracIdx++

      } else {
        if intAry1.intAry[iaFracIdx] > 0 {
          return 1, nil
        }

        iaFracIdx++
      }
    }
  }

  iaFracIdx = iaIntLen
  iAry2FracIdx = iAry2IntLen
  // Test for case where iAry2 precision Greater than ia precision
  if intAry2.precision > intAry1.precision {

    for i := 0; i < intAry2.precision; i++ {

      if i < intAry1.precision {

        if intAry1.intAry[iaFracIdx] > intAry2.intAry[iAry2FracIdx] {
          return 1, nil
        }

        if intAry2.intAry[iAry2FracIdx] > intAry1.intAry[iaFracIdx] {
          return -1, nil
        }

        iaFracIdx++
        iAry2FracIdx++

      } else {
        if intAry2.intAry[iAry2FracIdx] > 0 {
          return -1, nil
        }

        iAry2FracIdx++
      }
    }

  }

  // The two absolute numeric values must be equal
  return 0, nil
}

// SetIntAryToOne
//
//	Sets the value of the intAry object to one ('1').
func (iaQuark *intAryQuark) setIntAryToOne(
  ia *IntAry,
  numSepsSrcIntAry *IntAry,
  nsProfile NumSepsProfileSelection,
  precision int,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaQuark.lock == nil {
    iaQuark.lock = new(sync.Mutex)
  }

  iaQuark.lock.Lock()

  defer iaQuark.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryQuark.setIntAryToZero()",
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

  if precision < 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is INVALID!\n"+
        "'precision' is less than ZERO!\n"+
        "precision= '%v'", precision),
    }
  }

  var numSeps NumericSeparatorDto

  nsProfile.OutputNumSepsName = "numSeps"

  var actualNumSepsSrcIntAryPtr *IntAry

  if numSepsSrcIntAry == nil {

    nsProfile.SourceObjectName = "ia"
    actualNumSepsSrcIntAryPtr = ia

  } else {

    nsProfile.SourceObjectName = "numSepsSrcIntAry"
    actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
  }

  numSeps, err = new(intAryUtility).selectNumericSeparators(
    actualNumSepsSrcIntAryPtr,
    nsProfile,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
        "actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
      ErrContext: fmt.Sprintf("nsProfile.SourceObjectName= '%v'",
        nsProfile.SourceObjectName),
      ErrMessage: err.Error(),
    }
  }

  nsProfile2 := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          numSeps,
  }

  return new(intAryGluon).setIntAryWithInt(
    ia,
    nil,
    nsProfile2,
    1,
    uint(precision),
    true,
    ePrefix.XCpy("Set 'ia' = 1"))
}

// setIntAryToZero
//
//			Sets the value of the intAry object to zero ('0').
//
//		Input Parameters
//		================
//
//		intAry                   *IntAry
//		  A pointer to an IntAry object. This object will be
//		  reconfigured with a zero value.
//
//		numSepsSrcIntAry         *IntAry
//		  If this pointer is NOT 'nil', the Numeric Separators
//	   will be taken from this IntAry Object.
//
//	   If this pointer is 'nil', it will be ignored and
//	   the source of Numeric Separators will either be
//	   the 'intAry' object or standard defaults as specified
//	   by input parameter 'nsProfile'.
//		  reconfigured with a new value based on the following
//		  input parameters.
//
//		nsProfile                NumSepsProfileSelection
//		 This struct contains all the prameters and options
//		 necessary to generate the NumericSeparatorsDto which is
//		 required for configuration of Numeric Separators in the
//		 IntAry object returned by this method.
//
//		intDigits                int
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in parameter,
//		  'intAry'.
//
//		signVal                  int
//		  Input parameter 'signVal' must be set to one of two values:
//		  +1 or -1. This value is used to signal the sign of the
//		  resulting numeric value. +1 identifies a positive number and
//		  -1 identifies a negative number. 'signVal' determines the
//		  numeric sign of the resulting IntAry value, either plus or
//		  minus.
//
//		precision                uint
//		  'precision' specifies the number of fractional digits in the
//		  final numeric value stored in 'intAry'
//
//		  Although 'precision' is an unsigned integer type, the maximum
//		  value allowed for this parameter is 2,147,483,647.
//
//		Return Values
//		=============
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (iaQuark *intAryQuark) setIntAryToZero(
  intAry *IntAry,
  numSepsSrcIntAry *IntAry,
  nsProfile NumSepsProfileSelection,
  precision uint,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaQuark.lock == nil {
    iaQuark.lock = new(sync.Mutex)
  }

  iaQuark.lock.Lock()

  defer iaQuark.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryQuark.setIntAryToZero()",
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

  var numSeps NumericSeparatorDto

  nsProfile.OutputNumSepsName = "numSeps"

  var actualNumSepsSrcIntAryPtr *IntAry

  if numSepsSrcIntAry == nil {

    nsProfile.SourceObjectName = "intAry"
    actualNumSepsSrcIntAryPtr = intAry

  } else {

    nsProfile.SourceObjectName = "numSepsSrcIntAry"
    actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
  }

  numSeps, err = new(intAryUtility).selectNumericSeparators(
    actualNumSepsSrcIntAryPtr,
    nsProfile,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "finalNumSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
        "actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  intAry.intAryLen = 1 + int(precision)

  intAry.precision = int(precision)

  intAry.intAry = make([]uint8, intAry.intAryLen)

  intAry.signVal = 1

  err = new(intAryPhoton).setNumericSeparatorsDto(
    intAry, numSeps, false, ePrefix.XCpy("Setting 'intAry'"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryMechanics).\n" +
        "  setNumericSeparatorsDto(intAry, ePrefix.String())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(intAryNanobot).setInternalFlags(
    intAry, ePrefix.XCpy("Setting 'intAry' Flags"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).\n" +
        "  setInternalFlags(intAry, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// setIntAryWithNumStr
//
//	Receives a raw number string and sets the fields of the
//	internal intAry structure to the appropriate values.
func (iaQuark *intAryQuark) setIntAryWithNumStr(
  ia *IntAry,
  validateIa bool,
  numSepsSrcIntAry *IntAry,
  nsProfile NumSepsProfileSelection,
  str string,
  validateResult bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaQuark.lock == nil {
    iaQuark.lock = new(sync.Mutex)
  }

  iaQuark.lock.Lock()

  defer iaQuark.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryQuark.setIntAryToZero()",
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

  if len(str) == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "Error: Input parameter 'str' is a zero length number string",
    }

  }

  err = new(intAryUtility).selectIntAryValidation(
    ia,
    "ia",
    validateIa,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
        "ia, 'ia', validateIa= '%v', ePrefix)", validateIa),
      ErrContext: "Valiate on Startup",
      ErrMessage: err.Error(),
    }
  }

  var numSeps NumericSeparatorDto

  nsProfile.OutputNumSepsName = "numSeps"

  var actualNumSepsSrcIntAryPtr *IntAry

  if numSepsSrcIntAry == nil {

    nsProfile.SourceObjectName = "ia"
    actualNumSepsSrcIntAryPtr = ia

  } else {

    nsProfile.SourceObjectName = "numSepsSrcIntAry"
    actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
  }

  numSeps, err = new(intAryUtility).selectNumericSeparators(
    actualNumSepsSrcIntAryPtr,
    nsProfile,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "finalNumSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
        "actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nsProfile2 := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: false,
    ValidateNumSeps:          false,
    OverrideNumSeps:          numSeps,
  }

  new(intAryBoson).emptyBackUp(ia)

  err = new(intAryPhoton).setNumericSeparatorsDto(
    ia, numSeps, true, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryPhoton).setNumericSeparatorsDto(\n" +
        "ia, numSeps, true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  ia.signVal = 1
  baseRunes := []rune(str)
  lBaseRunes := len(baseRunes)
  isStartRunes := false
  isEndRunes := false
  isFractionalValue := false

  for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

    if baseRunes[i] == '+' ||
      baseRunes[i] == ' ' ||
      baseRunes[i] == ia.thousandsSeparator ||
      baseRunes[i] == ia.currencySymbol {

      continue

    }

    if baseRunes[i] == ',' && ia.decimalSeparator != ',' {
      continue
    }

    if isStartRunes == true &&
      isFractionalValue &&
      baseRunes[i] == ia.decimalSeparator {

      continue
    }

    if baseRunes[i] == '-' &&
      isStartRunes == false &&
      i+1 < lBaseRunes &&
      ((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
        baseRunes[i+1] == ia.decimalSeparator) {

      ia.signVal = -1
      isStartRunes = true
      continue

    } else if baseRunes[i] >= '0' && baseRunes[i] <= '9' {

      ia.intAry = append(ia.intAry, uint8(baseRunes[i]-48))
      isStartRunes = true

      if isFractionalValue {
        ia.precision++
      }

    } else if i+1 < lBaseRunes &&
      baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
      baseRunes[i] == ia.decimalSeparator {

      isFractionalValue = true
      continue

    } else if isStartRunes {

      isEndRunes = true

    }
  }

  err = new(intAryNanobot).setInternalFlags(
    ia, ePrefix.XCpy("Setting 'ia' Flags"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
        "  ia, ePrefix.XCpy(Setting 'ia' Flags))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if ia.intAryLen == 0 || ia.isZeroValue {

    err = new(intAryQuark).setIntAryToZero(
      ia,
      nil,
      nsProfile2,
      uint(ia.precision),
      ePrefix)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryToZero(\n"+
          "ia, nil, nsProfile2, uint(ia.precision)= '%v', ePrefix", uint(ia.precision)),
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  err = new(intAryUtility).selectIntAryValidation(
    ia,
    "ia",
    validateResult,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
        "ia, 'ia', validateResult= '%v', ePrefix)", validateResult),
      ErrContext: "Valiate on Exit",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
