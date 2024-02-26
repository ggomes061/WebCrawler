package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	var links []string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		links = append(links, link)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scrapping completed!")

		saveToFile(links, "links.txt")
	})

	err := c.Visit("https://nmap.org/")
	if err != nil {
		log.Fatal(err)
	}
}

func saveToFile(links []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error ao criar arquivo", err)
	}
	defer file.Close()

	for _, link := range links {
		_, err := file.WriteString(link + "\n")
		if err != nil {
			log.Fatal("Erro ao escrever arquivo", err)
		}
	}

	fmt.Printf("Links salvos em %s\n", filename)
}
