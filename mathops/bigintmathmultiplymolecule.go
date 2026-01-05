package mathops

import (
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathMultiplyMolecule struct {
	lock *sync.Mutex
}

// multiplyPair
//
//	Multiplication Operation
//	========================
//
//	This method receives a 'BigIntPair' instance labeled 'bPair'
//	and proceeds to multiply 'bPair.Big1' times 'bPair.Big2'. Both
//	'bPair.Big1' and 'bPair.Big2' are of type 'BigIntNum'.
//
//	  multiplier x multiplicand = product or result
//	  bPair.Big1 x bPair.Big2   = product or result
//
//	The result of this multiplication operation (a.k.a. 'product')
//	is returned as a BigIntNum type.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values.
//
//	Numeric Separator characters are typically encapsulated in a
//	NumericSeparatorDto type.
//
//	The returned BigIntNum multiplication 'Product' will contain
//	numeric separators (Decimal Separator, Thousands Separator and
//	Currency Symbol) derived from one of two possible sources.
//
//	Users have the option to supply an input parameter,
//	'outputNumSeps' of type NumericSeparatorDto. If this optional
//	parameter is provided, it will be used to configure the
//	returned BigIntNum multiplication 'product' from this method.
//	Note that the first valid NumericSeparatorDto in the
//	'outputNumSeps' series will be selected and used. There is no
//	need to provide more than one valid NumericSeparatorDto object
//	for parameter 'outputNumSeps'.
//
//	If the optional input parameter 'outputNumSeps' is NOT
//	provided, the returned BigIntNum instance will be
//	configured using numeric separators copied from input
//	parameter, 'bPair.Big1'.
//
//	Input Parameters
//	================
//
//	bPair                    BigIntPair
//
//	This instance of BigIntPair contains two numeric values of type
//	'BigIntNum'. These two 'BigIntNum' values are referred to as
//	bPair.Big1 and bPair.Big2. This method is designed to multiply
//	bPair.Big1 times bPair.Big2 and return the product as a
//	'BigIntNum' type.
//
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//	This object encapsulates an error prefix string
//	which is included in all returned error
//	messages. Usually, it contains the name of the
//	calling method or methods listed as a function
//	chain.
//
//	If no error prefix information is needed, set
//	this parameter to 'nil'.
//
//	Type ErrPrefixDto is included in the 'errpref'
//	software package:
//		"github.com/MikeAustin71/errpref".
//
//	outputNumSeps            ... NumericSeparatorDto
//
//	This method is defined as a variadic function in that
//	'outputNumSeps' is configured as an optional input parameter
//	meaning that it is NOT required. The user can choose to
//	provide a value for 'outputNumSeps', or not.
//
//	If the user chooses to provide a valid 'NumericSeparatorDto'
//	object for this parameter, it will be used to configure the
//	'BigIntNum' product value returned by this method.
//
//	Note that the first valid NumericSeparatorDto in the
//	'outputNumSeps' series will be selected and used. There is no
//	need to provide more than one valid NumericSeparatorDto object
//	for parameter 'outputNumSeps'.
//
//	Be advised that if the user chooses NOT to provide this
//	optional parameter, the 'BigIntNum' value returned by this
//	method will be configued using the 'NumericSeparatorDto'
//	copied from 'bPair.Big1'
//
//	Return Values
//	=============
//
//	BigIntNum
//
//	The product of the multiplication operation described above is
//	returned as BigIntNum type.
//
//	error
//
//	If no errors are encountered during method execution, this
//	return parameter is set to 'nil'.
func (bMathMulMolecule *bigIntMathMultiplyMolecule) multiplyPair(
	bPair BigIntPair,
	errPrefDto *ePref.ErrPrefixDto,
	outputNumSeps ...NumericSeparatorDto) (BigIntNum, error) {

	if bMathMulMolecule.lock == nil {
		bMathMulMolecule.lock = new(sync.Mutex)
	}

	bMathMulMolecule.lock.Lock()

	defer bMathMulMolecule.lock.Unlock()

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

	err = bPair.IsValid(ePrefix.XCpy("Validating 'bPair'.").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = bPair.IsValid(ePrefix.XCpy(\n" +
					"\"Validating 'bPair'.\").String())",
				ErrContext: "Input parameter 'bPair' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	var finalOutputNumSeps, bPairBig1NumSeps NumericSeparatorDto

	bPairBig1NumSeps, err = bPair.Big1.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bPairBig1NumSeps, err := bPair.Big1.\n" +
					"GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(outputNumSeps) > 0 {

		_, finalOutputNumSeps,
			err = new(numSepsDtoMechanics).selectValidNumSepInSeries(
			"outputNumSeps",
			"bPairBig1",
			&bPairBig1NumSeps,
			ePrefix,
			outputNumSeps...)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "_, finalOutputNumSeps, err = \n" +
						"    new(numSepsDtoMechanics).selectValidNumSepInSeries(\n" +
						"    \"outputNumSeps\", \"multiplier\")",
					ErrContext: "Error: Failed to Select 'finalOutputNumSeps'",
					ErrMessage: err.Error(),
				}
		}
	} else {

		// Checks validity of bPairBig1NumSeps

		err = new(numSepsDtoMechanics).copyNumSepsDto(
			&finalOutputNumSeps, // Destination
			&bPairBig1NumSeps,
			false, // Set Defaults if Empty
			ePrefix.XCpy("Copy 'bPairBig1NumSeps' Into 'finalOutputNumSeps'"))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numSepsDtoMechanics).copyNumSepsDto(\n" +
						"    &finalOutputNumSeps, &bPairBig1NumSeps, false,\n" +
						"    ePrefix.XCpy(\"Copy 'bPairBig1NumSeps' Into 'finalOutputNumSeps'\"))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return new(bigIntMathMultiplyNanobot).
		multiplyPairWithNumSeps(bPair, finalOutputNumSeps,
			ePrefix.XCpy("bPair & finalOutputNumSeps"))
}
