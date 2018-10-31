package main

import (
	"../examples"
	"../mathops"
	"fmt"
	"math/big"
	"time"
)

func main() {
	// integer residual value 1.430511474609375

	aNum := big.NewInt(24)
	aNumPrecision := big.NewInt(0)
	gNum := big.NewInt(6)
	gNumPrecision := big.NewInt(0)

	TestArithmeticGeometricMean(
		aNum,
		aNumPrecision,
		gNum,
		gNumPrecision)

}


func TestArithmeticGeometricMean(
	aNum,
	aNumPrecision,
	gNum,
	gNumPrecision *big.Int) {

	timeStart := time.Now()
	timeEnd := time.Now()
	ePrefix := "TestArithmeticGeometricMean() "

	timeStart = time.Now()

	agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err :=
		mathops.BigIntMath{}.ArithmeticGeometricMean(aNum, aNumPrecision, gNum, gNumPrecision)

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
	fmt.Println("           aNum: ", aNum.Text(10))
	fmt.Println("  aNumPrecision: ", aNumPrecision.Text(10))
	fmt.Println("      a Num Str: ", binANum.GetNumStr())
	fmt.Println("           gNum: ", gNum.Text(10))
	fmt.Println("  gNumPrecision: ", gNumPrecision.Text(10))
	fmt.Println("      g Num Str: ", binGNum.GetNumStr())
	fmt.Println("======================================================================")
	fmt.Println("         agMean: ", agMean.Text(10))
	fmt.Println("agMeanPrecision: ", agMeanPrecision.Text(10))
	fmt.Println("  agMean NumStr: ", binAGMean.GetNumStr())
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("         gValue: ", gValue.Text(10))
	fmt.Println("gValuePrecision: ", gValuePrecision.Text(10))
	fmt.Println("  gValue NumStr: ", binGValue.GetNumStr())
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("         Cycles: ", cycles)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println(" Execution Time: ", duration)
	fmt.Println("======================================================================")
	fmt.Println()

}

func TestBigIntLogBaseOfX(
	base,
	xNum,
	xNumPrecision,
	maxInternalPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestBigIntLogBaseOfX() "

	timeStart := time.Now()
	logResult, logResultPrecision, err :=
		mathops.BigIntMathLogarithms{}.BigIntLogBaseOfX(
			base,
			xNum,
			xNumPrecision,
			maxInternalPrecision,
			maxPrecision)

	timeEnd := time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v ", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

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

	fmt.Println()
	fmt.Println("------------------------------------------------")
	fmt.Println("      BigIntMathLogarithms.LogBaseOfX() ")
	fmt.Println("------------------------------------------------")
	fmt.Println("BigIntNum Log Result: ", biNumResult.GetNumStr())
	fmt.Println("     Expected Result: ", expectedValue)
	fmt.Println("      Execution Time: ", duration)
	fmt.Println("------------------------------------------------")
	fmt.Println()
}

func TestGetBigIntDecimalDigits(
	base,
	residualXNum,
	residualXNumPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestGetBigIntDigits() "

	timeStart := time.Now()

	digit, newResidualXNum, newResidualXNumPrecision, err :=
	 	mathops.BigIntMathLogarithms{}.BigIntGetNextDecimalDigit(
	 		base,
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
	xNum,
	xNumPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestGetBigIntDigits() "

	timeStart := time.Now()
	digits, newResidualXNum, newResidualXNumPrecision, err :=
		mathops.BigIntMathLogarithms{}.BigIntGetIntDigits(
			base,
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
