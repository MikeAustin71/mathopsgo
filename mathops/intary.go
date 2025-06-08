package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math"
  "math/big"
  "strconv"
)

/*

	IntAry Overview And General Usage:
	==================================

	Source file 'intary.go' contains a structure, 'IntAry', which is designed
 	to perform a variety of math operations on integer strings.

	This Type is capable of performing highly accurate operations on very large
  numbers. For example, the directory 'MikeAustin71/mathopsgo/examples/eulersnumbercalc'
	contains an example which calculates Euler's Number out to 1,000 digits.

	The IntAry Type also has a backup and restore feature.

	Dependencies
	============
	The 'IntAry' object has the following dependency:

	nthroot.go - MikeAustin71/mathopsgo/mathops/nthroot.go


	INumMgr
	========

	The IntAry Type implements the INumMgr interface.

 Source Code Repository:
 =======================
 	https://github.com/MikeAustin71/mathopsgo.git

 Local File:
 ===========
 MikeAustin71\mathopsgo\mathops\intary.txt

*/

// IntAry - Used to perform string
// based numeric math operations.
//
// Dependencies: NthRootOp - nthroot.go
type IntAry struct {
  intAry []uint8 // Storage is right to left.
  // Least significant digit is
  // stored in intAry[0]. Most
  // significant digit is stored
  // at intAry[intAryLen - 1]
  intAryLen              int
  integerLen             int
  significantIntegerLen  int
  significantFractionLen int
  firstDigitIdx          int
  lastDigitIdx           int
  isZeroValue            bool
  isIntegerZeroValue     bool
  precision              int
  signVal                int
  decimalSeparator       rune
  thousandsSeparator     rune
  currencySymbol         rune
  BackUp                 BackUpIntAry
}

// AddIntAryToThis
//
//	Adds the value of intAry parameter ia2 to the value of the
//	current intAry object.
//
//
//	Validation Testing
//	==================
//
//	This method will NOT perform validation tests on the current
//	instance of IntAry ('ia').
//
//	Input Parameters
//	================
//
//	ia2                      *IntAry
//	  The numeric value of this incoming IntAry object will be
//	  subtracted from the numeric valud of the current IntAry
//	  instance.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during proceesing the return
//	  value of this parameter will be set to 'nil'.
func (ia *IntAry) AddIntAryToThis(ia2 *IntAry) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddIntAryToThis",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).addIntAryToThis(ia, false, ia2, true, true, ePrefix)
}

// AddIntToThis
//
//	Adds an integer number to the value of the current IntAry
//	object.
//
//	Example
//	=======
//
//	  num         precision    result
//
//	 946254          3         946.254
//	 946254          0         946254
//	-946254          3        -946.254
//	-946254          0        -946254
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
//	Numeric Separators previous configured for the current IntAry
//	instance ('ia') will not be modified by this method.
//
//	Validation Testing
//	==================
//
//	This method will subject the current instance of IntAry ('ia')
//	to validation tests.
//
//	Input Parameters
//	================
//
//	num                      int
//	  The integer number to be added to the current IntAry object.
//
//	precision                uint
//	  The precision which should be applied to the int64 input
//	  parameter to designate the number of digits to the right of
//	  the decimal point.
//
//	  num = 123456, precision = 3	Result = 123.456
//
//	  The value 123.456 will be added to the current value of the
//	  current IntAry object. Note: If the value of input parameter
//	  'precision' is negative, an error will be returned.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered, this return value will be set
//	  to 'nil'.
func (ia *IntAry) AddIntToThis(num int, precision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddIntToThis",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).addIntToThis(
    ia, true, num, precision, true, ePrefix)
}

// AddInt64ToThis
//
//	Adds an integer (int64) to the value of the current IntAry
//	object.
//
//	Example
//	=======
//
//	                         Result Added
//	                         to Current
//	int64Num    precision      IntAry
//
//	 946254        3           946.254
//	 946254        0           946254
//	-946254        3          -946.254
//	-946254        0          -946254
//
//	Input Parameters
//	================
//
//	int64Num                 int64
//		  The integer number to be added to the current IntAry object.
//
//	precision               uint
//	  The precision which should be applied to the int64 input
//	  parameter to designate the number of digits to the right
//	  of the decimal point. Example:  num = 123456, precision = 3
//	  Result = 123.456 will be added to the current value of the
//	  current IntAry object. If precision is greater than
//	  2,147,483,647 (max int32 value), and error will be returned.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered, this return value will be set
//	  to 'nil'.
func (ia *IntAry) AddInt64ToThis(int64Num int64, precision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddInt64ToThis",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).addInt64ToThis(ia, true, int64Num, precision, true, ePrefix)
}

// AddBigIntToThis
//
//	Adds the value of the *big.Int input parameter, 'num' to the
//	value of the current IntAry object.
//
//	Example
//	=======
//
//	intDigits     precision     	    result
//
//	 946254           3                946.254
//	 946254           0                946254
//	-946254           3               -946.254
//	-946254           0               -946254
//
//	Usage
//	=====
//
//	num := big.NewInt(123456)
//	precision := uint(3)
//	err := ia.AddBigIntToThis(num, precision)
//
//	The result will equal the current value of the IntAry
//	object plus, '123.456'
//
//	Input Parameters
//	================
//
//	num                       *big.Int
//	  The numeric value to be added to the value of the current
//	  IntAry object. Note that 'num' may be positive or negative.
//
//	precision                 uint
//	  'precision' indicates the number of digits to be formatted
//	  to the right of the decimal	place.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered, this return value will be set
//	  to 'nil'.
func (ia *IntAry) AddBigIntToThis(num *big.Int, precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddBigIntToThis",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).addBigIntToThis(
    ia, true, num, precision, true, ePrefix)
}

// AddBigIntNumToThis - Adds the value of the current IntAry
// to the BigIntNum input parameter.
func (ia *IntAry) AddBigIntNumToThis(bINum BigIntNum) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddBigIntNumToThis",
    "")

  if err != nil {
    return err
  }

  ia2, err := new(IntAry).NewBigIntNum(bINum)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "ia2, err := new(IntAry).NewBigIntNum(bINum)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathAdd).RunTotal(ia, &ia2)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathAdd).RunTotal(ia, &ia2)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddFloat32ToThis - Adds the value of input number (float32)
// to the current value of this IntAry object.
//
// Input Parameters:
//
// num float32 		- The number which will be added to the current
//
//	IntAry object.
//
// precision int	-	Input parameter 'precision' is used
//
//	to set the input precision of 'num'.
//
//	Input parameter 'precision' must be set
//	to a number greater than or equal to zero.
//	It may also be set to a value of -1
//	which causes the number to be formatted to
//	the smallest number of digits to right of
//	the decimal point.
//
// Usage:
// num := float32(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// err := ia.AddFloat32ToThis(num, precision)
//
// This usage will result in a value of '123.456'
//
// If precision were set to -1, the resulting value
// would also be '123.456'
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding may occur.
func (ia *IntAry) AddFloat32ToThis(num float32, precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddFloat32ToThis",
    "")

  if err != nil {
    return err
  }

  if precision < -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "precision < -1",
      ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is invalid.\n"+
        "'precision' must be greater than or equal to -1.\n"+
        "'precision'= '%v'",
        precision),
    }
  }

  ia2, err := new(IntAry).NewFloat32(num, precision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "ia2, err := new(IntAry).NewFloat32(num, precision)",
      ErrContext: fmt.Sprintf("num = '%v'  precision= '%v'",
        num, precision),
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathAdd).RunTotal(ia, &ia2)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathAdd).RunTotal(ia, &ia2)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddFloat64ToThis - Adds a floating point number (float64) to the
// current value of this IntAry object.
//
// Input Parameters:
//
//	num float64 - This float64 value will be added to the value of the
//								current IntAry. Positive and negative values are
//								accepted.
//
// precision int -	'precision' is applied to the input parameter
//
//	'num' (float64) to determine the number of
//	digits to the right of the decimal point which
//	will be applied to the resulting value.
//
//	Input parameter 'precision' must be set to a
//	number greater than or equal to zero.  It may
//	also be set to a value of -1 which causes the
//	number to be formatted to the smallest number
//	of digits to right of the decimal point.
//
//	Note: if precision is a positive number and less
//	than the current number of digits to the right
//	of the decimal place, rounding may occur.
func (ia *IntAry) AddFloat64ToThis(num float64, precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddFloat64ToThis",
    "")

  if err != nil {
    return err
  }

  if precision < -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "precision < -1",
      ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is invalid.\n"+
        "'precision' must be greater than or equal to -1.\n"+
        "'precision'= '%v'", precision),
    }
  }

  ia2, err := new(IntAry).NewFloat64(num, precision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "ia2, err := new(IntAry).NewFloat64(num, precision)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathAdd).RunTotal(ia, &ia2)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathAdd).RunTotal(ia, &ia2)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddFloatBigToThis - Adds a *big.Float number to the current
// value of this IntAry object.
//
// Input Parameters:
//
//	num *big.Float	- This *big.Float value will be added to the value
//										of the current IntAry object. Positive and negative
//										values are accepted.
//
// precision int -	'precision' is applied to the input parameter
//
//	'num' (*big.Float) to determine the number of
//	digits to the right of the decimal point which
//	will be applied to the resulting value.
//
//	Input parameter 'precision' must be set to a
//	number greater than or equal to zero.  It may
//	also be set to a value of -1 which causes the
//	number to be formatted to the smallest number
//	of digits to right of the decimal point.
//
//	Note: if precision is a positive number and less
//	than the current number of digits to the right
//	of the decimal place, rounding may occur.
func (ia *IntAry) AddFloatBigToThis(num *big.Float, precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddFloatBigToThis",
    "")

  if err != nil {
    return err
  }

  if precision < -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "precision < -1",
      ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is invalid.\n"+
        "'precision' must be greater than or equal to -1.\n"+
        "'precision'= '%v'", precision),
    }
  }

  ia2, err := new(IntAry).NewFloatBig(num, precision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "ia2, err := new(IntAry).NewFloatBig(num, precision)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathAdd).RunTotal(ia, &ia2)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathAdd).RunTotal(ia, &ia2)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddMultipleToThis - Add the values of  multiple intAry objects to the current
// intAry value.
//
// convertToNumStr - boolean value determines whether the current intAry
//
//	object will convert the intAry value to a number string.
//	Set this parameter to 'false' if this method is called
//	multiple times in order to improve performance.
func (ia *IntAry) AddMultipleToThis(iaMany ...*IntAry) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddMultipleToThis",
    "")

  if err != nil {
    return err
  }

  for idx, iAry := range iaMany {

    err = new(IntAryMathAdd).RunTotal(ia, iAry)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(IntAryMathAdd).RunTotal(ia, iAry)",
        ErrContext: fmt.Sprintf("Error occurred on cycle= '%v'", idx),
        ErrMessage: err.Error(),
      }
    }

  }

  return nil
}

// AddArrayLengthLeft
//
//	 Adds leading zeros to the internal storage array holding the
//		numeric value for the current instance of IntAry.
func (ia *IntAry) AddArrayLengthLeft(addLen int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddArrayLengthLeft",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).addArrayLengthLeft(ia, true, addLen, true, ePrefix)
}

// AddArrayLengthRight
//
//	Adds trailing zeros to the right of the current intAry.
func (ia *IntAry) AddArrayLengthRight(addLen int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AddArrayLengthRight",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).addArrayLengthRight(
    ia, true, addLen, true, ePrefix)
}

// AppendToIntAry - appends an integer of
// type uint8 to the internal Integer Array
// of the current IntAry object.
func (ia *IntAry) AppendToIntAry(num uint8) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.AppendToIntAry()",
    "")

  if err != nil {
    return err
  }

  ia.intAry = append(ia.intAry, num)

  ia.intAryLen = len(ia.intAry)

  ia.integerLen = ia.intAryLen - ia.precision

  if num > 0 {
    ia.isZeroValue = false
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

  return nil
}

// Ceiling
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
//	with Numeric Separators copied from the original, current
//	IntAry instance, 'ia'.
//
//	Validation Testing
//	==================
//
//	This method will subject the current instance of IntAry ('ia')
//	to validation tests.
func (ia *IntAry) Ceiling() (IntAry, error) {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.Ceiling()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  return new(intAryNeutron).ceiling(ia, true, true, ePrefix)
}

// ChangeSign
//
//	Changes the sign of the current IntAry instance.
//
//	If the current IntAry numeric value is positive (+),
//	this method will change the sign value to negative (-).
//
//	Conversely, if the current IntAry is a negative (-)
//	numeric value, this method will change the sign to
//	positive (+).
func (ia *IntAry) ChangeSign() error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.ChangeSign",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).changeSign(ia, true, ePrefix)
}

// CompareSignedValues
//
//	Compares two IntAry signed numeric values.
//
//	Returns:
//	 0  = Current IntAry value is equal to the passed IntAry value.
//	 1  = Current IntAry value is greater than the passed IntAry value.
//	-1  = Current IntAry value is less than the passed IntAry value.
func (ia *IntAry) CompareSignedValues(iAry2 *IntAry) (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CompareSignedValues()",
    "")

  if err != nil {
    return -1, err
  }

  return new(intAryPhoton).compareSignedValues(
    ia, true, iAry2, true, ePrefix)
}

// CompareAbsoluteValues
//
//	Compares the absolute values of two IntAry instances.
//
//	Returns:
//	 0  = Current IntAry value is equal to the passed IntAry value.
//	 1  = Current IntAry value is greater than the passed IntAry value.
//	-1  = Current IntAry value is less than the passed IntAry value.
func (ia *IntAry) CompareAbsoluteValues(iAry2 *IntAry) (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CompareAbsoluteValues()",
    "")

  if err != nil {
    return -1, err
  }

  // The two absolute numeric values must be equal
  return new(intAryQuark).compareAbsoluteValues(
    ia, true, iAry2, true, ePrefix)
}

// CopyIn
//
//	Receives an input parameter 'IAry2', an instance of IntAry. The
//	data fields contained in 'IAry2' will be copied to the current
//	instance of IntAry ('ia'). Upon completion, the current IntAry
//	instance ('ia') will be an identical copy of 'iAry2'.
//
//	If input parameter 'copyBackUp' is set to 'true', a second copy
//	of 'IAry2' will be stored in the backup field of the current
//	IntAry instance ('ia.Backup')
func (ia *IntAry) CopyIn(iAry2 *IntAry, copyBackUp bool) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CopyIn()",
    "")

  if err != nil {
    return err
  }

  if iAry2 == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'iAry2'",
    }
  }

  err = new(intAryProton).copy(ia, iAry2, true, copyBackUp, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryProton).copy(\n"+
        "  ia, iAry2, validateSource=true, copyBackUp=%v, ePrefix)",
        copyBackUp),
      ErrContext: "iAry2 -> ia",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// CopyOut
//
//	 Makes a deep copy of the current IntAry instance, with backup,
//	 and returns it as a new IntAry object.
//
//		The term 'Backup' refers to the IntAry public field 'BackUp'.
//		Under this method, the returned 'IntAry' instance will have
//		a 'BackUp' field populated with a copy of the current IntAry
//		instance data.
func (ia *IntAry) CopyOut() (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CopyOut()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry2 := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry2, ia, true, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(\n" +
          "  &iAry2, ia, validateSource=true, copyToBackup=true, ePrefix)",
        ErrContext: "ia -> iAry2",
        ErrMessage: err.Error(),
      }
  }

  return iAry2, nil
}

// CopyOutNoBackup
//
//	Makes a deep copy of the current IntAry instance with NO backup.
//	The method then returns the deep copy as a new IntAry object.
//
//	The term 'Backup' refers to the IntAry public field 'BackUp'.
func (ia *IntAry) CopyOutNoBackup() (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CopyOutNoBackup()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry2 := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry2, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(\n" +
          "  &iAry2, ia, validateSource=true, copyToBackup=false, ePrefix)",
        ErrContext: "ia -> iAry2",
        ErrMessage: err.Error(),
      }
  }

  return iAry2, nil
}

// CopyOutDigits
//
//	Makes a deep copy of the current IntAry instance with NO
//	backup and returns that deep copy as a new instance of IntAry.
//
//	This method uses input parameter 'digitsToCopy' to copy a
//	specified number of digits to the new copy. The copy operation
//	is 'left to right' meaning that it starts with IntAry element
//	0 and proceeds to a length of 'digitsToCopy'.
//
//
//	No backup is included in the new, returned instance of IntAry.
//	of the original IntAry is returned in this copy. The term
//	'Backup' refers to the IntAry public field 'BackUp'.
func (ia *IntAry) CopyOutDigits(digitsToCopy int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CopyOutDigits()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  return new(intAryProton).copyOutDigits(
    ia, true, digitsToCopy, false, ePrefix)
}

// CopyOutPtr
//
//  Similar to CopyOut. In this case, however, the method returns
//  a pointer to a deep copy of the current IntAry instance.
//
//  Like method 'CopyOut', this method also includes a 'backup'.
//
//	The term 'Backup' refers to the IntAry public field 'BackUp'.
//	Under this method, the returned 'IntAry' instance will have
//	a 'BackUp' field populated with a duplicate copy of the
// 	current IntAry instance data.

func (ia *IntAry) CopyOutPtr() (*IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CopyOutPtr()",
    "")

  if err != nil {
    return &IntAry{}, err
  }

  iAry2 := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry2, ia, true, true, ePrefix)

  return &iAry2, err
}

// CopyToBackUp
//
//	Makes a copy of all data fields for the current IntAry instance
//	and saves a duplicate copy to the public 'BackUp' field of the
//	current IntAry instance.
//
//	To retrieve this backup data see method 'ResetFromBackUp()'.
//
//	Prior to initiating this 'copy' operation, the current instance
//	of IntAry will be subjected to validation tests.
func (ia *IntAry) CopyToBackUp() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.CopyToBackUp()",
    "")

  if err != nil {
    return err
  }

  return new(intAryLepton).copyToBackup(ia, ia, true, ePrefix)
}

// DecrementIntegerOne
//
//	Decrements the numeric value of the current IntAry instance by
//	subtracting '1'.
func (ia *IntAry) DecrementIntegerOne() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.DecrementIntegerOne()",
    "")

  if err != nil {
    return err
  }

  return new(intAryMechanics).decrementIntegerOne(
    ia, true, ePrefix)
}

// DivideByTwo
//
//	Divides the numeric value of the current IntAry instance by 2.
func (ia *IntAry) DivideByTwo() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.DivideByTwo()",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).divideByTwo(
    ia, true, ePrefix.XCpy("Divide 'ia' by 2"))
}

// DivideByInt64
//
//	Divide the current value of the intAry by an int64 'divisor'
//	parameter passed to the method.
//
//	If the quotient has a number of decimal places to the right of
//	the decimal point which is greater than 'maxPrecision', the
//	result is rounded to 'maxPrecision' decimal places.
//
//	If 'maxPrecision' is set equal to -1, 'maxPrecision' is
//	automatically set to 4,096.
//
//	If 'maxPrecision' is less than -1, an error will be returned.
func (ia *IntAry) DivideByInt64(
  divisor int64, maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.DivideByInt64()",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).divideByInt64(
    ia, true, divisor, maxPrecision, ePrefix)
}

