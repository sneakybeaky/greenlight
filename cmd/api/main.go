package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	port := flags.Int("port", 4000, "API server port")
	environment := flags.String("env", "development", "Environment (development|staging|production)")

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	srv, err := newServer(
		withPort(*port),
		withEnvironment(*environment),
	)
	if err != nil {
		return err
	}

	// Declare a HTTP server with some sensible timeout settings, which listens on the
	// port provided in the config struct and uses the servemux we created above as the
	// handler.
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", srv.config.port),
		Handler:      srv,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	srv.logger.Printf("starting %s server on %s", srv.config.env, s.Addr)
	return s.ListenAndServe()
}

// Declare a string containing the application version number. Later in the book we'll
// generate this automatically at build time, but for now we'll just store the version
// number as a hard-coded global constant.
const version = "1.0.0"

type config struct {
	port int
	env  string
}

type server struct {
	config config
	router *httprouter.Router
	logger *log.Logger
}

func newServer(opts ...serveroption) (*server, error) {

	// Initialize a new logger which writes messages to the standard out stream,
	// prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	srv := &server{
		router: httprouter.New(),
		config: config{
			port: 4000,
			env:  "development",
		},
		logger: logger,
	}

	for _, opt := range opts {
		opt(srv)
	}

	srv.routes()
	return srv, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

type serveroption func(*server)

func withPort(port int) serveroption {
	return func(s *server) {
		s.config.port = port
	}
}

func withEnvironment(env string) serveroption {
	return func(s *server) {
		s.config.env = env
	}
}
