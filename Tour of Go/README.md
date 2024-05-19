# Motivation 
Documenting my journey learning golang to keep myself accountable

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
    1. [Zero Values](#Para7.1)
    2. [Type Conversion](#Para7.2)
    3. [Constants](#Para7.3)
    4. [Numeric Constants](#Para7.4)
8. [Flow Control](#Para8)


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
* A `var` statement can be used both at a package level and function level.
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

**_NOTE_** 
* The uintptr type is typically used in conjunction with unsafe pointers, which are pointers that bypass the type safety checks provided by Go's garbage collector. unsafe pointers are used for low-level memory manipulation tasks, such as accessing hardware registers or manipulating raw memory.

### Zero Values <a name="Para7.1"></a>
Variables declared without an explicit initial zero value are given zero value
```
For Example:
Numeric types are given 0
Boolean types are given false
Strings are given ""
```

**_NOTE_**
* Make use of `"%q"` in `fmt.printf()` method to achieve the same

### Type Conversions <a name="Para7.2"></a>
Type conversion in Go are explicit and as such need to be specified
```
For Example:
i int = 23
f float = float(i)
u uint = uint(f)
```

Example use case in `Basic Types/type_convo.go`

### Constants   <a name="Para7.3"></a>
* Constants can be declared using `const` keyword
* Constants can be character, boolean, string or numeric values
* Constants cannot be declared using `:=`

```
For Example:
const truth = true
const world = "hello"
const pi = 3.14
```

### Numeric Constants   <a name="Para7.4"></a>
Numeric constants represent exact values of arbitrary precision and do not flow.

```
For example:
const a = 1 << 100
fmt.Println(a)  // Cannot print since it overflows but we can store it inside a numeric constant
```

## Flow Control <a name="Para8"></a>
### for loops <a name="Para7.4"></a>
* For loops in go is simliar to C except that it doesnt have any paranthesis and the initiator statement must make use of short hand declaration.

```
For example:
for i := 0; i < 10; i++ {
    sum += i
}
```

* Init and Post statements are optional
```
For example:
for sum < 100 {
    sum += 1
}
```

* And as you may have noticed from the previous syntax that it resembles `while`, its for a reason `for` is Golang's `while`. And offcourse it doesnt have the paranthesis.
* A simple infinite loop can be created by omitting the condition altogether
```
For Example:
for {
    //Infinite loop
}
```

### If conditions
* Similar to C statements except the expressions need not be surrounded by paranthesis () but braces are required
* Go's if statements are like its for loops

```
For example:
if i <= 2 {
    fmt.Println("i is bigger than or equal to 2")
}
```

If with a short statement
* The variable declared by the short statement wont be accessible outside the scope of the if condition
```
func pow(int x, n, lim float64) float64 {
    if v := math.pow(x, n); v < lim {
        return v
    }
    return lim
}
```
* The variable declared by the short hand is also accessible inside the else condition
```
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Println("%v >= %v", v, lim)
    }
// although v is not accessible here
    return lim
}
```

### Switch
Its similar to switch statement in C, C++ except few difference which are as follows:
* It supports automatic breaks for each case i.e only the specific case would run and not all of them
* Also Go's switch cases need not be constants or integers
* It also supports short hand statements

```
switch os := runtime.GOOS; os {
    case "Darwin":
        fmt.Println("MAC OS")
    case "linux":
        fmt.Println("Linux OS")
    default:
        fmt.Printf("%s \n", os)
}
```

#### Switch with no condition
* Its same as `switch true {...}`
* Its utilised to write long if-else conditions 

```
For Example:
t = time.Now()
switch {
case t.Hour < 12:
    fmt.Println("Good Morning")
case t.Hour < 17:
    fmt.Println("Good Afternoon")
default:
    fmt.Println("Good Evening")
}
```

### Defer
* A defer statement defers the execution of a function until the surrounding function returns.
* A defer statement's funciton argument is evaluated immediately but the funciton is not called until the surrounding function returns.

```
For example:
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

* When defer is used they are stacked in "last in first out" manner

## Pointers
* Two Differences from C's pointers are as follows:
    * Zero value for pointers in go is `nil`
    * Go does not support pointer arithmetic  

## Structs
* A struct is a collection of fields
* Syntax for defining structures: `type` <{name of structure}> `struct` {fields...} 

```
For example:
type Vertex struct {
    X int
    Y int
}
```

* Struct fields are accessed using `dot notation`
```
For example:
var v Vertex = Vertex{1, 2}
v.x = 2
```

### Pointers to structs
```
For Example:
a := Vertex{1, 2}
b := &a     // This is effectively (*b).X
b.x := 1e8
fmt.Println(b.X)
```

### Struct Literals
* In Golang, a struct literal is a concise way to create a new struct instance with the desired field values. It allows you to specify the values for each field within curly braces, following the struct name. Struct literals are often used to initialize struct instances directly rather than declaring variables and assigning values later.

* You can list just a subset of fields by using the Name: syntax. 
* And the order of named fields is irrelevant

```
For example:
var vertex1 = Vertex{X: 2, Y: 3}
```

## Arrays
* The type `[n]T` is an array of type `T` and `n values`
* An array's length is part of its type in go, so arrays cannot be resized. This seems limiting but dont worry go provides 
a convenient way of working with arrays.
* Go’s arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element 
(as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its 
contents. (To avoid the copy you could pass a pointer to the array, but then that’s a pointer to an array, not an array.) 
One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value. 

```
For Example:
var a [5]int
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## Slices
* An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays. 
* The type []T is a slice with elements of type T. 
* A slice is formed by specifying two indices, a low and high bound, separated by a colon':'. 
* syntax: a[low : high] -> This selects a half open range which selects the first element but excludes the last one.

```
For Example:
var a []int = primes[3:5]
```

* **Slices are like references to an array.**
* Changing the elements of a slice modifies the corresponding elements of its underlying array. 
* Other slices that share the same underlying array will see those changes. 

### Slice Literals
* A slice literal is like an array literal without the length. 
```
For Example:
var a []int = []int{1, 2, 3, 4, 5}
var s []struct{
    X int
    Y bool
} = []struct{
    X int
    Y bool
}{
    {1, true},
    {2, false},
    {3, false},
}

// s can also be re-written as follows inside a function using short hand declaration to make it more readable:
s := []struct{
    X int
    Y bool
}{
    {1, true},
    {2, false},
    {3, false},
}
```

* Slice defaults for lower and higher bounds is 0 and length of string.

### Slice length and capacity
* A slice has both a length and a capacity. 
* The length of a slice is the number of elements it contains. 
* The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. 
* The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s). 
* You can extend a slice's length by re-slicing it, provided it has sufficient capacity

```
For example:
// The following is when there is no underlying array
var s []int = []int{1, 2, 3, 4, 5}

// Decrease length to 0
s = s[:0]   // Length = 0, Capacity = 5

// Increase length
// Length = 4, Capacity = 5
s = s[:4]   // here '4' cannot be larger than the initial capacity or you would have to re-initialise the slice

// Drop 2 values
// Length = 1, Capacity = 1
s = s[3:]
```

* **Zero Value of slices is `nil` and the length and capacity of such slices is 0**

### Slices with `make` or Dynamically Created Arrays
* Slices can be created with the built-in make function; this is how you create dynamically-sized arrays. 
* The `make` function allocates a zeroed array and returns a slice that refers to that array 
* Syntax of make-> make(type, length, capacity) // `type` should be of type `slice type`

```
For example:
a := make([]int, 3, 5)
fmt.Println(a, len(a), cap(a))
```

### Slices with slices
* Slices can contain any type, including other slices. 
```
For example:
// We can print 2d slice with the help of the following code instead of traditional 2 loops:
for i = 0; i < len(s); i++ {
    fmt.Println(strings.Join(s[i], " "))
}
```

### Appeding to a slice
* syntax of append function:- append(s []T, vs...T)
    * `s` is a slice of type T
    * `vs` are values of type T that are to be appended to the existing slice
* If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point
to the newly allocated array.
* Refer this article for better understanding: https://go.dev/blog/slices-intro

### Range
* `range` form of for loop returns 2 values: 
    * index
    * value

```
For example:
for i, v in range slice {
    fmt.Printf("%d %d", i, v)
}
```

* You can skip the index or value by assigning `_` or if you only want index
```
For example:
1) for i,_ range slice
2) for _,v range slice
3) for i range slice
```

## Maps
* A map maps `keys` to `values`. 
* The zero value of a map is `nil`.
* A nil map has no keys, nor can keys be added. 
* It is initialised using `make`
* syntax: var := make(map[key type]"value type")

```
For Example:
a := make(map[string]int)
a["Hello"] = 1
fmt.Println(a["Hello"])
fmt.Println(a)
```

### Map literals
* Map Literals are like struct literals but with keys
```
For example:
var m = map[string]int{
    "Hello": 1,
    "World": 2,
    "Danny": 3,
}
```

* If the top-level type is just a type name, you can omit it from the elements of the literal. 
```
For example:
var a = map[string]Vertex {
    "Hello": Vertex{1, 2},
    "World": Vertex{2, 3},
}

// Can be replaced with the following code snippet
var b = map[string]Vertex{
    "Hello": {1, 2},
    "World": {2, 3},
}
```

### Mutating maps
```
For example:
m := make(map[string]int)

// insert/update
m["Hello"] = 2

// Retrieve an element
a := m["Hello"]

// Delete an element
delete(m, "Hello")  // delete("map variable", "key")

// 2 Value assignment
val, ok = m["Hello"]
```

## Function Values
* Function are values too. They can be passed around just like other values
* Function values may be used as function arguments or function returns.

```
For example:
func compute(fn func(x, y float64) float64) float64 {
    fn(3, 5)
}
...
hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(5, 12))
fmt.Println(compute(hypot))
fmt.Println(compute(math.Pow))
```

## Function closures
* These are a combination of 
    * anonymous function
    * Lexical scope:- variables defined within a function are only accessible from that function's scope
* A function closure is an anonymous function that referrences functions from outside its scope
* These variables are captured by the closure when it is created, and they remain accessible even after the enclosing function has returned.
* You cannot directly pass the argument to function that returns a function closer, instead to pass the argument to the closure
assign it to a variable and pass the argument through that variable.

```
For example:
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
...

pos, neg = adder(), adder()
for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-i))
}

```

## Class Characteristics 
### Methods and Interfaces
* Go does not have classes. However, you can define methods on types. 
* A method is a function with a special receiver argument. 
* The receiver appears in its own argument list between the func keyword and the method name. 

```
For Example:
type Vertex struct {
    X float64
    Y float64
}

func (v Vertex) Abs() float64 {     // value receiver
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
...
a := Vertex{3, 4}
fmt.Println(a.Abs())
```
* Remember: a method is just a function with a receiver argument. 
```
func Abs(v Vertex) float64 {    // A regular function
    ...
}

fmt.Println(Abs(a))  
```
* You can declare a method on non-struct types, too but define them as aliases using `type` to do so.

*__NOTE__* You can only declare a method with a receiver whose type is defined in the same package as the method.

### Pointer Receiver
* You can declare methods with pointer receivers.
* Methods with pointer receivers can modify the value to which the receiver points

```
For example:
func (v *Vertex) Scale(multiplier int) {
    v.X *= float64(multiplier)
    v.Y *= float64(multiplier)
}
...
a := Vertex{5, 12}
a.Scale(10)
```

* With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function. 

### Choosing a value or pointer receiver
* There are two reasons to use a pointer receiver. 
* The first is so that the method can modify the value that its receiver points to. 
* The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct. 

## Interfaces
* An interface type is defined as a set of method signatures. 
* A value of interface type can hold any value that implements those methods. 
* They implement polymorphism, decoupling and abstraction
```
For example:
type Base interface {
    getter() float64
}

type Vertex struct {
    x float64
    y float64
}

func (v *Vertex) getter() float64 {
    return v.x
}

type Myfloat float64

func (f float64) getter() float64 {
    return f
}
...
var a Base
b := Vertex{5, 12}
a = &b  // `&` since its pointer receiver

c := Myfloat(4)
a = c   // no need for `&` since its value receiver
```
* **Interfaces are implemented implicitly**
```
For example:
type I interface {
    M()
}

type T struct {
    S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
    fmt.Println(t.S)
}

var i I = T{"Hello"}
i.M()
```
* Interface value can be thought of as a tuple of:-
    * concrete value
    * concrete type
* Calling a method on an interface value executes the method of the same name on its underlying type. 

### Understanding Nil interfaces and interfaces with nil values in go
* If the concrete value inside the interface itself is nil, the method will be called with a nil receiver. 
* Note that an interface value that holds a nil concrete value is itself non-nil. 
* https://trstringer.com/go-nil-interface-and-interface-with-nil-concrete-value/
* Example implemented in "./Class Characteristics/interfaces_nil_val.go"

### Nil interface
* A nil interface value holds neither value nor concrete type. 
* Calling a method on nil interface is a run-time error because there in no type inside the interface tuple to indicate 
which concrete method to call.
```
For example:
type I interface {
    M()
}
...
var a I
fmt.Prinf("%v %T", i , i)
a.M() // ILLEGAL: calling a method on nil interface results in a runtime error
```

### Empty Interface
* Interfaces that specifies zero methods is known as empty interface.
* An empty interface may hold values of any type (Every type implements atleast zero methods)
* Empty interfaces are used by code that handles values of unknown type
```
For example:
var a interface{}
a = 43
fmt.Printf("%v %T", a, a)   // 43 int

a = "hello"
fmt.Println("%v %T", a, a)  // hello string
```

### Type Assertions
* A type assertion provides access to an interface value's underlying concrete value. 
```
a := i.(T)  // where T is a struct and i is an interface 
```
* This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
* If i does not hold a T, the statement will trigger a panic. 
* To test whether an interface value holds a specific type, a type assertion can return two values: 
    * the underlying value 
    * a boolean value that reports whether the assertion succeeded. 

```
For example:
var i interface{}
i = 32
a, ok := i.(int)
fmt.Println(a, ok)  // 32 true
```

### Type Switches
* A type switch is a construct that permits several type assertions in series. 
* A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value. 
```
For example:
func do(i interface{}) {
    switch v := i.(type) {
        case int:
            fmt.Println("Int", v)
        case string:
            fmt.Println("String", v)
        default:
            fmt.Println("Neither string nor int")
    }
}
...
// inside main
var i interface{} 
i = 32
do(i)
```
* The declaration in a type switch has the same syntax as a type assertion i.(int), but the specific type `int` is replaced
with the keyword `type`

### Stringer interface
* A Stringer is a type that can describe itself as a string. The fmt package(and many others) look for this interface to 
print values.
* Stringer: 
```
type Stringer interface {
    String() string
}
```
* For example:
```
type Vertex struct {
    x int
    y int
}

func (v Vertex) String() string {
    return fmt.Sprintf("Vertex contains (%v, %v)", v.x, v.y)
}
...
// In the main function
v := Vertex{3, 4}
fmt.Printf("%s", v)
```

### Error Interface
* Programs express error state with `error` values.
* The `error` type is a built-in interface is similar to fmt.Stringer
```
type error interface {
    Error() string
}
```
* Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
* By implementing Error() interface, it becomes a valid `error` type.
* **TLDR**: Its used to create a custom error

```
For example:
type MyError struct {
    When time.Time
    What string
}

func (e MyError) Error() string {
    return fmt.Sprintf("at %v, %v", e.When, e.What)
}

func run() error {
    return &MyError {
        time.Now,
        string("It doesn't work")
    }
}
...
// Inside main
if err := run(); err != nil {
    fmt.Println(err)
}
```

**_NOTE_**:
* The issue with calling `fmt.Sprint(e)/fmt.Sprintf("%v", e)` inside the Error() method of an error type in Go arises from circular recursion. Let's break down the problem and solution:
* When fmt.Sprint(e) is called within the Error() method, it attempts to format the error value itself.
* However, the Error() method is responsible for returning a string representation of the error.
* This creates a circular dependency: fmt.Sprint needs the error string, but the error string is generated by fmt.Sprint, leading to an infinite loop

### Readers Interface
* `Reader` interface in Golang is used to work with input streams of data. It defines a contract for objects that can provide bytes on demand, making it a versatile tool for reading from various sources like files, network connections or even in-memory buffers
* By implementing a Reader method for a particular type, it helps to provide a way to read input from the underlying data source i.e the type for which it is implemented.
* Example available at `"Class characteristics/readers.go"`
* A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way. 
* By implementing the io.Reader.Read method that struct essentially becomes of type `io.Reader`
* Examples available in `"Class characteristics/readers.go"` and `"Class characteristics/readers_exercise.go"`

**_NOTE_** **Underlying Resource, from which data is to be read is of type `io.Reader`**

### Images Interface
* The images interface in golang can be described using 3 methods:-
    * ColorModel() color.Model: This interface defines how colors are represented within an image. Different image formats might use different color models. Different Formats include color.RGBA(red, blue, green, alpha), color.Gray(Grayscale), color.YCbCr(luma, chroma components)
    * Bounds() image.Rectangle: This struct represents a rectangular region within an image defined by its top-left corner(x, y) and its width and height. It basically specifies the area occupied by iamge data.
    * At(x, y int) color.Color: method of the image interface returns a color.Color object representing the color at a specific point within the image bounds.
```
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}

```

## Generics
### Type Parameters [Generic Functions]
* Go functions can be written to work on multiple types using type parameters.
* The type parameters of a function appear between brackets, before the function's arguments. 
* Syntax: func func_name[T comparable](s []T, x T) int
    * here `T` is the generic type that basically simulates working with multiple types
```
For example:
func Index[T comparable](s []T, x T) int {
    for index, val := range s {
        if val == x {
            return index
        }
    }
    return -1
}
...
// in main
a := []int{1, 34, 456, 6562}
fmt.Println(a, 34)  // 1

b := []string{"schmeichel", "degea", "vandersar"}
fmt.Println(b, "barthez")   // -1
```

### Generic Types
* In addition to Generic functions, Go also supports generic types. A type can be parameterized with a type parameter, 
which could be useful for implementing generic data structure.
```
For example:
type List[T any] struct {
    next *List[T]
    val int
}

func (l *List[T]) Append(head *List[T]) {
    for l.next != nil {
        l = l.next
    }
    l.next = head
}

// inside main
a := List[int]{val: 2}
b := List[int]{val: 3}
a.Append(&b)
```

* Example in `Class characteristics/generic_types.go`

## Concurrency
### Go Routines
* A go routine is a lightweight thread managed by the Go runtime.
* `go f(x, y, z)`   // starts a new goroutine running `f(x, y, z)`
* Evaluation of `x`, `y` and `z` happens in the current goroutine and execution of f happens in the new goroutinge
* Goroutines run in the same address space, so access to shared memory must be synchronised.
* The `sync` package provides useful primitives.

```
For example:
func say(s string) {
    time.Sleep(100 * time.Millisecond)
}
...
// inside main package
go say("Hello")
say("World")
```

### Channels
* Channels are typed conduit through which you can send and receive values with the channel operator `<-`.
* By default there is a sender and receiver for a given channel and unless both are defined, trying to send and receive date using it will result in a `deadlock`. This allows goroutines to synchronize without explicit locks or condition variables. 
* Also untill the data is consumed in a channel, it cannot be accessed for writing again in the `main goroutine`, doing so will result in a `deadlock`.
*  __*A channel is always blocking if its full*__ i.e Initially channels can be thought of as size 0 and its size expands when it gets filled but for that needs to have both sender and receiver otherwise will be blocked.
```
channel := make(chan int)
go func() {
    channel <- 10       // assigning data
}()
fmt.Println(channel)    // consuming the channel
```

### Buffered Channels
* Channels can be buffered. 
* Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
* To pass on large amount of data use buffer channel without first consuming the data concurrently but can be done together
* Also buffered channels can be thought of like FIFO queues
* Overfilling the buffers will result in a deadlock
```
For example:
buf := make(chan int, 10)
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
}()
for i := 0; i < 10; i++ {
    fmt.Println(<-ch)
}
```

### Range and close
* A sender can close a channel to indicate that no more values will be sent.
* Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression. `val, ok := <-channel // here ok is a bool that represents if channel is closed or not`
* For a closed channel the value of `val`(in above) is zero value
```
func say(c chan int, n int) {
    for i := 0; i < n; i++ {
        c <- i
    }
    close(c)    // important for range to work without running into deadlock after accessing last value in channel
}
...
// In main
channel := make(chan int, 10)
go say(channel, 10)
for val := range channel {
    fmt.Println(val)
}
```

### Select
* The `select` statement lets a goroutine wait on multiple communication operations. 
* A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
* To understand `Class characteristics/select.go` watch [this](https://www.youtube.com/watch?v=42zwssL4YIE&list=PLve39GJ2D71wSwRQLp_h8B60pKgS85StC&index=4)

#### Default Selection
* The `default` case in a `select` is run if no other case is ready.
* Use a default case to try a send or receive without blocking
```
For example:
select {
    case i := <-c
        // use i
    default:
    // receive from c would block
}
```

### sync.Mutex
* This is used to make sure that one goroutine can access one resource at a time to avoid conflicts
* This concept is called mutual exclusion and the conventional name for the data structure that provides it is mutex.
* Go's standard library provides mutual exclusion with `sync.Mutex` and its 2 methods:
    * Lock
    * Unlock

```
For example:
type SafeCounter struct {
    mu sync.Mutex
    v map[string]int
}

func (s *SafeCounter) Inc(key string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.v[key]++
}

func (s* SafeCounter) Val(key string) int{
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.v[key]
}
...
// inside main
a := make(map[string]int)
for i := 0; i < 1000; i++ {
    a.Inc("somekey")
}

time.Sleep(time.Second) // since we are not using channels we need to sleep the main goroutine to enable the spawned goroutine to complete
a.Val("somekey")
```
