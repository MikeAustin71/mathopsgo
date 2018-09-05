package mathops

import (
	"testing"
)

func TestPrecisionScaleMode_String_01(t *testing.T) {

	r := SCALEPRECISIONRIGHT

	expectedStr := "ScalePrecisionRight"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected SCALEPRECISIONRIGHT string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestPrecisionScaleMode_String_02(t *testing.T) {

	r := SCALEPRECISIONLEFT
	expectedStr := "ScalePrecisionLeft"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected SCALEPRECISIONLEFT string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestPrecisionScaleMode_Value_01(t *testing.T) {

	var r PrecisionScaleMode

	var i int

	r = SCALEPRECISIONRIGHT

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'SCALEPRECISIONRIGHT' value = 0. Instead, got %v", i)
	}

}

func TestPrecisionScaleMode_Value_02(t *testing.T) {

	var r PrecisionScaleMode

	var i int

	r = SCALEPRECISIONLEFT

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'SCALEPRECISIONLEFT' value = 1. Instead, got %v", i)
	}

}
