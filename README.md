# go-query-rabbitMQ

**Require**
* Go 1.17
* RabbitMQ

**RabbitMQ**
* Run Server: `rabbitmq-server –detached`

**Run Search**
* Client: `go run main.go`
* Pathfinder queue: `go run pathfinder/pathfinder.go`
* Pathfinder server: `go run server/*.go`
