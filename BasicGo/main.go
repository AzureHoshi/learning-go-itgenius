package main

import (
	"fmt"

	"github.com/AzureHoshi/go-learning-itgenius/itgenius"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello World!")
	itgenius.HelloITGenius()
	fmt.Print("Add 1 + 2 = ", itgenius.Add(1, 2), "\n")

	id := uuid.New()
	fmt.Println(id)
}
