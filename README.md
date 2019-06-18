---
# API Users & Microservice in Go (golang)
CRUD example of users with Go and MongoDB 

#### For Run Microservice Application
Requeriments:
- Docker
- Go
- MongoDB

```bash
$ cd proyect_directory
$ docker-compose build
$ docker-compose up
```
Open in your browser:
http://localhost:8236

For run service locally without micro service change in connection.go  "mongo:27017" to "localhost:20017" and just run the app.