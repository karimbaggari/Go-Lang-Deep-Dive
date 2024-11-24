# Zero values blank identifier

## Zero Values:
The file demonstrates the concept of zero values in Go. When variables are declared but not initialized, they automatically receive a default value based on their type:
int defaults to 0
float64 defaults to 0.0
bool defaults to false
string defaults to an empty string ""
The program prints these zero values to the console, illustrating how Go handles uninitialized variables.

## Blank Identifier:
the blank identifier (_), which is used to ignore a value returned by a function. In this case, the divide function returns both a quotient and a remainder, but only the quotient is needed for further processing.
This demonstrates a common practice in Go to avoid unused variable warnings while still utilizing the function's output.

# Understanding Numerical Systems in Go

## Purpose
This Go program demonstrates how to represent and print numbers in different numerical systems: decimal, hexadecimal, and binary. It serves as an educational tool for understanding how numerical representations work in programming.

## Learning Objectives
By studying this file, you will learn:
- How to use the `fmt.Printf` function to format output in Go.
- The differences between decimal, hexadecimal, and binary number systems.
- How to define numbers in different bases in Go (e.g., binary using `0b`, hexadecimal using `0x`).
- How to print multiple numerical representations in a single formatted output.



# Go Type Conversion and Variable Scope

## File: `fundamentals/03_values_types_conversions_scope.go`

### Description
The main file showcases:
- **Type Conversion**: Explicitly converting one data type to another, such as converting an `int` to a `float64` for arithmetic operations.
- **Variable Scope**: Understanding the accessibility of variables defined in different scopes (global vs local).

### Key Concepts

#### Type Conversion
- Go requires explicit type conversion when performing operations between different types.
- Example:
  ```go
  var a int = 10
  var b float64 = 20.5
  c := float64(a) + b // 'a' is converted to float64 before addition
  ```

#### Variable Scope
- **Global Variables**: Defined outside of functions and accessible throughout the package.
- **Local Variables**: Defined within a function and only accessible within that function.
- Example:
  ```go
  var globalVar = "I am global" // Global variable
  localVar := "I am local"       // Local variable
  ```