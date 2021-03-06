package mathops

import (
	"fmt"
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

	PresentValue 	BigIntNum
	FutureValue		BigIntNum
	InterestRate	BigIntNum
	NPeriods			BigIntNum

}


// LumpSumFVBigIntNum - Computes and returns the future value of a present value
// lump sum amount given a risk free interest rate and a specific number of
// compounding periods.
//
// PV = Present Value
// FV = Future Value
//  i = risk free percentage interest rate paid by the investment per period. The
//      interest rate percentage must be expressed as a decimal fraction. For example,
//      3% would be passed to this method as 0.03 or percentage / 100
//
//  NthRoot = Number of compounding periods the investment will be held
//
// Future Value Formula:
// =====================
//
// FV = PV x (1 + i)^NthRoot
//
// Input Parameters:
// =================
//
// presentValue BigIntNum - Present Value (PV) amount.
//
// interestRate	BigIntNum	- Risk Free interest rate percentage (i) per period. The
//                          percentage must be expressed as a decimal value. For
//                          example 3% must be submitted as 0.03 or percentage/100.
//                          The interest rate may be a positive or negative numeric
//                          value.
//
//
// numOfPeriods BigIntNum - Number of periods (NthRoot) investment will be held with compounding
//                          applied for each period. Must be a positive value.
//
// futureValueMaxPrecision uint	- The maximum precision required in the resulting
// 																Future Value calculation result. Amounts will be
// 																rounded to the 'futureValueMaxPrecision' value,
// 																if necessary. 'precision' identifies the number
// 																of digits to the right of the decimal point.
//
// Return Values:
// ==============
//
// BigIntNum	- If the function completes successfully, the Future Value is returned
//							as a BigIntNum type.
//
// error			- If an error is encountered during the calculation an error message is
//              is returned in this 'error' type. If the calculation is successful, this
//              value is 'nil'.
//
func (tvm TimeValOfMoney) LumpSumFVBigIntNum(
		presentValue, interestRate, numOfPeriods BigIntNum,
		futureValueMaxPrecision uint ) (BigIntNum, error) {

	ePrefix := "TimeValOfMoney.LumpSumFVBigIntNum() "

	if numOfPeriods.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error: Input parameter 'numOfPeriods' is a negative value. " +
				"numOfPeriods='%v' ", presentValue.GetNumStr())
	}

	binumOne := BigIntNum{}.NewOne(0)

	onePlusI := BigIntMathAdd{}.AddBigIntNums(binumOne, interestRate)

	maxPrecision, _ := BigIntMathPower{}.MinimumRequiredPrecision(onePlusI, numOfPeriods)

	pwrN, err := BigIntMathPower{}.Pwr(onePlusI, numOfPeriods, maxPrecision)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMathPower{}.Pwr(onePlusI, numOfPeriods, maxPrecision) " +
				"onePlusI='%v, numOfPeriods='%v', maxPrecision='%v' Error='%v' ",
				onePlusI.GetNumStr(), numOfPeriods.GetNumStr(), maxPrecision, err.Error())
	}

	futureValue := BigIntMathMultiply{}.MultiplyBigIntNums(presentValue, pwrN)

	if futureValue.GetPrecisionUint() > futureValueMaxPrecision {
		futureValue.RoundToDecPlace(futureValueMaxPrecision)
	}

	return futureValue, nil
}


