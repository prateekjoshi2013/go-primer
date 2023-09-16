## Printing and Formatting

### Sprint Functions

- Functions starting with Sprint, such as fmt.Sprint, all return a formatted string

### Print Functions

- Functions starting with Print, such as fmt.Print, all print the formatted string to the standard output

### Sprint Functions

- Functions starting with Sprint, such as fmt.Sprint, all return a formatted string

### Fprint Functions

- Functions starting with Fprint, such as fmt.Fprint, all print the formatted string to the provided io.Writer (for ex. writing to a file)

### Multiple Arguments with Println

- When calling fmt.Println, multiple arguments are printed in the order they are passed, with a new line added to the end of each line.

- fmt.Println Automatically Adds a New Line to the End of Each Argument to Be Printed

## Formatting Functions

- All of the functions ending in f, such as fmt.Sprintf, allow for the use of formatting verbs. 

- These verbs will create different styles of formatting. Verbs are usually proceeded by a % or \ character.

### Escape Sequences

- Escape sequences, such as \n to create a new line and \t to create a tab, can be used to format the output

### Escaping Escape Sequences

- Formatting verbs begin with a % character. Escape sequences begin with a \ character.
   
- There may be times that you want to use those characters as themselves in the formatted string. To do this, each character needs to be escaped by repeating itself so \\\ and %%

### Formatting Strings

- When formatting a string, the two most common verbs are %s and %q. 
- The %s verb prints the string as it is
- the %q verb prints the string as it is, with quotes around it

```go
  s := "Hello, World!"

  // use '%s' to print a string
  fmt.Printf("%s\n", s)

  // use '%q' to print a string
  fmt.Printf("%q\n", s)
```

```console
Hello, World!
"Hello, World!"
```

### Formatting Integers

- Integers are printed with the %d verb

- The %d verb can be modified to add padding to the right or left of the integer

- This padding is considered a minimum padding. 

- If a number is larger than the minimum padding, the number is printed as it is.

- Width of the padding is declared by using an integer in front of the d in the %d verb


```go
package main

import "fmt"

func main() {
    d := 123

    // use '%5d' to print an integer
    // padded on the left with spaces
    // to a minimum of 5 characters wide
    fmt.Printf("Padded: '%5d'\n", d)

    // use '%5d' to print an integer
    // padded on the left with zeros
    // to a minimum of 5 characters wide
    fmt.Printf("Padded: '%05d'\n", d)

    // a number larger than the padding
    // is printed as is
    fmt.Printf("Padded: '%5d'\n", 1234567890)

    // use '%-5d' to print an integer
    // padded on the right with spaces

    // to a minimum of 5 characters wide
    fmt.Printf("Padded: '%-5d'\n", d)

    // a number larger than the padding
    // is printed as is
    fmt.Printf("Padded: '%-5d'\n", 1234567890)
}
```
```console
Padded: '  123'
Padded: '00123'
Padded: '1234567890'
Padded: '123 '
Padded: '1234567890'
```

### Formatting Floats

- Formatting floats is similar to formatting integers with the %f verb being used
- formatting a float’s decimal point to a specific number of places. 
- For example, the integer 2 is used to specify the number of decimal places: %.2f

```go
package main

import (
    "fmt"
)

func main() {

    // use '%f' to print a float
    fmt.Printf("%f\n", 1.2345)

    // use '%.2f' to print a float
    // with 2 decimal places
    fmt.Printf("%.2f\n", 1.2345)
}

```
```console
$ go run .

1.234500
1.23
```

### Printing a Value’s Type

- The %T verb prints the type of the value. 
  
- For example, the type of the variable u is printed using %T.

```go
u := User{
        Name: "Kurt",
        Age:  27,
    }

// use '%T' to print the type of a value
fmt.Printf("%T\n", u)

```

```console
$ go run .

main.User
```

### Printing a Value

- To print the value of a variable, you use the %v verb. 

- For example, the value of the variable u is printed using the %v formatting verb. 

- The %v formatting verb prints the values of each of the struct’s fields.

- The %v verb can be extended with the + operator to print more information about the value, if possible.

- For example, the %+v verb is used to print a struct’s field names, as well as the values of the fields.

- When you use a formatting verb incorrectly, such as trying to format an int with a %s verb, there will be no error to respond to. Instead, Go prints error information into the formatted string.


```go
    u := User{
      Name: "Kurt",
        Age:  27,
    }

    // use '%v' to print a value
    fmt.Printf("%v\n", u)
    
    // use '%+v' to print an extended
    // representation of a value, if possible
    fmt.Printf("%+v\n", u)

    // use '%#v' to print the
    // Go-syntax representation of a value
    fmt.Printf("%#v\n", u)

    // print an integer with a the '%s' verb
    fmt.Printf("This is an int: %s\n", 42)
```

```console
$ go run .

{Kurt 27}
{Name:Kurt Age:27}
main.User{Name:"Kurt", Age:27}
This is an int: %!s(int=42)
```

- To avoid using the wrong verbs, and other issues, you should always run **go vet** on your code

```console
$ go vet .

# demo
./main.go:8:2: fmt.Printf format %s has arg 42 of wrong type int
```

### Explicit Argument Indexes

- The default behavior for these functions is for each formatting verb to format successive arguments passed into the call
  
- We can use [N] with the verb to indicate which argument to use.
  
- Using the [N] syntax, we can use the same argument multiple times in the formatting string.

```go
func main() {
    fmt.Printf("in order: %s %s %s %s\n", "one", "two", "three", "four")
    fmt.Printf("explicit: %[4]s %[3]s %[2]s %[1]s\n", "one", "two", "three", "four")
    //use [1] to reuse the first argument in multiple places
    fmt.Printf("Value of name is %[1]q, type: %[1]T\n", "John")
}
```
```console
$ go run .

in order: one two three four
explicit: four three two one
Value of name is "John", type: string
```