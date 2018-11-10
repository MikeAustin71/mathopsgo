package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMath_TruncateTrailingFractionalZeros_01(t *testing.T) {

	num := big.NewInt(56123456000)
	numPrecision := big.NewInt(9)

	expectedNum := big.NewInt(56123456)
	expectedNumPrecision := big.NewInt(6)

	actualNum, actualNumPrecision, err :=
		BigIntMath{}.TruncateTrailingFractionalZeros(num, numPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.TruncateTrailingFractionalZeros() " +
			"Error: %v", err.Error())
	}

	if expectedNum.Cmp(actualNum)!=0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNum.Text(10), actualNum.Text(10))
	}

	if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNumPrecision.Text(10), actualNumPrecision.Text(10))
	}

}

func TestBigIntMath_TruncateTrailingFractionalZeros_02(t *testing.T) {

	num := big.NewInt(-56123456000)
	numPrecision := big.NewInt(9)

	expectedNum := big.NewInt(-56123456)
	expectedNumPrecision := big.NewInt(6)

	actualNum, actualNumPrecision, err :=
		BigIntMath{}.TruncateTrailingFractionalZeros(num, numPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.TruncateTrailingFractionalZeros() " +
			"Error: %v", err.Error())
	}

	if expectedNum.Cmp(actualNum)!=0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNum.Text(10), actualNum.Text(10))
	}

	if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNumPrecision.Text(10), actualNumPrecision.Text(10))
	}

}

func TestBigIntMath_TruncateTrailingFractionalZeros_03(t *testing.T) {

	num := big.NewInt(56123456)
	numPrecision := big.NewInt(6)

	expectedNum := big.NewInt(56123456)
	expectedNumPrecision := big.NewInt(6)

	actualNum, actualNumPrecision, err :=
		BigIntMath{}.TruncateTrailingFractionalZeros(num, numPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.TruncateTrailingFractionalZeros() " +
			"Error: %v", err.Error())
	}

	if expectedNum.Cmp(actualNum)!=0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNum.Text(10), actualNum.Text(10))
	}

	if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNumPrecision.Text(10), actualNumPrecision.Text(10))
	}

}

func TestBigIntMath_TruncateTrailingFractionalZeros_04(t *testing.T) {

	num := big.NewInt(-56123456)
	numPrecision := big.NewInt(6)

	expectedNum := big.NewInt(-56123456)
	expectedNumPrecision := big.NewInt(6)

	actualNum, actualNumPrecision, err :=
		BigIntMath{}.TruncateTrailingFractionalZeros(num, numPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.TruncateTrailingFractionalZeros() " +
			"Error: %v", err.Error())
	}

	if expectedNum.Cmp(actualNum)!=0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNum.Text(10), actualNum.Text(10))
	}

	if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNumPrecision.Text(10), actualNumPrecision.Text(10))
	}

}

func TestBigIntMath_TruncateTrailingFractionalZeros_05(t *testing.T) {

	num := big.NewInt(0)
	numPrecision := big.NewInt(6)

	expectedNum := big.NewInt(0)
	expectedNumPrecision := big.NewInt(0)

	actualNum, actualNumPrecision, err :=
		BigIntMath{}.TruncateTrailingFractionalZeros(num, numPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.TruncateTrailingFractionalZeros() " +
			"Error: %v", err.Error())
	}

	if expectedNum.Cmp(actualNum)!=0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNum.Text(10), actualNum.Text(10))
	}

	if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
		t.Errorf("Expected actualNum='%v'. Instead, actualNum='%v'",
			expectedNumPrecision.Text(10), actualNumPrecision.Text(10))
	}

}
