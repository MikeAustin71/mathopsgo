package mathops

import (
  "testing"
)

func TestStrMathOp_AddN1N2_01(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_01"

  inputNumStr1 := "457.3"

  inputNumStr2 := "22.2"

  expectedFinalNumStr := "479.5"

  expectedNRunes := []rune("4795")

  expectedIntArray := []int{4, 7, 9, 5}

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  expectedPrecisionInt := 1

  expectedSignValue := 1

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_02(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_02"

  inputNumStr1 := "457.325"

  inputNumStr2 := "-22.2"

  expectedFinalNumStr := "435.125"

  expectedNRunes := []rune("435125")

  expectedIntArray := []int{4, 3, 5, 1, 2, 5}

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_03(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_03"

  inputNumStr1 := "-457.325"

  inputNumStr2 := "22.2"

  expectedFinalNumStr := "-435.125"

  expectedNRunes := []rune("435125")

  expectedIntArray := []int{4, 3, 5, 1, 2, 5}

  expectedPrecisionInt := 3

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_04(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_04"

  inputNumStr1 := "-457.325"

  inputNumStr2 := "-22.2"

  expectedFinalNumStr := "-479.525"

  expectedNRunes := []rune("479525")

  expectedIntArray := []int{4, 7, 9, 5, 2, 5}

  expectedPrecisionInt := 3

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_05(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_05"

  inputNumStr1 := "0.000"

  inputNumStr2 := "-22.2"

  expectedFinalNumStr := "-22.200"

  expectedNRunes := []rune("22200")

  expectedIntArray := []int{2, 2, 2, 0, 0}

  expectedPrecisionInt := 3

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_06(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_06"

  inputNumStr1 := "0.000"

  inputNumStr2 := "0"

  expectedFinalNumStr := "0.000"

  expectedNRunes := []rune("0000")

  expectedIntArray := []int{0, 0, 0, 0}

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_07(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_07"

  inputNumStr1 := "99.225"

  inputNumStr2 := "-99.1"

  expectedFinalNumStr := "0.125"

  expectedNRunes := []rune("0125")

  expectedIntArray := []int{0, 1, 2, 5}

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_08(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_08"

  inputNumStr1 := "350"

  inputNumStr2 := "122"

  expectedFinalNumStr := "472"

  expectedNRunes := []rune("472")

  expectedIntArray := []int{4, 7, 2}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_09(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_09"

  inputNumStr1 := "-350"

  inputNumStr2 := "122"

  expectedFinalNumStr := "-228"

  expectedNRunes := []rune("228")

  expectedIntArray := []int{2, 2, 8}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_10(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_10"

  inputNumStr1 := "-350"

  inputNumStr2 := "-122"

  expectedFinalNumStr := "-472"

  expectedNRunes := []rune("472")

  expectedIntArray := []int{4, 7, 2}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_11(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_11"

  inputNumStr1 := "350"

  inputNumStr2 := "-122"

  expectedFinalNumStr := "228"

  expectedNRunes := []rune("228")

  expectedIntArray := []int{2, 2, 8}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_12(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_12"

  inputNumStr1 := "350"

  inputNumStr2 := "0"

  expectedFinalNumStr := "350"

  expectedNRunes := []rune("350")

  expectedIntArray := []int{3, 5, 0}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_13(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_13"

  inputNumStr1 := "-350"

  inputNumStr2 := "0"

  expectedFinalNumStr := "-350"

  expectedNRunes := []rune("350")

  expectedIntArray := []int{3, 5, 0}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_14(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_14"

  inputNumStr1 := "122"

  inputNumStr2 := "350"

  expectedFinalNumStr := "472"

  expectedNRunes := []rune("472")

  expectedIntArray := []int{4, 7, 2}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_15(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_15"

  inputNumStr1 := "-122"

  inputNumStr2 := "350"

  expectedFinalNumStr := "228"

  expectedNRunes := []rune("228")

  expectedIntArray := []int{2, 2, 8}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_16(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_16"

  inputNumStr1 := "-122"

  inputNumStr2 := "-350"

  expectedFinalNumStr := "-472"

  expectedNRunes := []rune("472")

  expectedIntArray := []int{4, 7, 2}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_17(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_17"

  inputNumStr1 := "122"

  inputNumStr2 := "-350"

  expectedFinalNumStr := "-228"

  expectedNRunes := []rune("228")

  expectedIntArray := []int{2, 2, 8}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_18(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_18"

  inputNumStr1 := "0"

  inputNumStr2 := "350"

  expectedFinalNumStr := "350"

  expectedNRunes := []rune("350")

  expectedIntArray := []int{3, 5, 0}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_19(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_19"

  inputNumStr1 := "0"

  inputNumStr2 := "-350"

  expectedFinalNumStr := "-350"

  expectedNRunes := []rune("350")

  expectedIntArray := []int{3, 5, 0}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_20(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_20"

  inputNumStr1 := "122"

  inputNumStr2 := "122"

  expectedFinalNumStr := "244"

  expectedNRunes := []rune("244")

  expectedIntArray := []int{2, 4, 4}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_21(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_21"

  inputNumStr1 := "-122"

  inputNumStr2 := "122"

  expectedFinalNumStr := "0"

  expectedNRunes := []rune("0")

  expectedIntArray := []int{0}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_22(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_22"

  inputNumStr1 := "-122"

  inputNumStr2 := "-122"

  expectedFinalNumStr := "-244"

  expectedNRunes := []rune("244")

  expectedIntArray := []int{2, 4, 4}

  expectedPrecisionInt := 0

  expectedSignValue := -1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_23(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_23"

  inputNumStr1 := "122"

  inputNumStr2 := "-122"

  expectedFinalNumStr := "0"

  expectedNRunes := []rune("0")

  expectedIntArray := []int{0}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_AddN1N2_24(t *testing.T) {

  ePrefix := "TestStrMathOp_AddN1N2_24"

  inputNumStr1 := "0"

  inputNumStr2 := "0"

  expectedFinalNumStr := "0"

  expectedNRunes := []rune("0")

  expectedIntArray := []int{0}

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedLengthNRunes := len(expectedNRunes)

  expectedLengthIntArray := len(expectedIntArray)

  mOps := new(StrMathOp).New()

  err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := mOps.N1.SetIntAryWithNumStr(inputNumStr1)\n"+
      "inputNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr1, err.Error())
    return
  }

  err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.N2.SetIntAryWithNumStr(inputNumStr2)\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumStr2, err.Error())
    return
  }

  err = mOps.AddN1N2()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = mOps.AddN1N2()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalNumStr, err := mOps.IFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalNumStr, err := mOps.IFinal.GetNumStr()\n"+
      "inputNumStr1= '%v'\n"+
      "inputNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr1,
      inputNumStr2,
      err.Error())

    return
  }

  actualFinalPrecisionInt := mOps.IFinal.GetPrecision()

  actualFinalSignValue, err := mOps.IFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualFinalSignValue, err := mOps.IFinal.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  actualLengthNRunes := mOps.IFinal.GetIntAryLength()

  actualLengthIntArray := mOps.IFinal.GetIntAryLength()

  if expectedFinalNumStr != actualFinalNumStr {
    t.Errorf("%v\n"+
      "Error: actualFinalNumStr Result is INVALID!\n"+
      "Because expectedFinalNumStr != actualFinalNumStr\n"+
      "Expected actualFinalNumStr = '%v'\n"+
      "  Actual actualFinalNumStr = '%v'\n\n",
      ePrefix, actualFinalNumStr, actualFinalNumStr)

    return
  }

  if expectedPrecisionInt != actualFinalPrecisionInt {
    t.Errorf("%v\n"+
      "Error: actualFinalPrecisionInt Result is INVALID!\n"+
      "Because expectedPrecisionInt != actualFinalPrecisionInt\n"+
      "Expected actualFinalPrecisionInt = '%v'\n"+
      "  Actual actualFinalPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, actualFinalPrecisionInt)

    return
  }

  if expectedSignValue != actualFinalSignValue {
    t.Errorf("%v\n"+
      "Error: actualFinalSignValue is INVALID!\n"+
      "Because expectedSignValue != actualFinalSignValue\n"+
      "Expected actualFinalSignValue = '%v'\n"+
      "  Actual actualFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, actualFinalSignValue)

    return
  }

  if expectedLengthNRunes != actualLengthNRunes {
    t.Errorf("%v\n"+
      "Error: actualLengthNRunes is INVALID!\n"+
      "Because expectedLengthNRunes != actualLengthNRunes\n"+
      "Expected actualLengthNRunes = '%v'\n"+
      "  Actual actualLengthNRunes = '%v'\n\n",
      ePrefix, expectedLengthNRunes, actualLengthNRunes)

    return
  }

  if expectedLengthIntArray != actualLengthIntArray {
    t.Errorf("%v\n"+
      "Error: actualLengthIntArray Result is INVALID!\n"+
      "Because expectedLengthIntArray != actualLengthIntArray\n"+
      "Expected actualLengthIntArray = '%v'\n"+
      "  Actual actualLengthIntArray = '%v'\n\n",
      ePrefix, expectedLengthIntArray, actualLengthIntArray)

    return
  }

  var actualRune rune

  for i := 0; i < expectedLengthNRunes; i++ {

    actualRune, err = mOps.IFinal.GetIntAryRune(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualRune, err := mOps.IFinal.GetIntAryRune(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    if expectedNRunes[i] != actualRune {
      t.Errorf("%v\n"+
        "Error: expected VS actual rune DO NOT MATCH!\n"+
        "Because expectedNRunes[i] != actualRune\n"+
        "Expected actualRune = '%v'\n"+
        "  Actual actualRune = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedNRunes[i], actualRune, i)

      return
    }

  }

  var elementUint8 uint8

  var actualArrayElement int

  for i := 0; i < expectedLengthIntArray; i++ {

    elementUint8, err = mOps.IFinal.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "elementUint8, err = mOps.IFinal.GetIntAryElement(i)\n"+
        "i= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    actualArrayElement = int(elementUint8)

    if expectedIntArray[i] != actualArrayElement {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }

    if expectedIntArray[i] != actualArrayElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected actualArrayElement = '%v'\n"+
        "  Actual actualArrayElement = '%v'\n"+
        "Index 'i' = '%v'\n\n",
        ePrefix, expectedIntArray[i], actualArrayElement, i)

      return
    }

  }

  return
}

func TestStrMathOp_Divide_01(t *testing.T) {

  ePrefix := "TestStrMathOp_Divide_01"

  inputDividendStr := "-9360"

  inputDivisorStr := "24.48"

  inputMaxPrecisionInt := 29

  //                                       1         2         3
  //                            0.1234567890123456789012345678901234567
  expectedQuotientNumStr := "-382.35294117647058823529411764706"

  expectedPrecisionInt := 29

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  smop := new(StrMathOp).New()

  err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)\n"+
      "inputDividendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDividendStr, err.Error())
    return
  }

  err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)\n"+
      "inputDivisorStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDivisorStr, err.Error())
    return
  }

  err = smop.Divide(inputMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divide(inputMaxPrecisionInt)\n"+
      "inputMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputMaxPrecisionInt, err.Error())
    return
  }

  err = smop.Quotient.IsValid("Validating Quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Quotient.IsValid(\"Validating Quotient\")\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientNumStr, err := smop.Quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumStr, err := smop.Quotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientPrecisionInt := smop.Quotient.GetPrecision()

  smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientSignValue, err := smop.Quotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientSignValue, err := smop.Quotient.GetSign()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientNumSeps, err := smop.Quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumSeps, err :=\n"+
      "  smop.Quotient.GetNumericSeparatorsDto()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  if expectedQuotientNumStr != smopQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Number Strings DON'T MATCH!\n"+
      "Because expectedQuotientNumStr != smopQuotientNumStr\n"+
      "Expected smopQuotientNumStr = '%v'\n"+
      "  Actual smopQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuotientNumStr, smopQuotientNumStr)

    return
  }

  if expectedPrecisionInt != smopQuotientPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Integer Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionInt != smopQuotientPrecisionInt\n"+
      "Expected smopQuotientPrecisionInt = '%v'\n"+
      "  Actual smopQuotientPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, smopQuotientPrecisionInt)

    return
  }

  if expectedPrecisionUint != smopQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Uint Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionUint != smopQuotientPrecisionUint\n"+
      "Expected smopQuotientPrecisionUint = '%v'\n"+
      "  Actual smopQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, smopQuotientPrecisionUint)

    return
  }

  if expectedSignValue != smopQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != smopQuotientSignValue\n"+
      "Expected smopQuotientSignValue = '%v'\n"+
      "  Actual smopQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, smopQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(smopQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Numeric Separator Values DON'T MATCH!\n"+
      "Because expectedNumSeps != smopQuotientNumSeps \n"+
      "Expected smopQuotientNumSeps = '%v'\n"+
      "  Actual smopQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), smopQuotientNumSeps.String())

    return
  }

  return
}

func TestStrMathOp_Divide_02(t *testing.T) {

  ePrefix := "TestStrMathOp_Divide_02"

  inputDividendStr := "48"

  inputDivisorStr := "24"

  inputMaxPrecisionInt := 0

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedQuotientNumStr := "2"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  smop := new(StrMathOp).New()

  err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)\n"+
      "inputDividendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDividendStr, err.Error())
    return
  }

  err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)\n"+
      "inputDivisorStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDivisorStr, err.Error())
    return
  }

  err = smop.Divide(inputMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divide(inputMaxPrecisionInt)\n"+
      "inputMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputMaxPrecisionInt, err.Error())
    return
  }

  err = smop.Quotient.IsValid("Validating Quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Quotient.IsValid(\"Validating Quotient\")\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientNumStr, err := smop.Quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumStr, err := smop.Quotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientPrecisionInt := smop.Quotient.GetPrecision()

  smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientSignValue, err := smop.Quotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientSignValue, err := smop.Quotient.GetSign()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientNumSeps, err := smop.Quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumSeps, err :=\n"+
      "  smop.Quotient.GetNumericSeparatorsDto()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  if expectedQuotientNumStr != smopQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Number Strings DON'T MATCH!\n"+
      "Because expectedQuotientNumStr != smopQuotientNumStr\n"+
      "Expected smopQuotientNumStr = '%v'\n"+
      "  Actual smopQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuotientNumStr, smopQuotientNumStr)

    return
  }

  if expectedPrecisionInt != smopQuotientPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Integer Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionInt != smopQuotientPrecisionInt\n"+
      "Expected smopQuotientPrecisionInt = '%v'\n"+
      "  Actual smopQuotientPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, smopQuotientPrecisionInt)

    return
  }

  if expectedPrecisionUint != smopQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Uint Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionUint != smopQuotientPrecisionUint\n"+
      "Expected smopQuotientPrecisionUint = '%v'\n"+
      "  Actual smopQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, smopQuotientPrecisionUint)

    return
  }

  if expectedSignValue != smopQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != smopQuotientSignValue\n"+
      "Expected smopQuotientSignValue = '%v'\n"+
      "  Actual smopQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, smopQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(smopQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Numeric Separator Values DON'T MATCH!\n"+
      "Because expectedNumSeps != smopQuotientNumSeps \n"+
      "Expected smopQuotientNumSeps = '%v'\n"+
      "  Actual smopQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), smopQuotientNumSeps.String())

    return
  }

  return
}

func TestStrMathOp_Divide_03(t *testing.T) {

  ePrefix := "TestStrMathOp_Divide_03"

  inputDividendStr := "54"

  inputDivisorStr := "24"

  inputMaxPrecisionInt := 7

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedQuotientNumStr := "2.25"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  smop := new(StrMathOp).New()

  err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)\n"+
      "inputDividendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDividendStr, err.Error())
    return
  }

  err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)\n"+
      "inputDivisorStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDivisorStr, err.Error())
    return
  }

  err = smop.Divide(inputMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divide(inputMaxPrecisionInt)\n"+
      "inputMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputMaxPrecisionInt, err.Error())
    return
  }

  err = smop.Quotient.IsValid("Validating Quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Quotient.IsValid(\"Validating Quotient\")\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientNumStr, err := smop.Quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumStr, err := smop.Quotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientPrecisionInt := smop.Quotient.GetPrecision()

  smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientSignValue, err := smop.Quotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientSignValue, err := smop.Quotient.GetSign()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientNumSeps, err := smop.Quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumSeps, err :=\n"+
      "  smop.Quotient.GetNumericSeparatorsDto()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  if expectedQuotientNumStr != smopQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Number Strings DON'T MATCH!\n"+
      "Because expectedQuotientNumStr != smopQuotientNumStr\n"+
      "Expected smopQuotientNumStr = '%v'\n"+
      "  Actual smopQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuotientNumStr, smopQuotientNumStr)

    return
  }

  if expectedPrecisionInt != smopQuotientPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Integer Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionInt != smopQuotientPrecisionInt\n"+
      "Expected smopQuotientPrecisionInt = '%v'\n"+
      "  Actual smopQuotientPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, smopQuotientPrecisionInt)

    return
  }

  if expectedPrecisionUint != smopQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Uint Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionUint != smopQuotientPrecisionUint\n"+
      "Expected smopQuotientPrecisionUint = '%v'\n"+
      "  Actual smopQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, smopQuotientPrecisionUint)

    return
  }

  if expectedSignValue != smopQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != smopQuotientSignValue\n"+
      "Expected smopQuotientSignValue = '%v'\n"+
      "  Actual smopQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, smopQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(smopQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Numeric Separator Values DON'T MATCH!\n"+
      "Because expectedNumSeps != smopQuotientNumSeps \n"+
      "Expected smopQuotientNumSeps = '%v'\n"+
      "  Actual smopQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), smopQuotientNumSeps.String())

    return
  }

  return
}

