package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type NumMgrContrlr struct {
	lock *sync.Mutex
}

func (nMgrContrlr *NumMgrContrlr) GetNumMgrFromNumValue(
	numericValue interface{},
	precision uint,
	numSeps NumericSeparatorDto,
	outputNMgrTypeCode NumMgrTypeCode) (numMgr INumMgr, err error) {

	if nMgrContrlr.lock == nil {
		nMgrContrlr.lock = new(sync.Mutex)
	}

	nMgrContrlr.lock.Lock()

	defer nMgrContrlr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"NumMgrContrlr.GetNumMgrFromNumValue",
		"")

	if err != nil {
		return new(BigIntNum), err
	}

	if numericValue == nil {

		return new(BigIntNum),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numericValue'",
			}
	}

	numSeps.SetDefaultsIfEmpty()

	switch outputNMgrTypeCode {

	case BigIntNumNMgrCode:

		numMgr = new(BigIntNum)

	case BigIntFixedDecimalNMgrCode:

		numMgr = new(BigIntFixedDecimal)

	case DecimalNMgrCode:

		numMgr = new(Decimal)

	case IntAryNMgrCode:

		numMgr = new(IntAry)

	case NumStrDtoNMgrCode:

		numMgr = new(NumStrDto)

	}

	return numMgr, nil
}
