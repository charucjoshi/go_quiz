package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	nameFlag := flag.String("name", "file.csv", "Name of the csv file.")
	timeFlag := flag.Int("time", 30, "Time in seconds.")
	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeFlag) * time.Second)

	file, err := os.Open(*nameFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Quiz Time!")

	score := 0
probloop:
	for qn, row := range data {
		answerCh := make(chan string)
		fmt.Printf("Q:%d %s\n", qn+1, row[0])
		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- input
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break probloop
		case input := <-answerCh:
			if input == strings.TrimSpace(row[1]) {
				score++
			}
		}
	}

	fmt.Printf("Your Score = %d", score)
}
