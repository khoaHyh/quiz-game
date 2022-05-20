package main

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type quizData struct {
	Question string
	Answer   string
}

func ReadCsvFile() ([][]string, error) {
	// open file
	csvFile, err := os.Open("test.csv") // TODO: need to do wildcard *.csv
	if err != nil {
		return nil, errors.Wrap(err, "opening file failed")
	}
	fmt.Println("Successfully Opened CSV file")
	// remember to close the file at the end of the program
	defer csvFile.Close()

	// read csv value
	csvReader, err := csv.NewReader(csvFile).ReadAll()

	return csvReader, err
}

func main() {
	readCsvFile, err := ReadCsvFile()
	if err != nil {
		errors.WithMessage(err, "something errored?")
		os.Exit(1)
	}
	correctAnswers := 0
	totalQuestions := 0
	for _, line := range readCsvFile {
		quiz := quizData{
			Question: line[0],
			Answer:   line[1],
		}
		// Print question
		fmt.Println(quiz.Question + " ")
		totalQuestions++
		var userInput string
		// Take input from user
		fmt.Scanln(&userInput)

		// Check if the answer is correct and increment variable
		// that keeps track of correct answers
		if userInput == line[1] {
			correctAnswers++
		}
	}
	fmt.Printf("Correct Answers: %v\n", correctAnswers)
	fmt.Printf("Total Questions: %v\n", totalQuestions)
}
