package mathops

import (
	"fmt"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type intAryMolecule struct {
	lock sync.Mutex
}

// PrefixToIntAry
//
//	Adds an integer of type uint8 to the front of the existing
//	internal integer array maintained by the current IntAry object.
func (iaMolecule *intAryMolecule) prefixToIntAry(
	ia *IntAry,
	num uint8) {

	iaMolecule.lock.Lock()

	defer iaMolecule.lock.Unlock()

	if ia == nil {
		return
	}

	ia.intAry = append([]uint8{num}, ia.intAry...)

	ia.intAryLen = len(ia.intAry)

	ia.integerLen = ia.intAryLen - ia.precision

	new(intAryNanobot).setInternalFlagsNoErrors(ia)

	return
}

// roundToPrecision
//
//	Rounds the value of the current IntAry instance to a precision
//	specified by the 'roundToPrecision' parameter.
func (iaMolecule *intAryMolecule) roundToPrecision(
	intAry *IntAry,
	validateIntAry bool,
	roundToPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMolecule.lock.Lock()

	defer iaMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMolecule.roundToPrecision()",
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

	if roundToPrecision < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input parameter 'roundToPrecision' is less than ZERO!\n"+
				"roundToPrecision= %v\n", roundToPrecision),
		}
	}

	if intAry.precision == 0 {
		return nil
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix.XCpy("Validate on Entry"))

	if err != nil {
		return err
	}

	numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, true, ePrefix.XCpy("Numeric Separators intAry -> numSeps"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err := new(intAryPhoton).\n" +
				"  getNumericSeparatorsDto(intAry, true,\n" +
				"  ePrefix.XCpy(Numeric Separators intAry -> numSeps))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	nsProfile := NumSepsProfileSelection{
		SourceObjectName:         "intAry",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	if intAry.isZeroValue {

		// Sets internal flags
		err = new(intAryQuark).setIntAryToZero(
			intAry,
			nil,
			nsProfile,
			uint(roundToPrecision),
			ePrefix.XCpy("Setting 'ia' to Zero"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  ia, nil, nsProfile-numSeps, 0,\n" +
					"  ePrefix.XCpy(Setting 'ia' to Zero))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	if roundToPrecision == intAry.precision {
		return nil
	}

	if roundToPrecision > intAry.precision {

		deltaPrecision := roundToPrecision - intAry.precision

		for i := 0; i < deltaPrecision; i++ {

			intAry.intAry = append(intAry.intAry, 0)
		}

		intAry.intAryLen = len(intAry.intAry)

		intAry.precision = roundToPrecision

		err = new(intAryNanobot).setInternalFlags(
			intAry, ePrefix.XCpy("Setting 'ia' Flags"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
					"  ia, ePrefix.XCpy(Setting 'ia' Flags))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// roundToPrecision must be < ia.precision

	intLen := intAry.intAryLen - intAry.precision

	newIntAryLen := intLen + roundToPrecision

	fracIdx := intLen

	fracRoundIdx := fracIdx + roundToPrecision

	t := make([]uint8, intAry.intAryLen+1)

	n1 := uint8(0)

	n2 := uint8(0)

	carry := uint8(0)

	for i := fracRoundIdx; i >= 0; i-- {

		n1 = intAry.intAry[i]

		if i == fracRoundIdx {

			n2 = n1 + 5

		} else {

			n2 = n1 + carry

		}

		if n2 > 9 {

			carry = 1

			n2 = n2 - 10

		} else {

			carry = 0
		}

		t[i+1] = n2
	}

	intAry.intAry = []uint8{}

	if carry > 0 {

		t[0] = carry

		intAry.intAry = t[0 : newIntAryLen+1]

	} else {

		intAry.intAry = t[1 : newIntAryLen+1]
	}

	intAry.precision = roundToPrecision

	intAry.intAryLen = len(intAry.intAry)

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix.XCpy("Setting 'ia' Flags"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
				"  ia, ePrefix.XCpy(Setting 'ia' Flags))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setPrecision
//
//	Sets the precision of the IntAry instance input parameter 'ia'
//	to the value of input parameter 'precision'.
//
//	If input parameter 'roundResult' is set to 'true', the
//	resulting numeric value will be rounded to the number of
//	decimal places specified by input parameter 'precision'.
//
//	If input parameter 'roundResult' is set to 'false', the
//	resulting numeric value will be truncated to the number of
//	decimal places specified by input parameter 'precision'.
//
//	If 'precision' is set to a value less than zero, an error will
//	be returned.
//
//	If 'precision' is greater than the existing precision, trailing
//	zeros will be added
//
//	Examples
//	========
//
//	   Original           'newPrecision'         Resulting
//	    Value             input parameter          Value
//	--------------        ---------------      --------------
//
//	  654.123456                 9              654.123456000
//	  654.123456                 4              654.1235
//
//	 -654.123456                 9             -654.123456000
//	 -654.123456                 4             -654.1235
//
//	    0                        3                0.000
//	    0.000000                 0                0
func (iaMolecule *intAryMolecule) setPrecision(
	ia *IntAry,
	validateIntAry bool,
	precision int,
	roundResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMolecule.lock.Lock()

	defer iaMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMolecule.setPrecision()",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIntAry,
		ePrefix.XCpy("Validate on Entry"))

	if err != nil {
		return err
	}

	if precision < 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'precision' value is less than ZERO!\n"+
			"precision= '%v'\n",
			ePrefix,
			precision)

	}

	err = new(intAryElectron).isValidIntAry(
		ia,
		ePrefix.XCpy("Validating 'ia'").String())

	if err != nil {
		return err
	}

	if ia.isZeroValue {

		nsProfile := NumSepsProfileSelection{
			SourceObjectName:         "ia",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          NumericSeparatorDto{},
		}

		err = new(intAryQuark).setIntAryToZero(
			ia,
			nil,
			nsProfile,
			uint(precision),
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  ia, nil, nsProfile,uint(precision), ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	if precision == ia.precision {
		return nil
	}

	if precision > ia.precision {

		deltaPrecision := precision - ia.precision

		for i := 0; i < deltaPrecision; i++ {
			ia.intAry = append(ia.intAry, 0)
		}

		ia.precision = precision
		ia.intAryLen = len(ia.intAry)

		err = new(intAryNanobot).setInternalFlags(
			ia, ePrefix.XCpy("Setting 'ia' Flags"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
					"  ia, ePrefix.XCpy(Setting 'ia' Flags))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// Must Be ia.precision > precision

	if roundResult {

		//err = ia.RoundToPrecision(precision)
		// Sets Internal flags
		err = new(intAryMolecule).roundToPrecision(
			ia,
			false,
			precision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryMolecule).roundToPrecision(\n" +
					"  ia, false, precision, ePrefix",
				ErrContext: fmt.Sprintf("precision= '%v'", precision),
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	intLen := ia.intAryLen - ia.precision

	newAryLen := intLen + precision

	ia.intAry = ia.intAry[0:newAryLen]

	ia.intAryLen = newAryLen

	ia.precision = precision

	err = new(intAryNanobot).setInternalFlags(
		ia, ePrefix.XCpy("Setting 'ia' Flags On Exit"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
				"  ia, ePrefix.XCpy(Setting 'ia' On Exit Flags))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// suffixToIntAry
//
//	Adds an integer of type uint8 to the end of the existing
//	internal integer array maintained by the current IntAry object.
func (iaMolecule *intAryMolecule) suffixToIntAry(
	ia *IntAry,
	num uint8) {

	if ia == nil {
		return
	}

	ia.intAry = append(ia.intAry, num)

	ia.intAryLen = len(ia.intAry)

	ia.integerLen = ia.intAryLen - ia.precision

	new(intAryNanobot).setInternalFlagsNoErrors(ia)

	return
}
