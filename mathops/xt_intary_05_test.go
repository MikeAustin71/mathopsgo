package mathops

import (
	"math/big"
	"testing"
)

func TestIntAry_MultiplyByTenToPower_01(t *testing.T) {

	nStr := "457.3"
	eNumStr := "45730"
	eIAry := []uint8{4, 5, 7, 3, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := new(IntAry).New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(2)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia intAry! index='%v'", i)
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_02(t *testing.T) {

	nStr := "457.3"
	power := uint(2)
	eNumStr := "45730"
	eIAry := []uint8{4, 5, 7, 3, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := new(IntAry).New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry! Index='%v'", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTenToPower_03(t *testing.T) {

	nStr := "457.3"
	power := uint(10)
	eNumStr := "4573000000000"
	eIAry := []uint8{4, 5, 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := new(IntAry).New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v'", i)
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_04(t *testing.T) {

	nStr := "457.3"
	power := uint(0)
	eNumStr := "457.3"
	eIAry := []uint8{4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	ia := new(IntAry).New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTenToPower_05(t *testing.T) {

	nStr := "-457.3"
	power := uint(1)
	eNumStr := "-4573"
	eIAry := []uint8{4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia := new(IntAry).New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTenToPower_06(t *testing.T) {

	nStr := "0"
	power := uint(2)
	eNumStr := "000"
	eIAry := []uint8{0, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := new(IntAry).New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTwoToPower_01(t *testing.T) {
	nStr1 := "23"
	expected := "12058624"
	power := uint(19)
	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= %v. Instead, ia.GetNumStr()= %v", expected, ia.GetNumStr())
	}

}

func TestIntAry_MultiplyByTwoToPower_02(t *testing.T) {
	nStr1 := "23"
	expected := "23"
	power := uint(0)
	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= %v. Instead, ia.GetNumStr()= %v", expected, ia.GetNumStr())
	}

}

func TestIntAry_MultiplyByTwoToPower_03(t *testing.T) {
	nStr1 := "23"
	expected := "46"
	power := uint(1)
	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= %v. Instead, ia.GetNumStr()= %v", expected, ia.GetNumStr())
	}

}

func TestIntAry_MultiplyThisBy_01(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("96.8524")
	ia2, _ := new(IntAry).NewNumStr("8574.21396845")
	expected := "830433.20095790678"

	this.MultiplyThisBy(&ia2, -1, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_02(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("-96.8524")
	ia2, _ := new(IntAry).NewNumStr("8574.21396845")
	expected := "-830433.20095790678"

	this.MultiplyThisBy(&ia2, -1, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_03(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("96.8524")
	ia2, _ := new(IntAry).NewNumStr("-8574.21396845")
	expected := "-830433.20095790678"

	this.MultiplyThisBy(&ia2, -1, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_04(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("-96.8524")
	ia2, _ := new(IntAry).NewNumStr("-8574.21396845")
	expected := "830433.20095790678"

	this.MultiplyThisBy(&ia2, -1, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_05(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("96.8524")
	ia2, _ := new(IntAry).NewNumStr("8574.21396845")
	expected := "830433.200958"
	maxPrecision := 6

	this.MultiplyThisBy(&ia2, maxPrecision, maxPrecision)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

	if maxPrecision != this.GetPrecision() {
		t.Errorf("Expected this.GetPrecisionInt()= %v.  Instead, this.GetPrecisionInt()= %v", maxPrecision, this.GetPrecision())
	}

}

func TestIntAry_MultiplyThisBy_06(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("0")
	ia2, _ := new(IntAry).NewNumStr("-8574.21396845")
	expected := "0.00000000"

	this.MultiplyThisBy(&ia2, 8, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_07(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("96.8524")
	ia2, _ := new(IntAry).NewNumStr("0")
	expected := "0"

	this.MultiplyThisBy(&ia2, 0, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_08(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("40")
	ia2, _ := new(IntAry).NewNumStr("2")
	expected := "80"

	this.MultiplyThisBy(&ia2, 0, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

	thisStats := this.GetIntAryStats()

	if thisStats.IntAryLen != 2 {
		t.Errorf("Expected this this IntAryLen = 2 .  Instead this IntAryLen= %v .", thisStats.IntAryLen)
	}

	if thisStats.IntegerLen != 2 {
		t.Errorf("Expected this this IntegerLen = 2 .  "+
			"Instead this IntegerLen= %v .", thisStats.IntegerLen)
	}

	if this.GetPrecision() != 0 {
		t.Errorf("Expected this this.GetPrecisionInt() = 0 .  "+
			"Instead this.GetPrecisionInt= %v .", this.GetPrecision())
	}
}

func TestIntAry_MultiplyThisBy_09(t *testing.T) {

	this, _ := new(IntAry).NewNumStr("999.99")
	ia2, _ := new(IntAry).NewNumStr("99.9")
	expected := "99899.001"

	this.MultiplyThisBy(&ia2, -1, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

	thisStats := this.GetIntAryStats()

	if thisStats.IntAryLen != 8 {
		t.Errorf("Expected this this IntAryLen = 8 .  Instead this IntAryLen= %v .", thisStats.IntAryLen)
	}

	if thisStats.IntegerLen != 5 {
		t.Errorf("Expected this this IntegerLen = 5 .  Instead this IntegerLen= %v .", thisStats.IntegerLen)
	}

	if this.GetPrecision() != 3 {
		t.Errorf("Expected this this.GetPrecisionInt() = 3 .  Instead this.GetPrecisionInt= %v .", this.GetPrecision())
	}

}

func TestIntAry_NewNumStr_01(t *testing.T) {
	nStr1 := "579.123456000"
	ePrecision := 9
	eSignVal := 1

	ia, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewNumStr_02(t *testing.T) {
	nStr1 := "-579.123456000"
	ePrecision := 9
	eSignVal := -1

	ia, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewNumStrWithNumSeps_01(t *testing.T) {

	nStr1 := "579,123456000"
	ePrecision := 9
	eSignVal := 1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	ia, err := new(IntAry).NewNumStrWithNumSeps(nStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	actualNumSeps := ia.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAry_NewNumStrWithNumSeps_02(t *testing.T) {

	nStr1 := "579.123456000"
	ePrecision := 9
	eSignVal := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	ia, err := new(IntAry).NewNumStrWithNumSeps(nStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	actualNumSeps := ia.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAry_NewNumStrWithNumSeps_03(t *testing.T) {

	nStr1 := "579.123456000"
	ePrecision := 9
	eSignVal := 1

	expectedNumSeps := NumericSeparatorDto{}

	ia, err := new(IntAry).NewNumStrWithNumSeps(nStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	actualNumSeps := ia.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAry_NewBigInt_01(t *testing.T) {
	num := big.NewInt(123456)
	precision := 3
	ia, err := new(IntAry).NewBigInt(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewBigInt(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewBigInt_02(t *testing.T) {
	num := big.NewInt(-123456)
	precision := 3
	ia, err := new(IntAry).NewBigInt(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewBigInt(num, precision). num= %v  precision= %v Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_01(t *testing.T) {
	num := big.NewFloat(123.456)
	precision := 3
	ia, err := new(IntAry).NewFloatBig(num, 3)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_02(t *testing.T) {
	num := big.NewFloat(-123.456)
	precision := 3
	ia, err := new(IntAry).NewFloatBig(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_03(t *testing.T) {
	num := big.NewFloat(123.456000)
	precision := 3
	ia, err := new(IntAry).NewFloatBig(num, -1)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_04(t *testing.T) {
	num := big.NewFloat(123.456700)
	precision := 3
	ia, err := new(IntAry).NewFloatBig(num, 3)
	eStr := "123.457"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloat64_01(t *testing.T) {
	num := float64(123.456111)
	precision := 3
	ia, err := new(IntAry).NewFloat64(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat64(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat64_02(t *testing.T) {
	num := float64(-123.456111)
	precision := 3
	ia, err := new(IntAry).NewFloat64(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat64(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat32_01(t *testing.T) {
	num := float64(123.456111)
	precision := 3
	ia, err := new(IntAry).NewFloat64(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat32(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat32_02(t *testing.T) {
	num := float64(-123.456111)
	precision := 3
	ia, err := new(IntAry).NewFloat64(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat32(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewInt_01(t *testing.T) {
	num := int(123456)
	precision := uint(3)
	ia := new(IntAry).NewInt(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt_02(t *testing.T) {
	num := int(-123456)
	precision := uint(3)
	ia := new(IntAry).NewInt(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewIntExponent_01(t *testing.T) {

	intNum := int(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_02(t *testing.T) {

	intNum := int(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_03(t *testing.T) {

	intNum := int(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_04(t *testing.T) {

	intNum := int(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_05(t *testing.T) {

	intNum := int(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_06(t *testing.T) {

	intNum := int(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32_01(t *testing.T) {
	num := int32(123456)
	precision := uint(3)
	eStr := "123.456"
	eSignVal := 1

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt32_02(t *testing.T) {

	num := int32(-123456)
	precision := uint(3)
	eStr := "-123.456"
	eSignVal := -1

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt32_03(t *testing.T) {

	num := int32(0)
	precision := uint(0)
	eStr := "0"

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewInt32_04(t *testing.T) {

	num := int32(0)
	precision := uint(3)
	eStr := "0.000"

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewInt32Exponent_01(t *testing.T) {

	int32Num := int32(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_02(t *testing.T) {

	int32Num := int32(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_03(t *testing.T) {

	int32Num := int32(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_04(t *testing.T) {

	int32Num := int32(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_05(t *testing.T) {

	int32Num := int32(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_06(t *testing.T) {

	int32Num := int32(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64_01(t *testing.T) {
	int64Num := int64(123456)
	precision := uint(3)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "123.456"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_02(t *testing.T) {
	int64Num := int64(-123456)
	precision := uint(3)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_03(t *testing.T) {
	int64Num := int64(0)
	precision := uint(3)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "0.000"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_04(t *testing.T) {
	int64Num := int64(0)
	precision := uint(0)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "0"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64Exponent_01(t *testing.T) {

	int64Num := int64(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_02(t *testing.T) {

	int64Num := int64(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_03(t *testing.T) {

	int64Num := int64(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_04(t *testing.T) {

	int64Num := int64(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_05(t *testing.T) {

	int64Num := int64(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_06(t *testing.T) {

	num := int64(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewInt64Exponent(num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewOne_01(t *testing.T) {

	expected := "1"
	ePrecision := 0

	ia := new(IntAry).NewOne(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewOne_02(t *testing.T) {

	expected := "1.00"
	ePrecision := 2

	ia := new(IntAry).NewOne(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewFive_01(t *testing.T) {

	expected := "5"
	ePrecision := 0

	ia := new(IntAry).NewFive(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewFive_02(t *testing.T) {

	expected := "5.00"
	ePrecision := 2

	ia := new(IntAry).NewFive(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTen_01(t *testing.T) {

	expected := "10"
	ePrecision := 0

	ia := new(IntAry).NewTen(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTen_02(t *testing.T) {

	expected := "10.00"
	ePrecision := 2

	ia := new(IntAry).NewTen(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewThree_01(t *testing.T) {

	expected := "3"
	ePrecision := 0

	ia := new(IntAry).NewThree(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewThree_02(t *testing.T) {

	expected := "3.00"
	ePrecision := 2

	ia := new(IntAry).NewThree(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTwo_01(t *testing.T) {

	expected := "2"
	ePrecision := 0

	ia := new(IntAry).NewTwo(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTwo_02(t *testing.T) {

	expected := "2.00"
	ePrecision := 2

	ia := new(IntAry).NewTwo(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewUint_01(t *testing.T) {

	ePrefix := "TestIntAry_NewUint_01"

	uintNum := uint(123456)
	precision := uint(3)
	signVal := 1

	ia, err := new(IntAry).NewUint(uintNum, signVal, precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint(uintNum, signVal, precision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	eStr := "123.456"

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n"+
			"Expected ia.GetNumStr()== %v\n"+
			"Instead ia.GetNumStr() == %v\n\n",
			ePrefix, iaNumStr, eStr)

		return
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("%v\n"+
			"Expected ia.GetPrecisionInt() == '%v'\n"+
			"Instead ia.GetPrecisionInt() == %v\n\n",
			ePrefix, precision, ia.GetPrecision())
	}

}

func TestIntAry_NewUint_02(t *testing.T) {
	uintNum := uint(123456)
	precision := uint(5)
	ia := new(IntAry).NewUint(uintNum, precision)
	eStr := "1.23456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint_03(t *testing.T) {
	uintNum := uint(0)
	precision := uint(3)
	ia := new(IntAry).NewUint(uintNum, precision)
	eStr := "0.000"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint_04(t *testing.T) {
	uintNum := uint(0)
	precision := uint(0)
	ia := new(IntAry).NewUint(uintNum, precision)
	eStr := "0"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUintExponent_01(t *testing.T) {

	uintNum := uint(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewUintExponent(uintNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUintExponent_02(t *testing.T) {

	uintNum := uint(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewUintExponent(uintNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUintExponent_03(t *testing.T) {

	uintNum := uint(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewUintExponent(uintNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUintExponent_04(t *testing.T) {

	uintNum := uint(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewUintExponent(uintNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUintExponent_05(t *testing.T) {

	uintNum := uint(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewUintExponent(uintNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUintExponent_06(t *testing.T) {

	num := uint(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewUintExponent(num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32_01(t *testing.T) {
	uint32Num := uint32(123456)
	precision := uint(3)
	ia := new(IntAry).NewUint32(uint32Num, precision)
	eStr := "123.456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewUint32_02(t *testing.T) {
	uint32Num := uint32(123456)
	precision := uint(5)
	ia := new(IntAry).NewUint32(uint32Num, precision)
	eStr := "1.23456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint32_03(t *testing.T) {
	uint32Num := uint32(0)
	precision := uint(3)
	ia := new(IntAry).NewUint32(uint32Num, precision)
	eStr := "0.000"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint32_04(t *testing.T) {
	uint32Num := uint32(0)
	precision := uint(0)
	ia := new(IntAry).NewUint32(uint32Num, precision)
	eStr := "0"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint32Exponent_01(t *testing.T) {

	uint32Num := uint32(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_02(t *testing.T) {

	uint32Num := uint32(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_03(t *testing.T) {

	uint32Num := uint32(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_04(t *testing.T) {

	uint32Num := uint32(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_05(t *testing.T) {

	uint32Num := uint32(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_06(t *testing.T) {

	num := uint32(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewUint32Exponent(num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint64_01(t *testing.T) {
	uint64Num := uint64(123456)
	precision := uint(3)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "123.456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewUint64_02(t *testing.T) {
	uint64Num := uint64(123456)
	precision := uint(5)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "1.23456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint64_03(t *testing.T) {
	uint64Num := uint64(0)
	precision := uint(3)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "0.000"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint64_04(t *testing.T) {
	uint64Num := uint64(0)
	precision := uint(0)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "0"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint64Exponent_01(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_01"

	uint64Num := uint64(123456)
	signValue := 1
	exponent := 3
	eStr := "123456.000"

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected ia.GetNumStr()== %v.\n"+
			"Instead ia.GetNumStr() == %v\n\n",
			ePrefix, eStr, iaNumStr)

	}

	return
}

func TestIntAry_NewUint64Exponent_02(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_02"

	uint64Num := uint64(123456)
	exponent := -3
	signValue := 1
	eStr := "123.456"

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected ia.GetNumStr()== '%v'\n"+
			"Instead ia.GetNumStr() == '%v'",
			ePrefix, eStr, iaNumStr)
	}
}

func TestIntAry_NewUint64Exponent_03(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_03"

	uint64Num := uint64(123456)
	exponent := 0
	eStr := "123456"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}


	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Error:\n"+
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}
}

func TestIntAry_NewUint64Exponent_04(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_04"
	uint64Num := uint64(0)
	exponent := 0
	eStr := "0"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Error:\n"+
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == %v\n\n",
			ePrefix, eStr, iaNumStr))
	}
	return
}

func TestIntAry_NewUint64Exponent_05(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_05"
	uint64Num := uint64(0)
	exponent := 3
	eStr := "0.000"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}

	return
}

func TestIntAry_NewUint64Exponent_06(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_06"
	num := uint64(0)
	exponent := -3
	eStr := "0.000"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}


	if eStr != iaNumStr {

		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)

	}

	return
}

func TestIntAry_OptimizeIntArrayLen_01(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456000"
	ePrecision := 9
	eLen := 12

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(false)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_02(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456"
	ePrecision := 6
	eLen := 9

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_03(t *testing.T) {
	nStr1 := "00579.000000000"
	expected := "579"
	ePrecision := 0
	eLen := 3

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_04(t *testing.T) {
	nStr1 := "00000.123450000"
	expected := "0.12345"
	ePrecision := 5
	eLen := 6

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_05(t *testing.T) {
	nStr1 := "00000.000123450000"
	expected := "0.00012345"
	ePrecision := 8
	eLen := 9

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v' .", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_06(t *testing.T) {
	nStr1 := "0.0"
	expected := "0.0"
	ePrecision := 1
	eLen := 2

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v' .", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}