func TestStrMathOp_Divide_04(t *testing.T) {

  ePrefix := "TestStrMathOp_Divide_04"

  inputDividendStr := "0"

  inputDivisorStr := "24"

  inputMaxPrecisionInt := 7

  //                                       1         2         3
  //                            0.1234567890123456789012345678901234567
  expectedQuotientNumStr := "0"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  smop := new(StrMathOp).New()

  err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)\n"+
      "inputDividendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDividendStr, err.Error())
    return
  }

  err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)\n"+
      "inputDivisorStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDivisorStr, err.Error())
    return
  }

  err = smop.Divide(inputMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divide(inputMaxPrecisionInt)\n"+
      "inputMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputMaxPrecisionInt, err.Error())
    return
  }

  err = smop.Quotient.IsValid("Validating Quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Quotient.IsValid(\"Validating Quotient\")\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientNumStr, err := smop.Quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumStr, err := smop.Quotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientPrecisionInt := smop.Quotient.GetPrecision()

  smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientSignValue, err := smop.Quotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientSignValue, err := smop.Quotient.GetSign()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientNumSeps, err := smop.Quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumSeps, err :=\n"+
      "  smop.Quotient.GetNumericSeparatorsDto()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  if expectedQuotientNumStr != smopQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Number Strings DON'T MATCH!\n"+
      "Because expectedQuotientNumStr != smopQuotientNumStr\n"+
      "Expected smopQuotientNumStr = '%v'\n"+
      "  Actual smopQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuotientNumStr, smopQuotientNumStr)

    return
  }

  if expectedPrecisionInt != smopQuotientPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Integer Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionInt != smopQuotientPrecisionInt\n"+
      "Expected smopQuotientPrecisionInt = '%v'\n"+
      "  Actual smopQuotientPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, smopQuotientPrecisionInt)

    return
  }

  if expectedPrecisionUint != smopQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Uint Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionUint != smopQuotientPrecisionUint\n"+
      "Expected smopQuotientPrecisionUint = '%v'\n"+
      "  Actual smopQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, smopQuotientPrecisionUint)

    return
  }

  if expectedSignValue != smopQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != smopQuotientSignValue\n"+
      "Expected smopQuotientSignValue = '%v'\n"+
      "  Actual smopQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, smopQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(smopQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Numeric Separator Values DON'T MATCH!\n"+
      "Because expectedNumSeps != smopQuotientNumSeps \n"+
      "Expected smopQuotientNumSeps = '%v'\n"+
      "  Actual smopQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), smopQuotientNumSeps.String())

    return
  }

  return
}

