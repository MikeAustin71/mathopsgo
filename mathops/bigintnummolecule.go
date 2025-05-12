package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumMolecule struct {
	lock *sync.Mutex
}

// formatCurrencyStr - Formats the current BigIntNum numeric value as a currency string.
//
// If the Currency Symbol was not previously set for this BigIntNum, the currency symbol
// is defaulted to the USA standard dollar sign, ('$'). To use other currency symbols, see
// method BigIntNum.SetCurrencySymbol(). For a list of Major Currency Unicode Symbols, see
// constants located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// If the Decimal Separator was not previously set for this BigIntNum, the Decimal Separator
// is defaulted to the USA standard period ('.'). To use another character for Decimal
// Separator, see method BigIntNum.SetDecimalSeparator().
//
// If the Thousands Separator was not previously set for this BigIntNum, the Thousands
// Separator is defaulted to the USA standard comma (','). To use another character for
// Thousands Separator, see method BigIntNum.SetThousandsSeparator().
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//
//	LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -$123,456.78
//
//	PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//															surrounding parentheses.
//															Example: ($123,456.78)
//
//
//	ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//															absolute (positive) integer value
//															and no decimal place separator.
//															Example: ($12,345,678)
func (bIntMolecule *bigIntNumMolecule) formatCurrencyStr(
	bNum *BigIntNum,
	negValMode NegativeValueFmtMode,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.formatCurrencyStr()",
		"")

	if err != nil {
		return "", err
	}

	if bNum == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing 'bNum'"))

	if err != nil {

		return "", err
	}

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	if bNum.currencySymbol == 0 {
		bNum.currencySymbol = '$'
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)

	if scratchNum.Cmp(baseZero) == 0 {
		bNum.sign = 1

		outRunes = append(outRunes, bNum.currencySymbol)

		outRunes = append(outRunes, '0')

		if bNum.precision > 0 {

			if negValMode != ABSOLUTEPURENUMSTRFMTMODE {
				outRunes = append(outRunes, bNum.decimalSeparator)
			}

			cnt := int(bNum.precision)

			if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
				cnt--
			}

			for h := 0; h < cnt; h++ {
				outRunes = append(outRunes, '0')
			}

		}

		return string(outRunes), nil
	}

	startIdx := 0
	modulo := big.NewInt(0)
	baseTen := big.NewInt(10)
	digitCnt := 0
	thouCnt := -1

	if bNum.precision == 0 {
		thouCnt = 0
	}

	for scratchNum.Cmp(baseZero) == 1 {

		if startIdx == 0 &&
			bNum.sign == -1 &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64()+int64(48)))
		digitCnt++
		startIdx++

		if thouCnt > -1 {
			thouCnt++
		}

		if scratchNum.Cmp(baseZero) == 1 &&
			thouCnt == 3 {

			outRunes = append(outRunes, bNum.thousandsSeparator)
			startIdx++
			thouCnt = 0
		}

		if bNum.precision > 0 &&
			int(bNum.precision) == startIdx &&
			negValMode != ABSOLUTEPURENUMSTRFMTMODE {

			outRunes = append(outRunes, bNum.decimalSeparator)
			startIdx++
			thouCnt = 0
		}

	}

	if int(bNum.precision) >= digitCnt {

		delta := int(bNum.precision) - digitCnt + 1

		if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
			delta--
		}

		for k := 0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bNum.precision > 0 &&
				int(bNum.precision) == startIdx &&
				negValMode != ABSOLUTEPURENUMSTRFMTMODE {

				outRunes = append(outRunes, bNum.decimalSeparator)
				startIdx++

			}
		}
	}

	startIdx--

	// append Currency Symbol
	outRunes = append(outRunes, bNum.currencySymbol)
	startIdx++

	// adjust for negative sign value
	if bNum.sign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		} else if negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, '(')
			startIdx += 2
		}

		// Must be negValMode == ABSOLUTEPURENUMSTRFMTMODE

	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i := startIdx; i > sortLimit; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes), nil
}

