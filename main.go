package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Food struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var foods []Food = []Food{
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

func addFood(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
	}

	var newFood Food
	json.Unmarshal(body, &newFood)
	newFood.Id = len(foods) + 1
	foods = append(foods, newFood)
	encoder := json.NewEncoder(w)
	encoder.Encode(newFood)
}

func deleteFood(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(parts[2])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var indiceFood int = -1
	for indice, food := range foods {
		if food.Id == id {
			indiceFood = indice
			break
		}
	}

	if indiceFood < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ladoEsquerdo := foods[0:indiceFood]
	ladoDireito := foods[indiceFood+1:]
	foods = append(ladoEsquerdo, ladoDireito...)
	w.WriteHeader(http.StatusNoContent)
}

func routeFoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jason")
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) == 2 || len(parts) == 3 && parts[2] == "" {
		if r.Method == "GET" {
			getFoods(w, r)
		} else if r.Method == "POST" {
			addFood(w, r)
		}
	} else if len(parts) == 3 {
		if r.Method == "GET" {
			searchFood(w, r)
		} else if r.Method == "DELETE" {
			deleteFood(w, r)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func searchFood(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(parts[2])

	for _, food := range foods {
		if food.Id == id {
			json.NewEncoder(w).Encode(food)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func setRoutes() {
	http.HandleFunc("/", mainRoute)
	http.HandleFunc("/foods", routeFoods)
	http.HandleFunc("/foods/", routeFoods)
}

func startServer() {
	setRoutes()
	fmt.Println("Server is running is port 1410")
	http.ListenAndServe(":1410", nil)
}

func main() {
	startServer()
}
