package examples

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
	"time"
)

func ExampleIntAryPwrByTwos_01(
	baseStr, exponentStr, expectedResult string,
	minResultPrecision, maxResultPrecision int) {

	iaBase, err := mathops.IntAry{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(numStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}

	iaExponent, err := mathops.IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return
	}

	var t1 time.Time
	var t0 time.Time

	t0 = time.Now()
	err = mathops.IntAryMathPower{}.Pwr(
		&iaBase,
		&iaExponent,
		minResultPrecision,
		maxResultPrecision)

	t1 = time.Now()

	if err != nil {
		fmt.Printf("Error returned by mathops.IntAryMathPower{}.Pwr(...). "+
			"Error='%v' \n", err.Error())
		return
	}

	actualResult := iaBase.GetNumStr()
	timeStr := CodeDurationToStr(t1.Sub(t0))

	fmt.Println()
	fmt.Println("IntAryMathPower{}.Pwr()")
	fmt.Println("=======================")
	if expectedResult != actualResult {
		fmt.Println("XXX FAILURE XXX")
	} else {
		fmt.Println("*** SUCCESS ***")
	}

	fmt.Println("=======================")
	fmt.Println("           base: ", baseStr)
	fmt.Println("       exponent: ", exponentStr)
	fmt.Println("  Actual Result: ", actualResult)
	fmt.Println("Expected Result: ", expectedResult)
	fmt.Println("   Elapsed Time: ", timeStr)

}

func ExampleIntAryGetMagnitude_01(numStr string, expectedMagnitude int) {

	iaNum, err := mathops.IntAry{}.NewNumStr(numStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(numStr). "+
			"numStr='%v' Error='%v' \n", numStr, err.Error())
		return
	}

	bInt, err := iaNum.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by iaNum.GetBigInt(). "+
			"Error='%v' \n", err.Error())
		return
	}

	bIMagnitude, err := mathops.BigIntMath{}.GetMagnitude(bInt)

	if err != nil {
		fmt.Printf("Error returned by BigIntMath{}.GetMagnitudeDigits(bInt). "+
			"Error='%v' \n", err.Error())
		return
	}

	magnitude := iaNum.GetMagnitudeDigits()

	iaMagnitude, err := iaNum.GetMagnitude()

	if err != nil {
		fmt.Printf("Error returned by iaNum.GetMagnitude(). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("       Target Number: ", numStr)
	fmt.Println("    Actual Magnitude: ", magnitude)
	fmt.Println("  Expected Magnitude: ", expectedMagnitude)
	fmt.Println(" IntAry.GetMagnitude: ", iaMagnitude)
	fmt.Println("-------------------------------------")
	fmt.Println("BigIntMath Magnitude: ", bIMagnitude.Text(10))

}

func ExampleIntAryMultiplyPower_01(baseStr, exponentStr, expectedResultStr string, minPrecision, maxPrecision int) {

	var t1 time.Time
	var t0 time.Time

	iaBase, err := mathops.IntAry{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := mathops.IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	t0 = time.Now()
	actualResult, err := mathops.IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)
	t1 = time.Now()

	if err != nil {
		fmt.Printf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	fmt.Println()
	fmt.Println("PwrByMultiplication()")
	fmt.Println("=====================")
	if expectedResultStr == actualResultStr {
		fmt.Println("*** SUCCESS ***")
	} else {
		fmt.Println("XXX FAILURE XXX")
	}
	fmt.Println("=====================")

	str := CodeDurationToStr(t1.Sub(t0))

	fmt.Println("     Input Base: ", baseStr)
	fmt.Println(" Input Exponent: ", exponentStr)
	fmt.Println("         Result: ", actualResultStr)
	fmt.Println("Expected Result: ", expectedResultStr)
	fmt.Println("   Elapsed Time: ", str)

}

func ExampleIntAryDivide_01(dividendStr, divisorStr, eQuotient string, minPrecision, maxPrecision int) {

	dividend, err := mathops.IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' \n", dividendStr, err.Error())
		return

	}

	divisor, err := mathops.IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(divisorStr). "+
			"divisorStr='%v' Error='%v' \n", divisorStr, err.Error())
		return

	}

	quotient, err := mathops.IntAryMathDivide{}.Divide(
		&dividend,
		&divisor,
		minPrecision,
		maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by IntAryMathDivide{}.Divide(). "+
			"dividendStr='%v' divisorStr='%v' minPrecision='%v' maxPrecision='%v'"+
			"Error='%v' \n",
			dividendStr, divisorStr, minPrecision, maxPrecision, err.Error())

		return
	}

	actualStr := quotient.GetNumStr()

	if eQuotient != actualStr {
		fmt.Printf("Error: Expected quotient='%v'.  Actual quotient='%v' ",
			eQuotient, actualStr)
		return
	}

	fmt.Println("*** SUCCESS ***")
	fmt.Println("Expected Quotient: ", eQuotient)
	fmt.Println("  Actual Quotient: ", actualStr)
	fmt.Println("         Dividend: ", dividendStr)
	fmt.Println("          Divisor: ", divisorStr)

}
