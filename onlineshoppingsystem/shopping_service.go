package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ShoppingService struct {
	usersMu sync.RWMutex
	users map[string]*User

	cartsMu sync.RWMutex
	carts   map[string]*Cart // userId -> Cart

	orderMu sync.RWMutex
	orders  map[string]*Order // orderID -> Order 

	idempotencyMu sync.RWMutex 
	idempotencyMap map[string]string // idempotencyKey -> orderId

	catalogSvc *CatalogService

	/*
	we could create userservice, cartService, orderservice, etc separately. 
	But for simplicity we have included them in one service. 
	*/
}

// Should be a single instance. Not handled here 
func NewShoppingService(catalogSvc *CatalogService) *ShoppingService {
	return &ShoppingService{
		users: make(map[string]*User),
		carts: make(map[string]*Cart),
		orders: make(map[string]*Order),
		idempotencyMap: make(map[string]string),
		catalogSvc: catalogSvc,
	}
}

func(s *ShoppingService) AddUser(u *User) {
	s.usersMu.Lock()
	defer s.usersMu.Unlock()
	s.users[u.id] = u 
}

func(s *ShoppingService) AddToCart(userId string, productId string, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	// Get the cart for this user 
	cart := s.GetOrCreateCart(userId)
	// Get the product 
	p, err := s.catalogSvc.GetProduct(productId)
	if err != nil {
		return err 
	}
	// Add this product to the cart  
	cart.AddProduct(p, quantity)
	return nil 

}

func(s *ShoppingService) GetOrCreateCart(userId string) *Cart {
	s.cartsMu.Lock()
	defer s.cartsMu.Unlock()

	cart, ok := s.carts[userId]
	if !ok {
		cartId := "cart:" + uuid.NewString()
		cart = &Cart{id: cartId, userId: userId, items: make(map[string]*CartItem)}
		s.carts[userId] = cart
	}
	return cart 
}

func(s *ShoppingService) GetCart(userId string) (*Cart, error ) {
	s.cartsMu.RLock()
	defer s.cartsMu.RUnlock()
	cart, ok := s.carts[userId]
	if !ok {
		return nil, fmt.Errorf("cart not found")
	}
	return cart, nil 
}

func(s *ShoppingService) ClearCart(userId string) {
	s.cartsMu.RLock()
	defer s.cartsMu.RUnlock()
	delete(s.carts, userId)
}

// Checkout: decrements inventory via catalog service, simulates payment, compensates on failure, records order idempotently.
func(s *ShoppingService) Checkout(userId string, idempotencyKey string, paymentStrategy PaymentStrategy ) (*Order, error) {
	// idempotency check 
	s.idempotencyMu.RLock() 
	orderId, ok := s.idempotencyMap[idempotencyKey]
	s.idempotencyMu.RUnlock()
	if ok {
		s.orderMu.RLock()
		order, ok := s.orders[orderId]
		s.orderMu.RUnlock()
		if ok {
			return order, nil 
		}
	}

	// Get the cart 	
	c, err := s.GetCart(userId)
	if err != nil {
		return nil, err
	}
	if len(c.items) == 0 {
		return nil, fmt.Errorf("no product found in the cart")
	}

	orderId = "order:" + uuid.NewString()
	now := time.Now() 
	var totalPrice float32

	orderItems := make([]*OrderItem, 0, len(c.items))
	for _, cartItem := range c.items {
		product, _ := s.catalogSvc.GetProduct(cartItem.productId)

		if (product.IsAvailable(cartItem.quantity)) {
			totalPrice += (product.pricePerItem *float32(cartItem.quantity))
			product.UpdateQuantity(-cartItem.quantity)
			orderItems = append(orderItems, &OrderItem{
				id: "order-item:" + uuid.NewString(),
				orderId: orderId,
				productId: cartItem.productId,
				quantity: cartItem.quantity,
				pricePerUnit: product.pricePerItem,
				createdAt: now,
			})
		} else {
			// remove all products 
			for _, orderItem := range orderItems {
				product, _ := s.catalogSvc.GetProduct(orderItem.productId)
				product.UpdateQuantity(orderItem.quantity)			
			}
			return nil, fmt.Errorf("one of the product quantity is not sufficient")
		}
	}

	// set cart to nil 
	s.ClearCart(userId)

	// Create an order 
	order := &Order{
		id: orderId, 
		userId: userId,
		totalPrice: totalPrice ,
		paymentStatus: PaymentStatus_Pending,
		items: orderItems,
		createdAt: now,
	}

	// map idempotency key to this order 
	s.idempotencyMu.Lock()
	s.idempotencyMap[idempotencyKey] = orderId
	s.idempotencyMu.Unlock()

	err = paymentStrategy.Pay(totalPrice)
	if err != nil {
		// update the quantities back to original
		for _, orderItem := range order.items {
			product, _ := s.catalogSvc.GetProduct(orderItem.productId)
			product.UpdateQuantity(orderItem.quantity)
		}
		order.paymentStatus = PaymentStatus_Failure
		return nil, fmt.Errorf("failed to make payment. %w", err)
	}
	order.paymentStatus = PaymentStatus_Success

	// Add order to the orders
	s.orderMu.Lock()
	s.orders[orderId] = order
	s.orderMu.Unlock()

	return order, nil
}

