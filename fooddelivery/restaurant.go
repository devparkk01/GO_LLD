package main

import "sync"

type Restaurant struct {
	id      string
	name    string
	address *Address
	menu    *Menu

	mu sync.RWMutex
}

func NewRestaurant(id, name string, address *Address) *Restaurant {
	return &Restaurant{
		id:      id,
		name:    name,
		address: address,
	}
}

func (r *Restaurant) GetOrCreateMenu() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.menu == nil {
		r.menu = NewMenu("menu:" + r.id)
	}
}

func (r *Restaurant) AddMenuItem(menuItem *MenuItem) {
	r.GetOrCreateMenu()
	r.menu.AddMenuItem(menuItem)
}

func (r *Restaurant) AddMenuItems(menuItems []*MenuItem) {
	r.GetOrCreateMenu()
	for _, item := range menuItems {
		r.menu.AddMenuItem(item)
	}
}