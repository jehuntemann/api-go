package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type food struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var foods []food = []food{
	{
		Id:   1,
		Name: "Banana",
		Type: "fruta",
	},
	{
		Id:   2,
		Name: "Arroz",
		Type: "Cereal",
	},
	{
		Id:   3,
		Name: "Feijão",
		Type: "Leguminosa",
	},
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func getFoods(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(foods)
}

func setRoutes() {
	http.HandleFunc("/", mainRoute)
	http.HandleFunc("/foods", getFoods)
}

func startServer() {
	setRoutes()

	fmt.Println("Server is running is port 1410")
	http.ListenAndServe(":1410", nil)
}

func main() {
	startServer()
}
