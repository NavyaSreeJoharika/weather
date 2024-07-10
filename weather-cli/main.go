package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=ec180872243c4f57a4f153631230105&q=Iasi&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
