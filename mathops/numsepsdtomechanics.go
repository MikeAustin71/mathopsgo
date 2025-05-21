package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numSepsDtoMechanics struct {
	lock *sync.Mutex
}

// copyNumSepsDto
//
//	Input Parameters
//	================
//
//	numSepsDtoDest           *NumericSeparatorDto
//	  The destination instance of NumericSeparatorDto which
//	  will be configured and returned as identical deep
//	  copy of the source NumericSeparatorDto instance,
//	  'numSepsDtoSrc'.
//
//	numSepsDtoSrc           *NumericSeparatorDto
//	  The source instance of NumericSeparatorDto. All
//	  member variable values will be copied to the
//	  destination NumericSeparatorDto,
//	  'numSepsDtoDest'.
//
//	setDefaultsIfEmpty       bool
//	  If this boolean parameter is set to 'true' any existging
//	  invalid values in the returned NumericSeparatorDto copy
//	  ('numSepsDtoDest') will be automatically reset to valid
//	  USA defaults.
//
//	  If this setDefaultsIfEmpty is set to 'false', any invalid
//	  values copied to 'numSepsDtoDest' will trigger an error
//	  return.
func (nSepsDtoMech *numSepsDtoMechanics) copyNumSepsDto(
	numSepsDtoDest *NumericSeparatorDto,
	numSepsDtoSrc *NumericSeparatorDto,
	setDefaultsIfEmpty bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nSepsDtoMech.lock == nil {
		nSepsDtoMech.lock = new(sync.Mutex)
	}

	nSepsDtoMech.lock.Lock()

	defer nSepsDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numSepsDtoMechanics.copyNumSepsDto()",
		"")

	if err != nil {
		return err
	}

	if numSepsDtoDest == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'numSepsDtoDest'",
		}
	}

	if numSepsDtoSrc == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'numSepsDtoSrc'",
		}
	}

	nSepsDtoElectron := new(numSepsDtoElectron)

	if !setDefaultsIfEmpty {
		err = nSepsDtoElectron.isValidNumStrDto(
			numSepsDtoSrc,
			ePrefix.XCpy("Testing 'numSepsDtoSrc'"))

		if err != nil {
			return err
		}

	}

	nSepsDtoElectron.emptyNumSepsDto(numSepsDtoDest)

	numSepsDtoDest.DecimalSeparator = numSepsDtoSrc.DecimalSeparator

	numSepsDtoDest.ThousandsSeparator = numSepsDtoSrc.ThousandsSeparator

	numSepsDtoDest.CurrencySymbol = numSepsDtoSrc.CurrencySymbol

	if setDefaultsIfEmpty {

		nSepsDtoElectron.setNumSepDtoDefaultsIfEmpty(numSepsDtoDest)

	}

	return nil

}
