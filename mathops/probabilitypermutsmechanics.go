package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type probabilityPermutsMechanics struct {
	lock *sync.Mutex
}

// permutationsNoRepsBigInt
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. This calculation assumes that repetitions ARE NOT
//	allowed.
//
//	Input parameters 'numOfItems' and 'numOfItemsPicked' are both
//	passed as type *big.Int. Both input parameters must be non-zero,
//	positive integer numbers. 'numOfItems' must be equal to or greater
//	than 'numOfItemsPicked'.
//
//	Input Parameters
//	================
//
//	numOfItems               *big.Int
//	  Must be a positive integer number greater than zero.
//	  'numOfItems' must be greater than or equal to 'numOfItemsPicked'.
//
//	numOfItemsPicked         *big.Int
//	  Must be a positive integer number greater than zero.
//	  'numOfItemsChosen' must be less than or equal to 'numOfItems'.
//
//	numSeps                  NumericSeparatorDto
//
//	  Numeric Separators consist of the decimal separator, thousands
//	  seprator, and currency symbol used to cofigure custom number
//	  strings.
//
//	  The new instance of BigIntNum returned by this method will
//	  contain a copy of the Numeric Separators contained in this
//	  input parameter (numSeps).
//
//	  If input parameter 'numSeps' is unconfigured, empty or
//	  invalid, an error will be returned.
//
//	  Be sure to call Probability.SetNumSeps() before you call this
//	  method, if Non-USA Numeric Separators are required.
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
//	BigIntNum
//	  If the calculation is successful, the result is returned as a
//	  BigIntNum type. If the calculation fails, the error return is
//	  populated.
//
//	error
//	  If the calculation is successful, this return value is 'nil'. If
//	  the calculation fails, 'error' is populated with an appropriate
//	  error message.
//
//	Calculation
//	===========
//
//	In the following permutation formula, n= 'numOfItems' and
//	r = 'numOfItemsPicked'
//
//	                    n!
//	          nPr  =  ------
//	                  (n-r)!
//
//
//
//	     Where n is the number of things to select from,
//	     and we pick r of them, repetition is NOT allowed,
//	     and order matters.
//
//	This permutation calculation assumes NO REPETITIONS!
//
//	Example
//	=======
//
//	How many ways can first and second place be awarded to 10 people?
//
//	                    10!
//	Answer =  10P2  =  ------  =  10 x 9 = 90
//	                     8!
//
//	Note: 0! = 1
func (probPermutnsMech *probabilityPermutsMechanics) permutationsNoRepsBigInt(
	numOfItems *big.Int,
	numOfItemsPicked *big.Int,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if probPermutnsMech.lock == nil {
		probPermutnsMech.lock = new(sync.Mutex)
	}

	probPermutnsMech.lock.Lock()

	defer probPermutnsMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"probabilityPermutsMechanics.permutationsNoRepsBigInt()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if numOfItems == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numOfItems'",
			}
	}

	if numOfItemsPicked == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numOfItemsPicked'",
			}
	}

	err = numSeps.IsValid("Validating 'numSeps")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(\"Validating 'numSeps\")",
				ErrContext: "Input Parameter 'numSeps' is invalid",
				ErrMessage: err.Error(),
			}
	}

	bigZero := big.NewInt(0)

	if numOfItems.Cmp(bigZero) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
			}
	}

	if numOfItemsPicked.Cmp(bigZero) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
			}
	}

	if numOfItems.Cmp(bigZero) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
			}
	}

	if numOfItemsPicked.Cmp(bigZero) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
			}
	}

	if numOfItemsPicked.Cmp(numOfItems) == 1 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'!",
			}
	}

	bigOne := big.NewInt(1)

	if numOfItemsPicked.Cmp(bigOne) == 0 {

		bINumOne, err := new(BigIntNum).NewBigIntNumSeps(numOfItems, 0, numSeps)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bINumOne, err := new(BigIntNum).NewBigIntNumSeps(numOfItems, 0, numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return bINumOne, nil
	}

	// Numerator

	nUpperLimit := big.NewInt(0).Set(numOfItems)
	nLowerLimit := big.NewInt(0).Set(bigOne)

	if numOfItems.Cmp(numOfItemsPicked) == 0 {

		//                             n!
		// This is equivalent to   ------------
		//                             1

		result, err := NFactorial{}.CalcFactorialValueBigInt(nUpperLimit, nLowerLimit)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: fmt.Sprintf("upperLimit='%v' lowerLimit= '%v'", numOfItems, 1),
					ErrMessage: err.Error(),
				}
		}

		return result, nil

	}

	// numOfItems MUST BE GREATER THAN numOfItemsPicked

	nLowerLimit = big.NewInt(0).Sub(nUpperLimit, numOfItemsPicked)

	result, err := NFactorial{}.CalcFactorialValueBigInt(nUpperLimit, nLowerLimit)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("upperLimit='%v' lowerLimit= '%v'",
					nUpperLimit.Text(10), nLowerLimit.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	err = result.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "Error setting final Numeric Separators",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}

