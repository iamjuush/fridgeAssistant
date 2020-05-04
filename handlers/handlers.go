package handlers

import (
	"fmt"
	"fridgeAssistant/database"
	"fridgeAssistant/models"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing here yet. Stay tuned!")
}

func renderAddItemForm(objectName string, w http.ResponseWriter) {
	templatePath := "./templates/form.html"
	item := struct {
		ObjectName string
		Route string
		Fridge bool
		Grocery bool
	}{
		ObjectName: objectName,
		Route: fmt.Sprintf("/add/%s", objectName),
	}

	if objectName == "fridge" {
		item.Fridge = true
	}
	if objectName == "grocery" {
		item.Grocery = true
	}

	t, _ := template.ParseFiles(templatePath)
	t.Execute(w, item)
}

func AddGroceryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderAddItemForm("grocery", w)
	} else {
		AddItem(r, w, "grocery")
	}
}

func AddFridgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderAddItemForm("fridge", w)
	} else {
		AddItem(r, w, "fridge")
	}
}

func AddItem(r *http.Request, w http.ResponseWriter, itemName string) {
	var model interface{}
	// print information on server side.
	fmt.Println("path:", r.URL.Path)

	if itemName == "fridge" {
		model = &models.Fridge{
			Name: r.PostFormValue("Name"),
		}
	} else if itemName == "grocery" {
		quantity, _ := strconv.ParseFloat(r.PostFormValue("Quantity"), 64)
		expiry, _ := time.Parse("2006-01-05", r.PostFormValue("Expiry"))
		model = &models.Grocery{
			Name:     r.PostFormValue("Name"),
			Quantity: quantity,
			Unit:     r.PostFormValue("Unit"),
			Category: r.PostFormValue("Category"),
			Expiry:   expiry,
			FridgeID: 0,
		}
	}

	if err := database.DBCon.Create(model).Error; err != nil {
		// error handling if adding items fail
		fmt.Fprintf(w, "Failed to add %s.\n"+
			"%v", r.PostFormValue("Name"), err) // Send error message
	} else {
		fmt.Fprintf(w, "Added a %s", itemName) // Send success message
	}
}
