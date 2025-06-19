package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"sync"
)

type numStrDtoMolecule struct {
	lock sync.Mutex
}

// compareSignedValues
//
//	Compares the signed numeric values of two NumStrDto objects.
//
//	The term 'signed numeric values' as used here means that the
//	two NumStrDto objects being compared may be either positive
//	or negative numeric values.
//
//	Examples
//	========
//
//	   n1         n2          Result
//
//	-9691.23     91.245         -1
//	 9691.23     91.245          1
//	   -5        82             -1
//	    5         5              0
//
//	Input Parameters
//	================
//
//	n1Dto                    *NumStrDto
//	  A pointer to an instance of NumStrDto. The signed numeric
//	  value of this object will be compared to input parameter.
//	  'n2Dto'.
//
//	n2Dto                    *NumStrDto
//	  A pointer to an instance of NumStrDto. The signed numeric
//	  value of this object will be compared to input parameter.
//	  'n1Dto'.
//
//	Return Values
//	=============
//
//	int
//	  This returned integer will be set to one of three values:
//
//	  -1 = n1Dto is less than n2Dto
//
//	   0 = n1Dto is equal to n2Dto
//
//	   1 = n1Dto is greater than n2Dto
func (nStrDtoMolecule *numStrDtoMolecule) compareSignedValues(
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

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
		return 0, err
	}

	if n1Dto == nil {

		return 0, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n1Dto'",
		}
	}

	if n2Dto == nil {

		return 0, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n2Dto'",
		}
	}

	if validateN1Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return 0, &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  n1Dto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateN2Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

		if err != nil {

			return 0, &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  n2Dto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	cmpAbs, err := new(numStrDtoAtom).compareAbsoluteValues(
		n1Dto, false, n2Dto, false, ePrefix.XCpy("n1Dto vs n2Dto"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "cmpAbs, err := new(numStrDtoAtom).compareAbsoluteValues(\n" +
					"  n1Dto, false, n2Dto, false, ePrefix)\n",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if cmpAbs == 0 {

		if n1Dto.signVal == n2Dto.signVal {

			return 0, nil

		} else {
			// n1Dto.signVal != n2Dto.signVal

			if n1Dto.signVal == 1 {
				return 1, nil
			}

			// n2Dto.signVal must == 1
			return -1, nil

		}
	}

	if cmpAbs == 1 {

		if n1Dto.signVal == n2Dto.signVal {

			if n1Dto.signVal == 1 {
				return 1, nil
			}

			// must be n1Dto.signVal == n2Dto.signVal && n1Dto.signVal == -1

			return -1, nil

		}

		// must be n1Dto.signVal != n2Dto.signVal
		if n1Dto.signVal == 1 {

			return 1, nil
		} else {
			// must be n2Dto.signVal == 1

			return -1, nil
		}
	}

	// MUST BE:
	// cmpAbs == -1

	if n2Dto.signVal == n1Dto.signVal {

		if n2Dto.signVal == 1 {
			// n1Dto.signVal && n2Dto.signVal must equal 1

			return -1, nil
		} else {
			// n1Dto.signVal && n2Dto.signVal must equal -1

			return 1, nil
		}

	}

	// must be n2Dto.signVal != n1Dto.signVal

	if n2Dto.signVal == -1 {
		return 1, nil
	}

	// must be n2Dto.signVal == 1
	return -1, nil
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
//	parameter 'precision' is set to a value greater than zero,
//	then an equal number of zero characters will be added to the
//	right of the decimal point.
//
//	Examples
//	========
//
//	precision      Results NumStrOut
//
//	     0                "0"
//	     2                "0.00"
//	     4                "0.0000"
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	If this value exceeds the maximum value for a 32-bit integer,
//	this value will be automatically reduced to the maximum
//	limit of 2,147,483,647 or	2^31 - 1.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators passed by input
//	parameter 'numSeps'. If these Numeric Separators are
//	determined to be invalid, they will be automatically reset
//	to default USA values.
func (nStrDtoMolecule *numStrDtoMolecule) newZeroNumStrDto(
	numSeps NumericSeparatorDto,
	precision uint) NumStrDto {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(precision) {

		precision = uint(math.MaxInt32)

	}

	numSeps.SetDefaultsIfEmpty()

	n2Dto := NumStrDto{}
	n2Dto.signVal = 1
	n2Dto.thousandsSeparator = numSeps.ThousandsSeparator
	n2Dto.decimalSeparator = numSeps.DecimalSeparator
	n2Dto.currencySymbol = numSeps.CurrencySymbol
	n2Dto.signVal = 1
	n2Dto.precision = 0
	n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')

	if precision > 0 {

		for i := uint(0); i < precision; i++ {
			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')
		}

		n2Dto.precision = precision
	}

	return n2Dto
}
