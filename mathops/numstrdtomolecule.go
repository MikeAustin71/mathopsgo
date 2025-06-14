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
