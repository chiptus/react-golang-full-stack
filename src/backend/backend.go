package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const port = ":8080"

type spaHandler struct {
	staticPath string
	indexPath  string
}

func usernameHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username string `json:"username"`
	}
	user := User{os.Getenv("USERNAME")}
	p, _ := json.Marshal(user)
	w.Write(p)
}

func main() {
	log.Println("Starting Backend")

	r := mux.NewRouter()
	// Define API routes
	r.HandleFunc("/api/username", usernameHandler).Methods("GET")
	r.HandleFunc("/test", testHandler).Methods("POST")

	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	distDir := filepath.Join(currentDir, "dist")
	spa := spaHandler{staticPath: distDir, indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)
	handler := csrf.Protect([]byte("my32ByteSecretKeyForCSRF123456"), csrf.HttpOnly(false), csrf.Secure(false), csrf.Path("/"))(r)

	log.Println("Http Listening")
	http.ListenAndServe(port, handler)
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	log.Println(path)
	indexFile := filepath.Join(h.staticPath, h.indexPath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, indexFile)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
