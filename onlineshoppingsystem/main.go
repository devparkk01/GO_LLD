package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	catSvc := NewCatalogService()
	shopSvc := NewShoppingService(catSvc)

	// Create user
	user1 := &User{id: "user1", name: "Alice", email: "alice@gmail.com"}
	user2 := &User{id: "user2", name: "Bob", email: "bob@gmail.com"}

	// Create category
	c1 := &Category{id: "cat1", name: "Laptop"}
	c2 := &Category{id: "cat2", name: "Smartphone"}

	catSvc.AddCategory(c1)
	catSvc.AddCategory(c2)

	shopSvc.AddUser(user1)
	shopSvc.AddUser(user2)

	p1 := &Product{
		id: "prod1", name: "macbook Pro", description: "Macbook Pro 16 inches",
		categoryId:   c1.id,
		pricePerItem: 800000,
		quantity:     10,
	}

	p2 := &Product{
		id: "prod2", name: "macbook air", description: "Macbook air 16 inches",
		categoryId:   c1.id,
		pricePerItem: 50000,
		quantity:     8,
	}

	p3 := &Product{
		id: "prod3", name: "iphone", description: "iphone 16 Pro",
		categoryId:   c2.id,
		pricePerItem: 55000,
		quantity:     1,
	}

	catSvc.AddProduct(p1)
	catSvc.AddProduct(p2)
	catSvc.AddProduct(p3)

	shopSvc.AddToCart(user1.id, p1.id, 1)
	shopSvc.AddToCart(user1.id, p3.id, 1)

	order, _ := shopSvc.Checkout(user1.id, "idem:"+ uuid.NewString(),  &CardPayment{})
	fmt.Println(order)
	fmt.Println(shopSvc.orders)

	// Print user1 cart again after placing first order. it should be nil 
	_ , err := shopSvc.GetCart(user1.id)
	if err != nil {
		fmt.Println(err.Error())
	}

	shopSvc.AddToCart(user1.id, p3.id, 1)

	failedOrder, err := shopSvc.Checkout(user1.id, "idem:"+ uuid.NewString(), &CardPayment{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(failedOrder)

	products := catSvc.SearchProductByTitle("macbook")
	fmt.Println(products)


}