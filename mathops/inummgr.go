package mathops

import "math/big"

type INumMgr interface {
	GetNumStr() string
	GetBigInt() (*big.Int, error)
	GetPrecisionUint() uint
	IsZero() bool
}
