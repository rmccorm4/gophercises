package main

import (
  "bufio"
  "encoding/csv"
  "flag"
  "fmt"
  "os"
  "strings"
  "time"
)

// Simple struct to hold Quiz Questions/Answers
type Problem struct {
  question string
  answer   string
}

// Expects a 2-column CSV file
func readCSV(filename string) ([]Problem) {
  csvfile, err := os.Open(filename)


  reader := csv.NewReader(csvfile)
  lines, err := reader.ReadAll()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }


  // Create a slice of Problems to fill from CSV file
  problems := make([]Problem, len(lines))
  for i, line := range lines {
    problems[i] = Problem { question: line[0], answer: strings.TrimSpace(line[1]) }
  }

  return problems
}

// Handle Event Loop for Quiz
func startQuiz(problems []Problem, timeLimit int) {
  input := bufio.NewReader(os.Stdin)
  correct := 0

  timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

  // Use a label so we can break out of for loop from select statement
  problemLoop:
    for _, p := range problems {
      fmt.Println("Question:", p.question, "?")
      responseChannel := make(chan string)
      // Go-Routine to read answers asynchronously
      go func() {
        response, _ := input.ReadString('\n')
        response = strings.TrimSpace(response)
        responseChannel <- response
      }()

      select {
        // Timer channel has given a value == Timer has expired
        case <-timer.C:
          fmt.Println("Out of time!")
          break problemLoop
        // Timer hasn't expired and we received a response from channel
        case response := <-responseChannel:
          if response == p.answer {
            correct++
          }
      }
    }

  fmt.Printf("Total Score: %d/%d\n", correct, len(problems))
}

// Get CLI args and Run Quiz
func main() {
  var filenamePtr = flag.String("filename", "quiz.csv",
                                "CSV file containing quiz questions/answers.")
  var timeLimitPtr = flag.Int("time", 30, "Time limit to finish the quiz (seconds).")
  flag.Parse()
  problems := readCSV(*filenamePtr)
  startQuiz(problems, *timeLimitPtr)
}
