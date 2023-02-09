# delos-aqua

Requirements
- Im running on windows 11
- Install Docker 
- Dbeaver - PostgreSQL
- Go Version 1.19.0

with docker i have succesfully running, if you have any error u can running in the number two 
1. How To start run with docker:
- Clone this apps
- Run this command on your cmd : docker run -d -p 80:80 docker/getting-started (for checking docker running)
- go install
- Docker compose Build
- Docker compose up
- Import postman collection to ur postman apps

This below for fail running on the top
2. How To start run with go run main.go :
- Clone this apps
- Create database postgres with name delos
- go install
- go run cmd/main.go
- - it will automigrate in db.go
- or u can import/restore database in folder resource/migrations
- for login after import database email : fauzi@delos.com, pass: fauzi123
- Import postman collection to ur postman apps

