package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal("OK")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	w.Write(data)
}

func (app *App) Products(w http.ResponseWriter, r *http.Request) {

	products, err := app.Database.GetProducts()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if products == nil {
		app.NotFound(w)
		return
	}

	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (app *App) Show(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	product, err := app.Database.GetProduct(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if product == nil {
		app.NotFound(w)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
