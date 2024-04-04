### TodoList ###
- I created todolist using golang and angular


# Install go
check go version `go version`

`go mod init golang`

# install Fiber 
- fastest HTTP engine for Go.

`go get github.com/gofiber/fiber/v2`

run application: `go run .`

![image](https://github.com/naveen5655/TodoList-using-Angular-and-Golang/assets/89301294/eb9ef67f-f984-4605-ab4a-228841228d84)


# install ORM
`go get -u gorm.io/gorm`

- connect to database (link: `https://gorm.io/docs/connecting_to_the_database.html`)
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

## Screenshots:
# Type task
![image](https://github.com/naveen5655/TodoList-using-Angular-and-Golang/assets/89301294/2a81f5b0-24ea-4d77-86b1-d1ea1fe6bbc5)


# Click on add task
![image](https://github.com/naveen5655/TodoList-using-Angular-and-Golang/assets/89301294/872a0c76-84a8-4fb0-b818-ed420e28811a)

# Delete Task
![image](https://github.com/naveen5655/TodoList-using-Angular-and-Golang/assets/89301294/d1789721-4277-4ec2-bd36-35fd2b77808c)


click on ok, it will delete the task
![image](https://github.com/naveen5655/TodoList-using-Angular-and-Golang/assets/89301294/eb8b27c2-a681-491e-b09f-4e71dc48129f)





