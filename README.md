# Alert-API
an API that enables users to send requests for reading and writing alert data to a data storage system.

The Project that configures a web server using the Chi router and connects to a MySQL database using the GORM library. For managing alerts data, the server provides two HTTP endpoints: one for posting (POST) and one for reading (GET). The alerts data is transmitted and received in JSON format.

GORM and database/sql libraries are used in a server structure to hold database connections.

A GORM database object is returned by the ConnectToDatabase method, which connects to a MySQL database.

It sets up a basic HTTP server with Chi routing to deal with HTTP queries. On port 8080, it is listening.

The WriteAlert function receives incoming POST requests, parses the JSON data, and puts it in the database.

The ReadAlerts function responds to incoming GET requests by retrieving alerts data from the database using service and time range criteria and returning it as JSON.

## Pre-requsities:
* Install mysql.
* For Database string use credentials while creating the database.
* username:password@tcp(localhost:port)/dbname?charset=utf8&parseTime=True&loc=Local
* Here, username: database account username, password: password of the database set while installing db, dbname: name of the db.

## How to run the application
* go mod tidy
* go run main.go
