# mathopsgo
Math Operations Library using the Go Programming Language

The source code repository of the 'mathops' Library 
is located here:

    https://github.com/MikeAustin71/mathopsgo.git

## Library Installation

### Installing 'mathops' Library Locally
Use this command to down load and install the 'mathops' library
locally. 

    go get github.com/MikeAustin71/mathopsgo/mathops
    
After installation, you may import and reference the library
as follows:

        import (
            "MikeAustin71/mathopsgo/mathops"
        )    

### Updating the 'mathops' Local Library Installation
To update the library run:
    
    go get -u github.com/MikeAustin71/mathopsgo/mathops


### Installing 'mathops' Library In Your Project
As an alternative you could clone the library to a local drive:

    git clone  https://github.com/MikeAustin71/mathopsgo.git

Thereafter just copy the 'MikeAustin71/mathopsgo/mathops'
directory to your local drive and reference it using the '../mathops'
syntax. Example:

    import (
        "../mathops"
    )


### Math Operations Utilities Written In The Go Programming Language

The Math Operations Utilities currently consist of separate libraries maintained
in the sub-directory, 'mathops'. This directory also contains tests used to 
validate these libraries. For additional documentation, see the source code
files identified below.


## Math Operations Types

### 1. IntAry
Type IntAry is designed	to perform a variety of math operations on number strings.
Support for both integer and fractional math operations is provided as well as 
'backup and restore' functionality.  

This Type is capable of performing highly accurate operations on very large numbers.
For example, the directory 'MikeAustin71/mathopsgo/examples/eulersnumbercalc'
contains an example which calculates Euler's Number out to 1,000 digits.

Location: MikeAustin71/mathopsgo/mathops/intary.go


### 2. Decimal
Type Decimal contains Date Fields and Methods for managing decimal numbers.  

The type is used to perform math operations which achieve a high degree of
accuracy and uniformity when dealing with fractional numbers containing
digits to the right of the decimal place.

Location: MikeAustin71/mathopsgo/mathops/decimal.go

### 3. NthRootOp
Type NthRootOp is used to extract square roots and nth roots of positive
and negative real numbers. Currently nth roots may only be passed as
integer values. The technique employed to calculate nth roots is known as
the "shifting nth-root algorithm".

Location: MikeAustin71/mathopsgo/mathops/nthroot.go

### 4. NumStrUtility
Type NumStrUtility provides a set of numeric conversion
and management routines primarily focused on number strings.

Location: MikeAustin71/mathopsgo/mathops/numstrutility.go


### 5. NumStrDto 
A 'lite' data transfer object containing data fields and methods
used to manage, store and transport number strings. 

Location: MikeAustin71/mathopsgo/mathops/numstrdto.go

### 6. StrMathOp 
This Type performs various math operations on a series of
IntAry's and stores the results internally.

Location: MikeAustin71/mathopsgo/mathops/strmathop.go

### 7. Math Constants
This source file is used to store constants used by various
Types in the 'mathops' Library.

Location: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
