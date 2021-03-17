# golibcontrol
golibcontrol is a Library Control API where you can control which books are available for lending or which student has lent a book.

# Installation
The first installation of this application you must run the command `make install`, but if you have already installed it, you can start this application with `make run` (with terminal attached) or `make rund` (with terminal dettached)

### Commands
`make install` install & run the application with terminal attached. Must be used the first time you run this application.

`make run` run the application with terminal attached.

`make rund` run the application with terminal dettached.

`make stop` stops the application.

`make tests` run all tests implemented.

# Architecture
This project is being constructed following the golang standards project layout. You can see more about the standard layout clicking on the reference bellow.

Reference: `https://github.com/golang-standards/project-layout`

The project follows the Hexagonal Architecture which allow us to isolate the dependencies, focus on a simplified code, focus on tests development and has an organized structure that includes better names for functions, packages and folders. You can also see more about the hexagonal architecture on the link bellow.

Reference: `https://idevoid.medium.com/stygis-golang-hexagonal-architecture-a2d89d01f84b` 


### Flow
 
`User Request > Middleware (work in progress) > Handler > Module > Repository`

### Project Structure
```
./golibcontrol/
├── build
│   └── docker    
│   │   └── json-server.go
│   │   │   └── db.json
│   │   └── Dockerfile
│   └── server    
│       └── server.go
├── cmd
│   └── main.go
├── internal
│   └── app
│       └── constants
│       │   └── model
│       │   │   ├── book_category_model.go
│       │   │   ├── book_model.go
│       │   │   ├── category_model.go
│       │   │   ├── lending_model.go
│       │   │   └── student_model.go
│       │   └── errors
│       │       └── api_errors.go
│       ├── handler
│       │   └── rest
│       │       ├── book_handler.go
│       │       ├── category_handler.go
│       │       ├── lending_handler.go
│       │       └── student_handler.go
│       ├── middleware
│       │   └── 
│       ├── module
│       │   ├── book_module.go
│       │   ├── category_module.go
│       │   ├── lending_module.go
│       │   └── student_module.go
│       └── repository
│           ├── book_repository.go
│           ├── category_repository.go
│           ├── lending_repository.go
│           └── student_repository.go
├── platform
│   └── mysql
│   │   └── mysql.go    
│   ├── redis
│   │
│   └── router
│       └── build
│       │   ├── book_routes.go               
│       │   └── category_routes.go
│       └──  router.go
├── scripts
│   ├── install.sh
│   ├── run.sh
│   ├── rund.sh
│   ├── stop.sh
│   └── tests.sh
├── vendor
│   └── dependencies
├── docker-compose.yml
├── go.mod
│   └── go.sum
└── README.md
```

### Directories

| Dir |Content|
| --- | --- |
| cmd | The cmd contains the main file of this application. |
| build | Build contains the code that will build this system  |
| handler | The handlers are responsible of redirecting the user requests. |
| module | The modules will process user requests manipulating the data and sending them to the repository. |
| repository | The repositories are the files which controll every operation with the Database. |
| constants | The constants folder contains all constants used by this application, that includes the modules which will be used on controllers, modules and repositories.  |
| pkg | The PKG folder contains all lib codes used on the application. |
| platform | The Platform folder contains all tools used by this application. |

# Dependencies
### GORM
GORM make database commands easier to understand and work. This lib is used to transform the models structs into tables in our database.

### GIN
GIN simplifies the HTTP functions making easy to get the context, create routes (with and without groups) and sending response. 

### REDIS
(redis is coming!)

### MYSQL
MYSQL is the database choosen for this application.

# Tests
The tests can be executed using the command `make tests`

# API
