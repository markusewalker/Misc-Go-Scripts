// Authored by		: Markus Walker
// Date Modified	: 6/4/22

// Description		: Perform division utilizing Cobra CLI tool.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// Provide CLI context on how to perform division.
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Perform division",
	Long:  `Perform division two between specified numbers, whole or floating point`,
	Run: func(cmd *cobra.Command, args []string) {
		division()
	},
}

// Function to allow dividing floating numbers.
func division() {
	var num1, _ = strconv.ParseFloat(os.Args[2], 64)
	var num2, _ = strconv.ParseFloat(os.Args[3], 64)
	var result float64

	result = num1 / num2
	fmt.Println("Division Result:", result)
}

func init() {
	rootCmd.AddCommand(divCmd)
}
