package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoPhoton struct {
	lock sync.Mutex
}

// formatForMathOps
//
//	Receives two NumStrDto objects and converts their number
//	strings such that both have the same number of integer and
//	fractional digits while maintaining their original numeric
//	values. This transform will facilitate the performance of
//	string based math operations such as addition and subtraction.
//
//	The return values represent the formatted NumStrDto objects.
//	The first NumStrDto returned always contains the larger
//	absolute value. The second NumStrDto always contains the
//	absolute numeric value which is less than or equal to the first
//	NumStrDto object returned.
//
//	The third parameter returned by this method is an integer
//	('compare') value which will always be set to '1' or '0'.
//
//	'1' indicates that the absolute value of the first NumStrDto
//	object ('n1DtoOut') returned by this method is greater than the
//	second NumStrDto object ('n2DtoOut') returned by this method.
//
//	If the returned integer ('compare') value returned is zero, it
//	signals that the absolute values (not the signed values) of
//	both returned NumStrDto objects are equal.
//
//	If the absolute value of 'n1Dto' is less than 'n2Dto', return
//	value 'n1DtoOut' will be populated with 'n2Dto' values, return
//	value 'n2DtoOut' will be populated with 'n1Dto' values and return
//	parameter 'isOrderReversed' will be set to 'true'.
func (nStrDtoPhoton *numStrDtoPhoton) formatForMathOps(
	numSeps NumericSeparatorDto,
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) (
	n1DtoOut NumStrDto,
	n2DtoOut NumStrDto,
	compare int,
	isOrderReversed bool,
	err error) {

	nStrDtoPhoton.lock.Lock()

	defer nStrDtoPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoPhoton.formatForMathOps()",
		"")

	if err != nil {
		return NumStrDto{},
			NumStrDto{},
			0,
			false,
			err
	}

	if n1Dto == nil {

		return NumStrDto{},
			NumStrDto{},
			0,
			false,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n1Dto'",
			}
	}

	if n2Dto == nil {

		return NumStrDto{},
			NumStrDto{},
			0,
			false,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n2Dto'",
			}
	}

	if validateN1Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return NumStrDto{},
				NumStrDto{},
				0,
				false,
				&FuncReturnError{
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

			return NumStrDto{},
				NumStrDto{},
				0,
				false,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n2Dto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{},
			NumStrDto{},
			0,
			false,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter 'numSeps' is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	lenN1AllRunes := 0
	lenN1IntRunes := 0
	lenN1FracRunes := 0
	lenN2AllRunes := 0
	lenN2IntRunes := 0
	lenN2FracRunes := 0

	compare = nDto.CompareAbsoluteValues(&n1Dto, &n2Dto)

	if compare == 1 {
		n1DtoOut = n1Dto.CopyOut()
		n2DtoOut = n2Dto.CopyOut()
	} else if compare == -1 {
		n1DtoOut = n2Dto.CopyOut()
		n2DtoOut = n1Dto.CopyOut()
		isOrderReversed = true
		compare = 1
	} else {
		// compare must be zero
		n1DtoOut = n1Dto.CopyOut()
		n2DtoOut = n2Dto.CopyOut()
	}

	n1DtoOutAbsIntRunes := n1DtoOut.GetAbsIntRunes()
	n1DtoOutAbsFracRunes := n1DtoOut.GetAbsFracRunes()

	n2DtoOutAbsIntRunes := n2DtoOut.GetAbsIntRunes()
	n2DtoOutAbsFracRunes := n2DtoOut.GetAbsFracRunes()

	if n1DtoOut.precision > n2DtoOut.precision {

		deltaPrecision := n1DtoOut.precision - n2DtoOut.precision

		for i := uint(0); i < deltaPrecision; i++ {
			n2DtoOut.absAllNumRunes = append(n2DtoOut.absAllNumRunes, '0')
			n2DtoOutAbsFracRunes = append(n2DtoOutAbsFracRunes, '0')
		}

		lenN2AllRunes = len(n2DtoOut.absAllNumRunes)
		lenN2IntRunes = len(n2DtoOutAbsIntRunes)
		lenN2FracRunes = len(n2DtoOutAbsFracRunes)

		n2DtoOut.precision = n1DtoOut.precision
		err = n2DtoOut.IsValid(ePrefix)

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

		lenN1AllRunes = len(n1DtoOut.absAllNumRunes)
		lenN1IntRunes = len(n1DtoOutAbsIntRunes)
		lenN1FracRunes = len(n1DtoOutAbsFracRunes)

	} else if n1DtoOut.precision < n2DtoOut.precision {

		deltaPrecision := n2DtoOut.precision - n1DtoOut.precision

		for i := uint(0); i < deltaPrecision; i++ {
			n1DtoOut.absAllNumRunes = append(n1DtoOut.absAllNumRunes, '0')
			n1DtoOutAbsFracRunes = append(n1DtoOutAbsFracRunes, '0')
		}

		lenN1AllRunes = len(n1DtoOut.absAllNumRunes)
		lenN1IntRunes = len(n1DtoOutAbsIntRunes)
		lenN1FracRunes = len(n1DtoOutAbsFracRunes)

		n1DtoOut.precision = n2DtoOut.precision
		err = n1DtoOut.IsValid(ePrefix)

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

		lenN2AllRunes = len(n2DtoOut.absAllNumRunes)
		lenN2IntRunes = len(n2DtoOutAbsIntRunes)
		lenN2FracRunes = len(n2DtoOutAbsFracRunes)

	} else {
		// n1DtoOut.precision == n2DtoOut.precision

		lenN1AllRunes = len(n1DtoOut.absAllNumRunes)
		lenN1IntRunes = len(n1DtoOutAbsIntRunes)
		lenN1FracRunes = len(n1DtoOutAbsFracRunes)

		lenN2AllRunes = len(n2DtoOut.absAllNumRunes)
		lenN2IntRunes = len(n2DtoOutAbsIntRunes)
		lenN2FracRunes = len(n2DtoOutAbsFracRunes)

	}

	if lenN2IntRunes > lenN1IntRunes {

		var absAllRunes []rune
		var absIntRunes []rune
		deltaRunes := lenN2IntRunes - lenN1IntRunes
		for i := 0; i < deltaRunes; i++ {
			absAllRunes = append(absAllRunes, '0')
			absIntRunes = append(absIntRunes, '0')
		}

		for j := 0; j < lenN1AllRunes; j++ {
			absAllRunes = append(absAllRunes, n1DtoOut.absAllNumRunes[j])

			if j < lenN1IntRunes {
				absIntRunes = append(absIntRunes, n1DtoOutAbsIntRunes[j])
			}

		}

		n1DtoOut.absAllNumRunes = absAllRunes
		n1DtoOutAbsIntRunes = absIntRunes
		lenN1AllRunes = len(n1DtoOut.absAllNumRunes)
		lenN1IntRunes = len(n1DtoOutAbsIntRunes)

		err = n1DtoOut.IsValid(ePrefix)

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

	} else if lenN1IntRunes > lenN2IntRunes {

		var absAllRunes []rune
		var absIntRunes []rune
		deltaRunes := lenN1IntRunes - lenN2IntRunes
		for i := 0; i < deltaRunes; i++ {
			absAllRunes = append(absAllRunes, '0')
			absIntRunes = append(absIntRunes, '0')
		}

		for j := 0; j < lenN2AllRunes; j++ {
			absAllRunes = append(absAllRunes, n2DtoOut.absAllNumRunes[j])

			if j < lenN2IntRunes {
				absIntRunes = append(absIntRunes, n2DtoOutAbsIntRunes[j])
			}

		}

		n2DtoOut.absAllNumRunes = absAllRunes
		n2DtoOutAbsIntRunes = absIntRunes
		lenN2AllRunes = len(n2DtoOut.absAllNumRunes)
		lenN2IntRunes = len(n2DtoOutAbsIntRunes)

		err := n2DtoOut.IsValid(ePrefix)

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

	}

	if lenN1AllRunes != lenN2AllRunes {
		return NumStrDto{}, NumStrDto{}, 0, false,
			fmt.Errorf("FormatForMathOps() - n1 and n2 AllNumRune arrays are NOT equal in length. "+
				"n1 length= '%v' n2 length= '%v'", lenN1AllRunes, lenN2AllRunes)
	}

	if lenN1IntRunes != lenN2IntRunes {
		return NumStrDto{}, NumStrDto{}, 0, false,
			fmt.Errorf("FormatForMathOps() - n1 and n2 IntRunes arrays are NOT equal in length. "+
				"n1 length= '%v' n2 length= '%v'", lenN1IntRunes, lenN2IntRunes)
	}

	if lenN1FracRunes != lenN2FracRunes {
		return NumStrDto{}, NumStrDto{}, 0, false,
			fmt.Errorf("FormatForMathOps() - n1 and n2 FracRunes arrays are NOT equal in length. "+
				"n1 length= '%v' n2 length= '%v'", lenN1FracRunes, lenN2FracRunes)
	}

	if n1DtoOut.precision != n2DtoOut.precision {
		return NumStrDto{}, NumStrDto{}, 0, false,
			fmt.Errorf("FormatForMathOps() - n1 and n2 FracRunes arrays are NOT equal in length. "+
				"n1 length= '%v' n2 length= '%v'", lenN1FracRunes, lenN2FracRunes)

	}

	err = n1DtoOut.IsValid(ePrefix + "n1DtoOut - ")

	if err != nil {
		return NumStrDto{}, NumStrDto{}, 0, false, err
	}

	err = n2DtoOut.IsValid(ePrefix + "n2DtoOut - ")

	if err != nil {
		return NumStrDto{}, NumStrDto{}, 0, false, err
	}

	return n1DtoOut, n2DtoOut, compare, isOrderReversed, nil
}
