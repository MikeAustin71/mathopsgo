package mathops

// IntAryStatsDto - used to transmit
// statistics and critical configuration
// data on an IntAry object.
type IntAryStatsDto struct {
	IntAryLen              int
	IntegerLen             int
	SignificantIntegerLen  int
	SignificantFractionLen int
	Precision              int
	SignVal                int
	FirstDigitIdx          int
	LastDigitIdx           int
	IsZeroValue            bool
	IsIntegerZeroValue     bool
	DecimalSeparator       rune
	ThousandsSeparator     rune
	CurrencySymbol         rune
}
