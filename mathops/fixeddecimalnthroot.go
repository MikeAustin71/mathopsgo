package mathops

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

type NthRootBeta struct {

	LastGuessIdx 		int
	LastGuessResult	int // -1 = r' less than zero; 0 = r' equals zero; 1 = r' greater than zero
	NextGuessIdx 		int
	Idx 						int
	Result          int // -1 = r' less than zero; 0 = r' equals zero; 1 = r' greater than zero
	Beta 						*big.Int
	RPrime					*big.Int
	YPrime					*big.Int

}

func (nRootBeta *NthRootBeta) CopyIn(beta2 NthRootBeta) {

	nRootBeta.LastGuessIdx 		= beta2.LastGuessIdx
	nRootBeta.LastGuessResult = beta2.LastGuessResult
	nRootBeta.NextGuessIdx 		= beta2.NextGuessIdx
	nRootBeta.Idx 						= beta2.Idx
	nRootBeta.Result          = beta2.Result
	nRootBeta.Beta 						= big.NewInt(0).Set(beta2.Beta)
	nRootBeta.RPrime					= big.NewInt(0).Set(beta2.RPrime)
	nRootBeta.YPrime					= big.NewInt(0).Set(beta2.YPrime)
	
}

func (nRootBeta NthRootBeta) New(idx, betaNum int) NthRootBeta {

	beta := NthRootBeta{}
	beta.LastGuessIdx = -1
	beta.LastGuessResult = -99
	beta.NextGuessIdx = -1
	beta.Idx 		= idx
	beta.Result = -99
	beta.Beta 	= big.NewInt(int64(betaNum))
	beta.RPrime = big.NewInt(0)
	beta.YPrime = big.NewInt(0)

	return beta
}

// FixedDecNthRootCalcFactors - Internal calculation values
// used by FixedDecimalNthRoot for nthRoot calculations.
//
type FixedDecNthRootCalcFactors struct {
	NthRoot          					*big.Int	// NthRoot as *big.Int type
	NthRootPrecision					*big.Int  // NthRoot Precision as *big.Int type
	Radicand				 					*big.Int	// The Radicand for which the nthRoot will
															 				//   be calculated
  RadicandPrecision 				*big.Int	// Precision specification for Radicand
	Root          						*big.Int	// Calculated nthRoot of radicand
	RootPrecision 						*big.Int	// Precision specification of calculated
																	 		//  nthRoot of radicand.
	MaxPrecision             	*big.Int 	// Maximum Precision for nthRoot result 'Root'
	IntRadicand              	*big.Int 	// Integer digits of the radicand
	IntRadicandTotalDigits   	*big.Int 	// Total number of digits in IntRadicand
	FmtFracRadicand          	*big.Int 	// Formatted Fractional digits of the radicand
	FmtFracRadicandPrecision 	*big.Int 	// Precision specification for FmtFracRadicand
	FracMask1                	*big.Int 	// 11x10^nthRoot
	FracMask2                	*big.Int 	// 10^(nthRoot+1)
	BBase                    	*big.Int 	// bBase = base number system (always 10)
	BPwrN                    	*big.Int 	// bBase^n
	Term1											*big.Int  // Used to calculate r' and y'
	TermBy										*big.Int	// Used to calculate r' and y'
	Term2b										*big.Int	// Used to calculate r' and y'
	Zero                     	*big.Int 	// zero
	One                      	*big.Int 	// one
	Two                      	*big.Int 	// two
	Ten                      	*big.Int 	// ten
	Eleven                   	*big.Int 	// eleven
	Betas                    	[] NthRootBeta  // beta guess array
}

// New - Creates a new FixedDecNthRootCalcFactors with all data variables set
// to their initial (zero) values.
//
func (rootCalcFacs FixedDecNthRootCalcFactors) New() FixedDecNthRootCalcFactors {

	calcFacs := FixedDecNthRootCalcFactors{}

	calcFacs.NthRoot 									= big.NewInt(0)
	calcFacs.NthRootPrecision					= big.NewInt(0)
	calcFacs.Radicand 								= big.NewInt(0)
	calcFacs.RadicandPrecision 				= big.NewInt(0)
	calcFacs.Root											= big.NewInt(0)
	calcFacs.RootPrecision						= big.NewInt(0)
	calcFacs.MaxPrecision   					= big.NewInt(0)
	calcFacs.IntRadicand							= big.NewInt(0)
	calcFacs.IntRadicandTotalDigits		= big.NewInt(0)
	calcFacs.FmtFracRadicand 					= big.NewInt(0)
	calcFacs.FmtFracRadicandPrecision = big.NewInt(0)
	calcFacs.FracMask1      					= big.NewInt(0)
	calcFacs.FracMask2     						= big.NewInt(0)
	calcFacs.BBase          					= big.NewInt(0)
	calcFacs.BPwrN          					= big.NewInt(0)
	calcFacs.Term1										= big.NewInt(0)
	calcFacs.TermBy										= big.NewInt(0)
	calcFacs.Term2b										= big.NewInt(0)
	calcFacs.Zero           					= big.NewInt(0)
	calcFacs.One            					= big.NewInt(0)
	calcFacs.Two            					= big.NewInt(0)
	calcFacs.Ten            					= big.NewInt(0)
	calcFacs.Eleven         					= big.NewInt(0)
	calcFacs.Betas          					= make([]NthRootBeta, 10)
	for i:=0; i < 10; i++ {
		calcFacs.Betas[i] = NthRootBeta{}.New(i, i)
	}

	return calcFacs
}

// NewCalcFacs - Creates, populates and returns a new FixedDecNthRootCalcFactors
// instance based on input parameter 'fixDecNthRoot'
func (rootCalcFacs FixedDecNthRootCalcFactors) NewCalcFacs(
	fixDecNthRoot *FixedDecimalNthRoot) FixedDecNthRootCalcFactors {

	calcFacs := FixedDecNthRootCalcFactors{}.New()

	calcFacs.Initialize(fixDecNthRoot)

	return calcFacs
}

