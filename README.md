# Reservation system backend
This is a hotel reservation system that allows users to authenticate and manage hotel bookings. Users can retrieve hotel information and book rooms using a secure JWT-based authentication system.


## Features
 
- Users:
  * Book rooms in selected hotels
- Admins:
  * View/check booking details
    
* User authentication
* JWT token-based authentication for secure API access

## Getting started
To run this project locally

## Prerequisites
- Go installed on your machine
- MongoDB set up / for install mongodb as Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```


```
git clone 
cd reservation-system
```
Rename the .env.default file to .env or create .env file and set the appropriate values for each variable

```
go mod vendor
go mod tidy
make build
make run
```
### Running the application from docker
```
make docker
```
### Running tests
```
make test
```
## Technologies used
- **Backend**: [Go](https://golang.org/) with [Fiber](https://gofiber.io) (a fast framework for Go)
- **Database**: [MongoDB](https://mongodb.com/)
Installing mongoDB client
```
go get go.mongodb.org/mongo-driver/v2/mongo
```
- **Authentication**: JSON web tokens


### Deployment to Google Cloud Run
To deploy this application to Google Cloud Run, we need to adjust the PORT variable, since it is set in Cloud Run configuration

To build the docker image in MacOS:
```
docker buildx build --platform linux/amd64 -t USERNAME/REPO:TAG ~/PATH
```