func TestStrMathOp_Divide_05(t *testing.T) {

  ePrefix := "TestStrMathOp_Divide_05"

  inputDividendStr := "5"

  inputDivisorStr := "24"

  inputMaxPrecisionInt := 32

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedQuotientNumStr := "0.20833333333333333333333333333333"

  expectedPrecisionInt := 32

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  smop := new(StrMathOp).New()

  err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)\n"+
      "inputDividendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDividendStr, err.Error())
    return
  }

  err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)\n"+
      "inputDivisorStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDivisorStr, err.Error())
    return
  }

  err = smop.Divide(inputMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divide(inputMaxPrecisionInt)\n"+
      "inputMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputMaxPrecisionInt, err.Error())
    return
  }

  err = smop.Quotient.IsValid("Validating Quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Quotient.IsValid(\"Validating Quotient\")\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientNumStr, err := smop.Quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumStr, err := smop.Quotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientPrecisionInt := smop.Quotient.GetPrecision()

  smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientSignValue, err := smop.Quotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientSignValue, err := smop.Quotient.GetSign()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientNumSeps, err := smop.Quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumSeps, err :=\n"+
      "  smop.Quotient.GetNumericSeparatorsDto()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  if expectedQuotientNumStr != smopQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Number Strings DON'T MATCH!\n"+
      "Because expectedQuotientNumStr != smopQuotientNumStr\n"+
      "Expected smopQuotientNumStr = '%v'\n"+
      "  Actual smopQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuotientNumStr, smopQuotientNumStr)

    return
  }

  if expectedPrecisionInt != smopQuotientPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Integer Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionInt != smopQuotientPrecisionInt\n"+
      "Expected smopQuotientPrecisionInt = '%v'\n"+
      "  Actual smopQuotientPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, smopQuotientPrecisionInt)

    return
  }

  if expectedPrecisionUint != smopQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Uint Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionUint != smopQuotientPrecisionUint\n"+
      "Expected smopQuotientPrecisionUint = '%v'\n"+
      "  Actual smopQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, smopQuotientPrecisionUint)

    return
  }

  if expectedSignValue != smopQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != smopQuotientSignValue\n"+
      "Expected smopQuotientSignValue = '%v'\n"+
      "  Actual smopQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, smopQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(smopQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Numeric Separator Values DON'T MATCH!\n"+
      "Because expectedNumSeps != smopQuotientNumSeps \n"+
      "Expected smopQuotientNumSeps = '%v'\n"+
      "  Actual smopQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), smopQuotientNumSeps.String())

    return
  }

  return
}

