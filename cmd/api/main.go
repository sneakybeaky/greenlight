package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare a string containing the Application version number. Later in the book we'll // generate this automatically at build time, but for now we'll just store the version // number as a hard-coded global constant.
const version = "1.0.0"

// Define a Config struct to hold all the configuration settings for our Application.
// For now, the only configuration settings will be the network Port that we want the
//server to listen on, and the name of the current operating environment for the
//Application (development, staging, production, etc.). We will read in these
//configuration settings from command-line flags when the Application starts.
type Config struct {
	Port int
	Env  string
}

// Define an Application struct to hold the dependencies for our HTTP handlers, helpers, // and middleware. At the moment this only contains a copy of the Config struct and a // Logger, but it will grow to include a lot more as our build progresses.

type Application struct {
	Config Config
	Logger *log.Logger
}

func main() {

	// Declare an instance of the Config struct.
	var cfg Config

	// Read the value of the Port and Env command-line flags into the Config struct. We // default to using the Port number 4000 and the environment "development" if no // corresponding flags are provided.

	flag.IntVar(&cfg.Port, "Port", 4000, "API server Port")
	flag.StringVar(&cfg.Env, "Env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new Logger which writes messages to the standard out stream, // prefixed with the current date and time.

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the Application struct, containing the Config struct and // the Logger.

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	// Declare a new servemux and add a /v1/healthcheck route which dispatches requests // to the healthcheckHandler method (which we will create in a moment).

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Declare a HTTP server with some sensible timeout settings, which listens on the // Port provided in the Config struct and uses the servemux we created above as the // handler.

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