// Initialize - Intializes the values of the current FixedDecNthRootCalcFactors instance
// to the values supplied by input paramter, 'fixDecNthRoot'.
func (rootCalcFacs *FixedDecNthRootCalcFactors) Initialize(fixDecNthRoot *FixedDecimalNthRoot) {
	
	fixDecNthRoot.validateCalcFactors()

	rootCalcFacs.NthRoot 									= big.NewInt(0).Set(fixDecNthRoot.NthRoot)
	rootCalcFacs.NthRootPrecision					= big.NewInt(0).Set(fixDecNthRoot.NthRootPrecision)
	rootCalcFacs.Radicand 								= big.NewInt(0).Set(fixDecNthRoot.Radicand)
	rootCalcFacs.RadicandPrecision				= big.NewInt(0).Set(fixDecNthRoot.Radicand)
	rootCalcFacs.Root											= big.NewInt(0).Set(fixDecNthRoot.Root)
	rootCalcFacs.RootPrecision						= big.NewInt(0).Set(fixDecNthRoot.RootPrecision)
	rootCalcFacs.MaxPrecision   					= big.NewInt(0).Set(fixDecNthRoot.maxPrecision)
	rootCalcFacs.IntRadicand   						= big.NewInt(0).Set(fixDecNthRoot.intRadicand)
	rootCalcFacs.IntRadicandTotalDigits 	= big.NewInt(0).Set(fixDecNthRoot.intRadicandTotalDigits)
	rootCalcFacs.FmtFracRadicand 					= big.NewInt(0).Set(fixDecNthRoot.fmtFracRadicand)
	rootCalcFacs.FmtFracRadicandPrecision = big.NewInt(0).Set(fixDecNthRoot.fmtFracRadicandPrecision)
	rootCalcFacs.FracMask1      					= big.NewInt(0).Set(fixDecNthRoot.fracMask1)
	rootCalcFacs.FracMask2     						= big.NewInt(0).Set(fixDecNthRoot.fracMask2)
	rootCalcFacs.BBase          					= big.NewInt(0).Set(fixDecNthRoot.bBase)
	rootCalcFacs.BPwrN          					= big.NewInt(0).Set(fixDecNthRoot.bPwrN)
	rootCalcFacs.Term1										= big.NewInt(0).Set(fixDecNthRoot.term1)
	rootCalcFacs.TermBy										= big.NewInt(0).Set(fixDecNthRoot.termBy)
	rootCalcFacs.Term2b										= big.NewInt(0).Set(fixDecNthRoot.term2b)
	rootCalcFacs.Zero           					= big.NewInt(0).Set(fixDecNthRoot.zero)
	rootCalcFacs.One            					= big.NewInt(0).Set(fixDecNthRoot.one)
	rootCalcFacs.Two            					= big.NewInt(0).Set(fixDecNthRoot.two)
	rootCalcFacs.Ten            					= big.NewInt(0).Set(fixDecNthRoot.ten)
	rootCalcFacs.Eleven         					= big.NewInt(0).Set(fixDecNthRoot.eleven)

	limit := len(fixDecNthRoot.betas)

	rootCalcFacs.Betas = make([]NthRootBeta, limit)

	for i:=0; i < limit; i++ {
		rootCalcFacs.Betas[i].CopyIn(fixDecNthRoot.betas[i])
	}

}

type FixedDecimalNthRoot struct {
	OriginalNthRoot  					BigIntFixedDecimal
	NthRoot          					*big.Int	// NthRoot as *big.Int type
	NthRootPrecision 					*big.Int	// NthRootPrecision as *big.Int type
	Radicand				 					*big.Int	// The Radicand for which the nthRoot will
	                           					//   be calculated
  RadicandPrecision     		*big.Int	// Precision specification for Radicand
	Root                  		*big.Int	// Calculated nthRoot of radicand
	RootPrecision         		*big.Int	// Precision specification for Root
	maxPrecision          		*big.Int	// The maximum precision specification for
	                               			//  the final root calculation result.
	intRadicand              	*big.Int 	// Integer digits of the original radicand
	intRadicandTotalDigits   	*big.Int 	// Total number of digits in intRadicand
	fmtFracRadicand          	*big.Int 	// Formatted Fractional digits of the original radicand
	fmtFracRadicandPrecision 	*big.Int 	// Precision specification for Formatted Fractional
	                               			//  digits of the original radicand
	fracMask1             		*big.Int	// 11x10^nthRoot
	fracMask2             		*big.Int	// 10^(nthRoot+1)
	bBase                 		*big.Int	// bBase = base number system (always 10)
	bPwrN                 		*big.Int	// bBase^n
	term1											*big.Int  // Used to calculate r' and y'
	termBy                    *big.Int  // Used to calculate r' and y'
	term2b										*big.Int  // Used to calculate r' and y'
	zero                  		*big.Int	// zero
	one                   		*big.Int	// one
	two                   		*big.Int	// two
	ten                   		*big.Int	// ten
	eleven                		*big.Int	// eleven
	betas                 		[] NthRootBeta // Array of beta guesses
}

func (fdNthRoot FixedDecimalNthRoot) BigIntFixedDecNthRoot(
	radicand,
	nthRoot BigIntFixedDecimal,
	maxPrecision uint) (root *big.Int, rootPrecision uint, err error) {

	root = big.NewInt(0)
	rootPrecision = 0
	err = nil


	return root, rootPrecision, err
}


/*
***********************************************************************************
***********************************************************************************
***********************************************************************************
			Nth Root Calculation Single Point of Entry
***********************************************************************************
***********************************************************************************
***********************************************************************************
*/

func (fdNthRoot FixedDecimalNthRoot) GetNthRoot(
	radicand *big.Int,
	radicandPrecision *big.Int,
	nthRoot *big.Int,
	nthRootPrecision *big.Int,
	maxPrecision *big.Int) (
													result *big.Int,
													resultPrecision *big.Int,
													err error) {

	ePrefix := "FixedDecimalNthRoot.GetNthRoot() "

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil
	bigZero := big.NewInt(0)

	if radicand.Cmp(bigZero) == 0 {
		return result, resultPrecision, err
	}

	if radicand.Cmp(big.NewInt(1)) == 0 {
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	if radicandPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error 'radicandPrecision' is Less Than Zero! " +
			"radicandPrecision='%v'", radicandPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtPrecisionCmpZero := nthRootPrecision.Cmp(bigZero)

	if nthRtPrecisionCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Less Than Zero! " +
			"nthRootPrecision='%v'", radicandPrecision.Text(10))
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error 'maxPrecision' is Less Than Zero! " +
			"maxPrecision='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtCmpZero := nthRoot.Cmp(bigZero)

	if nthRtCmpZero == 0 {
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	if nthRoot.Cmp(big.NewInt(1)) == 0 {
		result = big.NewInt(0).Set(radicand)
		resultPrecision = big.NewInt(0).Set(radicandPrecision)
	}

	var errx error

	if nthRtCmpZero == 1 &&
			nthRtPrecisionCmpZero == 0 {
			// nthRoot is > 0 and nthRootPrecision is 0
			// nthRoot is a positive integer value

			result, resultPrecision, errx =
				fdNthRoot.CalculatePositiveIntegerNthRoot(
				radicand,
				radicandPrecision,
				nthRoot,
				nthRootPrecision,
				maxPrecision)

	}

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"%v", errx.Error())

		return result, resultPrecision, err
	}


	return result, resultPrecision, err
}



/*
***********************************************************************************
***********************************************************************************
***********************************************************************************
			Nth Root Calculation Classification By Type
***********************************************************************************
***********************************************************************************
***********************************************************************************
*/


// CalculatePositiveIntegerNthRoot - Calculates roots for nthRoots which
// are positive integer values.
func (fdNthRoot *FixedDecimalNthRoot) CalculatePositiveIntegerNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) (result *big.Int,
													resultPrecision *big.Int ,
													err error) {

	ePrefix := "FixedDecimalNthRoot.CalculatePositiveIntegerNthRoot() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	bigZero := big.NewInt(0)

	radicandPrecisionZeroCmp := radicandPrecision.Cmp(bigZero)

	if radicandPrecisionZeroCmp == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: 'radicandPrecision' is Less Than Zero! " +
			"radicandPrecision='%v'", radicandPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtPrecisionCmpZero := nthRootPrecision.Cmp(bigZero)

	if nthRtPrecisionCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Less Than Zero! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	if nthRtPrecisionCmpZero == 1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Greater Than Zero! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	radicandZeroCmp := radicand.Cmp(bigZero)

	if radicandZeroCmp == 0 {
		return result, resultPrecision, err
	}

	if radicand.Cmp(big.NewInt(1)) == 0 {
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	if radicandZeroCmp == -1 {
		// Cannot calculate root of a negative radicand when
		// nthRoot is even.
		remainder := big.NewInt(0).Rem(nthRoot, big.NewInt(2))

		if remainder.Cmp(bigZero) == 0 {

			err = errors.New(ePrefix +
				"INVALID ENTRY - Cannot calculate nthRoot of a negative radicand when nthRoot is even.")
			return result, resultPrecision, err

		}
	}

	// nthRoot precision must be zero. This is an integer nthRoot

	if nthRoot.Cmp(big.NewInt(1)) == 0 {
		// nthRoot is 1, so result = radicand
		result = big.NewInt(0).Set(radicand)
		resultPrecision = big.NewInt(0).Set(radicandPrecision)
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error 'maxPrecision' is Less Than Zero! " +
			"maxPrecision='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtCmpZero := nthRoot.Cmp(bigZero)

	if nthRtCmpZero == 0 {
		err = errors.New(ePrefix +
			"Error: 'nthRoot' is Zero!")
		return result, resultPrecision, err
	}

	if nthRtCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRoot' is negative! " +
			"nthRoot='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	errx := fdNthRoot.FormatCalculationConstants(
		radicand,
		radicandPrecision,
		nthRoot,
		nthRootPrecision,
		maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		return result, resultPrecision, err
	}

	result, resultPrecision, errx = fdNthRoot.CalculateRoot()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		return result, resultPrecision, err
	}

	err = nil

	return result, resultPrecision, err
}


