package car_assemble

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	values := productionRate * int(successRate) / 100
	return float64(values)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	v := convInt(successRate)
	c := (productionRate * v) / 60
	return c
}

func convInt(f float64) int {
	v := int(f)
	return v
}