// DivideByTenToPower
//
//	Divide the numerical value of the current IntAry instance by
//	10 raised to the power of the input parameter, 'exponent'.
//
//	             'ia'
//	    ia =  -------------
//	         (10^exponent)
//
//	The result or quotient is stored in the current IntAry instance.
func (ia *IntAry) DivideByTenToPower(exponent uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.DivideByTenToPower()",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).divideByTenToPower(
    ia, true, exponent, ePrefix)
}

// DivideThisBy
//
//	Divides the current value of intAry by the parameter iAry2. The
//	result of this division is returned as an intAry.
//
//	Given a ÷ b = c, 'a' is the dividend, 'b' is the divisor and
//	'c' is the quotient. For this method
//
//	    a = Current Instance of IntAry ('ia')
//	    b = Input parameter 'iAry2'
//	    c = Quotient returned by this method
//
//	Maximum precision of the division result is controlled by the
//	input parameter, 'maxPrecision'.
//
//	If 'maxPrecision' is greater than or equal to zero ('0'), the
//	number of digits to the right of the decimal place will not
//	exceed 'maxPrecision'.
//
//	If 'maxPrecision' is set equal to minus one ('-1'),
//	'maxPrecision' will be automatically set to a maximum of 4,096
//	digits to the right of the decimal point.
//
//	'minPrecision' specifies the minimum precision of the final
//	result. If 'minPrecision' is less than zero, it is automatically
//	set to zero.
func (ia *IntAry) DivideThisBy(iAry2 *IntAry, minPrecision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.DivideThisBy()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if iAry2 == nil {

    return IntAry{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'iAry2'",
      }
  }

  return new(intAryNeutron).divideIntArys(
    ia, true, iAry2, true, minPrecision, maxPrecision, ePrefix)
}

// Empty - Basically resets all the fields of the intAry
// structure to their 'zero' values.
func (ia *IntAry) Empty() {

  new(intAryElectron).empty(ia)

}

// EmptyBackUp - Deletes the values
// currently stored as backup to the
// current intAry object.
func (ia *IntAry) EmptyBackUp() {

  new(intAryBoson).emptyBackUp(ia)
}

// Equal
//
//	Returns 'true' if all field values of the current intAry object
//	are equal to all field values of the input parameter intAry
//	object, 'iAry2'.
//
//	Note that the BackUp fields for both compared IntAry objects
//	are NOT included in the 'Equals' comparison.
//
//	This method ('Equal') differs from method 'Equals' in that
//	this method ('Equal') returns an error parameter in addition
//	to a boolean return value.
//
//	If either of the two IntAry objects fail validation testing,
//	an error will be returned.
func (ia *IntAry) Equal(iAry2 *IntAry) (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.Equal()",
    "")

  if err != nil {
    return false, err
  }

  return new(intAryAtom).equal(ia, true, iAry2, true, ePrefix)
}

// Equals
//
//	Returns 'true' if all field values of the current intAry object
//	are equal to all field values of the input parameter intAry
//	object, 'iAry2'.
//
//	If any errors are encountered, a boolean value of 'false' is
//	returned
//
//	Note that the BackUp fields for both compared IntAry objects
//	are NOT included in the 'Equals' comparison.
func (ia *IntAry) Equals(iAry2 *IntAry) bool {

  return new(intAryElectron).equals(ia, iAry2)
}

// Floor
//
//	Math 'Floor' function. Finds the integer number which is less
//	than or equal to the value of the current intAry.
//
//	Reference Wikipedia
//	  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//	Examples
//	========
//
//	  Initial     Floor
//	   Value      Value
//	  -------    -------
//	   5.95         5
//	   5.05         5
//	   5            5
//	  -5.05        -6
//	   2.4          2
//	   2.9          2
//	  -2.7         -3
//	  -2           -2
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
//	with Numeric Separators copied from the original, current
//	IntAry instance, 'ia'.
//
//	Validation Testing
//	==================
//
//	This method will subject the current instance of IntAry to
//	validation tests.
func (ia *IntAry) Floor() (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.Floor()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  return new(intAryNeutron).floor(ia, true, ePrefix)
}

// GetAbsoluteValue
//
//	Returns an IntAry object which represents the Absolute Value of
//	the current intAry instance.
func (ia *IntAry) GetAbsoluteValue() (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetAbsoluteValue()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  absIa := IntAry{}

  err = new(intAryProton).copy(&absIa, ia, true, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&absIa, ia, true, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = new(intAryMechanics).setAbsoluteValue(&absIa, ePrefix.XCpy("Set 'absIa' Absolute Value"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryMechanics).setAbsoluteValue(\n" +
          "  &absIa, ePrefix.XCpy(Set 'absIa' Absolute Value)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return absIa, nil
}

// GetBigInt
//
//	Returns the current value of this intAry object expressed as a
//	signed integer number of type *big.Int.
func (ia *IntAry) GetBigInt() (*big.Int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetBigInt()",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia,Validating 'ia')",
        ErrContext: "The current instance of IntAry ('ia') is INVALID!\n" +
          "'ia' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return new(intAryNeutron).getBigInt(
    ia,
    false,
    ePrefix)
}

// GetBigIntNum
//
//	Converts the numeric value of the current IntAry to 'BigIntNum'
//	instance and returns it to the calling function.
//
//	The returned BigIntNum will contain numeric separators (decimal
//	separator, thousands separator and currency symbol) copied from
//	the current IntAry instance.
func (ia *IntAry) GetBigIntNum() (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetBigIntNum()",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia,ePrefix.XCpy(\"Validating 'ia'\").String())\n",
        ErrContext: "Current Instance of 'IntAry' is INVALID!\n" +
          "'ia'FAILED Validation Tests",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(ia, false, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(intAryPhoton).\n" +
          "  getNumericSeparatorsDto(ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bInt, err := new(intAryNeutron).getBigInt(
    ia,
    false,
    ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bInt, err := new(intAryNeutron).getBigInt(ia, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bIntNum, err := new(BigIntNum).NewBigInt(bInt, uint(ia.precision))

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bIntNum, err := new(BigIntNum).NewBigInt(bInt, uint(ia.precision))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = bIntNum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = bIntNum.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return bIntNum, nil
}

// GetCurrencySymbol
//
//	Returns a type 'rune' which represents the currency symbol
//	as configured in the current IntAry object.
//
// To set the value of currency symbol, see the method
// 'IntAry.SetCurrencySymbol()'.
//
//	Note: if the current IntAry currency value
//	was not previously set, it will be automatically
//	set to the USA dollar sign ('$').
func (ia *IntAry) GetCurrencySymbol() rune {

  if ia.currencySymbol == 0 {
    ia.currencySymbol = '$'
  }

  return ia.currencySymbol
}

// GetDecimal
//
//	Converts the current IntAry instance to a Type, 'Decimal' and
//	returns it to the calling function.
func (ia *IntAry) GetDecimal() (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetDecimal",
    "")

  if err != nil {
    return Decimal{}, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Valildating 'ia'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia, ePrefix.XCpy(Valildating 'ia').String())",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  //numSeps, err := ia.GetNumericSeparatorsDto()
  numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(ia, false, ePrefix)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(\n" +
          "  ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iaNumStr, err := new(intAryAtom).getRawNumStr(ia, false, ePrefix)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "iaNumStr, err := new(intAryAtom).\n" +
          "  getRawNumStr(ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  dec, err := new(Decimal).NewNumStrWithNumSeps(
    iaNumStr, numSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "dec, err := new(Decimal).NewNumStrWithNumSeps(iaNumStr, numSeps)",
        ErrContext: fmt.Sprintf("iaNumStr= '%v'\nnumSeps= '%v\n",
          iaNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  return dec, nil
}

// GetDecimalSeparator
//
//	Returns a type 'rune' which represents the setting for decimal
//	separator in the current IntAry object.
//
//	A decimal separator separates integer digits from fractional
//	digits which the value of the current IntAry object is
//	expressed as a string.
//
//	To set the value of decimal separator, see method
//	'IntAry.SetDecimalSeparator()'.
func (ia *IntAry) GetDecimalSeparator() rune {

  if ia.decimalSeparator == 0 {
    ia.decimalSeparator = '.'
  }

  return ia.decimalSeparator
}

// GetFractionalDigits
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
//	with Numeric Separators copied from the original, current
//	instance of IntAry.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry
func (ia *IntAry) GetFractionalDigits() (IntAry, error) {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetFractionalDigits()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Valildating 'ia'").String())

  if err != nil {
    return IntAry{}, err
  }

  return new(intAryNeutron).getFractionalDigits(ia, true, ePrefix)
}

// GetIntegerDigits
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
//	with Numeric Separators copied from the original, current
//	instance of IntAry.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry
func (ia *IntAry) GetIntegerDigits() (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToZero",
    "")

  if err != nil {
    return IntAry{}, err
  }

  return new(intAryNeutron).getIntegerDigits(ia, true, ePrefix)
}

// GetInt
//
// Returns the value of the current intAry as a 32-bit integer. If
// the intAry exceeds the maximum or minimum values for int, an
// error will be returned.
//
// Minimum and Maximum Values for 32-bit Integer (int32 & int
// Types):
//
//	-2,147,483,648 to 2,147,483,647
//
//	Anything out this range will generate an error.
//
// Reference: https://golang.org/ref/spec#Numeric_types
func (ia *IntAry) GetInt() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetInt()",
    "")

  if err != nil {
    return 0, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Valildating 'ia'").String())

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "  ia, ePrefix.XCpy(Valildating 'ia').String())",
        ErrContext: "The current instance of IntAry ('ia') is INVALID!\n",
        ErrMessage: err.Error(),
      }
  }

  maxInt := big.NewInt(0).SetInt64(int64(math.MaxInt32))

  minInt := big.NewInt(0).SetInt64(int64(math.MinInt32))

  // result, err := ia.GetBigInt()
  result, err := new(intAryNeutron).getBigInt(ia, false, ePrefix)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "result, err := new(intAryNeutron).getBigInt(ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  compare := result.Cmp(maxInt)

  if compare == 1 {
    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: the value of the current intAry object\n"+
          "exceeds the maximum allowable value for the int type.\n"+
          "The Maximum allowable int value is: %v\n"+
          "The computed value of the current IntAry object is: %v",
          maxInt.Text(10), result.Text(10)),
      }
  }

  compare = result.Cmp(minInt)

  if compare == -1 {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: the value of the intAry object is less\n"+
          "than the minimum allowable value for the int type.\n"+
          "The Minimum allowable int value is: %v\n"+
          "The computed value of the current IntAry object is: %v",
          minInt.Text(10), result.Text(10)),
      }
  }

  return int(result.Int64()), nil
}

// GetInt64
//
//	Returns the value of the current IntAry as an int64 value. If
//	the IntAry value exceeds the maximum or minimum values for
//	int64, an error will be returned.
//
//	Minimum and Maximum Values for 64-bit Integer (int64 Type):
//
//	-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
//
//	Anything outside this range will generate an error.
//
// Reference: https://golang.org/ref/spec#Numeric_types
func (ia *IntAry) GetInt64() (int64, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetInt64()",
    "")

  if err != nil {
    return 0, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Valildating 'ia'").String())

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "  ia, ePrefix.XCpy(Valildating 'ia').String())",
        ErrContext: "The current instance of IntAry ('ia') is INVALID!\n",
        ErrMessage: err.Error(),
      }
  }

  maxI64 := big.NewInt(0).SetInt64(math.MaxInt64)

  minI64 := big.NewInt(0).SetInt64(math.MinInt64)

  result, err := new(intAryNeutron).getBigInt(ia, false, ePrefix)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "result, err := new(intAryNeutron).\n" +
          "  getBigInt(ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  compare := result.Cmp(maxI64)

  if compare == 1 {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: The value of the current intAry object exceeds the\n"+
          "maximum allowable value for the int64 type.\n"+
          "The Maximum allowable int64 value is: %v\n"+
          "The actual computed value for the current IntAry Object is: %v\n",
          maxI64.Text(10), result.Text(10)),
      }
  }

  compare = result.Cmp(minI64)

  if compare == -1 {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: the value of the current intAry object is less than the\n"+
          "minimum allowable value for the int64 type.\n"+
          "The Minimum allowable int value is: %v\n"+
          " %v\n",
          maxI64.Text(10), result.Text(10)),
      }
  }

  return result.Int64(), nil
}

// GetIntAryElement
//
//	Returns an element of the internal integer array maintained by
//	the current IntAry object.
//
//	Input Parameter
//	===============
//
//	index                    int
//	  The element returned is based on the integer index passed to
//	  the method. If the index is outside the range of the internal
//	  array, an error is returned.
//
//	Return Values
//	=============
//
//	uint8
//	  If the 'index' passed to the method is valid, this method
//	  will return an integer of type uint8 representing the value
//	  of the internal integer array at the specified index.
//
//	error
//	  If no errors are encountered, this method will set the
//	  returned error value to 'nil'.
func (ia *IntAry) GetIntAryElement(index int) (uint8, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetIntAryElement()",
    "")

  if err != nil {
    return 0, err
  }

  if index < 0 || index > (ia.intAryLen-1) {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "index < 0 || index > (ia.intAryLen-1)",
        ErrMessage: fmt.Sprintf("Error: Input parameter 'index' is INVALID!\n"+
          "'index' Out Of Array Bounds!\n"+
          "'index' value must be >= 0 and index cannot exceed\n"+
          "the interal integer array length +1\n"+
          "Internal integer array length: %v"+
          "index= '%v'", ia.intAryLen, index),
      }
  }

  result := ia.intAry[index]

  return result, nil
}

// GetIntAryInt
//
//	Returns an element of the internal integer array maintained by
//	the current IntAry object. The value returned is of Type 'int'.
//
//	Input Parameter
//	===============
//
//	index                    int
//	  The element returned is based on the integer index passed to
//	  the method. If the index is outside the range of the internal
//	  array, an error is returned.
//
//	Return Values
//	=============
//
//	int
//	  If the 'index' passed to the method is valid, this method
//	  will return an integer of type 'int' representing the value
//	  of the internal integer array at the specified index.
//
//	error
//	  If no errors are encountered, this method will set the
//	  returned error value to 'nil'.
func (ia *IntAry) GetIntAryInt(index int) (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetIntAryInt()",
    "")

  if err != nil {
    return 0, err
  }

  if index < 0 || index > (ia.intAryLen-1) {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "index < 0 || index > (ia.intAryLen-1)",
        ErrMessage: fmt.Sprintf("Error: Input parameter 'index' is INVALID!\n"+
          "'index' Out Of Array Bounds!\n"+
          "'index' value must be >= 0 and index cannot exceed\n"+
          "the interal integer array length +1\n"+
          "Internal integer array length: %v"+
          "index= '%v'", ia.intAryLen, index),
      }
  }

  return int(ia.intAry[index]), nil
}

// GetIntAryRune
//
//	 Returns an element of the internal integer array maintained by
//	 the current IntAry object. The value returned is of Type
//	 'rune'.
//
//		Input Parameter
//		===============
//
//		index                    int
//		  The element returned is based on the integer index passed to
//		  the method. If the index is outside the range of the internal
//		  array, an error is returned.
//
//		Return Values
//		=============
//
//		rune
//		  If the 'index' passed to the method is valid, this method
//		  will return an instance of type 'rune' representing the value
//		  of the internal integer array at the specified index.
//
//		error
//		  If no errors are encountered, this method will set the
//		  returned error value to 'nil'.
func (ia *IntAry) GetIntAryRune(index int) (rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetIntAryRune()",
    "")

  if err != nil {
    return 0, err
  }

  if index < 0 || index > (ia.intAryLen-1) {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "index < 0 || index > (ia.intAryLen-1)",
        ErrMessage: fmt.Sprintf("Error: Input parameter 'index' is INVALID!\n"+
          "'index' Out Of Array Bounds!\n"+
          "'index' value must be >= 0 and index cannot exceed\n"+
          "the interal integer array length +1\n"+
          "Internal integer array length: %v"+
          "index= '%v'", ia.intAryLen, index),
      }
  }

  return rune(ia.intAry[index] + 48), nil
}

// GetIntAry
//
//	Returns a deep copy of the current IntAry instance. This method
//	is necessary in order to comply with the requirements of the
//	INumMgr interface.
//
//	The returned IntAry copy will also contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from the current IntAry instance.
//
//	Before returning the new IntAry instance, this method performs
//	a validity test on the current IntAry instance.
func (ia *IntAry) GetIntAry() (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetIntAry()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry2 := new(IntAry).New()

  err = new(intAryProton).copy(&iAry2, ia, true, true, ePrefix)

  return iAry2, err
}

// GetIntAryElements
//
//		Returns the internal integer  array used by the IntAry object.
//
//		************************************************
//
//		WARNING!! This returns a reference (pointer)
//		to the internal integer array for this IntAry
//		object.
//
//		************************************************
//
//		Return parameters
//		=================
//
//		[]uint8
//		  An array of unsigned 8-bit integers
//
//		int
//		  The length of the []uint8 array returned above.
//
//	 error
//	   If no errors are encountered, this method will
//	   set this returned error value to 'nil'.
//
//		************************************************
//
//		          BE CAREFUL!!!!!!!
//
//		************************************************
//
//		Be careful!! - This returns a slice and therefore a reference
//		(i.e. pointer) to the internal array for this IntAry object. If
//		you alter this array externally, it will alter the internal
//		array values.
//
//		IMPORTANT: If you alter the returned array, the current IntAry
//		object may become invalid. After changing the returned array,
//		Call method 'IntAry.SetInternalFlags()' to ensure that the
//		IntAry object correctly reflects the changed values.
//
//	 Be advised that this method does NOT perform validation tests
//	 on the current IntAry instance. To validate this instance, see
//	 method:
//	           IntAry.IsValid()
//
//		For an array which can be safely manipulated, independent of
//		the current IntAry instance, see method:
//		          IntAry.GetIntAryDeepCopy()
//
//		To Append Elements to the internal integer array, see method:
//		          IntAry.AppendToIntAry()
func (ia *IntAry) GetIntAryElements() ([]uint8, int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetIntAryElements()",
    "")

  if err != nil {
    return []uint8{}, 0, err
  }

  err = new(intAryNanobot).setInternalFlags(
    ia, ePrefix.XCpy("Setting 'ia' Flags"))

  if err != nil {

    return []uint8{}, 0, &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
        "  ia, ePrefix.XCpy(Setting 'ia' Flags))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return ia.intAry, ia.intAryLen, nil
}

// GetIntAryDeepCopy
//
//	 Returns a deep copy of the internal integer array maintained by
//	 this IntAry object. Unlike the array returned by method
//	 GetIntAryElements(), the array returned by this method is not a
//	 reference or pointer.
//
//	 This means that the array returned by this method may be altered
//	 externally without changing the value of the original internal
//	 array maintained by this IntAry object.
//
//		Validation Testing
//		==================
//
//		This method will perform validation tests on the current instance
//		of IntAry ('ia').
//
//		Return parameters
//		=================
//
//		[]uint8
//		  An array of unsigned 8-bit integers representing a deep copy
//	   of the internal array for the current IntAry object.
//
//		int
//		  The length of the []uint8 array returned above.
//
//	 error
//	   If no errors are encountered, this method will
//	   set this returned error value to 'nil'.
func (ia *IntAry) GetIntAryDeepCopy() ([]uint8, int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetIntAryDeepCopy",
    "")

  if err != nil {
    return []uint8{}, 0, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return []uint8{}, 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "  ia, ePrefix.XCpy(Validating 'ia'))",
        ErrContext: "Error: The current instance of IntAry ('ia') is INVALID!\n" +
          "'ia' FAILED Validatin Tests.",
        ErrMessage: err.Error(),
      }
  }

  ary := make([]uint8, ia.intAryLen)

  for i := 0; i < ia.intAryLen; i++ {
    ary[i] = ia.intAry[i]
  }

  return ary, ia.intAryLen, nil
}

