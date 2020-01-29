<img src="https://github.com/ashishghogre/go-ang-todo/workflows/Back-end%20Test/badge.svg"/> <img src="https://github.com/ashishghogre/go-ang-todo/workflows/Client%20Test/badge.svg"/>

# TODO APP

A simple app to manage list of todo items built in Angular8 and GoLang

## Getting Started

The Repository contains the GoLang code for the back-end REST APIs and Angular, HTML, CSS code for the front-end

## Prerequisites

The Project uses GoLang for back-end, in order to install GoLang follow the instructions from the following Link

```
    https://golang.org/doc/install
```

For the front-end following tools are needed

- [NVM](https://github.com/creationix/nvm)

The project also uses dockerized builds for both APIs and Front-end.
In Case you wish to use dockerized builds for development make sure to have docker desktop installed and running.

## Installing and Running Back-end API

- With Go

  - change directory to src folder in api and run the following commands to build and run the app

  ```
      go build .
      go install .
      go run github.com/ashishghoge/goapi
  ```

- With Docker
  - Run the following commands from within api directory
  ```
      docker build -t <tagname> .
      docker run -p8080:8080 <tagname>
  ```
  - The first command builds a docker image with the name provided in place of tagname.
  - The second command runs the image created in first step and maps the port 8080 of host to the port 8080 inside the docker container.
  - In case port 8080 is currently in use in your host OS, change the second command to
  ```
      docker run -p<hostPort>:8080 <tagname>
  ```
  - By default the back-end api runs with an in-memory database, which gets viped off every time you restart the docker container. To get a persistent database use the following command while starting the docker container
  ```
      docker run -p8080:8080 -v$(pwd):/go/src/app <tagname>
  ```

## Installing and Running front-end

- Using Node

  - Run the following commands from within the web directory

  ```
      nvm install
      nvm use
      npm install
      npm start
  ```

  - This will start the front-end server on port 4200

- Using Docker
  - Run the following commands from within the web directory
  ```
      docker build . <tagname>
      docker run -p4200:4200 <tagname>
  ```
  - The first command builds a docker image with the name provided in place of tagname.
  - The second command runs the image created in first step and maps the port 4200 of host to the port 4200 inside the docker container.
  - In case port 4200 is currently in use in your host OS, change the second command to
  ```
      docker run -p<hostPort>:4200 <tagname>
  ```
- If you have used a port other than 8080 for running back-end, update the constant baseUrl inside the file item.service.ts to reflect correct port. E.g., you are using port 8081 change the constant from `"http://localhost:8080/"` to `"http://localhost:8081/"`
