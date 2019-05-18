package main

import (	
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type DataKtp struct {
    Id    int
    Name  string    
	Address string	
	PlaceOfBirth string
	DateOfBirth string	
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "db_go"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("views/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    rows, err := db.Query("SELECT * FROM tb_ktp ORDER BY id ASC")
    if err != nil {
        log.Println(err.Error())
    }
    dt := DataKtp{}
    res := []DataKtp{}
    for rows.Next() {
        var id int
        var name, address, place_of_birth, date_of_birth string
        err = rows.Scan(&id, &name, &address, &place_of_birth, &date_of_birth)
        if err != nil {
            log.Println(err.Error())
        }
        dt.Id = id
        dt.Name = name
        dt.Address = address
		dt.PlaceOfBirth = place_of_birth
		dt.DateOfBirth = date_of_birth
		
        res = append(res, dt)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    rows, err := db.Query("SELECT * FROM tb_ktp WHERE id=?", nId)
    if err != nil {
        log.Println(err.Error())
    }
    dt := DataKtp{}
    for rows.Next() {
        var id int
        var name, address, place_of_birth, date_of_birth string
		err = rows.Scan(&id, &name, &address, &place_of_birth, &date_of_birth)
        if err != nil {
            log.Println(err.Error())
        }
        dt.Id = id
        dt.Name = name
        dt.Address = address
		dt.PlaceOfBirth = place_of_birth
		dt.DateOfBirth = date_of_birth
    }
    tmpl.ExecuteTemplate(w, "Edit", dt)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        address := r.FormValue("address")
		place_of_birth := r.FormValue("place_of_birth")
		date_of_birth := r.FormValue("date_of_birth")		
        insForm, err := db.Prepare("INSERT INTO tb_ktp(name, address, place_of_birth, date_of_birth) VALUES(?, ?, ?, ?)")
        if err != nil {
            log.Println(err.Error())
        }
        insForm.Exec(name, address, place_of_birth, date_of_birth)
        log.Println("Insert data success...")
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
		id := r.FormValue("uid")
        name := r.FormValue("name")
        address := r.FormValue("address")
		place_of_birth := r.FormValue("place_of_birth")
		date_of_birth := r.FormValue("date_of_birth")
        insForm, err := db.Prepare("UPDATE tb_ktp SET name=?, address=?, place_of_birth=?, date_of_birth=? WHERE id=?")
        if err != nil {
            log.Println(err.Error())
        }
        insForm.Exec(name, address, place_of_birth, date_of_birth, id)
        log.Println("Update data success...")
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    dt := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM tb_ktp WHERE id=?")
    if err != nil {
        log.Println(err.Error())
    }
    delForm.Exec(dt)
    log.Println("Delete data success...")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {    
    http.HandleFunc("/", Index)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
	log.Println("Server started on port 9000")
    http.ListenAndServe(":9000", nil)
}