package mathops

import (
	"testing"
	"math/big"
	"errors"
)


func TestBigIntPair_Compare_01(t *testing.T) {
	n1Str := "4"
	n2Str := "-24"
	expectedBig1Compare := 1
	expectedPrecision1Compare := 0
	expectedBig1AbsCompare := -1

	b1, oK := big.NewInt(0).SetString(n1Str, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(n1Str, 10)" +
			"n1Str='%v' ", n1Str)
	}


	b2, oK := big.NewInt(0).SetString(n2Str, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(n2Str, 10)" +
			"n2Str='%v' ", n2Str)
	}

	bPair := BigIntPair{}.NewBase(b1, 0, b2, 0)

	if expectedBig1Compare != bPair.Big1Compare  {
		t.Errorf("Error: Expected bPair.Big1Compare='%d'. "+
			"Instead, bPair.Big1Compare='%v'. bPair.Big1='%s'  bPair.Big2='%s'",
			expectedBig1Compare, bPair.Big1Compare,
				bPair.Big1.BigInt.Text(10), bPair.Big2.BigInt.Text(10))
	}

	if expectedPrecision1Compare != bPair.Precision1Compare {
		t.Errorf("Error: Expected bPair.Precision1Compare='%d'. " +
			"Instead, bPair.Precision1Compare='%d'. bPair.Big1.Precision='%d' " +
			"bPair.Big2.Precision='%d'.",
			expectedBig1Compare, bPair.Precision1Compare,
				bPair.Big1.Precision, bPair.Big2.Precision)
	}

	if expectedBig1AbsCompare != bPair.Big1AbsCompare  {
		t.Errorf("Error: Expected bPair.Big1AbsCompare='%d'. " +
			"Instead, bPair.Big1AbsCompare='%d'. bPair.Big1.AbsBigInt='%s' " +
			"bPair.Big2.AbsBigInt='%s'",
				expectedBig1AbsCompare, bPair.Big1AbsCompare,
					bPair.Big1.AbsBigInt.Text(10),
						bPair.Big2.AbsBigInt.Text(10))
	}

}

func TestBigIntPair_Compare_02(t *testing.T) {
	n1Str := "-24"
	n2Str := "4"
	expectedBig1Compare := -1
	expectedPrecision1Compare := 0
	expectedBig1AbsCompare := 1

	b1, oK := big.NewInt(0).SetString(n1Str, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(n1Str, 10)" +
			"n1Str='%v' ", n1Str)
	}


	b2, oK := big.NewInt(0).SetString(n2Str, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(n2Str, 10)" +
			"n2Str='%v' ", n2Str)
	}

	bPair := BigIntPair{}.NewBase(b1, 0, b2, 0)

	if expectedBig1Compare != bPair.Big1Compare  {
		t.Errorf("Error: Expected bPair.Big1Compare='%d'. "+
			"Instead, bPair.Big1Compare='%v'. bPair.Big1='%s'  bPair.Big2='%s'",
			expectedBig1Compare, bPair.Big1Compare,
				bPair.Big1.BigInt.Text(10), bPair.Big2.BigInt.Text(10))
	}

	if expectedPrecision1Compare != bPair.Precision1Compare {
		t.Errorf("Error: Expected bPair.Precision1Compare='%d'. " +
			"Instead, bPair.Precision1Compare='%d'. bPair.Big1.Precision='%d' " +
			"bPair.Big2.Precision='%d'.",
			expectedBig1Compare, bPair.Precision1Compare,
				bPair.Big1.Precision, bPair.Big2.Precision)
	}

	if expectedBig1AbsCompare != bPair.Big1AbsCompare  {
		t.Errorf("Error: Expected bPair.Big1AbsCompare='%d'. " +
			"Instead, bPair.Big1AbsCompare='%d'. bPair.Big1.AbsBigInt='%s' " +
			"bPair.Big2.AbsBigInt='%s'",
				expectedBig1AbsCompare, bPair.Big1AbsCompare,
					bPair.Big1.AbsBigInt.Text(10),
						bPair.Big2.AbsBigInt.Text(10))
	}

}


