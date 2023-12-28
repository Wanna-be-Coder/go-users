db:
	docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=userdb -p 3306:3306 -d mysql:latest
server:
	go run main.go

.PHONY: server