package mathops

import (
	"math/big"
	"fmt"
)

// BigIntNum - wraps a *big.Int integer and its associated
// precision and Sign Value.
//
type BigIntNum struct {
	BigInt			*big.Int
	AbsBigInt		*big.Int
	Precision 	uint			// Number of digits to the right of the decimal point.
	ScaleFactor *big.Int	// Scale Factor =  10^precision
	Sign				int				// Valid values are -1 or +1. Indicates the sign of the
												// 	the 'BigInt' integer.
}


// New - returns a new BigIntNum instance initialized to
// zero.
//
func (bNum BigIntNum) New() BigIntNum {
	b := BigIntNum{}
	b.Empty()
	return b
}

// NewBigInt - Creates a new BigIntNum instance using a *big.Int type and its
// associated precision.
//
func (bNum BigIntNum) NewBigInt(bigI *big.Int, precision uint) BigIntNum {
	b := BigIntNum{}
	b.SetBigIntNum(bigI, precision)
	return b
}

// NewDecimal - Receives a 'Decimal' type as input and returns a BigIntNum.
//
func (bNum BigIntNum) NewDecimal(decNum Decimal) (BigIntNum, error) {
	ePrefix := "BigIntNum.NewIntAry() "

	err := decNum.IsDecimalValid()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error: Input Parameter 'decNum' is INVALID!. Error returned by " +
				"decNum.IsDecimalValid(). Error='%v'", err.Error())
	}

	bInt, err := decNum.GetSignedBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by decNum.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	precision := uint(decNum.GetPrecision())

	b := BigIntNum{}
	b.SetBigIntNum(bInt, precision)
	return b, nil
}

// NewIntAry - Creates a new BigIntNum instance from an input parameter
// IntAry.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bNum BigIntNum) NewIntAry(ia IntAry) (BigIntNum, error) {
	ePrefix := "BigIntNum.NewIntAry() "

	err := ia.IsIntAryValid("")

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error: Input Parameter 'ia' is INVALID!. Error returned by " +
			"ia.IsIntAryValid(\"\"). Error='%v'", err.Error())
	}

	bInt := ia.GetBigInt()
	precision := uint(ia.GetPrecision())

	b := BigIntNum{}
	b.SetBigIntNum(bInt, precision)
	return b, nil
}

// NewNumStr - Receives a number string as input and returns
// a new BigIntNum instance.
//
func (bNum BigIntNum) NewNumStr(numStr string) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStr() "

	nDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	b := BigIntNum{}

	b.SetBigIntNum(bigI, nDto.GetPrecision())

	return b, nil
}

// NewNumStrDto - Receives a NumStrDto instance as input and returns
// a new BigIntNum instance.
//
func (bNum BigIntNum) NewNumStrDto(nDto NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStrDto() "

	err := nDto.IsNumStrDtoValid("")

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error returned from nDto.IsNumStrDtoValid(\"\"). " +
			"NumStr='%v' Error='%v'", nDto.GetNumStr(), err.Error())
	}

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	b := BigIntNum{}

	b.SetBigIntNum(bigI, nDto.GetPrecision())

	return b, nil
}

// Empty - Resets the BigIntNum data fields to their
// uninitialized or zero state.
//
func (bNum *BigIntNum) Empty() {

	bNum.BigInt = big.NewInt(0)
	bNum.AbsBigInt = big.NewInt(0)
	bNum.ScaleFactor = big.NewInt(1)
	bNum.Sign = 1
	bNum.Precision = 0

}

// CopyIn - Receives an incoming BigIntNum type and
// copies the value into the current BigIntNum instance.
//
func (bNum *BigIntNum) CopyIn(bigN BigIntNum) {

	bNum.BigInt = big.NewInt(0).Set(bigN.BigInt)
	bNum.AbsBigInt = big.NewInt(0).Set(bigN.AbsBigInt)
	bNum.Precision = bigN.Precision
	bNum.ScaleFactor = big.NewInt(0).Set(bigN.ScaleFactor)
	bNum.Sign = bigN.Sign
}

// CopyOut - Makes a deep copy of the current BigIntNum instance
// and returns it as a new BigIntNum instance.
//
func (bNum *BigIntNum) CopyOut() BigIntNum {

	b2 := BigIntNum{}.NewBigInt(big.NewInt(0).Set(bNum.BigInt), bNum.Precision)

	return b2
}

