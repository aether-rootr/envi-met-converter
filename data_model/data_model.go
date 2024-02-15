package data_model

type DataPoint struct {
	GridX  int64
	GridY  int64
	MeterX float64
	MeterY float64
	Value  string // Prevent loss of accuracy
}

func GetMaxGridX(data_point_list []DataPoint) int64 {
	var max_grid_x int64 = 0
	for _, point := range data_point_list {
		if point.GridX > max_grid_x {
			max_grid_x = point.GridX
		}
	}
	return max_grid_x
}

func GetMaxGridY(data_point_list []DataPoint) int64 {
	var max_grid_y int64 = 0
	for _, point := range data_point_list {
		if point.GridY > max_grid_y {
			max_grid_y = point.GridX
		}
	}
	return max_grid_y
}
