package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

func init() {
  var filename string
  var rootCmd = &cobra.Command{
    Use:   "quiz",
    Short: "Quiz is a simple quiz game written in Go.",
    Long: `Quiz is a simple quiz game written in Go. The
           Questions and Answers are read from a CSV file.`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Simple Quiz Game!\n------------------")
    }
  }
  rootCmd.Flags().StringVarP(&filename, "f", "filename", "idk what goes here",
                             "CSV Filename to read questions/answers from.")
  viper.SetDefault("filename", "quiz.csv")
}
