package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

/*
	StrMathOp
	=========

	The source code repository for strmathop.go is located at:
			https://github.com/MikeAustin71/mathopsgo.git

	The source file strmathop.go is located in directory:
		MikeAustin71/mathopsgo/mathops/strmathop.go

*/

// StrMathOp
//
//	Used to perform string based math operations. The results of these
//	operateions are stored internally in a series of data fields.
//
// Dependencies:
//
//	IntAry
type StrMathOp struct {
	N1       IntAry
	N2       IntAry
	N3       IntAry
	IntMAry  [][]int
	IFinal   IntAry
	Dividend IntAry
	Divisor  IntAry
	Quotient IntAry
	Modulo   IntAry
}

func (sMathOp *StrMathOp) New() StrMathOp {
	iAry := StrMathOp{}
	iAry.N1 = IntAry{}
	iAry.N2 = IntAry{}
	iAry.N3 = IntAry{}
	iAry.IntMAry = make([][]int, 0)
	iAry.IFinal = IntAry{}
	iAry.Dividend = IntAry{}
	iAry.Divisor = IntAry{}
	iAry.Quotient = IntAry{}
	iAry.Modulo = IntAry{}
	return iAry
}

func (sMathOp *StrMathOp) Empty() {

	sMathOp.N1 = IntAry{}
	sMathOp.N2 = IntAry{}
	sMathOp.N3 = IntAry{}
	sMathOp.IntMAry = make([][]int, 0)
	sMathOp.IFinal = IntAry{}
	sMathOp.Dividend = IntAry{}
	sMathOp.Divisor = IntAry{}
	sMathOp.Quotient = IntAry{}
	sMathOp.Modulo = IntAry{}

}

// AddN1N2
//
//	Adds the values in the StrMathOp.N1 and StrMathOp.N2 arrays
//	(Type IntAry) and returns the sum in the StrMathOp.IFinal array
//	(Type IntAry).
//
//	Before calling this method, IntAry's StrMathOp.N1 and StrMathOp.N2
//	must be correctly initialized and configured with the desired base
//	values.
func (sMathOp *StrMathOp) AddN1N2() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.AddN1N2",
		"")

	if err != nil {
		return err
	}

	err = sMathOp.N1.IsValid(ePrefix.XCpy("Validating sMathOp.N1").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N1.IsValid(ePrefix)",
			ErrContext: "Error: 'sMathOp.N1' is invalid!\n" +
				"'sMathOp.N1' FAILED validation tests.",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.N2.IsValid(ePrefix.XCpy("Validating sMathOp.N2").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N2.IsValid(ePrefix)",
			ErrContext: "Error: 'sMathOp.N2' is invalid!\n" +
				"'sMathOp.N2' FAILED validation tests.",
			ErrMessage: err.Error(),
		}
	}

	sMathOp.IFinal, err = sMathOp.N1.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "sMathOp.IFinal, err =\n" +
				"  sMathOp.N1.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.IFinal.AddIntAryToThis(&sMathOp.N2)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.IFinal.AddIntAryToThis(&sMathOp.N2)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// RaiseToPower
