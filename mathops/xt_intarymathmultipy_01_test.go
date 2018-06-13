package mathops

import "testing"

func TestIntAryMathMultiply_MultiplyInPlace_01(t *testing.T) {
	nStr1 := "45.762"
	nStr2 := "12.851"
	expectedStr := "588.087462"
	minimumPrecision := 6
	maximumPrecision := 6

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). " +
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}


	IntAryMathMultiply{}.MultiplyInPlace(&ia1, &ia2, minimumPrecision, maximumPrecision)

	if expectedStr != ia1.GetNumStr() {
		t.Errorf("Error: Expected multiplication result='%v'. Instead, result='%v'",
			expectedStr, ia1.GetNumStr())
	}

}

func TestIntAryMathMultiply_MultiplyInPlace_02(t *testing.T) {
	nStr1 := "-45"
	nStr2 := "35"
	expectedStr := "-1575"
	minimumPrecision := -1
	maximumPrecision := -1

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). " +
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}


	IntAryMathMultiply{}.MultiplyInPlace(&ia1, &ia2, minimumPrecision, maximumPrecision)

	if expectedStr != ia1.GetNumStr() {
		t.Errorf("Error: Expected multiplication result='%v'. Instead, result='%v'",
			expectedStr, ia1.GetNumStr())
	}

}

func TestIntAryMathMultiply_Multiply_01(t *testing.T) {
	nStr1 := "92.135"
	nStr2 := "5.76405"
	expectedStr := "531.07074675"
	minimumPrecision := -1
	maximumPrecision := -1

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). " +
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAry{}.NewZero(0)

	IntAryMathMultiply{}.Multiply(&ia1, &ia2, &ia3, minimumPrecision, maximumPrecision)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected multiplication result='%v'. Instead, result='%v'",
			expectedStr, ia1.GetNumStr())
	}

}