func TestStrMathOp_Divide_06(t *testing.T) {

  ePrefix := "TestStrMathOp_Divide_06"

  inputDividendStr := "0.05"

  inputDivisorStr := "24"

  inputMaxPrecisionInt := 32

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedQuotientNumStr := "0.00208333333333333333333333333333"

  expectedPrecisionInt := 32

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  smop := new(StrMathOp).New()

  err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := smop.Dividend.SetIntAryWithNumStr(inputDividendStr)\n"+
      "inputDividendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDividendStr, err.Error())
    return
  }

  err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divisor.SetIntAryWithNumStr(inputDivisorStr)\n"+
      "inputDivisorStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputDivisorStr, err.Error())
    return
  }

  err = smop.Divide(inputMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Divide(inputMaxPrecisionInt)\n"+
      "inputMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputMaxPrecisionInt, err.Error())
    return
  }

  err = smop.Quotient.IsValid("Validating Quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = smop.Quotient.IsValid(\"Validating Quotient\")\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientNumStr, err := smop.Quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumStr, err := smop.Quotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  smopQuotientPrecisionInt := smop.Quotient.GetPrecision()

  smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientPrecisionUint, err := smop.Quotient.GetPrecisionUint()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientSignValue, err := smop.Quotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientSignValue, err := smop.Quotient.GetSign()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  smopQuotientNumSeps, err := smop.Quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "smopQuotientNumSeps, err :=\n"+
      "  smop.Quotient.GetNumericSeparatorsDto()\n"+
      "smopQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, smopQuotientNumStr, err.Error())
    return
  }

  if expectedQuotientNumStr != smopQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Number Strings DON'T MATCH!\n"+
      "Because expectedQuotientNumStr != smopQuotientNumStr\n"+
      "Expected smopQuotientNumStr = '%v'\n"+
      "  Actual smopQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuotientNumStr, smopQuotientNumStr)

    return
  }

  if expectedPrecisionInt != smopQuotientPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Integer Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionInt != smopQuotientPrecisionInt\n"+
      "Expected smopQuotientPrecisionInt = '%v'\n"+
      "  Actual smopQuotientPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, smopQuotientPrecisionInt)

    return
  }

  if expectedPrecisionUint != smopQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Uint Precision Values DON'T MATCH!\n"+
      "Because expectedPrecisionUint != smopQuotientPrecisionUint\n"+
      "Expected smopQuotientPrecisionUint = '%v'\n"+
      "  Actual smopQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, smopQuotientPrecisionUint)

    return
  }

  if expectedSignValue != smopQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != smopQuotientSignValue\n"+
      "Expected smopQuotientSignValue = '%v'\n"+
      "  Actual smopQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, smopQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(smopQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected VS Actual Numeric Separator Values DON'T MATCH!\n"+
      "Because expectedNumSeps != smopQuotientNumSeps \n"+
      "Expected smopQuotientNumSeps = '%v'\n"+
      "  Actual smopQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), smopQuotientNumSeps.String())

    return
  }

  return
}

