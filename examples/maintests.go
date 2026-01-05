package examples

import (
	"fmt"
	"math/big"
	"time"

	"github.com/mikeaustin71/mathops"
)

func TestBigIntToPositiveFractionalPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int,
	expectedResult string) {

	ePrefix := "TestBigIntToPositiveFractionalPower"

	timeStart := time.Now()

	result,
		resultPrecision,
		err := new(mathops.BigIntMathPower).BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			" result, resultPrecision, err := new(mathops.BigIntNum).\n"+
			"  base, basePrecision, exponent, exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	timeEnd := time.Now()

	binResult, err := new(mathops.BigIntNum).NewBigInt(result, uint(resultPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(result, uint(resultPrecision.Uint64()))\n"+
			"result= '%v'\n"+
			"uint(resultPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			uint(resultPrecision.Uint64()),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	codeDurationStr := CodeDurationToStr(timeDuration)

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("BigIntMathPower{}.BigIntToPositiveFractionalPower() ")
	fmt.Println("============================================================")
	fmt.Println("                  base: ", base.Text(10))
	fmt.Println("         basePrecision: ", basePrecision.Text(10))
	fmt.Println("              exponent: ", exponent.Text(10))
	fmt.Println("     exponentPrecision: ", exponentPrecision.Text(10))
	fmt.Println("     Maximum Precision: ", maxPrecision.Text(10))
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision)
	fmt.Println("         result NumStr: ", binResultNumStr)
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binBase, err := new(mathops.BigIntNum).NewBigIntBigPrecision(base, basePrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binBase, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(base, basePrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			err.Error())
		return
	}

	binBaseNumStr, err := binBase.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binBaseNumStr, err := binBase.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binExponent, err := new(mathops.BigIntNum).NewBigIntBigPrecision(exponent, exponentPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binExponent, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(exponent, exponentPrecision)\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponent.Text(10),
			exponentPrecision.Text(10),
			err.Error())
		return
	}

	binExponentNumStr, err := binExponent.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binExponentNumStr, err := binExponent.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	timeStart = time.Now()

	binPwr, err := new(mathops.BigIntMathPower).BigIntNumPwr(binBase, binExponent, uint(maxPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwr, err := new(mathops.BigIntMathPower).\n"+
			" BigIntNumPwr(binBase, binExponent, uint(maxPrecision.Uint64()))\n"+
			"binBase= '%v'\n"+
			"binExponent= '%v'\n"+
			"uint(maxPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binBaseNumStr,
			binExponentNumStr,
			uint(maxPrecision.Uint64()),
			err.Error())
		return
	}

	timeEnd = time.Now()

	timeDuration = timeEnd.Sub(timeStart)

	codeDurationStr = CodeDurationToStr(timeDuration)

	binPwrNumStr, err := binPwr.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwrNumStr, err := binPwr.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binPwrPrecisionInt, err := binPwr.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwrPrecisionInt, err := binPwr.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("               BigIntMathPower{}.BigIntNumPwr() ")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("     BigIntNum  result: ", binPwrNumStr)
	fmt.Println("   BigIntNum precision: ", binPwrPrecisionInt)
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

}

func TestBigIntNegativeFractionalNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestBigIntNegativeFractionalNthRoot"

	fdNr := new(mathops.FixedDecimalNthRoot)

	timeStart := time.Now()

	result, resultPrecision, err :=
		fdNr.CalculateNegativeFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err :=\n"+
			"  fdNr.CalculateNegativeFractionalNthRoot(\n"+
			"    radicand, radicandPrecision, nthRoot,\n"+
			"     nthRootPrecision, maxPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	timeEnd := time.Now()

	resultBiNum, err := new(mathops.BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultBiNum, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	codeDurationStr := CodeDurationToStr(timeDuration)

	resultBiNumNumStr, err := resultBiNum.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("FixedDecimalNthRoot{}.CalculateNegativeFractionalNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", radicand.Text(10))
	fmt.Println("     radicandPrecision: ", radicandPrecision.Text(10))
	fmt.Println("               nthRoot: ", nthRoot.Text(10))
	fmt.Println("      nthRootPrecision: ", nthRootPrecision)
	fmt.Println("     Maximum Precision: ", maxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNumNumStr)
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binRadicand, err := new(mathops.BigIntNum).NewBigIntBigPrecision(radicand, radicandPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicand, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(radicand, radicandPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			err.Error())
		return
	}

	binRadicandNumStr, err := binRadicand.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicandNumStr, err := binRadicand.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binNthRoot, err := new(mathops.BigIntNum).NewBigIntBigPrecision(nthRoot, nthRootPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binNthRoot, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(nthRoot, nthRootPrecision)\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			err.Error())
		return
	}

	binNthRootNumStr, err := binNthRoot.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binNthRootNumStr, err := binNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	uintMaxPrecisionUint := uint(maxPrecision.Uint64())

	timeStart = time.Now()

	binRoot, err := new(mathops.BigIntMathNthRoot).GetNthRoot(binRadicand, binNthRoot, uintMaxPrecisionUint)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRoot, err := new(mathops.BigIntMathNthRoot).\n"+
			"  GetNthRoot(binRadicand, binNthRoot, uintMaxPrecisionUint)\n"+
			"binRadicand= '%v'\n"+
			"binNthRoot= '%v'\n"+
			"uintMaxPrecisionUint= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binRadicandNumStr,
			binNthRootNumStr,
			uintMaxPrecisionUint,
			err.Error())
		return
	}

	timeEnd = time.Now()

	timeDuration = timeEnd.Sub(timeStart)

	codeDurationStr = CodeDurationToStr(timeDuration)

	binRadicandNumStr, err = binRadicand.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicandNumStr, err = binRadicand.GetNumStr())\n"+
			"Initialization #2\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binRadicandPrecisionInt, err := binRadicand.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicandPrecisionInt, err := \n"+
			" binRadicand.GetPrecisionInt()\n"+
			"binRadicand= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binRadicandNumStr,
			err.Error())
		return
	}

	binNthRootPrecisionInt, err := binNthRoot.GetPrecisionInt()

	if err != nil {
		fmt.Printf("Error returned by:\n"+
			"binNthRootPrecisionInt, err :=\n"+
			"  binNthRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n", err.Error())
		return
	}

	binRootNumStr, err := binRoot.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			" binRootNumStr, err := binRoot.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binRootPrecisionInt, err := binRoot.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"binRoot= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binRootNumStr,
			err.Error())
		return
	}

	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println("           BigIntMathNthRoot{}.GetNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", binRadicandNumStr)
	fmt.Println("     radicandPrecision: ", binRadicandPrecisionInt)
	fmt.Println("               nthRoot: ", binNthRootNumStr)
	fmt.Println("      nthRootPrecision: ", binNthRootPrecisionInt)
	fmt.Println("     Maximum Precision: ", uintMaxPrecisionUint)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", binRootNumStr)
	fmt.Println("       resultPrecision: ", binRootPrecisionInt)
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

}

func TestBigIntPositiveFractionalNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestBigIntPositiveFractionalNthRoot"

	fdNr := new(mathops.FixedDecimalNthRoot)

	timeStart := time.Now()

	result, resultPrecision, err :=
		fdNr.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err :=\n"+
			"  fdNr.CalculatePositiveFractionalNthRoot(\n"+
			"   radicand, radicandPrecision, nthRoot,\n"+
			"   nthRootPrecision, maxPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	timeEnd := time.Now()

	resultBiNum, err := new(mathops.BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultBiNum, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	codeDurationStr := CodeDurationToStr(timeDuration)

	resultBiNumNumStr, err := resultBiNum.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultBiNumNumStr, err := resultBiNum.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

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
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNumNumStr)
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binRadicand, err := new(mathops.BigIntNum).NewBigIntBigPrecision(radicand, radicandPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicand, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(radicand, radicandPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			err.Error())
		return
	}

	binRadicandNumStr, err := binRadicand.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicandNumStr, err := binRadicand.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binNthRoot, err := new(mathops.BigIntNum).NewBigIntBigPrecision(nthRoot, nthRootPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binNthRoot, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(nthRoot, nthRootPrecision)\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			err.Error())
		return
	}

	binNthRootNumStr, err := binNthRoot.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binNthRootNumStr, err := binNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	uintMaxPrecision := uint(maxPrecision.Uint64())

	timeStart = time.Now()

	binRoot, err := new(mathops.BigIntMathNthRoot).GetNthRoot(binRadicand, binNthRoot, uintMaxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRoot, err := new(mathops.BigIntMathNthRoot).\n"+
			"  GetNthRoot(binRadicand, binNthRoot, uintMaxPrecision)\n"+
			"binRadicand= '%v'\n"+
			"binNthRoot= '%v'\n"+
			"uintMaxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binRadicandNumStr,
			binNthRootNumStr,
			uintMaxPrecision,
			err.Error())
		return
	}

	timeEnd = time.Now()

	timeDuration = timeEnd.Sub(timeStart)

	codeDurationStr = CodeDurationToStr(timeDuration)

	binRadicandPrecisionInt, err := binRadicand.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRadicandPrecisionInt, err := binRadicand.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binNthRootPrecisionInt, err := binNthRoot.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binNthRootPrecisionInt, err := binNthRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binNthRootNumStr, err = binRoot.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binNthRootNumStr, err = binRoot.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binRootPrecisionInt, err := binRoot.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println("           BigIntMathNthRoot{}.GetNthRoot() ")
	fmt.Println("============================================================")
	fmt.Println("              radicand: ", binRadicandNumStr)
	fmt.Println("     radicandPrecision: ", binRadicandPrecisionInt)
	fmt.Println("               nthRoot: ", binNthRootNumStr)
	fmt.Println("      nthRootPrecision: ", binNthRootPrecisionInt)
	fmt.Println("     Maximum Precision: ", uintMaxPrecision)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", binNthRootNumStr)
	fmt.Println("       resultPrecision: ", binRootPrecisionInt)
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
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

	ePrefix := "TestBigIntNegativeIntNthRoot()"

	fdNr := new(mathops.FixedDecimalNthRoot)

	timeStart := time.Now()

	result, resultPrecision, err :=
		fdNr.CalculateNegativeIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err :=\n"+
			" fdNr.CalculateNegativeIntegerNthRoot(\n"+
			"  radicand, radicandPrecision, nthRoot,\n"+
			"   nthRootPrecision, maxPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"param3= '%v'\n"+
			"param4= '%v'\n"+
			"param5= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	timeEnd := time.Now()

	resultBiNum, err := new(mathops.BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultBiNum, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	codeDurationStr := CodeDurationToStr(timeDuration)

	resultBiNumStr, err := resultBiNum.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

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
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNumStr)
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
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

	ePrefix := "TestBigIntPositiveIntNthRoot()"

	var timeStart, timeEnd time.Time

	fdNr := new(mathops.FixedDecimalNthRoot)

	timeStart = time.Now()

	result, resultPrecision, err :=
		fdNr.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			nthRoot,
			nthRootPrecision,
			maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err :=\n"+
			"\n fdNr.CalculatePositiveIntegerNthRoot("+
			"  radicand, radicandPrecision, nthRoot,\n"+
			"   nthRootPrecision, maxPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '%v'\n"+
			"param5= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			nthRoot.Text(10),
			err.Error())
		return
	}

	timeEnd = time.Now()

	resultBiNum, err := new(mathops.BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultBiNum, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	codeDurationStr := CodeDurationToStr(timeDuration)

	resultBiNumStr, err := resultBiNum.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultBiNumStr, err := resultBiNum.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

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
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision.Text(10))
	fmt.Println("         result NumStr: ", resultBiNumStr)
	fmt.Println("       expected result: ", expectedValue)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
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

	ePrefix := "maintests.TestBigIntNegativeIntPower()"

	timeStart := time.Now()

	result,
		resultPrecision,
		err := new(mathops.BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(mathops.BigIntMathPower).\n"+
			" BigIntToNegativeIntegerPower(\n"+
			"  base, basePrecision, exponent, exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	timeEnd := time.Now()

	binResult, err := new(mathops.BigIntNum).NewBigInt(result, uint(resultPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(result, uint(resultPrecision.Uint64()))\n"+
			"result= '%v'\n"+
			"uint(resultPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			uint(resultPrecision.Uint64()),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)
	codeDurationStr := CodeDurationToStr(timeDuration)

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

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
	fmt.Println("         result NumStr: ", binResultNumStr)
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binBase, err := new(mathops.BigIntNum).NewBigIntBigPrecision(base, basePrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binBase, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(base, basePrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			err.Error())
		return
	}

	binBaseNumStr, err := binBase.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binBaseNumStr, err := binBase.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binExponent, err := new(mathops.BigIntNum).NewBigIntBigPrecision(exponent, exponentPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binExponent, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(exponent, exponentPrecision)\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponent.Text(10),
			exponentPrecision.Text(10),
			err.Error())
		return
	}

	binExponentNumStr, err := binExponent.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binExponentNumStr, err := binExponent.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	timeStart = time.Now()

	binPwr, err := new(mathops.BigIntMathPower).BigIntNumPwr(binBase, binExponent, uint(maxPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwr, err := new(mathops.BigIntMathPower).\n"+
			"  BigIntNumPwr(binBase, binExponent, uint(maxPrecision.Uint64()))\n"+
			"binBase= '%v'\n"+
			"binExponent= '%v'\n"+
			"uint(maxPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binBaseNumStr,
			binExponentNumStr,
			uint(maxPrecision.Uint64()),
			err.Error())
		return
	}

	timeEnd = time.Now()

	timeDuration = timeEnd.Sub(timeStart)

	binPwrNumStr, err := binPwr.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwrNumStr, err := binPwr.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binPwrPrecisionInt, err := binPwr.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwrPrecisionInt, err := binPwr.GetPrecisionInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	codeDurationStr = CodeDurationToStr(timeDuration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("               BigIntMathPower{}.BigIntNumPwr() ")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("     BigIntNum  result: ", binPwrNumStr)
	fmt.Println("   BigIntNum precision: ", binPwrPrecisionInt)
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", codeDurationStr)
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

	ePrefix := "maintests.TestBigIntPositiveIntPower()"

	timeStart := time.Now()
	result,
		resultPrecision,
		err := new(mathops.BigIntMathPower).BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(mathops.BigIntMathPower).\n"+
			"  BigIntToPositiveIntegerPower(\n"+
			"  base, basePrecision, exponent, exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	timeEnd := time.Now()

	binResult, err := new(mathops.BigIntNum).NewBigInt(result, uint(resultPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(result, uint(resultPrecision.Uint64()))\n"+
			"result= '%v'\n"+
			"uint(resultPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			result.Text(10),
			uint(resultPrecision.Uint64()),
			err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := CodeDurationToStr(timeDuration)

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

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
	fmt.Println("         result NumStr: ", binResultNumStr)
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binBase, err := new(mathops.BigIntNum).NewBigIntBigPrecision(base, basePrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binBase, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(base, basePrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			err.Error())
		return
	}

	binBaseNumStr, err := binBase.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binBaseNumStr, err := binBase.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	binExponent, err := new(mathops.BigIntNum).NewBigIntBigPrecision(exponent, exponentPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binExponent, err := new(mathops.BigIntNum).\n"+
			" NewBigIntBigPrecision(exponent, exponentPrecision)\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponent.Text(10),
			exponentPrecision.Text(10),
			err.Error())
		return
	}

	binExponentNumStr, err := binExponent.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binExponentNumStr, err := binExponent.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	timeStart = time.Now()

	binPwr, err := new(mathops.BigIntMathPower).BigIntNumPwr(binBase, binExponent, uint(maxPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwr, err := new(mathops.BigIntMathPower).\n"+
			"  BigIntNumPwr(binBase, binExponent, uint(maxPrecision.Uint64()))\n"+
			"binBase= '%v'\n"+
			"binExponent= '%v'\n"+
			"uint(maxPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			binBaseNumStr,
			binExponentNumStr,
			uint(maxPrecision.Uint64()),
			err.Error())
		return
	}

	timeEnd = time.Now()

	timeDuration = timeEnd.Sub(timeStart)

	binPwrNumStr, err := binPwr.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwrNumStr, err := binPwr.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	binPwrPrecisionInt, err := binPwr.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binPwrPrecisionInt, err := binPwr.GetPrecisionInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	duration = CodeDurationToStr(timeDuration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("               BigIntMathPower{}.BigIntNumPwr() ")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("     BigIntNum  result: ", binPwrNumStr)
	fmt.Println("   BigIntNum precision: ", binPwrPrecisionInt)
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

	ePrefix := "TestFixDecNthRootFmtFracDigits()"

	nthRootCalc := new(mathops.FixedDecimalNthRoot)

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

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"err := nthRootCalc.FormatCalculationConstants(\n"+
			"  radicand, radicandPrecision, nthRoot, nthRootPrecision, maxPrecision)\n"+
			"radicand= '%v'\n"+
			"radicandPrecision= '%v'\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			radicand.Text(10),
			radicandPrecision.Text(10),
			nthRoot.Text(10),
			nthRootPrecision.Text(10),
			maxPrecision,
			err.Error())
		return
	}

	timeEnd := time.Now()

	remainder := new(big.Int)

	intRadicand, fracRadicand := big.NewInt(0).QuoRem(radicand, radicandPrecision, remainder)

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

	duration := CodeDurationToStr(timeDuration)
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

	ePrefix := "TestFixDecNthRootGetNextFracBundle()"

	nthRootCalc := new(mathops.FixedDecimalNthRoot)

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

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"err := nthRootCalc.FormatCalculationConstants(\n"+
			"  fracNum, fracPrecision, nthRoot, 0, 9)\n"+
			"fracNum= '%v'\n"+
			"fracPrecision= '%v'\n"+
			"nthRoot= '%v'\n"+
			"nthRoot Precision= '0'\n"+
			"maxPrecision= '9'\n"+
			"Error='%v'\n\n",
			ePrefix,
			fracNum.Text(10),
			fracPrecision.Text(10),
			nthRoot.Text(10),
			err.Error())
		return
	}

	fmtFracNum, fmtFracPrecision, err :=
		nthRootCalc.FormatFractionalDigitsFromRadicand(
			fracNum,
			fracPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"fmtFracNum, fmtFracPrecision, err :=\n"+
			"  nthRootCalc.FormatFractionalDigitsFromRadicand(\n"+
			"    fracNum, fracPrecision)\n"+
			"fracNum= '%v'\n"+
			"xrayStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			fracNum.Text(10),
			fracPrecision.Text(10),
			err.Error())
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

	for fmtFracPrecision.Cmp(bigZero) == 1 {

		timeStart = time.Now()

		nextBundle, fmtFracNum, fmtFracPrecision, err =
			nthRootCalc.GetNextFractionalBundleFromRadicand(
				fmtFracNum,
				fmtFracPrecision)

		if err != nil {
			fmt.Printf("%v\n"+
				"Error returned by:\n"+
				"nextBundle, fmtFracNum, fmtFracPrecision, err =\n"+
				" nthRootCalc.GetNextFractionalBundleFromRadicand(\n"+
				"  fmtFracNum, fmtFracPrecision)\n"+
				"Error='%v'\n\n",
				ePrefix,
				err.Error())
			return
		}

		timeEnd = time.Now()

		cycle++
		timeDuration := timeEnd.Sub(timeStart)

		duration := CodeDurationToStr(timeDuration)
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

	ePrefix := "TestFixDecNthRootNextIntBundle"

	bigZero := big.NewInt(0)
	nextBundle := big.NewInt(0)
	nextBundleTotDigits := big.NewInt(0)
	var err error
	cycle := 0
	nthRootCalc := mathops.FixedDecimalNthRoot{}

	if integerNum == nil {
		fmt.Printf("%v\n"+
			"Error: Input Parameter is 'nil'\n"+
			"integerNum == nill\n\n",
			ePrefix)
		return

	}

	if intTotalDigits == nil {
		fmt.Printf("%v\n"+
			"Error: Input Parameter is 'nil'\n"+
			"intTotalDigits == nill\n\n",
			ePrefix)
		return

	}

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
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"err = nthRootCalc.FormatCalculationConstants(\n"+
			" integerNum, big.NewInt(0), nthRoot, big.NewInt(0),\n"+
			"  big.NewInt(9))\n"+
			"integerNum= '%v'\n"+
			"radicandPrecision= '0'\n"+
			"nthRoot= '%v'\n"+
			"nthRootPrecision= '0'\n"+
			"maxPrecision= '9'\n"+
			"Error='%v'\n\n",
			ePrefix,
			integerNum.Text(10),
			nthRoot.Text(10),
			err.Error())
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

	var lastIntegerNum, lastIntTotalDigits big.Int

	lastIntegerNum.Set(integerNum)

	lastIntTotalDigits.Set(intTotalDigits)

	for intTotalDigits.Cmp(bigZero) == 1 {

		timeStart = time.Now()

		nextBundle, integerNum, intTotalDigits, err =
			nthRootCalc.GetNextIntegerBundleFromRadicand(
				&lastIntegerNum,
				&lastIntTotalDigits)

		if err != nil {
			fmt.Printf("%v\n"+
				"Error returned by:\n"+
				"nextBundle, integerNum, intTotalDigits, err =\n"+
				" nthRootCalc.GetNextIntegerBundleFromRadicand(\n"+
				" lastIntegerNum, lastIntTotalDigits)\n"+
				"lastIntegerNum= '%v'\n"+
				"intTotalDigits= '%v'\n"+
				"cycle= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix,
				lastIntegerNum.Text(10),
				lastIntTotalDigits.Text(10),
				cycle,
				err.Error())
			return
		}

		timeEnd = time.Now()

		fmt.Println("                 Cycle: ", cycle)
		fmt.Println("           Next Bundle: ", nextBundle.Text(10))
		fmt.Println("Next Bundle Tot-Digits: ", nextBundleTotDigits.Text(10))
		fmt.Println("               integerNum: ", integerNum.Text(10))
		fmt.Println("     Num of Int Digits: ", intTotalDigits.Text(10))
		fmt.Println("               nthRoot: ", nthRoot.Text(10))
		timeDuration := timeEnd.Sub(timeStart)

		duration := CodeDurationToStr(timeDuration)
		fmt.Println("         Time Duration: ", duration)
		fmt.Println("--------------------------------------------------------")
		fmt.Println()

		lastIntegerNum.Set(integerNum)

		lastIntTotalDigits.Set(intTotalDigits)

		cycle++

	}

	return
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

func TestEPwrXFromTaylorSeriesFixedDecimal(
	exponentX mathops.BigIntFixedDecimal,
	a, nCycles uint, expectedResult string) {

	ePrefix := "maintests.TestEPwrXFromTaylorSeriesFixedDecimal()"

	timeStart := time.Now()

	exponentXNumStr, err := exponentX.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	result, err := new(mathops.BigIntMathLogarithms).EPwrXFromTaylorSeriesFixedDecimal(
		exponentX,
		a,
		nCycles)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(mathops.BigIntMathLogarithms).\n"+
			"  EPwrXFromTaylorSeriesFixedDecimal(\n"+
			"   exponentX, a, nCycles)\n"+
			"exponentX= '%v'\n"+
			"a= '%v'\n"+
			"nCycles= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponentXNumStr,
			a,
			nCycles,
			err.Error())
		return
	}

	timeEnd := time.Now()

	exponentXNumStr, err = exponentX.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"exponentXNumStr, err := exponentX.GetNumStr()\n"+
			"Initialization #2\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	exponentXPrecisionInt, err := exponentX.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"exponentXPrecisionInt, err := exponentX.GetPrecisionInt()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	resultPrecisionInt, err := result.GetPrecisionInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionInt, err := result.GetPrecisionInt()\n"+
			"result= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			resultNumStr,
			err.Error())
		return
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------------")
	fmt.Println("BigIntMathLogarithms{}.EPwrXFromTaylorSeriesFixedDecimal()")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("           nCycles: ", nCycles)
	fmt.Println("           A-Value: ", a)
	fmt.Println("          exponent: ", exponentXNumStr)
	fmt.Println("exponent Precision: ", exponentXPrecisionInt)
	fmt.Println("            xValue: ", resultNumStr)
	fmt.Println("    expectedXValue: ", expectedResult)
	fmt.Println("  xValue Precision: ", resultPrecisionInt)

	timeDuration := timeEnd.Sub(timeStart)

	duration := CodeDurationToStr(timeDuration)
	fmt.Println("   Time Duration: ", duration)

	return
}

