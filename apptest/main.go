package main

import (
	"../examples"
	"../mathops"
	"fmt"
	"math/big"
	"time"
)

/*
	numStr := "7"
	nthRootStr := "4"
	maxPrecision := uint64(31)
	expectedResult := "1.6265765616977857432112323454938"

	numStr := "4.2"
	nthRootStr := "3"
	maxPrecision := uint64(31)
	expectedResult := "1.6134286460245437640588422364822"

*/


func main() {
	radicand := big.NewInt(-1)
	radicandPrecision := big.NewInt(0)
	nthRoot := big.NewInt(53)
	nthRootPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(30)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1.163014545670929132492822181094"

	TestBigIntPositiveFractionalNthRoot(
		radicand,
		radicandPrecision,
		nthRoot,
		nthRootPrecision,
		maxPrecision,
		expectedResult)
}

func TestBigIntPositiveFractionalNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	fdNr :=  mathops.FixedDecimalNthRoot{}
	timeStart := time.Now()
	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)
	timeEnd := time.Now()
	if err != nil {
		fmt.Printf("Error returned by fdNr.CalculatePositiveFractionalNthRoot(...) " +
			"Error='%v' ", err.Error())
		return
	}

	resultBiNum, err := mathops.BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)
	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("FixedDecimalNthRoot{}.CalculatePositiveFractionalNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", radicand.Text(10))
	fmt.Println("     radicandPrecision: ", radicandPrecision.Text(10))
	fmt.Println("               nthRoot: ", nthRoot.Text(10))
	fmt.Println("      nthRootPrecision: ", nthRootPrecision)
	fmt.Println("     Maximum Precision: ", maxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10) )
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNum.GetNumStr())
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binRadicand, _ := mathops.BigIntNum{}.NewBigIntPrecision(radicand, radicandPrecision)
	binNthRoot, _ := mathops.BigIntNum{}.NewBigIntPrecision(nthRoot, nthRootPrecision)
	uintMaxPrecision := uint(maxPrecision.Uint64())
	timeStart = time.Now()
	binRoot, err := mathops.BigIntMathNthRoot{}.GetNthRoot(binRadicand, binNthRoot, uintMaxPrecision)
	timeEnd = time.Now()
	if err != nil {
		fmt.Printf("Error returned by BigIntMathNthRoot{}.GetNthRoot() " +
			"Error='%v' ", err.Error())
		return
	}

	timeDuration = timeEnd.Sub(timeStart)
	duration = examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println("           BigIntMathNthRoot{}.GetNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", binRadicand.GetNumStr())
	fmt.Println("     radicandPrecision: ", binRadicand.GetPrecision())
	fmt.Println("               nthRoot: ", binNthRoot.GetNumStr())
	fmt.Println("      nthRootPrecision: ", binNthRoot.GetPrecision())
	fmt.Println("     Maximum Precision: ", uintMaxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", binRoot.GetNumStr() )
	fmt.Println("       resultPrecision: ", binRoot.GetPrecision())
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

}


func TestBigIntNegativeIntNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	fdNr :=  mathops.FixedDecimalNthRoot{}
	timeStart := time.Now()
	result, resultPrecision, err :=
		fdNr.CalculateNegativeIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)
	timeEnd := time.Now()
	if err != nil {
		fmt.Printf("Error returned by fdNr.CalculateNegativeIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
		return
	}

	resultBiNum, err := mathops.BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)
	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("FixedDecimalNthRoot{}.CalculateNegativeIntegerNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", radicand.Text(10))
	fmt.Println("     radicandPrecision: ", radicandPrecision.Text(10))
	fmt.Println("               nthRoot: ", nthRoot.Text(10))
	fmt.Println("      nthRootPrecision: ", nthRootPrecision)
	fmt.Println("     Maximum Precision: ", maxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10) )
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNum.GetNumStr())
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()


}

func TestBigIntPositiveIntNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	fdNr :=  mathops.FixedDecimalNthRoot{}
	timeStart := time.Now()
	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)
	timeEnd := time.Now()
	if err != nil {
		fmt.Printf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) " +
			"Error='%v' ", err.Error())
		return
	}

	resultBiNum, err := mathops.BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("Error returned by .BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)
	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("FixedDecimalNthRoot{}.CalculatePositiveIntegerNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", radicand.Text(10))
	fmt.Println("     radicandPrecision: ", radicandPrecision.Text(10))
	fmt.Println("               nthRoot: ", nthRoot.Text(10))
	fmt.Println("      nthRootPrecision: ", nthRootPrecision)
	fmt.Println("     Maximum Precision: ", maxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10) )
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNum.GetNumStr())
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()


}

func TestBigIntNegativeIntPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int,
	expectedResult string) {


	timeStart := time.Now()
	result,
	resultPrecision,
	err := mathops.BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	timeEnd := time.Now()

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	binResult := mathops.BigIntNum{}.NewBigInt(result, uint(resultPrecision.Uint64()))

	timeDuration := timeEnd.Sub(timeStart)
	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("BigIntMathPower{}.BigIntToNegativeIntegerPower() ")
	fmt.Println("============================================================")
	fmt.Println("                  base: ", base.Text(10))
	fmt.Println("         basePrecision: ", basePrecision)
	fmt.Println("              exponent: ", exponent.Text(10))
	fmt.Println("     exponentPrecision: ", basePrecision)
	fmt.Println("     Maximum Precision: ", maxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision)
	fmt.Println("         result NumStr: ", binResult.GetNumStr())
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()


}

func TestBigIntPositiveIntPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int,
	expectedResult string) {

	timeStart := time.Now()
	result,
	resultPrecision,
	err := mathops.BigIntMathPower{}.BigIntToPositiveIntegerPower(
					base,
					basePrecision,
					exponent,
					exponentPrecision,
					maxPrecision)

	timeEnd := time.Now()

	if err != nil {
	fmt.Printf("%v ", err.Error())
	return
	}

	binResult := mathops.BigIntNum{}.NewBigInt(result, uint(resultPrecision.Uint64()))
	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("BigIntMathPower{}.BigIntToPositiveIntegerPower() ")
	fmt.Println("============================================================")
	fmt.Println("                  base: ", base.Text(10))
	fmt.Println("         basePrecision: ", basePrecision.Text(10))
	fmt.Println("              exponent: ", exponent.Text(10))
	fmt.Println("     exponentPrecision: ", exponentPrecision.Text(10))
	fmt.Println("     Maximum Precision: ", maxPrecision.Text(10))
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision)
	fmt.Println("         result NumStr: ", binResult.GetNumStr())
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

}


func TestFixDecNthRootFmtFracDigits(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) {

	nthRootCalc := mathops.FixedDecimalNthRoot{}

	/*
	func (fdNthRoot *FixedDecimalNthRoot) FormatCalculationConstants(
	radicand,
	radicandPrecision,
	intRadicand,
	fracRadicand,
	fracRadicandPrecision,
	nthRoot *big.Int,
	maxPrecision uint64) error
	 */

	timeStart := time.Now()

	err := nthRootCalc.FormatCalculationConstants(
		radicand,
		radicandPrecision,
		nthRoot,
		nthRootPrecision,
		maxPrecision)

	timeEnd := time.Now()


	if err != nil {
		fmt.Printf("Error returned from FixedDecimalNthRoot{}.FormatFractionalDigitsFromRadicand() " +
			"%v", err.Error())
		return
	}

	intRadicand, fracRadicand := big.NewInt(0).QuoRem(radicand, radicandPrecision, nil)

	calcFacs := nthRootCalc.GetInternalCalcFactors()

	fmt.Println()
	fmt.Println()
	fmt.Println("FixedDecimalNthRoot{}.FormatFractionalDigitsFromRadicand() ")
	fmt.Println("============================================================")
	fmt.Println("                  nthRoot: ", nthRoot.Text(10))
	fmt.Println("                 radicand: ", radicand.Text(10))
	fmt.Println("       radicand precision: ", radicandPrecision.Text(10))
	fmt.Println("         radicand integer: ", intRadicand.Text(10))
	fmt.Println("            radicand frac: ", fracRadicand.Text(10))
	fmt.Println("  radicand frac precision: ", radicandPrecision.Text(10))
	fmt.Println("------------------------------------------------------------")
	fmt.Println("   formatted frac integer: ", calcFacs.FmtFracRadicand.Text(10))
	fmt.Println(" formatted frac precision: ", calcFacs.FmtFracRadicandPrecision.Text(10))
	fmt.Println("------------------------------------------------------------")
	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)
	fmt.Println("            Time Duration: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                fracMask1: ", calcFacs.FracMask1.Text(10))
	fmt.Println("                fracMask2: ", calcFacs.FracMask2.Text(10))
	fmt.Println("============================================================")
	fmt.Println()

}

