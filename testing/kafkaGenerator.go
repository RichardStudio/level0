package main

import (
	"context"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/kafka-go"
	"level0/models"
	"log"
	"time"
)

func main() {
	gofakeit.Seed(0)

	// Конфигурация Kafka
	kafkaURL := "localhost:9092"
	topic := "orders"

	// Инициализация Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	// Генерация данных и отправка в Kafka
	for {
		order := generateOrder()
		data, err := json.Marshal(order)
		if err != nil {
			log.Printf("Error marshalling order: %v", err)
			continue
		}

		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(order.OrderUID),
				Value: data,
			},
		)
		if err != nil {
			log.Printf("Error writing message to Kafka: %v", err)
			continue
		}

		log.Printf("Sent order: %v", order.OrderUID)

		// Ждём перед генерацией следующего сообщения
		time.Sleep(1 * time.Second)
	}
}

func generateOrder() models.Order {
	return models.Order{
		OrderUID:          gofakeit.UUID(),
		TrackNumber:       gofakeit.UUID(),
		Entry:             gofakeit.Word(),
		Delivery:          generateDelivery(),
		Payment:           generatePayment(),
		Items:             generateItems(),
		Locale:            gofakeit.LanguageAbbreviation(),
		InternalSignature: gofakeit.Word(),
		CustomerID:        gofakeit.Word(),
		DeliveryService:   gofakeit.Word(),
		Shardkey:          gofakeit.Word(),
		SmID:              gofakeit.Number(1, 1000),
		DateCreated:       time.Now(),
		OofShard:          gofakeit.Word(),
	}
}

func generateDelivery() models.Delivery {
	return models.Delivery{
		Name:    gofakeit.Name(),
		Phone:   gofakeit.Phone(),
		Zip:     gofakeit.Zip(),
		City:    gofakeit.City(),
		Address: gofakeit.Street(),
		Region:  gofakeit.State(),
		Email:   gofakeit.Email(),
	}
}

func generatePayment() models.Payment {
	return models.Payment{
		Transaction:  gofakeit.UUID(),
		RequestID:    "",
		Currency:     gofakeit.CurrencyShort(),
		Provider:     gofakeit.Word(),
		Amount:       gofakeit.Number(100, 1000),
		PaymentDt:    time.Now().Unix(),
		Bank:         gofakeit.Word(),
		DeliveryCost: gofakeit.Number(100, 1000),
		GoodsTotal:   gofakeit.Number(100, 1000),
		CustomFee:    gofakeit.Number(0, 100),
	}
}

func generateItems() []models.Item {
	itemCount := gofakeit.Number(1, 5)
	items := make([]models.Item, itemCount)
	for i := 0; i < itemCount; i++ {
		items[i] = models.Item{
			ChrtID:      gofakeit.Number(1, 10000),
			TrackNumber: gofakeit.UUID(),
			Price:       gofakeit.Number(100, 1000),
			Rid:         gofakeit.UUID(),
			Name:        gofakeit.Word(),
			Sale:        gofakeit.Number(0, 100),
			Size:        gofakeit.Word(),
			TotalPrice:  gofakeit.Number(100, 1000),
			NmID:        gofakeit.Number(1, 10000),
			Brand:       gofakeit.Company(),
			Status:      gofakeit.HTTPStatusCode(),
		}
	}
	return items
}
