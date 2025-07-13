package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

type IntAryMathDivide struct {
	Input  IntAryPair
	Result IntAry
}

// Divide
//
//	Divides the parameter 'dividend' by the parameter 'divisor'.
//	The result of this division is a 'quotient' which is returned
//	as an IntAry type.
//
//	The dividend is the number being divided, the divisor is the
//	number by which the dividend is divided, and the quotient is
//	the result of the division.
//
//	Given a ÷ b = c, 'a' is the dividend, 'b' is the divisor and
//	'c' is the quotient.
//
//	Maximum precision of the division result is controlled by the
//	input parameter, 'maxPrecision'.
//
//	If 'maxPrecision' is greater than or equal to zero ('0'),
//	the number of digits to the right of the decimal place will
//	not exceed 'maxPrecision'.
//
//	If 'maxPrecision' is set equal to minus one ('-1'),
//	'maxPrecision' will be automatically set to a maximum of 4,096
//	digits to the right of the decimal point.
//
//	'minPrecision' specifies the minimum precision of the final
//	result. If 'minPrecision' is less than zero, an error will be
//	returned.
//
//	Return Values
//	=============
//
//	IntAry
//	  If the division calculation completes successfully, the
//	  computed quotient is returned as a type IntAry.	The return
//	  value 'IntAry' will contain the numeric separators (decimal
//	  separator, thousands separator and currency symbol) copied
//	  from input parameter 'dividend'.
//
//	error
//	  If the division calculation completes successfully this error
//	  type is set equal to 'nil'. Otherwise, it will contain an
//	  appropriate error message.
func (iaDivide *IntAryMathDivide) Divide(
	dividend *IntAry,
	divisor *IntAry,
	minPrecision int,
	maxPrecision int) (IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathDivide.Divide",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if dividend == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'dividend'",
			}
	}

	if divisor == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'divisor'",
			}
	}

	err = dividend.IsValid(ePrefix.XCpy("Validating dividend").String())

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = dividend.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'dividend' is invalid!\n" +
					"'dividend' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	err = divisor.IsValid(ePrefix.XCpy("Validating divisor").String())

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = divisor.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'divisor' is invalid!\n" +
					"'divisor' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	err = dividend.SetInternalFlags()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = dividend.SetInternalFlags()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = divisor.SetInternalFlags()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = divisor.SetInternalFlags()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if divisor.isZeroValue {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "divisor.isZeroValue = true",
				ErrMessage: "Error: 'divisor' is invalid!\n" +
					"'divisor' has a zero value generating divide by zero error.",
			}
	}

	if maxPrecision == -1 {
		maxPrecision = 4096
	}

	if maxPrecision < -1 {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("maxPrecision = %d", maxPrecision),
				ErrMessage: "Error: 'maxPrecision' is invalid!\n" +
					"'maxPrecision' has a value less than -1.",
			}
	}

	if minPrecision < 0 {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("minPrecision = %d", minPrecision),
				ErrMessage: "Error: 'minPrecision' is invalid!\n" +
					"'minPrecision' has a value less than zero ('0').",
			}
	}

	if minPrecision > maxPrecision {
		minPrecision = maxPrecision
	}

	quotient := new(IntAry).New()

	err = quotient.SetIntAryToZero(0)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = quotient.SetIntAryToZero(0)\n",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if dividend.isZeroValue {
		return quotient, nil
	}

	numSeps, err := dividend.GetNumericSeparatorsDto()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := dividend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating numSeps").String())

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix)",
				ErrContext: "Error: Numeric Separators ('numSeps') are invalid!\n" +
					"'numSeps' extracted from 'dividend' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	trialDividend, err := dividend.CopyOut()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "trialDividend, err := dividend.CopyOut()",
				ErrContext: "ID#1",
				ErrMessage: err.Error(),
			}
	}

	tempDivisor, err := divisor.CopyOut()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "tempDivisor, err := divisor.CopyOut()",
				ErrContext: "ID#1",
				ErrMessage: err.Error(),
			}
	}

	tensCount := new(IntAry).New()

	err = tensCount.SetIntAryToOne(0)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newSignVal := 1

	if trialDividend.signVal != tempDivisor.signVal {

		newSignVal = -1
	}

	if trialDividend.signVal == -1 {

		trialDividend.signVal = 1
	}

	if tempDivisor.signVal == -1 {

		tempDivisor.signVal = 1
	}

	dividendMag, err := trialDividend.GetMagnitudeDigits()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "dividendMag, err := trialDividend.GetMagnitudeDigits()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorMag, err := tempDivisor.GetMagnitudeDigits()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "divisorMag, err := tempDivisor.GetMagnitudeDigits()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	deltaMag := uint(0)

	incrementVal := new(IntAry).New()

	incrementVal, err = tempDivisor.CopyOut()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "incrementVal, err = tempDivisor.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if dividendMag > divisorMag {

		deltaMag = uint(dividendMag - divisorMag)

		err = tensCount.MultiplyByTenToPower(deltaMag)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = tensCount.MultiplyByTenToPower(deltaMag)",
					ErrContext: "dividendMag > divisorMag",
					ErrMessage: err.Error(),
				}
		}

		err = incrementVal.MultiplyThisBy(&tensCount, -1, -1)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = incrementVal.MultiplyThisBy(&tensCount, -1, -1)",
					ErrContext: "dividendMag > divisorMag",
					ErrMessage: err.Error(),
				}
		}

	} else if divisorMag > dividendMag {

		deltaMag = uint(divisorMag - dividendMag)

		err = trialDividend.MultiplyByTenToPower(deltaMag)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = trialDividend.MultiplyByTenToPower(deltaMag)",
					ErrContext: "if divisorMag > dividendMag",
					ErrMessage: err.Error(),
				}
		}

		err = tensCount.DivideByTenToPower(deltaMag)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = tensCount.DivideByTenToPower(deltaMag)",
					ErrContext: "if divisorMag > dividendMag",
					ErrMessage: err.Error(),
				}
		}
	}

	compare := 0

	precisionCutOff := maxPrecision + dividendMag + 1

	// for true
	for {

		if quotient.precision >= precisionCutOff {

			err = quotient.RoundToPrecision(maxPrecision)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.RoundToPrecision(maxPrecision)",
						ErrContext: "if quotient.precision >= precisionCutOff",
						ErrMessage: err.Error(),
					}
			}

			err = quotient.OptimizeIntArrayLen(true)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.OptimizeIntArrayLen(true)",
						ErrContext: "if quotient.precision >= precisionCutOff",
						ErrMessage: err.Error(),
					}
			}

			quotient.signVal = newSignVal

			if quotient.GetPrecision() > maxPrecision {

				err = quotient.RoundToPrecision(maxPrecision)

				if err != nil {

					return IntAry{},
						&FuncReturnError{
							ErrPrefix:  ePrefix.String(),
							ReturnFunc: "err = quotient.RoundToPrecision(maxPrecision)",
							ErrContext: "if quotient.GetPrecision() > maxPrecision",
							ErrMessage: err.Error(),
						}
				}

			}

			if quotient.GetPrecision() < minPrecision {

				err = quotient.SetPrecision(minPrecision, false)

				if err != nil {

					return IntAry{},
						&FuncReturnError{
							ErrPrefix:  ePrefix.String(),
							ReturnFunc: "err = quotient.SetPrecision(minPrecision, false)",
							ErrContext: "if quotient.GetPrecision() < minPrecision",
							ErrMessage: err.Error(),
						}
				}
			}

			err = quotient.SetNumericSeparatorsDto(numSeps)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.SetNumericSeparatorsDto(numSeps)",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			return quotient, nil
		}

		compare, err = incrementVal.CompareAbsoluteValues(&trialDividend)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "compare, err = incrementVal.\n" +
						"CompareAbsoluteValues(&trialDividend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		if compare == 0 {
			// incrementalVal is equal to trialDividend
			err = quotient.AddIntAryToThis(&tensCount)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.AddIntAryToThis(&tensCount)",
						ErrContext: "if compare == 0",
						ErrMessage: err.Error(),
					}
			}

			err = quotient.OptimizeIntArrayLen(true)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.OptimizeIntArrayLen(true)",
						ErrContext: "if compare == 0",
						ErrMessage: err.Error(),
					}
			}

			quotient.signVal = newSignVal

			if quotient.GetPrecision() > maxPrecision {

				err = quotient.RoundToPrecision(maxPrecision)

				if err != nil {

					return IntAry{},
						&FuncReturnError{
							ErrPrefix:  ePrefix.String(),
							ReturnFunc: "err = quotient.RoundToPrecision(maxPrecision)",
							ErrContext: "if compare == 0; if quotient.GetPrecision() > maxPrecision",
							ErrMessage: err.Error(),
						}
				}
			}

			if quotient.GetPrecision() < minPrecision {

				err = quotient.SetPrecision(minPrecision, false)

				if err != nil {

					return IntAry{},
						&FuncReturnError{
							ErrPrefix:  ePrefix.String(),
							ReturnFunc: "err = quotient.SetPrecision(minPrecision, false)",
							ErrContext: "if compare == 0; if quotient.GetPrecision() < minPrecision",
							ErrMessage: err.Error(),
						}
				}
			}

			err = quotient.SetNumericSeparatorsDto(numSeps)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.SetNumericSeparatorsDto(numSeps)",
						ErrContext: "if compare == 0",
						ErrMessage: err.Error(),
					}
			}

			return quotient, nil

		} else if compare == -1 {
			// incrementalVal < trialDividend
			err = quotient.AddIntAryToThis(&tensCount)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = quotient.AddIntAryToThis(&tensCount)",
						ErrContext: "else if compare == -1",
						ErrMessage: err.Error(),
					}
			}

			// Calc Remainder
			err = trialDividend.SubtractFromThis(&incrementVal)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = trialDividend.SubtractFromThis(&incrementVal)",
						ErrContext: "else if compare == -1",
						ErrMessage: err.Error(),
					}
			}

			continue

		} else {
			// Must Be compare == 1
			// incrementalVal > trialDividend

			err = tensCount.DivideByTenToPower(1)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = tensCount.DivideByTenToPower(1)",
						ErrContext: "else compare == 1",
						ErrMessage: err.Error(),
					}
			}

			err = incrementVal.DivideByTenToPower(1)

			if err != nil {

				return IntAry{},
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "err = incrementVal.DivideByTenToPower(1)",
						ErrContext: "else compare == 1",
						ErrMessage: err.Error(),
					}
			}
		}
	}

	//
	//if quotient.GetPrecision() > maxPrecision {
	//	quotient.RoundToPrecision(maxPrecision)
	//}
	//
	//if quotient.GetPrecision() < minPrecision {
	//	quotient.SetPrecision(minPrecision, false)
	//}
	//
	//err = quotient.SetNumericSeparatorsDto(numSeps)
	//
	//if err != nil {
	//	return IntAry{}.New(),
	//		fmt.Errorf(ePrefix+
	//			"Error returned by quotient.SetNumericSeparatorsDto(numSeps). "+
	//			"Error='%v' ", err.Error())
	//}
	//
	//return quotient, nil
}

