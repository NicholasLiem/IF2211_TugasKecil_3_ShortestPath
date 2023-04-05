package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/calculate.html"))
	tmpl.Execute(w, nil)
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
	tmpl.Execute(w, data)
}