// GetIntAryLength
//
//	Returns the length of the internal integer array maintained by
//	the current IntAry object.
func (ia *IntAry) GetIntAryLength() int {

  new(intAryNanobot).setInternalFlagsNoErrors(ia)

  return ia.intAryLen
}

// GetIntAryStats
//
//	Returns a series of descriptive statistics about the internal
//	Integer Array maintained by the current IntAry object.
//
//	This method returns an IntAryStatsDto object which includes the
//	following information on the current IntAry object:
//
//	IntAryLen                - Type int
//	  Length of the internal integer array
//
//	IntegerLen               - Type int
//	  Number of integer digits in the current IntAry value.
//
//	SignificantIntegerLen    - Type int
//	  The number of non-zero integer digits in the current IntAry
//	  value.
//
//	SignificantFractionLen   - Type int
//	  The number of non-zero digits to the right of the decimal
//	  point in the current IntAry value.
//
//	precision                - Type int
//	  The number of digits to the right of the decimal point in the
//	  current IntAry value.
//
//	SignVal                  - Type int
//	  Either +1 or -1 indicating the sign of the current IntAry
//	  value.
//
//	FirstDigitIdx            - Type int
//	  The IntAry index of the first non-zero digit in the current
//	  internal integer array.
//
//	LastDigitIdx             - Type int
//	  The IntAry index of the last non-zero digit in the internal
//	  integer array.
//
//	IsZero                   - bool
//	  A boolean value indicating whether the current IntAry value
//	  is zero.
//
//	IsIntegerZeroValue       - bool
//	  A boolean value indicating whether integer value to the left
//	  of the decimal point is zero.
//
//	DecimalSeparator         - Type rune
//	  The character used to separate integer and fractional elements
//	  of the current IntAry value when presented in string format.
//
//	ThousandsSeparator       - Type rune
//	  The character used to separate thousands when the current
//	  IntAry value is presented in string format.
//
//	CurrencySymbol           - Type rune
//	  The character used to designate currency when the current
//	  IntAry value is presented as a currency string.
func (ia *IntAry) GetIntAryStats() IntAryStatsDto {

  iStats := IntAryStatsDto{}

  new(intAryNanobot).setInternalFlagsNoErrors(ia)

  iStats.IntAryLen = ia.intAryLen
  iStats.IntegerLen = ia.integerLen
  iStats.SignificantIntegerLen = ia.significantIntegerLen
  iStats.SignificantFractionLen = ia.significantFractionLen
  iStats.Precision = ia.precision
  iStats.SignVal = ia.signVal
  iStats.FirstDigitIdx = ia.firstDigitIdx
  iStats.LastDigitIdx = ia.lastDigitIdx
  iStats.IsZeroValue = ia.isZeroValue
  iStats.IsIntegerZeroValue = ia.isIntegerZeroValue
  iStats.DecimalSeparator = ia.decimalSeparator
  iStats.ThousandsSeparator = ia.thousandsSeparator
  iStats.CurrencySymbol = ia.currencySymbol

  return iStats
}

// GetMagnitude
//
//	Returns the magnitude of the integer portion of the current
//	IntAry numeric value. The integer portion of the number is
//	represented by the digits to the left of the decimal point.
//
//	Magnitude is defined here as the power of 10 which generates a
//	value less than or equal to the integer portion of the current
//	IntAry numeric value.
//
//	   10^magnitude  <= IntAry value
//
//	Note
//	====
//
//	If the current IntAry value is negative, an error will be generated.
func (ia *IntAry) GetMagnitude() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetMagnitude()",
    "")

  if err != nil {
    return 0, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia, ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia, ePrefix.XCpy(Validating 'ia').String())",
        ErrContext: "Error: The current instance of IntAry ('ia') is INVALID!\n" +
          "'ia' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  err = new(intAryNanobot).setInternalFlags(ia, ePrefix.XCpy("Setting Flags 'ia'"))

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
          "  ia, ePrefix.XCpy(Setting Flags 'ia'))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iaNumStr, err := new(intAryAtom).getRawNumStr(ia, false, ePrefix)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNumStr, err := new(intAryAtom).getRawNumStr(ia, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if ia.signVal == -1 {

    return -1,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "ia.signVal == -1",
        ErrMessage: "Error: current IntAry value is negative!\n" +
          fmt.Sprintf("value= '%v'", iaNumStr),
      }
  }

  return ia.intAryLen - ia.precision - ia.firstDigitIdx - 1, nil
}

// GetMagnitudeDigits
//
//	Returns the number of digits in the integer portion of the
//	current IntAry numeric value.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry
func (ia *IntAry) GetMagnitudeDigits() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetMagnitudeDigits()",
    "")

  if err != nil {
    return 0, err
  }

  return new(intAryMechanics).getMagnitudeDigits(ia, true, ePrefix)
}

// GetNumericSeparatorsDto
//
//	Returns a NumericSeparatorDto structure containing the
//	character or rune values for decimal point separator, thousands
//	separator and currency symbol.
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
func (ia *IntAry) GetNumericSeparatorsDto() (NumericSeparatorDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetNumericSeparatorsDto()",
    "")

  if err != nil {
    return NumericSeparatorDto{}, err
  }

  return new(intAryPhoton).getNumericSeparatorsDto(ia, false, ePrefix)
}

// GetNumStr
//
//	Returns the current value of this intAry object as a number
//	string.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. These digits
//	are formatted in a way that facilitates conversion to a
//	corresponding numeric value.
//
//	The number string returned by this method will contain a
//	decimal separator to separate integer and fractional
//	components of the numeric value. The returned number string
//	will not contain 'thousands' separators or 'currency' symbols.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry.
func (ia *IntAry) GetNumStr() (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetNumStr()",
    "")

  if err != nil {
    return "", err
  }

  return new(intAryAtom).getRawNumStr(ia, true, ePrefix)
}

// GetNumStrDto
//
//	Converts the current IntAry to a NumStrDto instance and
//	returns it.
//
//	The returned NumStrDto will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from the current IntAry instance.
//
//	Before returning the new NumStrDto instance, this method
//	performs a validity check on the current IntAry instance.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry.
func (ia *IntAry) GetNumStrDto() (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetNumStrDto()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  return new(intAryNeutron).getNumStrDto(ia, true, ePrefix)
}

// GetNthRootOfThis
//
//	Returns an intAry object equal to the 'nth Root' of the current
//	intAry.
//
//	Input Parameters
//	================
//
//	nthRoot                  int
//	  Value which specifies the root to calculate.
//
//	maxPrecision             int
//	  Value which specifies the number of digits to the right of
//	  the decimal place in the result.
//
//	Return Values
//	=============
//
//	IntAry
//	  A new instance of IntAry configured with the result of the
//	  'nth Root' calculation.
//
//	error
//	  If no errors are encountered, this return value will be set
//	  to 'nil'.
func (ia *IntAry) GetNthRootOfThis(nthRoot int, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetNthRootOfThis()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot = '%v'", nthRoot),
        ErrMessage: err.Error(),
      }
  }

  nthRt := NthRootOp{}

  return nthRt.GetNthRootIntAry(ia, &iaNthRoot, maxPrecision)
}

// GetPrecision
//
//	Returns the precision value for the current intAry object.
//
//	'precision' is defined as the number of numeric digits to the
//	right of the decimal place. To compute the location of the
//	decimal point in a string of numeric digits, go to the right
//	most digit in the number string and count left 'precision'
//	digits.
//
//	The value of 'precision' returned by this method will always
//	be >= zero (greater than or equal to zero '0').
//
//	Example
//	=======
//
//	 1.234    GetPrecision() = 3
//	 5        GetPrecision() = 0
//	 0.12345  GetPrecision() = 5
//
//	Number String    precision    Fractional Number
//	   123456            3            123.456
func (ia *IntAry) GetPrecision() int {
  return ia.precision
}

// GetPrecisionUint - returns the precision value for the
// current intAry object as an unsigned integer (uint).
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// The value of 'precision' returned by this method will
// always be >= zero (greater than or equal to zero '0').
//
// Example:
//
//					1.234    	GetPrecision() = 3
//							5			GetPrecision() = 0
//				0.12345  		GetPrecision() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (ia *IntAry) GetPrecisionUint() (uint, error) {
  return uint(ia.precision), nil
}

// GetRuneArray
//
//	Returns all the elements of the current IntAry as an array of
//	runes. The returned array of runes represents a deep copy of
//	the internal array encapsulated by the current instance of
//	IntAry.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry.
func (ia *IntAry) GetRuneArray() ([]rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetRuneArray()",
    "")

  if err != nil {
    return []rune{}, err
  }

  err = new(intAryNanobot).setInternalFlags(
    ia, ePrefix.XCpy("Setting 'ia' Flags"))

  if err != nil {

    return []rune{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
          "  ia, ePrefix.XCpy(\"Setting 'ia' Flags\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  aLen := ia.intAryLen

  if aLen == 0 {
    return []rune{}, nil
  }

  outRunes := make([]rune, aLen)

  for i := 0; i < aLen; i++ {

    outRunes[i] = rune(ia.intAry[i] + 48)

  }

  return outRunes, nil
}

// GetScaleFactor
//
//	Returns a pointer to a Big Integer (*big.Int) which specifies
//	the scale factor associated with the numeric value of the
//	current IntAry instance.
//
//	Scale Value, or Scale Factor, is defined by 10 raised to the
//	power of IntAry precision.
func (ia *IntAry) GetScaleFactor() (*big.Int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetScaleFactor",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "ia, ePrefix.XCpy(Validating 'ia').String())",
        ErrContext: "Error: The current instance of IntAry ('ia') is INVALID!\n" +
          "'ia' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  if ia.precision == 0 {
    return big.NewInt(int64(1)), nil
  }

  base10 := big.NewInt(0).SetInt64(int64(10))

  bigPrecision := big.NewInt(0).SetInt64(int64(ia.precision))

  scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

  return scaleFactor, nil
}