/*
func TestGetNthRoot(radicand, nthRoot mathops.BigIntFixedDecimal, maxPrecision uint64, expectedResult string) {

	nthRootCalc := mathops.FixedDecimalNthRoot{}

	timeStart := time.Now()

	root, err := nthRootCalc.FixedDecNthRoot(radicand, nthRoot, maxPrecision)

	timeEnd := time.Now()

	if err != nil {
		fmt.Printf("Error retrned by nthRootCalc.FixedDecNthRoot(...). " +
			"Error='%v'", err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("FixedDecimalNthRoot.FixedDecNthRoot()")
	fmt.Println("=========================================================")
	fmt.Println("     radicand: ", radicand.GetNumStr())
	fmt.Println("      nthRoot: ", nthRoot.GetNumStr())
	fmt.Println(" maxPrecision: ", maxPrecision)
	fmt.Println("         root: ", root.GetNumStr())
	fmt.Println("expected root: ", expectedResult)
	fmt.Println("=========================================================")
	fmt.Println(" Time Elapsed: ", duration)
	fmt.Println("=========================================================")
	radBINum := mathops.BigIntNum{}.NewBigIntFixedDecimal(radicand)

	nthRootBINum := mathops.BigIntNum{}.NewBigIntFixedDecimal(nthRoot)

	umaxPrecision := uint(maxPrecision)
	timeStart = time.Now()
	root2, err := mathops.BigIntMathNthRoot{}.GetNthRoot(radBINum, nthRootBINum, umaxPrecision)
	timeEnd = time.Now()

	if err != nil {
		fmt.Printf("Error retrned by BigIntMathNthRoot{}.FixedDecNthRoot(...). " +
			"%v", err.Error())
		return
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)

	fmt.Println("BigIntNum Nth Root Calculation")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("         root: ", root2.GetNumStr())
	fmt.Println("---------------------------------------------------------")
	fmt.Println(" Time Elapsed: ", duration)
	fmt.Println("---------------------------------------------------------")



}
*/

func TestFixDecNthRootGetNextFracBundle(
									fracNum,
									fracPrecision,
									nthRoot *big.Int) {

	nthRootCalc := mathops.FixedDecimalNthRoot{}

	/*
	func (fdNthRoot *FixedDecimalNthRoot) FormatCalculationConstants(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) error {

 */

	err := nthRootCalc.FormatCalculationConstants(
		fracNum,
		fracPrecision,
		nthRoot,
		big.NewInt(0),
		big.NewInt(9))



	fmtFracNum, fmtFracPrecision, err :=
		nthRootCalc.FormatFractionalDigitsFromRadicand(
			fracNum,
			fracPrecision)

	if err != nil {
		fmt.Printf("Error returned from FixedDecimalNthRoot{}.FormatFractionalDigitsFromRadicand() " +
			"%v", err.Error())
		return
	}


	fmt.Println()
	fmt.Println("FixedDecimalNthRoot.GetNextFractionalBundleFromRadicand()")
	fmt.Println("=========================================================")
	fmt.Println("   fmtFracNum: ", fmtFracNum.Text(10))
	fmt.Println("fracPrecision: ", fmtFracPrecision.Text(10))
	fmt.Println("      nthRoot: ", nthRoot.Text(10))
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("---------------------------------------------------------")
	bigZero := big.NewInt(0)
	nextBundle := big.NewInt(0)
	cycle := 0
	timeStart := time.Now()
	timeEnd := time.Now()

	for fmtFracPrecision.Cmp(bigZero)==1 {

		timeStart = time.Now()

		nextBundle, fmtFracNum, fmtFracPrecision, err =
		 	nthRootCalc.GetNextFractionalBundleFromRadicand(
		 		fmtFracNum,
				fmtFracPrecision)

		timeEnd = time.Now()

		if err != nil {
			fmt.Printf("Error returned bynthRootCalc.GetNextFractionalBundleFromRadicand(...). " +
				"Error='%v' ", err.Error())
			return
		}

		cycle++
		timeDuration := timeEnd.Sub(timeStart)

		duration := examples.CodeDurationToStr(timeDuration)
		fmt.Println("           Cycle: ", cycle)
		fmt.Println("Next Frac Bundle: ", nextBundle.Text(10))
		fmt.Println("         fracNum: ", fmtFracNum.Text(10))
		fmt.Println("   fracPrecision: ", fmtFracPrecision.Text(10))
		fmt.Println("   Time Duration: ", duration)
		fmt.Println("---------------------------------------------------------")
		fmt.Println()
	}

}

