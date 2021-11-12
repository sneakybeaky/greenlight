package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 8080, "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	srv, err := newServer()
	if err != nil {
		return err
	}
	fmt.Printf("cccs listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	router *httprouter.Router
}

func newServer() (*server, error) {
	srv := &server{
		router: httprouter.New(),
	}
	srv.routes()
	return srv, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
