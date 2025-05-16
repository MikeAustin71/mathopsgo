package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type bigIntFixedDecElectron struct {
  lock *sync.Mutex
}

// changeSign
//
// This method will change the sign of the BigIntFixedDecimal
// numeric value passed as input parameter . If the value
// is negative, this method will change the sign to
// positive. Likewise, if the sign is currently positive,
// calling this method will change the sign to negative.
//
//	NOTE
//	====
//
//	This method does NOT test the validity of 'bigIFxDec', an
//	instance of type BigIntNum. The calling method must
//	do this!
func (bigIFdElectron *bigIntFixedDecElectron) changeSign(
  bigIFxDec *BigIntFixedDecimal,
  errPrefDto *ePref.ErrPrefixDto) error {

  if bigIFdElectron.lock == nil {
    bigIFdElectron.lock = new(sync.Mutex)
  }

  bigIFdElectron.lock.Lock()

  defer bigIFdElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntFixedDecElectron.changeSign",
    "")

  if err != nil {
    return err
  }

  if bigIFxDec == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec'",
    }
  }

  if bigIFxDec.integerNum == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec.integerNum'",
    }
  }

  if bigIFxDec.integerNum.Cmp(big.NewInt(0)) == 0 {
    return nil
  }

  bigIFxDec.integerNum.Neg(bigIFxDec.integerNum)

  return nil
}

// cmp
//
// Compares the numeric values of two BigIntFixedDecimal
// instances.
//
// If the current BigIntFixedDecimal insance value is greater than
// that of input parameter 'fd2', this method returns '1'.
//
// If the current BigIntFixedDecimal value is equal to that of
// input parameter 'fd2', the method returns '0'.
//
// If the current BigIntFixedDecimal value is less than that of
// input parameter 'fd2', the method returns '-1'.
//
//			Examples
//			========
//
//			  BigIntFixedDecimal     'fd2'      Return
//			        Value            Value      Value
//			  ------------------   ---------   -------
//			         5                 2          1
//			         5.2               5.1        1
//			         5.2               5.2        0
//			    837123.4          837123.5       -1
//			         0                 0.1       -1
//			        35.123456         40.5       -1
//			        35.123456          2.5        1
//	           -5                 5         -1
//
//		NOTE
//		====
//
//		This method does NOT test the validity of 'bigIFxDec', an
//		instance of type BigIntNum. The calling method must
//		do this!
//
//	 However, this method WILL test the validity of input
//	 parameter 'fd2' (BigIntFixedDecimal)
func (bigIFdElectron *bigIntFixedDecElectron) cmp(
  bigIFxDec *BigIntFixedDecimal,
  fd2 BigIntFixedDecimal,
  errPrefDto *ePref.ErrPrefixDto) (int, error) {

  if bigIFdElectron.lock == nil {
    bigIFdElectron.lock = new(sync.Mutex)
  }

  bigIFdElectron.lock.Lock()

  defer bigIFdElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntFixedDecElectron.cmp",
    "")

  if err != nil {
    return -1, err
  }

  if bigIFxDec == nil {

    return -1, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec'",
    }
  }

  if bigIFxDec.integerNum == nil {

    return -1, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec.integerNum'",
    }
  }

  err = fd2.IsValid(ePrefix.XCpy(
    "Testing 'fd2'").String())

  if err != nil {

    return 0, err
  }

  if bigIFxDec.precision == fd2.precision {
    return bigIFxDec.integerNum.Cmp(fd2.integerNum), nil
  }

  bigTen := big.NewInt(10)

  if fd2.precision > bigIFxDec.precision {

    delta := int64(fd2.precision - bigIFxDec.precision)

    fdValue := big.NewInt(0).Set(bigIFxDec.integerNum)

    scale := big.NewInt(0).Exp(bigTen, big.NewInt(delta), nil)

    fdValue.Mul(fdValue, scale)

    return fdValue.Cmp(fd2.integerNum), nil

  }

  // MUST BE bigIFd.precision > fd2.precision
  delta := int64(bigIFxDec.precision - fd2.precision)

  fd2Value := big.NewInt(0).Set(fd2.integerNum)

  scale := big.NewInt(0).Exp(bigTen, big.NewInt(delta), nil)

  fd2Value.Mul(fd2Value, scale)

  return bigIFxDec.integerNum.Cmp(fd2Value), nil
}

// cmpZero
//
// Compares the current BigIntFixedDecimal numeric value to Zero
// and returns an integer flag as follows:
//
//    +1 = BigIntFixedDecimal > 0
//
//     0 = BigIntFixedDecimal == 0
//
//    -1 = BigINtFixedDecimal < 0
func (bigIFdElectron *bigIntFixedDecElectron) cmpZero(
  bigIFxDec *BigIntFixedDecimal,
  errPrefDto *ePref.ErrPrefixDto) (int, error) {

  if bigIFdElectron.lock == nil {
    bigIFdElectron.lock = new(sync.Mutex)
  }

  bigIFdElectron.lock.Lock()

  defer bigIFdElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntFixedDecElectron.cmpZero",
    "")

  if err != nil {
    return -1, err
  }

  if bigIFxDec == nil {

    return -1, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec'",
    }
  }

  if bigIFxDec.integerNum == nil {

    return -1, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec.integerNum'",
    }
  }

  return bigIFxDec.integerNum.Cmp(big.NewInt(0)), nil
}
