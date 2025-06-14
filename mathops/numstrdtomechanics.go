package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoMechanics struct {
	lock sync.Mutex
}

// addNumStrDto
//
//	Adds the value of input parameter 'n2Dto' to the current
//	value of 'n2Dto'.
func (numStrDtoMech *numStrDtoMechanics) addNumStrDto(
	n1Dto *NumStrDto,
	n2Dto *NumStrDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.addNumStrDto()",
		"")

	if err != nil {
		return err
	}

	n1Dto := nDto.CopyOut()

	nResult, err := nDto.AddNumStrs(n1Dto, n2Dto)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by nDto.AddNumStrs(n1Dto, n2Dto). "+
			"Error='%v'", err.Error())
	}

	nDto.CopyIn(nResult)

	return nil
}
