package main

import (
	"fmt"
	"time"
)

type Fridge struct {
	Groceries []*Grocery
}

type Grocery struct {
	Name string
	Quantity float64
	Unit string
	Category string
	Expiry time.Time
}

func (f Fridge) getGroceries() []*Grocery {
	return f.Groceries
}

func (f *Fridge) addGrocery(grocery *Grocery) {
	fmt.Printf("Adding %.2f %s of %s \n", grocery.Quantity, grocery.Unit, grocery.Name)
	f.Groceries = append(f.Groceries, grocery)
}

//func dashboardHandler(w http.ResponseWriter, r *http.Request) {
//	groceries := getGroceries()
//	t, _ := template.ParseFiles("dashboard.html")
//	t.Execute(w, groceries)
//}

func main() {
	fridge := Fridge{}
	cheese := &Grocery{Name: "White cheddar", Quantity: 1, Unit: "Block", Category: "Dairy", Expiry: time.Now()}
	fridge.addGrocery(cheese)
	g := fridge.getGroceries()
	fmt.Println(g[0])

	// Serve
	//http.HandleFunc("/", dashboardHandler)
	//log.Fatal(http.ListenAndServe(":8778", nil))
}