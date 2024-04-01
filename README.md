# Microservices Architecture with gRPC and Kafka

This project implements a microservices architecture using gRPC for internal communication and Kafka for event streaming. The setup involves a gateway that routes client requests, microservices for handling different operations, and Kafka producers and consumers for asynchronous message processing.

## Services

- `gateway`: The entry point for all client requests, translating JSON HTTP requests to gRPC.
- `producer-micro`: Handles incoming driver data and produces messages to Kafka.
- `consumer-micro`: Consumes messages from Kafka and persists them into the database.
- `write-db-micro`: Microservice responsible for database write operations.
- `read-db-micro`: Microservice responsible for database read operations.
- `distance-micro`: Calculates the distance traveled by drivers.
- `spam`: This is a simple spam services that spams the gateway with locations.

## Build and Run

Each service can be built and run using the provided Makefile commands:

- Build and run gateway: `make gate`
- Build and run producer-micro: `make prod`
- Build and run consumer-micro: `make con`
- Build and run write-db-micro: `make write`
- Build and run read-db-micro: `make read`
- Build and run distance-micro: `make dist`
- Build and run spam: `make spam`

To generate gRPC code from `.proto` files, run: `make proto`

### Prerequisites

- Go installed on your machine.
- Kafka running locally or accessible via network.
- Protocol Buffer Compiler (`protoc`) with Go support.
- Docker

## Docker Setup

This project includes a `docker-compose` file for easy setup of the microservices and MongoDB as containers. The `docker-compose` file orchestrates the creation and interconnection of all the necessary containers.

- Build the container : `docker-compose up -d`

## API Endpoint

The following endpoint is available on the gateway for drivers to send their data:

```
- http://localhost:3000/api/v1/drivers
- http://localhost:3000/api/v1/drivers/1

```
