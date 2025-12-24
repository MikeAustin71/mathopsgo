package mathops

import (
  "fmt"
  "math/big"

  ePref "github.com/MikeAustin71/errpref"
)

/*

This file contains Probability source code for 'Permutations'.

For 'Combinations' see source code file:
		MikeAustin71\mathopsgo\mathops\probabilitycombinations.go

														Permutations
  													============

Reference
#########

	https://www.mathsisfun.com/combinatorics/combinations-permutations.html
	https://www.youtube.com/watch?v=XqQTXW7XfYA&list=PL06A16C388F14E6FE&index=21
	https://www.probabilitycourse.com/courses.php

$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
  *** This permutation calculation assumes NO REPETITIONS! ***
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

 In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'

              						n!
 							nPr	 =		------
 												(n-r)!
             ----------------------

 				Where n is the number of things to choose from,
				and we choose r of them, repetition is NOT allowed,
 				and order matters.


$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
 *** This Permutation calculation assumes REPETITIONS ARE ALLOWED! ***
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'


 									nPr	 =		n^r
                  -------------

 				Where n is the number of things to choose from,
				and we choose r of them, repetition is allowed,
 				and order matters.


*/

// PermutationsNoRepsBigInt
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
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
//	  'numOfItems' must be greater than or equal to 'numOfItemsPicked'.
//
//	numOfItemsPicked         *big.Int
//	  Must be a positive integer number greater than zero.
//	  'numOfItemsChosen' must be less than or equal to 'numOfItems'.
//
//	Returns
//	=======
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
func (prob *Probability) PermutationsNoRepsBigInt(
  numOfItems *big.Int, numOfItemsPicked *big.Int) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsNoRepsBigInt",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  return new(probabilityPermutsMechanics).permutationsNoRepsBigInt(
    numOfItems,
    numOfItemsPicked,
    numSeps,
    ePrefix)
}

// PermutationsWithRepsBigInt
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. This calculation assumes that REPETITIONS ARE ALLOWED.
//	This is also referred to as 'ordered sampling with replacement'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
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
//
//	numOfItemsPicked         *big.Int
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
func (prob *Probability) PermutationsWithRepsBigInt(
  numOfItems *big.Int, numOfItemsPicked *big.Int) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsNoRepsBigInt",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  return new(probabilityPermutsMechanics).permutationsWithRepsBigInt(
    numOfItems,
    numOfItemsPicked,
    numSeps,
    ePrefix)
}

// PermutationsBigIntNum
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type BigIntNum. Both input parameters must be
//	non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the permutation will therefore
//	vary depending on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                    n!
//	          nPr  =  ------
//	                  (n-r)!
//
//	              Note: 0! = 1
//
//	Where n is the number of things to choose from,
//	and we choose r of them, repetition is NOT allowed,
//	and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be
//	a positive integer number which is less than or equal to
//	'numOfItems'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	                nPr  =  n^r
//
//	Where n is the number of things to choose from,
//	and we choose r of them, repetition IS allowed,
//	and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive integer number.
//	'numOfItemsPicked' can be greater than, equal to or less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsBigIntNum(
  numOfItems, numOfItemsPicked BigIntNum,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsBigIntNum",
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

  numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedIsZero {

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

  numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedtSignValue == -1 {

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

  numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!",
      }
  }

  numOfItemsVsPickedCmp, err := numOfItems.Cmp(numOfItemsPicked)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsVsPickedCmp, err := numOfItems.Cmp(numOfItemsPicked)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && numOfItemsVsPickedCmp < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'!",
      }
  }

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).permutationsNoRepsBigInt(
      numOfItems.bigInt,
      numOfItemsPicked.bigInt,
      numSeps,
      ePrefix)

  }

  return new(probabilityPermutsMechanics).permutationsWithRepsBigInt(
    numOfItems.bigInt,
    numOfItemsPicked.bigInt,
    numSeps,
    ePrefix)
}

