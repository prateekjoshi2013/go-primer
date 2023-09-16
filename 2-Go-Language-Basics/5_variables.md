### Variables

- In Go, there are several ways to declare and initialize a variable. 

- New variables can be declared with the var keyword and given a name and a type.

- In Go, when you declare a variable or ask for an argument to a function, the name of the variable comes first, followed by the type of the variable: 

```go
func main() {
    // declare i as an int
    var i int     
    // declare f as a float64
    var f float64 
    // declare b as a bool
    var b bool    
    // declare s as a string
    var s string  
}
```
### Zero Values

- If a variable has been declared but no value has been assigned to it, the variable will have a zero value. 

- A zero value is a value that is the same type as the variable.

- Complex types, such as a **struct, have a zero value that is composed of the type’s individual fields’ zero values**.

- There are certain types in Go, such as **maps, interfaces, and pointers**, that have no obvious zero value. For these types, the **zero value is nil**.

```go
var s string    // defaults to ""
var r rune      // defaults to 0
var bt byte     // defaults to 0
var i int       // defaults to 0
var ui uint     // defaults to 0
var f float32   // defaults to 0
var c complex64 // defaults to 0
var b bool      // defaults to false
var arr [2]int  // defaults to [0 0]
var obj struct {
    b   bool
    arr [2]int
}                     // defaults to {false [0 0]}
var si []int          // defaults to []
var ch chan string    // defaults to nil
var mp map[int]string // defaults to nil map[] which you cannot add keys to
var fn func()         // defaults to nil
var ptr *string       // defaults to nil
var all any           // defaults to nil
```

```console
$ go run .

string              : ""
rune                : 0
byte                : 0
int                 : 0
uint                : 0
float32             : 0
complex64           : (0+0i)
bool                : false
array [2]int        : [0 0]
struct              : {false [0 0]}
slice []int         : []
channel chan string : <nil>
map map[int]string  : map[]
function func()     : <nil>
*string             : <nil>
any                 : <nil>
```

### Variable Declaration and Initialization

- When declaring a variable, we often want to initialize that variable with a value. 
- This is done by assigning a value to the variable at declaration time

```go
func main() {
    var i int = 42
    var f float64 = 3.14
    var b bool = true
    var s string = "hello world"
}
```
- When declaring and initializing a variable at the same time, we can use the $:=$ operator to shorten the declaration and initialization process.

```go
func main() {
    var i int = 42
    var f float64 = 3.14
    var b bool = true
    var s string = "hello world"
}
```

- When using the $:=$ operator, the variable is declared and initialized at the same time, with Go inferring the type of the variable from the value being assigned.

- For most types, the Go compiler has no trouble inferring the correct type for the variable. When it comes to numbers, however, **the compiler has to guess the type and by default assumes that the variable is of type int or float64**, if there is a decimal point in the value being assigned.

- casting can be used with the $:=$ operator to explicitly cast the value to the desired type

```go
func main() {
    i := uint32(42)
    f := float32(3.14)
}
```
- Go allows you to assign several values to several variables within the same line. Each of these values can be of a different data type. 

- We are declaring new variables, i, f, b, and s, and assigning them each a value.

```go
func main() {
    i, f, b, s := 42, 3.14, true, "hello world"

    fmt.Printf("var i %T = %v\n", i, i)
    fmt.Printf("var f %T = %f\n", f, f)
    fmt.Printf("var b %T = %v\n", b, b)
    fmt.Printf("var s %T = %q\n", s, s)
}
```

### Unused Variables

- The Go compiler does not allow for unused variables or imports

- Go also requires that you capture all of the results of a function call, even if you do not use the results

- We can tell the Go compiler to discard a variable by assigning it to the blank identifier _.

- The _ identifier acts as a placeholder for the variable that is not being used, allowing the compiler to run successfully.

```go
  func main() {
      _, _, _, s := Values()

      fmt.Println(s)
  }

  func Values() (int, float64, bool, string) {
      return 42, 3.14, true, "hello world"
  }
```

## Constants

- Constants are like variables, except they can’t be modified once they have been declared. 
  
- **Constants can only be a character, string, boolean, or numeric value**. 
  
- Constants are declared using the const keyword

- If you declare a constant with a type, it will be that exact type. In example, leapYear is defined as data type int32. This means it can only operate with int32 data types.

```go
package main

import "fmt"

const (
    year = 365            // untyped
    leapYear = int32(366) // typed
)

func main() {
    hours := 24
    minutes := int32(60)
    // multiplying an int and untyped
    fmt.Println(hours * year)
    // multiplying an int32 and untyped      
    fmt.Println(minutes * year)
    // multiplying both int32 types
    fmt.Println(minutes * leapYear) 
}
```

### Type Inference

- It is important to remember that the untyped const or var will be converted to the type it is combined with for any operation for that type. 

- In example we can see that the constants a and b, as well as the variable d, all of which are “untyped” are cast to the appropriate integer type needed for the operation.

```go
const (
    a = 2
    b = 2
    c = int32(2)
)

func main() {
    fmt.Printf("a = %[1]d (%[1]T)\n", a)
    fmt.Printf("b = %[1]d (%[1]T)\n", b)
    fmt.Printf("c = %[1]d (%[1]T)\n", c)

    fmt.Printf("a*b = %[1]d (%[1]T)\n", a*b)
    fmt.Printf("a*c = %[1]d (%[1]T)\n", a*c)

    d := 4
    e := int32(4)

    fmt.Printf("a*d = %[1]d (%[1]T)\n", a*d)
    fmt.Printf("a*e = %[1]d (%[1]T)\n", a*e)
}
```

### Naming Identifiers

#### Rules

- Variable names are case sensitive, meaning that userName, USERNAME, UserName, and uSERnAME are all completely different variables.

- Variable names cannot be reserved words.

- Variable names cannot begin with a number.

- Variable names cannot contain special characters.

- Variable names must be made up of only letters, numbers, and underscores (_).

- Variable names must only be one word (as in no spaces).

#### Convention

|Conventional Style|Unconventional Style|Why Unconventional|
|----|----|----|
|userName|user_name|Underscores are not conventional|
|i|index|Prefer i over index as it is shorter|
|serveHTTP|serveHttp|Acronyms should be capitalized|
|userID|UserId|Acronyms should be capitalized|

### Exporting through Capitalization

- The case of the first letter of an identifier has special meaning in Go. 

- If an identifier starts with an uppercase letter, that identifier is accessible outside the package it is declared in (or exported).

- If an identifier starts with a lowercase letter, it is only available within the package it is declared in.
