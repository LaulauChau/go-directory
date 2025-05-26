package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/LaulauChau/go-directory/internal/service"
	"github.com/LaulauChau/go-directory/web/templates"
)

type Handlers struct {
	directory *service.Directory
}

func NewHandlers(directory *service.Directory) *Handlers {
	return &Handlers{
		directory: directory,
	}
}

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	contacts := h.directory.ListContacts()
	component := templates.Index(contacts)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *Handlers) AddContact(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	phone := strings.TrimSpace(r.FormValue("phone"))

	if name == "" || phone == "" {
		http.Error(w, "Name and phone are required", http.StatusBadRequest)
		return
	}

	err := h.directory.AddContact(name, phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contacts := h.directory.ListContacts()
	component := templates.ContactList(contacts)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *Handlers) UpdateContact(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/contacts/")
	name, err := url.QueryUnescape(path)
	if err != nil {
		http.Error(w, "Invalid contact name", http.StatusBadRequest)
		return
	}

	phone := strings.TrimSpace(r.FormValue("phone"))
	if phone == "" {
		http.Error(w, "Phone is required", http.StatusBadRequest)
		return
	}

	err = h.directory.EditContact(name, phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contacts := h.directory.ListContacts()
	component := templates.ContactList(contacts)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *Handlers) DeleteContact(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/contacts/")
	name, err := url.QueryUnescape(path)
	if err != nil {
		http.Error(w, "Invalid contact name", http.StatusBadRequest)
		return
	}

	err = h.directory.DeleteContact(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contacts := h.directory.ListContacts()
	component := templates.ContactList(contacts)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *Handlers) SearchContact(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		w.WriteHeader(http.StatusOK)
		return
	}

	matches := h.directory.SearchContacts(query)
	component := templates.SearchResults(matches, query)
	if renderErr := component.Render(r.Context(), w); renderErr != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
