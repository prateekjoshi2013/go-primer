## Modules

- Go module is a collection of Go packages and their dependencies that can be built, versioned, and managed as a unit.
- Root of a Go module is a go.mod file, which specifies the module’s dependencies, required minimum version of Go, and other metadata.

> ```go
> module demo
> 
> go 1.19
> ```


## Toolchain

- Commands such as ***go build***, ***go test***, and ***go get*** have been written to work with modules.
- When we need to work with modules directly, you can use the ***go mod*** command to interact with them. 
- The go mod command has a number of ***sub-commands*** that can be useful for managing modules.
- The ***go help mod*** command will print a list of these commands

## Initializing a Module

- To initialize a new module, you can use the ***go mod init*** command. 
- This command requires the name of the module and will create a go.mod file in the current directory with that name.
- The name of the module follows the same naming rules as Go packages. 

## Managing Modules with VCS

- Developers use services like GitHub, GitLab, and Bitbucket to manage their code.
- If hosting modules on one of these services, or an internal VCS system, you will most likely want to use a module that can be fetched using the VCS system through the ***go get*** command.
- if you were to host your module at ***https://github.com/user/repo***, your module name would be ***github.com/user/repo***

> ```go
> module github.com/user/repo
> 
> go 1.19
> ```

## Packages

- A package is a collection of Go files that share the same import path. 
- These files are almost always in the same directory as each other, and the directory name is almost always the same as the package name
- Code within a package can reference any identifier, such as a function, variable, or type, within the package, regardless of whether that identifier is exported. 
- Any exported identifier can be used by any other package.

## Naming Packages

- The name of the package must be declared at the top of every Go source file using the package keyword.
- There are naming rules, and conventions, for naming packages in Go. 
- Package names are lowercased
- Should contain only contain letters and numbers.
- You are not allowed to start package names with a number.
- Your package name should be short and descriptive. For example, ***sql is a good package name, but structured-query-language is not***
- To keep package names short, you may have to use an abbreviation. When abbreviating package names, you should make sure that the abbreviation is something that will be easily recognized by the reader of your code.
- You should always try to avoid using the names of popular packages already in the Go standard library or names used for popular variables, such as user, conn, db, etc
- ***Don’t use generic, or "catch-all" packages***, such as util, helpers, misc, etc. These are not good names for packages. They are not descriptive, and they are an open invitation for dumping random code.

## Folders, Files, and Organization

- With a few exceptions, the name of the folder and the name of the package are the same.
- While the folder name and the package name don’t always have to match, it is generally a good idea, and idiomatic, to follow this convention.

## Multiple Packages in a Folder

- **All files within the same directory/package have to have the same package name in the file**.

## File Names

- File names are less strict than package names. The convention is to use ***lowercase letters and underscores***

## Importing Packages and Modules

- The import path for a package is the relative path from the module root to the package directory.
- The import path for a package, or module, always starts with a module name.

> ```go
> package main
> 
> // import the foo package using
> // its fully qualified name
> // starting with the module name
> import "github.com/user/repo/foo"
> 
> func main() {
>    foo.Greet()
> }
> ```

## Using Imports

- The import keyword is used along with the fully qualified package name to import

> ```go
> import "time"
> import "os/exec"
> import "github.com/user/repo/foo"
> ```

- Multiple import statements can be condensed into a single import 

> ```
> import (
>    "time"
>    "os/exec"
>    "github.com/user/repo/foo"
> )
> ```
## Resolving Import Name Clashes

- We can fix this collision by using an import alias for one

> ```go
> import (
>     "demo/bar"
>     pub "demo/foo/bar"
> )
> ```

## Dependencies

- To add dependencies to our module, we can use the ***go get*** command

> ```go
> go get example.com/pkg@v1.2.3
> ```

## Go Sum File

- When adding a third-party dependency to the module, the Go toolchain will generate and manage a go.sum file.
- The go.sum file contains a list of all the dependencies, direct and indirect, that are required by the module. 
- Additionally, the go.sum file contains a hash of the source code and go.mod file for each dependency.
- The go.sum file should be checked into version control. 
- The go.sum acts both as a security measure to prevent malicious code from being added to the module and as a source of truth for the dependencies that are required by the module.
- The go.sum file is managed by the Go toolchain and should not be modified by the module author.

## Updating Dependencies

- To update a direct dependency of our module, we can use the go get command
- When we use the go get command with the -u flag, the go.mod and go.sum files are updated to include the dependency, its version, and any other dependencies that it requires.
  
> ```go
> go get -u example.com/pkg@v1.2.3
> ```

- If we want to update all direct dependencies of our module, we can use the go get command with the -u flag, without specifying a dependency.

> ```go
> go get -u
> ```

## Semantic Versioning

- Go modules are versioned using semantic versioning.6 For example, v1.2.3 is a semantic version that tells us the major version of the release, 1, the minor version of the release, 2, and the patch version of the release
- When a breaking change is introduced in a module, that module should increase its major version by 1. For example, if the module is currently on v1.2.3, and a breaking change is introduced in the module, the module should now be tagged as v2.0.0.

## Semantic Import Versions

- Semantic import versioning states that if your major version changes, the import path to that new version should also change.


- **When we import the github.com/user/repo package**, Go modules will automatically find the **latest version of the package that is less than v2.0.0** and use that version. 


- For example, if the latest version of the package is v1.2.3, the import path should be github.com/user/repo.


- If the **github.com/user/repo module is updated to v2.0.0**, then we need to **change the import path to reflect the new version github.com/user/repo/v2**.

## Cyclical Imports

- Care must be take to avoid cyclical imports. Cyclical imports can cause a package to import itself, which is not allowed by the Go language

### bar.go

```go
package bar

import "demo/foo"

type Bar int

func Convert(f foo.Foo) Bar {
   return Bar(f)
}
```

### foo.go

```go
package foo

import "demo/bar"

type Foo = bar.Bar
```

### Console

```sh
$ go build ./...

package demo/bar
        imports demo/foo
        imports demo/bar: import cycle not allowed
```
