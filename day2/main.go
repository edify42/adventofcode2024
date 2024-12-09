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
	var safe []int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		// Split each line into fields by a delimiter (e.g., space or comma)
		row := strings.Fields(scanner.Text()) // Use Fields to split by any whitespace

		// loop through the column array
		rowSafe := true
		var ascending bool
		changed := false
		a, _ := strconv.Atoi(row[0]) // starting number
		b, _ := strconv.Atoi(row[1]) // second number
		if a < b {                   // type conversion - yes, it's a string but a number can be compared
			ascending = true
		} else {
			ascending = false
		}
		// fmt.Printf("Row %s is ascending %t\n", row, ascending)
		for i := 1; i < len(row); i++ { // start at 2nd column
			a, _ := strconv.Atoi(row[i])   // lead number
			b, _ := strconv.Atoi(row[i-1]) // trailing number
			if abs(a-b) > 3 || abs(a-b) == 0 {
				rowSafe = false
			}
			if (ascending && a < b) || (!ascending && a > b) {
				changed = true
			}
		}

		fmt.Printf("Row %s was safe? %t. Ascending? %t\n", row, rowSafe, ascending)
		if rowSafe && !changed {
			safe = append(safe, 1)
		} else {
			safe = append(safe, 0)
		}
	}
	total := 0
	fmt.Printf("Safe: %d\n", safe)
	for _, value := range safe {
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
