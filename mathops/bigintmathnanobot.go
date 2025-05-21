package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathNanobot struct {
	lock *sync.Mutex
}

// ArithmeticGeometricMean
//
// Computes the Arithmetic-Geometric Mean of two numbers. See:
//
// https://en.wikipedia.org/wiki/Arithmetic–geometric_mean
//
// Dev Note: maxInternalPrecision may need to be 60 to 75 times
// targetPrecision.
func (bIMathNanobot *bigIntMathNanobot) arithmeticGeometricMeanBigInt(
	aNum *big.Int,
	aNumPrecision *big.Int,
	gNum *big.Int,
	gNumPrecision *big.Int,
	maxInternalPrecision *big.Int,
	targetPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (agMean *big.Int,
	agMeanPrecision *big.Int,
	gValue *big.Int,
	gValuePrecision *big.Int,
	cycles uint64,
	err error) {

	if bIMathNanobot.lock == nil {
		bIMathNanobot.lock = new(sync.Mutex)
	}

	bIMathNanobot.lock.Lock()

	defer bIMathNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNanobot.arithmeticGeometricMeanBigInt",
		"")

	if err != nil {
		return big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			0, err
	}

	agMean = big.NewInt(0)
	agMeanPrecision = big.NewInt(0)
	gValue = big.NewInt(0)
	gValuePrecision = big.NewInt(0)
	cycles = 0
	err = nil

	if aNum == nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'aNum'",
			}
	}

	if aNumPrecision == nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'aNumPrecision'",
			}
	}

	if gNum == nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'gNum'",
			}
	}

	if gNumPrecision == nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'gNumPrecision'",
			}
	}

	if maxInternalPrecision == nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'maxInternalPrecision'",
			}
	}

	if targetPrecision == nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'targetPrecision'",
			}
	}

	bigZero := big.NewInt(0)

	if aNumPrecision.Cmp(bigZero) == -1 {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("aNumPrecision='%v'", aNumPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'aNumPrecision' is less than zero and INVALID!",
			}
	}

	if gNumPrecision.Cmp(bigZero) == -1 {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("gNumPrecision='%v'", gNumPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'gNumPrecision' is less than zero and INVALID!",
			}
	}

	if maxInternalPrecision.Cmp(bigZero) == -1 {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("maxInternalPrecision='%v'", maxInternalPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'maxInternalPrecision' is less than zero and INVALID!",
			}
	}

	if targetPrecision.Cmp(bigZero) == -1 {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("targetPrecision='%v'", targetPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'targetPrecision' is less than zero and INVALID!",
			}
	}

	minInternalPrecision := big.NewInt(0).Add(targetPrecision, big.NewInt(120))

	if maxInternalPrecision.Cmp(minInternalPrecision) == -1 {
		maxInternalPrecision.Set(minInternalPrecision)
	}

	cycleLimit := maxInternalPrecision.Uint64() + uint64(20)

	// Assume aNum is less than gNum

	a := big.NewInt(0).Set(aNum)
	aPrecision := big.NewInt(0).Set(aNumPrecision)
	g := big.NewInt(0).Set(gNum)
	gPrecision := big.NewInt(0).Set(gNumPrecision)

	bIMathElectron := new(bigIntMathElectron)

	agCmp, err :=
		bIMathElectron.precisionCmpBigInt(
			aNum,
			aNumPrecision,
			gNum,
			gNumPrecision,
			ePrefix)

	if err != nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "agCmp, err :=bIMathElectron.precisionCmpBigInt(\n" +
					"aNum, aNumPrecision, gNum,  gNumPrecision, ePrefix)",
				ErrMessage: err.Error(),
			}
	}

	if agCmp == 0 {
		agMean.Set(aNum)
		agMeanPrecision.Set(aNumPrecision)
		gValue.Set(gNum)
		gValuePrecision.Set(gNumPrecision)
		cycles = 0
		err = nil
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if agCmp == 1 {
		// aNum is greater than gNum
		a.Set(gNum)
		aPrecision.Set(gNumPrecision)
		g.Set(aNum)
		gPrecision.Set(aNumPrecision)

	}

	origA := big.NewInt(0).Set(a)
	origAPrecision := big.NewInt(0).Set(aPrecision)
	origG := big.NewInt(0).Set(g)
	origGPrecision := big.NewInt(0).Set(gPrecision)
	aCom := big.NewInt(0)
	aComPrecision := big.NewInt(0)
	gCom := big.NewInt(0)
	gComPrecision := big.NewInt(0)
	factor := big.NewInt(2)
	factorPrecision := big.NewInt(0)
	aTest := big.NewInt(0)
	aTestPrecision := big.NewInt(0)
	gTest := big.NewInt(0)
	gTestPrecision := big.NewInt(0)

	oneHalfFactor, oneHalfFactorPrecision, err :=
		new(BigIntMathDivide).BigIntFracQuotient(
			big.NewInt(1),
			bigZero,
			big.NewInt(2),
			bigZero,
			maxInternalPrecision)

	if err != nil {

		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "oneHalfFactor, oneHalfFactorPrecision, err := \n" +
					"    new(BigIntMathDivide).BigIntFracQuotient(\n" +
					"    big.NewInt(1), bigZero, big.NewInt(2), bigZero, maxInternalPrecision)",
				ErrMessage: err.Error(),
			}
	}

	fdNRoot := FixedDecimalNthRoot{}
	cycleNum := uint64(0)
	/*
	   aCycleValue := BigIntNum{}.NewZero(0)
	   gCycleValue := BigIntNum{}.NewZero(0)
	   tempACyclePreAddPrecision := big.NewInt(0)
	   tempACyclePrecision := big.NewInt(0)
	   tempGCyclePrecision := big.NewInt(0)
	   tempGComPrecision:= big.NewInt(0)
	   gComPreRoundValue := BigIntNum{}.NewZero(0)
	   gComValue := BigIntNum{}.NewZero(0)

	   aCycleResult := BigIntNum{}.NewZero(0)
	   gCycleResult := BigIntNum{}.NewZero(0)
	*/

	for i := uint64(0); i < cycleLimit; i++ {

		cycleNum = i + 1

		// Debug capture
		/*
			aCycleValue, errX = BigIntNum{}.NewBigIntBigPrecision(a, aPrecision)

			if errX != nil {
				err = fmt.Errorf("%v\n" + "aCycleValue- %v\n", ePrefix, errX.Error())
				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
			}

			gCycleValue, errX = BigIntNum{}.NewBigIntBigPrecision(g, gPrecision)

			if errX != nil {
				err = fmt.Errorf("%v\n" +  "gCycleValue- %v", ePrefix, errX.Error())
				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
			}
			tempACyclePreAddPrecision = big.NewInt(0).Set(aPrecision)

		*/

		aCom, aComPrecision, err =
			new(BigIntMathAdd).BigIntAdd(
				a,
				aPrecision,
				g,
				gPrecision)

		if err != nil {

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "aCom, aComPrecision, err =\n" +
						"    new(BigIntMathAdd).BigIntAdd(\n" +
						"    a, aPrecision, g, gPrecision)",
					ErrMessage: err.Error(),
				}
		}

		gCom.Mul(a, g)
		// tempACyclePrecision = big.NewInt(0).Set(aPrecision)
		// tempGCyclePrecision = big.NewInt(0).Set(gPrecision)
		gComPrecision.Add(aPrecision, gPrecision)
		// tempGComPrecision = big.NewInt(0).Set(gComPrecision)

		/*
			if cycleNum == 7 {
				gComPreRoundValue, errX = BigIntNum{}.NewBigIntBigPrecision(gCom, gComPrecision)

				if errX != nil {
					err = fmt.Errorf("%v\n" + "gComPreRoundValue- %v\n", ePrefix, errX.Error())
					return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
				}

			}
		*/

		gCom, gComPrecision, err =
			bIMathElectron.roundToMaxPrecisionBigInt(
				gCom,
				gComPrecision,
				maxInternalPrecision,
				true,
				ePrefix)

		if err != nil {

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "gCom, gComPrecision, err =\n" +
						"    bIMathElectron.roundToMaxPrecisionBigInt(\n" +
						"    gCom, gComPrecision, maxInternalPrecision, true, ePrefix)",
					ErrMessage: err.Error(),
				}
		}

		a.Mul(aCom, oneHalfFactor)
		aPrecision.Add(aComPrecision, oneHalfFactorPrecision)

		a, aPrecision, err =
			bIMathElectron.roundToMaxPrecisionBigInt(
				a,
				aPrecision,
				maxInternalPrecision,
				true,
				ePrefix)

		if err != nil {

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "a, aPrecision, err =\n" +
						"     bIMathElectron.roundToMaxPrecisionBigInt(\n" +
						"    a, aPrecision, maxInternalPrecision, true, ePrefix)",
					ErrContext: fmt.Sprintf("maxInternalPrecision= '%v'",
						maxInternalPrecision.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		/*

			a, aPrecision, errX = BigIntMathDivide{}.BigIntFracQuotient(aCom, aComPrecision, factor, factorPrecision, maxInternalPrecision)

			if errX != nil {

				err = fmt.Errorf(fmt.Errorf("%v\n" +
				"Error returned by BigIntMathDivide{}.BigIntFracQuotient()" +
				"Error: %v\n",
				ePrefix,
				errX.Error())

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
			}



			gComValue, errX = BigIntNum{}.NewBigIntBigPrecision(gCom, gComPrecision)

			if errX != nil {

				err = fmt.Errorf("%v\n" +
				"Error returned by BigIntNum{}.NewBigIntBigPrecision(gCom, gComPrecision)" +
				"Error: %v",
				ePrefix,
				errX.Error())

				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
			}
		*/

		g, gPrecision, err =
			fdNRoot.CalculatePositiveIntegerNthRoot(
				gCom,
				gComPrecision,
				factor,
				factorPrecision,
				maxInternalPrecision)

		if err != nil {

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "g, gPrecision, err =\n" +
						"    fdNRoot.CalculatePositiveIntegerNthRoot(\n" +
						"    gCom, gComPrecision, factor, factorPrecision, maxInternalPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("gCom= '%v'; gComPrecision= '%v'; factor= '%v'\n"+
						"factorPrecision= '%v'; maxInternalPrecision= '%v'",
						gCom.Text(10), gComPrecision.Text(10), factor.Text(10), factorPrecision.Text(10), maxInternalPrecision.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		/*
				aCycleResult, errX =
					BigIntNum{}.NewBigIntBigPrecision(a, aPrecision)

				gCycleResult, errX =
					BigIntNum{}.NewBigIntBigPrecision(g, gPrecision)

				if errX != nil {

					err = fmt.Errorf("%v\n" +
						"Error returned by BigIntNum{}.NewBigIntBigPrecision()\n" +
						"Error: %v",
						ePrefix,
						errX.Error())

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
				}


				if cycleNum == 7 {

					fmt.Println()
					fmt.Println("=================================================")
					fmt.Println("             aCycleValue: ", aCycleValue.GetNumStr())
					fmt.Println("    aCycleValuePrecision: ", aCycleValue.GetPrecision())
					fmt.Println("             gCycleValue: ", gCycleValue.GetNumStr())
					fmt.Println("    gCycleValuePrecision: ", gCycleValue.GetPrecision())
					fmt.Println("-------------------------------------------------")
					fmt.Println("  tACyclePreAddPrecision: ", tempACyclePreAddPrecision.Text(10))
					fmt.Println("     tempACyclePrecision: ", tempACyclePrecision.Text(10))
					fmt.Println("     tempGCyclePrecision: ", tempGCyclePrecision.Text(10))
					fmt.Println("       tempGComPrecision: ", tempGComPrecision.Text(10))
					fmt.Println("       gComPreRoundValue: ", gComPreRoundValue.GetNumStr())
					fmt.Println("gComPreRoundValPrecision: ", gComPreRoundValue.GetPrecision())
					fmt.Println("-------------------------------------------------")
					fmt.Println("               gComValue: ", gComValue.GetNumStr())
					fmt.Println("      gComValuePrecision: ", gComValue.GetPrecision())
					fmt.Println("-------------------------------------------------")
					fmt.Println("            aCycleResult: ", aCycleResult.GetNumStr())
					fmt.Println("  aCycleResult Precision: ", aCycleResult.GetPrecision())
					fmt.Println("-------------------------------------------------")
					fmt.Println("            gCycleResult: ", gCycleResult.GetNumStr())
					fmt.Println("  gCycleResult Precision: ", gCycleResult.GetPrecision())
					fmt.Println("=================================================")
					fmt.Println()
					return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err

				}
		*/

		/*
			aCycleValue.RoundToDecPlace(32)
			gCycleValue.RoundToDecPlace(32)
			aCycleResult.RoundToDecPlace(32)
			gCycleResult.RoundToDecPlace(32)
			fmt.Println()
			fmt.Println("=================================================")
			fmt.Println("             cycle No: ", cycleNum)
			fmt.Println("          aCycleValue: ", aCycleValue.GetNumStr())
			fmt.Println("      aCyclePrecision: ", aCycleValue.GetPrecision())
			fmt.Println("          gCycleValue: ", gCycleValue.GetNumStr())
			fmt.Println("      gCyclePrecision: ", gCycleValue.GetPrecision())
			fmt.Println("         aCycleResult: ", aCycleResult.GetNumStr())
			fmt.Println("aCycleResultPrecision: ", aCycleResult.GetPrecision())
			fmt.Println("         gCycleResult: ", gCycleResult.GetNumStr())
			fmt.Println("gCycleResultPrecision: ", gCycleResult.GetPrecision())
			fmt.Println("=================================================")
			fmt.Println()

		*/

		aTest, aTestPrecision, err =
			bIMathElectron.truncateToMaxPrecisionBigInt(
				a,
				aPrecision,
				targetPrecision,
				ePrefix)

		if err != nil {

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "aTest, aTestPrecision, err =\n" +
						"    bIMathElectron.truncateToMaxPrecisionBigInt(\n" +
						"    a, aPrecision, targetPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("a= '%v'; aPrecision= '%v'; targetPrecision= '%v'",
						a.Text(10), aPrecision.Text(10), targetPrecision.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		gTest, gTestPrecision, err =
			bIMathElectron.truncateToMaxPrecisionBigInt(g, gPrecision, targetPrecision, ePrefix)

		if err != nil {

			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "gTest, gTestPrecision, err =\n" +
						"    bIMathElectron.truncateToMaxPrecisionBigInt(\n" +
						"    g, gPrecision, targetPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("g= '%v'; gPrecision= '%v'; targetPrecision= '%v'",
						g.Text(10), gPrecision.Text(10), targetPrecision.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		if aTest.Cmp(gTest) == 0 &&
			aTestPrecision.Cmp(gTestPrecision) == 0 {

			agCmp, err =
				bIMathElectron.precisionCmpBigInt(
					gTest,
					gTestPrecision,
					origG,
					origGPrecision,
					ePrefix)

			if err != nil {

				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
					&FuncReturnError{
						ErrPrefix: ePrefix.String(),
						ReturnFunc: "agCmp, err =\n" +
							"    bIMathElectron.precisionCmpBigInt(\n" +
							"    gTest, gTestPrecision, origG, origGPrecision, ePrefix)",
						ErrContext: fmt.Sprintf("gTest= '%v'; gTestPrecision= '%v'\n"+
							"origG= '%v'; origGPrecision= '%v'",
							gTest.Text(10), gTestPrecision.Text(10), origG.Text(10), origGPrecision.Text(10)),
						ErrMessage: err.Error(),
					}
			}

			if agCmp == 1 {

				cycles = cycleNum

				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "",
						ErrContext: fmt.Sprintf("Result= '%v'; ResultPrecision= '%v'\n"+
							"gNum='%v' gNumPrecision='%v'",
							gTest.Text(10), gTestPrecision.Text(10), origG.Text(10), origGPrecision.Text(10)),
						ErrMessage: "Computation Failure: Result is greater than largest test value!",
					}
			}

			agCmp, err =
				bIMathElectron.precisionCmpBigInt(
					aTest,
					aTestPrecision,
					origA,
					origAPrecision,
					ePrefix)

			if err != nil {

				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
					&FuncReturnError{
						ErrPrefix: ePrefix.String(),
						ReturnFunc: "agCmp, err =\n        bIMathElect\n" +
							"    bIMathElectron.precisionCmpBigInt(\n" +
							"    aTest, aTestPrecision, origA, origAPrecision, ePrefix)",
						ErrContext: fmt.Sprintf("aTest= '%v'; aTestPrecision= '%v'\n"+
							"origA= '%v'; origAPrecision= '%v'",
							aTest.Text(10), aTestPrecision.Text(10), origA.Text(10), origAPrecision.Text(10)),
						ErrMessage: err.Error(),
					}
			}

			if agCmp == -1 {

				cycles = cycleNum

				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
					&FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "",
						ErrContext: fmt.Sprintf("Result= '%v'; ResultPrecision= '%v'\n"+
							"aNum='%v' aNumPrecision='%v'",
							gTest.Text(10), gTestPrecision.Text(10), origA.Text(10), origAPrecision.Text(10)),
						ErrMessage: "Computation Failure: Result is less than smallest test value!",
					}
			}

			agMean.Set(aTest)
			agMeanPrecision.Set(aTestPrecision)
			gValue.Set(gTest)
			gValuePrecision.Set(gTestPrecision)
			cycles = cycleNum

			// Successful method completion exits here!
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, nil
		}

	}

	err = fmt.Errorf("%v\n"+
		"Error: No Arithmetic Geometric Mean Computed.\n",
		ePrefix)

	agMean.Set(a)
	agMeanPrecision.Set(aPrecision)
	gValue.Set(g)
	gValuePrecision.Set(gPrecision)
	cycles = cycleLimit

	// Error Exit Here! No Arithmetic Geometric Mean Computed.
	return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
}
