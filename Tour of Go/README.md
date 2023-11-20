# Motivation
To learn go and contribute to kyverno

## Introduction <a name="Para2"></a>

### Table of Contents
1. [Motivation](#Para1)
2. [Introduction](#Para2)
3. [HelloWorld](#Para3)
4. [Packages](#Para4)
    1. [Exported Names](#Para4.1)
5. [Functions](#Para5)
6. [Variables](#Para6)
    1. [Short Variable Declaration](#Para6.1)
7. [Basic Types](#Para7)


## HelloWorld   <a name="Para3"></a>
Creating a helloworld program

```
package main    # declares the package as main
import "fmt"    # import "fmt" package

func main() {
    fmt.Println("helloworld!")
}
```

## Packages <a name="Para5"></a>
* Packages can be imported using `import` statements. For eg:
`import "fmt"`

* Example to import multiple packages -> make use of paranthesized "factored" import statement:
```
import (
    "fmt"
    "math/rand"
)
```
By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand

### Exported Names  <a name="Para4.1"></a>
If a name begin with a capital letter its exported. `For example: math.Pi` where Pi is imported from math

## Functions    <a name="Para5"></a>
* Syntax: func <func-name>(<param-name> <return type>, ...) [return type]
* Also if the successive parameters have the same type we can ommit writing it multiple times. For example: `func add(x, y int) int`
* Function can return any number of results. For example `func swap(a, b string) (string, string)`
* There's also the concept of naked returns, in which we define the return variable in the function signature itself. Eg `func add(x int, y int) (sum int)`

## Variables    <a name="Para6"></a>
* The `var` declares a list of variables: as in function argument list, the type is last.
* A `var` statement can be both at a package level and function level.
* Example: `var a int`

### Short Variable Declaration  <a name="Para6.1"></a>
* Inside of a function `:=` can be used instead of `var` with implicit type
* Outside of the function every statement must begin with a keyword( `var`, `package` `func` ...) so `:=` construct is not available

## Basic Types <a name="Para7"></a>
Broadly speaking there are 6 major classes of basic types in go:
1. Boolean (`bool`)
2. String (`string`)
3. Integer (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `byte`, `rune`)
4. float (`float32`, `float64`)
5. Complex (`complex64`, `complex128`)
6. Pointers (`uintpr`, `*int`, `*int8`, `*float64`, `*float128`, `*complex64`)