package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathMultiplyElectron struct {
	lock *sync.Mutex
}

// BigIntMultiply
//
// Receives two *big.Int numbers and their associated precision
// specifications. This method then proceeds to perform a
// multiplication operation by multiplying the 'multiplier' times
// the 'multiplicand' to generate the 'product'.
//
// 'multiplier', 'multiplicand' and 'product' are configured as
// pairs of *big.Int integer numbers and precision specifications.
//
// Taken together, an integer number and precision specification
// are used to define a fixed length floating point number.
//
//	Examples
//	========
//
//	In the multiplication operation, the number to be multiplied
//	is called the "multiplicand", while the number of times the
//	multiplicand is to be multiplied comes from the "multiplier".
//	Usually the multiplier is placed first and the multiplicand
//	is placed second.
//
//	For example, in the problem 5 x 3 equals 15, the 5 is the
//	'multiplier', 3 is the 'multiplicand' and 15 is the 'product',
//	or result.
//
//	  multiplier x multiplicand = product or result
//
//	Consider the following multiplication example.
//
//	  752.314 x 21.67894 = product
//
//	'multiplier' and 'multiplicand' would be configured as follows:
//
//	    multiplier            = 752314
//	    multiplierPrecision   = 3
//	    multiplicand          = 2167894
//	    multiplicandPrecision = 5
//
//	The 'product' value (16309.37006716) of 'multiplier' and
//	'multiplicand' would be calculated and configured as
//	follows:
//
//	    product           = 1630937006716
//	    productPrecision  = 8
//	    product value     = 16309.37006716
//
//	Input Parameters
//	================
//
//	multiplier               *big.Int
//
//	The number to be multiplied by 'multiplicand'
//
//
//	multiplierPrecision      *big.Int
//
//	The 'multiplier' precision or numeric digits to
//	the right of the decimal point.
//
//
//	multiplicand             *big.Int
//
//	The number to be multiplied by the 'multiplier'.
//
//
//	multiplicandPrecision    *big.Int
//
//	The 'multiplicand' precision or numeric digits to
//	the right of the decimal point.
//
//
//	Return Values
//	=============
//
//	product                  *big.Int
//
//	The product of the multiplier multiplied by
//	the multiplicand.
//
//
//	productPrecision         *big.Int
//
//	The precision specification for the returned product. Here,
//	the term precision is defined as the number of fractional
//	digits to the right of the decimal place in the returned
//	'product'. 'productPrecision' is always equal to or greater
//	than zero.
//
//
//	Note
//	====
//
//	This method removes trailing fractional zeros from the result.
//
//	  Example: 3.1200 is returned as 3.12
func (bIMathMultiplyElec *bigIntMathMultiplyElectron) multiplyBigInt(
	multiplier *big.Int,
	multiplierPrecision *big.Int,
	multiplicand *big.Int,
	multiplicandPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (product *big.Int, productPrecision *big.Int, err error) {

	if bIMathMultiplyElec.lock == nil {
		bIMathMultiplyElec.lock = new(sync.Mutex)
	}

	bIMathMultiplyElec.lock.Lock()

	defer bIMathMultiplyElec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyElectron.multiplyBigInt",
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

	if multiplicand == nil {

		return product, productPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplicand'",
			}
	}

	if multiplicandPrecision == nil {

		return product, productPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplicandPrecision'",
			}
	}

	bigZero := big.NewInt(0)

	if multiplierPrecision.Cmp(bigZero) == -1 {

		return product, productPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("multiplierPrecision='%v'",
					multiplierPrecision.Text(10)),
				ErrMessage: "Error: Input parameter multiplierPrecision is LESS THAN ZERO!",
			}
	}

	if multiplicandPrecision.Cmp(bigZero) == -1 {

		return product, productPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("multiplicandPrecision='%v'",
					multiplicandPrecision.Text(10)),
				ErrMessage: "Error: Input parameter multiplicandPrecision is LESS THAN ZERO!",
			}
	}

	productPrecision.Add(multiplierPrecision, multiplicandPrecision)

	product = big.NewInt(0).Mul(multiplier, multiplicand)

	if product.Cmp(bigZero) == 0 {
		productPrecision = big.NewInt(0)
	}

	// Delete trailing fractional zeros
	// If productPrecision > 0

	if productPrecision.Cmp(bigZero) == 1 {

		bigOne := big.NewInt(1)
		scrap := big.NewInt(0)
		biBase10 := big.NewInt(10)
		biBaseZero := big.NewInt(0)
		newProduct, mod10 := big.NewInt(0).QuoRem(product, biBase10, scrap)

		for mod10.Cmp(biBaseZero) == 0 && productPrecision.Cmp(bigZero) == 1 {
			product.Set(newProduct)
			productPrecision.Sub(productPrecision, bigOne)
			newProduct, mod10 = big.NewInt(0).QuoRem(product, biBase10, scrap)
		}
	}

	return product, productPrecision, nil
}

// newBigIntMathMultiplyZero
//
// Creates a BigIntMathMultiply instance with data variables
// initialized to zero.
func (bIMathMultiplyElec *bigIntMathMultiplyElectron) newBigIntMathMultiplyZero(
	errPrefDto *ePref.ErrPrefixDto) (BigIntMathMultiply, error) {

	if bIMathMultiplyElec.lock == nil {
		bIMathMultiplyElec.lock = new(sync.Mutex)
	}

	bIMathMultiplyElec.lock.Lock()

	defer bIMathMultiplyElec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyElectron.newBigIntMathMultiplyZero",
		"")

	if err != nil {
		return BigIntMathMultiply{}, err
	}

	b2Math := new(BigIntMathMultiply)

	b2Math.Input, err = new(BigIntPair).New()

	baseZero := big.NewInt(0)

	b2Math.Result, err = new(BigIntNum).NewBigInt(baseZero, 0)

	if err != nil {

		return BigIntMathMultiply{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "b2Math.Result, err = new(BigIntNum).NewBigInt(\n" +
					"    baseZero, 0)",
				ErrMessage: err.Error(),
			}
	}

	return *b2Math, nil
}
