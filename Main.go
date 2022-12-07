package main

//Importar
import (
	"fmt" //Formato de impresión
	"net/http" //Para el servidor
    "html/template" //Para el template de html
    "log" //Para el log
    "database/sql" //Para la base de datos
    _ "github.com/go-sql-driver/mysql" //Para la base de datos mysql
)

//Estructura Encuestado
type Encuestado struct {
	Nombre, Email, Edad, Genero, Gusto string //Campos string
    ID int //Campo int
}

//Función Conexión
func Conexión() (*sql.DB, error) {
    //String de conexión
	usuario := "root:@tcp(127.0.0.1:3306)/encuesta"

    //Abrir conexión
	db, err := sql.Open("mysql", usuario)
	if err != nil {
        //Si hay error, imprimir y regresar
        fmt.Printf("Error obteniendo base de datos: %v", err)
		return nil, err
	}

    //Si no hay error, regresar
	return db, nil
}

//Función Insertar
func Insertar(e Encuestado) (er error) {
    //Abrir conexión
    db, err := Conexión()
	if err != nil {
        //Si hay error, imprimir y regresar
		return err
	}
    //Si no hay error, cerrar conexión
	defer db.Close()

    //Preparar sentencia
    query, err := db.Prepare("INSERT INTO `encuestado`(`Nombre`, `Email`, `Edad`, `Genero`, `Gusto`) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
    //Si no hay error, cerrar query
	defer query.Close()

	// Ejecutar sentencia, un valor por cada '?'
	_, err = query.Exec(e.Nombre, e.Email, e.Edad, e.Genero, e.Gusto)
	if err != nil {
        //Si hay error, imprimir y regresar
		return err
	}
    //Si no hay error, regresar
	return nil
}

//Función Leer
func Leer() ([]Encuestado, error){
    //Arreglo de encuestados
    Encuestados := []Encuestado{}

    //Abrir conexión
	db, err := Conexión()
	if err != nil {
        //Si hay error, imprimir y regresar
		return nil, err
	}
    //Si no hay error, cerrar conexión
	defer db.Close()

    //Preparar sentencia
	filas, err := db.Query("SELECT * FROM encuestado")

	if err != nil {
        //Si hay error, imprimir y regresar
		return nil, err
	}
	//Entonces no ocurrió un error, cerramos las filas
	defer filas.Close()

	//Se prepara una variable para los registros
	var en Encuestado

	//Recorrer todas las filas
	for filas.Next() {
		err = filas.Scan(&en.ID, &en.Nombre, &en.Email, &en.Edad, &en.Genero, &en.Gusto)
		//Escanear los valores
		if err != nil {
            //Si hay error, imprimir y regresar
			return nil, err
		}
		//Y si no, entonces se agrega al arreglo
		Encuestados = append(Encuestados, en)
	}
	//Se regresa el arreglo
	return Encuestados, nil
}


//Función Index
func Index(rw http.ResponseWriter, r *http.Request) {
    //Vista de index
    template, err := template.ParseFiles("public/index.html")
    
    if err != nil {
        //Si hay error, imprimir y regresar
        panic(err)
    } else {
        //Si no hay error, ejecutar template
        template.Execute(rw, nil)
    }
}

//Función Guardar
func Guardar(rw http.ResponseWriter, r *http.Request) {
    //Vista de guardar
    template, err := template.ParseFiles("public/index.html")

    //Mostar los datos en la consola
    nombre, existN := r.URL.Query()["nombre"]
    email, existEm := r.URL.Query()["email"]
    edad, existEd := r.URL.Query()["edad"]
    genero, existGe := r.URL.Query()["genero"]
    gusta, existGu := r.URL.Query()["gusta"]

    //Si no existen los parametros
    if !existN || !existEm || !existEd || !existGe || !existGu {
        fmt.Println("No existe algún parametro")
    } else {
        //Si existen, imprimirlos
        fmt.Println("Nombre:", nombre[0])
        fmt.Println("Email:", email[0])
        fmt.Println("Edad:", edad[0])
        fmt.Println("Genero:", genero[0])
        fmt.Println("Gusto:", gusta[0])

        //Crear objeto
        e := Encuestado{
            Nombre: nombre[0],
            Email: email[0],
            Edad: edad[0],
            Genero: genero[0],
            Gusto: gusta[0],
        }

        //Insertar objeto en la base de datos
        err2 := Insertar(e)
        if err2 != nil {
            //Si hay error, imprimir y regresar
            fmt.Println("Error al insertar", err2)
        } else {
            //Si no hay error, imprimir
            fmt.Println("Insertado correctamente")
        }

        if err != nil {
            //Si hay error, imprimir y regresar
            panic(err)
        } else {
            //Si no hay error, ejecutar template
            template.Execute(rw, nil)
        }
    }
}

func Mostrar(rw http.ResponseWriter, r *http.Request) {
    //Vista de mostrar
    template, err := template.ParseFiles("public/tabla.html")

    //Obtener los encuestados
    encuestados, err2 := Leer()
    if err2 != nil {
        //Si hay error, imprimir y regresar
        fmt.Printf("Error obteniendo los encuestados:", err)
    } else {
        if err != nil {
            //Si hay error, imprimir y regresar
            panic(err)
        } else {
            //Si no hay error, ejecutar template
            template.Execute(rw, encuestados)
        }
    }
}

// Función main
func main() {
    //Rutas
    http.HandleFunc("/inicio", Index)
    http.HandleFunc("/guardar", Guardar)
    http.HandleFunc("/mostrar", Mostrar)

    //Servidor
    fmt.Println("Servidor corriendo en http://localhost:8000/inicio")
    log.Fatal(http.ListenAndServe(":8000", nil))
}