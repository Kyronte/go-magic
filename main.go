package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

type info struct {
	CardID  string `json:"ID"`
	Name    string `json:"Name"`
	Group   string `json:"Group"`
	Snippet string `json:"Snippet"`
}

type results struct {
	Results     []info
	SearchChars string
}

type collection map[string]card

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("allcards.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic(err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var col collection

	json.Unmarshal([]byte(byteValue), &col)

	requestCardInfo("Brushland")

}

func (col collection) prettyPrintCard(card string) {
	fmt.Println("Name:", col[card].Name)
	fmt.Println("ManaCost:", col[card].ManaCost)
	fmt.Println("Type:", col[card].Type)
	fmt.Println("Text:", col[card].Text)
	if containsType(col[card].Types, "Creature") {
		fmt.Println("Power", col[card].Power)
		fmt.Println("Toughness", col[card].Toughness)
	}
}

func containsType(t []string, s string) bool {
	for _, a := range t {
		if a == s {
			return true
		}
	}
	return false
}

func requestCardInfo(card string) {
	url := fmt.Sprintf("http://gatherer.wizards.com/Handlers/InlineCardSearch.ashx?nameFragment=%+v", card)
	formattedURL := strings.Replace(url, " ", "%20", -1)

	//Fetch and save response from call
	resp, err := http.Get(formattedURL)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var res results
	json.Unmarshal([]byte(body), &res)

	fmt.Println(fetchCardImageURL(res.Results[0].CardID))

	//io.Copy(os.Stdout, r.Body)
}

func fetchCardImageURL(id string) string {
	return fmt.Sprintf("http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=%+v&type=card", id)
}
