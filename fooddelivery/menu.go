package main

import (
	"fmt"
	"sync"
)

type Menu struct {
	id        string
	menuItems map[string]*MenuItem

	mu sync.RWMutex
}

func NewMenu(id string) *Menu {
	return &Menu{
		id:        id,
		menuItems: make(map[string]*MenuItem),
	}
}

func (m *Menu) AddMenuItem(menuItem *MenuItem) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.menuItems[menuItem.id] = menuItem
}

func (m *Menu) GetMenuItem(menuItemId string) (*MenuItem, error ) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	menuItem, ok := m.menuItems[menuItemId]
	if !ok {
		return nil, fmt.Errorf("menu item %s does not exist", menuItemId)
	}
	return menuItem, nil
}