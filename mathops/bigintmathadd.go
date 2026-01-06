package mathops

import (
	"fmt"
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

type BigIntMathAdd struct {
	Input  BigIntPair
	Result BigIntNum
}

// AddBigInts - Adds two *big.Int numbers. Each *big.Int number
// is passed to the method with an associated decimal place precision
// specification.
//
// The BigIntNum 'result' returned by this addition operation will contain
// USA default numeric separators (decimal separator, thousands separator and
// currency symbol).
func (bAdd *BigIntMathAdd) AddBigInts(
	b1 *big.Int,
	precision1 uint,
	b2 *big.Int,
	precision2 uint) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddBigInts()"

	// No error is possible because both precision parameters
	// are by definition, greater than or equal to zero.

	b1BigInNum, err := new(BigIntNum).NewBigIntBigPrecision(
		b1, big.NewInt(int64(precision1)))

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "b1BigInNum, err := new(BigIntNum).NewBigIntBigPrecision(\n" +
					"    b1, big.NewInt(int64(precision1)))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b2BigInNum, err := new(BigIntNum).NewBigIntBigPrecision(
		b2, big.NewInt(int64(precision2)))

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "b2BigInNum, err := new(BigIntNum).NewBigIntBigPrecision(\n" +
					"    b2, big.NewInt(int64(precision2)))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	biNum, err := new(BigIntMathAdd).AddBigIntNums(b1BigInNum, b2BigInNum)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "biNum, err := new(BigIntMathAdd).\n" +
					"    AddBigIntNums(b1BigInNum, b2BigInNum)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return biNum, nil
}

// AddBigIntNums - Adds two BigIntNums and returns the result in a new
// BigIntNum instance
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'b1'.
func (bAdd *BigIntMathAdd) AddBigIntNums(b1 BigIntNum, b2 BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddBigIntNums"

	bPair, err := new(BigIntPair).NewBigIntNum(b1, b2)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(b1, b2)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	addResult, err := bAdd.AddPair(bPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(b1, b2)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}
	return addResult, err
}