// CalculatePositiveFractionalNthRoot - Calculates roots for positive decimal
// value or fractional nthRoots.
//
func (fdNthRoot *FixedDecimalNthRoot) CalculatePositiveFractionalNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) (
														result *big.Int,
														resultPrecision *big.Int,
														err error) {

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.CalculatePositiveFractionalNthRoot() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	bigZero := big.NewInt(0)

	radicandPrecisionZeroCmp := radicandPrecision.Cmp(bigZero)

	if radicandPrecisionZeroCmp == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: 'radicandPrecision' is Less Than Zero! " +
			"radicandPrecision='%v'", radicandPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtPrecisionCmpZero := nthRootPrecision.Cmp(bigZero)

	if nthRtPrecisionCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is NEGATIVE! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	if nthRtPrecisionCmpZero == 0 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Zero. nthRoot is an integer! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	// nthRootPrecision Must Be > 0

	radicandZeroCmp := radicand.Cmp(bigZero)

	if radicandZeroCmp == 0 {
		// if radicand == 0 then result == 0
		return result, resultPrecision, err
	}

	if radicand.Cmp(big.NewInt(1)) == 0 {
		result = big.NewInt(1)
		return result, resultPrecision, err
	}


	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error 'maxPrecision' is Less Than Zero! " +
			"maxPrecision='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtCmpZero := nthRoot.Cmp(bigZero)

	if nthRtCmpZero == 0 {
		err = errors.New(ePrefix +
			"Error: 'nthRoot' is Zero!")
		return result, resultPrecision, err
	}

	if nthRtCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRoot' is negative! " +
			"nthRoot='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	// nthRoot is positive and fractional
	// 5 ^ 2/3


	internalPrecision := big.NewInt(0).Add(maxPrecision, big.NewInt(1000))

	rat := big.NewRat(1, 1).SetFrac(nthRoot, nthRootPrecision)

	tempFactor, tempFactorPrecision, errx :=
		BigIntMathPower{}.BigIntPwr(
			radicand,
			radicandPrecision,
			rat.Num(),
			big.NewInt(0),
			internalPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())

		return result, resultPrecision, errx
	}

	result, resultPrecision, errx =
		FixedDecimalNthRoot{}.CalculatePositiveIntegerNthRoot(
			tempFactor,
			tempFactorPrecision,
			rat.Denom(),
			big.NewInt(0),
			maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		return result, resultPrecision, errx
	}

	err = nil

	return result, resultPrecision, err
}


// CalculateNegativeFractionalNthRoot - Calculates roots for negative decimal
// value or fractional nthRoots.
//
func (fdNthRoot *FixedDecimalNthRoot) CalculateNegativeFractionalNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) (result *big.Int,
													resultPrecision *big.Int,
													err error) {

	ePrefix := "FixedDecimalNthRoot.CalculateNegativeFractionalNthRoot() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	bigZero := big.NewInt(0)

	radicandPrecisionZeroCmp := radicandPrecision.Cmp(bigZero)

	if radicandPrecisionZeroCmp == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: 'radicandPrecision' is Less Than Zero! " +
			"radicandPrecision='%v'", radicandPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtPrecisionCmpZero := nthRootPrecision.Cmp(bigZero)

	if nthRtPrecisionCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Less Than Zero! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error 'maxPrecision' is Less Than Zero! " +
			"maxPrecision='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	if nthRtPrecisionCmpZero == 0 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Zero. nthRoot is an integer value!" +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	radicandZeroCmp := radicand.Cmp(bigZero)

	if radicandZeroCmp == 0 {
		// If radicand is zero, result = zero
		return result, resultPrecision, err
	}

	if radicand.Cmp(big.NewInt(1)) == 0 {
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	nthRtCmpZero := nthRoot.Cmp(bigZero)

	if nthRtCmpZero == 0 {
		err = errors.New(ePrefix +
			"Error: 'nthRoot' is Zero!")
		return result, resultPrecision, err
	}

	if nthRtCmpZero == 1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRoot' is positive! " +
			"nthRoot='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	// nthRoot precision must be positive. This is a fractional nthRoot
	// with a negative value.

	// Example: 5^-0.666 = 1/ 5^0.666

	// Convert nthRoot to positive value
	tempNthRoot := big.NewInt(0).Neg(nthRoot)

	tempNthRootPrecision := big.NewInt(0).Set(nthRootPrecision)

	internalPrecision := big.NewInt(0).Add(maxPrecision, big.NewInt(500))

	// calculate positive nthRoot Solution
	tempRoot, tempRootPrecision, errx :=
		FixedDecimalNthRoot{}.CalculatePositiveFractionalNthRoot(
			radicand,
			radicandPrecision,
			tempNthRoot,
			tempNthRootPrecision,
			internalPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		return result, resultPrecision, err
	}


	result, resultPrecision, errx =
		BigIntMathDivide{}.BigIntFracQuotient(
			big.NewInt(1),
			big.NewInt(0),
			tempRoot,
			tempRootPrecision,
			maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		return result, resultPrecision, err
	}

	err = nil

	return result, resultPrecision, err
}

// CalculateNegativeIntegerNthRoot - Calculates roots for nthRoots which
// are negative integer values.
func (fdNthRoot *FixedDecimalNthRoot) CalculateNegativeIntegerNthRoot(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) (result *big.Int,
													resultPrecision *big.Int ,
													err error)  {

  ePrefix := "FixedDecimalNthRoot.CalculateNegativeIntegerNthRoot() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	bigZero := big.NewInt(0)

	radicandPrecisionZeroCmp := radicandPrecision.Cmp(bigZero)

	if radicandPrecisionZeroCmp == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: 'radicandPrecision' is Less Than Zero! " +
			"radicandPrecision='%v'", radicandPrecision.Text(10))
		return result, resultPrecision, err
	}

	nthRtPrecisionCmpZero := nthRootPrecision.Cmp(bigZero)

	if nthRtPrecisionCmpZero == -1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Less Than Zero! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error 'maxPrecision' is Less Than Zero! " +
			"maxPrecision='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	if nthRtPrecisionCmpZero == 1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRootPrecision' is Greater Than Zero! " +
			"nthRootPrecision='%v'", nthRootPrecision.Text(10))
		return result, resultPrecision, err
	}

	radicandZeroCmp := radicand.Cmp(bigZero)

	if radicandZeroCmp == 0 {
		return result, resultPrecision, err
	}

	if radicand.Cmp(big.NewInt(1)) == 0 {
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	if radicandZeroCmp == -1 {
		// Cannot calculate root of a negative radicand when
		// nthRoot is even.
		remainder := big.NewInt(0).Rem(nthRoot, big.NewInt(2))

		if remainder.Cmp(bigZero) == 0 {

			err = errors.New(ePrefix +
				"INVALID ENTRY - Cannot calculate nthRoot of a negative radicand when nthRoot is even.")
			return result, resultPrecision, err
		}
	}


	bigOne := big.NewInt(1)
	bigOnePrecision := big.NewInt(0)
	var errx error

	if nthRoot.Cmp(big.NewInt(-1)) == 0 {
		// nthRoot is 1, so result = 1/radicand

		result, resultPrecision, errx =
			BigIntMathDivide{}.BigIntFracQuotient(
				bigOne,
				bigOnePrecision,
				radicand,
				radicandPrecision,
				maxPrecision)

		if errx != nil {
			err = fmt.Errorf(ePrefix + "%v", errx.Error())
			result = big.NewInt(0)
			resultPrecision = big.NewInt(0)
			return result, resultPrecision, err
		}

		err = nil

		return result, resultPrecision, err
	}

	nthRtCmpZero := nthRoot.Cmp(bigZero)

	if nthRtCmpZero == 0 {
		err = errors.New(ePrefix +
			"Error: 'nthRoot' is Zero!")
		return result, resultPrecision, err
	}

	if nthRtCmpZero == 1 {
		err = fmt.Errorf(ePrefix + "Error 'nthRoot' is positive! " +
			"nthRoot='%v'", maxPrecision.Text(10))
		return result, resultPrecision, err
	}

	// nthRoot precision must be zero. This is an integer nthRoot
	// with a negative value.

	// Convert nthRoot to positive value
	tempNthRoot := big.NewInt(0).Neg(nthRoot)

	tempNthRootPrecision := big.NewInt(0).Set(nthRootPrecision)

	internalPrecision := big.NewInt(0).Add(maxPrecision, big.NewInt(500))

	// calculate positive nthRoot Solution
	tempRoot, tempRootPrecision, errx :=
		FixedDecimalNthRoot{}.CalculatePositiveIntegerNthRoot(
			radicand,
			radicandPrecision,
			tempNthRoot,
			tempNthRootPrecision,
			internalPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		return result, resultPrecision, err
	}

	result, resultPrecision, errx =
		BigIntMathDivide{}.BigIntFracQuotient(
			bigOne,
			bigOnePrecision,
			tempRoot,
			tempRootPrecision,
			maxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		return result, resultPrecision, err
	}

	err = nil

	return result, resultPrecision, err
}


