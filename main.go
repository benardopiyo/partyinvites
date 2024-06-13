package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name       string
	Email      string
	Phone      string
	WillAttend bool
}

var (
	responses = make([]*Rsvp, 0, 10)
	templates = make(map[string]*template.Template, 3)
)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for i, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", i, name)
		} else {
			panic(err)
		}
	}
	fmt.Println(templateNames)
}

func welcomeHandler(res http.ResponseWriter, req *http.Request) {
	templates["welcome"].Execute(res, nil)
}

func listHandler(res http.ResponseWriter, req *http.Request) {
	templates["list"].Execute(res, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		templates["form"].Execute(res, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	} else if req.Method == http.MethodPost {
		req.ParseForm()
		responseData := Rsvp{
			Name:       req.Form["name"][0],
			Email:      req.Form["email"][0],
			Phone:      req.Form["phone"][0],
			WillAttend: req.Form["willattend"][0] == "true",
		}

		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}

		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}

		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}

		if len(errors) > 0 {
			templates["form"].Execute(res, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {

			responses = append(responses, &responseData)

			if responseData.WillAttend {
				templates["thanks"].Execute(res, responseData.Name)
			} else {
				templates["sorry"].Execute(res, responseData.Name)
			}
		}
	}
}

func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
