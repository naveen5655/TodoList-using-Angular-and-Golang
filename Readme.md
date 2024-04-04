### TodoList ###
- I created todolist using golang and angular

### Golang ###

# Install go
check go version `go version`

`go mod init golang`

# install Fiber 
- fastest HTTP engine for Go.

`go get github.com/gofiber/fiber/v2`

run application: `go run .`

# install ORM
`go get -u gorm.io/gorm`

# connect to database (link: `https://gorm.io/docs/connecting_to_the_database.html`)
`go get -u gorm.io/driver/mysql`

`go get -u gorm.io/driver/postgres`


# APIS

- post ` localhost:8000/todos`

```
  {
    "title":"do something else",
    "completed":false
  }

```



### Angular ###

check version `ng version`

create component:
`ng g component todo-form`

create service:
`ng g service todo-list`