/*
***********************************************************************************
***********************************************************************************
***********************************************************************************
			Entry point for Nth Root Calculation
***********************************************************************************
***********************************************************************************
***********************************************************************************
*/

func (fdNthRoot *FixedDecimalNthRoot) CalculationController(
	radicand,
	nthRoot BigIntFixedDecimal,
	maxPrecision uint64) (result BigIntFixedDecimal, err error) {


	result = BigIntFixedDecimal{}.NewZero(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.CalculationController() "

	radCmpZero := radicand.CmpZero()

	if radCmpZero == 0 {
		// radicand is zero
		// result == 0
		return result, err
	}

	// fixDecOne = +1
	fixDecOne := BigIntFixedDecimal{}.NewInt(1, 0)

	nthRootCmpZero := nthRoot.CmpZero()

	if nthRootCmpZero == 0 {
		// nthRoot == 0
		// result = 1
		result = fixDecOne.CopyOut()
		return result, err
	}

	if nthRoot.Cmp(fixDecOne)==0 {
		// if nthRoot == 1; result = radicand
		result = radicand.CopyOut()
		return result, err
	}

	fixDecOne.ChangeSign() // now -1

	if nthRoot.Cmp(fixDecOne) == 0 {
		// nthRoot is -1 ; result = inverse of radicand
		result = radicand.CopyOut()
		result.Inverse(uint(maxPrecision))
		return result, err
	}

	//       nthRoot Status
	//  nthRoot is less than -1 or nthRoot is greater than +1

	if radCmpZero == -1 {
		// radicand is less than zero

		isEvenNum := nthRoot.IsEven()

		if isEvenNum {
			// Fatal error
			err = fmt.Errorf(ePrefix+
					"INVALID ENTRY - Cannot calculate nthRoot of a negative radicand when nthRoot is even. "+
					"radicand= %v  nthRoot= %v\n", radicand.GetNumStr(), nthRoot.GetNumStr())

			return result, err
		}
	}

	if nthRootCmpZero == 1 {
		// nthRoot is greater than +1


	}


	return result, err
}





// CalculationController - Central coordinating function for the Nth Root
// Calculation
//return result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, errreturn result, resultPrecision, err

/*
func (fdNthRoot *FixedDecimalNthRoot) CalculationController(
	radicand,
	radicandPrecision,
	intRadicand,
	fracRadicand,
	fracRadicandPrecision,
	nthRoot *big.Int,
	maxPrecision uint64 ) (result BigIntFixedDecimal, err error) {

	ePrefix := "FixedDecimalNthRoot.CalculationController() "

	result = BigIntFixedDecimal{}.NewZero(0)
	err = nil


	errx := fdNthRoot.FormatCalculationConstants(
		radicand,
		radicandPrecision,
		intRadicand,
		fracRadicand,
		fracRadicandPrecision,
		nthRoot,
		maxPrecision)

	if errx != nil {
		err= fmt.Errorf(ePrefix + "%v", err)
		return result, err
	}


	result, errx = fdNthRoot.CalculateRoot()

	if errx != nil {
		result = BigIntFixedDecimal{}.NewZero(0)
		err= fmt.Errorf(ePrefix + "%v", err)
		return result, err
	}


	return result, err
}

*/

// CalculateRoot - Calculate the nthRoot of a radicand
// *** The following Method MUST BE CALLED FIRST ***
// fdNthRoot.FormatCalculationConstants()
//
func (fdNthRoot *FixedDecimalNthRoot) CalculateRoot() (result *big.Int, resultPrecision *big.Int, err error) {

		ePrefix := "FixedDecimalNthRoot.CalculateRoot() "
		result =  big.NewInt(0)
		resultPrecision = big.NewInt(0)
		err = nil

		r := big.NewInt(0)
		y := big.NewInt(0)
		rPrime := big.NewInt(0)
		yPrime := big.NewInt(0)
		nextBundle := big.NewInt(0)
		residualInteger := big.NewInt(0).Set(fdNthRoot.intRadicand)
		residualIntegerTotalDigits := big.NewInt(0).Set(fdNthRoot.intRadicandTotalDigits)
		var errx error


		for residualInteger.Cmp(fdNthRoot.zero) == 1 {

			nextBundle,
				residualInteger,
				residualIntegerTotalDigits,
				errx = fdNthRoot.GetNextIntegerBundleFromRadicand(
				residualInteger,
				residualIntegerTotalDigits)

			if errx != nil {
				err = fmt.Errorf(ePrefix + "%v", errx.Error())
				return result, resultPrecision, err
			}

			rPrime, yPrime, errx = fdNthRoot.ComputeBeta(r, nextBundle, y)
			y.Mul(y, fdNthRoot.ten)
			y.Add(y, yPrime)
			r.Set(rPrime)
		}


		residualFracNum := big.NewInt(0).Set(fdNthRoot.fmtFracRadicand)
		residualFracPrecision := big.NewInt(0).Set(fdNthRoot.fmtFracRadicandPrecision)
		maxPrecision := big.NewInt(0).Set(fdNthRoot.maxPrecision)
		maxPrecision.Add(maxPrecision,fdNthRoot.one)

		for maxPrecision.Cmp(fdNthRoot.zero) == 1 {
			// fmt.Println("      residualFracNum: ", residualFracNum.Text(10))
			// fmt.Println("residualFracPrecision: ", residualFracPrecision.Text(10))
			nextBundle,
				residualFracNum,
				residualFracPrecision,
				errx = fdNthRoot.GetNextFractionalBundleFromRadicand(
				residualFracNum,
				residualFracPrecision)

			rPrime, yPrime, errx = fdNthRoot.ComputeBeta(r, nextBundle, y)
			y.Mul(y, fdNthRoot.ten)
			y.Add(y, yPrime)
			r.Set(rPrime)

			maxPrecision.Sub(maxPrecision, fdNthRoot.one)
		}

		// Round to maxPrecision decimal places.
		y.Add(y, big.NewInt(5))
		y.Quo(y, fdNthRoot.ten)
		fdNthRoot.Root.Set(y)
		fdNthRoot.RootPrecision.Set(fdNthRoot.maxPrecision)

		uintMax := big.NewInt(0).SetUint64(math.MaxUint32)

		if fdNthRoot.maxPrecision.Cmp(uintMax) == 1 {
				err = fmt.Errorf(ePrefix +
					"Error: Requested result maximum precision exceeds maximum for uint type." +
					"Requested maxPrecision='%v'. Mixmum uint type capacity='%v'",
					fdNthRoot.maxPrecision.Text(10), uintMax.Text(10))

				return result, resultPrecision, err
		}

		result.Set(fdNthRoot.Root)
		resultPrecision.Set(fdNthRoot.maxPrecision)
		err = nil

		return result, resultPrecision, err
}


func (fdNthRoot *FixedDecimalNthRoot) ComputeBeta(
	r,
	alpha,
	y *big.Int) (
	rPrime,
	yPrime *big.Int,
	err error) {

	ePrefix := "FixedDecimalNthRoot.ComputeBeta() "

	rPrime = big.NewInt(0)
	yPrime = big.NewInt(0)
	err = nil


	fdNthRoot.term2b = big.NewInt(0).Exp(y, fdNthRoot.NthRoot, nil)
	fdNthRoot.term2b.Mul(fdNthRoot.term2b, fdNthRoot.bPwrN)

	fdNthRoot.termBy = big.NewInt(0).Mul(fdNthRoot.bBase, y)

	fdNthRoot.term1 = big.NewInt(0).Mul(fdNthRoot.bPwrN, r)
	fdNthRoot.term1.Add(fdNthRoot.term1, alpha)

	correctGuessIdx := -1
	var errx error
	errx = nil

	lastResult := fdNthRoot.GuessBeta(&fdNthRoot.betas[5])

	if lastResult == 0 {

		correctGuessIdx = 5

	} else if lastResult == 1 {

		correctGuessIdx, errx = fdNthRoot.Guess5Positive()

	} else {

		//  lastResult must be  -1
		correctGuessIdx, errx = fdNthRoot.Guess5Negative()

	}

	if errx != nil {
		err = fmt.Errorf(ePrefix + "%v", errx.Error())
		rPrime.Set(fdNthRoot.zero)
		yPrime.Set(fdNthRoot.zero)
		return rPrime, yPrime, err
	}

	rPrime.Set(fdNthRoot.betas[correctGuessIdx].RPrime)
	yPrime.Set(fdNthRoot.betas[correctGuessIdx].YPrime)
	err = nil

	return rPrime, yPrime, err
}

// Guess 5 was positive. Next guess is 7
func (fdNthRoot *FixedDecimalNthRoot) Guess5Positive() (int, error) {

	result := fdNthRoot.GuessBeta(&fdNthRoot.betas[7])

	if result == 0 {

		return 7, nil

	} else if result == 1 {

		// 7 is positive. Check 9
		if fdNthRoot.GuessBeta(&fdNthRoot.betas[9]) > -1 {
			return 9, nil
		}

		// Nine is negative. Could be 8

		if fdNthRoot.GuessBeta(&fdNthRoot.betas[8]) == -1 {
			// 8 is negative, 7 is positive.
			return 7, nil
		}

		// 8 is positive - return 8
		return 8, nil

	}
	// result must be 7 is negative check 6

	if fdNthRoot.GuessBeta(&fdNthRoot.betas[6]) > -1 {
		// 7 is negative and 6 is positive.
		return 6, nil
	}

	// Six is negative and 5 is positive. So it must be 5
	return 5, nil

}


// Guess 5 was negative - next guess is 2
func (fdNthRoot *FixedDecimalNthRoot) Guess5Negative() (int, error) {

	result := fdNthRoot.GuessBeta(&fdNthRoot.betas[2])

	if result == 0 {
		// Nailed it
		return 2, nil

	} else if result == 1 {
		// 5 was negative 2 was positive.
		// Answer could be 3 or 4
		// Guess 4

		result = fdNthRoot.GuessBeta(&fdNthRoot.betas[4])

		if result == -1 {
			// 4 was negative and 2 was positive
			// check 3

			result = fdNthRoot.GuessBeta(&fdNthRoot.betas[3])

			if result == -1 {
				// 3 was negative and 2 was positive.
				// return 2
				return 2, nil
			}

			// 4 was negative and 3 was positive
			// return 3
			return 3, nil

		}

		// 5 was negative and 4 was positive
		return 4, nil

	}

	// Guess 2 was negative
	return fdNthRoot.Guess2Negative()

}

// Guess 2 was negative - next guess is 1
func (fdNthRoot *FixedDecimalNthRoot) Guess2Negative() (int, error) {
	ePrefix := "FixedDecimalNthRoot.Guess2Negative() "
	result := fdNthRoot.GuessBeta(&fdNthRoot.betas[1])

	if result == 0 {
		return 1, nil
	} else if result == 1 {

		// 1 was positive. 2 was negative
		// return 1
		return 1, nil
	}

	// one must be negative therefore zero must be positive

	if fdNthRoot.GuessBeta(&fdNthRoot.betas[0]) == -1 {
		return -1,
		fmt.Errorf(ePrefix +
			"Error. 1 is negative and zero is negative")
	}

	return 0, nil
}


func (fdNthRoot *FixedDecimalNthRoot)GuessBeta(
	beta *NthRootBeta) int {

	term2a := big.NewInt(0).Add(fdNthRoot.termBy, beta.Beta)
	term2a.Exp(term2a, fdNthRoot.NthRoot,nil)
	term2:= big.NewInt(0).Sub(term2a, fdNthRoot.term2b)
	beta.RPrime.Sub(fdNthRoot.term1, term2)
	beta.YPrime.Set(beta.Beta)
	beta.Result = beta.RPrime.Cmp(fdNthRoot.zero)
	return beta.Result
}


// FormatCalculationConstants - Generates calculation constants to be used in
// computing the nthRoot of the target radicand.
//
// Input Parameters
// ================
//
// radicand					  *big.Int			- The radicand used for processing. These methods will
//																		compute the nthRoot of radicand.
//
// maxPrecision					uint64			- The maximum precision for the result of this nthRoot
//                                    calculation
//
// nthRoot							*big.Int		- The nthRoot must be defined as positive value greater than
// 																		or equal to '2'.
//
func (fdNthRoot *FixedDecimalNthRoot) FormatCalculationConstants(
	radicand,
	radicandPrecision,
	nthRoot,
	nthRootPrecision,
	maxPrecision *big.Int) error {

	ePrefix := "FixedDecimalNthRoot.FormatCalculationConstants() "

	scale := big.NewInt(0).Exp(big.NewInt(10), radicandPrecision, nil)

	intRadicand, fracRadicand := big.NewInt(0).QuoRem(radicand, scale, nil)

	fracRadicandPrecision := big.NewInt(0).Set(radicandPrecision)

	fdNthRoot.initializeCalcFactors()

	// Initialize NthRoot
	fdNthRoot.NthRoot = big.NewInt(0).Set(nthRoot)

	// Intitialize NthRootPrecision
	fdNthRoot.NthRootPrecision = big.NewInt(0).Set(nthRootPrecision)

	// Initialize the radicand
	fdNthRoot.Radicand = big.NewInt(0).Set(radicand)

	// Initialize radicand precision
	fdNthRoot.RadicandPrecision = big.NewInt(0).Set(radicandPrecision)

		// Initialize Root calculation result
	fdNthRoot.Root = big.NewInt(0)

	// Initialize Root calculation result precision
	fdNthRoot.RootPrecision = big.NewInt(0)

	// Initialize Root Maximum Precision
	fdNthRoot.maxPrecision = big.NewInt(0).Set(maxPrecision)

	// Initialize integer portion of radicand
	fdNthRoot.intRadicand = big.NewInt(0).Set(intRadicand)

	totDigits, err := BigIntMath{}.GetMagnitude(intRadicand)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"%v", err.Error())
	}

	totDigits.Add(totDigits, big.NewInt(1))

	// Initialize integer Radicand total number of integer digits
	fdNthRoot.intRadicandTotalDigits = big.NewInt(0).Set(totDigits)

	// Set bigZero
	fdNthRoot.zero = big.NewInt(0)

	// Set bigOne
	fdNthRoot.one = big.NewInt(1)

	// Set bigTwo
	fdNthRoot.two = big.NewInt(2)

	// Set bigTen
	fdNthRoot.ten = big.NewInt(10)

	// Set bigEleven
	fdNthRoot.eleven = big.NewInt(11)

	// Initialize fractional portion of radicand
	// with "formatted fmtFracRadicand. Also initialize
	// fmtFracRadicandPrecision

	fdNthRoot.fmtFracRadicand,
	fdNthRoot.fmtFracRadicandPrecision,
	err = fdNthRoot.FormatFractionalDigitsFromRadicand(
		fracRadicand,
		fracRadicandPrecision)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"%v", err.Error())
	}

	// fracMask1 is a calculation constant
	// used in calculating bundles of fractional
	// digits. 11x10^nthRoot
	fdNthRoot.fracMask1 =
		big.NewInt(0).Mul(
			fdNthRoot.eleven,
			big.NewInt(0).Exp(
				fdNthRoot.ten,
				fdNthRoot.NthRoot,
				nil))

	// fracMask2 is a calculation constant
	// used in calculating bundles of fractional
	// digits. 10^(nthRoot+1)
	fdNthRoot.fracMask2 = big.NewInt(0).Exp(
		fdNthRoot.ten,
		big.NewInt(0).Add(
			fdNthRoot.NthRoot,
			big.NewInt(1)), nil)


	// Setting number system base = 10
	fdNthRoot.bBase = big.NewInt(10)

	// bBase^n
	fdNthRoot.bPwrN = big.NewInt(0).Exp(fdNthRoot.ten, fdNthRoot.NthRoot, nil)


	// fdNthRoot.betas Set up by call to 
	// fdNthRoot.initializeCalcFactors() above
	   
	return nil
}

