package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
	"math/big"
)

func main() {

	/*
	radicand := "0.027"
	nthRoot := "3"
	maxPrecision := uint(6)
	expectedResult := "0.300000"
	Corrections:
		1. Move fract digits to integer.  27.0
	  2. Precision Adjustment = -1
	 */
/*
	radicand := "0.0275"
	nthRoot := "3"
	expectedStr := "0.301840536839884"
	maxPrecision := uint(15)

	radicand := "0.45"
	nthRoot := "3"
	expectedStr := "0.766309432393553"
	maxPrecision := uint(15)

	radicand := "0.027"
	nthRoot := "3"
	expectedStr := "0.300000"
	maxPrecision := uint(6)

	radicand := "0.0275"
	nthRoot := "3"
	expectedStr := "0.301840536839884"
	maxPrecision := uint(15)



	radicand := "200000.000005"
	nthRoot := "2"
	expectedStr := "447.213595505548"
	maxPrecision := uint(12)
	radicand := "200001.100005"
	nthRoot := "2"
	expectedStr := "447.214825341245"
	maxPrecision := uint(12)

*/

	radicand := "209050307.020509033"
	nthRoot := "2"
	expectedStr := "14458.5720947993"
	maxPrecision := uint(10)

	// Error
	// 14456.8496914335^2 = 209,000,503.000701



	ExampleBigIntNumNthRoot_01(radicand, nthRoot, maxPrecision, expectedStr)
}

/*
func ExampleBiNumNthRootSetupBundles_01(radicand, nthRoot mathops.BigIntNum){

}
*/
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
	fmt.Println("        NthRoot: ", nthRoot.GetNumStr())

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

	result, err := mathops.BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). " +
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
		return
	}

	fmt.Println("*** BigIntMathPower{}.Pwr() ***")
	fmt.Println("Expected Result: ", expectedStr)
	fmt.Println("  Actual Result: ", result.GetNumStr())
	fmt.Println("           Base: ", bINumBase.GetNumStr())
	fmt.Println("       Exponent: ", bINumExponent.GetNumStr())

}