package itgenius

import (
	"fmt"

	personal "github.com/AzureHoshi/go-learning-itgenius/itgenius/internal"
)

func HelloITGenius() {
	fmt.Println("Hello ITGenius!")
	personal.HelloPersonal()

}

func Add(a, b int) int {
	return a + b
}
