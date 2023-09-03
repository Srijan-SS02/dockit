package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Define the Path to the templates directory | not used just for readebility
const templatesDir = "templates"

// Define the supported project types and their respective template files
var ProjectTemplates = map[string]string{
	"Django":  "django/Dockerfile.tmpl",
	"Node.js": "nodejs/Dockerfile.tmpl",
	// Add more project types as needed
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose a Project type")

	//List of Projects
	for projectType := range ProjectTemplates {
		fmt.Println("%s\n", projectType)
	}

	fmt.Println("Enter the project type: ")
	projectType, _ := reader.ReadString('\n')    //will read string until next line
	projectType = strings.TrimSpace(projectType) //to remove any extra space or tab in the "projectType" string

	templateFile, found := ProjectTemplates[projectType]
	if !found {
		fmt.Println("Unsupported project type:", projectType)
		fmt.Println("Raise an error with the type of Dockerfile you want on the github repo and we will try to resolve as soon as possible")
		return
	}

	// Initialize the template engine
	tmpl, err := template.ParseFiles(templateFile) //parses template file and returns a pointer
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Execute the template and print the Dockerfile
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
