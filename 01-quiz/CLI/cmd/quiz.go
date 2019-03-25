// Expects a 2-column CSV file
func readCSV(filename string) ([]string, []string) {
  csvfile, _ := os.Open(filename)
  reader := csv.NewReader(bufio.NewReader(csvfile))
  var questions []string
  var answers   []string

  for {
    line, error := reader.Read()
    if error == io.EOF { 
      break 
    } else if error != nil { 
      log.Fatal(error) 
    }

    questions = append(questions, line[0])
    answers   = append(answers, line[1])
  }

  return questions, answers
}

func startQuiz(filename string) {
  questions, answers := readCSV(filename)
  reader := bufio.NewReader(os.Stdin)
  correct := 0

  for i := range questions {
    fmt.Println("Question:", questions[i], "?")
    response, _ := reader.ReadString('\n')
    response = strings.TrimSpace(response)
    if response == answers[i] {
      correct++
    } else {
      fmt.Println("Correct Answer:", answers[i])
    }
  }
  fmt.Printf("Total Score: %d/%d\n", correct, len(questions))
}


