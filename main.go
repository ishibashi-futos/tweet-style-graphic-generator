package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	go serve()
	shots()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("input is", scanner.Text())
}
