package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryMinibot struct {
  lock sync.Mutex
}

// setIntAryInt64Exponent
//
//	Reconfigures an IntAry instance passed in as input parameter
//	'ia'. The new numeric value is calculated using an int64 value
//	multiplied by 10 raised to the power of the 'exponent' parameter.
//
//	    numeric value = int64 X 10^exponent
//
//	Examples
//	========
//
//		uint64Num      exponent      IntAry Result
//
//		  123456          -3           123.456
//		  123456           3           123456.000
//		  123456           0           123456
//
//	IMPORTANT
//	=========
//
//	In practice, the maximum limits for 'int64Num' and
//	'exponent' will be constrained by the maximum array size
//	permitted by your system. Type IntAry relies on arrays of
//	8-bit integers for numeric value storage.
//
//	Input Parameters
//	================
//
//	ia                       *IntAry
//	  This instance of IntAry will be overwritten with the new
//	  IntAry values computed by this method.
//
//	numSepsSrcIntAry         *IntAry
//	  This instance of IntAry may be populated by the calling
//	  function as a source of Numeric Separators
//
//	nsProfile                NumSepsProfileSelection
//	  This struct contains decision parameters for selecting and
//	  configuring Numeric Separators.
//
//	int64Num                 int64
//	  The numeric digits contained in this value comprise both
//		 the integer digits and the fractional digits which will be
//		 configured in the final numeric value stored in the IntAry
//		 object returned by this method.
//
//	exponent                 int
//	  'exponent' specifies the number of fractional digits in the
//		 final numeric value stored in the returned IntAry object.
//	  See the examples above.
//
//	validateResult           bool
//	  When set to 'true', the final numeric value calculated by
//	  this method will be subjected to validation testing.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//		 value will be set to 'nil'
func (iaMinibot *intAryMinibot) setIntAryInt64Exponent(
  ia *IntAry,
  numSepsSrcIntAry *IntAry,
  nsProfile NumSepsProfileSelection,
  int64Num int64,
  exponent int,
  validateResult bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaMinibot.lock.Lock()

  defer iaMinibot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMinibot.setIntAryInt64Exponent()",
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
      ErrContext: "",
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

  if exponent > 0 {
    for i := 0; i < exponent; i++ {

      int64Num *= 10

    }
  }

  if exponent < 0 {

    exponent = exponent * -1

  }

  err = new(intAryGluon).setIntAryWithInt64(
    ia, nil, nsProfile2, int64Num, uint(exponent), validateResult, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryGluon).setIntAryWithInt64(\n" +
        "  ia, nil, nsProfile2, int64Num, uint(exponent), validateResult, ePrefix)",
      ErrContext: fmt.Sprintf("validateResult = '%v'", validateResult),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// setIntAryUint64Exponent
//
//	Receives an IntAry instance ('ia') and reconfigures the
//	numeric values based on input parameters 'uint64Num',
//	'signValue' and 'exponent'.
//
//	The IntAry numeric value is configured using an uint64 value
//	multiplied by 10 raised to the power of the 'exponent'
//	parameter.
//
//		    Result Numeric Value = uint64 X 10^exponent
//
//	Be advised that the original value of input parameter 'ia'
//	will be overwritten and destroyed. The final calculation
//	result will be stored in the 'ia' IntAry instance.
//
//	Examples
//	========
//
//	uint64Num      exponent      IntAry Result
//
//	  123456          -3           123.456
//	  123456           3           123456.000
//	  123456           0           123456
//
//	IMPORTANT
//	=========
//
//	In practice, the maximum limits for 'uint64Num' and
//	'exponent' will be constrained by the maximum array size
//	permitted by your system. Type IntAry relies on arrays of
//	8-bit integers for numeric value storage.
//
//	Input Parameters
//	================
//
//	ia                       *IntAry
//	  This instance of IntAry will be overwritten with the new
//	  IntAry values computed by this method.
//
//	numSepsSrcIntAry         *IntAry
//	  This instance of IntAry may be populated by the calling
//	  function as a source of Numeric Separators
//
//	nsProfile                NumSepsProfileSelection
//	  This struct contains decision parameters for selecting and
//	  configuring Numeric Separators.
//
//	uint64Num                uint64
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
//	validateResult           bool
//	  When set to 'true', the final numeric value calculated by
//	  this method will be subjected to validation testing.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (iaMinibot *intAryMinibot) setIntAryUint64Exponent(
  ia *IntAry,
  numSepsSrcIntAry *IntAry,
  nsProfile NumSepsProfileSelection,
  uint64Num uint64,
  signValue int,
  exponent int,
  validateResult bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaMinibot.lock.Lock()

  defer iaMinibot.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMinibot.setInt64Exponent()",
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
      ErrContext: "",
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

  uint64Ten := uint64(10)

  if exponent > 0 {
    for i := 0; i < exponent; i++ {

      uint64Num *= uint64Ten

    }
  }

  if exponent < 0 {

    exponent = exponent * -1

  }

  err = new(intAryGluon).setIntAryWithUint64(
    ia,
    nil,
    nsProfile2,
    uint64Num,
    signValue,
    uint(exponent),
    validateResult,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryGluon).setIntAryWithUint64(\n" +
        "  &iAry, ia, nsProfile2, uint64Num, signValue, uint(exponent),\n" +
        "  validateResult, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