func TestEPwrXFromTaylorSeries(exponent, binA mathops.BigIntNum, nCycles int64, expectedXValue string) {

	ePrefix := "maintests.TestEPwrXFromTaylorSeries()"

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

	xValue, err := new(mathops.BigIntMathLogarithms).EPwrXFromTaylorSeries(exponent, binA, nCycles)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"xValue, err := new(mathops.BigIntMathLogarithms).\n"+
			"  EPwrXFromTaylorSeries(exponent, binA, nCycles)\n"+
			"exponent= '%v'\n"+
			"binA= '%v'\n"+
			"nCycles= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponent.String(),
			binA.String(),
			nCycles,
			err.Error())
		return
	}

	timeEnd := time.Now()

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"exponentNumStr, err := exponent.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	xValueNumStr, err := xValue.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"xValueNumStr, err := xValue.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fmt.Println(" nCycles: ", nCycles)
	fmt.Println("exponent: ", exponentNumStr)
	fmt.Println("  xValue: ", xValueNumStr)
	fmt.Println("expectedXValue: ", expectedXValue)
	fmt.Println("StartTime: ", timeStart.String())
	fmt.Println("  EndTime: ", timeEnd.String())

	timeDuration := timeEnd.Sub(timeStart)

	duration := CodeDurationToStr(timeDuration)
	fmt.Println("Time Duration: ", duration)

	return
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

	PowerTest01(base, exponent, maxPrecision)

