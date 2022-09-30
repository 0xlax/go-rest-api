# go-rest-api

### User Model
    - "id": "xxx",
    - "name": "test",
    - "dob": "",
    - "address": "",
    - "description": "",
    - "createdAt": ""

### Tech stack to be used

	Golang, NoSql DB - MongoDB suits the best :) 

### File Structure

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for better performance!

### Build 

	go build main.go

### Prerequisites
	
	1)Setup MongoDB locally on your machine
		curl -fsSL https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -
		echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list
		sudo apt install mongodb-org

	2) Install MongoDB 
		sudo apt install mongodb-org

	3) Start Service
		sudo systemctl start mongod.service

	4) Verify COnnection Status
		mongo --eval 'db.runCommand({ connectionStatus: 1 })'


