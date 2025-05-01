package mathops

import "bytes"

// BackUpIntAry
// This structure and associated methods are used to back up and store
// a copy of an intArray type.
type BackUpIntAry struct {
	intAry                 []uint8
	intAryLen              int
	integerLen             int
	significantIntegerLen  int
	significantFractionLen int
	firstDigitIdx          int
	lastDigitIdx           int
	isZeroValue            bool
	isIntegerZeroValue     bool
	precision              int
	signVal                int
	decimalSeparator       rune // https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
	thousandsSeparator     rune // https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
	currencySymbol         rune // See Currency Symbol References below:
	// https://gist.github.com/bzerangue/5484121
	// http://symbologic.info/currency.htm
	// http://www.xe.com/symbols.php
}

func (iBa *BackUpIntAry) New() BackUpIntAry {
	iAry := BackUpIntAry{}

	iAry.intAry = []uint8{}
	iAry.intAryLen = 0
	iAry.integerLen = 0
	iAry.significantIntegerLen = 0
	iAry.significantFractionLen = 0
	iAry.firstDigitIdx = -1
	iAry.lastDigitIdx = -1
	iAry.isZeroValue = true
	iAry.isIntegerZeroValue = true
	iAry.precision = 0
	iAry.signVal = 1
	iAry.decimalSeparator = '.'
	iAry.thousandsSeparator = ','
	iAry.currencySymbol = '$'

	return iAry
}

func (iBa *BackUpIntAry) Empty() {
	iBa.intAry = []uint8{}
	iBa.intAryLen = 0
	iBa.integerLen = 0
	iBa.significantIntegerLen = 0
	iBa.significantFractionLen = 0
	iBa.firstDigitIdx = -1
	iBa.lastDigitIdx = -1
	iBa.isZeroValue = true
	iBa.isIntegerZeroValue = true
	iBa.precision = 0
	iBa.signVal = 1
	iBa.decimalSeparator = '.'
	iBa.thousandsSeparator = ','
	iBa.currencySymbol = '$'

}

func (iBa *BackUpIntAry) CopyIn(iBa2 *BackUpIntAry) {
	iBa2.SetInternalFlags()
	iBa.Empty()

	iBa.intAry = make([]uint8, iBa2.intAryLen)
	for i := 0; i < iBa2.intAryLen; i++ {
		iBa.intAry[i] = iBa2.intAry[i]
	}

	iBa.intAryLen = iBa2.intAryLen
	iBa.integerLen = iBa2.integerLen
	iBa.significantIntegerLen = iBa2.significantIntegerLen
	iBa.significantFractionLen = iBa2.significantFractionLen
	iBa.firstDigitIdx = iBa2.firstDigitIdx
	iBa.lastDigitIdx = iBa2.lastDigitIdx
	iBa.isZeroValue = iBa2.isZeroValue
	iBa.isIntegerZeroValue = iBa2.isIntegerZeroValue
	iBa.precision = iBa2.precision
	iBa.signVal = iBa2.signVal
	iBa.decimalSeparator = iBa2.decimalSeparator
	iBa.thousandsSeparator = iBa2.thousandsSeparator
	iBa.currencySymbol = iBa2.currencySymbol

}

func (iBa *BackUpIntAry) CopyOut() BackUpIntAry {

	iBa.SetInternalFlags()

	iAry2 := new(BackUpIntAry).New()

	iAry2.intAry = make([]uint8, iBa.intAryLen)

	for i := 0; i < iBa.intAryLen; i++ {
		iAry2.intAry[i] = iBa.intAry[i]
	}

	iAry2.intAryLen = iBa.intAryLen
	iAry2.integerLen = iBa.integerLen
	iAry2.significantIntegerLen = iBa.significantIntegerLen
	iAry2.significantFractionLen = iBa.significantFractionLen
	iAry2.firstDigitIdx = iBa.firstDigitIdx
	iAry2.lastDigitIdx = iBa.lastDigitIdx
	iAry2.isZeroValue = iBa.isZeroValue
	iAry2.isIntegerZeroValue = iBa.isIntegerZeroValue
	iAry2.precision = iBa.precision
	iAry2.signVal = iBa.signVal
	iAry2.decimalSeparator = iBa.decimalSeparator
	iAry2.thousandsSeparator = iBa.thousandsSeparator
	iAry2.currencySymbol = iBa.currencySymbol

	return iAry2
}

