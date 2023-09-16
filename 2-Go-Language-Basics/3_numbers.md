## Numbers

- Go has two types of numeric types. 
  - The first type is an architecture-independent type. 
  - This means that regardless of the architecture you compile for, the type will have the correct size, bytes, for your type.

>> | type | description |
>> | ----- |-----|
>> | uint8 | unsigned 8-bit integers (0 >> to 255) |
>> |uint16|unsigned 16-bit integers (0 >> to 65535)|
>> |uint32|unsigned 32-bit integers (0 to 4294967295)|
>> |uint64|unsigned 64-bit integers (0 to 18446744073709551615)|
>> |int8|signed 8-bit integers (-128 to 127)|
>> |int16|signed 16-bit integers (-32768 to 32767)|
>> |int32|signed 32-bit integers (-2147483648 to 2147483647)|
>> |int64|signed 64-bit integers (-9223372036854775808 to 9223372036854775807)|
>> |float32|IEEE-754 32-bit floating-point numbers (+- 10-45 -> +- 3.4 * 1038)|
>> |float64|IEEE-754 64-bit floating-point numbers (+- 5 * 10-324 -> 1.7 * 10308)|
>> |complex64|complex numbers with float32 real and imaginary parts|
>> |complex128|complex numbers with float64 real and imaginary parts|
>> |byte|alias for uint8|
>> |rune|alias for int32|


  - The second type is an implementation-specific type, and the byte size of that numeric type can vary based on the architecture the program is built for.

>> |type|description|
>> |----|-----------|
>> |uint|either 32 or 64 bits|
>> |int|same size as uint|
>> |uintptr|an unsigned integer large enough to store the uninterpreted bits of a pointer value|

### Integer

- it's common in Go to use the implementation types like int or uint. This typically results in the fastest processing speed for your target architecture.
  
- If you know you won’t exceed a specific size range, picking an architecture-independent type can both increase speed and decrease memory usage. 
  
- Integer types come in two flavors: int and uint. An int is a signed integer, meaning it can store positive and negative numbers. An uint, on the other hand, is an unsigned integer, meaning it can only store positive numbers.

### Float

- Floats are available in two flavors: float32 and float64. float32 is a 32-bit floating point number, and float64 is a 64-bit floating point number.

### Overflow versus Wraparound

- At compile time, if the compiler can determine a value will be too large to hold in the data type specified, it will throw an overflow error. 
- At runtime if the operation on variable results in a number greater than what the variable can hold the number is wrapped around without throwing error