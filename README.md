# Go-React

I am building a golang backend and a react front end spa mainly as a learning experience

# Prerequisites 

At the moment it is pointing at a local Postgres server, I used this godaddy tutorial for setting up [postgres]  (https://www.godaddy.com/garage/how-to-install-postgresql-on-ubuntu-14-04/)  

table setup 
```bash
# log into postgres
psql -U postgres -W                
Password for user postgres: 

# create the table
postgres=# CREATE TABLE apps ( ID SERIAL PRIMARY KEY, Appname TEXT, Disabled boolean, GlobalDisableMessage TEXT); 
```
Optional install `pgAdmin` to view the local database [see](https://www.pgadmin.org/screenshots/)  
For the moment there is only The golang api server setup, standard [golang setup](https://golang.org/doc/install)

# Install

```bash
go get -u github.com/austincunningham/go-react
cd $GOPATH/src/github.com/austincuningham/go-react
dep ensure
go run main.go
```
Server is served on 8001 so
http://localhost:8001/ will return `hello world`

# API

- GET `/apps` Get all apps in apps table
- GET `/apps/:id` Get an app by id in apps table
- PUT `/apps/:id` Update an app in apps table
- POST `/apps` Create an new app in apps table
- DELETE `/apps/:id` Delete an app in apps table
