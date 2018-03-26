
package mathops

import (
	"testing"
)

func TestFracIntAry_GetRationalValue_01(t *testing.T) {
	numStrNum := "2"
	numStrDenom := "3"
  eFrac := "0.66666666666666666666666666666667"
  precision := 32
	fIary, _ := FracIntAry{}.NewNumStrs(numStrNum, numStrDenom)

	ratNum, err := fIary.GetRationalValue(precision)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	num := ratNum.Num()

	if numStrNum != num.String() {
		t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
	}

	denom := ratNum.Denom()

	if numStrDenom != denom.String() {
		t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
	}

	fStr := ratNum.FloatString(precision)

	if eFrac != fStr {
		t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
	}

}

func TestFracIntAry_GetRationalValue_02(t *testing.T) {
	numStrNum := "3"
	numStrDenom := "4"
  eFrac := "0.75"
  precision:= 2
	fIary, _ := FracIntAry{}.NewNumStrs(numStrNum, numStrDenom)

	ratNum, err := fIary.GetRationalValue(precision)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	num := ratNum.Num()

	if numStrNum != num.String() {
		t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
	}

	denom := ratNum.Denom()

	if numStrDenom != denom.String() {
		t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
	}

	fStr := ratNum.FloatString(2)

	if eFrac != fStr {
		t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
	}

}

func TestFracIntAry_GetRationalValue_03(t *testing.T) {
	numStrNum := "1000"
	eNum := "1"
	numStrDenom := "2000"
	eDenom:= "2"
  eFrac := "0.5"
  precision := 1
	fIary, _ := FracIntAry{}.NewNumStrs(numStrNum, numStrDenom)

	ratNum, err := fIary.GetRationalValue(precision)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	num := ratNum.Num()

	if eNum != num.String() {
		t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
	}

	denom := ratNum.Denom()

	if eDenom != denom.String() {
		t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
	}

	fStr := ratNum.FloatString(precision)

	if eFrac != fStr {
		t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
	}

}

func TestFracIntAry_GetRationalValue_04(t *testing.T) {
	numStrNum := "9.24"
	eNum := "291115311909262759924385633270321361058601"
	numStrDenom := "15.87"
	eDenom:= "500000000000000000000000000000000000000000"
  eFrac := "0.582230623818525519848771266540642722117202"
  precision := 42
	fIary, _ := FracIntAry{}.NewNumStrs(numStrNum, numStrDenom)

	ratNum, err := fIary.GetRationalValue(precision)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	num := ratNum.Num()

	if eNum != num.String() {
		t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", eNum, num.String())
	}

	denom := ratNum.Denom()

	if eDenom != denom.String() {
		t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
	}

	fStr := ratNum.FloatString(precision)

	if eFrac != fStr {
		t.Errorf("Expected %v-digit decimal string = %v  .   Instead, decimal string = %v .",precision, eFrac, fStr)
	}

}

func TestFracIntAry_GetRationalValue_05(t *testing.T) {
	numStrNum := "3"
	numStrDenom := "4"
	eFrac := "0.75"
	precision:= 2
	iaNum, _ := IntAry{}.NewNumStr(numStrNum)

	iaDenom, _ := IntAry{}.NewNumStr(numStrDenom)

	fIary := FracIntAry{}.NewIntArys(&iaNum, &iaDenom)

	ratNum, err := fIary.GetRationalValue(precision)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	num := ratNum.Num()

	if numStrNum != num.String() {
		t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
	}

	denom := ratNum.Denom()

	if numStrDenom != denom.String() {
		t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
	}

	fStr := ratNum.FloatString(precision)

	if eFrac != fStr {
		t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
	}

}

func TestFracIntAry_GetRationalValue_06(t *testing.T) {
	numStrNum := "-3"
	numStrDenom := "4"
	eFrac := "-0.75"
	precision:= 2
	iaNum, _ := IntAry{}.NewNumStr(numStrNum)

	iaDenom, _ := IntAry{}.NewNumStr(numStrDenom)

	fIary := FracIntAry{}.NewIntArys(&iaNum, &iaDenom)

	ratNum, err := fIary.GetRationalValue(precision)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	num := ratNum.Num()

	if numStrNum != num.String() {
		t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
	}

	denom := ratNum.Denom()

	if numStrDenom != denom.String() {
		t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
	}

	fStr := ratNum.FloatString(precision)

	if eFrac != fStr {
		t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
	}

}

func TestFracIntAry_GetRationalValue_07(t *testing.T) {
	numStrNum := "3.1"
	numStrDenom := "4.2"
	eFrac := "0.738095238095"
	precision := 12
	iaNum, _ := IntAry{}.NewNumStr(numStrNum)

	iaDenom, _ := IntAry{}.NewNumStr(numStrDenom)

	fIary := FracIntAry{}.NewIntArys(&iaNum, &iaDenom)

	ratNum, err := fIary.GetRationalValue(-1)

	if err != nil {
		t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
	}

	fStr := ratNum.FloatString(precision)

	if eFrac != fStr {
		t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
	}

}

