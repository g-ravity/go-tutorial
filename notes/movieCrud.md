To install dependencies in a Go project:

```bash
go mod init <module_name>
go get <package_name>
```

Packages in go can be discovered from pkg.go.dev

After installing packages, go.mod will look like this:

```go
module movie_crud

go 1.23.5

require github.com/gorilla/mux v1.8.1 // indirect
```

go.sum file will also be created, which basically serves the same function as package-lock.json in node projects. It stores the exact project version hash, so that whenever the project is installed, the same package version and code will be installed.

Additional commands for go packages:

```bash
go get -u github.com/gorilla/mux # updates the package to the latest version
go get github.com/gorilla/mux@v1.8.0 # installs a specific version of the package
go mod tidy # removes unused packages
```
