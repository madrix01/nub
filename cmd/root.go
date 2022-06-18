package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "nub",
    Short: "Manage your markdown notes using github gist",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to nub")
    }, 
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