// GetSciNotationNumber
//
//	Converts the numeric value of the current IntAry instance into
//	scientific notation and returns this value as a new instance of
//	type SciNotationNum.
//
//	Scientific Notation
//	===================
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//
//	Example Scientific Notation
//	===========================
//
//	scientific notation string: '2.652e+8'
//	significand = '2.652'
//	significand integer digit = '2'
//	mantissa    = significand factional digits = '.652'
//	exponent    = '8'  (10^8)
//
//	Note
//	====
//
//	The maximum number of digits which will be retained in the
//	significand is 50,000.
//
//	Validation Testing
//	==================
//
//	This method will subject the current instance of IntAry ('ia')
//	to validation tests.
//
//	Input Parameter
//	===============
//
//	mantissaLen              uint
//	  Specifies the length of the mantissa in the returned
//	  scientific notation string. If the value of 'mantissaLen' is
//	  less than two ('2'), this method will automatically set the
//	  'mantissaLen' to a default value of two ('2').
//
//	Return Values
//	=============
//
//	SciNotationNum
//	  This type encapsulates the IntAry numeric value expressed in
//	  Scientific Notation.
//
//	error
//	  If no errors are encountered, this method will return an
//	  'error' value of 'nil'.
func (ia *IntAry) GetSciNotationNumber(mantissaLen uint) (SciNotationNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetSciNotationNumber",
    "")

  if err != nil {
    return SciNotationNum{}, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return SciNotationNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia,ePrefix.XCpy(Validating 'ia').String())",
        ErrContext: "Error: The current instance of IntAry ('ia') is INVALID!\n" +
          "'ia' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  if mantissaLen < 2 {
    mantissaLen = 2
  }

  sciNotationNum := new(SciNotationNum).New()

  if ia.isZeroValue {

    newBINumZeroMantissa, err := new(BigIntNum).NewZero(mantissaLen)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "newBINumZeroMantissa, err := new(BigIntNum).NewZero(mantissaLen)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    newBINumZero, err := new(BigIntNum).NewZero(0)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "newBINumZero, err := new(BigIntNum).NewZero(0)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    err = sciNotationNum.SetBigIntNumElements(
      newBINumZeroMantissa, newBINumZero)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = sciNotationNum.SetBigIntNumElements(\n" +
            "  newBINumZeroMantissa, newBINumZero)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    return sciNotationNum, nil
  }

  if !ia.isIntegerZeroValue {

    magnitudeInt, err := new(intAryNeutron).getMagnitude(
      ia, false, ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "magnitudeInt, err := new(intAryNeutron).getMagnitude(\n" +
            "  ia, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    iaNew := IntAry{}

    err = new(intAryProton).copy(&iaNew, ia, false, false, ePrefix.XCpy("Copy ia->iaNew"))

    if err != nil {
      return SciNotationNum{}, err
    }

    // err = iaNew.DivideByTenToPower(uint(magnitudeInt))
    err = new(intAryNeutron).divideByTenToPower(&iaNew, true, uint(magnitudeInt), ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = err = new(intAryNeutron).\n" +
            "  divideByTenToPower( &iaNew, true, uint(magnitudeInt), ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    iaMagnitude := new(intAryElectron).newIntAry()

    nsProfile := NumSepsProfileSelection{
      SourceObjectName:         "ia",
      OutputNumSepsName:        "numSeps",
      UseDefaultNumSeps:        false,
      SetDefaultNumSepsIfEmpty: true,
      ValidateNumSeps:          false,
      OverrideNumSeps:          NumericSeparatorDto{},
    }

    err = new(intAryGluon).setIntAryWithInt(
      &iaMagnitude,
      ia,
      nsProfile,
      magnitudeInt,
      0,
      false,
      ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryGluon).setIntAryWithInt(\n" +
            "&iaMagnitude, ia, nsProfile, magnitudeInt, 0, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    if ia.precision > 50000 {

      err = new(intAryMolecule).
        setPrecision(&iaNew, false, 50000, true, ePrefix.XCpy("Set Precision 'iaNew'"))

      if err != nil {

        return SciNotationNum{},
          &FuncReturnError{
            ErrPrefix: ePrefix.String(),
            ReturnFunc: "err = new(intAryMolecule).\n" +
              "  setPrecision(&iaNew, false, 50000, true,\n" +
              "  ePrefix.XCpy(Set Precision 'iaNew'))",
            ErrContext: "",
            ErrMessage: err.Error(),
          }
      }
    }

    err = sciNotationNum.SetIntAryElements(iaNew, iaMagnitude)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = sciNotationNum.SetIntAryElements(iaNew, iaMagnitude)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {
    // Must be number with zero integers and fractional digits.
    // Example: 0.256

    //iaFracPart, err := ia.GetFractionalDigits()
    iaFracPart, err := new(intAryNeutron).getFractionalDigits(
      ia, false, ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "iaFracPart, err := new(intAryNeutron).getFractionalDigits(\n" +
            "ia, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    //iaFracPart.MultiplyByTenToPower(uint(iaFracPart.precision))
    err = new(IntAryMathMultiply).MultiplyByTenToPower(&iaFracPart, uint(iaFracPart.precision))

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(IntAryMathMultiply).MultiplyByTenToPower(\n" +
            "  &iaFracPart, uint(iaFracPart.precision))",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    err = new(intAryAtom).optimizeIntArrayLen(&iaFracPart, false, false, false, ePrefix.XCpy("Optimize iaFracPart"))

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryAtom).optimizeIntArrayLen(\n" +
            "  &iaFracPart, validateIntAry=false, optimizeFracDigits=false,\n" +
            "  validateResult=false ePrefix.XCpy(Optimize iaFracPart))",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    //intMagnitudeFrac, err := iaFracPart.GetMagnitude()
    intMagnitudeFrac, err := new(intAryNeutron).getMagnitude(
      &iaFracPart, false, ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "intMagnitudeFrac, err := new(intAryNeutron).getMagnitude(\n" +
            "  &iaFracPart, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    //iaFracPart.DivideByTenToPower(uint(intMagnitudeFrac))
    err = new(intAryNeutron).divideByTenToPower(
      &iaFracPart, false, uint(intMagnitudeFrac), ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryNeutron).divideByTenToPower(\n" +
            "  &iaFracPart, false, uint(intMagnitudeFrac), ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    intMagnitudeFrac = intMagnitudeFrac - ia.precision

    // Old Code
    //iaMagnitude, err := new(IntAry).NewInt(intMagnitudeFrac, 0)
    //
    //if err != nil {
    //
    //	return SciNotationNum{},
    //		&FuncReturnError{
    //			ErrPrefix:  ePrefix.String(),
    //			ReturnFunc: "iaMagnitude, err := new(IntAry).NewInt(intMagnitudeFrac, 0)",
    //			ErrContext: "",
    //			ErrMessage: err.Error(),
    //		}
    //}

    // -------------------------------------------------------
    //              New Code

    iaMagnitude := new(intAryElectron).newIntAry()

    nsProfile := NumSepsProfileSelection{
      SourceObjectName:         "ia",
      OutputNumSepsName:        "numSeps",
      UseDefaultNumSeps:        false,
      SetDefaultNumSepsIfEmpty: true,
      ValidateNumSeps:          false,
      OverrideNumSeps:          NumericSeparatorDto{},
    }

    err = new(intAryGluon).setIntAryWithInt(
      &iaMagnitude,
      ia,
      nsProfile,
      intMagnitudeFrac,
      0,
      false,
      ePrefix)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryGluon).setIntAryWithInt(\n" +
            "&iaMagnitude, ia, nsProfile, intMagnitudeFrac, 0, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    // -------------------------------------------------------

    err = sciNotationNum.SetIntAryElements(iaFracPart, iaMagnitude)

    if err != nil {

      return SciNotationNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = sciNotationNum.SetIntAryElements(iaFracPart, iaMagnitude)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  sciNotationNum.SetMantissaLength(mantissaLen)

  return sciNotationNum, nil
}

// GetSciNotationStr - Returns a string expressing the current IntAry
// numerical value as scientific notation.
//
// Input parameter 'mantissaLen' is used to express the number of
// fractional digits displayed in the returned scientific notation
// string. If 'mantissaLen' is less than '2' (two), 'mantissaLen'
// will be automatically set to '2' (two).
//
// Input Parameter
// ===============
//
// mantissaLen uint	-
//
//	Specifies the length of the mantissa in the returned
//	scientific notation string.
//
//	Example Scientific Notation:
//	----------------------------
//
//	scientific notation string: '2.652e+8'
//
//	significand = '2.652'
//	significand integer digit = '2'
//	mantissa		= significand factional digits = '.652'
//	exponent    = '8'  (10^8)
func (ia *IntAry) GetSciNotationStr(mantissaLen uint) (string, error) {

  ePrefix := "IntAry.GetSciNotationStr()"

  if mantissaLen < 2 {
    mantissaLen = 2
  }

  sciNotationNum, err := ia.GetSciNotationNumber(mantissaLen)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by ia.GetSciNotationNumber(mantissaLen)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  result, err := sciNotationNum.GetSciNotationStr(mantissaLen)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by sciNotationNum.GetSciNotationStr(mantissaLen)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return result, nil
}

// GetSign - returns the sign of the current
// intAry object as an integer value.
//
// The sign value returned by this method should
// always be one of two values: +1 or -1 .
func (ia *IntAry) GetSign() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.GetSign()",
    "")

  if err != nil {
    return 0, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia, ePrefix.XCpy("Validating 'ia'").String())

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia, ePrefix.(Validating 'ia'",
        ErrContext: "Error: The current instance of IntAry ('ia') is INVALID!\n" +
          "'ia' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return ia.signVal, nil
}

// GetSquareRootOfThis - Returns an intAry object equal to the 'square root'
// of the current intAry.
//
// Input Parameters:
// 'maxPrecision' - uint value which specifies the number of digits
//
//	to the right of the decimal place in the result.
func (ia *IntAry) GetSquareRootOfThis(maxPrecision int) (IntAry, error) {
  nthRt := NthRootOp{}

  return nthRt.GetSquareRootIntAry(ia, maxPrecision)
}

// GetThisPointer - Returns a pointer to the current IntAry instance
func (ia *IntAry) GetThisPointer() *IntAry {

  return ia
}

// GetThousandsSeparator - returns a value of type 'rune'
// which represents the thousands separator associated
// with the current IntAry object.
//
// The thousands separator is used to separate thousands
// in the integer digits of this IntAry value when that
// value is expressed as a string.
//
// In the US, the thousands separator is typically the comma
// character (','). In this example, the thousands separator
// is a comma: Example - '1,000,000,000'.
//
// To set the value of thousands separator, see the method
// SetThousandsSeparator().
func (ia *IntAry) GetThousandsSeparator() rune {

  if ia.thousandsSeparator == 0 {
    ia.thousandsSeparator = ','
  }

  return ia.thousandsSeparator
}

// HasFractionalDigits
//
//	This method examines the current intAry object to determine if
//	there are non-zero digits to the right of the decimal place. If
//	all digits to the right of the decimal place are zero, this
//	method returns 'false'
//
//	If non-zero digits are present to the right of the decimal
//	place, the method returns 'true'.
func (ia *IntAry) HasFractionalDigits() (bool, error) {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.HasFractionalDigits",
    "")

  if err != nil {
    return false, err
  }

  return new(intAryNanobot).hasFractionalDigits(
    ia, true, ePrefix)
}

// IncrementIntegerOne - Increment the value of the
// current intAry by adding '1'
func (ia *IntAry) IncrementIntegerOne() error {

  ePrefix := "IntAry.IncrementIntegerOne()"

  err := new(intAryElectron).isValidIntAry(
    ia,
    ePrefix+" Called on 'ia'")

  if err != nil {
    return err
  }

  if ia.isZeroValue || ia.isIntegerZeroValue {
    ia.signVal = 1
  }

  intLen := ia.intAryLen - ia.precision
  intIdx := intLen - 1
  lastIdx := ia.intAryLen - 1

  n1 := 0
  n2 := 0
  carry := 0

  ia.isZeroValue = true
  ia.isIntegerZeroValue = true

  for i := lastIdx; i >= 0; i-- {
    n1 = int(ia.intAry[i])

    if i > intIdx {
      //  i > intIdx
      // This must be a fractional digit
      // Retain fractional digits

      if n1 != 0 {
        ia.isZeroValue = false
      }

      continue

    } else if i == intIdx {

      n2 = n1 + (1 * ia.signVal)

      if n2 < 0 {
        n2 = n1 + 10 - 1
        carry = 1

      } else if n2 > 9 {
        n2 = n1 + 1 - 10
        carry = 1

      } else {
        carry = 0
      }

    } else {
      // Must be i < intIdx

      n2 = n1 + ((ia.signVal * carry) * 1)

      if n2 < 0 {
        n2 = n1 + 10 - carry
        carry = 1
      } else if n2 > 9 {
        n2 = n1 - 10 + carry
        carry = 1
      } else {
        carry = 0
      }

    }

    if n2 != 0 {
      ia.isZeroValue = false
      ia.isIntegerZeroValue = false
    }

    ia.intAry[i] = uint8(n2)

  }

  if ia.isZeroValue && carry == 0 {
    ia.signVal = 1
  }

  if carry > 0 {

    ia.intAry = append([]uint8{1}, ia.intAry...)
    ia.intAryLen++

  } else if ia.intAry[0] == 0 && intLen > 1 {
    ia.intAry = ia.intAry[1:]
    ia.intAryLen--
  }

  return nil
}

// IsValid - Examines the current intAry and returns
// an error if the intAry object is found to be invalid.
func (ia *IntAry) IsValid(errName string) error {

  if len(errName) == 0 {
    errName = "IntAry.IsValid()"
  }

  return new(intAryElectron).isValidIntAry(
    ia,
    errName)
}

// IsEvenNumber
//
//	Returns 'true' if the current IntAry numeric value is evenly
//	divisible by two (2) with no remainder.
//
//	Even Number Definition:
//	  https://www.mathsisfun.com/definitions/even-number.html
func (ia *IntAry) IsEvenNumber() (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.IsEvenNumber()",
    "")

  if err != nil {
    return false, err
  }

  return new(intAryMechanics).isIntAryEvenNumber(
    ia, true, ePrefix.XCpy("Is 'ia' Even Number"))
}

// IsMinusOne
//
//	Returns 'true' if the value of the current IntAry is minus one
//	(-1).
//
//	Examples:
//	=========
//
//	    Value      Result
//	    -----      ------
//	    -1.0        true
//	    -1          true
//	    -1.0000     true
//	    -1.0001     false
//	    -2.0        false
func (ia *IntAry) IsMinusOne() (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.IsMinusOne",
    "")

  if err != nil {
    return false, err
  }

  iaMinusOne, err := new(IntAry).NewOne(ia.precision)

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaMinusOne, err := new(IntAry).NewOne(ia.precision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = iaMinusOne.ChangeSign()

  isIaEqualToMinusOne, err := new(intAryAtom).equal(
    ia, true, &iaMinusOne, true, ePrefix.XCpy("'ia' == 'iaMinusOne' ?"))

  if isIaEqualToMinusOne {
    return true, nil
  }

  return false, nil
}

// IsOne - Returns 'true' if the value of the current
// IntAry is '1'.
//
// Examples:
// =========
//
// Value 			Result
// -----			------
// 1.0				true
// 1					true
// 1.0000			true
// 1.0001			false
// 2.0				false
func (ia *IntAry) IsOne() (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.IsOne",
    "")

  if err != nil {
    return false, err
  }

  iaOne, err := new(IntAry).NewOne(ia.precision)

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaOne, err := new(IntAry).NewOne(ia.precision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iaIsEqualOne, err := new(intAryAtom).equal(
    ia, true, &iaOne, false, ePrefix.XCpy("'ia' == 'iaOne' ?"))

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "iaIsEqualOne, err := new(intAryAtom).equal(\n" +
          "ia, true, &iaOne, false, ePrefix.XCpy('ia' == 'iaOne' ?))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if iaIsEqualOne {
    return true, nil
  }

  return false, nil
}

// IsZero - Analyzes the current IntAry to determine
// it is a zero value. If the IntAry is equal to a zero
// value, this method returns 'true'
//
// Even Number Definition:
//
//	https://www.mathsisfun.com/definitions/even-number.html
func (ia *IntAry) IsZero() (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.IsZero",
    "")

  if err != nil {
    return false, err
  }

  err = new(intAryElectron).isValidIntAry(
    ia,
    ePrefix.XCpy("Validating 'ia' of IsZero").String())

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia,\n" +
          "  ePrefix.XCpy(Validating 'ia' of IsZero).String()))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return ia.isZeroValue, nil

}

// Inverse
//
//	Returns the inverse BigIntNum of the current intAry's value.
//
//	  IntAry = current instance of IntAry
//	  Inverse = 1 ÷ IntAry
//
//	Input Parameter
//	===============
//
//	maxPrecision             int
//	  Determines the number of digits to the right of the decimal
//	  point in the result.
//
//	  If 'maxPrecision' is set equal to negative one (-1), the
//	  maximum number of decimal digits is automatically set to
//	  4096 digits to the right of the decimal	place.
func (ia *IntAry) Inverse(maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.IsOne",
    "")

  if err != nil {
    return IntAry{}, err
  }

  return new(intAryMechanics).inverseIntAry(ia, true, maxPrecision, ePrefix)
}

// MultiplyByTwoToPower
//
//		Multiply the existing value of the current IntAry instance by
//		2 to the power of the input parameter 'power'.
//
//	 Example
//	 =======
//
//	 IntAry Numeric Value x 2^power = result
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
//	Numeric Separators previous configured for the current IntAry
//	instance ('ia') will not be modified by this method.
func (ia *IntAry) MultiplyByTwoToPower(power uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.MultiplyByTwoToPower()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryNeutron).multiplyByTwoToPower(ia, true, power, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNeutron).\n" +
        "  multiplyByTwoToPower(ia, true, power, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// MultiplyByTenToPower
//
//	The value of intAry is multiplied by 10 to the power of the
//	input parameter 'power'.
//
//	Example
//	=======
//
//	IntAry Numeric Value x 10^power = result
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
//	Numeric Separators previous configured for the current IntAry
//	instance ('ia') will not be modified by this method.
func (ia *IntAry) MultiplyByTenToPower(power uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.MultiplyByTenToPower",
    "")

  if err != nil {
    return err
  }

  err = new(intAryNeutron).multiplyByTenToPower(ia, true, power, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNeutron).\n" +
        "  multiplyByTenToPower(ia, true, power, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// MultiplyThisBy
//
//	Multiplies the current IntAry instance ('ia') by IntAry input
//	parameter 'ia2' and stores the multiplication result in the
//	current IntAry instance ('ia').
//
//	Example
//	=======
//
//	  ia = ia x ia2
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
//	Numeric Separators previous configured for the current IntAry
//	instance ('ia') will not be modified by this method.
//
//	Input Parameters
//	================
//
//	ia2                      *IntAry
//	  Pointer to an IntAry object. In this multiplication
//	  operation, 'ia2' is the multiplier.
//
//	minimumResultPrecision   int
//	  'minimumResultPrecision' will determine the minimum number of
//	  digits computed to the right of the decimal place in the
//	  final result.
//
//	  If 'minimumResultPrecision' is set to a value of -1, all
//	  significant digits (digits greater than zero) will be
//	  returned to the right of the decimal place. Remember that
//	  the maximum number of decimal digits returned will be
//	  controlled by parameter 'maxResultPrecision'
//
//	maxResultPrecision       int
//	  'maxResultPrecision' will determine the maximum number of
//	  digits to the right of the decimal place in the result.
//
//	  Valid values are -1 and values >= zero ('0')
//
//	  Values less than -1 will trigger an error.
//
//	  A value of -1 signals that no limit will be placed on the
//	  number of decimals places to right of the decimal point in
//	  the result. Be advised that a very, very large number of
//	  decimal digits may be accommodated by the IntAry Type.
//
//	Return Value
//	============
//
//	error
//	  If no errors are encountered, this method will return an
//	  error value of 'nil'.
func (ia *IntAry) MultiplyThisBy(ia2 *IntAry, minimumPrecision, maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.MultiplyThisBy",
    "")

  if err != nil {
    return err
  }

  err = new(intAryNeutron).multiplyThisBy(
    ia, true, ia2, true, minimumPrecision, maxPrecision, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNeutron).multiplyThisBy(\n" +
        "  ia, true, ia2, true, minimumPrecision, maxPrecision, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// Multiply
//
//	This method receives three IntAry input parameters, 'ia1',
//	'ia2' and 'iaResult'.It then proceeds to multiply 'ia1' by
//	'ia2' and stores the multiplication result in intAry
//	'iaResult'.
//
//	Example
//	=======
//
//	      product = multiplicand x multipllier
//
//	      iaResult = ia1 x ia2
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
//	Numeric Separators for the calculated multiplication result
//	('iaResult') will be copied from current IntAry instance
//	('ia').
//
//	Input Parameters
//	================
//
//	ia1                      *IntAry
//	  In this multiplication operation, 'ia1' is the multiplicand.
//
//	ia2                      *IntAry
//	  In this multiplication operation, 'ia2' is the multiplier.
//
//	iaResult                 *IntAry
//
//	  This 'iaResult' IntArray object which will store the result
//	  of the multiplication operation. 'iaResult' is the 'product'.
//
//	minimumResultPrecision   int
//	  'minimumResultPrecision' will determine the minimum number of
//	  digits computed to the right of the decimal place in the
//	  final result or 'product'.
//
//	  If 'minimumResultPrecision' is set to a value of -1, all
//	  significant digits (digits greater than zero) will be
//	  returned to the right of the decimal place. Remember that
//	  the maximum number of decimal digits returned will be
//	  controlled by parameter 'maxResultPrecision'
//
//	maxResultPrecision       int
//	  'maxResultPrecision' will determine the maximum number of
//	  digits to the right of the decimal place in the result.
//
//	  Valid values are -1 and values >= zero ('0')
//
//	  Values less than -1 will trigger an error.
//
//	  A value of -1 signals that no limit will be placed on the
//	  number of decimals places to right of the decimal point in
//	  the result. Be advised that a very, very large number of
//	  decimal digits may be accommodated by the IntAry Type.
//
//	Return Value
//	============
//
//	error
//	  If no errors are encountered, this method will return an
//	  error value of 'nil'.
func (ia *IntAry) Multiply(
  ia1 *IntAry,
  ia2 *IntAry,
  iaResult *IntAry,
  minimumResultPrecision int,
  maxResultPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.Multiply",
    "")

  if err != nil {
    return err
  }

  err = new(intAryNeutron).
    multiply(ia1, true, ia2, true, iaResult,
      true, minimumResultPrecision, maxResultPrecision, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNeutron).multiply(\n" +
        "  ia1, true, ia2, true, iaResult,\n" +
        "  true, minimumResultPrecision, maxResultPrecision, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// New - Creates and returns a new blank intAry object.
//
// The returned IntAry instance will contain USA default numeric separators
// (decimal separator, thousands separator and currency symbol).
//
// Usage: ia := intAry{}.New()
func (ia *IntAry) New() IntAry {
  iAry := IntAry{}
  iAry.intAry = []uint8{}
  iAry.intAryLen = 0
  iAry.integerLen = 0
  iAry.significantIntegerLen = 0
  iAry.significantFractionLen = 0
  iAry.firstDigitIdx = -1
  iAry.lastDigitIdx = -1
  iAry.isZeroValue = true
  iAry.isIntegerZeroValue = true
  iAry.precision = 0
  iAry.signVal = 1
  iAry.decimalSeparator = '.'
  iAry.thousandsSeparator = ','
  iAry.currencySymbol = '$'
  iAry.BackUp = new(BackUpIntAry).New()

  return iAry
}

// NewWithNumSeps
//
//	 Creates and returns a new blank intAry object. The returned
//	 IntAry instance will contain numeric separators (decimal
//	 separator, thousands separator and currency symbol) as
//	 specified by input parameter, 'numSeps'.
//
//		Usage: ia := new(intAry).New(numSeps)
//
//		Note: If input parameter 'numSeps' is empty, it will be set to
//					default USA separators.
func (ia *IntAry) NewWithNumSeps(numSeps NumericSeparatorDto) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewWithNumSeps",
    "")

  if err != nil {
    return IntAry{}, err
  }

  numSeps.SetDefaultsIfEmpty()

  iAry := new(intAryElectron).newIntAry()

  err = iAry.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return iAry,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = iAry.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = new(intAryElectron).isValidIntAry(
    &iAry,
    ePrefix.XCpy("Validating 'iAry'").String())

  if err != nil {

    return iAry,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "&iAry, ePrefix.XCpy(\"Validating 'iAry'\").String())",
        ErrContext: "Newly generated IntAry object 'iAry' is INVALID!\n" +
          "'iAry' FAILED final validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewBigInt - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type '*big.Int'. Note that 'num' may be a positive
// or negative number.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place. 'precision'
// must be a positive value. Negative 'precision' values will
// trigger an error.
//
// Example:
//
//	intDigits     precision     	    result
//	946254  			   3							   946.254
//	946254				   0							   946254
//	-946254  			   3					      -946.254
//	-946254				   0						    -946254
//
// Usage:
// num := big.NewInt(123456)
// precision := uint(3)
// ia, err := intAry{}.NewBigInt(num, precision)
func (ia *IntAry) NewBigInt(num *big.Int, precision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewBigInt",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithBigInt(
    &iAry, nil, nsProfile, num, precision, true, ePrefix)

  return iAry, err
}

// NewBigIntNum - Creates a new intAry object initialized
// to the value of input parameter 'bINum' which is passed
// as type 'BigIntNum'.
//
// If the 'BigIntNum' precision value exceeds the positive
// maximum value of 'int' (32-bit integer = 2,147,483,647 ),
// an error will be returned.
//
// Usage:
// bINum, _ := new(BigIntNum).NewNumStr("1234.5678")
// ia, err := intAry{}.NewBigIntNim(bINum)
//
// The value of Numeric Separators contained in BigINum will NOT be copied
// into the current IntAry instance. IntAry will retain its
// current numeric separator values.
func (ia *IntAry) NewBigIntNum(bINum BigIntNum) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewBigIntNum",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryMinibot).setIntAryWithBigIntNum(
    &iAry, nil, nsProfile, &bINum, true, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryMinibot).setIntAryWithBigIntNum(\n" +
          "  &iAry, nil, nsProfile, &bINum,\n" +
          "  validateBINum=true, validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewFive
//
//	Creates a new IntAry object with a value of '5'.
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
//	for the current instance of IntAry to the new, returned
//	instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	IntAry
//	  This method returns a new IntAry object configured with the
//	  numeric value of five ('5') and the 'precision' and Numeric
//	  Separator specifications described above.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewFive(precision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewOne",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryQuark).setIntAryToFive(
    &iAry,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 5"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryQuark).setIntAryToFive(&iAry, nil, nsProfile, precisioin, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewFloat32 - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'float32'. Input parameter 'precision' is used
// to set the input precision of 'num'.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Usage:
// num := float32(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// ia, err := intAry{}.NewFloat32(num, precision)
//
// The value resulting from the usage example above would be '123.456'.
// If the value of 'precision' was set to -1, the resulting value would
// also be '123.456'
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding may occur.
func (ia *IntAry) NewFloat32(num float32, precision int) (IntAry, error) {

  ePrefix := "IntAry.NewFloat32()"

  iAry := new(intAryElectron).newIntAry()

  if precision < -1 {

    return iAry,
      fmt.Errorf("%v\n"+
        "Error: Input parameter 'precision' is INVALID!\n"+
        "'precision' must be greater than or equal to -1.\n"+
        "precision='%v'\n",
        ePrefix,
        precision)

  }

  err := iAry.SetIntAryWithFloat32(num, precision)

  if err != nil {
    return IntAry{},
      fmt.Errorf("%v\n"+
        "Error returned by iAry.SetIntAryWithFloat32(num, precision)\n"+
        "'num'='%v'\n"+
        "precision='%v'\n",
        ePrefix,
        num,
        precision)
  }

  err = new(intAryElectron).isValidIntAry(
    &iAry,
    ePrefix)

  return iAry, err
}

// NewFloat64
//
//	Creates a new intAry object initialized to the value of input parameter
//	'num' which is passed as type 'float64'. Input parameter 'precision' is
//	used to set the input precision of 'num'.
//
//	Input parameter 'precision' must be set to a number greater than or equal
//	to zero.  It may also be set to a value of -1 which causes the number
//	to be formatted to the smallest number of digits to right of the decimal
//	point.
//
//	Input parameter 'precision' must be set to a number greater than or equal
//	to zero.  It may also be set to a value of -1 which causes the number to be
//	formatted to the smallest number of digits to right of the decimal point.
//
//	Usage:
//		num := float64(123.456000)
//		precision := 3 - signals that only the decimals 456 will included as input
//		ia, err := intAry{}.NewFloat64(num, precision)
func (ia *IntAry) NewFloat64(num float64, precision int) (IntAry, error) {

  ePrefix := "IntAry.NewFloat64()"

  iAry := new(intAryElectron).newIntAry()

  if precision < -1 {

    return iAry,
      fmt.Errorf("%v\n"+
        "Error: Input parameter 'precision' is INVALID!\n"+
        "'precision' must be greater than or equal to -1.\n"+
        "precision='%v'",
        ePrefix,
        precision)

  }

  err := iAry.SetIntAryWithFloat64(num, precision)

  if err != nil {

    return IntAry{},
      fmt.Errorf("%v\n"+
        "Error returned by ia.SetIntAryWithFloat64(num, precision)\n"+
        "'num'='%v'\n"+
        "precision='%v'\n"+
        "Error= %v\n",
        ePrefix,
        num,
        precision,
        err.Error())

  }

  err = new(intAryElectron).isValidIntAry(
    &iAry,
    ePrefix)

  return iAry, err
}

// NewFloatBig - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type '*big.Float'.
//
// Input Parameters:
//
// num *big.Float -	This floating value will be used to
//
//	set the value of a new IntAry object
//
// precision int -	'precision' is applied to the input parameter
//
//	'num' (*big.Float) to determine the number of
//	digits to the right of the decimal point which
//	will be applied to the resulting value.
//
//	Input parameter 'precision' must be set to a
//	number greater than or equal to zero.  It may
//	also be set to a value of -1 which causes the
//	number to be formatted to the smallest number
//	of digits to right of the decimal point.
//
// Usage:
//
//	num, err := big.NewFloatBig(123.4560, 3)
//
// Result:
//
//	num = '123.456 - Note: precision parameter may result in rounding.
func (ia *IntAry) NewFloatBig(num *big.Float, precision int) (IntAry, error) {

  ePrefix := "IntAry.NewFloatBig()"

  iAry := new(intAryElectron).newIntAry()

  if num == nil {

    return iAry,
      fmt.Errorf("%v\n"+
        "Error: Input parameter 'num' is INVALID!\n"+
        "'num' is a nil pointer!\n",
        ePrefix)

  }

  if precision < -1 {

    return iAry,
      fmt.Errorf("%v\n"+
        "Error: Input parameter 'precision' is INVALID!\n"+
        "'precision' must be greater than or equal to -1.\n"+
        "precision='%v'\n",
        ePrefix,
        precision)

  }

  err := iAry.SetIntAryWithFloatBig(num, precision)

  if err != nil {

    return IntAry{},
      fmt.Errorf("%v\n"+
        "Error returned by iAry.SetIntAryWithBig(num, precision)\n"+
        "precision='%v'\n"+
        "Error= %v\n",
        ePrefix,
        precision,
        err.Error())

  }

  err = new(intAryElectron).isValidIntAry(
    &iAry,
    ePrefix)

  return iAry, err
}

// NewInt
//
//	Creates a new intAry object initialized to the value of input
//	parameter 'intNum' which is passed as type 'int'.
//
//	Input parameter 'precision' indicates the number of digits to
//	be formatted to the right of the decimal place. Input parameter
//	'precision' is of type uint. The maximum value allowed for
//	'precision' is 2,147,483,647 (the max int32 value). If
//	'precision' exceeds this maximum value, an error will be
//	returned.
//
//	Usage
//	=====
//
//	This method is designed to be used in conjunction with the
//	'new' keyword shown as follows:
//
//	    intNum := int(123456)
//	    precision := uint(3)
//	    iAry := new(IntAry).NewInt(intNum, precision)
//	    The numeric value of 'iAry' is now equal to 123.456
//
//	Examples
//	========
//
//	int Num    precision    IntAry Result
//	-------    ---------    -------------
//
//	123456         4           12.3456
//	123456         0           123456
//	123456         1           12345.6
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
//	with Numeric Separators copied from current instance of IntAry.
//
//	Input Parameters
//	================
//
//	intNum                   int
//	  The numeric digits contained in this value comprise both
//	  the integer digits and the fractional digits which will be
//	  configured in the final numeric value stored in the IntAry
//	  object returned by this method.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	 In practice, the maximum limit for 'precision' will be
//	 constrained by the maximum array size permitted by
//	 your system.
//
//	Return Values
//	=============
//
//	IntAry
//	  This new instance of IntAry will be returned configured with
//	  the numeric value calculated from input parameters, 'intNum'
//	  and 'precision'.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewInt(intNum int, precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewInt",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt(
    &iAry,
    nil,
    nsProfile,
    intNum,
    precision,
    true,
    ePrefix)

  return iAry, err
}

// NewIntExponent
//
//	Returns a new IntAry instance. The numeric value is set using
//	an int value multiplied by 10 raised to the power of the
//	'exponent' parameter.
//
//	    numeric value = int X 10^exponent
//
//	Usage
//	=====
//
//	This method may be used in conjunction with the 'new' keword
//	syntax.
//
//	    iAry := new(IntAry).NewIntExponent(123456, -3)
//	    -- iAry is now equal to "123.456", precision = 3
//
//	    iAry := new(IntAry).NewIntExponent(123456, 3)
//	    -- iAry is now equal to "123456.000", precision = 3
//
//	Examples
//	========
//
//	intNum      exponent      IntAry Result
//
//	123456         -3              123.456
//	123456          3           123456.000
//	123456          0           123456
//
//	Input Parameters
//	================
//
//	intNum                   int
//	  The numeric digits which will make up the returned IntAry
//	  numeric value.
//
//	exponent                 int
//	  This value will be used to determine the numeric digits in
//	  'intNum' which will be assigned to the right of the decimal
//	  point in the returned IntAry value.
//
//	Return Values
//	=============
//
//	IntAry
//	  This returned IntAry object will be configured with the
//	  numeric value computed from input parameters 'intNum' and
//	  'exponent' according to the conversion algorithm described
//	  above.
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (ia *IntAry) NewIntExponent(intNum int, exponent int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewIntExponent",
    "")

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryMinibot).setIntAryInt64Exponent(
    &iAry,
    nil,
    nsProfile,
    int64(intNum),
    exponent,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "new(intAryMinibot).setIntAryInt64Exponent(\n" +
          "  &iAry, nil, nsProfile, int64(intNum), exponent,\n" +
          "validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewInt32
//
//	Creates a new intAry object initialized to the value of input
//	parameter 'int32Num' which is passed as type 'int32'.
//
//	Input parameter 'precision' indicates the number of digits to
//	be formatted to the right of the decimal place. Input parameter
//	'precision' is of type uint. The maximum value allowed for
//	'precision' is 2147483645 (the max int32 value minus 2). If
//	'precision' exceeds this maximum value it will be reset to that
//	maximum value.
//
//	Usage
//	=====
//
//	This method may be used with the 'new' keyword syntax.
//
//	  int32Num := int64(123456)
//	  precision := uint(3)
//	  iAry := new(IntAry).NewInt32(int32Num, precision)
//	  iAry is now equal to 123.456
//
//	Examples
//	========
//
//	int32Num      precision      IntAry Result
//
//	 123456           4              12.3456
//	 123456           0              123456
//	 123456           1              12345.6
//
//	Input Parameters
//	================
//
//	int32Num                 int
//	  The numeric digits contained in this value comprise both
//	  the integer digits and the fractional digits which will be
//	  configured in the final numeric value stored in the IntAry
//	  object returned by this method.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  Although 'precision' is an unsigned integer type, the maximum
//	  value allowed for this parameter is 2,147,483,647. This
//	  is the maximum limit for a 32-bit integer which is the
//	  internal IntAry storage type for 'precision.
//
//	Return Values
//	=============
//
//	IntAry
//	  This new instance of IntAry will be returned configured with
//	  the numeric value calculated from input parameters, 'intNum'
//	  and 'precision'.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewInt32(int32Num int32, precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewInt32",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt(
    &iAry,
    ia,
    nsProfile,
    int(int32Num),
    precision,
    true,
    ePrefix)

  return iAry, err
}

