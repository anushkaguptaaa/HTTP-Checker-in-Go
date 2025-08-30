package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("http checker is running...")

	// CLI Argument Handling
	cliArgument := os.Args[1:]

	if len(cliArgument) == 0 {
		fmt.Println("Empty string found: Usage - go run main.go <url1> <url2>")
	} else {
		for _, url := range cliArgument {
			fmt.Printf("Checking Url: %s\n", url)
		}
	}

}
