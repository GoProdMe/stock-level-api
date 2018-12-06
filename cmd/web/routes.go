package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *App) Routes() http.Handler {

	// Change the signature so we're returning a http.Handler instead of a
	// *http.ServeMux.
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/products", http.HandlerFunc(app.Products))
	mux.Get("/product/:id", http.HandlerFunc(app.Show))

	/* Middleware BGN */
	// flow of control
	//LogRequest ↔ SecureHeaders ↔ Router ↔ Application Handler
	return LogRequest(SecureHeaders(mux))

}
