# wednesday

### Before starting run below command.

> docker run -d --name wednesday --restart="always" -e POSTGRESQL_USER="wednesday" -e POSTGRESQL_PASSWORD="wednesday" -e POSTGRESQL_DATABASE="wednesday_db" -p 5542:5432 centos/postgresql-10-centos7

Otherwise changes configs inside config/config.go file


### Then execute
> go run cmd/main.go
