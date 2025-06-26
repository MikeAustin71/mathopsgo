package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

/*
	The TimeValOfMoney type contains functions which apply the standard time value
  of money formulas.

	Reference :
		https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/time-value-of-money
		https://www.youtube.com/watch?v=m3azU7gYHc0
		http://www.tvmcalcs.com/index.php/tvm/formulas/tvm_formulas
		https://en.wikipedia.org/wiki/Time_value_of_money

*/

type TimeValOfMoney struct {
	PresentValue BigIntNum
	FutureValue  BigIntNum
	InterestRate BigIntNum
	NPeriods     BigIntNum
}

// LumpSumFVBigIntNum
//
//	Computes and returns the future value of a present value
//	lump sum amount given a risk-free interest rate and a specific number of
//	compounding periods.
//
//	PV = Present Value
//	FV = Future Value
//	 i = risk-free percentage interest rate paid by the investment per period. The
//	     interest rate percentage must be expressed as a decimal fraction. For example,
//	     3% would be passed to this method as 0.03 or percentage / 100
//
//	NthRoot = Number of compounding periods the investment will be held
//
//	Future Value Formula:
//	=====================
//
//	FV = PV x (1 + i)^NthRoot
//
//	Input Parameters:
//	=================
//
//	presentValue             BigIntNum
//	  Present Value (PV) amount.
//
//	interestRate             BigIntNum
//	  Risk Free interest rate percentage (i) per period. The
//	  percentage must be expressed as a decimal value. For example
//	  3% must be submitted as 0.03 or percentage/100. The interest
//	  rate may be a positive or negative numeric value.
//
//	numOfPeriods             BigIntNum
//	  Number of periods (NthRoot) investment will be held with
//	  compounding applied for each period. Must be a positive value.
//
//	futureValueMaxPrecision  uint
//	  The maximum precision required in the resulting Future Value
//	  calculation result. Amounts will be rounded to the
//	  'futureValueMaxPrecision' value, if necessary. 'precision'
//	  identifies the number of digits to the right of the decimal point.
//
//	Return Values
//	==============
//
//	BigIntNum
//	  If the function completes successfully, the Future Value is
//	  returned as a BigIntNum type.
//
//	error
//	  If an error is encountered during the calculation an error
//	  message is returned in this 'error' type. If the calculation
//	  is successful, this value is 'nil'.
func (tvm *TimeValOfMoney) LumpSumFVBigIntNum(
	presentValue, interestRate, numOfPeriods BigIntNum,
	futureValueMaxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"TimeValOfMoney.LumpSumFVBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = presentValue.IsValid(ePrefix.XCpy("presentValue").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = presentValue.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'presentValue' is invalid!\n" +
					"'presentValue' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	err = interestRate.IsValid(ePrefix.XCpy("interestRate").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = interestRate.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'interestRate' is invalid!\n" +
					"'interestRate' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	err = numOfPeriods.IsValid(ePrefix.XCpy("numOfPeriods").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numOfPeriods.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'numOfPeriods' is invalid!\n" +
					"'interestRate' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	numOfPeriodsSignValue, err := numOfPeriods.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numOfPeriodsSignValue, err := numOfPeriods.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numOfPeriodsNumStr, err := numOfPeriods.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numOfPeriodsNumStr, err := numOfPeriods.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if numOfPeriodsSignValue == -1 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("numOfPeriods= '%v'", numOfPeriodsNumStr),
				ErrMessage: "Input parameter 'numOfPeriods' is a negative value.",
			}
	}

	binumOne, err := new(BigIntNum).NewOne(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "binumOne, err := new(BigIntNum).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	onePlusI, err := new(BigIntMathAdd).AddBigIntNums(binumOne, interestRate)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "onePlusI, err := new(BigIntMathAdd).\n" +
					"AddBigIntNums(binumOne, interestRate)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	onePlusINumStr, err := onePlusI.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "onePlusINumStr, err := onePlusI.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	maxPrecision, err := new(BigIntMathPower).MinimumRequiredPrecision(onePlusI, numOfPeriods)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "maxPrecision, err := new(BigIntMathPower).\n" +
					"  MinimumRequiredPrecision(onePlusI, numOfPeriods)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	pwrN, err := new(BigIntMathPower).Pwr(onePlusI, numOfPeriods, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "pwrN, err := new(BigIntMathPower).Pwr(\n" +
					"  onePlusI, numOfPeriods, maxPrecision)",
				ErrContext: fmt.Sprintf("onePlusI= '%v'\n"+
					"numOfPeriods= '%v'\nmaxPrecision='%v'",
					onePlusINumStr, numOfPeriodsNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	futureValue, err := new(BigIntMathMultiply).MultiplyBigIntNums(presentValue, pwrN)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "futureValue, err := new(BigIntMathMultiply).\n" +
					"  MultiplyBigIntNums(presentValue, pwrN)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	futureValuePrecisionUint, err := futureValue.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "futureValuePrecisionUint, err := \n" +
					"futureValue.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if futureValuePrecisionUint > futureValueMaxPrecision {

		err = futureValue.RoundToDecPlace(futureValueMaxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = futureValue.RoundToDecPlace(\n" +
						"  futureValueMaxPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return futureValue, nil
}

// LumpSumPVBigIntNum
//
//	Computes and returns the Present Value (PV) of a
//	Future Value discounted at risk-free discount rate (DR).
//
//	Present Value Formula:
//	======================
//
//	Reference:
//	https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/introduction-to-present-value
//	https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/present-value-2
//	https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/present-value-3
//
//	PV = Present Value
//	FV = Future Value
//
//	i = risk-free discount rate percentage paid on the investment per period. The discount
//	    rate percentage must be submitted as a decimal fraction. For example, 3% must be
//	    passed to this method as 0.03 or percentage/100.
//
//	NthRoot = Number of compounding periods
//
//	 PV = FV / (1 + DR)^NthRoot
func (tvm *TimeValOfMoney) LumpSumPVBigIntNum(
	futureValue, discountRate, numOfPeriods BigIntNum,
	presentValueMaxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"TimeValOfMoney.LumpSumPVBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = futureValue.IsValid(ePrefix.XCpy("futureValue").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = futureValue.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'futureValue' is invalid!\n" +
					"'futureValue' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	futureValueNumStr, err := futureValue.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "futureValueNumStr, err := futureValue.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = discountRate.IsValid(ePrefix.XCpy("discountRate").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = discountRate.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'discountRate' is invalid!\n" +
					"'discountRate' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	err = numOfPeriods.IsValid(ePrefix.XCpy("numOfPeriods").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numOfPeriods.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'numOfPeriods' is invalid!\n" +
					"'numOfPeriods' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	biNumOne, err := new(BigIntNum).NewOne(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "biNumOne, err := new(BigIntNum).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	onePlusDr, err := new(BigIntMathAdd).AddBigIntNums(
		biNumOne, discountRate)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "onePlusDr, err :=	new(BigIntMathAdd).AddBigIntNums(\n" +
					"  biNumOne, discountRate)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	onePlusDrNumStr, err := onePlusDr.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "onePlusDrNumStr, err := onePlusDr.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	maxPrecision, err := new(BigIntMathPower).MinimumRequiredPrecision(
		onePlusDr, numOfPeriods)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "maxPrecision, err := new(BigIntMathPower).MinimumRequiredPrecision(\n" +
					"  onePlusDr, numOfPeriods)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	DRtoPwr, err := new(BigIntMathPower).Pwr(onePlusDr, numOfPeriods, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "DRtoPwr, err := new(BigIntMathPower).Pwr(\n" +
					"  onePlusDr, numOfPeriods, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSeps := new(NumericSeparatorDto).NewUSADefaults()

	presentValue, err :=
		new(BigIntMathDivide).BigIntNumFracQuotient(futureValue, DRtoPwr, numSeps, presentValueMaxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "presentValue, err := new(BigIntMathDivide).\n" +
					"  BigIntNumFracQuotient(futureValue, DRtoPwr, numSeps, presentValueMaxPrecision)",
				ErrContext: fmt.Sprintf("futureValue= '%v'\nonePlusDr= '%v'\nmaxPrecision= '%v'",
					futureValueNumStr, onePlusDrNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return presentValue, nil
}

// SimpleInterestBigIntNum
//
//	Computes future value of an investment using simple
//	the interest formula.
//
//	Simple Interest Formula:
//	========================
//
//	PV = Present Value
//	FV = Future Value
//
//	i = risk-free interest rate percentage paid by the investment
//	    per period. The interest rate percentage must be submitted
//	    as a decimal fraction. For example, 3% must be passed to
//	    this method as 0.03 or percentage/100.
//
//	NthRoot = Number of compounding periods the investment will be held
//
//	FV = PV x ( 1 + ( i x NthRoot) )
//
//	Input Parameters
//	================
//
//	initialInvestment        BigIntNum
//	  The present value amount or Initial Investment.
//		(PV). Must be expressed as a positive value.
//
//	interestRate             BigIntNum
//	  The risk-free percentage interest rate per period (i).
//		The percentage must be expressed as a decimal fraction.
//		For example, 3% must be passed to this method as 0.03 or
//		percentage/ 100. The interest rate can be expressed
//		as either a positive or negative value.
//
//	numOfPeriods             BigIntNum
//	  Number of compounding periods (NthRoot). Must be expressed as
//	  a positive value.
//
//	futureValueMaxPrecision  uint
//	  The maximum precision allowed in the future value amount
//	  resulting from this simple interest calculation. Amounts
//	  will be rounded to the futureValueMaxPrecision' value,
//	  if necessary. 'precision' identifies the number of decimal
//	  places to the right of the decimal point.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  If successful the function will return the future value
//	  resulting from the simple interest calculation as a BigIntNum
//	  type.
//
//	error
//	  If an error is encountered during the calculation an appropriate
//	  error message is returned in this 'error' type. If the
//	  calculation is successful, this value is 'nil'.
func (tvm *TimeValOfMoney) SimpleInterestBigIntNum(
	initialInvestment, interestRate, numOfPeriods BigIntNum,
	futureValueMaxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"TimeValOfMoney.LumpSumFVBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = initialInvestment.IsValid(ePrefix.XCpy("initialInvestment").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = initialInvestment.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'initialInvestment' is invalid!\n" +
					"'initialInvestment' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	initialInvestmentNumStr, err := initialInvestment.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "initialInvestmentNumStr, err := initialInvestment.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = interestRate.IsValid(ePrefix.XCpy("interestRate").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = interestRate.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'interestRate' is invalid!\n" +
					"'interestRate' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	interestRateNumStr, err := interestRate.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "interestRateNumStr, err := interestRate.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = numOfPeriods.IsValid(ePrefix.XCpy("numOfPeriods").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numOfPeriods.IsValid(ePrefix)",
				ErrContext: "Error: Input parameter 'numOfPeriods' is invalid!\n" +
					"'numOfPeriods' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	numOfPeriodsNumStr, err := numOfPeriods.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numOfPeriodsNumStr, err := numOfPeriods.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iTimesN, err := new(BigIntMathMultiply).MultiplyBigIntNums(interestRate, numOfPeriods)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "iTimesN, err := new(BigIntMathMultiply).MultiplyBigIntNums(\n" +
					"  interestRate, numOfPeriods)",
				ErrContext: fmt.Sprintf("interestRate= '%v'\nnumOfPeriods= '%v'",
					interestRateNumStr, numOfPeriodsNumStr),
				ErrMessage: err.Error(),
			}
	}

	bINumOne, err := new(BigIntNum).NewOne(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumOne, err := new(BigIntNum).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	onePlusITimesN, err :=
		new(BigIntMathAdd).AddBigIntNums(bINumOne, iTimesN)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "onePlusITimesN, err := new(BigIntMathAdd).\n" +
					" AddBigIntNums(bINumOne, iTimesN)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	onePlusITimesNNumStr, err := onePlusITimesN.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "onePlusITimesNNumStr, err := onePlusITimesN.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	futureValue, err :=
		new(BigIntMathMultiply).MultiplyBigIntNums(initialInvestment, onePlusITimesN)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "futureValue, err := new(BigIntMathMultiply).\n" +
					" MultiplyBigIntNums(initialInvestment, onePlusITimesN)",
				ErrContext: fmt.Sprintf("initialInvestment= '%v'\nonePlusITimesN= '%v'",
					initialInvestmentNumStr, onePlusITimesNNumStr),
				ErrMessage: err.Error(),
			}
	}

	futureValuePrecisionUint, err := futureValue.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "futureValuePrecisionUint, err := futureValue.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if futureValuePrecisionUint > futureValueMaxPrecision {

		err = futureValue.RoundToDecPlace(futureValueMaxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "futureValuePrecisionUint, err := futureValue.GetPrecisionUint(err = futureValue.\n" +
						"  RoundToDecPlace(futureValueMaxPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return futureValue, nil
}
