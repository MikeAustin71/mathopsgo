package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathDivideNanobot struct {
	lock *sync.Mutex
}

// pairFracQuotient
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the divisor.
//
// 'BigIntPair.maxPrecision' is used to control the maximum
// precision of the resulting fractional quotient. Be advised that
// this method is capable of calculating quotients with very long
// strings of fractional digits. Therefore, the user is advised to
// set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//	  Big1          BigIntNum  // The Dividend
//	  Big2          BigIntNum  // The Divisor
//	  maxPrecision  uint       // Controls Precision
//	}
//
// This method performs a division operation on BigIntNum parameters
// 'dividend'(BigIntPair.Big1) and 'divisor' (BigIntPair.Big2).
//
//	Dividend (BigIntPair.Big1) divided Divisor (BigIntPair.Big2) = quotient
//
// The resulting quotient is returned as a BigIntNum type
// representing the result 0f the division operation expressed as
// integer and fractional digits. The maximum number of fractional
// digits output to the result is controlled by BigIntPair.maxPrecision.
// Remember that the BigIntNum type specifies 'precision'. Precision
// is defined as the number of fractional digits to the right of the
// decimal place.
//
//		Examples
//		=========
//
//		Note: For all examples BigIntPair.maxPrecision is specified as '15'.
//
//		                                        Quotient
//		Dividend  divided by  Divisor   =   BigIntNum Integer    Precision    Result
//
//		  10.5         /         2      =                 525         2      5.25
//		  10           /         2      =                   5         0      5
//		  11.5         /         2.5    =                  46         1      4.6
//		  2.5          /        12.555  =     199123855037834        15      0.199123855037834
//		-12.555        /         2.5    =    -           5022         3     -5.022
//		-12.555        /         2      =    -          62775         4     -6.2775
//		- 2.5          /        12.555  =    -199123855037834        15     -0.199123855037834
//		 12.555        /       - 2.5    =    -           5022         3     -5.022
//		 12.555        /       - 2      =    -          62775         4     -6.2775
//		  2.5          /       -12.555  =    -199123855037834        15     -0.199123855037834
//		-12.555        /       - 2.5    =                5022         3      5.022
//		-12.555        /       - 2      =               62775         4      6.2775
//		- 2.5          /       -12.555  =     199123855037834        15      0.199123855037834
//		-10            /       - 2      =                   5         5      5.00000
//
//	 Numeric Separators
//	 ==================
//
//	 Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	 instance. A NumericSeparatorDto contains symbols or characters
//	 for the decimal separator, thousands separator and currency
//	 symbol. These separators are used when presenting numeric
//	 values in number strings.
//
//	 If any of the 'numSeps' Numeric Separator Components are set
//	 to zero, those components will be automatically reset to USA
//	 default values.
//
//	 The returned values ('quotient' and 'modulo') will be
//	 configured with 'numSeps'.
func (bIMathDivideNano *bigIntMathDivideNanobot) pairFracQuotient(
	bPair *BigIntPair,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (fracQuotient BigIntNum, err error) {

	if bIMathDivideNano.lock == nil {
		bIMathDivideNano.lock = new(sync.Mutex)
	}

	bIMathDivideNano.lock.Lock()

	defer bIMathDivideNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathDivideNanobot.pairFracQuotient",
		"")

	if err != nil {
		return fracQuotient, err
	}

	if bPair == nil {

		return fracQuotient,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bPair'",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return fracQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bPairBig2IsZero, err := bPair.Big2.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2IsZero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bPairBig2IsZero {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bPairBig2IsZero {",
				ErrMessage: "Attempted Divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	rDividend := big.NewRat(1, 1).SetInt(bPair.Big1.bigInt)

	rDivisor := big.NewRat(1, 1).SetInt(bPair.Big2.bigInt)

	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	numStr := rQuotient.FloatString(int(bPair.MaxPrecision))

	fracQuotient, err =
		new(BigIntNum).NewNumStr(numStr)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntNum).NewNumStr(numStr)",
				ErrContext: fmt.Sprintf("numStr= '%v'", numStr),
				ErrMessage: err.Error(),
			}
	}

	err = fracQuotient.TrimTrailingFracZeros()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fracQuotient.TrimTrailingFracZeros()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = fracQuotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = fracQuotient.SetNumericSeparatorsDto(\n" +
					"    numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// PairIntQuotient
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the Divisor. This method performs integer division
// on input parameters, 'Dividend' (BigIntPair.Big1) and 'Divisor'
// (BigIntPair.Big2).
//
// The result of this division operation returns an integer quotient
// of Dividend (BigIntPair.Big1) divided by Divisor
// (BigIntPair.Big2).
//
// The division operation performed by this method is T-Division or
// truncated division. See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathops/notes/divmodnote-letter.pdf.
//
// After completing the division operation, an integer quotient of
// type BigIntNum is returned.
//
//	Examples
//	=========
//
//	                                    Return Value
//	Dividend  divided by  Divisor  =  Integer Quotient
//
//	    5          /          2    =           2
//	    5.25       /          2    =           2
//	    2          /          4    =           0
//	  - 5          /          2    =          -2
//	  - 5.25       /          2    =          -2
//	  - 2          /          4    =           0
//	    5          /         -2    =          -2
//	    5.25       /         -2    =          -2
//	    2          /         -4    =           0
//	  - 5          /         -2    =           2
//	  - 5.25       /         -2    =           2
//	  - 2          /         -4    =           0
//	   12.555      /         -2.5  =          -5
//	  -12.555      /         -2.5  =           5
//	   12.555      /         -2    =          -6
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned value 'intQuotient' will be configured with
//	'numSeps'.
func (bIMathDivideNano *bigIntMathDivideNanobot) pairIntQuotient(
	bPair *BigIntPair,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (intQuotient BigIntNum, err error) {

	if bIMathDivideNano.lock == nil {
		bIMathDivideNano.lock = new(sync.Mutex)
	}

	bIMathDivideNano.lock.Lock()

	defer bIMathDivideNano.lock.Unlock()

	defer bIMathDivideNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathDivideNanobot.pairIntQuotient",
		"")

	if err != nil {
		return intQuotient, err
	}

	if bPair == nil {

		return intQuotient,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bPair'",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bPairBig2IsZero, err := bPair.Big2.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2IsZero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bPairBig2IsZero {

		return intQuotient,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bPairBig2IsZero {",
				ErrMessage: "Attempted Divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigIQuotient := big.NewInt(0).Quo(bPair.Big1.bigInt, bPair.Big2.bigInt)

	intQuotient, err = new(BigIntNum).NewBigInt(bigIQuotient, 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "intQuotient, err = new(BigIntNum).\n" +
					"    NewBigInt(bigIQuotient, 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = intQuotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = intQuotient.SetNumericSeparatorsDto(\n" +
					"    numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return intQuotient, nil
}

// pairMod
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the Divisor. The method proceeds to perform a
// modulo operation on input parameters 'dividend' (BigIntPair.Big1)
// and 'divisor' (BigIntPair.Big2). The modulo result is returned
// as a type BigIntNum.
//
// 'BigIntPair.maxPrecision' is used to control the precision of
// the resulting fractional modulo returned by this method. Be
// advised that this method is capable of calculating modulo values
// with very long strings of fractional digits. Therefore, the user
// is advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//	  Big1              BigIntNum  // The Dividend
//	  Big2              BigIntNum  // The Divisor
//	  maxPrecision      uint       // Controls Precision
//	}
//
// The modulo operation finds the remainder after division of one
// number by another (sometimes called modulus).
//
//	Wikipedia https://en.wikipedia.org/wiki/Modulo_operation
//
// The calculation of 'modulo' is based on T-Division (Truncate
// Division). See "Division and Modulus for Computer Scientists",
// DAAN LEIJEN, University of Utrecht Dept. of Computer Science,
// PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d)
//
//	r = D mod d = D − d ·q
//
// The modulo operation finds the remainder after division of one
// number by another. (r = D mod d = D − d ·q)
//
//	Examples
//	========
//
//	Dividend    mod by    Divisor    =  Modulo/Remainder
//
//	  12.555      %           2.5    =        0.055
//	  12.555      %           2      =        0.555
//	   2.5        %          12.555  =        2.5
//	 -12.555      %           2.5    =       -0.055
//	 -12.555      %           2      =       -0.555
//	 - 2.5        %          12.555  =       -2.5
//	  12.555      %         - 2.5    =        0.055
//	  12.555      %         - 2      =        0.555
//	  2.5         %         -12.555  =        2.5
//	 -12.555      %         - 2.5    =       -0.055
//	 -12.555      %         - 2      =       -0.555
//	 - 2.5        %         -12.555  =       -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIMathDivideNano *bigIntMathDivideNanobot) pairMod(
	bPair *BigIntPair,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (modulo BigIntNum, err error) {

	if bIMathDivideNano.lock == nil {
		bIMathDivideNano.lock = new(sync.Mutex)
	}

	bIMathDivideNano.lock.Lock()

	defer bIMathDivideNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathDivideNanobot.pairMod",
		"")

	if err != nil {
		return modulo, err
	}

	if bPair == nil {

		return modulo,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bPair'",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bPairBig2IsZero, err := bPair.Big2.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2IsZero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bPairBig2IsZero {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bPairBig2IsZero {",
				ErrMessage: "Attempted Divide by ZERO!\n" +
					"'bPair.Big2' has a ZERO value.",
			}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	moduloBigI := big.NewInt(0).Rem(bPair.Big1.bigInt, bPair.Big2.bigInt)

	bPairBig2Precision, err := bPair.Big2.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2Precision, err := bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	modulo, err = new(BigIntNum).NewBigInt(moduloBigI, bPairBig2Precision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(BigIntNum).NewBigInt(\n" +
					"    moduloBigI, bPairBig2Precision)",
				ErrContext: fmt.Sprintf("bPairBig2Precision= '%v'",
					bPairBig2Precision),
				ErrMessage: err.Error(),
			}
	}

	if modulo.precision > bPair.MaxPrecision {

		err = modulo.RoundToDecPlace(bPair.MaxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = modulo.RoundToDecPlace(bPair.MaxPrecision)\n" +
						fmt.Sprintf(" bPair.MaxPrecision= '%v'", bPair.MaxPrecision),
					ErrContext: "if modulo.precision > bPair.MaxPrecision {",
					ErrMessage: err.Error(),
				}
		}
	}

	err = modulo.TrimTrailingFracZeros()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = modulo.TrimTrailingFracZeros()",
				ErrMessage: err.Error(),
			}
	}

	err = modulo.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = modulo.SetNumericSeparatorsDto(\n" +
					"    numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, err
}

// pairQuotientMod
//
// Receives a BigIntPair type as an input parameter.
// 'BigIntPair.Big1' is treated as the Dividend. 'BigIntPair.Big2'
// is considered the Divisor. 'BigIntPair.maxPrecision' is used to
// control the precision of the resulting fractional quotient. Be
// advised that this method is capable of calculating quotients with
// very long strings of fractional digits. Therefore, the user is
// advised to set a relevant 'BigIntPair.maxPrecision' value.
//
//	type BigIntPair struct {
//	  Big1              BigIntNum  // The Dividend
//	  Big2              BigIntNum  // The Divisor
//	  maxPrecision      uint       // Controls Precision
//	}
//
// The method performs a division operation on BigIntNum input
// parameters 'dividend' (BigIntPair.Big1) and 'divisor'
// (BigIntPair.Big2). The result is a quotient and modulo
// (remainder). These values are returned as two BigIntNum types,
// 'quotient' and 'modulo'.
//
// The calculation of 'quotient' and 'modulo' is based on T-Division
// (Truncate Division). See "Division and Modulus for Computer
// Scientists", DAAN LEIJEN, University of Utrecht Dept. of Computer
// Science, PO.Box 80.089, 3508 TB Utrecht The Netherlands:
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
//
// This information is also available at:
//
//	mathopsgo/notes/divmodnote-letter.pdf.
//
// So, for q=quotient; D=Dividend d=Divisor r=Remainder or 'modulo' :
//
//	q = D div d = f(D/d) r = D mod d = D − d ·q
//
//	'quotient'  -  The integer result of dividing the 'dividend'
//	               by the 'divisor'.
//
//	'modulo'    -  The modulo operation finds the remainder after
//	               division of one number by another.
//	               (r = D mod d = D − d ·q)
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting 'modulo'. Precision is defined as the
// number of fractional digits to the right of the decimal point.
// Be advised that these calculations can support very large precision
// values.
//
//	Examples
//	=========
//
//	 Dividend  divided by  Divisor  =  Quotient  Modulo/Remainder
//
//	  12.555        /        2.5    =      5           0.055
//	  12.555        /        2      =      6           0.555
//	   2.5          /       12.555  =      0           2.5
//	 -12.555        /        2.5    =     -5          -0.055
//	 -12.555        /        2      =     -6          -0.555
//	 - 2.5          /       12.555  =      0          -2.5
//	  12.555        /      - 2.5    =     -5           0.055
//	  12.555        /      - 2      =     -6           0.555
//	   2.5          /      -12.555  =      0           2.5
//	 -12.555        /      - 2.5    =      5          -0.055
//	 -12.555        /      - 2      =      6          -0.555
//	 - 2.5          /      -12.555  =      0          -2.5
//
//	Numeric Separators
//	==================
//
//	Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	instance. A NumericSeparatorDto contains symbols or characters
//	for the decimal separator, thousands separator and currency
//	symbol. These separators are used when presenting numeric
//	values in number strings.
//
//	If any of the 'numSeps' Numeric Separator Components are set
//	to zero, those components will be automatically reset to USA
//	default values.
//
//	The returned values ('quotient' and 'modulo') will be
//	configured with 'numSeps'.
func (bIMathDivideNano *bigIntMathDivideNanobot) pairQuotientMod(
	bPair *BigIntPair,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (quotient BigIntNum, modulo BigIntNum, err error) {

	if bIMathDivideNano.lock == nil {
		bIMathDivideNano.lock = new(sync.Mutex)
	}

	bIMathDivideNano.lock.Lock()

	defer bIMathDivideNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathDivideNanobot.pairQuotientMod",
		"")

	if err != nil {
		return quotient, modulo, err
	}

	if bPair == nil {

		return quotient, modulo,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bPair'",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = numSeps.IsValid(ePrefix.XCpy("Testing numSeps").String())

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Testing numSeps\").String())",
				ErrContext: "numSeps is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	bPairBig2IsZero, err := bPair.Big2.IsZero()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bPairBig2IsZero, err := bPair.Big2.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bPairBig2IsZero {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bPairBig2IsZero {",
				ErrMessage: "Attempted Divide by ZERO!\n" +
					"'bPair.Big2.bigInt' has a ZERO value.",
			}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	scratch := big.NewInt(0)

	quotientBigI, moduloBigI := big.NewInt(0).QuoRem(
		bPair.Big1.bigInt,
		bPair.Big2.bigInt,
		scratch)

	quotient, err = new(BigIntNum).NewBigInt(quotientBigI, 0)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "quotient, err = new(BigIntNum).NewBigInt\n" +
					"    (quotientBigI, 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bPairBig2Precision, err := bPair.Big2.GetPrecisionUint()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPairBig2Precision, err := \n" +
					"    bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	modulo, err = new(BigIntNum).NewBigInt(moduloBigI, bPairBig2Precision)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = new(BigIntNum).\n" +
					"    NewBigInt(moduloBigI, bPairBig2Precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if modulo.precision > bPair.MaxPrecision {

		err = modulo.RoundToDecPlace(bPair.MaxPrecision)

		if err != nil {

			return quotient, modulo,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = modulo.RoundToDecPlace(bPair.MaxPrecision)",
					ErrContext: "if modulo.precision > bPair.MaxPrecision {",
					ErrMessage: err.Error(),
				}
		}

	}

	err = modulo.TrimTrailingFracZeros()

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = modulo.TrimTrailingFracZeros()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = quotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = quotient.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = modulo.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return quotient, modulo,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = modulo.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return quotient, modulo, err
}
