# Tinble

# Welcome to StackEdit!

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
