# go-query-rabbitMQ

docker hub: https://hub.docker.com/r/dedazamarks/go-query-rabbitmq

**Require**
* Go 1.17
* RabbitMQ

**RabbitMQ**
* Run Server: `rabbitmq-server –detached`

**Run Search**
* Client: `go run main.go`
* Pathfinder queue: `go run pathfinder/pathfinder.go`
* Pathfinder server: `go run server/*.go`