// DivideByInt64
//
//	Divide the IntAry parameter 'ia' by an int64 divisor and
//	returns the quotient in the pointer to the 'ia' parameter.
//	Consequently, the original value of 'ia' will be overwritten
//	and replaced by the resulting quotient.
//
//	The returned value 'ia' will contain the same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as that of the original 'ia' instance. 'ia' numeric
//	separators will therefore remain unchanged.
//
//	If the quotient has a number of decimal places to the right of
//	the decimal point which is greater than 'maxPrecision', the
//	result is rounded to 'maxPrecision' decimal places.
//
//	If 'maxPrecision' is set equal to -1, 'maxPrecision' is
//	automatically set to 4,096.
//
//	If 'maxPrecision' is less than -1, an error will be returned.
func (iaDivide *IntAryMathDivide) DivideByInt64(
	ia *IntAry, divisor int64, maxPrecision int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathDivide.DivideByInt64()",
		"")

	if err != nil {
		return err
	}

	if divisor == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "divisor == 0",
			ErrMessage: "Error: Input parameter 'divisor' is zero!\n" +
				"Attempted divide by zero.",
		}
	}

	if maxPrecision < -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "maxPrecision < -1",
			ErrMessage: "Error: Input parameter 'maxPrecision' is less than -1 !\n" +
				fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
		}
	}

	if maxPrecision == -1 {
		maxPrecision = 4096
	}

	err = ia.OptimizeIntArrayLen(false)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia.OptimizeIntArrayLen(false)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if ia.isZeroValue {

		iaPrecisionUint, err := ia.GetPrecisionUint()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaPrecisionUint, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		err = ia.SetIntAryToZero(iaPrecisionUint)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia.SetIntAryToZero(iaPrecisionUint)",
				ErrContext: fmt.Sprintf("iaPrecisionUint= '%v'", iaPrecisionUint),
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	dSignVal := 1

	if divisor < 0 {

		dSignVal = -1

		divisor = divisor * -1

	}

	ia.signVal = dSignVal * ia.signVal

	n1 := int64(0)

	n2 := int64(0)

	carry := int64(0)

	iMaxPrecision := int(maxPrecision) + 1

	newAryLen := ia.intAryLen

	intAryLen := ia.intAryLen - ia.precision

	precisionCnt := 0

	for i := 0; i < newAryLen; i++ {

		if i >= intAryLen {

			precisionCnt++
		}

		if i < ia.intAryLen {

			n1 = int64(ia.intAry[i]) + carry

		} else {

			n1 = int64(0) + carry

		}

		n2 = n1 / divisor

		carry = (n1 - (n2 * divisor)) * 10

		if i < ia.intAryLen {

			ia.intAry[i] = uint8(n2)

		} else {

			ia.intAry = append(ia.intAry, uint8(n2))

		}

		if i == newAryLen-1 &&
			carry > 0 && precisionCnt <= iMaxPrecision {

			newAryLen++

		}

	}

	ia.precision = precisionCnt

	ia.intAryLen = newAryLen

	if precisionCnt >= iMaxPrecision {

		iMaxPrecision--

		err = ia.RoundToPrecision(iMaxPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia.RoundToPrecision(iMaxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if ia.intAry[0] == 0 {

		err = ia.SetSignificantDigitIdxs()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia.SetSignificantDigitIdxs()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		ia.intAry = ia.intAry[ia.firstDigitIdx:]

		err = ia.SetIntAryLength()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia.SetIntAryLength()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	return nil
}

// DivideByTenToPower
//
//	Divide input parameter 'ia' of type IntAry by 10 raised to the
//	power of input parameter 'exponent'.
//
//	               'ia'
//	    'ia' =  -----------
//	            10^exponent
//
//	The result, or quotient, is returned via the pointer to input
//	parameter 'ia'. Consequently, the original value of 'ia' will
//	be overwritten and replaced by the resulting quotient.
//
//	The returned value, 'ia', will contain the same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as that of the original 'ia' instance. 'ia' numeric
//	separators will therefore remain unchanged.
func (iaDivide *IntAryMathDivide) DivideByTenToPower(
	ia *IntAry, exponent uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathDivide.DivideByTenToPower()",
		"")

	if err != nil {
		return err
	}

	if exponent == 0 {
		return nil
	}

	ia.precision += int(exponent)

	ia.intAryLen = len(ia.intAry)

	newLen := ia.precision + 1

	if ia.intAryLen < newLen {

		t := make([]uint8, newLen)

		deltaLen := newLen - ia.intAryLen

		for i := 0; i < newLen; i++ {

			if i < deltaLen {
				t[i] = 0
			} else {
				t[i] = ia.intAry[i-deltaLen]
			}

		}

		ia.intAry = make([]uint8, newLen)

		for i := 0; i < newLen; i++ {

			ia.intAry[i] = t[i]
		}

		ia.intAryLen = newLen
	}

	err = ia.OptimizeIntArrayLen(false)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia.OptimizeIntArrayLen(false)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// DivideByTwo
//
//	Receives an input parameter of pointer to an IntAry instance
//	('ia'). 'ia' is then divided by two (2). The result, or
//	quotient, is returned via the pointer to input parameter 'ia'.
//	Consequently, the original value of 'ia' will be overwritten
//	and replaced by the resulting quotient.
//
//	The returned value 'ia' will contain the same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as that of the original 'ia' instance. 'ia' numeric
//	separators will therefore remain unchanged.
func (iaDivide *IntAryMathDivide) DivideByTwo(ia *IntAry) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathDivide.DivideByTwo()",
		"")

	if err != nil {
		return err
	}

	err = ia.OptimizeIntArrayLen(false)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia.OptimizeIntArrayLen(false)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if ia.isZeroValue {

		iaPrecisionUint, err := ia.GetPrecisionUint()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaPrecisionUint, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		err = ia.SetIntAryToZero(iaPrecisionUint)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaPrecisionUint, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	n1 := uint8(0)
	n2 := uint8(0)
	carry := uint8(0)

	for i := 0; i < ia.intAryLen; i++ {

		n1 = ia.intAry[i] + carry
		n2 = n1 / 2
		carry = (n1 - (n2 * 2)) * 10
		ia.intAry[i] = n2

	}

	if carry > 0 {
		ia.intAry = append(ia.intAry, 5)
		ia.intAryLen++
		ia.precision++
	}

	if ia.intAry[0] == 0 {

		err = ia.SetSignificantDigitIdxs()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia.SetSignificantDigitIdxs()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		ia.intAry = ia.intAry[ia.firstDigitIdx:]

		err = ia.SetIntAryLength()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia.SetIntAryLength()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	return nil
}
