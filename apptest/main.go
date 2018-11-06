package main

import (
	"../examples"
	"../mathops"
	"fmt"
	"math/big"
	"time"
)

func main() {

	fdEuler := mathops.GetEulersNum1050()

	xNum := big.NewInt(1500000)
	xNumPrecision := big.NewInt(0)
	maxInternalPrecision := big.NewInt(500)
	maxPrecision := big.NewInt(30)
	//                            1         2         3
	//                   1234567890123456789012345678901
	expectedValue := "14.220975666072438486085961843571"

	TestBigIntLogBaseOfX(
		fdEuler.GetInteger(),
		fdEuler.GetPrecisionBigInt(),
		xNum,
		xNumPrecision,
		maxInternalPrecision,
		maxPrecision,
		expectedValue)

}

func TestBigIntLogBaseOfX(
	base,
	basePrecision,
	xNum,
	xNumPrecision,
	maxInternalPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestBigIntLogBaseOfX() "
	timeStart := time.Now()
	timeEnd := time.Now()

	timeStart = time.Now()
	logResult, logResultPrecision, err :=
		mathops.BigIntMathLogarithms{}.BigIntLogBaseOfX(
			base,
			basePrecision,
			xNum,
			xNumPrecision,
			maxInternalPrecision,
			maxPrecision)

	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v ", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	status := "Success!!!!! Expected Matches Actual Value"

	if err == nil {
		binLogResult, _ := mathops.BigIntNum{}.NewBigIntPrecision(logResult, logResultPrecision)

		fmt.Println()
		fmt.Println("================================================")
		fmt.Println("    BigIntMathLogarithms{}.BigIntLogBaseOfX() ")
		fmt.Println("================================================")
		fmt.Println("           base: ", base.Text(10))
		fmt.Println("           xNum: ", xNum.Text(10))
		fmt.Println("  xNumPrecision: ", xNumPrecision.Text(10))
		fmt.Println("      logResult: ", logResult.Text(10))
		fmt.Println("   logPrecision: ", logResultPrecision.Text(10))
		fmt.Println("      logNumStr: ", binLogResult.GetNumStr())
		fmt.Println("Expected Result: ", expectedValue)
		fmt.Println(" Execution Time: ", duration)
		fmt.Println("================================================")
		if expectedValue != binLogResult.GetNumStr() {
			status = "FAILURE****** Expected DOES NOT MATCH Actual Value"
		}

		fmt.Println(status)
		fmt.Println("================================================")
		fmt.Println()
	}

	biBase := mathops.BigIntNum{}.NewBigInt(base, 0)
	biXNum, _ := mathops.BigIntNum{}.NewBigIntPrecision(xNum, xNumPrecision)

	uiMaxPrecision := uint(maxPrecision.Uint64())

	timeStart = time.Now()

	biNumResult, err := mathops.BigIntMathLogarithms{}.LogBaseOfX(
		biBase,
		biXNum,
		uiMaxPrecision)
	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v ", err.Error())
		return
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)
	status = "Success!!!!! Expected Matches Actual Value"

	fmt.Println()
	fmt.Println("------------------------------------------------")
	fmt.Println("      BigIntMathLogarithms.LogBaseOfX() ")
	fmt.Println("------------------------------------------------")
	fmt.Println("BigIntNum Log Result: ", biNumResult.GetNumStr())
	fmt.Println("     Expected Result: ", expectedValue)
	fmt.Println("      Execution Time: ", duration)
	fmt.Println("------------------------------------------------")
	if expectedValue != biNumResult.GetNumStr() {
		status = "FAILURE****** Expected DOES NOT MATCH Actual Value"
	}
	fmt.Println(status)
	fmt.Println("------------------------------------------------")
	fmt.Println()
}



