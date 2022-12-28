package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

	for index, problem := range problems {

		var answer string

		fmt.Printf("Problem %v: %v = ", index+1, problem[0])
		fmt.Scan(&answer)

		if strings.TrimSpace(answer) == strings.TrimSpace(problem[1]) {
			correct++
		}
		// fmt.Println()
	}

	fmt.Printf("Your score is %v out of %v", correct, len(problems))

}
