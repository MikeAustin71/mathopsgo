package mathops

import (
	"errors"
	"fmt"
	"math/big"
)

type BigIntMath struct {
	Input  *big.Int
	Output *big.Int
}


// ArithmeticGeometricMean - Computes the Arithmetic-Geometric Mean of
// two numbers. See:
// 		https://en.wikipedia.org/wiki/Arithmetic–geometric_mean
//
// Dev Note: maxInternalPrecision may need to be 60 to 75 times targetPrecision.
func (bIntMath BigIntMath) ArithmeticGeometricMean(
	aNum,
	aNumPrecision,
	gNum,
	gNumPrecision,
	maxInternalPrecision,
	targetPrecision *big.Int) (agMean,
														agMeanPrecision,
														gValue,
														gValuePrecision *big.Int,
														cycles uint64,
														err error) {

	ePrefix := "BigIntMath.ArithmeticGeometricMean() "
	agMean = big.NewInt(0)
	agMeanPrecision = big.NewInt(0)
	gValue= big.NewInt(0)
	gValuePrecision = big.NewInt(0)
	cycles = 0
	err =  nil

	if aNum == nil {
			err = errors.New(ePrefix +
				"Error: Input parameter 'aNum' is nil and INVALID!")
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if aNumPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'aNumPrecision' is nil and INVALID!")
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if gNum == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'gNum' is nil and INVALID!")
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if gNumPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'gNumPrecision' is nil and INVALID!")
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if maxInternalPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'maxInternalPrecision' is nil and INVALID!")
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if targetPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'targetPrecision' is nil and INVALID!")
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	bigZero := big.NewInt(0)

	if aNumPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'aNumPrecision' is less than zero and INVALID! " +
			"aNumPrecision='%v' ", aNumPrecision.Text(10))
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if gNumPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'gNumPrecision' is less than zero and INVALID! " +
			"gNumPrecision='%v' ", gNumPrecision.Text(10))
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if maxInternalPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'maxInternalPrecision' is less than zero and INVALID! " +
			"maxInternalPrecision='%v' ", maxInternalPrecision.Text(10))
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	if targetPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'targetPrecision' is less than zero and INVALID! " +
			"targetPrecision='%v' ", targetPrecision.Text(10))
		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
	}

	minInternalPrecision := big.NewInt(0).Add(targetPrecision, big.NewInt(120))

	if maxInternalPrecision.Cmp(minInternalPrecision) == -1 {
		maxInternalPrecision.Set(minInternalPrecision)
	}

	cycleLimit := maxInternalPrecision.Uint64() + uint64(200)

	var errX error

	// Assume aNum is less than gNum

	a := big.NewInt(0).Set(aNum)
	aPrecision := big.NewInt(0).Set(aNumPrecision)
	g := big.NewInt(0).Set(gNum)
	gPrecision := big.NewInt(0).Set(gNumPrecision)

	agCmp :=
		BigIntMath{}.BigIntPrecisionCmp(
			aNum,
			aNumPrecision,
			gNum,
			gNumPrecision)

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

	fdNRoot := FixedDecimalNthRoot{}


	for i:= uint64(0); i < cycleLimit; i++ {

		aCom, aComPrecision, errX = BigIntMathAdd{}.BigIntAdd(a, aPrecision, g, gPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v",errX.Error())
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

		gCom, gComPrecision, errX = BigIntMathMultiply{}.BigIntMultiply(a, aPrecision, g, gPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v",errX.Error())
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

		a, aPrecision, errX = BigIntMathDivide{}.BigIntFracQuotient(aCom, aComPrecision, factor, factorPrecision, maxInternalPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v",errX.Error())
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

		g, gPrecision, errX =
				fdNRoot.CalculatePositiveIntegerNthRoot(
					gCom,
					gComPrecision,
					factor,
					factorPrecision,
					maxInternalPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v",errX.Error())
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

		aTest, aTestPrecision, errX =
		 	BigIntMath{}.TruncateToMaxPrecision(a, aPrecision, targetPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v",errX.Error())
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

		gTest, gTestPrecision, errX =
			BigIntMath{}.TruncateToMaxPrecision(g, gPrecision, targetPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v",errX.Error())
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

		if aTest.Cmp(gTest) == 0 &&
			aTestPrecision.Cmp(gTestPrecision) == 0 {

			agCmp =
				BigIntMath{}.BigIntPrecisionCmp(
					gTest,
					gTestPrecision,
					origG,
					origGPrecision)

			if agCmp==1 {

				err = fmt.Errorf(ePrefix +
					"Computation Failure: Result is greater than largest test value! " +
					"Result='%v' ResultPrecision='%v' gNum='%v' gNumPrecision='%v'",
					gTest.Text(10), gTestPrecision.Text(10),
					origG.Text(10), origGPrecision.Text(10))
				cycles = i + 1
				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
			}

			agCmp =
				BigIntMath{}.BigIntPrecisionCmp(
					aTest,
					aTestPrecision,
					origA,
					origAPrecision)

			if agCmp == -1 {
				err = fmt.Errorf(ePrefix +
					"Computation Failure: Result is less than smallest test value! " +
					"Result='%v' ResultPrecision='%v' aNum='%v' aNumPrecision='%v'",
					gTest.Text(10), gTestPrecision.Text(10),
					origA.Text(10), origAPrecision.Text(10))
				cycles = i + 1
				return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err

			}

			agMean.Set(aTest)
			agMeanPrecision.Set(aTestPrecision)
			gValue.Set(gTest)
			gValuePrecision.Set(gTestPrecision)
			cycles = i + 1
			err = nil
			return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
		}

	}

	err = errors.New(ePrefix +
		"Error: No Arithmetic Geometric Mean Computed. ")

	agMean.Set(a)
	agMeanPrecision.Set(aPrecision)
	gValue.Set(g)
	gValuePrecision.Set(gPrecision)
	cycles = cycleLimit

	return agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err
}

// BigIntPrecisionCmp - Compares two *big.Int number pairs and adjusts the
// comparison for precision.
//
// Return Values
// num 1 <  num2 == -1
// num 1 == num2 == 0
// num 1 >  num2 == 1
//
func (bIntMath BigIntMath)BigIntPrecisionCmp(
	num1,
	num1Precision,
	num2,
	num2Precision *big.Int) int {

	bigZero := big.NewInt(0)

	if num1 == nil {
		num1 = big.NewInt(0)
		num1Precision = big.NewInt(0)
	}

	if num2 == nil {
		num2 = big.NewInt(0)
		num2Precision = big.NewInt(0)
	}

	if num1Precision.Cmp(bigZero) == -1 {
		num1Precision = big.NewInt(0)

	}

	if num2Precision.Cmp(bigZero) == -1 {
		num2Precision = big.NewInt(0)
	}

	if num1Precision.Cmp(num2Precision) == 0 {
		return num1.Cmp(num2)
	}

	bigTen := big.NewInt(10)
	tNum1 := big.NewInt(0).Set(num1)
	tNum2 := big.NewInt(0).Set(num2)

	if num1Precision.Cmp(num2Precision) == 1 {
		delta := big.NewInt(0).Sub(num1Precision, num2Precision)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		tNum2.Mul(tNum2,scale)
		return tNum1.Cmp(tNum2)

	}

	// MUST BE num2Precision > num1Precision
	delta := big.NewInt(0).Sub(num2Precision, num1Precision)
	scale := big.NewInt(0).Exp(bigTen, delta, nil)
	tNum1.Mul(tNum1, scale)

	return tNum1.Cmp(tNum2)
}

// GetMagnitude - Returns the magnitude of a *big.Int number passed as an input
// parameter ('initialValue'. Magnitude is defined here as the power of 10 which generates a value
// less than or equal to the 'target' *big.Int number.
//
// 													10^magnitude  <= initialValue
//
// The value of magnitude is returned as a *big.Int number to the calling function.
//
//
// Examples:
// =========
//
//			   			target				magnitude
//              ------        ---------
//
//			  		 963,256						5
//									 2						0
//									32						1
// 			 8,456,123,921					  9
//
//
// Input Parameters
// ================
//
// initialValue	*big.Int 	- An integer of type *big.Int. This method will analyze this
//                         	integer and return it's magnitude.
//
// Return Values
// =============
//
// magnitude 		*big.Int	- 10 raised to the power of magnitude will yield a value which
//                          is less than or equal to the input parameter 'initialValue'.
//
// err					error			- If during the completion of this calculation an error is encountered,
//           								return value 'magnitude' will be set to zero and this error object
//                          will be populated with an appropriate error message. If the method
//                          completes successfully, this return value, 'err' will be set to
//                          'nil'.
//
func (bIntMath BigIntMath) GetMagnitude(initialValue *big.Int) (magnitude *big.Int, err error) {

	ePrefix := "BigIntMath) GetMagnitude() "

	magnitude = big.NewInt(0)
	err = nil

	if initialValue == nil {
		err = errors.New(ePrefix + "Error: 'initialValue' is NOT Initialized! initialValue=nil")
		return magnitude, err
	}

	target := big.NewInt(0).Set(initialValue)
	compareResult := target.Cmp(big.NewInt(0))

	// Convert to absolute value
	if compareResult == -1 {
		target.Neg(target)
	}

	if compareResult == 0 {
		// magnitude = 0; err=nil
		return magnitude, err
	}

	bigTen := big.NewInt(10)

	if target.Cmp(bigTen) == -1 {
		// magnitude = 0; err=nil
		return magnitude, err
	}

	bitLen := target.BitLen()

	if bitLen <= 0 {
		err = fmt.Errorf(ePrefix+"Error: target.BitLen() = %v - Negative value!",
			bitLen)
		return magnitude, err
	}

	tenToPowerPrecision := big.NewInt(0)
	bigZero := big.NewInt(0)

	// Note: log10of2To20k is a global constant

	var errx error

	// ************************************
	// target MUST BE <= 2^(bit length)
	// ************************************
	magnitude, tenToPowerPrecision, errx =
		BigIntMathMultiply{}.BigIntMultiply(
			big.NewInt(int64(bitLen)),
			big.NewInt(0),
			log10of2To20k.GetInteger(),
			log10of2To20k.GetPrecisionBigInt())

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"%v", errx.Error())
		magnitude = big.NewInt(0)
		return magnitude, err
	}

	if tenToPowerPrecision.Cmp(bigZero) == 1 {
		scale :=
			big.NewInt(0).Exp(
				bigTen,
				tenToPowerPrecision,
				nil)

		magnitude.Quo(magnitude, scale)
	}

	testNum := big.NewInt(0).Exp(bigTen, magnitude, nil)

	if testNum.Cmp(target) == 1 {
		magnitude.Sub(magnitude, big.NewInt(1))
	}

	return magnitude, err
}


// RoundToMaxPrecision - Applies maximum precision to a *big.Int number
// and associated numeric precision, 'bigIntNum' and 'bigIntNumPrecision'.
//
// Precision as used here defines the number of digits to the right of the
// decimal place. If 'bigIntNumPrecision' exceeds 'maxPrecision', the
// 'bigIntNum' and 'bigIntNumPrecision' pair are rounded to 'maxPrecision'
// and returned as 'result' and 'resultPrecision'.
//
// Note that if 'bigIntNumPrecision' exceeds 'maxPrecision', the returned
// value will be rounded to 'maxPrecision' decimal digits to the right of
// the decimal place.
//
// Examples:
// =========
//
//  bigIntNum	bigIntNumPrecision	maxPrecision	result		resultPrecision
//	5255						3                  	2					526					2
//  52671						4										6					52671				4
//
func (bIntMath BigIntMath) RoundToMaxPrecision(
	bigIntNum,
	bigIntNumPrecision,
	maxPrecision *big.Int) (result, resultPrecision *big.Int, err error) {

	ePrefix := "BigIntMath.RoundToMaxPrecision() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	if bigIntNum == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'bigIntNum' is nil. INVALID!")
		return result, resultPrecision, err
	}

	if bigIntNumPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'bigIntNumPrecision' is nil. INVALID!")
		return result, resultPrecision, err
	}

	if maxPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'maxPrecision' is nil. INVALID!")
		return result, resultPrecision, err
	}

	bigZero := big.NewInt(0)

	if bigIntNum.Cmp(bigZero) == 0 {
		return result, resultPrecision, err
	}

	if bigIntNumPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'bigIntNumPrecision' is LESS THAN ZERO! " +
			"bigIntNumPrecision='%v' ", bigIntNumPrecision.Text(10))
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'maxPrecision' is LESS THAN ZERO! " +
			"maxPrecision='%v' ", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	result = big.NewInt(0).Set(bigIntNum)
	resultPrecision = big.NewInt(0).Set(bigIntNumPrecision)

	if resultPrecision.Cmp(maxPrecision) == 1 {
		delta := big.NewInt(0).Sub(resultPrecision, maxPrecision)
		delta.Sub(delta, big.NewInt(1))
		bigTen := big.NewInt(10)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		result.Quo(result, scale)
		bigFive := big.NewInt(5)
		if result.Cmp(bigZero) == -1 {
			bigFive.Neg(bigFive)
		}
		result.Add(result, bigFive)
		result.Quo(result, bigTen)
		resultPrecision = big.NewInt(0).Set(maxPrecision)
	}

	err = nil

	return result, resultPrecision, err
}

