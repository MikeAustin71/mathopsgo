package mathops

import (
  "fmt"
  "math/big"

  ePref "github.com/MikeAustin71/errpref"
)

/*

This file contains Probability source code for 'Combinations'.

For 'Permutations' see source code file:
	MikeAustin71\mathopsgo\mathops\probabilitypermutations.go


                                Combinations
																============

		Reference:
    ##########
		https://www.mathsisfun.com/combinatorics/combinations-permutations.html
		https://www.youtube.com/watch?v=bCxMhncR7PU&list=PL06A16C388F14E6FE&index=22
		https://www.youtube.com/watch?v=ZcSSI6VY1kM
		https://www.probabilitycourse.com/courses.php

$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
  *** This combination calculation assumes NO REPETITIONS! ***
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

 	In the following combination formula, n= 'numOfItems'  and r = 'numOfItemsChosen'

  Formula WITHOUT Repetitions
  ---------------------------
																			 n!
											 nCr	 =		-----------
																	 (n-r)! r!

	Example WITHOUT Repetitions
  ---------------------------
  16 pool balls. How many ways to choose 3 pool balls with NO repetitions or No repeats.

                  16!            16 x 15 x 14
	16C3 =    -------------  =    --------------  =    560
            (16-3)! x 3!            3 x 2


$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
  *** This combination calculation assumes REPETITIONS ARE ALLOWED! ***
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

 	In the following combination formula, n= 'numOfItems'  and r = 'numOfItemsChosen'

  Formula WITH Repetitions
  ------------------------
																	(r + n - 1)!
											nCr    =  -------------------
																	  r! (n-1)!

	Example WITH Repetitions:
  -------------------------
	Let us say there are five flavors of icecream:
  (1) banana, (2) chocolate, (3) lemon, (4) strawberry and (5) vanilla.

	We can have three scoops. How many variations will there be?
         nCr = (n==5 r==3) = 5C3 = Answer: 35


*/

type Probability struct {
  NoOfTrials                   BigIntNum
  PossibleOutcomesPerTrial     BigIntNum
  NoSuccessfulOutcomesPerTrial BigIntNum
  TotPossibleOutcomes          BigIntNum
  TotSuccessfulOutcomes        BigIntNum
  PercentCertainty             BigIntNum
  NumSeps                      NumericSeparatorDto
}

// CombinationsNoRepsBigInt
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
//	contain a copy of the Numeric Separators current configured for
//	this instance of type Probability ( Probability.NumSeps ). If
//	Probability.NumSeps is unconfigured or empty, it will be
//	automatically to USA default values (decimal separator ('.'),
//	thousands separator (','), currency symbol ('$')). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
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
func (prob *Probability) CombinationsNoRepsBigInt(
  numOfItems, numOfItemsChosen *big.Int) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsNoRepsBigInt",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  bigZero := big.NewInt(0)
  numSeps := new(NumericSeparatorDto).NewUSADefaults()

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

    biNumOfItems, err := new(BigIntNum).NewBigInt(numOfItems, 0)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "biNumOfItems, err := new(BigIntNum).\n" +
            "NewBigInt(numOfItems, 0)",
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

// CombinationsWithRepsBigInt
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
func (prob *Probability) CombinationsWithRepsBigInt(
  numOfItems, numOfItemsChosen *big.Int) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsWithRepsBigInt",
    "")

  if err != nil {
    return BigIntNum{}, err
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

  if numOfItemsChosen.Cmp(bigZero) == 0 {
    return BigIntNum{},
      fmt.Errorf("%v\n"+
        "Error: 'numOfItemsChosen' is ZERO!\n",
        ePrefix)
  }

  bigOne := big.NewInt(1)

  numSeps := new(NumericSeparatorDto).NewUSADefaults()

  // If 'numOfItemsChosen' == 1, result is always equal to 'numOfItems'.
  if numOfItemsChosen.Cmp(bigOne) == 0 {

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

// CombinationsBigIntNum
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'BigIntNum'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether
//	repetitions are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a
//	type 'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or
//	'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                      n!
//	          nCr  =  -----------
//	                   (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	      (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers. In
//	addition, 'numOfItems' MUST be greater than or equal to
//	'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                     (r + n - 1)!
//	          nCr  =  -------------------
//	                      r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS ARE
//	ALLOWED!
//	         (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers.
//	'numOfItems' can be greater than, equal to or less than
//	'numOfItemsChosen'.
func (prob *Probability) CombinationsBigIntNum(
  numOfItems, numOfItemsChosen BigIntNum, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsBigIntNum",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  numOfItemsIsZero, err := numOfItems.IsZero()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsIsZero, err := numOfItems.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }

  }

  if numOfItemsChosenIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
  }

  numOfItemsSignValue, err := numOfItems.GetSign()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsSignValue, err := numOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }

  }

  if numOfItemsSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }

  }

  if numOfItemsChosenSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  numOfItemsVsChosenCmp, err := numOfItems.Cmp(numOfItemsChosen)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsCmpResult, err := numOfItems.Cmp(numOfItemsChosen)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && numOfItemsVsChosenCmp < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "allowRepetitions= 'false'",
        ErrMessage: "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'!",
      }
  }

  if !allowRepetitions {

    return new(Probability).CombinationsNoRepsBigInt(numOfItems.bigInt, numOfItemsChosen.bigInt)
  }

  return new(Probability).CombinationsWithRepsBigInt(numOfItems.bigInt, numOfItemsChosen.bigInt)
}

