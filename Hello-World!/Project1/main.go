package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

/**
	1. flag.Parse - Parse parses the command line flags from the os.Args. This must be defined before
		after all the flags are defined and all the flags are accessed by the program.
	2. flag.String is used to define the flags.
	3. The Timer type represents a single event. When the timer expires, the current time will be
		sent on Channel c, unless the Timer was created by After func.
	4. Continously sending messages, use the Ticker and not the Timer.

 */

func main(){
	csvFileName := flag.String("csv","problems.csv","A CSV File in the form of values" +
		" and answers.")
	timeLimit := flag.Int("limit",30,"The time limit for the Quiz in seconds")

	flag.Parse()

	file,err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n",*csvFileName))
	}

	/**
		We have the Reader interface. As long as the file implements that, We can successfully
		read the data into the byte array. And we return number of bytes read and an error if occurs.
	 */

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll() //We read

	if err != nil {
		exit(fmt.Sprintf("Some error occured while reading the file %s",*csvFileName))
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i,p := range problems {
		fmt.Printf("Problem %d: %s\n",i+1,p.question)
		answerCh := make(chan string) //per question
		go func() {
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh <- answer
		}()
		select {
			case <- timer.C:
				fmt.Println("Times up!!")
				result(fmt.Sprintf("\nYou scored %d out of %d\n",correct,len(lines)))
			case answer := <-answerCh:
				if p.answer == answer {
					correct += 1
				}
		}

	}
	result(fmt.Sprintf("\nYou scored %d out of %d\n",correct,len(lines)))
}


//Defining the problem struct,
type Problem struct {
	question string
	answer string
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem,len(lines))

	for i,line := range lines {
		ret[i] = Problem{
			line[0],
			strings.TrimSpace(line[1]),
		}
	}
	return ret
}


//Better code style would be that we have an exit function specifically.
func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func result(score string){
	fmt.Println(score)
	os.Exit(0)
}