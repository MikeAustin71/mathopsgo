package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type intAryGluon struct {
	lock sync.Mutex
}

// setIntAryWithBigInt
//
//	Sets the current value of the intAry to the value of input
//	parameter 'intDigits'. The sign value (plus or minus) is taken
//	from the input parameter, 'intDigits'.
//
//	The precision or number of digits to the right of the decimal
//	point, is determined by the input parameter, 'precision'. Input
//	parameter 'precision' must be passed as a positive value.
//
//	In practice, the maximum limit for 'precision' will be
//	constrained by the maximum array size permitted by	your
//	system.
//
//	Negative 'precision' values will trigger an error.
//
//	Example
//	=======
//
//	intDigits      precision      result
//
//	 946254            3           946.254
//	 946254            0           946254
//	-946254            3          -946.254
//	-946254            0          -946254
func (iaGluon *intAryGluon) setIntAryWithBigInt(
	ia *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	intDigits *big.Int,
	precision int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaGluon.lock.Lock()

	defer iaGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithBigInt()",
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

	if intDigits == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intDigits'",
		}
	}

	if precision < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is a negative value!\n"+
				"precision='%v'", precision),
		}

	}

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "ia"
		actualNumSepsSrcIntAryPtr = ia

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

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	bigZero := big.NewInt(0)

	quotient := big.NewInt(0)

	mod := big.NewInt(0)

	big10 := big.NewInt(10)

	modX := big.NewInt(0)

	xIntDigits := big.NewInt(0).Set(intDigits)

	compare := bigZero.Cmp(xIntDigits)

	ia.intAry = []uint8{}

	ia.intAryLen = 0

	ia.precision = precision

	ia.signVal = 1

	if compare == 1 {

		bigMinus1 := big.NewInt(0).SetInt64(int64(-1))

		xIntDigits = big.NewInt(0).Mul(xIntDigits, bigMinus1)

		ia.signVal = -1
	}

	if compare == 0 {

		err = new(intAryQuark).setIntAryToZero(
			ia,
			nil,
			nsProfile2,
			uint(precision),
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryToZero(\n"+
					"ia, nil, nsProfile2, uint(precision)= '%v', ePrefix", precision),
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	for {

		compare := bigZero.Cmp(xIntDigits)

		if compare == 0 {
			break
		}

		quotient, mod = big.NewInt(0).QuoRem(xIntDigits, big10, modX)

		ia.intAry = append(ia.intAry, uint8(mod.Int64()))

		ia.intAryLen++

		xIntDigits.Set(quotient)

	}

	n1 := uint8(0)

	lastIdx := ia.intAryLen - 1

	totalLen := ia.intAryLen / 2

	for i := 0; i < totalLen; i++ {

		n1 = ia.intAry[i]

		ia.intAry[i] = ia.intAry[lastIdx]

		ia.intAry[lastIdx] = n1

		lastIdx--
	}

	// If no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia,\"ia\", validateResult= '%v', ePrefix)", validateResult),
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithInt
//
//		Receives an IntAry instance ('intAry') and reconfigures the
//		numeric values based on input parameters 'intDigits' and
//		'precision'.
//
//		Input parameter 'precision' to indicate the number of digits to
//		the right of the decimal place. Input parameter 'precision' is
//		of type uint.
//
//		The numeric sign (plus or minus) of the resulting intAry value
//		is determined by the sign of input parameter,'intDigits'.
//
//		Example
//		=======
//
//		intDigits      precision      result
//		---------      ---------      ------
//
//		  946254            3          946.254
//		  946254            0          946254
//		 -946254            3         -946.254
//		 -946254            0         -946254
//
//		IMPORTANT
//		=========
//
//		  In practice, the maximum limit for 'precision' will be
//		  constrained by the maximum array size permitted by
//		  your system. Type IntAry relies on arrays of 8-bit
//		  integers for numeric value storage.
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
//		  This instance of IntAry may be populated by the calling
//		  function as a source of Numeric Separators
//
//		nsProfile                NumSepsProfileSelection
//		 This struct contains all the prameters and options
//		 necessary to generate the NumericSeparatorsDto. This
//	  Data Transfer Object (Dto) is required for configuration
//	  of Numeric Separators in the IntAry object returned by
//	  this method.
//
//		intDigits                int
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in parameter,
//		  'intAry'.
//
//		precision                uint
//		  'precision' specifies the number of fractional digits in the
//		  final numeric value stored in 'intAry'
//
//		  In practice, the maximum limit for 'precision' will be
//		  constrained by the maximum array size permitted by
//		  your system. Type IntAry relies on arrays of 8-bit
//		  integers for numeric value storage.
//
//		validateResult           bool
//		  When set to 'true', the final numeric value calculated by
//		  this method will be subjected to validation testing.
//
//		Return Values
//		=============
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (iaGluon *intAryGluon) setIntAryWithInt(
	intAry *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	intDigits int,
	precision uint,
	validateResult bool,
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

	// If no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"intAry,\"intAry\", validateResult='%v', ePrefix)", validateResult),
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithIntAry
//
//	Sets the value of the current intAry based on an array of
//	integers i.e. []int.
//
//	Input parameter 'precision' will determine the number of
//	digits to the right of the decimal place.
//
//	In practice, the maximum limit for 'precision' will be
//	constrained by the maximum array size permitted by	your
//	system.
//
//	Input parameter 'signVal' must be either +1 or -1 indicating
//	the sign of the number represented by the integer array.
//
//	If signVal is not equal to +1 or -1,	an error will be
//	generated.
func (iaGluon *intAryGluon) setIntAryWithIntAry(
	ia *IntAry,
	validateIa bool,
	iAry2 []int,
	precision uint,
	signVal int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaGluon.lock.Lock()

	defer iaGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithIntAry",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia, 'ia', validateIa= '%v', ePrefix)", validateIa),
			ErrContext: "Valiate on Startup",
			ErrMessage: err.Error(),
		}
	}

	if signVal != 1 && signVal != -1 {

		return fmt.Errorf("SetIntAryWithIntAry()\n"+
			"Error: Input parameter 'signVal' parameter is INVALID!\n"+
			"signVal must be -1 or +1.\n"+
			"signVal='%v'", signVal)

	}

	lIAry2 := len(iAry2)

	ia.intAry = make([]uint8, lIAry2)

	for i := 0; i < lIAry2; i++ {

		ia.intAry[i] = uint8(iAry2[i])
	}

	ia.intAryLen = lIAry2

	ia.precision = int(precision)

	ia.signVal = signVal

	// Even if no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia,\"ia\", validateResult='%v', ePrefix)", validateResult),
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithIntAryObj
//
//	Sets the value of 'iAry1Destination' equal to that of
//	'iAry2Source'.
func (iaGluon *intAryGluon) setIntAryWithIntAryObj(
	iAry1Destination *IntAry,
	iAry2Source *IntAry,
	validateIAry2Source bool,
	copyBackUp bool,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaGluon.lock.Lock()

	defer iaGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithIntAry",
		"")

	if err != nil {
		return err
	}

	if iAry1Destination == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'iAry1Destination'",
		}
	}

	if iAry2Source == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'iAry2Source'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		iAry2Source,
		"iAry2Source",
		validateIAry2Source,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"  iAry2Source, 'iAry2Source', \n"+
				"  validateIAry2Source= '%v', ePrefix)", validateIAry2Source),
			ErrContext: "Valiate 'iAry2Source' on Startup",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryProton).copy(iAry1Destination, iAry2Source, false, copyBackUp, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryProton).copy(\n"+
				"  iAry1Destination, iAry2Source, validateSource=false, copyBackUp=%v, ePrefix)",
				copyBackUp),
			ErrContext: "iAry2 -> ia",
			ErrMessage: err.Error(),
		}
	}

	// Even if no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		iAry1Destination,
		"iAry1Destination",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"  iAry1Destination,\"iAry1Destination\",\n"+
				"  validateResult='%v', ePrefix)", validateResult),
			ErrContext: "Validating Final Result 'iAry1Destination' on Exit",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithInt64
