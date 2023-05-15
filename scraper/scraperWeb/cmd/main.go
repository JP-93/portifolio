package main

import (
	"github.com/gocolly/colly"
	"github.com/zolamk/colly-postgres-storage/colly/postgres"
)

func main() {
	c := colly.NewCollector()
	storage := &postgres.Storage{
		URI:                "",
		VisitedTable:       "",
		CookiesTable:       "",
		MaxOpenConnections: 0,
	}

}
