// Tagging Git Commits with Versions
// This is a conceptual example showing how to tag a Git commit and use versioned modules.

package main

import "fmt"

func main() {
	fmt.Println("Steps for versioning your Go module:")
	fmt.Println("1. Commit your changes: `git commit -m 'Add new feature'`")
	fmt.Println("2. Tag the commit: `git tag v1.0.0`")
	fmt.Println("3. Push the tags: `git push origin --tags`")
	fmt.Println("4. Use the version in other projects: `go get mymodule@v1.0.0`")
}
