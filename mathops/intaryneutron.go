package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type intAryNeutron struct {
	lock sync.Mutex
}

// addIntAryToThis
//
//	Adds the value of intAry parameter ia2 to the value of 'ia'.
//
//	Input Parameters
//	================
//
//	ia2                      *IntAry
//	  The numeric value of this incoming IntAry object will be
//	  subtracted from the numeric valud of the current IntAry
//	  instance.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during proceesing the return
//	  value of this parameter will be set to 'nil'.
func (iaNeutron *intAryNeutron) addIntAryToThis(
	ia *IntAry,
	validateIa bool,
	ia2 *IntAry,
	validateIa2 bool,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.addArrayLengthLeft()",
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

	if ia2 == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia2'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix.XCpy("Validating 'ia' on Startup"))

	if err != nil {
		return err
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia2,
		"ia2",
		validateIa2,
		ePrefix.XCpy("Validating 'ia2' on Startup"))

	if err != nil {
		return err
	}

	err = new(IntAryMathAdd).RunTotal(ia, ia2)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathAdd).RunTotal(ia, ia2)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix.XCpy("Final Result Validation"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia, \"ia\", validateResult='%v' ePrefix", validateResult),
			ErrContext: "Error: The Final Result is INVALID!\n" +
				"Final Result 'ia' FAILED Validation Tests",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// addArrayLengthLeft
//
//	 Adds leading zeros to the internal storage array holding the
//		numeric value for the current instance of IntAry.
func (iaNeutron *intAryNeutron) addArrayLengthLeft(
	ia *IntAry,
	validateIa bool,
	addLen int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.addArrayLengthLeft()",
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

	if addLen < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'addLen' is INVALID!\n" +
				"'addLen' has value less than zero.\n" +
				fmt.Sprintf("addLen= '%v'", addLen),
		}

	}

	if addLen == 0 {
		// Nothing to add
		return nil
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(intAryElectron).setIntAryLength(ia, ePrefix.XCpy("Setting 'ia' IntAry Length"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryElectron).setIntAryLength(ia,\n" +
				"  ePrefix.XCpy(Setting 'ia' IntAry Length))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	newLen := addLen + ia.intAryLen

	t := make([]uint8, newLen)

	for i := 0; i < newLen; i++ {

		if i < addLen {

			t[i] = 0

		} else {

			t[i] = ia.intAry[i-addLen]

		}

	}

	ia.intAry = t

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix.XCpy("Validating Final Calculation Result"))

	if err != nil {
		return err
	}

	return nil
}

// AddArrayLengthRight
//
//	Adds trailing zeros to the right of the current intAry.
func (iaNeutron *intAryNeutron) addArrayLengthRight(
	ia *IntAry,
	validateIa bool,
	addLen int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.addArrayLengthLeft()",
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

	if addLen < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'addLen' is INVALID!\n" +
				"'addLen' has value less than zero.\n" +
				fmt.Sprintf("addLen= '%v'", addLen),
		}

	}

	if addLen == 0 {
		// Nothing to add
		return nil
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(intAryElectron).setIntAryLength(ia, ePrefix.XCpy("Setting 'ia' IntAry Length"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryElectron).setIntAryLength(ia,\n" +
				"  ePrefix.XCpy(Setting 'ia' IntAry Length))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	for i := 0; i < addLen; i++ {

		ia.intAry = append(ia.intAry, 0)

	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix.XCpy("Validating Final Calculation Result"))

	if err != nil {
		return err
	}

	return nil
}

// ceiling
//
//	Returns an IntAry which constitutes the mathematical ceiling of
//	the current IntAry.
//
//	Examples
//	========
//
//	      Initial      Ceiling
//	       Value        Value
//	      -------      -------
//	        5.95          6
//	        5.05          6
//	        5             5
//	       -5.05         -5
//	        2.4           3
//	        2.9           3
//	       -2.7          -2
//	       -2            -2
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
//	The IntAry object returned by this method will be configured
//	with Numeric Separators copied from the original input
//	parameter, 'intAry'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) ceiling(
	intAry *IntAry,
	validateIntAry bool,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getBigInt()",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if intAry == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	numSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := intAry.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSeps.SetDefaultsIfEmpty()

	iAry2 := new(intAryElectron).newIntAry()

	intLen := intAry.intAryLen - intAry.precision

	intIdx := intLen - 1

	hasFracDigits, err := new(intAryNanobot).hasFractionalDigits(
		intAry, false, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "hasFracDigits, err := new(intAryNanobot).\n" +
					"  hasFractionalDigits(ia, false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !hasFracDigits {

		// iAry2, err = ia.CopyOut()
		// Acquires 'numSeps' from 'intAry'
		err = new(intAryProton).copy(&iAry2, intAry, false, true, ePrefix)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = new(intAryProton).copy(&iAry2, intAry, false, true, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return iAry2, nil
	}

	err = new(intAryPhoton).setNumericSeparatorsDto(
		&iAry2, numSeps, true, ePrefix.XCpy("intAry numSeps -> iAry2"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryPhoton).setNumericSeparatorsDto(\n" +
					"  &iAry2, numSeps, true, ePrefix.XCpy(intAry numSeps -> iAry2))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if intAry.signVal < 0 {

		t := make([]uint8, intAry.intAryLen)

		for i := 0; i < intLen; i++ {
			t[i] = intAry.intAry[i]
		}

		iAry2.intAry = t[0:]
		iAry2.intAryLen = intAry.intAryLen
		iAry2.precision = intAry.precision
		iAry2.signVal = intAry.signVal
		return iAry2, nil
	}

	t := make([]uint8, intAry.intAryLen+1)

	n1 := 0
	n2 := 0
	carry := 0
	adjFac := 1 * intAry.signVal
	for i := intIdx; i >= 0; i-- {

		n1 = int(intAry.intAry[i])

		if i == intIdx {

			if n1+adjFac < 0 {

				n2 = 10 + n1 + adjFac
				carry = -1

			} else if n1+adjFac > 9 {

				n2 = n1 + adjFac - 10
				carry = 1

			} else {

				n2 = n1 + adjFac
				carry = 0

			}

		} else {

			if n1+carry < 0 {

				n2 = 10 + n1
				carry = -1

			} else if n1+carry > 9 {

				n2 = n1 - 10
				carry = 1

			} else {

				n2 = n1 + carry
				carry = 0
			}
		}

		t[i+1] = uint8(n2)

	}

	if carry != 0 {

		t[0] = uint8(carry)
		iAry2.intAry = t[0 : intAry.intAryLen+1]

	} else {

		iAry2.intAry = t[1 : intAry.intAryLen+1]
	}

	iAry2.intAryLen = len(iAry2.intAry)

	iAry2.precision = intAry.precision

	iAry2.signVal = intAry.signVal

	err = new(intAryUtility).selectIntAryValidation(
		&iAry2,
		"iAry2",
		validateResult,
		ePrefix.XCpy("Validating Final Calculation Result"))

	if err != nil {
		return IntAry{}, err
	}

	return iAry2, nil
}

// changeSign
//
//	Changes the sign of the IntAry instance passed as input
//	parameter 'intAry'.
//
//	If the 'intAry' numeric value is positive (+), this method will
//	change the sign value to negative (-).
//
//	Conversely, if the 'intAry' is a negative (-)	numeric value,
//	this method will change the sign to	positive (+).
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
//	The Numeric Separators originally configured for input
//	parameter IntAry object (intAry), will remain unchanged.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) changeSign(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.changeSign()",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return err
	}

	if intAry.isZeroValue {

		intAry.signVal = 1

		return nil
	}

	if intAry.signVal < 1 {

		intAry.signVal = 1

	} else {

		intAry.signVal = -1

	}

	return nil
}

