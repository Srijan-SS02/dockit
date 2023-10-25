# dockit
dockit is a command-line tool written in Go that simplifies the process of creating Dockerfiles for different project types. With this tool, you can quickly generate Dockerfiles tailored to your specific project, making it easier to containerize your applications.

## Features

- Generates Dockerfiles for various project types.
- Simple command-line interface for project selection.
- Option to specify the output directory for the generated Dockerfile.

## Prerequisites

- Go (Golang) installed on your system.
- Clone or download this repository to your local machine

## Getting Started

1. Clone or download this repository to your local machine:
    ```
    git clone [https://github.com/Srijan-SS02/docki.git](https://github.com/Srijan-SS02/dockit.git)
    ```
    
2. Navigate to the repository directory:
   ```
   cd dockit
   ```
3. Run the Dockerfile Generator tool:
   ```
    go run main.go
   ```
4. Choose a project type from the list of available options (e.g., "Django" or "Node.js") and press Enter.
   ```
    Choose a Project type
    Django
    Node.js
    Mojo
    Enter the project type:
   ```
   The generated Dockerfile will be displayed in the terminal.

5. To save it to a different directory:
    ```
     go run main.go /path/to/destination/directory
    ```
### Supported Project Types
- Django
- Node.js
- Go
- Flask
- Mojo 
### Extending Project Types
You can easily extend this tool to support additional project types by adding Dockerfile templates and updating the ProjectTemplates map in the code.
Or you can also raise an issue and we will fix it as soon as possible.












