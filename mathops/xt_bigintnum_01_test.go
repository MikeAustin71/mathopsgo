package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_BigInt_01(t *testing.T) {

	ePrefix := "TestBigIntNum_BigInt_01"

	expectedBigINumStr := "123.456"

	expectedBigINumPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nbStr := "123456"

	expectedScale := big.NewInt(1000)

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
			"nbStr= '%v'\n\n",
			ePrefix, nbStr)
		return
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewBigInt(\n"+
			"  bOriginal, expectedPrecision)\n"+
			"bOriginal= '%v'\n"+
			"expectedPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginal.Text(10),
			expectedBigINumPrecisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_BigInt_02(t *testing.T) {

	ePrefix := "TestBigIntNum_BigInt_02"

	expectedBigINumStr := "-123.456"

	expectedBigINumPrecisionUint := uint(3)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nbStr := "-123456"

	expectedScale := big.NewInt(1000)

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
			"nbStr= '%v'\n\n",
			ePrefix, nbStr)
		return
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewBigInt(\n"+
			"  bOriginal, expectedPrecision)\n"+
			"bOriginal= '%v'\n"+
			"expectedPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginal.Text(10),
			expectedBigINumPrecisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_BigInt_03(t *testing.T) {

	ePrefix := "TestBigIntNum_BigInt_03"

	expectedBigINumStr := "0.000123456"

	expectedBigINumPrecisionUint := uint(9)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nbI64 := int64(123456)

	expectedScale := big.NewInt(1000000000)

	bOriginal := big.NewInt(nbI64)

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewBigInt(\n"+
			"  bOriginal, expectedPrecision)\n"+
			"bOriginal= '%v'\n"+
			"expectedPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginal.Text(10),
			expectedBigINumPrecisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_BigInt_04(t *testing.T) {

	ePrefix := "TestBigIntNum_BigInt_04"

	expectedBigINumStr := "-0.000123456"

	expectedBigINumPrecisionUint := uint(9)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nbI64 := int64(-123456)

	expectedScale := big.NewInt(1000000000)

	bOriginal := big.NewInt(nbI64)

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewBigInt(\n"+
			"  bOriginal, expectedPrecision)\n"+
			"bOriginal= '%v'\n"+
			"expectedPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginal.Text(10),
			expectedBigINumPrecisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_01"

	originalNumStr := "5.95"

	expectedBigINumStr := "6"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_02"

	originalNumStr := "5.05"

	expectedBigINumStr := "6"

	expectedBigINumPrecisionUint := uint(0)

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_03"

	originalNumStr := "5"

	expectedBigINumStr := "5"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_04(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_04"

	originalNumStr := "-5.05"

	expectedBigINumStr := "-5"

	expectedBigINumSign := -1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_05(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_05"

	originalNumStr := "2.4"

	expectedBigINumStr := "3"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_06(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_06"

	originalNumStr := "2.9"

	expectedBigINumStr := "3"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_07(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_07"

	originalNumStr := "-2.7"

	expectedBigINumStr := "-2"

	expectedBigINumSign := -1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_08(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_08"

	originalNumStr := "-2"

	expectedBigINumStr := "-2"

	expectedBigINumSign := -1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_09(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_09"

	originalNumStr := "0"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_10(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_10"

	originalNumStr := "0.00000"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_11(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_11"

	originalNumStr := "159876231.9999999999"

	expectedBigINumStr := "159876232"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_12(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_12"

	originalNumStr := "-159876231.9999999999"

	expectedBigINumStr := "-159876231"

	expectedBigINumSign := -1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_13(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_13"

	originalNumStr := "159876231.0000000000000001"

	expectedBigINumStr := "159876232"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_14(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_14"

	originalNumStr := "-159876231.0000000000000001"

	expectedBigINumStr := "-159876231"

	expectedBigINumSign := -1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_15(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_15"

	originalNumStr := "-0.0000000000000001"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_16(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_16"

	originalNumStr := "0.0000000000000001"

	expectedBigINumStr := "1"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Ceil_17(t *testing.T) {

	ePrefix := "TestBigIntNum_Ceil_17"

	originalNumStr := "-0.0000000000000001"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	expectedBigINumPrecisionUint := uint(0)

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeiling, err := bINum.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeiling, err := bINum.Ceiling()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINumCeiling.IsValid("Validating bINumCeiling")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINumCeiling')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingPrecisionUint, err :=\n"+
			"  bINumCeiling.GetPrecisionUint()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingSignValue, err := bINumCeiling.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumCeilingNumSeps, err := \n"+
			"  bINumCeiling.GetNumericSeparatorsDto()\n"+
			"bINumCeiling= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
		return
	}

	if expectedBigINumStr != bINumCeilingNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumCeilingNumStr\n"+
			"Expected bINumCeilingNumStr = '%v'\n"+
			"  Actual bINumCeilingNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumCeilingNumStr)

		return
	}

	if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected bINumCeilingPrecisionUint = '%v'\n"+
			"  Actual bINumCeilingPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

		return
	}

	if expectedBigINumSign != bINumCeilingSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumCeilingSignValue\n"+
			"Expected bINumCeilingSignValue = '%v'\n"+
			"  Actual bINumCeilingSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumCeilingSignValue)

		return
	}

	if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumCeilingNumSeps = '%v'\n"+
			"  Actual bINumCeilingNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ChangeSign_01(t *testing.T) {

	ePrefix := "TestBigIntNum_ChangeSign_01"

	originalNumStr := "123.456"

	expectedBigINumStr := "-123.456"

	originalNumSign := 1

	expectedBigINumSign := -1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumOriginalSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	err = bINum.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Sign Change")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Sign Change')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumNewNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumSeps, err :=\n"+
			"  bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumNewNumStr, err.Error())
		return
	}

	if originalNumSign != bINumOriginalSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumSign != bINumOriginalSignValue\n"+
			"Expected bINumOriginalSignValue = '%v'\n"+
			"  Actual bINumOriginalSignValue = '%v'\n"+
			"Original bINumStr= '%v'\n\n",
			ePrefix, originalNumSign, bINumOriginalSignValue, bIOriginalNumStr)

		return
	}

	if expectedBigINumSign != bINumNewSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumNewSignValue\n"+
			"Expected bINumNewSignValue = '%v'\n"+
			"  Actual bINumNewSignValue = '%v'\n"+
			"New/Final bINumStr= '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumNewSignValue, bINumNewNumStr)

		return
	}

	if originalNumStr != bIOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bIOriginalNumStr \n"+
			"Expected bIOriginalNumStr = '%v'\n"+
			"  Actual bIOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bIOriginalNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumNewNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumNewNumSeps \n"+
			"Expected bINumNewNumSeps = '%v'\n"+
			"  Actual bINumNewNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumNewNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ChangeSign_02(t *testing.T) {

	ePrefix := "TestBigIntNum_ChangeSign_02"

	originalNumStr := "-123.456"

	expectedBigINumStr := "123.456"

	originalNumSign := -1

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumOriginalSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	err = bINum.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Sign Change")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Sign Change')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumNewNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumSeps, err :=\n"+
			"  bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumNewNumStr, err.Error())
		return
	}

	if originalNumSign != bINumOriginalSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumSign != bINumOriginalSignValue\n"+
			"Expected bINumOriginalSignValue = '%v'\n"+
			"  Actual bINumOriginalSignValue = '%v'\n"+
			"Original bINumStr= '%v'\n\n",
			ePrefix, originalNumSign, bINumOriginalSignValue, bIOriginalNumStr)

		return
	}

	if expectedBigINumSign != bINumNewSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumNewSignValue\n"+
			"Expected bINumNewSignValue = '%v'\n"+
			"  Actual bINumNewSignValue = '%v'\n"+
			"New/Final bINumStr= '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumNewSignValue, bINumNewNumStr)

		return
	}

	if originalNumStr != bIOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bIOriginalNumStr \n"+
			"Expected bIOriginalNumStr = '%v'\n"+
			"  Actual bIOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bIOriginalNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumNewNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumNewNumSeps \n"+
			"Expected bINumNewNumSeps = '%v'\n"+
			"  Actual bINumNewNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumNewNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ChangeSign_03(t *testing.T) {

	ePrefix := "TestBigIntNum_ChangeSign_03"

	originalNumStr := "0.00"

	expectedBigINumStr := "0.00"

	originalNumSign := 1

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumOriginalSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	err = bINum.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Sign Change")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Sign Change')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bIOriginalNumStr, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumNewNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumSeps, err :=\n"+
			"  bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumNewNumStr, err.Error())
		return
	}

	if originalNumSign != bINumOriginalSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumSign != bINumOriginalSignValue\n"+
			"Expected bINumOriginalSignValue = '%v'\n"+
			"  Actual bINumOriginalSignValue = '%v'\n"+
			"Original bINumStr= '%v'\n\n",
			ePrefix, originalNumSign, bINumOriginalSignValue, bIOriginalNumStr)

		return
	}

	if expectedBigINumSign != bINumNewSignValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumSign != bINumNewSignValue\n"+
			"Expected bINumNewSignValue = '%v'\n"+
			"  Actual bINumNewSignValue = '%v'\n"+
			"New/Final bINumStr= '%v'\n\n",
			ePrefix, expectedBigINumSign, bINumNewSignValue, bINumNewNumStr)

		return
	}

	if originalNumStr != bIOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bIOriginalNumStr \n"+
			"Expected bIOriginalNumStr = '%v'\n"+
			"  Actual bIOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bIOriginalNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumNewNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumNewNumSeps \n"+
			"Expected bINumNewNumSeps = '%v'\n"+
			"  Actual bINumNewNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumNewNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Cmp_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_01"

	originalNumStr1 := "123.456"

	originalNumStr2 := "123.455"

	expectedCompareResult := 1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_02"

	originalNumStr1 := "123.456"

	originalNumStr2 := "123.457"

	expectedCompareResult := -1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_03"

	originalNumStr1 := "123.456"

	originalNumStr2 := "123.456"

	expectedCompareResult := 0

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_04(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_04"

	originalNumStr1 := "-123.456"

	originalNumStr2 := "-123.457"

	expectedCompareResult := 1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_05(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_05"

	originalNumStr1 := "-123.456"

	originalNumStr2 := "-123.455"

	expectedCompareResult := -1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_06(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_06"

	originalNumStr1 := "-123.456"

	originalNumStr2 := "-123.456"

	expectedCompareResult := 0

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_07(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_07"

	originalNumStr1 := "123456000643218"

	originalNumStr2 := "123456000643217"

	expectedCompareResult := 1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_08(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_08"

	originalNumStr1 := "123456000643218"

	originalNumStr2 := "123456000643219"

	expectedCompareResult := -1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_09(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_09"

	originalNumStr1 := "123456000643218"

	originalNumStr2 := "123456000643218"

	expectedCompareResult := 0

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_10(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_10"

	originalNumStr1 := "-123456000643218"

	originalNumStr2 := "-123456000643217"

	expectedCompareResult := -1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_11(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_11"

	originalNumStr1 := "-123456000643218"

	originalNumStr2 := "-123456000643219"

	expectedCompareResult := 1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_12(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_12"

	originalNumStr1 := "-123456000643218"

	originalNumStr2 := "-123456000643218"

	expectedCompareResult := 0

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_13(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_13"

	originalNumStr1 := "1.7"

	originalNumStr2 := "-12345.6000643218"

	expectedCompareResult := 1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_14(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_14"

	originalNumStr1 := "-1.7"

	originalNumStr2 := "0.01"

	expectedCompareResult := -1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_15(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_15"

	originalNumStr1 := "17"

	originalNumStr2 := "-1"

	expectedCompareResult := 1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_16(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_16"

	originalNumStr1 := "-17"

	originalNumStr2 := "1"

	expectedCompareResult := -1

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Cmp_17(t *testing.T) {

	ePrefix := "TestBigIntNum_Cmp_17"

	originalNumStr1 := "0"

	originalNumStr2 := "0.000"

	expectedCompareResult := 0

	bINum1, err := new(BigIntNum).NewNumStr(originalNumStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr1)\n"+
			"originalNumStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr1,
			err.Error())

		return
	}

	err = bINum1.IsValid("Validating bINum1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum1.IsValid('Validating bINum1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum1NumStr, err := bINum1.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1NumStr, err := bINum1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr1 != bINum1NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr1 != bINum1NumStr\n"+
			"Expected bINum1NumStr = '%v'\n"+
			"  Actual bINum1NumStr = '%v'\n\n",
			ePrefix, originalNumStr1, bINum1NumStr)

		return
	}

	bINum2, err := new(BigIntNum).NewNumStr(originalNumStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr2)\n"+
			"originalNumStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr2,
			err.Error())

		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumStr, err := bINum2.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumStr, err := bINum2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr2 != bINum2NumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr2 != bINum2NumStr\n"+
			"Expected bINum2NumStr = '%v'\n"+
			"  Actual bINum2NumStr = '%v'\n\n",
			ePrefix, originalNumStr2, bINum2NumStr)

		return
	}

	compareResult, err := bINum1.Cmp(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"compareResult, err := bINum1.Cmp(bINum2)\n"+
			"bINum1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINum1NumStr,
			bINum2NumStr,
			err.Error())

		return
	}

	if expectedCompareResult != compareResult {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedCompareResult != compareResult\n"+
			"Expected compareResult = '%v'\n"+
			"  Actual compareResult = '%v'\n\n",
			ePrefix, expectedCompareResult, compareResult)

		return
	}

	return
}

func TestBigIntNum_Decimal_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Decimal_01"

	expectedBigINumStr := "123.456"

	expectedBigINumPrecisionUint := uint(3)

	nbStr := "123456"

	absoluteNbStr := "123456"

	expectedScale := big.NewInt(1000)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
			"nbStr= '%v'\n\n",
			ePrefix, nbStr)
		return
	}

	expectedAbsBigInt, isOk := big.NewInt(0).SetString(absoluteNbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigInt, isOk := big.NewInt(0).\n"+
			"  SetString(absoluteNbStr, 10)\n"+
			"absoluteNbStr= '%v'\n\n",
			ePrefix, absoluteNbStr)
		return
	}

	decNum, err := new(Decimal).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decNum, err := new(Decimal).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = decNum.IsValid("Validating decNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decNum.IsValid('Validating decNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decNumStr, err := decNum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decNumStr, err := decNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINum, err := new(BigIntNum).NewDecimal(decNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewDecimal(decNum)\n"+
			"decNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsoluteBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsoluteBigInt, err :=\n"+
			"  bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsoluteBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Absolute Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsoluteBigInt) != 0\n"+
			"Expected bINumAbsoluteBigInt = '%v'\n"+
			"  Actual bINumAbsoluteBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsoluteBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Decimal_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Decimal_02"

	expectedBigINumStr := "-123.456"

	expectedBigINumPrecisionUint := uint(3)

	nbStr := "-123456"

	absoluteNbStr := "-123456"

	expectedScale := big.NewInt(1000)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
			"nbStr= '%v'\n\n",
			ePrefix, nbStr)
		return
	}

	expectedAbsBigInt, isOk := big.NewInt(0).SetString(absoluteNbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigInt, isOk := big.NewInt(0).\n"+
			"  SetString(absoluteNbStr, 10)\n"+
			"absoluteNbStr= '%v'\n\n",
			ePrefix, absoluteNbStr)
		return
	}

	decNum, err := new(Decimal).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decNum, err := new(Decimal).\n"+
			"  NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	err = decNum.IsValid("Validating decNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decNum.IsValid('Validating decNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decNumStr, err := decNum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decNumStr, err := decNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINum, err := new(BigIntNum).NewDecimal(decNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewDecimal(decNum)\n"+
			"decNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsoluteBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsoluteBigInt, err :=\n"+
			"  bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsoluteBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Absolute Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsoluteBigInt) != 0\n"+
			"Expected bINumAbsoluteBigInt = '%v'\n"+
			"  Actual bINumAbsoluteBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsoluteBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Decimal_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Decimal_03"

	expectedBigINumStr := "-123,456"

	expectedBigINumPrecisionUint := uint(3)

	nbStr := "-123456"

	absoluteNbStr := "-123456"

	expectedScale := big.NewInt(1000)

	expectedSignVal := -1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
			"nbStr= '%v'\n\n",
			ePrefix, nbStr)
		return
	}

	expectedAbsBigInt, isOk := big.NewInt(0).SetString(absoluteNbStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigInt, isOk := big.NewInt(0).\n"+
			"  SetString(absoluteNbStr, 10)\n"+
			"absoluteNbStr= '%v'\n\n",
			ePrefix, absoluteNbStr)
		return
	}

	decNum, err := new(Decimal).NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decNum, err := new(Decimal).\n"+
			"  NewNumStr(expectedBigINumStr, expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = decNum.IsValid("Validating decNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decNum.IsValid('Validating decNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decNumStr, err := decNum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decNumStr, err := decNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINum, err := new(BigIntNum).NewDecimal(decNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewDecimal(decNum)\n"+
			"decNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumAbsoluteBigInt, err := bINum.GetAbsoluteBigIntValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsoluteBigInt, err :=\n"+
			"  bINum.GetAbsoluteBigIntValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	if expectedScale.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors ARE NOT Equal!\n"+
			"Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

		return
	}

	if bOriginal.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because bOriginal.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

		return
	}

	if expectedAbsBigInt.Cmp(bINumAbsoluteBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Absolute Values ARE NOT Equal!\n"+
			"Because expectedAbsBigInt.Cmp(bINumAbsoluteBigInt) != 0\n"+
			"Expected bINumAbsoluteBigInt = '%v'\n"+
			"  Actual bINumAbsoluteBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigInt.Text(10), bINumAbsoluteBigInt.Text(10))

		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: Sign Values ARE NOT Equal!\n"+
			"Because expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	if expectedBigINumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps\n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Decrement_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Decrement_01"

	originalNumStr := "8"

	expectedBigINumStr := "7"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bINumOriginalNumStr\n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.Decrement()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Decrement()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Decrement")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Decrement')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Note: 'bINum' After Decrement"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumFinalNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, bINumNewNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumFinalNumSeps \n"+
			"Expected bINumFinalNumSeps = '%v'\n"+
			"  Actual bINumFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumFinalNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Decrement_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Decrement_02"

	originalNumStr := "8.2"

	expectedBigINumStr := "7.2"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bINumOriginalNumStr\n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.Decrement()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Decrement()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Decrement")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Decrement')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Note: 'bINum' After Decrement"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumFinalNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, bINumNewNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumFinalNumSeps \n"+
			"Expected bINumFinalNumSeps = '%v'\n"+
			"  Actual bINumFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumFinalNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Decrement_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Decrement_03"

	originalNumStr := "-8.2"

	expectedBigINumStr := "-9.2"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bINumOriginalNumStr\n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.Decrement()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Decrement()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Decrement")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Decrement')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Note: 'bINum' After Decrement"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumFinalNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, bINumNewNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumFinalNumSeps \n"+
			"Expected bINumFinalNumSeps = '%v'\n"+
			"  Actual bINumFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumFinalNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Decrement_04(t *testing.T) {

	ePrefix := "TestBigIntNum_Decrement_04"

	originalNumStr := "8"

	expectedBigINumStr := "7"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumStr != bINumOriginalNumStr\n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.Decrement()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Decrement()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Decrement")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Decrement')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNewNumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNewNumStr, err := bINum.GetNumStr()\n"+
			"Note: 'bINum' After Decrement"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	bINumFinalNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, bINumNewNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if expectedBigINumStr != bINumNewNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedBigINumStr != bINumNewNumStr\n"+
			"Expected bINumNewNumStr = '%v'\n"+
			"  Actual bINumNewNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, bINumNewNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumFinalNumSeps \n"+
			"Expected bINumFinalNumSeps = '%v'\n"+
			"  Actual bINumFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumFinalNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTwo_01(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTwo_01"

	numStr := "658.78562347"

	expectedQuoStr := "329"

	expectedModStr := "0.78562347"

	maxPrecision := uint(10)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedQuotientBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuotientBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)\n"+
			"expectedQuoStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedQuotientBigINum.IsValid("Validating expectedQuotientBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedQuotientBigINum.IsValid('Validating expectedQuotientBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedQuotientBigINumberStr, err := expectedQuotientBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuotientBigINumberStr, err := expectedQuotientBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedModStr, &expectedNumSeps)\n"+
			"expectedModStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedModuloBigINum.IsValid("Validating expectedModuloBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBigINum.IsValid('Validating expectedModuloBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloBigINumberStr, err := expectedModuloBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBigINumberStr, err := expectedModuloBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	base, err := new(BigIntNum).NewNumStr(numStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"base, err := new(BigIntNum).NewNumStr(numStr)\n"+
			"numStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStr, err.Error())
		return
	}

	err = base.IsValid("Validating base")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = base.IsValid('Validating base')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	baseNumberStr, err := base.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseNumberStr, err := base.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)\n"+
			"base= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseNumberStr, maxPrecision, err.Error())
		return
	}

	err = quotient.IsValid("Validating quotient")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = quotient.IsValid('Validating quotient')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumberStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumberStr, err := quotient.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumSeps, err := quotient.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
			"quotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, quotientNumberStr, err.Error())
		return
	}

	err = modulo.IsValid("Validating modulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = modulo.IsValid('Validating modulo')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	moduloNumberStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"moduloNumberStr, err := modulo.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	moduloNumSeps, err := modulo.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"moduloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
			"modulo= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, moduloNumberStr, err.Error())
		return
	}

	if expectedQuoStr != quotientNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuoStr != quotientNumberStr\n"+
			"Expected quotientNumberStr = '%v'\n"+
			"  Actual quotientNumberStr = '%v'\n\n",
			ePrefix, expectedQuoStr, quotientNumberStr)

		return
	}

	expectedEqualsQuotient, err := expectedQuotientBigINum.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsQuotient, err := \n"+
			"  expectedQuotientBigINum.Equal(quotient)\n"+
			"expectedQuotientBigINum= '%v'\n"+
			"quotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedQuotientBigINumberStr, quotientNumberStr, err.Error())
		return
	}

	if !expectedEqualsQuotient {
		t.Errorf("%v\n"+
			"Error: Expected Quotient and 'quotient' values NOT Equal\n"+
			"Because expectedEqualsQuotient = 'false' \n"+
			"Expected quotient = '%v'\n"+
			"  Actual quotient = '%v'\n\n",
			ePrefix, expectedQuotientBigINumberStr, quotientNumberStr)

		return
	}

	if !expectedNumSeps.Equal(quotientNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != quotientNumSeps\n"+
			"Expected quotientNumSeps = '%v'\n"+
			"  Actual quotientNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), quotientNumSeps.String())

		return
	}

	if expectedModStr != moduloNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModStr != moduloNumberStr\n"+
			"Expected moduloNumberStr = '%v'\n"+
			"  Actual moduloNumberStr = '%v'\n\n",
			ePrefix, expectedModStr, moduloNumberStr)

		return
	}

	expectedEqualsModulo, err := expectedModuloBigINum.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsModulo, err := \n"+
			"  expectedModuloBigINum.Equal(modulo)\n"+
			"expectedModuloBigINum= '%v'\n"+
			"modulo= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModuloBigINumberStr, moduloNumberStr, err.Error())
		return
	}

	if !expectedEqualsModulo {
		t.Errorf("%v\n"+
			"Error: Expected Modulo and 'modulo' values NOT Equal!\n"+
			"Because expectedEqualsModulo = 'false' \n"+
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloBigINumberStr, moduloNumberStr)

		return
	}

	if !expectedNumSeps.Equal(moduloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != moduloNumSeps\n"+
			"Expected moduloNumSeps = '%v'\n"+
			"  Actual moduloNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), moduloNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTwo_02(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTwo_02"

	numStr := "8"

	expectedQuoStr := "4"

	expectedModStr := "0"

	maxPrecision := uint(10)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedQuotientBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuotientBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)\n"+
			"expectedQuoStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedQuotientBigINum.IsValid("Validating expectedQuotientBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedQuotientBigINum.IsValid('Validating expectedQuotientBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedQuotientBigINumberStr, err := expectedQuotientBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuotientBigINumberStr, err := expectedQuotientBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedModStr, &expectedNumSeps)\n"+
			"expectedModStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedModuloBigINum.IsValid("Validating expectedModuloBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBigINum.IsValid('Validating expectedModuloBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloBigINumberStr, err := expectedModuloBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBigINumberStr, err := expectedModuloBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	base, err := new(BigIntNum).NewNumStr(numStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"base, err := new(BigIntNum).NewNumStr(numStr)\n"+
			"numStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStr, err.Error())
		return
	}

	err = base.IsValid("Validating base")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = base.IsValid('Validating base')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	baseNumberStr, err := base.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseNumberStr, err := base.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)\n"+
			"base= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseNumberStr, maxPrecision, err.Error())
		return
	}

	err = quotient.IsValid("Validating quotient")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = quotient.IsValid('Validating quotient')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumberStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumberStr, err := quotient.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumSeps, err := quotient.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
			"quotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, quotientNumberStr, err.Error())
		return
	}

	err = modulo.IsValid("Validating modulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = modulo.IsValid('Validating modulo')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	moduloNumberStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"moduloNumberStr, err := modulo.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	moduloNumSeps, err := modulo.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"moduloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
			"modulo= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, moduloNumberStr, err.Error())
		return
	}

	if expectedQuoStr != quotientNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuoStr != quotientNumberStr\n"+
			"Expected quotientNumberStr = '%v'\n"+
			"  Actual quotientNumberStr = '%v'\n\n",
			ePrefix, expectedQuoStr, quotientNumberStr)

		return
	}

	expectedEqualsQuotient, err := expectedQuotientBigINum.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsQuotient, err := \n"+
			"  expectedQuotientBigINum.Equal(quotient)\n"+
			"expectedQuotientBigINum= '%v'\n"+
			"quotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedQuotientBigINumberStr, quotientNumberStr, err.Error())
		return
	}

	if !expectedEqualsQuotient {
		t.Errorf("%v\n"+
			"Error: Expected Quotient and 'quotient' values NOT Equal\n"+
			"Because expectedEqualsQuotient = 'false' \n"+
			"Expected quotient = '%v'\n"+
			"  Actual quotient = '%v'\n\n",
			ePrefix, expectedQuotientBigINumberStr, quotientNumberStr)

		return
	}

	if !expectedNumSeps.Equal(quotientNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != quotientNumSeps\n"+
			"Expected quotientNumSeps = '%v'\n"+
			"  Actual quotientNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), quotientNumSeps.String())

		return
	}

	if expectedModStr != moduloNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModStr != moduloNumberStr\n"+
			"Expected moduloNumberStr = '%v'\n"+
			"  Actual moduloNumberStr = '%v'\n\n",
			ePrefix, expectedModStr, moduloNumberStr)

		return
	}

	expectedEqualsModulo, err := expectedModuloBigINum.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsModulo, err := \n"+
			"  expectedModuloBigINum.Equal(modulo)\n"+
			"expectedModuloBigINum= '%v'\n"+
			"modulo= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModuloBigINumberStr, moduloNumberStr, err.Error())
		return
	}

	if !expectedEqualsModulo {
		t.Errorf("%v\n"+
			"Error: Expected Modulo and 'modulo' values NOT Equal!\n"+
			"Because expectedEqualsModulo = 'false' \n"+
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloBigINumberStr, moduloNumberStr)

		return
	}

	if !expectedNumSeps.Equal(moduloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != moduloNumSeps\n"+
			"Expected moduloNumSeps = '%v'\n"+
			"  Actual moduloNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), moduloNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTwo_03(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTwo_03"

	numStr := "658,78562347"

	expectedQuoStr := "329"

	expectedModStr := "0,78562347"

	maxPrecision := uint(10)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedQuotientBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuotientBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)\n"+
			"expectedQuoStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedQuotientBigINum.IsValid("Validating expectedQuotientBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedQuotientBigINum.IsValid('Validating expectedQuotientBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedQuotientBigINumberStr, err := expectedQuotientBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuotientBigINumberStr, err := expectedQuotientBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedModStr, &expectedNumSeps)\n"+
			"expectedModStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedModuloBigINum.IsValid("Validating expectedModuloBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBigINum.IsValid('Validating expectedModuloBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloBigINumberStr, err := expectedModuloBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBigINumberStr, err := expectedModuloBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	base, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"base, err := new(BigIntNum).\n"+
			"  NewNumStr(numStr, &expectedNumSeps)\n"+
			"numStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = base.IsValid("Validating base")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = base.IsValid('Validating base')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	baseNumberStr, err := base.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseNumberStr, err := base.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)\n"+
			"base= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseNumberStr, maxPrecision, err.Error())
		return
	}

	err = quotient.IsValid("Validating quotient")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = quotient.IsValid('Validating quotient')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumberStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumberStr, err := quotient.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumSeps, err := quotient.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
			"quotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, quotientNumberStr, err.Error())
		return
	}

	err = modulo.IsValid("Validating modulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = modulo.IsValid('Validating modulo')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	moduloNumberStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"moduloNumberStr, err := modulo.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	moduloNumSeps, err := modulo.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"moduloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
			"modulo= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, moduloNumberStr, err.Error())
		return
	}

	if expectedQuoStr != quotientNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuoStr != quotientNumberStr\n"+
			"Expected quotientNumberStr = '%v'\n"+
			"  Actual quotientNumberStr = '%v'\n\n",
			ePrefix, expectedQuoStr, quotientNumberStr)

		return
	}

	expectedEqualsQuotient, err := expectedQuotientBigINum.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsQuotient, err := \n"+
			"  expectedQuotientBigINum.Equal(quotient)\n"+
			"expectedQuotientBigINum= '%v'\n"+
			"quotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedQuotientBigINumberStr, quotientNumberStr, err.Error())
		return
	}

	if !expectedEqualsQuotient {
		t.Errorf("%v\n"+
			"Error: Expected Quotient and 'quotient' values NOT Equal\n"+
			"Because expectedEqualsQuotient = 'false' \n"+
			"Expected quotient = '%v'\n"+
			"  Actual quotient = '%v'\n\n",
			ePrefix, expectedQuotientBigINumberStr, quotientNumberStr)

		return
	}

	if !expectedNumSeps.Equal(quotientNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != quotientNumSeps\n"+
			"Expected quotientNumSeps = '%v'\n"+
			"  Actual quotientNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), quotientNumSeps.String())

		return
	}

	if expectedModStr != moduloNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModStr != moduloNumberStr\n"+
			"Expected moduloNumberStr = '%v'\n"+
			"  Actual moduloNumberStr = '%v'\n\n",
			ePrefix, expectedModStr, moduloNumberStr)

		return
	}

	expectedEqualsModulo, err := expectedModuloBigINum.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsModulo, err := \n"+
			"  expectedModuloBigINum.Equal(modulo)\n"+
			"expectedModuloBigINum= '%v'\n"+
			"modulo= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModuloBigINumberStr, moduloNumberStr, err.Error())
		return
	}

	if !expectedEqualsModulo {
		t.Errorf("%v\n"+
			"Error: Expected Modulo and 'modulo' values NOT Equal!\n"+
			"Because expectedEqualsModulo = 'false' \n"+
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloBigINumberStr, moduloNumberStr)

		return
	}

	if !expectedNumSeps.Equal(moduloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps != moduloNumSeps\n"+
			"Expected moduloNumSeps = '%v'\n"+
			"  Actual moduloNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), moduloNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTenToPower_01(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTenToPower_01"

	var err error

	originalNumStr := "654.123"

	expectedNumStr := "0.654123"

	exponent := uint(3)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.DivideByTenToPower(exponent)\n"+
			"bINum Original= '%v'\n"+
			"exponent= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, exponent, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedNumStr= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTenToPower_02(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTenToPower_02"

	var err error

	originalNumStr := "-654.123"

	expectedNumStr := "-0.654123"

	exponent := uint(3)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.DivideByTenToPower(exponent)\n"+
			"bINum Original= '%v'\n"+
			"exponent= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, exponent, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedNumStr= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTenToPower_03(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTenToPower_03"

	var err error

	originalNumStr := "654123"

	expectedNumStr := "0.000654123"

	exponent := uint(9)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.DivideByTenToPower(exponent)\n"+
			"bINum Original= '%v'\n"+
			"exponent= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, exponent, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedNumStr= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTenToPower_04(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTenToPower_04"

	var err error

	originalNumStr := "654123"

	expectedNumStr := "654123"

	exponent := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.DivideByTenToPower(exponent)\n"+
			"bINum Original= '%v'\n"+
			"exponent= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, exponent, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedNumStr= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTenToPower_05(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTenToPower_05"

	var err error

	originalNumStr := "-654123"

	expectedNumStr := "-654123"

	exponent := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.DivideByTenToPower(exponent)\n"+
			"bINum Original= '%v'\n"+
			"exponent= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, exponent, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedNumStr= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_DivideByTenToPower_06(t *testing.T) {

	ePrefix := "TestBigIntNum_DivideByTenToPower_06"

	var err error

	originalNumStr := "654,123"

	expectedNumStr := "0,654123"

	exponent := uint(3)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr,&expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.DivideByTenToPower(exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.DivideByTenToPower(exponent)\n"+
			"bINum Original= '%v'\n"+
			"exponent= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumOriginalNumStr, exponent, err.Error())
		return
	}

	err = bINum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedNumStr= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_01(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_01"

	var err error

	originalNumStr := "654.123"

	expectedNumStr := "654.12300"

	extendPrecisionDigitsUint := uint(2)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_02(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_02"

	var err error

	originalNumStr := "-654.123"

	expectedNumStr := "-654.12300"

	extendPrecisionDigitsUint := uint(2)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_03(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_03"

	var err error

	originalNumStr := "7"

	expectedNumStr := "7.000"

	extendPrecisionDigitsUint := uint(3)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_04(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_04"

	var err error

	originalNumStr := "-654.123"

	expectedNumStr := "-654.12300000"

	extendPrecisionDigitsUint := uint(5)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_05(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_05"

	var err error

	originalNumStr := "654.123"

	expectedNumStr := "654.123000000"

	extendPrecisionDigitsUint := uint(6)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_06(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_06"

	var err error

	originalNumStr := "0"

	expectedNumStr := "0.00"

	extendPrecisionDigitsUint := uint(2)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_ExtendPrecision_07(t *testing.T) {

	ePrefix := "TestBigIntNum_ExtendPrecision_07"

	var err error

	originalNumStr := "-654,123"

	expectedNumStr := "-654,1230000"

	extendPrecisionDigitsUint := uint(4)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumOriginalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOriginalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumOriginalNumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumOriginalNumStr)

		return
	}

	err = bINum.ExtendPrecision(extendPrecisionDigitsUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.ExtendPrecision(extendPrecisionDigitsUint)\n"+
			"extendPrecisionDigitsUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, extendPrecisionDigitsUint, err.Error())
		return
	}

	err = bINum.IsValid("Validating Final bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating Final bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFinalNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFinalNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsBINum, err := expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINumFinalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumFinalNumStr, err.Error())
		return
	}

	if !expectedEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if expectedNumStr != bINumFinalNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != bINumFinalNumStr\n"+
			"Expected bINumFinalNumStr = '%v'\n"+
			"  Actual bINumFinalNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFinalNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}
