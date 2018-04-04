package mathops

import "testing"

func TestNegativeValueFmtMode_String_01(t *testing.T) {

	r := LEADMINUSNEGVALFMTMODE

	expectedStr := "LeadingMinusSign"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected LEADMINUSNEGVALFMTMODE string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestNegativeValueFmtMode_String_02(t *testing.T) {

	r := PARENTHESESNEGVALFMTMODE

	expectedStr := "SurroundingParentheses"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected PARENTHESESNEGVALFMTMODE string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestNegativeValueFmtMode_Value_01(t *testing.T) {

	var r NegativeValueFmtMode

	var i int

	r = LEADMINUSNEGVALFMTMODE

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'LEADMINUSNEGVALFMTMODE' value = 0. Instead, got %v", i)
	}

}

func TestNegativeValueFmtMode_Value_02(t *testing.T) {

	var r NegativeValueFmtMode

	var i int

	r = PARENTHESESNEGVALFMTMODE

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'PARENTHESESNEGVALFMTMODE' value = 1. Instead, got %v", i)
	}

}