// AddBigIntNumArray - Adds an Array of 'BigIntNum' types and returns the result
// as type 'BigIntNum'
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from the first element of the bNums array (bNums[0]).
func (bAdd *BigIntMathAdd) AddBigIntNumArray(bNums []BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddBigIntNumArray()"

	finalResult, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "finalResult, err := new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenBNums := len(bNums)

	if lenBNums == 0 {

		return finalResult, nil

	}

	numSeps, err := bNums[0].GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
					"    big.NewInt(0).Set(bNum.bigInt), bNum.precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	for i := 0; i < lenBNums; i++ {

		if i == 0 {

			finalResult, err = bNums[i].CopyOut()

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = bNums[i].CopyOut()",
						ErrContext: fmt.Sprintf("bNums is type []BigIntNum. bNums[%v]", i),
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bNums[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bNums[i])",
					ErrContext: fmt.Sprintf("bNums is type []BigIntNum. bNums[%v]", i),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddBigIntNumOutputToArray - The first input parameter to this method
// is a BigIntNum Type labeled, 'addend'.  The second element is an
// array of BigIntNum types labeled 'bNums'. The 'addend' is added to
// each element of the 'bNums' array with the result output to another
// array of BigIntNum types ([]BigIntNum) which is returned to the calling
// function.
//
// Example
// =======
//
//											Multiplicands												Output
//	 Addend   				    	Array														Array
//
//			3			+					bNums[0] = 2			=				  outputarray[0] =  5
//			3			+					bNums[1] = 3			=				  outputarray[1] =  6
//			3			+					bNums[2] = 4			=				  outputarray[2] =  7
//			3			+					bNums[3] = 5			=				  outputarray[3] =  8
//			3			+					bNums[4] = 6			=				  outputarray[4] =  9
//			3			+					bNums[5] = 9			=				  outputarray[5] = 12
//
// Each element of the []BigIntNum 'result' array returned by this addition
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) which were copied from input parameter
// 'addend'.
func (bAdd *BigIntMathAdd) AddBigIntNumOutputToArray(
	addend BigIntNum,
	bNums []BigIntNum) ([]BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddBigIntNumOutputToArray()"

	lenBNums := len(bNums)

	if lenBNums == 0 {

		return []BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ErrContext: "if lenBNums == 0",
				ErrMessage: "lenBNums is 0",
			}

	}

	numSeps, err := addend.GetNumericSeparatorsDto()

	if err != nil {
		return []BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultArray := make([]BigIntNum, lenBNums)

	for i := 0; i < lenBNums; i++ {

		bPair, err := new(BigIntPair).NewBigIntNum(addend, bNums[i])

		if err != nil {
			return []BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(addend, bNums[i])",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		result, err := bAdd.addPairNoNumSeps(bPair)

		if err != nil {
			return []BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "result, err := bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		resultArray[i], err = result.CopyOut()

		if err != nil {
			return []BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "resultArray[i], err = result.CopyOut()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return resultArray, nil
}

// AddBigIntNumSeries - Adds a series of BigIntNum types and returns the total in a
// BigIntNum instance.
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from the first element in input series 'bNums'.
func (bAdd *BigIntMathAdd) AddBigIntNumSeries(bNums ...BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddBigIntNumSeries"

	finalResult := new(BigIntNum).New()

	numSeps := NumericSeparatorDto{}

	var err error

	for i, bNum := range bNums {

		if i == 0 {

			finalResult, err = bNum.CopyOut()

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
						ErrContext: "for i, bNum := range bNums",
						ErrMessage: err.Error(),
					}
			}

			numSeps, err = finalResult.GetNumericSeparatorsDto()

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "numSeps, err = finalResult.GetNumericSeparatorsDto()",
						ErrContext: "for i, bNum := range bNums",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bNum)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bNum)",
					ErrContext: "for i, bNum := range bNums",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "for i, bNum := range bNums",
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddDecimal - Receives two Decimal instances and adds their numeric values.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'dec1'.
//

func (bAdd *BigIntMathAdd) AddDecimal(dec1, dec2 Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto()"

	// This method tests the validity of dec1 and dec2
	bPair, err := new(BigIntPair).NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by BigIntPair{}.NewDecimal(dec1, dec2).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	finalResult, err := bAdd.AddPair(bPair)

	if err != nil {
		return finalResult,
			fmt.Errorf("%v\n"+
				"Error returned by bAdd.AddPair(bPair).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return finalResult, nil
}

// AddDecimalArray - Adds an array of 'Decimal' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from the first element of the input (decs[0]).
func (bAdd *BigIntMathAdd) AddDecimalArray(decs []Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddDecimalArray() "

	finalResult := new(BigIntNum).New()

	var err error

	lenDecs := len(decs)

	if lenDecs == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenDecs == 0 {",
				ErrMessage: "    lenDecs == 0\n" +
					"Length of input parameter 'decs' array is Zero.\n",
			}
	}

	numSeps, err := decs[0].GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := decs[0].GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSeps.SetDefaultsIfEmpty()

	var decNStr string

	for i := 0; i < lenDecs; i++ {

		err = decs[i].IsValid(fmt.Sprintf("Validating decs[%v]", i))

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = decs[i].IsValid(fmt.Sprintf(\"Validating decs[%v]\", i))",
					ErrContext: fmt.Sprintf("decs[%v] Failed Validation Tests", i),
					ErrMessage: err.Error(),
				}
		}

		if i == 0 {

			decNStr, err = decs[i].GetNumStr()

			if err != nil {
				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: fmt.Sprintf("decNStr, err = decs[%v].GetNumStr()", i),
						ErrContext: "i==0",
						ErrMessage: err.Error(),
					}
			}

			// This method tests the validity of decs[i]
			finalResult, err = new(BigIntNum).NewDecimal(decs[i])

			if err != nil {
				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: fmt.Sprintf("finalResult, err = new(BigIntNum).NewDecimal(decs[%v])", i),
						ErrContext: fmt.Sprintf("decs[%v].GetNumStr()='%v'", i, decNStr),
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		decNStr, err = decs[i].GetNumStr()

		if err != nil {
			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: fmt.Sprintf("decNStr, err = decs[%v].GetNumStr()", i),
					ErrContext: fmt.Sprintf("i== '%v'", i),
					ErrMessage: err.Error(),
				}
		}
		// This method tests the validity of decs[i]
		bigINumNextAddend, err := new(BigIntNum).NewDecimal(decs[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bigINumNextAddend, err := new(BigIntNum).NewDecimal(decs[i])",
					ErrContext: fmt.Sprintf("decs[%v] = %v", i, decNStr),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bigINumNextAddend)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bigINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddDecimalOutputToArray - The first input parameter to this method
// is a Decimal Type labeled, 'addend'.  The second element is an
// array of Decimal types labeled 'decs'. The 'addend' is added to
// each element of the 'decs' array with the result output to another
// array of Decimal types which is returned to the calling function.
//
// Example
// =======
//
//											    decs										 Output
//	 Addend   				    	Array											Array
//
//			3			+					decs[0] = 2			=				  outputarray[0] =  5
//			3			+					decs[1] = 3			=				  outputarray[1] =  6
//			3			+					decs[2] = 4			=				  outputarray[2] =  7
//			3			+					decs[3] = 5			=				  outputarray[3] =  8
//			3			+					decs[4] = 6			=				  outputarray[4] =  9
//			3			+					decs[5] = 9			=				  outputarray[5] = 12
//
// Each element in the []Decimal array 'result' returned by this addition
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) which were copied from input parameter
// 'addend'.
func (bAdd *BigIntMathAdd) AddDecimalOutputToArray(
	addend Decimal,
	decs []Decimal) ([]Decimal, error) {

	ePrefix := "BigIntMathAdd.AddDecimalOutputToArray()"

	lenDecs := len(decs)

	if lenDecs == 0 {

		return []Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenDecs == 0 {",
				ErrMessage: "Error: Input parameter 'decs' array is Empty!",
			}
	}

	numSeps, err := addend.GetNumericSeparatorsDto()

	if err != nil {

		return []Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultsArray := make([]Decimal, lenDecs)

	for i := 0; i < lenDecs; i++ {

		err = decs[i].IsValid(ePrefix + fmt.Sprintf("decs[%v]", i))

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = decs[i].IsValid(ePrefix + fmt.Sprintf(\"decs[%v]\", i))",
					ErrContext: fmt.Sprintf("decs[%v] Failed Validation Tests", i),
					ErrMessage: err.Error(),
				}
		}

		// This method tests the validity of addend and decs[i]
		bPair, err := new(BigIntPair).NewDecimal(addend, decs[i])

		if err != nil {

			decNStr, _ := decs[i].GetNumStr()

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewDecimal(addend, decs[%v])", i),
					ErrContext: fmt.Sprintf("dec[i].GetNumStr()='%v'", decNStr),
					ErrMessage: err.Error(),
				}
		}

		result, err := bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "result, err := bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		resultsArray[i], err = result.GetDecimal()

		if err != nil {

			dNumStr, _ := result.GetNumStr()

			return []Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: fmt.Sprintf("resultsArray[%v], err = result.GetDecimal()", i),
					ErrContext: fmt.Sprintf("decs[i].GetNumStr()='%v'", dNumStr),
					ErrMessage: err.Error(),
				}
		}
	}

	return resultsArray, nil
}

// AddDecimalSeries - Adds a series of 'Decimal' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The BigIntNum 'result' returned by this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) which
// were copied from input series element 'decs[0]'.
func (bAdd *BigIntMathAdd) AddDecimalSeries(decs ...Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddDecimalSeries()"

	finalResult := new(BigIntNum).New()
	var err error

	if len(decs) == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if len(decs) == 0 {",
				ErrMessage: "Error: Input parameter 'decs' series is Empty!",
			}
	}

	numSeps := NumericSeparatorDto{}

	for i, dec := range decs {

		err = dec.IsValid(ePrefix + fmt.Sprintf(" Testing decs[%v]", i))

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = dec.IsValid(ePrefix + fmt.Sprintf(\" Testing decs[%v]\", i))",
					ErrContext: fmt.Sprintf("decs[%v] Failed Validation Tests", i),
					ErrMessage: err.Error(),
				}
		}

		if i == 0 {

			// This method tests the validity of 'dec'
			finalResult, err = new(BigIntNum).NewDecimal(dec)

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewDecimal(dec)",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			numSeps, err = dec.GetNumericSeparatorsDto()

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "numSeps, err = dec.GetNumericSeparatorsDto()",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		bigINumNextAddend, err := dec.GetBigIntNum()

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bigINumNextAddend, err := dec.GetBigIntNum()",
					ErrContext: fmt.Sprintf("decs[%v] generated error.", i),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bigINumNextAddend)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bigINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bigINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddIntAry - Receives two IntAry instances and adds their numeric values.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) which
// were copied from input parameter 'ia1'.
func (bAdd *BigIntMathAdd) AddIntAry(ia1, ia2 IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto()"

	// This method will test the validity of ia1 and ia2
	bPair, err := new(BigIntPair).NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by BigIntPair{}.NewIntAry(ia1, ia2).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	finalResult, err := bAdd.AddPair(bPair)

	if err != nil {
		return finalResult,
			fmt.Errorf("%v\n"+
				"Error returned by bAdd.AddPair(bPair).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return finalResult, nil
}

// AddIntAryArray
// Receives an array of IntAry objects and totals their numeric values.
// The total numeric value is returned in a BigIntNum instance.
//
// The BigIntNum 'result' returned by this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) which
// were copied from the first element of the input array, iarys[0].
func (bAdd *BigIntMathAdd) AddIntAryArray(iarys []IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddIntAryArray()"

	finalResult := new(BigIntNum).New()

	lenIaArray := len(iarys)

	if lenIaArray == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenIaArray == 0 {",
				ErrMessage: "Error: Input parameter 'iarys' array is Empty!",
			}
	}

	var err error

	numSeps, err := iarys[0].GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(&finalResult, nums[i])",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	for i := 0; i < lenIaArray; i++ {

		err = iarys[i].IsValid(ePrefix + fmt.Sprintf(" Testing iarys[%v]", i))

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = iarys[i].IsValid(ePrefix + fmt.Sprintf(\" Testing iarys[%v]\", i))",
					ErrContext: fmt.Sprintf("iarys[%v] Failed Validation Testing.", i),
					ErrMessage: err.Error(),
				}
		}

		if i == 0 {

			// This method will test the validity of iarys[i]
			finalResult, err = new(BigIntNum).NewIntAry(iarys[i])

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewIntAry(iarys[i])",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		// This method will test the validity of iarys[i]
		bINumNextAddend, err := new(BigIntNum).NewIntAry(iarys[i])

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bINumNextAddend, err := new(BigIntNum).NewIntAry(iarys[i])",
					ErrContext: fmt.Sprintf("iarys[%v] generated error", i),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bINumNextAddend)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddIntAryOutputToArray - The first input parameter to this method
// is an IntAry Type labeled, 'addend'.  The second element is an
// array of IntAry types labeled 'iarys'. The 'addend' is added to
// each element of the 'iarys' array with the result output to another
// array of IntAry types which is returned to the calling function.
//
// Example
// =======
//
//											    decs										 Output
//	 Addend   				    	Array											Array
//
//			3			+					iarys[0] = 2			=				  outputarray[0] =  5
//			3			+					iarys[1] = 3			=				  outputarray[1] =  6
//			3			+					iarys[2] = 4			=				  outputarray[2] =  7
//			3			+					iarys[3] = 5			=				  outputarray[3] =  8
//			3			+					iarys[4] = 6			=				  outputarray[4] =  9
//			3			+					iarys[5] = 9			=				  outputarray[5] = 12
//
// Each element of the []IntAry returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'addend'.
func (bAdd *BigIntMathAdd) AddIntAryOutputToArray(
	addend IntAry,
	iarys []IntAry) ([]IntAry, error) {

	ePrefix := "BigIntMathAdd.AddIntAryOutputToArray() "

	// This method will test the validity of 'addend'
	bINumAddend, err := new(BigIntNum).NewIntAry(addend)

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bINumAddend, err := new(BigIntNum).NewIntAry(addend)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenIaArray := len(iarys)

	if lenIaArray == 0 {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenIaArray == 0 {",
				ErrMessage: "Error: Input parameter 'iarys' array is Empty!",
			}
	}

	numSeps, err := addend.GetNumericSeparatorsDto()

	if err != nil {

		return []IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultsArray := make([]IntAry, lenIaArray)

	for i := 0; i < lenIaArray; i++ {

		// This method tests the validity of iarys[i]
		bINumNextAddend, err := new(BigIntNum).NewIntAry(iarys[i])

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bINumNextAddend, err := new(BigIntNum).NewIntAry(iarys[i])",
					ErrContext: fmt.Sprintf("This method tests the validity of iarys[%v]", i),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(bINumAddend, bINumNextAddend)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(bINumAddend, bINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		result, err := bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "result, err := bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		resultsArray[i], err = result.GetIntAry()

		if err != nil {

			iNumStr, _ := result.GetNumStr()

			return []IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: fmt.Sprintf("resultsArray[%v], err = result.GetIntAry()", i),
					ErrContext: fmt.Sprintf("result= %v", iNumStr),
					ErrMessage: err.Error(),
				}
		}
	}

	return resultsArray, nil
}

// AddIntArySeries - Adds a series of IntAry objects and returns the total in a
// BigIntNum.
//
// The BigIntNum result of this addition operation will contain numeric separators
// (decimal separator, thousands separator and currency symbol) which were copied
// from the first element input series 'iarys'.
func (bAdd *BigIntMathAdd) AddIntArySeries(iarys ...IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddIntArySeries() "

	finalResult := new(BigIntNum).New()

	if len(iarys) == 0 {
		return finalResult,
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'iarys' series is Empty!\n",
				ePrefix)
	}

	var err error

	numSeps := NumericSeparatorDto{}

	for i, ia := range iarys {

		if i == 0 {

			numSeps, err = ia.GetNumericSeparatorsDto()

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "numSeps, err = ia.GetNumericSeparatorsDto()",
						ErrContext: "index= 0",
						ErrMessage: err.Error(),
					}
			}

			// This method tests the validity of 'ia'
			finalResult, err = new(BigIntNum).NewIntAry(ia)

			if err != nil {

				iaNumStr, _ := ia.GetNumStr()

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewIntAry(ia)",
						ErrContext: fmt.Sprintf("index= %v ia NumStr = %v", i, iaNumStr),
						ErrMessage: err.Error(),
					}
			}
			continue
		}

		bINumNextAddend, err := new(BigIntNum).NewIntAry(ia)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bINumNextAddend, err := new(BigIntNum).NewIntAry(ia)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bINumNextAddend)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddINumMgr - Receives two objects which implement the INumMgr Interface
// and adds their numeric values.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry.
//
// The result is returned as an instance of Type, 'BigIntNum'.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) which were copied from the input parameter 'num1'.
func (bAdd *BigIntMathAdd) AddINumMgr(num1, num2 INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddINumMgr() "

	num1Str, err := num1.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "num1Str, err := num1.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	num2Str, err := num2.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "num2Str, err := num2.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// This method will test the validity of num1 and num2
	bPair, err := new(BigIntPair).NewINumMgr(num1, num2)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(num1, num2)",
				ErrContext: fmt.Sprintf("num1= '%v' num2= '%v'", num1Str, num2Str),
				ErrMessage: err.Error(),
			}
	}

	finalResult, err := bAdd.AddPair(bPair)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "finalResult, err := bAdd.AddPair(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddINumMgrArray
//
//	Adds an array of objects which implement the 'INumMgr'
//	interface. The combined total of the numeric values from these
//	objects is returned as an instance of Type, 'BigIntNum'.
//
//	The INumMgr interface is implemented by types, BigIntNum, Decimal,
//	NumStrDto and IntAry. This allows the user to mix different types in
//	a single array and add their numeric values.
//
//	The returned BigIntNum result of this addition operation will contain
//	numeric separators (decimal separator, thousands separator and currency
//	symbol) which were copied from the first element of the input parameter
//	array, nums (nums[0]).
//
//	BE CAREFUL - Arrays are tricky
//
//	-- THIS WORKS --
//	for i:= 0; i < something; i++ {
//
//	dec, err := new(Decimal).NewNumStr(numStrAry[i])
//
//	if err != nil {
//	  err := fmt.Errorf("Some Error")
//	  return err
//	}
//
//	inumMgrAry[i] = &dec
//
//	....
//
// }
//
//	-- THIS FAILS --
//
//	var dec Decimal
//
//	for i:= 0; i < something; i++ {
//
//	dec, err = new(Decimal).NewNumStr(numStrAry[i])
//	// DON'T DO THIS!
//
//	if err != nil {
//	  err := fmt.Errorf("Some Error")
//	  return err
//	}
//
//	inumMgrAry[i] = &dec
//
//	....
//
// }
func (bAdd *BigIntMathAdd) AddINumMgrArray(nums []INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrArray() "

	var finalResultNumStr string

	var err error

	finalResult := new(BigIntNum).New()

	lenNums := len(nums)

	if lenNums == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenNums == 0 {",
				ErrMessage: "Error: Input parameter 'nums' array is Empty!",
			}
	}

	numSeps, err := nums[0].GetNumericSeparatorsDto()

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := nums[0].GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	for i := 0; i < lenNums; i++ {

		if i == 0 {
			// This method will test the validity of nums[i]
			finalResult, err = new(BigIntNum).NewINumMgr(nums[i])

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: fmt.Sprintf("finalResult, err = new(BigIntNum).NewINumMgr(nums[%v])", i),
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			err = finalResult.IsValid("Validating finalResult")

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "err = finalResult.IsValid(\"Validating finalResult\")",
						ErrContext: fmt.Sprintf("i= '%v'", i),
						ErrMessage: err.Error(),
					}
			}

			finalResultNumStr, err = finalResult.GetNumStr()

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResultNumStr, err = finalResult.GetNumStr()",
						ErrContext: fmt.Sprintf("i= '%v'", i),
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		bPair, err := new(BigIntPair).NewINumMgr(&finalResult, nums[i])

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(&finalResult, nums[i])",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = bPair.IsValid("Validating finalResult nums[i]")

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = bPair.IsValid(\"Validating finalResult nums[i]\")",
					ErrContext: fmt.Sprintf("i='%v'", i),
					ErrMessage: "Error: bPair is INVALID!\n",
				}

		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(&finalResult, nums[i])",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResultNumStr, err = finalResult.GetNumStr()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResultNumStr, err = finalResult.GetNumStr()",
					ErrContext: fmt.Sprintf("i= '%v'", i),
					ErrMessage: err.Error(),
				}
		}

	} // End of 'for' statement

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = finalResult.IsValid("Validating finalResult")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.IsValid(\"Validating finalResult\")",
				ErrContext: fmt.Sprintf("Final Validation finalResult= '%v'", finalResultNumStr),
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddINumMgrOutputToArray - The first input parameter to this method
// is an object that implements the  INumMgr interface and is labeled,
// 'addend'.  The second element is an array of objects which implement
// the INumMgr interface. This array is labeled, 'numMgrs'. The 'addend'
// is added to each element of the 'numMgrs' array with the result
// output to another INumMgr array which is returned to the calling function.
//
// Example
// =======
//
//											    numMgrs										Output
//	 Addend   				    	Array											Array
//
//			3			+					numMgrs[0] = 2			=				  outputarray[0] =  5
//			3			+					numMgrs[1] = 3			=				  outputarray[1] =  6
//			3			+					numMgrs[2] = 4			=				  outputarray[2] =  7
//			3			+					numMgrs[3] = 5			=				  outputarray[3] =  8
//			3			+					numMgrs[4] = 6			=				  outputarray[4] =  9
//			3			+					numMgrs[5] = 9			=				  outputarray[5] = 12
//
// Each element of the returned []INumMgr array will contain numeric separators
// (decimal separator, thousands separator and currency symbol) which were
// copied from the input parameter 'addend'.
func (bAdd *BigIntMathAdd) AddINumMgrOutputToArray(
	addend INumMgr,
	numMgrs []INumMgr) ([]INumMgr, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrOutputToArray()"

	lenDecs := len(numMgrs)

	if lenDecs == 0 {
		return []INumMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'numMgrs' array is Empty!\n",
				ePrefix)
	}

	numSeps, err := addend.GetNumericSeparatorsDto()

	if err != nil {

		return []INumMgr{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultsArray := make([]INumMgr, lenDecs)

	for i := 0; i < lenDecs; i++ {

		// This method will test the validity of numMgrs[i]
		bPair, err := new(BigIntPair).NewINumMgr(addend, numMgrs[i])

		if err != nil {

			numMgrStr, _ := numMgrs[i].GetNumStr()

			addendStr, _ := addend.GetNumStr()

			return []INumMgr{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: fmt.Sprintf("bPair, err := new(BigIntPair).NewINumMgr(addend, numMgrs[%v])", i),
					ErrContext: fmt.Sprintf("numMgrs[%v].GetNumStr()= '%v' addend.GetNumStr()= '%v'", i, numMgrStr, addendStr),
					ErrMessage: err.Error(),
				}
		}

		bIntNum, err := bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return []INumMgr{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bIntNum, err := bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = bIntNum.SetNumericSeparatorsDto(numSeps)

		if err != nil {

			return []INumMgr{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = bIntNum.SetNumericSeparatorsDto(numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		resultsArray[i] = &bIntNum
	}

	return resultsArray, nil
}

// AddINumMgrSeries - Adds a series of objects which implement the 'INumMgr'
// interface. The combined total of the numeric values from these objects
// is returned as an instance of Type, 'BigIntNum'.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) which were copied from the first element of the input parameter
// 'nums'.
func (bAdd *BigIntMathAdd) AddINumMgrSeries(nums ...INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrSeries()"

	finalResult := new(BigIntNum).New()

	if len(nums) == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if len(nums) == 0 {",
				ErrMessage: "Error: Input parameter 'nums' series is Empty!",
			}
	}

	var err error

	numSeps := NumericSeparatorDto{}

	for i, num := range nums {

		if i == 0 {

			numSeps, err = num.GetNumericSeparatorsDto()

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "numSeps, err = num.GetNumericSeparatorsDto()",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			// This method will test the validity of 'num'
			finalResult, err = new(BigIntNum).NewINumMgr(num)

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewINumMgr(num)",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		bPair, err := new(BigIntPair).NewINumMgr(&finalResult, num)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(&finalResult, num)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddNumStr - Receives two number strings and adds their numeric values.
//
// The first two input parameters, 'n1NumStr' and 'n2NumStr' must be formatted
// as strings of numeric digits, or number strings. Number strings may have
// a leading minus sign ('-') to indicate the numeric sign value. In addition,
// the string of numeric digits may include a delimiting decimal separator
// to identify fractional digits. The number strings are parsed based on the
// decimal separator character specified by input parameter 'numSeps'.
//
// Input parameter 'numSeps' is a type NumericSeparatorDto and is used to
// parse the number strings 'n1NumStr' and 'n2NumStr'. 'numSeps' represents the
// applicable decimal separator, thousands separator and currency symbol. In
// addition, 'numSeps' is also used in configuring the return value for this
// addition operation.
//
// The result is returned as type BigIntNum.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) specified by input parameter, 'numSeps'.
func (bAdd *BigIntMathAdd) AddNumStr(
	n1NumStr string, n2NumStr string, numSeps NumericSeparatorDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStr() "

	numSeps.SetDefaultsIfEmpty()

	bPair, err := new(BigIntPair).NewNumStrWithNumSeps(n1NumStr, n2NumStr, numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by BigIntPair{}.NewNumStr(n1NumStr, n2NumStr).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	finalResult, err := bAdd.AddPair(bPair)

	if err != nil {
		return finalResult,
			fmt.Errorf("%v\n"+
				"Error returned by bAdd.AddPair(bPair).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return finalResult, nil
}

// AddNumStrArray - Adds a series of number strings and returns the combined total
// as an instance of Type 'BigIntNum'.
//
// All the elements of the 'numStrs' array must be formatted as strings of numeric
// digits or number strings. Number strings may have a leading minus sign ('-')
// to indicate the numeric sign value.  In addition, the string of numeric digits
// may include a delimiting decimal separator to identify fractional digits. The
// number strings are parsed based on the decimal separator character specified
// by input parameter 'numSeps'.
//
// Input parameter 'numSeps' is a type NumericSeparatorDto and is used to parse
// the number strings contained in the 'numStrs' array. Input parameter, 'numSeps',
// represents the applicable decimal separator, thousands separator and currency
// symbol. 'numSeps' is also used in configuring the BigIntNum return value for
// this addition operation.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) as specified by input parameter 'numSeps'.
func (bAdd *BigIntMathAdd) AddNumStrArray(
	numStrs []string, numSeps NumericSeparatorDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrArray()"

	numSeps.SetDefaultsIfEmpty()

	finalResult := new(BigIntNum).New()
	var err error

	lenNumStrs := len(numStrs)

	if lenNumStrs == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenNumStrs == 0 {",
				ErrMessage: "Error: Input parameter 'numStrs' array is Empty!",
			}
	}

	for i := 0; i < lenNumStrs; i++ {

		if i == 0 {

			finalResult, err = new(BigIntNum).NewNumStrWithNumSeps(numStrs[i], &numSeps)

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewNumStrWithNumSeps(numStrs[i], numSeps)",
						ErrContext: "index= 0",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		b2Num, err := new(BigIntNum).NewNumStrWithNumSeps(numStrs[i], &numSeps)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "b2Num, err := new(BigIntNum).NewNumStrWithNumSeps(numStrs[i], numSeps)",
					ErrContext: fmt.Sprintf("index= %d nNumStr='%v'", i, numStrs[i]),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.AddBigIntNums(finalResult, b2Num)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.AddBigIntNums(finalResult, b2Num)",
					ErrContext: fmt.Sprintf("index= %d nNumStr='%v'", i, numStrs[i]),
					ErrMessage: err.Error(),
				}
		}
	}

	return finalResult, nil
}

// AddNumStrOutputToArray - The first input parameter to this method
// is a string Type labeled, 'addend'.  The second element is an
// array of string types labeled 'numStrs'. The 'addend' is added to
// each element of the 'numStrs' array with the result output to another
// array of string types. This output array of strings is then returned
// to the calling function.
//
// Input parameters 'addend' and all elements of the 'numStrs' array must
// be formatted as strings of numeric digits or number strings. Number strings
// may have a leading minus sign ('-') to indicate the numeric sign value. In
// addition, the string of numeric digits may include a delimiting decimal
// separator to identify fractional digits. The number strings are parsed based
// on the decimal separator character specified by input parameter 'numSeps'.
//
// Input parameter 'numSeps' is a type NumericSeparatorDto and is used to parse
// the number strings contained in 'addend' and the 'numStrs' array . Input
// parameter, 'numSeps', represents the applicable decimal separator, thousands
// separator and currency symbol. 'numSeps' is also used in configuring the output
// string array returned by this addition operation.
//
// Example
// =======
//
//											    numStrs										 Output
//	 Addend   				    	Array											Array
//
//			3			+					numStrs[0] = 2			=				  outputarray[0] =  5
//			3			+					numStrs[1] = 3			=				  outputarray[1] =  6
//			3			+					numStrs[2] = 4			=				  outputarray[2] =  7
//			3			+					numStrs[3] = 5			=				  outputarray[3] =  8
//			3			+					numStrs[4] = 6			=				  outputarray[4] =  9
//			3			+					numStrs[5] = 9			=				  outputarray[5] = 12
func (bAdd *BigIntMathAdd) AddNumStrOutputToArray(
	addend string,
	numStrs []string,
	numSeps NumericSeparatorDto) ([]string, error) {

	ePrefix := "BigIntMathAdd.AddNumStrOutputToArray()"

	numSeps.SetDefaultsIfEmpty()

	lenNumStrs := len(numStrs)

	if lenNumStrs == 0 {
		return []string{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'numStrs' array is Empty!\n",
				ePrefix)
	}

	bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addend, &numSeps)

	if err != nil {

		return []string{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addend, numSeps)",
				ErrContext: fmt.Sprintf("addend= %v ", addend),
				ErrMessage: err.Error(),
			}
	}

	resultsArray := make([]string, lenNumStrs)

	for i := 0; i < lenNumStrs; i++ {

		b2Num, err := new(BigIntNum).NewNumStrWithNumSeps(numStrs[i], &numSeps)

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addend, numSeps)",
					ErrContext: fmt.Sprintf("i='%v' NumStr='%v'", i, numStrs[i]),
					ErrMessage: err.Error(),
				}
		}

		bigPair, err := new(BigIntPair).NewBigIntNum(bINumAddend, b2Num)

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bigPair, err := new(BigIntPair).NewBigIntNum(bINumAddend, b2Num)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		result, err := bAdd.AddPair(bigPair)

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "result, err := bAdd.AddPair(bigPair)",
					ErrContext: fmt.Sprintf("i='%v' NumStr='%v'", i, numStrs[i]),
					ErrMessage: err.Error(),
				}
		}

		resultsArray[i], err = result.GetNumStr()

		if err != nil {

			return []string{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "resultsArray[i], err = result.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return resultsArray, nil
}

// AddNumStrSeries - Adds a series of number strings and returns
// the combined total as an instance of Type, 'BigIntNum'.
//
// The second Input Parameter ,'numStrs', is series strings of numeric digits,
// or number strings. Number strings may have a leading minus sign ('-')
// to indicate the numeric sign value. In addition, the string of numeric digits
// may include a delimiting decimal separator to identify fractional digits. The
// number strings are parsed based on the decimal separator character specified
// by input parameter 'numSeps'.
//
// Input parameter 'numSeps' is a type NumericSeparatorDto and is used to parse
// the number strings contained in the 'numStrs' series . Input parameter, 'numSeps',
// represents the applicable decimal separator, thousands separator and currency symbol.
// 'numSeps' is also used in configuring the output string array returned by this addition
// operation.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) specified by input parameter 'numSeps'.
func (bAdd *BigIntMathAdd) AddNumStrSeries(
	numSeps NumericSeparatorDto, numStrs ...string) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrSeries() "

	numSeps.SetDefaultsIfEmpty()

	var err error

	finalResult, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "finalResult, err := new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if len(numStrs) == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if len(numStrs) == 0 {",
				ErrMessage: "Error: Input parameter 'numStrs' series is Empty!",
			}
	}

	for i, numStr := range numStrs {

		if i == 0 {

			finalResult, err = new(BigIntNum).NewNumStrWithNumSeps(numStr, &numSeps)

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewNumStrWithNumSeps(numStr, numSeps)",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		b2Num, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, &numSeps)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "b2Num, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.AddBigIntNums(finalResult, b2Num)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "b2Num, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, numSeps)",
					ErrContext: fmt.Sprintf("i='%v'\nNumStr='%v'", i, numStr),
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsToDefaultIfEmpty()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddNumStrDto - Receives two NumStrDto instances and adds their numeric values.
//
// The result is returned as type BigIntNum.
//
// The returned BigIntNum result of this addition operation will contain
// the numeric separators (decimal separator, thousands separator and currency
// symbol) copied from input parameter 'n1Dto'.
func (bAdd *BigIntMathAdd) AddNumStrDto(n1Dto, n2Dto NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	// This method will test the validity of n1Dto and n2Dto
	bPair, err := new(BigIntPair).NewNumStrDto(n1Dto, n2Dto)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bPair, err := new(BigIntPair).NewNumStrDto(n1Dto, n2Dto)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	finalResult, err := bAdd.AddPair(bPair)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by bAdd.AddPair(bPair).\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return finalResult, nil
}

// AddNumStrDtoArray - Adds an array of 'NumStrDto' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) copied from the first element of the input array, nDtos[0].
func (bAdd *BigIntMathAdd) AddNumStrDtoArray(nDtos []NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddDecimalArray() "

	finalResult := new(BigIntNum).New()

	var err error

	lenNDtos := len(nDtos)

	if lenNDtos == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if lenNDtos == 0 {",
				ErrMessage: "Error: Input parameter 'nDtos' array is Empty!",
			}
	}

	numSeps := NumericSeparatorDto{}

	for i := 0; i < lenNDtos; i++ {

		if i == 0 {

			// This method will test the validity of 'nDtos[i]'
			finalResult, err = new(BigIntNum).NewNumStrDto(nDtos[i])

			if err != nil {

				numStr, _ := nDtos[i].GetNumStr()

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewNumStrDto(nDtos[i])",
						ErrContext: fmt.Sprintf("nDtos[%v].GetNumStr()='%v'", i, numStr),
						ErrMessage: err.Error(),
					}
			}

			numSeps, err = nDtos[0].GetNumericSeparatorsDto()

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "numSeps, err = nDtos[0].GetNumericSeparatorsDto()",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		bPair, err := new(BigIntPair).NewINumMgr(&finalResult, &nDtos[i])

		if err != nil {

			numStr, _ := nDtos[i].GetNumStr()

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewINumMgr(&finalResult, &nDtos[i])",
					ErrContext: fmt.Sprintf("nDtos[%v].GetNumStr()='%v'", i, numStr),
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddNumStrDtoOutputToArray - The first input parameter to this method
// is a NumStrDto Type labeled, 'addend'.  The second element is an
// array of NumStrDto types labeled 'nDtos'. The 'addend' is added to
// each element of the 'nDtos' array with the result output to another
// array of NumStrDto types which is returned to the calling function.
//
// Example
// =======
//
//	                        nDtos                      Output
//	  Addend                Array                      Array
//
//				3			+					nDtos[0] = 2			=				  outputarray[0] =  5
//				3			+					nDtos[1] = 3			=				  outputarray[1] =  6
//				3			+					nDtos[2] = 4			=				  outputarray[2] =  7
//				3			+					nDtos[3] = 5			=				  outputarray[3] =  8
//				3			+					nDtos[4] = 6			=				  outputarray[4] =  9
//				3			+					nDtos[5] = 9			=				  outputarray[5] = 12
//
// Each element in the returned array of []NumStrDto resulting of this addition
// operation will contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter 'addend'.
func (bAdd *BigIntMathAdd) AddNumStrDtoOutputToArray(
	addend NumStrDto,
	nDtos []NumStrDto) ([]NumStrDto, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDtoOutputToArray()"

	lenNDtos := len(nDtos)

	if lenNDtos == 0 {
		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
				ErrContext: "if lenNDtos == 0 {",
				ErrMessage: "Error: Input parameter 'nDtos' array is Empty!",
			}
	}

	numSeps, err := addend.GetNumericSeparatorsDto()

	if err != nil {

		return []NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	resultsArray := make([]NumStrDto, lenNDtos)

	for i := 0; i < lenNDtos; i++ {

		err := nDtos[i].IsValid(ePrefix +
			fmt.Sprintf("nDtos[%v] INVALID! ", i))

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix,
					ReturnFunc: "err := nDtos[i].IsValid(ePrefix +\n" +
						"    fmt.Sprintf(\"nDtos[%v] INVALID! \", i))",
					ErrContext: fmt.Sprintf("nDtos[%v] FAILED Validation Testing", i),
					ErrMessage: err.Error(),
				}
		}

		// This method tests the validity of addend and nDtos[i]
		bPair, err := new(BigIntPair).NewNumStrDto(addend, nDtos[i])

		if err != nil {

			numStr, _ := nDtos[i].GetNumStr()

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "numSeps, err := addend.GetNumericSeparatorsDto()",
					ErrContext: fmt.Sprintf("nDtos[%v].GetNumStr()='%v'", i, numStr),
					ErrMessage: err.Error(),
				}
		}

		result, err := bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "result, err := bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		resultsArray[i], err = result.GetNumStrDto()

		if err != nil {

			numStr, _ := nDtos[i].GetNumStr()

			return []NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
					ErrContext: fmt.Sprintf("nDtos[%v].GetNumStr()='%v'", i, numStr),
					ErrMessage: err.Error(),
				}
		}
	}

	return resultsArray, nil
}

// AddNumStrDtoSeries - Adds a series of 'NumStrDto' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The returned BigIntNum resulting of this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied
// from the first element of the input series 'nDtos'.
func (bAdd *BigIntMathAdd) AddNumStrDtoSeries(nDtos ...NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDtoSeries() "

	finalResult := new(BigIntNum).New()

	if len(nDtos) == 0 {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "if len(nDtos) == 0 {\n\t",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'nDtos' series is Empty!",
			}
	}

	var err error
	numSeps := NumericSeparatorDto{}

	for i, nDto := range nDtos {

		if i == 0 {

			// This method will test the validity of 'nDto'
			finalResult, err = new(BigIntNum).NewNumStrDto(nDto)

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "finalResult, err = new(BigIntNum).NewNumStrDto(nDto)",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			numSeps, err = nDto.GetNumericSeparatorsDto()

			if err != nil {

				return finalResult,
					&FuncReturnError{
						ErrPrefix:  ePrefix,
						ReturnFunc: "numSeps, err = nDto.GetNumericSeparatorsDto()",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
			}

			continue
		}

		// This method will test the validity of 'nDto'
		bINumNextAddend, err := new(BigIntNum).NewNumStrDto(nDto)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bINumNextAddend, err := new(BigIntNum).NewNumStrDto(nDto)",
					ErrContext: fmt.Sprintf("i='%v'", i),
					ErrMessage: err.Error(),
				}
		}

		bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bINumNextAddend)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(finalResult, bINumNextAddend)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		finalResult, err = bAdd.addPairNoNumSeps(bPair)

		if err != nil {

			return finalResult,
				&FuncReturnError{
					ErrPrefix:  ePrefix,
					ReturnFunc: "finalResult, err = bAdd.addPairNoNumSeps(bPair)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// AddPair - Receives a BigIntPair instance and proceeds to add b1.BigIntNum
// to b2.BigIntNum.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) copied from b1.BigIntNum.
func (bAdd *BigIntMathAdd) AddPair(bPair BigIntPair) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddPair()"

	numSeps, err := bPair.Big1.GetNumericSeparatorsDto()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numSeps, err := bPair.Big1.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	finalResult, err := bAdd.addPairNoNumSeps(bPair)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "finalResult, err := bAdd.addPairNoNumSeps(bPair)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}

// BigIntAdd - Adds two fixed length floating point numbers and generates
// the sum or total resulting from that addition. The two numbers added
// together are configured as pairs of *big.Int integer numbers and precision
// specifications. Each integer precision pair is used to define a fixed
// length floating point number.
//
// Examples:
// =========
//
// In the addition operation:
//
//	b1 + b2 = total or sum
//
// This method provides for the addition of fixed length
// floating point values by means of integer and precision
// specification pairs.
//
// As an example, consider the following addition operation
//
//							752.314 + 21.67894 = 773.99294 = total
//	             b1    +     b2   =   total
//
// In this case 'b1', 'b2' and 'total' would be configured as integer
// precision pairs:
//
//										b1 							= 752314
//	                 b1Precision			= 3
//	                 b2 							= 2167894
//	                 b2Precision			= 5
//
//	                 total						= 77399294
//	                 totalPrecision 	= 5
//
// In this way, the method uses integer, precision pairs to define fixed
// length floating point numbers.
//
// Input Parameters
// ================
//
//		b1 					*big.Int	- The first number which will be added to 'b2' to
//	                   			generate a total.
//
//		b1Precision	*big.Int	- Specifies the precision for input parameter 'b1'.
//														Precision defines the number of fractional digits
//														after the decimal place. 'b1Precision' must be equal
//	                         to or greater than zero.
//
//		b2 					*big.Int	- The second number which is added to 'b1' in order to
//	                					generate a total.
//
//		b2Precision	*big.Int  - The 'b2' precision or the number of fractional digits
//														after the decimal place. 'b2Precision' must be equal
//	                         to or greater than zero.
//
// Return Values
// =============
//
// total 					*big.Int		- The sum or total of 'b1' and 'b2' input values.
//
// totalPrecision *big.Int   	- The 'total' precision or the number of fractional
//
//	digits after the decimal place.
//
// err						error				- If input parameters 'b1Precision' or 'b2Precision'
//
//	are less than zero, an error will be returned.
//
// Taken together, 'total' and 'totalPrecision' can define a fixed length floating point number.
func (bAdd *BigIntMathAdd) BigIntAdd(
	b1,
	b1Precision,
	b2,
	b2Precision *big.Int) (total *big.Int, totalPrecision *big.Int, err error) {

	ePrefix := "BigIntMathAdd.BigIntAdd()"

	total = big.NewInt(0)
	totalPrecision = big.NewInt(0)
	err = nil

	if b1 == nil {
		b1 = big.NewInt(0)
	}

	if b2 == nil {
		b2 = big.NewInt(0)
	}

	bigZero := big.NewInt(0)

	if b1Precision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'b1Precision' is LESS THAN ZERO!\n"+
			"b1Precision='%v'\n",
			ePrefix,
			b1Precision.Text(10))

		return total, totalPrecision, err
	}

	if b2Precision.Cmp(bigZero) == -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'b2Precision' is LESS THAN ZERO!\n"+
			"b2Precision='%v'\n",
			ePrefix,
			b2Precision.Text(10))

		return total, totalPrecision, err
	}

	if b1.Cmp(bigZero) == 0 &&
		b2.Cmp(bigZero) == 0 {
		total = big.NewInt(0)
		totalPrecision = big.NewInt(0)
		return total, totalPrecision, nil
	}

	bigTen := big.NewInt(10)
	delta := big.NewInt(0)
	scale := big.NewInt(0)

	if b1Precision.Cmp(b2Precision) == 0 {
		total = big.NewInt(0).Add(b1, b2)
		totalPrecision = big.NewInt(0).Set(b1Precision)

	} else if b1Precision.Cmp(b2Precision) == 1 {
		// b1Precision > b2Precision
		delta = big.NewInt(0).Sub(b1Precision, b2Precision)
		scale = big.NewInt(0).Exp(bigTen, delta, nil)

		b2ToScale := big.NewInt(0).Mul(b2, scale)

		total = big.NewInt(0).Add(b1, b2ToScale)
		totalPrecision = big.NewInt(0).Set(b1Precision)

	} else {
		// b2Precision must be GREATER than b1Precision
		delta = big.NewInt(0).Sub(b2Precision, b1Precision)

		scale = big.NewInt(0).Exp(bigTen, delta, nil)

		b1ToScale := big.NewInt(0).Mul(b1, scale)

		total = big.NewInt(0).Add(b1ToScale, b2)

		totalPrecision = big.NewInt(0).Set(b2Precision)

	}

	if total.Cmp(bigZero) == 0 {
		totalPrecision = big.NewInt(0)
	}

	// Delete trailing fractional zeros
	if totalPrecision.Cmp(bigZero) == 1 {
		//totalPrecision > 0
		scrap := big.NewInt(0)
		biBase10 := big.NewInt(10)
		biBaseZero := big.NewInt(0)
		newTotal, mod10 := big.NewInt(0).QuoRem(total, biBase10, scrap)
		bigOne := big.NewInt(1)
		for mod10.Cmp(biBaseZero) == 0 && totalPrecision.Cmp(bigZero) == 1 {
			total.Set(newTotal)
			totalPrecision.Sub(totalPrecision, bigOne)
			newTotal, mod10 = big.NewInt(0).QuoRem(total, biBase10, scrap)
		}
	}

	return total, totalPrecision, nil
}

// FixedDecimalAdd - Performs an addition operation using two BigIntFixedDecimal
// types. The addition result or total is also returned as a BigIntFixedDecimal
// type.
//
// Examples:
// =========
//
// In the addition operation:
//
//	b1 + b2 = total or sum
//
// For this method 'b1', 'b2' and 'total' are all configured as BigIntFixedDecimal
// types.
//
// The BigIntFixedDecimal type is used to defined fixed length floating point
// numbers and is defined as follows:
//
// type BigIntFixedDecimal struct {
//
//	integerNum *big.Int  -	All the numeric digits, both integer and fractional,
//													necessary to define a fixed length floating point number.
//													The number of digits to the right of the decimal place
//													is specified by the data field,
//													BigIntFixedDecimal.precision.
//
//	precision  uint				- Specifies the number of digits to the right of the decimal
//													place in the series of numeric digits represented by data
//													field BigIntFixedDecimal.integerNum.
//
// }
//
//	To represent the floating point number 52.459	a BigIntDecimal Structure
//	would be configured as follows:
//
//			BigIntFixedDecimal.integerNum	= 52459
//			BigIntFixedDecimal.precision	= 3
//
// As an example consider the following addition operation:
//
//							752.314 + 21.67894 = 773.99294 = total
//	             b1    +     b2   =   total
//
// In this case 'b1', 'b2' and 'total' would be configured as BigIntDecimal
// types:
//
//										b1.integerNum			= 752314
//	                 b1.precision			= 3
//	                 b2.integerNum			= 2167894
//	                 b2.precision 			= 5
//
//	                 total.integerNum	= 77399294
//	                 total.precision 	= 5
//
// In this way, the method uses BigIntFixedDecimal types to define fixed
// length floating point numbers.
//
// Input Parameters
// ================
//
//		b1 BigIntFixedDecimal	- The first number which will be added to 'b2' to
//	                   			generate a total.
//
//		b2 BigIntFixedDecimal	- The second number which is added to 'b1' in order to
//	                					generate a total.
//
// Return Values
// =============
//
// total BigIntFixedDecimal	- The sum or total of 'b1' and 'b2' input values.
func (bAdd *BigIntMathAdd) FixedDecimalAdd(
	b1,
	b2 BigIntFixedDecimal) (total BigIntFixedDecimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathAdd.FixedDecimalAdd",
		"")

	if err != nil {
		return total, err
	}

	total = new(BigIntFixedDecimal).NewZero(0)

	err = b1.IsValid(ePrefix.XCpy("Validating 'b1'").String())

	if err != nil {

		return total,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = b1.IsValid(ePrefix.XCpy(\"Validating 'b1'\").String())",
				ErrContext: "Input parameter 'b1' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	err = b2.IsValid(ePrefix.XCpy("Validating 'b2'").String())

	if err != nil {

		return total,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = b1.IsValid(ePrefix.XCpy(\"Validating 'b1'\").String())",
				ErrContext: "Input parameter 'b1' is invalid!",
				ErrMessage: err.Error(),
			}
	}

	// No error is possible because by definition, both precision
	// values must be equal to or greater than zero.

	var b1BigIntValue, b1BigIntPrecision,
		b2BigIntValue, b2BigIntPrecision *big.Int

	b1BigIntValue, err = b1.GetIntegerValue()

	if err != nil {

		return total,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "b1BigIntValue, err = b1.GetIntegerValue()",
				ErrContext: "b1 *big.Int is invalid!",
				ErrMessage: err.Error(),
			}
	}

	b1BigIntPrecision, err = b1.GetPrecisionBigInt()

	if err != nil {

		return total,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "b1BigIntPrecision, err = b1.GetPrecisionBigInt()",
				ErrContext: "b1 Precision Value is invalid!",
				ErrMessage: err.Error(),
			}
	}

	b2BigIntValue, err = b2.GetIntegerValue()

	if err != nil {

		return total,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "b2BigIntValue, err = b2.GetIntegerValue()",
				ErrContext: "b2 *big.Int is invalid!",
				ErrMessage: err.Error(),
			}
	}

	b2BigIntPrecision, err = b2.GetPrecisionBigInt()

	if err != nil {

		return total,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "b2BigIntPrecision, err = b2.GetPrecisionBigInt()",
				ErrContext: "b2 Precision Value is invalid!",
				ErrMessage: err.Error(),
			}
	}

	bIResult, bIPrecision, err :=
		new(BigIntMathAdd).BigIntAdd(
			b1BigIntValue,
			b1BigIntPrecision,
			b2BigIntValue,
			b2BigIntPrecision)

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bIResult, bIPrecision, err :=\n" +
					"    b1BigIntValue, b1BigIntPrecision,\n" +
					"    b2BigIntValue, b2BigIntPrecision,\n",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = total.SetNumericValue(bIResult, uint(bIPrecision.Uint64()))

	if err != nil {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = total.SetNumericValue(bIResult,\n" +
					"uint(bIPrecision.Uint64()))\n",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return total, err
}

// addPairNoNumSeps - Receives a BigIntPair and proceeds to add b1.BigIntNum to
// b2.BigIntNum.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// default numeric separators (decimal separator, thousands separator and
// currency symbol).
func (bAdd *BigIntMathAdd) addPairNoNumSeps(bPair BigIntPair) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddPairNoNumSeps()"

	err := bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err := bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigI1, err := bPair.GetBig1BigInt()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bigI1, err := bPair.GetBig1BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigI2, err := bPair.GetBig2BigInt()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bigI2, err := bPair.GetBig2BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b3 := big.NewInt(0).Add(bigI1, bigI2)

	big2Precision, err := bPair.Big2.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "big2Precision, err := bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bResult, err := new(BigIntNum).NewBigInt(b3, big2Precision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bResult, err := new(BigIntNum).NewBigInt(b3, big2Precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bResult, nil
}
