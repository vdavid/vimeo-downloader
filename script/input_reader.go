package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readHeaders reads headers from a file and returns them as a map
func readHeaders(filename string) (map[string]string, error) {
	headers := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header line: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		headers[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headers, nil
}

// readUrls reads URLs from a file and returns them as a slice of strings
func readUrls(filename string) ([]string, error) {
	var urls []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines
		if line != "" {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

// readChapters reads the CSV file and returns a slice of Chapter structs
func readChapters(filename string) ([]Chapter, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = ','       // Set the delimiter
	reader.LazyQuotes = true // Allows quotes in the data
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var chapters []Chapter
	for _, record := range records {
		if len(record) != 4 {
			return nil, fmt.Errorf("invalid record length: %v", record)
		}

		index, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid index: %v", record[0])
		}

		chapter := Chapter{
			Index:       index,
			Filename:    record[1],
			Title:       record[2],
			Description: record[3],
		}
		chapters = append(chapters, chapter)
	}

	return chapters, nil
}
