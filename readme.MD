## Go JWT

A sample demo about how provide JWT basic authentication using Go.

[Source](https://www.udemy.com/course/build-jwt-authenticated-restful-apis-with-golang/)

## Running the project

Go to the root directory

`cd ./go-jwt` 

Get the dependencies

`go mod tidy`

Build the database image

`docker build -t jwt-db-img .`

Start your container

`docker run -p 5432:5432 --name jwt-db -d jwt-db-img`

Start the application

`go run ./`

Use the `requests.http` file to perform the requests using `REST Client` extension on your `VSCode`.