// formatBigIntNumStr - Formats the numeric value of the current BigIntNum
// instance as number string consisting of integer digits to the left
// of the decimal place and fractional digits to the right of the decimal
// point, if such fractional digits exist. The resulting number string
// will NOT contain a currency symbol or thousands separators.
//
// If the Decimal Separator was not previously set for this BigIntNum,
// the Decimal Separator is defaulted to the USA standard period ('.').
// To use another character for Decimal Separator, see method
// BigIntNum.SetDecimalSeparator().
//
// Output Examples: 123456.789 or -123456.789
//
//	NOTE:
//
// ================
//
// This method does NOT test the validity of 'bNum' BigIntNum
// instance. The calling method must do this!
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//
//	LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -123456.78
//
//	PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//															surrounding parentheses.
//															Example: (123456.78)
//
//	ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//															absolute (positive) integer value
//															and no decimal place separator.
//															Example: (12345678)
func (bIntMolecule *bigIntNumMolecule) formatBigIntNumStr(
	bNum *BigIntNum,
	negValMode NegativeValueFmtMode,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.formatBigIntNumStr",
		"")

	if err != nil {
		return "", err
	}

	if bNum == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)

	if scratchNum.Cmp(baseZero) == 0 {
		bNum.sign = 1

		outRunes = append(outRunes, '0')

		if bNum.precision > 0 {

			if negValMode != ABSOLUTEPURENUMSTRFMTMODE {
				outRunes = append(outRunes, bNum.decimalSeparator)
			}

			cnt := int(bNum.precision)

			if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
				cnt--
			}

			for h := 0; h < cnt; h++ {
				outRunes = append(outRunes, '0')
			}

		}

		return string(outRunes), nil
	}

	startIdx := 0
	modulo := big.NewInt(0)
	baseTen := big.NewInt(10)
	digitCnt := 0

	for scratchNum.Cmp(baseZero) == 1 {

		if startIdx == 0 &&
			bNum.sign == -1 &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64()+int64(48)))
		digitCnt++
		startIdx++

		if bNum.precision > 0 &&
			int(bNum.precision) == startIdx &&
			negValMode != ABSOLUTEPURENUMSTRFMTMODE {

			outRunes = append(outRunes, bNum.decimalSeparator)
			startIdx++
		}

	}

	if int(bNum.precision) >= digitCnt {

		delta := int(bNum.precision) - digitCnt + 1

		if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
			delta--
		}

		for k := 0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bNum.precision > 0 &&
				int(bNum.precision) == startIdx &&
				negValMode != ABSOLUTEPURENUMSTRFMTMODE {

				outRunes = append(outRunes, bNum.decimalSeparator)
				startIdx++
			}
		}
	}

	startIdx--

	// adjust for negative sign value
	if bNum.sign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		} else if negValMode == PARENTHESESNEGVALFMTMODE {
			outRunes = append(outRunes, '(')
			startIdx += 2
		}

		/*
				MUST BE negValMode == ABSOLUTEPURENUMSTRFMTMODE
			  Do NOT Display Sign Character

		*/
	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i := startIdx; i > sortLimit; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes), nil
}

