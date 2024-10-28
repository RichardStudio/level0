package database

import (
	"errors"
	"level0/models"
	"sync"
)

type MyCache struct {
	cache map[string]models.Order
	mutex sync.Mutex
}

func NewMyCache() *MyCache {
	return &MyCache{cache: make(map[string]models.Order)}
}

func (m *MyCache) Add(order models.Order) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.cache[order.OrderUID] = order
}

func (m *MyCache) Get(id string) (models.Order, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	order, ok := m.cache[id]
	if !ok {
		return models.Order{}, errors.New("order not found")
	}
	return order, nil
}
