# **Deep Dive into Go: Section 2 - Go Mod and Dependency Management**

## **1. Introduction to Go Modules**
- **File:** [01_go_modules_intro.go](01_go_modules_intro.go)  
- **Description:** This file demonstrates creating a new Go module and importing external packages. It covers the basic commands to initialize a module and add dependencies.

## **2. Using Go Mod Commands**
- **File:** [03_go_mod_commands.go](03_go_mod_commands.go)  
- **Description:** This file explains various `go mod` commands for managing Go dependencies effectively, including initializing a module, adding dependencies, cleaning up unused dependencies, and verifying modules.

## **3. Modular Code Example**
- **File:** [02_modular_code_example.go](02_modular_code_example.go)  
- **Description:** This file demonstrates the concept of modular code in Go, showcasing package visibility and the difference between exported and unexported identifiers. It includes an example of accessing an exported function from a helper package.

## **4. Helpers Package**
- **File:** [helpers/helpers.go](helpers/helpers.go)  
- **Description:** This file defines a helper package that includes both an exported function (`Add`) and an unexported function (`subtract`). The exported function can be accessed from other packages, while the unexported function is only available within the package itself.

## **5. Tagging Git Commits with Versions**
- **File:** [04_tagging_versions.go](04_tagging_versions.go)  
- **Description:** This file provides a conceptual example of how to tag Git commits and use versioned modules in Go. It outlines the steps to commit changes, tag the commit, and use the version in other projects.

## **6. Go Modules in Action**
- **File:** [go.mod](go.mod)  
- **Description:** This file defines the module name and Go version for the project. It also lists the required dependencies, such as `github.com/google/uuid`.

## **7. Go Module Dependency Management**
- **File:** [go.sum](go.sum)  
- **Description:** This file contains the checksums for the module dependencies, ensuring the integrity of the modules used in the project.