// formatThousandsStr - Returns the number string delimited with the
// BigIntNum ThousandsSeparator character plus the Decimal Separator
// character if applicable. See methods BigIntNum.SetThousandsSeparator()
// and BigIntNum.SetDecimalSeparator().
//
// If the Decimal Separator was not previously set for this BigIntNum,
// the Decimal Separator is defaulted to the USA standard period ('.').
// To use another character for Decimal Separator, see method
// BigIntNum.SetDecimalSeparator().
//
// If the Thousands Separator was not previously set for this BigIntNum,
// the Thousands Separator is defaulted to the USA standard comma (',').
// To use another character for Thousands Separator, see method
// BigIntNum.SetThousandsSeparator().
//
// Example:
// numStr = 1000000.234 converted to 1,000,000.234
//
//	NOTE:
//
// ================
//
// This method does NOT test the validity of 'bNum' BigIntNum
// instance. The calling method must do this!
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//
//	LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -123,456.78
//
//	PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//															surrounding parentheses.
//															Example: (123,456.78)
//
//
//	ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//															absolute (positive) integer value
//															and no decimal place separator.
//															Example: (12,345,678)
func (bIntMolecule *bigIntNumMolecule) formatThousandsStr(
	bNum *BigIntNum,
	negValMode NegativeValueFmtMode,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.formatThousandsStr",
		"")

	if err != nil {
		return "", err
	}

	if bNum == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing 'bNum'"))

	if err != nil {

		return "", err
	}

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)

	if scratchNum.Cmp(baseZero) == 0 {
		bNum.sign = 1

		outRunes = append(outRunes, '0')

		if bNum.precision > 0 {

			if negValMode != ABSOLUTEPURENUMSTRFMTMODE {
				outRunes = append(outRunes, bNum.decimalSeparator)
			}

			cnt := int(bNum.precision)

			if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
				cnt--
			}

			for h := 0; h < cnt; h++ {
				outRunes = append(outRunes, '0')
			}

		}

		return string(outRunes), nil
	}

	startIdx := 0
	modulo := big.NewInt(0)
	baseTen := big.NewInt(10)
	digitCnt := 0
	thouCnt := -1

	if bNum.precision == 0 {
		thouCnt = 0
	}

	for scratchNum.Cmp(baseZero) == 1 {

		if startIdx == 0 &&
			bNum.sign == -1 &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64()+int64(48)))
		digitCnt++
		startIdx++

		if thouCnt > -1 {
			thouCnt++
		}

		if scratchNum.Cmp(baseZero) == 1 &&
			thouCnt == 3 {

			outRunes = append(outRunes, bNum.thousandsSeparator)
			startIdx++
			thouCnt = 0
		}

		if bNum.precision > 0 &&
			int(bNum.precision) == startIdx &&
			negValMode != ABSOLUTEPURENUMSTRFMTMODE {

			outRunes = append(outRunes, bNum.decimalSeparator)
			startIdx++
			thouCnt = 0
		}

	}

	if int(bNum.precision) >= digitCnt {

		delta := int(bNum.precision) - digitCnt + 1

		if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
			delta--
		}

		for k := 0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bNum.precision > 0 &&
				int(bNum.precision) == startIdx &&
				negValMode != ABSOLUTEPURENUMSTRFMTMODE {

				outRunes = append(outRunes, bNum.decimalSeparator)
				startIdx++
			}
		}
	}

	startIdx--

	// adjust for negative sign value
	if bNum.sign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		} else if negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, '(')
			startIdx += 2
		}

		// Must Be negValMode == ABSOLUTEPURENUMSTRFMTMODE

	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i := startIdx; i > sortLimit; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes), nil

}

// getActualNumberOfDigits - Returns the number of numeric digits
// in the absolute value of this BigIntNum instance. In addition,
// a boolean value is returned indicating whether the absolute value
// is zero.
//
// Examples
// ========
//
//	      123.45														5
//	1,234,567                               7
//
// -1,234,567                               7
//
//					 0															1
//	         0.00                           1
//	       012.34                           4
//	         0.1234													4
//	       - 0.1234													4
//	         0.123400												4
//	         0.0123400											4
//	 1,234,567.800												  8
//	         5                              1
func (bIntMolecule *bigIntNumMolecule) getActualNumberOfDigits(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (
	numberOfDigits *big.Int, isZeroValue bool, err error) {

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	numberOfDigits = big.NewInt(0)

	isZeroValue = false

	err = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.getActualNumberOfDigits",
		"")

	if err != nil {
		return numberOfDigits, isZeroValue, err
	}

	if bNum == nil {

		return numberOfDigits, isZeroValue,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {

		return numberOfDigits, isZeroValue, err

	}

	numOfDigits, errx := BigIntMath{}.GetMagnitude(bNum.absBigInt)

	if errx != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			"numOfDigits, errx := BigIntMath{}.GetMagnitude(bNum.absBigInt)\n"+
			"bNum.absBigInt='%v' Error='%v' ",
			ePrefix.String(),
			bNum.absBigInt.Text(10),
			errx.Error())

		return numberOfDigits, isZeroValue, err
	}

	numberOfDigits = big.NewInt(0).Add(numOfDigits, big.NewInt(1))

	if bNum.absBigInt.Cmp(big.NewInt(0)) == 0 {

		isZeroValue = true

	}

	return numberOfDigits, isZeroValue, err
}

