package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Chapter represents a row in the CSV file
type Chapter struct {
	Index       int
	Filename    string
	Title       string
	Description string
}

func main() {
	cookiesFilename := "../input/cookies.txt"
	urlsFilename := "../input/urls.txt"
	headersFilename := "../input/headers.txt"
	chaptersFilename := "../input/chapters.csv"

	urls, err := readUrls(urlsFilename)
	if err != nil {
		fmt.Printf("Error reading URLs: %v\n", err)
		return
	}

	headers, err := readHeaders(headersFilename)
	if err != nil {
		fmt.Printf("Error reading headers: %v\n", err)
		return
	}

	chapters, err := readChapters(chaptersFilename)
	if err != nil {
		fmt.Printf("Error reading chapters: %v\n", err)
		return
	}

	for i, url := range urls {
		command := assembleYtDlpCommand(url, cookiesFilename, headers, "../output/"+chapters[i].Filename)
		fmt.Printf("Downloading chapter %d into %s\n", chapters[i].Index, "../output/"+chapters[i].Filename)

		// Execute the command
		cmd := exec.Command(command[0], command[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command. The error was: %v\n", err)
			fmt.Printf("Command: %s\n", strings.Join(command, " "))
			fmt.Printf("Output: %s\n", output)

			// If output contains "has already been downloaded", don't return
			if strings.Contains(string(output), "has already been downloaded") {
				continue
			}
			fmt.Printf("Skipping %d\n", chapters[i].Index)
			continue
		}
		fmt.Printf("Output: %s\n", output)

		fmt.Printf("Waiting for 5 seconds...\n")
		time.Sleep(5 * time.Second)
	}
}
