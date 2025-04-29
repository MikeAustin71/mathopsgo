package mathops

import (
	"fmt"
	"math/big"
)

// FracIntAry - A fraction represented by a numerator and a denominator.
// Both numerator and denominator are of type intAry
type FracIntAry struct {
	Numerator   IntAry
	Denominator IntAry
}

// NewBigInts - Creates a new FracIntAry type from two *big.Int types passed
// as input parameters.
func (fIa *FracIntAry) NewBigInts(numerator, denominator *big.Int) (FracIntAry, error) {
	ePrefix := "FracIntAry.NewBigInts() "

	iaNumerator, err := IntAry{}.NewBigInt(numerator, 0)

	if err != nil {
		return FracIntAry{},
			fmt.Errorf(ePrefix+"- Error returned by IntAry{}.NewBigInt(numerator, 0) "+
				"Error='%v' ", err)
	}

	iaDenominator, err := IntAry{}.NewBigInt(denominator, 0)

	if err != nil {
		return FracIntAry{},
			fmt.Errorf(ePrefix+"- Error returned by IntAry{}.NewBigInt(denominator, 0) "+
				"Error='%v' ", err)
	}

	newFracIntAry := new(FracIntAry).NewIntArys(&iaNumerator, &iaDenominator)

	return newFracIntAry, nil
}

// NewNumStrs - Creates a new FracIntAry type by passing input parameters numerator and
// denominator as number strings.
func (fIa *FracIntAry) NewNumStrs(numerator, denominator string) (FracIntAry, error) {

	var err error

	ePrefix := "FracIntAry.NewNumStrs() "
	fIa2 := FracIntAry{}

	fIa2.Numerator, err = IntAry{}.NewNumStr(numerator)

	if err != nil {
		return FracIntAry{}, fmt.Errorf(ePrefix+
			"- Error returned from intAry{}.NewNumStr(numerator). Error= %v", err)
	}

	fIa2.Denominator, err = IntAry{}.NewNumStr(denominator)

	if err != nil {
		return FracIntAry{},
			fmt.Errorf(ePrefix+
				"- Error returned from intAry{}.NewNumStr(denominator). Error= %v", err)
	}

	return fIa2, nil
}

// NewIntArys - Creates a type FracIntAry by passing numerator and denominator
// input parameters of type *intAry
func (fIa *FracIntAry) NewIntArys(numerator, denominator *IntAry) FracIntAry {

	fIa2 := FracIntAry{}

	fIa2.Numerator = numerator.CopyOut()
	fIa2.Denominator = denominator.CopyOut()

	return fIa2
}

// NewFracIntAry - Creates a FracIntAry instance from a single IntAry object.
// The IntAry input parameter is converted into an equivalent fraction.
func (fIa *FracIntAry) NewFracIntAry(ia *IntAry) FracIntAry {

	fIa2 := FracIntAry{}

	if ia.GetPrecision() == 0 {

		fIa2.Numerator = ia.CopyOut()
		fIa2.Denominator = IntAry{}.NewOne(0)

		return fIa2
	}

	precision := ia.GetPrecision()
	fIa2.Numerator = ia.CopyOut()

	if precision > 0 {
		fIa2.Numerator.ShiftPrecisionRight(uint(precision))
	}

	fIa2.Denominator = IntAry{}.NewOne(0)
	IntAryMathMultiply{}.MultiplyByTenToPower(&fIa2.Denominator, uint(precision))

	return fIa2
}

// CopyOut - Creates and returns a copy of the current
// FracIntAry.
func (fIa *FracIntAry) CopyOut() FracIntAry {

	newFrac := FracIntAry{}

	newFrac.Numerator = fIa.Numerator.CopyOut()
	newFrac.Denominator = fIa.Denominator.CopyOut()

	return newFrac
}

// CopyIn - Receives a pointer to an incoming FracIntAry and copies
// the values into the current FracIntAry.
func (fIa *FracIntAry) CopyIn(fIa2 *FracIntAry) {

	fIa.Numerator = fIa2.Numerator.CopyOut()

	fIa.Denominator = fIa2.Denominator.CopyOut()

}

