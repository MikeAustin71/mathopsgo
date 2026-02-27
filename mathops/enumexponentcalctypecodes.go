package mathops

type ExponentCalcTypeCode int

// String returns a string representation of the ExponentCalcTypeCode
// or "Unknown" for invalid codes.
// Example:
//
//	code := ExponentCalcTypeCode(1)
//	fmt.Println(code.String())
//	Output:
//	BasePlusIntExpoPlusInt
func (code ExponentCalcTypeCode) String() string {
	if int(code) < 0 || int(code) >= len(ExponentCalcTypeCodeLabels) {
		return "Unknown"
	}
	return ExponentCalcTypeCodeLabels[code]
}

const (
	ExpoCalcBasePlusIntExpoPlusInt ExponentCalcTypeCode = iota + 1
	ExpoCalcBasePlusIntExpoPlusFrac
	ExpoCalcBasePlusIntExpoMinusInt
	ExpoCalcBasePlusIntExpoMinusFrac
	ExpoCalcBasePlusFracExpoPlusInt
	ExpoCalcBasePlusFracExpoPlusFrac
	ExpoCalcBasePlusFracExpoMinusInt
	ExpoCalcBasePlusFracExpoMinusFrac
	ExpoCalcBaseMinusIntExpoPlusInt
	ExpoCalcBaseMinusIntExpoPlusFrac
	ExpoCalcBaseMinusIntExpoMinusInt
	ExpoCalcBaseMinusIntExpoMinusFrac
	ExpoCalcBaseMinusFracExpoPlusInt
	ExpoCalcBaseMinusFracExpoPlusFrac
	ExpoCalcBaseMinusFracExpoMinusInt
	ExpoCalcBaseMinusFracExpoMinusFrac
)

var ExponentCalcTypeCodeLabels = []string{"Invalid-None",
	"BasePlusIntExpoPlusInt",     // 1
	"BasePlusIntExpoPlusFrac",    // 2
	"BasePlusIntExpoMinusInt",    // 3
	"BasePlusIntExpoMinusFrac",   // 4
	"BasePlusFracExpoPlusInt",    // 5
	"BasePlusFracExpoPlusFrac",   // 6
	"BasePlusFracExpoMinusInt",   // 7
	"BasePlusFracExpoMinusFrac",  // 8
	"BaseMinusIntExpoPlusInt",    // 9
	"BaseMinusIntExpoPlusFrac",   // 10
	"BaseMinusIntExpoMinusInt",   // 11
	"BaseMinusIntExpoMinusFrac",  // 12
	"BaseMinusFracExpoPlusInt",   // 13
	"BaseMinusFracExpoPlusFrac",  // 14
	"BaseMinusFracExpoMinusInt",  // 15
	"BaseMinusFracExpoMinusFrac"} // 16
