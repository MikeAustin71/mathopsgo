package mathops

import (
  "fmt"
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
