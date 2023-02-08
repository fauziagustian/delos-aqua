# delos-aqua

Requirements
- Im running on windows
- Install Docker 
- Dbeaver - PostgreSQL
- Go Version 1.19.0

How To start run with docker:
- Clone this apps
- Run this command on your cmd : docker run -d -p 80:80 docker/getting-started (for checking docker running)
- go install
- Docker compose Build
- Docker compose up
- Import postman collection to ur postman apps

How To start run with go run main.go :
- Clone this apps
- Create database postgres with name delos
- go install
- go run cmd/main.go
- - it will automigrate in db.go
- Import postman collection to ur postman apps

