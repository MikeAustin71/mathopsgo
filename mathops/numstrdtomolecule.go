package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoMolecule struct {
	lock sync.Mutex
}

// copy
//
//	Copies the member data elements from the source NumStrDto
//	object to the destination NumStrDto object.
//
// NumStrDto fields and returns a completely
// new instance of NumStrDto
func (nStrDtoMolecule *numStrDtoMolecule) copy(
	destinationNStrDto *NumStrDto,
	sourceNStrDto *NumStrDto,
	validateSourceDto bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.copy()",
		"")

	if err != nil {
		return err
	}

	if sourceNStrDto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'sourceNStrDto'",
		}
	}

	if destinationNStrDto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'destinationNStrDto'",
		}
	}

	if validateSourceDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			sourceNStrDto, ePrefix.XCpy("Validating 'sourceNStrDto'"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  sourceNStrDto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	lenSrcRunes := len(sourceNStrDto.absAllNumRunes)

	destinationNStrDto.absAllNumRunes =
		make([]rune, lenSrcRunes)

	for i := 0; i < lenSrcRunes; i++ {
		destinationNStrDto.absAllNumRunes[i] = sourceNStrDto.absAllNumRunes[i]
	}

	destinationNStrDto.signVal = sourceNStrDto.signVal

	destinationNStrDto.precision = sourceNStrDto.precision

	destinationNStrDto.thousandsSeparator = sourceNStrDto.thousandsSeparator

	destinationNStrDto.decimalSeparator = sourceNStrDto.decimalSeparator

	destinationNStrDto.currencySymbol = sourceNStrDto.currencySymbol

	return nil
}

// newZeroNumStrDto
//
//	Returns a new NumStrDto initialized to zero value. If the
//	parameter numFracDigits is set to a value greater than zero,
//	then an equal number of zero characters will be added to the
//	right of the decimal point.
//
//	Examples
//	========
//
//	numFracDigits    Results NumStrOut
//
//	     0                "0"
//	     2                "0.00"
//	     4                "0.0000"
func (nStrDtoMolecule *numStrDtoMolecule) newZeroNumStrDto(
	numSeps NumericSeparatorDto,
	numFracDigits uint) NumStrDto {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	numSeps.SetDefaultsIfEmpty()

	n2Dto := NumStrDto{}
	n2Dto.signVal = 1
	n2Dto.thousandsSeparator = numSeps.ThousandsSeparator
	n2Dto.decimalSeparator = numSeps.DecimalSeparator
	n2Dto.currencySymbol = numSeps.CurrencySymbol
	n2Dto.signVal = 1
	n2Dto.precision = 0
	n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')

	if numFracDigits > 0 {

		for i := uint(0); i < numFracDigits; i++ {
			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')
		}

		n2Dto.precision = numFracDigits
	}

	return n2Dto
}
