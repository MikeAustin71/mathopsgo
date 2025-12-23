package mathops

import "testing"

func TestNumStrFmtMode_String_01(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_String_01"

	r := PUREINTEGERFMT

	expectedStr := "PureIntegerString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("%v\n"+
			"Expected PUREINTEGERFMT string='%v'.\n"+
			"Instead, string='%v'\n",
			ePrefix, expectedStr, s)
	}

	return
}

func TestNumStrFmtMode_String_02(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_String_02"

	r := INTSTRDECIMALFMT

	expectedStr := "IntegerDecimalString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("%v\n"+
			"Expected INTSTRDECIMALFMT string='%v'.\n"+
			"Instead, string='%v'\n",
			ePrefix, expectedStr, s)
	}

	return
}

func TestNumStrFmtMode_String_03(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_String_03"

	r := THOUSANDSNUMSTRFMT

	expectedStr := "ThousandsNumString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("%v\n"+
			"Expected THOUSANDSNUMSTRFMT string='%v'.\n"+
			"Instead, string='%v'\n",
			ePrefix, expectedStr, s)
	}

	return
}

func TestNumStrFmtMode_String_04(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_String_04"

	r := CURRENCYNUMSTRFMT

	expectedStr := "CurrencyNumString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("%v\n"+
			"Expected THOUSANDSNUMSTRFMT string='%v'.\n"+
			"Instead, string='%v'\n",
			ePrefix, expectedStr, s)
	}

	return
}

func TestNumStrFmtMode_Value_01(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_Value_01"

	var r NumStrFmtMode

	var i int

	r = PUREINTEGERFMT

	i = int(r)

	if i != 0 {
		t.Errorf("%v\n"+
			"Expected 'PUREINTEGERFMT' value = 0.\n"+
			"Instead, received value= %v",
			ePrefix, i)
	}

	return
}

func TestNumStrFmtMode_Value_02(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_Value_02"

	var r NumStrFmtMode

	var i int

	r = INTSTRDECIMALFMT

	i = int(r)

	if i != 1 {
		t.Errorf("%v\n"+
			"Expected 'INTSTRDECIMALFMT' value = 1.\n"+
			"Instead, received value= %v",
			ePrefix, i)
	}

	return
}

func TestNumStrFmtMode_Value_03(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_06"

	var r NumStrFmtMode

	var i int

	r = THOUSANDSNUMSTRFMT

	i = int(r)

	if i != 2 {
		t.Errorf("%v\n"+
			"Expected 'THOUSANDSNUMSTRFMT' value = 2.\n"+
			"Instead, got %v", ePrefix, i)
	}

	return
}

func TestNumStrFmtMode_Value_04(t *testing.T) {

	ePrefix := "TestNumStrFmtMode_Value_04"

	var r NumStrFmtMode

	var i int

	r = CURRENCYNUMSTRFMT

	i = int(r)

	if i != 3 {
		t.Errorf("%v\nExpected 'CURRENCYNUMSTRFMT' value = 3.\n"+
			"Instead, got %v",
			ePrefix, i)
	}

	return
}
