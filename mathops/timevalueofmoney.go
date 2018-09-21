package mathops

import "fmt"

/*
	The TimeValOfMoney type contains functions which apply the standard time value
  of money formulas.

 */

type TimeValOfMoney struct {

	PresentValue 	BigIntNum
	FutureValue		BigIntNum
	InterestRate	BigIntNum
	NPeriods			BigIntNum

}


// FutureValueBigIntNum - Computes the future value of a present value amount
// given an interest rate and a number of periods.
//
// PV = Present Value
// FV = Future Value
//  i = interest paid by the investment per period (expressed as a decimal fraction)
//  N = Number of periods the investment will be held
//
// Future Value Formula:
// =====================
//
// FV = PV x (1 + i)^N
//
// Input Parameters:
// =================
//
// presentValue BigIntNum - Present Value (PV) amount. Must be a positive integer
//                          or fractional value.
//
// interestRate	BigIntNum	- Interest Rate (i) per period. Can be  expressed either
// 													a positive or negative value.
//
// numOfPeriods BigIntNum - Number of periods (N) investment will be held. Must
//                          be a positive value.
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
func (tvm TimeValOfMoney) FutureValueBigIntNum(
		presentValue, interestRate, numOfPeriods BigIntNum,
		futureValueMaxPrecision uint ) (BigIntNum, error) {

	ePrefix := "TimeValOfMoney.FutureValueBigInt() "

	if presentValue.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error: Input parameter 'presentValue' is a negative value. " +
				"presentValue='%v' ", presentValue.GetNumStr())
	}

	if numOfPeriods.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error: Input parameter 'numOfPeriods' is a negative value. " +
				"numOfPeriods='%v' ", presentValue.GetNumStr())
	}

	binumOne := BigIntNum{}.NewOne(0)

	onePlusI := BigIntMathAdd{}.AddBigIntNums(binumOne, interestRate)

	maxPrecision:= interestRate.GetPrecisionUint() + uint(500)

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


// SimpleInterestBigIntNum - Computes future value of an investment using simple
// the interest formula.
//
// Simple Interest Formula:
// ========================
//
// PV = Initial Investment (initialInvestment)
//  i = Interest Rate per period (interestRate)
//  N = Number of periods (numOfPeriods)
// FV = Future Value. The result of applying the simple interest formula
//
// FV = PV x ( 1 + ( i x N) )
//
// Input Parameters:
// =================
//
// initialInvestment	BigIntNum - The present value amount or Initial Investment (PV).
//                                Must be expressed as a positive value.
//
// interestRate				BigIntNum - The interest rate per period (i). Can be expressed
//                                as either a positive or negative value.
//
// numOfPeriods				BigIntNum - Number of periods (N). Must be expressed as
//                                a positive value.
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