//
//	Raises the value of StrMathOp.N1 Int Array to the power specified
//	by the 'power' parameter passed to this method.
//
//	     StrMathOp.N1 ^ power = StrMathOp.IFinal
//
//	Requirements
//	============
//
//	Before calling this method, StrMathOp.N1 Int Array must
//	be properly initialized and configured with the desired base
//	value.
//
//	If input parameter 'power' is less than zero, an error will be
//	returned.
//
//	If input parameter 'internalPrecisionIncrement' is less than 50
//	'internalPrecisionIncrement' will be automatically reset to 50.
//
//	Return Values
//	=============
//
//	Upon completion, calculated values are returned in
//	StrMathOp.IFinal.
func (sMathOp *StrMathOp) RaiseToPower(
	power int,
	internalPrecisionIncrement int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.RaiseToPower",
		"")

	if err != nil {
		return err
	}

	if power < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("power=%d", power),
			ErrMessage: "Error: Input parameter 'power' is out of range!\n" +
				"'power' is less than zero.",
		}
	}

	err = sMathOp.N1.IsValid(ePrefix.XCpy("Validating sMathOp.N1").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N1.IsValid(ePrefix)",
			ErrContext: "sMathOp.N1 is invalid!\n" +
				"You must configure sMathOp.N1 correctly before calling\n" +
				"this method.",
			ErrMessage: err.Error(),
		}
	}

	if internalPrecisionIncrement < 50 {
		internalPrecisionIncrement = 50
	}

	err = sMathOp.N1.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N1.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	sMathOp.IFinal, err = sMathOp.N1.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOp.IFinal, err = sMathOp.N1.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	resultPrecision := sMathOp.N1.GetPrecision()

	sMathOpN1IsZero, err := sMathOp.N1.IsZero()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOpN1IsZero, err := sMathOp.N1.IsZero()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if sMathOpN1IsZero {

		err = sMathOp.IFinal.SetIntAryToZero(uint(resultPrecision))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.IFinal.SetIntAryToZero(uint(resultPrecision))",
				ErrContext: fmt.Sprintf("resultPrecision= '%v'\n"+
					"sMathOp.N1 = 0", uint(resultPrecision)),
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	if power == 0 {

		err = sMathOp.IFinal.SetIntAryToOne(resultPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.IFinal.SetIntAryToOne(resultPrecision)",
				ErrContext: "power= 0",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	if power == 1 {

		return nil
	}

	resultPrecision = resultPrecision * power

	internalPrecision := resultPrecision + internalPrecisionIncrement

	err = sMathOp.IFinal.Pow(power, resultPrecision, internalPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.IFinal.Pow(power, resultPrecision, internalPrecision)",
			ErrContext: fmt.Sprintf("resultPrecision= '%v'\n"+
				"internalPrecisionIncrement= '%v'\n"+
				"internalPrecision= '%v'",
				resultPrecision, internalPrecisionIncrement, internalPrecision),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// MultiplyN1N2
//
//		Mulitplies IntAry's StrMathOp.N1 by StrMathOp.N2 and places the
//		result in StrMathOp.IFinal (Type IntAry).
//
//	 Requirements
//	 ============
//
//	 Both StrMathOp.N1 and StrMathOp.N2 must be properly initialized
//	 and configured with the desired values before calling this method.
func (sMathOp *StrMathOp) MultiplyN1N2() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.MultiplyN1N2",
		"")

	if err != nil {
		return err
	}

	err = sMathOp.N1.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N1.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.N2.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N2.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.N1.IsValid(ePrefix.XCpy("Validating sMathOp.N1").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N1.IsValid(ePrefix)",
			ErrContext: "Errror: 'sMathOp.N1' is invalid!\n" +
				"You must configure sMathOp.N1 correctly before calling\n" +
				"this method.",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.N2.IsValid(ePrefix.XCpy("Validating sMathOp.N2").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N2.IsValid(ePrefix)",
			ErrContext: "Errror: 'sMathOp.N2' is invalid!\n" +
				"You must configure sMathOp.N2 correctly before calling\n" +
				"this method.",
			ErrMessage: err.Error(),
		}
	}

	sMathOp.IFinal, err = sMathOp.N1.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOp.IFinal, err = sMathOp.N1.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.IFinal.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.IFinal.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	n1Precision := sMathOp.N1.GetPrecision()

	n2Precision := sMathOp.N2.GetPrecision()

	greatestPrecision := n1Precision

	if n2Precision > n1Precision {

		greatestPrecision = n2Precision
	}

	maxPrecision := n1Precision + n2Precision

	sMathOpN2NumStr, err := sMathOp.N2.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOpN2NumStr, err := sMathOp.N2.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	sMathOpIFinalNumStr, err := sMathOp.IFinal.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOpIFinalNumStr, err := sMathOp.IFinal.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.IFinal.MultiplyThisBy(&sMathOp.N2, greatestPrecision, maxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = sMathOp.IFinal.MultiplyThisBy(\n" +
				"  &sMathOp.N2, greatestPrecision, maxPrecision)",
			ErrContext: fmt.Sprintf("sMathOp.IFinal= '%v'\n"+
				"sMathOp.N2= '%v'\n"+
				"greatestPrecision= '%v'\n"+
				"maxPrecision= '%v'",
				sMathOpIFinalNumStr, sMathOpN2NumStr, greatestPrecision, maxPrecision),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// Divide
//
//	Divides the Dividend IntAry, StrMathOp.Dividend, by the Divisor
//	IntAry, StrMathOp.Divisor.
//
//	  StrMathOp.Dividend / StrMathOp.Divisor = StrMathOp.Quotient
//
//	The results of this division operation are stored in the StrMathOp
//	member variable, StrMathOp.Quotient.
//
//	Before calling this method, StrMathOp.Dividend and
//	StrMathOp.Divisor must be properly initialized and configured with
//	the desired values.
//
//	If input parameter 'maxPrecision' is less than zero ('0'), an
//	error will be returned.
func (sMathOp *StrMathOp) Divide(maxPrecision int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.Divide",
		"")

	if err != nil {
		return err
	}

	if maxPrecision < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
			ErrMessage: "Error: Input parameter 'maxPrecision' is invalid!\n" +
				"'maxPrecision' must be greater than or equal to '0'.",
		}

	}

	err = sMathOp.Quotient.SetIntAryToZero(0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = err = sMathOp.Quotient.SetIntAryToZero(0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.Modulo.SetIntAryToZero(0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.Modulo.SetIntAryToZero(0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	tensCount := new(IntAry).New()

	err = tensCount.SetIntAryToOne(0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = tensCount.SetIntAryToOne(0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	newSignVal := 1

	divisorSignValue, err := sMathOp.Divisor.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "divisorSignValue, err := sMathOp.Divisor.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	dividendSignValue, err := sMathOp.Dividend.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "dividendSignValue, err := sMathOp.Dividend.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if divisorSignValue != dividendSignValue {

		newSignVal = -1
	}

	if divisorSignValue == -1 {

		err = sMathOp.Divisor.SetSign(1)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.Divisor.SetSign(1)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if dividendSignValue == -1 {

		err = sMathOp.Dividend.SetSign(1)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.Dividend.SetSign(1)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	sMathOp.Divisor.SetIsZeroValue()

	sMathOpDivisorIsZero, err := sMathOp.Divisor.IsZero()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOpDivisorIsZero, err := sMathOp.Divisor.IsZero()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if sMathOpDivisorIsZero {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: 'sMathOp.Divisor' is invalid!\n" +
				"'sMathOp.Divisor' has a zero value. This will\n" +
				"result in a 'divide by zero' error.",
		}
	}

	sMathOp.Dividend.SetIsZeroValue()

	sMathOpDividendIsZero, err := sMathOp.Dividend.IsZero()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOpDividendIsZero, err := sMathOp.Dividend.IsZero()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if sMathOpDividendIsZero {
		return nil
	}

	trialDividend, err := sMathOp.Dividend.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "trialDividend, err := sMathOp.Dividend.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	dividendMag, err := sMathOp.Dividend.GetMagnitudeDigits()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "dividendMag, err := sMathOp.Dividend.GetMagnitudeDigits()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	divisorMag, err := sMathOp.Divisor.GetMagnitudeDigits()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "divisorMag, err := sMathOp.Divisor.GetMagnitudeDigits()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	deltaMag := uint(0)

	incrementVal := new(IntAry).New()

	sMathOpDivisorNumStr, err := sMathOp.Divisor.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOpDivisorNumStr, err := sMathOp.Divisor.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = incrementVal.SetIntAryWithNumStr(sMathOpDivisorNumStr)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = incrementVal.SetIntAryWithNumStr(sMathOpDivisorNumStr)",
			ErrContext: fmt.Sprintf("sMathOpDivisorNumStr= '%v'",
				sMathOpDivisorNumStr),
			ErrMessage: err.Error(),
		}
	}

	var tensCountNumStr, incrementValNumStr, trialDividendNumStr string

	if dividendMag > divisorMag {

		deltaMag = uint(dividendMag - divisorMag)

		err = tensCount.MultiplyByTenToPower(deltaMag)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = tensCount.MultiplyByTenToPower(deltaMag)",
				ErrContext: fmt.Sprintf("deltaMag= '%v'",
					deltaMag),
				ErrMessage: err.Error(),
			}
		}

		tensCountNumStr, err = tensCount.GetNumStr()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.N2.SetInternalFlags()",
				ErrContext: "tensCountNumStr, err = tensCount.GetNumStr()",
				ErrMessage: err.Error(),
			}
		}

		incrementValNumStr, err = incrementVal.GetNumStr()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "incrementValNumStr, err = incrementVal.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		err = incrementVal.MultiplyThisBy(&tensCount, -1, -1)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = incrementVal.MultiplyThisBy(&tensCount, -1, -1)",
				ErrContext: fmt.Sprintf("incrementVal= '%v'\n"+
					"tensCount= '%v'", incrementValNumStr, tensCountNumStr),
				ErrMessage: err.Error(),
			}
		}

	} else if divisorMag > dividendMag {

		deltaMag = uint(divisorMag - dividendMag)

		trialDividendNumStr, err = trialDividend.GetNumStr()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "trialDividendNumStr, err = trialDividend.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		err = trialDividend.MultiplyByTenToPower(deltaMag)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = trialDividend.MultiplyByTenToPower(deltaMag)",
				ErrContext: fmt.Sprintf("deltaMag= '%v'\n"+
					"trialDividend= '%v'", deltaMag, trialDividendNumStr),
				ErrMessage: err.Error(),
			}
		}

		err = tensCount.DivideByTenToPower(deltaMag)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = tensCount.DivideByTenToPower(deltaMag)",
				ErrContext: fmt.Sprintf("deltaMag= '%v'\n"+
					"tensCount= '%v'", deltaMag, tensCountNumStr),
				ErrMessage: err.Error(),
			}
		}
	}

	compareAbsValues := 0

	precisionCutOff := maxPrecision + dividendMag + 1

	var sMathOpQuotientPrecisionInt int

	idx := 0

	for {

		idx++

		sMathOpQuotientPrecisionInt = sMathOp.Quotient.GetPrecision()

		if sMathOpQuotientPrecisionInt == precisionCutOff {

			err = sMathOp.Quotient.SetSign(newSignVal)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = sMathOp.Quotient.SetSign(newSignVal)",
					ErrContext: fmt.Sprintf("newSignVal= '%v'", newSignVal),
					ErrMessage: err.Error(),
				}
			}

			err = sMathOp.Quotient.RoundToPrecision(maxPrecision)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = sMathOp.Quotient.RoundToPrecision(maxPrecision)",
					ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
					ErrMessage: err.Error(),
				}
			}

			return nil
		}

		trialDividendNumStr, err = trialDividend.GetNumStr()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.N2.SetInternalFlags()",
				ErrContext: fmt.Sprintf("Cycle No. = '%v'\n", idx),
				ErrMessage: err.Error(),
			}
		}

		incrementValNumStr, err = incrementVal.GetNumStr()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "incrementValNumStr, err = incrementVal.GetNumStr()",
				ErrContext: fmt.Sprintf("Cycle No. = '%v'\n", idx),
				ErrMessage: err.Error(),
			}
		}

		compareAbsValues, err = incrementVal.CompareAbsoluteValues(&trialDividend)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "compareAbsValues, err = \n" +
					"  incrementVal.CompareAbsoluteValues(&trialDividend)",
				ErrContext: fmt.Sprintf("Cycle No. = '%v'\n"+
					"trialDividend= '%v'\n"+
					"incrementVal= '%v'",
					idx, trialDividendNumStr, incrementValNumStr),
				ErrMessage: err.Error(),
			}
		}

		if compareAbsValues == 0 {
			// incrementalVal is equal to trialDividend

			tensCountNumStr, err = tensCount.GetNumStr()

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "tensCountNumStr, err = tensCount.GetNumStr()",
					ErrContext: fmt.Sprintf("tensCount= '%v'\n"+
						"Cycle No. = '%v'", tensCountNumStr, idx),
					ErrMessage: err.Error(),
				}
			}

			err = sMathOp.Quotient.AddIntAryToThis(&tensCount)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = sMathOp.Quotient.AddIntAryToThis(&tensCount)",
					ErrContext: fmt.Sprintf("tensCount= '%v'\n"+
						"Cycle No. = '%v'", tensCountNumStr, idx),
					ErrMessage: err.Error(),
				}
			}

			err = sMathOp.Quotient.SetSign(newSignVal)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = sMathOp.Quotient.SetSign(newSignVal)",
					ErrContext: fmt.Sprintf("newSignVal= '%v'\n"+
						"Cycle No. = '%v'", newSignVal, idx),
					ErrMessage: err.Error(),
				}
			}

			return nil

		} else if compareAbsValues == -1 {
			// incrementalVal < trialDividend

			tensCountNumStr, err = tensCount.GetNumStr()

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "tensCountNumStr, err = tensCount.GetNumStr()",
					ErrContext: fmt.Sprintf("compareAbsValues == -1 \n"+
						"Cycle No. = '%v'", idx),
					ErrMessage: err.Error(),
				}
			}

			err = sMathOp.Quotient.AddIntAryToThis(&tensCount)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = sMathOp.Quotient.AddIntAryToThis(&tensCount)",
					ErrContext: fmt.Sprintf("compareAbsValues == -1 \n"+
						"tensCount= '%v'\n"+
						"Cycle No. = '%v'", tensCountNumStr, idx),
					ErrMessage: err.Error(),
				}
			}

			incrementValNumStr, err = incrementVal.GetNumStr()

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "incrementValNumStr, err = incrementVal.GetNumStr()",
					ErrContext: fmt.Sprintf("compareAbsValues == -1 \n"+
						"Cycle No. = '%v'", idx),
					ErrMessage: err.Error(),
				}
			}

			trialDividendNumStr, err = trialDividend.GetNumStr()

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "trialDividendNumStr, err = trialDividend.GetNumStr()",
					ErrContext: fmt.Sprintf("compareAbsValues == -1 \n"+
						"Cycle No. = '%v'", idx),
					ErrMessage: err.Error(),
				}
			}

			// Calc Remainder
			err = trialDividend.SubtractFromThis(&incrementVal)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = trialDividend.SubtractFromThis(&incrementVal)",
					ErrContext: fmt.Sprintf("compareAbsValues == -1 \n"+
						"incrementVal= '%v'\n"+
						"trialDividend= '%v'\n"+
						"Cycle No. = '%v'",
						incrementValNumStr, trialDividendNumStr, idx),
					ErrMessage: err.Error(),
				}
			}

			continue

		} else {
			// Must Be compareAbsValues == 1
			// incrementalVal > trialDividend

			tensCountNumStr, err = tensCount.GetNumStr()

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "tensCountNumStr, err = tensCount.GetNumStr()",
					ErrContext: fmt.Sprintf("compareAbsValues == 1 \n"+
						"Cycle No. = '%v'", idx),
					ErrMessage: err.Error(),
				}
			}

			incrementValNumStr, err = incrementVal.GetNumStr()

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "incrementValNumStr, err = incrementVal.GetNumStr()",
					ErrContext: fmt.Sprintf("compareAbsValues == 1 \n"+
						"Cycle No. = '%v'", idx),
					ErrMessage: err.Error(),
				}
			}

			err = tensCount.DivideByTenToPower(1)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = tensCount.DivideByTenToPower(1)",
					ErrContext: fmt.Sprintf("compareAbsValues == 1 \n"+
						"tensCount= '%v'\n"+
						"Cycle No. = '%v'",
						tensCountNumStr, idx),
					ErrMessage: err.Error(),
				}
			}

			err = incrementVal.DivideByTenToPower(1)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = incrementVal.DivideByTenToPower(1)",
					ErrContext: fmt.Sprintf("compareAbsValues == -1 \n"+
						"incrementVal= '%v'\n"+
						"Cycle No. = '%v'",
						incrementValNumStr, idx),
					ErrMessage: err.Error(),
				}
			}
		}
	} // for 'true' loop

	//return nil
}