func TestFixDecNthRootNextIntBundle(
	integerNum,
	intTotalDigits,
	nthRoot *big.Int) {

	bigZero := big.NewInt(0)
	nextBundle := big.NewInt(0)
	nextBundleTotDigits := big.NewInt(0)
	var err error
	cycle := 0
	nthRootCalc := mathops.FixedDecimalNthRoot{}

	/*
	func (fdNthRoot *FixedDecimalNthRoot) FormatCalculationConstants(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) error {

	 */


	err = nthRootCalc.FormatCalculationConstants(
		integerNum,
		big.NewInt(0),
		nthRoot,
		big.NewInt(0),
		big.NewInt(9))

	if err != nil {
		fmt.Printf("Error '%v' ", err.Error())
		return
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("FixedDecimalNthRoot{}.GetNextIntegerBundleFromRadicand() ")
	fmt.Println("=============================================")
	fmt.Println("       integerNum: ", integerNum.Text(10))
	fmt.Println("intTotalDigits: ", intTotalDigits.Text(10))
	fmt.Println("       nthRoot: ", nthRoot.Text(10))
	fmt.Println("=============================================")
	fmt.Println("--------------------------------------------------------")

	timeStart := time.Now()
	timeEnd := time.Now()

	for intTotalDigits.Cmp(bigZero)==1{

		timeStart = time.Now()

		nextBundle, integerNum, intTotalDigits, err =
			nthRootCalc.GetNextIntegerBundleFromRadicand(
				integerNum,
			intTotalDigits)

		timeEnd = time.Now()

		if err != nil {
			fmt.Printf("Error returned from FixedDecimalNthRoot{}.GetNextIntegerBundleFromRadicand() " +
				"Error='%v' ", err.Error())
			return
		}

		cycle++

		fmt.Println("                 Cycle: ", cycle)
		fmt.Println("           Next Bundle: ", nextBundle.Text(10))
		fmt.Println("Next Bundle Tot-Digits: ",nextBundleTotDigits.Text(10))
		fmt.Println("               integerNum: ", integerNum.Text(10))
		fmt.Println("     Num of Int Digits: ", intTotalDigits.Text(10))
		fmt.Println("               nthRoot: ", nthRoot.Text(10))
		timeDuration := timeEnd.Sub(timeStart)

		duration := examples.CodeDurationToStr(timeDuration)
		fmt.Println("         Time Duration: ", duration)
		fmt.Println("--------------------------------------------------------")
		fmt.Println()

	}

}

func LogTest001(base, xNum mathops.BigIntNum, maxPrecision uint, expectedLogValue string) {
	// Logarithm test code


	logValue, err := mathops.BigIntMathLogarithms{}.LogBaseOfX(base, xNum, maxPrecision)

	if err != nil {
		fmt.Printf("Error: error returned from BigIntMathLogarithms{}.GetIntDigits(" +
			"base, xNum, maxPrecision) "+
			"base='%v' xNum='%v' Error='%v'",
			base.GetNumStr(), xNum.GetNumStr(), err.Error())

		return
	}

	/*

	checkValue, err := mathops.BigIntMathPower{}.Pwr(base, logValue, 50)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(base, logValue, 50). ")
		return
	}
	errorVariance := mathops.BigIntMathSubtract{}.SubtractBigIntNums(xNum, checkValue)

	errSciNot, err := errorVariance.GetSciNotationStr(20)

	if err != nil {
		fmt.Printf("Error returned by errorVariance.GetSciNotationStr(20). ")
		return
	}
	*/


	fmt.Println("GetNextDecimalDigit() ")
	fmt.Println("          log Value: ", logValue.GetNumStr())
	fmt.Println(" Expected log Value: ", expectedLogValue)
	//fmt.Println("					Check Value: ", checkValue.GetNumStr())
	//fmt.Println("     Error Variance: ", errorVariance.GetNumStr())
	//fmt.Println(" Error Sci-Notation: ", errSciNot)
	fmt.Println("               base: ", base.GetNumStr())
	fmt.Println("               xNum: ", xNum.GetNumStr())
	fmt.Println("       maxPrecision: ", maxPrecision)


}


/*
func main() {

	num := 654
	expectedNumStr := "654.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := mathops.BigIntFixedDecimal{}.NewInt(num, precision)


	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	fmt.Println("Expected Result: ", expectedNumStr)
	fmt.Println("  Actual Result: ", actualNumStr)

	if expectedNumStr != actualNumStr {
		fmt.Printf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}
*/

/*
func main() {
//             1	       2         3
//    123456789012345678901234567890
// 14.220975666072438486085961843571

exponentStr:= "14.220975666072438486085961843571"
exponentBigIntNum, err := mathops.BigIntNum{}.NewNumStr(exponentStr)

if err != nil {
	fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
		"exponentStr='%v' err='%v' ", exponentStr, err.Error())
}

exponentX := mathops.BigIntFixedDecimal{}.New(
	exponentBigIntNum.GetIntegerValue(),
	exponentBigIntNum.GetPrecisionUint())


aValue := uint(14)
nyCycles := uint(20)
//xValue:  1500000.00000000000000000000000069842673210096714730191927902756767566221044250079066725104319314718146994887137115606630049648677768442699335030922790812019670241810296532052911125642686515441075033029443544985829219134240541457261865552972313790517758927060887095319054568787905518252632621321481449139361504466251586088509990238677768615813699307919896788360901752254114404836468448172504734768548195745297400040354523658702608708071939461214809351415323715083214351186140083204203595589902644394289238026
expectedValue:="1500000.000000000000000000000000000"

	// expectedResult := "0.41436922386282680615600815856503"
	//                "0.41436922386282680615600815856503"
	//               "-0.41436922386282680615600815856503"
	// expectedResult := "-0.41436922386282680615600815856503"

	TestEPwrXFromTaylorSeriesBigInt(
		exponentX,
		aValue,
		nyCycles,
		expectedValue)
}
*/

func TestEPwrXFromTaylorSeriesBigInt(
		exponentX mathops.BigIntFixedDecimal,
			a, nCycles uint, expectedResult string) {

	timeStart := time.Now()

	result, err := mathops.BigIntMathLogarithms{}.EPwrXFromTaylorSeriesBigInt(
		exponentX,
		a,
		nCycles)

	timeEnd := time.Now()

	if err != nil {
			fmt.Printf("Error returned by BigIntMathLogarithms{}.EPwrXFromTaylorSeriesBigInt() " +
				"Error='%v' ", err.Error())
	}

	fmt.Println("           nCycles: ", nCycles)
	fmt.Println("           A-Value: ", a)
	fmt.Println("          exponent: ", exponentX.GetNumStr())
	fmt.Println("exponent Precision: ", exponentX.GetPrecision())
	fmt.Println("            xValue: ", result.GetNumStr())
	fmt.Println("    expectedXValue: ", expectedResult)
	fmt.Println("  xValue Precision: ", result.GetPrecision())
	fmt.Println("         StartTime: ", timeStart.String())
	fmt.Println("           EndTime: ", timeEnd.String())

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)
	fmt.Println("   Time Duration: ", duration)
}