// PermutationsDecimal
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type Decimal. Both input parameters must be non-zero,
//	positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which will
//	determine whether the calculation results will allow repetitions or
//	not. The formula for the permutation will therefore vary depending
//	on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	Decimal.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on whether
//	input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                    n!
//	          nPr  =  ------
//	                  (n-r)!
//
//								Note: 0! = 1
//
//	Where n is the number of things to choose from,
//	and we choose r of them, repetition is NOT allowed,
//	and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	          'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	Where n is the number of things to choose from,
//	and we choose r of them, repetition IS allowed,
//	and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators consist of the decimal separator, thousands
//	seprator, and currency symbol used to cofigure custom number
//	strings.
//
//	The new instance of Decimal returned by this method will
//	contain a copy of the Numeric Separators currently configured
//	for this instance of type Probability (Probability.NumSeps). If
//	Probability.NumSeps is unconfigured or empty, it will be
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsDecimal(
  numOfItems, numOfItemsPicked Decimal,
  allowRepetitions bool) (Decimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsDecimal",
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

  numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()

  if err != nil {
    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedtSignValue == -1 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()

  if err != nil {
    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedIsZero {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
  }

  numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()

  if err != nil {
    return Decimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numOfItemsPickedPrecisionUint, err :=\n" +
          "numOfItemsPicked.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedPrecisionUint > 0 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!",
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

  r, err := numOfItemsPicked.GetBigInt()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsPicked.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && r.Cmp(n) == 1 {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsPicked= '%v'",
          numOfItemsNumStr, numOfItemsPickedNumStr),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return Decimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    result, err = new(probabilityPermutsMechanics).permutationsNoRepsBigInt(
      n, r, numSeps, ePrefix)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            "  permutationsNoRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(probabilityPermutsMechanics).
      permutationsWithRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return Decimal{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            " permutationsWithRepsBigInt(n, r, numSeps, ePrefix)",
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

// PermutationsIntAry
//
//	 Calculates the number of permutations associated with a collection
//	 of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	 significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	 are passed as type IntAry. Both input parameters must be non-zero,
//	 positive integer numbers.
//
//	 The input parameter 'allowRepetitions' is a boolean value which
//	 will determine whether the calculation results will allow
//	 repetitions or not. The formula for the permutation will therefore
//	 vary depending on whether repetitions are allowed.
//
//	 The result of this permutation calculation is returned as a type
//	 IntAry.
//
//	 In the following permutation formulas, n= 'numOfItems' and
//	 r = 'numOfItemsPicked'. The actual formula applied depends on
//	 whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//		========================================================================
//
//		          'allowRepetitions' = false
//
//		========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	               Note: 0! = 1
//
//	       Where n is the number of things to choose from,
//	       and we choose r of them, repetition is NOT allowed,
//	       and order matters.
//
//	 When 'allowRepetitions' = false, 'numOfItemsPicked' must be a
//	 positive integer number which is less than or equal to
//	 'numOfItems'.
//
//		========================================================================
//
//		          'allowRepetitions' = true
//
//		========================================================================
//
//	             nPr  =  n^r
//
//	       Where n is the number of things to choose from,
//	       and we choose r of them, repetition IS allowed,
//	       and order matters.
//
//	 When 'allowRepetitions' = true, 'numOfItemsPicked' must be a
//	 positive integer number. 'numOfItemsPicked' can be greater than,
//	 equal to or less than 'numOfItems'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators consist of the decimal separator, thousands
//	seprator, and currency symbol used to cofigure custom number
//	strings.
//
//	The new instance of IntAry returned by this method will
//	contain a copy of the Numeric Separators currently configured
//	for this instance of type Probability (Probability.NumSeps). If
//	Probability.NumSeps is unconfigured or empty, it will be
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsIntAry(
  numOfItems, numOfItemsPicked IntAry,
  allowRepetitions bool) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsIntAry",
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

  numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()

  if err != nil {
    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedtSignValue == -1 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()

  if err != nil {
    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedIsZero {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
  }

  numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()

  if err != nil {
    return IntAry{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numOfItemsPickedPrecisionUint, err :=\n" +
          "numOfItemsPicked.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedPrecisionUint > 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!",
      }
  }

  n, err := numOfItems.GetBigInt()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n, err := numOfItems.GetBigInt()",
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

  r, err := numOfItemsPicked.GetBigInt()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsPicked.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && r.Cmp(n) == 1 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems='%v'\nnumOfItemsPicked='%v'",
          numOfItemsNumStr, numOfItemsPickedNumStr),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    result, err = new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            "  permutationsNoRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    // result, err = new(Probability).PermutationsWithRepsBigInt(n, r)

    result, err = new(probabilityPermutsMechanics).
      permutationsWithRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return IntAry{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            "  permutationsWithRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  iaResult, err := result.GetIntAry()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error retruned on final calculation result.",
        ErrMessage: err.Error(),
      }
  }

  return iaResult, nil
}

// PermutationsINumMgr
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type INumMgr. Both input parameters must be non-zero,
//	positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which will
//	determine whether the calculation results will allow repetitions or
//	not. The formula for the permutation will therefore vary depending
//	on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on whether
//	input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                    n!
//	          nPr  =  ------
//	                  (n-r)!
//
//	               Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	          nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsINumMgr(
  numOfItems, numOfItemsPicked INumMgr,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsINumMgr",
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

  numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedtSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
  }

  numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!",
      }
  }

  n, err := numOfItems.GetBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "n, err := numOfItems.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
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

  r, err := numOfItemsPicked.GetBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsPicked.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && r.Cmp(n) == 1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsPicked= '%v'",
          numOfItemsNumStr, numOfItemsPickedNumStr),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    result, err = new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            " permutationsNoRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(probabilityPermutsMechanics).
      permutationsWithRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            "  permutationsWithRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  return result, nil
}