// CombinationsDecimal
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'Decimal'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'Decimal'.
//
//	In the following combination formulas, n= 'numOfItems'  and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                      n!
//	          nCr  =  -----------
//	                   (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes
//	NO REPETITIONS!
//
//	(a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers.
//	In addition, 'numOfItems' MUST be greater than or equal to
//	'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//
//	                    (r + n - 1)!
//	          nCr  =  -------------------
//	                      r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//
//	(a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers.
//	'numOfItems' can be greater than, equal to or less than
//	'numOfItemsChosen'.
func (prob *Probability) CombinationsDecimal(
  numOfItems, numOfItemsChosen Decimal, allowRepetitions bool) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsDecimal",
    "")

  if err != nil {
    return Decimal{}, err
  }

  numOfItemsSignValue, err := numOfItems.GetSign()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsSignValue, err := numOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsSignValue == -1 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  numOfItemsIsZero, err := numOfItems.IsZero()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsIsZero, err := numOfItems.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsIsZero {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPrecisionUint > 0 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  numOfItemsChosenSignValue, err := numOfItems.GetSign()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenSignValue, err := numOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenSignValue == -1 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenIsZero {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  numOfItemsChosenPrecisionUint, err := numOfItems.GetPrecisionUint()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numOfItemsChosenPrecisionUint, err := \n" +
          "  numOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenPrecisionUint > 0 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  n, err := numOfItems.GetBigInt()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n, err := numOfItems.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  r, err := numOfItemsChosen.GetBigInt()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsChosen.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsNumStr, err := numOfItems.GetNumStr()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsNumStr, err := numOfItems.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && n.Cmp(r) < 0 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\n"+
          "numOfItemsChosen= '%v'",
          numOfItemsNumStr, numOfItemsChosenNumStr),
        ErrMessage: "Error: 'numOfItems' is LESS THAN 'numOfItemsChosen'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "resultDecimal, err := result.GetDecimal()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return resultDecimal, nil
}

// CombinationsIntAry
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'IntAry'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'IntAry'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                      n!
//	          nCr  =  -----------
//	                  (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	(a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen' must both
//	be positive integer numbers. In addition, 'numOfItems' MUST be greater than or
//	equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                    (r + n - 1)!
//	          nCr  =  -------------------
//	                      r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//
//	(a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers.
//	'numOfItems' can be greater than, equal to or less than
//	'numOfItemsChosen'.
func (prob *Probability) CombinationsIntAry(
  numOfItems, numOfItemsChosen IntAry, allowRepetitions bool) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsIntAry",
    "")

  if err != nil {
    return IntAry{}, err
  }

  numOfItemsSignValue, err := numOfItems.GetSign()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsSignValue, err := numOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsSignValue == -1 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  numOfItemsIsZero, err := numOfItems.IsZero()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsIsZero, err := numOfItems.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsIsZero {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPrecisionUint > 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenSignValue == -1 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenIsZero {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenPrecisionUint > 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  n, err := numOfItems.GetBigInt()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n, err := numOfItems.GetBigInt()\n",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  r, err := numOfItemsChosen.GetBigInt()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsChosen.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsNumStr, err := numOfItems.GetNumStr()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsNumStr, err := numOfItems.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && r.Cmp(n) == 1 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItemsNumStr, numOfItemsChosenNumStr),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  resultIntAry, err := result.GetIntAry()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "resultIntAry, err := result.GetIntAry()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return resultIntAry, nil
}

// CombinationsINumMgr
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'INumMgr'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'INumMgr'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                      n!
//	          nCr  =  -----------
//	                   (n-r)! r!
//
//
//	       Where n is the number of things to choose from,
//	       and we choose r of them, order does NOT matter
//	       and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	(a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. In addition, 'numOfItems' MUST
//	be greater than or equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//
//	                     (r + n - 1)!
//	          nCr  =  -------------------
//	                       r! (n-1)!
//
//	       Where n is the number of things to choose from,
//	       and we choose r of them, order does NOT matter
//	       and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//	     (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. 'numOfItems' can be greater
//	than, equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsINumMgr(
  numOfItems, numOfItemsChosen INumMgr, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsINumMgr",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  numOfItemsSignValue, err := numOfItems.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsSignValue, err := numOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  numOfItemsIsZero, err := numOfItems.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsIsZero, err := numOfItems.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  n, err := numOfItems.GetBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n, err := numOfItems.GetBigInt()",
        ErrContext: "",
        ErrMessage: "",
      }
  }

  r, err := numOfItemsChosen.GetBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsChosen.GetBigInt()",
        ErrContext: "",
        ErrMessage: "",
      }
  }

  numOfItemsNumStr, err := numOfItems.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsNumStr, err := numOfItems.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItemsNumStr, numOfItemsChosenNumStr),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: "",
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: "",
        }
    }
  }

  return result, nil
}

