package mathops

import (
	"testing"
	"math/big"
	"errors"
)

func TestBigIntNum_NewNumStr_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
	}

}