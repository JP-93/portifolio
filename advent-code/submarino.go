package main

import (
	"fmt"
)

func main() {
	//filecsv := "data.csv"
	//opencsv, err := os.Open(filecsv)
	//errors(err)
	//defer opencsv.Close()
	//
	//data, err := csv.NewReader(opencsv).ReadAll()
	//errors(err)

	slice := []int{1, 2, 3, 4, 5}
	if slice[0] < slice[1] {
		fmt.Println("ok")
	}
}

func errors(err error) {
	if err != nil {
		fmt.Printf("find error", err)
	}
}
