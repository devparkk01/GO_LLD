package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Cart struct {
	id         string
	customerId string
	restaurantId string // restaurant id from which items have been added 
	items      []*OrderItem

	mu sync.RWMutex
}

func NewCart(id, customerId string) *Cart {
	return &Cart{
		id:         id,
		customerId: customerId,
	}
}

func (c *Cart) AddItem(restaurantId string, item *MenuItem, quantity int ) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.restaurantId == "" {
		c.restaurantId = restaurantId
	}

	if c.restaurantId != restaurantId {
		return fmt.Errorf("different restaurnt selected. you must clear cart")
	}

	c.items = append(c.items, &OrderItem{id: "orderItem:" + uuid.NewString(), name: item.name, price: item.price , quantity:  quantity} )
	return nil 
}


func(c *Cart) ClearCart() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.restaurantId = ""
	c.items = nil 
}
