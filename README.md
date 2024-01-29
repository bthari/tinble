# Tinble

This is a Tinder Like Backend Service! This service, as of now, only handle the creation and the signin function of the user. For more information on how to run the program, please see below:

## Prerequisite

**The project is made with:**

- Go Version 1.17
- MongoDB 7.0.2

**Configuration Setting**:

The deployment, mongoDB config, and JWT Secret is configurable within `config.yaml`.

## How To Run


- Create Database: firstly, you need to make the mongodb database to be used, along with its collection named "user"
- Update the configuration in the `config.yaml` file according to the auth/configuration in your device.
- Get the depedencies file using this command
```  
go mod vendor  
```  
- To run the code, you can directly run this in command line
```  
go run main.go  
```  
or, build and run the binary file
```  
go build  
./tinble  
```


## Structure

The three main layer for the application are:
```
cmd
-- api
---- handler
---- router
internal
-- model
-- store
-- usecase
-- ...
pkg
-- config
-- util
main.go
```

- **cmd**: this is the main interface of the application, which in this application it will be REST API, if the app serves a gRPC then there will be another folder the same level as `api`
    - **handler**: the handler main responsibility is to validate request/response and to transformit into an internal struct/parameter used to communicate with the usecase layer
- **internal**: the internal library code, which couldn't be imported by another application
    - **usecase**: will handle the business logic of the request
    - **store**: the only layer that has a direct access with the mongodb and will handle the query logic to the database
- **pkg**: the pkg directory will contain a code that can be used by another application

### The flow of handling request:
- handler > usecase > store 

### The API
- #### For new user want to register `{url}/register`
Body:
```json
{
  "username": "username",
  "email": "validemail@gmail.com",
  "password": "password"
}
```
- #### For new user want to register `{url}/sign-in`
Body:
```json
{
  "username": "username",
  "password": "password"
}
```