// DivideDividendByDivisor
//
//	Performs division on two StrMathOp member variables.
//	StrMathOp.Dividend is divided by StrMathOp.Divisor. The results
//	are stored in StrMathOp member variables, StrMathOp.Quotient and
//	StrMathOp.Modulo.
//
//	Before calling this method, StrMathOp.Dividend and
//	StrMathOp.Divisor must be properly initialized and configured with
//	the desired values.
func (sMathOp *StrMathOp) DivideDividendByDivisor() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.DivideDividendByDivisor",
		"")

	if err != nil {
		return err
	}

	sMathOp.N1 = sMathOp.Divisor

	compare := -1

	quotient := 0

	sN := fmt.Sprintf("%v", quotient)

	err = sMathOp.N2.SetIntAryWithNumStr(sN)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.N2.SetIntAryWithNumStr(sN)",
			ErrContext: fmt.Sprintf("sN= '%v'", sN),
			ErrMessage: err.Error(),
		}
	}

	currentCompare := 0

	for compare < 1 {

		currentCompare = compare

		quotient++

		err = sMathOp.N2.IncrementIntegerOne()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.N2.IncrementIntegerOne()",
				ErrContext: fmt.Sprintf("compare= '%v'", compare),
				ErrMessage: err.Error(),
			}
		}

		sMathOp.N3 = sMathOp.IFinal

		err = sMathOp.MultiplyN1N2()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.MultiplyN1N2()",
				ErrContext: fmt.Sprintf("compare= '%v'", currentCompare),
				ErrMessage: err.Error(),
			}
		}

		compare, err = sMathOp.IFinal.CompareSignedValues(&sMathOp.Dividend)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.MultiplyN1N2()",
				ErrContext: fmt.Sprintf("compare= '%v'", currentCompare),
				ErrMessage: err.Error(),
			}
		}
	}

	quotient = quotient - 1

	sN = fmt.Sprintf("%v", quotient)

	err = sMathOp.Quotient.SetIntAryWithNumStr(sN)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.Quotient.SetIntAryWithNumStr(sN)",
			ErrContext: fmt.Sprintf("sN= '%v'", sN),
			ErrMessage: err.Error(),
		}
	}

	// Always false
	//if compare == 0 {
	//	sMathOp.Modulo.SetIntAryToZero(0)
	//	return nil
	//}

	sMathOp.N1 = sMathOp.Dividend

	sMathOp.N2 = sMathOp.N3

	err = sMathOp.SubtractN1N2()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.SubtractN1N2()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	sMathOp.Modulo = sMathOp.IFinal

	return nil
}

