package main

import "net/http"

type Pagina struct {
	Titulo string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/teste", Pagina{Titulo: "Olá, seja muito bem vindo"}.ServerHTTP)
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func (p Pagina) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(p.Titulo))
}
