package http

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "createuserviper/go-api/internal/storage"
)

func newTestServer() *Server {
    store := storage.NewMemoryStore()
    return NewServer(store)
}

func TestHealth(t *testing.T) {
    s := newTestServer()
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    w := httptest.NewRecorder()
    s.Router().ServeHTTP(w, req)
    if w.Code != http.StatusOK { t.Fatalf("status %d", w.Code) }
}

func TestCreateUserSuccess(t *testing.T) {
    s := newTestServer()
    body, _ := json.Marshal(map[string]string{"name":"John","email":"j@e.com"})
    req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    s.Router().ServeHTTP(w, req)
    if w.Code != http.StatusCreated { t.Fatalf("status %d", w.Code) }
    var out map[string]any
    if err := json.Unmarshal(w.Body.Bytes(), &out); err != nil { t.Fatal(err) }
    if out["name"] != "John" { t.Fatalf("unexpected name %v", out["name"]) }
}

func TestCreateUserInvalidEmail(t *testing.T) {
    s := newTestServer()
    body, _ := json.Marshal(map[string]string{"name":"John","email":"invalid"})
    req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    s.Router().ServeHTTP(w, req)
    if w.Code != http.StatusBadRequest { t.Fatalf("status %d", w.Code) }
}

func TestGetUserNotFound(t *testing.T) {
    s := newTestServer()
    req := httptest.NewRequest(http.MethodGet, "/users/does-not-exist", nil)
    w := httptest.NewRecorder()
    s.Router().ServeHTTP(w, req)
    if w.Code != http.StatusNotFound { t.Fatalf("status %d", w.Code) }
}
