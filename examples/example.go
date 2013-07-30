package main

import (
	"fmt"
	"github.com/johnwesonga/wbdata"
)

func main() {

	client := wbdata.NewClient()

	countries, err := client.Countries.GetCountries()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	for _, v := range countries {
		fmt.Printf("%v\n", v.Name)
	}

	country, err := client.Countries.GetCountry("BR")
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
	fmt.Println(country)

}
