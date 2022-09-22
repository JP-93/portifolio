package scraper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://lista.mercadolivre.com.br/computador-gamer#D[A:computador%20gamer]"

type Products struct {
	Description string `json:"description"`
	Price       string `json:"price"`
}

func getResponse(url string) *http.Response {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(fmt.Sprint("status code error: %d %s", resp.StatusCode, resp))
	}

	return resp
}

func StarScraper() {
	res := getResponse(url)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	doc.Find("div.shops-custom-secondary-font").Each(func(i int, s *goquery.Selection) {
		pr := Products{
			Description: s.Find("h2.shops__item-title").Text(),
		}
		if pr.Description != "" {
			fmt.Println(pr)
		}

	})
	doc.Find("div.shops-custom-secondary-font").Each(func(i int, s *goquery.Selection) {
		pr := Products{
			Price: s.Find("span.price-tag-fraction").Text(),
		}
		if pr.Price != "" {
			conv, _ := strconv.ParseFloat(pr.Price, 64)
			if conv != 0 {
				fmt.Printf("%.4f\n", conv)
			}

		}
	})

}
