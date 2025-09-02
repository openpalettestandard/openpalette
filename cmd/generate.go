package cmd

import (
	"fmt"
	"os"

	"github.com/openpalettestandard/openpalette/internal/palette"
	"github.com/openpalettestandard/openpalette/internal/types"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate various color palettes and themes",
	Long:  `Generate command allows you to generate palette.json with your complete palette implementation`,
}

var paletteCmd = &cobra.Command{
	Use:   "palette",
	Short: "Generate your color palette",
	Long:  `Generate your complete color palette in JSON format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		outputFile, _ := cmd.Flags().GetString("output")
		configFile, _ := cmd.Flags().GetString("config")
		versionFlag, _ := cmd.Flags().GetString("version")

		var paletteData types.PaletteResult
		var err error

		if configFile != "" {
			paletteData, err = palette.GenerateFromConfig(configFile)
			if err != nil {
				return fmt.Errorf("failed to generate from config: %w", err)
			}
		} else {
			paletteData = palette.Generate()
		}

		if versionFlag != "" {
			paletteData.Version = versionFlag
		}

		if outputFile == "" {
			if err := types.WriteJSON(paletteData, os.Stdout); err != nil {
				return fmt.Errorf("error writing JSON to stdout: %w", err)
			}
		} else {
			if err := types.WriteJSONFile(paletteData, outputFile); err != nil {
				return fmt.Errorf("error writing JSON file: %w", err)
			}
		}

		return nil
	},
}

var exampleCmd = &cobra.Command{
	Use:   "example-config",
	Short: "Generate an example configuration file",
	Long:  `Generate an example JSON configuration file that you can customize with your own colors.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, _ := cmd.Flags().GetString("output")
		if filename == "" {
			filename = "palette-config.json"
		}

		if err := palette.GenerateExampleConfig(filename); err != nil {
			return fmt.Errorf("failed to generate example config: %w", err)
		}

		fmt.Printf("Generated example config: %s\n", filename)
		fmt.Println("Edit this file with your custom colors, then use:")
		fmt.Printf("  openpalette generate palette -c %s\n", filename)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.AddCommand(paletteCmd)
	generateCmd.AddCommand(exampleCmd)

	paletteCmd.Flags().StringP("output", "o", "", "Output file path")
	paletteCmd.Flags().StringP("config", "c", "", "Configuration file (JSON format)")
	paletteCmd.Flags().StringP("version", "v", "", "Palette version (overrides config)")

	exampleCmd.Flags().StringP("output", "o", "", "Output config file")
}
