package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

type bigIntMathMultiplyMechanics struct {
	lock *sync.Mutex
}

// multiplyBigIntsBigIntNum
//
// Receives two *big.Int numbers and their associated precision
// specifications. This method then proceeds to perform a
// multiplication operation by multiplying the 'multiplier' times
// the 'multiplicand'.
//
// In the multiplication operation, the number to be multiplied
// is called the "multiplicand", while the number of times the
// multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is
// placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the
// 'multiplier' and 3 is the 'multiplicand', The result of this
// calculation, '15', is designated as the 'product'.
//
//							multiplier x multiplicand = product or result
//
//	 Input Parameters
//	 ================
//
//	 multiplier               *big.Int
//	   The number to be multiplied by 'multiplicand'
//
//	 multiplierPrecision      uint
//	   The 'multiplier' precision or numeric digits after the
//	   decimal point in 'multiplier'.
//
//	 multiplicand             *big.Int
//	   The number to be multiplied by the 'multiplier'.
//
//	 multiplicandPrecision    uint
//	   The 'multiplicand' precision or numeric digits after the
//	   decimal point.
//
//	 numSepsDto               NumericSeparatorDto
//	   Numeric sepatators include the Decimal separator, Thousands
//	   separator and Currency symbol characters. Numeric separators
//	   are used to parse number strings and display numeric values
//	   formatted as number strings.
//
//	   This method will configure the returned BigIntNum 'product'
//	   with Numeric Separators current configured in the
//	   'multiplier'.
//
//	   If any elements of 'numSepsDto' are invalid, they will be
//	   automatically reset to USA defaults.
//
//	 Return Values
//	 =============
//
//	 BigIntNum
//	   This method performs the multiplication operation and
//	   returns the result or 'product' as a BigIntNum type.
//
//	 error
//	  If no errors are encountered by this method during
//	  execution, the returned error value will be set to
//	  'nil'.
func (bigIMathMultiplyMech *bigIntMathMultiplyMechanics) multiplyBigIntsBigIntNum(
	multiplier *big.Int,
	multiplierPrecision uint,
	multiplicand *big.Int,
	multiplicandPrecision uint,
	numSepsDto NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bigIMathMultiplyMech.lock == nil {
		bigIMathMultiplyMech.lock = new(sync.Mutex)
	}

	bigIMathMultiplyMech.lock.Lock()

	defer bigIMathMultiplyMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyMechanics.multiplyBigIntsBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if multiplier == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplier'",
			}
	}

	if multiplicand == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplicand'",
			}
	}

	multiplierPrecisionBigInt :=
		big.NewInt(0).SetUint64(uint64(multiplierPrecision))

	multiplicandPrecisionBigInt :=
		big.NewInt(0).SetUint64(uint64(multiplicandPrecision))

	productBInt, productPrecisionBInt, err :=
		new(bigIntMathMultiplyElectron).multiplyBigInt(
			multiplier,
			multiplierPrecisionBigInt,
			multiplicand,
			multiplicandPrecisionBigInt,
			ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "productBInt, productPrecisionBInt, err :=\n" +
					"    new(bigIntMathMultiplyElectron).multiplyBigInt(\n" +
					"    multiplier, multiplierPrecisionBigInt, multiplicand,\n" +
					"    multiplicandPrecisionBigInt, ePrefix)",
				ErrContext: fmt.Sprintf("multiplier= '%v'; multiplierPrecisionBigInt= '%v'\n"+
					"multiplicand= '%v'; multiplicandPrecisionBigInt= '%v'",
					multiplier.Text(10), multiplierPrecisionBigInt.Text(10),
					multiplicand.Text(10), multiplicandPrecisionBigInt.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	biNumProduct, err := new(BigIntNum).NewBigIntBigPrecision(
		productBInt,
		productPrecisionBInt)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "biNumProduct, err := new(BigIntNum).NewBigIntBigPrecision(\n" +
					"    productBInt, productPrecisionBInt)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSepsDto.SetDefaultsIfEmpty()

	err = biNumProduct.SetNumericSeparatorsDto(numSepsDto)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = biNumProduct.SetNumericSeparatorsDto(numSepsDto)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return biNumProduct, nil
}

// multiplyByTenToPowerBigInt
//
// Multiplies a *big.Int number by ten to the power of 'exponent'.
// 'exponent' is an input parameter of type *big.Int.
//
//	product = multiplier x 10^exponent
//
// The result is returned as a *big.Int type ('product') with a
// precision specification ('productPrecision').
//
//	Input Parameters
//	================
//
//	multiplier            *big.Int
//	  'multiplier' will be multiplied by 10 to power of 'exponent'
//	  to generate the result or 'product'.
//
//	multiplierPrecision   *big.Int
//	  The precision specification for 'multiplier'. Precision
//	  specifies the number of digits to the right of the decimal
//	  place in 'multiplier'. This value must be greater than or
//	  equal to zero.
//
//	exponent              *big.Int
//	  Ten will be raised to the power of exponent and multiplied by
//	  'multiplier' to generate the result or product. 'exponent'
//	  can be a negative value.
//
//	errPrefDto             *ePref.ErrPrefixDto
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
//
//	Return Values
//	=============
//
//	product               *big.Int
//	  The result generated by multiplying 'multiplier' by ten to
//	  the power of 'exponent'.
//
//	productPrecision      *big.Int
//	  The precision specification for 'product'. Precision specifies
//	  the number of digits to the right of the decimal place in
//	  'product'. This value will always be greater than or equal to
//	  zero.
//
//	err                   error
//	  If 'multiplierPrecision' or 'exponent' are less than zero, an
//	  error will be returned. If no processing errors are encountered
//	  during execution, value of 'nil' will be returned for 'err'.
//
//	Note
//	====
//
//	This method will remove trailing fractional zeros from the final
//	result (product).
func (bigIMathMultiplyMech *bigIntMathMultiplyMechanics) multiplyByTenToPowerBigInt(
	multiplier *big.Int,
	multiplierPrecision *big.Int,
	exponent *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (
	product *big.Int, productPrecision *big.Int, err error) {

	if bigIMathMultiplyMech.lock == nil {
		bigIMathMultiplyMech.lock = new(sync.Mutex)
	}

	bigIMathMultiplyMech.lock.Lock()

	defer bigIMathMultiplyMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyMechanics.multiplyByTenToPowerBigInt",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	product = big.NewInt(0)
	productPrecision = big.NewInt(0)

	if multiplier == nil {

		return product, productPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplier'",
			}
	}

	if multiplierPrecision == nil {

		return product, productPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplierPrecision'",
			}
	}

	if exponent == nil {

		return product, productPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'exponent'",
			}
	}

	bigZero := big.NewInt(0)

	if multiplierPrecision.Cmp(bigZero) == -1 {

		return product, productPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("multiplierPrecision='%v'", multiplierPrecision.Text(10)),
				ErrMessage: "Error: Input Parameter 'multiplierPrecision' is LESS THAN ZERO!",
			}
	}

	bigTen := big.NewInt(10)

	bigOne := big.NewInt(1)

	if exponent.Cmp(bigZero) == -1 {
		// if exponent is less than zero
		// apply special processing...

		exponent.Neg(exponent)

		divisor := big.NewInt(0).Exp(bigTen, exponent, nil)

		quoFrac, quoFracPrecision, err := new(BigIntMathDivide).
			BigIntFracQuotient(
				bigOne,
				big.NewInt(0),
				divisor,
				big.NewInt(0),
				exponent)

		if err != nil {

			return product, productPrecision,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "quoFrac, quoFracPrecision, err := \n" +
						"    new(BigIntMathDivide). BigIntFracQuotient(\n" +
						"    bigOne, big.NewInt(0), divisor,  big.NewInt(0), exponent)",
					ErrContext: fmt.Sprintf("divisor= '%v'; exponent= '%v'",
						divisor.Text(10), exponent.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		product, productPrecision, err =
			new(bigIntMathMultiplyElectron).multiplyBigInt(
				multiplier,
				multiplierPrecision,
				quoFrac,
				quoFracPrecision,
				ePrefix)

		if err != nil {

			return product, productPrecision,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "product, productPrecision, err =\n" +
						"    new(bigIntMathMultiplyElectron).multiplyBigInt(\n" +
						"    multiplier, multiplierPrecision, quoFrac, quoFracPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("multiplier= '%v'; multiplierPrecision= '%v'\n"+
						"quoFrac= '%v'; quoFracPrecision= '%v'",
						multiplier.Text(10), multiplierPrecision.Text(10),
						quoFrac.Text(10), quoFracPrecision.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		// Delete trailing fractional zeros
		// if productPrecision > 0
		if productPrecision.Cmp(bigZero) == 1 {

			scrap := big.NewInt(0)
			newProduct, mod10 := big.NewInt(0).QuoRem(product, bigTen, scrap)

			for mod10.Cmp(bigZero) == 0 && productPrecision.Cmp(bigZero) == 1 {
				product.Set(newProduct)
				productPrecision.Sub(productPrecision, bigOne)
				newProduct, mod10 = big.NewInt(0).QuoRem(product, bigTen, scrap)
			}
		}

		// Successful Completion
		return product, productPrecision, nil
	}

	delta := big.NewInt(0)
	scale := big.NewInt(0)

	if multiplier.Cmp(bigZero) == 0 {
		product = big.NewInt(0)
		productPrecision = big.NewInt(0)

	} else if exponent.Cmp(bigZero) == 0 {
		product.Set(multiplier)
		productPrecision.Set(multiplierPrecision)

	} else if exponent.Cmp(multiplierPrecision) == 0 {
		product.Set(multiplier)
		productPrecision = big.NewInt(0)

	} else if exponent.Cmp(multiplierPrecision) == 1 {
		// exponent > multiplierPrecision
		delta = big.NewInt(0).Sub(exponent, multiplierPrecision)
		scale = big.NewInt(0).Exp(bigTen, delta, nil)
		product = big.NewInt(0).Mul(multiplier, scale)
		productPrecision = big.NewInt(0)

	} else {
		// multiplierPrecision must be GREATER THAN exponent
		product.Set(multiplier)
		productPrecision.Sub(multiplierPrecision, exponent)
	}

	// Delete trailing fractional zeros
	// if productPrecision > 0
	if productPrecision.Cmp(bigZero) == 1 {
		scrap := big.NewInt(0)
		newProduct, mod10 := big.NewInt(0).QuoRem(product, bigTen, scrap)

		for mod10.Cmp(bigZero) == 0 && productPrecision.Cmp(bigZero) == 1 {
			product.Set(newProduct)
			productPrecision.Sub(productPrecision, bigOne)
			newProduct, mod10 = big.NewInt(0).QuoRem(product, bigTen, scrap)
		}
	}

	//Successful completion
	return product, productPrecision, nil
}

// multiplyByTwoToPowerBigIntNum
//
// Multiplies a *big.Int number by powers of two and returns
// the result as a BigIntNum type.
//
//	 Be Advised
//	 ==========
//
//	 This method differs BigIntMathMultiply.BigIntMultiplyByTwoToPower.
//	 BigIntMathMultiply.BigIntMultiplyByTwoToPower returns *big.Int
//	 types.
//
//	 This method returns a BigIntNum type for the result.
//
//	 Examples
//	 ========
//
//		product = multiplier X 2^exponent
//
//	 multiplier   multiplierPrecision    exponent       product
//
//	 12345                 5                15          4045.2096
//	            (0.12345 x 2^15 = 4045.2096)
//
//
//	 571                   1                 8          14617.6
//	              (57.1 x 2^8 = 14617.6)
//
//	 Note
//	 ====
//
//	 This method will delete trailing fractional zeros from the
//	 returned product.
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
//	 The returned value ('BigIntNum') will be configured with
//	 'numSeps' Numeric Separators.
func (bigIMathMultiplyMech *bigIntMathMultiplyMechanics) multiplyByTwoToPowerBigIntNum(
	multiplier *big.Int,
	multiplierPrecision uint,
	exponent uint,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bigIMathMultiplyMech.lock == nil {
		bigIMathMultiplyMech.lock = new(sync.Mutex)
	}

	bigIMathMultiplyMech.lock.Lock()

	defer bigIMathMultiplyMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyMechanics.multiplyByTwoToPowerBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	product, productPrecision, err :=
		new(bigIntMathMultiplyNanobot).multiplyByTwoToPowerBigInt(
			multiplier,
			big.NewInt(0).SetUint64(uint64(multiplierPrecision)),
			exponent,
			ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	// error should never trigger because productPrecision will
	// never be greater than multiplierPrecision.
	biNum, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "biNum, err := new(BigIntNum).NewBigIntBigPrecision(\n" +
					"    product, productPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSeps.SetDefaultsIfEmpty()

	err = biNum.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = biNum.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return biNum, nil
}

// multiplyBigIntFixedDecimals
//
// This method receives two BigIntFixedDecimal types and then
// proceeds to perform a multiplication operation by multiplying
// the 'multiplier' by the 'multiplicand' to generate the
// 'product'.
//
// In the multiplication operation, the number to be multiplied is
// called the "multiplicand", while the number of times the
// multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is
// placed second.
//
//	 Examples
//	 ========
//
//	 For example, in the problem 5 x 3 equals 15, the 5 is the
//	 'multiplier', 3 is the 'multiplicand' and 15 is the 'product'
//	 or result.
//
//	   multiplier x multiplicand = product or result
//
//	 'multiplier', 'multiplicand' and 'product' are
//	 BigIntFixedDecimal types which may be used to defined fixed
//	 length floating point numbers.
//
//	 BigIntFixedDecimal
//	 ==================
//
//	 The BigIntFixedDecimal structure is defined as
//
//	 type BigIntFixedDecimal struct {
//	   integerNum *big.Int  -  All the numeric digits, both integer and fractional,
//	                           necessary to define a fixed length floating point number.
//	                           The number of digits to the right of the decimal place
//	                           is specified by the data field,
//	                           BigIntFixedDecimal.precision.
//
//	   precision  uint      -  Specifies the number of digits to the right of the decimal
//	                           place in the series of numeric digits represented by data
//	                           field BigIntFixedDecimal.integerNum.
//
//	 }
//
//	 To represent the floating point number 52.459, a
//	 BigIntFixedDecimal Structure would be configured as follows:
//
//	    BigIntFixedDecimal.integerNum = 52459
//	    BigIntFixedDecimal.precision  = 3
//
//	 Consider the following multiplication example:
//	    product = 752.314 x 21.67894 = 16309.37006716
//
//	 'multiplier' and 'multiplicand' would be configured as follows:
//
//	    multiplier.integerNum   = 752314
//	    multiplier.precision    = 3
//	    multiplicand.integerNum = 2167894
//	    multiplicand.precision  = 5
//
//	 The 'product' would be calculated as follows:
//
//	    product.integerNum  = 1630937006716
//	    product.precision   = 8
//
//	 Input Parameters
//	 ================
//
//	 multiplier             BigIntFixedDecimal
//	   The number to be multiplied by 'multiplicand'
//
//		multiplicand          BigIntFixedDecimal
//		  The number to be multiplied by the 'multiplier'.
//
//
//	 Return Values
//	 =============
//
//	 product                BigIntFixedDecimal
//	   The product of the 'multiplier' multiplied by	the 'multiplicand'.
//
//	 err                    error
//	   If no errors are encountered during execution, this returned
//	   error value will be set to 'nil'.
//
//	 Numeric Separators
//	 ==================
//
//	 Numeric sepatators include the Decimal separator, Thousands
//	 separator and Currency symbol characters. Numeric separators
//	 are used to parse number strings and display numeric values
//	 formatted as number strings.
//
//	 This method will configure the returned BigIntFixedDecimal
//	 'product' with Numeric Separators current configured in the
//	 'multiplier'.
func (bigIMathMultiplyMech *bigIntMathMultiplyMechanics) multiplyBigIntFixedDecimals(
	multiplier BigIntFixedDecimal,
	multiplicand BigIntFixedDecimal,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (product BigIntFixedDecimal, err error) {

	if bigIMathMultiplyMech.lock == nil {
		bigIMathMultiplyMech.lock = new(sync.Mutex)
	}

	bigIMathMultiplyMech.lock.Lock()

	defer bigIMathMultiplyMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyMechanics.multiplyBigIntFixedDecimals",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	product = new(BigIntFixedDecimal).NewZero(0)

	err = multiplier.IsValid(ePrefix.XCpy("Testing multiplier").String())

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = multiplier.IsValid(ePrefix.XCpy(\"Testing multiplier\").String())",
				ErrContext: "Input parameter 'multiplier' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	err = multiplicand.IsValid(ePrefix.XCpy("Testing multiplicand").String())

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = multiplicand.IsValid(ePrefix.XCpy(\"Testing multiplicand\").String())",
				ErrContext: "Input parameter 'multiplicand' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	multiplierBigInt, err := multiplier.GetInteger()

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierBigInt, err := multiplier.GetInteger()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplierPrecision, err := multiplier.GetPrecision()

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplierPrecision, err := multiplier.GetPrecision()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplierPrecisionBigInt := big.NewInt(0).
		SetUint64(uint64(multiplierPrecision))

	multiplicandBigInt, err := multiplicand.GetInteger()

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplicandBigInt, err := multiplicand.GetInteger()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplicandPrecision, err := multiplicand.GetPrecision()

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "multiplicandPrecision, err := multiplicand.GetPrecision()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	multiplicandPrecisionBigInt := big.NewInt(0).
		SetUint64(uint64(multiplicandPrecision))

	result, resultPrecision, err :=
		new(bigIntMathMultiplyElectron).multiplyBigInt(
			multiplierBigInt,
			multiplierPrecisionBigInt,
			multiplicandBigInt,
			multiplicandPrecisionBigInt,
			ePrefix)

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "result, resultPrecision, err :=\n" +
					"    new(bigIntMathMultiplyElectron).multiplyBigInt(\n" +
					"    multiplierBigInt, multiplierPrecisionBigInt, multiplicandBigInt,\n" +
					"    multiplicandPrecisionBigInt, ePrefix)",
				ErrContext: fmt.Sprintf("multiplierBigInt= '%v'; multiplierPrecisionBigInt= '%v'\n"+
					"multiplicandBigInt= '%v'; multiplicandPrecisionBigInt= '%v'",
					multiplierBigInt.Text(10), multiplierPrecisionBigInt.Text(10),
					multiplicandBigInt.Text(10), multiplicandPrecisionBigInt.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	biMaxUint := big.NewInt(int64(math.MaxUint32))

	if resultPrecision.Cmp(biMaxUint) > 1 {
		delta := big.NewInt(0).Sub(resultPrecision, biMaxUint)
		delta.Sub(delta, big.NewInt(1))
		bigTen := big.NewInt(10)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		result.Quo(result, scale)
		bigFive := big.NewInt(5)

		if result.Cmp(big.NewInt(0)) == -1 {
			bigFive.Neg(bigFive)
		}

		result.Add(result, bigFive)
		result.Quo(result, bigTen)
		resultPrecision.Set(biMaxUint)
	}

	numSeps.SetDefaultsIfEmpty()

	err = new(bigIntFixedDecAtom).setNumericValue(
		&product,
		result,
		uint(resultPrecision.Uint64()),
		numSeps,
		ePrefix.XCpy("Setting 'product'"))

	return product, err
}

// multiplyPairNoNumSeps
//
// Receives a BigIntPair instance and proceeds to multiply
// bPair.Big1 by bPair.Big2. Both 'Big1' and 'Big2' are of
// type 'BigIntNum'.
//
//	bPair.Big1 x bPair.Big2 = Result
//
// The result of this multiplication operation is returned as a
// BigIntNum type.
//
// The returned BigIntNum multiplication 'Result' will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// configured for the BigIntNum, 'bPair.Big1'
func (bigIMathMultiplyMech *bigIntMathMultiplyMechanics) multiplyPairNoNumSeps(
	bPair BigIntPair,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bigIMathMultiplyMech.lock == nil {
		bigIMathMultiplyMech.lock = new(sync.Mutex)
	}

	bigIMathMultiplyMech.lock.Lock()

	defer bigIMathMultiplyMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyMechanics.multiplyPairNoNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = bPair.Big1.IsValid(ePrefix.XCpy("Testing  bPair.Big1").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "Error: bPair.Big1 is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bPair.Big1,
		ePrefix.XCpy("bPair.Big1->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	return new(bigIntMathMultiplyNanobot).
		multiplyPairWithNumSeps(bPair, numSepsDto, ePrefix)
}
