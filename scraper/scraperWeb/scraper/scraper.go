package scraper

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const url = "https://www.amazon.com.br/s?k=iphone&__mk_pt_BR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=1JVNU41JERTN0&sprefix=iphone%2Caps%2C221&ref=nb_sb_noss_1"

type Description struct {
	Title string
	Price string
}

func StarScraper() {
	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal("Error creating", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"title", "Price"}
	writer.Write(header)

	c := colly.NewCollector()
	c.OnHTML("div.a-spacing-base", func(e *colly.HTMLElement) {
		data := Description{}
		data.Title = e.ChildText("span.a-text-normal")
		data.Price = e.ChildText("span.a-price-whole")
		row := []string{data.Title, data.Price}
		writer.Write(row)

	})
	c.Visit(url)
}