// CombinationsInt
//
//	Calculates the number of combinations associated with a	collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'int'. Both input parameters
//	must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether
//	repetitions are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling
//	WITH replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems'  and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//		                    n!
//		        nCr  =  -----------
//		                 (n-r)! r!
//
//
//		     Where n is the number of things to choose from,
//		     and we choose r of them, order does NOT matter
//		     and repetition is NOT allowed.
//
//
//		This version of the combinations calculation assumes NO REPETITIONS!
//		        (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. In addition, 'numOfItems' MUST
//	be greater than or equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//		                    (r + n - 1)!
//		        nCr  =  -------------------
//		                     r! (n-1)!
//
//		     Where n is the number of things to choose from,
//		     and we choose r of them, order does NOT matter
//		     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS ARE
//	ALLOWED!
//
//		(a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. 'numOfItems' can be greater
//	than, equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsInt(
  numOfItems, numOfItemsChosen int, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsInt",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  if numOfItems == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsChosen < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  if numOfItemsChosen == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsChosen))

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  return result, nil
}

// CombinationsInt32
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'int32'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether
//	repetitions are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                      n!
//	          nCr  =  -----------
//	                   (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	      (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers. In
//	addition, 'numOfItems' MUST be greater than or equal to
//	'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//
//	                  (r + n - 1)!
//	          nCr = -------------------
//	                  r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//		   (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers.
//	'numOfItems' can be greater than, equal to or less than
//	'numOfItemsChosen'.
func (prob *Probability) CombinationsInt32(
  numOfItems, numOfItemsChosen int32, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsInt32",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'", numOfItems),
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  if numOfItems == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsChosen < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItemsChosen= '%v'", numOfItemsChosen),
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  if numOfItemsChosen == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsChosen))

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return result, nil
}

// CombinationsInt64
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'int32'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether
//	repetitions are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems'  and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                    n!
//	       nCr  =  -------------
//	                (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	     (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. In addition, 'numOfItems'
//	MUST be greater than or equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                    (r + n - 1)!
//	          nCr  =  -------------------
//	                      r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//	        (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be positive integer numbers.
//	'numOfItems' can be greater than, equal to or less than
//	'numOfItemsChosen'.
func (prob *Probability) CombinationsInt64(
  numOfItems, numOfItemsChosen int64, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsInt64",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'", numOfItems),
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  if numOfItems == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsChosen < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItemsChosen= '%v'", numOfItemsChosen),
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  if numOfItemsChosen == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  n := big.NewInt(numOfItems)

  r := big.NewInt(numOfItemsChosen)

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return result, nil
}

