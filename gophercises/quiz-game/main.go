package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	var filename string
	var timeLimit int64

	flag.StringVar(&filename, "csv", "problems.csv", "The csv file with the problems to be processed")
	flag.Int64Var(&timeLimit, "limit", 30, "The amount of time to answer the quiz")
	flag.Parse()

	var fileReader, err = os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	csvData := csv.NewReader(fileReader)

	problems, err := csvData.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var correct int

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	// <-timer.C

problemLoop:
	for index, problem := range problems {

		fmt.Printf("Problem %v: %v = ", index+1, problem[0])
		answerChan := make(chan string)

		go func() {
			var answer string

			fmt.Scan(&answer)
			answerChan <- answer

		}()

		select {

		case <-timer.C:
			fmt.Printf("\nYour score is %v out of %v", correct, len(problems))
			break problemLoop

		case answer := <-answerChan:
			if strings.TrimSpace(answer) == strings.TrimSpace(problem[1]) {
				correct++
			}

		}

	}

	fmt.Printf("Your score is %v out of %v", correct, len(problems))

}