// LumpSumPVBigIntNum - Computes and returns the Present Value (PV) of a
// Future Value discounted at risk free discount rate (DR).
//
// Present Value Formula:
// ======================
//
// Reference:
// https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/introduction-to-present-value
// https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/present-value-2
// https://www.khanacademy.org/economics-finance-domain/core-finance/interest-tutorial/present-value/v/present-value-3
//
// PV = Present Value
// FV = Future Value
//  i = risk free discount rate percentage paid on the investment per period. The discount
//      rate percentage must be submitted as a decimal fraction. For example, 3% must be
//      passed to this method as 0.03 or percentage/100.
//
//  NthRoot = Number of compounding periods
//
// 					PV = FV / (1 + DR)^NthRoot
//
//
func (tvm TimeValOfMoney) LumpSumPVBigIntNum(
	futureValue, discountRate, numOfPeriods BigIntNum,
	presentValueMaxPrecision uint) (BigIntNum, error) {

	ePrefix := "TimeValOfMoney.LumpSumPVBigIntNum() "
	onePlusDr :=
		BigIntMathAdd{}.AddBigIntNums(
			BigIntNum{}.NewOne(0), discountRate)


	maxPrecision, _ := BigIntMathPower{}.MinimumRequiredPrecision(onePlusDr, numOfPeriods)

	DRtoPwr, err := BigIntMathPower{}.Pwr(onePlusDr, numOfPeriods, maxPrecision )

	presentValue, err :=
		BigIntMathDivide{}.BigIntNumFracQuotient(futureValue, DRtoPwr, presentValueMaxPrecision)

	if err != nil {
		return BigIntNum{}.NewZero(0),
		fmt.Errorf(ePrefix +
			"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient( " +
			"futureValue, onePlusDr, maxPrecision) " +
			"futureValue='%v' onePlusDr='%v' maxPrecision='%v' Error='%v'",
			futureValue.GetNumStr(), onePlusDr.GetNumStr(), maxPrecision, err.Error())
	}

	return presentValue, nil
}


// SimpleInterestBigIntNum - Computes future value of an investment using simple
// the interest formula.
//
// Simple Interest Formula:
// ========================
//
// PV = Present Value
// FV = Future Value
//  i = risk free interest rate percentage paid by the investment per period. The
//      interest rate percentage must be submitted as a decimal fraction. For example,
//      3% must be passed to this method as 0.03 or percentage/100.
//
//  NthRoot = Number of compounding periods the investment will be held
//
// FV = PV x ( 1 + ( i x NthRoot) )
//
// Input Parameters:
// =================
//
// initialInvestment	BigIntNum - The present value amount or Initial Investment
// 																(PV). Must be expressed as a positive value.
//
// interestRate				BigIntNum - The risk free percentage interest rate per period (i).
//                                The percentage must be expressed as a decimal fraction.
//                                For example, 3% must be passed to this method as 0.03 or
//                                percentage/ 100. The interest rate can be expressed
//                                as either a positive or negative value.
//
// numOfPeriods				BigIntNum - Number of compounding periods (NthRoot). Must be
// 																expressed as a positive value.
//
// futureValueMaxPrecision uint - The maximum precision allowed in the future
//                                value amount resulting from this simple interest
// 																calculation. Amounts will be rounded to the
//                                'futureValueMaxPrecision' value, if necessary.
//                                'precision' identifies the number of decimal places
//                                to the right of the decimal point.
//
// Return Values:
// ==============
//
// BigIntNum 	- If successful the the function will return the future value resulting
//              from the simple interest calculation as a BigIntNum type.
//
// error			- If an error is encountered during the calculation an error message is
//              is returned in this 'error' type. If the calculation is successful, this
//              value is 'nil'.
//
func (tvm TimeValOfMoney) SimpleInterestBigIntNum(
		initialInvestment, interestRate, numOfPeriods BigIntNum,
		futureValueMaxPrecision uint) (BigIntNum, error) {


	iTimesN := BigIntMathMultiply{}.MultiplyBigIntNums(interestRate, numOfPeriods)

	onePlusITimesN :=
		BigIntMathAdd{}.AddBigIntNums(BigIntNum{}.NewOne(0),	iTimesN)

	futureValue :=
		BigIntMathMultiply{}.MultiplyBigIntNums(initialInvestment, onePlusITimesN)

	if futureValue.GetPrecisionUint() > futureValueMaxPrecision {
		futureValue.RoundToDecPlace(futureValueMaxPrecision)
	}

	return futureValue, nil
}