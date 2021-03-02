# golang-web-api-postgres-simple

## Description
This returns a data dump of the table `pop`
as a JSON encoded string. An ORM is used to
create and seed the database.

## Tech stack
- bash
- golang 1.13

## Docker stack
- ubuntu:latest
- mariadb

## Build notes
The service takes about 1 min before callable.

## To run
`sudo ./install.sh -u`
Available at http://localhost:8080

## To stop
`sudo ./install.sh -d`

## To see help
`sudo ./install.sh -h`

## Credits
- https://tutorialedge.net/golang/golang-mysql-tutorial/
- https://levelup.gitconnected.com/build-a-rest-api-using-go-mysql-gorm-and-mux-a02e9a2865ee
