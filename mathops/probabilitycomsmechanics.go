package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type probabilityComsMechanics struct {
	lock *sync.Mutex
}

// combinationsNoRepsBigInt
//
//	Calculates the number of combinations from 'numOfItems' and
//	'numOfItemsChosen' or 'n' things chosen 'r' at a time with NO
//	repetitions and order does NOT matter. This is also referred to
//	as an unordered sampling WITHOUT replacement. The calculation
//	result is returned as a BigIntNum type.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators consist of the decimal separator, thousands
//	seprator, and currency symbol used to cofigure custom number
//	strings.
//
//	The new instance of BigIntNum returned by this method will
//	contain a copy of the Numeric Separators currently configured
//	for this instance of type Probability (Probability.NumSeps). If
//	Probability.NumSeps is unconfigured or empty, it will be
//	automatically set to USA default values (decimal separator ('.'),
//	thousands separator (','), currency symbol ('$')). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
//
//	Input Parameters
//	================
//
//	numOfItems               *big.Int
//	  Must be a positive integer number greater than zero.
//	  'numOfItems' must be greater than or equal to
//	  'numOfItemsChosen'.
//
//	numOfItemsChosen         *big.Int
//	  Must be a positive integer number greater than zero.
//	 'numOfItemsChosen' must be less than or equal to 'numOfItems'.
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
//	  If input parameter 'numSeps' is unconfigured or empty, an error
//	  will be returned.
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
//	 If the calculation is successful, the result is returned as a
//		BigIntNum type. If the calculation fails, the error return is
//		populated.
//
//	error
//	 If an error is encountered during processing, this returend
//	 'error' object will be configured with an appropriate error
//	 message.
//
//	Calculation
//	===========
//
//	The calculation performed by this method uses the following
//	combinations formula, n= 'numOfItems' and r = 'numOfItemsChosen'.
//
//	                       n!
//	           nCr  =  -----------
//	                   (n-r)! r!
//
//	--------------------------------------------------
//
//	Where n is the number of things to choose from,
//	and we choose r of them, repetition is NOT allowed,
//	and order does NOT matter.
//
//	*** This calculation assumes NO REPETITIONS! ***
//
//	(a.k.a. as unordered sampling WITHOUT replacement)
//
//	Example WITHOUT Repetitions
//	===========================
//
//	There are 16 pool balls. How many ways to choose 3 pool balls with NO repetitions
//	or No repeats.
//
//	               16!             16 x 15 x 14
//	16C3 =   -------------   =    --------------  =  560
//	          (16-3)! x 3!             3 x 2
//
//	Note:  0! = 1
func (probComsMech *probabilityComsMechanics) combinationsNoRepsBigInt(
	numOfItems *big.Int,
	numOfItemsChosen *big.Int,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if probComsMech.lock == nil {
		probComsMech.lock = new(sync.Mutex)
	}

	probComsMech.lock.Lock()

	defer probComsMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"probabilityComsMechanics.combinationsNoRepsBigInt()",
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

	if numOfItemsChosen == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numOfItemsChosen'",
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

	if numOfItems.Cmp(bigZero) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("numOfItems= '%v'",
					numOfItems.Text(10)),
				ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
			}
	}

	if numOfItemsChosen.Cmp(bigZero) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("numOfItemsChosen= '%v'",
					numOfItemsChosen.Text(10)),
				ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
			}
	}

	if numOfItems.Cmp(numOfItemsChosen) < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: 'numOfItems' is LESS THAN 'numOfItemsChosen'!",
			}
	}

	if numOfItemsChosen.Cmp(bigZero) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: 'numOfItemsChosen' is ZERO!",
			}
	}

	bigOne := big.NewInt(1)

	if numOfItemsChosen.Cmp(bigOne) == 0 {

		biNumOfItems, err := new(BigIntNum).NewBigIntNumSeps(numOfItems, 0, numSeps)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "biNumOfItems, err := new(BigIntNum).\n" +
						"NewBigIntNumSeps(numOfItems, 0, numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return biNumOfItems, nil
	}

	numeratorNUpperLimit := big.NewInt(0).Set(numOfItems)

	numeratorNLowerLimit := big.NewInt(0).Set(bigOne)

	nMinusR := big.NewInt(0).Sub(numOfItems, numOfItemsChosen)

	if nMinusR.Cmp(bigZero) == 0 {

		nMinusR = big.NewInt(0).Set(bigOne)
	}

	var nFactorial BigIntNum

	var rFactorial BigIntNum

	if nMinusR.Cmp(numOfItemsChosen) == 1 {
		// nMinusR is Greater Than numOfItemsChosen
		numeratorNLowerLimit = big.NewInt(0).Set(nMinusR)

		nFactorial, err = NFactorial{}.CalcFactorialValueBigInt(
			numeratorNUpperLimit, numeratorNLowerLimit)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nFactorial, err = NFactorial{}.CalcFactorialValueBigInt(\n" +
						"  numeratorNUpperLimit, numeratorNLowerLimit)",
					ErrContext: fmt.Sprintf("numeratorNUpperLimit= '%v'\n"+
						"numeratorNLowerLimit=nMinusR= '%v'",
						numeratorNUpperLimit.Text(10), numeratorNLowerLimit.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		rUpperLimit := big.NewInt(0).Set(numOfItemsChosen)

		rLowerLimit := big.NewInt(0).Set(bigOne)

		rFactorial, err = NFactorial{}.CalcFactorialValueBigInt(
			rUpperLimit, rLowerLimit)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "rFactorial, err = NFactorial{}.CalcFactorialValueBigInt(\n" +
						"  rUpperLimit, rLowerLimit)",
					ErrContext: fmt.Sprintf("rUpperLimit= '%v'\n"+
						"rLowerLimit= '%v'",
						rUpperLimit.Text(10), rLowerLimit.Text(10)),
					ErrMessage: err.Error(),
				}
		}

	} else {
		// nMinusR is Less Than OR Equal to numOfItemsChosen

		numeratorNLowerLimit = big.NewInt(0).Set(numOfItemsChosen)

		nFactorial, err = NFactorial{}.CalcFactorialValueBigInt(
			numeratorNUpperLimit, numeratorNLowerLimit)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nFactorial, err = NFactorial{}.CalcFactorialValueBigInt(\n" +
						"  numeratorNUpperLimit, numeratorNLowerLimit)",
					ErrContext: fmt.Sprintf("numeratorNUpperLimit= '%v'\n"+
						"numeratorNLowerLimit= '%v'",
						numeratorNUpperLimit.Text(10), numeratorNLowerLimit.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		rFactorial, err = NFactorial{}.CalcFactorialValueBigInt(
			nMinusR, big.NewInt(0).Set(bigOne))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "rFactorial, err = NFactorial{}.CalcFactorialValueBigInt(\n" +
						"  nMinusR, big.NewInt(0).Set(bigOne))",
					ErrContext: fmt.Sprintf("UpperLimit=nMinusR= '%v'",
						nMinusR.Text(10)),
					ErrMessage: err.Error(),
				}
		}
	}

	nFactorialNumStr, err := nFactorial.GetNumStr()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nFactorialNumStr, err := nFactorial.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	rFactorialNumStr, err := rFactorial.GetNumStr()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "rFactorialNumStr, err := rFactorial.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	combinationsResult, err :=
		new(BigIntMathDivide).BigIntNumFracQuotient(nFactorial, rFactorial, numSeps, 10)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "combinationsResult, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n" +
					"  nFactorial, rFactorial, numSeps, 10)",
				ErrContext: fmt.Sprintf("nFactorial= '%v'\n"+
					"rFactorial= '%v'",
					nFactorialNumStr, rFactorialNumStr),
				ErrMessage: err.Error(),
			}
	}

	return combinationsResult, nil
}