// SetBigIntNum - Sets the value of the current BigIntNum instance using
// the input parameters *big.Int integer and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal point.
//
func (bNum *BigIntNum) SetBigIntNum(bigI *big.Int, precision uint) {

	bNum.BigInt = big.NewInt(0).Set(bigI)
	bNum.Precision = precision
	base10 := big.NewInt(0).SetInt64(int64(10))
	bigPrecision := big.NewInt(0).SetInt64(int64(bNum.Precision))
	bNum.ScaleFactor = big.NewInt(0).Exp(base10, bigPrecision, nil)
	baseZero := big.NewInt(0).SetInt64(0)
	result := bNum.BigInt.Cmp(baseZero)

	if result == -1 {
		bNum.Sign = -1
		minusOne := big.NewInt(0).SetInt64(-1)
		bNum.AbsBigInt = big.NewInt(0).Mul(bNum.BigInt, minusOne)
	} else {
		bNum.Sign = 1
		bNum.AbsBigInt = big.NewInt(0).Set(bNum.BigInt)
	}
	
}

// BigIntPair - contains a pair of 'BitIntNum' types. This structure
// is used to set up calculations involving *big.Int types.
type BigIntPair struct {
	Big1 							BigIntNum
	Big1Compare 			int
	Big1AbsCompare 		int
	Precision1Compare int
	Big2							BigIntNum
}


// New - Creates an Empty BigIntPair instance. Both
// 'Big1' and 'Big2' are set to zero.  Both Precision
// values are also set to zero.
func (bPair BigIntPair) New() BigIntPair {
		base1Zero := big.NewInt(0)
		base2Zero := big.NewInt(0)
		b2Pair := BigIntPair{}.NewBase(base1Zero, 0, base2Zero, 0)
		return b2Pair
}

// NewBase - Creates a BigIntPair instance using two sets of
// *big.Int integers and their associated precision specifications.
func (bPair BigIntPair) NewBase(
						b1 *big.Int,
						b1Precision uint,
						b2 *big.Int,
						b2Precision uint) BigIntPair {


	b1BigIntNum := BigIntNum{}.NewBigInt(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.NewBigInt(b2, b2Precision)

	return BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

}


// NewBigIntNum - Creates a new BigIntPair instance from input parameters
// consisting of two 'BigIntNum' types.
func (bPair BigIntPair) NewBigIntNum(b1, b2 BigIntNum ) BigIntPair {

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1, b2)

	return bd2
}


// NewIntAry - Creates a new BigIntPair instance from two
// Decimal instances passed as input parameters.
//
func (bPair BigIntPair) NewDecimal(dec1, dec2 Decimal) (BigIntPair, error) {

	ePrefix := "BigIntPair.NewDecimal() "

	b1Num, err := BigIntNum{}.NewDecimal(dec1)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewDecimal(dec1). " +
				"Error='%v' ", err.Error())
	}

	b2Num, err := BigIntNum{}.NewDecimal(dec2)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewDecimal(dec2). " +
				"Error='%v' ", err.Error())
	}

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1Num, b2Num)

	return bd2, nil
}

// NewIntAry - Creates a new BigIntPair instance from two
// IntAry instances passed as input parameters.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bPair BigIntPair) NewIntAry(ia1, ia2 IntAry) (BigIntPair, error) {

	ePrefix := "BigIntPair.NewIntAry() "

	b1Num, err := BigIntNum{}.NewIntAry(ia1)

	if err != nil {
		return BigIntPair{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(ia1). " +
			"Error='%v' ", err.Error())
	}

	b2Num, err := BigIntNum{}.NewIntAry(ia2)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(ia2). " +
				"Error='%v' ", err.Error())
	}

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1Num, b2Num)

	return bd2, nil
}

// NewNumStr - Creates a new BigIntPair instance from two number strings
// passed as input parameters.
//
func (bPair BigIntPair) NewNumStr(n1NumStr, n2NumStr string) (BigIntPair, error) {
	ePrefix := "BigIntPair.NewNumStrDto() "
	b1Num, err := BigIntNum{}.NewNumStr(n1NumStr)

	if err != nil {
		return BigIntPair{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(n1NumStr). " +
			"numStr='%v' Error='%v' ", n1NumStr, err.Error())
	}

	b2Num, err := BigIntNum{}.NewNumStr(n2NumStr)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(n2NumStr). " +
				"numStr='%v' Error='%v' ", n2NumStr, err.Error())
	}

	b2Pair := BigIntPair{}

	b2Pair.SetBigIntPair(b1Num, b2Num)

	return b2Pair, nil

}

