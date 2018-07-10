package main

import (
	"../mathops"
	"fmt"
	"math/big"
	"time"
	"../examples"
)

func main() {

	// Time: 0-Nanoseconds
	baseStr := "-45.1"
	exponentStr := "1"
	//                      12345678901234567890123456789012345
	expectedNumStr := "-45.1"
	maxPrecision := 5
	minPrecision := 0

	/*
	fmt.Println("     int64 Minute ", int64(time.Minute))
	fmt.Println("     int64 Second ", int64(time.Second))
	fmt.Println("int64 MilliSecond ", int64(time.Millisecond))
	fmt.Println("int64 MicroSecond ", int64(time.Microsecond))
	*/

	//                        12345648901234567890123456789012

	ExampleIntAryMultiplyPower_01(baseStr, exponentStr, expectedNumStr, minPrecision, maxPrecision)
	fmt.Println("------------------------------")
	ExampleBigIntNumPower_01(baseStr, exponentStr, expectedNumStr, uint(maxPrecision))

}

func ExampleIntAryMultiplyPower_01(baseStr, exponentStr, expectedResultStr string, minPrecision, maxPrecision int) {

	var t1 time.Time
	var t0 time.Time


 iaBase, err := mathops.IntAry{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

 iaExponent, err := mathops.IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	t0 = time.Now()
  actualResult, err := mathops.IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)
	t1 = time.Now()

  if err != nil {
		fmt.Printf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

  actualResultStr := actualResult.GetNumStr()

  fmt.Println()
	fmt.Println("=====================")
  fmt.Println("PwrByMultiplication()")
  fmt.Println("=====================")
	if expectedResultStr == actualResultStr {
		fmt.Println("*** SUCCESS ***")
	} else {
		fmt.Println("@@@@ ERROR @@@@")
	}

	str := examples.CodeDurationToStr(t1.Sub(t0))

	fmt.Println("     Input Base: ", baseStr)
	fmt.Println(" Input Exponent: ", exponentStr)
	fmt.Println("         Result: ", actualResultStr)
	fmt.Println("Expected Result: ", expectedResultStr)
	fmt.Println("   Elapsed Time: ", str)

}

func ExampleIntAryDivide_01(dividendStr, divisorStr, eQuotient string, minPrecision, maxPrecision int) {

	dividend, err := mathops.IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(dividendStr). " +
			"dividendStr='%v' Error='%v' \n", dividendStr, err.Error())
		return

	}

	divisor, err := mathops.IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(divisorStr). " +
			"divisorStr='%v' Error='%v' \n", divisorStr, err.Error())
		return

	}

	quotient, err := mathops.IntAryMathDivide{}.Divide(
										&dividend,
											&divisor,
		minPrecision,
													maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by IntAryMathDivide{}.Divide(). " +
			"dividendStr='%v' divisorStr='%v' minPrecision='%v' maxPrecision='%v'" +
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

func ExampleBigIntNumNthRoot_01(
			radicandStr, nthRootStr string,
				maxPrecision uint,
					expectedNumStr string) {


	radicand, err := mathops.BigIntNum{}.NewNumStr(radicandStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(radicandStr). " +
			"radicandStr='%v' Error='%v' \n", radicandStr, err.Error())
		return
	}


	nthRoot, err := mathops.BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(nthRootStr). " +
			"nthRootStr='%v' Error='%v' \n", nthRootStr, err.Error())
		return
	}


	result, err := mathops.BigIntMathNthRoot{}.GetNthRoot(radicand, nthRoot, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by mathNthRootOp.GetNthRoot(radicand, nthRoot, maxPrecision). " +
			"Error='%v' \n", err.Error())
		return
	}

	expectedResult, err := mathops.BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		fmt.Printf("Error returned by mathNthRootOp.BigIntNum{}.NewNumStr(expectedNumStr). " +
			"expectedNumStr='%v' Error='%v' \n", expectedNumStr, err.Error())
		return
	}


	fmt.Println("*** BigIntMathNthRoot ***")
	fmt.Println("Expected Result: ", expectedNumStr)
	fmt.Println("  Actual Result: ", result.GetNumStr())
	fmt.Println("  Max Precision: ", maxPrecision)
	if !expectedResult.Equal(result) {
		fmt.Println("*** ERROR - Actual Result Does NOT Match Expected Result! ***")
	} else {
		fmt.Println("SUCCESS - Expected Result Matched Actual Result!")
		fmt.Println()
	}

	fmt.Println("           Base: ", radicand.GetNumStr())
	fmt.Println("        NthRootInt: ", nthRoot.GetNumStr())

}

func ExampleBigIntNumPower_01(baseStr, exponentStr, expectedStr string, maxPrecision uint) {

	bINumBase, err := mathops.BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}

	bINumExponent, err := mathops.BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return
	}

	var t0 time.Time
	var t1 time.Time

	t0 = time.Now()
	result, err := mathops.BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)
	t1 = time.Now()

	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). " +
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
		return
	}

	str := examples.CodeDurationToStr(t1.Sub(t0))

	fmt.Println("*** BigIntMathPower{}.Pwr() ***")
	fmt.Println("Expected Result: ", expectedStr)
	fmt.Println("  Actual Result: ", result.GetNumStr())
  fmt.Println("   Elapsed Time: ", str)
	if expectedStr != result.GetNumStr() {
		fmt.Println("XXX FAILURE XXX Expected Result NOT EQUAL to Actual Result!")
	} else {
		fmt.Println("*** SUCCESS *** Expected Result Equals Actual Result!")
	}


	fmt.Println("           Base: ", bINumBase.GetNumStr())
	fmt.Println("       Exponent: ", bINumExponent.GetNumStr())

}



