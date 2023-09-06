package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Define the supported project types and their respective template files
var ProjectTemplates = map[string]string{
	"Django":  "templates/django/Dockerfile.tmpl",
	"Node.js": "nodejs/Dockerfile.tmpl",
	// Add more project types as needed
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose a Project type")

	// List of Projects
	for projectType := range ProjectTemplates {
		fmt.Printf("%s\n", projectType)
	}

	fmt.Print("Enter the project type: ")
	projectType, _ := reader.ReadString('\n')
	projectType = strings.TrimSpace(projectType)

	templateFile, found := ProjectTemplates[projectType]
	if !found {
		fmt.Println("Unsupported project type:", projectType)
		fmt.Println("Raise an error with the type of Dockerfile you want on the GitHub repo, and we will try to resolve it as soon as possible")
		return
	}

	// Initialize the template engine
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the target directory from the command-line argument (if provided)
	var targetDir string
	if len(os.Args) > 1 {
		targetDir = os.Args[1]
	} else {
		// If no target directory is provided, use the current working directory
		targetDir, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			return
		}
	}

	// Create or overwrite the Dockerfile in the target directory
	outputFile, err := os.Create(filepath.Join(targetDir, "Dockerfile"))
	if err != nil {
		fmt.Println("Error creating Dockerfile:", err)
		return
	}
	defer outputFile.Close()

	// Execute the template and write the Dockerfile content to the file
	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Dockerfile has been generated in %s.\n", targetDir)
}
