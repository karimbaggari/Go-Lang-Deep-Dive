# **Deep Dive into Go: Section 4 - Control Flow**

## **1. If Statements and Comparisons**
- **File:** [01_if_statements.go](control_flow/01_if_statements.go)  
- **Description:** This file demonstrates how to use `if` statements to make decisions based on conditions. It includes examples of initializing variables within an `if` statement and evaluating different conditions for branching.

## **2. Logical Operators**
- **File:** [02_logical_operators.go](control_flow/02_logical_operators.go)  
- **Description:** This program illustrates the use of logical operators (`AND`, `OR`, `NOT`) in Go. These operators are used to evaluate multiple conditions in a single statement, enabling complex decision-making.

## **3. Switch Statements**
- **File:** [03_switch_statements.go](control_flow/03_switch_statements.go)  
- **Description:** This file provides examples of using `switch` statements for cleaner multi-way branching. It includes examples of:
  - Switch with conditions.
  - Switch without a condition to evaluate multiple expressions.

## **4. Select Statements**
- **File:** [04_select_statements.go](control_flow/04_select_statements.go)  
- **Description:** This example demonstrates the use of `select` statements to handle multiple channel operations. It is particularly useful in concurrent programming for managing communication between goroutines.

## **5. For Loops**
- **File:** [05_for_loops.go](control_flow/05_for_loops.go)  
- **Description:** This file explores the versatility of `for` loops in Go, including:
  - The classic `for` loop.
  - A `while`-like loop (condition-only `for` loop).
  - An infinite loop.

## **6. Nested Loops**
- **File:** [06_nested_loops.go](control_flow/06_nested_loops.go)  
- **Description:** This file shows how to work with nested loops to perform repeated iterations within other loops. It provides insights into managing loop complexity and controlling flow with `break` and `continue`.

## **7. For Range Loops**
- **File:** [07_for_range_loops.go](control_flow/07_for_range_loops.go)  
- **Description:** This example highlights the `for range` loop, which is used for iterating over collections like slices, maps, and strings. It showcases how to retrieve elements and their indices in a concise way.

## **Conclusion**
Control flow constructs are fundamental for building robust applications in Go. This section provides a detailed understanding of decision-making (`if` and `switch`), iteration (`for`, `nested`, and `range` loops), and concurrent communication (`select` statements). Mastering these concepts ensures efficient and logical code execution in your Go programs.
