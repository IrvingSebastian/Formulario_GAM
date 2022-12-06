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
func Conexión() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	base := "encuesta"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, base))
	if err != nil {
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
    names, exist := r.URL.Query()["nombre"]

    if (!exist || len(names) < 0) {
        fmt.Println("No existe el parametro")
    } else {
        fmt.Println("Nombre: ", names[0])

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

    err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}

	fmt.Printf("Conectado correctamente")

    http.HandleFunc("/inicio", Index)
    http.HandleFunc("/guardar", Guardar)

    //Servidor
    fmt.Println("Servidor corriendo en http://localhost:8000/inicio")
    log.Fatal(http.ListenAndServe(":8000", nil))
}