// NewInt32Exponent
//
//	Returns a new IntAry instance. The numeric value is set using
//	an int32 value multiplied by 10 raised to the power of the
//	'exponent' parameter.
//
//	    numeric value = int32 X 10^exponent
//
//	Usage
//	=====
//
//	This method may be used in conjunction with the 'new' keword
//	syntax.
//
//	    iAry := new(IntAry).NewInt32Exponent(123456, -3)
//	    -- iAry is now equal to "123.456", precision = 3
//
//	    iAry := new(IntAry).NewInt32Exponent(123456, 3)
//	    -- iAry is now equal to "123456.000", precision = 3
//
//	Examples
//	========
//
//	intNum      exponent      IntAry Result
//
//	123456         -3              123.456
//	123456          3           123456.000
//	123456          0           123456
//
//	Input Parameters
//	================
//
//	intNum                   int
//	  The numeric digits which will make up the returned IntAry
//	  numeric value.
//
//	exponent                 int
//	  This value will be used to determine the numeric digits in
//	  'intNum' which will be assigned to the right of the decimal
//	  point in the returned IntAry value.
//
//	Return Values
//	=============
//
//	IntAry
//	  This returned IntAry object will be configured with the
//	  numeric value computed from input parameters 'intNum' and
//	  'exponent' according to the conversion algorithm described
//	  above.
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (ia *IntAry) NewInt32Exponent(int32Num int32, exponent int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewInt32Exponent",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryMinibot).setIntAryInt64Exponent(
    &iAry,
    nil,
    nsProfile,
    int64(int32Num),
    exponent,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "new(intAryMinibot).setIntAryInt64Exponent(\n" +
          "  &iAry, nil, nsProfile, int64(int32Num), exponent,\n" +
          "validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewInt64
//
//	Creates a new intAry object initialized to the value of input
//	parameter 'int64Num' which is passed as a type 'int64'.
//
//	Input parameter 'precision' indicates the number of digits to
//	be formatted to the right of the decimal place. Input parameter
//	precision' is of type uint. The maximum value allowed for
//	precision' is 2147483647 (the max int32 value). If 'precision'
//	exceeds this maximum value it will be reset to that maximum
//	value.
//
//	Usage
//	=====
//
//	This method may be used with the 'new' keyword syntax.
//
//	  int64Num := int64(123456)
//	  precision := uint(3)
//	  iAry := new(IntAry).NewInt64(int64Num, precision)
//	  -- iAry is now equal to 123.456 --
//
//	Examples
//	========
//
//	int64Num      precision      IntAry Result
//
//	123456            4            12.3456
//	123456            0            123456
//	123456            1            12345.6
//
//	Input Parameters
//	================
//
//	int64Num                 int64
//	  The numeric digits contained in this value comprise both
//	  the integer digits and the fractional digits which will be
//	  configured in the final numeric value stored in the IntAry
//	  object returned by this method.
//
//	precision uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  Although 'precision' is an unsigned integer type, the maximum
//	  value allowed for this parameter is 2,147,483,647. This
//	  is the maximum limit for a 32-bit integer which is the
//	  internal IntAry storage type for 'precision.
//
//	Return Values
//	=============
//
//	IntAry
//	  This returned IntAry object will be configured with the
//	  numeric value computed from input parameters 'int64Num' and
//	  'precision' according to the conversion algorithm described
//	  above.
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (ia *IntAry) NewInt64(int64Num int64, precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithUint64",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt64(
    &iAry,
    nil,
    nsProfile,
    int64Num,
    precision,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryGluon).setIntAryWithInt64(\n" +
          "  &iAry, nil, nsProfile, int64Num, precision,\n" +
          "  validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewInt64Exponent
//
//	 Returns a new IntAry instance. The numeric value is set using
//		an int64 value multiplied by 10 raised to the power of the
//		'exponent' parameter.
//
//	     numeric value = int64 X 10^exponent
//
//		Usage
//		=====
//
//		This method is may be used with the 'new' keyword syntax.
//
//		  iAry := new(IntAry).NewInt64Exponent(123456, -3)
//		  -- iAry is now equal to "123.456", exponent = 3
//
//		  iAry := new(IntAry).NewInt64Exponent(123456, 3)
//		  -- iAry is now equal to "123456.000", exponent = 3
//
//		Examples
//		========
//
//		uint64Num      exponent      IntAry Result
//
//		  123456          -3           123.456
//		  123456           3           123456.000
//		  123456           0           123456
//
//		IMPORTANT
//		=========
//
//	 In practice, the maximum limits for 'int64Num' and
//	 'exponent' will be constrained by the maximum array size
//	 permitted by your system. Type IntAry relies on arrays of
//	 8-bit integers for numeric value storage.
//
//		Input Parameters
//		================
//
//		int64Num                 int64
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in the IntAry
//		  object returned by this method.
//
//		exponent                 int
//		  'exponent' specifies the number of fractional digits in the
//		  final numeric value stored in the returned IntAry object.
//	   See the examples above.
//
//		Return Values
//		=============
//
//		IntAry
//		  This new instance of IntAry will be returned configured with
//		  the numeric value calculated from input parameters, 'intNum'
//		  and 'precision'.
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (ia *IntAry) NewInt64Exponent(int64Num int64, exponent int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewInt64Exponent()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryMinibot).setIntAryInt64Exponent(
    &iAry, nil, nsProfile, int64Num, exponent, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryMinibot).setIntAryInt64Exponent(\n" +
          "  &iAry, nil, nsProfile, int64Num, exponent, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, err
}

// NewIntFracStr
//
//	Creates a new IntAry instance based on a numeric value
//	represented by separate integer and fractional components.
//
//	Input parameters 'intStr' and 'fracStr' are strings
//	representing the integer and fractional components. They are
//	combined by this method to create a numeric value which is
//	assigned to the new IntAry instance.
//
//	Input parameter 'signVal' must be set to one of two values:
//	+1 or -1. This value is used to signal the sign of the
//	resulting numeric value. +1 generates a positive number and -1
//	generates a negative number. If input parameters 'inStr' or
//	'fracStr' contain a leading minus or plus sign character, it
//	will be ignored. The sign of the resulting numeric value is
//	controlled strictly by input parameter, 'signVal'.
func (ia *IntAry) NewIntFracStr(intStr, fracStr string, signVal int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewIntFracStr()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = new(intAryPhoton).setNumericSeparatorsToDefaultIfEmpty(
    &iAry, ePrefix.XCpy("Set 'iAry' Num Seps To Default if Empty"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err := new(intAryPhoton).\n" +
          "  setNumericSeparatorsToDefaultIfEmpty(&iAry,\n" +
          "  ePrefix.XCpy(\"Set 'iAry' Num Seps To Default if Empty\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryNeutron).setIntAryWithIntFracStr(&iAry, false, nil, nsProfile, intStr, fracStr, signVal, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNeutron).setIntAryWithIntFracStr(\n" +
          "  &iAry, validateiAry=false, nil, nsProfile, intStr, fracStr, signVal, validateResult=true, ePrefix)",
        ErrContext: fmt.Sprintf("intStr= '%v'\nfracStr= '%v'\nsignVal= '%v'",
          intStr, fracStr, signVal),
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewNumStr
//
//		Creates a new intAry object initialized to the value of input
//		parameter 'numStr', a string of numbers and delimiters which is
//		passed as type 'string'.
//
//		This method assumes that the input parameter 'numStr' is a
//		string of numeric digits which may be delimited by default
//		USA numeric separators. Default USA numeric separators are
//		defined as:
//
//			  decimal separator = '.'
//			  thousands separator = ','
//			  currency symbol = '$'
//	 The USA default numeric separators will be used to parse the
//	 number string 'numStr'.
//
//		If the subject 'numStr' employs other national or cultural
//		numeric separators, see method IntAry.NewNumStrWithNumSeps(),
//		below.
//
//		Usage
//		=====
//
//		  ia := new(IntAry).NewNumStr("123.456")
func (ia *IntAry) NewNumStr(numStr string) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewNumStr",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  numSeps := new(NumericSeparatorDto).NewUSADefaults()

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          numSeps,
  }

  err = new(intAryQuark).
    setIntAryWithNumStr(&iAry, false, nil,
      nsProfile, numStr, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryWithNumStr(\n"+
          "  &iAry, validateiAry=false, nil, nsProfile, numStr,\n"+
          "  validateResult=true, ePrefix)\n"+
          "  numStr= '%v'", numStr),
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewNumStrWithNumSeps
//
//	Receives a number string as input and returns a new IntAry
//	instance. The input parameter 'numSeps' contains numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) which will be used to parse the number string.
//
// In addition, the numeric separators contained in input parameter
// 'numSeps' will be copied to the returned IntAry instance.
func (ia *IntAry) NewNumStrWithNumSeps(
  numStr string,
  numSeps NumericSeparatorDto) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewNumStrWithNumSeps",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(\n" +
          "  ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FALIED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          numSeps,
  }

  err = new(intAryQuark).setIntAryWithNumStr(&iAry, false, nil, nsProfile, numStr, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryWithNumStr(\n"+
          "  &iAry, validateiAry=false, nil, nsProfile, numStr,\n"+
          "  validateResult=true, ePrefix)\n"+
          "  numStr= '%v'", numStr),
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewNumStrMaxPrecision
//
//	Creates a new intAry object initialized to the value of input
//	parameter 'numStr' which is passed as type 'string'. 'numStr'
//	is assumed to be a valid number string
//
//	Input parameter 'numSeps' consists of a type
//	NumericSeparatorDto. This structure encapsulate numeric
//	serparators which will be used to parse the number string
//	passed by input parameter, 'numStr'.
//
//	Numeric separators specify the decimal separator, thousands
//	separator and currency symbol which are essential for parsing
//	number strings.
//
//	If the number of decimal points to the right of the decimal
//	place exceeds input parameter 'maxPrecision', the resulting
//	numeric value will be rounded to 'maxPrecision' decimal places.
//
// Usage: ia := intAry{}.NewNumStr("123.456", 3)
func (ia *IntAry) NewNumStrMaxPrecision(
  numStr string,
  numSeps NumericSeparatorDto,
  maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewNumStrMaxPrecision",
    "")

  if err != nil {
    return IntAry{}, err
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(\n" +
          "  ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FALIED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          numSeps,
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryNeutron).
    setIntAryWithNumStrMaxPrecision(&iAry, false, nil,
      nsProfile, numStr, maxPrecision, true, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: fmt.Sprintf("err = new(intAryNeutron).setIntAryWithNumStrMaxPrecision(\n"+
          "  &iAry, validateiAry=false, nil, nsProfile, numStr,\n"+
          "  maxPrecision, validateResult=true, ePrefix)\n"+
          "  numStr= '%v' maxPrecision= '%v'", numStr, maxPrecision),
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewNumStrDto - Creates, initializes and returns an IntAry
// Type using an input parameter of Type NumStrDto.
func (ia *IntAry) NewNumStrDto(numDto NumStrDto) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewNumStrDto",
    "")

  if err != nil {
    return IntAry{}, err
  }

  err = numDto.IsValid(ePrefix.XCpy("Validating 'numDto'").String())

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numDto.IsValid(ePrefix.XCpy(Validating 'numDto').String())",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iAry := new(intAryElectron).newIntAry()

  numDtoNumStr, err := numDto.GetNumStr()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numDtoNumStr, err := numDto.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = iAry.SetIntAryWithNumStr(numDtoNumStr)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = iAry.SetIntAryWithNumStr(numDtoNumStr)",
        ErrContext: fmt.Sprintf("numDtoNumStr= '%v' ", numDtoNumStr),
        ErrMessage: err.Error(),
      }
  }

  err = new(intAryElectron).isValidIntAry(
    &iAry, ePrefix.XCpy("Validating 'iAry' Final Result").String())

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
          "&iAry, ePrefix.XCpy(Validating 'iAry' Final Result).String())",
        ErrContext: "Error: Final Result 'iAry' is INVALID!\n" +
          "'iAry' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewOne
//
//	Creates a new IntAry object with a value of '1'.
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
//	for the current instance of IntAry to the new, returned
//	instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	IntAry
//	  This method returns a new IntAry object configured with the
//	  numeric value of one ('1') and the 'precision' and Numeric
//	  Separator specifications described above.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewOne(precision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewOne",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryQuark).setIntAryToOne(
    &iAry,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 1"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryQuark).setIntAryToOne(&iAry, nil, nsProfile, precisioin, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewPtr - Returns a pointer to a new IntAry instance.
func (ia *IntAry) NewPtr() *IntAry {

  ia2 := new(intAryElectron).newIntAry()

  return &ia2
}

// NewTen
//
//	Creates a new IntAry object with a value of '1'.
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
//	for the current instance of IntAry to the new, returned
//	instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	IntAry
//	  This method returns a new IntAry object configured with the
//	  numeric value of ten ('10') and the 'precision' and Numeric
//	  Separator specifications described above.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewTen(precision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewTen",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryQuark).setIntAryToTen(
    &iAry,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 10"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryQuark).setIntAryToTen(\n" +
          "&iAry, nil, nsProfile, precisioin, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewThree
//
//	Creates a new IntAry instance with a value of '3'.
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
//	for the current instance of IntAry to the new, returned
//	instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	IntAry
//	  This method returns a new IntAry object configured with the
//	  numeric value of three ('3') and the 'precision' and Numeric
//	  Separator specifications described above.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewThree(precision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewThree",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(\n" +
          "  &iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryQuark).setIntAryToThree(
    &iAry,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 3"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryQuark).setIntAryToThree(&iAry, nil, nsProfile, precisioin, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewTwo
//
//	Creates a new IntAry instance with a value of '2'.
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
//	for the current instance of IntAry to the new, returned
//	instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	IntAry
//	  This method returns a new IntAry object configured with the
//	  numeric value of two ('2') and the 'precision' and Numeric
//	  Separator specifications described above.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) NewTwo(precision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewTwo",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  err = new(intAryQuark).setIntAryToTwo(
    &iAry,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 2"))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryQuark).setIntAryToTwo(\n" +
          "&iAry, nil, nsProfile, precisioin, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, err
}

// NewUint - Creates a new intAry object initialized to the
// value of input parameter 'uintNum' which is passed as type
// 'uint'.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place. Input
// parameter 'precision' is of type uint. The maximum value
// allowed for 'precision' is 2147483645 (the max int32 value
// minus 2). If 'precision' exceeds this maximum value it will
// be reset to that maximum value.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the
// IntAry{} syntax thereby allowing IntAry type creation and
// initialization in one step.
//
//					uintNum := uint(123456)
//					precision := uint(3)
//					iAry := IntAry{}.NewUint(uintNum, precision)
//	       iAry is now equal to 123.456
//
// Examples:
// ---------
//
//	 uintNum			precision			 IntAry Result
//		 123456		 		   4							12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (ia *IntAry) NewUint(
  uintNum uint, signVal int, precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewUint",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if signVal != 1 && signVal != -1 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'signVal' is INVALID!\n" +
          "'signVal' must be set to +1 or -1\n" +
          fmt.Sprintf("signVal= '%v'", signVal),
      }

  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithUint64(
    &iAry,
    nil,
    nsProfile,
    uint64(uintNum),
    signVal,
    precision,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryGluon).setIntAryWithUint64(\n" +
          "  &iAry, nil, nsProfile, uint64(uintNum), signVal, precision,\n" +
          "  validateFinalResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, err
}

// NewUintExponent - Returns a new IntAry instance. The numeric
// value is set using an uint value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = uint X 10^exponent
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the IntAry{}
// syntax thereby allowing IntAry type creation and initialization in
// one step.
//
//		iAry := IntAry{}.NewUintExponent(123456, -3)
//	 -- iAry is now equal to "123.456", precision = 3
//
//		iAry := IntAry{}.NewUintExponent(123456, 3)
//	 -- iAry is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 uintNum		    exponent		  	IntAry Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (ia *IntAry) NewUintExponent(uintNum uint, signValue int, exponent int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewUintExponent",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryMinibot).setIntAryUint64Exponent(
    &iAry,
    nil,
    nsProfile,
    uint64(uintNum),
    signValue,
    exponent,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setIntAryUint64Exponent(\n" +
          "  &iAry, nil, nsProfile, uint64Num, signValue,\n" +
          "  exponent, validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, err
}

// NewUint32
//
//		Creates a new intAry object initialized to the value of input
//		parameter 'uint32Num' which is passed as type 'uint32'.
//
//		Input parameter 'precision' indicates the number of digits
//		to be formatted to the right of the decimal place. Input
//		parameter 'precision' is of type uint.
//
//		Usage
//		=====
//
//		This method may be used with the 'new' keyword syntax.
//
//		    uint32Num := uint32(123456)
//		    precision := uint(3)
//		    iAry := new(IntAry).NewUint32(uint32Num, precision)
//		    -- iAry is now equal to 123.456 --
//
//		Examples
//		========
//
//		uint32Num      precision      IntAry Result
//
//		 123456            4             12.3456
//		 123456            0             123456
//		 123456            1             12345.6
//
//		Input Parameters
//		================
//
//		uint32Num                uint32
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in IntAry object
//		  returned by this method.
//
//		signValue                int
//		 Input parameter 'signValue' must be set to one of two values:
//		 +1 or -1. This value is used to signal the sign of the
//		 resulting numeric value. +1 identifies a positive number and
//		 -1 identifies a negative number. 'signValue' determines the
//		 numeric sign of the resulting IntAry value, either plus or
//		 minus.
//
//		precision                uint
//		  'precision' specifies the number of fractional digits in the
//		  final numeric value stored in 'intAry'
//
//	   In practice, the maximum limit for 'precision' will be
//	   constrained by the maximum array size permitted by
//	   your system.
//
//		Return Values
//		=============
//
//		IntAry
//		  This new instance of IntAry will be returned configured with
//		  the numeric value calculated from input parameters, 'intNum'
//		  'signVal' and 'precision'.
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (ia *IntAry) NewUint32(uint32Num uint32, signValue int, precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewUint32()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithUint64(
    &iAry,
    nil,
    nsProfile,
    uint64(uint32Num),
    signValue,
    precision,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryGluon).setIntAryWithUint64(\n" +
          "  &iAry, nil, nsProfile, uint64Num, signValue, precision,\n" +
          "  ValidateFinalResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewUint32Exponent
//
//	Returns a new IntAry instance based on input parameters,
//	'uint32Num', 'signValue' and 'exponent'.
//
//	The returned IntAry numeric value is set using an uint64 value
//	multiplied by 10 raised to the power of the 'exponent' parameter.
//
//		    Result Numeric Value = uint32Num X 10^exponent
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
//	uint32Num      exponent      IntAry Result
//
//	  123456          -3              123.456
//	  123456           3           123456.000
//	  123456           0           123456
//
//	Input Parameters
//	================
//
//	uint32Num                uint32
//	  The uint32 value holds the numeric digits which will make up the
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
//	  numeric value computed from input parameters 'uint32Num',
//	  signValue and 'exponent'.
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (ia *IntAry) NewUint32Exponent(uint32Num uint32, signValue int, exponent int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewUint64Exponent",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryMinibot).setIntAryUint64Exponent(
    &iAry,
    nil,
    nsProfile,
    uint64(uint32Num),
    signValue,
    exponent,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setIntAryUint64Exponent(\n" +
          "  &iAry, nil, nsProfile, uint64(uint32Num), signValue,\n" +
          "  exponent, validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewUint64
//
//		Creates a new intAry object initialized to the value of input
//		parameters 'uint64Num', 'signValue', and 'precision'.
//
//		Input parameter 'precision' indicates the number of digits to
//		be formatted to the right of the decimal place. Input
//		parameter 'precision' is of type uint. The maximum value
//		allowed for 'precision' is 2,147,483,647 (the max int32 value).
//		If 'precision' exceeds this maximum value it will be reset to
//		that maximum value.
//
//		Usage
//		=====
//
//		This method is designed to be used in conjunction with the
//		'new' keyword shown as follows:
//
//		    uint64Num := uint64(123456)
//		    precision := uint(3)
//		    iAry := new(IntAry).NewUint64(
//		                      uint64Num, signVal, precision)
//		    The numeric value of 'iAry' is now equal to 123.456
//
//		Examples
//		========
//
//		uint64Num    precision    IntAry Result
//		---------    ---------    -------------
//
//		 123456          4           12.3456
//		 123456          0           123456
//		 123456          1           12345.6
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
//		The IntAry object returned by this method will be configured
//		with Numeric Separators copied from current instance of IntAry.
//
//		Input Parameters
//		================
//
//		uint64Num                uint64
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in IntAry object
//		  returned by this method.
//
//		signValue                int
//		 Input parameter 'signValue' must be set to one of two values:
//		 +1 or -1. This value is used to signal the sign of the
//		 resulting numeric value. +1 identifies a positive number and
//		 -1 identifies a negative number. 'signValue' determines the
//		 numeric sign of the resulting IntAry value, either plus or
//		 minus.
//
//		precision                uint
//		  'precision' specifies the number of fractional digits in the
//		  final numeric value stored in 'intAry'
//
//	   In practice, the maximum limit for 'precision' will be
//	   constrained by the maximum array size permitted by
//	   your system.
//
//		Return Values
//		=============
//
//		IntAry
//		  This new instance of IntAry will be returned configured with
//		  the numeric value calculated from input parameters, 'intNum'
//		  'signVal' and 'precision'.
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (ia *IntAry) NewUint64(
  uint64Num uint64, signValue int, precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewUint64()",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithUint64(
    &iAry,
    nil,
    nsProfile,
    uint64Num,
    signValue,
    precision,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryGluon).setIntAryWithUint64(\n" +
          "  &iAry, nil, nsProfile, uint64Num, signValue, precision,\n" +
          "  ValidateFinalResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewUint64Exponent
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
//	  The uint64 value holds the numeric digits which will make up
//	  the returned IntAry numeric value.
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
func (ia *IntAry) NewUint64Exponent(
  uint64Num uint64, signValue int, exponent int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewUint64Exponent",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryMinibot).setIntAryUint64Exponent(
    &iAry,
    nil,
    nsProfile,
    uint64Num,
    signValue,
    exponent,
    true,
    ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setIntAryUint64Exponent(\n" +
          "  &iAry, nil, nsProfile, uint64Num, signValue,\n" +
          "  exponent, validateResult=true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iAry, nil
}

// NewZero - Creates a new IntAry instance and sets
// the value to zero.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place. Input
// parameter 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the
// IntAry{} syntax thereby allowing IntAry type creation and
// initialization in one step.
//
//		iAry := IntAry{}.NewZero(0)
//	 -- iAry is now equal to "0", precision = 0
//
//		iAry := IntAry{}.NewZero(3)
//	 -- iAry is now equal to "0.000", precision = 3
func (ia *IntAry) NewZero(precision uint) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.NewZero",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iAry := new(intAryElectron).newIntAry()

  err = new(intAryProton).copy(&iAry, ia, false, false, ePrefix)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryProton).copy(&iAry, ia, true, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt(
    &iAry,
    nil,
    nsProfile,
    0,
    precision,
    true,
    ePrefix)

  return iAry, err
}

// OptimizeIntArrayLen
//
//	Eliminates Leading zeros from the front or integer portion
//	of the integer string.
//
//	If parameter 'optimizeFracDigits' is set equal to 'true',
//	trailing zeros to the right of the decimal place will also be
//	eliminated.
//
//	Validation Testing
//	==================
//
//	This method will perform validation testing on the current
//	instance of IntAry
func (ia *IntAry) OptimizeIntArrayLen(optimizeFracDigits bool) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.OptimizeIntArrayLen",
    "")

  if err != nil {
    return err
  }

  err = new(intAryAtom).optimizeIntArrayLen(ia, true, optimizeFracDigits, false, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryAtom).optimizeIntArrayLen(\n"+
        "  ia, validateIntAry=true, optimizeFracDigits=%v,\n"+
        "   validateResult=false, ePrefix)",
        optimizeFracDigits),
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// Pow
//
//	Raises the value of the current intAry to the power designated
//	by the input parameter 'power'. This method uses the
//	'Power By Twos' technique which raises a number to a specified
//	power using the exponentiation by squaring algorithm.
//
// Input Parameters:
// =================
// 'power' int -
//
//	The input parameter 'power' may be either
//	a positive or negative integer.
//
// 'maxResultPrecision' int -
//
//					'maxResultPrecision' will determine the maximum
//					number of digits to the right of the decimal
//					place in the result.
//
//					Valid values are -1 and values >= zero ('0')
//	       Values less than -1 will trigger an error.
//
//					A value of -1 signals that no limit will be placed on
//					the number of decimals places to right of the decimal
//					point in the result.
//
// TODO Determine if internal precision is required
//
//	'internalPrecision' int -
//				'internalPrecision' will control the number of digits of
//				accuracy to the right of the decimal point maintained by
//				internal multiplication operations used in raising the intAry
//				value to the designated power.
//
//				Valid values are -1 and values >= zero ('0')
//				Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point during internal multiplication operations.
func (ia *IntAry) Pow(power int, maxResultPrecision int, internalPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.Pow",
    "")

  if err != nil {
    return err
  }

  if internalPrecision < -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "internalPrecision < -1",
      ErrMessage: fmt.Sprintf("Error: Parameter 'internalPrecision' is less than -1.\n"+
        "internalPrecision= '%v'", internalPrecision),
    }
  }

  iaPower := new(intAryElectron).newIntAry()

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt(
    &iaPower,
    ia,
    nsProfile,
    power,
    0,
    false,
    ePrefix)

  err = new(IntAryMathPower).Pwr(ia, &iaPower, 0, maxResultPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathPower).Pwr(ia, &iaPower, 0, maxResultPrecision)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// PowThisSquared - Raises the value of the current intAry object
// to a power of '2'.  Essentially the new value of this intAry object
// is equal to the original value squared.
func (ia *IntAry) PowThisSquared() error {

  err := ia.MultiplyThisBy(ia, ia.GetPrecision(), -1)

  if err != nil {
    return fmt.Errorf("IntAry.PowThisSquared\n"+
      "Error returned by IntAry.MultiplyThisBy().\n"+
      "Error='%v'\n",
      err.Error())
  }

  return nil
}

// PowByTwos - Raises the value of the current intAry to the power indicated by the parameter,
// 'power'.
//
// Input Parameters:
// =================
// 'power' int -
//
//	The input parameter 'power' may be either
//	a positive or negative integer.
//
// 'maxResultPrecision' int -
//
//					'maxResultPrecision' will determine the maximum
//					number of digits to the right of the decimal
//					place in the result.
//
//					Valid values are -1 and values >= zero ('0')
//	       Values less than -1 will trigger an error.
//
//					A value of -1 signals that no limit will be placed on
//					the number of decimals places to right of the decimal
//					point in the result.
//
//		'internalPrecision' int -
//					'internalPrecision' will control the number of digits of
//					accuracy to the right of the decimal point maintained by
//					internal multiplication operations used in raising the intAry
//					value to the designated power.
//
//					Valid values are -1 and values >= zero ('0')
//	       Values less than -1 will trigger an error.
//
//					A value of -1 signals that no limit will be placed on
//					the number of decimals places to right of the decimal
//					point during internal multiplication operations.
func (ia *IntAry) PowByTwos(power *big.Int, maxResultPrecision, internalPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.PowByTwos",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "iAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  return new(intAryMinibot).pwrByTwos(ia, true, nil, nsProfile, power, maxResultPrecision, internalPrecision, true, ePrefix)
}

// PrefixToIntAry
//
//	Adds an integer of type uint8 to the front of the existing
//	internal integer array maintained by the current IntAry object.
func (ia *IntAry) PrefixToIntAry(num uint8) {

  new(intAryMolecule).prefixToIntAry(ia, num)

  return
}

// ResetFromBackUp
//
//	Retrieves data from the last saved backup and populates the
//	current intAry instance.
func (ia *IntAry) ResetFromBackUp() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.ResetFromBackUp",
    "")

  if err != nil {
    return err
  }

  return new(intAryBoson).resetFromBackUp(ia, ePrefix)
}

// RoundToPrecision
//
//	Rounds the value of the current IntAry instance to a precision
//	specified by the 'roundToPrecision' parameter.
//
//	TODO - Find a better way to round to zero precision
func (ia *IntAry) RoundToPrecision(roundToPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.RoundToPrecision",
    "")

  if err != nil {
    return err
  }

  err = new(intAryMolecule).roundToPrecision(
    ia,
    true,
    roundToPrecision,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryMolecule).roundToPrecision(\n"+
        "ia, validateIa='true',\n"+
        "roundPrecision= '%v', ePrefix",
        roundToPrecision),
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetAbsoluteValueThis - Converts the current
// value of this intAry object to its
// absolute value.
func (ia *IntAry) SetAbsoluteValueThis() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetAbsoluteValueThis()",
    "")

  if err != nil {
    return err
  }

  return new(intAryMechanics).setAbsoluteValue(ia, ePrefix.XCpy("Setting 'ia' to absolute value"))
}

