# Go RabbitMQ Sample

*Project created to improve knowledge in RabbiMQ with Go Lang*

## Description

This Project contains two scripts, Publisher and Subscriber. The Publisher send one simple message to "hello" queue and the Subscriber make subscription on a "hello" queue and print all messages to received.


## Get Started

- First run docker-compose to create local RabbitMQ Broker

```
docker-compose up -d
```

- Second, start Subscriber server:

```
cd subscriber && go run src/*.go
```

- Third execute publisher script

```
cd publisher && go run src/*.go
```