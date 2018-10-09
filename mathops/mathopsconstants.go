package mathops

import "fmt"

/*
	Math Constants
	==============

	The source code repository for math constants used by

	the 'mathops' library is located at:
			https://github.com/MikeAustin71/mathopsgo.git

	This source file mathopsconstants.go is located in directory:
		MikeAustin71/mathopsgo/mathops/mathopsconstants.go

	Overview and General Usage
  ==========================
	This source file is used to store constants used by various Types
	in the 'mathops' Library.

*/

func GetEulersNum1050() BigIntFixedDecimal {
	return eulersNumber1050.GetFixedDecimal()
}

func GetPiTo1000() BigIntFixedDecimal {
	return piNumber1000.GetFixedDecimal()
}

func GetNatLogOfTwoTo99() BigIntFixedDecimal {
	return natLogTwo99.GetFixedDecimal()
}

var eulersNumber1050 FixedDecimalReadOnly

var piNumber1000 FixedDecimalReadOnly

var natLogTwo99 FixedDecimalReadOnly

// A002162
//Decimal expansion of the natural logarithm of 2. (Formerly M4074 N1689)
// https://oeis.org/A002162
var NatLog2_99Str =
 "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"
//  1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
//           1         2         3         4         5         6         7         8         9         0
//                                                                                                     1
/*
6, 9, 3, 1, 4, 7, 1, 8, 0, 5, 5, 9, 9, 4, 5, 3, 0, 9, 4, 1, 7, 2, 3, 2, 1, 2, 1, 4, 5, 8, 1, 7, 6, 5, 6, 8, 0, 7, 5, 5, 0, 0, 1, 3, 4, 3, 6, 0, 2, 5, 5, 2, 5, 4, 1, 2, 0, 6, 8, 0, 0, 0, 9, 4, 9, 3, 3, 9, 3, 6, 2, 1, 9, 6, 9, 6, 9, 4, 7, 1, 5, 6, 0, 5, 8, 6, 3, 3, 2, 6, 9, 9, 6, 4, 1, 8, 6, 8, 7
*/

// Pi to 1000 digits
// St Andrews University
// http://www-groups.dcs.st-and.ac.uk/history/
// http://www-groups.dcs.st-and.ac.uk/history/HistTopics/1000_places.html
// 50-digits per row
var Pi1000Str = "3." +
"14159265358979323846264338327950288419716939937510" +
"58209749445923078164062862089986280348253421170679" +
"82148086513282306647093844609550582231725359408128" +
"48111745028410270193852110555964462294895493038196" +
"44288109756659334461284756482337867831652712019091" +
"45648566923460348610454326648213393607260249141273" +
"72458700660631558817488152092096282925409171536436" +
"78925903600113305305488204665213841469519415116094" +
"33057270365759591953092186117381932611793105118548" +
"07446237996274956735188575272489122793818301194912" +
"98336733624406566430860213949463952247371907021798" +
"60943702770539217176293176752384674818467669405132" +
"00056812714526356082778577134275778960917363717872" +
"14684409012249534301465495853710507922796892589235" +
"42019956112129021960864034418159813629774771309960" +
"51870721134999999837297804995105973173281609631859" +
"50244594553469083026425223082533446850352619311881" +
"71010003137838752886587533208381420617177669147303" +
"59825349042875546873115956286388235378759375195778" +
"18577805321712268066130019278766111959092164201989"

// Arranged 75-digits per line
// Source: http://www-history.mcs.st-and.ac.uk/HistTopics/e_10000.html
// 10,000 digits of Euler's integerNum are provided. Only, 1,000 digits
// to the right of the decimal place are used here.
// First line 75-digits to the right of the decimal place
// Total 14-lines @ 75-decimal digits each = 1050-decimal digits to the
// right of the decimal place.
//
var EulersNum1050Str =
"2.718281828459045235360287471352662497757247093699959574966967627724076630353" +
  "547594571382178525166427427466391932003059921817413596629043572900334295260" +
  "595630738132328627943490763233829880753195251019011573834187930702154089149" +
  "934884167509244761460668082264800168477411853742345442437107539077744992069" +
  "551702761838606261331384583000752044933826560297606737113200709328709127443" +
  "747047230696977209310141692836819025515108657463772111252389784425056953696" +
  "770785449969967946864454905987931636889230098793127736178215424999229576351" +
  "482208269895193668033182528869398496465105820939239829488793320362509443117" +
  "301238197068416140397019837679320683282376464804295311802328782509819455815" +
  "301756717361332069811250996181881593041690351598888519345807273866738589422" +
  "879228499892086805825749279610484198444363463244968487560233624827041978623" +
  "209002160990235304369941849146314093431738143640546253152096183690888707016" +
  "768396424378140592714563549061303107208510383750510115747704171898610687396" +
  "965521267154688957035035402123407849819334321068170121005627880235193033225"



