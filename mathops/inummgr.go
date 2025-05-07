package mathops

import "math/big"

type INumMgr interface {
  GetNumStr() (string, error)
  GetBigInt() (*big.Int, error)
  GetBigIntNum() (BigIntNum, error)
  GetDecimal() (Decimal, error)
  GetIntAry() (IntAry, error)
  GetNumericSeparatorsDto() (NumericSeparatorDto, error)
  GetNumStrDto() (NumStrDto, error)
  GetPrecisionUint() (uint, error)
  GetSign() (int, error)
  IsValid(errName string) error
  IsZero() (bool, error)
  SetNumericSeparatorsToDefaultIfEmpty() error
  SetNumericSeparatorsDto(customSeparators NumericSeparatorDto) error
}
