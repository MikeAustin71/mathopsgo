package mathops

import "testing"

func TestIntAryMathAdd_Add_01(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "12.851"
	expectedStr := "58.613"

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

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr() )
	}
}

func TestIntAryMathAdd_Add_02(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "-12.851"
	expectedStr := "32.911"

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

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr() )
	}
}

func TestIntAryMathAdd_Add_03(t *testing.T) {

	nStr1 := "-45.762"
	nStr2 := "-12.851"
	expectedStr := "-58.613"

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

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr() )
	}
}

func TestIntAryMathAdd_Add_04(t *testing.T) {

	nStr1 := "0"
	nStr2 := "12.851"
	expectedStr := "12.851"

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

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr() )
	}
}

func TestIntAryMathAdd_Add_05(t *testing.T) {

	nStr1 := "45.762"
	nStr2 := "0"
	expectedStr := "45.762"

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

	ia3 := IntAryMathAdd{}.Add(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, ia3.GetNumStr() )
	}
}

func TestIntAryMathAdd_RunTotal_01(t *testing.T) {
	total := IntAry{}.NewZero(0)
	
	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

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
	
	ia3, err := IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). " +
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	ia4, err := IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). " +
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}
	
}