func init() {
	InitializeEulerNum1050()
	InitializePi1000()
	InitializeNatLogTwo()
}

func InitializeEulerNum1050() {

  ePrefix := "mathopsconstants.go eulersNumber1050 Initialization Failed! "



	xEuler, err := BigIntFixedDecimal{}.NewNumStr(EulersNum1050Str)

	if err != nil {
		erx :=
			fmt.Sprintf(ePrefix +
				" Error='%v'", err.Error())

		panic(erx)
	}


	eulersNumber1050 = FixedDecimalReadOnly{}.NewFixedDecimal(xEuler)
	if !eulersNumber1050.IsValid() {
		erx := ePrefix + "eulersNumber1050 object INVALID!"
		panic(erx)
	}

}

func InitializePi1000() {

	ePrefix := "mathopsconstants.go eulersNumber1050 Initialization Failed! "


	xPiNum, err := BigIntFixedDecimal{}.NewNumStr(Pi1000Str)

	if err != nil {
		erx :=
			fmt.Sprintf(ePrefix +
				" Error='%v'", err.Error())

		panic(erx)
	}

	piNumber1000 = FixedDecimalReadOnly{}.NewFixedDecimal(xPiNum)
	if !piNumber1000.IsValid() {
		erx := ePrefix + "piNumber1000 INVALID!"
		panic(erx)
	}

}

func InitializeNatLogTwo() {

	ePrefix := "mathopsconstants.go Natural Log of 2 Initialization Failed! "

	xNatLog2, err := BigIntFixedDecimal{}.NewNumStr(NatLog2_99Str)

	if err != nil {
		erx :=
			fmt.Sprintf(ePrefix +
				" Error='%v'", err.Error())

		panic(erx)
	}

	natLogTwo99 = FixedDecimalReadOnly{}.NewFixedDecimal(xNatLog2)
	if !natLogTwo99.IsValid() {
		erx := ePrefix + "natLogTwo99 INVALID!"
		panic(erx)
	}

}

// Source Currency Info
// https://gist.github.com/bzerangue/5484121
// http://symbologic.info/currency.htm
// http://www.xe.com/symbols.php
//
// NumStrCurrencySymbols - an array containing
// most of the world's major currency symbols
// stored as type 'rune'.
//
var NumStrCurrencySymbols = []rune{
	'\U00000024', // Australia Dollar 								 0
	'\U00008236', // Brazil Real											 1
	'\U00000024', // Canada Dollar 										 2
	'\U000000a5', // China Yuan												 3
	'\U00000024', // Colombia Peso										 4
	'\U0004b10d', // Czech Republic Koruna						 5
	'\U000000a3', // Egypt Pound											 6
	'\U000020ac', // Euro €  													 7
	'\U00070116', // Hungary Forint										 8
	'\U00107114', // Iceland Krona										 9
	'\U00082112', // Indonesia Rupiah									10
	'\U000020aa', // Israel Shekel  									11
	'\U000000a5', // Japan Yen  											12
	'\U000020a9', // Korea Won  											13
	'\U0000524d', // Malaysia Ringgit									14
	'\U00000024', // Mexico Peso  										15
	'\U00006b72', // Norway Krone											16
	'\U00000192', // Netherlands Antilles Guilder			17
	'\U000020a8', // Pakistan Rupee 									18
	'\U000020bd', // Russian Ruble  									19
	'\U0000fdfc', // Saudi Arabia Riyal 							20
	'\U00000082', // South Africa Rand								21
	'\U00006b72', // Sweden Krona											22
	'\U000020a3', // Switzerland Franc								23
	'\U00000024', // Taiwan NewBigIntNum Dollar								24
	'\U000020ba', // TURKISH LIRA											25
	'\U00066115', // Venezuela Bolivar								26
	'\U00008363', // Viet Nam Dong										27
	'\U00000024', // United States Dollar  						28
	'\U000000a3', // United Kingdom Pound (£)					29
	'\U000020a3', // French Franc  						        30
	'\U000020a4', // Italy Lira  						          31
	'\U000020bf', // Bitcoin  						            32
	'\U000000a2'} // United States Cent		            33

// PrecisionScaleMode - Specifies the scaling mode used in
// setting precision or the number of digits to the right
// of the decimal place.
type PrecisionScaleMode int