/*
func main() {
	xNum := big.NewInt(2)
	xNumPrecision := big.NewInt(0)
	m := big.NewInt(175)
	s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
	agMeanMaxInternalPrecision := big.NewInt(2000)
	agMeanMaxOutputPrecision := big.NewInt(150)
	maxPiDividePrecision := big.NewInt(2000)
	maxFinalOutputPrecision := big.NewInt(99)
	expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418688"

	TestNatLogOfXArithmeticGeometricMean(
		xNum,
		xNumPrecision,
		m,
		s4DivPrecisionMaxInternalPrecision,
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		maxPiDividePrecision,
		maxFinalOutputPrecision,
		expectedValue)

}


func main() {
	ePrefix := "main() "
	m:= big.NewInt(18)
	xNum := big.NewInt(2)
	xNumPrecision := big.NewInt(0)
	twoToM := big.NewInt(0).Exp(big.NewInt(2), m, nil)

	s, sPrecision, err :=


/*
		mathops.BigIntMathMultiply{}.BigIntMultiply(
			twoToM,
			big.NewInt(0),
			xNum,
			xNumPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)

	fourDivS, fourDivSPrecision, err :=
		mathops.BigIntMathDivide{}.BigIntFracQuotient(
			big.NewInt(4),
			big.NewInt(0),
			s,
			sPrecision,
			s4DivPrecisionMaxInternalPrecision)

	agMeanMaxInternalPrecision := big.NewInt(15000)
	agMeanMaxOutputPrecision := big.NewInt(200)

	fmt.Println("m: ", m.Text(10))
	fmt.Println()
	TestArithmeticGeometricMean(
		fourDivS,
		fourDivSPrecision,
		big.NewInt(1),
		big.NewInt(0),
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		"Unknown")
}

func main() {
	ePrefix := "main() "

	num1 := big.NewInt(762939453125)
	num1Precision := big.NewInt(17)

	num2 := big.NewInt(1)
	num2Precision := big.NewInt(0)

	product, productPrecision, err :=
		mathops.BigIntMathMultiply{}.BigIntMultiply(num1, num1Precision, num2, num2Precision)

	if err !=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binProduct, err := mathops.BigIntNum{}.NewBigIntPrecision(product, productPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "binProduct - %v", err.Error())
		return
	}

	fmt.Println("Product: ", binProduct.GetNumStr())
	maxInternalPrecision := big.NewInt(6000)


	fdNRoot := mathops.FixedDecimalNthRoot{}
	g, gPrecision, err :=
		fdNRoot.CalculatePositiveIntegerNthRoot(
			product,
			productPrecision,
			big.NewInt(2),
			big.NewInt(0),
			maxInternalPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "nthRoot - %v", err.Error())
		return
	}

	binG, err := mathops.BigIntNum{}.NewBigIntPrecision(g, gPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "binG - %v", err.Error())
		return
	}

	fmt.Println("binG: ", binG.GetNumStr())

}



func main() {
	xNum := big.NewInt(2)
	xNumPrecision := big.NewInt(0)
	m := big.NewInt(193)
	s4DivPrecisionMaxInternalPrecision := big.NewInt(8000)
	agMeanMaxInternalPrecision := big.NewInt(6000)
	agMeanMaxOutputPrecision := big.NewInt(800)
	maxPiDividePrecision := big.NewInt(15000)
	maxFinalOutputPrecision := big.NewInt(99)
	expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418688"

	TestNatLogOfXArithmeticGeometricMean(
		xNum,
		xNumPrecision,
		m,
		s4DivPrecisionMaxInternalPrecision,
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		maxPiDividePrecision,
		maxFinalOutputPrecision,
		expectedValue)

}

*/

