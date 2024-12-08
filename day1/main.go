package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var column1, column2, sum []int

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

	sort.Ints(column1)
	sort.Ints(column2)

	// Output the contents of the two variables
	fmt.Println("Column 1:", column1)
	fmt.Println("Column 2:", column2)

	// fmt.Print(column1[0])
	// fmt.Print(column2[0])

	// should be equal length columns
	for i, j := range column1 {
		sum = append(sum, j-column2[i])
	}
	fmt.Printf("%d", sum)

	total := 0
	for _, value := range sum {
		total += abs(value)
	}
	fmt.Printf("Total: %d", total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
