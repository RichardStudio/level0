package kafka

import (
	"context"
	"encoding/json"
	"level0/database"
	"level0/models"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeKafkaMessages(dbClient *database.DatabaseClient, broker string, topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "order_service_group",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		var order models.Order
		if err := json.Unmarshal(m.Value, &order); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		if !isOrderComplete(order) {
			log.Printf("Incomplete order data: %v", order)
			continue
		}

		database.Cache.Add(order)
		dbClient.SaveOrder(order)
	}
}

func isOrderComplete(order models.Order) bool {
	err := order.Validate()
	if err != nil {
		log.Printf("Error validating order: %v", err)
	}
	return err == nil
}
