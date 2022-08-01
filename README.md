# Ports manager

Ports manager is a containerized ports database app.

Ports manager loads data from a json file into a map in memory.

# Build and run from terminal

go build .  
go run .  

# Build and run with docker

docker build -t ports .  
docker run ports  

Packages
```
ports  
|  
|------- data  
|           |------- ports.json - json input file with ports data         
|------- database   
|           |------- db.go - contains a Database struct   
|                    with a map to hold data and a mutex for parallel processing,   
|                    plus CRUD methods: Get, Insert, Update, Delete  
|           |------- db_test.go - contains tests for database Get, Insert, Update, Delete  
|------- managers  
|           |------- manager.go - contains Manager struct   
|                    and NewManager and LoadData to create a parser and a database,   
|                    and load data from parser output into database  
|           |------- manager_test.go - contains tests for LoadData  
|------- parsers  
|           |------- parser.go - contains Parser struct with a file pointer and a json decoder and functions OpenFile, ParseNextRecord, CloseFile, and readPortCode  
|           |------- parser_test.go - contains tests for OpenFile, ParseNextRecord, CloseFile functions  
|------- signals  
|           |------- signal.go - terminates the service upon os signals such as SIGINT, SIGTERM SIGQUIT, SIGKILL  
|------- testdata - contains test data for various edge cases  
```

# Notes:

I feel like I should have added http api for CRUD operations on the db, but since it wasn't in the assignment, and it was recommended not to take more than 2 hours, I decided to leave the http api out.

Initially I planned to read the input file using multiple goroutines, and that's why I used mutex read/write locks in the database. However, I did not have enough time to implement parsing the input file in parallel, so the lock/unlock functionality is not used.

