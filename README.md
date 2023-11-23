### SportVenue-App-Go-Backend


> This is a Venue API built with Golang

## Getting Started

> [[Technologies](#technologies-used) &middot;  &middot; [Installations](#installations) &middot; &middot; [Tests](#tests) &middot; [Author](#author)


## Technologies Used

[golang]: (https://golang.org)

- [Golang](https://golang.org).
- [Gin Framework](https://github.com/gin-gonic/gin).


## Installations

### Clone

- Clone this project to your local machine `https://github.com/cutebex/sportvenue-go.git`

### Setup
  > In the root directory, run the command
  ```shell
  $ go run server.go
  ```

  - Use `http://localhost:8080` as base url for endpoints

## How can I improve the current code base
1. In a typical backend architecture using the Router-Controller-Service-Model structure, the Router handles incoming requests and routes them to appropriate controllers, the Controller manages the request, calling upon services for business logic, the Service layer contains business logic and interacts with the Model, and the Model represents the data structure and directly manages the data, logic, and rules of the application.
   But Because the current logic is not complicated router and service were not used.

2. Implementing Swagger UI involves setting up the Swagger middleware in your application, configuring it to parse your API routes and models, and then serving the Swagger-generated UI, which provides a comprehensive and interactive documentation of all available API endpoints, their parameters, expected responses, and even allows for testing API calls in real-time.
