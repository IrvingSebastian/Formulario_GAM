package main

//Importar
import (
	"fmt"
	"net/http"
    "html/template"
    "log"   
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

//Función Conexión
func Conexión() (*sql.DB, error) {
	usuario := "root:@tcp(127.0.0.1:3306)/encuesta"

	db, err := sql.Open("mysql", usuario)
	if err != nil {
        fmt.Printf("Error obteniendo base de datos: %v", err)
		return nil, err
	}

    err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return nil, err
	}

	return db, nil
}

//Función Index
func Index(rw http.ResponseWriter, r *http.Request) {
    template, err := template.ParseFiles("public/index.html")
    
    if err != nil {
        panic(err)
    } else {
        template.Execute(rw, nil)
    }
}

//Función Guardar
func Guardar(rw http.ResponseWriter, r *http.Request) {
    template, err := template.ParseFiles("public/index.html")

    nombre, existN := r.URL.Query()["nombre"]
    email, existEm := r.URL.Query()["email"]
    edad, existEd := r.URL.Query()["edad"]
    genero, existGe := r.URL.Query()["genero"]
    gusta, existGu := r.URL.Query()["gusta"]

    if !existN || !existEm || !existEd || !existGe || !existGu {
        fmt.Println("No existe algún parametro")
    } else {
        fmt.Println("Nombre:", nombre[0])
        fmt.Println("Email:", email[0])
        fmt.Println("Edad:", edad[0])
        fmt.Println("Genero:", genero[0])
        fmt.Println("Gusta:", gusta[0])

        if err != nil {
            panic(err)
        } else {
            template.Execute(rw, nil)
        }
    }
}

// Función main
func main() {
    //Rutas
    db, err := Conexión()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}

	fmt.Println("Conectado correctamente")

    db.Close()

    http.HandleFunc("/inicio", Index)
    http.HandleFunc("/guardar", Guardar)

    //Servidor
    fmt.Println("Servidor corriendo en http://localhost:8000/inicio")
    log.Fatal(http.ListenAndServe(":8000", nil))
}