package main

import (
    "fmt"
    "net/http"

    "database/sql"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "morhill_user"
    password = "12345"
    dbname   = "bird_encyclopedia_test"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s " +
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

  staticFileDirectory := http.Dir("./assets/")
  staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
  r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

  r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}


func main() {
  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

  r := newRouter()

  http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hellow World!")
}
