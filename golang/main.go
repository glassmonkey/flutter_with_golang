package main

import (
	"fmt"
	"github.com/glassmonkey/flutter_with_golang/golang/pkg/compute"
)

func main() {
	price, _ := compute.Add("123.456", "678.9")
	fmt.Print(price)
}