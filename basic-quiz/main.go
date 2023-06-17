package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

type flagParams struct {
	csvFileName string
	timeLimit   int
}

func main() {
	flagParams := getFlagParams()
	csvFilename := flagParams.csvFileName
	timeLimit := flagParams.timeLimit

	quizProblems := getQuizProblems(csvFilename)
	conductQuiz(quizProblems, timeLimit)
}

func getQuizProblems(csvFilename string) []problem {
	records := readProblemsFromCSV(csvFilename)
	problems := parseRecords(records)
	return problems
}

func conductQuiz(quizProblems []problem, timeLimit int) {
	correctCount := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	fmt.Println()

problemLoop:
	for index, problem := range quizProblems {
		fmt.Printf("Problem %d: %s = ", index+1, problem.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.answer {
				correctCount++
			}
		}
	}
	fmt.Printf("\nYou answered %d out of %d questions correctly\n", correctCount, len(quizProblems))
}

func getFlagParams() flagParams {
	csvFilenamePtr := flag.String("csv", "problems.csv", "Give data in `question,answer` format")
	timeLimit := flag.Int("time", 10, "Give time limit in seconds")
	flag.Parse()
	fmt.Println("CSV file name: ", *csvFilenamePtr)
	fmt.Println("Total time limit: ", *timeLimit)
	flags := flagParams{
		csvFileName: *csvFilenamePtr,
		timeLimit:   *timeLimit,
	}
	return flags
}

func readProblemsFromCSV(csvFilename string) [][]string {
	// open csv file
	file, err := os.Open(csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s", csvFilename))
	}
	fileReader := csv.NewReader(file)
	records, err := fileReader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to read content from file: %s", csvFilename))
	}
	return records
}

func parseRecords(records [][]string) []problem {
	problems := make([]problem, len(records))
	for index, line := range records {
		problems[index] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