// divideByTwo
//
//	Divides the numeric value of the input parameter 'intAry' by 2.
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
//	The Numeric Separators originally configured for input
//	parameter IntAry object (intAry), will remain unchanged.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) divideByTwo(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.divideByTwo()",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(IntAryMathDivide).DivideByTwo(intAry)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathDivide).DivideByTwo(ia)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// divideByInt64
//
//	Divide the current value of the IntAry parameter 'intAry' by an
//	int64 'divisor'	parameter passed to the method. The result or
//	'quotient' produced by this divsion operation is stored in the
//	original 'intAry' object.
//
//	If the quotient has a number of decimal places to the right of
//	the decimal point which is greater than 'maxPrecision', the
//	result is rounded to 'maxPrecision' decimal places.
//
//	If 'maxPrecision' is set equal to -1, 'maxPrecision' is
//	automatically set to 4,096.
//
//	If 'maxPrecision' is less than -1, an error will be returned.
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
//	The Numeric Separators originally configured for input
//	parameter IntAry object (intAry), will remain unchanged.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) divideByInt64(
	intAry *IntAry,
	validateIntAry bool,
	divisor int64,
	maxPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.divideByInt64()",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return err
	}

	intAryNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "intAryNumStr, err := new(intAryAtom).getRawNumStr(\n" +
				"  getRawNumStr(intAry, false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(IntAryMathDivide).DivideByInt64(intAry, divisor, maxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathDivide).DivideByInt64(\n" +
				"intAry, divisor, maxPrecision)",
			ErrContext: fmt.Sprintf("intAry = '%v'\ndivisor = '%v'\nmaxPrecision= '%v'",
				intAryNumStr, divisor, maxPrecision),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// divideByTenToPower
