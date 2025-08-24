package http

import (
    "context"
    "encoding/json"
    "errors"
    "net/http"
    "strings"
    "time"
    "createuserviper/go-api/internal/domain"
    "createuserviper/go-api/internal/storage"
)

type Server struct {
    store storage.Store
    mux   *http.ServeMux
}

func NewServer(store storage.Store) *Server {
    s := &Server{store: store, mux: newMux()}
    s.routes()
    return s
}

func (s *Server) Router() *http.ServeMux { return s.mux }

func (s *Server) routes() {
    s.mux.HandleFunc("GET /health", s.handleHealth)
    s.mux.HandleFunc("GET /users", s.handleListUsers)
    s.mux.HandleFunc("POST /users", s.handleCreateUser)
    s.mux.HandleFunc("GET /users/", s.handleGetUser)
}

func decode[T any](r *http.Request, v *T) error {
    dec := json.NewDecoder(r.Body)
    return dec.Decode(v)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
    writeJSON(w, status, map[string]string{"error": msg})
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, map[string]string{"status":"ok"})
}

func (s *Server) handleListUsers(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
    defer cancel()
    out := s.store.ListUsers(ctx)
    writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var in domain.CreateUserInput
    if err := decode(r, &in); err != nil { writeError(w, http.StatusBadRequest, "invalid body"); return }
    in.Name = strings.TrimSpace(in.Name)
    in.Email = strings.TrimSpace(in.Email)
    if in.Name == "" { writeError(w, http.StatusBadRequest, "name is required"); return }
    if in.Email == "" { writeError(w, http.StatusBadRequest, "email is required"); return }
    if !domain.ValidEmail(in.Email) { writeError(w, http.StatusBadRequest, "invalid email"); return }
    ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
    defer cancel()
    u, err := s.store.CreateUser(ctx, in)
    if err != nil { writeError(w, http.StatusInternalServerError, "server error"); return }
    writeJSON(w, http.StatusCreated, u)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/users/")
    if id == "" || strings.Contains(id, "/") { writeError(w, http.StatusNotFound, "not found"); return }
    ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
    defer cancel()
    u, ok := s.store.GetUser(ctx, id)
    if !ok { writeError(w, http.StatusNotFound, "not found"); return }
    writeJSON(w, http.StatusOK, u)
}

var _ = errors.New
