package mathops

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
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

		}
		// n1Dto.signVal != n2Dto.signVal

		if n1Dto.signVal == 1 {
			return 1, nil
		}

		// n2Dto.signVal must == 1
		return -1, nil

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
		}
		// must be n2Dto.signVal == 1

		return -1, nil

	}

	// MUST BE:
	// cmpAbs == -1

	if n2Dto.signVal == n1Dto.signVal {

		if n2Dto.signVal == 1 {
			// n1Dto.signVal && n2Dto.signVal must equal 1

			return -1, nil
		}
		// n1Dto.signVal && n2Dto.signVal must equal -1

		return 1, nil
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

// getSciNotationStr
//
//	Returns a string expressing the current NumStrDto numerical value as
//	scientific notation.
//
//	Example Scientific Notation
//	===========================
//
//	    scientific notation string: '2.652e+8'
//	    significand  = '2.652'
//	    significand integer digit  = '2'
//	    mantissa  = significand factional digits = '.652'
//	    exponent  = '8' (10^8)
//
//	Input Parameter
//	===============
//
//	nDto                     *NumStrDto
//	  The returned SciNotationNum string contains the numeric
//	  value extracted from this instance of NumStrDto
//
//	mantissaLen              uint
//	 Specifies the length of the mantissa in the returned
//	 scientific notation string. If the value of 'mantissaLen' is
//	 less than two ('2'), this method will automatically set the
//	 'mantissaLen' to a default value of two ('2').
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//	  This object encapsulates an error prefix string
//	  which is included in all returned error
//	  messages. Usually, it contains the name of the
//	  calling method or methods listed as a function
//	  chain.
//
//	  If no error prefix information is needed, set
//	  this parameter to 'nil'.
//
//	  Type ErrPrefixDto is included in the 'errpref'
//	  software package:
//	    "github.com/MikeAustin71/errpref".
//
//	Return Values
//	=============
//
//	string
//	  This returned strings contains the numeric representaion of
//	  the NumStrDto instance 'nDto' formatted in Scientific
//	  Notation.
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nStrDtoMolecule *numStrDtoMolecule) getSciNotationStr(
	nDto *NumStrDto,
	validateNumStrDto bool,
	mantissaLen uint,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.getSciNotationStr()",
		"")

	if err != nil {
		return "", err
	}

	if nDto == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	sciNotation, err := new(numStrDtoGluon).getSciNotationNumber(
		nDto, false, mantissaLen, ePrefix)

	if err != nil {

		return "",
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "sciNotation, err := new(numStrDtoGluon).getSciNotationNumber(\n" +
					"  nDto, false, mantissaLen, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	sciNotationStr, err := sciNotation.GetSciNotationStr(mantissaLen)

	if err != nil {

		return "",
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "sciNotationStr, err := sciNotation.\n" +
					"  GetSciNotationStr(mantissaLen)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return sciNotationStr, nil
}

// newBigFloat
//
//	Receives an incoming Big Float (*big.Float) type and a
//	precision specification which is converted and returned as a
//	new NumStrDto instance.
//
//	The 'precision' specification designates the number of digits
//	to the right of the decimal point in the final numeric value.
//
//	"A negative 'precision' selects the smallest number of
//	decimal digits necessary to identify the value x uniquely
//	using x.Prec() mantissa bits."
//	  Go Documentation: https://pkg.go.dev/math/big#Float.Text
func (nStrDtoMolecule *numStrDtoMolecule) newBigFloat(
	numSeps NumericSeparatorDto,
	bigFloat *big.Float,
	precision int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newBigFloat()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if bigFloat == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigFloat'",
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
				"'numSeps' Numeric Separators FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	numStr := bigFloat.Text('f', precision)

	n2, err := new(numStrDtoQuark).parseNumStr(numSeps, numStr, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: fmt.Sprintf("numStr = '%s'\n"+
					"numSeps = '%s'", numStr, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newBigInt
//
//	Receives a signed Bit Int Number (*big.Int) and precision
//	specification. This method then proceeds to create and return
//	a new instace of NumStrDto instance.
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	  signedBigInt    precision           result
//
//	    946254            3               946.254
//	    946254            1               94625.4
//	    946254            0               946254
//	   -946254            3              -946.254
//	   -946254            2              -9462.54
//	   -946254            0              -946254
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	The maximum limit for a 'precision' uint value is
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If the 'precision' value exceeds the maximum allowable limit,
//	an error will be returned.
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
//	configured with the Numeric Separators copied from the current
//	NumStrDto instance ('nDto'). If these Numeric Separators prove
//	to be invalid, an error will be returned.
func (nStrDtoMolecule *numStrDtoMolecule) newBigInt(
	numSeps NumericSeparatorDto,
	signedBigInt *big.Int,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newBigInt()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(precision) {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
					fmt.Sprintf("precision= '%v'", precision),
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: fmt.Sprintf("Error: Numeric Separators input paramter ('numSeps') is INVALID!\n"+
					"'numSeps' FAILED Validation Tests.\n"+
					"numSeps = '%v'", numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	n2, err := new(numStrDtoQuark).parseSignedBigInt(
		numSeps, signedBigInt, precision, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseSignedBigInt(\n" +
					"  numSeps, signedBigInt, precision, ePrefix)",
				ErrContext: fmt.Sprintf("signedBigInt = '%s'\n"+
					"precision = '%v'\n"+
					"numSeps = '%s'",
					signedBigInt.Text(10), precision, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating final result 'n2'"))

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Error: Final result NumStrDto ('n2') is INVALID!\n" +
					"'n2' FAILED Final Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newBigIntNum
//
//	Receives a type BigIntNum numeric value, converts it to a
//	type NumStrDto and then returns that NumStrDto instance.
func (nStrDtoMolecule *numStrDtoMolecule) newBigIntNum(
	biNum BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newBigIntNum()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		&biNum,
		ePrefix.XCpy("Validating Input Parameter 'biNum'"))

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).isBigIntNumValid(\n" +
				"  biNum, ePrefix)",
			ErrContext: "Error: Input parameter 'biNum' (BigIntNum) is INVALID!\n" +
				"'biNum' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	numSeps, err := biNum.GetNumericSeparatorsDto()

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := biNum.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	biNumStr, err := biNum.GetNumStr()

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "biNumStr, err := biNum.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// n2, err := new(NumStrDto).ParseBigIntNum(bINum)
	n2, err := new(numStrDtoQuark).parseBigIntNum(numSeps, &biNum, false, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseBigIntNum(\n" +
					"  numSeps, &biNum, false, ePrefix)",
				ErrContext: fmt.Sprintf("biNum= '%s'\n"+
					"numSeps= '%s'", biNumStr, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newFloat32
//
//	Receives a type float32 numeric value and precision
//	specification. This method then converts these parameters into
//	a new NumStrDto instance and returns that NumStrDto instance to
//	the calling functin.
//
//	The 'precision' specification designates the number of digits
//	to the right of the decimal point in the final numeric value.
//
//	'precision' MUST BE >= -1
//
//	"The special precision -1 uses the smallest number of digits
//	necessary such that ParseFloat will return f exactly. The
//	exponent is written as a decimal integer ..."
//	  Go Documentation: https://pkg.go.dev/strconv#FormatFloat
func (nStrDtoMolecule *numStrDtoMolecule) newFloat32(
	numSeps NumericSeparatorDto,
	f32 float32,
	precision int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newFloat32()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	if precision < -1 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' must be >= -1\n" +
					fmt.Sprintf("precision = %d", precision),
			}
	}

	numStr := strconv.FormatFloat(float64(f32), 'f', precision, 32)

	n2, err := new(numStrDtoQuark).parseNumStr(numSeps, numStr, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: fmt.Sprintf("numStr= '%s'\n"+
					"numSeps= '%s'", numStr, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newFloat64
//
//	Receives a type float32 numeric value and precision
//	specification. This method then converts these parameters into
//	a new NumStrDto instance and returns that NumStrDto instance to
//	the calling functin.
//
//	The 'precision' specification designates the number of digits
//	to the right of the decimal point in the final numeric value.
//
//	'precision' MUST BE >= -1
//
//	"The special precision -1 uses the smallest number of digits
//	necessary such that ParseFloat will return f exactly. The
//	exponent is written as a decimal integer ..."
//	  Go Documentation: https://pkg.go.dev/strconv#FormatFloat
func (nStrDtoMolecule *numStrDtoMolecule) newFloat64(
	numSeps NumericSeparatorDto,
	f64 float64,
	precision int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newFloat64()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	if precision < -1 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' must be >= -1\n" +
					fmt.Sprintf("precision = %d", precision),
			}
	}

	numStr := strconv.FormatFloat(f64, 'f', precision, 64)

	n2, err := new(numStrDtoQuark).parseNumStr(numSeps, numStr, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: fmt.Sprintf("numStr= '%s'\n"+
					"numSeps= '%s'", numStr, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newInt64
//
//	Creates a new NumStrDto instance from an int64 value and a
//	precision specification.
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	   int64         precision          result
//
//	   946254            3               946.254
//	   946254            1               94625.4
//	   946254            0               946254
//	  -946254            3              -946.254
//	  -946254            2              -9462.54
//	  -946254            0              -946254
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	The maximum limit for a 'precision' uint value is
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If the 'precision' value exceeds the maximum allowable limit,
//	an error will be returned.
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
//	configured with the Numeric Separators copied from the input
//	parameter ('numSeps'). If these Numeric Separators prove
//	to be invalid, an error will be returned.
//
//	Useage
//	======
//
//	        nDto, err := new(NumStrDto).NewInt64(123456, 3)
//	      'nDto' is returned with a numeric value of 123.456.
func (nStrDtoMolecule *numStrDtoMolecule) newInt64(
	numSeps NumericSeparatorDto,
	i64 int64,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newInt64()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(precision) {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
					fmt.Sprintf("precision= '%v'", precision),
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
				"'numSeps' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	numStr := strconv.FormatInt(i64, 10)

	n2, err := new(numStrDtoQuark).parseNumStr(
		numSeps, numStr, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoMuon).setPrecisionNumStrDto(
		numSeps, &n2, precision, true, ePrefix.XCpy("n2"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoMuon).setPrecisionNumStrDto(\n" +
					"  numSeps, &n2, precision, true, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newInt64Exponent
//
//	Returns a new NumStrDto instance. The numeric value is set
//	using an int64 value multiplied by 10 raised to the power of
//	the 'exponent' parameter.
//
//	      numeric value = int64 X 10^exponent
//
//	Usage
//	=====
//
//		nDto := new(NumStrDto).NewInt64Exponent(123456, -3)
//	     nDto is now equal to "123.456", precision = 3
//
//		nDto := new(NumStrDto).NewInt64Exponent(123456, 3)
//	  nDto is now equal to "123456.000", precision = 3
//
//	Examples
//	========
//
//	int64Num      exponent      Result
//
//	 123456         -3          123.456
//	 123456          3          123456.000
//	 123456          0          123456
func (nStrDtoMolecule *numStrDtoMolecule) newInt64Exponent(
	numSeps NumericSeparatorDto,
	int64Num int64,
	exponent int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newInt64Exponent",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	numStr := strconv.FormatInt(int64Num, 10)

	if exponent > 0 {

		for i := 0; i < exponent; i++ {

			numStr += "0"

		}
	}

	if exponent < 0 {

		exponent = exponent * -1
	}

	var n2 NumStrDto

	if exponent == 0 {

		n2, err = new(numStrDtoQuark).parseNumStr(
			numSeps, numStr, ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "n2, err = new(numStrDtoQuark).parseNumStr(\n" +
						"  numSeps, numStr, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	} else {

		n2, err = new(numStrDtoTau).shiftPrecisionLeft(
			numSeps, numStr, uint(exponent), ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "n2, err = new(numStrDtoTau).shiftPrecisionLeft(\n" +
						"  numSeps, numStr, uint(exponent), ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newNumStrWithNumSeps
//
//	Receives a number string as input and returns a new NumStrDto
//	instance. This number string is assumed to contain numeric
//	digits which may include a decimal separator character to
//	separate integer and fractional digits. In addtion, negative
//	numeric values will include a minus sign ('-') in the first
//	string position or the negative value may be surrouned with
//	parentheses ('()').
//
//	The input parameter 'numSeps' contains numeric	separators
//	(decimal separator, thousands separator and currency symbol)
//	which will be used to parse the number string.
//
//	The numeric separators contained in inputparameter 'numSeps'
//	will be used to parse the number string and configured the
//	returned NumStrDto instance.
func (nStrDtoMolecule *numStrDtoMolecule) newNumStrWithNumSeps(
	signedNumStr string,
	numStrNumSeps IGetNumSeparators,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newNumStrWithNumSeps",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if len(signedNumStr) == 0 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(numStr) == 0",
				ErrMessage: "Error: Input parameter 'numStr' is INVALID!\n" +
					"'numStr' is an empty, zero length string.",
			}

	}

	inputNumSeps, err := numStrNumSeps.GetInputSeparators()

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "inputNumSeps, err = numStrNumSeps.GetInputSeparators()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	err = inputNumSeps.IsValid(ePrefix.XCpy("Validating 'inputNumSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = inputNumSeps.IsValid(ePrefix.XCpy(\"Validating 'inputNumSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('inputNumSeps') is INVALID!\n" +
					"'inputNumSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	outputNumSeps, err := numStrNumSeps.GetOutputSeparators()

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "outputNumSeps, err = numStrNumSeps.GetOutputSeparators()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	err = outputNumSeps.IsValid(ePrefix.XCpy("Validating 'outputNumSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = outputNumSeps.IsValid(ePrefix.XCpy(\"Validating 'outputNumSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('outputNumSeps') is INVALID!\n" +
					"'outputNumSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	//n2, err := n.ParseNumStr(numStr)
	n2, err := new(numStrDtoQuark).parseNumStr(
		*inputNumSeps, signedNumStr, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"*inputNumSeps, signedNumStr, ePrefix)",
				ErrContext: fmt.Sprintf("inputNumSeps= '%s'\n"+
					"signedNumStr= '%s'", inputNumSeps.String(), signedNumStr),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoAtom).setNumericSeparatorsDto(
		&n2, *outputNumSeps, ePrefix)

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoAtom).setNumericSeparatorsDto(\n" +
					"&n2, *outputNumSeps, ePrefix)",
				ErrContext: fmt.Sprintf("outputNumSeps= '%s'\n",
					outputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Error: Final Calculated NumStrDto Result 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newRational
//
//	Creates a new NumStrDto instance from a big rational number
//	(*big.Rat) and a precision specification.
//
//	For information on Big Rational Numbers (*big.Rat), see:
//	  https://golang.org/pkg/math/big/
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	  bigRat         precision     NumStrDto Result
//
//	   946254/1          3               946.254
//	   946254/1          1               94625.4
//	   946254/1          0               946254
//	  -946254/1          3              -946.254
//	  -946254/1          2              -9462.54
//	  -946254/1          0              -946254
//	       40/2          2               20.00
//	      480/5          1               81.6
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	The maximum limit for a 'precision' uint value is
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If this value exceeds the maximum value for a 32-bit integer,
//	this 'precision' value will be automatically reduced to the
//	maximum limit of 2,147,483,647 or	2^31 - 1.
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
//	configured with the Numeric Separators copied from the current
//	NumStrDto instance ('nDto'). If these Numeric Separators prove
//	to be invalid, they will be automatically reset to USA default
//	values.
func (nStrDtoMolecule *numStrDtoMolecule) newRational(
	numSeps NumericSeparatorDto,
	bigRat *big.Rat,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newRational()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if bigRat == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigRat'",
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(precision) {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
					fmt.Sprintf("precision= '%v'", precision),
			}
	}

	numStr := bigRat.FloatString(int(precision))

	n2, err := new(numStrDtoQuark).parseNumStr(
		numSeps,
		numStr,
		ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: fmt.Sprintf("numSeps= '%s'\n"+
					"numStr= '%s'",
					numSeps.String(), numStr),
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newUint64
//
//	Receives an uint64 value and a precision specification. This
//	method then proceeds to create and return a new instance of
//	NumStrDto.
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	  intNum         precision     NumStrDto Result
//
//	   946254            3               946.254
//	   946254            1               94625.4
//	   946254            0               946254
//	  -946254            3              -946.254
//	  -946254            2              -9462.54
//	  -946254            0              -946254
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	The maximum limit for a 'precision' uint value is
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If this value exceeds the maximum value for a 32-bit integer,
//	this 'precision' value will be automatically reduced to the
//	maximum limit of 2,147,483,647 or	2^31 - 1.
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
//	configured with the Numeric Separators copied from the current
//	NumStrDto instance ('nDto'). If these Numeric Separators prove
//	to be invalid, they will be automatically reset to USA default
//	values.
//
//	Usage
//	=====
//
//	          uint64Num := uint64(123456)
//	          precision := uint(3)
//	          nDto, err := new(NumStrDto).NewUint64(uint64Num, precision)
//	                'nDto' is now equal to 123.456
func (nStrDtoMolecule *numStrDtoMolecule) newUint64(
	numSeps NumericSeparatorDto,
	uint64Num uint64,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newUint64()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(precision) {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
					fmt.Sprintf("precision= '%v'", precision),
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
				"'numSeps' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	numStr := strconv.FormatUint(uint64Num, 10)

	n2, err := new(numStrDtoQuark).parseNumStr(
		numSeps,
		numStr,
		ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: fmt.Sprintf("numSeps= '%s'\n"+
					"numStr = '%s'\n"+
					"uint64Num= '%v'\n",
					numSeps.String(), numStr, uint64Num),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoMuon).setPrecisionNumStrDto(
		numSeps, &n2, precision, true, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
					"  numSeps, numStr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
}

// newUint64Exponent
//
//	 Returns a new NumStrDto instance. The numeric value for this
//	 new NumStrDto is set using an uint64 value multiplied by 10
//	 raised to the power of the 'exponent' parameter.
//
//	          numeric value = uint64Num X 10^exponent
//
//		Usage
//		=====
//
//		  nDto := new(NumStrDto).NewIntExponent(uint64(123456), -3)
//		     nDto is now equal to "123.456", precision = 3
//
//		  nDto := new(NumStrDto).NewIntExponent(uint64(123456), 3)
//		     nDto is now equal to "123456.000", precision = 3
//
//		Examples
//		========
//
//		intNum        exponent        NumStrDto Result
//
//		123456          -3                123.456
//		123456           3                123456.000
//		123456           0                123456
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and
//		convert them into numeric values.
//
//		The final NumStrDto result returned by this method will be
//		configured with the Numeric Separators copied from the current
//		NumStrDto instance ('nDto'). If these Numeric Separators prove
//		to be invalid, an error will be returned.
func (nStrDtoMolecule *numStrDtoMolecule) newUint64Exponent(
	numSeps NumericSeparatorDto,
	uint64Num uint64,
	exponent int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.newUint64Exponent()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
				"'numSeps' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	numStr := strconv.FormatUint(uint64Num, 10)

	if exponent > 0 {

		for i := 0; i < exponent; i++ {

			numStr += "0"

		}
	}

	if exponent < 0 {

		exponent = exponent * -1

	}

	var n2 NumStrDto

	if exponent == 0 {

		n2, err = new(numStrDtoQuark).parseNumStr(
			numSeps, numStr, ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "n2, err := new(numStrDtoQuark).parseNumStr(\n" +
						"  numSeps, numStr, ePrefix)",
					ErrContext: fmt.Sprintf("numSeps= '%s'\n"+
						"numStr = '%s'\n"+
						"uint64Num= '%v'\n",
						numSeps.String(), numStr, uint64Num),
					ErrMessage: err.Error(),
				}
		}

	} else {

		//n2, err = nDto.ShiftPrecisionLeft(numStr, uint(exponent))

		n2, err = new(numStrDtoTau).shiftPrecisionLeft(
			numSeps, numStr, uint(exponent), ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "n2, err = new(numStrDtoTau).shiftPrecisionLeft(\n" +
						"  numSeps, numStr, uint(exponent), ePrefix)",
					ErrContext: fmt.Sprintf("numSeps= '%s'\n"+
						"numStr = '%s'\n"+
						"uint64Num= '%v'\n"+
						"exponent= '%d'",
						numSeps.String(), numStr, uint64Num, exponent),
					ErrMessage: err.Error(),
				}
		}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2, ePrefix.XCpy("Validating Final Result 'n2'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2' is INVALID!\n" +
					"'n2' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2, nil
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

// scaleNumStr
//
//	Shifts the position of the decimal point left or right
//	depending on the value of input parameter 'scaleMode'.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. These digits
//	must be formatted in a way that facilitates conversion to a
//	corresponding numeric value.
//
//	Number String Negative Values
//	=============================
//
//	The 'numStr' number string parameter passed to this method must
//	consist of a string of numeric digits representing a numeric
//	value. A leading minus sign (-), or surrounding parentheses
//	'()', may be included in this number string to indicate a
//	negative numeric value.
//
//	Fractional Digits in Number Strings
//	===================================
//
//	The 'numStr' number string of numeric digits may also include
//	a delimiting decimal separator to identify fractional digits to
//	the right of the decimal separator. In the USA, the default
//	decimal separator is the period character ('.'). The actual
//	decimal separator character used to parse the 'numStr' number
//	string is determined by the Numeric Separators parameter,
//	'numStrNumSeps'.
//
//	Input Parameters
//	================
//
//	signedNumStr					 string
//	  This parameter should be formatted as a string of numeric
//	  digits as outlined above. Using the Numeric
//	  Separators provided by input parameter 'numStrNumSeps',
//	  this method will parse the 'numStr' number string and
//	  convert it to a numeric value which will be returned as a
//	  NumStrDto.
//
//	numStrNumSeps            IGetNumSeparators
//	  The IGetNumSeparators interface type gives users the option
//	  of submitting one of two different concrete types.
//
//	  User may choose to submit a type NumericSeparatorDto
//	  consisting of one set of Numeric Separators. These Numeric
//	  Separators will be used to both parse the number string
//	  provided by input parameter 'signedNumStr' and format the
//	  returned NumStrDto type containing the converted numeric
//	  value.
//
//	  The second alternative allows the user to submit a type
//	  NumericSeparatorPairDto for this parameter. This type
//	  encapsulates two separate instances of NumericSeparatorDto.
//	  The 'input' NumericSeparatorDto instance will be used to
//	  parse number string 'signedNumStr' while the 'output'
//	  instance will be used to format the numeric value contained
//	  in the returned instance of NumStrDto.
//
//	shiftPrecision           uint
//	  The number of positions which the decimal point will be
//	  shifted. If 'shiftPrecision' is Equal to zero, no action will
//	  be taken, no error will be issued and the original
//	  'signedNumStr' will be configured in the returned NumStrDto
//	  instance.
//
//	scaleMode                PrecisionScaleMode
//	  A constant with one of two Scale Mode values.
//
//	  SCALEPRECISIONLEFT -  Shifts the decimal point from its
//	                        current position to the left.
//
//	  SCALEPRECISIONRIGHT - Shifts the decimal point from its
//	                        current position to the right.
//
//	  Note: See Methods NumStrDto.ShiftPrecisionRight() and
//	  NumStrDto.ShiftPrecisionLeft() for additional information.
//
//	Return Values
//	=============
//
//	NumStrDto
//	  This new NumStrDto instance contains the numeric value
//	  extracted from the 'signedNumStr' and transformed as
//	 described above.
//
//	error
//	  If errors are encountered during processng, this returned
//	  error object will be configured with an appropriate error
//	  message.
func (nStrDtoMolecule *numStrDtoMolecule) scaleNumStr(
	signedNumStr string,
	numStrNumSeps IGetNumSeparators,
	shiftPrecision uint,
	scaleMode PrecisionScaleMode,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMolecule.lock.Lock()

	defer nStrDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMolecule.scaleNumStr",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if len(signedNumStr) == 0 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(numStr) == 0",
				ErrMessage: "Error: Input parameter 'signedNumStr' is INVALID!\n" +
					"'signedNumStr' is an empty, zero length string.",
			}

	}

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(shiftPrecision) {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'shiftPrecision' is INVALID!\n" +
					"'precision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
					fmt.Sprintf("shiftPrecision= '%v'", shiftPrecision),
			}
	}

	inputNumSeps, err := numStrNumSeps.GetInputSeparators()

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "inputNumSeps, err = numStrNumSeps.GetInputSeparators()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	err = inputNumSeps.IsValid(ePrefix.XCpy("Validating 'inputNumSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = inputNumSeps.IsValid(ePrefix.XCpy(\"Validating 'inputNumSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('inputNumSeps') is INVALID!\n" +
					"'inputNumSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	outputNumSeps, err := numStrNumSeps.GetOutputSeparators()

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "outputNumSeps, err = numStrNumSeps.GetOutputSeparators()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = outputNumSeps.IsValid(ePrefix.XCpy("Validating 'outputNumSeps'").String())

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = outputNumSeps.IsValid(ePrefix.XCpy(\"Validating 'outputNumSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('outputNumSeps') is INVALID!\n" +
					"'outputNumSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	n2Dto := NumStrDto{}

	if scaleMode == SCALEPRECISIONLEFT {

		// n2Dto, err = nDto.ShiftPrecisionLeft(signedNumStr, shiftPrecision)
		n2Dto, err = new(numStrDtoTau).shiftPrecisionLeft(
			*inputNumSeps, signedNumStr, shiftPrecision, ePrefix)

		if err != nil {
			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "n2Dto, err = new(numStrDtoTau).shiftPrecisionLeft(\n" +
						"*inputNumSeps, signedNumStr, shiftPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("inputNumSeps= '%s'\n"+
						"signedNumStr= '%s'\n"+
						"shiftPrecision= '%v'\n"+
						"Scale Mode= SCALEPRECISIONLEFT",
						inputNumSeps.String(), signedNumStr, shiftPrecision),
					ErrMessage: err.Error(),
				}
		}

	} else if scaleMode == SCALEPRECISIONRIGHT {

		//n2Dto, err = nDto.ShiftPrecisionRight(signedNumStr, shiftPrecision)
		n2Dto, err = new(numStrDtoTau).shiftPrecisionRight(
			*inputNumSeps, signedNumStr, shiftPrecision, ePrefix)

		if err != nil {
			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "n2Dto, err = new(numStrDtoTau).shiftPrecisionRight(\n" +
						"*inputNumSeps, signedNumStr, shiftPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("inputNumSeps= '%s'\n"+
						"signedNumStr= '%s'\n"+
						"shiftPrecision= '%v'\n"+
						"Scale Mode= SCALEPRECISIONRIGHT",
						inputNumSeps.String(), signedNumStr, shiftPrecision),
					ErrMessage: err.Error(),
				}
		}

	} else {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error! Scale Mode is INVALID!\n" +
					"Scale Mode is NOT Equal to SCALEPRECISIONLEFT or SCALEPRECISIONRIGHT.",
			}
	}

	err = new(numStrDtoAtom).setNumericSeparatorsDto(
		&n2Dto, *outputNumSeps, ePrefix.XCpy("outputNumSeps->n2Dto"))

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoAtom).setNumericSeparatorsDto(\n" +
					"&n2Dto, *outputNumSeps, ePrefix)",
				ErrContext: fmt.Sprintf("outputNumSeps= '%s'\n",
					outputNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2Dto, ePrefix.XCpy("Validating Final Result 'n2Dto'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2Dto, ePrefix)",
				ErrContext: "Final Calculated Result NumStrDto 'n2Dto' is INVALID!\n" +
					"'n2Dto' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2Dto, nil
}
