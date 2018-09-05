package mathops

import (
	"testing"
)

func TestStrMathOp_AddN1N2_01(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "457.3"
	nStr2 := "22.2"
	expected := "479.5"
	nRunes := []rune("4795")
	eIAry := []int{4, 7, 9, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v",
			nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'",
			ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'",
			lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		element, _ := mOps.IFinal.GetIntAryRune(i)

		if nRunes[i] != element {
			t.Errorf("Error: Expected nRunes Array does NOT match ia.NumRunes Array! "+
				" nRunes[i]='%v' element='%v'",
				nRunes[i], element)
			return
		}

	}

	for i := 0; i < lEArray; i++ {

		element, _ := mOps.IFinal.GetIntAryElement(i)

		if eIAry[i] != int(element) {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_02(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "457.325"
	nStr2 := "-22.2"
	expected := "435.125"
	nRunes := []rune("435125")
	eIAry := []int{4, 3, 5, 1, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		element, _ := mOps.IFinal.GetIntAryRune(i)

		if nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		element, _ := mOps.IFinal.GetIntAryElement(i)
		if eIAry[i] != int(element) {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_03(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "-457.325"
	nStr2 := "22.2"
	expected := "-435.125"
	nRunes := []rune("435125")
	eIAry := []int{4, 3, 5, 1, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		element, _ := mOps.IFinal.GetIntAryRune(i)

		if nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {

		element, _ := mOps.IFinal.GetIntAryElement(i)

		if eIAry[i] != int(element) {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_04(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "-457.325"
	nStr2 := "-22.2"
	expected := "-479.525"
	nRunes := []rune("479525")
	eIAry := []int{4, 7, 9, 5, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	actualRunes := mOps.IFinal.GetRuneArray()

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != actualRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {

		element, _ := mOps.IFinal.GetIntAryElement(i)

		if eIAry[i] != int(element) {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_05(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "0.000"
	nStr2 := "-22.2"
	expected := "-22.200"
	nRunes := []rune("22200")
	eIAry := []int{2, 2, 2, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {

		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_06(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "0.000"
	nStr2 := "0"
	expected := "0.000"
	nRunes := []rune("0000")
	eIAry := []int{0, 0, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_07(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "99.225"
	nStr2 := "-99.1"
	expected := "0.125"
	nRunes := []rune("0125")
	eIAry := []int{0, 1, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_08(t *testing.T) {
	// N1 > N2 + and +
	mOps := StrMathOp{}.New()
	nStr1 := "350"
	nStr2 := "122"
	expected := "472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_09(t *testing.T) {
	// N1 > N2 - and +
	mOps := StrMathOp{}.New()
	nStr1 := "-350"
	nStr2 := "122"
	expected := "-228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_10(t *testing.T) {
	// N1 > N2 - and -
	mOps := StrMathOp{}.New()
	nStr1 := "-350"
	nStr2 := "-122"
	expected := "-472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_11(t *testing.T) {
	// N1 > N2 + and -
	mOps := StrMathOp{}.New()
	nStr1 := "350"
	nStr2 := "-122"
	expected := "228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_12(t *testing.T) {
	// N1 > N2  350 + 0
	mOps := StrMathOp{}.New()
	nStr1 := "350"
	nStr2 := "0"
	expected := "350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_13(t *testing.T) {
	// N1 > N2  -350 + 0
	mOps := StrMathOp{}.New()
	nStr1 := "-350"
	nStr2 := "0"
	expected := "-350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_14(t *testing.T) {
	// N2 > N1  + and +
	mOps := StrMathOp{}.New()
	nStr1 := "122"
	nStr2 := "350"
	expected := "472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_15(t *testing.T) {
	// N2 > N1  - and +
	mOps := StrMathOp{}.New()
	nStr1 := "-122"
	nStr2 := "350"
	expected := "228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_16(t *testing.T) {
	// N2 > N1  - and -
	mOps := StrMathOp{}.New()
	nStr1 := "-122"
	nStr2 := "-350"
	expected := "-472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_17(t *testing.T) {
	// N2 > N1  + and -
	mOps := StrMathOp{}.New()
	nStr1 := "122"
	nStr2 := "-350"
	expected := "-228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_18(t *testing.T) {
	// N2 > N1  0 and +350
	mOps := StrMathOp{}.New()
	nStr1 := "0"
	nStr2 := "350"
	expected := "350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_19(t *testing.T) {
	// N2 > N1  0 and -350
	mOps := StrMathOp{}.New()
	nStr1 := "0"
	nStr2 := "-350"
	expected := "-350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_20(t *testing.T) {
	// N1 == N2  +122 and +122
	mOps := StrMathOp{}.New()
	nStr1 := "122"
	nStr2 := "122"
	expected := "244"
	nRunes := []rune("244")
	eIAry := []int{2, 4, 4}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_21(t *testing.T) {
	// N1 == N2  -122 and +122
	mOps := StrMathOp{}.New()
	nStr1 := "-122"
	nStr2 := "122"
	expected := "0"
	nRunes := []rune("0")
	eIAry := []int{0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_22(t *testing.T) {
	// N1 == N2  -122 and -122
	mOps := StrMathOp{}.New()
	nStr1 := "-122"
	nStr2 := "-122"
	expected := "-244"
	nRunes := []rune("244")
	eIAry := []int{2, 4, 4}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_23(t *testing.T) {
	// N1 == N2  +122 and -122
	mOps := StrMathOp{}.New()
	nStr1 := "122"
	nStr2 := "-122"
	expected := "0"
	nRunes := []rune("0")
	eIAry := []int{0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_AddN1N2_24(t *testing.T) {
	// N1 == N2  0 and 0
	mOps := StrMathOp{}.New()
	nStr1 := "0"
	nStr2 := "0"
	expected := "0"
	nRunes := []rune("0")
	eIAry := []int{0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.AddN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_Divide_01(t *testing.T) {

	dividend := "-9360"
	divisor := "24.48"
	eQuotient := "-382.35294117647058823529411764706"
	eSignVal := -1
	maxPrecision := 29

	smop := StrMathOp{}.New()
	smop.Dividend.SetIntAryWithNumStr(dividend)
	smop.Divisor.SetIntAryWithNumStr(divisor)
	smop.Divide(maxPrecision)

	if eQuotient != smop.Quotient.GetNumStr() {
		t.Errorf("Error - Expected smop.Quotient.GetNumStr()= '%v'. Instead, smop.Quotient.GetNumStr()= '%v' .", eQuotient, smop.Quotient.GetNumStr())
	}

	if maxPrecision != smop.Quotient.GetPrecision() {
		t.Errorf("Error - Expected smop.Quotient.GetPrecision()= '%v'. Instead, smop.Quotient.GetPrecision()= '%v' .", maxPrecision, smop.Quotient.GetPrecision())
	}

	if eSignVal != smop.Quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, smop.Quotient.GetSign())
	}

}

func TestStrMathOp_Divide_02(t *testing.T) {

	dividend := "48"
	divisor := "24"
	eQuotient := "2"
	eSignVal := 1
	maxPrecision := 0

	smop := StrMathOp{}.New()
	smop.Dividend.SetIntAryWithNumStr(dividend)
	smop.Divisor.SetIntAryWithNumStr(divisor)
	smop.Divide(maxPrecision)

	if eQuotient != smop.Quotient.GetNumStr() {
		t.Errorf("Error - Expected smop.Quotient.GetNumStr()= '%v'. Instead, smop.Quotient.GetNumStr()= '%v' .", eQuotient, smop.Quotient.GetNumStr())
	}

	if maxPrecision != smop.Quotient.GetPrecision() {
		t.Errorf("Error - Expected smop.Quotient.GetPrecision()= '%v'. Instead, smop.Quotient.GetPrecision()= '%v' .", maxPrecision, smop.Quotient.GetPrecision())
	}

	if eSignVal != smop.Quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, smop.Quotient.GetSign())
	}

}

func TestStrMathOp_Divide_03(t *testing.T) {

	dividend := "54"
	divisor := "24"
	eQuotient := "2.25"
	eSignVal := 1
	maxPrecision := 7
	ePrecision := 2

	smop := StrMathOp{}.New()
	smop.Dividend.SetIntAryWithNumStr(dividend)
	smop.Divisor.SetIntAryWithNumStr(divisor)
	smop.Divide(maxPrecision)

	if eQuotient != smop.Quotient.GetNumStr() {
		t.Errorf("Error - Expected smop.Quotient.GetNumStr()= '%v'. Instead, smop.Quotient.GetNumStr()= '%v' .", eQuotient, smop.Quotient.GetNumStr())
	}

	if ePrecision != smop.Quotient.GetPrecision() {
		t.Errorf("Error - Expected smop.Quotient.GetPrecision()= '%v'. Instead, smop.Quotient.GetPrecision()= '%v' .", ePrecision, smop.Quotient.GetPrecision())
	}

	if eSignVal != smop.Quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, smop.Quotient.GetSign())
	}

}

func TestStrMathOp_Divide_04(t *testing.T) {

	dividend := "0"
	divisor := "24"
	eQuotient := "0"
	eSignVal := 1
	maxPrecision := 7
	ePrecision := 0

	smop := StrMathOp{}.New()
	smop.Dividend.SetIntAryWithNumStr(dividend)
	smop.Divisor.SetIntAryWithNumStr(divisor)
	smop.Divide(maxPrecision)

	if eQuotient != smop.Quotient.GetNumStr() {
		t.Errorf("Error - Expected smop.Quotient.GetNumStr()= '%v'. Instead, smop.Quotient.GetNumStr()= '%v' .", eQuotient, smop.Quotient.GetNumStr())
	}

	if ePrecision != smop.Quotient.GetPrecision() {
		t.Errorf("Error - Expected smop.Quotient.GetPrecision()= '%v'. Instead, smop.Quotient.GetPrecision()= '%v' .", ePrecision, smop.Quotient.GetPrecision())
	}

	if eSignVal != smop.Quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, smop.Quotient.GetSign())
	}

}

func TestStrMathOp_Divide_05(t *testing.T) {

	dividend := "5"
	divisor := "24"
	eQuotient := "0.20833333333333333333333333333333"
	eSignVal := 1
	maxPrecision := 32
	ePrecision := 32

	smop := StrMathOp{}.New()
	smop.Dividend.SetIntAryWithNumStr(dividend)
	smop.Divisor.SetIntAryWithNumStr(divisor)
	smop.Divide(maxPrecision)

	if eQuotient != smop.Quotient.GetNumStr() {
		t.Errorf("Error - Expected smop.Quotient.GetNumStr()= '%v'. Instead, smop.Quotient.GetNumStr()= '%v' .", eQuotient, smop.Quotient.GetNumStr())
	}

	if ePrecision != smop.Quotient.GetPrecision() {
		t.Errorf("Error - Expected smop.Quotient.GetPrecision()= '%v'. Instead, smop.Quotient.GetPrecision()= '%v' .", ePrecision, smop.Quotient.GetPrecision())
	}

	if eSignVal != smop.Quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, smop.Quotient.GetSign())
	}

}

func TestStrMathOp_Divide_06(t *testing.T) {

	dividend := "0.05"
	divisor := "24"
	eQuotient := "0.00208333333333333333333333333333"
	eSignVal := 1
	maxPrecision := 32
	ePrecision := 32

	smop := StrMathOp{}.New()
	smop.Dividend.SetIntAryWithNumStr(dividend)
	smop.Divisor.SetIntAryWithNumStr(divisor)
	smop.Divide(maxPrecision)

	if eQuotient != smop.Quotient.GetNumStr() {
		t.Errorf("Error - Expected smop.Quotient.GetNumStr()= '%v'. Instead, smop.Quotient.GetNumStr()= '%v' .", eQuotient, smop.Quotient.GetNumStr())
	}

	if ePrecision != smop.Quotient.GetPrecision() {
		t.Errorf("Error - Expected smop.Quotient.GetPrecision()= '%v'. Instead, smop.Quotient.GetPrecision()= '%v' .", ePrecision, smop.Quotient.GetPrecision())
	}

	if eSignVal != smop.Quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, smop.Quotient.GetSign())
	}

}

func TestStrMathOp_MultiplyN1N2_01(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "457.3"
	nStr2 := "22.2"
	expected := "10152.06"
	nRunes := []rune("1015206")
	eIAry := []int{1, 0, 1, 5, 2, 0, 6}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_02(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "457.3"
	nStr2 := "-22.2"
	expected := "-10152.06"
	nRunes := []rune("1015206")
	eIAry := []int{1, 0, 1, 5, 2, 0, 6}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := -1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_03(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "-457.3"
	nStr2 := "-22.2"
	expected := "10152.06"
	nRunes := []rune("1015206")
	eIAry := []int{1, 0, 1, 5, 2, 0, 6}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_04(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "-457.3"
	nStr2 := "0"
	expected := "0.0"
	nRunes := []rune("00")
	eIAry := []int{0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_05(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "0.0"
	nStr2 := "0.0"
	expected := "0.0"
	nRunes := []rune("00")
	eIAry := []int{0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_06(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "9999.9999"
	nStr2 := "9899.9899"
	expected := "98999898.01000101"
	nRunes := []rune("9899989801000101")
	eIAry := []int{9, 8, 9, 9, 9, 8, 9, 8, 0, 1, 0, 0, 0, 1, 0, 1}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 8
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_07(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "9899999.99991234"
	nStr2 := "7989899.98995678"
	expected := "79100009899871.7273668803886652"
	nRunes := []rune("791000098998717273668803886652")
	eIAry := []int{7, 9, 1, 0, 0, 0, 0, 9, 8, 9, 9, 8, 7, 1, 7, 2, 7, 3, 6, 6, 8, 8, 0, 3, 8, 8, 6, 6, 5, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 16
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_MultiplyN1N2_08(t *testing.T) {
	mOps := StrMathOp{}.New()
	nStr1 := "79100009899871.7273668803886652"
	nStr2 := "7989899.98995678"
	expected := "632001168304566313062.047887670481022949890056"
	nRunes := []rune("632001168304566313062047887670481022949890056")
	eIAry := []int{6, 3, 2, 0, 0, 1, 1, 6, 8, 3, 0, 4, 5, 6, 6, 3, 1, 3, 0, 6, 2, 0, 4, 7, 8, 8, 7, 6, 7, 0, 4, 8, 1, 0, 2, 2, 9, 4, 9, 8, 9, 0, 0, 5, 6}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 24
	eSignVal := 1

	mOps.N1.SetIntAryWithNumStr(nStr1)

	mOps.N2.SetIntAryWithNumStr(nStr2)

	err := mOps.MultiplyN1N2()

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := mOps.IFinal.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if mOps.IFinal.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
	}

	if eSignVal != mOps.IFinal.GetSign() {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
	}

	if lNRunes != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
	}

	if lEArray != mOps.IFinal.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
	}

	for i := 0; i < lNRunes; i++ {

		if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestStrMathOp_RaiseToPower_01(t *testing.T) {
	nStr1 := "625"
	power := 10
	expected := "9094947017729282379150390625"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_02(t *testing.T) {
	nStr1 := "625.25"
	power := 3
	expected := "244433710.953125"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_03(t *testing.T) {
	nStr1 := "5.3"
	power := 9
	expected := "3299763.591802133"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_04(t *testing.T) {
	nStr1 := "5"
	power := 0
	expected := "1"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_06(t *testing.T) {
	nStr1 := "5745"
	power := 1
	expected := "5745"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_07(t *testing.T) {
	nStr1 := "-625.25"
	power := 3
	expected := "-244433710.953125"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_08(t *testing.T) {
	nStr1 := "2"
	power := 8
	expected := "256"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_09(t *testing.T) {
	nStr1 := "0"
	power := 8
	expected := "0"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_RaiseToPower_10(t *testing.T) {
	nStr1 := "-25.25"
	power := 4
	expected := "406485.94140625"
	sMOp := StrMathOp{}.New()

	sMOp.N1.SetIntAryWithNumStr(nStr1)
	err := sMOp.RaiseToPower(power)
	if err != nil {

		t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

	}

	s := sMOp.IFinal.GetNumStr()
	if expected != s {
		t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
	}

}

func TestStrMathOp_SubtractN1N2_01(t *testing.T) {
	nStr1 := "900.777"
	nStr2 := "901.000"
	eNumStr := "-0.223"
	ePrecision := 3
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_02(t *testing.T) {
	nStr1 := "350"
	nStr2 := "122"
	eNumStr := "228"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_03(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "122"
	eNumStr := "-472"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_04(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "-122"
	eNumStr := "-228"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_05(t *testing.T) {
	nStr1 := "350"
	nStr2 := "-122"
	eNumStr := "472"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_06(t *testing.T) {
	nStr1 := "350"
	nStr2 := "0"
	eNumStr := "350"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_07(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "0"
	eNumStr := "-350"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_08(t *testing.T) {
	nStr1 := "122"
	nStr2 := "350"
	eNumStr := "-228"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_09(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "350"
	eNumStr := "-472"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_10(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "-350"
	eNumStr := "228"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_11(t *testing.T) {
	nStr1 := "122"
	nStr2 := "-350"
	eNumStr := "472"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_12(t *testing.T) {
	nStr1 := "0"
	nStr2 := "350"
	eNumStr := "-350"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_13(t *testing.T) {
	nStr1 := "0"
	nStr2 := "-350"
	eNumStr := "350"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_14(t *testing.T) {
	nStr1 := "122"
	nStr2 := "122"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_15(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "122"
	eNumStr := "-244"
	ePrecision := 0
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_16(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "-122"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_17(t *testing.T) {
	nStr1 := "122"
	nStr2 := "-122"
	eNumStr := "244"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_18(t *testing.T) {
	nStr1 := "0"
	nStr2 := "0"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}

func TestStrMathOp_SubtractN1N2_19(t *testing.T) {
	nStr1 := "1.122"
	nStr2 := "4.5"
	eNumStr := "-3.378"
	ePrecision := 3
	eSignVal := -1

	smop := StrMathOp{}.New()
	smop.N1.SetIntAryWithNumStr(nStr1)
	smop.N2.SetIntAryWithNumStr(nStr2)
	err := smop.SubtractN1N2()

	if err != nil {
		t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
	}

	if eNumStr != smop.IFinal.GetNumStr() {
		t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
	}

	if ePrecision != smop.IFinal.GetPrecision() {
		t.Errorf("Error - Expected IFinal.GetPrecision()= '%v' .  Instead, IFinal.GetPrecision()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
	}

	if eSignVal != smop.IFinal.GetSign() {
		t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
	}

}