*/

func TestBigIntDivide(
	dividend,
	dividendPrecision,
	divisor,
	divisorPrecision,
	maxPrecision *big.Int,
	expectedResult string) {

	ePrefix := "maintests.TestBigIntDivide()"

	quotient, quotientPrecision, err :=
		new(mathops.BigIntMathDivide).BigIntFracQuotient(
			dividend,
			dividendPrecision,
			divisor,
			divisorPrecision,
			maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"quotient, quotientPrecision, err :=\n"+
			" .BigIntFracQuotient(\n"+
			"   dividend, dividendPrecision, divisor,\n"+
			"    divisorPrecision, maxPrecision)\n"+
			"dividend= '%v'\n"+
			"dividendPrecision= '%v'\n"+
			"divisor= '%v'\n"+
			"divisorPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividend.Text(10),
			dividendPrecision.Text(10),
			divisor.Text(10),
			divisorPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())
		return
	}

	biNumResult, err := new(mathops.BigIntNum).NewBigInt(quotient, uint(quotientPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumResult, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(quotient, uint(quotientPrecision.Uint64()))\n"+
			"quotient= '%v'\n"+
			"uint(quotientPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			quotient.Text(10),
			uint(quotientPrecision.Uint64()),
			err.Error())
		return
	}

	biNumResultNumStr, err := biNumResult.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumResultNumStr, err := biNumResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	timeStart := time.Now()

	biNumExpectedResult, err := new(mathops.BigIntNum).NewNumStr(expectedResult)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumExpectedResult, err := new(mathops.BigIntNum).\n"+
			" NewNumStr(expectedResult)\n"+
			"expectedResult= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedResult,
			err.Error())
		return
	}

	timeEnd := time.Now()

	biNumExpectedResultNumStr, err := biNumExpectedResult.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumExpectedResultNumStr, err := biNumExpectedResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biNumResultPrecisionInt, err := biNumResult.GetPrecisionUint()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumResultPrecisionInt, err := \n"+
			" biNumResult.GetPrecisionUint()\n"+
			"biNumResult= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			biNumExpectedResultNumStr,
			err.Error())
		return
	}

	biNumExpectedResultPrecisionInt, err := biNumExpectedResult.GetPrecisionUint()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumExpectedResultPrecisionInt, err := \n"+
			" biNumExpectedResult.GetPrecisionUint()\n"+
			"biNumExpectedResult= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			biNumExpectedResultNumStr,
			err.Error())
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
	fmt.Println("          quotient Result: ", biNumResultNumStr)
	fmt.Println("quotient Result Precision: ", biNumResultPrecisionInt)
	fmt.Println("          expected Result: ", expectedResult)
	fmt.Println("   bi Num Expected Result: ", biNumExpectedResultNumStr)
	fmt.Println("bi Num Expected Precision: ", biNumExpectedResultPrecisionInt)
	fmt.Println("====================================================")

	if biNumExpectedResult.GetNumStr() == biNumResult.GetNumStr() {

		fmt.Println("Success! Expected Result Matches Actual Result!")

	} else {

		fmt.Println("Failure! Expected Result DOES NOT Match Actual Result!")

	}
	fmt.Println("====================================================")

	timeDuration := timeEnd.Sub(timeStart)

	duration := CodeDurationToStr(timeDuration)

	fmt.Println("Time Duration: ", duration)
	fmt.Println("====================================================")

	biNumDividend, err := new(mathops.BigIntNum).NewBigInt(dividend, uint(dividendPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumDividend, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(dividend, uint(dividendPrecision.Uint64()))\n"+
			"dividend= '%v'\n"+
			"uint(dividendPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			dividend.Text(10),
			uint(dividendPrecision.Uint64()),
			err.Error())
		return
	}

	biNumDivisor, err := new(mathops.BigIntNum).NewBigInt(divisor, uint(divisorPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumDivisor, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(divisor, uint(divisorPrecision.Uint64()))\n"+
			"divisor= '%v'\n"+
			"uint(divisorPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			divisor.Text(10),
			uint(divisorPrecision.Uint64()),
			err.Error())
		return
	}

	numSeps := new(mathops.NumericSeparatorDto).NewUSADefaults()

	timeStart = time.Now()

	biNumQuotient, err := new(mathops.BigIntMathDivide).BigIntNumFracQuotient(biNumDividend, biNumDivisor, numSeps, uint(maxPrecision.Uint64()))

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumQuotient, err := new(mathops.BigIntMathDivide).\n"+
			"  BigIntNumFracQuotient(biNumDividend, biNumDivisor,\n"+
			"    numSeps, uint(maxPrecision.Uint64()))\n"+
			"biNumDividend= '%v'\n"+
			"biNumDivisor= '%v'\n"+
			"numSeps= '%v'\n"+
			"uint(maxPrecision.Uint64())= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			biNumDividend.String(),
			biNumDivisor.String(),
			numSeps.String(),
			uint(maxPrecision.Uint64()),
			err.Error())
		return
	}

	timeEnd = time.Now()

	timeDuration = timeEnd.Sub(timeStart)
	duration = CodeDurationToStr(timeDuration)

	biNumDividendNumStr, err := biNumDividend.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumDividendNumStr, err := biNumDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biNumDivisorNumStr, err := biNumDivisor.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumDivisorNumStr, err := biNumDivisor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	biNumQuotientNumStr, err := biNumQuotient.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"biNumQuotientNumStr, err := biNumQuotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fmt.Println("-- BigIntMathDivide{}.BigIntNumFracQuotient() --")
	fmt.Println("--------------------------------------------------")
	fmt.Println("dividend: ", biNumDividendNumStr)
	fmt.Println(" divisor: ", biNumDivisorNumStr)
	fmt.Println("quotient: ", biNumQuotientNumStr)
	fmt.Println("====================================================")
	fmt.Println("Time Duration: ", duration)

	return
}

