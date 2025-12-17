package mathops

import (
	"sync"
)

type UtilityMathExponents struct {
	lock *sync.Mutex
}

// Int64ExponentByLoop
//
//	Computes int64 base rasied to the power of int64 exponent.
//	The calculation uses a loop to perform multiple multiplications
//	necessary to compute the result.
func (utilMathExpo *UtilityMathExponents) Int64ExponentByLoop(
	base int64, exponent int64) (result int64) {

	result = 1

	if exponent == 0 || base == 1 {
		return result
	}

	for i := int64(0); i < exponent; i++ {
		result *= base
	}

	return result
}

// Int64ExponentBySquare
//
//	Computes int64 base rasied to the power of int64 exponent.
//	The calculation uses a 'square' technique to compute the
//	result.
func (utilMathExpo *UtilityMathExponents) Int64ExponentBySquare(
	base int64, exponent int64) (result int64) {

	result = int64(1)

	if exponent == 0 || base == 1 {
		return result
	}

	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}

	return result
}