// PermutationsInt
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type int. Both input parameters must be non-zero,
//	positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which will
//	determine whether the calculation results will allow repetitions or
//	not. The formula for the permutation will therefore vary depending
//	on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on whether
//	input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	                Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsInt(
  numOfItems, numOfItemsPicked int,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsInt",
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

  if numOfItemsPicked == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
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

  if !allowRepetitions && numOfItemsPicked > numOfItems {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems='%v'\nnumOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsPicked))

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)

  }

  return new(probabilityPermutsMechanics).
    permutationsWithRepsBigInt(n, r, numSeps, ePrefix)
}

// PermutationsInt32
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type int32. Both input parameters must be non-zero,
//	positive integer numbers. 'numOfItems' must be equal to or greater
//	than 'numOfItemsPicked'.
//
//	The input parameter 'allowRepetitions' is a boolean value which will
//	determine whether the calculation results will allow repetitions or
//	not. The formula for the permutation will therefore vary depending
//	on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on whether
//	input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                   n!
//	         nPr  =  ------
//	                 (n-r)!
//
//	             Note: 0! = 1
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, repetition is NOT allowed,
//	     and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	         nPr  =  n^r
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, repetition is allowed,
//	     and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsInt32(
  numOfItems, numOfItemsPicked int32,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsInt32",
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

  if numOfItemsPicked == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
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

  if numOfItemsPicked < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  if !allowRepetitions && numOfItemsPicked > numOfItems {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems='%v'\nnumOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsPicked))

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)
  }

  return new(probabilityPermutsMechanics).
    permutationsWithRepsBigInt(n, r, numSeps, ePrefix)
}

// PermutationsInt64
//
//	Calculates the number of permutations associated with a collection of
//	'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
//	parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int64. Both input parameters
//	must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
//	'numOfItemsPicked'.
//
//	The input parameter 'allowRepetitions' is a boolean value which will determine whether the
//	calculation results will allow repetitions or not. The formula for the permutation will
//	therefore vary depending on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
//	formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                   n!
//	         nPr  =  ------
//	                 (n-r)!
//
//	           Note: 0! = 1
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, repetition is NOT allowed,
//	     and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	         nPr  =  n^r
//
//	     Where n is the number of things to choose from,
//	     and we choose r of them, repetition is allowed,
//	     and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsInt64(
  numOfItems, numOfItemsPicked int64,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsInt64",
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

  if numOfItemsPicked == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
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

  if !allowRepetitions && numOfItemsPicked > numOfItems {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems='%v'\nnumOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  n := big.NewInt(numOfItems)

  r := big.NewInt(numOfItemsPicked)

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)
  }

  return new(probabilityPermutsMechanics).
    permutationsWithRepsBigInt(n, r, numSeps, ePrefix)
}

