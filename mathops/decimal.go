package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
)

/*
	Decimal
	=======

	The source code repository for decimal.go is located at:
			https://github.com/MikeAustin71/mathopsgo.git

	The source file decimal.go is located in directory:
		MikeAustin71/mathopsgo/mathops/decimal.go

*/

// Decimal - This type is used to perform math operations
// which achieve a high degree of accuracy when dealing
// with fractional numbers containing digits to the right
// of the decimal place.
//
// While storage operations are provided by a type NumStrDto,
// math operations are performed using types *big.Int and
// *big.Rat.
//
// The Decimal Type implements the INumMgr interface.
type Decimal struct {
  bigINum BigIntNum
}

var _ INumMgr = (*Decimal)(nil)

// Add - Adds the value of the current Decimal to that of
// the incoming Decimal and returns in the result in a
// Decimal Type.
//
// Note that Numeric separators remain unchanged and are
// set to the values of the current Decimal instance.
func (dec *Decimal) Add(d2 Decimal) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.Add()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
          "Validation Error on Decimal instance 'dec'.",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.IsValid(ePrefix.XCpy("Validating Input Parameter 'd2'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Input parameter d2 (type Decimal) is INVALID!\n" +
          "Validation Error on Decimal instance 'd2'.",
        ErrMessage: err.Error(),
      }
  }

  bINumResult, err := new(BigIntMathAdd).AddBigIntNums(dec.bigINum, d2.bigINum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINumResult, err := new(BigIntMathAdd).AddBigIntNums(dec.bigINum, d2.bigINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d3, err := new(Decimal).NewBigIntNum(bINumResult)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d3, err := new(Decimal).NewBigIntNum(bINumResult)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d3.IsValid(ePrefix.XCpy("Validating final result 'd3'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Final result variable 'd3' (type Decimal) is INVALID!\n" +
          "Validation Error on Decimal instance 'd3'.",
        ErrMessage: err.Error(),
      }
  }

  err = d3.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d3.SetNumericSeparatorsDto(numSeps)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d3, nil
}

// AddToThis - adds the value of the incoming Decimal to that
// of the current Decimal object. The new total is retained
// in the current Decimal object.
func (dec *Decimal) AddToThis(d2 Decimal) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.AddToThis()",
    "")

  if err != nil {
    return err
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  dec.bigINum, err = new(BigIntMathAdd).AddBigIntNums(dec.bigINum, d2.bigINum)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "dec.bigINum, err = new(BigIntMathAdd).\n" +
        "   AddBigIntNums(dec.bigINum, d2.bigINum)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = dec.bigINum.SetNumericSeparatorsDto(numSeps)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddToThisArray - Receives an array of Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisArray(decs []Decimal) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.AddToThisArray()",
    "")

  if err != nil {
    return err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
      ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
        "Validation Error on Decimal instance 'dec'.",
      ErrMessage: err.Error(),
    }
  }

  if len(decs) == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "len(decs) == 0",
      ErrMessage: "Error: Input parameter array 'decs' is EMPTY!",
    }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  bINumResult, err := dec.bigINum.CopyOut()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "bINumResult, err := dec.bigINum.CopyOut()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  for i, dx := range decs {

    err = dx.IsValid(ePrefix.XCpy("Validating 'dx'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bINumResult, err := dec.bigINum.CopyOut()",
        ErrContext: fmt.Sprintf("Error: Array element decs[%v] is INVALID!\n"+
          "decs[%v] Failed Decimal Validation Test.", i, i),
        ErrMessage: err.Error(),
      }
    }

    bINumResult, err = new(BigIntMathAdd).AddBigIntNums(bINumResult, dx.bigINum)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bINumResult, err = new(BigIntMathAdd).AddBigIntNums(bINumResult, dx.bigINum)",
        ErrContext: fmt.Sprintf("Error occurred at array element decs[%v].", i),
        ErrMessage: err.Error(),
      }
    }
  } // End of for loop

  err = bINumResult.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = bINumResult.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.SetBigIntNum(bINumResult)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.SetBigIntNum(bINumResult)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
      ErrContext: "Error: Decimal type resulting from Array Addition is INVALID!\n" +
        "Decimal instance 'dec' FAILED Validation Test.",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddToThisSeries - Receives multiple Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisSeries(decs ...Decimal) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.AddToThisSeries()",
    "")

  if err != nil {
    return err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
      ErrContext: "Error: The current Decimal instance (dec) is INVALID! ",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  bINumResult, err := dec.bigINum.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " bINumResult, err := dec.bigINum.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  for i, dx := range decs {

    err = dx.IsValid(ePrefix.XCpy("Validating 'dx'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dx.IsValid(ePrefix.XCpy(\"Validating 'dx'\").String())",
        ErrContext: fmt.Sprintf("Error: Series element decs[%v] is INVALID!\n", i),
        ErrMessage: err.Error(),
      }
    }

    bINumResult, err = new(BigIntMathAdd).AddBigIntNums(bINumResult, dx.bigINum)

    if err != nil {

      return fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINumResult, err = new(BigIntMathAdd).AddBigIntNums(bINumResult, dx.bigINum)\n"+
        "Error occurred at series element decs[%v]\n"+
        "Error= %v\n",
        ePrefix,
        i,
        err.Error())
    }

  }

  err = bINumResult.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = bINumResult.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.SetBigIntNum(bINumResult)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.SetBigIntNum(bINumResult)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
      ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
        "Validation Error on Decimal instance 'dec'.",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AllDigitsNumStr - parses the incoming string and returns
// a pure number string consisting of all numeric digits. No
// sign characters, decimals or thousands separators are returned.
// The returned value is the absolute numeric value extracted from
// the 'numStr' input parameter.
func (dec *Decimal) AllDigitsNumStr(numStr string) (string, error) {

  ePrefix := "Decimal.AllDigitsNumStr()"

  bigIntNum, err := new(BigIntNum).NewNumStr(numStr)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bigIntNum, err := new(BigIntNum).NewNumStr(numStr)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  var xStr string

  xStr, err = bigIntNum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " numStr, err = bigIntNum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return xStr, nil
}

