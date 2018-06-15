package mathops

import "errors"

type IntAryMathDivide struct {
	Input  IntAryPair
	Result IntAry
}

// Divide - Divides the parameter 'dividend' by the parameter 'divisor'.
// The result of this division is a 'quotient' which is returned as an
// IntAry type.
//
// The dividend is the number being divided, the divisor is
// the number by which the dividend is divided, and the quotient
// is the result of the division.
//
// Given a ÷ b = c, a is the dividend, b is the divisor and c is
// the quotient.
//
// Maximum precision of the division result is controlled by the input
// parameter, 'maxPrecision'.
//
// If 'maxPrecision' is greater than or equal to zero ('0'),
// the number of digits to the right of the decimal place will
// not exceed 'maxPrecision'.
//
// If 'maxPrecision' is set equal to minus one ('-1'), 'maxPrecision'
// will be automatically set to a maximum of 4,096 digits to the right
// of the decimal point.
//
// 'minPrecision' specifies the minimum precision of the final result.
// If 'minPrecision' is less than zero, it is automatically set to zero.
//
// Returns
// =======
// IntAry - If the division calculation completes successfully, the
//						computed quotient is returned as a type IntAry.
//
// error	- If the division calculation completes successfully this
//						error type is set equal to 'nil'. Otherwise, it will
//						contain an error message.
//
func (iaDivide IntAryMathDivide) Divide(
						dividend, divisor *IntAry,
									minPrecision,  maxPrecision int) (IntAry, error) {

	dividend.SetInternalFlags()
	divisor.SetInternalFlags()

	if divisor.isZeroValue {
		return IntAry{}.New(), errors.New("Error: divide by zero")
	}

	if maxPrecision < -1 {
		return IntAry{}.New(), errors.New("Error: Input parameter 'maxPrecision' is INVALID. 'maxPrecision' is less than -1")
	}

	if maxPrecision == -1 {
		maxPrecision = 4096
	}

	if minPrecision < 0 {
		minPrecision = 0
	}

	if minPrecision > maxPrecision {
		minPrecision = maxPrecision
	}

	quotient := IntAry{}.New()
	quotient.SetIntAryToZero(0)

	if dividend.isZeroValue {
		return quotient, nil
	}

	trialDividend := dividend.CopyOut()

	tempDivisor := divisor.CopyOut()

	tensCount := IntAry{}.New()
	tensCount.SetIntAryToOne(0)

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

	dividendMag := trialDividend.GetMagnitude()
	divisorMag := tempDivisor.GetMagnitude()
	deltaMag := uint(0)
	incrementVal := IntAry{}.New()
	incrementVal = tempDivisor.CopyOut()

	if dividendMag > divisorMag {
		deltaMag = uint(dividendMag - divisorMag)
		tensCount.MultiplyByTenToPower(deltaMag)
		incrementVal.MultiplyThisBy(&tensCount, -1, -1)

	} else if divisorMag > dividendMag {
		deltaMag = uint(divisorMag - dividendMag)
		trialDividend.MultiplyByTenToPower(deltaMag)
		tensCount.DivideByTenToPower(deltaMag)

	}

	compare := 0
	precisionCutOff := maxPrecision + dividendMag + 1

	for true {

		if quotient.precision >= precisionCutOff {
			quotient.RoundToPrecision(maxPrecision)
			quotient.OptimizeIntArrayLen(true)
			quotient.signVal = newSignVal

			if quotient.GetPrecision() > maxPrecision {
				quotient.RoundToPrecision(maxPrecision)
			}

			if quotient.GetPrecision() < minPrecision {
				quotient.SetPrecision(minPrecision, false)
			}

			return quotient, nil
		}

		compare = incrementVal.CompareAbsoluteValues(&trialDividend)

		if compare == 0 {
			// incrementalVal is equal to trialDividend
			quotient.AddToThis(&tensCount)

			quotient.OptimizeIntArrayLen(true)
			quotient.signVal = newSignVal

			if quotient.GetPrecision() > maxPrecision {
				quotient.RoundToPrecision(maxPrecision)
			}

			if quotient.GetPrecision() < minPrecision {
				quotient.SetPrecision(minPrecision, false)
			}

			return quotient, nil

		} else if compare == -1 {
			// incrementalVal < trialDividend
			quotient.AddToThis(&tensCount)

			// Calc Remainder
			trialDividend.SubtractFromThis(&incrementVal)

			continue

		} else {
			// Must Be compare == 1
			// incrementalVal > trialDividend

			tensCount.DivideByTenToPower(1)
			incrementVal.DivideByTenToPower(1)
		}

	}

	if quotient.GetPrecision() > maxPrecision {
		quotient.RoundToPrecision(maxPrecision)
	}

	if quotient.GetPrecision() < minPrecision {
		quotient.SetPrecision(minPrecision, false)
	}

	return quotient, nil
}

// DivideByTwo - Receives an input parameter of pointer to
// an IntAry instance ('ia'). 'ia' is then divided by two (2)
// and the result is returned in the input parameter 'ia'.
//
func (iaDivide IntAryMathDivide) DivideByTwo(ia *IntAry) {

	ia.OptimizeIntArrayLen(false)

	if ia.isZeroValue {

		ia.SetIntAryToZero(ia.precision)

		return
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
		ia.SetSignificantDigitIdxs()
		ia.intAry = ia.intAry[ia.firstDigitIdx:]
		ia.SetIntAryLength()
	}

}
