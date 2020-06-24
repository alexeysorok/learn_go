package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)

	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	// fmt.Println(&mustang)
	fmt.Println(mustang.topSpeed)
}
