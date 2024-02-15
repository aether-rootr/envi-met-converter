package io

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	data_model "github.com/aetherrootr/envi-met-converter/data_model"
)

func ReadDataFromCsv(input_file_path string) []data_model.DataPoint {
	var data_point_list []data_model.DataPoint

	f, err := os.Open(input_file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(rec) == 5 {
			var grid_x, grid_y int64
			var meter_x, meter_y float64
			var value string = rec[4]
			grid_x, err = strconv.ParseInt(rec[0], 10, 64)
			if err != nil {
				continue
			}
			grid_y, err = strconv.ParseInt(rec[1], 10, 64)
			if err != nil {
				continue
			}
			meter_x, err = strconv.ParseFloat(rec[2], 64)
			if err != nil {
				continue
			}
			meter_y, err = strconv.ParseFloat(rec[3], 64)
			if err != nil {
				continue
			}
			data_point_list = append(data_point_list,
				data_model.DataPoint{
					GridX:  grid_x,
					GridY:  grid_y,
					MeterX: meter_x,
					MeterY: meter_y,
					Value:  value,
				})
		}
	}
	return data_point_list
}