func TestStrMathOp_MultiplyN1N2_01(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "457.3"
  nStr2 := "22.2"
  expected := "10152.06"
  nRunes := []rune("1015206")
  eIAry := []int{1, 0, 1, 5, 2, 0, 6}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 2
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_02(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "457.3"
  nStr2 := "-22.2"
  expected := "-10152.06"
  nRunes := []rune("1015206")
  eIAry := []int{1, 0, 1, 5, 2, 0, 6}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 2
  eSignVal := -1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_03(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "-457.3"
  nStr2 := "-22.2"
  expected := "10152.06"
  nRunes := []rune("1015206")
  eIAry := []int{1, 0, 1, 5, 2, 0, 6}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 2
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_04(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "-457.3"
  nStr2 := "0"
  expected := "0.0"
  nRunes := []rune("00")
  eIAry := []int{0, 0}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 1
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_05(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "0.0"
  nStr2 := "0.0"
  expected := "0.0"
  nRunes := []rune("00")
  eIAry := []int{0, 0}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 1
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_06(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "9999.9999"
  nStr2 := "9899.9899"
  expected := "98999898.01000101"
  nRunes := []rune("9899989801000101")
  eIAry := []int{9, 8, 9, 9, 9, 8, 9, 8, 0, 1, 0, 0, 0, 1, 0, 1}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 8
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_07(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "9899999.99991234"
  nStr2 := "7989899.98995678"
  expected := "79100009899871.7273668803886652"
  nRunes := []rune("791000098998717273668803886652")
  eIAry := []int{7, 9, 1, 0, 0, 0, 0, 9, 8, 9, 9, 8, 7, 1, 7, 2, 7, 3, 6, 6, 8, 8, 0, 3, 8, 8, 6, 6, 5, 2}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 16
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_MultiplyN1N2_08(t *testing.T) {
  mOps := StrMathOp{}.New()
  nStr1 := "79100009899871.7273668803886652"
  nStr2 := "7989899.98995678"
  expected := "632001168304566313062.047887670481022949890056"
  nRunes := []rune("632001168304566313062047887670481022949890056")
  eIAry := []int{6, 3, 2, 0, 0, 1, 1, 6, 8, 3, 0, 4, 5, 6, 6, 3, 1, 3, 0, 6, 2, 0, 4, 7, 8, 8, 7, 6, 7, 0, 4, 8, 1, 0, 2, 2, 9, 4, 9, 8, 9, 0, 0, 5, 6}
  lNRunes := len(nRunes)
  lEArray := len(eIAry)
  ePrecision := 24
  eSignVal := 1

  mOps.N1.SetIntAryWithNumStr(nStr1)

  mOps.N2.SetIntAryWithNumStr(nStr2)

  err := mOps.MultiplyN1N2()

  if err != nil {
    t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
  }

  s := mOps.IFinal.GetNumStr()

  if s != expected {
    t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
  }

  if mOps.IFinal.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, mOps.IFinal.GetPrecision())
  }

  if eSignVal != mOps.IFinal.GetSign() {
    t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, mOps.IFinal.GetSign())
  }

  if lNRunes != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, mOps.IFinal.GetIntAryLength())
  }

  if lEArray != mOps.IFinal.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, mOps.IFinal.GetIntAryLength())
  }

  for i := 0; i < lNRunes; i++ {

    if element, _ := mOps.IFinal.GetIntAryRune(i); nRunes[i] != element {
      t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
      return
    }

  }

  for i := 0; i < lEArray; i++ {
    if element, _ := mOps.IFinal.GetIntAryInt(i); eIAry[i] != element {

      t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
      return

    }
  }

}