// combinationsWithRepsBigInt
//
//	Calculates the number of combinations from 'numOfItems' and
//	'numOfItemsChosen' or 'n' things chosen 'r' at a time WITH
//	repetitions and order does NOT matter. This is also referred to as
//	an unordered sampling WITH replacement. The calculation result is
//	returned as a BigIntNum type.
//
//	Input Parameters
//	================
//
//	numOfItems               *big.Int
//	  Must be a positive integer number greater than zero.
//
//	numOfItemsChosen         *big.Int
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
//	  If input parameter 'numSeps' is unconfigured or empty, an error
//	  will be returned.
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
//	  If the calculation is successful, this return value is 'nil'.
//	  If the calculation fails, 'error' is populated with an
//	  appropriate error message.
//
//	Calculation
//	===========
//
//	The calculation performed by this method uses the following
//	combinations formula, n= 'numOfItems' and r = 'numOfItemsChosen'.
//	Since this calculation applies to unordered sampling WITH
//	replacement, 'numOfItemsChosen' may exceed 'numOfItems'.
//
//	                     (r + n - 1)!
//	         nCr  =  -------------------
//	                       r! (n-1)!
//
//	Example WITH Repetitions
//	========================
//
//	Let us say there are five flavors of icecream:
//	(1) banana, (2) chocolate, (3) lemon, (4) strawberry and (5) vanilla.
//
//	We can have three scoops. How many variations will there be?
//	nCr = (n==5 r==3) = 5C3 = Answer: 35
//
//	Note: 0! = 1
func (probComsMech *probabilityComsMechanics) combinationsWithRepsBigInt(
	numOfItems *big.Int,
	numOfItemsChosen *big.Int,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if probComsMech.lock == nil {
		probComsMech.lock = new(sync.Mutex)
	}

	probComsMech.lock.Lock()

	defer probComsMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"probabilityComsMechanics.combinationsWithRepsBigInt()",
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

	if numOfItemsChosen == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numOfItemsChosen'",
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

	bigOne := big.NewInt(1)

	// If 'numOfItemsChosen' == 1, result is always equal to 'numOfItems'.
	if numOfItemsChosen.Cmp(bigOne) == 0 {

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

	temp1 := big.NewInt(0).Add(numOfItems, numOfItemsChosen)

	numeratorUpperLimit := big.NewInt(0).Sub(temp1, bigOne)

	numeratorLowerLimit := big.NewInt(0).Sub(numOfItems, bigOne)

	nFactorial, err := NFactorial{}.CalcFactorialValueBigInt(
		numeratorUpperLimit, numeratorLowerLimit)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nFactorial, err = NFactorial{}.CalcFactorialValueBigInt(\n" +
					"  numeratorUpperLimit, numeratorLowerLimit)",
				ErrContext: fmt.Sprintf("numeratorUpperLimit= '%v'\n"+
					"numeratorLowerLimit= '%v'",
					numeratorUpperLimit.Text(10), numeratorLowerLimit.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	rUpperLimit := big.NewInt(0).Set(numOfItemsChosen)

	rLowerLimit := big.NewInt(1)

	rFactorial, err := NFactorial{}.CalcFactorialValueBigInt(
		rUpperLimit, rLowerLimit)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "rFactorial, err = NFactorial{}.CalcFactorialValueBigInt(\n" +
					"  rUpperLimit, rLowerLimit)",
				ErrContext: fmt.Sprintf("rUpperLimit= '%v'\n"+
					"rLowerLimit= '%v'",
					rUpperLimit.Text(10), rLowerLimit.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	nFactorialNumStr, err := nFactorial.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nFactorialNumStr, err := nFactorial.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	rFactorialNumStr, err := rFactorial.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "rFactorialNumStr, err := rFactorial.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	combinationsResult, err :=
		new(BigIntMathDivide).BigIntNumFracQuotient(nFactorial, rFactorial, numSeps, 10)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "combinationsResult, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n" +
					"  nFactorial, rFactorial, numSeps, 10)",
				ErrContext: fmt.Sprintf("nFactorial= '%v'\n"+
					"rFactorial= '%v'",
					nFactorialNumStr, rFactorialNumStr),
				ErrMessage: err.Error(),
			}
	}

	return combinationsResult, nil
}
