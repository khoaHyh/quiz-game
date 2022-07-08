package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
)

type quizData struct {
	Question string
	Answer   string
}

// TODO: Need to allow user to customize filename via a flag
// 1. figure out how to use flags
// 2. figure out how to customize filename in golang
// 3. figure out how to run only the method to customize filename
//      when the flag is present
// *bonus*: make renaming method a module that we import

// Reads files in current directory to find a csv file to use
// and if there is more than one csv file then prompt the user to
// choose one
func SelectCsvFile() (string, error) {
	files, err := os.ReadDir("./")
	if err != nil {
		return "", errors.New("reading files in './' directory failed")
	}

	csvFileNames := make(map[int]string)
	var csvFileCounter = 1

	for _, f := range files {
		if path.Ext(f.Name()) == ".csv" {
			csvFileNames[csvFileCounter] = f.Name()
			fmt.Printf("%v %v\n", csvFileCounter, f.Name())
			csvFileCounter++
		}
	}
	fmt.Println("Multiple csv files were found. Enter the number that corresponds with the file of your choice.")

	var fileKeyChoice int
	fmt.Scanln(&fileKeyChoice)

	return csvFileNames[fileKeyChoice], err
}

func ReadCsvFile(csvFileName string) ([][]string, error) {
	// open file
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		return nil, errors.New("opening file failed")
	}
	fmt.Println("Successfully Opened CSV file")
	// remember to close the file after reading in the data
	defer csvFile.Close()

	csvReader, err := csv.NewReader(csvFile).ReadAll()

	return csvReader, err
}

func RunQuiz(readCsvFile [][]string) map[string]int {
	// Initiate Map to hold response values
	response := map[string]int{"correctAnswers": 0, "totalQuestions": len(readCsvFile)}

	// Iterate through each line of the csv file and prompt user with a question
	// while incrementing the appropriate response values
	for _, line := range readCsvFile {
		quiz := quizData{
			Question: line[0],
			Answer:   line[1],
		}

		fmt.Println(quiz.Question + " ")

		var userInput string
		fmt.Scanln(&userInput)

		// Check if the answer is correct
		if userInput == quiz.Answer {
			response["correctAnswers"]++
		}
	}
	return response
}

func main() {
	csvFileName, err := SelectCsvFile()

	if err != nil {
		log.Fatal(err)
	}

	readCsvFile, err := ReadCsvFile(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	response := RunQuiz(readCsvFile)

	fmt.Printf("Correct Answers: %v\n", response["correctAnswers"])
	fmt.Printf("Total Questions: %v\n", response["totalQuestions"])
}
