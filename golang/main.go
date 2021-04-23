package main

import (
	"fmt"
	"github.com/glassmonkey/flutter_with_golang/golang/pkg/api"
)

func main() {
	r, _ := api.Evaluate("1+1")
	fmt.Print(r)
}