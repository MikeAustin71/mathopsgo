package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type scinotationnumMolecule struct {
	lock sync.Mutex
}

// setNumStr
//
//	Receives a properly formatted scientific notation string as an
//	input parameter and converts to the data fields of the current
//	SciNotationNum instance.
//
//	Examples of properly formatted scientific notation strings are
//	provided below:
//
//	  2.652e+8
//	  2.652E+8
//	  2.652e8
//	  2.652E8
//
//	The use of the decimal point may be customized by first setting
//	the desired decimal separator using method,
//	SciNotationNum.SetDecimalSeparatorChar().
//
//	Also, note that exponent digits must be integer numbers. Used of fractional
//	digits in the exponent will trigger an error. Example:
//
//	2.652E9.24 = ERROR fractional digits in exponent!
func (sciNotMolecule *scinotationnumMolecule) setNumStr(
	sciNotan *SciNotationNum,
	sciNotationStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	sciNotMolecule.lock.Lock()

	defer sciNotMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"scinotationnumMolecule.setNumStr",
		"")

	if err != nil {
		return err
	}

	if sciNotan == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'sciNotan'",
		}
	}

	if len(sciNotationStr) == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "len(sciNotationStr) == 0",
			ErrMessage: "Error: Input parameter 'sciNotationStr' is an EMPTY string!",
		}
	}

	sciNotElectron := new(scinotationnumElectron)

	//sciNotan.SetDecimalSeparatorIfEmpty()
	sciNotElectron.setDecimalSeparatorIfEmpty(sciNotan)

	//sciNotan.SetExponentCharIfEmpty()

	sciNotElectron.setExponentCharIfEmpty(sciNotan)

	i := strings.Index(sciNotationStr, "e")

	if i == -1 {

		i = strings.Index(sciNotationStr, "E")
	}

	if i == -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if i == -1 {",
			ErrMessage: "Error: Input parameter 'sciNotationStr' does NOT contain an\n" +
				"Exponent Character ('e' or 'E')",
		}
	}

	significandStr := sciNotationStr[:i]

	if len(significandStr) == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if len(significandStr) == 0 {",
			ErrMessage: "Error: Input parameter 'sciNotationStr' does NOT contain any\n" +
				"digits in the significand!",
		}
	}

	exponentStr := sciNotationStr[i+1:]

	if len(exponentStr) == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if len(exponentStr) == 0  {",
			ErrMessage: "Error: Input parameter 'sciNotationStr' does NOT contain any\n" +
				"digits in the exponent!",
		}

	}

	//sciNotan.SetDecimalSeparatorIfEmpty()
	sciNotElectron.setDecimalSeparatorIfEmpty(sciNotan)

	bINumSignificand := new(BigIntNum).New()

	err = bINumSignificand.SetDecimalSeparator(sciNotan.decimalSeparator)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = bINumSignificand.SetDecimalSeparator(\n" +
				"  sciNotan.decimalSeparator)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	sciNotanNumSeps := NumericSeparatorDto{
		ThousandsSeparator: ',',
		DecimalSeparator:   sciNotan.decimalSeparator,
		CurrencySymbol:     '$',
	}

	err = bINumSignificand.SetNumStr(significandStr, sciNotanNumSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = bINumSignificand.SetNumStr(\n" +
				"  significandStr, sciNotanNumSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bINumExponent, err := new(BigIntNum).NewNumStrWithNumSeps(
		exponentStr, &sciNotanNumSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bINumExponent, err := new(BigIntNum).NewNumStrWithNumSeps(\n" +
				"exponentStr, sciNotanNumSeps)",
			ErrContext: fmt.Sprintf("exponentStr= '%v'\n"+
				"sciNotanNumSeps= '%v'", exponentStr, sciNotanNumSeps.String()),
			ErrMessage: err.Error(),
		}
	}

	bINumExponentPrecision, err := bINumExponent.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bINumExponentPrecision, err := bINumExponent.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if bINumExponentPrecision > 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: " bINumExponent.GetPrecisionUint()  > 0 ",
			ErrMessage: "Error: The exponent component of the input parameter 'sciNotationStr'\n" +
				"contains fractional digits!",
		}
	}

	err = sciNotan.significand.CopyIn(&bINumSignificand)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sciNotan.significand.CopyIn(&bINumSignificand)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bINumSignificandPrecision, err := bINumSignificand.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bINumSignificandPrecision, err := bINumSignificand.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	//sciNotan.SetMantissaLength(bINumSignificandPrecision)
	sciNotElectron.setMantissaLength(sciNotan, bINumSignificandPrecision)

	err = sciNotan.exponent.CopyIn(&bINumExponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sciNotan.exponent.CopyIn(&bINumExponent)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
