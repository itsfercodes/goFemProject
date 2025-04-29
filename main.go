package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/itsfercodes/goFemProject/internal/app"
	"github.com/itsfercodes/goFemProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Application backend server port, default is 8080")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// defer makes the statement to be executed at the end of life of the function
	defer app.DB.Close()

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("App initated correctly at port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
