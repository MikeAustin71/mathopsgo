package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

type IntAryMathMultiply struct {
	Input  IntAryPair
	Result IntAry
}

// MultiplyByTenToPower - Input parameter 'base' is multiplied by 10 to the exponent
// power. The result is stored in base.
//
// base = base X 10^power
//
// Input Parameters
// ================
//
// base		*IntAry		- Pointer to an IntAry numeric value. 'base' is multiplied
//
//	by 10 to the exponent 'power' and the result is stored
//	in base. The original value of 'base' will therefore be
//	overwritten and destroyed.
//
// power  uint			- 10 is raised to the exponent 'power' and multiplied by 'base'.
func (iaMultiply *IntAryMathMultiply) MultiplyByTenToPower(base *IntAry, power uint) error {
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathMultiply.MultiplyByTenToPower()",
		"")

	if err != nil {
		return err
	}

	if power == 0 {
		return nil
	}

	for i := uint(0); i < power; i++ {

		if base.precision > 0 {
			base.precision--
			continue
		}

		base.intAry = append(base.intAry, 0)
	}

	base.intAryLen = len(base.intAry)

	if base.precision < 0 {
		base.precision = 0
	}

	err = base.OptimizeIntArrayLen(false)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.OptimizeIntArrayLen(false)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// MultiplyByTwoToPower - Multiplies 'base' by 2 to the exponent 'power'. The result
// is stored in 'base'.
//
// base = base X (2^power)
//
// Input Parameters
// ================
//
// base		*IntAry		- Pointer to an IntAry numeric value. 'base' is multiplied
//
//	by 2 to the exponent 'power' and the result is stored
//	in base. The original value of 'base' will therefore be
//	overwritten and destroyed.
//
// power  uint			- 2 is raised to the exponent 'power' and multiplied by 'base'.
func (iaMultiply *IntAryMathMultiply) MultiplyByTwoToPower(base *IntAry, power uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathMultiply.MultiplyByTwoToPower()",
		"")

	if err != nil {
		return err
	}

	err = base.SetIntAryLength()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.SetIntAryLength()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if power == 0 {
		return nil
	}

	for h := uint(0); h < power; h++ {

		n1 := uint8(0)

		carry := uint8(0)

		for i := base.intAryLen - 1; i >= 0; i-- {

			n1 = (base.intAry[i] * 2) + carry

			if n1 > 9 {

				n1 = n1 - 10
				carry = 1

			} else {

				carry = 0
			}

			base.intAry[i] = n1
		}

		if carry > 0 {

			base.intAry = append([]uint8{1}, base.intAry...)

			base.intAryLen++
		}

	}

	return nil
}

// MultiplyInPlace - Multiplies IntAry parameter 'ia1' by IntAry parameter 'ia2' and
// returns the result in parameter 'ia1'.
//
// Input Parameters
// ================
//
// 'ia1' -				Pointer to an intAry object. 'ia1' will be multiplied by 'ia2' and the
//
//	result will be stored here in 'ia1'.
//
// 'ia2' - 				Pointer to an intAry object. In this multiplication operation, 'ia2'
//
//	is the multiplier.
//
// 'minimumResultPrecision' int -
//
//	'minimumResultPrecision' will determine the minimum number of digits computed
//	to the right of the decimal place in the final result.
//
//	If 'minimumResultPrecision' is set to a value of -1, all significant digits (digits
//	greater than zero) will be returned to the right of the decimal place. Remember
//	that the maximum number of decimal digits returned will be controlled by parameter
//	'maxResultPrecision'
//
//	If 'minimumResultPrecision' is set to a value less than -1, an error will be triggered.
//
// 'maxResultPrecision' 		int -
//
//									'maxResultPrecision' will determine the maximum
//									number of digits to the right of the decimal
//									place in the result.
//
//									Valid values are -1 and values >= zero ('0')
//	       				Values less than -1 will trigger an error.
//
//									A value of -1 signals that no limit will be placed on
//									the number of decimals places to right of the decimal
//									point in the result. Be advised that a very, very large number
//									of decimal digits may be accommodated by the IntAry Type.
func (iaMultiply *IntAryMathMultiply) MultiplyInPlace(
	ia1, ia2 *IntAry,
	minimumResultPrecision, maxResultPrecision int) error {

	ePrefix := "IntAryMathMultiply.MultiplyInPlace() "

	err := iaMultiply.Multiply(ia1, ia2, ia1, minimumResultPrecision, maxResultPrecision)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by iaMultiply.Multiply(). Error='%v'", err.Error())
	}

	return nil
}

