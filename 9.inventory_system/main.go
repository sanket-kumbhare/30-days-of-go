// Problem: Build type Product struct { Name string; stock int } with methods AddStock(n int), Sell(n int) error. If stock goes negative, return error. Use pointer receiver - if you use value receiver, stock won't update.

// Learn: Why (p *Product) is needed, errors.New

package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	stock int
}

func (p *Product) AddStock(n int) {
	p.stock += n
}

func (p *Product) Sell(n int) error {
	if p.stock < n {
		return errors.New("Less Stock")
	}
	p.stock -= n
	return nil
}

func (p Product) Stock() int {
	return p.stock
}

func main() {
	product := &Product{
		Name:  "Laptop",
		stock: 10,
	}
	product.AddStock(10)
	fmt.Println("Stocks Present:", product.Stock())

	err := product.Sell(10)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("Stocks Present:", product.Stock())
}
