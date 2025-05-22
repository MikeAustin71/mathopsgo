package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numSepsDtoElectron struct {
	lock *sync.Mutex
}

// emptyNumSepsDto
//
// Sets all member variables in 'numSepsDto' to their 'inital'
// states.
func (nSepsDtoElectron *numSepsDtoElectron) emptyNumSepsDto(
	numSepsDto *NumericSeparatorDto) {

	if nSepsDtoElectron.lock == nil {
		nSepsDtoElectron.lock = new(sync.Mutex)
	}

	nSepsDtoElectron.lock.Lock()

	defer nSepsDtoElectron.lock.Unlock()

	if numSepsDto == nil {
		return
	}

	numSepsDto.DecimalSeparator = 0

	numSepsDto.ThousandsSeparator = 0

	numSepsDto.CurrencySymbol = 0

	return
}

// isValidNumStrDto
//
// Returns an error if 'numSepsDto' is found to be invalid.
func (nSepsDtoElectron *numSepsDtoElectron) isValidNumStrDto(
	numSepsDto *NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nSepsDtoElectron.lock == nil {
		nSepsDtoElectron.lock = new(sync.Mutex)
	}

	nSepsDtoElectron.lock.Lock()

	defer nSepsDtoElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numSepsDtoElectron.isValidNumStrDto",
		"")

	if err != nil {
		return err
	}

	if numSepsDto == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'numSepsDto'",
		}
	}

	var errStr string

	if numSepsDto.DecimalSeparator == 0 {

		errStr = "ERROR: NumericSeparatorDto DecimalSeparator value is ZERO!\n"
	}

	if numSepsDto.ThousandsSeparator == 0 {

		errStr += "ERROR: NumericSeparatorDto ThousandsSeparator value is ZERO!\n"

	}

	if numSepsDto.CurrencySymbol == 0 {
		errStr += "ERROR: NumericSeparatorDto CurrencySymbol value is ZERO!\n"
	}

	if len(errStr) > 0 {

		return fmt.Errorf("%v\n"+
			"NumericSeparatorDto is Invalid!\n"+
			"%v\n",
			ePrefix.String(),
			errStr)

	}

	return nil
}

// setNumSepDtoDefaultsIfEmpty
//
// If any of the NumericSeparatorDto member elements are invalid,
// this method will reset those member elements to USA Defaults.
func (nSepsDtoElectron *numSepsDtoElectron) setNumSepDtoDefaultsIfEmpty(
	numSepsDto *NumericSeparatorDto) {

	if numSepsDto == nil {
		return
	}

	if numSepsDto.DecimalSeparator == 0 {
		numSepsDto.DecimalSeparator = '.'
	}

	if numSepsDto.ThousandsSeparator == 0 {
		numSepsDto.ThousandsSeparator = ','
	}

	if numSepsDto.CurrencySymbol == 0 {
		numSepsDto.CurrencySymbol = '$'
	}

	return
}

// setNumSepDtoToUSADefaults
//
// If any of the NumericSeparatorDto member elements are invalid,
// this method will reset those member elements to USA Defaults.
//
// This method will arbitrarily set all member variables of
// the current NumericSeparatorDto instance to USA defaults.
// USA default values are listed as follows:
//
//		Decimal Separator   = '.' (period)
//	 Thousands Separator = ',' (comma)
//	 Currency Symbol     = '$' (dollar sign)
func (nSepsDtoElectron *numSepsDtoElectron) setNumSepDtoToUSADefaults(
	numSepsDto *NumericSeparatorDto) {

	if numSepsDto == nil {
		return
	}

	numSepsDto.DecimalSeparator = '.'

	numSepsDto.ThousandsSeparator = ','

	numSepsDto.CurrencySymbol = '$'

	return
}
