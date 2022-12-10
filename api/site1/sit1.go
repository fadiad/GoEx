package site1

import (
	"fmt"
	"strconv"
	"strings"

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
	item := getItem(xDays, cityName)
	fmt.Println(item.Chance)
}

func WeatherSummary(cityName string) {	
	itemObj := getItem(0, cityName)
	var weatherSummary = item{}
	weatherSummary.Temperature=itemObj.Temperature
	weatherSummary.Wind=itemObj.Wind
	weatherSummary.Humidity=itemObj.Humidity
	weatherSummary.Chance=itemObj.Chance
	fmt.Println(weatherSummary)
}

func NextRainDay(xDays int, cityName string) {
	for i := xDays; i < 15; i++ {
		item := getItem(i, cityName)

		if item.Chance > "50%" {
			fmt.Println(item)
			break
		}
	}

}

func AverageTemp(xDays int, cityName string) {
	tempSum := 0
	for i := 0; i < xDays; i++ {
		item := getItem(i, cityName)
		tempArr := strings.Split(item.Temperature, " ")
		tempValStr := strings.Trim(strings.Split(tempArr[2], "°C")[0], "\u00a0")

		intTemp, err := strconv.Atoi(tempValStr)
		if err != nil {
			fmt.Println("Error during conversion", err)
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
	 	maxMinTemps=append(maxMinTemps, "Max\\min: "+item.Temperature) 
    }
	fmt.Println(maxMinTemps)

}


func getItem(xDays int, cityName string) item {
	c := colly.NewCollector()

	var arr []string
	var newItem = item{}
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(day int, h *colly.HTMLElement) {

			if day == xDays {
				//fmt.Println("================== day :", day, "==========================")

				arr = append(arr, h.ChildText("th span"))

				h.ForEach("td", func(j int, b *colly.HTMLElement) {
					arr = append(arr, b.Text)
				})

				// newItem =  {
				newItem.Day = arr[0]
				newItem.Temperature = arr[2]
				newItem.FeelsLike = arr[3]
				newItem.Wind = arr[5]
				newItem.Humidity = arr[7]
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
