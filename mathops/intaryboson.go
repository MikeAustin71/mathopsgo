package mathops

import (
	"sync"
)

type intAryBoson struct {
	lock sync.Mutex
}

// dataField Equality Test
//
//	Returns 'true' of all data fields in two IntAry objects
//	are equivalent.
//
//	No validation testing is performed on the two IntAry
//	objects
func (iaBoson *intAryBoson) dataFieldEqualityTest(
	iAry1 *IntAry,
	iAry2 *IntAry) bool {

	iaBoson.lock.Lock()

	defer iaBoson.lock.Unlock()

	if iAry1 == nil || iAry2 == nil {
		return false
	}

	if iAry1.intAryLen != iAry2.intAryLen {

		return false
	}

	for i := 0; i < iAry2.intAryLen; i++ {

		if iAry1.intAry[i] != iAry2.intAry[i] {

			return false
		}
	}

	if iAry1.integerLen != iAry2.integerLen ||
		iAry1.significantIntegerLen != iAry2.significantIntegerLen ||
		iAry1.significantFractionLen != iAry2.significantFractionLen ||
		iAry1.firstDigitIdx != iAry2.firstDigitIdx ||
		iAry1.lastDigitIdx != iAry2.lastDigitIdx ||
		iAry1.isZeroValue != iAry2.isZeroValue ||
		iAry1.isIntegerZeroValue != iAry2.isIntegerZeroValue ||
		iAry1.precision != iAry2.precision ||
		iAry1.signVal != iAry2.signVal ||
		iAry1.decimalSeparator != iAry2.decimalSeparator ||
		iAry1.thousandsSeparator != iAry2.thousandsSeparator ||
		iAry1.currencySymbol != iAry2.currencySymbol {

		return false
	}

	return true

}

// emptyBackUp
//
//	Deletes the values currently stored as backup for the
//	intAry object passed as input parameter 'intAry'
func (iaBoson *intAryBoson) emptyBackUp(
	intAry *IntAry) {

	iaBoson.lock.Lock()

	defer iaBoson.lock.Unlock()

	if intAry == nil {
		return
	}

	intAry.BackUp = new(BackUpIntAry).New()
}
