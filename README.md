# Reservation system backend

## Project outline
- users -> book room from an hotel
- admins -> going to check reservations/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> Database management -> seeding, migration

## Resources
### Mongodb driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-river/mongo
```

###gofiber
Documentation
```
https://gofiber.io
```

## Docker
### Installing mongodb as Docker container
docker run --name mongodb -d mongo:latest -p 27017:27017
