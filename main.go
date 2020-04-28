package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Fridge struct {
	Groceries map[string]*Grocery
}

type Grocery struct {
	Name string
	Quantity float64
	Unit string
	Category string
	Expiry time.Time
}

func (f Fridge) getGroceries() map[string]*Grocery {
	return f.Groceries
}

func (f *Fridge) addGrocery(grocery *Grocery) {
	fmt.Printf("Adding %.2f %s of %s \n", grocery.Quantity, grocery.Unit, grocery.Name)

	// Check if exists. If exists, then just add the grocery.
	if item, ok := f.Groceries[grocery.Name]; ok {
		fmt.Printf("%s already exists in the fridge.\n", grocery.Name)
		item.Quantity += grocery.Quantity
	} else {
		fmt.Printf("%s doesn't exist in the fridge yet.\n", grocery.Name)
		f.Groceries[grocery.Name] = grocery
	}

	db, err := sql.Open("sqlite3", "./fridge.db")
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO main.groceries (Name, Quantity, Unit, Category, Expiry) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range f.getGroceries() {
		_, err2 := stmt.Exec(item.Name, item.Quantity, item.Unit, item.Category, item.Expiry)
		if err2 != nil { panic(err2) }
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fridge := Fridge{}
	cheese := &Grocery{Name: "White cheddar", Quantity: 1, Unit: "Block", Category: "Dairy", Expiry: time.Now()}
	milk :=  &Grocery{Name: "Yellow cheddar", Quantity: 1, Unit: "Block", Category: "Dairy", Expiry: time.Now()}
	fridge.addGrocery(milk)
	fridge.addGrocery(cheese)
	groceries := fridge.getGroceries()
	t, _ := template.ParseFiles("dashboard.html")
	t.Execute(w, groceries)
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	check(err)
	if db == nil {
		panic("db nil")
	} else {
		creationSql, err := ioutil.ReadFile("./sql/script.sql")
		check(err)
		_, err = db.Exec(string(creationSql))
		check(err)
	}
	return db
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Initialise the db if it is the first run.
	const dbPath = "./fridge.db"
	initDB(dbPath)

	// Serve
	http.HandleFunc("/", dashboardHandler)
	log.Fatal(http.ListenAndServe(":8778", nil))
}

//fridge := Fridge{Groceries: make(map[string]*Grocery)}
//cheese := &Grocery{Name: "White cheddar", Quantity: 1, Unit: "Block", Category: "Dairy", Expiry: time.Now()}
//fridge.addGrocery(cheese)
//cheese = &Grocery{Name: "White cheddar", Quantity: 1, Unit: "Block", Category: "Dairy", Expiry: time.Now()}
//fridge.addGrocery(cheese)