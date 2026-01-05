package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type numStrDtoGluon struct {
	lock sync.Mutex
}

// getAbsAllNumRunes
//
//	Returns an array of runes representing all the integer and
//	fractional digits included in the current NumStrDto instance.
//	The rune array returned will consist of numeric digits with no
//	sign value prefixed. This effectively returns the absolute
//	value of all integer and fractional digits combined in one rune
//	array (there is no decimal point).
func (nStrDtoGluon *numStrDtoGluon) getAbsAllNumRunes(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) ([]rune, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsAllNumRunes()",
		"")

	if err != nil {
		return []rune{}, err
	}

	if numStrDto == nil {

		return []rune{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return []rune{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if lenAbsAllNumRunes == 0 {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "len(numStrDto.absAllNumRunes) == 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The internal rune array of numeric characters is empty.",
				}
		}

		precision := int(numStrDto.precision)

		if precision < 0 {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision < 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is less than zero.",
				}
		}

		if precision > lenAbsAllNumRunes {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is greater than the internal numeric digits array.",
					ErrMessage: "",
				}
		}
	}

	outRunes := make([]rune, lenAbsAllNumRunes)

	for i := 0; i < lenAbsAllNumRunes; i++ {
		outRunes[i] = numStrDto.absAllNumRunes[i]
	}

	return outRunes, nil
}

// getAbsoluteBigInt
//
//	Returns the absolute value of all numeric digits in the number
//	string (nDto.absAllNumRunes) extracted from input parameter
//	'nDto'. As such, Fractional digits to the right of the decimal
//	are included in the consolidate integer	number.
//
//	IMPORTANT
//	=========
//
//	If the NumStrDto instance 'nDto' is invalid, an error will be
//	returned.
//
//	All the numeric digits in the number string are
//	therefore returned as a type *big.Int numeric value.
//
//	Input Parameter
//	===============
//
//	nDto                     *NumStrDto
//	  The returned Big Int Number (*big.Int) contains the numeric
//	  value extracted from this instance of NumStrDto.
//
//	validateNumStrDto        bool
//	  When set to 'true', the NumStrDto parameter 'nDto' will be
//	  subjected to validation tests.
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
//	*big.Int
//	  A positive (+) big int number representing the absolute value
//	  of both the integer and  fractional digits extracted from the
//	  current NumStrDto instance, 'nDto'.
//
//	uint
//	  An unsigned integer value representing the number of
//	  fractional digits to the right of the decimal point in the
//	  numeric value returned by '*big.Int'.
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nStrDtoGluon *numStrDtoGluon) getAbsoluteBigInt(
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, uint, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsoluteBigInt()",
		"")

	if err != nil {
		return big.NewInt(0), 0, err
	}

	if nDto == nil {

		return big.NewInt(0), 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return big.NewInt(0), 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	lenAllNumRunes := len(nDto.absAllNumRunes)

	if lenAllNumRunes == 0 {

		return big.NewInt(0), 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "len(nDto.absAllNumRunes) == 0",
				ErrMessage: "Error: Input parameter 'nDto' (NumStrDto) is INVALID!\n" +
					"'nDto' contains a zero length internal runes array.",
			}

	}

	base10 := big.NewInt(int64(10))

	absBigInt := big.NewInt(0)

	for i := 0; i < lenAllNumRunes; i++ {

		absBigInt = big.NewInt(0).Mul(absBigInt, base10)

		absBigInt = big.NewInt(0).Add(absBigInt,
			big.NewInt(int64(nDto.absAllNumRunes[i]-48)))

	}

	return absBigInt, nDto.precision, nil
}

// GetAbsIntRunesLength
//
//	Returns the length of the integer portion of the number string.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Examples
//	========
//
//	Numeric Value    Returned Length
//
//	  -123456              6
//	   123456              6
//	   123.456             6
//	   0.123456            6
func (nStrDtoGluon *numStrDtoGluon) getAbsIntRunesLength(
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var err error
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsPureNumStr()",
		"")

	if err != nil {
		return 0, err
	}

	if nDto == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	lenAllNums := len(nDto.absAllNumRunes)

	return lenAllNums - int(nDto.precision), nil
}