// permutationsWithRepsBigInt
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. This calculation assumes that REPETITIONS ARE ALLOWED.
//	This is also referred to as 'ordered sampling with replacement'.
//
//	Input Parameters
//	================
//
//	numOfItems               *big.Int
//	  Must be a positive integer number greater than zero.
//
//	numOfItemsPicked         *big.Int
//	  Must be a positive integer number greater than zero.
//
//	numSeps                  NumericSeparatorDto
//
//	  Numeric Separators consist of the decimal separator, thousands
//	  seprator, and currency symbol used to cofigure custom number
//	  strings.
//
//	  The new instance of BigIntNum returned by this method will
//	  contain a copy of the Numeric Separators contained in this
//	  input parameter (numSeps).
//
//	  If input parameter 'numSeps' is unconfigured, empty or
//	  invalid, an error will be returned.
//
//	  Be sure to call Probability.SetNumSeps() before you call this
//	  method, if Non-USA Numeric Separators are required.
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
//	BigIntNum
//	  If the calculation is successful, the result is returned as a
//	  BigIntNum type. If the calculation fails, the error return is
//	  populated.
//
//	error
//	  If processing errors are encountered, this returned error object
//	  will be configured with an appropriate error message.
//
//	Calculation
//	===========
//
//	In the following permutation formula, n= 'numOfItems'
//	and r = 'numOfItemsPicked'
//
//	           nPr  =  n^r
//
//	Where n is the number of things to choose from,
//	and we choose r of them, repetition is allowed,
//	and order matters.
//
//	This permutation calculation assumes REPETITIONS ARE ALLOWED!
//
//	This is also known as 'ordered sampling with replacement'. Since
//	this calculation allows repetitions or replacements, the
//	'numOfItemsPicked' may be greater than 'numOfItems'.
//
//	Example
//	=======
//
//	Combination theLock with 3 numbers: there are 10 numbers to choose from
//	(0,1,2,3,4,5,6,7,8,9) and we choose 3 of them (repetitions allowed):
//
//	      10 × 10 × ... (3 times) = 10^3 = 1,000 permutations
func (probPermutnsMech *probabilityPermutsMechanics) permutationsWithRepsBigInt(
	numOfItems *big.Int,
	numOfItemsPicked *big.Int,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if probPermutnsMech.lock == nil {
		probPermutnsMech.lock = new(sync.Mutex)
	}

	probPermutnsMech.lock.Lock()

	defer probPermutnsMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"probabilityPermutsMechanics.permutationsWithRepsBigInt()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if numOfItems == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numOfItems'",
			}
	}

	if numOfItemsPicked == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numOfItemsPicked'",
			}
	}

	err = numSeps.IsValid("Validating 'numSeps")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(\"Validating 'numSeps\")",
				ErrContext: "Input Parameter 'numSeps' is invalid",
				ErrMessage: err.Error(),
			}
	}

	bigZero := big.NewInt(0)

	if numOfItems.Cmp(bigZero) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
			}
	}

	if numOfItemsPicked.Cmp(bigZero) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
			}
	}

	if numOfItems.Cmp(bigZero) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
			}
	}

	if numOfItemsPicked.Cmp(bigZero) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
			}
	}

	bigOne := big.NewInt(1)

	if numOfItemsPicked.Cmp(bigOne) == 0 {

		bINumOne, err := new(BigIntNum).NewBigInt(numOfItems, 0)

		if err != nil {
			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bINumOne, err := new(BigIntNum).NewBigInt(numOfItems, 0)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return bINumOne, nil
	}

	resultBigInt := big.NewInt(0).Exp(numOfItems, numOfItemsPicked, nil)

	result, err := new(BigIntNum).NewBigIntNumSeps(resultBigInt, 0, numSeps)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "result, err := new(BigIntNum).\n" +
					"  NewBigInt(resultBigInt, 0, numSeps)",
				ErrContext: fmt.Sprintf("resultBigInt= '%v'\n"+
					"numSeps = '%v'\n",
					resultBigInt.Text(10), numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}
