# Contributing

## Checking for outdated dependencies

Launch `go list -m -u all` for all dependencies, or for project's dependencies:

```sh
go list -m -u github.com/gin-gonic/gin
go list -m -u github.com/joho/godotenv
go list -m -u gorm.io/driver/postgres
go list -m -u gorm.io/gorm
```
