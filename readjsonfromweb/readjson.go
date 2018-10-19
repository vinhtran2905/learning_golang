package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name  string
	Price string
}

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)
	tours := TourFromJson(content)
	//	fmt.Println(tours)
	//	fmt.Println(contentFromServer(url))

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("Tour name: %v (%v)\n", tour.Name, price)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("response type %T\n ", response)

	defer response.Body.Close()

	byte, err := ioutil.ReadAll(response.Body)
	checkError(err)

	content := string(byte)
	return content
}

func TourFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)

	}
	return tours
}
