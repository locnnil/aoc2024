package main

import (
	"fmt"

	"github.com/locnnil/aoc2024.git/pkg/env"
	"github.com/locnnil/aoc2024.git/pkg/request"
)

func main() {
	env.LoadEnv()

	token := env.GetOrDie("SESSION_TOKEN")
	in, err := request.ReadInput(1, token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
	fmt.Println(in)
}
