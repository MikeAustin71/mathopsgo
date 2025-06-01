package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryBoson struct {
	lock *sync.Mutex
}

// setNumericSeparatorsToDefaultIfEmpty
//
//	 Receives a pointer to an IntAry object. By defintion, the
//	 IntAry contains fields for Numeric Separators.
//
//		If numeric separators are set to zero or nil, this method will
//		set those numeric separators to the USA defaults. This means
//		that the Decimal separator is set to a period ('.'), the
//		Thousands separator is set to a comma (',') and the currency
//		symbol is set to the dollar sign ('$').
//
//		If the numeric separators were previously set to a value other
//		than zero or nil, that value is not altered by this method.
//
//		Effectively, this method ensures that numeric separators are
//		set to valid values.
func (iaBoson *intAryBoson) setNumericSeparatorsToDefaultIfEmpty(
	intAry *IntAry,
	callingFunction string) error {

	if iaBoson.lock == nil {
		iaBoson.lock = new(sync.Mutex)
	}

	iaBoson.lock.Lock()

	defer iaBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		callingFunction,
		"IntAry.DivideByInt64()",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	if intAry.decimalSeparator == 0 {
		intAry.decimalSeparator = '.'
	}

	if intAry.thousandsSeparator == 0 {
		intAry.thousandsSeparator = ','
	}

	if intAry.currencySymbol == 0 {
		intAry.currencySymbol = '$'
	}

	return nil
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

	if iaBoson.lock == nil {
		iaBoson.lock = new(sync.Mutex)
	}

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

	if iaBoson.lock == nil {
		iaBoson.lock = new(sync.Mutex)
	}

	iaBoson.lock.Lock()

	defer iaBoson.lock.Unlock()

	if intAry == nil {
		return
	}

	intAry.BackUp = new(BackUpIntAry).New()
}
