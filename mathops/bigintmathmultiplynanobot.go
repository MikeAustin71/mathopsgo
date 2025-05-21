package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathMultiplyNanobot struct {
	lock *sync.Mutex
}

// multiplyByTwoToPowerBigInt
//
// Multiplies a *big.Int number by powers of two and returns the
// result as a *big.Int value and associated precision
// specification.
//
//				 product = multiplier X 2^exponent
//
//				 productPrecision = multiplierPrecision
//
//	 Examples
//	 ========
//
//	 multiplier  multiplierPrecision  exponent   product
//
//	 12345               5               15     4045.2096
//	          (0.12345 x 2^15 = 4045.2096)
//
//	 571                 1                8     14617.6
//							            (57.1 x 2^8 = 14617.6)
//
//
//	 Input Parameters
//	 ================
//
//	 multiplier               *big.Int
//
//	 'multiplier' will be multiplied by two to power of 'exponent'
//	 to generate the result or 'product'.
//
//
//	 multiplierPrecision	     *big.Int
//
//	 The precision specification for 'multiplier'. Precision
//	 specifies the number of digits to the right of the decimal
//	 place in 'multiplier'. This value must be greater than or
//	 equal to zero.
//
//	 exponent                 uint
//
//	 Two will be raised to the power of exponent and multiplied
//	 by 'multiplier' to generate the result or product.
//
//
//	 errPrefDto					      *ePref.ErrPrefixDto
//
//	 This object encapsulates an error prefix string
//	 which is included in all returned error
//	 messages. Usually, it contains the name of the
//	 calling method or methods listed as a function
//	 chain.
//
//	 If no error prefix information is needed, set
//	 this parameter to 'nil'.
//
//	 Type ErrPrefixDto is included in the 'errpref'
//	 software package:
//	   "github.com/MikeAustin71/errpref".
//
//
//	 Return Values
//	 =============
//
//	 product                  *big.Int
//
//	 The multiplication result. product will be set equal to
//	 multiplier times two to the power of exponent.
//
//
//	 productPrecision         *big.Int
//
//	 The precision specification for 'product'. Precision
//	 specifies the number of digits to the right of the
//	 	decimal place in 'product'. This value will always be
//	 	greater than or equal to zero.
//
//
//	 err                      error
//
//	 If 'multiplierPrecision' is less than zero, an error will
//	 be returned. If no processing errors are encountered, this
//	 error return value will be set to 'nil'.
//
//
//	 Note
//	 ====
//
//	 This method will delete trailing fractional zeros from the
//	 returned result (product).
func (bigIMathMultiplyNanobot *bigIntMathMultiplyNanobot) multiplyByTwoToPowerBigInt(
	multiplier *big.Int,
	multiplierPrecision *big.Int,
	exponent uint,
	errPrefDto *ePref.ErrPrefixDto) (product *big.Int, productPrecision *big.Int, err error) {

	if bigIMathMultiplyNanobot.lock == nil {
		bigIMathMultiplyNanobot.lock = new(sync.Mutex)
	}

	bigIMathMultiplyNanobot.lock.Lock()

	defer bigIMathMultiplyNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyNanobot.multiplyByTwoToPowerBigInt",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	product = big.NewInt(0)

	productPrecision = big.NewInt(0)

	err = nil

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

	return product, productPrecision, nil
}

// multiplyPairWithNumSeps
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
//		Numeric Separators
//		==================
//
//		Numeric Separators specify the symbols or characters (runes)
//		used for the decimal separator, thousands separator and
//		currency symbol. These separators are used when displaying
//		numeric values in number strings.
//
//	 Input parameter 'numSepsDto' is an instance of
//	 NumericSepartorsDto. 'numSepDto' will be used to configure
//	 Numeric Separators in the returned BigIntNum instance.
//
//	 If any member elements of 'numSepsDto' are invalid, they
//	 will be automatically reset to USA default values.
func (bigIMathMultiplyNanobot *bigIntMathMultiplyNanobot) multiplyPairWithNumSeps(
	bPair BigIntPair,
	numSepsDto NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bigIMathMultiplyNanobot.lock == nil {
		bigIMathMultiplyNanobot.lock = new(sync.Mutex)
	}

	bigIMathMultiplyNanobot.lock.Lock()

	defer bigIMathMultiplyNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathMultiplyNanobot.multiplyPairWithNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = bPair.IsValid(ePrefix.XCpy("Testing bPair").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.IsValid(ePrefix.XCpy(\"Testing bPair\").String())",
				ErrContext: "Error: Input parameter 'bPair' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	big1BigInt, err := bPair.GetBig1BigInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "big1BigInt, err := bPair.GetBig1BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	big2BigInt, err := bPair.GetBig2BigInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "big2BigInt, err := bPair.GetBig2BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b3 := big.NewInt(0).Mul(big1BigInt, big2BigInt)

	big1Precision, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "big1Precision, err := bPair.Big1.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	big2Precision, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "big2Precision, err := bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	bResult, err := new(BigIntNum).NewBigInt(
		b3,
		big1Precision+big2Precision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bResult, err := new(BigIntNum).NewBigInt(\n" +
					"    b3, big1Precision+big2Precision)",
				ErrMessage: err.Error(),
			}

	}

	err = bResult.TrimTrailingFracZeros()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bResult.TrimTrailingFracZeros()",
				ErrMessage: err.Error(),
			}

	}

	numSepsDto.SetDefaultsIfEmpty()

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		&bResult, numSepsDto, ePrefix.XCpy("numSepsDto->bResult"))

	return bResult, err
}
