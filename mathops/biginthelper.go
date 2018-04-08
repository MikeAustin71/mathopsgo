package mathops

import (
	"math/big"
)

type BigIntNum struct {
	BigInt			*big.Int
	AbsBigInt		*big.Int
	Precision 	uint
	ScaleFactor *big.Int
	Sign				int
}

func (bNum BigIntNum) New(bigI *big.Int, precision uint) BigIntNum {
	b := BigIntNum{}
	b.SetBigIntNum(bigI, precision)
	return b
}

func (bNum *BigIntNum) CopyIn(bigN BigIntNum) {

	bNum.BigInt = big.NewInt(0).Set(bigN.BigInt)
	bNum.AbsBigInt = big.NewInt(0).Set(bigN.AbsBigInt)
	bNum.Precision = bigN.Precision
	bNum.ScaleFactor = big.NewInt(0).Set(bigN.ScaleFactor)
	bNum.Sign = bigN.Sign
}

func (bNum *BigIntNum) CopyOut() BigIntNum {

	b2 := BigIntNum{}.New(big.NewInt(0).Set(bNum.BigInt), bNum.Precision)

	return b2
}

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


type BigIntDuo struct {
	Big1 							BigIntNum
	Big1Compare 			int
	Big1AbsCompare 		int
	Precision1Compare int
	Big2							BigIntNum
}

func (bDuo BigIntDuo) New(b1, b2 BigIntNum ) BigIntDuo {

	bd2 := BigIntDuo{}

	bd2.SetDuo(b1, b2)

	return bd2
}


func (bDuo *BigIntDuo) CopyIn(bd2 BigIntDuo) {

	bDuo.Big1 = bd2.Big1.CopyOut()
	bDuo.Big2 = bd2.Big2.CopyOut()
	bDuo.Big1Compare = bd2.Big1Compare
	bDuo.Big1AbsCompare = bd2.Big1AbsCompare
	bDuo.Precision1Compare = bd2.Precision1Compare

}

func (bDuo *BigIntDuo) CopyOut() BigIntDuo {

	bd2 := BigIntDuo{}.New(bDuo.Big1, bDuo.Big2)

	return bd2
}


func (bDuo *BigIntDuo) SetDuo(b1, b2 BigIntNum ) {
	bDuo.Big1 = b1.CopyOut()
	bDuo.Big2 = b2.CopyOut()

	bDuo.Big1Compare = bDuo.Big1.BigInt.Cmp(bDuo.Big2.BigInt)
	bDuo.Big1AbsCompare = bDuo.Big1.AbsBigInt.Cmp(bDuo.Big2.AbsBigInt)

	if bDuo.Big1.Precision == bDuo.Big2.Precision {
		bDuo.Precision1Compare = 0
	} else if bDuo.Big1.Precision > bDuo.Big2.Precision {
		bDuo.Precision1Compare = 1
	} else {
		// Must be bDuo.Big1.Precision < bDuo.Big2.Precision
		bDuo.Precision1Compare = -1
	}

}

type BigIntDto struct {
	Input 			BigIntDuo
	Output 			BigIntDuo
}

func (bDto BigIntDto) New(bDuoInput BigIntDuo) BigIntDto {

	b2Dto := BigIntDto{}

	b2Dto.Input = bDuoInput.CopyOut()

	return b2Dto
}

func (bDto BigIntDto) NewBase(b1 *big.Int, p1 uint, b2 *big.Int, p2 uint) BigIntDto {

	b1IntNum := BigIntNum{}.New(b1, p1)

	b2IntNum := BigIntNum{}.New(b2, p2)

	bIntDuo := BigIntDuo{}.New(b1IntNum, b2IntNum)

	b2Dto := BigIntDto{}.New(bIntDuo)

	return b2Dto

}