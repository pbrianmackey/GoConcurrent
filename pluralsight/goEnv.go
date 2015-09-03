package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Environ())
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	printEnvironment("USER")
	return
}

func printEnvironment(input string) {
	fmt.Println("Environment variable lookup result", os.Getenv(input))
}
