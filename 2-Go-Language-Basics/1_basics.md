## Go Language Basics

- Go is a 
  1. statically typed

  2. garbage collected

  3. compiled language

  4. It is capable of producing highly concurrent, thread-safe programs that will scale as needed.

  5. Go binaries are statically linked and self-contained. This means they don’t need any runtime or development libraries to run.

  6. These binaries can be compiled to run a variety of architectures and operating systems.

### Static Typing

- In a statically typed language, each statement in the program is checked at compile time, when the binary is built, rather than at runtime. 
- It also means that the data type is bound to the variable.
- In dynamically typed languages, the data type is bound to the value.

### Garbage Collection

- Managing memory and applications can be tedious, and many times Because of this, Go employs the use of a garbage collector to manage memory. 
- This means that developers do not need to explicitly allocate and deallocate memory for storing variables.

### Compilation

- Go language was fast compile times. 
- The advantage of this is that it allows you to get immediate feedback to program changes and retain the benefits of a compiled language.
- This is done by the go build command