// NewNumStrDto - Creates a new BigIntPair instance from two NumStrDto
// instances passed as input parameters.
//
func (bPair BigIntPair) NewNumStrDto(n1Dto, n2Dto NumStrDto) (BigIntPair, error) {

	ePrefix := "BigIntPair.NewNumStrDto() "
	b1Num, err := BigIntNum{}.NewNumStrDto(n1Dto)

	if err != nil {
		return BigIntPair{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(n1Dto). " +
			"numStr='%v' Error='%v' ", n1Dto.GetNumStr(), err.Error())
	}

	b2Num, err := BigIntNum{}.NewNumStrDto(n2Dto)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(n2Dto). " +
				"numStr='%v' Error='%v' ", n2Dto.GetNumStr(), err.Error())
	}

	b2Pair := BigIntPair{}

	b2Pair.SetBigIntPair(b1Num, b2Num)

	return b2Pair, nil

}

// CopyIn - Copies the values provided by incoming BigIntPair
// parameter into the current BigIntPair instance.
func (bPair *BigIntPair) CopyIn(bd2 BigIntPair) {

	bPair.Big1.CopyIn(bd2.Big1)
	bPair.Big2.CopyIn(bd2.Big2)
	bPair.Big1Compare = bd2.Big1Compare
	bPair.Big1AbsCompare = bd2.Big1AbsCompare
	bPair.Precision1Compare = bd2.Precision1Compare

}

// CopyOut - Makes a deep copy of the current BigIntPair
// instance and returns it as a new BigIntPair object.
//
func (bPair *BigIntPair) CopyOut() BigIntPair {

	bd2 := BigIntPair{}.NewBigIntNum(bPair.Big1, bPair.Big2)

	return bd2
}

// SetBigIntPairSetBigIntPair -Sets the values of the current
// BigIntPair instance to the input values of b1 and b2
// respectively.
//
func (bPair *BigIntPair) SetBigIntPair(b1, b2 BigIntNum ) {
	bPair.Big1.CopyIn(b1)
	bPair.Big2.CopyIn(b2)

	bPair.Big1Compare = bPair.Big1.BigInt.Cmp(bPair.Big2.BigInt)
	bPair.Big1AbsCompare = bPair.Big1.AbsBigInt.Cmp(bPair.Big2.AbsBigInt)

	if bPair.Big1.Precision == bPair.Big2.Precision {
		bPair.Precision1Compare = 0
	} else if bPair.Big1.Precision > bPair.Big2.Precision {
		bPair.Precision1Compare = 1
	} else {
		// Must be bPair.Big1.Precision < bPair.Big2.Precision
		bPair.Precision1Compare = -1
	}

}

// MakePrecisionsEqual - Analyzes the two component
// BigIntNum's, b1 and b2, and then converts the
// one with the least precision to a new value
// with equivalent precision. When completed, this
// method insures that both component BigIntNum's
// are formated to the largest precision.
//
func(bPair *BigIntPair) MakePrecisionsEqual() {

	if bPair.Big1.Precision == bPair.Big2.Precision {
		// Nothing to do. Precisions are equal.
		return
	}

	base10 := big.NewInt(10)

	if bPair.Big1.Precision > bPair.Big2.Precision {
		deltaPrecision := big.NewInt(int64(bPair.Big1.Precision - bPair.Big2.Precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		newB2Int := big.NewInt(0).Mul(bPair.Big2.BigInt, deltaPrecisionScale)
		newB2Num := BigIntNum{}.NewBigInt(newB2Int, bPair.Big1.Precision)
		newBPair := BigIntPair{}.NewBigIntNum(bPair.Big1, newB2Num)
		bPair.CopyIn(newBPair)
		return

	}

	// Must be bPair.Big2.Precision > bPair.Big1.Precision
	deltaPrecision := big.NewInt(int64(bPair.Big2.Precision - bPair.Big1.Precision))
	deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
	newB1Int := big.NewInt(0).Mul(bPair.Big1.BigInt, deltaPrecisionScale)
	newB1Num := BigIntNum{}.NewBigInt(newB1Int, bPair.Big2.Precision)
	newBPair := BigIntPair{}.NewBigIntNum(newB1Num, bPair.Big2)

	bPair.CopyIn(newBPair)

	return
}

// BigIntBasicMathResult - Used to return the result
// of an Addition, Subtraction or Multiplication operation.
//
type BigIntBasicMathResult struct {
	Input BigIntPair
	Result BigIntNum
}

// BigIntDivideModResult - Used to return the result
// of Big Int division with a Quotient and Modulo
// Values.
//
type BigIntDivideModResult struct {
  Dividend BigIntNum
  Divisor BigIntNum
  Quotient BigIntNum
  Modulo BigIntNum
}

type BigIntDivideFracResult struct {
	Dividend BigIntNum
	Divisor BigIntNum
	Quotient *big.Rat
}
