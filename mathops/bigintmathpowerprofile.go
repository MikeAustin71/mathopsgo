package mathops

type BigIntMathPowerProfile struct {
	ExpoCalcTypeCode ExponentCalcTypeCode

	BaseIsZero bool

	BaseIsPlusOne bool

	BaseIsMinusOne bool

	BaseIsAbsOne bool // True if Base is +/- 1

	BaseIsInteger bool

	BaseIsNegative bool

	ExponentIsZero bool

	ExponentIsPlusOne bool

	ExponentIsMinusOne bool

	ExponentIsAbsOne bool // True if Exponent is +/- 1

	ExponentIsInteger bool

	ExponentIsNegative bool
}