// getAbsPureNumStr
//
//	Returns all digits in the current NumStrDto numeric value as a
//	pure, unsigned number string. If fractional digits exists, they
//	are included in the string and NOT separated by a decimal
//	separator.
//
//	All the numeric digits in the numeric value are returned as a
//	type string.
//
//	Examples
//	========
//
//	Numeric
//	Value       pureNumStr    precision    numberSign
//	------      ----------    ---------    ----------
//
//	 123.45      12345            2             1
//	 12345       12345            0             1
//	-123.45      11245            2            -1
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	nDto                     *NumStrDto
//	  The returned number string will be extracted from the numeric
//	  value contained in this instance of NumStrDto
//
//	validateNumStrDto        bool
//	  When set to 'true', the NumStrDto parameter 'nDto' will be
//	  subjected to validation tests.
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
//	pureNumStr               string
//	  This returned string value contains the pure number string
//	  representaion of the absolue numeric value for the current
//	  NumStrDto instance.
//
//	precision                 uint
//	  An unsigned integer value representing the number of
//	  fractional digits to the right of the decimal point in the
//	  numeric value returned by 'pureNumStr'.
//
//	numberSign                int
//	  The returned number sign will designate the 'purNumStr' as
//	  either a positive value or a negative value.
//
//	  Possible return values:
//	    +1 = Positive numeric value (including zero)
//	                OR
//	    -1 = Negative numeric value
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nStrDtoGluon *numStrDtoGluon) getAbsPureNumStr(
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (
	pureNumStr string, precision uint, numberSign int, err error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsPureNumStr()",
		"")

	if err != nil {
		return pureNumStr, precision, numberSign, err
	}

	if nDto == nil {

		return pureNumStr, precision, numberSign,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return pureNumStr, precision, numberSign,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	pureNumStr = string(nDto.absAllNumRunes)

	precision = nDto.precision

	numberSign = nDto.signVal

	return pureNumStr, precision, numberSign, err
}

// getAbsFracRunes
//
//	Returns all the fractional digits to the right of the decimal
//	place in the current NumStrDto instance as an array of runes.
//	The rune array is not signed; that is, the rune array does not
//	contain a '+' or '-' character in the first array position. The
//	rune array is therefore said to represent the absolute value of
//	the fractional digits in the current NumStrDto numeric value.
func (nStrDtoGluon *numStrDtoGluon) getAbsFracRunes(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) ([]rune, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsFracRunes()",
		"")

	if err != nil {
		return []rune{}, err
	}

	if numStrDto == nil {

		return []rune{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	precision := int(numStrDto.precision)

	lenAllNums := len(numStrDto.absAllNumRunes)

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return []rune{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if lenAllNums == 0 {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "len(numStrDto.absAllNumRunes) == 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The internal rune array of numeric characters is empty.",
				}
		}

		if precision < 0 {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision < 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is less than zero.",
				}
		}

		if precision > lenAllNums {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is greater than the internal numeric digits array.",
					ErrMessage: "",
				}
		}

	}

	absFracRunes := make([]rune, precision)

	lenIntNums := lenAllNums - precision

	for i := lenIntNums; i < lenAllNums; i++ {
		absFracRunes[i-lenIntNums] = numStrDto.absAllNumRunes[i]
	}

	return absFracRunes, nil
}

// getAbsFracRunesLength
//
//	Returns the length of the fractional digits in the number
//	string.
func (nStrDtoGluon *numStrDtoGluon) getAbsFracRunesLength(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsIntRunes()",
		"")

	if err != nil {
		return 0, err
	}

	if numStrDto == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	lenAllNums := len(numStrDto.absAllNumRunes)

	precision := int(numStrDto.precision)

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return 0,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if lenAllNums == 0 {

			return 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "len(numStrDto.absAllNumRunes) == 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The internal rune array of numeric characters is empty.",
				}
		}

		if precision < 0 {

			return 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision < 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is less than zero.",
				}
		}

		if precision > lenAllNums {

			return 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is greater than the internal numeric digits array.",
					ErrMessage: "",
				}
		}

	}

	return precision, nil
}

// getAbsIntRunes
//
//	Returns all the integer digits included in the current
//	NumStrDto numeric value as an array of runes. The returned rune
//	array does not contain a sign value in the first position and
//	therefore represents the absolute or positive value of all the
//	integer digits. The integer digits of a NumStrDto numeric
//	includes all the digits to the left of the decimal point.
//
//	If the current NumStrDto consists of zero integers and
//	fractional digits (Example: '0.123456'), this method will
//	return a rune array consisting one array element with a '0'
//	value.
func (nStrDtoGluon *numStrDtoGluon) getAbsIntRunes(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) ([]rune, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getAbsIntRunes()",
		"")

	if err != nil {
		return []rune{}, err
	}

	if numStrDto == nil {

		return []rune{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	lenAllNums := len(numStrDto.absAllNumRunes)

	precision := int(numStrDto.precision)

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return []rune{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "Error: Input parameter 'numStrDto' is INVALID!",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if lenAllNums == 0 {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "len(numStrDto.absAllNumRunes) == 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The internal rune array of numeric characters is empty.",
				}
		}

		if precision < 0 {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision < 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is less than zero.",
				}
		}

		if precision > lenAllNums {

			return []rune{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision > lenAllNums",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is greater than the internal numeric digits array.",
				}
		}

	}

	lenIntNum := lenAllNums - precision

	absIntRunes := make([]rune, lenIntNum)

	for i := 0; i < lenIntNum; i++ {
		absIntRunes[i] = numStrDto.absAllNumRunes[i]
	}

	return absIntRunes, nil
}

