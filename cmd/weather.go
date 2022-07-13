package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type weather struct {
	Location struct {
		Name    string
		Country string
	}
	Current struct {
		Temp_c float32
		Is_day float32
	}
}

func (w weather) GetIsDay() string {
	if w.Current.Is_day == 0 {
		return "Day"
	} else {
		return "Night"
	}
}

type weatherApiResponse struct {
	err  error
	data *weather
}

var WeatherCmd = &cobra.Command{
	Use: "weather [location]",
	Run: func(cmd *cobra.Command, args []string) {
		location := args[0]
		ch := make(chan weatherApiResponse)

		go fetchData(ch, location)
		fmt.Printf("Waiting on Response...\n\n")
		res := <-ch

		if res.err != nil {
			log.Fatalln(res.err)
			os.Exit(1)
		}

		fmt.Printf("Weather in %s, %s:\n", res.data.Location.Name, res.data.Location.Country)
		fmt.Printf("It's %f°c and it's %s time\n", res.data.Current.Temp_c, res.data.GetIsDay())
	},
}

func fetchData(c chan weatherApiResponse, location string) {
	key := getApiKey()
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", key, location))

	if err != nil {
		c <- weatherApiResponse{
			err:  errors.New("could not fetch the api"),
			data: nil,
		}
	}

	defer res.Body.Close()

	var data weather
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		c <- weatherApiResponse{
			data: nil,
			err:  err,
		}
	}

	c <- weatherApiResponse{
		data: &data,
		err:  nil,
	}
}

func getApiKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading .env File")
	}

	return os.Getenv("WEATHER_API_KEY")
}
