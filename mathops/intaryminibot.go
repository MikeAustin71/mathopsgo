package mathops

import (
	"fmt"
	"math"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type intAryMinibot struct {
	lock sync.Mutex
}

// pwrByTwos
//
//	Raises a *big.Int 'base', to the specified 'power' using the
//	Exponentiation by squaring algorithm.
//	See:
//	https://en.wikipedia.org/wiki/Exponentiation_by_squaring
//	https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Computation_by_powers_of_2
//
//	This method is based on revised code taken in part from Ye Lin Aung.
//	https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
//	Algorithm modified by Mike Rapp to achieve improved performance.
//
//	Input Parameters
//	=================
//
//	power                    int
//	  The input parameter 'power' may be either	a positive or
//	  negative integer.
//
//	maxResultPrecision       int
//	  'maxResultPrecision' will determine the maximum number of
//	  digits to the right of the decimal place in the result.
//
//	  Valid values are -1 and values >= zero ('0')
//
//	  Values less than -1 will trigger an error.
//
//	  A value of -1 signals that no limit will be placed on the
//	  number of decimals places to right of the decimal point in
//	  the result.
//
//	internalPrecision        int
//	  'internalPrecision' will control the number of digits of
//	  accuracy to the right of the decimal point maintained by
//	  internal multiplication operations used in raising the intAry
//	  value to the designated power.
//
//	  Valid values are -1 and values >= zero ('0')
//
//	  Values less than -1 will trigger an error.
//
//	  A value of -1 signals that no limit will be placed on	the
//	  number of decimals places to right of the decimal	point
//	  during internal multiplication operations.
//
//	Return Value
//	============
//
//	error
//	  If no errors are encountered during processing, this returned
//	  error parameter will be set to 'nil'.
func (iaMinibot *intAryMinibot) pwrByTwos(
	ia *IntAry,
	validateIa bool,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	power *big.Int,
	maxResultPrecision int,
	internalPrecision int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMinibot.lock.Lock()

	defer iaMinibot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMinibot.setInt64Exponent()",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	if power == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'power'",
		}
	}

	if maxResultPrecision < -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Parameter maxResultPrecision is less than -1.\n"+
				"maxResultPrecision= %v", maxResultPrecision),
		}
	}

	if internalPrecision < -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Parameter internalPrecision is less than -1.\n"+
				"internalPrecision= %v", internalPrecision),
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateIa,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia, 'ia', validateIa= '%v', ePrefix)", validateIa),
			ErrContext: "Valiate on Startup",
			ErrMessage: err.Error(),
		}
	}

	var finalNumSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "finalNumSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "ia"
		actualNumSepsSrcIntAryPtr = ia

	} else {

		nsProfile.SourceObjectName = "numSepsSrcIntAry"
		actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
	}

	finalNumSeps, err = new(intAryUtility).selectNumericSeparators(
		actualNumSepsSrcIntAryPtr,
		nsProfile,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "finalNumSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if internalPrecision < maxResultPrecision {
		internalPrecision = maxResultPrecision + 20
	}

	// Internal flags alread set on 'ia'

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: false,
		ValidateNumSeps:          false,
		OverrideNumSeps:          finalNumSeps,
	}

	iaHlprNeutron := new(intAryNeutron)

	if ia.isZeroValue {
		//ia.SetIntAryToZero(0)

		err = new(intAryQuark).setIntAryToZero(
			ia,
			nil,
			nsProfile2,
			uint(ia.precision),
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
					"  ia, nsProfile2, uint(ia.precision), ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	tPower := big.NewInt(0).Set(power)

	one := big.NewInt(1)

	zero := big.NewInt(0)

	two := big.NewInt(2)

	if tPower.Cmp(two) == -1 {

		if tPower.Cmp(zero) == -1 {

			ia2, err := new(intAryMechanics).inverseIntAry(ia, false, internalPrecision, ePrefix)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "ia2, err := new(intAryMechanics).inverseIntAry(\n" +
						"ia, vallidate=false, internalPrecision, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
			}

			err = new(intAryProton).copy(ia, &ia2, false, false, ePrefix)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = new(intAryProton).copy(ia, &ia2, validate=false, copyBackup=false, ePrefix)",
					ErrContext: "ia2->ia",
					ErrMessage: err.Error(),
				}
			}

			tPower = big.NewInt(0).Mul(tPower, big.NewInt(-1))

		} else if tPower.Cmp(one) == 0 {
			// no change in value. x^1 == x
			return nil
		} else if tPower.Cmp(zero) == 0 {

			err = new(intAryGluon).setIntAryWithInt(ia, nil, nsProfile2, 1, 0, true, ePrefix)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryGluon).setIntAryWithInt(\n" +
						"  ia, numSepsSrc=nil, nsProfile2, 1, 0, validateResult=true, ePrefix)",
					ErrContext: "Set 'ia' to '1' 0 Precision",
					ErrMessage: err.Error(),
				}
			}

			return nil
		}
	}

	tBase := new(IntAry).New()

	err = new(intAryProton).copy(&tBase, ia, false, false, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryProton).copy(\n" +
				"  &iAry2, ia, validateSource=false, copyToBackup=false, ePrefix)",
			ErrContext: "ia -> iAry2",
			ErrMessage: err.Error(),
		}
	}

	// Set 'ia' = 1
	err = new(intAryGluon).setIntAryWithInt(ia, nil, nsProfile2, 1, 0, true, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryGluon).setIntAryWithInt(\n" +
				"  ia, numSepsSrc=nil, nsProfile2, 1, 0, validateResult=true, ePrefix)",
			ErrContext: "Set 'ia' to '1' 0 Precision",
			ErrMessage: err.Error(),
		}
	}

	for tPower.Cmp(zero) == 1 {
		//temp, _:= intAry{}.NewNumStr("0")

		if big.NewInt(0).Mod(tPower, two).Cmp(one) == 0 {
			//temp = big.NewInt(0).Mul(result, tBase)
			//result = big.NewInt(0).Set(temp)

			err = iaHlprNeutron.multiplyThisBy(ia, false, &tBase, false, -1, internalPrecision, ePrefix)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = iaHlprNeutron.multiplyThisBy(\n"+
						"ia, validateIa=false, &tBase, validateTbase=false,\n"+
						"minPrecision=-1, internalPrecision= '%v', ePrefix)",
						internalPrecision),
					ErrContext: "Multiply 'ia' by tBase (ia x tBase)",
					ErrMessage: err.Error(),
				}
			}

			//fmt.Println("ia precision = ", ia.GetPrecisionInt())

			if tPower.Cmp(one) == 0 {

				if maxResultPrecision > -1 && maxResultPrecision < ia.GetPrecision() {

					err = new(intAryMolecule).
						setPrecision(ia, true, maxResultPrecision, true, ePrefix)

					if err != nil {

						return &FuncReturnError{
							ErrPrefix: ePrefix.String(),
							ReturnFunc: "err = new(intAryMolecule).setPrecision(\n" +
								"ia, validate=true, maxResultPrecision, roundResult=true, ePrefix)",
							ErrContext: fmt.Sprintf("maxResultPrecision= '%v'",
								maxResultPrecision),
							ErrMessage: err.Error(),
						}
					}

					return nil
				}
				// Keep going!
			}

			err = iaHlprNeutron.multiplyThisBy(&tBase, false, &tBase, false, -1, internalPrecision, ePrefix)

			if err != nil {

				return &FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: fmt.Sprintf("err = iaHlprNeutron.multiplyThisBy(\n"+
						"&tBase, validateTbase=false, &tBase, validateTbase=false,\n"+
						"minPrecision=-1, internalPrecision= '%v', ePrefix)",
						internalPrecision),
					ErrContext: "Multiply 'ia' by tBase (ia x tBase)",
					ErrMessage: err.Error(),
				}
			}
		}

		tPower = big.NewInt(0).Div(tPower, two)
	}

	if maxResultPrecision < ia.GetPrecision() &&
		maxResultPrecision == -1 {

		err = new(intAryMolecule).
			setPrecision(ia, true, maxResultPrecision, true, ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryMolecule).setPrecision(\n" +
					"ia, validate=true, maxResultPrecision, roundResult=true, ePrefix)",
				ErrContext: fmt.Sprintf("maxResultPrecision= '%v'",
					maxResultPrecision),
				ErrMessage: err.Error(),
			}
		}
	}

	// If no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		ia,
		"ia",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryUtility).selectIntAryValidation(\n" +
				"ia,\"ia\", validateResult, ePrefix)",
			ErrContext: "Validating Final Result",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

func (iaMinibot *intAryMinibot) setIntAryWithBigIntNum(
	ia *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	bINum *BigIntNum,
	validateBINum bool,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMinibot.lock.Lock()

	defer iaMinibot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMinibot.setIntAryWithBigIntNum",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	if bINum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bINum'",
		}
	}

	if validateBINum {

		err = bINum.IsValid(ePrefix.XCpy("Validating bINum").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bINum.IsValid(ePrefix.XCpy(\"Validating bINum\").String())",
				ErrContext: "Input parameter 'bINum' is invalid!\n" +
					"'bINum' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bINumStr, err := bINum.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if bINum.precision > uint(math.MaxInt) {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: Input parameter bINum has a 'precision' value\n"+
				"which exceeds the Maximum Integer (MaxInt) Value.\n"+
				"MaxInt= '%v'\n"+
				"bINum.precision= '%v'\n"+
				"bINum= '%v\n",
				math.MaxInt, bINum.precision, bINumStr),
		}
	}

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "ia"
		actualNumSepsSrcIntAryPtr = ia

	} else {

		nsProfile.SourceObjectName = "numSepsSrcIntAry"
		actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
	}

	numSeps, err = new(intAryUtility).selectNumericSeparators(
		actualNumSepsSrcIntAryPtr,
		nsProfile,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	bInt, err := bINum.GetBigInt()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bInt, err := bINum.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bigINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bigINumPrecisionUint, err := bigINum.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryGluon).setIntAryWithBigInt(
		ia, nil, nsProfile2, bInt, int(bigINumPrecisionUint), validateResult, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryGluon).setIntAryWithBigInt(\n"+
				"ia, nil, nsProfile2, bInt, int(bigINumPrecisionUint), validateResult='%v', ePrefix)",
				validateResult),
			ErrContext: fmt.Sprintf("bINum= '%v' bInt= '%v'",
				bINumStr, bInt.Text(10)),
		}
	}

	return nil
}