// PermutationsNumStrDto
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type NumStrDto. Both input parameters must be non-zero,
//	positive integer numbers. 'numOfItems' must be equal to or greater
//	than 'numOfItemsPicked'.
//
//	The input parameter 'allowRepetitions' is a boolean value which will
//	determine whether the calculation results will allow repetitions or
//	not. The formula for the permutation will therefore vary depending on
//	whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	NumStrDto.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on whether
//	input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	                Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators consist of the decimal separator, thousands
//	seprator, and currency symbol used to cofigure custom number
//	strings.
//
//	The new instance of NumStrDto returned by this method will
//	contain a copy of the Numeric Separators currently configured
//	for this instance of type Probability (Probability.NumSeps). If
//	Probability.NumSeps is unconfigured or empty, it will be
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsNumStrDto(
  numOfItems, numOfItemsPicked NumStrDto,
  allowRepetitions bool) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsNumStrDto",
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

  numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedtSignValue, err := numOfItemsPicked.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedtSignValue == -1 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedIsZero, err := numOfItemsPicked.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedIsZero {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
  }

  numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedPrecisionUint, err := numOfItemsPicked.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if numOfItemsPickedPrecisionUint > 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!",
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

  r, err := numOfItemsPicked.GetBigInt()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "r, err := numOfItemsPicked.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numOfItemsPickedNumStr, err := numOfItemsPicked.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // r > n
  if !allowRepetitions && r.Cmp(n) == 1 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItems= '%v'\nnumOfItemsPicked= '%v'",
          numOfItemsNumStr, numOfItemsPickedNumStr),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  var result BigIntNum

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    result, err = new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            " permutationsNoRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    result, err = new(probabilityPermutsMechanics).
      permutationsWithRepsBigInt(n, r, numSeps, ePrefix)

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, err = new(probabilityPermutsMechanics).\n" +
            " permutationsWithRepsBigInt(n, r, numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  numStrDtoResult, err := result.GetNumStrDto()

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numStrDtoResult, err := result.GetNumStrDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return numStrDtoResult, nil
}

