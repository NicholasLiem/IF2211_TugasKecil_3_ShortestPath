package main

import (
	"fmt"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/utils"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calculate", calculateHandler)
	utils.Hello()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// localhost
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	num1, err1 := strconv.Atoi(r.FormValue("num1"))
	num2, err2 := strconv.Atoi(r.FormValue("num2"))
	operator := r.FormValue("operator")

	if err1 != nil || err2 != nil {
		fmt.Fprintln(w, "Invalid input")
		return
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Fprintln(w, "Invalid operator")
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/calculate.html"))
	data := struct {
		Num1     int
		Operator string
		Num2     int
		Result   int
	}{
		num1,
		operator,
		num2,
		result,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
