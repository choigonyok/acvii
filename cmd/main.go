package main

import (
	"fmt"

	"github.com/choigonyok/goopt/internal/test"
	example "github.com/choigonyok/goopt/pkg"
	"k8s.io/client-go/rest"
)

// go doc test
func main() {
	println(test.Test())
	example.Example()
	init := rest.Config{}
	fmt.Printf("init: %v\n", init)
}
