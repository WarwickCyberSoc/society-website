package main

import (
	"fmt"
	"html/template"
	"os"
	"gopkg.in/yaml.v2"
)

var templates map[string]*template.Template
var config Config

func init() {
		// Remove build directory
	err := os.RemoveAll("build")
	if err != nil {
		fmt.Println("Error removing build directory:", err)
		os.Exit(1)
	}
	// Create build directory
	err = os.Mkdir("build", 0755)
	if err != nil {
		fmt.Println("Error creating build directory:", err)
		os.Exit(1)
	}


	templates = make(map[string]*template.Template)
	config = Config{}
	file, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		os.Exit(1)
	}
	
	// If NavBar is empty, add default values
	if len(config.Navbar) == 0 {
		config.Navbar = append(config.Navbar, struct {
			Name string `yaml:"name"`
			Link string `yaml:"link"`
		}{Name: "Home", Link: "/"})
	}


	// Save the config to a file
	file, err = os.Create("build/config.yaml")
	if err != nil {
		fmt.Println("Error creating config file:", err)
		os.Exit(1)
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		fmt.Println("Error encoding config file:", err)
		os.Exit(1)
	}
}

func main() {

	// Load every template
	// layout
	templates["layout"] = template.Must(template.ParseFiles("templates/layout.tmpl"))
	
	// Load every file in templates folder
	files, err := os.ReadDir("templates")
	if err != nil {
		fmt.Println("Error reading templates directory:", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() {
			continue	
		}
		if file.Name() == "layout.tmpl" {
			continue
		}
		// Generate the templa
		templates[file.Name()] = template.Must(templates["layout"].Clone())
		templates[file.Name()] = template.Must(templates[file.Name()].ParseFiles("templates/" + file.Name()))
		// Execute the template (swap tmpl with html)
		outFile, err := os.Create("build/" + file.Name()[:len(file.Name())-5] + ".html")
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		err = templates[file.Name()].ExecuteTemplate(outFile, "layout", config)
		if err != nil {
			fmt.Println("Error executing template:", err)
			os.Exit(1)
		}
		outFile.Close()
	}



	// Copy static files from public folder
	// List all files in public folder
	files, err = os.ReadDir("public")
	if err != nil {
		fmt.Println("Error reading public directory:", err)
		os.Exit(1)
	}
	// Copy each file to build directory
	for _, file := range files {
		data, err := os.ReadFile("public/" + file.Name())
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		err = os.WriteFile("build/" + file.Name(), data, 0644)
		if err != nil {
			fmt.Println("Error copying file:", err)
			os.Exit(1)
		}
	}
}
