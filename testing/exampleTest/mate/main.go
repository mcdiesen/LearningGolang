package main

import (
	"fmt"

	_ "github.com/mcdiesen/LearningGolang/testing/exampleTest/"
)

func main() {
	mate.Sum(1)
	fmt.Println(mate.Sum(2, 4, 6))
	fmt.Println(mate.Sum(4, 7, 8, 0))
}
