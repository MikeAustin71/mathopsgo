package mathops

import "testing"

func TestExponentCalcTypeCode_AllValues(t *testing.T) {

	ePrefix := "TestExponentCalcTypeCode_AllValues"

	tests := []struct {
		name     string
		code     ExponentCalcTypeCode
		expected int
		label    string
	}{
		{"ExpoCalcBasePlusIntExpoPlusInt", ExpoCalcBasePlusIntExpoPlusInt, 1, "BasePlusIntExpoPlusInt"},
		{"ExpoCalcBasePlusIntExpoPlusFrac", ExpoCalcBasePlusIntExpoPlusFrac, 2, "BasePlusIntExpoPlusFrac"},
		{"ExpoCalcBasePlusIntExpoMinusInt", ExpoCalcBasePlusIntExpoMinusInt, 3, "BasePlusIntExpoMinusInt"},
		{"ExpoCalcBasePlusIntExpoMinusFrac", ExpoCalcBasePlusIntExpoMinusFrac, 4, "BasePlusIntExpoMinusFrac"},
		{"ExpoCalcBasePlusFracExpoPlusInt", ExpoCalcBasePlusFracExpoPlusInt, 5, "BasePlusFracExpoPlusInt"},
		{"ExpoCalcBasePlusFracExpoPlusFrac", ExpoCalcBasePlusFracExpoPlusFrac, 6, "BasePlusFracExpoPlusFrac"},
		{"ExpoCalcBasePlusFracExpoMinusInt", ExpoCalcBasePlusFracExpoMinusInt, 7, "BasePlusFracExpoMinusInt"},
		{"ExpoCalcBasePlusFracExpoMinusFrac", ExpoCalcBasePlusFracExpoMinusFrac, 8, "BasePlusFracExpoMinusFrac"},
		{"ExpoCalcBaseMinusIntExpoPlusInt", ExpoCalcBaseMinusIntExpoPlusInt, 9, "BaseMinusIntExpoPlusInt"},
		{"ExpoCalcBaseMinusIntExpoPlusFrac", ExpoCalcBaseMinusIntExpoPlusFrac, 10, "BaseMinusIntExpoPlusFrac"},
		{"ExpoCalcBaseMinusIntExpoMinusInt", ExpoCalcBaseMinusIntExpoMinusInt, 11, "BaseMinusIntExpoMinusInt"},
		{"ExpoCalcBaseMinusIntExpoMinusFrac", ExpoCalcBaseMinusIntExpoMinusFrac, 12, "BaseMinusIntExpoMinusFrac"},
		{"ExpoCalcBaseMinusFracExpoPlusInt", ExpoCalcBaseMinusFracExpoPlusInt, 13, "BaseMinusFracExpoPlusInt"},
		{"ExpoCalcBaseMinusFracExpoPlusFrac", ExpoCalcBaseMinusFracExpoPlusFrac, 14, "BaseMinusFracExpoPlusFrac"},
		{"ExpoCalcBaseMinusFracExpoMinusInt", ExpoCalcBaseMinusFracExpoMinusInt, 15, "BaseMinusFracExpoMinusInt"},
		{"ExpoCalcBaseMinusFracExpoMinusFrac", ExpoCalcBaseMinusFracExpoMinusFrac, 16, "BaseMinusFracExpoMinusFrac"},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if int(tt.code) != tt.expected {
				t.Errorf("Error Returned by: %v\n"+
					"Expected value %d, got %d\n",
					ePrefix, tt.expected, tt.code)
			}

			if tt.code.String() != tt.label {
				t.Errorf("Error Returned by: %v\n"+"Expected label %q, got %q\n", ePrefix, tt.label, tt.code.String())
			}
		})
	}

}
