package main 

import (
	"fmt"
)


func main() {
	
	gun, err := getGun("AK47")
	if err != nil {
		fmt.Printf("Error while creating gun. %s\n", err.Error())
		return 
	}
	fmt.Printf("Created Gun: Name: %s, Power: %d\n", gun.getName(), gun.getPower())

	anotherGun, err := getGun("Musket")
	if err != nil {
		fmt.Printf("Error while creating gun. %s\n", err.Error())
		return 
	}
	fmt.Printf("Created Gun: Name: %s, Power: %d\n", anotherGun.getName(), anotherGun.getPower())
	
	faultyGun, err := getGun("faulty")
	if err != nil {
		fmt.Printf("Error while creating gun. %s\n", err.Error())
		return 
	}
	fmt.Printf("Created Gun: Name: %s, Power: %d\n", faultyGun.getName(), faultyGun.getPower())


}