// PermutationsNumberStr
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant.
//
//	Input parameters 'numOfItems' and 'numOfItemsPicked' are passed as
//	strings. These strings must be formatted as valid number strings.
//	Number strings may be prefixed by a plus (+) or minus (-) and must
//	consist of a string of numeric digits which may be delimited by a
//	thousand separator. If the numeric value is a fractional value,
//	the fractional digits must be preceded by a period ('.') or
//	decimal separator.
//
//	Both input parameters must be non-zero, positive integer numbers.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the permutation will
//	therefore vary depending on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	                Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsNumberStr(
  numOfItems, numOfItemsPicked string,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsNumberStr",
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
        ErrMessage: "Error: Input Parameter 'numOfItems' is an EMPTY string!",
      }
  }

  if numOfItemsPicked == "" {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input Parameter 'numOfItemsPicked' is an EMPTY string!",
      }
  }

  nBigINumOfItems, err := new(BigIntNum).NewNumStr(numOfItems)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "nBigINumOfItems, err := new(BigIntNum).\n" +
          "  NewNumStr(numOfItems)",
        ErrContext: fmt.Sprintf("numOfItems= '%v'", numOfItems),
        ErrMessage: err.Error(),
      }
  }

  rBigINumOfItemsPicked, err := new(BigIntNum).NewNumStr(numOfItemsPicked)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "rBigINumOfItemsPicked, err := new(BigIntNum).\n" +
          "  NewNumStr(numOfItemsPicked)",
        ErrContext: fmt.Sprintf("numOfItemsPicked= '%v'", numOfItemsPicked),
        ErrMessage: err.Error(),
      }
  }

  nBigINumOfItemsSignValue, err := nBigINumOfItems.GetSign()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nBigINumOfItemsSignValue, err :=  nBigINumOfItems.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nBigINumOfItemsSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is LESS THAN ZERO!",
      }
  }

  nBigINumOfItemsIsZero, err := nBigINumOfItems.IsZero()

  if err != nil {
    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nBigINumOfItemsIsZero, err :=  nBigINumOfItems.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nBigINumOfItemsIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is ZERO!!",
      }
  }

  nBigINumOfItemsPrecisionUint, err := nBigINumOfItems.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "nBigINumOfItemsPrecisionUint, err := \n" +
          "  nBigINumOfItems.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nBigINumOfItemsPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItems' is NOT an Integer!",
      }
  }

  rBigINumOfItemsPickedSignValue, err := rBigINumOfItemsPicked.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "rBigINumOfItemsPickedSignValue, err := \n" +
          "  rBigINumOfItemsPicked.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if rBigINumOfItemsPickedSignValue == -1 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("numOfItemsPicked='%v' ", numOfItemsPicked),
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!",
      }
  }

  rBigINumOfItemsPickedIsZero, err := rBigINumOfItemsPicked.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "rBigINumOfItemsPickedIsZero, err := \n" +
          "  rBigINumOfItemsPicked.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if rBigINumOfItemsPickedIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is ZERO!",
      }
  }

  rBigINumOfItemsPickedPrecisionUint, err := rBigINumOfItemsPicked.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "rBigINumOfItemsPickedPrecisionUint, err :=\n" +
          "  rBigINumOfItemsPicked.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if rBigINumOfItemsPickedPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!",
      }
  }

  nBigINumOfItemsCmprBigINumOfItemsPicked, err := nBigINumOfItems.Cmp(rBigINumOfItemsPicked)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "nBigINumOfItemsCmprBigINumOfItemsPicked, err := \n" +
          "  nBigINumOfItems.Cmp(rBigINumOfItemsPicked)",
        ErrContext: "Comparing nBigINumOfItems vs rBigINumOfItemsPicked",
        ErrMessage: err.Error(),
      }
  }

  if !allowRepetitions && nBigINumOfItemsCmprBigINumOfItemsPicked < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("allowRepetitions= 'false'\n"+
          "numOfItems='%v'\n"+
          "numOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'.",
      }
  }

  var resultBINum BigIntNum

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    resultBINum, err = new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(
        nBigINumOfItems.bigInt,
        rBigINumOfItemsPicked.bigInt,
        numSeps,
        ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "resultBINum, err = new(probabilityPermutsMechanics).\n" +
            "  permutationsNoRepsBigInt(\n" +
            "   nBigINumOfItems.bigInt, rBigINumOfItemsPicked.bigInt,\n" +
            "      numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  } else {

    resultBINum, err = new(probabilityPermutsMechanics).
      permutationsWithRepsBigInt(
        nBigINumOfItems.bigInt,
        rBigINumOfItemsPicked.bigInt,
        numSeps,
        ePrefix)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "resultBINum, err = new(probabilityPermutsMechanics)." +
            " permutationsWithRepsBigInt(\n" +
            "  nBigINumOfItems.bigInt, rBigINumOfItemsPicked.bigInt\n" +
            "      numSeps, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return resultBINum, nil
}

// PermutationsUint
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type uint. 'numOfItems' must be equal to or greater
//	than 'numOfItemsPicked'.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the permutation will
//	therefore vary depending on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n = 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	                Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsUint(
  numOfItems, numOfItemsPicked uint,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsUint",
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
        ErrMessage: "Error: 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsPicked == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: 'numOfItemsPicked' is ZERO!",
      }
  }

  if !allowRepetitions && numOfItemsPicked > numOfItems {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("allowRepetitions= 'false'\n"+
          "numOfItems='%v'\n"+
          "numOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsPicked))

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)
  }

  // allowRepetitions = 'true'
  return new(probabilityPermutsMechanics).
    permutationsWithRepsBigInt(n, r, numSeps, ePrefix)
}

