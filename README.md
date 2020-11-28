# TodoList_BE
A simple todo backend API, written in golang.

### Project structure
```
main.go
models/
    todo_model.go
    connector.go
controllers/
    todo_controllers.go
```
**main.go** contains
 - middleware, intermediate steps of the request for logging or debugging
 - db connection function
 - controller specifications and path
 - kick start of the gin framework

**todo_model.go** *[tablename_model.go]* contains
 - todo table model structure
 - insert model structure of the table
 - update model structure of the table

**connector.go** contains
 - the db connection string
 - the db connection method
 - the auto migration method of the model todo

**todo_connector.go** *[tablename_controller.go]* contains
 - the controllers of a specific model API



##TODO
 - add documentation api with gin-sawgger