func TestEPwrXFromTaylorSeries(exponent, binA mathops.BigIntNum, nCycles int64, expectedXValue string) {

	/*
exponent:  14.220975666072438486085961843571
	  binA:  12
  xValue:  1500000.00000000000000000000000069842673210096714730191927902756767566221044250079066725104319314718146994887137115606630049648677768442699335030922790812019670241810296532052911125642686515441075033029443544985829219134240541457261865552972313790517758927060887095319054568787905518252632621321481449139361504466251586088509990238677768615813699307919896788360901752254114404836468448172504734768548195745297400040354523658702608708071939461214809351415323715083214351186140083204203595589902644394289238026
expectedXValue:  1500000
StartTime:  2018-10-01 23:47:04.3386444 -0500 CDT m=+0.009993801
  EndTime:  2018-10-01 23:47:04.372625 -0500 CDT m=+0.043974401
Time Duration:  33-Milliseconds 980-Microseconds 600-Nanoseconds
	 */

	timeStart := time.Now()

	xValue, err := mathops.BigIntMathLogarithms{}.EPwrXFromTaylorSeries(exponent, binA, nCycles)
	if err != nil {
		fmt.Printf("Error returned by " +
			"BigIntMathLogarithms{}.EPwrXFromTaylorSeries(exponent, binA, nCycles) " +
			"Error='%v' ", err.Error())
		return
	}

	timeEnd := time.Now()

	fmt.Println(" nCycles: ", nCycles)
	fmt.Println("exponent: ", exponent.GetNumStr())
	fmt.Println("  xValue: ", xValue.GetNumStr())
	fmt.Println("expectedXValue: ", expectedXValue)
	fmt.Println("StartTime: ", timeStart.String())
	fmt.Println("  EndTime: ", timeEnd.String())

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)
	fmt.Println("Time Duration: ", duration)

}

