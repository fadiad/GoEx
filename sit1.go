package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Day         string
	Temperature string
	Weather     string
	FeelsLike   string
	Wind        string
	Humidity    string
	Chance      string
	Amount      string
}

func WillItRain(xDays int, cityName string) {
	c := colly.NewCollector()

	var arr []string

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(day int, h *colly.HTMLElement) {

			if day == xDays {
				fmt.Println("================== day :", day, "==========================")

				arr = append(arr, h.ChildText("th span"))

				h.ForEach("td", func(j int, b *colly.HTMLElement) {
					arr = append(arr, b.Text)
				})

				item := item{
					Day:         arr[0],
					Temperature: arr[1],
					Weather:     arr[2],
					FeelsLike:   arr[3],
					Wind:        arr[4],
					Humidity:    arr[5],
					Chance:      arr[8],
					Amount:      arr[9],
				}

				fmt.Println(item.Chance)
			}
		})
	})

	c.Visit("https://www.timeanddate.com/weather/israel/" + cityName + "/ext")

}
