package main

import (
	"fmt"
	"sync"
)

// Cart - shopping cart
type Cart struct {
	Items []CartItem
}

func (cart *Cart) addItem(name string) {
	cart.Items = append(cart.Items, CartItem{Name: name})
}

func (cart *Cart) printItems() {
	fmt.Println("-- Printing cart items --")
	for i, item := range cart.Items {
		fmt.Println(fmt.Sprintf("%d - %s", i, item.Name))
	}
}

// CartItem - item
type CartItem struct {
	Name string
}

var (
	once     sync.Once
	instance *Cart
)

// GetInstance of shopping cart
func GetInstance() *Cart {
	once.Do(func() {
		fmt.Println("Instantiating a new cart!")
		instance = &Cart{}
	})
	return instance
}

func main() {
	cartInstance1 := GetInstance()
	cartInstance1.addItem("Banana")

	cartInstance2 := GetInstance()
	cartInstance2.printItems() // Should print banana

}