func (iBa *BackUpIntAry) Equals(iBa2 *BackUpIntAry) bool {

	iBa.SetInternalFlags()

	iBa2.SetInternalFlags()

	if iBa.intAryLen != iBa2.intAryLen {
		return false
	}

	for i := 0; i < iBa2.intAryLen; i++ {
		if iBa.intAry[i] != iBa2.intAry[i] {
			return false
		}
	}

	if iBa.integerLen != iBa2.integerLen ||
		iBa.significantIntegerLen != iBa2.significantIntegerLen ||
		iBa.significantFractionLen != iBa2.significantFractionLen ||
		iBa.firstDigitIdx != iBa2.firstDigitIdx ||
		iBa.lastDigitIdx != iBa2.lastDigitIdx ||
		iBa.isZeroValue != iBa2.isZeroValue ||
		iBa.isIntegerZeroValue != iBa2.isIntegerZeroValue ||
		iBa.precision != iBa2.precision ||
		iBa.signVal != iBa2.signVal ||
		iBa.decimalSeparator != iBa2.decimalSeparator ||
		iBa.thousandsSeparator != iBa2.thousandsSeparator ||
		iBa.currencySymbol != iBa2.currencySymbol {

		return false
	}

	return true
}

func (iBa *BackUpIntAry) GetIntAryStats() IntAryStatsDto {

	iBa.SetInternalFlags()

	iStats := IntAryStatsDto{}

	iStats.IntAryLen = iBa.intAryLen
	iStats.IntegerLen = iBa.integerLen
	iStats.SignificantIntegerLen = iBa.significantIntegerLen
	iStats.SignificantFractionLen = iBa.significantFractionLen
	iStats.Precision = iBa.precision
	iStats.SignVal = iBa.signVal
	iStats.FirstDigitIdx = iBa.firstDigitIdx
	iStats.LastDigitIdx = iBa.lastDigitIdx
	iStats.IsZeroValue = iBa.isZeroValue
	iStats.IsIntegerZeroValue = iBa.isIntegerZeroValue
	iStats.DecimalSeparator = iBa.decimalSeparator
	iStats.ThousandsSeparator = iBa.thousandsSeparator
	iStats.CurrencySymbol = iBa.currencySymbol

	return iStats

}

func (iBa *BackUpIntAry) GetNumStr() string {

	iBa.SetInternalFlags()

	if iBa.decimalSeparator == 0 {
		iBa.decimalSeparator = '.'
	}

	var buffer bytes.Buffer

	if iBa.signVal < 0 {
		buffer.WriteRune('-')
	}

	intLen := iBa.intAryLen - iBa.precision

	for i := 0; i < intLen; i++ {
		buffer.WriteRune(rune(iBa.intAry[i] + 48))
	}

	if iBa.precision > 0 {
		buffer.WriteRune(iBa.decimalSeparator)

		for j := 0; j < iBa.precision; j++ {
			buffer.WriteRune(rune(iBa.intAry[intLen] + 48))
			intLen++
		}

	}

	return buffer.String()

}

func (iBa *BackUpIntAry) GetPrecision() int {
	return iBa.precision
}

func (iBa *BackUpIntAry) GetSignValue() int {
	return iBa.signVal
}

func (iBa *BackUpIntAry) SetInternalFlags() {

	iBa.intAryLen = len(iBa.intAry)

	if iBa.intAryLen == iBa.precision {
		iBa.intAry = append([]uint8{0}, iBa.intAry...)
		iBa.intAryLen++
	}

	if iBa.intAryLen < iBa.precision {

		deltaZeros := iBa.precision - iBa.intAryLen + 1
		zeroAry := make([]uint8, deltaZeros)
		iBa.intAry = append(zeroAry, iBa.intAry...)
		iBa.intAryLen += deltaZeros
	}

	iBa.firstDigitIdx = -1
	iBa.lastDigitIdx = -1

	lastIntIdx := iBa.intAryLen - iBa.precision - 1
	iBa.isZeroValue = true
	iBa.isIntegerZeroValue = true
	iBa.integerLen = iBa.intAryLen - iBa.precision

	for i := 0; i < iBa.intAryLen; i++ {
		if iBa.intAry[i] > 0 {
			iBa.isZeroValue = false

			if i < iBa.integerLen {
				iBa.isIntegerZeroValue = false
			}
		}
		// At minimum, there should be a single
		// leading zero before the decimal point.
		// Example 0.000.
		if i == lastIntIdx && iBa.intAry[i] == 0 {

			if iBa.firstDigitIdx == -1 {
				iBa.firstDigitIdx = i
			}

		}

		if iBa.intAry[i] > 0 {

			if iBa.firstDigitIdx == -1 {
				iBa.firstDigitIdx = i
			}

			iBa.lastDigitIdx = i
		}

	}

	iBa.significantIntegerLen = iBa.intAryLen - iBa.precision - iBa.firstDigitIdx

	if iBa.lastDigitIdx >= iBa.integerLen {
		iBa.significantFractionLen = iBa.precision - (iBa.lastDigitIdx - iBa.integerLen + 1)
	} else {
		iBa.significantFractionLen = 0
	}

}

func (iBa *BackUpIntAry) SetSignValue(signVal int) {

	if signVal < 0 {
		panic("BackUpIntAry.SetSignValue() - sign value less than zero!")
	}

	iBa.signVal = signVal

}
