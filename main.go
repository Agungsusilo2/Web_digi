package main

import (
	"DigiHero/controller"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/Register", controller.Started)
	http.HandleFunc("/Register/Heroes", controller.RegHeroes)
	http.HandleFunc("/RegisterHeroes", controller.ActRegHeroes)
	http.HandleFunc("/NextStep/Heroes", controller.NextHeroes)
	http.HandleFunc("/NextStepHeroes", controller.NextRegHeroes)
	http.HandleFunc("/Succes", controller.RegSuccess)

	log.Print("Server started on: http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}
