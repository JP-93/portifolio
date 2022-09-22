package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var url = "https://www.espn.com.br/nfl/classificacao"

type Nacional struct {
	cn string
}
type Americana struct {
	ca string
}

func scraper(url string) {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("request url", request.URL)
	})
	c.OnHTML("div.flex", func(e *colly.HTMLElement) {
		t := e.ChildText("abbr")
		edit := strings.Replace(t, "BUFMIANYJNECLEBALPITCINHOUINDTENJAXLVDENLACKC", "Conferencial Americana", -1)
		fi := strings.Replace(edit, "NYGWSHPHIDALCHIDETGBMINATLCARTBNOSFARILARSEA", "Conferencial Nacional", -18)
		r := Nacional{cn: fi}
		fmt.Println(r)
	})
	c.OnHTML("tbody.Table__TBODY", func(i *colly.HTMLElement) {
		r := i.ChildText("span.stat-cell")
		t := Americana{ca: r}
		fmt.Println(t)
	})

	c.Visit(url)
}
func main() {
	scraper(url)

}