// CombinationsNumStrDto
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'NumStrDto'. Both input
//	parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed. 'allowRepetitions' == false signals unordered
//	sampling WITHOUT replacement. 'allowRepetitions' == true signals
//	unordered sampling WITH replacement.
//
//	The result of this combination calculation is returned as a type
//	'NumStrDto'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                      n!
//	          nCr  =  -----------
//	                   (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	        (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen' must both
//	be positive integer numbers. In addition, 'numOfItems' MUST be greater than or
//	equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//
//	               (r + n - 1)!
//	     nCr  =  -------------------
//	                r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS ARE
//	ALLOWED!
//	      (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. 'numOfItems' can be greater
//	than, equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsNumStrDto(
  numOfItems, numOfItemsChosen NumStrDto, allowRepetitions bool) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsNumStrDto",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numOfItemsSignValue, err := numOfItems.GetSign()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsSignValue, err := numOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsSignValue == -1 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  numOfItemsIsZero, err := numOfItems.IsZero()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsIsZero, err := numOfItems.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsIsZero {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPrecisionUint, err := numOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPrecisionUint > 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  numOfItemsChosenSignValue, err := numOfItemsChosen.GetSign()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numOfItemsChosenSignValue, err :=\n" +
          "  numOfItemsChosen.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenSignValue == -1 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenIsZero, err := numOfItemsChosen.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenIsZero {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenPrecisionUint, err := numOfItemsChosen.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsChosenPrecisionUint > 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  n, err := numOfItems.GetBigInt()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n, err := numOfItems.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  r, err := numOfItemsChosen.GetBigInt()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsChosen.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsNumStr, err := numOfItems.GetNumStr()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsNumStr, err := numOfItems.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsChosenNumStr, err := numOfItemsChosen.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && r.Cmp(n) == 1 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems='%v'\n"+
          "numOfItemsChosen='%v'",
          numOfItemsNumStr, numOfItemsChosenNumStr),
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  resultNumStrDto, err := result.GetNumStrDto()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "esultNumStrDto, err := result.GetNumStrDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return resultNumStrDto, nil
}

// CombinationsNumberStr
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant.
//
//	Input parameters 'numOfItems' and 'numOfItemsChosen' are passed as
//	strings. These strings must be formatted as valid number strings.
//	Number strings may be prefixed by a plus (+) or minus (-) and must
//	consist of a string of numeric digits which may be delimited by
//	the 'thousands' separator. If the numeric value is a fractional
//	value, the fractional digits must be preceded by a period ('.') or
//	decimal separator. However, for purposes of this calculation, both
//	input parameters must be passed as positive. integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                   n!
//	      nCr  =  -----------
//	               (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	    (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. In addition, 'numOfItems'
//	MUST be greater than or equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//
//	                   (r + n - 1)!
//	        nCr  =  -------------------
//	                    r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS ARE
//	ALLOWED!
//	      (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and 'numOfItemsChosen'
//	must both be positive integer numbers. 'numOfItems' can be greater
//	than, equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsNumberStr(
  numOfItems, numOfItemsChosen string, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsNumberStr",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems == "" {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is an EMPTY string!",
      }
  }

  if numOfItemsChosen == "" {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is an EMPTY string!",
      }
  }

  nBigIntNum, err := new(BigIntNum).NewNumStr(numOfItems)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "nBigIntNum, err := new(BigIntNum).\n" +
          "  NewNumStr(numOfItems)",
        ErrContext: fmt.Sprintf("  numOfItems= '%v'", numOfItems),
        ErrMessage: err.Error(),
      }
  }

  nBigIntNumIsZero, err := nBigIntNum.IsZero()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nBigIntNumIsZero, err := nBigIntNum.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nBigIntNumIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  nBigIntNumSignValue, err := nBigIntNum.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nBigIntNumSignValue, err := nBigIntNum.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nBigIntNumSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  nBigIntNumPrecisionUint, err := nBigIntNum.GetPrecisionUint()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "nBigIntNumPrecisionUint, err := \n" +
          "  nBigIntNum.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nBigIntNumPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  rBigIntNum, err := new(BigIntNum).NewNumStr(numOfItemsChosen)

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "rBigIntNum, err := new(BigIntNum).\n" +
          "  NewNumStr(numOfItemsChosen)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  rBigIntNumIsZero, err := rBigIntNum.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "rBigIntNumIsZero, err := rBigIntNum.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if rBigIntNumIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  rBigIntNumSignValue, err := rBigIntNum.GetSign()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "rBigIntNumSignValue, err := rBigIntNum.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if rBigIntNumSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is LESS THAN ZERO!",
      }
  }

  rBigIntNumPrecisionUint, err := rBigIntNum.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "rBigIntNumPrecisionUint, err := rBigIntNum.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if rBigIntNumPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  nBigIntNumVsRBigIntNumCmp, err := nBigIntNum.Cmp(rBigIntNum)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nBigIntNumVsRBigIntNum, err := nBigIntNum.Cmp(rBigIntNum)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && nBigIntNumVsRBigIntNumCmp < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems='%v' numOfItemsChosen='%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is NOT an Integer!",
      }
  }

  if !allowRepetitions {

    return new(Probability).CombinationsNoRepsBigInt(nBigIntNum.bigInt, rBigIntNum.bigInt)
  }

  return new(Probability).CombinationsWithRepsBigInt(nBigIntNum.bigInt, rBigIntNum.bigInt)
}

