package mathops

import (
	"math/big"
	"fmt"
)

type IntAryMathPower struct {
	Input  IntAryPair
	Result IntAry
}

func (iaPwr IntAryMathPower) Pwr(
					base, exponent *IntAry,
							minResultPrecision, maxResultPrecision int) error {

	ePrefix := "IntAryMathPower.Pwr() "

	if exponent.GetPrecision() == 0 {
		// This is an integer Exponent!
		bInt, err := exponent.GetBigInt()

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error returned by exponent.GetBigInt() Error='%v'\n",
				err.Error())
		}

		err = iaPwr.PwrByTwos(base, bInt, maxResultPrecision, maxResultPrecision + 10)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error returned by iaPwr.PwrByTwos(...) Error='%v' ", err.Error())
		}

		if base.GetPrecision() < minResultPrecision {
			base.SetPrecision(minResultPrecision, false)
		}

		return nil
	}

	// exponent must be a fractional exponent
	// TODO - Finish fractional exponent
	return nil
}

// pwrByTwos - Raises a *big.Int 'base', to the specified 'power'
// using the Exponentiation by squaring algorithm.
//
// See:
// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
// https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Computation_by_powers_of_2
// This method is based on revised code taken in part from Ye Lin Aung.
// https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
// This algorithm modified by Mike Rapp to achieve improved performance.
//
// Input Parameters:
// =================
//
// 'ia' *IntAry -
//			The base which will be raised to an exponent specified by input
//			parameter 'power'
//
// 'power' *big.Int -
// 				The input parameter 'power' may be either
// 				a positive or negative integer.
//
// 'maxResultPrecision' int -
//				'maxResultPrecision' will determine the maximum
// 				number of digits to the right of the decimal
//				place in the result.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point in the result.
//
//	'internalPrecision' int -
// 				'internalPrecision' will control the number of digits of
//				accuracy to the right of the decimal point maintained by
//				internal multiplication operations used in raising the intAry
//				value to the designated power.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point during internal multiplication operations.
//
func (iaPwr IntAryMathPower) PwrByTwos(
					ia *IntAry,
						power *big.Int,
							maxResultPrecision,
								internalPrecision int ) error {

  ePrefix := "IntAryMathPower.PwrByTwos() "

	if maxResultPrecision < -1 {
		return fmt.Errorf(ePrefix +
			"Error: Parameter maxResultPrecision is less than -1. " +
			"maxResultPrecision= %v", maxResultPrecision)
	}

	if internalPrecision < -1 {
		return fmt.Errorf(ePrefix +
			"Error: Parameter internalPrecision is less than -1. " +
			"internalPrecision= %v", internalPrecision)
	}

	ia.SetInternalFlags()
	tPower := big.NewInt(0).Set(power)
	one := big.NewInt(1)
	zero := big.NewInt(0)
	two := big.NewInt(2)

	if ia.isZeroValue {
		ia.SetIntAryToZero(0)
		return nil
	}

	if tPower.Cmp(two) == -1 {

		if tPower.Cmp(zero) == -1 {

			ia2, err := ia.Inverse(internalPrecision)

			if err != nil {
				return fmt.Errorf(ePrefix +
					"Error returned from ia.Inverse(-1): Error= %v", err.Error())
			}

			ia.CopyIn(&ia2, false)
			tPower = big.NewInt(0).Mul(tPower, big.NewInt(-1))

		} else if tPower.Cmp(one) == 0 {
			// no change in value. x^1 == x
			return nil
		} else if tPower.Cmp(zero) == 0 {
			ia.SetIntAryToOne(0)
			return nil
		}

	}

	tBase := ia.CopyOut()

	ia.SetIntAryToOne(0)

	for tPower.Cmp(zero) == 1 {
		//temp, _:= intAry{}.NewNumStr("0")

		if big.NewInt(0).Mod(tPower, two).Cmp(one) == 0 {
			//temp = big.NewInt(0).Mul(result, tBase)
			//result = big.NewInt(0).Set(temp)
			err := ia.MultiplyThisBy(&tBase, -1, internalPrecision)
			//fmt.Println("ia precision = ", ia.GetPrecision())

			if err != nil {
				return fmt.Errorf(ePrefix +
					"Error From result.MultiplyThisBy(&tBase, true). Error= %v", err.Error())
			}

			if tPower.Cmp(one) == 0 {
				if maxResultPrecision > -1 && maxResultPrecision < ia.GetPrecision() {
					ia.SetPrecision(maxResultPrecision, true)
				}

				return nil
			}
		}

		err:= tBase.MultiplyThisBy(&tBase, -1,  internalPrecision)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error From tBase.MultiplyThisBy(&temp, true). Error= %v", err.Error())
		}


		tPower = big.NewInt(0).Div(tPower, two)
	}

	if maxResultPrecision == -1 {
		return nil
	}

	if maxResultPrecision > -1 && maxResultPrecision < ia.GetPrecision()  {
		ia.SetPrecision(maxResultPrecision, true)
	}

	return nil
}