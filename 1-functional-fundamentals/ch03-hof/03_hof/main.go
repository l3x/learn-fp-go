package main

import (
	. "hof"
	"log"
	"os"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.SetOutput(os.Stdout)
}


func main() {

	if os.Getenv("RUN_HTTP_SERVER") == "TRUE" {
		router := httprouter.New()
		router.GET("/cars", CarsIndexHandler)
		router.GET("/cars/:id", CarHandler)
		log.Println("Listening on port 8000")
		log.Fatal(http.ListenAndServe(":8000", router))
	} else {

		cars := LoadCars()
		log.Printf("cars: %+v", cars)

		//PrintCars("ByMake - Honda", cars.Filter(ByMake("Honda")))
		//
		//PrintCars("Numeric", cars.Filter(ByHasNumber()))
		//
		//PrintCars("Foreign, Numeric, Toyota",
		//	cars.Filter(ByForeign()).
		//		Filter(ByHasNumber()).
		//		Filter(ByMake("Toyota")))
		//
		//PrintCars("Domestic, Numeric, GM",
		//	cars.Filter(ByDomestic()).
		//		Filter(ByHasNumber()).
		//		Filter(ByMake("GM")))
		//
		//moreCars := LoadMoreCars()
		//
		//PrintCars("More Cars, Domestic, Numeric, GM",
		//	cars.AddCars(moreCars).
		//		Filter(ByDomestic()).
		//		Filter(ByHasNumber()).
		//		Filter(ByMake("GM")))
		//
		//PrintCars("More Cars, Domestic, Numeric, Ford",
		//	cars.AddCars(moreCars).
		//		Filter(ByDomestic()).
		//		Filter(ByHasNumber()).
		//		Filter(ByMake("Ford")))
		//
		//
		//PrintCars("Numeric, Foreign, Map Upgraded",
		//	cars.Filter(ByHasNumber()).
		//		Filter(ByForeign()).
		//		Map(Upgrade()))
		//
		//PrintCars("Filter Honda, Reduce JSON",
		//	cars.Filter(ByMake("Honda")).
		//		Reduce(JsonReducer(cars), Collection{}))
		//
		//PrintCars("Reduce, Honda, JSON",
		//	cars.Reduce(MakeReducer("Honda", cars), Collection{}).
		//		Reduce(JsonReducer(cars), Collection{}))
		//
		//PrintCars2("Reduce - Lexus",
		//	cars.Filter(ByMake("Lexus")).
		//		Reduce2(CarTypeReducer(cars), []CarType{}))
		//
		//PrintCars("ByModel - Accord up/downgraded",
		//	cars.Filter(ByModel("Accord")).
		//		Map(Upgrade()).
		//		Map(Downgrade()))

		PrintCars("GenerateCars(1, 3)",
			cars.GenerateCars(1, 3))

		PrintCars("GenerateCars(1, 14), Domestic, Numeric, JSON",
			cars.GenerateCars(1, 14).
				Filter(ByDomestic()).
				Map(Upgrade()).
				Filter(ByHasNumber()).
				Reduce(JsonReducer(cars), Collection{}))

	}
}