// GetRationalValue - Converts the fraction and returns the value as a
// big rational number (*big.Rat).
//
// Input parameter maxPrecision determines the maximum number of decimal
// places to the right of the decimal point contained in the result.
//
// If the value of maxPrecision is -1, maximum precision will default to
// 4096 decimal places. maxPrecision values less than -1 will trigger an
// error.
func (fIa *FracIntAry) GetRationalValue(maxPrecision int) (*big.Rat, error) {

	if maxPrecision < -1 {
		return big.NewRat(1, 1), fmt.Errorf("GetRationalValue() - maxPrecision is less than -1 and therefore INVALID. maxPrecision= %v", maxPrecision)
	}

	if maxPrecision == -1 {
		maxPrecision = 4096
	}

	if fIa.Numerator.GetPrecision() == 0 && fIa.Denominator.GetPrecision() == 0 {

		fRat, ok := big.NewRat(1, 1).SetString(fIa.Numerator.GetNumStr() + "/" + fIa.Denominator.GetNumStr())

		if !ok {

			return big.NewRat(1, 1),
				fmt.Errorf("GetRationalValue()\n" +
					"Method FAILED! :\n" +
					"  big.NewFloat(0).SetString(big.NewRat(1, 1).SetString(fIa.Numerator.GetNumStr() + \"/\" + fIa.Denominator.GetNumStr())\n")

		}

		return fRat, nil

	}

	newFloat, err := fIa.Numerator.DivideThisBy(&fIa.Denominator, 0, maxPrecision)

	if err != nil {
		return big.NewRat(1, 1), fmt.Errorf("GetRationalValue() - Error returned from fIa.Numerator.DivideThisBy(&fIa.Denominator, 42). Error= %v", err)
	}

	fRat, ok := big.NewRat(1, 1).SetString(newFloat.GetNumStr())

	if !ok {
		return big.NewRat(1, 1),
			fmt.Errorf("GetRationalValue()\n" +
				"Method big.NewFloat(0).SetString(fracIa.GetNumStr()) Failed!\n")
	}

	return fRat, nil

}

// GetLowestCommonDenom - Returns a FracIntAry which represents the lowest common
// denominator for the current FracIntAry.
//
// Note: if 'maxPrecision' is less than 0, it is automatically converted to '4,096'
// decimal places.
func (fIa *FracIntAry) GetLowestCommonDenom(maxPrecision int) (FracIntAry, error) {

	if maxPrecision < 0 {
		maxPrecision = 4096
	}

	ePrefix := "FracIntAry.GetLowestCommonDenom() "

	ratFrac, err := fIa.GetRationalValue(maxPrecision)

	if err != nil {
		return FracIntAry{},
			fmt.Errorf(ePrefix +
				"Error returned by fIa.GetRationalValue(4096) ")
	}

	newFAry, err := fIa.NewBigInts(ratFrac.Num(), ratFrac.Denom())

	if err != nil {
		return FracIntAry{},
			fmt.Errorf(ePrefix +
				"Error returned by fIa.NewBigInts(ratFrac.Num(), ratFrac.Denom()) ")
	}

	return newFAry, nil
}

// ReduceToLowestCommonDenom - Converts the value of the current FracIntAry
// to its lowest common denominator.
//
// Note: if 'maxPrecision' is less than 0, it is automatically converted to '4,096'
// decimal places.
func (fIa *FracIntAry) ReduceToLowestCommonDenom(maxPrecision int) error {

	if maxPrecision < 0 {
		maxPrecision = 4096
	}

	fIaLCD, err := fIa.GetLowestCommonDenom(maxPrecision)

	if err != nil {
		ePrefix := "FracIntAry.ReduceToLowestCommonDenom() "
		return fmt.Errorf(ePrefix+
			"Error returned by fIa.GetLowestCommonDenom(maxPrecision). "+
			"Error='%v' ", err.Error())
	}

	fIa.CopyIn(&fIaLCD)

	return nil
}
