package main

import (
	"sync"

	"github.com/google/uuid"
)

// we have one cart for each user
// once the order has been placed, cart information like items, userid, etc is stored in
// order. User's cart is then cleared.
type Cart struct {
	id     string // cart id 
	userId string // user id it belongs to 
	items  map[string]*CartItem // key: productId, value: *CartItem 
	mu     sync.RWMutex
}


func(c *Cart) AddProduct(p *Product, quantity int)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	// get the cartItem corresponding to this product ID
	ci, ok := c.items[p.id]
	if ok {
		ci.quantity += quantity
	} else {
		c.items[p.id] = &CartItem{
			id : "cart-item:" + uuid.NewString(),
			cartId: c.id,
			productId: p.id,
			quantity: quantity,
			pricePerItem: p.pricePerItem,
		}	
	}
}