// TruncateToMaxPrecision - Applies maximum precision to a *big.Int number
// and associated numeric precision, 'bigIntNum' and 'bigIntNumPrecision'.
//
// Precision as used here defines the number of digits to the right of the
// decimal place. If 'bigIntNumPrecision' exceeds 'maxPrecision', the
// 'bigIntNum' and 'bigIntNumPrecision' pair are rounded to 'maxPrecision'
// and returned as 'result' and 'resultPrecision'.
//
// Note that if 'bigIntNumPrecision' exceeds 'maxPrecision', the returned
// value will be truncated to (not rounded to) 'maxPrecision' decimal digits
// to the right of the decimal place.
//
// Examples:
// =========
//
//  bigIntNum	bigIntNumPrecision	maxPrecision	result		resultPrecision
//	5255						3                  	2					525					2
//  52671						4										6					52671				4
//
func (bIntMath BigIntMath) TruncateToMaxPrecision(
	bigIntNum,
	bigIntNumPrecision,
	maxPrecision *big.Int) (result, resultPrecision *big.Int, err error) {

	ePrefix := "BigIntMath.RoundToMaxPrecision() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	if bigIntNum == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'bigIntNum' is nil. INVALID!")
		return result, resultPrecision, err
	}

	if bigIntNumPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'bigIntNumPrecision' is nil. INVALID!")
		return result, resultPrecision, err
	}

	if maxPrecision == nil {
		err = errors.New(ePrefix +
			"Error: Input parameter 'maxPrecision' is nil. INVALID!")
		return result, resultPrecision, err
	}

	bigZero := big.NewInt(0)

	if bigIntNum.Cmp(bigZero) == 0 {
		return result, resultPrecision, err
	}

	if bigIntNumPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'bigIntNumPrecision' is LESS THAN ZERO! " +
			"bigIntNumPrecision='%v' ", bigIntNumPrecision.Text(10))
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'maxPrecision' is LESS THAN ZERO! " +
			"maxPrecision='%v' ", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	result = big.NewInt(0).Set(bigIntNum)
	resultPrecision = big.NewInt(0).Set(bigIntNumPrecision)

	if resultPrecision.Cmp(maxPrecision) == 1 {
		delta := big.NewInt(0).Sub(resultPrecision, maxPrecision)
		bigTen := big.NewInt(10)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		result.Quo(result, scale)
		resultPrecision = big.NewInt(0).Set(maxPrecision)
	}

	err = nil

	return result, resultPrecision, err
}
