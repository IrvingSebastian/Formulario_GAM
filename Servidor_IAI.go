package main

//Importar
import (
	"fmt"
	"net/http"
    "html/template"
    "log"   
)

var i = 10

func Index(rw http.ResponseWriter, r *http.Request) {
    template, err := template.ParseFiles("public/index.html")
    
    if err != nil {
        panic(err)
    } else {
        template.Execute(rw, nil)
    }
}

// Función main
func main() {
    //Rutas
    http.HandleFunc("/inicio", Index)

    //Servidor
    fmt.Println("Servidor corriendo en http://localhost:8000/inicio")
    log.Fatal(http.ListenAndServe(":8000", nil))
}