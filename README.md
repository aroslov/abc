# ABC

[![Gitter](https://badges.gitter.im/aroslov/abc.svg)](https://gitter.im/aroslov/abc?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

# Requirements

* Golang v1.5
* Rabbitmq

# Run 

* Check config files in /config
* Run the mock webserver: `go run main/mockserver.go`
* Run the generator: `go run main/producer.go 10`
* Run the consumer: `go run main/consumer.go`