// PermutationsUint32
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type uint32. 'numOfItems' must be equal to or
//	greater than 'numOfItemsPicked'.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the permutation will
//	therefore vary depending on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n= 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	                Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
func (prob *Probability) PermutationsUint32(
  numOfItems, numOfItemsPicked uint32,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsUint32",
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
        ErrMessage: "Error: 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsPicked == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: 'numOfItemsPicked' is ZERO!",
      }
  }

  if !allowRepetitions && numOfItemsPicked > numOfItems {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("allowRepetitions= 'false'\n"+
          "numOfItems='%v'\n"+
          "numOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsPicked))

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)
  }

  // allowRepetitions = 'true'
  return new(probabilityPermutsMechanics).
    permutationsWithRepsBigInt(n, r, numSeps, ePrefix)
}

// PermutationsUint64
//
//	Calculates the number of permutations associated with a collection
//	of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
//	significant. Input parameters 'numOfItems' and 'numOfItemsPicked'
//	are passed as type uint64. 'numOfItems' must be equal to or
//	greater than 'numOfItemsPicked'.
//
//	The input parameter 'allowRepetitions' is a boolean value which
//	will determine whether the calculation results will allow
//	repetitions or not. The formula for the permutation will
//	therefore vary depending on whether repetitions are allowed.
//
//	The result of this permutation calculation is returned as a type
//	BigIntNum.
//
//	In the following permutation formulas, n = 'numOfItems' and
//	r = 'numOfItemsPicked'. The actual formula applied depends on
//	whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//	========================================================================
//
//	          'allowRepetitions' = false
//
//	========================================================================
//
//	                     n!
//	           nPr  =  ------
//	                   (n-r)!
//
//	                Note: 0! = 1
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition is NOT allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = false, 'numOfItemsPicked' must be a positive
//	integer number which is less than or equal to 'numOfItems'.
//
//
//	========================================================================
//
//	         'allowRepetitions' = true
//
//	========================================================================
//
//	           nPr  =  n^r
//
//	      Where n is the number of things to choose from,
//	      and we choose r of them, repetition IS allowed,
//	      and order matters.
//
//	When 'allowRepetitions' = true, 'numOfItemsPicked' must be a positive
//	integer number. 'numOfItemsPicked' can be greater than, equal to or
//	less than 'numOfItems'.
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
//	automatically set to USA default values: decimal separator ('.'),
//	thousands separator (','), currency symbol ('$'). If other
//	Numeric Separators are required, set this instance of type
//	Probability using method Probability.SetNumSeps().
//
//	Be sure to call Probability.SetNumSeps() before you call this
//	method, if Non-USA Numeric Separators are required.
func (prob *Probability) PermutationsUint64(
  numOfItems, numOfItemsPicked uint64,
  allowRepetitions bool) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "Probability.PermutationsUint64",
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
        ErrMessage: "Error: 'numOfItems' is ZERO!",
      }
  }

  if numOfItemsPicked == 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: 'numOfItemsPicked' is ZERO!",
      }
  }

  if !allowRepetitions && numOfItemsPicked > numOfItems {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("allowRepetitions= 'false'\n"+
          "numOfItems='%v'\n"+
          "numOfItemsPicked='%v'",
          numOfItems, numOfItemsPicked),
        ErrMessage: "Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.",
      }
  }

  n := big.NewInt(int64(numOfItems))

  r := big.NewInt(int64(numOfItemsPicked))

  prob.NumSeps.SetDefaultsIfEmpty()

  numSeps, err := prob.NumSeps.CopyOut(true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := prob.NumSeps.CopyOut(true)",
        ErrContext: "Probability.NumSeps Copy Out FAILED!",
        ErrMessage: err.Error(),
      }

  }

  if !allowRepetitions {

    return new(probabilityPermutsMechanics).
      permutationsNoRepsBigInt(n, r, numSeps, ePrefix)
  }

  // allowRepetitions = 'true'
  return new(probabilityPermutsMechanics).
    permutationsWithRepsBigInt(n, r, numSeps, ePrefix)
}
