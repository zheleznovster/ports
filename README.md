# Ports manager

Ports manager is a containerized ports database app.

Ports manager loads data from a json file into a map in memory.

# Build and run from terminal

go build .  
go run .  

# Build and run with docker

docker build -t ports .  
docker run ports  

# Structure

```
ports
├──data
│   ├── ports.json - input file for database
│──database
│   ├── db.go - synced map with CRUD funcs
│   ├── db_test.go - tests for CRUD funcs
├── managers
│   ├── manager.go - func LoadData - creates a parser and a database, 
│   │                then uses parsed data to fill the database  
│   ├── manager_test.go - tests for func LoadData
├── parsers
│   ├── parser.go - parser to process ports input json one record at a time 
│   │               to avoid loading whole file into memory
│   ├── parser_test.go - tests for parser funcs
├── signals
│   ├── signal.go - handle os signals
├── testdata - test json files for various edge cases
├── Dockerfile - build container using docker 
├── go.mod
├── main.go - entry point
├── README.md
```


# Notes:

I feel like I should have added http api for CRUD operations on the db, but since it wasn't in the assignment, and it was recommended not to take more than 2 hours, I decided to leave the http api out.

Initially I planned to read the input file using multiple goroutines, and that's why I used mutex read/write locks in the database. However, I did not have enough time to implement parsing the input file in parallel, so the lock/unlock functionality is not used.

...