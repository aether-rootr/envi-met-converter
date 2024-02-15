package io

import (
	"strconv"

	data_model "github.com/aetherrootr/envi-met-converter/data_model"
)

func GenerateOutputMatrix(data_point_list []data_model.DataPoint) [][]string {
	var x_size = data_model.GetMaxGridX(data_point_list) + 1
	var y_size = data_model.GetMaxGridY(data_point_list) + 1
	var matrix = make([][]string, y_size+1)
	for i := range matrix {
		matrix[i] = make([]string, x_size+1)
	}

	var x_id = 0
	// Generate row and column ids
	for i := 1; i < int(x_size+1); i++ {
		matrix[0][i] = strconv.FormatInt(int64(x_id), 10)
		x_id++
	}
	var y_id = y_size - 1
	for i := 1; i < int(y_size+1); i++ {
		matrix[i][0] = strconv.FormatInt(int64(y_id), 10)
		y_id--
	}
	// convert data
	for _, data_point := range data_point_list {
		matrix[y_size-data_point.GridY][data_point.GridX+1] = data_point.Value
	}

	return matrix
}
