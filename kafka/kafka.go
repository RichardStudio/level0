package kafka

import (
	"context"
	"encoding/json"
	"level0/database"
	"level0/models"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeKafkaMessages(dbClient *database.DatabaseClient) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
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

		database.Cache[order.OrderUID] = order
		dbClient.SaveOrder(order)
	}
}

func isOrderComplete(order models.Order) bool {
	// Проверяем все поля
	return order.OrderUID != "" && order.TrackNumber != "" && order.Entry != "" &&
		order.Delivery.Name != "" && order.Delivery.Phone != "" && order.Delivery.Zip != "" &&
		order.Delivery.City != "" && order.Delivery.Address != "" && order.Delivery.Region != "" &&
		order.Delivery.Email != "" && order.Payment.Transaction != "" && order.Payment.Currency != "" &&
		order.Payment.Provider != "" && order.Payment.Amount != 0 && order.Payment.PaymentDt != 0 &&
		order.Payment.Bank != "" && order.Payment.DeliveryCost != 0 && order.Payment.GoodsTotal != 0 &&
		len(order.Items) > 0 && order.Locale != "" && order.CustomerID != "" && order.DeliveryService != "" &&
		order.Shardkey != "" && order.SmID != 0 && !order.DateCreated.IsZero() && order.OofShard != ""
}
