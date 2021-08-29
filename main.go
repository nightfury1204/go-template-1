package main

import (
	"log"

	"go-template/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "template",
	Short:   "A simplate go template",
	Version: "1.0.0",
}

func init() {
	rootCmd.AddCommand(cmd.ServeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
