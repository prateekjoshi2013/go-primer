
### String

- A string is a sequence of one or more characters (letters, numbers, symbols) that can be either a constant or a variable.
   
- Strings exist within either back quotes (`) or double quotes (") in Go and have different characteristics depending on which quotes you use.

#### Interpreted String Literals

- If you use the double quotes, you are creating an interpreted string literal.

- Within the quotes, any character may appear except newline and unescaped double quote.

- You will almost always use interpreted string literals because they allow for escape characters, such as \t and \n

```go
package main

import "fmt"

func main() {
    a := "Say \"hello\"\n\t\tto Go!\n\n\nHi!"
    fmt.Println(a)
}
```

Output

```sh
$ go run .

Say "hello"

            to Go!


Hi!
```

#### Raw String Literals

- If you use the back quotes, you are creating a raw string literal.
-  Within the quotes, any character may appear except a back quote.
-  Backslashes have no special meaning inside of raw string literals. For example, trying to insert newline or tab characters will not work
-  Those characters will be interpreted as part of the string.

```go
package main

import "fmt"

func main() {
    
    a := 'Say "hello"\n\t\tto Go!\n\n\nHi!'
    fmt.Println(a)

    a = '# json data for testing
{
    "id": 1,
    "name": "Janis",
    "email": "pearl@example.com"
}
'
    fmt.Println(a)
}
```
Output

```console
$ go run .

Say "hello"\n\t\tto Go!\n\n\nHi!

# json data for testing
{
    "id": 1,
    "name": "Janis",
    "email": "pearl@example.com"
}
```

### UTF-8

- Go supports UTF-8 characters out of the box, without any special setup, libraries, or packages. example, we are using English characters for the word “Hello” and Chinese characters for “World.” 
- In this example, Go is simply printing the output to the console.
- Go is able to properly iterate correctly over both the English and Chinese characters.

```go
func main() {
    a := "Hello, 世界"
     fmt.Println(a)
}
```
```console
$ go run .

Hello, 世界
```

## Runes

- A rune is an alias for int32 and is used to represent individual characters. 

- A rune can be made up of 1 to 3 int32 values. This allows for both single and multibyte characters. 

- A rune can be defined using the single quote (') character.

```go
func main() {
    a := 'A'
    fmt.Printf("%v (%T)\n", a, a)
}
```
```console
$ go run .

65 (int32)
```

## Iterating Over UTF-8 chars

- When iterating over a string, Listing 2.20, and trying to access individual characters, you encounter issues with UTF-8 characters, if you’re not careful. 

- One approach is to use a traditional for loop, where you iterate over the string’s bytes.When iterating this way, the output is unpredictable.

```go
func main() {
    a := "Hello, 世界" // 9 characters (including the space and comma)

    // incorrect way
    for i := 0; i < len(a); i++ {
        fmt.Printf("%d: %s\n", i, string(a[i]))
    }

    //correct way
    for i, c := range a {
        fmt.Printf("%d: %s\n", i, string(c))
    }

}
```

```console
$ go run .

0: H
1: e
2: l
3: l
4: o
5: ,
6:
7: ä
8: ¸
9: ⊠
10: ç
11: ⊠
12: ⊠

0: H
1: e
2: l
3: l
4: o
5: ,
6:
7: 世
10: 界

```