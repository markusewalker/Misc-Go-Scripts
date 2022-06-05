// Authored by		: Markus Walker
// Date Modified	: 6/4/22

// Description		: Perform subtraction utilizing Cobra CLI tool.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// Provide CLI context on how to perform subtraction.
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Perform subtraction",
	Long:  `Perform subtraction two between specified numbers, whole or floating point`,
	Run: func(cmd *cobra.Command, args []string) {
		subtraction()
	},
}

// Function to allow subtracting floating numbers.
func subtraction() {
	var num1, _ = strconv.ParseFloat(os.Args[2], 64)
	var num2, _ = strconv.ParseFloat(os.Args[3], 64)
	var result float64

	result = num1 - num2
	fmt.Println("Subtraction Result:", result)
}

func init() {
	rootCmd.AddCommand(subCmd)
}
