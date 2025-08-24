package http

import (
    "net/http"
)

func newMux() *http.ServeMux {
    mux := http.NewServeMux()
    return mux
}
