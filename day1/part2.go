package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize slices to store values from both columns
	var column1, column2, replicas []int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		// Split each line into fields by a delimiter (e.g., space or comma)
		fields := strings.Fields(scanner.Text()) // Use Fields to split by any whitespace
		if len(fields) == 2 {                    // Ensure the line has exactly two columns
			a, _ := strconv.Atoi(fields[0])
			b, _ := strconv.Atoi(fields[1])
			column1 = append(column1, a) // First column
			column2 = append(column2, b) // Second column
		}
	}

	// Check for errors in the file scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// should be equal length columns
	for _, j := range column1 {
		replicas = append(replicas, countSpecificNumber(column2, j))
	}
	fmt.Printf("%d", replicas)

	total := 0
	for i, value := range replicas {
		total += value * column1[i]
	}
	fmt.Printf("Total: %d", total)
}

func countSpecificNumber(arr []int, target int) int {
	count := 0
	for _, num := range arr {
		if num == target {
			count++
		}
	}
	return count
}
