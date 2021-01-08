package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in := os.Open("oscar_age_male.csv")

	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < range records; x++ {
		for y := 0; y < dx; y++ {
		 rec[x][y] = uint8((x+y)/2)
		}
	   } 

	fmt.Print(records)
}