package CSV

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ShoppingRecord struct {
	Vegetable string
	Fruit     string
}

func createShoppingList(data [][]string) []ShoppingRecord {
	var shoppingList []ShoppingRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec ShoppingRecord
			for j, field := range line {
				if j == 0 {
					rec.Vegetable = field
				} else if j == 1 {
					rec.Fruit = field
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

func ReadCSV() {
	// open file
	f, err := os.Open("./fileOperation/CSV/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	shoppingList := createShoppingList(data)

	// print the array
	// fmt.Printf("%+v\n", shoppingList)

	for _, val := range shoppingList {
		fmt.Println(val)
		b1, err := strconv.ParseFloat(val.Vegetable, 64)
		if err != nil {
			log.Println("convert vegetable error")
		}
		fmt.Println(b1)

		b2, err := strconv.ParseFloat(val.Fruit, 64)
		if err != nil {
			log.Println("convert fruit error")
		}
		fmt.Println(b2)
	}
}
