package mathops

import "math/big"

type INumMgr interface {
	GetNumStr() string
	GetBigInt() (*big.Int, error)
	GetNumericSeparatorsDto() NumericSeparatorDto
	GetPrecisionUint() uint
	IsZero() bool
	SetNumericSeparatorsToDefaultIfEmpty()
	SetNumericSeparatorsDto(customSeparators NumericSeparatorDto) error
}