// getRationalNumber
//
//	Returns the sign value of the number string, plus the numeric
//	value of the number string expressed as a Rational Number.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameter
//	===============
//
//	nDto                     *NumStrDto
//	  The returned rational number contains the numeric value
//	  extracted from this instance of NumStrDto.
//
//	validateNumStrDto        bool
//	  When set to 'true', the NumStrDto parameter 'nDto' will be
//	  subjected to validation tests.
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
//	signValue                int
//	  Sign value of the current instance NumStrDto numeric value.
//	  This returned value either be a +1 or a -1.
//
//	    Possible Sign Values
//	      1 = NumStrDto numeric value is a Positive Number
//	     -1 = NumStrDto numeric value is a Negative Number
//
//	bigRat                   *big.Rat
//	  NumStrDto numeric value expressed as a rational number.
//
//	err                      error
//	  If an error is encountered during processing, this error
//	  object will be returned configured with an appropriate error
//	  message.
func (nStrDtoGluon *numStrDtoGluon) getRationalNumber(
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (signValue int, bigRat *big.Rat, err error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	bigRat = big.NewRat(1, 1)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getRationalNumber()",
		"")

	if err != nil {
		return signValue, bigRat, err
	}

	if nDto == nil {

		return signValue, bigRat,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return signValue, bigRat,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	signVal := nDto.signVal

	absInt, isOk := big.NewInt(0).SetString(string(nDto.absAllNumRunes), 10)

	if !isOk {

		return signValue, bigRat,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "absInt, isOk := big.NewInt(0).SetString(string(nDto.absAllNumRunes), 10)",
				ErrContext: fmt.Sprintf("nDto.absAllNumRunes= '%v'",
					string(nDto.absAllNumRunes)),
				ErrMessage: "Error: SetString Failed to convert NumStrDto runes to *big.Int!",
			}
	}

	base10 := big.NewInt(10)

	bigPrecision := big.NewInt(int64(nDto.precision))

	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

	rationalNum := big.NewRat(1, 1).SetFrac(absInt, scaleFactor)

	return signVal, rationalNum, nil
}

// getScaleFactor
//
//	Returns the scale factor for the number string encapsulated by
//	the current instance of NumStrDto.
//
//	Scale factor is defined by 10 raised to the power of
//	'precision' (nDto.precision).  nDto.precision is the number of
//	digits to the right of the decimal point.
//
//	This method will fail and return an error if the current instance
//	of NumStrDto is invalid.
func (nStrDtoGluon *numStrDtoGluon) getScaleFactor(
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getScaleFactor()",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if nDto == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return big.NewInt(0),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if nDto.precision == 0 {
		return big.NewInt(int64(1)), nil
	}

	base10 := big.NewInt(0).SetInt64(int64(10))

	bigPrecision := big.NewInt(0).SetInt64(int64(nDto.precision))

	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

	return scaleFactor, nil
}

// getSciNotationNumber
//
//	Converts the numeric value of the current NumStrDto instance
//	into scientific notation and returns this value as an instance
//	of type SciNotationNum.
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
//	  The returned SciNotationNum instance contains the numeric
//	  value extracted from this instance of NumStrDto.
//
//	validateNumStrDto        bool
//	  When set to 'true', the NumStrDto parameter 'nDto' will be
//	  subjected to validation tests.
//
//	mantissaLen              uint
//	  Specifies the length of the mantissa in the returned
//	  scientific notation string. If the value of 'mantissaLen' is
//	  less than two ('2'), this method will automatically set the
//	  'mantissaLen' to a default value of two ('2').
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
//	SciNotationNum
//	  This returned SciNotationNum instance contains the numeric
//	  representaion of the NumStrDto instance 'nDto'.
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nStrDtoGluon *numStrDtoGluon) getSciNotationNumber(
	nDto *NumStrDto,
	validateNumStrDto bool,
	mantissaLen uint,
	errPrefDto *ePref.ErrPrefixDto) (SciNotationNum, error) {

	nStrDtoGluon.lock.Lock()

	defer nStrDtoGluon.lock.Unlock()

	var err error
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoGluon.getSciNotationNumber()",
		"")

	if err != nil {
		return SciNotationNum{}, err
	}

	if nDto == nil {

		return SciNotationNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
		nDto, ePrefix.XCpy("nDto -> numSeps"))

	if err != nil {
		return SciNotationNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
					"  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
				ErrContext: "Error extracting Numeric Separators from from NumStrDto ('nDto').",
				ErrMessage: err.Error(),
			}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return SciNotationNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: The NumStrDto instance ('nDto') is INVALID!\n" +
					"Numeric Separators from 'nDto' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	//bINum, err := nDto.GetBigIntNum()
	bINum, err := new(numStrDtoMechanics).getBigIntNumWithNumStr(
		numSeps, nDto, false, ePrefix)

	if err != nil {

		return SciNotationNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINum, err := new(numStrDtoMechanics).getBigIntNumWithNumStr(\n" +
					"  numSeps, nDto, false, ePrefix)",
				ErrContext: "Error extracting BigIntNum from NumStrDto ('nDto').",
				ErrMessage: err.Error(),
			}
	}

	if mantissaLen < 2 {
		mantissaLen = 2
	}

	sciNotation, err := bINum.GetSciNotationNumber(mantissaLen)

	if err != nil {

		return SciNotationNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "sciNotation, err := bINum.\n" +
					"  GetSciNotationNumber(mantissaLen)",
				ErrContext: "Error converting BigIntNum ('bINum') to Scientific Notation.",
				ErrMessage: err.Error(),
			}
	}

	return sciNotation, nil
}
