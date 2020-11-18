package main

import (
	"bufio"
	"fmt"
	"os"
)

// For tracing on small inputs
func TestClient() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	fmt.Println(text)
}
