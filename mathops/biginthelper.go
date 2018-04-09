package mathops

import (
	"math/big"
)

// BigIntNum - wraps a *big.Int integer and its associated
// precision and Sign Value.
type BigIntNum struct {
	BigInt			*big.Int
	AbsBigInt		*big.Int
	Precision 	uint			// Number of digits to the right of the decimal point.
	ScaleFactor *big.Int	// Scale Factor =  10^precision
	Sign				int				// Valid values are -1 or +1. Indicates the sign of the
												// 	the 'BigInt' integer.
}

// New - Creates a new BigIntNum instance using a *big.Int type and its
// associated precision.
//
func (bNum BigIntNum) New(bigI *big.Int, precision uint) BigIntNum {
	b := BigIntNum{}
	b.SetBigIntNum(bigI, precision)
	return b
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

	b2 := BigIntNum{}.New(big.NewInt(0).Set(bNum.BigInt), bNum.Precision)

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


	b1BigIntNum := BigIntNum{}.New(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.New(b2, b2Precision)

	return BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

}


// NewBigIntNum - Creates a new BigIntPair instance from input parameters
// consisting of two 'BigIntNum' types.
func (bPair BigIntPair) NewBigIntNum(b1, b2 BigIntNum ) BigIntPair {

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1, b2)

	return bd2
}


// CopyIn - Copies the values provided by incoming BigIntPair
// parameter into the current BigIntPair instance.
func (bPair *BigIntPair) CopyIn(bd2 BigIntPair) {

	bPair.Big1 = bd2.Big1.CopyOut()
	bPair.Big2 = bd2.Big2.CopyOut()
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
	bPair.Big1 = b1.CopyOut()
	bPair.Big2 = b2.CopyOut()

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
		return
	}

	base10 := big.NewInt(10)

	if bPair.Big1.Precision > bPair.Big2.Precision {
		deltaPrecision := big.NewInt(int64(bPair.Big1.Precision - bPair.Big2.Precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		newB2Int := big.NewInt(0).Mul(bPair.Big2.BigInt, deltaPrecisionScale)
		newB2Num := BigIntNum{}.New(newB2Int, bPair.Big1.Precision)
		newBPair := BigIntPair{}.NewBigIntNum(bPair.Big1, newB2Num)
		bPair.CopyIn(newBPair)
		return

	}

	// Must be bPair.Big2.Precision > bPair.Big1.Precision
	deltaPrecision := big.NewInt(int64(bPair.Big2.Precision - bPair.Big1.Precision))
	deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
	newB1Int := big.NewInt(0).Mul(bPair.Big1.BigInt, deltaPrecisionScale)
	newB1Num := BigIntNum{}.New(newB1Int, bPair.Big2.Precision)
	newBPair := BigIntPair{}.NewBigIntNum(newB1Num, bPair.Big2)

	bPair.CopyIn(newBPair)

	return
}

// BigIntMathDto - This type is used to perform math operations
// using the *big.Int Type.
//
// If you are unfamiliar with the *big.Int type, reference:
// 						https://golang.org/pkg/math/big/
//
type BigIntMathDto struct {
	Input  BigIntPair
	Result BigIntNum
}

func (bMath BigIntMathDto) New() BigIntMathDto {

	b2Math := BigIntMathDto{}

	b2Math.Input = BigIntPair{}.New()

	baseZero := big.NewInt(0)

	b2Math.Result = BigIntNum{}.New(baseZero, 0)

	return b2Math

}

// NewNewBigIntPair - Creates a new BigIntMathDto based on input parameter
// type, 'BigIntPair'
//
func (bMath BigIntMathDto) NewBigIntPairResult(bPair BigIntPair) BigIntMathDto {

	b2Math := BigIntMathDto{}

	b2Math.Input = bPair.CopyOut()

	return b2Math
}



func (bMath BigIntMathDto) NewAdd(
										b1 *big.Int,
											precision1 uint,
												b2 *big.Int,
													precision2 uint ) BigIntMathDto {

	b1Pair := BigIntPair{}.NewBase(b1, precision1, b2, precision2)

	b1Pair.MakePrecisionsEqual()

	b3 := big.NewInt(0).Add(b1Pair.Big1.BigInt, b1Pair.Big2.BigInt)

	b2Math := BigIntMathDto{}.New()
	b2Math.Input = b1Pair.CopyOut()
	b2Math.Result = BigIntNum{}.New(b3, b1Pair.Big2.Precision)
	return b2Math
}