// DivideBySubtraction
//
//		Performs a division operation by dividing sMathOp.Dividend by
//	 sMathOp.Divisor.
//
//	 This method differs other StrMathOp division methods in that
//	 this division operation is performed by series of substraction
//	 operations.
//
//		Before calling this method, StrMathOp.Dividend and
//		StrMathOp.Divisor must be properly initialized and configured with
//		the desired values.
func (sMathOp *StrMathOp) DivideBySubtraction() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.DivideBySubtraction",
		"")

	if err != nil {
		return err
	}

	sMathOp.Modulo, err = sMathOp.Dividend.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOp.Modulo, err = sMathOp.Dividend.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.Quotient.SetIntAryToZero(0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.Quotient.SetIntAryToZero(0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	//compare := 1
	compare, err := sMathOp.Modulo.CompareSignedValues(&sMathOp.Divisor)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "compare, err := sMathOp.Modulo.CompareSignedValues(&sMathOp.Divisor)",
			ErrContext: "Error on initial comparison of Dividend and Divisor",
			ErrMessage: err.Error(),
		}
	}

	idx := 0

	for compare >= 0 {

		idx++

		err = sMathOp.Modulo.SubtractFromThis(&sMathOp.Divisor)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.Modulo.SubtractFromThis(&sMathOp.Divisor)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		err = sMathOp.Quotient.IncrementIntegerOne()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = sMathOp.Quotient.IncrementIntegerOne()",
				ErrContext: fmt.Sprintf("Cycle Index= '%d'", idx),
				ErrMessage: err.Error(),
			}
		}

		compare, err = sMathOp.Modulo.CompareSignedValues(&sMathOp.Divisor)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "compare, err = sMathOp.Modulo.CompareSignedValues(&sMathOp.Divisor)",
				ErrContext: fmt.Sprintf("Cycle Index= '%d'", idx),
				ErrMessage: err.Error(),
			}
		}
	}

	return nil
}

//func (sMathOp *StrMathOp) SubtractDivArys() {
//
//}

// SubtractN1N2
//
//	 Subtracts StrMathOp.N2 (IntAry) from StrMathOp.N1 (IntAry).
//
//	 The result is returned in the StrMathOp.IFinal IntAry.
//
//		Before calling this method, StrMathOp.N1 and StrMathOp.N2 must be
//		properly initialized and configured with the desired values.
func (sMathOp *StrMathOp) SubtractN1N2() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"StrMathOp.SubtractN1N2",
		"")

	if err != nil {
		return err
	}

	sMathOp.IFinal, err = sMathOp.N1.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sMathOp.IFinal, err = sMathOp.N1.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sMathOp.IFinal.SubtractFromThis(&sMathOp.N2)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sMathOp.IFinal.SubtractFromThis(&sMathOp.N2)",
			ErrContext: "sMathOp.IFinal contains the value of sMathOp.N1",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