// FormatFractionalDigitsFromRadicand - Formats the fractional digits of
// the radicand for bundle allocation.
//
// The fractionalRadicand is received as a type *big.Int integer with
// an associated precision specification. Together these two parameters
// define a decimal fraction such as 0.0123456. This represents the fraction
// part of the radicand. The method then proceeds to reformat this decimal
// fraction as an integer value with a leading '11' value. For example the
// decimal fraction 0.0123456 would be sent to this method as an integer value
// of '123456' with a precision value of '7'. This value would then be reformatted
// as the integer value 11012345600. In addition to the prefix '11', notice that
// two trailing zeros were added. This is due to the requirement that all fractional
// values must be evenly divisible by the nthRoot. In this example, the nthRoot is
// '3'. Upon completion, this method will return a formattedFracInteger of '11012345600'
// and a fracPrecision of '9'.
//
// Example:
// 	decimal fraction:  0.0123456
//
// 	Input Values
//  ------------
// 	fmtFracRadicand = 123456
//  fmtFracRadicandPrecision = 7
//  nthRoot = 3
//
//  Return Values
//  -------------
//  formattedFracInteger = 11012345600 - Notice the leading integer '1' placeholder
//                                       in the return value. Aso notice that the
//                                       fractional digits have trailing zeros added
//                                       such that the fractional digits are evenly
//                                       divisible by the nthRoot.
//
// fracPrecision = 9									 	 Only the number of fractional digits in the
//                                       'formattedFracInteger' are counted. The
//                                       leading '11' integers are NOT counted.
//
// Input Parameters
// ================
// fmtFracRadicand					 	*big.Int	- The fractional numeric digits of the radicand expressed
//                            				as a type *big.Int integer value. fmtFracRadicand must be
//                                		greater than or equal to zero
//
// fmtFracRadicandPrecision 	*big.Int	- The precision specification for input parameter 'fmtFracRadicand'.
//																		Taken together, 'fmtFracRadicand' and 'fmtFracRadicandPrecision'
// 																		defined	a fixed length decimal value.
//                                		Example: fmtFracRadicand = 123456; fmtFracRadicandPrecision=7 defines
//                                         	 a value of 0.0123456.
//
//                                  	Note: If fmtFracRadicandPrecision is less than zero, an error
//                                  	will be returned.
//
// Return Values
// =============
//
// formattedFracInteger 	*big.Int	- If this function completes successfully, this return value
//                                  	will be populated with the formatted fractional integer
// 																		value. In the example above, this would be '11012345600'.
//
// fracPrecision        	*big.Int  - The precision associated with the returned
// 																		'formattedFracInteger'.
//
// err										error			- If the function completes successfully, this return value
// 																		will be set to 'nil'. If an error occurs, the returned
// 																		error instance will include an appropriate error message.
//
// ** IMPORTANT **
// This function must be called prior to calling GetNextFractionalBundleFromRadicand()
//
func (fdNthRoot *FixedDecimalNthRoot) FormatFractionalDigitsFromRadicand(
	fracRadicand,
	fracRadicandPrecision *big.Int) (formattedFracInteger, fracPrecision *big.Int, err error) {

	ePrefix := "FixedDecimalNthRoot.FormatFractionalDigitsFromRadicand() "

	formattedFracInteger = big.NewInt(0)
	fracPrecision = big.NewInt(0)
	err = nil

	if fracRadicandPrecision.Cmp(fdNthRoot.zero) == - 1{
		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'fmtFracRadicandPrecision' is less than ZERO!. "+
			"fmtFracRadicandPrecision='%v' ", fracRadicandPrecision.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	cmpFracRadicandZero := fracRadicand.Cmp(fdNthRoot.zero)


	if cmpFracRadicandZero == -1 {

		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'fmtFracRadicand' is less than ZERO!. "+
			"fmtFracRadicand='%v' ", fracRadicand.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	if cmpFracRadicandZero == 0 {

		// formattedFracInteger == 0; fracPrecision==0; err==nil
		return formattedFracInteger, fracPrecision, err
	}

	fracPrecision.Set(fracRadicandPrecision)

	scale := big.NewInt(0).Exp(fdNthRoot.ten, fracPrecision, nil)

	leadingOneOne := big.NewInt(0).Mul(fdNthRoot.eleven, scale)
	formattedFracInteger = big.NewInt(0).Add(fracRadicand, leadingOneOne)

	precisionRemainder := big.NewInt(0).Rem(fracPrecision, fdNthRoot.NthRoot)

	if precisionRemainder.Cmp(fdNthRoot.zero) == 1 {
		nthRootMinusRmdr := big.NewInt(0).Sub(fdNthRoot.NthRoot, precisionRemainder)
		scale = big.NewInt(0).Exp(fdNthRoot.ten, nthRootMinusRmdr, nil)
		formattedFracInteger.Mul(formattedFracInteger, scale)
		fracPrecision.Add(fracPrecision, nthRootMinusRmdr)
	}

	return formattedFracInteger, fracPrecision, err
}

// GetInternalCalcFactors - Returns the internal root calculation factors computed
// for this root calculation thus far.
func (fdNthRoot *FixedDecimalNthRoot) GetInternalCalcFactors() FixedDecNthRootCalcFactors {
	
	fdNthRoot.validateCalcFactors()

	return FixedDecNthRootCalcFactors{}.NewCalcFacs(fdNthRoot)
}

// GetNextFractionalBundleFromRadicand - Receives the fractional digits of a
// radicand and extracts the next value. If all the digits have been allocated
// to bundles and the value of input parameter 'fracInteger' is zero,
// this method returns 'fracIntResidue' as zero.
//
// This method assumes that input parameter 'fracInteger' containing the fractional
// digits of a radicand has been properly formatted. This means that number of
// digits in the 'fracInteger' is evenly divisible by 'nthRoot'. It also assumes
// a placeholder value of '1' as the only digit to the left of the decimal
//
// Example:
//    For actual fractional digits '123456'
//		InputValues =
//								formattedFracDigits = 1.000123456
//						    fracPrecision = 9
//								nthRoot = 3
//
//    Return Values:
//								nextBundle = 123
//                residualFracNum = 1.456
//                residualFracPrecision = 3
//
// Input Parameters
// ================
//
// fmtFracNum					*big.Int	- The OriginalRadicand fractional digits formatted with a
//                                  leading integer, '11'. Example: radicand fractional
//                                  digits, '123456' with a precision of '9', must be
//                                  formatted as '11.000123456'. If 'fmtFracNum'
//                                  is less than '11', an error will be returned.
//                                  A value of '11' signals that all fractional digits
//                                  have been processed and the next bundle will be
//                                  set to zero.
//
// fracPrecision				*big.Int	- The number of digits to the right of the decimal
//                                  place in fmtFracNum. If fracPrecision is
//                                  NOT evenly divisible by nthRoot, an error will
//                                  be returned.
//
// Return Values
// =============
//
// nextBundle						*big.Int	- The next bundle of digits to be processed for the
//                                  nthRoot calculation.
//
// residualFracNum			*big.Int	- The remaining fractional digits to be processed
//
// residualFracPrecision *big.Int	- The precision specification for 'residualFracNum.
//
func (fdNthRoot FixedDecimalNthRoot) GetNextFractionalBundleFromRadicand(
	fmtFracNum,
	fmtFracPrecision *big.Int) (
	nextBundle,
	residualFracNum,
	residualFracPrecision *big.Int,
	err error) {

	nextBundle = big.NewInt(0)
	residualFracNum = big.NewInt(0)
	residualFracPrecision = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextFractionalBundleFromRadicand() "

	if fmtFracPrecision.Cmp(fdNthRoot.zero) == 0 ||
		fmtFracNum.Cmp(fdNthRoot.eleven) == 0 {
		// If fmtFracPrecision is zero, it signals there are no more digits
		// left to process. fmtFracNum == 11 also signals
		// that there are no fractional digits to process. In either
		// of these cases, nextBundle == zero
		err = nil
		return nextBundle, residualFracNum, residualFracPrecision, err
	}

	fracPrecisionRemainder := big.NewInt(0).Rem(fmtFracPrecision, fdNthRoot.NthRoot)

	if fracPrecisionRemainder.Cmp(fdNthRoot.zero) == 1 {
		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'fmtFracPrecision' is NOT evely divisible by 'nthRoot'! "+
			"fmtFracPrecision='%v' nthRoot='%v' ",
			fmtFracPrecision.Text(10), fdNthRoot.NthRoot.Text(10))

		return nextBundle, residualFracNum, residualFracPrecision, err
	}

	// fdNthRoot.fracMask1 and fdNthRoot.fracMask2 are static and
	// were previously calculated in the call to FormatFractionalDigitsFromRadicand()

	fracMask3 := big.NewInt(0).Exp(
		fdNthRoot.ten,
		big.NewInt(0).Sub(
			fmtFracPrecision,
			fdNthRoot.NthRoot), nil)

	fracMask4 := big.NewInt(0).Mul(fdNthRoot.eleven, fracMask3)

	fracMask5 := big.NewInt(0).Exp(
		fdNthRoot.ten,
		big.NewInt(0).Add(
			fmtFracPrecision,
			fdNthRoot.one),
		nil)

	adjustedBundle := big.NewInt(0).Quo(fmtFracNum, fracMask3)
	nextBundle = big.NewInt(0).Sub(adjustedBundle, fdNthRoot.fracMask1)
	adjustedBundle2 := big.NewInt(0).Sub(adjustedBundle, fdNthRoot.fracMask2)
	adjustedBundle3 := big.NewInt(0).Mul(adjustedBundle2, fracMask3)
	adjustedFracNum1 := big.NewInt(0).Sub(fmtFracNum, adjustedBundle3)
	adjustedFracNum2 := big.NewInt(0).Add(fracMask4, adjustedFracNum1)
	residualFracNum = big.NewInt(0).Sub(adjustedFracNum2, fracMask5)
	residualFracPrecision = big.NewInt(0).Sub(fmtFracPrecision, fdNthRoot.NthRoot)
	err = nil

	return nextBundle, residualFracNum, residualFracPrecision, err
}

// GetNextIntegerBundleFromRadicand - Returns the next bundle for an integer value.
//
// Input Parameters
// ================
//
// integerNum 		*big.Int	- The integer digits which will parsed into a bundle for
//                          	purposes of the nthRoot calculation. 'integerNum' must
//                            be greater than or equal to zero.
//
// intTotalDigits	*big.Int	- The number of numeric digits comprising input parameter
//														'integerNum'. 'intTotalDigits' must be greater than or
//                            equal to zero.
//
// Return Values
// =============
//
// nextBundle							*big.Int	- The next bundle of digits to be processed for
// 																		the	nthRoot calculation.
//
// residualInteger				*big.Int	- The remaining integer digits to be processed.
//
// residualIntTotalDigits *big.Int	- The number of numeric digits comprising return
// 																		value 'residualInteger'.
//
func (fdNthRoot *FixedDecimalNthRoot) GetNextIntegerBundleFromRadicand(
	integerNum,
	intTotalDigits *big.Int) (
	nextBundle,
	residualInteger,
	residualIntTotalDigits *big.Int,
	err error) {

	nextBundle = big.NewInt(0)
	residualInteger = big.NewInt(0)
	residualIntTotalDigits = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundleFromRadicand() "

	if integerNum.Cmp(fdNthRoot.zero) == -1 {
		err = fmt.Errorf(ePrefix+"Error: Input Parameter 'integerNum' is less than zero! "+
			"integerNum='%v' ", integerNum.Text(10))

		return nextBundle, residualInteger, residualIntTotalDigits, err
	}

	if intTotalDigits.Cmp(fdNthRoot.zero) == -1 {
		err = fmt.Errorf(ePrefix+"Error: Input Parameter 'intTotalDigits' is less than zero! "+
			"intTotalDigits='%v' ", intTotalDigits.Text(10))

		return nextBundle, residualInteger, residualIntTotalDigits, err
	}

	scratch := big.NewInt(0)

	numOfBundlesInThisInteger, remainingDigits :=
		big.NewInt(0).QuoRem(intTotalDigits, fdNthRoot.NthRoot, scratch)

	cmpNumOfBundles := numOfBundlesInThisInteger.Cmp(fdNthRoot.zero)
	cmpRemainDigits := remainingDigits.Cmp(fdNthRoot.zero)

	if cmpNumOfBundles == 0 && cmpRemainDigits == 0 {
		return nextBundle, residualInteger, residualIntTotalDigits, err
	}

	nextBundleTotalDigits := big.NewInt(0)

	if cmpRemainDigits == 1 {
		// We have remaining digits.
		// or an incomplete bundle.
		// The number of digits here
		// is less than nthRoot.

		nextBundleTotalDigits.Set(remainingDigits)

	} else {
		// cmpNumOfBundles MUST BE == 1
		// The number of digits in this
		// bundle will be = nthRoot
		nextBundleTotalDigits.Set(fdNthRoot.NthRoot)
	}

	residualIntTotalDigits.Sub(intTotalDigits, nextBundleTotalDigits)
	scale := big.NewInt(0).Exp(big.NewInt(10), residualIntTotalDigits, nil)
	nextBundle, residualInteger = big.NewInt(0).QuoRem(integerNum, scale, scratch)

	err = nil
	return nextBundle, residualInteger, residualIntTotalDigits, err
}

func (fdNthRoot *FixedDecimalNthRoot) initializeCalcFactors() {

	fdNthRoot.NthRoot       						= big.NewInt(0)
	fdNthRoot.NthRootPrecision					= big.NewInt(0)
	fdNthRoot.Radicand									= big.NewInt(0)
	fdNthRoot.RadicandPrecision 				= big.NewInt(0)
	fdNthRoot.Root          						= big.NewInt(0)
	fdNthRoot.RootPrecision 						= big.NewInt(0)
	fdNthRoot.maxPrecision  						= big.NewInt(0)
	fdNthRoot.intRadicand 							= big.NewInt(0)
	fdNthRoot.intRadicandTotalDigits		= big.NewInt(0)
	fdNthRoot.fmtFracRadicand 					= big.NewInt(0)
	fdNthRoot.fmtFracRadicandPrecision 	= big.NewInt(0)
	fdNthRoot.fracMask1     						= big.NewInt(0)
	fdNthRoot.fracMask2     						= big.NewInt(0)
	fdNthRoot.bBase         						= big.NewInt(0)
	fdNthRoot.bPwrN         						= big.NewInt(0)
	fdNthRoot.term1											= big.NewInt(0)
	fdNthRoot.termBy										= big.NewInt(0)
	fdNthRoot.term2b										= big.NewInt(0)
	fdNthRoot.zero          						= big.NewInt(0)
	fdNthRoot.one           						= big.NewInt(0)
	fdNthRoot.two           						= big.NewInt(0)
	fdNthRoot.ten           						= big.NewInt(0)
	fdNthRoot.eleven        						= big.NewInt(0)
	fdNthRoot.betas         						= make([] NthRootBeta, 10)

	for i:=0; i < 10; i++ {
		fdNthRoot.betas[i] = NthRootBeta{}.New(i, i)
	}

}

func (fdNthRoot *FixedDecimalNthRoot) validateCalcFactors() {
	
	if fdNthRoot.NthRoot == nil {
		fdNthRoot.NthRoot = big.NewInt(0)
	}

	if fdNthRoot.NthRootPrecision == nil {
		fdNthRoot.NthRootPrecision = big.NewInt(0)
	}

	if fdNthRoot.Radicand == nil {
		fdNthRoot.Radicand = big.NewInt(0)
	}

	if fdNthRoot.RadicandPrecision == nil {
		fdNthRoot.RadicandPrecision = big.NewInt(0)
	}

	if fdNthRoot.Root == nil {
		fdNthRoot.Root = big.NewInt(0)
	}

	if fdNthRoot.RootPrecision == nil {
		fdNthRoot.RootPrecision = big.NewInt(0)
	}

	if fdNthRoot.maxPrecision == nil {
		fdNthRoot.maxPrecision = big.NewInt(0)
	}

	if fdNthRoot.intRadicand == nil {
		fdNthRoot.intRadicand = big.NewInt(0)
	}

	if fdNthRoot.intRadicandTotalDigits == nil {
		fdNthRoot.intRadicandTotalDigits = big.NewInt(0)
	}

	if fdNthRoot.fmtFracRadicand == nil {
		fdNthRoot.fmtFracRadicand = big.NewInt(0)
	}

	if fdNthRoot.fmtFracRadicandPrecision == nil {
		fdNthRoot.fmtFracRadicandPrecision = big.NewInt(0)
	}

	if fdNthRoot.fracMask1 == nil {
		fdNthRoot.fracMask1 = big.NewInt(0)
	}

	if fdNthRoot.fracMask2 == nil {
		fdNthRoot.fracMask2 = big.NewInt(0)
	}

	if fdNthRoot.bBase == nil {
		fdNthRoot.bBase = big.NewInt(0)
	}

	if fdNthRoot.bPwrN == nil {
		fdNthRoot.bPwrN = big.NewInt(0)
	}

	if fdNthRoot.term1 == nil {
		fdNthRoot.term1 = big.NewInt(0)
	}

	if fdNthRoot.termBy == nil {
		fdNthRoot.termBy = big.NewInt(0)
	}

	if fdNthRoot.term2b == nil {
		fdNthRoot.term2b = big.NewInt(0)
	}

	if fdNthRoot.zero == nil {
		fdNthRoot.zero = big.NewInt(0)
	}
	
	if fdNthRoot.one == nil {
		fdNthRoot.one = big.NewInt(0)
	}
	
	if fdNthRoot.two == nil {
		fdNthRoot.two = big.NewInt(0)
	}
	
	if fdNthRoot.ten == nil {
		fdNthRoot.ten = big.NewInt(0)
	}
	
	if fdNthRoot.eleven == nil {
		fdNthRoot.eleven = big.NewInt(0)
	}
	
	if len(fdNthRoot.betas) != 10 {
		
		fdNthRoot.betas = make([] NthRootBeta, 10)

		for i:=0; i < 10; i++ {
			fdNthRoot.betas[i] = NthRootBeta{}.New(i, i)
		}
		
	}
	
}
