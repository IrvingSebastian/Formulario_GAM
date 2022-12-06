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

type Encuesta struct {
	Nombre, Email, Edad, Genero, Gusto string
}

//Función Conexión
func Conexión() (*sql.DB, error) {
	usuario := "root:@tcp(127.0.0.1:3306)/encuesta"

	db, err := sql.Open("mysql", usuario)
	if err != nil {
        fmt.Printf("Error obteniendo base de datos: %v", err)
		return nil, err
	}

	return db, nil
}

//Función Insertar
func Insertar(e Encuesta) (er error) {
    db, err := Conexión()
	if err != nil {
		return err
	}
	defer db.Close()

    query, err := db.Prepare("INSERT INTO `encuestado`(`Nombre`, `Email`, `Edad`, `Genero`, `Gusto`) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer query.Close()

	// Ejecutar sentencia, un valor por cada '?'
	_, err = query.Exec(e.Nombre, e.Email, e.Edad, e.Genero, e.Gusto)
	if err != nil {
		return err
	}
	return nil
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

        e := Encuesta{
            Nombre: nombre[0],
            Email: email[0],
            Edad: edad[0],
            Genero: genero[0],
            Gusto: gusta[0],
        }

        err2 := Insertar(e)
        if err2 != nil {
            fmt.Println("Error al insertar", err2)
        } else {
            fmt.Println("Insertado correctamente")
        }

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
    http.HandleFunc("/inicio", Index)
    http.HandleFunc("/guardar", Guardar)

    //Servidor
    fmt.Println("Servidor corriendo en http://localhost:8000/inicio")
    log.Fatal(http.ListenAndServe(":8000", nil))
}