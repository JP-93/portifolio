package car_assemble

import (
	"log"
	"testing"
)

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	want := 1392.00
	got := CalculateWorkingCarsPerHour(1547, 90)

	if want != got {
		log.Fatalf("want %b, but got %b", want, got)
	}
}

func TestCalculateWorkingCarsPerMinute(t *testing.T) {
	
}
