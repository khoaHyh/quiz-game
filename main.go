package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

type quizData struct {
	Question string
	Answer   string
}

type Color string

const (
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
)

func colorizeMessage(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func customizeFileName(fileName string) string {
	fmt.Println("Enter the new name of the csv file (include .csv at the end):")

	var newFileName string
	fmt.Scanln(&newFileName)

	os.Rename(fileName, newFileName)

	return newFileName
}

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

	requestToCustomizeFileName := flag.Bool("file", false, "Option to customize file name chosen")
	flag.Parse()

	if *requestToCustomizeFileName {
		csvFileName = customizeFileName(csvFileName)
	}

	if err != nil {
		log.Fatal(err)
	}

	csvFileContent, err := ReadCsvFile(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	response := RunQuiz(csvFileContent)

	colorizeMessage(ColorGreen, "Correct Answers:"+strconv.Itoa(response["correctAnswers"]))
	colorizeMessage(ColorYellow, "Total Questions:"+strconv.Itoa(response["totalQuestions"]))
}
