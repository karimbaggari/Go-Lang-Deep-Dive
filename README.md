# **Deep Dive into Go**  

## **1. Fundamentals of Go Lang**  
- Variables: zero values, blank identifier.  
- Using `printf` for decimal and hexadecimal values.  
- Numerical Systems: decimal, binary, and hexadecimal.  
- Values, Types, Conversions, Scope, and Housekeeping.  
- Built-in Types: aggregate types.  
- Composition vs. Inheritance.  
- Composition.  

## **2. Go Mod and Dependency Management**  
- Introduction to Go Modules: dependency management.  
- Modular Code: dependency management and `go get`.  
- Go Modules in Action:  
  - `go mod init`.  
  - `go mod tidy`.  
- Package Visibility: visible, exported, and not exported.  
- Tagging Git Commits with Versions.  

## **3. Encryption**  
- Hash Algorithms.  
- Symmetric and Asymmetric Encryption.  
- Network Communication.  

## **4. Control Flow**  
- `if` Statements and Comparisons.  
- Logical Operators.  
- Statement; Statement & Comma, OK Idioms.  
- Using `switch` Statements to Make Decisions.  
- Using `select` Statements for Concurrency Communication.  
- Understanding & Using `for` Loops.  
- Nested Loops: Multiple Iterations Inside Loops.  
- Understanding & Using `for range` Loops.  
- Finding a Modulus Remainder.  

## **5. Grouping Data Values: Arrays & Slices**  
- Arrays.  
- Slices:  
  - Composite Literals.  
  - Accessing Values by Index with `for range`.  
  - Appending to a Slice.  
  - Slicing a Slice.  
  - Deleting from a Slice.  
  - Using `make` with Slices.  
  - Multidimensional Slices.  
  - Slice Internals & Underlying Arrays.  

## **6. Grouping Data Values: Maps**  
- Maps.  
- Iterating Over a Map with `for range`.  
- Deleting Elements from a Map.  
- Using the Comma-OK Idiom with Maps.  

## **7. Grouping Data Values: Structs**  
- Structs: Introduction.  
- Embedded Structs.  
- Anonymous Structs.  
- Composition.  

## **8. Functions in Go**  
- Introduction to Functions.  
- Syntax of Functions in Go.  
- Variadic Parameters.  
- Unfurling Slices.  
- `defer`.  
- Methods.  
- Interfaces and Polymorphism:  
  - Exploring the `Stringer` Interface.  
  - Expanding on the `Stringer` Interface: Wrapper Functions for Logging.  
  - `Writer` Interface: Writing to a Byte Buffer or File.  
- Anonymous Functions.  
- Function Expressions.  
- Returning Functions.  
- Callbacks.  
- Closures.  
- Function Fundamentals.  
- Recursion.  
- Wrapper Functions.  

## **9. Pointers**  
- Introduction to Pointers.  
- Seeing Types & Values for Pointers.  
- Dereferencing Pointers.  
- Pass by Value, Pointers, Reference Types, and Mutability.  
- Pointer and Value Semantics Defined.  
- Pointer & Value Semantics Heuristic.  
- Pointers and Values: The Stack and the Heap.  

## **10. Generics**  
- Type Constraints.  
- Type Constraints & Types.  
- Type Alias and Underlying Types.  
- Package Constraints.  
- Generics.  
- Concrete Type vs. Interface Type.  

## **11. Applications**  
- JSON Documentation.  
- JSON Marshalling.  
- JSON Unmarshalling.  
- Writer Interface.  
- Sorting.  
- Custom Sorting.  
- Cryptography.  

## **12. Concurrency**  
- Concurrency vs. Parallelism.  
- Using `WaitGroup`.  
- Methods Set Revisited.  
- Race Conditions.  
- Mutex.  
- Atomic Operations.  

## **13. Channels**  
- Understanding Channels.  
- Directional Channels.  
- Using Channels.  
- Ranging Over Channels.  
- Using `select` with Channels.  
- Comma-OK Idiom with Channels.  
- Fan-In Pattern.  
- Fan-Out Pattern.  
- Context in Channels.  

## **14. Error Handling**  
- Checking Errors.  
- Printing and Logging Errors.  
- Recovering from Errors.  
- Errors with Additional Info.  

## **15. Documentation with Go**  
- Using `go doc`.  
- `godoc` Tool.  
- Hosting on `godoc.org`.  

## **16. Testing and Benchmarking**  
- Table Tests.  
- Using `golint` for Code Linting.  
- Writing Benchmarks.  
- Coverage Reports.  
- Benchmarking Examples.  

