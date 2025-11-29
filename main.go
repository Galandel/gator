package main

import (
	"fmt"
)

func main() {
	path, err := config.getConfigFilePath()
	if err {
		fmt.Println("Error occurred")
	}
	fmt.Println(path)
}