func TestArithmeticGeometricMean(
	aNum,
	aNumPrecision,
	gNum,
	gNumPrecision,
	maxInternalPrecision,
	targetPrecision *big.Int,
	expectedValue string) {

	timeStart := time.Now()
	timeEnd := time.Now()
	ePrefix := "TestArithmeticGeometricMean() "

	timeStart = time.Now()

	agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err :=
		mathops.BigIntMath{}.ArithmeticGeometricMean(
			aNum,
			aNumPrecision,
			gNum,
			gNumPrecision,
			maxInternalPrecision,
			targetPrecision)

	timeEnd = time.Now()

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binANum, err := mathops.BigIntNum{}.NewBigIntPrecision(aNum, aNumPrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binGNum, err := mathops.BigIntNum{}.NewBigIntPrecision(gNum, gNumPrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binAGMean, err := mathops.BigIntNum{}.NewBigIntPrecision(agMean, agMeanPrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binGValue, err := mathops.BigIntNum{}.NewBigIntPrecision(gValue, gValuePrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("               BigIntMath{}.ArithmeticGeometricMean() ")
	fmt.Println("======================================================================")
	fmt.Println("                aNum: ", aNum.Text(10))
	fmt.Println("       aNumPrecision: ", aNumPrecision.Text(10))
	fmt.Println("           a Num Str: ", binANum.GetNumStr())
	fmt.Println("                gNum: ", gNum.Text(10))
	fmt.Println("       gNumPrecision: ", gNumPrecision.Text(10))
	fmt.Println("           g Num Str: ", binGNum.GetNumStr())
	fmt.Println("maxInternalPrecision: ", maxInternalPrecision.Text(10))
	fmt.Println("     targetPrecision: ", targetPrecision.Text(10))
	fmt.Println("======================================================================")
	fmt.Println("              agMean: ", agMean.Text(10))
	fmt.Println("     agMeanPrecision: ", agMeanPrecision.Text(10))
	fmt.Println("       agMean NumStr: ", binAGMean.GetNumStr())
	fmt.Println("      expected value: ", expectedValue)
	status := "FAILURE - Actual DOES NOT MATCH Expected Value!!"
	if binAGMean.GetNumStr() == expectedValue{
		status = "SUCCESS - Actual Matches Expected Value!"
	}

	fmt.Println("              Status: ", status)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("              gValue: ", gValue.Text(10))
	fmt.Println("     gValuePrecision: ", gValuePrecision.Text(10))
	fmt.Println("       gValue NumStr: ", binGValue.GetNumStr())
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("              Cycles: ", cycles)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("      Execution Time: ", duration)
	fmt.Println("======================================================================")
	fmt.Println()

}


func TestNatLogOfXArithmeticGeometricMean(
	xNum,
	xNumPrecision,
	m,
	s4DivPrecisionMaxInternalPrecision,
	agMeanMaxInternalPrecision,
	agMeanMaxOutputPrecision,
	piDivideMaxPrecision,
	maxFinalResultPrecision *big.Int,
	expectedValue string) {

	/*

		xNum := big.NewInt(2)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(144)
		s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
		agMeanMaxInternalPrecision := big.NewInt(2000)
		agMeanMaxOutputPrecision := big.NewInt(800)
		maxPiDividePrecision := big.NewInt(2000)
		maxFinalOutputPrecision := big.NewInt(99)
		expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"


		xNum := big.NewInt(2)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(1029)
		s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
		agMeanMaxInternalPrecision := big.NewInt(2000)
		agMeanMaxOutputPrecision := big.NewInt(800)
		maxPiDividePrecision := big.NewInt(2000)
		maxFinalOutputPrecision := big.NewInt(99)
		expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"

		xNum=14
		expectedValue := "2.6390573296152586145225848649014"
		 xNum=256
		expectedValue := "5.5451774444795624753378569716654"


		xNum := big.NewInt(2)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(512)
		agMeanMaxInternalPrecision := big.NewInt(8000)
		agMeanMaxOutputPrecision := big.NewInt(500)
		maxPiDividePrecision := big.NewInt(150)
		maxFinalOutputPrecision := big.NewInt(99)
		expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"



		xNum := big.NewInt(14)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(256)
		maxInternalPrecision := big.NewInt(1000)
		maxPrecision := big.NewInt(31)
		expectedValue := "2.6390573296152586145225848649014"

	 */



	timeStart := time.Now()
	timeEnd := time.Now()
	ePrefix := "TestNatLogOfXArithmeticGeometricMean() "

	timeStart = time.Now()

	lnOfX, lnOfXPrecision, err :=
		mathops.BigIntMathLogarithms{}.NatLogOfXArithmeticGeometricMean(
			xNum,
			xNumPrecision,
			m,
			s4DivPrecisionMaxInternalPrecision,
			agMeanMaxInternalPrecision,
			agMeanMaxOutputPrecision,
			piDivideMaxPrecision,
			maxFinalResultPrecision)

		timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	bINumXNum, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(xNum, xNumPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "bINumXNum Error= %v", err.Error())
		return
	}

	bINumlnOfX, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(lnOfX, lnOfXPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "bINumlnOfX Error= %v", err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("        BigIntMathLogarithms{}.NatLogOfXArithmeticGeometricMean() ")
	fmt.Println("======================================================================")
	fmt.Println("                              xNum: ", xNum.Text(10))
	fmt.Println("                     xNumPrecision: ", xNumPrecision.Text(10))
	fmt.Println("                       xNum NumStr: ", bINumXNum.GetNumStr())
	fmt.Println("                                 m: ", m.Text(10))
	fmt.Println("s4DivPrecisionMaxInternalPrecision: ", s4DivPrecisionMaxInternalPrecision.Text(10))
	fmt.Println("        agMeanMaxInternalPrecision: ", agMeanMaxInternalPrecision.Text(10))
	fmt.Println("          agMeanMaxOutputPrecision: ", agMeanMaxOutputPrecision.Text(10))
	fmt.Println("              piDivideMaxPrecision: ", piDivideMaxPrecision.Text(10))
	fmt.Println("           maxFinalResultPrecision: ", maxFinalResultPrecision.Text(10))
	fmt.Println("========================== Result ====================================")
	fmt.Println("                             ln(x): ", lnOfX.Text(10))
	fmt.Println("                   ln(x) Precision: ", lnOfXPrecision.Text(10))
	fmt.Println("                      ln(x) NumStr: ", bINumlnOfX.GetNumStr())
	fmt.Println("              ln(x) Expected Value: ", expectedValue)

	status := "FAILURE!!! Expected Value DOES NOT MATCH Actual Value"
	if expectedValue == bINumlnOfX.GetNumStr() {
		status = "SUCCESS**** Expected Value MATCHES Actual Value"
	}

	fmt.Println("                    Status: ", status)
	fmt.Println("======================================================================")
	fmt.Println("      Execution Time: ", duration)

}

func TestGetBigIntDecimalDigits(
	base,
	basePrecision,
	residualXNum,
	residualXNumPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestGetBigIntDigits() "

	bigOne := big.NewInt(1)
	bigZero := big.NewInt(0)

	inverseBase, inverseBasePrecision, errX :=
		mathops.BigIntMathDivide{}.BigIntFracQuotient(
			bigOne,
			bigZero,
			base,
			basePrecision,
			maxPrecision)

	if errX != nil {
		fmt.Printf(ePrefix + "%v", errX.Error())
		return
	}


	timeStart := time.Now()

	digit, newResidualXNum, newResidualXNumPrecision, err :=
	 	mathops.BigIntMathLogarithms{}.BigIntGetNextDecimalDigit(
	 		base,
	 		basePrecision,
	 		inverseBase,
	 		inverseBasePrecision,
	 		residualXNum,
	 		residualXNumPrecision,
	 		maxPrecision)


	timeEnd := time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	if err == nil {
		numStrXNumVerify, _ :=
			mathops.BigIntNum{}.NewBigIntPrecision(newResidualXNum, newResidualXNumPrecision)

		fmt.Println()
		fmt.Println("==================================================")
		fmt.Println("BigIntMathLogarithms{}.BigIntGetNextDecimalDigit()")
		fmt.Println("==================================================")
		fmt.Println("                    base: ", base.Text(10))
		fmt.Println("            residualXNum: ", residualXNum.Text(10))
		fmt.Println("   residualXNumPrecision: ", residualXNumPrecision.Text(10))
		fmt.Println("            maxPrecision: ", maxPrecision.Text(10))
		fmt.Println("=================== Results ======================")
		fmt.Println("                  digit: ", digit.Text(10))
		fmt.Println("         newResidualXNum: ", newResidualXNum.Text(10))
		fmt.Println("newResidualXNumPrecision: ", newResidualXNumPrecision.Text(10))
		fmt.Println("  newResidualXNum NumStr: ", numStrXNumVerify.GetNumStr())
		fmt.Println("           expectedValue: ", expectedValue)
		fmt.Println("          Execution Time: ", duration)
		fmt.Println("===================================================")
		fmt.Println()

	}

	binBase := mathops.BigIntNum{}.NewBigInt(base, 0)

	binXNum, _ := mathops.BigIntNum{}.NewBigIntPrecision(residualXNum, residualXNumPrecision)

	uintMaxPrecision := uint(maxPrecision.Uint64())

	timeStart = time.Now()
	binDigits, binNewResidualXNum, err :=
		mathops.BigIntMathLogarithms{}.GetNextDecimalDigit(binBase, binXNum, uintMaxPrecision)
	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println("BigIntMathLogarithms{}.BigIntGetNextDecimalDigit()")
	fmt.Println("--------------------------------------------------")
	fmt.Println("                    base: ", binBase.GetNumStr())
	fmt.Println("                    xNum: ", binXNum.GetNumStr())
	fmt.Println("            maxPrecision: ", uintMaxPrecision)
	fmt.Println("------------------ Results -----------------------")
	fmt.Println("                  digit: ", binDigits.GetNumStr())
	fmt.Println("         newResidualXNum: ", binNewResidualXNum.GetNumStr())
	fmt.Println("           expectedValue: ", expectedValue)
	fmt.Println("          Execution Time: ", duration)
	fmt.Println("------------------------------------------------")

}

func TestGetBigIntDigits(
	base,
	basePrecision,
	xNum,
	xNumPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestGetBigIntDigits() "

	bigOne := big.NewInt(1)
	bigZero := big.NewInt(0)

	inverseBase, inverseBasePrecision, errX :=
		mathops.BigIntMathDivide{}.BigIntFracQuotient(
			bigOne,
			bigZero,
			base,
			basePrecision,
			maxPrecision)

	if errX != nil {
		fmt.Printf(ePrefix + "%v", errX.Error())
		return
	}

	timeStart := time.Now()
	digits, newResidualXNum, newResidualXNumPrecision, err :=
		mathops.BigIntMathLogarithms{}.BigIntGetIntDigits(
			base,
			basePrecision,
			inverseBase,
			inverseBasePrecision,
			xNum,
			xNumPrecision,
			maxPrecision)

	timeEnd := time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	if err == nil {
		numStrXNumVerify, _ :=
			mathops.BigIntNum{}.NewBigIntPrecision(newResidualXNum, newResidualXNumPrecision)

		fmt.Println()
		fmt.Println("===============================================")
		fmt.Println("  BigIntMathLogarithms{}.BigIntGetIntDigits()")
		fmt.Println("===============================================")
		fmt.Println("                    base: ", base.Text(10))
		fmt.Println("                    xNum: ", xNum.Text(10))
		fmt.Println("           xNumPrecision: ", xNumPrecision.Text(10))
		fmt.Println("            maxPrecision: ", maxPrecision.Text(10))
		fmt.Println("================= Results =====================")
		fmt.Println("                  digits: ", digits.Text(10))
		fmt.Println("         newResidualXNum: ", newResidualXNum.Text(10))
		fmt.Println("newResidualXNumPrecision: ", newResidualXNumPrecision.Text(10))
		fmt.Println("  newResidualXNum NumStr: ", numStrXNumVerify.GetNumStr())
		fmt.Println("           expectedValue: ", expectedValue)
		fmt.Println("          Execution Time: ", duration)
		fmt.Println("================================================")
		fmt.Println()

	}

	binBase := mathops.BigIntNum{}.NewBigInt(base, 0)

	binXNum, _ := mathops.BigIntNum{}.NewBigIntPrecision(xNum, xNumPrecision)

	uintMaxPrecision := uint(maxPrecision.Uint64())

	timeStart = time.Now()
	binDigits, binNewResidualXNum, err := mathops.BigIntMathLogarithms{}.GetIntDigits(binBase, binXNum, uintMaxPrecision)
	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("-----------------------------------------------")
	fmt.Println("  BigIntMathLogarithms{}.GetIntDigits()")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                    base: ", binBase.GetNumStr())
	fmt.Println("                    xNum: ", binXNum.GetNumStr())
	fmt.Println("            maxPrecision: ", uintMaxPrecision)
	fmt.Println("----------------- Results ---------------------")
	fmt.Println("                  digits: ", binDigits.GetNumStr())
	fmt.Println("         newResidualXNum: ", binNewResidualXNum.GetNumStr())
	fmt.Println("           expectedValue: ", expectedValue)
	fmt.Println("          Execution Time: ", duration)
	fmt.Println("------------------------------------------------")

}


func TestBigIntToNegativeFractionalPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int,
	expectedResult string) {


	timeStart := time.Now()
	result,
	resultPrecision,
	err := mathops.BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	timeEnd := time.Now()

	if err != nil {
		fmt.Printf("%v ", err.Error())
	}

	binResult := mathops.BigIntNum{}.NewBigInt(result, uint(resultPrecision.Uint64()))
	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("BigIntMathPower{}.BigIntToNegativeFractionalPower() ")
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

	binBase, _ := mathops.BigIntNum{}.NewBigIntPrecision(base, basePrecision)
	binExponent, _ := mathops.BigIntNum{}.NewBigIntPrecision(exponent, exponentPrecision)

	timeStart = time.Now()
	binPwr, err := mathops.BigIntMathPower{}.Pwr(binBase, binExponent, uint(maxPrecision.Uint64()))
	timeEnd = time.Now()
	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(...) " +
			"Error='%v' ", err.Error())
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println( "               BigIntMathPower{}.Pwr() ")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("     BigIntNum  result: ", binPwr.GetNumStr())
	fmt.Println("   BigIntNum precision: ", binPwr.GetPrecision())
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
}