//
//	Divide the numerical value of an IntAry instance by
//	10 raised to the power of the input parameter, 'exponent'.
//
//	             'ia'
//	    ia =  -------------
//	         (10^exponent)
//
//	The result or quotient is stored in the IntAry instance passed
//	as input parameter 'intAry'.
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
//	The Numeric Separators originally configured for input
//	parameter IntAry object (intAry), will remain unchanged.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) divideByTenToPower(
	intAry *IntAry,
	validateIntAry bool,
	exponent uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.divideByTenToPower()",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return err
	}

	intAryNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "intAryNumStr, err := new(intAryAtom).getRawNumStr(\n" +
				"  getRawNumStr(intAry, false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(IntAryMathDivide).DivideByTenToPower(intAry, exponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathDivide).DivideByTenToPower(intAry, exponent)",
			ErrContext: fmt.Sprintf("intAry = '%v'\nexponent= '%v'",
				intAryNumStr, exponent),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// divideIntArys
//
//	Divides input parameters iAry1 by iAry2. The result of this
//	division is returned as an IntAry (Quotient).
//
//	Given a ÷ b = c, 'a' is the dividend, 'b' is the divisor and
//	'c' is the quotient. For this method
//
//	    dividend = a = IntAry input parameter 'iAry1'
//	    divisor  = b = IntAry input parameter 'iAry2'
//	    quotient = c = Quotient returned by this method
//
//	               iAry1 ÷ iAry2 = quotient
//
//	Maximum precision of the division result is controlled by the
//	input parameter, 'maxPrecision'.
//
//	If 'maxPrecision' is greater than or equal to zero ('0'), the
//	number of digits to the right of the decimal place will not
//	exceed 'maxPrecision'.
//
//	If 'maxPrecision' is set equal to minus one ('-1'),
//	'maxPrecision' will be automatically set to a maximum of 4,096
//	digits to the right of the decimal point.
//
//	'minPrecision' specifies the minimum precision of the final
//	result. If 'minPrecision' is less than zero, it is automatically
//	set to zero.
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
//	Numeric Separators for the returned 'quotient' are copied from
//	'iAry1'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry1' is set to true, this
//	method will subject 'iAry1' to validation tests.
//
//	Likewise, if input parameter 'validateIntAry2' is set to true, this
//	method will subject 'iAry2' to validation tests.
func (iaNeutron *intAryNeutron) divideIntArys(
	iAry1 *IntAry,
	validateIntAry1 bool,
	iAry2 *IntAry,
	validateIntAry2 bool,
	minPrecision,
	maxPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (quotient IntAry, err error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.divideIntArys",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if iAry1 == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'iAry1'",
			}
	}

	if iAry2 == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'iAry2'",
			}
	}

	iaUtils := new(intAryUtility)

	err = iaUtils.selectIntAryValidation(
		iAry1,
		"iAry1",
		validateIntAry1,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	err = iaUtils.selectIntAryValidation(
		iAry2,
		"iAry2",
		validateIntAry2,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	iAry1NumStr, err := iAry1.GetNumStr()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iAry1NumStr, err := iAry1.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iAry2NumStr, err := iAry2.GetNumStr()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iAry2NumStr, err := iAry2.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	quotient, err = new(IntAryMathDivide).Divide(
		iAry1, iAry2, minPrecision, maxPrecision)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, err = new(IntAryMathDivide).Divide(\n" +
					"  iAry1, iAry2, minPrecision, maxPrecision)",
				ErrContext: fmt.Sprintf("iAry1= '%v'\n iAry2= '%v'\n minPrecision= '%v'\nmaxPrecision= '%v'",
					iAry1NumStr, iAry2NumStr, minPrecision, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	err = new(intAryElectron).isValidIntAry(
		&quotient,
		ePrefix.XCpy("Validating 'quotient'").String())

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "Quotient returned by IntAryMathDivide.Divide is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	return quotient, err
}

// floor
//
//	Math 'Floor' function. Finds the integer number which is less
//	than or equal to the value of the current intAry.
//
//	Reference Wikipedia
//	  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//	Examples
//	========
//
//	  Initial     Floor
//	   Value      Value
//	  -------    -------
//	   5.95         5
//	   5.05         5
//	   5            5
//	  -5.05        -6
//	   2.4          2
//	   2.9          2
//	  -2.7         -3
//	  -2           -2
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
//	The Numeric Separators originally configured for input
//	parameter IntAry object (intAry), will remain unchanged.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) floor(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.floor",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if intAry == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, true, ePrefix.XCpy("IntAry NumSeps"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numSeps, err := new(intAryPhoton).\n" +
					"  getNumericSeparatorsDto(intAry, ePrefix.XCpy(IntAry NumSeps))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iAry2 := new(intAryElectron).newIntAry()

	err = new(intAryPhoton).setNumericSeparatorsDto(
		&iAry2, numSeps, false, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryPhoton).\n" +
					"  setNumericSeparatorsDto(&iAry2, numSeps, false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if intAry.isZeroValue {

		nsProfile := NumSepsProfileSelection{
			SourceObjectName:         "ia",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          numSeps,
		}

		err = new(intAryQuark).
			setIntAryToZero(&iAry2,
				nil,
				nsProfile,
				uint(intAry.precision),
				ePrefix)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
						"intAry, nil, nsProfile-numSeps, uint(intAry.precision), ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	hasFracDigits, err := new(intAryNanobot).hasFractionalDigits(
		intAry, false, ePrefix.XCpy("intAry Frac Digits"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "hasFracDigits, err := new(intAryNanobot).hasFractionalDigits(\n" +
					"intAry, false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !hasFracDigits {
		// There are NO non-zero digits to the
		// right of the decimal place

		err = new(intAryProton).copy(&iAry2, intAry, false, true, ePrefix)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryProton).copy(\n" +
						"  &iAry2, ia, false, true, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return iAry2, err
	}

	intLen := intAry.intAryLen - intAry.precision

	intIdx := intLen - 1

	if intAry.signVal > 0 {
		// There ARE non-zero digits to the right of the
		// decimal place
		t := make([]uint8, intAry.intAryLen)

		for i := 0; i < intLen; i++ {

			t[i] = intAry.intAry[i]
		}

		iAry2.intAry = t[0:]

		iAry2.intAryLen = intAry.intAryLen

		iAry2.precision = intAry.precision

		iAry2.signVal = intAry.signVal

		return iAry2, nil
	}

	// The number has non-zero digits to
	// the right of the decimal place and
	// the number sign is minus (- or ia.signVal = -1)

	t := make([]uint8, intAry.intAryLen+1)

	n1 := uint8(0)

	n2 := uint8(0)

	carry := uint8(0)

	for i := intIdx; i >= 0; i-- {

		n1 = intAry.intAry[i]

		if i == intIdx {

			if n1+1 > 9 {

				n2 = n1 + 1 - 10

				carry = 1

			} else {

				n2 = n1 + 1

				carry = 0

			}

		} else {

			if n1+carry > 9 {

				n2 = n1 + carry - 10

				carry = 1

			} else {

				n2 = n1 + carry

				carry = 0
			}
		}

		t[i+1] = n2

	}

	if carry != 0 {

		t[0] = carry

		iAry2.intAry = t[0 : intAry.intAryLen+1]

	} else {

		iAry2.intAry = t[1 : intAry.intAryLen+1]
	}

	iAry2.precision = intAry.precision

	iAry2.signVal = intAry.signVal

	iAry2.intAryLen = len(iAry2.intAry)

	//iAry2.SetIsZeroValue()
	err = new(intAryNanobot).setInternalFlags(
		&iAry2, ePrefix.XCpy("Setting 'iAry2' Flags"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
					"  &iAry2, ePrefix.XCpy(Setting 'iAry2' Flags)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return iAry2, nil
}

// getBigInt
//
//	Returns the current value of this intAry object expressed
//	as a signed integer number of type *big.Int.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getBigInt(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getBigInt()",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if intAry == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return big.NewInt(0), err
	}

	lenIntAry := len(intAry.intAry)

	if lenIntAry != intAry.intAryLen {

		return big.NewInt(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: " if lenIntAry != ia.intAryLen",
				ErrMessage: "Error: The actual length of 'intAry' does not match ia.intAryLen.\n" +
					"This instance of 'ia' is INVALID!",
			}

	}

	if lenIntAry == 0 {

		return big.NewInt(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: " if lenIntAry != ia.intAryLen",
				ErrMessage: "Error: The actual length of 'intAry' is ZERO.\n" +
					"This instance of 'ia' is INVALID!",
			}
	}

	result := big.NewInt(0).SetInt64(0)

	big10 := big.NewInt(0).SetInt64(10)

	for i := 0; i < intAry.intAryLen; i++ {
		result = big.NewInt(0).Mul(result, big10)
		result = big.NewInt(0).Add(result, big.NewInt(0).SetInt64(int64(intAry.intAry[i])))

	}

	if intAry.signVal == -1 {

		result = big.NewInt(0).Neg(result)
	}

	return result, nil
}

// getAbsoluteValue
//
//	Returns an IntAry object which represents the Absolute Value of
//	input parameter intAry.
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
//	The IntAry object returned by this method will be configured
//	with Numeric Separators copied from the original input
//	parameter, 'intAry'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getAbsoluteValue(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getAbsoluteValue",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if intAry == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	absIa := new(intAryElectron).newIntAry()

	// intAry numSeps -> absIa
	err = new(intAryProton).copy(&absIa, intAry, false, true, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryProton).copy(&absIa, ia, true, true, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(intAryMechanics).setAbsoluteValue(&absIa, ePrefix.XCpy("Set 'absIa' Absolute Value"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryMechanics).setAbsoluteValue(\n" +
					"  &absIa, ePrefix.XCpy(Set 'absIa' Absolute Value)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return absIa, nil
}

// getFractionalDigits
//
//	Examines the current IntAry instanace and returns a new IntAry
//	object consisting of the fractional digits to the right of the
//	decimal point from the original, current IntAry object.
//
//	Note: The sign Value of the returned int Ary is always
//	positive or +1.
//
//	The returned IntAry instance will display fractional digits
//	with a leading integer digit of zero. Example '0.5678'
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
//	The IntAry object returned by this method will be configured
//	with Numeric Separators copied from the original input
//	parameter, 'intAry'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getFractionalDigits(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getFractionalDigits()",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if intAry == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, true, ePrefix.XCpy("Numeric Separators intAry -> numSeps"))

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := new(intAryElectron).newIntAry()

	nsProfile := NumSepsProfileSelection{
		SourceObjectName:         "intAry",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	err = new(intAryQuark).setIntAryToZero(
		&iAry2,
		nil,
		nsProfile,
		0,
		ePrefix.XCpy("Setting iAry2 to Zero"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  &iAry2, nil, nsProfile-numSeps, 0,\n" +
					"  ePrefix.XCpy(Setting iAry2 to Zero))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if intAry.precision == 0 {
		return iAry2, nil
	}

	fracIdx := intAry.intAryLen - intAry.precision

	iAry2.intAry = make([]uint8, intAry.precision+1)

	idx := 1

	for i := fracIdx; i < intAry.intAryLen; i++ {

		iAry2.intAry[idx] = intAry.intAry[i]

		idx++
	}

	iAry2.precision = intAry.precision

	iAry2.signVal = 1

	err = new(intAryNanobot).setInternalFlags(
		&iAry2, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(&iAry2, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return iAry2, nil
}

// getIntegerDigits
//
//	Examines the current intAry object and returns a new intAry
//	consisting of only the integer digits to the left of the
//	decimal point in the current intAry object.
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
//	The IntAry object returned by this method will be configured
//	with Numeric Separators copied from the original input
//	parameter, 'intAry'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getIntegerDigits(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getIntegerDigits()",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if intAry == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, true, ePrefix.XCpy("Numeric Separators intAry -> numSeps"))

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := new(intAryElectron).newIntAry()

	nsProfile := NumSepsProfileSelection{
		SourceObjectName:         "intAry",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	// Copies intAry numSeps -> iAry2
	err = new(intAryQuark).setIntAryToZero(
		&iAry2, nil, nsProfile, 0, ePrefix.XCpy("Set iAry2 to Zero"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  &iAry2, nil, nsProfile-numSeps , 0, ePrefix.XCpy(Set iAry2 to Zero))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if intAry.isZeroValue {

		return iAry2, nil
	}

	intLen := intAry.intAryLen - intAry.precision

	iAry2.intAry = make([]uint8, intLen)

	for i := 0; i < intLen; i++ {

		iAry2.intAry[i] = intAry.intAry[i]
	}

	iAry2.signVal = intAry.signVal

	iAry2.precision = 0

	err = new(intAryNanobot).setInternalFlags(
		&iAry2, ePrefix.XCpy("Setting Flags on 'iAry2"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
					"  &iAry2, ePrefix.XCpy(Setting Flags on 'iAry2))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if iAry2.isZeroValue {
		iAry2.signVal = 1
	}

	return iAry2, nil
}

// getMagnitude
//
//	Returns the magnitude of the integer portion of the current
//	IntAry numeric value. The integer portion of the number is
//	represented by the digits to the left of the decimal point.
//
//	Magnitude is defined here as the power of 10 which generates a
//	value less than or equal to the integer portion of the current
//	IntAry numeric value.
//
//	   10^magnitude  <= IntAry value
//
//	Note
//	====
//
//	If the current IntAry value is negative, an error will be generated.
func (iaNeutron *intAryNeutron) getMagnitude(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getMagnitude",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return 0, err
	}

	iaNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix)

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaNumStr, err := new(intAryAtom).getRawNumStr(ia, false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if intAry.signVal == -1 {

		return -1,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "ia.signVal == -1",
				ErrMessage: "Error: current IntAry value is negative!\n" +
					fmt.Sprintf("value= '%v'", iaNumStr),
			}
	}

	return intAry.intAryLen - intAry.precision - intAry.firstDigitIdx - 1, nil
}

// getNumStrDto
//
//	Converts the IntAry input parameter ('intAry') to a returned
//	instance of NumStrDto.
//
//	The returned NumStrDto will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from the current IntAry instance.
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
//	The NumStrDto object returned by this method will be configured
//	with Numeric Separators copied from the original input
//	parameter, 'intAry'.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaNeutron *intAryNeutron) getNumStrDto(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getNumStrDto()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if intAry == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSeps, err := new(intAryPhoton).
		getNumericSeparatorsDto(intAry, true, ePrefix.XCpy("numSeps<-intAry"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := new(intAryNanobot).getNumericSeparatorsDto(intAry, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iaNumStr, err := new(intAryAtom).getRawNumStr(intAry, false, ePrefix.XCpy("iaNumStr<-intAry"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "iaNumStr, err := new(intAryAtom).getRawNumStr(\n" +
					"  intAry, false, ePrefix.XCpy(iaNumStr<-intAry))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nDto, err := new(NumStrDto).NewNumStrWithNumSeps(iaNumStr, numSeps)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nDto, err := new(NumStrDto).NewNumStrWithNumSeps(iaNumStr, numSeps)",
				ErrContext: fmt.Sprintf("iaNumStr= '%v'\nnumSeps= '%v'", iaNumStr, numSeps),
				ErrMessage: err.Error(),
			}
	}

	return nDto, nil
}

// Multiply
//
//	This method receives three IntAry input parameters, 'ia1',
//	'ia2' and 'iaResult'.It then proceeds to multiply 'ia1' by
//	'ia2' and stores the multiplication result in intAry
//	'iaResult'.
//
//	Example
//	=======
//
//	      product = multiplicand x multipllier
//
//	      iaResult = ia1 x ia2
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
//	Numeric Separators for the calculated multiplication result
//	('iaResult') will be copied from current IntAry instance
//	('ia').
//
//	Input Parameters
//	================
//
//	ia1                      *IntAry
//	  In this multiplication operation, 'ia1' is the multiplicand.
//
//	validateIa1              bool
//	   When set to 'true', input parameter 'ia' will be subjected
//	   to Validation Tests.
//
//	ia2                      *IntAry
//	  In this multiplication operation, 'ia2' is the multiplier.
//
//	validateIa2              bool
//	   When set to 'true', input parameter 'ia2' will be subjected
//	   to Validation Tests.
//
//	iaResult                 *IntAry
//
//	  This 'iaResult' IntArray object which will store the result
//	  of the multiplication operation. 'iaResult' is the 'product'.
//
//	validateIaResult         bool
//	   When set to 'true', input parameter 'iaResult' will be
//	   subjected to Validation Tests.
//
//	minimumResultPrecision   int
//	  'minimumResultPrecision' will determine the minimum number of
//	  digits computed to the right of the decimal place in the
//	  final result or 'product'.
//
//	  If 'minimumResultPrecision' is set to a value of -1, all
//	  significant digits (digits greater than zero) will be
//	  returned to the right of the decimal place. Remember that
//	  the maximum number of decimal digits returned will be
//	  controlled by parameter 'maxResultPrecision'
//
//	maxResultPrecision       int
//	  'maxResultPrecision' will determine the maximum number of
//	  digits to the right of the decimal place in the result.
//
//	  Valid values are -1 and values >= zero ('0')
//
//	  Values less than -1 will trigger an error.
//
//	  A value of -1 signals that no limit will be placed on the
//	  number of decimals places to right of the decimal point in
//	  the result. Be advised that a very, very large number of
//	  decimal digits may be accommodated by the IntAry Type.
//
//	Return Value
//	============
//
//	error
//	  If no errors are encountered, this method will return an
//	  error value of 'nil'.
func (iaNeutron *intAryNeutron) multiply(
	ia1 *IntAry,
	validateIa1 bool,
	ia2 *IntAry,
	validateIa2 bool,
	iaResult *IntAry,
	validateIaResult bool,
	minimumResultPrecision int,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.multiply",
		"")

	if err != nil {
		return err
	}

	if ia1 == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia1'",
		}
	}

	if ia2 == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia2'",
		}
	}

	if iaResult == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia2'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia1,
		"ia1",
		validateIa1,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia2,
		"ia2",
		validateIa2,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(IntAryMathMultiply).Multiply(ia1, ia2, iaResult, minimumResultPrecision, maxResultPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathMultiply).Multiply(\n" +
				"  ia1, ia2, iaResult, minimumResultPrecision, maxResultPrecision)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		iaResult,
		"iaResult",
		validateIaResult,
		ePrefix)

	if err != nil {
		return err
	}

	return nil
}

// multiplyByTenToPower
//
//	The value of intAry is multiplied by 10 to the power of the
//	input parameter 'power'.
//
//	Example
//	=======
//
//	IntAry Numeric Value x 10^power = result
func (iaNeutron *intAryNeutron) multiplyByTenToPower(
	intAry *IntAry,
	validateIntAry bool,
	power uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.multiplyByTenToPower",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(IntAryMathMultiply).MultiplyByTenToPower(intAry, power)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathMultiply).\n" +
				"  MultiplyByTenToPower(intAry, power)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// multiplyByTwoToPower
//
//		Multiply the existing value of the IntAry by 2 to the power
//		of the input parameter 'power'.
//
//	 Example
//	 =======
//
//	 IntAry Numeric Value x 2^power = result
func (iaNeutron *intAryNeutron) multiplyByTwoToPower(
	intAry *IntAry,
	validateIntAry bool,
	power uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.multiplyByTwoToPower",
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

	err = new(intAryUtility).selectIntAryValidation(
		intAry,
		"intAry",
		validateIntAry,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(IntAryMathMultiply).MultiplyByTwoToPower(intAry, power)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathMultiply).\n" +
				"  MultiplyByTwoToPower(intAry, power)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// multiplyThisBy
//
//	Multiplies the IntAry parameter 'ia' by 'ia2' and
//	stores the multiplication result in 'ia'.
//
//	Example
//	=======
//
//	  ia = ia x ia2
//
//	Input Parameters
//	================
//
//	ia                       *ia
//	  Pointer to an IntAry object. As shown above, this IntAry
//	  object will store the final multiplication result.
//
//	validateIa               bool
//	   When set to 'true', input parameter 'ia' will be subjected
//	   to Validation Tests.
//
//	ia2                      *IntAry
//	  Pointer to an IntAry object. In this multiplication
//	  operation, 'ia2' is the multiplier.
//
//	validateIa2              bool
//	   When set to 'true', input parameter 'ia2' will be subjected
//	   to Validation Tests.
//
//	minimumResultPrecision   int
//	  'minimumResultPrecision' will determine the minimum number of
//	  digits computed to the right of the decimal place in the
//	  final result.
//
//	  If 'minimumResultPrecision' is set to a value of -1, all
//	  significant digits (digits greater than zero) will be
//	  returned to the right of the decimal place. Remember that
//	  the maximum number of decimal digits returned will be
//	  controlled by parameter 'maxResultPrecision'
//
//	maxResultPrecision       int
//	  'maxResultPrecision' will determine the maximum number of
//	  digits to the right of the decimal place in the result.
//
//	  Valid values are -1 and values >= zero ('0')
//
//	  Values less than -1 will trigger an error.
//
//	  A value of -1 signals that no limit will be placed on the
//	  number of decimals places to right of the decimal point in
//	  the result. Be advised that a very, very large number of
//	  decimal digits may be accommodated by the IntAry Type.
//
//	Return Value
//	============
//
//	error
//	  If no errors are encountered, this method will return an
//	  error value of 'nil'.
func (iaNeutron *intAryNeutron) multiplyThisBy(
	ia *IntAry,
	validateIa bool,
	ia2 *IntAry,
	validateIa2 bool,
	minimumPrecision,
	maxPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.multiplyThisBy",
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

	if ia2 == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia2'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia2,
		"ia2",
		validateIa2,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(IntAryMathMultiply).
		Multiply(ia, ia2, ia, minimumPrecision, maxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(IntAryMathMultiply).\n" +
				"  Multiply(ia, ia2, ia, minimumPrecision, maxPrecision)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// SetIntAryWithIntFracStr
//
//	Sets the value of the current IntAry instance based on a
//	numeric value represented by separate integer and fractional
//	components.
//
//	Input parameters 'intStr' and 'fracStr' are strings
//	representing the integer and fractional components. They are
//	combined by this method to create a numeric value which is
//	assigned to, and stored in, the current IntAry instance.
//
//	Input parameter 'signVal' must be set to one of two values:
//	+1 or -1. This value is used to signal the sign of the
//	resulting numeric value. +1 generates a positive number and -1
//	generates a negative number. If input parameters 'inStr' or
//	'fracStr' contain a leading minus or plus sign character, it
//	will be ignored. The sign of the resulting numeric value is
//	controlled strictly by input parameter, 'signVal'.
//
//	IMPORTANT
//	=========
//
//	This method will use the numeric separators in the current
//	instance of IntAry to convert the integer and fractional
//	components into a consolidated number string for internal
//	calculation purposes.
func (iaNeutron *intAryNeutron) setIntAryWithIntFracStr(
	ia *IntAry,
	validateIa bool,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	intStr string,
	fracStr string,
	signVal int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.SetIntAryWithIntFracStr()",
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

	if len(intStr) == 0 && len(fracStr) == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameters 'intStr' and 'fracStr' empty strings!",
		}

	}

	if len(intStr) == 0 && len(fracStr) != 0 {
		intStr = "0"
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix)

	if err != nil {
		return err
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
			ReturnFunc: "finalNumSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\".String())",
			ErrContext: fmt.Sprintf("Generated 'numSeps' is INVALILD!\n"+
				"numSeps.DecimalSeparator= '%c'\n"+
				"numSeps.ThousandsSeparator= '%c'\n"+
				"numSeps.CurrencySymbol= '%c'",
				numSeps.DecimalSeparator, numSeps.ThousandsSeparator, numSeps.CurrencySymbol),
			ErrMessage: err.Error(),
		}
	}

	localDecimalSeparator := numSeps.DecimalSeparator

	if localDecimalSeparator == 0 {
		localDecimalSeparator = '.'
	}

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: false,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	cleanIntRuneAry := make([]rune, 0, 100)

	zeroChar := uint8('0')
	nineChar := uint8('9')

	lStr := len(intStr)

	if lStr == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'intStr' is zero Length!\n",
			ePrefix)
	}

	isFirstRune := true

	// Create pure number string from 'intStr'
	for i := 0; i < lStr; i++ {

		if intStr[i] >= zeroChar &&
			intStr[i] <= nineChar {

			if isFirstRune && signVal == -1 {

				cleanIntRuneAry = append(cleanIntRuneAry, '-')
			}

			isFirstRune = false

			cleanIntRuneAry = append(cleanIntRuneAry, rune(intStr[i]))
		}
	}

	if len(cleanIntRuneAry) == 0 {

		cleanIntRuneAry = append(cleanIntRuneAry, '0')
	}

	lStr = len(fracStr)

	if lStr > 0 {

		isFirstRune = true

		for j := 0; j < lStr; j++ {

			if fracStr[j] >= zeroChar &&
				fracStr[j] <= nineChar {

				if isFirstRune {
					cleanIntRuneAry = append(cleanIntRuneAry, localDecimalSeparator)
					isFirstRune = false
				}

				cleanIntRuneAry = append(cleanIntRuneAry, rune(fracStr[j]))
			}

		}
	}

	err = new(intAryQuark).setIntAryWithNumStr(
		ia,
		validateIa,
		numSepsSrcIntAry,
		nsProfile2,
		string(cleanIntRuneAry),
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryWithNumStr(\n"+
				"ia, validateIa= '%v', numSepsSrcIntAry, nsProfile2, numStr, validateResult= '%v', ePrefix)",
				validateIa, validateResult),
			ErrContext: fmt.Sprintf("numStr= %v", string(cleanIntRuneAry)),
			ErrMessage: err.Error(),
		}

	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryUtility).selectIntAryValidation(ia,\"ia\", validateResult, ePrefix)",
			ErrContext: "The final calculation result of intAryNeutron.setIntAryWithIntFracStr()\n" +
				"FAILED Validation Tests",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryWithNumStrMaxPrecision
//
//	Receives a raw number string and sets the fields of the
//	internal intAry structure to the appropriate values.
//
//	A second input parameter specifies the maximum allowable
//	precision for the IntAry. If IntAry precision exceeds
//	'maxPrecision', IntAry precision is rounded to 'maxPrecision'.
//
//	The term 'precision' defines the number of numeric digits to
//	the right of the decimal point or decimal separator.
func (iaNeutron *intAryNeutron) setIntAryWithNumStrMaxPrecision(
	ia *IntAry,
	validateIa bool,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	numStr string,
	maxPrecision int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.setIntAryWithNumStrMaxPrecision()",
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
		validateIa,
		ePrefix)

	if err != nil {
		return err
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
			ReturnFunc: "finalNumSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: false,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	err = new(intAryQuark).setIntAryWithNumStr(
		ia,
		validateIa,
		numSepsSrcIntAry,
		nsProfile2,
		numStr,
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryWithNumStr(\n"+
				"ia, validateIa= '%v', numSepsSrcIntAry, nsProfile2, numStr, validateResult= '%v', ePrefix)",
				validateIa, validateResult),
			ErrContext: "",
			ErrMessage: err.Error(),
		}

	}

	if ia.precision > maxPrecision {

		err = new(intAryMolecule).roundToPrecision(
			ia,
			true,
			maxPrecision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: fmt.Sprintf("err = new(intAryMolecule).roundToPrecision(\n"+
					"ia, validateIa='true',\n"+
					"maxPrecision= '%v', ePrefix",
					maxPrecision),
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia, 'ia', validateResult='%v', ePrefix",
				validateResult),
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setSign
//
//	Used to change the sign value of the intAry  object passed as
//	input parameter 'intAry'. The new sign value will be set
//	according the input parameter, 'signVal'.
//
//	'signVal' has only two valid values, -1 or +1. If
//	any value other than -1 or +1 is detected, an error
//	will be returned.
func (iaNeutron *intAryNeutron) setSign(
	intAry *IntAry,
	signVal int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.equal()",
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

	if signVal != -1 && signVal != 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("'signVal' == %d", signVal),
			ErrMessage: "Error: Input parameter 'signVal' is INVALID.\n" +
				"'signVal' must be either -1 or +1.",
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

	if intAry.isZeroValue {
		return nil
	}

	intAry.signVal = signVal

	return nil
}
