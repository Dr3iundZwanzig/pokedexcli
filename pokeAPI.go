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

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (cnf *config) getLocationAreas(link string) (LocationAreas, error) {
	var locations LocationAreas
	if val, ok := cnf.Cache.Get(link); ok {
		fmt.Println("getting cached area entrys")
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
	} else {
		cnf.Next = "end"
	}

	return locations, nil
}

func (cnf *config) getLocationArea(wantedLocation string) (LocationArea, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/" + wantedLocation
	var location LocationArea
	if val, ok := cnf.Cache.Get(baseUrl); ok {
		fmt.Println("getting entrys for an area")
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationArea{}, err
		}
	} else {
		res, err := http.Get(baseUrl)
		if err != nil {
			return LocationArea{}, err
		}
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return LocationArea{}, err
		}

		err = json.Unmarshal(body, &location)
		if err != nil {
			return LocationArea{}, err
		}
		cnf.Cache.Add(baseUrl, body)
	}
	return location, nil
}
