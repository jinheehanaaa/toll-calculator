# toll-calculator

# Kafka
- docker run --name kafka -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest


# On Board Unit (OBU)
- sends GPS coordinates

# data_receiver
- receives GPS coordinates

# types
- Shared data

# 3rd Party
- Gorilla Websocket
- Confluent Kafka Go