# Project Structure

This Go project is organized as follows:
- **bin/**: used for storing binary executables
- **cmd/**: store the entry points for your application
    - **myapp/**:
        - `main.go`

- **internal/**:is a directory where you can place packages that are specific to your project and not intended to be used by external projects.
    - **Scanner/**:
        - `token.go`
        - `scanner.go`
    - **Parser/**:
        - `parser.go`
        - `node.go`

- **pkg/**: is for packages that can be used by other projects.
    - **utils/**:
        - `retrieval.go`
        - `store.go`
- **tests/**: Contains test files 
- **ui/**: Contains code for the main application window and other top-level windows
    - **windows/**
        - `main_window.go`

    - **dialogs/**: Includes custom dialogs and pop-up windows
        -  `error_dialog.go`

- **go.mod**: The Go module file for dependency management.

## Getting Started

To run the application, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/Ahmed-Khaled24/Compilers-Project.git

2. Navigate to the application folder:
    ```shell
    cd Compilers-Project/cmd/myapp
    ```
3. Run the application:
    ```shell
    go run main.go
    ```