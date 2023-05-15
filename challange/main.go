package main

import (
	"fmt"

	"challange/m/database_ch"
	"challange/m/route_ch"
)

func main() {
	fmt.Println("Running")
	database_ch.ConnectDb()
	route_ch.HandlerRoute()
}
