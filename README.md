# Kafka Microservices with Go

## Overview

This project implements a Kafka-based producer and consumer microservice using Go. The producer sends messages to Kafka, and the consumer reads and processes them.

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Prerequisites

Ensure you have the following installed:

- Kafka (running on localhost:9092)
- Zookeeper (running on localhost:2181)
- Go (latest version)
- Required Go packages: `github.com/segmentio/kafka-go`

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Setting Up Kafka

#### **Step 1: Open a Terminal and Start Zookeeper**

Open your terminal and run the following commands:

```sh
cd /home/dev41/kafka_2.13-3.7.0
bin/zookeeper-server-start.sh config/zookeeper.properties
```

#### **Step 2: Open a New Terminal and Start Kafka Broker**

In a new terminal, navigate to the Kafka directory and start the Kafka broker:

```sh
cd /home/dev41/kafka_2.13-3.7.0
bin/kafka-server-start.sh config/server.properties
```

### 2. Create a Kafka Topic

Run the following command to create a new Kafka topic named `my-topic`:

```sh
cd /home/dev41/kafka_2.13-3.7.0
bin/kafka-topics.sh --create --topic my-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

### 3. Verify the Created Topic

To verify that the topic was created successfully, run:

```sh
cd /home/dev41/kafka_2.13-3.7.0
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

If Kafka is running, this command will list available topics.

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Running the Services

### 1. Run the Producer Service

```sh
go run producer/main.go
```

You can then enter messages, and they will be sent to Kafka.

### 2. Run the Consumer Service

```sh
go run consumer/main.go
```

The consumer will listen for and print incoming messages.

-----------------------------------------------------------------------------------------------------------------------------------------------------

## How It Works

1. **Producer:** Takes input, converts it to JSON, and sends it to Kafka.
2. **Consumer:** Listens for messages, decodes JSON, and prints them.

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Example Run

### **Producer Input:**

```sh
Enter message: Hello, Kafka!
Enter message: This is a test message.
Enter message: {"event": "user_signup", "user": "john_doe"}
```

### **Consumer Output:**

```sh
Consumer started... Listening for messages.
Received: {ID:1 Content:Hello, Kafka! Time:2025-03-27T14:30:45Z}
Received: {ID:2 Content:This is a test message. Time:2025-03-27T14:31:10Z}
Received: {ID:3 Content:{"event": "user_signup", "user": "john_doe"} Time:2025-03-27T14:32:05Z}
```

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Testing the Setup

1. **Run the consumer first:**

```sh
go run consumer/main.go
```

2. **Run the producer and enter messages:**

```sh
go run producer/main.go
```
-----------------------------------------------------------------------------------------------------------------------------------------------------