func TestBigIntPwr(
	base *big.Int,
	basePrecision,
	exponent,
	internalMaxPrecision,
	outputMaxPrecision uint,
	expectedResult string) {

	ePrefix := "TestBigIntPwr"

	baseToPwr, baseToPwrPrecision, err := new(mathops.BigIntMathPower).
		BigIntPwrIteration(
			base,
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(mathops.BigIntMathPower).\n"+
			"  BigIntPwrIteration(base, basePrecision, exponent,\n"+
			"   internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())
		return
	}

	bINumResult, err := new(mathops.BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bINumResult, err := new(mathops.BigIntNum).\n"+
			" NewBigInt(baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
			err.Error())
		return
	}

	bINumResultNumStr, err := bINumResult.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bINumResultNumStr, err := bINumResult.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	fmt.Println("TestBigIntPwr")
	fmt.Println("                     base: ", base.Text(10))
	fmt.Println("            basePrecision: ", basePrecision)
	fmt.Println("                 exponent: ", exponent)
	fmt.Println("     internalMaxPrecision: ", internalMaxPrecision)
	fmt.Println("       outputMaxPrecision: ", outputMaxPrecision)
	fmt.Println("----------------------------------------------")
	fmt.Println("         Result baseToPwr: ", baseToPwr.Text(10))
	fmt.Println("Result baseToPwrPrecision: ", baseToPwrPrecision)
	fmt.Println("  Result BigIntNum NumStr: ", bINumResultNumStr)
	fmt.Println("          Expected Result: ", expectedResult)

	return
}

func PowerTest01(base, exponent mathops.BigIntNum, maxPrecision, i uint) {

	ePrefix := "PowerTest01"

	powerValue, err := new(mathops.BigIntMathPower).BigIntNumPwr(base, exponent, maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigTarget, err := new(mathops.BigIntNum).\n"+
			"  BigIntNumPwr(base, exponent, maxPrecision)\n"+
			"base= '%v'\n"+
			"exponent= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Input Parameter 'i'= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			base.String(),
			exponent.String(),
			maxPrecision,
			i,
			err.Error())
		return
	}

	baseNumStr, err := base.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"baseNumStr, err := base.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"exponentNumStr, err := exponent.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	powerValueNumStr, err := powerValue.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"powerValueNumStr, err := powerValue.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Println("     Index: ", i)
	fmt.Println("      base: ", baseNumStr)
	fmt.Println("  exponent: ", exponentNumStr)
	fmt.Println("powerValue: ", powerValueNumStr)
	fmt.Println("--------------------------------------")
	fmt.Println()

	return
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
