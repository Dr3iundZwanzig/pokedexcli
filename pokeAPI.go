package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreas(link string) LocationAreas {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatal(fmt.Errorf("Response failed with status code %v", res.StatusCode))
	}
	if err != nil {
		log.Fatal(err)
	}

	var locations LocationAreas
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}
	return locations
}
