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

func (cnf *config) getLocationAreas(link string) (LocationAreas, error) {
	var locations LocationAreas
	if val, ok := cnf.Cache.Get(link); ok {
		fmt.Println("getting cached entrys")
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationAreas{}, err
		}
	} else {
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

		err = json.Unmarshal(body, &locations)
		if err != nil {
			log.Fatal(err)
		}
		cnf.Cache.Add(link, body)
	}
	if locations.Previous != nil {
		cnf.Previous = *locations.Previous
	} else {
		cnf.Previous = ""
	}

	if locations.Next != nil {
		cnf.Next = *locations.Next
	}

	return locations, nil
}
