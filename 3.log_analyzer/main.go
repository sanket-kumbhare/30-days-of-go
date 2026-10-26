// Problem: You have a app.log file like:

// app.log
// INFO: User logged in
// ERROR: DB failed
// INFO: Payment done
// ERROR: Timeout
// Count how many INFO, ERROR, WARN lines. Also find the most frequent ERROR message.

// You will learn: os.ReadFile, strings.Split, map[string]int, error handling if err != nil

// Deliverable: go run main.go app.log

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("app.log")

	if err != nil {
		fmt.Println("Err:", err)
	}

	fileContent := string(file)

	lines := strings.Split(fileContent, "\n")

	logsMap := make(map[string]int)
	errMsgCount := make(map[string]int)

	for _, line := range lines {
		logs := strings.SplitN(line, ": ", 2)

		if len(logs) != 2 {
			continue
		}

		level := logs[0]
		msg := logs[1]

		if level == "ERROR" {
			errMsgCount[msg]++
		}
		logsMap[logs[0]]++
	}

	for key, value := range logsMap {
		fmt.Println(key, ": ", value)
	}

	var mostFreqMsg string
	max := 0
	for msg, count := range errMsgCount {
		if count > max {
			max = count
			mostFreqMsg = msg
		}
	}

	fmt.Printf("Most Frequent ERROR: %s (%d times)", mostFreqMsg, max)
}
