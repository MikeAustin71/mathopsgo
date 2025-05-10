package mathops

type NumMgrTypeCode int

func (nMgrTypeValue NumMgrTypeCode) String() string {
	return NumMgrTypeCodeLabels[nMgrTypeValue]
}

const (
	BigIntNumNMgrCode NumMgrTypeCode = iota

	BigIntFixedDecimalNMgrCode

	DecimalNMgrCode

	IntAryNMgrCode

	NumStrDtoNMgrCode
)

var NumMgrTypeCodeLabels = []string{"BigIntNum", "BigIntFixedDecimal", "Decimal", "IntAry", "NumStrDto"}
