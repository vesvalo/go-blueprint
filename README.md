
## How to start on Windows
* Install **MSYS2** https://www.msys2.org/ and append `bin` directory of MSYS2 to the `PATH` environment variable
* Install `docker`
* All commands should be execute in **DIND** container run `make dind`
* See **Development** section below for examine commands of make

## How to start on Mac, Linux
* Install `make`, `docker`
* There are two way for development: 
    * Native (recommend) 
        * Install `go` and see **Development** section below
    * DIND container (slow file system on Mac)
        * Run `make dind` and see **Development** section below

## Development
* For help, run `make`
* Download dependencies `make vendor`
* Generate source files from resource `make generate`
* Build and run application `dev-build-up` (also it's usage for rebuild && recreate containers)

## Usage
Create new item  
```
mutation{
  ms {
    new(name: "item1") {
      status
      id
    }
  }
}
```

Search items by query  
```
query{
  ms {
    search(query: "item1", cursor: {limit: 10, offset: 0, cursor:""}, order: ASC) {
      status
      id
      cursor {
        count
        limit
        offset
        cursor
      }
    }
  }
}
```

## Migration
* new: `sql-migrate new -env="local" {name}`
* up: `sql-migrate up -env="local"`
* down: `sql-migrate down -env="local"`
* redo: `sql-migrate redo -env="local"`
* skip: `sql-migrate skip -env="local"`
* status: `sql-migrate status -env="local"`

## Requirements
* GoLang 1.12+
* PostgreSQL 11+

## ENV
* APP_WD - work directory, default: application directory 
* APP_POSTGRES_DSN - example `postgres://blueprint:insecure@localhost:5567?sslmode=disable`
