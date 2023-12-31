package main

import "log"

const kafkaTopic = "obudata"

// Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
