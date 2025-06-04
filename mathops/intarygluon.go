package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

type intAryGluon struct {
	lock sync.Mutex
}

// setIntAryWithInt
//
//	Sets the value of the current intAry object to that of the
//	input parameter 'intDigits', an integer of type 'int'.
//
//	Input parameter 'precision' to indicate the number of digits to
//	the right of the decimal place. Input parameter 'precision' is
//	of type uint.
//
//	The numeric sign (plus or minus) of the resulting intAry value
//	is determined by the sign of input parameter,'intDigits'.
//
//	Example
//	=======
//
//	intDigits      precision      result
//	---------      ---------      ------
//
//	  946254            3          946.254
//	  946254            0          946254
//	 -946254            3         -946.254
//	 -946254            0         -946254
//
//	IMPORTANT
//	=========
//
//	The maximum value for input parameter 'precision' is
//	2,147,483,647. This is the maximum value for a 32-bit
//	signed integer. The limitation derives from the maximum
//	length of arrays in 'Go'. Type IntAry relies on arrays
//	of 8-bit integers to store numeric values.
//
//	Input Parameters
//	================
//
//	intAry                   *IntAry
//	  A pointer to an IntAry object. This object will be
//	  reconfigured with a new value based on the following
//	  input parameters.
//
//	nsProfile                NumSepsProfileSelection
//	 This struct contains all the prameters and options
//	 necessary to generate the NumericSeparatorsDto which is
//	 required for configuration of Numeric Separators in the
//	 IntAry object returned by this method.
//
//	intDigits                int
//	  The numeric digits contained in this value comprise both
//	  the integer digits and the fractional digits which will be
//	  configured in the final numeric value stored in parameter,
//	  'intAry'.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in 'intAry'
//
//	  Although 'precision' is an unsigned integer type, the maximum
//	  value allowed for this parameter is 2,147,483,647.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (iaGluon *intAryGluon) setIntAryWithInt(
	intAry *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	intDigits int,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaGluon.lock.Lock()

	defer iaGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithInt",
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

	maxInt := big.NewInt(math.MaxInt)

	precisionParam := big.NewInt(0).SetUint64(uint64(precision))

	if precisionParam.Cmp(maxInt) > 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input precision 'parameter' is INVALID!\n"+
				"The maximum value allowed for 'precision' is %v\n"+
				"Input parameter 'precision' is %v",
				maxInt.Text(10), precisionParam.Text(10)),
		}
	}

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "intAry"
		actualNumSepsSrcIntAryPtr = intAry

	} else {

		nsProfile.SourceObjectName = "numSepsSrcIntAry"
		actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
	}

	numSeps, err = new(intAryUtility).selectNumericSeparators(
		actualNumSepsSrcIntAryPtr,
		nsProfile,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	quotient := 0
	mod := 0

	intAry.intAry = []uint8{}
	intAry.intAryLen = 0
	intAry.precision = int(precision)
	intAry.signVal = 1

	if intDigits < 0 {

		intDigits = intDigits * -1

		intAry.signVal = -1
	}

	if intDigits == 0 {

		nsProfile = NumSepsProfileSelection{
			SourceObjectName:         "intAry",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          numSeps,
		}

		// ia.SetIntAryToZero(precision)
		err = new(intAryQuark).setIntAryToZero(
			intAry,
			nil,
			nsProfile,
			precision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  intAry, nil, nsProfile-numSeps, precision, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	for {

		if intDigits == 0 {
			break
		}

		quotient = intDigits / 10

		mod = intDigits - (quotient * 10)

		intAry.intAry = append(intAry.intAry, uint8(mod))

		intAry.intAryLen++

		intDigits = quotient

	}

	n1 := uint8(0)

	lastIdx := intAry.intAryLen - 1

	totalLen := intAry.intAryLen / 2

	for i := 0; i < totalLen; i++ {

		n1 = intAry.intAry[i]

		intAry.intAry[i] = intAry.intAry[lastIdx]

		intAry.intAry[lastIdx] = n1

		lastIdx--
	}

	err = new(intAryPhoton).setNumericSeparatorsDto(
		intAry, numSeps, false, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryPhoton).setNumericSeparatorsDto(\n" +
				"  intAry, numSeps, true, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	//ia.SetInternalFlags()
	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix.XCpy("Setting 'intAry' Flags"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
				"  intAry, ePrefix.XCpy(Setting 'intAry' Flags))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithUint8Ary
//
//	This method is designed to set the value of the current IntAry
//	object by passing in a pointer to an unsigned integer array
//	([]uint8). []uint8 is the type of array native to the IntAry
//	object.
//
//	Input parameter 'precision' will determine the number of digits
//	to the right of the decimal place.
//
//	Input parameter 'signVal' must be either +1 or -1 indicating
//	the	sign of the number represented by the integer array.
//
//	If signVal is not equal to +1 or -1, an error will be
//	generated.
//
//	The Numeric Separators originaly configured for the current
//	IntAry instance will NOT be modified.
func (iaGluon *intAryGluon) setIntAryWithUint8Ary(
	intAry *IntAry,
	iAry2 []uint8,
	precision uint,
	signVal int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaGluon.lock.Lock()

	defer iaGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithUint8Ary",
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

	if signVal != 1 && signVal != -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input parameter 'signVal' is INVALID!\n"+
				"signVal MUST HAVE a value of -1 or +1.\n"+
				"signVal='%v'", signVal),
		}
	}

	lIAry2 := len(iAry2)

	intAry.intAry = make([]uint8, lIAry2)

	for i := 0; i < lIAry2; i++ {

		intAry.intAry[i] = iAry2[i]
	}

	intAry.intAryLen = lIAry2

	intAry.precision = int(precision)

	intAry.signVal = signVal

	err = new(intAryNanobot).setInternalFlags(intAry, ePrefix.XCpy("Set 'intAry' Flags"))

	if err != nil {
		return err
	}

	if intAry.isIntegerZeroValue && intAry.integerLen > 1 {

		err = new(intAryAtom).optimizeIntArrayLen(intAry, false, false, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}

// setIntAryWithUint64
//
//		Sets the value of the current IntAry object equal to that of
//		the input parameter 'intDigits', a 64-bit unsigned integer.
//
//		Note: Input parameter 'precision' to indicate the number of
//		digits to the right of the decimal place.
//
//		Input parameter, 'signVal' must be set to one of two values:
//		-1 or +1. 'signVal' determines the numeric sign of the
//		resulting IntAry value, either plus or minus.
//
//		Example
//		=======
//
//		intDigits  precision  signVal    result
//
//		 946254        3         1       946.254
//		 946254        0         1       946254
//		 946254        3        -1      -946.254
//		 946254        0        -1      -946254
//
//		Input Parameters
//		================
//
//		intAry                   *IntAry
//		  A pointer to an IntAry object. This object will be
//		  reconfigured with a new value based on the following
//		  input parameters.
//
//		numSepsSrcIntAry         *IntAry
//		  If this pointer is NOT 'nil', the Numeric Separators
//	   will be taken from this IntAry Object.
//
//	   If this pointer is 'nil', it will be ignored and
//	   the source of Numeric Separators will either be
//	   the 'intAry' object or standard defaults as specified
//	   by input parameter 'nsProfile'.
//		  reconfigured with a new value based on the following
//		  input parameters.
//
//		nsProfile                NumSepsProfileSelection
//		 This struct contains all the prameters and options
//		 necessary to generate the NumericSeparatorsDto which is
//		 required for configuration of Numeric Separators in the
//		 IntAry object returned by this method.
//
//		intDigits                int
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in parameter,
//		  'intAry'.
//
//		signVal                  int
//		  Input parameter 'signVal' must be set to one of two values:
//		  +1 or -1. This value is used to signal the sign of the
//		  resulting numeric value. +1 identifies a positive number and
//		  -1 identifies a negative number. 'signVal' determines the
//		  numeric sign of the resulting IntAry value, either plus or
//		  minus.
//
//		precision                uint
//		  'precision' specifies the number of fractional digits in the
//		  final numeric value stored in 'intAry'
//
//		  Although 'precision' is an unsigned integer type, the maximum
//		  value allowed for this parameter is 2,147,483,647.
//
//		Return Values
//		=============
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (iaGluon *intAryGluon) setIntAryWithUint64(
	intAry *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	intDigits uint64,
	signVal int,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithInt",
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

	if signVal != 1 && signVal != -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input parameter 'signVal' is INVALID!\n"+
				"The only valid values for 'signVal' are -1 or +1.\n"+
				"signVal='%v'", signVal),
		}
	}

	maxInt := big.NewInt(math.MaxInt)

	precisionParam := big.NewInt(0).SetUint64(uint64(precision))

	if precisionParam.Cmp(maxInt) > 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input precision 'parameter' is INVALID!\n"+
				"The maximum value allowed for 'precision' is %v\n"+
				"Input parameter 'precision' is %v",
				maxInt.Text(10), precisionParam.Text(10)),
		}
	}

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "intAry"
		actualNumSepsSrcIntAryPtr = intAry

	} else {

		nsProfile.SourceObjectName = "numSepsSrcIntAry"
		actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
	}

	numSeps, err = new(intAryUtility).selectNumericSeparators(
		actualNumSepsSrcIntAryPtr,
		nsProfile,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	intAry.signVal = signVal

	if intDigits == 0 {

		nsProfile = NumSepsProfileSelection{
			SourceObjectName:         "intAry",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          numSeps,
		}

		// ia numSeps -> final ia
		err = new(intAryQuark).setIntAryToZero(
			intAry,
			nil,
			nsProfile,
			precision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  ia, nil, nsProfile-numSeps, precision, ePrefix)",
				ErrContext: fmt.Sprintf("precision= '%v'", precision),
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	intAry.precision = int(precision)

	quotient := uint64(0)
	mod := uint64(0)
	ten := uint64(10)

	intAry.intAry = []uint8{}
	intAry.intAryLen = 0
	for {

		if intDigits == 0 {
			break
		}

		quotient = intDigits / ten

		mod = intDigits - (quotient * ten)

		intAry.intAry = append(intAry.intAry, uint8(mod))

		intAry.intAryLen++

		intDigits = quotient

	}

	n1 := uint8(0)

	lastIdx := intAry.intAryLen - 1

	totalLen := intAry.intAryLen / 2

	for i := 0; i < totalLen; i++ {

		n1 = intAry.intAry[i]

		intAry.intAry[i] = intAry.intAry[lastIdx]

		intAry.intAry[lastIdx] = n1

		lastIdx--
	}

	// original ia numSeps -> final ia
	err = new(intAryPhoton).setNumericSeparatorsDto(
		intAry, numSeps, false, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryPhoton).setNumericSeparatorsDto(\n" +
				"ia, numSeps, false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

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