func (scaleMode PrecisionScaleMode) String() string {

	return PrecisionScaleModeLabels[scaleMode]
}

const (
	SCALEPRECISIONRIGHT PrecisionScaleMode = iota

	SCALEPRECISIONLEFT
)

var PrecisionScaleModeLabels = [...]string{"ScalePrecisionRight", "ScalePrecisionLeft"}

// NumStrFmtMode - Designates the type of number string formatting
// applied when converting a number to a string.
type NumStrFmtMode int

func (nstrFmtMode NumStrFmtMode) String() string {
	return NumStrFmtModeLabels[nstrFmtMode]
}

const (

	// PUREINTEGERFMT - Specifies a pure number string with no decimal
	// point, no thousands separators and no currency symbol.
	// Example: 123456789
	//
	PUREINTEGERFMT NumStrFmtMode = iota

	// INTSTRDECIMALFMT - Specifies an integer string, decimal point and
	// fractional digits. No currency symbol or thousands separator are
	// injected in to the final number string.
	// Example: 12345.678
	//
	INTSTRDECIMALFMT

	// THOUSANDSNUMSTRFMT - Specifies that the output number string will
	// have a decimal point separating fractional digits. Integer numbers
	// to the left of the decimal point will have a thousands separator
	// character injected after every third character.
	// Example: 123,456,789.23
	//
	THOUSANDSNUMSTRFMT

	// CURRENCYNUMSTRFMT - Specifies a Currency String. The output number string
	// will include a currency symbol, thousands separators and a decimal point.
	CURRENCYNUMSTRFMT
)

var NumStrFmtModeLabels = [...]string{"PureIntegerString", "IntegerDecimalString", "ThousandsNumString", "CurrencyNumString"}

type NegativeValueFmtMode int

func (negValFmtMode NegativeValueFmtMode) String() string {
	return NegativeValueFmtModeLabels[negValFmtMode]
}

const (

	// LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
	//													 		a leading minus sign.
	//															Example: -123456.78
	//
	LEADMINUSNEGVALFMTMODE NegativeValueFmtMode = iota

	// PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
	//														surrounding parentheses.
	//														Example: (123456.78)
	//
	PARENTHESESNEGVALFMTMODE

	// ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
	//														 absolute (positive) integer value
	//														 and no decimal point separator.
	//														Example: (12345678)
	ABSOLUTEPURENUMSTRFMTMODE
)

var NegativeValueFmtModeLabels = [...]string{"LeadingMinusSign", "SurroundingParentheses", "AbsolutePureNumberString"}

// Numeric Separator Data Transfer Object
// Used to transmit symbols used for decimal point,
// thousands separator and currency symbol.
//
type NumericSeparatorDto struct {
	DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
	ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
	CurrencySymbol     rune // Currency Symbol
}

// Equal - Compares two NumericSeparatorDto's and returns 'true' if they
// are equivalent.
//
func (numSep *NumericSeparatorDto) Equal(numSep2 NumericSeparatorDto) bool {

	if numSep.DecimalSeparator != numSep2.DecimalSeparator {
		return false
	}

	if numSep.ThousandsSeparator != numSep2.ThousandsSeparator {
		return false
	}

	if numSep.CurrencySymbol != numSep2.CurrencySymbol {
		return false
	}

	return true
}

// New - Returns a new instance of NumericSeparatorDto. The
// rune values are automatically set to USA defaults.
func (numSep NumericSeparatorDto) New() NumericSeparatorDto {

	n2 := NumericSeparatorDto{}

	n2.SetDefaultsIfEmpty()

	return n2
}

// SetDefaultsIfEmpty - If any of the NumericSeparatorDTo rune values
// are zero, this method will set those elements to USA default values.
//
func (numSep *NumericSeparatorDto) SetDefaultsIfEmpty() {

	if numSep.DecimalSeparator == 0 {
		numSep.DecimalSeparator = '.'
	}

	if numSep.ThousandsSeparator == 0 {
		numSep.ThousandsSeparator = ','
	}

	if numSep.CurrencySymbol == 0 {
		numSep.CurrencySymbol = '$'
	}
}

// String - Provides a formatted listing of the contents from the current
// NumericSeparatorDto instance.
//
func (numSep *NumericSeparatorDto) String() string {
	return fmt.Sprintf("Decimal Separator: %q  Thousands Separator: %q  Currency Symbol: %q",
		numSep.DecimalSeparator, numSep.ThousandsSeparator, numSep.CurrencySymbol)
}
