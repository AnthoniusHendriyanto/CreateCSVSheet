package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Function: generateRandomPhoneNumber
// Description: Function for generating random phone number
// Returns: It will return a string of phone number that starts with 628 and followed by 8 random digits
func generateRandomPhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	phonePrefix := "628"
	randomNumber := rand.Intn(90000000) + 10000000
	return phonePrefix + strconv.Itoa(randomNumber)
}

// Function: generateRandomAmount
// Description: Function for generating random amount
// Returns: It will return integer of random amount between 10000 and 50000
func generateRandomAmount() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(40001) + 10000
}

func main() {
	var filename string
	fmt.Print("Enter the filename for the CSV: ")
	_, err := fmt.Scanf("%s", &filename)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fileNameWithExtension := filename + ".csv"

	// We're set the column names
	columnNames := []string{
		"Beneficiary Bank Code",
		"Beneficiary Account Number",
		"Mobile Number",
		"Amount",
	}

	// Prompt the user to enter the number of rows they want to generate
	var numRows int
	fmt.Print("Enter the number of rows: ")
	_, err = fmt.Scanf("%d", &numRows)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Rows slice to store the generated data
	rows := make([][]string, numRows)

	totalAmount := 0

	// Generate rows
	for i := 0; i < numRows; i++ {
		rows[i] = []string{
			"GNESIDJA",
			"510654320",
			generateRandomPhoneNumber(),
		}

		// Sum the total amount
		amount := generateRandomAmount()

		totalAmount += amount

		rows[i] = append(rows[i], fmt.Sprintf("%d", amount))
	}

	// Create a new CSV file
	file, err := os.Create(fileNameWithExtension)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column names to the CSV file
	err = writer.Write(columnNames)
	if err != nil {
		fmt.Println("Error writing column names:", err)
		return
	}

	// Write rows to the CSV file
	for _, row := range rows {
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row:", err)
			return
		}
	}

	fmt.Println("Total amount:", totalAmount)
	fmt.Println("CSV file created successfully.")
}