// SetCurrencySymbol is used to set the value of the currency
// separator character. The currency separator character is used
// to designate currency when formatting the value of this IntAry
// is expressed as a string. In the US, the currency is the dollar
// character ('$'). In the following example, the '$' character
// designates a currency value. Example - '$1,000,000,000.00' .
//
// The SetCurrencySymbol method can be used to change the currency
// separator character in accordance with the customs of countries other
// than the US. For example, in England use of the pound sterling currency
// symbol ('£') would be appropriate.
//
// To examine the current setting for currency symbol, see
// the method GetCurrencySymbol().
//
// Note the default thousands separator character is the comma (',').
func (ia *IntAry) SetCurrencySymbol(currencySymbol rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetCurrencySymbol()",
    "")

  if err != nil {
    return err
  }

  return new(intAryPhoton).
    setNumSepSymbol(ia,
      CURRENCYSYMBOL,
      currencySymbol,
      ePrefix.XCpy("Setting 'ia' currency symbol"))
}

// SetDecimalSeparator - sets the decimal separator character
// (usually a decimal point) which is used to separate
// integer digits from fractional digits when the value
// of this IntAry is expressed as a string.
//
// For US usage, the '.' in the following example is
// the decimal separator. Example: 123.456
//
// SetDecimalSeparator can be used to change the default
// decimal separator to a value other than '.'
//
// To examine the current setting for decimal separator, see
// method GetDecimalSeparator().
//
// Note: The default decimal separator character is '.'
func (ia *IntAry) SetDecimalSeparator(decimalSeparator rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetDecimalSeparator()",
    "")

  if err != nil {
    return err
  }

  return new(intAryPhoton).
    setNumSepSymbol(ia,
      DECIMALSYMBOL,
      decimalSeparator,
      ePrefix.XCpy("Setting 'ia' decimal symbol"))
}