// CombinationsUint
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'uint'. Both input
//	parameters must be non-zero, integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                n!
//	    nCr  =  -----------
//	             (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	       (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be integer numbers. In addition, 'numOfItems' MUST be
//	greater than or	equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                (r + n - 1)!
//	      nCr  =  -------------------
//	                 r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//	       (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and 'numOfItemsChosen'
//	must both be integer numbers. 'numOfItems' can be greater than,
//	equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsUint(
  numOfItems, numOfItemsChosen uint, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsUint",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsChosen == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  n := big.NewInt(0).SetUint64(uint64(numOfItems))

  r := big.NewInt(0).SetUint64(uint64(numOfItemsChosen))

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return result, nil
}

// CombinationsUint32
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'uint32'. Both input
//	parameters must be non-zero, integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether
//	repetitions are allowed.
//
//	'allowRepetitions' == false signals unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	        nCr  =  -----------
//	                 (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	        (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be integer numbers. In addition, 'numOfItems' MUST be
//	greater than or equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                     (r + n - 1)!
//	          nCr  =  -------------------
//	                      r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//	          (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be integer numbers. 'numOfItems' can
//	be greater than, equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsUint32(
  numOfItems, numOfItemsChosen uint32, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsUint32",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsChosen == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  n := big.NewInt(0).SetUint64(uint64(numOfItems))

  r := big.NewInt(0).SetUint64(uint64(numOfItemsChosen))

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return result, nil
}

// CombinationsUint64
//
//	Calculates the number of combinations associated with a collection
//	of 'numOfItems' from which one chooses 'numOfItemsChosen'. Order
//	IS NOT significant. Input parameters 'numOfItems' and
//	'numOfItemsChosen' are passed as type 'uint64'. Both input
//	parameters must be non-zero, integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the calculation of
//	combinations will therefore vary depending on whether repetitions
//	are allowed.
//
//	'allowRepetitions' == false signals	unordered sampling WITHOUT
//	replacement.
//
//	'allowRepetitions' == true signals unordered sampling WITH
//	replacement.
//
//	The result of this combination calculation is returned as a type
//	'BigIntNum'.
//
//	In the following combination formulas, n= 'numOfItems' and
//	r = 'numOfItemsChosen'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                       n!
//	          nCr  =  -----------
//	                  (n-r)! r!
//
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition is NOT allowed.
//
//	This version of the combinations calculation assumes NO REPETITIONS!
//	       (a.k.a. as unordered sampling WITHOUT replacement)
//
//	When 'allowRepetitions' = false, 'numOfItems' and 'numOfItemsChosen'
//	must both be integer numbers. In addition, 'numOfItems' MUST be
//	greater than or equal to 'numOfItemsChosen'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                    (r + n - 1)!
//	         nCr  =  -------------------
//	                     r! (n-1)!
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, order does NOT matter
//	     and repetition IS NOT allowed.
//
//	This version of the combinations calculation assumes REPETITIONS
//	ARE ALLOWED!
//	       (a.k.a. as unordered sampling WITH replacement)
//
//	When 'allowRepetitions' = true, 'numOfItems' and
//	'numOfItemsChosen' must both be integer numbers. 'numOfItems' can
//	be greater than, equal to or less than 'numOfItemsChosen'.
func (prob *Probability) CombinationsUint64(
  numOfItems, numOfItemsChosen uint64, allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.CombinationsUint64",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if numOfItems == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsChosen == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsChosen' is ZERO!",
      }
  }

  n := big.NewInt(0).SetUint64(numOfItems)

  r := big.NewInt(0).SetUint64(numOfItemsChosen)

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsChosen= '%v'",
          numOfItems, numOfItemsChosen),
        ErrMessage: "Error: 'numOfItemsChosen' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  if !allowRepetitions {

    result, err = new(Probability).CombinationsNoRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsNoRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(Probability).CombinationsWithRepsBigInt(n, r)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = Probability{}.\n" +
            "  CombinationsWithRepsBigInt(n, r)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return result, nil
}
