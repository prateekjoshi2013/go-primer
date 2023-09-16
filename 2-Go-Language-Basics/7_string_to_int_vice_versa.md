### Converting Strings to & from Numbers

- func ParseInt(s string, base int, bitSize int) (i int64, err error)
- ParseInt interprets a string s in the given base (0, 2 to 36) and bit
- If the base argument is 0, the true base is implied by the string's prefix following the sign  (if present): 2 for "0b", 8 for "0" or "0o", 16 for "0x", and 10 otherwise.
- Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16,int32, and int64

```go
package main

import (
    "fmt"
    "log"
    "strconv"
)

func main() {

  // parsing a string to a negative integer
  i, err := strconv.ParseInt("-42", 0, 64)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("-42: ", i)

  // parsing a octal integer where base is inferred from prefix
      i, err = strconv.ParseInt("0x2A", 0, 64)
  if err != nil {
      log.Fatal(err)
  }


  fmt.Println("0x2A: ", i)
  
  // parsing an octal integer by specifiying base
  i, err = strconv.ParseInt("2A", 16, 64)

  // parsing an unsigned integer
  u, err := strconv.ParseUint("42", 0, 64)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("42 (uint): ", u)

  // parsing a float
  f, err := strconv.ParseFloat("42.12345", 64)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("42.12345: ", f)

  // converting a string to an integer
  a, err := strconv.Atoi("42")
  fmt.Printf("%[1]v [%[1]T]\n", a)

  if err != nil {
      log.Fatal(err)
  }

}
```

```console
$ go run .

-42:  -42
0x2A:  42
42 (uint): 42
42.12345: 42.12345
42 [int]
```