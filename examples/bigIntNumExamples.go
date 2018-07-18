package examples

import (
	"fmt"
	"math/big"
	"../mathops"
)

func ExampleShiftPrecisionRight(baseNumStr string, shiftPlacesLeft uint, expectedResult string) {

	bigIntNum, err := mathops.BigIntNum{}.NewNumStr(baseNumStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseNumStr). " +
			"baseNumStr='%v' Error='%v' \n",
			baseNumStr, err.Error())
		return
	}

	bigIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bigIntNum.GetNumStr()

	fmt.Println()
	fmt.Println("BigIntNum{}.ShiftPrecisionRight()")
	fmt.Println("---------------------------------")
	if expectedResult == actualResult {
		fmt.Println("*** SUCCESS ***")
	} else {
		fmt.Println("@@@ FAILURE @@@")
	}
	fmt.Println("-------------------------")
	fmt.Println("      Base Number: ", baseNumStr)
	fmt.Println("Shift Places Left: ", shiftPlacesLeft)
	fmt.Println("    Actual Result: ", actualResult)
	fmt.Println("  Expected Result: ", expectedResult)

}

func ExampleShiftPrecisionLeft(baseNumStr string, shiftPlacesLeft uint, expectedResult string) {

	bigIntNum, err := mathops.BigIntNum{}.NewNumStr(baseNumStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseNumStr). " +
			"baseNumStr='%v' Error='%v' \n",
			baseNumStr, err.Error())
		return
	}

	bigIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bigIntNum.GetNumStr()

	fmt.Println()
	fmt.Println("BigIntNum{}.ShiftPrecisionLeft()")
	fmt.Println("-------------------------")
	if expectedResult == actualResult {
		fmt.Println("*** SUCCESS ***")
	} else {
		fmt.Println("@@@ FAILURE @@@")
	}
	fmt.Println("-------------------------")
	fmt.Println("      Base Number: ", baseNumStr)
	fmt.Println("Shift Places Left: ", shiftPlacesLeft)
	fmt.Println("    Actual Result: ", actualResult)
	fmt.Println("  Expected Result: ", expectedResult)
}

func ExampleSetBigFloat_01(bigFloat *big.Float, maxPrecision uint, expectedResult string) {

	bINum := mathops.BigIntNum{}.NewZero(0)

	err := bINum.SetBigFloat(bigFloat, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by bINum.SetBigFloat(bigFloat, maxPrecision). " +
			"bigFloat='%v' maxPrecision='%v' Error='%v' \n",
				bigFloat.Text('f', -1), maxPrecision, err.Error())
		return
	}

	actualNumStr := bINum.GetNumStr()

	fmt.Println()
	fmt.Println("BigIntNum{}.SetBigFloat()")
	fmt.Println("-------------------------")
	if expectedResult == actualNumStr {
		fmt.Println("*** SUCCESS ***")
	} else {
		fmt.Println("@@@ FAILURE @@@")
	}
	fmt.Println("-------------------------")
	fmt.Println("       bigFloat: ", bigFloat.Text('f', -1))
	fmt.Println("   maxPrecision: ", maxPrecision)
	fmt.Println("  Actual Result: ", actualNumStr)
	fmt.Println("Expected Result: ", expectedResult)


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
			"Error returned by BigIntMath{}.GetMagnitudeDigits(target). " +
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
			"Error returned by BigIntMath{}.GetMagnitudeDigits(target). " +
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


