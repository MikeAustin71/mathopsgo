package mathops

import (
	"fmt"
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

// selectValidNumSepInSeries
//
// This method is designed to select the first valid
// NumericSeparatorDto instance from a variadic function parameter
// series ('primarySourceNumSeps').
//
// If no valid NumericSeparatorDto is found in the
// 'primarySourceNumSeps' series, the alternative source NumStrDto
// will be selected.
//
// In the event that the alternative source NumStrDto is selected
// and proves to be invalid, an error will be returned.
func (nSepsDtoMech *numSepsDtoMechanics) selectValidNumSepInSeries(
	primarySourceName string,
	alternateSourceName string,
	alternateSourceNumSep *NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto,
	primarySourceNumSeps ...NumericSeparatorDto) (
	foundPrimaryNumSep bool,
	validSelectedNumSep NumericSeparatorDto, err error) {

	if nSepsDtoMech.lock == nil {
		nSepsDtoMech.lock = new(sync.Mutex)
	}

	nSepsDtoMech.lock.Lock()

	defer nSepsDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyElectron.multiplyBigInt",
		"")

	if err != nil {
		return false, NumericSeparatorDto{}, err
	}

	if len(primarySourceName) == 0 {

		return false, NumericSeparatorDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input Parameter 'primarySourceName' is an empty string!",
			}
	}

	if len(alternateSourceName) == 0 {

		return false, NumericSeparatorDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input Parameter 'alternateSourceName' is an empty string!",
			}
	}

	if alternateSourceNumSep == nil {

		return false, NumericSeparatorDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: fmt.Sprintf("alternateSourceNumSep (a.k.a %v)", alternateSourceName),
			}

	}

	foundPrimaryNumSep = false

	if len(primarySourceNumSeps) > 0 {

		for _, primarySourceNumSep := range primarySourceNumSeps {

			err = primarySourceNumSep.IsValid(ePrefix.String())

			if err != nil {
				continue
			}

			err = validSelectedNumSep.CopyIn(&primarySourceNumSep, true)

			if err != nil {

				validSelectedNumSep.SetUSADefaults()

				continue

			}

			foundPrimaryNumSep = true

			break
		}
	}

	if !foundPrimaryNumSep {

		err = validSelectedNumSep.CopyIn(alternateSourceNumSep, false)

		if err != nil {

			return false, NumericSeparatorDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = %vNumSep.GetNumericSeparatorsDto()",
						alternateSourceName),
					ErrContext: fmt.Sprintf("Error: Failed to acquire %vNumSep or %vNumSep Numeric Separators",
						primarySourceName, alternateSourceName),
					ErrMessage: err.Error(),
				}
		}
	}

	return foundPrimaryNumSep, validSelectedNumSep, nil
}
