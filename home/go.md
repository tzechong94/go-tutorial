## TYPES

bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
// represents a Unicode code point
float32 float64
complex64 complex128

Outside of a function (in the global/package scope), every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

### conversion of types

Some types can be converted the following way:

```go
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
```

Unless you have a good reason to, stick to the following types:

bool
string
int
uint
byte
rune
float64
complex128

### Print

fmt.Printf - Prints a formatted string to standard out.
fmt.Sprintf() - Returns the formatted string
%v - value
%d - decimal, integer
%s - string
%f - float

### Conditionals

```go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
```

== equal to
!= not equal to
< less than

> greater than
> <= less than or equal to
> = greater than or equal to

```go
if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
```

this makes things more concise instead of declaring a length variable

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _

use explicit returns instead of implicit

```go
type Wheel struct {
    Radius int
    Material string
}

type car struct {
    FrontWheel Wheel
    BackWheel Wheel
    Make string
    Model string
}

myCar := car{}
```