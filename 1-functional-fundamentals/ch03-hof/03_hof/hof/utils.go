package hof

import (
	"path/filepath"
	"encoding/csv"
	"os"
	"log"
	"fmt"
	s "strings"
)

const DASHES = "-----------------------"

func PrintCars(title string, cars Collection) {
	log.Printf("\n%s\n%s\n", title, DASHES)
	for _, car := range cars {
		log.Printf("car: %v\n", car)
	}
}

func PrintCars2(title string, cars CarCollection) {
	log.Printf("\n%s\n%s\n", title, DASHES)
	for _, car := range cars {
		log.Printf("car: %v\n", car)
	}
}

func CsvToStruct(fileName string) Collection {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//pwd = "./1-functional-fundamentals/ch03-hof/hof"
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	csvfile, err := os.Open(fmt.Sprintf("%s/%s", pwd, fileName))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var cars Collection
	for _, car := range rawCSVdata {
		cars = append(cars, car[0])
	}
	return cars
}

func GetMake(sentence string) string {
	ret := sentence
	posSpace := s.Index(sentence, " ")
	if posSpace >= 0 {
		ret = sentence[:(posSpace)]
	}
	return ret
}

func GetModel(sentence string) string {
	ret := sentence
	posSpace := s.Index(sentence, " ")
	if posSpace >= 0 {
		ret = sentence[posSpace:]
	}
	return ret
}
