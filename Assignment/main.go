package main

import (
	"./datastore/BrandStore"
	"./datastore/ModelStore"
	"./datastore/VariantStore"
	"./entity"
	"./delivery/BrandDelivery"
	"./service/BrandService"
	"./service/ModelService"
	"./service/VariantService"
	"./delivery/ModelDelivery"
	"./delivery/VariantDelivery"
	"./driver"
	"log"
	"net/http"
)

func main(){
	conf := entity.MySQLConfig{
		DbName: "testDB",
		DbUser: "raghav",
		DbPass: "raghav@123M",
	}

	db,err := driver.ConnectToDB(conf)
	if err != nil {
		log.Println("Could not connect to server ",err)
		return
	}
	brandStore := BrandStore.New(db)
	modelStore := ModelStore.New(db)
	variantStore := VariantStore.New(db)

	brandServicer := BrandService.New(brandStore)
	modelServicer := ModelService.New(modelStore)
	variantServicer := VariantService.New(variantStore)

	brandHandler := BrandDelivery.New(brandServicer)
	modelHandler := ModelDelivery.New(modelServicer)
	variantHandler := VariantDelivery.New(variantServicer)

	http.HandleFunc("/brand",brandHandler.Handle)
	http.HandleFunc("/model",modelHandler.Handle)
	http.HandleFunc("/variant",variantHandler.Handle)

	log.Println("Starting server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000",nil))
}