// SetElement - Sets the value of an IntAry element
// at a given index.
func (ia *IntAry) SetElement(index, val int) error {

  ePrefix := "IntAry.SetElement() "

  if val > math.MaxUint8 {

    return fmt.Errorf("%v\n"+
      "Error: Input parameter 'val' Exceeds Maximum for Unsigned Integer!\n"+
      "MaxUint8='%v'",
      ePrefix,
      math.MaxUint8)

  }

  if val < 0 {

    return fmt.Errorf("%v\n"+
      "Error: Input parameter 'val' is less than ZERO!\n"+
      "val='%v'\n", ePrefix, val)

  }

  if index < 0 {

    return fmt.Errorf("%v\n"+
      "Error: Input parameter 'index' is less than ZERO!\n"+
      "index='%v'\n", ePrefix, index)
  }

  if index > ia.GetIntAryLength()-1 {

    return fmt.Errorf("%v\n"+
      "Error: Input parameter 'index' EXCEEDS Int Array Length!\n"+
      "index='%v'\n"+
      "Actual Int Array Length='%v'\n",
      ePrefix,
      index,
      ia.GetIntAryLength()-1)
  }

  ia.intAry[index] = uint8(val)

  return nil
}

// SetEqualArrayLengths
//
//	Compares an intAry object to the current intAry and ensures
//	that the lengths of both IntArrays are equal.
func (ia *IntAry) SetEqualArrayLengths(iAry2 *IntAry) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetEqualArrayLengths",
    "")

  if err != nil {
    return err
  }

  return new(intAryAtom).setEqualArrayLengths(
    ia, false, iAry2, false, false, ePrefix)
}

// SetIntAryLength
//
//	Calculates the current IntAry string length and sets internal
//	variable 'ia.intAryLen'.
func (ia *IntAry) SetIntAryLength() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryLength()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryElectron).setIntAryLength(ia, ePrefix.XCpy("Setting 'ia' IntAry Length"))

  return err
}

// SetIntAryToFive
//
//	Sets the value of the intAry object to one ('5').
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
//	This method will not alter the Numeric Separators configured
//	for the current instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryToFive(precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToFive",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  return new(intAryQuark).setIntAryToFive(
    ia,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 5"))
}

// SetIntAryToOne
//
//	Sets the value of the intAry object to one ('1').
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
//	This method will not alter the Numeric Separators configured
//	for the current instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryToOne(precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToOne",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  return new(intAryQuark).setIntAryToOne(
    ia,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 1"))
}

// SetIntAryToTwo
//
//	Sets the value of the current IntAry object to two ('2').
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
//	This method will not alter the Numeric Separators configured
//	for the current instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryToTwo(precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToTwo",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  return new(intAryQuark).setIntAryToTwo(
    ia,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 2"))
}

// SetIntAryToThree
//
//	Sets the value of the intAry object to three ('3').
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
//	This method will not alter the Numeric Separators configured
//	for the current instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryToThree(precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToThree",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  return new(intAryQuark).setIntAryToThree(
    ia,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 3"))
}

// SetIntAryToTen
//
//	Sets the value of the current IntAry object to ten ('10').
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
//	This method will not alter the Numeric Separators configured
//	for the current instance of IntAry.
//
//	Input Parameters
//	================
//
//	precision                int
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the returned IntAry object.
//
//	  If the value of 'precision' is less than zero, an error will
//	  be returned.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryToTen(precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToTen",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  return new(intAryQuark).setIntAryToTen(
    ia,
    nil,
    nsProfile,
    precision,
    ePrefix.XCpy("Set 'ia' = 10"))
}

// SetIntAryToZero
//
//		Sets the value of the current IntAry object to zero ('0').
//
//	 The existing Numeric Separators configured for the current
//	 IntAry instance will remain unchanged.
func (ia *IntAry) SetIntAryToZero(precision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryToZero",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryQuark).setIntAryToZero(
    ia,
    nil,
    nsProfile,
    precision,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryToZero(\n"+
        "ia, nil, nsProfile, precision= '%v', ePrefix", precision),
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithInt
//
//	Sets the value of the current intAry object to that of the
//	input parameter 'intDigits', an integer of type 'int'.
//
//	Input parameter 'precision' to indicate the number of digits to
//	the right of the decimal place. Input parameter 'precision' is
//	of type uint.
//
//	The numeric sign (plus or minus) of the resulting intAry value
//	is determined by the sign of input parameter,'intDigits'.
//
//	Example
//	=======
//
//	intDigits      precision      result
//	---------      ---------      ------
//
//	  946254            3          946.254
//	  946254            0          946254
//	 -946254            3         -946.254
//	 -946254            0         -946254
//
//	Input Parameters
//	================
//
//	intDigits                int
//	  The numeric digits contained in this value comprise both
//	  the integer digits and the fractional digits which will be
//	  configured in the final numeric value stored in the current
//	  instance of IntAry.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in the current instance of
//	  'IntAry'.
//
//	  Although 'precision' is an unsigned integer type, the maximum
//	  value allowed for this parameter is 2,147,483,647.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryWithInt(intDigits int, precision uint) error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithInt",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
  }

  return new(intAryGluon).setIntAryWithInt(
    ia,
    nil,
    nsProfile,
    intDigits,
    precision,
    true,
    ePrefix)
}

// SetIntAryWithInt32 - Sets the value of the current intAry object
// to that of the input parameter 'intDigits', a 32-bit integer.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place. Input
// parameter 'precision' is of type uint. The maximum value
// allowed for 'precision' is 2147483645 (the max int32 value
// minus 2). If 'precision' exceeds this maximum value it will
// be reset to that maximum value.
//
// The numeric sign (plus or minus) of the resulting intAry value
// is determined by the sign of input parameter 'int32Num'.
//
// Example:
//
//	int32Num     precision     	       result
//	946254  			   3							   946.254
//	946254				   0							   946254
//	-946254  			   3					      -946.254
//	-946254				   0						    -946254
func (ia *IntAry) SetIntAryWithInt32(int32Num int32, precision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithInt32",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt(
    ia,
    nil,
    nsProfile,
    int(int32Num),
    precision,
    true,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryGluon).setIntAryWithInt(\n" +
        "ia, nil, nsProfile, int(int32Num), precision,\n" +
        "validateResult=true, ePrefix",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithInt64
//
//		Sets the value of the current intAry object to that of the
//		input parameter 'int64Num', a 64-bit integer.
//
//		Input parameter 'precision' indicates the number of digits
//		to be formatted to the right of the decimal place. Input
//		parameter 'precision' is of type uint.
//
//	  In practice, the maximum limit for 'precision' will be
//	  constrained by the maximum array size permitted by
//	  your system.
//
//		The numeric sign (plus or minus) of the resulting intAry value
//		is determined by the sign of input parameter 'int64Num'.
//
//		Example
//		=======
//
//		int64Num      precision      result
//
//		 946254           3           946.254
//		 946254           0           946254
//		-946254           3          -946.254
//		-946254           0          -946254
//
//		Input Parameters
//		================
//
//		int64Num                 int64
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in the IntAry
//		  object returned by this method.
//
//		precision                uint
//		  'precision' specifies the number of fractional digits in the
//		  final numeric value stored in the returned IntAry object.
//
//	    In practice, the maximum limit for 'precision' will be
//	    constrained by the maximum array size permitted by
//	    your system.
//
//		Return Values
//		=============
//
//		IntAry
//		  This new instance of IntAry will be returned configured with
//		  the numeric value calculated from input parameters, 'intNum'
//		  and 'precision'.
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (ia *IntAry) SetIntAryWithInt64(int64Num int64, precision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithInt64()",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryGluon).setIntAryWithInt64(
    ia,
    nil,
    nsProfile,
    int64Num,
    precision,
    true,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryGluon).setIntAryWithInt64(\n" +
        "  ia, nil, nsProfile, int64Num, precision,\n" +
        "  validateResult=true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithIntFracStr
//
//	Sets the value of the current IntAry instance based on a
//	numeric value represented by separate integer and fractional
//	components.
//
//	Input parameters 'intStr' and 'fracStr' are strings
//	representing the integer and fractional components. They are
//	combined by this method to create a numeric value which is
//	assigned to, and stored in, the current IntAry instance.
//
//	Input parameter 'signVal' must be set to one of two values:
//	+1 or -1. This value is used to signal the sign of the
//	resulting numeric value. +1 generates a positive number and -1
//	generates a negative number. If input parameters 'inStr' or
//	'fracStr' contain a leading minus or plus sign character, it
//	will be ignored. The sign of the resulting numeric value is
//	controlled strictly by input parameter, 'signVal'.
//
//	IMPORTANT
//	=========
//
//	This method will use the numeric separators in the current
//	instance of IntAry to convert the integer and fractional
//	components into a consolidated number string for internal
//	calculation purposes.
func (ia *IntAry) SetIntAryWithIntFracStr(intStr, fracStr string, signVal int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithIntFracStr",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryNeutron).setIntAryWithIntFracStr(ia, false, nil, nsProfile, intStr, fracStr, signVal, true, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNeutron).setIntAryWithIntFracStr(\n" +
        "  ia, validateIa=false, nil, nsProfile, intStr, fracStr, signVal, validateResult=true, ePrefix)",
      ErrContext: fmt.Sprintf("intStr= '%v'\nfracStr= '%v'\nsignVal= '%v'",
        intStr, fracStr, signVal),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithUint64
//
//	Sets the value of the current IntAry object equal to that of
//	the input parameter 'intDigits', a 64-bit unsigned integer.
//
//	Note: Input parameter 'precision' to indicate the number of
//	digits to the right of the decimal place.
//
//	Input parameter, 'signVal' must be set to one of two values:
//	-1 or +1. 'signVal' determines the numeric sign of the
//	resulting IntAry value, either plus or minus.
//
//	Example
//	=======
//
//	intDigits  precision  signVal    result
//
//	 946254        3         1       946.254
//	 946254        0         1       946254
//	 946254        3        -1      -946.254
//	 946254        0        -1      -946254
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
//	with Numeric Separators copied from current instance of IntAry.
//
//	Input Parameters
//	================
//
//	uint64Num                uint64
//	  The numeric digits contained in this value comprise both
//	  the integer digits and the fractional digits which will be
//	  configured in the final numeric value stored in IntAry object
//	  returned by this method.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in 'intAry'
//
//	  Although 'precision' is an unsigned integer type, the maximum
//	  value allowed for this parameter is 2,147,483,647.
//
//	signVal                  int
//	 Input parameter 'signVal' must be set to one of two values:
//	 +1 or -1. This value is used to signal the sign of the
//	 resulting numeric value. +1 identifies a positive number and
//	 -1 identifies a negative number. 'signVal' determines the
//	 numeric sign of the resulting IntAry value, either plus or
//	 minus.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in 'intAry'
//
//	  Although 'precision' is an unsigned integer type, the maximum
//	  value allowed for this parameter is 2,147,483,647.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (ia *IntAry) SetIntAryWithUint64(
  intDigits uint64, signVal int, precision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithUint64",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  return new(intAryGluon).setIntAryWithUint64(
    ia,
    nil,
    nsProfile,
    intDigits,
    signVal,
    precision,
    true,
    ePrefix)
}

// SetIntAryWithBigInt
//
//	Sets the current value of the intAry to the value of input
//	parameter 'intDigits'. The sign value (plus or minus) is taken
//	from the input parameter, 'intDigits'.
//
//	The precision or number of digits to the right of the decimal
//	point, is determined by the input parameter, 'precision'. Input
//	parameter 'precision' must be passed as a positive value.
//
//	Negative 'precision' values will trigger an error.
//
//	Example
//	=======
//
//	intDigits      precision      result
//
//	 946254            3           946.254
//	 946254            0           946254
//	-946254            3          -946.254
//	-946254            0          -946254
func (ia *IntAry) SetIntAryWithBigInt(intDigits *big.Int, precision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithBigInt",
    "")

  if err != nil {
    return err
  }

  if intDigits == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intDigits'",
    }
  }

  if precision < 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is a negative value!\n"+
        "precision='%v'", precision),
    }

  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  return new(intAryGluon).setIntAryWithBigInt(
    ia, nil, nsProfile, intDigits, precision, true, ePrefix)
}

// SetIntAryWithBigIntNum - Sets the current value of the intAry to the value
// of input parameter 'bigINum', a BigIntNum type.
//
// If the input parameter 'BigIntNum' precision value exceeds the positive
// maximum value of 'int' (32-bit integer = 2,147,483,647 ), an error will
// be returned.
//
// The value of Numeric Separators contained in BigINum will NOT be copied
// into the current IntAry instance. IntAry will retain its
// current numeric separator values.
func (ia *IntAry) SetIntAryWithBigIntNum(bigINum BigIntNum) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithBigIntNum",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryMinibot).setIntAryWithBigIntNum(
    ia, nil, nsProfile, &bigINum, true, true, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryMinibot).setIntAryWithBigIntNum(\n" +
        "  ia, nil, nsProfile, &bigINum,\n" +
        "  validateBINum=true, validateResult=true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithDecimal - Sets the value of the current IntAry
// instance to the numeric value of input parameter, 'dec' which
// is of Type Decimal.
//
// The value of Numeric Separators contained in 'dec' will NOT be
// copied into the current IntAry instance. IntAry will retain its
// current numeric separator values.
func (ia *IntAry) SetIntAryWithDecimal(dec Decimal) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithDecimal",
    "")

  if err != nil {
    return err
  }

  err = dec.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = dec.IsValid(ePrefix.XCpy(Validating 'dec').String())",
      ErrContext: "Error: Input parameter 'dec' is INVALID!\n" +
        "'dec' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  decNumStr, err := dec.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "decNumStr, err := dec.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = ia.SetIntAryWithNumStr(decNumStr)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "err = ia.SetIntAryWithNumStr(decNumStr)",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithFloat32 - Sets the current value of the intAry based on the
// input parameter, 'floatNum'. 'floatNum is of type float32, a 32-bit
// floating point number.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Usage:
// num := float32(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// err := ia.SetIntAryWithFloat32(num, precision)
//
// The value resulting from the usage example above would be '123.456'.
// If the value of 'precision' was set to -1, the resulting value would
// also be '123.456'
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding may occur.
func (ia *IntAry) SetIntAryWithFloat32(floatNum float32, precision int) error {

  ePrefix := "IntAry.SetIntAryWithFloat32()"

  if precision < -1 {

    return fmt.Errorf("%v\n"+
      "Input Parameter 'precision' invalid.\n"+
      "'precision' must be greater than or equal to -1.\n"+
      "precison= %v\n",
      ePrefix,
      precision)

  }

  numStr := strconv.FormatFloat(float64(floatNum), 'f', -1, 32)

  err := ia.SetIntAryWithNumStr(numStr)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned from ia.SetIntAryWithNumStr(numStrDto).\n"+
      "numStrDto= '%v'\nError=  %v\n",
      ePrefix,
      numStr,
      err.Error())
  }

  if precision > -1 {

    err = ia.SetPrecision(precision, true)

    if err != nil {

      return fmt.Errorf("%v\n"+
        "Error returned from ia.SetPrecision(precision, true).\n"+
        "precision='%v'\nError= '%v'\n",
        ePrefix,
        precision,
        err.Error())
    }

  }

  return nil
}

// SetIntAryWithFloat64 - Sets the current value of the intAry based on the
// input parameter, 'floatNum'. 'floatNum' is of type float64, a 64-bit
// floating point number.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding	may occur.
func (ia *IntAry) SetIntAryWithFloat64(floatNum float64, precision int) error {

  ePrefix := "IntAry.SetIntAryWithFloat64()"

  if precision < -1 {

    return fmt.Errorf("%v\n"+
      "Error: Invalid input parameter 'precision'.\n"+
      "'precision' must be greater than or equal to -1.\n"+
      "precision= '%v'",
      ePrefix,
      precision)
  }

  numStr := strconv.FormatFloat(floatNum, 'f', -1, 64)

  err := ia.SetIntAryWithNumStr(numStr)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned from ia.SetIntAryWithNumStr(numStrDto).\n"+
      "numStrDto='%v'\nError= %v\n",
      ePrefix,
      numStr,
      err.Error())

  }

  if precision > -1 {

    err = ia.SetPrecision(precision, true)

    if err != nil {

      return fmt.Errorf("%v\n"+
        "Error returned from ia.SetPrecision(precision, true).\n"+
        "precision='%v'\nError= %v\n",
        ePrefix,
        precision,
        err.Error())

    }

  }

  return nil
}

// SetIntAryWithFloatBig - Sets the current value of the intAry based on the
// input parameter, 'floatNum'.  'floatNum' is of type *big.Float.
//
// Input Parameters:
//
// floatNum *big.Float - 	The floating point value to which the current IntAry object
//
//	will be set.
//
// precision int 			 - 	The number of digits to the right of the decimal point
//
//	which will be used to set the value of the current IntAry
//	object
//
//	Input parameter 'precision' must be set to a number greater
//	than or equal to zero.  It may also be set to a value of -1
//	which causes the number to be formatted to the smallest number
//	of digits to right of the decimal point.
//
//	Note: if precision is a positive value, and less than the current
//	number of digits to the right of the decimal place, rounding may occur.
func (ia *IntAry) SetIntAryWithFloatBig(floatNum *big.Float, precision int) error {

  ePrefix := "IntAry.SetIntAryWithFloatBig()"

  if precision < -1 {

    return fmt.Errorf("%v\n"+
      "Error: Invalid input parameter 'precision'.\n"+
      "'precision' must be greater than -1."+
      "precision= '%v'",
      ePrefix,
      precision)

  }

  var numStr string

  if precision == -1 {
    numStr = floatNum.Text('f', precision)
  } else {
    numStr = floatNum.Text('f', precision+1)
  }

  err := ia.SetIntAryWithNumStr(numStr)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned from ia.SetIntAryWithNumStr(numStrDto)."+
      "numStrDto='%v'\nError= %v\n",
      ePrefix,
      numStr,
      err.Error())

  }

  if precision > -1 {

    err = ia.SetPrecision(precision, true)

    if err != nil {

      return fmt.Errorf("%v\n"+
        "Error returned from ia.SetPrecision(precision, true).\n"+
        "precision='%v'\nError= %v\n",
        ePrefix,
        precision,
        err.Error())
    }

  }

  return nil
}

// SetIntAryWithIntAry - Sets the value of the current intAry based on an array of
// integers i.e. []int.
//
// Input parameter 'precision' will determine the number of digits to the right
// of the decimal place.
//
// Input parameter 'signVal' must be either +1 or -1 indicating the sign of the
// number represented by the integer array. Note: If signVal is not equal to +1 or -1,
// an error is generated.
func (ia *IntAry) SetIntAryWithIntAry(iAry2 []int, precision uint, signVal int) error {

  if signVal != 1 && signVal != -1 {

    return fmt.Errorf("SetIntAryWithIntAry()\n"+
      "Error: Input parameter 'signVal' parameter is INVALID!\n"+
      "signVal must be -1 or +1.\n"+
      "signVal='%v'", signVal)

  }

  lIAry2 := len(iAry2)

  ia.intAry = make([]uint8, lIAry2)

  for i := 0; i < lIAry2; i++ {

    ia.intAry[i] = uint8(iAry2[i])
  }

  ia.intAryLen = lIAry2

  ia.precision = int(precision)

  ia.signVal = signVal

  ia.SetInternalFlags()

  return nil
}

