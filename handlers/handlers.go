package handlers

import (
	"encoding/json"
	"fmt"
	"fridgeAssistant/database"
	"fridgeAssistant/models"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing here yet. Stay tuned!")
}

func AddFridgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/form.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() // Parse url parameters passed, then parse the response packet for the POST body (request body)
		input := r.Form

		// print information on server side.
		fmt.Println("path:", r.URL.Path)
		fmt.Println("scheme:", r.URL.Scheme)

		// pretty print the user input.
		bytes, err := json.MarshalIndent(input, "", "  ")
		if err != nil {
			panic(err)
		} else {
			fmt.Println("User added a fridge:", string(bytes))
		}

		database.DBCon.Create(&models.Fridge{
			Name: input["name"][0],
		})

		fmt.Fprintf(w, "Added a fridge!") // write data to response
	}




}
