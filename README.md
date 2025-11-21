start docker with mysql container

`docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:latest`

in a terminal run: `mysqlsh --uri root@localhost:3306` and then go to sql moce by running `\sql` keep it running just in case yo need to interact using CLI
from the terminal first create the database links. use normal mysql commands such as `create database mylinks;` or `drop database mylinks;` etc.

run `migrate create -ext sql -dir database/migrations -seq links` to create golang-migrate files
run `migrate -path database/migrations -database "mysql://root:my-secret-pw@tcp(localhost:3306)/linksdb" up` to run the up migration