func TestStrMathOp_RaiseToPower_01(t *testing.T) {
  nStr1 := "625"
  power := 10
  expected := "9094947017729282379150390625"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_02(t *testing.T) {
  nStr1 := "625.25"
  power := 3
  expected := "244433710.953125"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_03(t *testing.T) {
  nStr1 := "5.3"
  power := 9
  expected := "3299763.591802133"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_04(t *testing.T) {
  nStr1 := "5"
  power := 0
  expected := "1"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_06(t *testing.T) {
  nStr1 := "5745"
  power := 1
  expected := "5745"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_07(t *testing.T) {
  nStr1 := "-625.25"
  power := 3
  expected := "-244433710.953125"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_08(t *testing.T) {
  nStr1 := "2"
  power := 8
  expected := "256"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_09(t *testing.T) {
  nStr1 := "0"
  power := 8
  expected := "0"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_RaiseToPower_10(t *testing.T) {
  nStr1 := "-25.25"
  power := 4
  expected := "406485.94140625"
  sMOp := StrMathOp{}.New()

  sMOp.N1.SetIntAryWithNumStr(nStr1)
  err := sMOp.RaiseToPower(power)
  if err != nil {

    t.Errorf("Error received from sMOp.RaiseThisToPower(power). nStr1= '%v' power= '%v' Error= %v", nStr1, power, err)

  }

  s := sMOp.IFinal.GetNumStr()
  if expected != s {
    t.Errorf("Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, s)
  }

}

func TestStrMathOp_SubtractN1N2_01(t *testing.T) {
  nStr1 := "900.777"
  nStr2 := "901.000"
  eNumStr := "-0.223"
  ePrecision := 3
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_02(t *testing.T) {
  nStr1 := "350"
  nStr2 := "122"
  eNumStr := "228"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_03(t *testing.T) {
  nStr1 := "-350"
  nStr2 := "122"
  eNumStr := "-472"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_04(t *testing.T) {
  nStr1 := "-350"
  nStr2 := "-122"
  eNumStr := "-228"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_05(t *testing.T) {
  nStr1 := "350"
  nStr2 := "-122"
  eNumStr := "472"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_06(t *testing.T) {
  nStr1 := "350"
  nStr2 := "0"
  eNumStr := "350"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_07(t *testing.T) {
  nStr1 := "-350"
  nStr2 := "0"
  eNumStr := "-350"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_08(t *testing.T) {
  nStr1 := "122"
  nStr2 := "350"
  eNumStr := "-228"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_09(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "350"
  eNumStr := "-472"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_10(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "-350"
  eNumStr := "228"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_11(t *testing.T) {
  nStr1 := "122"
  nStr2 := "-350"
  eNumStr := "472"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_12(t *testing.T) {
  nStr1 := "0"
  nStr2 := "350"
  eNumStr := "-350"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_13(t *testing.T) {
  nStr1 := "0"
  nStr2 := "-350"
  eNumStr := "350"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_14(t *testing.T) {
  nStr1 := "122"
  nStr2 := "122"
  eNumStr := "0"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_15(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "122"
  eNumStr := "-244"
  ePrecision := 0
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_16(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "-122"
  eNumStr := "0"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_17(t *testing.T) {
  nStr1 := "122"
  nStr2 := "-122"
  eNumStr := "244"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_18(t *testing.T) {
  nStr1 := "0"
  nStr2 := "0"
  eNumStr := "0"
  ePrecision := 0
  eSignVal := 1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}

func TestStrMathOp_SubtractN1N2_19(t *testing.T) {
  nStr1 := "1.122"
  nStr2 := "4.5"
  eNumStr := "-3.378"
  ePrecision := 3
  eSignVal := -1

  smop := StrMathOp{}.New()
  smop.N1.SetIntAryWithNumStr(nStr1)
  smop.N2.SetIntAryWithNumStr(nStr2)
  err := smop.SubtractN1N2()

  if err != nil {
    t.Errorf("Error returned from smop.SubtractN1N2(). Error= %v", err)
  }

  if eNumStr != smop.IFinal.GetNumStr() {
    t.Errorf("Error - Expected IFinal.GetNumStr()= '%v' .  Instead, IFinal.GetNumStr()= '%v' .", eNumStr, smop.IFinal.GetNumStr())
  }

  if ePrecision != smop.IFinal.GetPrecision() {
    t.Errorf("Error - Expected IFinal.GetPrecisionInt()= '%v' .  Instead, IFinal.GetPrecisionInt()= '%v' .", ePrecision, smop.IFinal.GetPrecision())
  }

  if eSignVal != smop.IFinal.GetSign() {
    t.Errorf("Error - Expected IFinal.GetSign()= '%v' .  Instead, IFinal.GetSign()= '%v' .", eSignVal, smop.IFinal.GetSign())
  }

}
