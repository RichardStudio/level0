// database/cache_test.go
package database

import (
	"github.com/stretchr/testify/assert"
	"level0/models"
	"testing"
)

func TestMyCache(t *testing.T) {
	cache := NewMyCache()
	order := models.Order{
		OrderUID: "test_order_uid",
	}

	// Add order to cache
	cache.Add(order)

	// Get order from cache
	retrievedOrder, err := cache.Get("test_order_uid")
	assert.NoError(t, err)
	assert.Equal(t, order, retrievedOrder)

	// Get non-existent order
	_, err = cache.Get("non_existent_order_uid")
	assert.Error(t, err)
	assert.Equal(t, "order not found", err.Error())
}
