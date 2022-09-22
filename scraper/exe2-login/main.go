package main

import (
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	err := c.Post("https://isdes.mrooms.net/login/index.php", map[string]string{"Identificação de usuário": "2085094", "Senha": "42808239874"})
	if err != nil {
		log.Fatal(err)
	}

	c.OnResponse(func(response *colly.Response) {
		log.Print("Status response ", response.StatusCode)
	})
	c.Visit("https://isdes.mrooms.net/")
}
