// Authored by		: Markus Walker
// Date Modified	: 6/4/22

// Description		: Perform addition utilizing Cobra CLI tool.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// Provide CLI context on how to perform addition.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Perform addition",
	Long:  `Perform addition between two specified numbers, whole or floating point`,
	Run: func(cmd *cobra.Command, args []string) {
		addition()
	},
}

// Function to allow adding floating numbers.
func addition() {
	var num1, _ = strconv.ParseFloat(os.Args[2], 64)
	var num2, _ = strconv.ParseFloat(os.Args[3], 64)
	var result float64

	result = num1 + num2
	fmt.Println("Addition Result:", result)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
