package mathops

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

// MathProcessUtility
// This structure encapsulates utility methods used
// by all elements of the mathops project.
type MathProcessUtility struct {
	lock *sync.Mutex
}

// CodeDurationToStr
// Private method used to convert code timings to meaningful string output
// expressing said timings as hour minutes and seconds etc.
func (mathProcUtil *MathProcessUtility) CodeDurationToStr(tDuration time.Duration) (string, error) {

	ePrefix := "MathProcessUtility.CodeDurationToStr()"

	tMinutes := int64(0)
	tSeconds := int64(0)
	tMilliseconds := int64(0)
	tMicroseconds := int64(0)
	tNanoseconds := int64(0)

	i64TDur := int64(tDuration)
	outStr := ""
	totNanoSecs := int64(0)

	if i64TDur >= int64(time.Minute) {

		tMinutes = i64TDur / int64(time.Minute)
		outStr += fmt.Sprintf("%v-Minutes ", tMinutes)
		i64TDur -= tMinutes * int64(time.Minute)
		totNanoSecs = tMinutes * int64(time.Minute)
	}

	if i64TDur >= int64(time.Second) {
		tSeconds = i64TDur / int64(time.Second)
		outStr += fmt.Sprintf("%v-Seconds ", tSeconds)
		i64TDur -= tSeconds * int64(time.Second)
		totNanoSecs += tSeconds * int64(time.Second)
	}

	if i64TDur >= int64(time.Millisecond) {
		tMilliseconds = i64TDur / int64(time.Millisecond)
		i64TDur -= tMilliseconds * int64(time.Millisecond)
		totNanoSecs += tMilliseconds * int64(time.Millisecond)
	}

	if i64TDur >= int64(time.Microsecond) {
		tMicroseconds = i64TDur / int64(time.Microsecond)
		i64TDur -= tMicroseconds * int64(time.Microsecond)
		totNanoSecs += tMicroseconds * int64(time.Microsecond)
	}

	tNanoseconds = i64TDur
	totNanoSecs += tNanoseconds

	if totNanoSecs != int64(tDuration) {

		return "ERROR",
			fmt.Errorf("%v\n"+
				"Error: Total Time Duration does NOT match Total Nanoseconds Duration!\n"+
				"Total Calculated Duration= '%v'\n"+
				"Total Actual Duration= '%v'\n",
				ePrefix,
				totNanoSecs,
				int64(tDuration))
	}

	outStr += fmt.Sprintf("%v-Milliseconds ", tMilliseconds)

	outStr += fmt.Sprintf("%v-Microseconds ", tMicroseconds)

	outStr += fmt.Sprintf("%v-Nanoseconds ", tNanoseconds)

	return outStr, nil
}

// DoesBigIntExceedMax32BitInt
//
//	Tests a signed *Big Int value to determine if it exceeds the
//	maximum allowable value for a 32-bit integer.
//
//	The maximum value for a 32-bit integer is 2,147,483,647 or
//	2^31 - 1.
func (mathProcUtil *MathProcessUtility) DoesBigIntExceedMax32BitInt(
	candidateBitIntValue *big.Int) (
	bigIntIsNilPtr bool, bigIntExceedMax32BitInt bool, bigIntLessThanZero bool) {

	if candidateBitIntValue == nil {

		bigIntIsNilPtr = true
		bigIntExceedMax32BitInt = false
		bigIntLessThanZero = false

		return bigIntIsNilPtr, bigIntExceedMax32BitInt, bigIntLessThanZero
	}

	max32BitIntBigInt := big.NewInt(math.MaxInt32)

	if candidateBitIntValue.Cmp(max32BitIntBigInt) == 1 {

		bigIntIsNilPtr = false
		bigIntExceedMax32BitInt = true
		bigIntLessThanZero = false

		return bigIntIsNilPtr, bigIntExceedMax32BitInt, bigIntLessThanZero
	}

	bigZero := big.NewInt(0)

	if candidateBitIntValue.Cmp(bigZero) == -1 {

		bigIntIsNilPtr = false
		bigIntExceedMax32BitInt = false
		bigIntLessThanZero = true

		return bigIntIsNilPtr, bigIntExceedMax32BitInt, bigIntLessThanZero
	}

	bigIntIsNilPtr = false
	bigIntExceedMax32BitInt = false
	bigIntLessThanZero = false

	return bigIntIsNilPtr, bigIntExceedMax32BitInt, bigIntLessThanZero
}

// DoesUintExceedMax32BitInt
//
//	Tests an unsigned integer value to determine if it exceeds the
//	maximum allowable value for a 32-bit integer.
//
//	The maximum value for a 32-bit integer is 2,147,483,647 or
//	2^31 - 1.
func (mathProcUtil *MathProcessUtility) DoesUintExceedMax32BitInt(
	candidateUintValue uint) (exceedsMaxInt bool) {

	max32BitIntBigInt := big.NewInt(math.MaxInt32)

	candidateBigInt := big.NewInt(0).SetUint64(uint64(candidateUintValue))

	if candidateBigInt.Cmp(max32BitIntBigInt) == 1 {
		return true
	}

	return false
}