// setIntAryInt64Exponent
//
//	Reconfigures an IntAry instance passed in as input parameter
//	'ia'. The new numeric value is calculated using an int64 value
//	multiplied by 10 raised to the power of the 'exponent' parameter.
//
//	    numeric value = int64 X 10^exponent
//
//	Examples
//	========
//
//		uint64Num      exponent      IntAry Result
//
//		  123456          -3           123.456
//		  123456           3           123456.000
//		  123456           0           123456
//
//	IMPORTANT
//	=========
//
//	In practice, the maximum limits for 'int64Num' and
//	'exponent' will be constrained by the maximum array size
//	permitted by your system. Type IntAry relies on arrays of
//	8-bit integers for numeric value storage.
//
//	Input Parameters
//	================
//
//	ia                       *IntAry
//	  This instance of IntAry will be overwritten with the new
//	  IntAry values computed by this method.
//
//	numSepsSrcIntAry         *IntAry
//	  This instance of IntAry may be populated by the calling
//	  function as a source of Numeric Separators
//
//	nsProfile                NumSepsProfileSelection
//	  This struct contains decision parameters for selecting and
//	  configuring Numeric Separators.
//
//	int64Num                 int64
//	  The numeric digits contained in this value comprise both
//		 the integer digits and the fractional digits which will be
//		 configured in the final numeric value stored in the IntAry
//		 object returned by this method.
//
//	exponent                 int
//	  'exponent' specifies the number of fractional digits in the
//		 final numeric value stored in the returned IntAry object.
//	  See the examples above.
//
//	validateResult           bool
//	  When set to 'true', the final numeric value calculated by
//	  this method will be subjected to validation testing.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered during processing, this returned
//		 value will be set to 'nil'
func (iaMinibot *intAryMinibot) setIntAryInt64Exponent(
	ia *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	int64Num int64,
	exponent int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMinibot.lock.Lock()

	defer iaMinibot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMinibot.setIntAryInt64Exponent()",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "ia"
		actualNumSepsSrcIntAryPtr = ia

	} else {

		nsProfile.SourceObjectName = "numSepsSrcIntAry"
		actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
	}

	numSeps, err = new(intAryUtility).selectNumericSeparators(
		actualNumSepsSrcIntAryPtr,
		nsProfile,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	if exponent > 0 {
		for i := 0; i < exponent; i++ {

			int64Num *= 10

		}
	}

	if exponent < 0 {

		exponent = exponent * -1

	}

	err = new(intAryGluon).setIntAryWithInt64(
		ia, nil, nsProfile2, int64Num, uint(exponent), validateResult, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryGluon).setIntAryWithInt64(\n" +
				"  ia, nil, nsProfile2, int64Num, uint(exponent), validateResult, ePrefix)",
			ErrContext: fmt.Sprintf("validateResult = '%v'", validateResult),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryUint64Exponent
//
//	Receives an IntAry instance ('ia') and reconfigures the
//	numeric values based on input parameters 'uint64Num',
//	'signValue' and 'exponent'.
//
//	The IntAry numeric value is configured using an uint64 value
//	multiplied by 10 raised to the power of the 'exponent'
//	parameter.
//
//		    Result Numeric Value = uint64 X 10^exponent
//
//	Be advised that the original value of input parameter 'ia'
//	will be overwritten and destroyed. The final calculation
//	result will be stored in the 'ia' IntAry instance.
//
//	Examples
//	========
//
//	uint64Num      exponent      IntAry Result
//
//	  123456          -3           123.456
//	  123456           3           123456.000
//	  123456           0           123456
//
//	IMPORTANT
//	=========
//
//	In practice, the maximum limits for 'uint64Num' and
//	'exponent' will be constrained by the maximum array size
//	permitted by your system. Type IntAry relies on arrays of
//	8-bit integers for numeric value storage.
//
//	Input Parameters
//	================
//
//	ia                       *IntAry
//	  This instance of IntAry will be overwritten with the new
//	  IntAry values computed by this method.
//
//	numSepsSrcIntAry         *IntAry
//	  This instance of IntAry may be populated by the calling
//	  function as a source of Numeric Separators
//
//	nsProfile                NumSepsProfileSelection
//	  This struct contains decision parameters for selecting and
//	  configuring Numeric Separators.
//
//	uint64Num                uint64
//	  The uint64 holds the numeric digits which will make up the
//	  returned IntAry numeric value.
//
//	signValue                int
//	  This parameter must be set to one of two possible values:
//	  +1 or -1.
//
//	  Final numeric values less than zero must be tagged with
//	  signValue= -1.
//
//	  Final numeric values greater than or equal to zero must be
//	  tagged with signValue= +1.
//
//	exponent                 int
//	  This value will be used to determine the numeric digits in
//	  'uint64Num' which will be assigned to the right of the
//	  decimal point in the final calculation result returned as
//	  IntAry instance.
//
//	validateResult           bool
//	  When set to 'true', the final numeric value calculated by
//	  this method will be subjected to validation testing.
//
//	Return Values
//	=============
//
//	error
//	  If no errors are encountered the return value for this
//	  parameter will be set to 'nil'.
func (iaMinibot *intAryMinibot) setIntAryUint64Exponent(
	ia *IntAry,
	numSepsSrcIntAry *IntAry,
	nsProfile NumSepsProfileSelection,
	uint64Num uint64,
	signValue int,
	exponent int,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMinibot.lock.Lock()

	defer iaMinibot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMinibot.setInt64Exponent()",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	var numSeps NumericSeparatorDto

	nsProfile.OutputNumSepsName = "numSeps"

	var actualNumSepsSrcIntAryPtr *IntAry

	if numSepsSrcIntAry == nil {

		nsProfile.SourceObjectName = "ia"
		actualNumSepsSrcIntAryPtr = ia

	} else {

		nsProfile.SourceObjectName = "numSepsSrcIntAry"
		actualNumSepsSrcIntAryPtr = numSepsSrcIntAry
	}

	numSeps, err = new(intAryUtility).selectNumericSeparators(
		actualNumSepsSrcIntAryPtr,
		nsProfile,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err = new(intAryUtility).selectNumericSeparators(\n" +
				"actualNumSepsSrcIntAryPtr, nsProfile, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	nsProfile2 := NumSepsProfileSelection{
		SourceObjectName:         "ia",
		OutputNumSepsName:        "numSeps",
		UseDefaultNumSeps:        false,
		SetDefaultNumSepsIfEmpty: true,
		ValidateNumSeps:          false,
		OverrideNumSeps:          numSeps,
	}

	uint64Ten := uint64(10)

	if exponent > 0 {
		for i := 0; i < exponent; i++ {

			uint64Num *= uint64Ten

		}
	}

	if exponent < 0 {

		exponent = exponent * -1

	}

	err = new(intAryGluon).setIntAryWithUint64(
		ia,
		nil,
		nsProfile2,
		uint64Num,
		signValue,
		uint(exponent),
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryGluon).setIntAryWithUint64(\n" +
				"  &iAry, ia, nsProfile2, uint64Num, signValue, uint(exponent),\n" +
				"  validateResult, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