/*
	//xNumInt := 1500000
	baseInt := 10
	maxPrecision := uint(15)
	// "6.1760912590556812420812890085306"
	expectedLogValue := "6.17609125905568"

	exponent, err := mathops.BigIntNum{}.NewNumStr(expectedLogValue)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(expectedLogValue). " +
			"Error='%v'", err.Error())
	}

	base := mathops.BigIntNum{}.NewInt(baseInt, 0)

	//xNum := mathops.BigIntNum{}.NewInt(xNumInt, 0)

	PowerTest_01(base, exponent, maxPrecision)

 */


func TestBigIntDivide(
	dividend,
	dividendPrecision,
	divisor,
	divisorPrecision,
	maxPrecision *big.Int,
	expectedResult string) {

	quotient, quotientPrecision, err :=
		mathops.BigIntMathDivide{}.BigIntFracQuotient(
			dividend,
			dividendPrecision,
			divisor,
			divisorPrecision,
			maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by "+
			"BigIntMathDivide{}.BigIntFracQuotient() " +
			"Error='%v'", err.Error())
		return
	}

	biNumResult := mathops.BigIntNum{}.NewBigInt(quotient, uint(quotientPrecision.Uint64()))
	timeStart := time.Now()
	biNumExpectedResult, err := mathops.BigIntNum{}.NewNumStr(expectedResult)
	timeEnd := time.Now()

	if err != nil {
		fmt.Printf("Error returned by "+
			"BigIntNum{}.NewNumStr(expectedResult) " +
			"Error='%v'", err.Error())
		return
	}

	fmt.Println("** Test BigIntMathDivide{}.BigIntFracQuotient() **")
	fmt.Println("--------------------------------------------------")
	fmt.Println("                 dividend: ", dividend.Text(10))
	fmt.Println("        dividendPrecision: ", dividendPrecision.Text(10))
	fmt.Println("                  divisor: ", divisor.Text(10))
	fmt.Println("         divisorPrecision: ", divisorPrecision.Text(10))
	fmt.Println("                 quotient: ", quotient.Text(10))
	fmt.Println("             maxPrecision: ", maxPrecision.Text(10))
	fmt.Println("        quotientPrecision: ", quotientPrecision)
	fmt.Println("          quotient Result: ", biNumResult.GetNumStr())
	fmt.Println("quotient Result Precision: ", biNumResult.GetPrecisionUint())
	fmt.Println("          expected Result: ", expectedResult)
	fmt.Println("   bi Num Expected Result: ", biNumExpectedResult.GetNumStr())
	fmt.Println("bi Num Expected Precision: ", biNumExpectedResult.GetPrecisionUint())
	fmt.Println("====================================================")
	if biNumExpectedResult.GetNumStr() == biNumResult.GetNumStr() {
		fmt.Println("Success! Expected Result Matches Actual Result!")
	} else {
		fmt.Println("Failure! Expected Result DOES NOT Match Actual Result!")
	}
	fmt.Println("====================================================")

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)
	fmt.Println("Time Duration: ", duration)
	fmt.Println("====================================================")
	biNumDividend := mathops.BigIntNum{}.NewBigInt(dividend, uint(dividendPrecision.Uint64()))
	biNumDivisor := mathops.BigIntNum{}.NewBigInt(divisor, uint(divisorPrecision.Uint64()))
	timeStart = time.Now()
	biNumQuotient, err :=mathops.BigIntMathDivide{}.BigIntNumFracQuotient(biNumDividend, biNumDivisor, uint(maxPrecision.Uint64()))
	if err != nil {
		fmt.Printf("Error returned by BigIntMathDivide{}.BigIntNumFracQuotient() " +
			"Error='%v' ", err.Error())
		return
	}
	timeEnd = time.Now()
	timeDuration = timeEnd.Sub(timeStart)
	duration = examples.CodeDurationToStr(timeDuration)
	fmt.Println("-- BigIntMathDivide{}.BigIntNumFracQuotient() --")
	fmt.Println("--------------------------------------------------")
	fmt.Println("dividend: ", biNumDividend.GetNumStr())
	fmt.Println(" divisor: ", biNumDivisor.GetNumStr())
	fmt.Println("quotient: ", biNumQuotient.GetNumStr())
	fmt.Println("====================================================")
	fmt.Println("Time Duration: ", duration)

}

