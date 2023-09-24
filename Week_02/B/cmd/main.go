package main

import "github.com/ivan-penchev/30-minutes-go-learning/Week_02/B/controller"

// instantiate the controller
// mocked for demonstration
// in reality it cc.New() would be a router action, or a payload received function

func main() {
	cc := controller.CustomerController{}
	cc.New()
}
