package static

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed webapp/*
var static embed.FS

func Serve(mux *http.ServeMux) {
	webappFS, err := fs.Sub(static, "webapp")
	if err != nil {
		panic(err)
	}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		filepath := fmt.Sprintf("webapp%s", r.RequestURI)
		if r.RequestURI != "/" && !strings.HasPrefix(r.RequestURI, "/api") && !fileExists(filepath) {
			content, _ := static.ReadFile("webapp/index.html")
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(content))
			return
		}

		http.FileServer(http.FS(webappFS)).ServeHTTP(w, r)
	})
}

func fileExists(filename string) bool {
	_, err := static.Open(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
	}
	return true
}
