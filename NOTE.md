# Intro
- On-Board Unit (OBU) sends data to data_receiver
- Calculate distance from Kafka receiver

# Kafka Receiver & Logging
- Kafka produces topic
- On-Board Unit (OBU) gives their data to data_receiver
- NewDataReceiver produces new kafka data from On-Board Unit
- Logging is included in LogMiddleware instead of NewKafkaProducer (More manageable)

# Kafka Calculate Distance
- Consumer calculates distance