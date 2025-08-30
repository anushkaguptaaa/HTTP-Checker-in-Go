package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("http checker is running...")

	// CLI Argument Handling
	cliArgument := os.Args[1:]

	if len(cliArgument) == 0 {
		fmt.Println("Usage: go run main.go <url1> <url2>")
		os.Exit(1)
	} else {
		for _, url := range cliArgument {
			fmt.Printf("Checking Url: %s\n", url)
			start := time.Now()
			res, err := http.Get(url)
			if err != nil {
				fmt.Println("ERROR:\n", err)
				continue
			}
			defer res.Body.Close()
			duration := time.Since(start)
			fmt.Println(res.Status, "(", duration.Milliseconds(), "ms )")
		}
	}

}
