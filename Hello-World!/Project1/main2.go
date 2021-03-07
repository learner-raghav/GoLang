package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem1 struct {
	question string
	answer string
}

func main(){

	csvFileName := flag.String("csv","Problems.csv","Enter the fileName to be parsed")
	timeLimit := flag.Int("limit",30,"The time limit for the Quiz in seconds")
	_ = timeLimit
	flag.Parse()

	file,err := os.Open(*csvFileName)
	if err != nil {
		exit1(fmt.Sprint("An error occured while opening the file %s",*csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit1(fmt.Sprintf("An error occured while parsing the file %s",*csvFileName))
	}

	problems := parseLines1(&lines)
	score := askQuestions(problems)

	fmt.Printf("You scored %d out of %d\n",score,len(lines))
}


func askQuestions(problems []Problem1) int {
	correct := 0
	for i,problem := range problems {
		fmt.Printf("Problem %d: %s\n",i+1,problem.question)
		var ans string
		fmt.Scanf("%s\n",&ans)
		if ans == problem.answer {
			correct += 1
		}
	}
	return correct
}

func parseLines1(lines *[][]string) []Problem1{
	problems := make([]Problem1,len(*lines))
	for i,line := range *lines {
		problem := Problem1{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
		problems[i] = problem
	}
	return problems
}


func exit1(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

