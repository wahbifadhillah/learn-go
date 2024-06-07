
# 01 Getting Started
## What is Go?
Its a open-source programming language developed and published by **Google**. Its relatively new programming language, its developed in 2007 and made public in 2009. Go is extremely popular because:
- **Focus on simplicty, clarity & scalability.**
 Go is inspired by Python for its simplicity, to provide a clean and understandable syntax.
- **High performance**
 Similar to C or C++ that is popular for its capability to run multi-threading tasks.
- **Batteries included**
 Comes with standard library and built-in core features. Avoiding bunch of library uses.
- **Static typing**
 Its static typing allow a type safefety to catch many errors early.
### Popular Use-cases
- Networking & APIs
- Microservices
- CLI Tools

## Installation
### Compiler
I am using Windows 11, so I used Microsoft WIndows Installer from [Go Website](https://go.dev/dl/) file named `go1.22.4.windows-amd64.msi`.
### Code Editor
I currently have Visual Studio Code as code editor it can be used to code Go.
Therefore there is a good Go extensions named `Go` by Go Team at Google.
## Hello World
Create file named `app.go`, `.go` is Go Languange extension.
[Source code](./app.go)
To run my first Go code, run this command in VSCode terminal:
```
go run app.go
```

## Initiating GO
Define a module name with:
```
go mod init example.com/example
```
This command is to tell compiler, that the code we define is belong to module `example.com/example`. Its neccesary to build a Go app that can be used across operating systems.

## Building GO
After we define the module, now we can build it so it can be executed without go compiler. To do so, run command:
```
go build
```