func TestBigIntPwr(
	base *big.Int,
	basePrecision,
	exponent,
	internalMaxPrecision,
	outputMaxPrecision uint,
	expectedResult string) {

	baseToPwr, baseToPwrPrecision := mathops.BigIntMathPower{}.BigIntPwrIteration(
																			base,
																			basePrecision,
																			exponent,
																			internalMaxPrecision,
																			outputMaxPrecision)


	fmt.Println("TestBigIntPwr")
	fmt.Println("                     base: ", base.Text(10))
	fmt.Println("            basePrecision: ", basePrecision)
	fmt.Println("                 exponent: ", exponent)
	fmt.Println("     internalMaxPrecision: ", internalMaxPrecision)
	fmt.Println("       outputMaxPrecision: ", outputMaxPrecision)
	fmt.Println("----------------------------------------------")
	fmt.Println("         Result baseToPwr: ", baseToPwr.Text(10))
	fmt.Println("Result baseToPwrPrecision: ", baseToPwrPrecision)
	bINumResult := mathops.BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)
	fmt.Println("  Result BigIntNum NumStr: ", bINumResult.GetNumStr())
	fmt.Println("          Expected Result: ", expectedResult)

}

func PowerTest_01(base, exponent mathops.BigIntNum, maxPrecision, i uint) {

	powerValue, err := mathops.BigIntMathPower{}.Pwr(base, exponent, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(base, exponent, maxPrecision) " +
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Println("     Index: ", i)
	fmt.Println("      base: ", base.GetNumStr())
	fmt.Println("  exponent: ", exponent.GetNumStr())
	fmt.Println("powerValue: ", powerValue.GetNumStr())
	fmt.Println("--------------------------------------")
	fmt.Println()
}

/*
		xNumInt := 1500000
		baseInt := 4
		maxPrecision := uint(15)
		//expectedLogValue := "10.2582655350227"
		// closer expectedLogValue = "10.258265535022667"
	  expectedLogValue := "10.258265535022667"

		xNumInt := 150
		baseInt := 3
		maxPrecision := uint(6)
	  expectedLogValue := "4.560877"
------------------------------------------------
 log Value:  10.258265535016932
 Expected log Value:  10.258265535022667
               base:  4
               xNum:  1500000

		xNumInt := 1500000
		baseInt := 10
		maxPrecision := uint(15)
		expectedLogValue := "6.17609125905568"


 */

