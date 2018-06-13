package mathops

type IntAryMathSubtract struct {
	Input  IntAryPair
	Result IntAry
}

func (iaSubtract IntAryMathSubtract) Subtract(ia1, ia2 *IntAry) IntAry {

	ia3 := ia1.CopyOut()

	iaSubtract.SubtractTotal(&ia3, ia2)

	return ia3
}

func (iaSubtract IntAryMathSubtract) SubtractTotal(ia1, ia2 *IntAry) {

	ia1.SetEqualArrayLengths(ia2)

	if ia1.isZeroValue && ia2.isZeroValue {
		ia1.SetIntAryToZero(ia1.precision)
		return
	}

	compare := ia1.CompareAbsoluteValues(ia2)
	isZeroResult := false

	// Largest Value in now in N1 slot
	newSignVal := ia1.signVal
	doAdd := false
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if ia1.signVal == 1 && ia2.signVal == 1 {
			doAdd = false
			newSignVal = 1
		} else if ia1.signVal == -1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia1.signVal == -1 && ia2.signVal == -1 {
			doAdd = false
			newSignVal = -1
		} else {
			// Must Be ia1.signVal == 1 && ia2.signVal == -1
			doAdd = true
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if ia1.signVal == 1 && ia2.signVal == 1 {
			doAdd = false
			doReverseNums = true
			newSignVal = -1
		} else if ia1.signVal == -1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia1.signVal == -1 && ia2.signVal == -1 {
			doAdd = false
			doReverseNums = true
			newSignVal = 1
		} else {
			// Must Be ia1.signVal == 1 && ia2.signVal == -1
			doAdd = true
			newSignVal = 1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if ia1.signVal == 1 && ia2.signVal == 1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else if ia1.signVal == -1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia1.signVal == -1 && ia2.signVal == -1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else {
			// Must Be ia1.signVal == 1 && ia2.signVal == -1
			doAdd = true
			newSignVal = 1
		}

	}

	iaSubtract.addToSubtract(ia1, ia2, newSignVal, doAdd, isZeroResult, doReverseNums)

	return
}


// addToSubtract - Adds or subtracts two IntAry instances
// and returns the result in the first IntAry parameter.
func (iaSubtract *IntAryMathSubtract) addToSubtract(
																				ia1, ia2 *IntAry,
																					newSignVal int,
																						doAdd bool,
																							isZeroResult bool,
																								doReverseNums bool) {

	if isZeroResult {
		ia1.SetIntAryToZero(ia1.precision)
		return
	}

	ia1.signVal = newSignVal

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := ia1.intAryLen - 1; j >= 0; j-- {

		if doReverseNums {

			n2 = int(ia1.intAry[j])
			n1 = int(ia2.intAry[j])

		} else {
			n1 = int(ia1.intAry[j])
			n2 = int(ia2.intAry[j])

		}

		if doAdd {
			// doAdd == true
			// Do Addition

			n3 = n1 + n2 + carry

			if n3 > 9 {
				n3 = n1 + n2 + carry - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// doAdd == false
			// Do Subtraction
			n3 = n1 - n2 - carry

			if n3 < 0 {
				n3 = n1 + 10 - n2 - carry
				carry = 1
			} else {
				carry = 0
			}
		}

		ia1.intAry[j] = uint8(n3)

	}

	if carry > 0 {
		ia1.intAry = append([]uint8{1}, ia1.intAry...)
		ia1.intAryLen++
	}

	if ia1.intAry[0] == 0 {
		ia1.SetSignificantDigitIdxs()
		ia1.intAry = ia1.intAry[ia1.firstDigitIdx:]
	}

	ia1.SetInternalFlags()

	return
}