func ExampleBundleCount_01(bINumTarget, bINumNthRoot mathops.BigIntNum, expectedBundleCnt string) {

	target, _ := bINumTarget.GetBigInt()
	nthRoot, _ := bINumNthRoot.GetBigInt()

	fmt.Println( "      Original Target: ", target.Text(10))
	newTarget, err :=ExampleBundleCount_02(target, nthRoot)
	if err != nil {
		fmt.Printf("Error returned by ExampleBundleCount_02(target, nthRoot). " +
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println("           New Target: ", newTarget.Text(10))

	bundleCnt, err := ExampleBundleCount_03(newTarget, nthRoot)

	if err != nil {
		fmt.Printf("Error returned by ExampleBundleCount_03(newTarget, nthRoot). " +
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println("         Bundle Count: ", bundleCnt.Text(10))
	fmt.Println("Expected Bundle Count: ", expectedBundleCnt)

}

func ExampleBundlePrecisionCount_03(
	intBundleCnt *big.Int,
		maxPrecision uint )	(totalBundleCnt, precisionBundleCnt *big.Int, err error) {

	err = nil
	totalBundleCnt = big.NewInt(0).Add(intBundleCnt, big.NewInt(int64(maxPrecision)))
	precisionBundleCnt = big.NewInt(0)

	return
}

func ExampleBundleCount_03(target, nthRoot *big.Int) (bundleCnt *big.Int, err error) {
	// Guaranteed: Magnitude of target is Greater Than or Equal to nthRoot
	bundleCnt = big.NewInt(0)
	err = nil

	ePrefix := "ExampleBundleCount_03() "
	magnitude, errx := mathops.BigIntMath{}.GetMagnitude(target)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMath{}.GetMagnitude(target). " +
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
		return bundleCnt, err
	}

	bigOne := big.NewInt(1)

	numOfDigits := big.NewInt(0).Add(magnitude, bigOne)

	quotient, mod := big.NewInt(0).QuoRem(
											numOfDigits,
											nthRoot,
											big.NewInt(0))


	bundleCnt = big.NewInt(0).Set(quotient)

	if mod.Cmp(big.NewInt(0)) == 1 {
		bundleCnt = big.NewInt(0).Add(bundleCnt, bigOne)
	}


	return bundleCnt, nil
}

func ExampleBundleCount_02(target, nthRoot *big.Int) (newTarget *big.Int, err error) {
	// FormatTargetInt
	ePrefix := "ExampleBundleCount_02() "
	newTarget = big.NewInt(0)
	err = nil
	magnitude, errx := mathops.BigIntMath{}.GetMagnitude(target)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMath{}.GetMagnitude(target). " +
			"target='%v' Error='%v' ",
				target.Text(10), err.Error())
		return newTarget, err
	}

	baseTen := big.NewInt(10)
	newTarget = big.NewInt(0).Set(target)
	numOfDigits := big.NewInt(0).Add(magnitude, big.NewInt(1))

	for numOfDigits.Cmp(nthRoot) == -1 {

		newTarget = big.NewInt(0).Mul(newTarget, baseTen )
		numOfDigits = big.NewInt(0).Add(numOfDigits, big.NewInt(1))

	}

	return newTarget, nil
}

func ExampleBigIntMagnitude_01(target *big.Int) {

	magnitude, err := mathops.BigIntMath{}.GetMagnitude(target)

	if err != nil {
		fmt.Printf("Error returned by BigIntMath{}.GetMagnitude(target). " +
			"target='%v' Error='%v' ", target.Text(10), err.Error())
		return
	}

	bigTen := big.NewInt(10)
	val := big.NewInt(0).Exp(bigTen, magnitude, nil)

	fmt.Println("   target= ", target.Text(10))
	fmt.Println("magnitude= ", magnitude.Text(10))
	fmt.Println("Ten Value= ", val.Text(10))

	bigTarget := mathops.BigIntNum{}.NewBigInt(target, 0)
	bigMagnitude := mathops.BigIntNum{}.NewBigInt(magnitude, 0)
	bigTenValue := mathops.BigIntNum{}.NewBigInt(val, 0)
	fmt.Println("========================================")
	fmt.Println("   Target Str: ", bigTarget.FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE))
	fmt.Println("Ten Value Str: ", bigTenValue.FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE))
	fmt.Println("    Magnitude: ", bigMagnitude.GetNumStr())

}

func ExampleBigIntExpTest_01(baseStr, exponentStr, expectedStr string) {

	bINumBase, err := mathops.BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}

	bIBase, err := bINumBase.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by bINumBase.GetBigInt(). " +
			"bINumBase='%v' Error='%v' \n", bINumBase.GetNumStr(), err.Error())
		return
	}


	bINumExponent, err := mathops.BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return
	}

	bIExponent, err := bINumExponent.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by bINumExponent.GetBigInt(). " +
			"bINumExponent='%v' Error='%v' \n", bINumExponent.GetNumStr(), err.Error())
		return
	}
	modM := big.NewInt(0)
	result := big.NewInt(0).Exp(bIBase, bIExponent, modM)

	fmt.Println("big.NewInt(0).Exp(...)")
	fmt.Println("-------------------------------")
	fmt.Println("    Base: ", bIBase.Text(10))
	fmt.Println("Exponent: ", bIExponent.Text(10))
	fmt.Println("-------------------------------")
	fmt.Println("Expected: ", expectedStr)
	fmt.Println("  Result: ", result.Text(10))
	fmt.Println("-------------------------------")
	fmt.Println("    ModM: ", modM.Text(10))
}

