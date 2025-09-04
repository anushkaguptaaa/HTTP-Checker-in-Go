package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type CheckResult struct {
	Url      string
	Status   string
	Duration time.Duration
	Error    error
}

func main() {
	fmt.Println("http checker is running...")

	// Intializing Struct Slice
	var AllResults []CheckResult

	// CLI Argument Handling
	cliArgument := os.Args[1:]

	if len(cliArgument) == 0 {
		fmt.Println("Usage: go run main.go <url1> <url2>")
		os.Exit(1)
	} else {
		for _, url := range cliArgument {
			start := time.Now()
			res, err := http.Get(url)
			if err != nil {
				badResult := CheckResult{
					Url:      url,
					Status:   "ERROR",
					Duration: 0,
					Error:    err,
				}
				AllResults = append(AllResults, badResult)
				continue
			}
			defer res.Body.Close()
			duration := time.Since(start)
			singleResult := CheckResult{
				Url:      url,
				Status:   res.Status,
				Duration: duration,
				Error:    err,
			}
			AllResults = append(AllResults, singleResult)
		}
	}

	//printing struct
	fmt.Printf("\nURL Results:")

	for _, result := range AllResults {
		fmt.Printf("\n---\n")
		fmt.Println("URL:", result.Url)
		fmt.Println("Status:", result.Status)
		fmt.Println("Duration:", result.Duration.Milliseconds())
		fmt.Println("Error:", result.Error)
	}
	// fmt.Printf("%+v\n", AllResults)
	// fmt.Printf("%#v\n", AllResults)
	// fmt.Println(AllResults)
}
