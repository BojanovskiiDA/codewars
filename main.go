package main

import (
	"CodeWars/kata"
	"fmt"
)

 

 
func main() {
	fmt.Println(kata.Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1},map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}))
}