// Multiply - Multiplies intAry ia1 by intAry ia2 and stores the multiplication result
// in intAry iaResult.
//
// Parameters
// ==========
//
// 'ia1' - 				Pointer to an intAry object. In this multiplication operation, 'ia1'
//
//	is the multiplicand.
//
// 'ia2' - 				Pointer to an intAry object. In this multiplication operation, 'ia2'
//
//									is the multiplier.
//
//	 'iaResult' -	Pointer to an intAry object which will be populated with the result
//									of the multiplication operation. The multiplication operation is achieved
//									by multiplying 'ia1' by 'ia2'.
//
// 'minimumResultPrecision' int -
//
//	'minimumResultPrecision' will determine the minimum number of digits computed
//	to the right of the decimal place in the final result.
//	If 'minimumResultPrecision' is set to a value of -1, all significant digits (digits
//	greater than zero) will be returned to the right of the decimal place. Remember
//	that the maximum number of decimal digits returned will be controlled by parameter
//	'maxResultPrecision'
//
// 'maxResultPrecision' 		int -
//
//									'maxResultPrecision' will determine the maximum
//									number of digits to the right of the decimal
//									place in the result.
//
//									Valid values are -1 and values >= zero ('0')
//	       				Values less than -1 will trigger an error.
//
//									A value of -1 signals that no limit will be placed on
//									the number of decimals places to right of the decimal
//									point in the result. Be advised that a very, very large number
//									of decimal digits may be accommodated by the IntAry Type.
//
// The returned parameter 'iaResult' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from input parameter, 'ia1'.
func (iaMultiply *IntAryMathMultiply) Multiply(
	ia1, ia2, iaResult *IntAry,
	minimumResultPrecision,
	maxResultPrecision int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathMultiply.Multiply()",
		"")

	if err != nil {
		return err
	}

	if maxResultPrecision < -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "maxResultPrecision < -1",
			ErrMessage: fmt.Sprintf("Error: Input Parameter 'maxResultPrecision' is less than -1.\n"+
				"  maxResultPrecision= %v", maxResultPrecision),
		}
	}

	if minimumResultPrecision < -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "minimumResultPrecision < -1",
			ErrMessage: fmt.Sprintf("Error: Input Parameter 'minimumResultPrecision is less than -1.\n"+
				"minimumResultPrecision='%v'", minimumResultPrecision),
		}
	}

	err = ia1.IsValid(ePrefix.XCpy("'ia1' is INVALID!").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia1.IsValid(ePrefix.XCpy('ia1' is INVALID!).String())",
			ErrContext: "Input parameter 'ia1' is INVALID!\n" +
				"'ia1' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	err = ia2.IsValid(ePrefix.XCpy("'ia2' is INVALID!").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia2.IsValid(ePrefix.XCpy('ia2' is INVALID!).String())",
			ErrContext: "Input parameter 'ia2' is INVALID!\n" +
				"'ia2' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	numSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "numSeps, err := ia1.GetNumericSeparatorsDto()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if minimumResultPrecision > maxResultPrecision &&
		maxResultPrecision != -1 {

		maxResultPrecision = minimumResultPrecision
	}

	err = ia1.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia1.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = ia2.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia2.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if ia1.isZeroValue || ia2.isZeroValue {

		if minimumResultPrecision < 1 {
			minimumResultPrecision = 0
		}

		err = iaResult.SetIntAryToZero(uint(minimumResultPrecision))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaResult.SetIntAryToZero(uint(minimumResultPrecision))",
				ErrContext: fmt.Sprintf("minimumResultPrecision= %v", minimumResultPrecision),
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	newSignVal := 1

	if ia1.signVal != ia2.signVal {
		newSignVal = -1
	}

	newPrecision := ia1.precision + ia2.precision

	newIntAryLen := ia1.intAryLen + ia2.intAryLen

	resultAry := make([]uint8, newIntAryLen)

	carry := uint8(0)
	multiplier := uint8(0)
	multiplicand := uint8(0)
	product := uint8(0)
	resultIdx := 0
	offset := 0

	for i := ia2.intAryLen - 1; i >= 0; i-- {

		multiplicand = ia2.intAry[i]

		offset++

		nextResultIdx := newIntAryLen - offset

		for j := ia1.intAryLen - 1; j >= 0; j-- {

			multiplier = ia1.intAry[j]

			product = multiplier * multiplicand

			resultIdx = nextResultIdx

			resultAry[resultIdx] += product

			for resultAry[resultIdx] > 9 {

				carry = resultAry[resultIdx] / 10

				resultAry[resultIdx] = resultAry[resultIdx] - (carry * 10)

				resultIdx--

				resultAry[resultIdx] += carry

			}

			carry = 0

			nextResultIdx--
		}

	}

	if newIntAryLen-newPrecision > 1 && resultAry[0] == 0 {

		iaResult.intAry = resultAry[1:]

		newIntAryLen--

	} else {

		iaResult.intAry = resultAry
	}

	iaResult.precision = newPrecision
	iaResult.signVal = newSignVal
	iaResult.intAryLen = newIntAryLen
	iaResult.integerLen = newIntAryLen - newPrecision
	iaResult.isZeroValue = false

	if minimumResultPrecision < 0 {

		err = iaResult.OptimizeIntArrayLen(true)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaResult.OptimizeIntArrayLen(true)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		newPrecision = iaResult.precision

	} else if newPrecision < minimumResultPrecision {

		err = iaResult.SetPrecision(minimumResultPrecision, false)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaResult.SetPrecision(minimumResultPrecision, false)",
				ErrContext: fmt.Sprintf("minimumResultPrecision= %v", minimumResultPrecision),
				ErrMessage: err.Error(),
			}
		}

	}

	if maxResultPrecision > -1 && maxResultPrecision < newPrecision {

		err = iaResult.SetPrecision(maxResultPrecision, true)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaResult.SetPrecision(maxResultPrecision, true)",
				ErrContext: fmt.Sprintf("maxResultPrecision= %v", maxResultPrecision),
				ErrMessage: err.Error(),
			}
		}

	}

	err = iaResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = iaResult.SetNumericSeparatorsDto(numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