//
//		Sets the value of the current intAry object to that of the
//		input parameter 'int64Num', a 64-bit integer.
//
//		Input parameter 'precision' indicates the number of digits
//		to be formatted to the right of the decimal place. Input
//		parameter 'precision' is of type uint.
//
//		In practice, the maximum limit for 'precision' will be
//		constrained by the maximum array size permitted by
//		your system.
//
//		The numeric sign (plus or minus) of the resulting intAry value
//		is determined by the sign of input parameter 'int64Num'.
//
//		Example
//		=======
//
//		int64Num      precision      result
//
//		 946254           3           946.254
//		 946254           0           946254
//		-946254           3          -946.254
//		-946254           0          -946254
//
//		Input Parameters
//		================
//
//		ia                       *IntAry
//		  This instance of IntAry will be overwritten with the new
//		  IntAry values computed by this method.
//
//		numSepsSrcIntAry         *IntAry
//		  This instance of IntAry may be populated by the calling
//		  function as a source of Numeric Separators
//
//		nsProfile                NumSepsProfileSelection
//		  This struct contains decision parameters for selecting and
//		  configuring Numeric Separators.
//
//		int64Num                 int64
//		  The numeric digits contained in this value comprise both
//		  the integer digits and the fractional digits which will be
//		  configured in the final numeric value stored in the IntAry
//		  object returned by this method.
//
//		precision                uint
//	   'precision' specifies the number of fractional digits in the
//	   final numeric value stored in the returned IntAry object.
//
//	   In practice, the maximum limit for 'precision' will be
//	   constrained by the maximum array size permitted by
//	   your system.
//
//		validateResult           bool
//		  When set to 'true', the final numeric value calculated by
//		  this method will be subjected to validation testing.
//
//		Return Values
//		=============
//
//		IntAry
//		  This new instance of IntAry will be returned configured with
//		  the numeric value calculated from input parameters, 'intNum'
//		  and 'precision'.
//
//		error
//		  If no errors are encountered during processing, this returned
//		  value will be set to 'nil'
func (iaGluon *intAryGluon) setIntAryWithInt64(
	ia *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	int64Num int64,
	precision uint,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaGluon.lock.Lock()

	defer iaGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryGluon.setIntAryWithInt64",
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

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "ia"
		actualNumSepsSrcIntAryPtr = ia

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

	quotient := int64(0)

	mod := int64(0)

	i64Ten := int64(10)

	ia.intAry = []uint8{}

	ia.intAryLen = 0

	ia.precision = int(precision)

	ia.signVal = 1

	if int64Num < 0 {

		int64Num = int64Num * int64(-1)

		ia.signVal = -1
	}

	if int64Num == 0 {

		//ia.SetIntAryToZero(precision)

		nsProfile = NumSepsProfileSelection{
			SourceObjectName:         "ia",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          numSeps,
		}

		// ia.SetIntAryToZero(precision)
		err = new(intAryQuark).setIntAryToZero(
			ia,
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

		if int64Num == 0 {
			break
		}

		quotient = int64Num / i64Ten

		mod = int64Num - (quotient * i64Ten)

		ia.intAry = append(ia.intAry, uint8(mod))

		ia.intAryLen++

		int64Num = quotient

	}

	n1 := uint8(0)

	lastIdx := ia.intAryLen - 1

	totalLen := ia.intAryLen / 2

	for i := 0; i < totalLen; i++ {

		n1 = ia.intAry[i]

		ia.intAry[i] = ia.intAry[lastIdx]

		ia.intAry[lastIdx] = n1

		lastIdx--
	}

	//ia.SetInternalFlags()

	err = new(intAryPhoton).setNumericSeparatorsDto(
		ia, numSeps, false, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryPhoton).setNumericSeparatorsDto(\n" +
				"  ia, numSeps, true, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	// If no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryUtility).selectIntAryValidation(\n" +
				"ia,\"ia\", validateResult, ePrefix)",
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

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryAtom).optimizeIntArrayLen(\n" +
				"intAry, validateIntAry=false, optimizeFracDigits=false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithUint64
//
//	Sets the value of the current IntAry object equal to that of
//	the input parameter 'intDigits', a 64-bit unsigned integer.
//
//	Note: Input parameter 'precision' to indicate the number of
//	digits to the right of the decimal place.
//
//	Input parameter, 'signVal' must be set to one of two values:
//	-1 or +1. 'signVal' determines the numeric sign of the
//	resulting IntAry value, either plus or minus.
//
//	Example
//	=======
//
//	intDigits  precision  signVal    result
//
//	 946254        3         1       946.254
//	 946254        0         1       946254
//	 946254        3        -1      -946.254
//	 946254        0        -1      -946254
//
//	Input Parameters
//	================
//
//	intAry                   *IntAry
//	  A pointer to an IntAry object. This object will be
//	  reconfigured with a new value based on the following
//	  input parameters.
//
//	numSepsSrcIntAry         *IntAry
//	  If this pointer is NOT 'nil', the Numeric Separators
//	 will be taken from this IntAry Object.
//
//	 If this pointer is 'nil', it will be ignored and
//	 the source of Numeric Separators will either be
//	 the 'intAry' object or standard defaults as specified
//	 by input parameter 'nsProfile'.
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
//	signVal                  int
//	  Input parameter 'signVal' must be set to one of two values:
//	  +1 or -1. This value is used to signal the sign of the
//	  resulting numeric value. +1 identifies a positive number and
//	  -1 identifies a negative number. 'signVal' determines the
//	  numeric sign of the resulting IntAry value, either plus or
//	  minus.
//
//	precision                uint
//	  'precision' specifies the number of fractional digits in the
//	  final numeric value stored in 'intAry'
//
//	  In practice, the maximum limit for 'precision' will be
//	  constrained by the maximum array size permitted by
//	  your system.
//
//	validateResult           bool
//	  When this parameter is set to 'true' the final calculation
//	  result generated by this method will be subjected to
//	  validation tests.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  value will be set to 'nil'
func (iaGluon *intAryGluon) setIntAryWithUint64(
	intAry *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	intDigits uint64,
	signVal int,
	precision uint,
	validateResult bool,
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

	// If no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryUtility).selectIntAryValidation(\n" +
				"intAry,\"intAry\", validateResult, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
