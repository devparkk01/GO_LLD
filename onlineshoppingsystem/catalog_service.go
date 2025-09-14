package main

import (
	"fmt"
	"strings"
	"sync"
)

// manages categories, Products and inventories
type CatalogService struct {
	//categories
	categoriesMu sync.RWMutex
	categories map[string]*Category // categoryId -> Category 
	// products
	productsMu          sync.RWMutex
	products    map[string]*Product        // productId -> Product

}


func NewCatalogService() *CatalogService {
	return &CatalogService{
		categories: make(map[string]*Category),
		products: make(map[string]*Product),
	}
}

func(cs *CatalogService) AddCategory(c *Category) {
	cs.categoriesMu.Lock()
	defer cs.categoriesMu.Unlock()

	cs.categories[c.id] = c 
}

func(cs *CatalogService) AddProduct(p *Product) {
	cs.productsMu.Lock()
	defer cs.productsMu.Unlock()

	cs.products[p.id] = p 
}

func(cs *CatalogService) GetProduct(productId string) (*Product, error)  {
	cs.productsMu.RLock()
	defer cs.productsMu.RUnlock()

	p, ok := cs.products[productId]
	if !ok {
		return nil, fmt.Errorf("product %v not found", productId)
	}
	return p , nil 
}

// Increment increases available qty (used for restock or compensation).
func(cs *CatalogService) Increment(productId string, quantity int) error {
	cs.productsMu.Lock()
	defer cs.productsMu.Unlock()

	p, ok := cs.products[productId]
	if !ok {
		return fmt.Errorf("product not found")
	}
	p.UpdateQuantity(quantity)
	return nil 
}

// Decrement reduces available qty. returns error if insufficient.
func(cs *CatalogService) Decrement(productId string, quantity int) error {
	cs.productsMu.Lock()
	defer cs.productsMu.Unlock()

	p, ok := cs.products[productId]
	if !ok {
		return fmt.Errorf("product not found")
	}
	if p.GetQuantity() < quantity {
		return fmt.Errorf("requested quantity can not be fulfilled")
	}
	p.UpdateQuantity(-quantity)
	return nil 
}


func(cs *CatalogService) SearchProductByTitle(title string) []*Product {
	cs.productsMu.Lock()
	defer cs.productsMu.Unlock()
	products := make([]*Product, 0)

	for _, product := range cs.products  {
		if strings.Contains(strings.ToLower(product.description), strings.ToLower(title)) {
			products = append(products, product)
		}
	}
	return products 
}
