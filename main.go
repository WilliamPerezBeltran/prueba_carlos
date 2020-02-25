package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryValues := r.URL.Query()
    fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("N"))
    fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("cadena"))
    
    numero := queryValues.Get("N")
    cadena := queryValues.Get("cadena")
    fmt.Fprintf(w, "hello, %s!\n", numero)
    fmt.Fprintf(w, "hello, %s!\n", cadena)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s %s!\n", ps.ByName("first_name"), ps.ByName("last_name"))
}

func main() {
    router := httprouter.New()
    router.GET("/practica2.1.php", Index)
    router.GET("/hello/:first_name/:last_name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}