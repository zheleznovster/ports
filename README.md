# Ports manager

Ports manager is a containerized ports database app.

Ports manager loads data from a json file into a map in memory.

# Build and run from terminal

go build .  
go run .  

# Build and run with docker

docker build -t ports .  
docker run ports  


# Notes:

I feel like I should have added http api for CRUD operations on the db, but since it wasn't in the assignment, and it was recommended not to take more than 2 hours, I decided to leave the http api out.

Initially I planned to read the input file using multiple goroutines, and that's why I used mutex read/write locks made the database. However, I did not have enough time to implement parsing the input file in parallel, so the lock/unlock functionality is not used.

