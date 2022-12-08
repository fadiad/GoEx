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

func willItRain(xDays int, cityName string) {
	item := getItem(xDays, cityName)
	fmt.Println(item.Chance)
}

func nextRainDay(xDays int, cityName string) {
	for i := xDays; i < 15; i++ {
		item := getItem(i, cityName)

		// fmt.Println(item)
		if item.Chance > "50%" {
			fmt.Println(item)
			break
		}
	}

}

func getItem(xDays int, cityName string) item {
	c := colly.NewCollector()

	var arr []string
	var newItem = item{}
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(day int, h *colly.HTMLElement) {

			if day == xDays {
				fmt.Println("================== day :", day, "==========================")

				arr = append(arr, h.ChildText("th span"))

				h.ForEach("td", func(j int, b *colly.HTMLElement) {
					arr = append(arr, b.Text)
				})

				// newItem =  {
				newItem.Day = arr[0]
				newItem.Temperature = arr[1]
				newItem.Weather = arr[2]
				newItem.FeelsLike = arr[3]
				newItem.Wind = arr[4]
				newItem.Humidity = arr[5]
				newItem.Chance = arr[8]
				newItem.Amount = arr[9]
				// }

				// fmt.Println(item.Chance)
			}
		})
	})

	c.Visit("https://www.timeanddate.com/weather/israel/" + cityName + "/ext")

	return newItem
}