// isBIntNumZero - Returns a boolean signaling whether a
// BigIntNum value is zero.
//
// NOTE:
// This method will first test the passed instance of BigIntNum
// to determine if that instance is valid, or not.
func (bIntMolecule *bigIntNumMolecule) isBIntNumZero(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.isBIntNumZero",
		"")

	if err != nil {
		return false, err
	}

	if bNum == nil {

		return false, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Vallidity Test on 'bNum'"))

	if err != nil {
		return false, err
	}

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}

	return false, nil
}

// newOne - Returns a BigIntNum Type with a value equal to '1' (one).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								1
//			1								1.0
//			2								1.00
//			3								1.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bIntMolecule *bigIntNumMolecule) newOne(
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.newOne",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	b, err := new(bigIntNumMechanics).newZero(
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	bNumNanobot := new(bigIntNumNanobot)

	if precision == 0 {

		err = bNumNanobot.setBigInt(
			&b,
			big.NewInt(1),
			0,
			ePrefix)

		if err != nil {

			return BigIntNum{}, err

		}

		return b, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	newVal := big.NewInt(0).Mul(big.NewInt(1), scaleVal)

	err = bNumNanobot.setBigInt(
		&b,
		newVal,
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return b, nil
}

// setBigIntExponent - Sets the numeric value using an integer
// multiplied by 10 raised to the power of the 'exponent'
// parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'bigI' is of type *big.Int.
//
// Input parameter 'exponent' is of type int.
//
// If exponent is less than +1, precision is set equal to exponent and
// bigI is unchanged. Example:
//
//	   bigI				exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//
// If exponent is greater than 0, bigI is multiplied by 10 raised to the
// power of exponent and precision is set equal to exponent.
//
//	   bigI				exponent			BigIntNum Result
//		 123456		 		   3							123456.000
func (bIntMolecule *bigIntNumMolecule) setBigIntExponent(
	bNum *BigIntNum,
	bigI *big.Int,
	exponent int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setBigIntExponent",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if bigI == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	if exponent < 1 {

		precision := uint(exponent * -1)

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			bigI,
			precision,
			ePrefix)

		if err != nil {
			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(\n" +
					"    bNum, bigI, precision, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// exponent must be greater than zero.
	// scale left exponent places and set precision to zero

	big10 := big.NewInt(10)
	scale := big.NewInt(int64(exponent))
	scaleValue := big.NewInt(0).Exp(big10, scale, nil)
	newBigI := big.NewInt(0).Mul(bigI, scaleValue)

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		newBigI,
		uint(exponent),
		ePrefix)

	if err != nil {
		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(\n" +
				"    bNum, newBigI, uint(exponent), ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setBigRat
//
// Sets the value of BigIntNum instance, passed as input parameter
// 'bNum', to that of input parameter 'ratNum', a rational number
// of type *big.Rat.
//
//	Input Parmeters
//	===============
//
//	bNum *BigIntNum
//
//	This instance of BigIntNum will be reconfigured using input
//	parameters 'ratNum', 'maxPrecision' and 'numSeps'
//
//
//	ratNum 				*big.Rat
//
//	The value of ratNum will be used to configure the current
//	BigIntNum instance and reset its value.
//
//
//	maxPrecision 	uint
//
//	The maximum precision for the resulting BigIntNum value
//	after it is reset to the value of input parameter 'ratNum'.
//	Precision will never be greater than 'maxPrecision'; however,
//	actual precision may be less than 'maxPrecision'.
//
//	Existing numeric separators (decimal separator, thousands separator
//	and currency symbol) in the current BigIntNum instance will remain
//	unchanged and will not be altered by this method.
//
//
//	numSeps				NumericSeparatorDto
//
//	The current BigIntNum instance will be reconfigured with
//	numeric separators provided by 'numSeps' an instance of
//	NumericSeparatorDto. Type NumericSeparatorDto contains
//	the decimal separator, thousands separator and currency
//	symbol.
//
//	type NumericSeparatorDto struct {
//		DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//		ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
//		CurrencySymbol     rune // Currency Symbol
//	}
//
//
//	IMPORTANT NOTE
//	==============
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntMolecule *bigIntNumMolecule) setBigRat(
	bNum *BigIntNum,
	ratNum *big.Rat,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setBigRat",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if ratNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	bIntNumAtom := new(bigIntNumAtom)

	err = bIntNumAtom.setNumericSeparatorsToDefaultIfEmpty(
		bNum, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = bIntNumAtom.setNumericSeparatorsToDefaultIfEmpty(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSeps, err := bIntNumAtom.getNumericSeparatorsDto(
		bNum,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numerator := big.NewInt(0).Set(ratNum.Num())

	denominator := big.NewInt(0).Set(ratNum.Denom())

	biPair, err := new(BigIntPair).
		NewBase(numerator, 0, denominator, 0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "biPair, err := new(BigIntPair).\n" +
				"NewBase(numerator, 0, denominator, 0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	biPair.MaxPrecision = maxPrecision

	biNum, err := BigIntMathDivide{}.PairFracQuotientNoNumSeps(biPair, numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "biNum, err := BigIntMathDivide{}.PairFracQuotientNoNumSeps(biPair, numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	biNumPrecision, err := biNum.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "biNumPrecision, err := biNum.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if biNumPrecision > maxPrecision {

		err = biNum.SetPrecision(maxPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = biNum.SetPrecision(maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	err = new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		&biNum,
		ePrefix.XCpy("bNum <- biNum"))

	return err
}

// setBigIntNumSeps
//
// Reconfigures a BigIntNum instance new values extracted from a
// *big.Int type and its associated precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//	Input Parmeters
//	===============
//
//	bNum				*BigIntNum
//
//	The instance of BigIntNum will be reconfigured with new
//	values based on the input parameters 'bigI', 'precision',
//
//
//	bigI *big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without decimal
//	digits.
//
//
//	precision		int
//
//	This unsigned integer (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places.
//
//	Example:
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
//
//	numSeps				NumericSeparatorDto
//
//	Input parameter 'bNum' (type BigIntNum) will be reconfigured
//	with numeric separators provided by input prameter 'numSeps',
//	an instance of NumericSeparatorDto. Type NumericSeparatorDto
//	contains the decimal separator, thousands separator and currency
//	symbol.
//
//		type NumericSeparatorDto struct {
//			DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//			ThousandsSeparator rune // Character used to separate thousands (1,000,000,000)
//			CurrencySymbol     rune // Currency Symbol
//		}
//
//	IMPORTANT NOTE
//	==============
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntMolecule *bigIntNumMolecule) setBigIntNumSeps(
	bNum *BigIntNum,
	bigI *big.Int,
	precision uint,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setBigIntNumSeps()",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if bigI == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	biNum := new(BigIntNum)

	err = new(bigIntNumNanobot).setBigInt(
		biNum,
		bigI,
		precision,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(\n" +
				"    biNum, bigI, precision, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSeps.SetDefaultsIfEmpty()

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		biNum,
		numSeps,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"    biNum, numSeps, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		biNum,
		ePrefix.XCpy("bNum <- biNum"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumUtility).bigIntNumCopyIn(\n" +
				"    bNum, biNum, ePrefix.XCpy(\"bNum <- biNum\")))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setBigRatNumSeps
//
// Sets the value of BigIntNum instance, passed as input parameter
// 'bNum', to that of input parameter 'ratNum', a rational number
// of type *big.Rat.
//
//	Input Parmeters
//	===============
//
//	bNum				*BigIntNum
//
//	The instance of BigIntNum will be reconfigured with new
//	values based on the following input parameters.
//
//
//	ratNum 			*big.Rat
//
//	The value of ratNum will be used to configure the current
//	BigIntNum instance and reset its value.
//
//
//	maxPrecision uint
//
//	The maximum precision for the resulting BigIntNum value
//	after it is reset to the value of input parameter 'ratNum'.
//	Precision will never be greater than 'maxPrecision'; however,
//	actual precision may be less than 'maxPrecision'.
//
//
//	numSeps				NumericSeparatorDto
//
//	The current BigIntNum instance will be reconfigured with
//	numeric separators provided by 'numSeps' an instance of
//	NumericSeparatorDto. Type NumericSeparatorDto contains
//	the decimal separator, thousands separator and currency
//	symbol.
//
//		type NumericSeparatorDto struct {
//			DecimalSeparator   rune // Character used to separate integer
//															//   and fractional digits ('.')
//			ThousandsSeparator rune // Character used to separate thousands
//															//   (1,000,000,000)
//			CurrencySymbol     rune // Currency Symbol
//		}
//
//	IMPORTANT NOTE
//	==============
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntMolecule *bigIntNumMolecule) setBigRatNumSeps(
	bNum *BigIntNum,
	bigRatNum *big.Rat,
	maxPrecision uint,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setBigRatNumSeps",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if bigRatNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	numerator := big.NewInt(0).Set(bigRatNum.Num())

	denominator := big.NewInt(0).Set(bigRatNum.Denom())

	biPair, err := new(BigIntPair).
		NewBase(numerator, 0, denominator, 0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "biPair, err := new(BigIntPair).\n" +
				"NewBase(numerator, 0, denominator, 0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	biPair.MaxPrecision = maxPrecision

	biNum, err := BigIntMathDivide{}.PairFracQuotientNoNumSeps(biPair, numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "biNum, err := BigIntMathDivide{}.PairFracQuotientNoNumSeps(\n" +
				"    biPair, numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	biNumPrecision, err := biNum.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "biNumPrecision, err := biNum.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if biNumPrecision > maxPrecision {

		err = biNum.SetPrecision(maxPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = biNum.SetPrecision(maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	numSeps.SetDefaultsIfEmpty()

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		&biNum,
		numSeps,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"    &biNum, numSeps, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		&biNum,
		ePrefix.XCpy("bNum <- biNum"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumUtility).bigIntNumCopyIn(\n" +
				"    bNum, &biNum, ePrefix.XCpy(\"bNum <- biNum\")))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setExpectedNumberOfDigits
//
// Sets the number of expected digits associated with the Absolute
// Value of input parameter 'bNum.absBigInt'. The number of expected
// digits value is stored in the data field,
// 'bNum.numberOfExpectedDigits'.
//
// Useful in tracking leading zeros.
func (bIntMolecule *bigIntNumMolecule) setExpectedNumberOfDigits(
	bNum *BigIntNum,
	numOfDigits *big.Int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setExpectedNumberOfDigits",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if numOfDigits == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfDigits' is a nil pointer!\n",
			ePrefix.String())

	}

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			bNum.precision,
			ePrefix)

		if err != nil {
			return err
		}
	}

	bNum.numberOfExpectedDigits = big.NewInt(0).Set(numOfDigits)

	return nil
}

// setNumStr - Initializes the BigIntNum instance
// for the numeric value of the number string input parameter.
// A number string is a string of numeric digits which may
// or may not be prefixed with a minus sign ('-'). The numeric
// string of digits may also contain a decimal separator such
// as a period ('.'). The decimal separator may be set by the
// user. See Method BigIntNum.SetDecimalSeparator(). The decimal
// separator is used to separate integer and fractional numeric
// digits within the number string.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bIntMolecule *bigIntNumMolecule) setNumStr(
	bNum *BigIntNum,
	numStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setNumStr()",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			0,
			ePrefix)

		if err != nil {
			return err
		}

	}

	if len(numStr) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is an EMPTY string!\n",
			ePrefix.String())
	}

	baseRunes := []rune(numStr)
	lBaseRunes := len(baseRunes)

	numSeps := NumericSeparatorDto{}
	numSeps.DecimalSeparator = bNum.decimalSeparator
	numSeps.ThousandsSeparator = bNum.thousandsSeparator
	numSeps.CurrencySymbol = bNum.currencySymbol

	numSeps.SetDefaultsIfEmpty()

	newSign := 1

	newPrecision := uint(0)

	newAbsBigInt := big.NewInt(0)

	baseTen := big.NewInt(10)

	isStartNumericDigits := false

	isEndNumericDigits := false

	isFractionalValue := false

	hasMinusSign := false

	hasLeftParen := false

	hasRightParen := false

	numOfNumericDigits := 0

	for i := 0; i < lBaseRunes; i++ {

		if isEndNumericDigits {
			continue
		}

		if baseRunes[i] == ',' && bNum.decimalSeparator != ',' {
			continue
		}

		if baseRunes[i] == '-' {
			hasMinusSign = true
			continue
		}

		if baseRunes[i] == '(' {

			if isStartNumericDigits == false {
				hasLeftParen = true
			}
			continue
		}

		if baseRunes[i] == ')' {

			if isStartNumericDigits == true &&
				hasLeftParen == true {
				hasRightParen = true
				isEndNumericDigits = true
			}

			continue
		}

		if baseRunes[i] == bNum.decimalSeparator {
			isFractionalValue = true
			continue
		}

		if baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			newAbsBigInt = big.NewInt(0).Mul(newAbsBigInt, baseTen)

			newAbsBigInt = big.NewInt(0).Add(newAbsBigInt,
				big.NewInt(int64(baseRunes[i]-48)))

			isStartNumericDigits = true
			numOfNumericDigits++

			if isFractionalValue {
				newPrecision++
			}
		}
	}

	if numOfNumericDigits == 0 {
		return fmt.Errorf("%v\n"+
			"Error: No numeric digits were found in input parameter 'numStr'.\n"+
			"numStr='%v'\n",
			ePrefix.String(),
			numStr)
	}

	if hasMinusSign == true ||
		(hasLeftParen == true && hasRightParen == true) {
		newSign = -1
	}

	bNum.Empty()
	bNum.sign = newSign
	bNum.precision = newPrecision
	bNum.absBigInt = big.NewInt(0).Set(newAbsBigInt)

	if bNum.sign == 1 {
		bNum.bigInt = big.NewInt(0).Set(newAbsBigInt)
	} else {
		bNum.bigInt = big.NewInt(0).Neg(newAbsBigInt)
	}

	bNum.scaleFactor = big.NewInt(0).Exp(baseTen,
		big.NewInt(int64(newPrecision)),
		nil)

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	return nil
}

// setINumMgr
//
// Receives an input parameter implementing the INumMgr interface and
// proceeds to set the current BigIntNum instance to its equivalent
// numeric value.
//
// Currently, the following 'mathops' Types implement the INumMgr
// interface:
//
//		Decimal,
//		IntAry,
//		NumStrDto,
//		BigIntNum
//	 BigIntFixedDecimal
//
// 'numMgr' must be a pointer to a type. This method will not accept
// 'numMgr' as a value. The pointer to the type is needed in or order to
// call methods on 'numMgr'.
//
// This method will test the validity of input parameter, 'numMgr'.
//
// Example 1:
//
//	dec, err := new(Decimal).NewNumStr(nStr)
//	bINum := new(BigIntNum)
//	err := bINum.SetINumMgr(&dec)
//
// Example 2:
// dec, err := new(Decimal).NewNumStr(nStr)
// bINum := BigIntNum{}
// err := bINum.SetINumMgr(dec.GetThisPointer())
//
// Example 3:
// dec := new(Decimal).NewPtr()
// err := dec.SetNumStr(nStr)
// bINum := BigIntNum{}
// err := bINum.SetINumMgr(dec)
//
// Example 4:
// fd := new(BigIntFixedDecimal).NewZero()
// err := fd.SetNumStr(numStr string)
// bINum := BigIntNum{}
// err := bINum.SetINumMgr(fd)
func (bIntMolecule *bigIntNumMolecule) setINumMgr(
	bNum *BigIntNum,
	numMgr INumMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntMolecule.lock == nil {
		bIntMolecule.lock = new(sync.Mutex)
	}

	bIntMolecule.lock.Lock()

	defer bIntMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMolecule.setNumStr()",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	err = numMgr.IsValid(ePrefix.XCpy("Testing 'numMgr'").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bigInt, err := numMgr.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bigInt, err := numMgr.GetBigInt()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bigInt, err := numMgr.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	precisionUint, err := numMgr.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "precisionUint, err := numMgr.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		bigInt,
		precisionUint,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(\n" +
				"    bNum, bigInt, precisionUint, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSepsDto, err := numMgr.GetNumericSeparatorsDto()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "numSepsDto, err := numMgr.GetNumericSeparatorsDto()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSepsDto,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"bNum, numSepsDto, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
