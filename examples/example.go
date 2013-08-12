package main

import (
	"fmt"
	"github.com/johnwesonga/wbdata"
)

func main() {

	client := wbdata.NewClient()

	countries, err := client.Countries.ListCountries()
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

	incomelevels, err := client.IncomeLevels.ListIncomeLevels()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
	fmt.Println(incomelevels)
	for _, v := range incomelevels {
		fmt.Printf("%v ", v)
	}

}
