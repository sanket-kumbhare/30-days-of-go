// Problem: Combine everything. Build a tool that takes a file urls.txt with 10 URLs and tells which are UP/DOWN and response time.

// File:
// https://google.com
// https://github.com
// https://thisurldoesnotexist12345.com
// Output:

// Code
// google.com -> UP (120ms)
// github.com -> UP (300ms)
// thisurldoesnotexist12345.com -> DOWN (Error: no such host)
// You will learn: net/http, time, reading files, structs, error handling - everything from Week 1.

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	urlBytes, err := os.ReadFile("urls.txt")
// 	if err != nil {
// 		fmt.Println("Err:", err)
// 	}

// 	urls := strings.SplitSeq(string(urlBytes), "\n")

// 	for url := range urls {
// 		start := time.Now()
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Printf("%s -> DOWN (%s) \n", url, err)
// 			continue
// 		}

// 		timeTaken := time.Since(start)
// 		domain := resp.Request.URL

// 		fmt.Printf("%s -> UP statusCode: %d (%dms) \n", domain, resp.StatusCode, timeTaken.Milliseconds())

//		req.Body.Close()
// 	}
// }

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urlBytes, err := os.ReadFile("urls.txt") // you wrote url.txt, standard is urls.txt
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	lines := strings.SplitSeq(string(urlBytes), "\n")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for url := range lines {
		url = strings.TrimSpace(url)
		if url == "" {
			continue
		}

		start := time.Now()
		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("%s -> DOWN (%s)\n", url, err)
			continue
		}

		timeTaken := time.Since(start)
		fmt.Printf("%s -> UP statusCode: %d (%dms)\n", resp.Request.URL, resp.StatusCode, timeTaken.Milliseconds())

		resp.Body.Close()
	}
}
