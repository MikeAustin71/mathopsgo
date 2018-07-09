package examples

import (
	"time"
	"fmt"
)

func CodeDurationToStr(tDuration time.Duration) string {

	tMinutes := int64(0)
	tSeconds := int64(0)
	tMilliseconds := int64(0)
	tMicroseconds := int64(0)
	tNanoseconds := int64(0)

	i64TDur := int64(tDuration)
	outStr := ""

	if i64TDur >= int64(time.Minute) {

		tMinutes = i64TDur / int64(time.Minute)
		outStr += fmt.Sprintf("%v-Minutes ", tMinutes)
		i64TDur -= tMinutes
	}

  if i64TDur >= int64(time.Second) {
  	tSeconds = i64TDur / int64(time.Second)
		outStr += fmt.Sprintf("%v-Seconds ", tSeconds)
  	i64TDur -= tSeconds
	}

	if i64TDur >= int64(time.Millisecond) {
		tMilliseconds = i64TDur / int64(time.Millisecond)
		i64TDur -= tMilliseconds
	}

	if i64TDur >= int64(time.Microsecond) {
		tMicroseconds = i64TDur / int64(time.Microsecond)
		i64TDur -= tMilliseconds
	}

	tNanoseconds = i64TDur

	outStr += fmt.Sprintf("%v-Milliseconds ", tMilliseconds)

	outStr += fmt.Sprintf("%v-Microseconds ", tMicroseconds)

	outStr += fmt.Sprintf("%v-Nanoseconds ", tNanoseconds)

	return outStr
}