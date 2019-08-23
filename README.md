
## How to start
* Install `go`, `make`, `docker`
* Examine make commands, just run `make`
* `make all` must call only once
* Actual documentation `godoc -http=:6060`
* Build application `make build`
* Enjoy!

## Start application
* Up services `make dev-docker-compose-up`
* Apply migration `sql-migrate up -env="local"`
* Start API Gateway on :8080 `./artifacts/bin gateway -c ./artifacts/configs/local.yaml -b :8080 -d`
* Start MS on :8081 `./artifacts/bin daemon -c ./artifacts/configs/local.yaml -b :8081 -d`
* Navigate to http://localhost:8080/client in your browser and paste queries below from "Usage" section

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
* APP_DB_POSTGRES_DSN - example `postgres://blueprint:insecure@localhost:5567?sslmode=disable`
