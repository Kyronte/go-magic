package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type card struct {
	Layout        string   `json:"layout"`
	Name          string   `json:"name"`
	ManaCost      string   `json:"manaCost"`
	Cmc           int      `json:"cmc"`
	Colors        []string `json:"colors"`
	Type          string   `json:"type"`
	Types         []string `json:"types"`
	Subtypes      []string `json:"subtypes"`
	Text          string   `json:"text"`
	Power         string   `json:"power"`
	Toughness     string   `json:"toughness"`
	ImageName     string   `json:"imageName"`
	ColorIdentity []string `json:"colorIdentity"`
}

type collection struct {
	Card card
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("singlecard.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var col collection

	//fmt.Println("Byteslice contains:", &byteValue)

	error := json.Unmarshal(byteValue, &col)
	if error != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", col)

}
