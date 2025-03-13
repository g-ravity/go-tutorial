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

require github.com/gorilla/mux v1.8.1
```

go.sum file will also be created, which basically serves the same function as package-lock.json in node projects. It stores the exact project version hash, so that whenever the project is installed, the same package version and code will be installed.

Additional commands for go packages:

```bash
go get -u github.com/gorilla/mux # updates the package to the latest version
go get github.com/gorilla/mux@v1.8.0 # installs a specific version of the package
go mod tidy # removes unused packages
```

Go doesn't store dependencies locally inside a project. It stores all the dependencies in a central location, so that the dependencies are shared among projects. So there is no node_modules like folder in Go.

On MacOS, you can check the go dependencies here:

```bash
cd ~/go/pkg/mod
```

For hot reloading, we can use air package:

```bash
go install github.com/air-verse/air@latest
air init # inside your project root directory
```

If for some reason, air command is not recognized, you can use this command to add air to your PATH variable:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Now `air init` will create .air.toml file with basic config

In the .air.toml file, include_ext means which extensions air would look out for hot reloading. You can add extensions as you want. Basic would be "go"

exclude_dir is for directories which you don't want to be watched for hot reloading. Normally, these would be your frontend folders, scripts, api docs etc. Evaluate folders as you add them.

Same thing goes for exclude_file and include_dir.

Other important configs to watch out for are tmp_dir, bin and cmd.
tmp_dir is the temporary directory where air stores the compiled files for hot reloading
bin specifies the binary file name after compilation.

If you use a different folder name other than default tmp (for example, bin), then update it in the .air.toml files

cmd specifies the command to build your go project. If your go files are not in root folder (most commonly, main.go would be stored in something like cmd/api/main.go), then you can specify the path to it here. Otherwise, you can leave it as it is.

With gin, you can simply setup server like this:

```go
r := gin.Default()

r.run()
```

It takes the PORT value from environment variable.
Use godotenv package to set and use environment variables. If no PORT env value is found, then gin uses 8080 by default

Example usage of godotenv

```go
go get github.com/joho/godotenv

if err := godotenv.Load(); err != nil {
	log.Fatal("Error loading .env file")
}

host := os.Getenv("DB_HOST")
```

In Go, init() is a special function, which is automatically called by the Go runtime. It is used to initialize the package (for example, load env variables, do other initializations etc). It is called before the main function is called.

You can connect to any managed DB solution like Firebase or Supabase for database access. I spinned a local Postgres instance using Docker.

```bash
docker pull postgres
docker run --name <service-name> -e POSTGRES_PASSWORD=<password> -p 5432:5432 -d <userName>
```

Now, to connect to the Postgres instance, you can use GORM

```bash
go get gorm.io/gorm
got get gorm.io/driver/<driver> # postgres / mysql / sqlite accordingly
```

Copy and paste the following syntax to get your DB connection

```go
DB, err = gorm.Open(postgres.Open(<connectionString>), &gorm.Config{}) // get this from gorm docs, and replace your creds
```

Sprintf is used in Go for string formatting

```go
str := fmt.Sprintf("Sample test string: %s", "test")
```

%s : string
%d : integer
%f : floats
%t : boolean
%v : any value (go compiler will determine format)
