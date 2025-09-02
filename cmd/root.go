package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "openpalette",
    Short: "A CLI tool for generating color palettes",
    Long:  `OpenPalette is a CLI tool for generating various color palettes and themes.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
}