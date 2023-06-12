# toll-calculator

# Docker (Kafka)
- docker run --name kafka -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest
- docker container prune


# Folder Structure
## On Board Unit (OBU)
- sends GPS coordinates

## data_receiver
- receives GPS coordinates

## distance_calculator

# types
- Shared data

# 3rd Party
- Gorilla Websocket
- [Confluent Kafka Go](https://github.com/confluentinc/confluent-kafka-go)
- sirupsen's logrus





# Resources
- [This](https://github.com/jinheehanaaa/toll-calculator)
- [Official Github](https://github.com/fulltimegodev/toll-calculator)