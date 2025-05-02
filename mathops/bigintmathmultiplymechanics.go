package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntMathMultiplyMechanics struct {
	lock *sync.Mutex
}

// bigIntMultiplyByTwoToPower - Multiplies a *big.Int number by powers
// of two and returns the result as a *big.Int value and associated
// precision specification.
//
//						product = multiplier X 2^exponent
//	          productPrecision = multiplierPrecision
//
// Examples:
// =========
//
//	multiplier			multiplierPrecision 	exponent		product
//
//	12345										5								15				4045.2096
//								(0.12345 x 2^15 = 4045.2096)
//
// -------------------------------------------------------------
//
//	571                     1                8        14617.6
//	            (57.1 x 2^8 = 14617.6)
//
// -------------------------------------------------------------
//
// Input Parameters
// ================
//
// multiplier						*big.Int -
//
//	'multiplier' will be multiplied by two to power of
//	'exponent' to generate the result or 'product'.
//
// multiplierPrecision	*big.Int	-
//
//	   The precision specification for 'multiplier'.
//			Precision specifies the number of digits to the
//	   right of the decimal place in 'multiplier'. This
//	   value must be greater than or equal to zero.
//
// exponent							uint			-
//
//	Two will be raised to the power of exponent and
//	multiplied by 'multiplier' to generate the result
//	or product.
//
// Return Values
// =============
//
// product 							*big.Int	- The multiplication result. product will be set equal
//
//	to multiplier times two to the power of exponent.
//
// productPrecision			*big.Int  - The precision specification for 'product'. Precision
//
//	specifies the number of digits to the right of the
//	decimal place in 'product'. This value will always be
//	greater than or equal to zero.
//
// err									error			- If 'multiplierPrecision' is less than zero, an error
//
//	will be returned.
//
// Note: This method will delete trailing fractional zeros from
//
//	the returned result (product).
func (bigIMathMultiplyMech *bigIntMathMultiplyMechanics) bigIntMultiplyByTwoToPower(
	multiplier *big.Int,
	multiplierPrecision *big.Int,
	exponent uint,
	methodCallChain string) (product *big.Int, productPrecision *big.Int, err error) {

	if bigIMathMultiplyMech.lock == nil {
		bigIMathMultiplyMech.lock = new(sync.Mutex)
	}

	bigIMathMultiplyMech.lock.Lock()

	defer bigIMathMultiplyMech.lock.Unlock()

	ePrefix := "bigIntMathMultiplyMechanics.bigIntMultiplyByTwoToPower()"

	product = big.NewInt(0)

	productPrecision = big.NewInt(0)

	err = nil

	if multiplier == nil {

		err = fmt.Errorf("%v\n"+
			"Method Call Chain: %v\n"+
			"Error: Input Parameter 'multiplier' is a nil pointer!\n",
			ePrefix, methodCallChain)

		return product, productPrecision, err
	}

	if multiplierPrecision == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'multiplierPrecision' is a nil pointer!\n",
			ePrefix)

		return product, productPrecision, err
	}

	bigZero := big.NewInt(0)

	if multiplierPrecision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'multiplierPrecision' is LESS THAN ZERO!\n"+
			"multiplierPrecision='%v'\n",
			ePrefix,
			multiplierPrecision.Text(10))

		return product, productPrecision, err
	}

	if multiplier.Cmp(big.NewInt(0)) == 0 {
		product = big.NewInt(0)
		productPrecision = big.NewInt(0)
		err = nil
		return product, productPrecision, err
	}

	product = big.NewInt(0).Lsh(multiplier, exponent)

	productPrecision.Set(multiplierPrecision)

	// Delete trailing fractional zeros
	if productPrecision.Cmp(bigZero) == 1 {
		// productPrecision > 0

		scrap := big.NewInt(0)

		biBase10 := big.NewInt(10)

		bigOne := big.NewInt(1)

		biBaseZero := big.NewInt(0)

		newProduct, mod10 := big.NewInt(0).QuoRem(product, biBase10, scrap)

		for mod10.Cmp(biBaseZero) == 0 && productPrecision.Cmp(bigZero) == 1 {

			product.Set(newProduct)

			productPrecision.Sub(productPrecision, bigOne)

			newProduct, mod10 = big.NewInt(0).QuoRem(product, biBase10, scrap)
		}
	}

	return product, productPrecision, err
}
