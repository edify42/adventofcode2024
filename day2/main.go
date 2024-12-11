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
	var obj [][]string

	// Read the file line by line
	for scanner.Scan() {
		// Split each line into fields by a delimiter (e.g., space or comma)
		row := strings.Fields(scanner.Text()) // Use Fields to split by any whitespace
		obj = append(obj, row)
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

		// fmt.Printf("Row %s was safe? %t. Ascending? %t\n", row, rowSafe, ascending)
		if rowSafe && !changed {
			safe = append(safe, 1)
		} else {
			safe = append(safe, 0)
		}
	} // end file scanner
	total := 0
	// fmt.Printf("Safe: %d\n", safe)
	for _, value := range safe {
		total += abs(value)
	}
	fmt.Printf("Total Good: %d\n", total)

	// part 2 - builds on part 1 as we already have the unsafe count
	// var newSafe []int

	// remove first bad number in a row and mark it as 'damped'. If we go through the rest it's ok.
	// if we find another bad number, we need to start again
	badRowId := 0
	fmt.Print("looping through obj")
	for i, j := range obj {
		row := j // Use Fields to split by any whitespace
		if safe[i] == 0 {
			fmt.Printf("Row %s is unsafe. Index is %d - Damping...\n", row, badRowId)
			var ascending bool
			changed := false
			intRow, err := convertToIntSlice(row)
			if err != nil {
				fmt.Println("Error converting string slice to int slice:", err)
				return
			}
			if intRow[] < b {                   // type conversion - yes, it's a string but a number can be compared
				ascending = true
			} else {
				ascending = false
			}
			for i := 0; i < len(row); i++ {

			}
		}
		badRowId++
	}
}

// check row will measure the average and say whether it is ascending or descending
func checkRowAscending(row []int) bool {
	ascending := true
	start, _ := strconv.Atoi(row[0])
	end, _ := strconv.Atoi(row[len(row)-1])
	return ascending
}

// checks whether two numbers are indeed increasing/decreasing
func checkGoodness(a int, b int, increasing bool) bool {
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func convertToIntSlice(strSlice []string) ([]int, error) {
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting string to int at index %d: %v", i, err)
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
