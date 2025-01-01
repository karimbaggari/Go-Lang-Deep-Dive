# **Deep Dive into Go: Section 1 - Fundamentals**

## **1. Zero Values and Blank Identifier**
- **File:** [01_zero_values_blank_identifier.go](fundamentals/01_zero_values_blank_identifier.go)  
- **Description:** This file demonstrates the concept of zero values in Go. When variables are declared but not initialized, they automatically receive a default value based on their type:
  - `int` defaults to 0
  - `float64` defaults to 0.0
  - `bool` defaults to false
  - `string` defaults to an empty string ""
  
  The program prints these zero values to the console, illustrating how Go handles uninitialized variables. It also covers the use of the blank identifier (_) to ignore values returned by functions.

## **2. Using `fmt.Printf` for Numerical Systems**
- **File:** [02_printf_and_numerical_systems.go](fundamentals/02_printf_and_numerical_systems.go)  
- **Description:** This Go program demonstrates how to represent and print numbers in different numerical systems: decimal, hexadecimal, and binary. It serves as an educational tool for understanding how numerical representations work in programming.

## **3. Go Type Conversion and Variable Scope**
- **File:** [03_values_types_conversions_scope.go](fundamentals/03_values_types_conversions_scope.go)  
- **Description:** This file showcases:
  - **Type Conversion:** Explicitly converting one data type to another, such as converting an `int` to a `float64` for arithmetic operations.
  - **Variable Scope:** Understanding the accessibility of variables defined in different scopes (global vs local).

## **4. Built-in Aggregate Types**
- **File:** [04_builtin_aggregate_types.go](fundamentals/04_builtin_aggregate_types.go)  
- **Description:** This file explains the built-in aggregate types in Go, including arrays, slices, maps, and structs. It provides examples of how to declare and use these types effectively.

## **5. Composition vs. Inheritance**
- **File:** [05_composition_vs_inheritance.go](fundamentals/05_composition_vs_inheritance.go)  
- **Description:** This file discusses the differences between composition and inheritance in Go, illustrating how composition promotes flexibility and code reuse.