# Advent of Code 2023 - Day 8: Treetop Tree House

This repository contains my solution for Day 8 of the **Advent of Code 2023** challenge, titled **"Treetop Tree House"**. I am using this challenge to further develop my skills in **Go** while tackling interesting coding problems.

## Challenge Summary

In this problem, you're given a grid of digits representing the heights of trees in a forest. Your goal is to determine how many trees are visible from outside the grid, considering both rows and columns.

### Problem Breakdown:

1. You have a **grid** of digits where each digit represents the height of a tree.
2. A tree is visible from the edge of the grid if there are no taller trees blocking it in the same row or column.
3. You need to determine how many trees are visible from the outside of the grid.

For example, with this input:
 **30373 25512 65332 33549 35390**


You are tasked with calculating visibility by analyzing both **rows** and **columns**.

## Solution in Go

I'm using **Go** to read the input grid, process the rows and columns, and determine which trees are visible. The solution involves:

1. **Reading the input file** into a 2D slice (`[][]string`).
2. **Extracting columns** from the grid for analysis (similar to how rows are processed).
3. Implementing the logic to check if a tree is visible by comparing it to other trees in the same row or column.

## Why Go?

I'm using this challenge to improve my understanding of **Go**:
- Working with slices and 2D arrays.
- Efficiently handling file input and processing.
- Learning more about Go's built-in features for error handling, memory management, and slices.

## How to Run

1. Clone this repository.
2. Run the Go program using:

```bash
go run main.go
