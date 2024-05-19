package main

import (
	"fmt"
)

// assembleYtDlpCommand assembles the command to download a video using yt-dlp
func assembleYtDlpCommand(url, cookiesFilename string, headers map[string]string, outputFilename string) []string {
	command := []string{"yt-dlp", "-o", outputFilename, "--cookies", cookiesFilename}
	for key, value := range headers {
		command = append(command, "--add-header", fmt.Sprintf("%s: %s", key, value))
	}
	command = append(command, url)
	return command
}
