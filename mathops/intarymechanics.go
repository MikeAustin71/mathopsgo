package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryMechanics struct {
	lock *sync.Mutex
}

// decrementIntegerOne
//
//	Decrements the numeric value of the IntAry instance ('intAry')
//	by subtracting '1'.
func (iaMech *intAryMechanics) decrementIntegerOne(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaMech.lock == nil {
		iaMech.lock = new(sync.Mutex)
	}

	iaMech.lock.Lock()

	defer iaMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMechanics.decrementIntegerOne()",
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

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(
			intAry, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy(Validating 'intAry').String())",
				ErrContext: "IntAry instanace 'intAry' is INVALID!\n" +
					"'intAry' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
		}

	}

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(intAry, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if intAry.isZeroValue || intAry.isIntegerZeroValue {
		intAry.signVal = -1
	}

	intLen := intAry.intAryLen - intAry.precision
	intIdx := intLen - 1
	lastIdx := intAry.intAryLen - 1

	n1 := 0
	n2 := 0
	carry := 0

	intAry.isZeroValue = true
	intAry.isIntegerZeroValue = true

	for i := lastIdx; i >= 0; i-- {
		n1 = int(intAry.intAry[i])

		if i > intIdx {
			//  i > intIdx
			// This must be a fractional digit
			// Retain fractional digits

			if n1 != 0 {
				intAry.isZeroValue = false
			}

			continue

		} else if i == intIdx {

			n2 = n1 + (-1 * intAry.signVal)

			if n2 < 0 {
				n2 = n1 + 10 - 1
				carry = 1

			} else if n2 > 9 {
				n2 = n1 + 1 - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// Must be i < intIdx

			n2 = n1 + ((intAry.signVal * carry) * -1)

			if n2 < 0 {
				n2 = n1 + 10 - carry
				carry = 1
			} else if n2 > 9 {
				n2 = n1 - 10 + carry
				carry = 1
			} else {
				carry = 0
			}

		}

		if n2 != 0 {
			intAry.isZeroValue = false
			intAry.isIntegerZeroValue = false
		}

		intAry.intAry[i] = uint8(n2)

	}

	if intAry.isZeroValue && carry == 0 {
		intAry.signVal = 1
	}

	if carry > 0 {

		intAry.intAry = append([]uint8{1}, intAry.intAry...)
		intAry.intAryLen++

	} else if intAry.intAry[0] == 0 && intLen > 1 {
		intAry.intAry = intAry.intAry[1:]
		intAry.intAryLen--
	}

	return nil
}

// getMagnitudeDigits
//
//	Returns the number of digits in the integer portion of the
//	numeric value in the IntAry instance passed as input parameter
//	'ia'.
//
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaMech *intAryMechanics) getMagnitudeDigits(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if iaMech.lock == nil {
		iaMech.lock = new(sync.Mutex)
	}

	iaMech.lock.Lock()

	defer iaMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMechanics.getMagnitudeDigits()",
		"")

	if err != nil {
		return 0, err
	}

	if intAry == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return 0,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
						"  intAry, ePrefix.XCpy(Validating 'intAry').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix)

	if err != nil {

		return 0, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	intAryMagnitude :=
		intAry.intAryLen - intAry.precision - intAry.firstDigitIdx

	return intAryMagnitude, nil
}

// inverseIntAry
//
//	Returns the inverse BigIntNum of the input parameter IntAry
//	('intAry') value.
//
//	  IntAry = input parameter 'intAry'
//	  Inverse = 1 ÷ IntAry
//
//	Input Parameter
//	===============
//
//	maxPrecision             int
//	  Determines the number of digits to the right of the decimal
//	  point in the result.
//
//	  If 'maxPrecision' is set equal to negative one (-1), the
//	  maximum number of decimal digits is automatically set to
//	  4096 digits to the right of the decimal	place.
func (iaMech *intAryMechanics) inverseIntAry(
	ia *IntAry,
	validateIntAry bool,
	maxPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	if iaMech.lock == nil {
		iaMech.lock = new(sync.Mutex)
	}

	iaMech.lock.Lock()

	defer iaMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMechanics.inverseIntAry()",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if ia == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'ia'",
			}
	}

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(ia, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
						"  ia, ePrefix.XCpy(Validating 'ia').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	if maxPrecision < 0 {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: fmt.Sprintf("ERROR: Input parameter 'maxPrecision' is INVALID.\n"+
					"'maxPrecision' cannot be less than zero.\n"+
					"maxPrecision= '%v'", maxPrecision),
			}
	}

	internalPrecision := maxPrecision + 50

	iaOne := new(intAryElectron).newIntAry()

	nsProfile := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          NumericSeparatorDto{},
	}

	err = new(intAryGluon).setIntAryWithInt(
		&iaOne,
		ia,
		nsProfile,
		1,
		0,
		true,
		ePrefix.XCpy("Creating iaOne"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryGluon).setIntAryWithInt(\n" +
					"  &iaOne, ia,nsProfile, 1, 0, ePrefix.XCpy(Creating iaOne))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iaInverse, err := new(intAryNeutron).
		divideIntArys(&iaOne, true, ia, false, 0, internalPrecision, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "iaInverse, err := new(intAryNeutron).divideIntArys(\n" +
					"  &iaOne, true, ia, false, 0, internalPrecision, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if iaInverse.GetPrecision() > maxPrecision {

		err = new(intAryMolecule).roundToPrecision(
			&iaInverse,
			false,
			maxPrecision,
			ePrefix)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = new(intAryMolecule).roundToPrecision(&iaInverse, false, maxPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
					ErrMessage: err.Error(),
				}
		}
	}

	return iaInverse, nil
}

// isIntAryEvenNumber
//
//	Returns 'true' if the current IntAry numeric value is evenly
//	divisible by two (2) with no remainder.
//
//	Even Number Definition:
//	  https://www.mathsisfun.com/definitions/even-number.html
func (iaMech *intAryMechanics) isIntAryEvenNumber(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	if iaMech.lock == nil {
		iaMech.lock = new(sync.Mutex)
	}

	iaMech.lock.Lock()

	defer iaMech.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMechanics.setAbsoluteValue()",
		"")

	if err != nil {
		return false, err
	}

	if intAry == nil {

		return false,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'ia'",
			}
	}

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return false,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
						"  intAry, ePrefix.XCpy(Validating 'intAry').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix)

	if err != nil {

		return false, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if intAry.precision > 0 {

		return false, nil

	}

	if intAry.isZeroValue {
		return true, nil
	}

	tNum := IntAry{}

	err = new(intAryProton).copy(
		&tNum,
		intAry,
		false,
		true,
		ePrefix.XCpy("Copy 'ia' -> 'tNum'"))

	if err != nil {
		return false, err
	}

	err = new(IntAryMathDivide).DivideByTwo(&tNum)

	if err != nil {

		return false, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathDivide).DivideByTwo(&tNum)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if tNum.precision > 0 {
		return false, nil
	}

	return true, nil
}

// setAbsoluteValue
//
//	Converts the numeric value of the IntAry object passed as
//	input parameter 'IntAry' to its absolute value.
func (iaMech *intAryMechanics) setAbsoluteValue(
	intAry *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaMech.lock == nil {
		iaMech.lock = new(sync.Mutex)
	}

	iaMech.lock.Lock()

	defer iaMech.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMechanics.setAbsoluteValue()",
		"")

	if err != nil {
		return err
	}

	err = new(intAryNeutron).setSign(intAry, 1, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNeutron).\n" +
				"   setSign(intAry, 1, ePrefix)\n",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
