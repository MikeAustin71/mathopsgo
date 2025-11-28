package mathops

import "testing"

func TestNegativeValueFmtMode_String_01(t *testing.T) {

  ePrefix := "TestNegativeValueFmtMode_String_01"

  r := LEADMINUSNEGVALFMTMODE

  expectedStr := "LeadingMinusSign"

  s := r.String()

  if expectedStr != s {
    t.Errorf("%v\n"+
      "Expected LEADMINUSNEGVALFMTMODE string='%v'.\n"+
      "Instead, string='%v'\n\n",
      ePrefix, expectedStr, s)

    return
  }

  return
}

func TestNegativeValueFmtMode_String_02(t *testing.T) {

  ePrefix := "TestNegativeValueFmtMode_String_02"

  r := PARENTHESESNEGVALFMTMODE

  expectedStr := "SurroundingParentheses"

  s := r.String()

  if expectedStr != s {

    t.Errorf("%v\n"+
      "Expected PARENTHESESNEGVALFMTMODE string='%v'.\n"+
      "Instead, string='%v'\n\n",
      ePrefix, expectedStr, s)

    return
  }

  return
}

func TestNegativeValueFmtMode_String_03(t *testing.T) {

  ePrefix := "TestNegativeValueFmtMode_String_03"

  r := ABSOLUTEPURENUMSTRFMTMODE

  expectedStr := "AbsolutePureNumberString"

  s := r.String()

  if expectedStr != s {

    t.Errorf("%v\n"+
      "Expected ABSOLUTEPURENUMSTRFMTMODE string='%v'.\n"+
      "Instead, string='%v'\n\n",
      ePrefix, expectedStr, s)

    return
  }

  return
}

func TestNegativeValueFmtMode_Value_01(t *testing.T) {

  ePrefix := "TestNegativeValueFmtMode_Value_01"

  var r NegativeValueFmtMode

  var i int

  r = LEADMINUSNEGVALFMTMODE

  i = int(r)

  if i != 0 {

    t.Errorf("%v\n"+
      "Expected 'LEADMINUSNEGVALFMTMODE' value = 0.\n"+
      "Instead, value = %v\n\n",
      ePrefix, i)

    return
  }

  return
}

func TestNegativeValueFmtMode_Value_02(t *testing.T) {

  ePrefix := "TestNegativeValueFmtMode_Value_02"

  var r NegativeValueFmtMode

  var i int

  r = PARENTHESESNEGVALFMTMODE

  i = int(r)

  if i != 1 {

    t.Errorf("%v\n"+
      "Expected 'PARENTHESESNEGVALFMTMODE' value = 1.\n"+
      "Instead, value = %v\n\n",
      ePrefix, i)

    return
  }

  return
}

func TestNegativeValueFmtMode_Value_03(t *testing.T) {

  ePrefix := "TestNegativeValueFmtMode_Value_03"

  var r NegativeValueFmtMode

  var i int

  r = ABSOLUTEPURENUMSTRFMTMODE

  i = int(r)

  if i != 2 {

    t.Errorf("%v\n"+
      "Expected 'ABSOLUTEPURENUMSTRFMTMODE' value = 2.\n"+
      "Instead, got %v\n\n",
      ePrefix, i)

    return

  }

  return
}
