package main

import (
	"fmt"
	"github.com/ivkoandrv/simple-rest-go/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
