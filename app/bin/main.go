package main

import (

  "encoding/json"

  "net/http"

  "fmt"

  "github.com/gorilla/mux"

  "github.com/jinzhu/gorm"

  "github.com/rs/cors"

  _ "github.com/go-sql-driver/mysql"
)

type Pop struct {

  Id     int

  Name   string

  Color  string
}

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/", GetPops).Methods("GET")

  handler := cors.Default().Handler(router)

  http.ListenAndServe(":8080", handler)
}

func GetPops(w http.ResponseWriter, r *http.Request) {

  var db *gorm.DB

  const (
    host     = "db"
    port     = 3306
    user     = "maria"
    password = "pass"
    dbname   = "beverage"
  )


    mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
       user, password, host, port, dbname)

    db, err := gorm.Open("mysql", mysqlInfo)

  if err != nil {
    panic(err)
  }

  if db != nil {
    db.Debug().DropTableIfExists(&Pop{})
    //Drops table if already exists
    var items = []Pop{

            {Name: "Cola", Color: "carmel"},

            {Name: "Gingerale",  Color: "brown"},

            {Name: "Lime",  Color: "green"},

            {Name: "Cherry",  Color: "red"},

            {Name: "Grape",  Color: "purple"},
        }
    db.Debug().AutoMigrate(&Pop{})
    //Auto create table based on Model
    for index := range items {

          db.Create(&items[index])

      }

    var pops []Pop

    db.Find(&pops)

    db.Close()

    json.NewEncoder(w).Encode(&pops)
  } else {
    fmt.Fprintf(w, "Hello astaxie!")
  }
}
