package site2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type item2 struct {
	Day    string
	Date   string
	Max    string
	Min    string
	Chance string
}

func WillItRain(xDays int, cityName string) {
	item := getItem(xDays, cityName)
	fmt.Println(item.Chance)
}

func NextRainDay(xDays int, cityName string) {
	for i := xDays; i < 15; i++ {
		item := getItem(i, cityName)
		fmt.Println(item)
		if item.Chance > "5 %" { // we dont have up of 50% so i tried 5 %
			fmt.Println(item)
			break
		}
	}

}


func AverageTemp(xDays int, cityName string) {
	tempSum := 0
	for i := 0; i < xDays; i++ {
		item := getItem(i, cityName)
		tempValStr := strings.Split(item.Max, "\n °")[0]
		intTemp, err := strconv.Atoi(tempValStr)

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}

		tempSum += intTemp

	}

	fmt.Printf("the vverage temperature of %s city is %d °C", cityName, tempSum/(xDays))

}

func TempArray(xDays int, cityName string) {

	var maxMinTemps [] string
	for i := 0; i < xDays; i++ {
	    item := getItem(i, cityName)
		tempValStrMax := strings.Split(item.Max, "\n °")[0]
		tempValStrMin := strings.Split(item.Min, "\n °")[0]
		maxMinTemp := tempValStrMax + " \\ "+ tempValStrMin+" °C" 
	 	maxMinTemps=append(maxMinTemps, "Max\\min: "+ maxMinTemp) 
    }
	fmt.Println(maxMinTemps)

}

func WeatherSummary(xDays int, cityName string) {	
   item := getItem(0, cityName)
   fmt.Println(item)
}

func getItem(xDays int, cityName string) item2 {
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
