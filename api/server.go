package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/LaulauChau/go-directory/internal/service"
)

type Server struct {
	handlers *Handlers
	port     string
}

func NewServer(directory *service.Directory, port string) *Server {
	return &Server{
		handlers: NewHandlers(directory),
		port:     port,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handlers.Index)
	mux.HandleFunc("/contacts", s.handleContacts)
	mux.HandleFunc("/contacts/", s.handleContactsWithPath)
	mux.HandleFunc("/search", s.handlers.SearchContact)

	server := &http.Server{
		Addr:         ":" + s.port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Starting web server on http://localhost:%s\n", s.port)
	return server.ListenAndServe()
}

func (s *Server) handleContacts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.handlers.AddContact(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleContactsWithPath(w http.ResponseWriter, r *http.Request) {

	if !strings.HasPrefix(r.URL.Path, "/contacts/") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		s.handlers.UpdateContact(w, r)
	case http.MethodDelete:
		s.handlers.DeleteContact(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) StartWithGracefulShutdown() {
	if err := s.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
