package mathops

// NumSepSymbolCode
// Used to provide a pseudo enum series for
// identifying the numeric separator symbols
// used by the 'mathops' package.

type NumSepSymbolCode int

func (nSepSymbolValue NumSepSymbolCode) String() string {
	return NumSepSymbolCodeLabels[nSepSymbolValue]
}

const (
	// DECIMALSYMBOL
	// Symbol for the separator character used to
	// separate integer and fractional segments of
	// a floating point number or curreny value.
	DECIMALSYMBOL NumSepSymbolCode = iota

	// THOUSANDSYMBOL
	// Symbol for the separator character used to
	// separate thousands in a numeric presentation
	// where the numeric value is greater than 999
	THOUSANDSYMBOL

	// CURRENCYSYMBOL
	// Symbol for the character used to designate a
	// numeric value as currency.
	CURRENCYSYMBOL
)

var NumSepSymbolCodeLabels = [...]string{"Decimal", "Thousands", "Currency"}
