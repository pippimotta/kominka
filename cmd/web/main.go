package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pippimotta/kominka/pkg/config"
	"github.com/pippimotta/kominka/pkg/handlers"
	"github.com/pippimotta/kominka/pkg/render"
)

var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannnot create Template Cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port: %v\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

/*
func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divided by 0")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divided by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
*/
