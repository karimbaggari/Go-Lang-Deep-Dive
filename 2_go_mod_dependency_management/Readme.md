# **Deep Dive into Go: Section 2 - Go Mod and Dependency Management**

## **1. Introduction to Go Modules**
- **File:** [01_go_modules_intro.go](01_go_modules_intro.go)  
- **Description:** This file demonstrates creating a new Go module and importing external packages. It covers:
  - **Initializing a Go Module:** Using `go mod init <module-name>` to set up a module.
  - **Importing External Packages:** Demonstrating the import of external packages using `go get`.
  - **Basic Module Setup:** Setting up the foundational Go module for managing dependencies.

---

## **2. Modular Code Example**
- **File:** [02_modular_code_example.go](02_modular_code_example.go)  
- **Description:** This file demonstrates the concept of modular code in Go, showcasing:
  - **Package Visibility:** Explaining the difference between exported and unexported identifiers.
  - **Exported Functions:** Demonstrates how functions in packages can be accessed if they start with an uppercase letter.
  - **Modular Code Design:** Shows how code can be structured into modular components.

---

## **3. Using Go Mod Commands**
- **File:** [03_go_mod_commands.go](03_go_mod_commands.go)  
- **Description:** This file explains various `go mod` commands for managing Go dependencies effectively, including:
  - **`go mod init`:** Initializing a new module.
  - **`go mod tidy`:** Removing unused dependencies.
  - **`go mod vendor`:** This command downloads all the dependencies your project needs and stores them in a special folder called vendor. This ensures that your project can access the necessary packages without needing to fetch them from the internet every time. It's like making a copy of the dependencies inside your project to keep everything together.
  - **`go mod verify`:** This command checks that all the dependencies your project uses are correct and haven't been changed or broken. It's like making sure that all the ingredients in a recipe are exactly as they should be and haven't been swapped out by mistake.

---

## **helpers. Helpers Package**
- **File:** [helpers/helpers.go](helpers/helpers.go)  
- **Description:** This file defines a helper package that includes:
  - **Exported Function (`Add`):** A public function that can be accessed outside the package.
  - **Unexported Function (`subtract`):** A private function that is not accessible outside the package.
  - **Purpose:** Demonstrates the concept of encapsulation in Go by managing internal and external logic.

---

## **5. Tagging Git Commits with Versions**
- **File:** [04_tagging_versions.go](04_tagging_versions.go)  
- **Description:** This file provides a conceptual example of how to tag Git commits and use versioned modules in Go:
  - **Commit Changes:** Shows how to commit changes in Git.
  - **Tagging Commits:** Demonstrates how to use `git tag` to version your project.
  - **Versioning Modules:** Explains how to use tagged versions of dependencies in Go projects.


This **Section 2** demonstrates the key concepts of Go modules, including initializing a module, managing dependencies, and structuring your code in a modular way. To use the files, follow the instructions provided in the description for each file. You can run the Go files using the `go run <filename>` command and manage dependencies with the `go mod` commands.
