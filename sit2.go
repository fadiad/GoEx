package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item2 struct {
	Day    string
	Date   string
	Max    string
	Min    string
	Chance string
}

func willItRain2(xDays int, cityName string) {
	item := getItem2(xDays, cityName)
	fmt.Println(item.Chance)
}

func nextRainDay2(xDays int, cityName string) {
	for i := xDays; i < 15; i++ {
		item := getItem2(i, cityName)
		// fmt.Println(item)
		if item.Chance > "5 %" { // we dont have up of 50% so i tried 5 %
			fmt.Println(item)
			break
		}
	}

}

func getItem2(xDays int, cityName string) item2 {
	c := colly.NewCollector()

	var item = item2{}
	c.OnHTML(".forecasts.days.wo-scrollbars", func(e *colly.HTMLElement) {
		e.ForEach("wo-forecast-day", func(i int, h *colly.HTMLElement) {
			if i == xDays {
				item.Day = h.ChildText("wo-date-day-of-week")
				item.Date = h.ChildText("wo-date-day-and-month")
				item.Max = h.ChildText(".max")
				item.Min = h.ChildText(".min")
				item.Chance = h.ChildText("wo-weather-characteristics-precipitation")
			}
		})

	})

	c.Visit("https://www.weatherandradar.co.uk/weather/" + cityName)

	return item
}
