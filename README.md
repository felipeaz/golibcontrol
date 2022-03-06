# FutureLib
FutureLib is a Library Platform developed to educational institutions. It provides an interface which the students can 
interact with each other inside the platform, giving their reviews about a book, comments about it and even
participate in Book Groups.
It also provides a management interface which the institution employees can manage the content available in the platform, 
and by content it means Students, Books, Lendings, Meetings and Book Groups.

This application is compound by four services:

`Account Service` Service that contains all accounts and which is responsible for user authentication.

`Management Service` Service that adds content to the digital library. The content can be Students, Books and Books Lending. 

`Platform Service` Service that contains all platform functionalities - Comments, Replies, Reviews and Reservations.

To integrate the services, the application uses Kong and Konga as API Gateway.

# Installation
The first installation of this application you must run the command `make install`, but if you have already installed it, you can start this application with `make run` (with terminal attached) or `make rund` (with terminal dettached)

### Commands
`make install` install & run the application with terminal attached. Must be used the first time you run this application.

`make run` run the application with terminal attached.

`make stop` stops the application.

`make tests` run all tests implemented.


### Running isolated servers

`make account` initialize account service.

`make management` initialize management service.

`make platform` initialize platform service.

### Setup API Gateway

Make sure that Kong is up and running then run

`make gateway` to setup all services and plugins.

# Architecture
This project is being constructed following the golang standards project layout. You can see more about the standard layout clicking on the reference bellow.

Reference: `https://github.com/golang-standards/project-layout`

The project follows the Hexagonal Architecture which allow us to isolate the dependencies, focus on a simplified code, focus on tests development and has an organized structure that includes better names for functions, packages and folders. You can also see more about the hexagonal architecture on the link bellow.

Reference: `https://idevoid.medium.com/stygis-golang-hexagonal-architecture-a2d89d01f84b`

### Flow
 
`User Request > Middleware (JWT) > Handler > Module > Repository`

### Project Structure
```
./golibcontrol/
├── build
│   └── server    
│       └── domain_item
│           └── router
│           │   └── build
│           │   │   └── domain_routes
│           │   └── router.go
│           └── router
│               └── server
│                   └── server.go
├── docker
│   └── images    
│   │   └── images_folder
│   │   │   └── Dockerfile
│   └── compose.yaml
├── cmd
│   └── domain_item
│       └── main.go
├── internal
│   └── app
│   │   └── constants
│   │    │   ├── errors
│   │    │   │   └── api_error.go
│   │    │   ├── login
│   │    │   │   └── login.go
│   │    │   └── logs
│   │    │       └── logs.go
│   │    ├── database
│   │    │   └── database_interface
│   │    ├── domain
│   │    │   └── domain_service
│   │    │       └── domain_item
│   │    │           └── handler
│   │    │           │   └── domain_item_handler.go
│   │    │           └── model
│   │    │           │   └── domain_item_model.go
│   │    │           └── module
│   │    │           │   └── domain_item_module.go
│   │    │           └── repository
│   │    │               └── interface
│   │    │               │    └── domain_item_repository_interface.go
│   │    │               └── mock
│   │    │               │    └── domain_item_repository_mock.go
│   │    │               └── domain_item_repository.go
│   │    └── middleware
│   │        └── middleware.go
│   └── pkg
│       ├── category_extractor.go
│       ├── category_extractor_test.go
│       ├── input_associator.go
│       └── input_associator_test.go
├── platform
│   └── jwt
│   │   └── model.go    
│   │   │   ├── jwt_token_details.go 
│   │   │   └── jwt_token_details.go 
│   │   └── jwt.go    
│   ├── logger
│   │   └── logger.go
│   ├── mysql
│   │   └── service    
│   │   │   └── mysql_service.go 
│   │   └── mysql.go    
│   ├── redis
│   │   └── redis.go
│   └── router
│       └── build
│       │   ├── account_routes.go               
│       │   ├── book_routes.go               
│       │   ├── category_routes.go               
│       │   ├── lending_routes.go               
│       │   └── student_routes.go
│       └──  router.go
├── scripts
│   ├── install.sh
│   ├── run.sh
│   ├── rund.sh
│   ├── stop.sh
│   └── tests.sh
├── vendor
│   └── dependencies
├── .env_example
├── docker-compose.yml
├── go.mod
│   └── go.sum
├── makefile
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
Redis is a cache tecnology which I choose to keep the user authentication stored. Doing this, I'm able to invalidate the JWT token even if it's time has not finished yet.

### MYSQL
MYSQL is the application database.

### KONG
KONG is an API Gateway which will be responsible to centralize the application services into a single host.

### KONGA
KONGA used on Api Gateway Configuration.

### POSTGRESQL
POSTGRESQL is the Api Gateway Database

# Tests
The tests can be executed using the command `make tests`

# GRPC
gRPC provides a communication server using http2 which is used to exchange messages between the FutureLib services.

To generate grpc messages and services, make sure you're in the root folder and run the command bellow

`protoc --proto_path=internal/app/plugins/protofiles --go_out=internal/app/plugins/grpc --go_opt=paths=source_relative --go-grpc_out=internal/app/plugins/grpc --go-grpc_opt=paths=source_relative ${protofilename.pro}.proto`

# API
The API Documentation was generated by Postman and it's available at: 
``https://documenter.getpostman.com/view/7958753/UVkvJXxt``

