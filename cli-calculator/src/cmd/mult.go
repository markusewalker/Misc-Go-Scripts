// Authored by		: Markus Walker
// Date Modified	: 6/4/22

// Description		: Perform multiplication utilizing Cobra CLI tool.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// Provide CLI context on how to perform multiplication.
var multCmd = &cobra.Command{
	Use:   "mult",
	Short: "Perform multiplication",
	Long:  `Perform multiplication two between specified numbers, whole or floating point`,
	Run: func(cmd *cobra.Command, args []string) {
		multiplication()
	},
}

// Function to allow multiplying floating numbers.
func multiplication() {
	var num1, _ = strconv.ParseFloat(os.Args[2], 64)
	var num2, _ = strconv.ParseFloat(os.Args[3], 64)
	var result float64

	result = num1 * num2
	fmt.Println("Multiplication Result:", result)
}

func init() {
	rootCmd.AddCommand(multCmd)
}
