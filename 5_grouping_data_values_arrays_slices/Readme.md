# **Deep Dive into Go: Section 5 - Grouping Data Values: Arrays & Slices**

## **1. Arrays**
- **File:** [01_arrays.go](grouping_data_values_arrays_slices/01_arrays.go)  
- **Description:** This file demonstrates how to declare and initialize arrays in Go. It explains the fixed size of arrays and how to access individual elements using their index.

## **2. Slices - Composite Literals**
- **File:** [02_slices_composite_literals.go](grouping_data_values_arrays_slices/02_slices_composite_literals.go)  
- **Description:** This example introduces slices, a more flexible alternative to arrays. It shows how to create slices using composite literals and append new elements dynamically.

## **3. Accessing Values by Index and Iterating with `for range`**
- **File:** [03_slices_index_and_range.go](grouping_data_values_arrays_slices/03_slices_index_and_range.go)  
- **Description:** This file demonstrates how to access slice elements using their index and iterate over a slice with a `for range` loop, retrieving both values and their indices.

## **4. Appending to a Slice**
- **File:** [04_slices_append.go](grouping_data_values_arrays_slices/04_slices_append.go)  
- **Description:** This example shows how to add elements to a slice using the built-in `append` function, emphasizing the dynamic resizing capabilities of slices.

## **5. Slicing a Slice**
- **File:** [05_slices_slicing.go](grouping_data_values_arrays_slices/05_slices_slicing.go)  
- **Description:** This file explains how to create sub-slices from an existing slice, demonstrating how slices can share the same underlying array.

## **6. Deleting from a Slice**
- **File:** [06_slices_delete.go](grouping_data_values_arrays_slices/06_slices_delete.go)  
- **Description:** This example illustrates how to remove elements from a slice by rearranging and resizing it, ensuring efficient manipulation of data.

## **7. Using `make` with Slices**
- **File:** [07_slices_make.go](grouping_data_values_arrays_slices/07_slices_make.go)  
- **Description:** This file demonstrates how to use the `make` function to create slices with specified length and capacity, providing better control over memory allocation.

## **8. Multidimensional Slices**
- **File:** [08_slices_multidimensional.go](grouping_data_values_arrays_slices/08_slices_multidimensional.go)  
- **Description:** This example explores the creation and manipulation of multidimensional slices, showing how to store and process data in a grid-like structure.

## **9. Slice Internals & Underlying Arrays**
- **File:** [09_slices_internals.go](grouping_data_values_arrays_slices/09_slices_internals.go)  
- **Description:** This file explains how slices are structured in Go. It shows how slices point to arrays in memory, how much data they can hold (capacity), and how they share memory with those arrays.
- **Note:** the code does not explicitly show an array; it only demonstrates slices. However, in Go, slices are built on top of arrays, even if the array is not directly visible in the code.

## **Conclusion**
Arrays and slices are fundamental tools for managing collections of data in Go. While arrays provide fixed-size storage, slices offer flexibility and dynamic resizing. This section offers a comprehensive overview, from basic usage to advanced concepts like slicing, deletion, and memory internals, ensuring you have a strong grasp of these essential data structures.
