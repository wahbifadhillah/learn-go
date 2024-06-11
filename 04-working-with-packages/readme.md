## Working with Packages in Go

### Functions within the Same Package

To use a function defined within the same package, simply call it directly. No import statement is needed.

```
// Example: Calling function "calculateArea" within the same package
result := calculateArea(5, 10) 
```
**Code Example:** [communication.go](./04-working-with-packages/communication.go)

### Functions from Different Packages

1. **Create a Package:** Organize your Go code into packages (directories). For instance, the `fileops` package holds functions for file operations. 

2. **Import the Package:** Use the `import` statement to bring in functions from other packages:

   ```go
   import "your-module-path/fileops" 
   ```

   * Replace `your-module-path` with the actual module path for your project (specified in `go.mod`).

3. **Call the Function:**  You can now access functions from the imported package:

   ```go
   result := fileops.ReadFloat("data.txt")
   ```
**Code Example:**  
* Package: [fileops/float.go](./04-working-with-packages/fileops/float.go)
* Module Definition: [go.mod](./04-working-with-packages/go.mod)

### Third-Party Packages

Go's package ecosystem offers a vast collection of libraries. Here's how to use them:

1. **Discover:** Find packages on the official Go package repository: [https://pkg.go.dev/](https://pkg.go.dev/)

2. **Install:**  Use `go get` to download and install the package (and its dependencies):

   ```bash
   go get github.com/Pallinder/go-randomdata
   ```
   * **Note:** Packages are typically installed globally to avoid duplication across projects.

3. **Import and Use:**  Import the package as you would any other package, and then call its functions:

   ```go
   import "github.com/Pallinder/go-randomdata"

   fullName := go-randomdata.FullName(go-randomdata.Male)
   ```

#### Installing All Dependencies

To conveniently install all dependencies listed in your project's `go.mod` file:

```bash
go get
```