// NumStrToDecimal - Creates a Decimal type from a number
// string.
//
// The returned Decimal contains teh same numeric separators (decimal separator,
// thousands separator and currency symbol) as the current Decimal instance.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) NumStrToDecimal(numStr string) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NumStrToDecimal()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := Decimal{}

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrMessage: err.Error(),
      }
  }

  d2.bigINum, err = new(BigIntNum).NewNumStrWithNumSeps(numStr, &numSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "d2.bigINum, err = new(BigIntNum).NewNumStrWithNumSeps(numStr, &numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// CopyIn - Receives an incoming Decimal object
// as an input parameter and sets the Current Decimal
// equal to that of the incoming Decimal object.
func (dec *Decimal) CopyIn(d2 Decimal) error {

  ePrefix := "Decimal.CopyIn()"

  dec.Empty()

  var err error

  dec.bigINum, err = d2.bigINum.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = d2.bigINum.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// CopyOut - Returns a deep copy of the current Decimal
// instance.
func (dec *Decimal) CopyOut() (Decimal, error) {

  ePrefix := "Decimal.CopyOut()"

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d2 := new(Decimal).New()

  bINum2, err := dec.bigINum.CopyOut()

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINum2, err := dec.bigINum.CopyOut()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2.SetBigIntNum(bINum2)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(bINum2)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// Cmp - Performs a comparison of two Decimal numeric values
// and returns an integer value indicating the relationship
// between the two numeric values (i.e. greater than, less than,
// or equal).
//
// Return Values:
// dec == dec2 							Return  0
// dec > 	dec2							Return +1
// dec < 	dec2							Return -1
func (dec *Decimal) Cmp(dec2 Decimal) (int, error) {

  ePrefix := "Decimal.Cmp()"

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return -99,
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  cmpResult, err := dec.bigINum.Cmp(dec2.bigINum)

  if err != nil {

    return -99,
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " cmpResult, err := dec.bigINum.Cmp(dec2.bigINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return cmpResult, nil
}

// CubeRoot - Returns a Decimal instance with a numeric value equal to the
// cube root of the current Decimal numeric value. The current Decimal instance
// is the radicand.
//
// Returns:
// ========
// The calculation result is returned as a Decimal instance. The returned Decimal instance
// will contain	numeric separators (decimal separator, thousands separator and currency symbol)
// copied from the current Decimal instance (dec).
func (dec *Decimal) CubeRoot(maxPrecision uint) (Decimal, error) {

  ePrefix := "Decimal.CubeRoot() "

  err := dec.bigINum.IsValid(ePrefix + "Current Decimal instance is INVALID! ")

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  bINumThree, err := new(BigIntNum).NewThree(0)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINumThree, err := new(BigIntNum).NewThree(0)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  decCubeRoot := new(Decimal).New()

  decCubeRoot.bigINum, err =
    new(BigIntMathNthRoot).GetNthRoot(dec.bigINum, bINumThree, maxPrecision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " decCubeRoot.bigINum, err = BigIntMathNthRoot{}.GetNthRoot(dec.bigINum, bINumThree, maxPrecision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = decCubeRoot.IsValid(ePrefix + "decCubeRoot INVALID! ")

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Result value 'decCubeRoot' (type Decimal) is INVALID!\n"+
        "Validation Error on Decimal instance 'decCubeRoot'.\n"+
        "Error returned by: \n"+
        " err = decCubeRoot.IsValid(ePrefix + \"decCubeRoot INVALID! \")\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return decCubeRoot, nil
}

// Divide - Divides the current decimal value by the input parameter 'divisor'
// and returns the quotient as a new Decimal instance.
//
// The calculation result is returned as a Decimal instance. The returned Decimal
// instance will contain	numeric separators (decimal separator, thousands separator
// and currency symbol) copied from the current Decimal instance (dec).
func (dec *Decimal) Divide(divisor Decimal, maxPrecision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.Divide()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
          "Validation Error on Decimal instance 'dec'.",
        ErrMessage: err.Error(),
      }
  }

  err = divisor.IsValid(ePrefix.XCpy("Validating input param 'divisor'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = divisor.IsValid(ePrefix.XCpy(\"Validating input param 'divisor'\").String())",
        ErrContext: "Input parameter 'divisor' (type Decimal) is INVALID!\n" +
          "Validation Error on Decimal instance 'divisor'.",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  d2Quotient := new(Decimal).New()

  d2Quotient.bigINum, err =
    new(BigIntMathDivide).BigIntNumFracQuotient(dec.bigINum, divisor.bigINum, numSeps, maxPrecision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d2Quotient.bigINum, err = new(BigIntMathDivide).BigIntNumFracQuotient(dec.bigINum, divisor.bigINum, maxPrecision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2Quotient.IsValid(ePrefix.XCpy("Validating d2Quotient").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2Quotient.IsValid(ePrefix.XCpy(\"Validating d2Quotient\").String())",
        ErrContext: "Intermediate variable 'd2Quotient' (type Decimal) is INVALID!\n" +
          "Validation Error on Decimal instance 'd2Quotient'.",
        ErrMessage: err.Error(),
      }
  }

  return d2Quotient, nil
}

// Equal - Returns true if the input Decimal instance is equal
// in all respects to the current Decimal instance.
//
// Note: this method will return false if two decimal instances
// being compared have equal numeric values but different precisions
// or different numeric separators.
//
// To perform an equivalency test of numeric values, see method
// Decimal{}.EqualValue() below.
func (dec *Decimal) Equal(dec2 Decimal) (bool, error) {

  return dec.bigINum.Equal(dec2.bigINum)
}

// EqualValue - Returns 'true' if the input Decimal instance holds
// a numeric value equal to that of the current Decimal
// instance.
func (dec *Decimal) EqualValue(dec2 Decimal) (bool, error) {

  ePrefix := "Decimal.EqualValue()"

  areEqual, err := dec.bigINum.EqualValue(dec2.bigINum)

  if err != nil {
    return false,
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " areEqual, err := dec.bigINum.EqualValue(dec2.bigINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return areEqual, nil
}

// Empty - Sets all values of the current Decimal's
// fields to their 'zero' values.
func (dec *Decimal) Empty() {

  new(bigIntNumElectron).empty(&dec.bigINum)

  return
}

// GetAbsoluteValue - returns the absolute value of the
// decimal expressed as a string. If the decimal value is
// '-123.456', this method will return '123.456'.
func (dec *Decimal) GetAbsoluteValue() (Decimal, error) {

  ePrefix := "Decimal.GetAbsoluteValue()"

  bi2, err := dec.bigINum.GetAbsoluteBigIntNumValue()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix,
        ReturnFunc: "bi2, err := dec.bigINum.GetAbsoluteBigIntNumValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bi3, err := new(Decimal).NewBigIntNum(bi2)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix,
        ReturnFunc: "bi3, err := new(Decimal).NewBigIntNum(bi2)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return bi3, nil
}

// GetAbsoluteAllDigitsStr - Returns the absolute value of the Decimal integer.
// Fractions are not returned, only the string of signed numeric digits which
// constitutes the entire number. In other words, if the value of the decimal
// is '-123.456', this method will return '123456'.
func (dec *Decimal) GetAbsoluteAllDigitsStr() (string, error) {

  ePrefix := "Decimal.GetAbsoluteAllDigitsStr()"

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  numStr, err := dec.bigINum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " numStr, err := dec.bigINum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return numStr, nil
}

// GetBigFloat - returns big Float representation of the Decimal Value.
func (dec *Decimal) GetBigFloat() (*big.Float, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetBigFloat()",
    "")

  if err != nil {
    return big.NewFloat(0.0), err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return big.NewFloat(0.0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
          "Validation Error on Decimal instance 'dec'.",
        ErrMessage: err.Error(),
      }
  }

  bigFloat, err := dec.bigINum.GetBigFloat()

  if err != nil {

    return big.NewFloat(0.0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bigFloat, err := dec.bigINum.GetBigFloat()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return bigFloat, nil
}

// GetBigFloatString - returns a signed number string which is accurate out
// to a large number of decimal places.
func (dec *Decimal) GetBigFloatString(precision uint) (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetBigFloatString()",
    "")

  if err != nil {
    return "", err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
          "Validation Error on Decimal instance 'dec'.",
        ErrMessage: err.Error(),
      }
  }

  biNum2, err := dec.bigINum.CopyOut()

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "biNum2, err := dec.bigINum.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = biNum2.SetPrecision(precision)

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = biNum2.SetPrecision(precision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  biNum2NumStr, err := biNum2.GetNumStr()

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "biNum2NumStr, err := biNum2.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return biNum2NumStr, nil
}

// GetBigInt - returns the Decimal value expressed as an
// integer value using type *big.Int. No factional values are included.
// For example, the value '-123.456' would be returned as the integer
// value '-123456'.  To compute the precise value of the Decimal, this
// integer value would need to be divided by the 'precision Value'. See
// GetScaleVal() below.
func (dec *Decimal) GetBigInt() (*big.Int, error) {

  ePrefix := "Decimal.GetBigInt() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {
    return big.NewInt(0),
      fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
        "Error='%v' ", err.Error())
  }

  bInt, err := dec.bigINum.GetBigInt()

  if err != nil {
    return big.NewInt(0),
      fmt.Errorf(ePrefix+"Error returned by dec.bigINum.GetBigInt() "+
        "Error='%v' ", err.Error())
  }

  return bInt, nil
}

// GetBigIntNum - Converts the current Decimal numeric value to
// an instance of type 'BigIntNum' and returns it to the calling
// function.
//
// The returned BigIntNum type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal
// instance.
//
// This method performs a validity test on the current Decimal instance.
func (dec *Decimal) GetBigIntNum() (BigIntNum, error) {

  ePrefix := "Decimal.GetBigIntNum()"

  err := dec.bigINum.IsValid(ePrefix + "Decimal INVALID! ")

  if err != nil {

    return BigIntNum{},
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix + \"Decimal INVALID! \")\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())

  }

  bigIntNum2, err := dec.bigINum.CopyOut()

  if err != nil {

    return BigIntNum{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bigIntNum2, err := dec.bigINum.CopyOut()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return bigIntNum2, nil
}

// GetBigRat - Returns the current Decimal's numeric value expressed
// as a rational number of type *big.Rat.
func (dec *Decimal) GetBigRat() (*big.Rat, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntNum.NewNumStrWithNumSeps()",
    "")

  if err != nil {
    return big.NewRat(1, 1), err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return big.NewRat(1, 1),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Error: Current instance of Decimal (dec) is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  decBigRat, err := dec.bigINum.GetBigRat()

  if err != nil {

    return big.NewRat(1, 1),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigRat, err := dec.bigINum.GetBigRat()",
        ErrContext: "Error: Current instance of Decimal (dec) is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  return decBigRat, nil
}

// GetCurrencySymbol - Returns the Decimal's current
// value for Currency Symbol.
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// In the USA, the currency symbol is the dollar sign
// ('$').
func (dec *Decimal) GetCurrencySymbol() (rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetCurrencySymbol()",
    "")

  if err != nil {
    return 0, err
  }

  currencySymbol, err := dec.bigINum.GetCurrencySymbol()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "currencySymbol, err := dec.bigINum.GetCurrencySymbol()",
        ErrContext: "Error: Current instance of Decimal (dec) is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  return currencySymbol, nil
}

// GetCurrencyStr - Returns the Decimal's numeric value expressed
// as number string delimited with the Decimal's Thousands Separator
// and prefixed with the designated Currency Symbol characters.
//
// Note: The file mathopsconstants.go file contains Unicode characters
// for most of the world's major currencies. This file is located at:
//
//	MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. These Separators and Currency
// Symbol are variable and may be controlled by the user.
//
// If the numeric value is negative, a leading minus sign will be prefixed
// to the currency display.
//
// Example:
// numstr = 1000000.23
// GetCurrencyStr() = $1,000,000.23
//
// numstr = -1000000.23
// GetCurrencyStr() = -$1,000,000.23
//
// Note: If the current Decimal is invalid, this method
// returns an empty string.
func (dec *Decimal) GetCurrencyStr() (string, error) {

  ePrefix := "Decimal.GetCurrencyStr()"

  currencyStr, err := dec.bigINum.FormatCurrencyStr(LEADMINUSNEGVALFMTMODE)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " currencyStr, err := dec.bigINum.FormatCurrencyStr(LEADMINUSNEGVALFMTMODE)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return currencyStr, nil
}

// GetCurrencyParen - Returns the Decimal's numeric value expressed
// as number string delimited with the Decimal's Thousands Separator
// and prefixed with the designated Currency Symbol characters.
//
// Note: The file mathopsconstants.go file contains Unicode characters
// for most of the world's major currencies. This file is located at:
//
//	MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. These Separators and Currency
// Symbol are variable and may be controlled by the user.
//
// If the numeric value is negative, the resulting number string is
// surrounded by parentheses.
//
// Example:
// numstr = 1000000.23
// GetCurrencyParen() = $1,000,000.23
//
// numstr = -1000000.23
// GetCurrencyParen() = ($1,000,000.23)
//
// Note: If the current Decimal is invalid, this method
// returns an empty string.
func (dec *Decimal) GetCurrencyParen() (string, error) {
  return dec.bigINum.FormatCurrencyStr(PARENTHESESNEGVALFMTMODE)
}

// GetDecimal - Returns a deep copy of the current Decimal instance.
//
// The returned BigIntNum type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal
// instance.
//
// This method performs a validity test on the current Decimal instance.
func (dec *Decimal) GetDecimal() (Decimal, error) {

  ePrefix := "Decimal.GetDecimal()"

  err := dec.bigINum.IsValid(ePrefix + "Decimal INVALID! ")

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix + \"Decimal INVALID! \")\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())

  }

  dec2, err := dec.CopyOut()

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " dec2, err := dec.CopyOut()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return dec2, nil
}

// GetDecimalSeparator - returns the Decimal's current
// value for Decimal Separator (i.e. '.')
func (dec *Decimal) GetDecimalSeparator() (rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetDecimalSeparator()",
    "")

  if err != nil {
    return 0, err
  }

  decDecimalSep, err := dec.bigINum.GetDecimalSeparator()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decDecimalSep, err := dec.bigINum.GetDecimalSeparator()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return decDecimalSep, nil
}

// GetFloat32 - Returns the current value of the Decimal as a
// float32. There may be a loss of accuracy during this conversion.
// The return parameter big.Accuracy will describe the level of
// accuracy provided.
//
// See big.Accuracy:
//
//	Below Accuracy = -1
//	Exact Accuracy = 0
//	Above Accuracy = +1
func (dec *Decimal) GetFloat32() (float32, big.Accuracy, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetFloat32()",
    "")

  if err != nil {
    return float32(0.0), big.Accuracy(0), err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return float32(0.0), big.Accuracy(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "The current instane of 'Decimal' is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  decBINumStr, err := dec.bigINum.GetNumStr()

  if err != nil {

    return float32(0.0), big.Accuracy(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bf, status := big.NewFloat(0.0).SetString(decBINumStr)

  if !status {

    return float32(0.0), big.Accuracy(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bf, status := big.NewFloat(0.0).SetString(decBINumStr)",
        ErrContext: "'status' == false",
        ErrMessage: fmt.Sprintf("SetString() Failed. NumStr= %v", decBINumStr),
      }
  }

  f32, accuracy := bf.Float32()

  return f32, accuracy, nil

}

// GetFloat64 - Returns the current value of the Decimal as a
// float64. There may be a loss of accuracy during this conversion.
// The return parameter big.Accuracy will describe the level of
// accuracy provided.
//
// See big.Accuracy:
//
//	Below Accuracy = -1
//	Exact Accuracy = 0
//	Above Accuracy = +1
func (dec *Decimal) GetFloat64() (float64, big.Accuracy, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetFloat64()",
    "")

  if err != nil {
    return 0.0, big.Accuracy(0), err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return 0.0, big.Accuracy(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
          "Validation Error on Decimal instance 'dec'.",
        ErrMessage: err.Error(),
      }
  }

  decBigINumStr, err := dec.bigINum.GetNumStr()

  if err != nil {

    return 0.0, big.Accuracy(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumStr, err := dec.bigINum.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bf, status := big.NewFloat(0.0).SetString(decBigINumStr)

  if !status {

    return 0.0, big.Accuracy(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bf, status := big.NewFloat(0.0).SetString(decBigINumStr)",
        ErrContext: "'status'= false",
        ErrMessage: fmt.Sprintf("SetString() Failed. NumStr= %v", decBigINumStr),
      }
  }

  f64, accuracy := bf.Float64()

  return f64, accuracy, nil
}

// GetIntAry - Returns an IntAry structure initialized
// to the value of the current 'Decimal' object.
//
// The returned IntAry contains numeric separators (decimal
// separator, thousands separator and currency symbol) copied
// from the current Decimal instance.
func (dec *Decimal) GetIntAry() (IntAry, error) {

  ePrefix := "Decimal.GetIntAryElements()"

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return IntAry{},
      fmt.Errorf("%v\n"+
        "Current instance of Decimal (dec) is INVALID!\n"+
        "Validation Error on Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  ia, err := dec.bigINum.GetIntAry()

  if err != nil {

    return IntAry{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " ia, err := dec.bigINum.GetIntAry()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return ia, nil
}

// GetNumericSeparatorsDto - returns a NumericSeparatorDto structure
// containing the current characters (runes) used to specify
// decimal point separator, thousands separator and currency symbol.
func (dec *Decimal) GetNumericSeparatorsDto() (NumericSeparatorDto, error) {

  return dec.bigINum.GetNumericSeparatorsDto()
}

// GetNumStr - Returns the internal value of the Decimal
// expressed as a signed numeric string. precision, or
// placement of the decimal point, is controlled by
// the Decimal's precision setting.
//
// Example Output:
// ===============
//
//	123
//	123.4
//	123456789
//	123456789.44
//
// -123
// -123.4
// -123456789
// -123456789.44
func (dec *Decimal) GetNumStr() (string, error) {

  ePrefix := "Decimal.GetNumStr()"

  numStr, err := dec.bigINum.FormatNumStr(LEADMINUSNEGVALFMTMODE)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " numStr, err := dec.bigINum.FormatNumStr(LEADMINUSNEGVALFMTMODE)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return numStr, nil
}

// GetNumParen - Returns the internal value of the
// Decimal expressed as number string. precision
// or placement of the decimal point is controlled
// by the Decimal's 'precision' setting.
//
// If the numeric value is less than zero, a negative
// number, the number string is surrounded in parentheses.
//
// Example Output:
// ===============
//
//	123
//	123.4
//
// (123)
// (123.4)
func (dec *Decimal) GetNumParen() (string, error) {

  ePrefix := "Decimal.GetNumParen()"

  numStr, err := dec.bigINum.FormatNumStr(PARENTHESESNEGVALFMTMODE)

  if err != nil {

    return "",
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " numStr, err := dec.bigINum.FormatNumStr(PARENTHESESNEGVALFMTMODE)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return numStr, nil
}

// GetNumStrDto - returns a NumStrDto structure initialized
// to the value of the current Decimal object.
//
// The returned NumStrDto type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal
// instance.
func (dec *Decimal) GetNumStrDto() (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetNumStrDto()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  decBigINum, err := dec.bigINum.GetNumStr()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINum, err := dec.bigINum.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nDto, err := dec.bigINum.GetNumStrDto()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nDto, err := dec.bigINum.GetNumStrDto()",
        ErrContext: fmt.Sprintf("dec.bigINum= '%v'", decBigINum),
        ErrMessage: err.Error(),
      }
  }

  return nDto, nil
}

// GetPrecision - returns the Decimal's current precision
// value. The Decimal structure maintains precision as an
// unsigned integer.
//
// The value of 'precision' returned by this method will
// always be >= zero (greater than or equal to zero '0').
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//					1.234    	GetPrecision() = 3
//							5			GetPrecision() = 0
//				0.12345  		GetPrecision() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (dec *Decimal) GetPrecision() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetPrecision()",
    "")

  if err != nil {
    return 0, err
  }

  decBigINumIntPrecision, err := dec.bigINum.GetPrecisionInt()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumIntPrecision, err :=  dec.bigINum.GetPrecisionInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return decBigINumIntPrecision, nil
}

// GetPrecisionUint - Returns precision as an
// unsigned integer.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//					1.234    	GetPrecision() = 3
//							5			GetPrecision() = 0
//				0.12345  		GetPrecision() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (dec *Decimal) GetPrecisionUint() (uint, error) {

  ePrefix := "Decimal.GetPrecisionUint()"

  precisionUint, err := dec.bigINum.GetPrecisionUint()

  if err != nil {
    return 0, &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "precisionUint, err := dec.bigINum.GetPrecisionUint()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return precisionUint, nil
}

// GetRational - returns a big Rational number type which
// is capable of very high accuracy.
//
// The returned *big.Rat number is initialized to the
// current value of the Decimal object.
func (dec *Decimal) GetRational() (*big.Rat, error) {

  ePrefix := "Decimal.GetRational()"

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {
    return big.NewRat(1, 1),
      fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
        "Error='%v' ", err.Error())
  }

  x, err := dec.bigINum.GetBigInt()

  if err != nil {
    return big.NewRat(1, 1),
      fmt.Errorf(ePrefix+"Error returned by dec.bigINum.GetBigInt() "+
        "Error='%v' ", err.Error())
  }

  decSignedAllDigitsBigInt := big.NewInt(0).Set(x)

  rDividend := big.NewRat(1, 1).SetInt(decSignedAllDigitsBigInt)

  decScaleFactor, err := dec.bigINum.GetScaleFactor()

  if err != nil {
    return big.NewRat(1, 1), &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "decScaleFactor, err := dec.bigINum.GetScaleFactor()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  rDivisor := big.NewRat(1, 1).SetInt(decScaleFactor)

  rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

  return rQuotient, nil
}

// GetSign - Returns the sign of the current
// Decimal Value. Return values are one of two
// integers: +1 or -1.
func (dec *Decimal) GetSign() (int, error) {

  ePrefix := "Decimal.GetSign()"

  sign, err := dec.bigINum.GetSign()

  if err != nil {
    return 0, &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "sign, err := dec.bigINum.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return sign, nil
}

// GetRelevantPrecision - Returns an unsigned integer representing
// the number of decimal places to the right of the decimal which
// are occupied by non-zero values.
//
// Example: Value = 1.640700000.  The number of relevant decimal
// places to the right of the decimal is '4'. In the case of an
// integer number the relevant precision is zero ('0') because
// there are no digits to the right of the decimal.
func (dec *Decimal) GetRelevantPrecision() (uint, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntNum.NewNumStrWithNumSeps()",
    "")

  if err != nil {
    return 0, err
  }

  bI2, err := dec.bigINum.CopyOut()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bI2, err := dec.bigINum.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = bI2.TrimTrailingFracZeros()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = bI2.TrimTrailingFracZeros()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  bI2PrecisionUint, err := bI2.GetPrecisionUint()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bI2PrecisionUint, err := bI2.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return bI2PrecisionUint, nil
}

// GetScaleVal - Returns the scale value associated with this decimal value. The
// scale value is expressed as 10 to an exponent. Example 10^2 == 100.
//
// Scale Value, or Scale Factor, is defined by 10 raised to the power
// of Decimal precision.
//
// The return scale value is of type big integer (*big.Int)
func (dec *Decimal) GetScaleVal() (*big.Int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetScaleVal()",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Error: The current instance of Decimal is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  decBigINumScaleFactor, err := dec.bigINum.GetScaleFactor()

  if err != nil {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumScaleFactor, err := dec.bigINum.GetScaleFactor()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return decBigINumScaleFactor, nil
}

// GetSciNotationNumber - Converts the numeric value of the current
// Decimal instance into scientific notation and returns this value
// as an instance of type SciNotationNum.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//
//											scientific notation string. If the value of 'mantissaLen'
//											is less than two ('2'), this method will automatically set
//											the 'mantissaLen' to a default value of two ('2').
//
//											Example Scientific Notation:
//											----------------------------
//
//	 										scientific notation string: '2.652e+8'
//
//	 										significand = '2.652'
//	 										significand integer digit = '2'
//												mantissa		= significand factional digits = '.652'
//	 										exponent    = '8'  (10^8)
func (dec *Decimal) GetSciNotationNumber(mantissaLen uint) (SciNotationNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetSciNotationNumber()",
    "")

  if err != nil {
    return SciNotationNum{}, err
  }

  sciNotationNum, err := dec.bigINum.GetSciNotationNumber(mantissaLen)

  if err != nil {

    return SciNotationNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "sciNotationNum, err := dec.bigINum.GetSciNotationNumber(mantissaLen)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return sciNotationNum, nil
}

// GetSciNotationStr - Returns a string expressing the current Decimal
// numerical value as scientific notation.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//
//											scientific notation string. If the value of 'mantissaLen'
//											is less than two ('2'), this method will automatically set
//											the 'mantissaLen' to a default value of two ('2').
//
//											Example Scientific Notation:
//											----------------------------
//
//	 										scientific notation string: '2.652e+8'
//
//	 										significand = '2.652'
//	 										significand integer digit = '2'
//												mantissa		= significand factional digits = '.652'
//												exponent    = '8'  (10^8)
func (dec *Decimal) GetSciNotationStr(mantissaLen uint) (string, error) {

  ePrefix := "BigIntNum.GetSciNotationStr() "

  sciNotation, err := dec.bigINum.GetSciNotationNumber(mantissaLen)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by dec.bigINum.GetSciNotationNumber(mantissaLen). "+
        "Error='%v'", err.Error())
  }

  result, err := sciNotation.GetSciNotationStr(mantissaLen)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by sciNotation.GetSciNotationStr(mantissaLen). "+
        "Error='%v'", err.Error())
  }

  return result, nil
}

// GetSignedAllDigitsStr - Returns the Decimal's internal
// Signed All Digits Integer Value expressed as a string.
// No Fractional Digits are included, this is a signed
// integer number string. Example: The value '-123.456'
// would be returned as '-123456'
func (dec *Decimal) GetSignedAllDigitsStr() (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntNum.NewNumStrWithNumSeps()",
    "")

  if err != nil {
    return "", err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "The current instance of Decimal is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  nDto, err := dec.bigINum.GetNumStrDto()

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nDto, err := dec.bigINum.GetNumStrDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  result := ""

  decBigINumSignValue, err := dec.bigINum.GetSign()

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumSignValue, err := dec.bigINum.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if decBigINumSignValue < 0 {
    result += "-"
  }

  result += string(nDto.absAllNumRunes)

  return result, nil
}

// GetSignedBigInt - Returns the numeric value of the current Decimal
// instance as a signed *big.Int.
func (dec *Decimal) GetSignedBigInt() (*big.Int, error) {

  ePrefix := "Decimal.GetBigInt() "

  bInt, err := dec.bigINum.GetBigInt()

  if err != nil {

    return big.NewInt(0),
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bInt, err := dec.bigINum.GetBigInt()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return big.NewInt(0).Set(bInt), nil
}

// GetThisPointer - Returns a pointer to the current Decimal instance
func (dec *Decimal) GetThisPointer() *Decimal {

  return dec
}

// GetThousandsSeparator - Gets the current value of
// the Thousands Separator for the current Decimal
// object.
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// For U.S.A. - The thousands separator is a comma (',')
func (dec *Decimal) GetThousandsSeparator() (rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetThousandsSeparator()",
    "")

  if err != nil {
    return 0, err
  }

  thousandsSeparator, err := dec.bigINum.GetThousandsSeparator()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "thousandsSeparator, err := dec.bigINum.GetThousandsSeparator()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return thousandsSeparator, nil
}

// GetThouStr - Returns a number string which represents the Decimal's
// numeric value. Thousands are separated by the Decimal's Thousands
// Separator. In the USA, the Thousands Separator is a comma character
// (',').
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// Negative numbers are preceded by a minus sign.
//
// Example Output
// ==============
//
//	123,456,789
//	123,456,789.12
//
// -123,456,789
// -123,456,789.12
func (dec *Decimal) GetThouStr() (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetThouStr()",
    "")

  if err != nil {
    return "", err
  }

  decBigINumThouStr, err := dec.bigINum.FormatThousandsStr(LEADMINUSNEGVALFMTMODE)

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumThouStr, err := dec.bigINum.FormatThousandsStr(LEADMINUSNEGVALFMTMODE)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return decBigINumThouStr, nil
}

// GetThouParen - Returns the Decimal's numeric value formatted
// as a number string with thousands separated by the Decimal's
// Thousands Separator.  In the USA, the Thousands Separator is
// a comma character (',').
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// Negative numbers are surrounded with parentheses.
//
// Example Output
// ==============
//
//	123,456,789
//	123,456,789.12
//
// (123,456,789)
// (123,456,789.12)
func (dec *Decimal) GetThouParen() (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.GetThouParen()",
    "")

  if err != nil {
    return "", err
  }

  fmtThouStrParen, err := dec.bigINum.FormatThousandsStr(PARENTHESESNEGVALFMTMODE)

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "fmtThouStrParen, err := dec.bigINum.FormatThousandsStr(PARENTHESESNEGVALFMTMODE)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return fmtThouStrParen, nil
}

// GetIsValid - returns a boolean indicating
// the current state of the Decimal information.
// If the current Decimal object is VALID, the
// method returns 'true'.
//
// If the current Decimal object is INVALID, the
// method returns 'false'.
//
// Notice that this method relies on
// Decimal.IsValid(ePrefix) which returns an
// 'error' type.
func (dec *Decimal) GetIsValid() bool {

  ePrefix := "Decimal.GetIsValid()"
  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {
    return false
  }

  return true
}

// Inverse - Returns the inverseBigIntNum of the current
// Decimal.  Inverse = 1/Current Decimal Value.
//
// Input Parameters:
// maxPrecision uint - determines the number of digits
//
//	to the right of the decimal point in the result.
//	if maxPrecision is less than zero, an error will
//	be triggered
func (dec *Decimal) Inverse(maxPrecision uint) (Decimal, error) {

  ePrefix := "Decimal.Inverse()"

  bin2, err := dec.bigINum.GetInverse(maxPrecision)

  if err != nil {
    return Decimal{},
      fmt.Errorf(ePrefix+
        "Error returned by dec.bigINum.GetInverse(uint(maxPrecision)) "+
        "Error='%v' \n", err.Error())
  }

  d2, err := bin2.GetDecimal()

  if err != nil {
    return Decimal{},
      fmt.Errorf(ePrefix+
        "Error returned by bin2.GetDecimal() "+
        "Error='%v' \n", err.Error())
  }

  return d2, nil
}

// IsEvenNumber - If the numeric value of the current Decimal instance
// is evenly divisible by two, with no remainder, it is classified as
// an even number and this method will return 'true'.
func (dec *Decimal) IsEvenNumber() (bool, error) {

  return dec.bigINum.IsEvenNumber()
}

// IsFraction - returns a boolean value. If 'true',
// it signals that the Decimal has digits to the
// right of the decimal place. If 'false', it
// signals that the decimal value is an integer with
// no digits to the right of the decimal place.
func (dec *Decimal) IsFraction() (bool, error) {

  ePrefix := "Decimal.IsFraction() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {
    return false,
      fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
        "Error='%v' ", err.Error())
  }

  if dec.bigINum.precision != 0 {
    return true, nil
  }

  return false, nil
}

// IsValid - Performs an internal diagnostic on the current
// Decimal instance and returns an 'error' if the instance is INVALID.
func (dec *Decimal) IsValid(callingMethodName string) error {

  ePrefix := "Decimal.IsValid()"

  if len(callingMethodName) > 0 {
    ePrefix = ePrefix + "\n" + callingMethodName
  }

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {
    return err
  }

  return nil
}

// IsZero - Returns true if the numeric value of the current
// 'Decimal' instance is zero.
func (dec *Decimal) IsZero() (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.IsZero()",
    "")

  if err != nil {
    return false, err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of 'dec.bigINum' is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  decBigINumIsZero, err := dec.bigINum.IsZero()

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumIsZero, err := dec.bigINum.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return decBigINumIsZero, nil
}

// MakeDecimalFromIntAry - generates a Decimal Type based on string information
// provided by the 'ia' *IntAry input parameter.
func (dec *Decimal) MakeDecimalFromIntAry(ia *IntAry) (Decimal, error) {

  ePrefix := "Decimal.MakeDecimalFromIntAry() "

  var err error

  if ia == nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "ERROR: Input parameter 'ia' is nil pointer!\n"+
        "Error= %v\n",
        ePrefix)
  }

  ia.SetInternalFlags()

  err = ia.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Input parameter 'ia' is INVALID!\n"+
        "Validation Error on IntAry instance 'ia'.\n"+
        "Error returned by: \n"+
        " err = ia.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d2 := Decimal{}

  d2.bigINum, err = ia.GetBigIntNum()

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d2.bigINum, err = ia.GetBigIntNum()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// Mod - performs a modulo operation where the current BigIntNum numeric value is the
// dividend and the divisor is the input parameter, 'divisor'.  The modulo operation finds
// the remainder after division of one number by another (sometimes called modulus).
// (Wikipedia: https://en.wikipedia.org/wiki/Modulo_operation)
//
//		 									dividend = bNum
//	  									dividend % divisor = modulo
//
// The result of this modulo operation is returned as a BigIntNum, 'modulo'. 'modulo' may
// consist of an integer or a floating point value consisting of integer and fractional
// digits.
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// floating point 'modulo'. Precision is defined as the number of fractional digits to
// the right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// The returned BigIntNum instance, 'modulo', will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from the current BigIntNum
// instance (bNum).
func (dec *Decimal) Mod(divisor Decimal,
  maxPrecision uint) (modulo Decimal, err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.Mod()",
    "")

  if err != nil {
    return modulo, err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return modulo,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Error: The current instance of Decimal ('dec') is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  bigIntNum2, err := dec.bigINum.CopyOut()

  if err != nil {

    return modulo,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bigIntNum2, err := dec.bigINum.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return modulo,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  modulo.bigINum, err = new(BigIntMathDivide).BigIntNumModulo(
    bigIntNum2, divisor.bigINum, numSeps, maxPrecision)

  if err != nil {

    return modulo,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "modulo.bigINum, err = new(BigIntMathDivide).\n" +
          "  BigIntNumModulo(bigIntNum2, divisor.bigINum, numSeps, maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision='%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  return modulo, nil
}

// Multiply - Multiplies the numeric value of the current Decimal instance
// (multiplier) by input parameter Decimal type 'multiplicand' and returns
// the product as a new Decimal instance.
//
// Before the multiplication operation is initiated, this method performs a
// validity check on both the current Decimal instance and the input
// parameter.
//
// The returned 'product' Decimal instance will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from
// the current Decimal instance.
func (dec *Decimal) Multiply(
  multiplicand Decimal) (product Decimal, err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.Multiply()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  product = new(Decimal).New()

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Current instance of Decimal (dec) is INVALID!\n" +
          "Validation Error on Decimal instance 'dec'.",
        ErrMessage: err.Error(),
      }
  }

  err = multiplicand.IsValid(ePrefix.XCpy("Validating Input Param 'multiplicand'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = multiplicand.IsValid(ePrefix.XCpy(\"Validating Input Param 'multiplicand'\").String())",
        ErrContext: "Input parameter multiplicand (type Deciimal) is INVALID!\n" +
          "Validation Error on Decimal instance 'multiplicand'.",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  product.bigINum, err =
    new(BigIntMathMultiply).MultiplyBigIntNums(
      dec.bigINum, multiplicand.bigINum)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "product.bigINum, err = new(BigIntMathMultiply).\n" +
          "   MultiplyBigIntNums(dec.bigINum, multiplicand.bigINum)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = product.bigINum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = product.bigINum.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return product, nil
}

// New - Creates and returns a Decimal type. The Decimal numeric value is
// initialized to zero.
//
// The returned new Decimal instance will contain USA default numeric separators
// (decimal separator, thousands separator and currency symbol).
//
// Example Usage:
//
//	d := Decimal{}.New()
//
// This is the recommended procedure for creating
// a Decimal type.
func (dec *Decimal) New() Decimal {

  d := Decimal{}

  d.Empty()

  return d
}

// NewWithNumSeps - Creates and returns a new Decimal instance.
// The returned new Decimal instance will contain numeric separators
// (decimal separator, thousands separator and currency symbol)
// copied from the input parameter, 'numSeps'.
//
// The Decimal value is initialized to zero.
//
// Example Usage:
//
//	d := Decimal{}.NewWithNumSeps(numSeps)
func (dec *Decimal) NewWithNumSeps(
  numSeps NumericSeparatorDto) (Decimal, error) {

  ePrefix := "Decimal.NewWithNumSeps()"

  d := Decimal{}

  d.Empty()

  numSeps.SetDefaultsIfEmpty()

  err := d.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err := d.SetNumericSeparatorsDto(numSeps)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d, nil
}

// NewPtr - Creates and returns a pointer to
// a new Decimal type. Can be used for initializing
// a Decimal Type and calling pointer methods in
// one statement.
//
// Example:
// d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision,true)
func (dec *Decimal) NewPtr() *Decimal {
  d := Decimal{}
  d.Empty()
  return &d
}

// NewBigInt - Returns a Decimal type based on Big Int and
// precision input parameters. If an error is encountered, it
// will trigger a panic condition.
//
// The 'NewBigInt' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewBigInt(bigI, precision)
// bigI := big.NewInt(123456)
// Decimal{}.NewBigInt(bigI, 3) = 123.456
func (dec *Decimal) NewBigInt(
  bigI *big.Int, precision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewBigInt()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  if bigI == nil {

    return Decimal{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'bigI'",
      }
  }

  err = d2.SetBigInt(bigI, precision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetBigInt(bigI, precision)",
        ErrContext: fmt.Sprintf("bigI= '%v' precision= '%v'",
          bigI.Text(10), precision),
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewBigIntNum - Returns a Decimal instance based on input
// parameter,
func (dec *Decimal) NewBigIntNum(
  bigINum BigIntNum) (Decimal, error) {

  ePrefix := "Decimal.NewBigIntNum()"

  var err error

  d2 := Decimal{}

  d2.bigINum, err = bigINum.CopyOut()

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d2.bigINum, err = bigINum.CopyOut()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewInt - Returns a Decimal type based on input parameters 'intNum'
// and 'precision'.
//
// Input parameter 'intNum' is of type int.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//					intNum := int(123456)
//					precision := uint(3)
//					dec := Decimal{}.NewInt(intNum, precision)
//	       dec is now equal to 123.456
//
// Examples:
// ---------
//
//	  intNum				precision			Decimal Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (dec *Decimal) NewInt(
  intNum int, precision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewInt()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  err = d2.SetInt(intNum, precision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetInt(intNum, precision)",
        ErrContext: fmt.Sprintf("intNum= '%v' precision= '%v'",
          intNum, precision),
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewIntExponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('intNum') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'intNum' is of type int.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//		decNum := Decimal{}.NewIntExponent(123456, -3)
//	 -- decNum is now equal to "123.456", precision = 3
//
//		decNum := Decimal{}.NewIntExponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  intNum			 exponent			  	Decimal Result
//		123456		 		  -3							123.456
//		123456		 		   3							123456.000
//	  123456					 0              123456
func (dec *Decimal) NewIntExponent(
  intNum int, exponent int) (Decimal, error) {

  ePrefix := "Decimal.NewIntExponent()"

  d2 := new(Decimal).New()

  err := d2.bigINum.SetBigIntExponent(big.NewInt(int64(intNum)), exponent)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err := d2.bigINum.SetBigIntExponent(big.NewInt(int64(intNum)), exponent)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewInt32 - Returns a new Decimal instance based on input parameters
// 'int32Num' and 'precision'.
//
// Input parameter 'int32Num' is of type int32.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//					int32Num := int32(123456)
//					precision := uint(3)
//					dec := Decimal{}.NewInt32(int32Num, precision)
//	       dec is now equal to 123.456
//
// Examples:
// ---------
//
//	 int32Num			precision			Decimal Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (dec *Decimal) NewInt32(
  int32Num int32, precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewInt32()"

  d2 := new(Decimal).New()

  err := d2.SetInt64(int64(int32Num), precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err := d2.SetInt64(int64(int32Num), precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, err
}

// NewInt32Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('int32Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int32Num' is of type int32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction
// with the Decimal{} syntax thereby allowing Decimal
// type creation and initialization in one step.
//
//		decNum := Decimal{}.NewInt32Exponent(123456, -3)
//	 -- decNum is now equal to "123.456", precision = 3
//
//		decNum := Decimal{}.NewInt32Exponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  int32Num		 exponent			  	Decimal Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (dec *Decimal) NewInt32Exponent(
  int32Num int32, exponent int) (Decimal, error) {

  ePrefix := "Decimal.NewInt32Exponent()"

  d2 := new(Decimal).New()

  bigI := big.NewInt(int64(int32Num))

  err := d2.bigINum.SetBigIntExponent(bigI, exponent)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err := d2.bigINum.SetBigIntExponent(bigI, exponent)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewInt64 - Returns a new Decimal instance based on input parameters
// 'int64Num' and 'precision'.
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'precision' indicates the number
// of digits to be formatted to the right of the decimal
// place. Input parameter 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction
// with the Decimal{} syntax thereby allowing Decimal
// type creation and initialization in one step.
//
//					int64Num := int64(123456)
//					precision := uint(3)
//					dec := Decimal{}.NewInt64(int64Num, precision)
//	       dec is now equal to 123.456
//
// Examples:
// ---------
//
//	  int64Num			precision			 Decimal Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (dec *Decimal) NewInt64(int64Num int64, precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewInt64()"

  d2 := new(Decimal).New()

  err := d2.SetInt64(int64Num, precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err := err := d2.SetInt64(int64Num, precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewInt64Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('int64Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//		decNum := Decimal{}.NewInt64Exponent(123456, -3)
//	 -- decNum is now equal to "123.456", precision = 3
//
//		decNum := Decimal{}.NewInt64Exponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  int64Num		 exponent			  	Decimal Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (dec *Decimal) NewInt64Exponent(
  int64Num int64, exponent int) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewInt64Exponent()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  bigI := big.NewInt(int64Num)

  err = d2.bigINum.SetBigIntExponent(bigI, exponent)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigIntExponent(bigI, exponent)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewFive - Returns a Decimal Type with a value equal to '5' (five).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '5', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								5
//			1								5.0
//	   	2								5.00
//			3								5.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
func (dec *Decimal) NewFive(precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewFive()"

  bINum, err := new(BigIntNum).NewFive(precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINum, err := new(BigIntNum).NewFive(precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())

  }

  d2 := new(Decimal).New()

  err = d2.SetBigIntNum(bINum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(bINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewFloat32 - Creates a new Decimal instance based on a float32
// input.
func (dec *Decimal) NewFloat32(f32 float32) (Decimal, error) {

  ePrefix := "Decimal.NewFloat32()"

  d2 := new(Decimal).New()
  err := d2.SetFloat32(f32)

  if err != nil {
    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by d2.SetFloat32(f32).\n"+
        "f32='%v' Error='%v' ",
        ePrefix,
        f32,
        err.Error())
  }

  return d2, nil
}

// NewFloat64 - Creates a new Decimal instance based on a float64
// input.
func (dec *Decimal) NewFloat64(
  f64 float64, maxPrecision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewFloat64()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  d2 := new(Decimal).New()

  bigFloat := big.NewFloat(f64)

  err = d2.bigINum.SetBigFloat(bigFloat, maxPrecision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigFloat(bigFloat, maxPrecision)",
        ErrContext: fmt.Sprintf("bigFloat= '%v'  maxPrecision= '%v'",
          bigFloat.Text('f', int(maxPrecision)), maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  err = d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewNumStrsMultiple - Used to create and return an array of Decimal Types.
// Input parameters are a series of number strings.
func (dec *Decimal) NewNumStrsMultiple(numStrs ...string) ([]Decimal, error) {

  ePrefix := "Decimal.NewNumStrsMultiple()"

  lenNumStrs := len(numStrs)

  decAry := make([]Decimal, lenNumStrs, lenNumStrs+100)

  var dec2 Decimal

  for i, numStr := range numStrs {

    dec := new(Decimal).New()

    err := dec.SetNumStr(numStr)

    if err != nil {

      return []Decimal{},
        fmt.Errorf("%v\n"+
          "Error returned by dec.SetNumStr(bigINum).\n"+
          "bigINum='%v'\nIndex='%v'\nError= %v\n",
          ePrefix,
          numStr,
          i,
          err.Error())
    }

    dec2, err = dec.CopyOut()

    if err != nil {

      return []Decimal{},
        fmt.Errorf("%v\n"+
          "Error returned by:\n"+
          "  dec2, err = dec.CopyOut().\n"+
          "Array Element: decAry[%v]\nError= %v\n",
          ePrefix,
          i,
          err.Error())
    }

    decAry[i] = dec2
  }

  return decAry, nil
}

// NewNumStrArray - Used to create and return an array of Decimal Types.
// The input parameter is an array of number strings.
func (dec *Decimal) NewNumStrArray(numStrs []string) ([]Decimal, error) {

  ePrefix := "Decimal.NewNumStrArray()"

  lenNumStrs := len(numStrs)

  decAry := make([]Decimal, lenNumStrs, lenNumStrs+100)

  var dec2 Decimal

  for i, numStr := range numStrs {

    dec := new(Decimal).New()

    err := dec.SetNumStr(numStr)

    if err != nil {

      return []Decimal{},
        fmt.Errorf("%v\n"+
          "Error returned by: \n"+
          " err := dec.SetNumStr(numStr)\n"+
          "Error= %v\n",
          ePrefix,
          err.Error())
    }

    dec2, err = dec.CopyOut()

    if err != nil {

      return []Decimal{},
        fmt.Errorf("%v\n"+
          "Error returned by: \n"+
          " dec2, err = dec.CopyOut()\n"+
          "Error= %v\n",
          ePrefix,
          err.Error())
    }

    decAry[i] = dec2

  }

  return decAry, nil
}

// NewNumStr - Returns a Decimal type based on a number string
// input parameter.
//
// This method assumes that the input parameter 'numStr' is a string
// of numeric digits which may be delimited by default USA numeric
// separators. Default USA numeric separators are defined as:
//
//	 	decimal separator = '.'
//	   thousands separator = ','
//			currency symbol = '$'
//
// If the subject 'numStr' employs other national or cultural numeric
// separators, see method Decimal.NewNumStrWithNumSeps(), below.
//
// The 'NewNumStr' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewNumStr("123.456")
func (dec *Decimal) NewNumStr(numStr string) (Decimal, error) {

  ePrefix := "Decimal.NewNumStr()"

  d2 := new(Decimal).New()

  err := d2.SetNumStr(numStr)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by\n"+
        " err := d2.SetNumStr(numStr)\n"+
        "numStr='%v'\nError= %v\n",
        ePrefix,
        numStr,
        err.Error())
  }

  return d2, nil
}

// NewNumStrWithNumSeps - Receives a number string as input and returns a
// new Decimal instance. The input parameter 'numSeps' contains numeric
// separators (decimal separator, thousands separator and currency symbol)
// which will be used to parse the number string.
//
// In addition, the numeric separators contained in input parameter 'numSeps'
// will be copied to the returned Decimal instance.
func (dec *Decimal) NewNumStrWithNumSeps(
  numStr string,
  numSeps NumericSeparatorDto) (Decimal, error) {

  ePrefix := "Decimal.NewNumStrWithNumSeps()"

  numSeps.SetDefaultsIfEmpty()

  d2 := new(Decimal).New()

  err := d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err := d2.SetNumericSeparatorsDto(numSeps)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2.SetNumStr(numStr)

  if err != nil {
    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by:\n"+
        "d2.SetNumStr(numStr)\n"+
        "numStr='%v'\nError= %v\n",
        ePrefix,
        numStr,
        err.Error())
  }

  return d2, nil
}

// NewNumStrDto - Returns a Decimal type based on a NumStrDto
// input parameter. If an error is encountered, it will trigger
// a panic condition.
//
// The 'NewNumStrDto' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewNumStrDto(numDto)
func (dec *Decimal) NewNumStrDto(
  numDto NumStrDto) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewNumStrDto()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  err = numDto.IsValid(ePrefix.XCpy("Validating input param 'numDto'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numDto.IsValid(ePrefix.XCpy(\"Validating input param 'numDto'\").String())",
        ErrContext: "Error: Input parameter 'numDto' is INVALID! ",
        ErrMessage: err.Error(),
      }
  }

  numDtoNumStr, err := numDto.GetNumStr()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numDtoNumStr, err := numDto.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  d2 := new(Decimal).New()

  err = d2.SetNumStrDto(numDto)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetNumStrDto(numDto)",
        ErrContext: fmt.Sprintf("numDto= '%v'", numDtoNumStr),
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewNumStrPrecision - Returns a Decimal type based on a number string
// and a precision value received as input parameters.
//
// The 'NewNumStrPrecision' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation and initialization
// in one step.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example: Decimal{}.NewNumStrPrecision('123456', 3, false) = 123.456
func (dec *Decimal) NewNumStrPrecision(numStr string, precision uint, roundResult bool) (Decimal, error) {

  ePrefix := "Decimal.NewNumStrPrecision() "

  d2, err := new(Decimal).NewPtr().NumStrPrecisionToDecimal(numStr, precision, roundResult)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d2, err := new(Decimal).NewPtr().NumStrPrecisionToDecimal(numStr, precision, roundResult)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewOne - Returns a Decimal Type with a value equal to '1' (one).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								1
//			1								1.0
//			2								1.00
//			3								1.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
func (dec *Decimal) NewOne(precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewOne()"

  bINum, err := new(BigIntNum).NewOne(precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINum, err := new(BigIntNum).NewOne(precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d2 := new(Decimal).New()

  err = d2.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetNumericSeparatorsToDefaultIfEmpty()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2.SetBigIntNum(bINum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(bINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewRationalNum - Creates a new Decimal instance based on input parameters consisting
// of a Rational Number (*big.Rat) and the specified 'precision' to be implemented in
// the resulting Decimal number value.
//
// For information on *big.Rat see https://golang.org/pkg/math/big/
func (dec *Decimal) NewRationalNum(
  bigRat *big.Rat, maxPrecision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewRationalNum()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  if bigRat == nil {

    return Decimal{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'bigRat'",
      }
  }

  numSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  d2, err := new(Decimal).NewZero(0)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "d2, err := new(Decimal).NewZero(0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.bigINum.SetBigRat(bigRat, maxPrecision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.bigINum.SetBigRat(bigRat, maxPrecision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetNumericSeparatorsDto(numSeps)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewTen - Returns a Decimal Type with a value equal to '10' (ten).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '10', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0							 10
//			1							 10.0
//			2							 10.00
//			3							 10.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
func (dec *Decimal) NewTen(precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewTen()"

  bINum, err := new(BigIntNum).NewTen(precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINum, err := new(BigIntNum).NewTen(precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d2 := new(Decimal).New()

  err = d2.SetBigIntNum(bINum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(bINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewThree - Returns a Decimal Type with a value equal to '3' (three).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '3', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								3
//			1								3.0
//			2								3.00
//			3								3.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
func (dec *Decimal) NewThree(precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewThree()"

  bINum, err := new(BigIntNum).NewThree(precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINum, err := new(BigIntNum).NewThree(precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d2 := new(Decimal).New()

  err = d2.SetBigIntNum(bINum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(bINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewTwo - Returns a Decimal Type with a value equal to '2' (two).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '2', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								2
//			1								2.0
//			2								2.00
//			3								2.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
func (dec *Decimal) NewTwo(precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewTwo()"

  bINum, err := new(BigIntNum).NewTwo(precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINum, err := new(BigIntNum).NewTwo(precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d2 := new(Decimal).New()

  err = d2.SetBigIntNum(bINum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(bINum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NewUint - Returns a new Decimal instance based on input parameters
// 'uintNum' and 'precision'.
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//					uintNum := uint(123456)
//					precision := uint(3)
//					dec := Decimal{}.NewUint(int32Num, precision)
//	       dec is now equal to 123.456
//
// Examples:
// ---------
//
//	 uintNum			 precision	  	  Decimal Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (dec *Decimal) NewUint(
  uintNum uint, precision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewUint()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  err = d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigInt(\n" +
          "  big.NewInt(0).SetUint64(uint64(uintNum)), precision)",
        ErrContext: fmt.Sprintf("uintNum= '%v'  precision= '%v'",
          uintNum, precision),
        ErrMessage: err.Error(),
      }
  }

  decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.bigINum.SetNumericSeparatorsDto(decNumSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetNumericSeparatorsDto(decNumSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewUintExponent - This method returns a new Decimal instance in which
// the numeric value is set using an integer ('uintNum') multiplied by 10
// raised to the power of the input parameter, 'exponent'.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//		decNum := Decimal{}.NewUintExponent(123456, -3)
//	 -- decNum is now equal to "123.456", precision = 3
//
//		decNum := Decimal{}.NewUintExponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 uintNum		    exponent		  	Decimal Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456					 0              123456
func (dec *Decimal) NewUintExponent(
  uintNum uint, exponent int) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewUintExponent()",
    "")

  if err != nil {
    return Decimal{}, err
  }
  d2 := new(Decimal).New()

  err = d2.bigINum.SetBigIntExponent(
    big.NewInt(0).SetUint64(uint64(uintNum)),
    exponent)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigIntExponent(\n" +
          "  big.NewInt(0).SetUint64(uint64(uintNum)),exponent)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  decBigINumNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "decBigINumNumSeps, err := \n" +
          "  dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.SetNumericSeparatorsDto(decBigINumNumSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetNumericSeparatorsDto(decBigINumNumSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewUint32 - Returns a new Decimal instance based on input parameters
// 'uint32Num' and 'precision'.
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//					uint32Num := uint32(123456)
//					precision := uint(3)
//					dec := Decimal{}.NewUint32(int32Num, precision)
//	       dec is now equal to 123.456
//
// Examples:
// ---------
//
//	 uint32Num			precision			Decimal Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (dec *Decimal) NewUint32(
  uint32Num uint32, precision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewUint32()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  err = d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)",
        ErrContext: fmt.Sprintf("uint32Num= '%v'  precision= '%v'",
          uint32Num, precision),
        ErrMessage: err.Error(),
      }
  }

  decBigINumNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.bigINum.SetNumericSeparatorsDto(decBigINumNumSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetNumericSeparatorsDto(decBigINumNumSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewUint32Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('uint32Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//		decNum := Decimal{}.NewUint32Exponent(123456, -3)
//	 -- decNum is now equal to "123.456", precision = 3
//
//		decNum := Decimal{}.NewUint32Exponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 uint32Num		 exponent			  	Decimal Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (dec *Decimal) NewUint32Exponent(
  uint32Num uint32, exponent int) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewUint32Exponent()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  err = d2.bigINum.SetBigIntExponent(
    big.NewInt(0).SetUint64(uint64(uint32Num)),
    exponent)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigIntExponent(\n" +
          "  big.NewInt(0).SetUint64(uint64(uint32Num)),exponent)",
        ErrContext: fmt.Sprintf("uint32Num= '%v'  exponent= '%v'",
          uint32Num, exponent),
        ErrMessage: err.Error(),
      }
  }

  decBigINumNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decBigINumNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.SetNumericSeparatorsDto(decBigINumNumSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.SetNumericSeparatorsDto(decBigINumNumSeps)",
        ErrContext: fmt.Sprintf("decBigINumNumSeps= '%v'",
          decBigINumNumSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewUint64 - Returns a new Decimal instance based on input parameters
// 'uint64Num' and 'precision'.
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//					uint64Num := uint64(123456)
//					precision := uint(3)
//					dec := Decimal{}.NewUint64(int32Num, precision)
//	       dec is now equal to 123.456
//
// Examples:
// ---------
//
//	 uint64Num			precision			Decimal Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (dec *Decimal) NewUint64(
  uint64Num uint64, precision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewUint64()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  err = d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  decNumStr, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decNumStr, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.bigINum.SetNumericSeparatorsDto(decNumStr)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetNumericSeparatorsDto(decNumStr)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewUint64Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('uint64Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//		decNum := Decimal{}.NewUint64Exponent(123456, -3)
//	 -- decNum is now equal to "123.456", precision = 3
//
//		decNum := Decimal{}.NewUint64Exponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 uint64Num		 exponent			  	Decimal Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (dec *Decimal) NewUint64Exponent(
  uint64Num uint64, exponent int) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NewUint64Exponent()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  d2 := new(Decimal).New()

  err = d2.bigINum.SetBigIntExponent(
    big.NewInt(0).SetUint64(uint64Num),
    exponent)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetBigIntExponent(big.NewInt(0).SetUint64(uint64Num),exponent)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = d2.SetNumericSeparatorsDto(decNumSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// NewZero - Creates a New Decimal Instance with a value of zero. Input
// parameter 'precision' indicates the number of zeros formatted to the
// right of the decimal place.
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								0
//			1								0.0
//			2								0.00
//			3								0.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
func (dec *Decimal) NewZero(precision uint) (Decimal, error) {

  ePrefix := "Decimal.NewZero()"

  biNum, err := new(BigIntNum).NewZero(precision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " biNum, err := new(BigIntNum).NewZero(precision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())

  }

  d2 := new(Decimal).New()

  err = d2.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetNumericSeparatorsToDefaultIfEmpty()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2.SetBigIntNum(biNum)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d2.SetBigIntNum(biNum)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d2, nil
}

// NthRoot - Calculates the nth root of the current Decimal value. The numeric value of
// the current Decimal instance constitutes the radicand.
//
// Input Parameters:
// =================
//
//	 nthRoot Decimal - Nth root specifies the root which will be calculated using the current
//											Decimal instance as the radicand.
//											Example, square root, cube root, 4th root, 9th root etc.
//
// maxPrecision uint -  Specifies the maximum number of decimals to the right of the decimal
//
//	place to which the Nth root will be calculated. If the internal
//	calculation exceeds the limit the nth root result will be rounded
//	to 'maxPrecision' decimal places.
//
// Returns:
// ========
// The nth root calculation result is returned as a Decimal instance. The returned Decimal
// instance will contain	numeric separators (decimal separator, thousands separator and
// currency symbol) copied from the current Decimal instance (radicand).
func (dec *Decimal) NthRoot(
  nthRoot Decimal, maxPrecision uint) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NthRoot()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  err = dec.bigINum.IsValid(ePrefix.XCpy("Validating 'dec'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = dec.bigINum.IsValid(ePrefix.XCpy(\"Validating 'dec'\").String())",
        ErrContext: "Error: Current Decimal instance 'dec' is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "'nthRoot' input parameter is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  // If the radicand is zero, the result will always be zero
  var decIsZero, nthRootIsZero bool

  decIsZero, err = dec.bigINum.IsZero()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decIsZero, err = dec.bigINum.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if decIsZero {

    dec3, err := new(Decimal).NewZero(0)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "dec3, err := new(Decimal).NewZero(0)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    return dec3, nil
  }

  nthRootIsZero, err = nthRoot.IsZero()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootIsZero, err = nthRoot.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // If nth root is zero, the result is always one.
  if nthRootIsZero {

    dec4, err := new(Decimal).NewOne(0)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "dec4, err := new(Decimal).NewOne(0)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    return dec4, nil
  }

  nthRootIsEven, err := nthRoot.IsEvenNumber()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootIsEven, err := nthRoot.IsEvenNumber()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  decSignValue, err := dec.bigINum.GetSign()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decSignValue, err := dec.bigINum.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if decSignValue == -1 && nthRootIsEven {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "if decSignValue == -1 && nthRootIsEven {",
        ErrMessage: "INVALID ENTRY! Cannot calculate nth root of a\n" +
          "negative radicand when nthRoot is even.\n" +
          "The sign of 'dec' is -1 and the nthRoot is 'even'.",
      }
  }

  decNthRoot := new(Decimal).New()

  decNthRoot.bigINum, err =
    new(BigIntMathNthRoot).GetNthRoot(dec.bigINum, nthRoot.bigINum, maxPrecision)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "decNthRoot.bigINum, err = new(BigIntMathNthRoot).\n" +
          "  GetNthRoot(dec.bigINum, nthRoot.bigINum, maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision='%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  err = decNthRoot.IsValid(ePrefix.XCpy("Validating 'decNthRoot'").String())

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = decNthRoot.IsValid(ePrefix.XCpy(\"Validating 'decNthRoot'\").String())",
        ErrContext: "Validation Error on Decimal instance 'decNthRoot'.",
        ErrMessage: err.Error(),
      }
  }

  return decNthRoot, nil
}

// NumStrPrecisionToDecimal - receives a number string and a
// precision value as parameters. This method creates a Decimal
// Type containing the converted numeric value and returns it.
// For example, if passed the string ('str') '123456' and a precision
// value of '3', the resulting Decimal value would be 123.456.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example Usage:
// d := Decimal{}.NewBigIntNum()
// d2, err := d.NumStrPrecisionToDecimal("123456", 3, false)
// d2 is Now Equal to 123.456
func (dec *Decimal) NumStrPrecisionToDecimal(
  numStr string,
  requiredPrecision uint,
  roundResult bool) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Decimal.NumStrPrecisionToDecimal()",
    "")

  if err != nil {
    return Decimal{}, err
  }

  decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "decNumSeps, err := dec.bigINum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  d2 := Decimal{}

  d2.bigINum, err = new(BigIntNum).NewNumStr(numStr)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "d2.bigINum, err = new(BigIntNum).NewNumStr(numStr)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if roundResult == true {

    err = d2.bigINum.SetPrecision(requiredPrecision)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = d2.bigINum.SetPrecision(requiredPrecision)",
          ErrContext: fmt.Sprintf("requiredPrecision='%v'", requiredPrecision),
          ErrMessage: err.Error(),
        }
    }

  } else {

    err = d2.bigINum.TruncToDecPlace(requiredPrecision)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = d2.bigINum.TruncToDecPlace(requiredPrecision)",
          ErrContext: fmt.Sprintf("requiredPrecision='%v'", requiredPrecision),
          ErrMessage: err.Error(),
        }
    }
  }

  err = d2.bigINum.SetNumericSeparatorsDto(decNumSeps)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = d2.bigINum.SetNumericSeparatorsDto(decNumSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return d2, nil
}

// Pow - Raises the numeric value of the current Decimal instance to the power of
// input parameter Decimal Type 'exponent'. The result is returned as a new Decimal
// instance.
//
// Note that this method can process positive, negative, integer and fractional
// exponents.
//
// exponent Decimal -		The numerical value of the current Decimal instance
//
//	will be raised to the power of 'exponent'.
//	  			 result = dec^exponent
//
// maxPrecision uint - 	Determines the maximum number of digits
//
//	to the right of the	decimal point returned
//	in the result. The actual precision may be
//	less than 'maxPrecision'.
//
// The returned Decimal type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal instance.
func (dec *Decimal) Pow(exponent Decimal, maxPrecision uint) (Decimal, error) {

  ePrefix := "Decimal.Pow() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error returned by: \n"+
        " err := dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  d3 := new(Decimal).New()

  d3.bigINum, err = new(BigIntMathPower).Pwr(dec.bigINum, exponent.bigINum, maxPrecision)

  if err != nil {
    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d3.bigINum, err = BigIntMathPower{}.Pwr(dec.bigINum, exponent.bigINum, maxPrecision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d3.IsValid(ePrefix + "d3 INVALID! ")

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on Decimal instance 'dec3'.\n"+
        "Error returned by: \n"+
        " err = d3.IsValid(ePrefix + \"d3 INVALID! \") \n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d3, nil
}

// PowInt - raises the current Decimal to the power of an integer 'exponent'.
// The result is returned as a Decimal type.
//
// Input Parameters:
//
// exponent 		int	 -	The numerical value of the current Decimal instance
//
//	will be raised to the power of 'exponent'.
//	  			 result = dec^exponent
//
// maxPrecision uint - 	Determines the maximum number of digits
//
//	to the right of the	decimal point returned
//	in the result. The actual precision may be
//	less than 'maxPrecision'.
//
// The returned Decimal type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal instance.
func (dec *Decimal) PowInt(
  exponent int, maxPrecision uint) (Decimal, error) {

  ePrefix := "Decimal.PowInt() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on current Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err = dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  biNumExponent, err := new(BigIntNum).NewBigInt(big.NewInt(int64(exponent)), 0)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " biNumExponent, err := new(BigIntNum).NewBigInt(big.NewInt(int64(exponent)), 0)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  bINumResult, err := new(BigIntMathPower).Pwr(dec.bigINum, biNumExponent, maxPrecision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINumResult, err := BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision)\n"+
        "dec.bigINum=%v\n"+
        "biNumExponent=%v\n"+
        "maxPrecision=%v\n"+
        "Error= %v\n",
        ePrefix,
        dec.bigINum.GetNumStr(),
        biNumExponent.GetNumStr(),
        maxPrecision,
        err.Error())
  }

  d3, err := new(Decimal).NewBigIntNum(bINumResult)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d3, err := new(Decimal).NewBigIntNum(bINumResult)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d3.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d3.SetNumericSeparatorsDto(numSeps)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d3, nil
}

// SetBigIntNum - Sets the numeric value of the current Decimal
// instance to that of input parameter, 'bigINum'
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// The numeric separators associated with the incoming 'bigINum' are
// not copied to the current Decimal instance.
func (dec *Decimal) SetBigIntNum(bigINum BigIntNum) error {

  ePrefix := "Decimal.SetBigIntNum()"

  var err error

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  dec.bigINum, err = bigINum.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = bigINum.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetBigInt - Sets the value of the current Decimal instance to the
// input parameter 'iBig' scaled to the value of precision. In other
// words, if 'iBig' is set to a value of '123456' and precision is
// set to '3', the current Decimal will be set to a value of '123.456'.
//
// Using the same example, a big int value of '123456' with a precision of
// zero ('0') will yield an integer of '123456'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// iBig := big.NewInt(int64(123))
// d.SetBigInt(iBig, 1)
// This yields a numeric value of d = 12.3
func (dec *Decimal) SetBigInt(
  iBig *big.Int, precision uint) error {

  ePrefix := "Decimal.SetBigInt()"

  var err error

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  dec.bigINum, err = new(BigIntNum).NewBigInt(iBig, precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = new(BigIntNum).NewBigInt(iBig, precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.SetNumericSeparatorsDto(numSeps)\n\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetCurrencySymbol - sets the character which serves as the
// currency symbol for this Decimal value. Currency defaults
// to '$'.
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// For a listing of Major World Currency Symbols in Unicode format,
// see array 'NumStrCurrencySymbols' in source file:
//
//	MikeAustin71/mathopsgo/mathops/mathopsconstants.go
func (dec *Decimal) SetCurrencySymbol(currencySymbol rune) error {

  dec.bigINum.SetCurrencySymbol(currencySymbol)

  return nil
}

// SetDecimalSeparator - sets the character which separates
// the number into integer and fractional components. This
// defaults to the period '.'
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// In the USA the Decimal Separator character is a period ('.')
func (dec *Decimal) SetDecimalSeparator(decimalSeparator rune) error {

  dec.bigINum.SetDecimalSeparator(decimalSeparator)

  return nil
}

// SetFloat32 - Sets the value of the current decimal to
// that of the passed-in float32 parameter.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// f32:= float32(123.456)
// d.SetFloat32(f32)
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetFloat32(f32 float32) error {

  ePrefix := "Decimal.SetFloat32() "

  err := dec.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  bigFloat := big.NewFloat(float64(f32))

  precision := bigFloat.Prec()

  err = dec.bigINum.SetBigFloat(bigFloat, precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.SetBigFloat(bigFloat, precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetFloat64 - Sets the value of the current decimal to
// that of the passed-in float64 parameter.
//
// Example usage:
// d:= Decimal{}.NewZero(0)
// f64:= float64(123.456)
// d.SetFloat32(f64)
// Number String = "123.456"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetFloat64(f64 float64) error {

  ePrefix := "Decimal.SetFloat64()"

  err := dec.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  bigFloat := big.NewFloat(f64)

  precision := bigFloat.Prec()

  err = dec.bigINum.SetBigFloat(bigFloat, precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.SetBigFloat(bigFloat, precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetFloatBig - Sets the value of the current Decimal to the
// passed-in *big.Float parameter.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// bigFloat:= big.NewFloat(float64(123.456))
// d.SetBigFloat(bigFloat)
// Number String = "123.456"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetFloatBig(bigFloat *big.Float) error {

  ePrefix := "Decimal.SetFloatBig() "

  err := dec.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.SetBigFloat(bigFloat, bigFloat.Prec())

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.SetBigFloat(bigFloat, bigFloat.Prec())\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetInt - Sets the value of the current Decimal to the input parameter 'iNum'
// scaled to the value of input parameter 'precision'. In other words, if 'iNum'
// is set to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(9876)
// The numeric value of 'd' is now, 9876.
//
// d.SetInt(123456, 3) yields a new value for 'd' of 123.456.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetInt(
  iNum int, precision uint) error {

  ePrefix := "Decimal.SetInt()"

  iBig := big.NewInt(int64(iNum))

  err := dec.bigINum.SetBigInt(iBig, precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetBigInt(iBig, precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetInt32 - Sets the value of the current Decimal to the input parameter
// 'int32Num' scaled to the value of input parameter, 'precision'. In other
// words, if 'int64Num' is set to a value of '123456' and precision is set
// to '3', the current Decimal will be set to a numeric value of '123.456'.
//
// Using the same example, an int32 value of '123456' and a precision
// value of zero ('0') will yield an integer value of '123456'.
//
// Input parameter 'int32Num' is an int32 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) yields numeric value of decimal instance 'd'
//
//	equals 956789.
//
// int32Num:= int32(123456)
// d.SetInt32(int32Num, 3) sets the numeric value of Decimal instance 'd' to 123.456.
func (dec *Decimal) SetInt32(
  int32Num int32, precision uint) error {

  ePrefix := "Decimal.SetInt32()"

  err := dec.bigINum.SetBigInt(big.NewInt(int64(int32Num)), precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetBigInt(big.NewInt(int64(int32Num)), precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetInt64 - Sets the value of the current Decimal to the input parameter
// 'i64' scaled to the value of precision. In other words, if 'i64' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a numeric value of '123.456'.
//
// Using the same example, an int64 value of '123456' and a precision
// value of zero ('0') will yield an integer value of '123456'.
//
// Input parameter 'i64' is an int64 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//
//	equals 956789.
//
// i64:= int64(123456)
// d.SetInt64(i64, 3) sets the numeric value of Decimal instance 'd' to 123.456.
func (dec *Decimal) SetInt64(
  i64 int64, precision uint) error {

  ePrefix := "Decimal.SetInt64()"

  err := dec.bigINum.SetBigInt(big.NewInt(i64), precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum.SetBigInt(big.NewInt(i64), precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetIntFracStrings - Sets the value of a decimal based on separate
// integer and fraction strings passed as input parameters. 'intNum'
// parameter will constitute the integer component of the Decimal value
// while 'fracNum' will be converted to the fractional component of the
// new Decimal value. 'fracNum' represents all the numeric digits to
// the right of the decimal place, while 'intNum' represents all the
// integer digits to the left of the decimal place.
//
// The parameter 'signVal' will determine the sign Value for the returned
// Decimal type. It should be set to either +1 or -1.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetIntFracStrings(
  intNum, fracNum string, signVal int) error {

  ePrefix := "Decimal.SetIntFracStrings() "

  var err error

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  binIntNum := BigIntNum{}

  err = binIntNum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = binIntNum.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = binIntNum.SetIntFracStrings(intNum, fracNum, signVal)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = binIntNum.SetIntFracStrings(intNum, fracNum, signVal)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.CopyIn(&binIntNum)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.CopyIn(&binIntNum)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumericSeparatorsDto - Sets the values of numeric separators:
//
//	decimal point separator
//	thousands separator
//	currency symbol
//
// based on values transmitted through input parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators' is set
// to zero, an error will be returned.
func (dec *Decimal) SetNumericSeparatorsDto(customSeparators NumericSeparatorDto) error {

  ePrefix := "Decimal.SetNumericSeparatorsDto()"

  err := dec.bigINum.SetNumericSeparatorsDto(customSeparators)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetNumericSeparatorsDto(customSeparators)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumericSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Note: If zero values are submitted as input, the values will default to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
func (dec *Decimal) SetNumericSeparators(
  decimalSeparator, thousandsSeparator, currencySymbol rune) error {

  ePrefix := "Decimal.SetNumericSeparators()"

  err := dec.bigINum.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumericSeparatorsToDefaultIfEmpty - If numeric separators are
// set to zero or nil, this method will set those numeric
// separators to the USA defaults. This means that the
// Decimal separator is set to ('.'), the Thousands separator
// is set to (',') and the currency symbol is set to '$'.
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that numeric separators
// are set to valid values.
func (dec *Decimal) SetNumericSeparatorsToDefaultIfEmpty() error {

  ePrefix := "Decimal.SetNumericSeparatorsToDefaultIfEmpty()"

  err := dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumericSeparatorsToUSADefault - Sets Numeric separators:
//
//	Decimal Point Separator
//	Thousands Separator
//	Currency Symbol
//
// to the United States of America (USA) defaults.
//
// Call specific methods to set numeric separators for other countries or
// cultures:
//
//	dec.SetDecimalSeparator()
//	dec.SetThousandsSeparator()
//	dec.SetCurrencySymbol()
func (dec *Decimal) SetNumericSeparatorsToUSADefault() error {

  ePrefix := "Decimal.SetNumericSeparatorsToUSADefault()"

  err := dec.SetDecimalSeparator('.')

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.SetDecimalSeparator('.')\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.SetThousandsSeparator(',')

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.SetThousandsSeparator(',')\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.SetCurrencySymbol('$')

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.SetCurrencySymbol('$')\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumStrPrecision - Sets the Decimal's value to a number string and
// applies the appropriate 'precision' in order to determine placement
// of the decimal point. For example, if the number string ('str') is passed
// in as "123456" with a 'precision value of '3', the result is a value of
// 123.456.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example Usage:
// d:= Decimal{}.NewBigIntNum()
// d.SetNumStrPrecision("123456", 3)
// Resulting Decimal Value = 123.456
func (dec *Decimal) SetNumStrPrecision(str string, precision uint, roundResult bool) error {

  ePrefix := "Decimal.SetNumStrPrecision() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance 'dec'.\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  d2, err := dec.NewNumStrPrecision(str, precision, roundResult)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " d2, err := dec.NewNumStrPrecision(str, precision, roundResult)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = d2.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.CopyIn(d2)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.CopyIn(d2)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumStr - Set's the Decimal's value to the input
// parameter 'str'. For example if 'str' is set equal
// to '123.456', this method will set the Decimal's
// value to 123.456.
//
// Example Usage:
// d := Decimal{}.NewBigIntNum()
// d.SetNumStr("123.456")
// Decimal Value = 123.456
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetNumStr(str string) error {

  ePrefix := "Decimal.SetNumStr() "

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  d2, err := dec.NumStrToDecimal(str)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " d2, err := dec.NumStrToDecimal(str)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = d2.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = d2.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.CopyIn(d2)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.CopyIn(d2)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetNumStrDto - Sets the value of the current Decimal type
// to the value represented by the incoming NumStrDto parameter.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetNumStrDto(nDto NumStrDto) error {

  ePrefix := "Decimal.SetNumStrDto() "

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  bIntNum, err := new(BigIntNum).NewNumStrDto(nDto)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " bIntNum, err := new(BigIntNum).NewNumStrDto(nDto)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = bIntNum.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = bIntNum.SetNumericSeparatorsDto(numSeps)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  dec.bigINum, err = bIntNum.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = bIntNum.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetPrecisionRound - Sets the precision or
// scale of the Decimal value. precision determines
// the number of digits displayed to the right of
// the decimal place. Note that precision is
// processed as an unsigned integer.
//
// When reducing precision, existing digits
// are ROUNDED! When increasing precision,
// additional zeros ('0') are added to the right
// of the decimal place.
func (dec *Decimal) SetPrecisionRound(precision uint) error {

  ePrefix := "Decimal.SetPrecisionRound() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance 'dec'.\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.SetPrecision(precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.SetPrecision(precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetPrecisionTrunc - Sets the precision or
// scale of the Decimal value. precision determines
// the number of digits displayed to the right of
// the decimal place. Note that precision is
// processed as an unsigned integer.
//
// When reducing precision, existing digits
// are TRUNCATED! When increasing precision,
// additional zeros ('0') are added to the right
// of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SetPrecisionTrunc(precision uint) error {

  ePrefix := "Decimal.SetPrecisionTrunc() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance 'dec'.\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.TruncToDecPlace(precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.TruncToDecPlace(precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetSign - Sets the sign of the numeric value
// for the current Decimal instance. Only two values
// are allowed: +1 and -1.
//
// If any other value is passed an error is thrown.
func (dec *Decimal) SetSign(newSignVal int) error {

  ePrefix := "Decimal.SetSign()"

  err := dec.bigINum.SetSignValue(newSignVal)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetSignValue(newSignVal)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetThousandsSeparator - sets the character which serves
// as the 'thousands' separator.
//
// In the USA, the Thousands Separator character is the comma (',').
// The Decimal Thousands Separator value defaults to the comma (',').
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
func (dec *Decimal) SetThousandsSeparator(thousandsSeparator rune) error {

  dec.bigINum.SetThousandsSeparator(thousandsSeparator)

  return nil
}

// SetUint - Sets the value of the current Decimal instance to the input
// parameter 'uintNum' scaled to the value of input parameter 'precision'.
// In other words, if 'uintNum' is set to a value of '123456' and precision
// is set to '3', the current Decimal will be set to a numeric value of
// '123.456'.
//
// Using the same example, an uint32 value of '123456' and a precision value
// of zero ('0') will yield a numeric value of '123456'.
//
// Input parameter 'uintNum' is an uint type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//
//	equals 956789.
//
// uintNum := uint(123456)
// d.SetUint(uintNum, 3) sets the numeric value of Decimal instance 'd' to 123.456.
func (dec *Decimal) SetUint(uintNum uint, precision uint) error {

  ePrefix := "Decimal.SetUint()"

  err := dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err:= dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetUint32 - Sets the value of the current Decimal instance to the input
// parameter 'uint32Num' scaled to the value of input parameter 'precision'.
// In other words, if 'uint32Num' is set to a value of '123456' and precision
// is set to '3', the current Decimal will be set to a numeric value of
// '123.456'.
//
// Using the same example, an uint32 value of '123456' and a precision value
// of zero ('0') will yield a numeric value of '123456'.
//
// Input parameter 'uint32Num' is an uint32 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//
//	equals 956789.
//
// uint32Num := uint32(123456)
// d.SetUint32(uint32Num, 3) sets the numeric value of Decimal instance 'd' to 123.456.
func (dec *Decimal) SetUint32(uint32Num uint32, precision uint) error {

  ePrefix := "Decimal.SetUint32() "

  err := dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SetUint64 - Sets the value of the current Decimal instance to the input
// parameter 'uint64Num' scaled to the value of input parameter 'precision'.
// In other words, if 'uint64Num' is set to a value of '123456' and precision
// is set to '3', the current Decimal will be set to a numeric value of
// '123.456'.
//
// Using the same example, an uint64 value of '123456' and a precision value
// of zero ('0') will yield a numeric value of '123456'.
//
// Input parameter 'uint64Num' is an uint64 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//
//	equals 956789.
//
// uint64Num := uint64(123456)
// d.SetUint64(uint64Num, 3) sets the numeric value of Decimal instance 'd' to 123.456.
func (dec *Decimal) SetUint64(uint64Num uint64, precision uint) error {

  ePrefix := "Decimal.SetUint64()"

  err := dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// ShiftPrecisionLeft - Shifts precision of the current Decimal instance
// numeric value to the left by 'shiftLeftPlaces' decimal places. This
// is a 'relative' shift-left operation. The shift left operation is
// therefore performed with the current decimal point position as the
// starting point.
//
// This operation is equivalent to:	result = Decimal value / 10^shiftLeftPlaces
// or signed number divided by 10 raised to the power of shiftLeftPlaces.
//
// This method performs a relative shift left of the decimal point position.
// Be careful, this is NOT Shift Number Left operation. This is Shift Precision
// Left which means that the decimal point will be shifted left.
//
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftLeftPlaces int	- The number of positions the decimal point will be
//												shifted left from its current position.
//
// Examples:
// =========
//
//	shift-left
//
// signed Number		  places				Result
//
//	"123456.789"				3						"123.456789"
//	"123456.789"				2						"1234.56789"
//	"123456.789"        6					  "0.123456789"
//	"123456789"	 			  6						"123.456789"
//	"123"               5	          "0.00123"
//	"0"								  3						"0"
//	"123456.789"				0						"123456.789"		- zero has no effect on original number string
//
// "-123456.789"        0          "-123456.789"
// "-123456.789"        3          "-123.456789"
// "-123456789"			    6					 "-123.456789"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) ShiftPrecisionLeft(shiftLeftPlaces uint) error {

  ePrefix := "Decimal.ShiftPrecisionLeft() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance 'dec'.\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.ShiftPrecisionLeft(shiftLeftPlaces)

  err = dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance (dec) returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "This validation error occurred after a Shift Left Operation!\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// ShiftPrecisionRight - Shifts precision of the current Decimal instance
// numeric value to the right by 'shiftRightPlaces' decimal places. This
// is a 'relative' shift-right operation. The shift right operation is
// therefore performed with the current decimal point position as the
// starting point.
//
// This is equivalent to: result = Decimal value X 10^shiftRightPrecision or
// Decimal numeric value multiplied by 10 raised to the power of
// shiftRightPrecision.
//
// This method performs a relative shift right of the decimal point position.
// Be careful, this is NOT a Shift Number Right operation. This is Shift Precision
// Right which means that the decimal point will be shifted right.
//
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftRightPlaces int	- The number of positions the decimal point will be
//													shifted right from its current position.
//
// Examples:
// =========
//
//	shift-right
//
// signed Number		  places				Result
//
//	"123456.789"				3						"123456789"
//	"123456.789"				2						"12345678.9"
//	"123456.789"        6					  "123456789000"
//	"123456789"	 			  6						"123456789000000"
//	"123"               5	          "12300000"
//	"0"								  3						"0"
//	"123456.789"				0						"123456.789"		- zero has no effect on original number string
//
// "-123456.789"        0          "-123456.789"
// "-123456.789"        3          "-123456789"
// "-123456789"			    6					 "-123456789000000"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) ShiftPrecisionRight(shiftRightPlaces uint) error {

  ePrefix := "Decimal.ShiftPrecisionRight() "

  err := dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance 'dec'.\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.ShiftPrecisionRight(shiftRightPlaces)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err = dec.bigINum.ShiftPrecisionRight(shiftRightPlaces)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance (dec) returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "This validation error occurred after a Shift Right Operation!\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SquareRoot - Returns a Decimal instance with a numeric value equal to the
// square root of the current Decimal numeric value. The current Decimal instance
// is the radicand.
//
// Note: If the current Decimal value is a negative value, an error will be generated.
// You cannot take the square root of a negative number.
//
// Returns:
// ========
// The calculation result is returned as a Decimal instance. The returned Decimal instance
// will contain	numeric separators (decimal separator, thousands separator and currency symbol)
// copied from the current Decimal instance (dec).
func (dec *Decimal) SquareRoot(maxPrecision uint) (Decimal, error) {

  ePrefix := "Decimal.SquareRoot() "

  err := dec.bigINum.IsValid(ePrefix + "Current Decimal instance is INVALID! ")

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on current Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err = dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  if dec.GetSign() == -1 {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "INVALID ENTRY!\n"+
        "Cannot calculate nth root of a negative radicand.\n"+
        "Decimal sign == -1\n",
        ePrefix)
  }

  bINumTwo, err := new(BigIntNum).NewTwo(0)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " bINumTwo, err := new(BigIntNum).NewTwo(0)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  decSqRoot := new(Decimal).New()

  decSqRoot.bigINum, err =
    new(BigIntMathNthRoot).GetNthRoot(dec.bigINum, bINumTwo, maxPrecision)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " decSqRoot.bigINum, err = BigIntMathNthRoot{}.GetNthRoot(dec.bigINum, bINumTwo, maxPrecision)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())

  }

  err = decSqRoot.IsValid(ePrefix + "decSqRoot INVALID! ")

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on current Decimal instance 'decSqRoot'.\n"+
        "Error returned by: \n"+
        " err = err = decSqRoot.IsValid(ePrefix + \"decSqRoot INVALID! \") \n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return decSqRoot, nil
}

// Subtract - Subtracts the incoming Decimal from the current
// Decimal and returns the result as Decimal Type.
//
// The returned Decimal Type contains the same numeric separators
// (decimal separator, thousands separator and currency symbol)
// as those of the current Decimal instance. The numeric separators
// are copied form the current Decimal instance to the returned
// Decimal instance.
func (dec *Decimal) Subtract(d2 Decimal) (Decimal, error) {

  ePrefix := "Decimal.Subtract() "
  var err error

  err = dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on current Decimal instance 'dec'.\n"+
        "Error returned by: \n"+
        " err = dec.bigINum.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d2.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on Decimal instance d2 returned by: \n"+
        " err = d2.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  bINumResult := BigIntMathSubtract{}.SubtractBigIntNums(dec.bigINum, d2.bigINum)

  d3, err := new(Decimal).NewBigIntNum(bINumResult)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " d3, err := new(Decimal).NewBigIntNum(bINumResult)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d3.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Error returned by: \n"+
        " err = d3.SetNumericSeparatorsDto(numSeps)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  err = d3.IsValid(ePrefix)

  if err != nil {

    return Decimal{},
      fmt.Errorf("%v\n"+
        "Validation Error on Decimal instance dec3.\n"+
        "Error returned by: \n"+
        " err = d3.IsValid(ePrefix)\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  return d3, nil
}

// SubtractFromThis - Subtracts the value of the incoming Decimal type
// from the current Decimal type. The updated value is stored and retained
// in the current Decimal instance.
//
// The numeric separators (decimal separator, thousands separator and
// currency symbol) for the current Decimal instance remain unchanged
// and are not modified by this method.
func (dec *Decimal) SubtractFromThis(d2 Decimal) error {

  ePrefix := "Decimal.Subtract()"

  var err error

  err = dec.bigINum.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on current Decimal instance (dec) returned by: \n"+
      " err = dec.bigINum.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  err = d2.IsValid(ePrefix)

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Validation Error on input parameter\n"+
      "Decimal instance (d2). Error returned by: \n"+
      " err = d2.IsValid(ePrefix)\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  numSeps := dec.bigINum.GetNumericSeparatorsDto()

  bINumResult := BigIntMathSubtract{}.SubtractBigIntNums(dec.bigINum, d2.bigINum)

  err = bINumResult.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by bINumResult.SetNumericSeparatorsDto(numSeps) "+
      "Error='%v' \n", err.Error())
  }

  dec.bigINum, err = bINumResult.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = bINumResult.CopyOut()\n\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SubtractFromThisMultiple - Subtracts the value of multiple incoming
// Decimal instances from the current Decimal type. The updated value
// is stored and retained in the current Decimal instance.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SubtractFromThisMultiple(decs ...Decimal) error {

  ePrefix := "SubtractFromThisMultiple"

  err := dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  bINumResult, err := dec.bigINum.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " bINumResult, err := dec.bigINum.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  for _, dx := range decs {

    bINumResult = BigIntMathSubtract{}.SubtractBigIntNums(bINumResult, dx.bigINum)
  }

  dec.bigINum, err = bINumResult.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = bINumResult.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}

// SubtractFromThisArray - Subtracts the values of an Array of incoming
// Decimal instances from the current Decimal type. The updated value
// is stored and retained in the current Decimal instance.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (dec *Decimal) SubtractFromThisArray(decs []Decimal) error {

  ePrefix := "Decimal.SubtractFromThisArray()"

  err := dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " err := dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  bINumResult, err := dec.bigINum.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " bINumResult, err := dec.bigINum.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  for _, dx := range decs {

    bINumResult = BigIntMathSubtract{}.SubtractBigIntNums(bINumResult, dx.bigINum)

  }

  dec.bigINum, err = bINumResult.CopyOut()

  if err != nil {

    return fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      " dec.bigINum, err = bINumResult.CopyOut()\n"+
      "Error= %v\n",
      ePrefix,
      err.Error())
  }

  return nil
}
