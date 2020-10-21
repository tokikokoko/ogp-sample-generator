package main

import (
	"flag"
	"html/template"
	"os"
)

var baseUrl = flag.String("u", "https://google.com/", "Base url")
var filePath = flag.String("p", "1.html", "File path")

func main() {
	// Parse flags
	flag.Parse()
	args := flag.Args()
	imageUrl := args[0]

	tpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	m := map[string]string{
		"BaseUrl":  *baseUrl,
		"FilePath": *filePath,
		"ImageUrl": imageUrl,
	}
	err = tpl.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
