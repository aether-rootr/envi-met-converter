package main

import (
	"fmt"
	"os"

	io "github.com/aetherrootr/envi-met-converter/io"
	"github.com/spf13/cobra"
)

var input_file string
var output_file string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A command line tool to convert envi-met output data",
	Run: func(cmd *cobra.Command, args []string) {
		var data_point_list = io.ReadDataFromCsv(input_file)
		var matrix = io.GenerateOutputMatrix(data_point_list)
		io.OutputToCsvFile(output_file, matrix)
	},
}

func init() {
	convertCmd.PersistentFlags().StringVarP(&input_file, "input_file", "i", "", "Input file path")
	convertCmd.PersistentFlags().StringVarP(&output_file, "output_file", "o", "", "Output file path")
}

func Execute() error {
	return convertCmd.Execute()
}

func main() {
	if err := Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
