package main 

import (
	"fmt"
)

func main() {
	ad1 := NewAddress("indiranagar", "Bengaluru", "560008", 23.34, 43.33)
	ad2 := NewAddress("indiranagar", "Bengaluru", "560009", 23.88, 43.40)
	ad3 := NewAddress("indiranagar", "Bengaluru", "560011", 67.34, 87.33)
	ad4 := NewAddress("indiranagar", "Bengaluru", "560008", 98.34, 43.33)
	ad5 := NewAddress("indiranagar", "Bengaluru", "5600021", 56.34, 22.33)

	rest1 := NewRestaurant("rest-1" , "Punjabi restaurant", ad1)
	rest2 := NewRestaurant("rest-2" , "Hyderabadi restaurant", ad2)

	// Add menu and menu items 
	menuItems1 := []*MenuItem{
		NewMenuItem("item-1", "Paneer butter" , 213, true),
		NewMenuItem("item-2", "egg rice", 73, true),
		NewMenuItem("item-3", "Gulab jamun", 90, true), 
	}
	menuItems2 := []*MenuItem{
		NewMenuItem("item-1", "Chicken butter" , 213, true),
		NewMenuItem("item-2", "Aloo rice", 73, true),
		NewMenuItem("item-3", "Rasgulla", 90, true), 
	}


	rest1.AddMenuItem(menuItems1[0])
	rest1.AddMenuItem(menuItems1[1])
	rest1.AddMenuItem(menuItems1[2])
	rest2.AddMenuItems(menuItems2)


	// Add customers
	cust1 := NewCustomer("cust-1", "Rohan", "1243564", "rohan@gmail.com", ad3)

	// Add delivery partners 
	del1 := NewDeliveryPartner("del-1", "Sandeep", "123456765", true, ad4)
	del2 := NewDeliveryPartner("del-2", "Mandeep", "123456735", true, ad5)


	foodService := NewFoodDeliveryService()
	foodService.AddRestaurant(rest1)
	foodService.AddRestaurant(rest2)

	foodService.AddDeliveryPartner(del1)
	foodService.AddDeliveryPartner(del2)

	foodService.AddCustomer(cust1)

	foodService.AddToCart(cust1.id, rest1.id, menuItems1[0].id, 1)
	foodService.AddToCart(cust1.id, rest1.id, menuItems1[1].id, 2)

	order, err := foodService.PlaceOrder(cust1.id, &UPIPayment{})
	if err != nil {
		panic(err)
	}

	fmt.Println(order)

}