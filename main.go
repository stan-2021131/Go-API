package main

import (
	"encoding/json"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strings"
)

type Videojuego struct{
	ID int `json:"id"`
	NOMBRE string `json:"nombre"`
	PUBLICADO int `json:"publicado"`
	GENERO string `json:"genero"`
	PLATAFORMAS string `json:"plataformas"`
	DESARROLLADOR string `json:"desarrollador"`
}

type Message struct {
	MESSAGE string `json:"Message"`
}

func connectionSQLite() *sql.DB{
	//Creando conexión con bd
	db, err := sql.Open("sqlite3", "./Database/db24759.db")
	if err != nil {
		log.Fatal(err)
	}

	//Verificación de conexión funcional
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}


func main(){
	db := connectionSQLite()
	http.HandleFunc("/api/videojuegos/", func(w http.ResponseWriter, r *http.Request) {videojuegosHandler(db, w, r)})

	log.Println("Api sobre videojuegos favoritos en :24759")
	log.Fatal(http.ListenAndServe(":24759", nil))
}


func videojuegosHandler(db *sql.DB, w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		getVideojuegos(db, w, r)	
	case http.MethodPost:
		saveVideojuego(db, w, r)
	}
	
}

func getVideojuegos(db *sql.DB, w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	path := strings.TrimPrefix(r.URL.Path, "/api/videojuegos/")
	id := strings.Trim(path, "/")
	generoParam := query.Get("genero")
	publicadoParam := query.Get("publicado")

	newQuery := "SELECT * FROM Videojuegos WHERE 1=1"
	args := []interface{}{}

	if id != "" {
		newQuery = "SELECT * FROM Videojuegos WHERE id = ?"
		args = args[:0]
		args = append(args, id)
	} else{
		if generoParam != "" {
			newQuery += " AND genero = ?"
			args = append(args, generoParam)
		}
	
		if publicadoParam != "" {
			newQuery += " AND publicado = ?"
			args = append(args, publicadoParam)
		}
	}

	log.Println(newQuery, args)
	rows, err := db.Query(newQuery, args...)
	if(err != nil){
		http.Error(w, "Error al obtener videojuegos", http.StatusInternalServerError)
		return 
	}
	defer rows.Close()
	
	var juegos []Videojuego
	
	for rows.Next(){
		var v Videojuego
		err := rows.Scan(&v.ID, &v.NOMBRE, &v.PUBLICADO, &v.GENERO, &v.PLATAFORMAS, &v.DESARROLLADOR)
		if err != nil{
			http.Error(w, "Error en estructura de los videojuegos", http.StatusBadRequest)
			return 
		}
		juegos = append(juegos, v)
	}

	if len(juegos) ==  0{
		writeJSON(w, http.StatusNotFound, Message{"Videojuegos no encontrados"})
		return
	}
	writeJSON(w, http.StatusOK, juegos)
	return  
}

func saveVideojuego(db *sql.DB, w http.ResponseWriter, r *http.Request){
	var newVideojuego Videojuego
	
	err := json.NewDecoder(r.Body).Decode(&newVideojuego)
	if err != nil{
		http.Error(w, "Estructura invalida JSON", http.StatusBadRequest)
		return
	}

	if newVideojuego.NOMBRE == "" || newVideojuego.PUBLICADO == 0 || newVideojuego.PLATAFORMAS == "" || newVideojuego.GENERO == "" || newVideojuego.DESARROLLADOR == "" {
		http.Error(w, "Parametros faltantes", http.StatusBadRequest)
	}
	result, err := db.Exec("INSERT INTO Videojuegos(nombre, publicado, genero, plataformas, desarrollador) VALUES (?, ?, ?, ?, ?)",
		newVideojuego.NOMBRE,
		newVideojuego.PUBLICADO,
		newVideojuego.GENERO,
		newVideojuego.PLATAFORMAS,
		newVideojuego.DESARROLLADOR,
	)
	if(err != nil){
		http.Error(w, "Error al obtener videojuegos", http.StatusInternalServerError)
		return 
	}

	id, err := result.LastInsertId()
	if err == nil{
		newVideojuego.ID = int(id)
	}

	writeJSON(w, http.StatusOK, newVideojuego)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}