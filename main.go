package main

import (
	"os"
	"html/template"
)



var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	content := map[string]string{"Variable": "MyName"}
	templates.ExecuteTemplate(os.Stdout, "lol", content)
}