func TestBigIntPair_Compare_03(t *testing.T) {
	n1Str := "40000"
	n1Precision := uint(4)
	n2Str := "240"
	n2Precision := uint(1)

	expectedMaxPrecision := uint(4)
	expectedBig1Compare := 1
	expectedPrecision1Compare := 1
	expectedBig1AbsCompare := 1

	b1, oK := big.NewInt(0).SetString(n1Str, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(n1Str, 10)" +
			"n1Str='%v' ", n1Str)
	}


	b2, oK := big.NewInt(0).SetString(n2Str, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(n2Str, 10)" +
			"n2Str='%v' ", n2Str)
	}

	bPair := BigIntPair{}.NewBase(b1, n1Precision, b2, n2Precision)

	if expectedBig1Compare != bPair.Big1Compare  {
		t.Errorf("Error: Expected bPair.Big1Compare='%d'. "+
			"Instead, bPair.Big1Compare='%v'. bPair.Big1='%s'  bPair.Big2='%s'",
			expectedBig1Compare, bPair.Big1Compare,
				bPair.Big1.BigInt.Text(10), bPair.Big2.BigInt.Text(10))
	}

	if expectedPrecision1Compare != bPair.Precision1Compare {
		t.Errorf("Error: Expected bPair.Precision1Compare='%d'. " +
			"Instead, bPair.Precision1Compare='%d'. bPair.Big1.Precision='%d' " +
			"bPair.Big2.Precision='%d'.",
			expectedBig1Compare, bPair.Precision1Compare,
				bPair.Big1.Precision, bPair.Big2.Precision)
	}

	if expectedBig1AbsCompare != bPair.Big1AbsCompare  {
		t.Errorf("Error: Expected bPair.Big1AbsCompare='%d'. " +
			"Instead, bPair.Big1AbsCompare='%d'. bPair.Big1.AbsBigInt='%s' " +
			"bPair.Big2.AbsBigInt='%s'",
				expectedBig1AbsCompare, bPair.Big1AbsCompare,
					bPair.Big1.AbsBigInt.Text(10),
						bPair.Big2.AbsBigInt.Text(10))
	}

	bPair.MakePrecisionsEqual()
	expectedBig1Compare = -1
	expectedPrecision1Compare = 0
	expectedBig1AbsCompare = -1

	if expectedMaxPrecision != bPair.Big2.Precision {
		t.Errorf("Error: After Equalizing Precision, Expected Big2.Precision='%v'. " +
			"Instead, Big2.Precision='%v'.", expectedMaxPrecision, bPair.Big2.Precision )
	}

	if expectedBig1Compare != bPair.Big1Compare {
		t.Errorf("Error: After Equalizing Precision, Expected Big1Compare='%v'. " +
			"Instead, Big1Compare='%v'. bPair.Big1='%s' bPair.Big2='%s'. ",
				expectedBig1Compare, bPair.Big1Compare, bPair.Big1.BigInt.Text(10),
					bPair.Big2.BigInt.Text(10))
	}

	if expectedBig1AbsCompare != bPair.Big1AbsCompare {
		t.Errorf("Error: After Equalizing Precision, Expected Big1AbsCompare='%v'. " +
			"Instead, Big1AbsCompare='%v'. bPair.Big1='%s' bPair.Big2='%s'. ",
				expectedBig1AbsCompare, bPair.Big1AbsCompare, bPair.Big1.BigInt.Text(10),
					bPair.Big2.BigInt.Text(10))
	}

	if expectedPrecision1Compare != bPair.Precision1Compare {
		t.Errorf("Error: After Equalizing Precision, Expected Precision1Compare='%v'. " +
			"Instead, bPair.Precision1Compare='%v'. bPair.Big1.Precision='%v' " +
			"bPair.Big2.Precision='%v'",
				expectedPrecision1Compare, bPair.Precision1Compare,
					bPair.Big1.Precision, bPair.Big2.Precision)
	}
}

func TestBigIntPair_NewBase_01(t *testing.T) {

	//n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	//n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	bPair := BigIntPair{}.NewBase(b1, b1Precision, b2, b2Precision)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}
func TestBigIntPair_NewBase_02(t *testing.T) {

	//n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	//n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	bPair := BigIntPair{}.NewBase(b1, b1Precision, b2, b2Precision)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewBigIntNum_01(t *testing.T) {

	//n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	//n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	b1BigIntNum := BigIntNum{}.NewBigInt(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.NewBigInt(b2, b2Precision)

	bPair := BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewBigIntNum_02(t *testing.T) {

	//n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	//n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	b1BigIntNum := BigIntNum{}.NewBigInt(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.NewBigInt(b2, b2Precision)

	bPair := BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStr_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	bPair, err := BigIntPair{}.NewNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStr(n1Str, n2Str). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStr_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	bPair, err := BigIntPair{}.NewNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStr(n1Str, n2Str). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStrDto_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	nx1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	nx2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStrDto_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	nx1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	nx2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewIntAry_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewIntAry_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_Decimal_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_Decimal_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