// SetIntAryWithUint8Ary
//
//		This method is designed to set the value of the current IntAry
//		object by passing in a pointer to an unsigned integer array
//		([]uint8). []uint8 is the type of array native to the IntAry
//		object.
//
//		Input parameter 'precision' will determine the number of digits
//		to the right of the decimal place.
//
//		Input parameter 'signVal' must be either +1 or -1 indicating
//		the	sign of the number represented by the integer array.
//
//		If signVal is not equal to +1 or -1, an error will be
//		generated.
//
//	 The Numeric Separators originaly configured for the current
//	 IntAry instance will NOT be modified.
func (ia *IntAry) SetIntAryWithUint8Ary(iAry2 []uint8, precision uint, signVal int) error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithUint8Ary()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryGluon).setIntAryWithUint8Ary(
    ia, iAry2, precision, signVal, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryGluon).setIntAryWithUint8Ary(\n" +
        "  ia, iAry2, precision, signVal, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithIntAryObj
// Sets the value of the current IntAry with a pointer to an IntAry object.
func (ia *IntAry) SetIntAryWithIntAryObj(iAry2 *IntAry, copyBackup bool) error {

  ePrefix := "IntAry.SetIntAryWithIntAryObj()"

  err := new(intAryElectron).isValidIntAry(
    iAry2,
    ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Input parameter 'iAry2' is INVALID!\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())

  }

  ia.CopyIn(iAry2, copyBackup)

  return nil
}

// SetIntAryWithNumStrMaxPrecision
//
//	Receives a raw number string and sets the fields of the
//	internal intAry structure to the appropriate values.
//
//	A second input parameter specifies the maximum allowable
//	precision for the IntAry. If IntAry precision exceeds
//	'maxPrecision', IntAry precision is rounded to 'maxPrecision'.
//
//	The term 'precision' defines the number of numeric digits to
//	the right of the decimal point or decimal separator.
func (ia *IntAry) SetIntAryWithNumStrMaxPrecision(numStr string, maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithNumStrMaxPrecision()",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryNeutron).
    setIntAryWithNumStrMaxPrecision(ia, false, nil,
      nsProfile, numStr, maxPrecision, true, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryNeutron).setIntAryWithNumStrMaxPrecision(\n"+
        "  ia, validateIa=false, nil, nsProfile, numStr,\n"+
        "  maxPrecision, validateResult=true, ePrefix)\n"+
        "  numStr= '%v' maxPrecision= '%v'", numStr, maxPrecision),
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithNumStr
//
//		Receives a raw number string and sets the fields of the
//		internal intAry structure to the appropriate values.
//
//	 The numeric separators (decimal separator, thousands separator,
//	 and currency symbol) used to parse the number string are
//	 taken from the current instance of IntAry ('ia').
func (ia *IntAry) SetIntAryWithNumStr(numStr string) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithNumStr()",
    "")

  if err != nil {
    return err
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "ia",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          NumericSeparatorDto{},
  }

  err = new(intAryQuark).setIntAryWithNumStr(ia, false, nil, nsProfile, numStr, true, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryWithNumStr(\n"+
        "  ia, validateIa=false, nil, nsProfile, numStr,\n"+
        "  validateResult=true, ePrefix)\n"+
        "  numStr= '%v'", numStr),
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetIntAryWithNumStrDto - Sets the numeric value of the current
// IntAry instance to that of input parameter 'nDto' which is of
// Type NumStrDto.
//
// Note that the Numeric Separator values are NOT copied into the
// current IntAry instance. The current IntAry numeric separators
// remain unchanged.
func (ia *IntAry) SetIntAryWithNumStrDto(nDto NumStrDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetIntAryWithNumStrDto",
    "")

  if err != nil {
    return err
  }

  err = nDto.IsValid(ePrefix.XCpy("Validating 'nDto'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nDto.IsValid(ePrefix.XCpy(Validating 'nDto').String())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nDtoNumStr, err := nDto.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nDtoNumStr, err := nDto.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = ia.SetIntAryWithNumStr(nDtoNumStr)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = ia.SetIntAryWithNumStr(nDtoNumStr)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetInternalFlags
//
//	Sets Array Lengths and tests for zero values
func (ia *IntAry) SetInternalFlags() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetInternalFlags",
    "")

  if err != nil {
    return err
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

  return nil
}

// SetInternalFlagsNoError
//
//	Sets Array Lengths and tests for zero values.
//
//	This method is identical to IntAry.SetInternalFlags();
//	however, this method does NOT return and error.
//
//	This method is depreicated but retained for backward
//	compatibility. Know that the preferred method is:
//	   IntAry.SetInternalFlags()
func (ia *IntAry) SetInternalFlagsNoError() {

  new(intAryNanobot).setInternalFlagsNoErrors(ia)

  return
}

// SetIsZeroValue
//
//	Analyzes the value of the intAry and sets a flag if the value
//	of intAry evaluates to zero.
func (ia *IntAry) SetIsZeroValue() {

  new(intAryNanobot).setIsZeroValue(ia)
}

// SetNumericSeparators
//
//	Used to assign values for the Decimal and Thousands separators
//	as well as the Currency Symbol to be used in displaying the
//	current number string.
//
//	Different nations and cultures use different symbols to delimit
//	numerical values. In the USA and many other countries, a period
//	character ('.') is used to delimit integer and fractional digits
//	within a numeric value (123.45). Likewise, thousands may be
//	delimited by a comma (','). Currency signs very by nationality.
//	For instance, the USA, Canada and several other countries use
//	the dollar sign ($) as a currency symbol.
//
//	For a list of major world currency symbols see:
//
//	  MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//
//	  http://www.xe.com/symbols.php
//
//	Note: If zero values are submitted as input for separator values,
//	those values will default to USA standards.
//
//	USA Examples
//	============
//
//	Decimal Separator period ('.')    = 123.456
//	Thousands Separator comma (',')   = 1,000,000,000
//	Currency Symbol dollar sign ('$') = $123
func (ia *IntAry) SetNumericSeparators(
  decimalSeparator,
  thousandsSeparator,
  currencySymbol rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetNumericSeparators()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryPhoton).setNumericSeparators(
    ia,
    decimalSeparator,
    thousandsSeparator,
    currencySymbol,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryMechanics).setNumericSeparators(\n" +
        "  ia, decimalSeparator, thousandsSeparator, currencySymbol, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetNumericSeparatorsToDefaultIfEmpty
//
//	If numeric separators are set to zero or nil, this method will
//	set those numeric separators to the USA defaults. This means
//	that the Decimal separator is set to a period ('.'), the
//	Thousands separator is set to a comma (',') and the currency
//	symbol is set to the dollar sign ('$').
//
//	If the numeric separators were previously set to a value other
//	than zero or nil, that value is not altered by this method.
//
//	Effectively, this method ensures that numeric separators are
//	set to valid values.
func (ia *IntAry) SetNumericSeparatorsToDefaultIfEmpty() error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetNumericSeparatorsToDefaultIfEmpty()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryPhoton).setNumericSeparatorsToDefaultIfEmpty(
    ia, ePrefix.XCpy("Set 'ia' Num Seps To Default if Empty"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err := new(intAryPhoton).\n" +
        "  setNumericSeparatorsToDefaultIfEmpty(ia,\n" +
        "  ePrefix.XCpy(\"Set 'ia' Num Seps To Default if Empty\"))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetNumericSeparatorsToUSADefault
//
//	Sets Numeric Separators to the United States of America (USA)
//	defaults.
//
//	Sets Numeric separators:
//	  Decimal Point Separator
//	  Thousands Separator
//	  Currency Symbol
//
// Call specific methods to set numeric separators for other countries or
// cultures:
//
//	ia.SetDecimalSeparator()
//	ia.SetThousandsSeparator()
//	ia.SetCurrencySymbol()
func (ia *IntAry) SetNumericSeparatorsToUSADefault() error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetNumericSeparatorsToUSADefault()",
    "")

  if err != nil {
    return err
  }
  err = new(intAryPhoton).setNumericSeparatorsToUSADefault(
    ia, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err := new(intAryQuark).\n" +
        " setNumericSeparatorsToUSADefault(ia, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetNumericSeparatorsDto
//
// Sets the values of numeric separators:
//
//	  decimal point separator
//	  thousands separator
//	  currency symbol
//
//	These vales are based on data transmitted through input
//	parameter 'customSeparators' of type NumericSeparatorDto.
//
//	If any of the values contained in input parameter
//	'customSeparators' are set to zero, an error will be returned.
func (ia *IntAry) SetNumericSeparatorsDto(customSeparators NumericSeparatorDto) error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetNumericSeparatorsDto()",
    "")

  if err != nil {
    return err
  }

  return new(intAryPhoton).setNumericSeparatorsDto(
    ia, customSeparators, true, ePrefix)
}

// SetPrecision
//
//	Sets the precision of the current IntAry instance to the value
//	of input parameter 'precision'.
//
//	If input parameter 'roundResult' is set to 'true', the
//	resulting numeric value will be rounded to the number of
//	decimal places specified by input parameter 'precision'.
//
//	If input parameter 'roundResult' is set to 'false', the
//	resulting numeric value will be truncated to the number of
//	decimal places specified by input parameter 'precision'.
//
//	If 'precision' is set to a value less than zero, an error will
//	be returned.
//
//	If 'precision' is greater than the existing precision, trailing
//	zeros will be added
//
//	Examples
//	========
//
//	   Original           'newPrecision'         Resulting
//	    Value             input parameter          Value
//	--------------        ---------------      --------------
//
//	  654.123456                 9              654.123456000
//	  654.123456                 4              654.1235
//
//	 -654.123456                 9             -654.123456000
//	 -654.123456                 4             -654.1235
//
//	    0                        3                0.000
//	    0.000000                 0                0
func (ia *IntAry) SetPrecision(precision int, roundResult bool) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetPrecision",
    "")

  if err != nil {
    return err
  }

  return new(intAryMolecule).
    setPrecision(ia, true, precision, roundResult, ePrefix)
}

// SetSign - Can be used to change the sign value
// of the current intAry value. The new sign value
// will be set according the input parameter, 'signVal'.
//
// 'signVal' has only two valid values, -1 or +1. If
// any value other than -1 or +1 is detected, an error
// will be thrown.
func (ia *IntAry) SetSign(signVal int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetSign()",
    "")

  if err != nil {
    return err
  }

  return new(intAryNeutron).setSign(ia, signVal, ePrefix)
}

// SetSignificantDigitIdxs - Finds the first
// significant digit (the first numeric digit
// greater than zero) and sets index value in
// the local field variable, 'firstDigitIdx'.
//
// In addition, this method also identifies the
// Last Significant Digit (the last non-zero value
// in the intAry) and records that index in the
// local field variable, 'lastDigitIdx'.
func (ia *IntAry) SetSignificantDigitIdxs() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetSignificantDigitIdxs",
    "")

  if err != nil {
    return err
  }

  err = ia.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = ia.SetNumericSeparatorsToDefaultIfEmpty()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  ia.intAryLen = len(ia.intAry)

  if ia.intAryLen == ia.precision {

    ia.intAry = append([]uint8{0}, ia.intAry...)

    ia.intAryLen++
  }

  if ia.intAryLen < ia.precision {

    deltaZeros := ia.precision - ia.intAryLen + 1

    zeroAry := make([]uint8, deltaZeros)

    ia.intAry = append(zeroAry, ia.intAry...)

    ia.intAryLen += deltaZeros
  }

  ia.firstDigitIdx = -1
  ia.lastDigitIdx = -1

  ia.integerLen = 0
  ia.significantIntegerLen = 0
  ia.significantFractionLen = 0

  lastIntIdx := ia.intAryLen - ia.precision - 1
  ia.isZeroValue = true
  ia.isIntegerZeroValue = true
  ia.integerLen = ia.intAryLen - ia.precision

  for i := 0; i < ia.intAryLen; i++ {
    if ia.intAry[i] > 0 {

      ia.isZeroValue = false

      if i < ia.integerLen {

        ia.isIntegerZeroValue = false
      }
    }

    // At minimum, there should be a single
    // leading zero before the decimal point.
    // Example 0.000.
    if i == lastIntIdx && ia.intAry[i] == 0 {

      if ia.firstDigitIdx == -1 {

        ia.firstDigitIdx = i
      }

    }

    if ia.intAry[i] > 0 {

      if ia.firstDigitIdx == -1 {

        ia.firstDigitIdx = i
      }

      ia.lastDigitIdx = i
    }

  }

  ia.significantIntegerLen = ia.intAryLen - ia.precision - ia.firstDigitIdx

  if ia.lastDigitIdx >= ia.integerLen {

    ia.significantFractionLen = ia.lastDigitIdx - ia.integerLen + 1
  } else {

    ia.significantFractionLen = 0
  }

  return nil
}

// SetSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current intAry number string.
//
// Different nations and cultures use different symbols to delimit numerical values. In the
// USA and many other countries, a period character ('.') is used to delimit integer and
// fractional digits within a numeric value (123.45). Likewise, thousands may be delimited
// by a comma (','). Currency signs very by nationality. For instance, the USA, Canada and
// several other countries use the dollar sign ($) as a currency symbol.
//
// For a list of major world currency symbols see:
//
//		MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//	 http://www.xe.com/symbols.php
//
// Note: If zero values are submitted as input for separator values, those values will default
// to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
func (ia *IntAry) SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol rune) error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetSeparators()",
    "")

  if err != nil {
    return err
  }

  return new(intAryPhoton).setNumericSeparators(
    ia,
    decimalSeparator,
    thousandsSeparator,
    currencySymbol,
    ePrefix.XCpy("Setting 'ia' Numeric Separators"))
}

// ShiftPrecisionLeft
//
//	Shifts the relative position of a decimal point within a number
//	string. The position of the decimal point is shifted
//	'shiftPrecision' positions to the left of the current decimal
//	point position.
//
//	This is equivalent to:
//	    result = signedNumStr / 10^precision or signedNumStr divided
//	             by 10 raised to the power of precision.
//
//	See the Examples section below.
//
//	Input Parameters
//	================
//
//	shiftPrecision           uint
//	  The number of digits by which the current decimal point
//	  position in the current IntAry numeric value, will be shifted
//	  to the left.
//
//	Return Values
//	=============
//
//	error
//	  If processing errors are encountered, an appropriate error
//	  message will be configured through this return parameter.
//
//	Examples
//	========
//
//	                    Requested
//	                      Shift
//	    signedNumStr    precision      Result
//	     "123456.789"       3           "123.456789"
//	     "123456.789"       2          "1234.56789"
//	     "123456.789"       6             "0.123456789"
//	  "123456789"           6           "123.456789"
//	        "123"           5             "0.00123"
//	          "0"           3             "0.000"
//	          "0.000"       2             "0.00000"
//	    "-123456.789"       3          "-123.456789"
//	 "-123456789"           6          "-123.456789"
//
//	          zero 'shiftPrecision' has no effect
//	             on the original number string
//
//	     "123456.789"       0         "123456.789"
//	    "-123456.789"       0         "-123.456789"
func (ia *IntAry) ShiftPrecisionLeft(shiftPrecision uint) error {
  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.ShiftPrecisionLeft()",
    "")

  if err != nil {
    return err
  }

  err = new(IntAryMathDivide).DivideByTenToPower(ia, shiftPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathDivide).DivideByTenToPower(ia, shiftPrecision)",
      ErrContext: fmt.Sprintf("shiftPrecision= '%v'", shiftPrecision),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// ShiftPrecisionRight
//
//	Shifts the existing precision of the current IntAry numeric
//	value. The position of the decimal point is shifted 'shiftPrecision'
//	positions to the right.
//
//	This is equivalent to:
//	   result = IntAry X 10^shiftPrecision
//	                   or
//	   IntAry Multiplied by 10 raised to the power of
//	   input parameter 'shiftPrecision'.
//
//	Examples:
//	=========
//
//	                  Input
//	                 Parameter
//	IntAry Value   shiftPrecision    Result
//	------------   --------------    ------
//
//	 "123456.789"        3           "123456789"
//	 "123456.789"        2           "12345678.9"
//	 "123456.789"        6           "123456789000"
//	 "123456789"         6           "123456789000000"
//	 "123"               5           "12300000"
//	 "0"                 3           "0"
//
//	 zero shiftPrecision has no effect on original number string
//
//	 "123456.789"        0           "123456.789"
//	"-123456.789"        0           "-123456.789"
//	"-123456.789"        3           "-123456789"
//	"-123456789"         6           "-123456789000000"
//
//	Input Parameters:
//	=================
//
//	shiftPrecision            uint
//	   The number of digits by which the current decimal point
//	   position in the current IntAry numeric value will be shifted
//	   to the right.
func (ia *IntAry) ShiftPrecisionRight(shiftPrecision uint) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.ShiftPrecisionRight",
    "")

  if err != nil {
    return err
  }

  err = new(IntAryMathMultiply).MultiplyByTenToPower(ia, shiftPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathMultiply).MultiplyByTenToPower(ia, shiftPrecision)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SetThousandsSeparator is used to set the value of the thousands
// separator character. The thousands separator character is used
// to separate thousands when the value of this IntAry object is
// expressed as a string. In the US, the thousands separator character
// is a comma. In the following example, the ',' character separates
// thousands in the integer number: Example - '1,000,000,000'.
//
// The SetThousandsSeparator method can be used to change the thousands
// separator character in accordance with the customs of countries other
// than the US.
//
// To examine the current setting for thousands separator, see
// method GetIntAryStats().
//
// Note the default thousands separator character is the comma (',').
func (ia *IntAry) SetThousandsSeparator(thousandsSeparator rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAry.SetDecimalSeparator()",
    "")

  if err != nil {
    return err
  }

  return new(intAryPhoton).
    setNumSepSymbol(ia,
      THOUSANDSYMBOL,
      thousandsSeparator,
      ePrefix.XCpy("Setting 'ia' thousands symbol"))
}

// SuffixToIntAry
//
//	Adds an integer of type uint8 to the end of the existing
//	internal integer array maintained by the current IntAry object.
func (ia *IntAry) SuffixToIntAry(num uint8) {

  new(intAryMolecule).suffixToIntAry(ia, num)

}

// String - Returns the numeric value of the current IntAry as
// a string.
func (ia *IntAry) String() string {

  str, err := ia.GetNumStr()

  if err != nil {
    return ""
  }

  return str
}

// SubtractFromThis - Subtracts the value of parameter
// 'ia2' from the current intAry object.
//
// Input Parameters:
// =================
//
// ia2 *intAry - Incoming intAry object whose value will be subtracted
//
//	from this current intAry value.
func (ia *IntAry) SubtractFromThis(ia2 *IntAry) error {

  IntAryMathSubtract{}.SubtractTotal(ia, ia2)

  return nil

}

// SubtractMultipleFromThis - This method will subtract multiple intAry values from the
// current intAry value. There are two input parameters:
//
// convertToNumStr bool - If true the result will be converted to a number string after
//
//	the final subtraction operation.
//
// iaMany ...*intAry - An unlimited series of pointers to intAry objects which will be
//
//	subtracted from the current intAry Value.
func (ia *IntAry) SubtractMultipleFromThis(iaMany ...*IntAry) error {

  for _, iAry := range iaMany {

    IntAryMathSubtract{}.SubtractTotal(ia, iAry)

  }

  return nil
}
