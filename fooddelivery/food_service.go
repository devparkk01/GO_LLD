package main

import (
	// "fmt"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type FoodDeliveryService struct {
	restaurants      map[string]*Restaurant
	deliveryPartners map[string]*DeliveryPartner
	customers        map[string]*Customer
	carts  map[string]*Cart // customerId -> Cart
	orders []*Order

	mu sync.RWMutex
}

var (
	FoodDeliveryServiceInstance *FoodDeliveryService 
	FoodDeliverOnce sync.Once
)

func NewFoodDeliveryService() *FoodDeliveryService {
	FoodDeliverOnce.Do(func(){
		FoodDeliveryServiceInstance = &FoodDeliveryService{
			restaurants: make(map[string]*Restaurant),
			deliveryPartners: make(map[string]*DeliveryPartner),
			customers: make(map[string]*Customer),
			carts: make(map[string]*Cart),
		}
	})
	return FoodDeliveryServiceInstance
}

func (f *FoodDeliveryService) AddRestaurant(rest *Restaurant) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.restaurants[rest.id] = rest
}

func (f *FoodDeliveryService) AddCustomer(cust *Customer) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.customers[cust.id] = cust
}

func (f *FoodDeliveryService) AddDeliveryPartner(del *DeliveryPartner) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.deliveryPartners[del.id] = del
}

func(f *FoodDeliveryService) getOrCreateCart(customerId string) *Cart{
	f.mu.RLock()
	cart, ok := f.carts[customerId]
	f.mu.RUnlock()
	if !ok {
		f.mu.Lock()
		defer f.mu.Unlock()
		cart = NewCart("cart+" + customerId, customerId)
		f.carts[customerId] = cart 
	}
	return cart 
}

func (f *FoodDeliveryService) AddToCart(customerId, restaurantId, menuItemId string, quantity int) error {
	// Get the cart for this customer
	cart := f.getOrCreateCart(customerId)

	f.mu.RLock()
	restaurant := f.restaurants[restaurantId]
	f.mu.RUnlock()

	// check if this item is available in this restaurant 
	item, err := restaurant.menu.GetMenuItem(menuItemId)
	if err != nil {
		return err 
	} 
	err = cart.AddItem(restaurantId, item, quantity)
	if err != nil {
		return err 
	}
	return nil 

}

func(f *FoodDeliveryService) PlaceOrder(customerId string, paymentStrategy PaymentStrategy) (*Order, error) {
	// Get the cart for this customer
	cart := f.getOrCreateCart(customerId)

	if len(cart.items) == 0 {
		return nil, fmt.Errorf("cart is empty")
	}
	var amount float32
	amount = 0 
	for _, item := range cart.items {
		amount += item.price * float32(item.quantity)
	}
	// Create a new order 
	o := &Order{
		id : "order:" + uuid.NewString(),
		customerId: customerId,
		restaurantId: cart.restaurantId,
		items: cart.items,
		createdAt: time.Now(),
		amount: amount,
		orderStatus: OrderStatusPending,
		deliveryPartner: nil,
	}

	f.mu.Lock()
	f.orders = append(f.orders, o)
	defer f.mu.Unlock()

	cart.ClearCart()

	err := paymentStrategy.ProcessPayment(o.amount)
	if err != nil {
		o.orderStatus = OrderStatusCancelled
		return o, err 
	}

	o.orderStatus = OrderStatusConfirmed
	
	var assignedDel *DeliveryPartner
	for _, del := range f.deliveryPartners {
		if del.checkAvailable() {
			assignedDel = del 
			break 
		}
	}
	if assignedDel == nil {
		fmt.Println("Unable to find delivery partner now. retrying after few mins.")
		return o, nil 
	}
	o.deliveryPartner = assignedDel
	fmt.Printf("assigned delivery partner %s\n", assignedDel.name)
	assignedDel.SetAvailability(false)
	// notify delivery partner
	return o , nil 
}